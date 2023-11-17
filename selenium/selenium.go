package selenium

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/go-errors/errors"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/or"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type MakeWebDriverOptions struct {
	Verbose          bool
	Headless         bool
	ChromeDriverPath string
	SeleniumPath     string
	Port             int
	Host             string
}

type WebDriverProvider func() (selenium.WebDriver, func(), error)

func MakeWebDriverProvider(opts MakeWebDriverOptions) WebDriverProvider {
	var wd selenium.WebDriver
	var cancel func()
	return func() (selenium.WebDriver, func(), error) {
		if wd != nil && cancel != nil {
			return wd, cancel, nil
		}
		w, c, err := makeWebDriver(opts)
		wd = w
		cancel = c
		return w, c, err
	}
}

func MakeWebDriver(opts MakeWebDriverOptions) (selenium.WebDriver, func(), error) {
	return makeWebDriver(opts)
}

func makeWebDriver(opts MakeWebDriverOptions) (selenium.WebDriver, func(), error) {
	seleniumPath := opts.SeleniumPath
	if seleniumPath == "" {
		seleniumPathTmp, err := getSeleniumPath()
		if err != nil {
			return nil, nil, errors.Errorf("couldn't find selenium jar")
		}
		seleniumPath = seleniumPathTmp
	}
	port := or.Int(opts.Port, 8082)
	chromeDriverPath := or.String(opts.ChromeDriverPath, findChromeDriver())
	if chromeDriverPath == "" {
		return nil, nil, errors.Errorf("couldn't find chromedriver")
	}
	selOpts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
	}
	if opts.Verbose {
		selOpts = append(selOpts, selenium.Output(os.Stderr))
		selenium.SetDebug(true)
	} else {
		selenium.SetDebug(false)
	}
	service, err := selenium.NewSeleniumService(seleniumPath, port, selOpts...)
	if err != nil {
		return nil, nil, err
	}

	args := []string{
		"--no-sandbox",
		"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
	}
	if opts.Headless {
		args = append(args, "--headless")
	}
	if !opts.Verbose {
		args = append(args, "--silent")
	}
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: args,
	}
	caps.AddChrome(chromeCaps)
	var host string
	if opts.Host != "" {
		host = opts.Host
	} else {
		host = fmt.Sprintf("http://localhost:%d/wd/hub", port)
	}
	wd, err := selenium.NewRemote(caps, host)
	if err != nil {
		return nil, nil, err
	}

	return wd, func() {
		wd.Quit()
		service.Stop()
	}, nil
}

func findChromeDriver() string {
	paths := []string{
		"/opt/homebrew/bin/chromedriver",
		"/usr/local/bin/chromedriver",
	}
	for _, f := range paths {
		if io.FileExists(f) {
			return f
		}
	}
	return ""
}

func getSeleniumPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	res := path.Join(home, ".selenium.jar")
	if io.FileExists(res) {
		log.Printf("selenium jar present: %s", res)
	} else {
		decoded, err := base64.StdEncoding.DecodeString(seleniumServerJar)
		if err != nil {
			return "", errors.Errorf("base64.StdEncoding: %v", err)
		}
		log.Printf("writing selenium jar to %s", res)
		if err := ioutil.WriteFile(res, decoded, 0755); err != nil {
			return "", errors.Errorf("ioutil.WriteFile(%q): %v", res, err)
		}
	}
	return res, nil
}

func FindElement(wd selenium.WebDriver, tag, text string) (selenium.WebElement, error) {
	els, err := wd.FindElements(selenium.ByTagName, tag)
	if err != nil {
		return nil, err
	}
	for _, el := range els {
		txt, err := el.Text()
		if err != nil {
			return nil, err
		}
		if txt == text {
			return el, nil
		}
	}
	return nil, nil
}

func FindElements(wd selenium.WebDriver, tag, text string) ([]selenium.WebElement, error) {
	els, err := wd.FindElements(selenium.ByTagName, tag)
	if err != nil {
		return nil, err
	}
	res := []selenium.WebElement{}
	for _, el := range els {
		txt, err := el.Text()
		if err != nil {
			return nil, err
		}
		if txt == text {
			res = append(res, el)
		}
	}
	return res, nil
}

// TODO: This is stupid, should remove and have callers specify "button"
func FindButton(wd selenium.WebDriver, text string) (selenium.WebElement, error) {
	return FindElement(wd, "button", text)
}

