#!/bin/python3

import os
from dotenv import load_dotenv
import psycopg2

# make this use .env variables
'''
dbname=input("Enter db name: ")
username=input("Enter user name: ")
password=input("Enter password: ")
conn = psycopg2.connect("dbname={0} user={1} password={2}".format(dbname, username, password))
'''

load_dotenv()

DB_NAME = os.getenv('DB_NAME')
USERNAME = os.getenv('DB_USERNAME')
PASSWORD = os.getenv('DB_PASSWORD')

print(DB_NAME)
print(DB_USERNAME)
conn = psycopg2.connect("dbname={0} user={1} password={2}".format(DB_NAME, DB_USERNAME, DB_PASSWORD))


conn.close()
