# Distributed-System

## Description

This is a distributed system running on Owl, Owl2, Owl3, and Falcon (master).
By default, the master has port number "3000". You can change it by modifying the variable "masterAddr" on line 29 in "daemon.go".
It is working on a simple task: print current time (on receiving this task, a worker will sleep for 6 seconds before print the current time).
Create more complicated tasks by changing the code under the two "to do"s.

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

