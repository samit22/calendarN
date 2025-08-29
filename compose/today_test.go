package main

import (
	"fmt"
	"testing"
)

func Test_getToday_OutputFormat(t *testing.T) {
	t.Log("run test for getToday() output format")
	output := getToday()
	if output == "" {
		t.Errorf("getToday() returned empty string")
	}
	if !containsAll(output, []string{"आज:", "साल", "महिनाको", "गते", "%", "दिन"}) {
		t.Errorf("getToday() output missing expected Nepali date or progress bar: %s", output)
	}
}

func containsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !contains(s, sub) {
			return false
		}
	}
	return true
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (len(substr) == 0 || (len(s) > 0 && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}

func Test_printProgressBar(t *testing.T) {
	tests := []struct {
		name      string
		progress  float64
		wantPct   int
		wantStart string
	}{
		{"zero progress", 0.0, 0, "\r[░"},
		{"half progress", 0.5, 50, "\r[▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░"},
		{"full progress", 1.0, 100, "\r[▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printProgressBar(tt.progress)
			if got[:3] != tt.wantStart[:3] {
				t.Errorf("printProgressBar() = %v, want start %v", got, tt.wantStart)
			}
			if tt.wantPct != 0 && !containsPct(got, tt.wantPct) {
				t.Errorf("printProgressBar() = %v, want percentage %d%%", got, tt.wantPct)
			}
		})
	}
}

func containsPct(s string, pct int) bool {
	return fmt.Sprintf("%d%%", pct) == s[len(s)-len(fmt.Sprintf("%d%%", pct)):]
}

func Test_repeat(t *testing.T) {
	tests := []struct {
		s     string
		count int
		want  string
	}{
		{"a", 0, ""},
		{"b", 1, "b"},
		{"c", 3, "ccc"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%d", tt.s, tt.count), func(t *testing.T) {
			got := repeat(tt.s, tt.count)
			if got != tt.want {
				t.Errorf("repeat(%q, %d) = %q, want %q", tt.s, tt.count, got, tt.want)
			}
		})
	}
}
