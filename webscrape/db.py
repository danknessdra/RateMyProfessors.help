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

load_dotenv()

dbname = os.getenv('DB_NAME')
user = os.getenv('DB_USERNAME')
password = os.getenv('DB_PASSWORD')

conn = psycopg2.connect("dbname={0} user={1} password={2}".format(dbname, user, password))

cur = conn.cursor()

#cur.execute("CREATE TABLE test (num integer);")
cur.execute("DROP TABLE test")
cur.execute("CREATE TABLE test (id serial PRIMARY KEY, num integer, data varchar);")
cur.execute("INSERT INTO test (num, data) VALUES (%s, %s)", (100, "abc'def"))
#cu.execute("INSERT INTO test (num) VALUES (3);")
cur.execute("SELECT * FROM test;")
print(cur.fetchone())

conn.commit()

cur.close()
conn.close()
