package libs

import (
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var base_url string

func Main(url string) {
	base_url = url
	// Run Chrome browser
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--start-maximized",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	Init()
	for {
		state := GetBrowserState(driver)
		if state == false {
			time.Sleep(time.Duration(Timers["main_loop"]) * time.Second)
			continue
		}
		form_code := GetFormCode(driver)
		if form_code == 0 {
			Login()
		} else if form_code == 1 {
			NavigateTo(driver, base_url)
		}

		time.Sleep(time.Duration(Timers["main_loop"]) * time.Second)
	}
}

func NavigateTo(driver selenium.WebDriver, base_url string) {
	driver.Get(base_url)
	Handle(driver)
}

func GetBrowserState(driver selenium.WebDriver) bool {
	return true
}
