package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"apm/helpers"
	"apm/models"
)

func main() {
	managementFile, managementFileErr := helpers.GetManagementFile()

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguements was specified!")
		os.Exit(1)
	}

	if managementFileErr != nil {
		fmt.Println("Management file not found, creating new one.")

		initErr := helpers.InitializeManagementFile()

		if initErr != nil {
			fmt.Println("An erro occured while creating management file.")
			os.Exit(1)
		}

		fmt.Println("Initialized management file.")
	}

	processes := []models.Process{}

	for _, m := range managementFile.ProcessList {
		proc, err := os.FindProcess(m.Pid)

		if err == nil {
			process := models.Process{
				Process: *proc,
				Name: m.Name,
				Logs: m.Logs,
			}
			processes = append(processes, process)
		}
	}

	switch(os.Args[1]) {
		case string(models.INIT_PROCESS):
			fmt.Println("Initializing process...")

			procName := "myApp"
			if len(os.Args) >= 4 {
				procName = os.Args[3]
			}

			dir, _ := os.Getwd()
			_, err := helpers.InitializeProcess(filepath.Join(dir, os.Args[2]), []string{}, procName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Successfully started process.");
			break

		case string(models.SHOW_PROCESSLIST):
			if len(processes) == 0 {
				fmt.Println("No process to list!")
			} else {
 				fmt.Println("Listing current processes...")
				for _, proc := range processes {
					fmt.Printf("[PID]: %d, [Name]: %s, [Uptime]: %d\n", proc.Pid, proc.Name, 0)
				}
			}
			break

		case string(models.KILL_PROCESS):
			fmt.Println("Killing process...")

			targetProc := os.Args[2]	
			for _, p := range processes {
				if p.Name == targetProc || strconv.FormatInt(int64(p.Pid), 10) == targetProc {
					helpers.DeleteProcessFromManagementFile(p.Pid)
					p.Kill()
				}
			}
			break

		default:
			fmt.Println("Invalid action specified")
	}
}