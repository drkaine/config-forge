package tests

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestRunNanoWithSudo(t *testing.T) {
	path := "/path/to/file"
	cmd := exec.Command("sudo", "nano", path)

	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	expectedArgs := []string{"sudo", "nano", path}
	if !reflect.DeepEqual(cmd.Args, expectedArgs) {
		t.Errorf("Expected args %v but got %v", expectedArgs, cmd.Args)
	}

	if cmd.Stdin != nil {
		t.Errorf("Expected nil for Stdin but got %v", cmd.Stdin)
	}
	if cmd.Stdout != nil {
		t.Errorf("Expected nil for Stdout but got %v", cmd.Stdout)
	}
	if cmd.Stderr != nil {
		t.Errorf("Expected nil for Stderr but got %v", cmd.Stderr)
	}
}
