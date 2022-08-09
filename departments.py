#!/bin/python3

import requests
from bs4 import BeautifulSoup

departments=requests.get('https://www.deanza.edu/schedule/')

soup = BeautifulSoup(departments.content,'html.parser')



select = soup.find(id='dept-select')

for option in select.find_all('option'):
    print(option)

