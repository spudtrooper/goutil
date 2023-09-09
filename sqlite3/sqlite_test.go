package util

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := map[string]string{
		"eventID":        "event_id",
		"EventID":        "event_id",
		"myStringVar":    "my_string_var",
		"JSONData":       "json_data",
		"AValue":         "a_value",
		"SimpleCase":     "simple_case",
		"URLParser":      "url_parser",
		"ID":             "id",
		"GoToHome":       "go_to_home",
		"XMLHTTPRequest": "xmlhttp_request",
		"UserID":         "user_id",
		"AbC":            "ab_c",
		"aB":             "a_b",
		"AB":             "ab",
		"UserIdList":     "user_id_list",
		"ABCData":        "abc_data",
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			result := toSnakeCase(input)
			if result != expected {
				t.Errorf("Expected %s but got %s", expected, result)
			}
		})
	}
}
