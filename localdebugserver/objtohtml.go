package localdebugserver

import (
	"encoding/json"
	"html/template"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	mapTmplString = `
<ul>
	{{- range $key, $value := .}}
		<li><b>{{$key}}:</b> {{.}}</li>
	{{- end}}
</ul>
`
	listTmplString = `
	<ol>
		{{- range .}}
			<li>{{.}}</li>
		{{- end}}
	</ol>
`
)

var (
	mapTmpl  = template.Must(template.New("map").Parse(mapTmplString))
	listTmpl = template.Must(template.New("list").Parse(listTmplString))
)

// ObjectToHTML converts an object to an HTML representation
func ObjectToHTML(data interface{}) (string, error) {
	var result string

	// Check the type of the data
	switch v := reflect.ValueOf(data); v.Kind() {
	case reflect.String:
		result = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		result = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		result = strconv.FormatBool(v.Bool())
	case reflect.Array, reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Struct || v.Type().Elem().Kind() == reflect.Interface {
			// Handle arrays/slices of objects (structs)
			r, err := arrayToTableHTML(v)
			if err != nil {
				return "", err
			}
			result = r
		} else {
			// Handle arrays/slices of primitives
			var items []string
			for i := 0; i < v.Len(); i++ {
				itemHTML, err := ObjectToHTML(v.Index(i).Interface())
				if err != nil {
					return "", err
				}
				items = append(items, itemHTML)
			}
			var buf strings.Builder
			if err := listTmpl.Execute(&buf, items); err != nil {
				return "", err
			}
			result = buf.String()
		}
	case reflect.Map:
		// Handle maps
		items := make(map[string]interface{})
		keys := v.MapKeys()
		for _, key := range keys {
			keyStr := key.String()
			valueHTML, err := ObjectToHTML(v.MapIndex(key).Interface())
			if err != nil {
				return "", err
			}
			items[keyStr] = resultToHTML(valueHTML)
		}
		var buf strings.Builder
		if err := mapTmpl.Execute(&buf, items); err != nil {
			return "", err
		}
		result = buf.String()
	default:
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		result = string(jsonBytes)
	}

	return result, nil
}

// arrayToTableHTML converts an array/slice of objects (structs or interfaces) to an HTML table
func arrayToTableHTML(arr reflect.Value) (string, error) {
	if arr.Len() == 0 {
		return "", nil
	}

	// Create the table header
	var fieldNames []string
	if arr.Index(0).Kind() == reflect.Struct {
		// If the first element is a struct, extract field names
		firstElem := arr.Index(0)
		for i := 0; i < firstElem.NumField(); i++ {
			fieldNames = append(fieldNames, firstElem.Type().Field(i).Name)
		}
	} else {
		// If the first element is an interface, use the keys as field names
		interfaceElem := arr.Index(0).Interface().(map[string]interface{})
		for key := range interfaceElem {
			fieldNames = append(fieldNames, key)
		}
	}
	sort.Strings(fieldNames)
	tableHeader := "<tr>"
	for _, fieldName := range fieldNames {
		tableHeader += "<th>" + fieldName + "</th>"
	}
	tableHeader += "</tr>"

	// Create rows with object field values
	rows := ""
	for i := 0; i < arr.Len(); i++ {
		row := "<tr>"
		elem := arr.Index(i)
		if elem.Kind() == reflect.Struct {
			// Handle struct elements
			for j := 0; j < elem.NumField(); j++ {
				fieldValue, err := ObjectToHTML(elem.Field(j).Interface())
				if err != nil {
					return "", err
				}
				row += "<td>" + fieldValue + "</td>"
			}
		} else {
			// Handle interface elements
			interfaceElem := elem.Interface().(map[string]interface{})
			for _, fieldName := range fieldNames {
				fieldValue, err := ObjectToHTML(interfaceElem[fieldName])
				if err != nil {
					return "", err
				}
				row += "<td>" + fieldValue + "</td>"
			}
		}
		row += "</tr>"
		rows += row
	}

	return "<table border=\"1\">" + tableHeader + rows + "</table>", nil
}

// resultToHTML converts a string to an HTML template.HTML type
func resultToHTML(s string) template.HTML {
	return template.HTML(s)
}
