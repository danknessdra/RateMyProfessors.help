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

dict = {}

for tr in soup.find_all('tr'):
    '''
    debugging
    #print("\n---")
    #print(tr)
    #print(tr.get_text())
    '''
    

    list = tr.get_text().split()
    
    # https://stackoverflow.com/questions/58146077/how-to-add-href-contains-condition-in-beautifulsoup
    # intializes course/prof since it loses scope if only in the for loop
    course = ''
    prof = ''
    for td in tr.find_all('td'):
        if td.get_text().__contains__("MATH"):
            course = td.get_text()
            #print(course)

    for a in tr.select('a[href^="/directory"]'):
        if a != None:
            prof = a.get_text()
            #print(prof)

    if  course and prof: 
        dict[course] = prof

print(dict)
    



'''
    if len(list) > 4:
        course = list[2]
        prof = list[-3] +' ' + list[-2]
    
    print(course)
    print(prof)
'''

'''
    if course != None:
        print(course)

    if prof != None:
        print(prof)
'''



departments = [ ]
