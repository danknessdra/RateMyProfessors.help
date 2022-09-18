#!/usr/bin/env python3

import ratemyprofessor


professor = ratemyprofessor.get_professor_by_school_and_name(
    ratemyprofessor.get_school_by_name("De Anza College"),
    "Sherby")
if professor is not None:
    print("professor: {0}".format(professor.name))
    print("rating: {0}".format(professor.rating))
    print("difficulty: {0}".format(professor.difficulty))
    print("num of rating: {0}".format(professor.num_ratings))
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
