package main

import (
	"fmt"
	"my_backup_project/internal/backup"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("My_backup input: my_backup <full | incremental> <workDir> <backupDir>")
		return
	}
	typeOfBackUp := os.Args[1]
	workDir := os.Args[2]
	backupDir := os.Args[3]
	timeCreate := time.Now().Format("2006-01-02_15-04-05")
	if typeOfBackUp == "full" {
		err := backup.CreateFullBackUp(workDir, backupDir, timeCreate)
		if err != nil {
			fmt.Println("error creatintg full backup", err)
		}
	} else if typeOfBackUp == "incremental" {
		err := backup.CreateIncrementalBackUp(workDir, backupDir, timeCreate)
		if err != nil {
			fmt.Println("error creatintg full backup", err)
		}
	} else {
		fmt.Println("unknown backup type", typeOfBackUp)
		return
	}
}
