package libs

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

var Timers map[string]int

func Init() {
	Timers = LoadJsonFromFile("timing.json")
}

func GetFormCode(driver selenium.WebDriver) int {
	return 1 // right now, I dont need auto login yet
}

func Login() {
	panic("unimplemented")
}

func Handle(driver selenium.WebDriver) {
	operations := LoadOperations("sample.json")
	fmt.Println(operations)
	process_operations(driver, operations)
}

func process_operations(driver selenium.WebDriver, operations []Operation) {
	for i, oper := range operations {
		fmt.Println(i)
		fmt.Println(oper.Action + " " + oper.Target + " : " + oper.Tag + " => " + oper.Value)

		if oper.Action == "open" {
			driver.Get(base_url + oper.Target)
		} else if oper.Action == "click" {
			if oper.Target == "css" {
				click_by_css(driver, oper.Tag)
			} else if oper.Target == "id" {
				click_by_id(driver, oper.Tag)
			} else if oper.Target == "linkText" {
				click_by_link(driver, oper.Tag)
			}
		} else if oper.Action == "type" {
			if oper.Target == "id" {
				set_input_by_id(driver, oper.Tag, oper.Value)
			}
		} else if oper.Action == "sendKeys" {
			if oper.Target == "id" {
				set_key_by_id(driver, oper.Tag, oper.Value)
			}
		}
	}
	time.Sleep(time.Duration(Timers["main_loop"]) * time.Second)
}

func set_key_by_id(driver selenium.WebDriver, id, key string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(time.Duration(Timers["after_input"]) * time.Second)
	}()

	input, err := driver.FindElement(selenium.ByID, id)
	if err == nil {
		input.SendKeys(Vkeys[key])
	}
}

func set_input_by_id(driver selenium.WebDriver, id, value string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(time.Duration(Timers["after_input"]) * time.Second)
	}()

	input, err := driver.FindElement(selenium.ByID, id)
	if err == nil {
		input.SendKeys(value)
	}
}

func click_by_link(driver selenium.WebDriver, link string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(time.Duration(Timers["after_click"]) * time.Second)
	}()

	btn, err := driver.FindElement(selenium.ByLinkText, link)
	if err == nil {
		btn.Click()
	}
}

func click_by_id(driver selenium.WebDriver, id string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(time.Duration(Timers["after_click"]) * time.Second)
	}()

	btn, err := driver.FindElement(selenium.ByID, id)
	if err == nil {
		btn.Click()
	}
}

func click_by_css(driver selenium.WebDriver, css string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(time.Duration(Timers["after_click"]) * time.Second)
	}()

	btn, err := driver.FindElement(selenium.ByCSSSelector, css)
	if err == nil {
		btn.Click()
	}

}
