package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	env := readEnv()
	createLogger()
	trackerInfo := createTrackerInfo(env)
	tracker := NewTracker(trackerInfo)

	fmt.Println(tracker.String())
}

func readEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprint("Unable to read .env file: ", err))
	}

	env := make(map[string]string)

	env["PORT"] = os.Getenv("PORT")

	return env
}

func createTrackerInfo(env map[string]string) *TrackerInfo {
	strPort := env["PORT"]
	ip := "::1"
	ipversion := IPv6

	port, err := strconv.Atoi(strPort)

	if err != nil {
		logrus.Fatal("Cannot convert port to int", err)
	}

	trackerInfo := &TrackerInfo{
		ip:        ip,
		port:      port,
		ipversion: ipversion,
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
