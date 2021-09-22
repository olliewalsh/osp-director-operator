#!/bin/bash
set -eux

# add cloud-admin ssh keys to $HOME/.ssh
mkdir -p $HOME/.ssh
cp /mnt/ssh-config/* $HOME/.ssh/
chmod 600 $HOME/.ssh/git_id_rsa
chown -R cloud-admin: $HOME/.ssh

GIT_HOST=$(echo $GIT_URL | sed -e 's|^git@\(.*\):.*|\1|g')
GIT_USER=$(echo $GIT_URL | sed -e 's|^git@.*:\(.*\)/.*|\1|g')

cat <<EOF > $HOME/.ssh/config
Host $GIT_HOST
    User $GIT_USER
    IdentityFile $HOME/.ssh/git_id_rsa
    StrictHostKeyChecking no
EOF
chmod 644 $HOME/.ssh/config

unset OS_CLOUD
export OS_AUTH_TYPE=none
export OS_ENDPOINT="http://{{ .HeatServiceName }}:8004/v1/admin"

HEAT_COUNT=0
until openstack stack list &> /dev/null || [ "$HEAT_COUNT" -gt 180 ]; do
  HEAT_COUNT=$(($HEAT_COUNT + 1))
  echo "waiting for Heat API to startup..."
  sleep 2
done

# delete the stack if it exists (should only happen on retries)
openstack stack delete overcloud -y --wait &>/dev/null || true

# create a temporary scratch directory to assemble the Heat templates
TEMPLATES_DIR=$HOME/tripleo-deploy-scratch/tripleo-heat-installer-templates
rm -Rf "$TEMPLATES_DIR"
mkdir -p $TEMPLATES_DIR

cp -a /usr/share/openstack-tripleo-heat-templates/* $TEMPLATES_DIR
pushd $TEMPLATES_DIR
# extract any tar files into the $TEMPLATES_DIR
{{- if .TripleoTarballFiles }}
{{- range $key, $value := .TripleoTarballFiles }}
tar -xvf $HOME/tht-tars/{{ $key }}
{{- end }}
{{- end }}

# copy to editable dir config-tmp
rm -Rf $HOME/config-tmp
mkdir -p $HOME/config-tmp
cp $HOME/config/* $HOME/config-tmp
cp $HOME/config-custom/* $HOME/config-tmp
cp $HOME/config-fencing/* $HOME/config-tmp
#FIXME: get rid of /usr/share/openstack-tripleo-heat-templates/ and use relative paths in dev-tools!
sed -e "s|/usr/share/openstack\-tripleo\-heat\-templates|\.|" -i $HOME/config-tmp/*.yaml
# copy to our temp t-h-t dir
cp -a $HOME/config-tmp/* "$TEMPLATES_DIR/"

python3 tools/process-templates.py -r $TEMPLATES_DIR/roles_data.yaml -n $TEMPLATES_DIR/network_data.yaml


# NOTE: only applies to OSP 16, on OSP 17+ we set NetworkSafeDefaults: false in the Heat ENV
if [ -e ./network/scripts/run-os-net-config.sh ]; then
  # disable running dhcp on all interfaces, setting disable_configure_safe_defaults in the interface template does not work
  sudo sed -i '/^set -eux/a disable_configure_safe_defaults=true' ./network/scripts/run-os-net-config.sh
fi

# only use env files that have ContainerImagePrepare in them, if more than 1 the last wins
PREPARE_ENV_ARGS=""
for ENV_FILE in $(grep -rl "ContainerImagePrepare:" *.yaml | grep -v overcloud-resource-registry-puppet); do
  PREPARE_ENV_ARGS="-e $ENV_FILE"
done

# if no container image prepare env files are provided generate the defaults
if [ -z "$PREPARE_ENV_ARGS" ]; then
  openstack tripleo container image prepare default --output-env-file container-image-prepare.yaml
  PREPARE_ENV_ARGS="-e container-image-prepare.yaml"
fi

if [ "$OSPVersion" == "16.2" ]; then
  PREPARE_ENV_ARGS="$PREPARE_ENV_ARGS -e hostnamemap.yaml"
else
  PREPARE_ENV_ARGS="$PREPARE_ENV_ARGS -e rendered-tripleo-config.yaml"
fi
openstack tripleo container image prepare $PREPARE_ENV_ARGS -r roles_data.yaml --output-env-file=tripleo-overcloud-images.yaml

mkdir -p $HOME/tripleo-deploy
rm -rf $HOME/tripleo-deploy/overcloud-ansible*

time openstack stack create --wait \
    -e $TEMPLATES_DIR/overcloud-resource-registry-puppet.yaml \
    -e $TEMPLATES_DIR/tripleo-overcloud-images.yaml \
    -e $TEMPLATES_DIR/environments/deployed-server-environment.yaml \
    -e $TEMPLATES_DIR/environments/docker-ha.yaml \
    -e ~/config-passwords/tripleo-overcloud-passwords.yaml \
    -e $TEMPLATES_DIR/fencing.yaml \
{{- range $key, $value := .TripleoDeployFiles }}
    -e {{ $key }} \
{{- end }}
{{- range $key, $value := .TripleoCustomDeployFiles }}
    -e {{ $key }} \
{{- end }}
    -t overcloud.yaml overcloud

mkdir -p $HOME/ansible

if [ "$OSPVersion" == "16.2" ]; then
# FIXME: there is no local 'config-download' command in OSP 16.2 (use tripleoclient config-download in OSP 17)
/usr/bin/python3 - <<"EOF_PYTHON"
from tripleoclient import utils as oooutils
from osc_lib import utils
from tripleo_common.inventory import TripleoInventory
from tripleo_common.actions import ansible
import sys
import os

API_NAME = 'tripleoclient'
API_VERSIONS = {
    '1': 'heatclient.v1.client.Client',
}
api_port='8004'
heat_client = utils.get_client_class(
    API_NAME,
    '1',
    API_VERSIONS)
client = heat_client(
    endpoint='http://{{ .HeatServiceName }}:%s/v1/admin' % api_port,
    username='admin',
    password='fake',
    region_name='regionOne',
    token='fake',
)
out_dir = oooutils.download_ansible_playbooks(client, 'overcloud', output_dir='/home/cloud-admin/ansible')

inventory = TripleoInventory(
    hclient=client,
    plan_name='overcloud',
    ansible_ssh_user='cloud-admin')

extra_vars = {
    'Standalone': {
        'ansible_connection': 'local',
        'ansible_python_interpreter': sys.executable,
        }
    }
inv_path = os.path.join(out_dir, 'tripleo-ansible-inventory.yaml')
inventory.write_static_inventory(inv_path, extra_vars)

# NOTE: we don't use transport=local like tripleoclient standalone
ansible.write_default_ansible_cfg(
    out_dir,
    'cloud-admin',
    ssh_private_key=None)

EOF_PYTHON

# with OSP17 the rendered templates get into overcloud directory, create link to have the same dst
output_dir=$(ls -dtr $HOME/ansible/tripleo-ansible-* | tail -1)
ln -sf ${output_dir} $HOME/ansible/overcloud

else

TMP_DIR_ANSIBLE=$(mktemp -d)
pushd $TMP_DIR_ANSIBLE

cat <<EOF > $TMP_DIR_ANSIBLE/vars.yaml
plan: overcloud
ansible_ssh_user: cloud-admin
ansible_ssh_private_key_file: $HOME/.ssh/id_rsa
output_dir: $HOME/ansible
EOF

echo -e "localhost ansible_connection=local\n\n[convergence_base]\nlocalhost" > hosts

ANSIBLE_FORCE_COLOR=true ansible-playbook -i hosts -e @vars.yaml /usr/share/ansible/tripleo-playbooks/cli-config-download.yaml
popd

# remove the .git directory as it conflicts with the repo below
rm -Rf $HOME/ansible/overcloud/.git

fi

TMP_DIR=$(mktemp -d)
git clone $GIT_URL $TMP_DIR
pushd $TMP_DIR
git checkout -b $ConfigHash
# add directory for playbooks
mkdir tripleo-ansible
cp -a $HOME/ansible/overcloud/* tripleo-ansible

# add directory for templates
mkdir source-templates
cp -a $TEMPLATES_DIR/* source-templates

# custom config
mkdir config-custom
cp -L $HOME/config-custom/* config-custom

# add tarball files
mkdir tarball
{{- if .TripleoTarballFiles }}
{{- range $key, $value := .TripleoTarballFiles }}
tar -xvf $HOME/tht-tars/{{ $key }} -C tarball
{{- end }}
{{- end }}

git config --global user.email "dev@null.io"
git config --global user.name "OSP Director Operator"

git add *
git commit -a -m "Generated playbooks for $ConfigHash"
git push -f origin $ConfigHash
popd