//go:generate genopts --function WaitForButton --extends Wait
func WaitForButton(wd selenium.WebDriver, btnText string, optss ...WaitForButtonOption) (selenium.WebElement, error) {
	opts := MakeWaitForButtonOptions(optss...)
	msg, limit := opts.Message(), opts.Limit()

	prefix := ""
	if msg != "" {
		prefix = fmt.Sprintf("[%s] ", msg)
	}

	var res selenium.WebElement
	var cnt int
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		if limit > 0 && cnt == limit {
			log.Printf("%shit limit=%d after %d tries, quitting ...", prefix, limit, cnt+1)
			return true, nil
		}
		log.Printf("[%d] %swaiting for button %s ...", cnt+1, prefix, btnText)
		cnt++
		btn, err := FindButton(wd, btnText)
		if err != nil {
			return false, err
		}
		if btn != nil {
			res = btn
			return true, nil
		}
		return false, nil
	})
	if res == nil {
		return nil, fmt.Errorf("couldn't find button with text: %s", btnText)
	}
	return res, nil
}

//go:generate genopts --function Wait message:string limit:int
func Wait(wd selenium.WebDriver, fn func(wd selenium.WebDriver) (bool, error), optss ...WaitOption) error {
	opts := MakeWaitOptions(optss...)
	msg, limit := opts.Message(), opts.Limit()

	prefix := ""
	if msg != "" {
		prefix = fmt.Sprintf("[%s] ", msg)
	}
	var cnt int
	err := wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		if limit > 0 && cnt == limit {
			log.Printf("%shit limit=%d after %d tries, quitting ...", prefix, limit, cnt+1)
			return true, nil
		}

		log.Printf("[%d/%d] %swaiting ...", cnt+1, limit, prefix)
		cnt++
		return fn(wd)
	})
	return err
}

// TODO: Refactor Wait* to use Wait above

//go:generate genopts --function WaitForElement --extends Wait
func WaitForElement(wd selenium.WebDriver, tagName, text string, optss ...WaitForElementOption) (selenium.WebElement, error) {
	opts := MakeWaitForElementOptions(optss...)
	msg, limit := opts.Message(), opts.Limit()

	prefix := ""
	if msg != "" {
		prefix = fmt.Sprintf("[%s] ", msg)
	}

	var res selenium.WebElement
	var cnt int
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		if limit > 0 && cnt == limit {
			log.Printf("%shit limit=%d after %d tries, quitting ...", prefix, limit, cnt+1)
			return true, nil
		}

		log.Printf("[%d] %swaiting for div %s %q...", cnt+1, prefix, text, tagName)
		cnt++
		btn, err := FindElement(wd, tagName, text)
		if err != nil {
			return false, err
		}
		if btn != nil {
			res = btn
			return true, nil
		}
		return false, nil
	})
	if res == nil {
		return nil, fmt.Errorf("couldn't find div with text: %s", text)
	}
	return res, nil
}

//go:generate genopts --function WaitForElements --extends Wait
func WaitForElements(wd selenium.WebDriver, tagName, text string, optss ...WaitForElementsOption) ([]selenium.WebElement, error) {
	opts := MakeWaitForElementsOptions(optss...)
	msg, limit := opts.Message(), opts.Limit()

	prefix := ""
	if msg != "" {
		prefix = fmt.Sprintf("[%s] ", msg)
	}

	var res []selenium.WebElement
	var cnt int
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		if limit > 0 && cnt == limit {
			log.Printf("%shit limit=%d after %d tries, quitting ...", prefix, limit, cnt+1)
			return true, nil
		}

		log.Printf("[%d] %swaiting for div %s %q [%d] ...", cnt+1, prefix, text, tagName)
		cnt++
		els, err := FindElements(wd, tagName, text)
		if err != nil {
			return false, err
		}
		if len(els) > 0 {
			res = els
			return true, nil
		}
		return false, nil
	})
	if res == nil {
		return nil, fmt.Errorf("couldn't find div with text: %s", text)
	}
	return res, nil
}

// From: https://github.com/lucasmdomingues/go-selenium-screenshot/blob/master/print.go
func TakeScreenshot(wd selenium.WebDriver, outFile string) error {
	ss, err := wd.Screenshot()
	if err != nil {
		return err
	}
	r := bytes.NewReader(ss)

	im, err := png.Decode(r)
	if err != nil {
		return err
	}

	log.Printf("writing to %s...", outFile)
	f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	png.Encode(f, im)

	return nil
}

// https://stackoverflow.com/questions/24267413/how-to-get-parent-of-webelement
func Parent(el selenium.WebElement) (selenium.WebElement, error) {
	return el.FindElement(selenium.ByXPATH, "./..")
}
