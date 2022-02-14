package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartRefresh(utility string, app string, pass string, logFile string) {
	cmdAllArg := fmt.Sprintf("%s -f:%s /U:admin /A:%s /R /D /FS >> %s", utility, pass, app, logFile)
	cmd := exec.Command("powershell.exe", "/C", cmdAllArg)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
		time.Sleep(7 * time.Second)
		log.Fatalf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
	}

	fmt.Println("Refresh was successfully executed.\n")

}
