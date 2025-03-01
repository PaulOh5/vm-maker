package utils

import (
	"fmt"
	"os/user"
	"strings"
	"testing"
)

func TestRunShellCommand(t *testing.T) {
	tests := []struct {
		bin      string
		args     []string
		expected string
		wantErr  bool
	}{
		{
			bin:      "echo",
			args:     []string{"hello"},
			expected: "hello\n",
			wantErr:  false,
		},
		{
			bin:      "ls",
			args:     []string{"nonexistent-folder-path"},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got, err := RunShellCommand(tt.bin, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RunShellCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Fatalf("RunShellCommand() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRunShellProccess(t *testing.T) {
	tests := []struct {
		bin     string
		args    []string
		wantErr bool
	}{
		{
			bin:     "sleep",
			args:    []string{"10"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			cmd, err := RunShellProcess(tt.bin, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RunShellProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, err := RunShellCommand("bash", "-c", "ps -ef | grep "+fmt.Sprint(cmd.Process.Pid)); err != nil {
				t.Fatalf("RunShellProcess() = %v, want %v", err, nil)
			}
		})
	}
}

func TestCheckGroup(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("user.Current() error = %v", err)
	}

	output, err := RunShellCommand("groups", currentUser.Username)
	if err != nil {
		t.Fatalf("RunShellCommand() error = %v", err)
	}

	groupBulk := strings.Split(strings.TrimSpace(output), ":")
	groups := strings.Split(strings.TrimSpace(groupBulk[1]), " ")

	for _, group := range groups {
		if check, err := CheckGroup(group); err != nil {
			t.Errorf("CheckGroup(%s) error = %v", group, err)
		} else if !check {
			t.Errorf("CheckGroup(%s) = %v, want %v", group, check, true)
		}
	}

	if check, err := CheckGroup("nonexistent-group"); err != nil {
		t.Errorf("CheckGroup(%s) error = %v", "nonexistent-group", err)
	} else if check {
		t.Errorf("CheckGroup(%s) = %v, want %v", "nonexistent-group", check, false)
	}
}
