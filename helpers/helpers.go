package helpers

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"time"

	"apm/models"
)

func InitializeManagementFile() (error) {
	managementFile := models.ManagementFile{}

	// Overwriting properties to prevent them from becoming null instead of slices
	managementFile.ProcessList = make([]models.Process, 0)
	managementFile.Errors = make([]error, 0)

	bs, err := json.Marshal(managementFile)

	if err != nil {
		return err
	}

	return os.WriteFile("apm.json", bs, 0777)
}

func GetManagementFile() (models.ManagementFile, error) {
	bs, readErr := os.ReadFile("apm.json")

	if readErr != nil {
		return models.ManagementFile{}, readErr
	}

	managementFile := models.ManagementFile{}
	parseErr := json.NewDecoder(strings.NewReader(string(bs))).Decode(&managementFile)

	if parseErr != nil  {
		return models.ManagementFile{}, parseErr
	}

	return managementFile, nil
}

func WriteProcessToManagementFile(p models.Process) (error) {
	managementFile, getErr := GetManagementFile()

	if getErr != nil {
		return getErr
	}

	managementFile.ProcessList = append(managementFile.ProcessList, p)

	bs, parseErr := json.Marshal(managementFile)

	if parseErr != nil {
		return parseErr
	}

	return os.WriteFile("apm.json", bs, 0777)
}

func DeleteProcessFromManagementFile(pid int) (error) {
	managementFile, getErr := GetManagementFile()

	if getErr != nil {

		return getErr
	}

	// TODO: Refactor this code
	procList := []models.Process{}
	for _, proc := range managementFile.ProcessList {
		if proc.Pid != pid {
			procList = append(procList, proc)
		}
	}

	managementFile.ProcessList = procList

	bs, err := json.Marshal(managementFile)

	if err != nil {
		return err
	}

	return os.WriteFile("apm.json", bs, 0777)
}

func InitializeProcess(filePath string, args []string, name string) (models.Process, error) {
	fileExts := strings.Split(filePath, ".")
	parsedPath := strings.Split(filePath, "\\")

	if fileExts[len(fileExts) - 1] == "js" {
		command := exec.Command(`C:\Program Files\nodejs\node.exe`, filePath)
		command.Start()

		process := models.Process{
			Process: *command.Process,
			Name: name,
			File: parsedPath[len(parsedPath) - 1],
			StartedAt: time.Now().String(),
		}

		WriteProcessToManagementFile(process)

		return process, nil
	} else {
		
		return models.Process{}, nil
	}
}