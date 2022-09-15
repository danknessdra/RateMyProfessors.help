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

class db:

    def __init__(self):
        # AUTHENTICATION
        load_dotenv()
        self.dbname = os.getenv('DB_NAME')
        self.user = os.getenv('DB_USERNAME')
        self.password = os.getenv('DB_PASSWORD')
        self.host = os.getenv('HOST')
    
    def connect(self):
        self.conn = psycopg2.connect("dbname={0} user={1} password={2} host={3}".format(self.dbname, self.user, self.password, self.host))
        self.cur = self.conn.cursor()
        
        self.cur.execute("CREATE TABLE IF NOT EXISTS courses ( \
                name    varchar(10), \
                department  varchar(5), \
                course  varchar(5), \
                prof    varchar(30), \
                difficulty  real, \
                size    int \
                );")
    def execute(self, command):
        self.cur.execute(command)
    
    def close(self):
        self.conn.commit()
        self.cur.close()
        self.conn.close()

'''
EXAMPLE
DB = db()
DB.connect()
DB.execute("SELECT * FROM courses;")
print(DB.cur.fetchall())
DB.close()
'''

''' 
DONT DELETE THIS, THIS IS FOR REST OF MAN DOCUMENTATION FOR LATER
'''

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

function: (function time)
> webscraping is much faster than getting data from rmp api (* means all data)

refresh function: overwrites/refreshes all data (webscrape * + rmp api *)
get function: only gets new entries, doesnt overwrite current entries (webscrape * + rmp api (newdata) )
update function: only updates current data (rmp api (current data) || webscrape (current data))

> either update ratings, or update professor data like maybe timings in the future
> could be a cron task, where every week webscrape data is updated, and rmp data updated only every quarter 
'''

# if var char size is not enough: https://stackoverflow.com/questions/22668024/how-to-change-column-size-of-varchar-type-in-mysql
# EOL format: https://stackoverflow.com/questions/59387001/mysql-command-in-python-yields-syntaxerror-eol-while-scanning-string-literal

'''
insert if not exists
https://www.postgresql.org/docs/current/sql-insert.html

 Im not sure if this is really better than official postgresql documentation (where not exists): https://stackoverflow.com/questions/5288283/sql-server-insert-if-not-exists-best-practice

do we even need this for our functions for overriding we dont need it, for updating, we should be first be comparing
 the webscrape entry to the database

'get function: only gets new entries, doesnt overwrite current entries (webscrape * + rmp api (newdata) )'
if course is not in SQL database: run api, insert into db (the check is already happening in python, no need
for sql to check with if not)

'update function: only updates current data (rmp api (current data) || webscrape (current data))'
UPDATE SQL function

on conflict will not work without a index in the table, however we are checking for a string tho
cur.execute("INSERT into courses (name, department, course, prof, difficulty, size) \
        VALUES ('CALC 1C', 'CALC', '1C', 'RICK', 3.2, 10) \
        ON CONFLICT (name) DO NOTHING;")
'''

'''
cur.execute("INSERT into courses (name, department, course, prof, difficulty, size) \
        VALUES ('CALC 1C', 'CALC', '1C', 'RICK', 3.2, 10);")
        '''

'''
cur.execute("SELECT * FROM courses;")
print(cur.fetchall())
print("DISTINCT")
cur.execute("SELECT DISTINCT ON (name) * FROM courses;")
print(cur.fetchall())
'''

'''
TEST
#cur.execute("CREATE TABLE IF NOT EXISTS test (id serial PRIMARY KEY, num integer, data varchar);")
cur.execute("INSERT INTO test (num, data) VALUES (%s, %s)", (100, "abc'def"))
#cu.execute("INSERT INTO test (num) VALUES (3);")
cur.execute("SELECT * FROM test;")
print(cur.fetchone())
'''

