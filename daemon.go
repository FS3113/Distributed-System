package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"sync"
	"strings"
	// "os/exec"
	// "database/sql"
)

type Daemon struct {
	ID        string
	IPAddress string
	Port      string
}

type Message struct {
	ID      string
	Type    string
	Payload string
}

var masterAddr = "128.174.246.108:3000"
var isMaster = false
var workerTimeStamp = map[string]int {}
var workerStatus = map[string]bool {}
var currentTask = map[string]string {}
var working = false
var idToServerName = map[string]string { "172.22.224.119": "Owl2", "128.174.246.108": "Falcon", "172.22.224.10": "Owl", "172.22.224.120": "Owl3" }
var mutex = &sync.Mutex{}

func (daemon Daemon) receiver() {
	servaddr, err := net.ResolveUDPAddr("udp", ":" + daemon.Port)
	serv, err := net.ListenUDP("udp", servaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer serv.Close()
	buf := make([]byte, 1024)
	fmt.Println("Receiver is ready!")

	for {
		n, _, err := serv.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		var message Message
		err = json.Unmarshal(buf[:n], &message)
		if err != nil {
			fmt.Println(err)
			return
		}
		if message.Type == "HEARTBEAT" && isMaster {
			mutex.Lock()
			workerTimeStamp[message.ID] = getTime()
			if message.Payload == "working" {
				workerStatus[message.ID] = true
			} else {
				workerStatus[message.ID] = false
			}
			mutex.Unlock()
		}
		if message.Type == "TASK" {
			working = true
			fmt.Println("Start working on: " + message.Payload)

			// to do: 
			// change the following code to working on more complicated tasks
			// instructions for tasks can be stored in "message.Payload"
			time.Sleep(6 * time.Second)
			fmt.Println(time.Now())

			fmt.Println("Finished " + message.Payload + "\n")
			working = false
		}
	}
}

func (daemon Daemon) sender(addr string, messageType string, payload string) {
	s, err := net.ResolveUDPAddr("udp4", addr)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	message := Message{ID: daemon.ID, Type: messageType, Payload: payload}
	m, err := json.Marshal(message)
	_, err = c.Write(m)
	if err != nil {
		fmt.Println(1, err)
		return
	}
}

func convertIdToServerName(id string) string {
	s := strings.Split(id, ":")
	return idToServerName[s[0]]
}

func (daemon Daemon) heartbeatManager() {
	for {
		if isMaster {
			currentTime := getTime()
			mutex.Lock()
			for k, _ := range workerStatus {
				if currentTime - workerTimeStamp[k] > 10 {
					log.Println(convertIdToServerName(k), "fails on task:", currentTask[k])
					delete(workerStatus, k)
					delete(workerTimeStamp, k)
					delete(currentTask, k)
				}
			}
			mutex.Unlock()
		} else {
			var m = "not working"
			if working {
				m = "working"
			}
			go daemon.sender(masterAddr, "HEARTBEAT", m)
		}
		time.Sleep(1 * time.Second)
	}
}

func (daemon Daemon) taskScheduler() {
	for {
		mutex.Lock()
		for k, v := range(workerStatus) {
			if !v {

				// to do: 
				// modify the instructions for tasks here (maybe read from a database containing a list tasks)
				var task = "print current time"

				go daemon.sender(k, "TASK", task)
				currentTask[k] = task
				log.Println("Assign task", task, "to", convertIdToServerName(k))
			}
		}
		mutex.Unlock()
		time.Sleep(time.Second)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	var daemon = new(Daemon)
	daemon.Port = os.Args[1]
	daemon.IPAddress = getLocalIP()
	daemon.ID = daemon.IPAddress + ":" + daemon.Port

	if daemon.ID == masterAddr {
		fmt.Println("This is Master!")
		isMaster = true
	}
	go daemon.heartbeatManager()
	go daemon.receiver()
	fmt.Printf("Port:%s, IP: %s\n", daemon.Port, daemon.IPAddress)

	if isMaster {
		go daemon.taskScheduler()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "ls" {
			for k, v := range(currentTask) {
				fmt.Println(convertIdToServerName(k), "is working on", v)
			}
		}
	}
}
