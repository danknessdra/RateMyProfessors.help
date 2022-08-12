#https://deanza.edu/articulation/documents/ge-csuge_2223.pdf

import requests
from bs4 import BeautifulSoup

url='https://deanza.edu/articulation/documents/ge-csuge_2223.pdf'
request=requests.get(url)

'''
# parse html
soup = BeautifulSoup(request.content,'html.parser')

print(soup.get_text)
'''