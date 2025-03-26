package libs

import (
	"errors"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Main() {
	// Run Chrome browser
	if _, err := os.Stat("./chromedriver.exe"); errors.Is(err, os.ErrNotExist) {
		panic("chromedriver.exe is not found, visit https://googlechromelabs.github.io/chrome-for-testing/ to download the driver which is the same version with your Chrome. ")
	}

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
			time.Sleep(3 * time.Second)
			continue
		}
		form_code := GetFormCode(driver)
		if form_code == 0 {
			Login()
		} else if form_code == 1 {
			NavigateTo(driver, base_url)
		}
		time.Sleep(3 * time.Second)
	}
}

func NavigateTo(driver selenium.WebDriver, base_url string) {
	Handle(driver)
}

func GetBrowserState(driver selenium.WebDriver) bool {
	return true
}
