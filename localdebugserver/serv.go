package localdebugserver

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type serv struct {
	del Delegate
}

//go:generate genopts --function Start port:int:8000 wait
func (s *serv) Start(optss ...StartOption) {
	opts := MakeStartOptions(optss...)

	var mu sync.RWMutex
	s.startHTTPServer(opts.Port(), &mu)

	if opts.Wait() {
		select {}
	}
}

func (s *serv) renderHTML(w http.ResponseWriter, mu *sync.RWMutex) error {
	mu.RLock()
	defer mu.RUnlock()

	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
		<script>
			setTimeout(function() {
				document.location=document.location;
			}, 1000);
		</script>
</head>
<body>
	<h1>Nmap Result</h1>
	<h2>{{.Now}}</h2>
	{{- if .HTMLContent}}
		<div>
			{{.HTMLContent}}
		</div>
	{{- end}}
	{{- if .TextContent}}
		<pre style="overflow:auto">
			{{.TextContent}}
		</pre>
	{{- end}}
</body>
</html>
`
	textContent, err := s.del.TextContent()
	if err != nil {
		return err
	}
	htmlContent, err := s.del.HTMLContent()
	if err != nil {
		return err
	}
	title, err := s.del.Title()
	if err != nil {
		return err
	}
	t := template.Must(template.New("serv").Parse(tmpl))
	data := struct {
		Title       string
		Now         string
		TextContent string
		HTMLContent template.HTML
	}{
		Title:       title,
		Now:         time.Now().Format("2006 15:04:05 MST"),
		TextContent: textContent,
		HTMLContent: template.HTML(htmlContent),
	}
	t.Execute(w, data)
	return nil
}

func (s *serv) startHTTPServer(port int, mu *sync.RWMutex) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := s.renderHTML(w, mu)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering HTML: %v", err), http.StatusInternalServerError)
		}
	})

	log.Printf("Starting HTTP server on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
