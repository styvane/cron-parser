package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		data  string
		want  *Result
		isErr bool
	}{
		{
			"valid string",
			`*/15 0 1,15 1-4 1-3 /usr/bin/sh`,
			&Result{"0 15 30 45", "0", "1 15", "1 2 3 4", "1 2 3", "/usr/bin/sh"},
			false,
		},

		{
			"valid string with extra range",
			`*/15 0 1,15,2-3 1-4 1-3 /usr/bin/sh`,
			&Result{"0 15 30 45", "0", "1 2 3 15", "1 2 3 4", "1 2 3", "/usr/bin/sh"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.data)
			r, err := p.Parse()
			if (tt.isErr && err == nil) || (!tt.isErr && err != nil) {
				t.Errorf("failed")
			}

			switch {
			case r.minute != tt.want.minute:
				t.Errorf("unexpected minute, want: %+v; got %+v", tt.want.minute, r.minute)
			case r.hour != tt.want.hour:
				t.Errorf("unexpected hour, want: %+v; got %+v", tt.want.hour, r.hour)
			case r.dayOfMonth != tt.want.dayOfMonth:
				t.Errorf("unexpected day of month, want: %+v; got %+v", tt.want.dayOfMonth, r.dayOfMonth)
			case r.month != tt.want.month:
				t.Errorf("unexpected month, want: %+v; got %+v", tt.want.month, r.month)

			case r.dayOfWeek != tt.want.dayOfWeek:
				t.Errorf("unexpected day of week, want: %+v; got %+v", tt.want.dayOfWeek, r.dayOfWeek)
			case r.cmd != tt.want.cmd:
				t.Errorf("unexpected command, want: %+v; got %+v", tt.want.cmd, r.cmd)

			}

		})
	}
}
