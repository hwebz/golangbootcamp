package main

import (
	"log"
	"os"
)

// type Job struct {
// 	Command string
// 	Logger *log.Logger
// }

type Job struct {
	Command string
	*log.Logger
}

func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	// same as
	// job := &Job{Command: "demo", Logger: log.New(os.Stderr, "Job: ", log.Ldate)}
	// job.Logger.Print("test")
	job.Print("starting now...")
}