#!/bin/python3

''' documentation
https://beautiful-soup-4.readthedocs.io/en/latest/
https://docs.python.org/3/library/json.html
https://pypi.org/project/RateMyProfessorAPI/
'''

import requests
from bs4 import BeautifulSoup
import json
import ratemyprofessor

school = ratemyprofessor.get_school_by_name("De Anza College")

json_file = open('../data/courses.json', 'w')
departments = open('../data/departments.txt', 'r').read().splitlines()


courses = []
count = 0 # for counting num of teachers gone through


for department in departments:
    url = 'https://www.deanza.edu/schedule/listings.html?dept='+department+'&t=F2022'
    #print(url)

    # http request
    request=requests.get(url)

    # parse html
    soup = BeautifulSoup(request.content,'html.parser')

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
            if td.get_text().__contains__(department):
                course = td.get_text()
                #print(course)

        for a in tr.select('a[href^="/directory"]'):
            if a != None:
                prof = a.get_text()
                #print(prof)
        '''
        if  course and prof: 
            dict[course] = prof
        '''

        dictionary = {}

        # trying to remove duplicate entries: and (dict.get(course) == prof in courses
        if course and prof:
            professor = ratemyprofessor.get_professor_by_school_and_name(school,prof)
            if professor is not None:
                count = count + 1
                print("{0} professor: {1}".format(count, professor.name))
                #print("rating: {0}".format(professor.rating))
                #print("difficulty: {0}".format(professor.difficulty))
                #print("num of rating: {0}".format(professor.num_ratings))
                dictionary['department'] = department
                dictionary['course'] = course
                dictionary['prof'] = prof
                dictionary['rating'] = professor.rating
                dictionary['size'] = professor.num_ratings
                dictionary['difficulty'] = professor.difficulty
                courses.append(dictionary)
            else:
                print("N/A for {0} {1}".format(course, prof))

        #print(courses)

            #print(json.dumps({'department': department, 'course': course, 'prof': prof}, sort_keys=True, indent=4))
    #print(dict)
    
# remove duplicate courses

# we should convert this so it does updates instead of overwriting each time
# script takes too long to run
# if having to run again, risks losing data
# also requesting data from API is quite slow, maybe we can collaborate with rate my professors
# converting based off a database model: still questions, how do we check if entry is there
# without a huge run time

courses = [dict(tupleized) for tupleized in set(tuple(item.items()) for item in courses)]
json.dump({"courses": courses}, json_file, sort_keys=True, indent=4)

json_file.close()

print("FINISHED")
print("FINISHED")
print("FINISHED")


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

'''
if professor is not None:
    print("professor: {0}".format(professor.name))
    print("rating: {0}".format(professor.rating))
    print("difficulty: {0}".format(professor.difficulty))
    print("num of rating: {0}".format(professor.num_ratings))
'''
'''
    print("%sworks in the %s Department of %s." % (professor.name, professor.department, professor.school.name))
    print("Rating: %s / 5.0" % professor.rating)
    print("Difficulty: %s / 5.0" % professor.difficulty)
    print("Total Ratings: %s" % professor.num_ratings)
    if professor.would_take_again is not None:
        print(("Would Take Again: %s" % round(professor.would_take_again, 1)) + '%')
    else:
        print("Would Take Again: N/A")
'''
