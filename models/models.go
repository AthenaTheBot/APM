package models

import "os"

type ManagementFile struct {
	ProcessList []Process
	Errors      []error
}

type Process struct {
	Name string   `json:"name"`
	Logs []string `json:"logs"`
	os.Process
}

type Actions string

const (
	INIT_PROCESS      	Actions = "start"
	STOP_PROCESS      	Actions = "stop"
	KILL_PROCESS 		Actions = "kill"
	SHOW_PROCESSLIST 	Actions = "ls"
	SHOW_LOGS_PROCESS 	Actions = "logs"
)