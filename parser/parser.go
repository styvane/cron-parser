// Package parser defines the parser operations

package parser

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
)

var (
	asterisk       = "*"
	rangeSeparator = "-"
	stepSeparator  = "/"
	listSeparator  = ","

	fieldAllowedValue = []*allowedValue{
		{0, 59},
		{0, 23},
		{1, 31},
		{1, 12},
		{0, 7},
	}
)

// Parser represents the cron parser data
type Parser struct {
	data string
}

type field struct {
	index int
	value string
	step  int
}

type allowedValue struct {
	start,
	end int
}

// Result represents a parsed cron data
type Result struct {
	minute,
	hour,
	dayOfMonth,
	month,
	dayOfWeek string
	cmd string
}

// New create a new parser
func New(data string) *Parser {
	return &Parser{data: data}
}

// create a new field by separating value from step.
func newField(s string, index int) *field {
	f := &field{index: index, step: 1}
	if strings.Contains(s, stepSeparator) {
		ss := strings.Split(s, stepSeparator)
		f.value = ss[0]
		f.step, _ = strconv.Atoi(ss[1]) // TODO(styvane) handle error
	} else {
		f.value = s
	}
	return f
}

// parse value parse an allowed field value
func (f *field) parseValue() (value string, err error) {
	var b strings.Builder
	switch {
	case f.value == asterisk:
		for d := fieldAllowedValue[f.index].start; d <= fieldAllowedValue[f.index].end; d += f.step {
			_, err = b.WriteString(fmt.Sprintf("%d ", d))
			if err != nil {
				return
			}
		}
	case strings.Contains(f.value, listSeparator):
		_, err = b.WriteString(strings.ReplaceAll(f.value, listSeparator, " "))
		if err != nil {
			return
		}
	case strings.Contains(f.value, rangeSeparator):
		ss := strings.Split(f.value, rangeSeparator)
		s, _ := strconv.Atoi(ss[0])
		e, _ := strconv.Atoi(ss[1])
		for s <= e {
			_, err = b.WriteString(fmt.Sprintf("%d ", s))
			if err != nil {
				return
			}
			s = s + f.step
		}
	default:
		b.WriteString(f.value)
	}

	value = strings.TrimSpace(b.String())

	return
}

// Parse parse a cron string and return a result.
// It also return any encountered error
func (p *Parser) Parse() (result *Result, err error) {
	fields := strings.Fields(p.data)
	n := len(fields) - 1
	result = &Result{cmd: fields[n]}
	for i, v := range fields[:n] {
		f := newField(v, i)
		str, err := f.parseValue()
		if err != nil {
			return nil, err
		}
		switch {
		case f.index == 0:
			result.minute = str
		case f.index == 1:
			result.hour = str
		case f.index == 2:
			result.dayOfMonth = str
		case f.index == 3:
			result.month = str
		case f.index == 4:
			result.dayOfWeek = str
		}
	}
	return
}

// Print print the formatted result into  the specified writer.
// It also returns any encountered error
func (r *Result) Print(out io.Writer) error {
	w := tabwriter.NewWriter(out, 0, 14, 2, ' ', 0)
	fmt.Fprintf(w, "minute\t%s\n", r.minute)
	fmt.Fprintf(w, "hour\t%s\n", r.hour)
	fmt.Fprintf(w, "day of month\t%s\n", r.dayOfMonth)
	fmt.Fprintf(w, "day of week\t%s\n", r.dayOfWeek)
	fmt.Fprintf(w, "command\t%s\n", r.cmd)

	w.Flush()
	return nil
}
