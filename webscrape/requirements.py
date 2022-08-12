#!/bin/python3
'''
available python extraction libraries

textract
pypdf2 : using this, more documentation
pdfminer: im gonna try this one first
https://pdfminersix.readthedocs.io/en/latest/

pdfrw
'''

from pdfrw import PdfReader

x = PdfReader('../data/ge-csuge_2223.pdf')


'''
i dont think works because of perms or smth, get ignore '/Perms' verify failed
from PyPDF2 import PdfReader


reader = PdfReader("../data/ge-csuge_2223.pdf")

meta = reader.metadata

print(meta.author)

'''
'''
from pdfminer.high_level import extract_text
from pdfminer.high_level import extract_text_to_fp
from io import StringIO
from pdfminer.layout import LAParams

output_string = StringIO()

with open('../data/ge-csuge_2223.pdf', 'rb') as fin:
    extract_text_to_fp(fin, output_string, laparams=LAParams(),
    output_type='html', codec=None),
    
print(output_string.getvalue().strip())

text = extract_text("../data/DAC_to_UCB_for_2021-2022_by_Computer_Science,_Lower_Division_B.A..pdf")

print(repr(text))
'''