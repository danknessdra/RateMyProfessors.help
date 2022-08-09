#!/bin/python3

import requests
from bs4 import BeautifulSoup

r= requests.get('https://www.deanza.edu/')
math=requests.get('https://www.deanza.edu/schedule/listings.html?dept=MATH&t=F2022')

soup = BeautifulSoup(math.content,'html.parser')

#print(math)
#print(math.content)

#print(soup.prettify())


for a in soup.find_all('a'):
    course = a.get('data-content')

    if course != None:
        print(course)
