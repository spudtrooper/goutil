package request

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// FormBuilder is used to create multipart forms for requests
type FormBuilder struct {
	b      *bytes.Buffer
	w      *multipart.Writer
	closed bool
}

// NewFormBuilder creates a new FormBuilder
func NewFormBuilder() *FormBuilder {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	return &FormBuilder{
		b: b,
		w: w,
	}
}

// Field adds an entry with the key and value
func (f *FormBuilder) Field(key, val string) error {
	fw, err := f.w.CreateFormField(key)
	if err != nil {
		return err
	}
	if _, err := io.Copy(fw, strings.NewReader(val)); err != nil {
		return err
	}
	return nil
}

// File adds an entry with the key and file
func (f *FormBuilder) File(key, file string) error {
	fw, err := f.w.CreateFormFile(key, path.Base(file))
	if err != nil {
		return err
	}
	r, err := os.Open(file)
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
