package libs

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func Init() {
	fmt.Println("For testing")
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
		}

		// if operation['command'] == 'type':
		//     if cmds[0] == 'id':
		//         if header is not None and header == cmds[1]:
		//             set_input_by_id(driver, header, hijack_input)
		//         else:
		//             set_input_by_id(driver, cmds[1], operation['value'])

		// elif cmds[0] == 'Glb':
		//     times = 1
		//     if 'times' in operation:
		//         times = operation['times']
		//     while times > 0:
		//         global_key(driver, operation['value'])
		//         times -= 1

		// # elif cmds[0] == 'Cmb':
		// elif cmds[0] == 'Output':
		//     value = ''

		//     if cmds[1] == 'Path':
		//         value = get_input_by_xpath(driver, widgets[cmds[2]]['xpath'])
		//     elif cmds[1] == 'ID':
		//         value = get_input_by_id(driver, widgets[cmds[2]]['id'])

		//     print(value)

		// elif cmds[0] == 'Exit':
		//     driver.close()
		//     # sys.exit(0)

	}
	time.Sleep(20 * time.Second)
}

func click_by_link(driver selenium.WebDriver, link string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		time.Sleep(3 * time.Second)
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
		time.Sleep(3 * time.Second)
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
		time.Sleep(3 * time.Second)
	}()

	btn, err := driver.FindElement(selenium.ByCSSSelector, css)
	if err == nil {
		btn.Click()
	}

}
