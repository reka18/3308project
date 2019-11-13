#!/bin/python
import sys, re


def main():
    '''
    This script dynamically replaces database owner information with
    that of the current user.
    '''
    user = sys.argv[1]

    sql_script = 'sql/restore_database.sql'

    with open(sql_script, 'r') as f:
        content = f.read()
        new = re.sub(r'(Owner:.*)', 'Owner: {}'.format(user), content)
        new = re.sub(r'(OWNER\sTO.*)(?=;)', 'OWNER TO {}'.format(user), new)

    with open(sql_script, 'w') as f:
        f.write(new)


if __name__ == '__main__':
    main()
