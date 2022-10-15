import requests
from bs4 import BeautifulSoup
import ratemyprofessor
import json
from time import sleep
import re
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36 OPR/43.0.2442.991"
}
regex = re.compile('.*handlebarData theme_is_whitehot.*')
count = 0
# first iteration is without pages in url, leading will include pages
request=requests.get('https://classes.berkeley.edu/search/class?f%5B0%5D=ts_course_level%3Agrad&f%5B1%5D=im_field_term_name%3A2729&f%5B2%5D=ts_course_level%3Alow&f%5B3%5D=ts_course_level%3Aup', headers=headers)
berksoup = BeautifulSoup(request.content,'html.parser')
results = berksoup.find_all('li', class_="search-result")
totalData = []
for result in results:
    sleep(0.12)
    data = result.find("div", {"class" : regex})
    if (data):
        try:
            unparsed = json.loads(data['data-json'])
        except:
            unparsed = None
        try:
            prof = unparsed['meetings'][0]["assignedInstructors"][0]["instructor"]["names"][0]['givenName'] + " " + unparsed['meetings'][0]["assignedInstructors"][0]["instructor"]["names"][0]['familyName']
        except:
            prof = None
        try:
            course = unparsed["course"]["classDisplayName"]
        except:
            course = None
        
        dictionary = {}
        if(prof and course):
            response = ratemyprofessor.get_professor_by_id_and_name(1072,prof.replace("  "," "))
            if response is not None:
                count = count + 1
                dictionary['department'] = course.split()[0]
                dictionary['course'] = course
                dictionary['prof'] = prof
                dictionary['rating'] = response.rating
                dictionary['size'] = response.num_ratings
                dictionary['difficulty'] = response.difficulty
                totalData.append(dictionary)
# leading iterations

# for page in range(1,332):


totalData = [dict(tupleized) for tupleized in set(tuple(item.items()) for item in totalData)]
sourceFile = open('BerkeleyCourses.txt', 'w')
jsonStr = json.dumps(totalData)
print(jsonStr, file = sourceFile)
sourceFile.close()