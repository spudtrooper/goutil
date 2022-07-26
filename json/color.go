package json

import (
	"encoding/json"
	realjson "encoding/json"

	colorjson "github.com/TylerBrock/colorjson"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/or"
)

//go:generate genopts --function ColorMarshal "indent:int"
func ColorMarshal(x interface{}, optss ...ColorMarshalOption) (string, error) {
	opts := MakeColorMarshalOptions(optss...)
	indent := or.Int(opts.Indent(), 2)
	b, err := realjson.Marshal(x)
	if err != nil {
		return "", err
	}
	var obj map[string]interface{}
	json.Unmarshal(b, &obj)
	f := colorjson.NewFormatter()
	f.Indent = indent
	s, err := f.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func MustColorMarshal(x interface{}, optss ...ColorMarshalOption) string {
	s, err := ColorMarshal(x, optss...)
	check.Err(err)
	return s
}
