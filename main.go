package main

import (
	"fmt"
	"log"
	"os"
	"refreshapp/helpers"
	"time"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	apps, err := helpers.GetOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Elevate Privileges to Administrator level
	if !helpers.AmIAdmin() {
		helpers.ElevateAsAdmin()
	}

	utility := "C:\\Middleware\\user_projects\\epmsystem1\\Planning\\planning1\\CubeRefresh.cmd"
	passFile := fmt.Sprintf("%q", "C:\\epmpi\\enc_pass")

	logFile := fmt.Sprintf("%s\\epmpi_%s.log", currentDir, time.Now().Format("20060102_150405"))
	fmt.Printf("\n\nLog file is located: %s\n\n", logFile)
	logFile = fmt.Sprintf("%q", logFile)

	fmt.Println("\nStarting refresh following applications:\n")

	for i := range apps {
		fmt.Printf("%d: %s\n", i, apps[i])
		helpers.StartRefresh(utility, apps[i], passFile, logFile)
	}

}
