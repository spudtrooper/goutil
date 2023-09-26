package localdebugserver

import (
	"encoding/json"

	"github.com/fatih/structs"
)

type Delegate interface {
	Title() (string, error)
	TextContent() (string, error)
	HTMLContent() (string, error)
}

type baseDel struct {
	title string
}

func (d *baseDel) Title() (string, error) { return d.title, nil }

type htmlDel struct {
	*baseDel
	getContent func() (string, error)
}

func (d *htmlDel) TextContent() (string, error) { return "", nil }
func (d *htmlDel) HTMLContent() (string, error) { return d.getContent() }

func NewHtmlDelegate(title string, getContent func() (string, error)) Delegate {
	return &htmlDel{
		baseDel:    &baseDel{title},
		getContent: getContent,
	}
}

type textDel struct {
	*baseDel
	getContent func() (string, error)
}

func (d *textDel) TextContent() (string, error) { return d.getContent() }
func (d *textDel) HTMLContent() (string, error) { return "", nil }

func NewTextDelegate(title string, getContent func() (string, error)) Delegate {
	return &textDel{
		baseDel:    &baseDel{title},
		getContent: getContent,
	}
}

func NewServer(del Delegate) *serv {
	return &serv{del}
}

type watchDel struct {
	res any
}

func (s *watchDel) Title() (string, error) {
	return "Nmap Result", nil
}

func (s *watchDel) TextContent() (string, error) {
	jsonData, err := json.MarshalIndent(s.res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (s *watchDel) HTMLContent() (string, error) {
	html, err := ObjectToHTML(structs.Map(s.res))
	if err != nil {
		return "", err
	}
	return html, nil
}

// NewWatchDelegate	creates a new Delegate that watches the given object and returns both HTML and text content
func NewWatchDelegate(res any) Delegate { return &watchDel{res} }
