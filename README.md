# Distributed-System

## Description

This is a distributed system running on Owl, Owl2, Owl3, and Falcon (master).\
By default, the master has port number "3000". You can change it by modifying the variable "masterAddr" on line 31 in "daemon.go".\
Run your own tasks by changing the code under the three "to do"s in "daemon.go".\
\
To run the sample task:\
Change the database configuration in initialize_table.py and script.py.\
Initialize database by "python initialize_table.py". You can run this on any server.\
initialize_table.py creates two tables, including an input list of tasks and an output table.\
script.py is a very simple algorithm accepting a single input task and storing the output into the database.\
See "Usage" for how to run the system.\
\
\
To use this system:\
1. Create a table containing a list of tasks (see initialize_table.py for an example)\
\
2. Create a script for running tasks.\
\
3. Modify the code under the four "todo"s in daemon.go

4. Currently with the code under "todo 4", the master will print an error message when it detects a worker is disconnected from the system when working on task "currentTask[k]".


## Usage

Build by

```bash
go build daemon.go helperFunctions.go
```

Run by (use different port number on different servers)
```bash
./daemon PORT_NUMBER
```
Exampe for the master (Falcon):
```bash
./daemon 3000
```

When master is running, input "ls" to see the working status of workers.

