package main

import (
	"fmt"
	"os"
	"regexp"
)
import "flag"

func main() {
	fmt.Println("Program Start")
	var folder = flag.String("f", "", "foldername")
	flag.Parse()

	fmt.Println("'" + *folder + "'")

	if checkName(*folder) != nil {
		panic("folder name is invalid")
	}

	err := os.MkdirAll("downloads/"+*folder, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished")
}

func checkName(name string) error {
	pattern, err := regexp.Compile("^[a-z0-9]+$")
	if err != nil {
		return err
	}

	if !pattern.MatchString(name) {
		return fmt.Errorf("foldername is invlid")
	}

	return nil
}
