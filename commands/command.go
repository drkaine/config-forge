package commands

import (
	"log"
	"os/exec"
)

func MoveFile(stdout string) {
	chemin := "/etc/nginx/sites-available/test"
	cmd := exec.Command("sudo", "mv", "/"+chemin)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
