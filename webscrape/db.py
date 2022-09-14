#!/bin/python3

import os
from dotenv import load_dotenv
import psycopg2


'''
create/edit .env file (name has to be exactly ".env")
sample config:

DB_NAME="testdb"
DB_USERNAME="test"
DB_PASSWORD="password"
'''

# AUTHENTICATION
load_dotenv()

dbname = os.getenv('DB_NAME')
user = os.getenv('DB_USERNAME')
password = os.getenv('DB_PASSWORD')
host = os.getenv('HOST')

conn = psycopg2.connect("dbname={0} user={1} password={2} host={3}".format(dbname, user, password, host))

cur = conn.cursor()

'''
DESIGN
TABLE: courses
course name (PHYS 4A)
department (PHYS)
course number (4A)
difficulty (3)
prof (Rick Taylor)
size (# of ratings)

there should be a different functions based off for writing, follow how jellyfin updates metadata

refresh function: overwrites/refreshes all data
update function: only gets new entries
update function overload: only checks/updates certain parameters

getting data takes long time with RateMyProfessors api so update function should be much faster than refresh
'''

# if var char size is not enough: https://stackoverflow.com/questions/22668024/how-to-change-column-size-of-varchar-type-in-mysql
# EOL format: https://stackoverflow.com/questions/59387001/mysql-command-in-python-yields-syntaxerror-eol-while-scanning-string-literal
cur.execute("CREATE TABLE IF NOT EXISTS courses ( \
        name    varchar(10), \
        department  varchar(5), \
        course  varchar(5), \
        prof    varchar(30), \
        difficulty  real, \
        size    int \
        );")

# where not exists: https://stackoverflow.com/questions/5288283/sql-server-insert-if-not-exists-best-practice
cur.execute("INSERT courses \
        SELECT DISTINCT name \
        FROM (name, department, course, prof, difficulty, size) \
        VALUES ('CALC 1C', 'CALC', '1C', 'RICK', 3.2, 10);")

cur.execute("SELECT * FROM courses;")
cur.execute("SELECT * FROM courses;")
print(cur.fetchall())

# TEST
#cur.execute("CREATE TABLE IF NOT EXISTS test (id serial PRIMARY KEY, num integer, data varchar);")
cur.execute("INSERT INTO test (num, data) VALUES (%s, %s)", (100, "abc'def"))
#cu.execute("INSERT INTO test (num) VALUES (3);")
cur.execute("SELECT * FROM test;")
print(cur.fetchone())

conn.commit()

cur.close()
conn.close()
