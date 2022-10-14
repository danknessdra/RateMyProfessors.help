import requests
from bs4 import BeautifulSoup
import ratemyprofessor
import json
from time import sleep
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36 OPR/43.0.2442.991"
}
request=requests.get('https://classes.berkeley.edu/search/class/?f%5B0%5D=ts_course_level%3Agrad&f%5B1%5D=im_field_term_name%3A2729&f%5B2%5D=ts_course_level%3Alow&f%5B3%5D=ts_course_level%3Aup', headers=headers)
berksoup = BeautifulSoup(request.content,'html.parser')
results = berksoup.find_all('li', class_="search-result")
totalData = []
count = 0
for result in results:
    sleep(0.12)
    data = result.find(class_="handlebarData theme_is_whitehot")
    if (data):
        try:
            unparsed = json.loads(data['data-json'])
        except:
            unparsed = None
        try:
            prof = unparsed['meetings'][0]["assignedInstructors"][0]["instructor"]["names"][0]['formattedName']
        except:
            prof = None
        try:
            course = unparsed["course"]["classDisplayName"]
        except:
            course = None
        dictionary = {}
        if(prof and course):
            professor = ratemyprofessor.get_professor_by_id_and_name(1072,prof.replace("  "," "))
            if professor is not None:
                count = count + 1
                dictionary['department'] = course.split()[0]
                dictionary['course'] = course
                dictionary['prof'] = prof
                dictionary['rating'] = professor.rating
                dictionary['size'] = professor.num_ratings
                dictionary['difficulty'] = professor.difficulty
                totalData.append(dictionary)
                print (dictionary)