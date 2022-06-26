package request

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
)

// FormBuilder is used to create multipart forms for requests
type FormBuilder struct {
	b             *bytes.Buffer
	w             *multipart.Writer
	closed        bool
	debugInfo     *debugInfo
	debugStringMu sync.Mutex
}

type debugInfo struct {
	fields map[string]interface{}
	files  map[string]string
}

// NewFormBuilder creates a new FormBuilder
func NewFormBuilder() *FormBuilder {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	return &FormBuilder{
		b: b,
		w: w,
		debugInfo: &debugInfo{
			fields: map[string]interface{}{},
			files:  map[string]string{},
		},
	}
}

// Field adds an entry with the key and value
func (f *FormBuilder) Field(key string, val interface{}) error {
	f.debugInfo.fields[key] = val
	fw, err := f.w.CreateFormField(key)
	if err != nil {
		return err
	}
	if _, err := io.Copy(fw, strings.NewReader(fmt.Sprintf("%v", val))); err != nil {
		return err
	}
	return nil
}

// File adds an entry with the key and file
func (f *FormBuilder) File(key, file string) error {
	f.debugInfo.files[key] = file
	r, err := os.Open(file)
	if err != nil {
		return err
	}
	fw, err := f.w.CreateFormFile(key, path.Base(file))
	if err != nil {
		return err
	}
	if _, err := io.Copy(fw, r); err != nil {
		return err
	}
	return nil
}

func (f *FormBuilder) close() {
	if !f.closed {
		f.w.Close()
		f.closed = true
	}
}

// String closes the underlying form and returns a string version of the form
func (f *FormBuilder) String() string {
	f.close()
	return f.b.String()
}

// Bytes closes the underlying form and returns a []byte version of the form
func (f *FormBuilder) Bytes() []byte {
	f.close()
	return f.b.Bytes()
}

// ContentType closes the underlying form and returns the "content-type" header value used in the form
func (f *FormBuilder) ContentType() string {
	f.close()
	return f.w.FormDataContentType()
}

// DebugInfo returns a struct for debugging.
func (f *FormBuilder) DebugInfo() debugInfo {
	return *f.debugInfo
}

// DebugString returns a string for debugging.
func (f *FormBuilder) DebugString() string {
	f.debugStringMu.Lock()
	defer f.debugStringMu.Unlock()

	var buf bytes.Buffer
	out := func(tmpl string, args ...interface{}) {
		buf.WriteString(fmt.Sprintf(tmpl+"\n", args...))
	}

	out("Fields (%d)", len(f.debugInfo.fields))
	{
		var keys []string
		for k := range f.debugInfo.fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			out(" [%d] %s: %v", i+1, k, f.debugInfo.fields[k])
		}
	}

	out("Files (%d)", len(f.debugInfo.files))
	{
		var keys []string
		for k := range f.debugInfo.files {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			out(" [%d] %s: %s", i+1, k, f.debugInfo.files[k])
		}
	}

	return strings.TrimSpace(buf.String())
}
