package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	env := readEnv()
	createLogger()
	trackerInfo := createTrackerInfo(env)
	tracker := NewTracker(trackerInfo)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go startServer(tracker)

	wg.Wait()
}

func readEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprint("Unable to read .env file: ", err))
	}

	env := make(map[string]string)

	env["PORT"] = os.Getenv("PORT")
	env["IP"] = os.Getenv("IP")

	return env
}

// Uses the environment variables to create a TrackerInfo object
func createTrackerInfo(env map[string]string) *TrackerInfo {
	strPort := env["PORT"]
	ip := env["IP"]
	var ipversion IpVersion

	ipv := os.Getenv("IP_VERSION")

	if ipv == "IPv4" {
		ipversion = IPv4
	} else if ipv == "IPv6" {
		ipversion = IPv6
	} else {
		logrus.Fatal("Invalid IP version", ipv)
	}

	port, err := strconv.Atoi(strPort)

	if err != nil {
		logrus.Fatal("Cannot convert port to int", err)
	}

	trackerInfo := &TrackerInfo{
		Ip:        ip,
		Port:      port,
		Ipversion: ipversion,
	}

	return trackerInfo
}

func createLogger() {
	//Create logger with logrus
	logFile, err := os.Create("logs/log.txt")

	if err != nil {
		log.Fatal("Cannot create log file", err)
	}

	logrus.SetOutput(logFile)
}
