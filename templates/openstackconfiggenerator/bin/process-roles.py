#! /usr/bin/env python3

import argparse
import collections
import yaml

class TemplateLoader(yaml.SafeLoader):
    def construct_mapping(self, node):
        self.flatten_mapping(node)
        return collections.OrderedDict(self.construct_pairs(node))

class TemplateDumper(yaml.SafeDumper):
    def represent_ordered_dict(self, data):
        return self.represent_dict(data.items())

    def description_presenter(self, data):
        if not len(data) > 80:
            return self.represent_scalar('tag:yaml.org,2002:str', data)
        return self.represent_scalar('tag:yaml.org,2002:str', data, style='>')


TemplateDumper.add_representer(str,
                               TemplateDumper.description_presenter)
TemplateDumper.add_representer(collections.OrderedDict,
                               TemplateDumper.represent_ordered_dict)
TemplateLoader.add_constructor(yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG,
                               TemplateLoader.construct_mapping)


def get_role_counts(environment):
    role_counts = {}
    for env_file in environment:
        with open(env_file, 'r', encoding='utf-8') as envfile:
            data = yaml.safe_load(envfile.read())
        if data is not None:
            for param, val in data.get('parameter_defaults', {}).items():
                if param.endswith('Count'):
                    role_name = param[:-5]
                    try:
                        role_counts[role_name] = int(val)
                    except ValueError:
                        pass
    return role_counts

def filter_roles(roles_file, environment):
    role_counts = get_role_counts(environment)

    with open(roles_file, 'r', encoding='utf-8') as orig:
        original_data = orig.read()
        original_roles = yaml.load(original_data, Loader=TemplateLoader)

    # Omit roles that are not in use or count is 0
    new_roles = [x for x in original_roles if role_counts.get(x['name'], 0) > 0]

    with open(roles_file, 'w', encoding='utf-8') as new:
        new.write(yaml.dump(new_roles, Dumper=TemplateDumper, default_flow_style=False))


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description='Filter unused roles in role_data'
    )
    parser.add_argument(
        '-r', '--roles-data',
        metavar='<roles_data>',
        required=True,
        help='Path to the roles data'
    )
    parser.add_argument(
        '-e', '--environment',
        metavar='<environment>',
        action='append',
        required=True,
        help='Path to the environment. Can be specified multiple times'
    )
    args = parser.parse_args()
    filter_roles(args.roles_data, args.environment)
