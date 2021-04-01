/*
Copyright © 2020 Styvane Soukossi <styvane@acm.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
