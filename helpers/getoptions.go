package helpers

import (
	"errors"
	"fmt"
	"log"
)

func GetOptions(arguments []string) ([]string, error) {

	allApps := []string{"Console", "Feed", "Meat", "MEZ", "Plant", "Pork", "Poultry", "TradeCo", "Turkey"}
	var apps []string

	if len(arguments) < 1 {
		return nil, errors.New("error: you must pass at least one argument")

	} else if arguments[0] == "all" && len(arguments) == 1 {
		fmt.Println("\n\nStart execute: refreshapp all")
		return allApps, nil

	} else if arguments[0] != "all" && len(arguments) >= 1 {

		// check each argument if it's equal to any application name
		for i := 0; i <= len(arguments)-1; i++ {

			tempAppList := allApps
			for j := 0; j <= len(allApps)-1; j++ {
				if arguments[i] == tempAppList[j] {
					// we found a name, so delete application j from the temp list
					tempAppList = append(tempAppList[:j], tempAppList[j+1:]...)
					break
				} else if j == len(allApps)-1 {
					log.Fatalf("error: wrong application name: %s", arguments[i])
				}
			}

			apps = append(apps, arguments[i])
		}

		return apps, nil
	}

	return nil, errors.New("error: unrecognized arguments")
}
