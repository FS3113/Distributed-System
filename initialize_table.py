import mysql.connector

mydb = mysql.connector.connect(
	host="owl2.cs.illinois.edu",
	user="juefeic2",
	password="0202141208",
	database="juefeic2_educationtoday"
)
mycursor = mydb.cursor()

# column priority can be used to prioritize tasks,
# then the master server can assign tasks follow this priority
mycursor.execute("DROP TABLE IF EXISTS Tasks")
mycursor.execute("CREATE TABLE Tasks (task VARCHAR(20), priority int)")

sql = 'insert into Tasks (task, priority) values (%s, %s)'
for i in range(1000):
	mycursor.execute(sql, ('{} plus {}'.format(i, i), i))

mycursor.execute('drop table if exists Output')
mycursor.execute("CREATE TABLE Output (task varchar(20), result varchar(20), timestamp varchar(100), status varchar(100))")
