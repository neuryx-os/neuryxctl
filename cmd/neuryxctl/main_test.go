package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{"version", []string{"version"}, exitOK, "dev\n", ""},
		{"help", []string{"help"}, exitOK, "Uso:", ""},
		{"sem argumentos", nil, exitOK, "Uso:", ""},
		{"não implementado", []string{"gpu", "detect"}, exitNotImplemented, "", "comando ainda não implementado\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("código = %d, esperado %d", code, tt.wantCode)
			}
			if !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, esperado conter %q", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != tt.wantStderr {
				t.Errorf("stderr = %q, esperado %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}
