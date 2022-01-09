package html

import (
	"bytes"
	"encoding/base64"
	"strings"
	"text/template"

	"github.com/yosssi/gohtml"
)

type TableRowData []string

type DataEntity interface {
	Render(buf *bytes.Buffer) error
}

type TableData struct {
	Head TableRowData
	Rows []TableRowData
}

type dataEntityTable struct {
	data TableData
}

func (t *dataEntityTable) Render(buf *bytes.Buffer) error {
	tr := func(s ...string) { outputTag(buf, trTag, s...) }
	td := func(s ...string) { outputTag(buf, tdTag, s...) }
	th := func(s string, attrs ...attr) {
		tagStart(buf, thTag, append(attrs, attr{key: "class", val: "th-sm"})...)
		out(buf, s)
		tagEnd(buf, thTag)
	}
	table := func() {
		out(buf, `<table class="sortable-table table table-striped table-bordered table-sm" cellspacing="0" width="100%">`)
	}

	table()
	tagStart(buf, theadTag)
	tr()
	for _, h := range t.data.Head {
		th(h)
	}
	tagEnd(buf, trTag)
	tagEnd(buf, theadTag)
	tagStart(buf, tbodyTag)
	for _, row := range t.data.Rows {
		tr()
		for _, d := range row {
			td(d)
		}
		tagEnd(buf, trTag)
	}
	tagEnd(buf, tbodyTag)
	tagEnd(buf, tableTag)

	return nil
}

func MakeDataEntityFromTable(data TableData) DataEntity {
	return &dataEntityTable{data}
}

type dataEntityTableSimple struct {
	data TableData
}

func (t *dataEntityTableSimple) Render(buf *bytes.Buffer) error {
	tr := func(s ...string) { outputTag(buf, trTag, s...) }
	td := func(s ...string) { outputTag(buf, tdTag, s...) }
	th := func(s string, attrs ...attr) {
		tagStart(buf, thTag)
		out(buf, s)
		tagEnd(buf, thTag)
	}
	table := func() {
		out(buf, `<table border=1>`)
	}

	table()
	tagStart(buf, theadTag)
	tr()
	for _, h := range t.data.Head {
		th(h)
	}
	tagEnd(buf, trTag)
	tagEnd(buf, theadTag)
	tagStart(buf, tbodyTag)
	for _, row := range t.data.Rows {
		tr()
		for _, d := range row {
			td(d)
		}
		tagEnd(buf, trTag)
	}
	tagEnd(buf, tbodyTag)
	tagEnd(buf, tableTag)

	return nil
}

func MakeSimpleDataEntityFromTable(data TableData) DataEntity {
	return &dataEntityTableSimple{data}
}

type Data struct {
	Entities []DataEntity
}

type tag string
type attr struct {
	key, val string
}

const (
	trTag    tag = "tr"
	tdTag    tag = "td"
	thTag    tag = "th"
	tableTag tag = "table"
	theadTag tag = "thead"
	tbodyTag tag = "tbody"
)

func out(buf *bytes.Buffer, ss ...string) {
	for _, s := range ss {
		buf.WriteString(s)
	}
	buf.WriteString("\n")
}

func tagStart(buf *bytes.Buffer, t tag, attrs ...attr) {
	s := "<" + string(t)
	for _, at := range attrs {
		s += " " + at.key + "='" + at.val + "'"
	}
	s += ">"
	out(buf, s)
}

func tagEnd(buf *bytes.Buffer, t tag) {
	out(buf, "</"+string(t)+">")
}

func outputTag(buf *bytes.Buffer, t tag, ss ...string) {
	tagStart(buf, t)
	if len(ss) > 0 {
		for _, s := range ss {
			out(buf, s)
		}
		tagEnd(buf, t)
	}
}

func Render(data Data) (string, error) {
	var buf bytes.Buffer
	if err := renderHTML(&buf, data); err != nil {
		return "", err
	}
	formatted := gohtml.Format(buf.String())
	return formatted, nil
}

func renderHTML(buf *bytes.Buffer, data Data) error {
	pageStart := func() error {
		css, err := base64.StdEncoding.DecodeString(cssAssets)
		if err != nil {
			return err
		}
		js, err := base64.StdEncoding.DecodeString(jsAssets)
		if err != nil {
			return err
		}
		head, err := renderTemplate(`
<!doctype html>
<html lang="en">
<head>
<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.1/css/all.min.css" rel="stylesheet"/>
<link href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap" rel="stylesheet"/>	
<style>
{{.Css}}
</style>
<script>
{{.Js}}
</script>
<script>
	$(document).ready(function () {
		$('.sortable-table').DataTable();
		$('.dataTables_length').addClass('bs-select');
		$('select[name="DataTables_Table_0_length"]').append($('<option>').attr('value', 9999999999999).text('All'));
	});		
</script>
</head>
<body>
	<div class="container-fluid">
		<a name="top"></a>
	`, "head", struct {
			Css, Js string
		}{
			Css: string(css),
			Js:  string(js),
		})
		if err != nil {
			return err
		}
		out(buf, head)
		return nil
	}

	pageEnd := func() {
		out(buf, `
		</div>
	</body>
</html>
		`)
	}

	if err := pageStart(); err != nil {
		return err
	}
	for _, e := range data.Entities {
		if err := e.Render(buf); err != nil {
			return err
		}
	}
	pageEnd()

	return nil
}

func RenderSimple(data Data) (string, error) {
	var buf bytes.Buffer
	if err := renderSimpleHTML(&buf, data); err != nil {
		return "", err
	}
	formatted := gohtml.Format(buf.String())
	return formatted, nil
}

func renderSimpleHTML(buf *bytes.Buffer, data Data) error {
	pageStart := func() error {
		out(buf, `
		<!doctype html>
		<body>
			`)
		return nil
	}

	pageEnd := func() {
		out(buf, `
	</body>
</html>
		`)
	}

	if err := pageStart(); err != nil {
		return err
	}
	for _, e := range data.Entities {
		if err := e.Render(buf); err != nil {
			return err
		}
	}
	pageEnd()

	return nil
}

func renderTemplate(t string, name string, data interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
