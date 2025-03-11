// Package term provides information about the terminal that the current process is connected to (if any),
// for example measuring the dimensions of the terminal and inspecting whether it's safe to output color.
package terminal

import (
	"testing"
)

func TestFromEnv(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		wantTerminal  bool
		wantColor     bool
		want256Color  bool
		wantTrueColor bool
	}{
		{
			name: "default",
			env: map[string]string{
				"CLICOLOR":       "",
				"CLICOLOR_FORCE": "",
				"NO_COLOR":       "",
				"TERM":           "",
				"COLORTERM":      "",
			},
			wantTerminal:  false,
			wantColor:     false,
			want256Color:  false,
			wantTrueColor: false,
		},
		{
			name: "force color",
			env: map[string]string{
				"CLICOLOR":       "",
				"CLICOLOR_FORCE": "1",
				"NO_COLOR":       "",
				"TERM":           "",
				"COLORTERM":      "",
			},
			wantTerminal:  false,
			wantColor:     true,
			want256Color:  false,
			wantTrueColor: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			terminal := FromEnv()
			if got := terminal.IsTerminalOutput(); got != tt.wantTerminal {
				t.Errorf("expected terminal %v, got %v", tt.wantTerminal, got)
			}
			if got := terminal.IsColorEnabled(); got != tt.wantColor {
				t.Errorf("expected color %v, got %v", tt.wantColor, got)
			}
			if got := terminal.Is256ColorSupported(); got != tt.want256Color {
				t.Errorf("expected 256-color %v, got %v", tt.want256Color, got)
			}
			if got := terminal.IsTrueColorSupported(); got != tt.wantTrueColor {
				t.Errorf("expected truecolor %v, got %v", tt.wantTrueColor, got)
			}
		})
	}
}
