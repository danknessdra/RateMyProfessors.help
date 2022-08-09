#!/bin/python3

''' documentation
https://beautiful-soup-4.readthedocs.io/en/latest/
https://docs.python.org/3/library/csv.html
https://pypi.org/project/RateMyProfessorAPI/
'''

import requests
from bs4 import BeautifulSoup
import csv
import ratemyprofessor

# csv file
file = "courses.csv"

fields = [ 'Department', 'Course', 'Professor', 'Rating' ]
with open(file, 'w') as csvfile:
    csvwriter = csv.writer(csvfile)

    csvwriter.writerow(fields)



# http request
r= requests.get('https://www.deanza.edu/')
math=requests.get('https://www.deanza.edu/schedule/listings.html?dept=MATH&t=F2022')


# parse html
soup = BeautifulSoup(math.content,'html.parser')

'''
#print(math)
#print(math.content)
#print(soup.prettify())
'''

for a in soup.find_all('a'):
    course = a.get('data-content')

    if course != None:
        print(course)
