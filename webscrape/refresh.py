#!/usr/bin/env python3
from db import *

print("test")

DB = db()
DB.connect()
DB.execute("SELECT * FROM courses;")
print(DB.cur.fetchall())
DB.close()
