import mysql.connector
import sys
import json
import os
from datetime import datetime
import time


def algorithm(task):
    t = task.split(' plus ')
    return int(t[0]) + int(t[1])

if __name__ == "__main__":
    mydb = mysql.connector.connect(
        host="owl2.cs.illinois.edu",
        user="juefeic2",
        password="0202141208",
        database="juefeic2_educationtoday"
    )
    mycursor = mydb.cursor()

    task = ' '.join(sys.argv[1:])
    time.sleep(6)

    try:
        res = algorithm(task)
        status = 'success'
    except:
        res = '???'
        status = 'fail'

    timestamp = datetime.now()
    
    sql = 'insert into Output (task, result, timestamp, status) values (%s, %s, %s, %s)'
    mycursor.execute(sql, (task, res, timestamp, status))
    mydb.commit()
