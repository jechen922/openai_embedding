package cmd

import (
	"log"
	"os/exec"
)

func Exec(cmdArgs ...string) error {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(cmdArgs)
		log.Println(err)
		log.Println(string(output))
		return err
	} else {
		log.Println(string(output))
	}
	return nil
}
