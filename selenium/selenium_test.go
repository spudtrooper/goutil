package selenium

import "testing"

func TestMakeWebDriver(t *testing.T) {
	wd, cancel, err := MakeWebDriver(MakeWebDriverOptions{
		Verbose:  true,
		Headless: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer cancel()
	if wd == nil {
		t.Errorf("nil wd")
	}
}
