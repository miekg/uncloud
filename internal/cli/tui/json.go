package tui

import (
	"strings"
	"unicode"

	"charm.land/lipgloss/v2/table"
)

type Objects []map[string]string

// TableToObjects convert the table to Objects that can then be used for json marshalling. Basic usage:
//
//	 type Machines struct { // define a wrapper type
//		 Machines []map[string]string `json:"machines"`
//	 }
//
//	 machines := Machines{Machines: tui.TableToObjects(t)} // convert the table
//	 data, _ := json.MarshalIndent(machines, "", "  ") // marshal to json
//	 fmt.Println(string(data)) // output
func TableToObjects(t *table.Table) Objects {
	rows := table.DataToMatrix(t.GetData())
	headers := t.GetHeaders()
	objects := make(Objects, len(rows))

	for i, row := range rows {
		obj := make(map[string]string)
		for j, val := range row {
			obj[camelcase(headers[j])] = val
		}
		objects[i] = obj
	}

	return objects
}

// camelcase returns a strings the is camel cased for use as JSON keys.
func camelcase(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}

	sb := strings.Builder{}
	transform := unicode.ToLower
	for i, r := range s {
		if i == 0 {
			sb.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsSpace(r) {
			transform = unicode.ToUpper
			continue
		}

		sb.WriteRune(transform(r))
		transform = unicode.ToLower
	}
	return sb.String()
}
