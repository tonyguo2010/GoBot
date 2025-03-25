package libs

import (
	"fmt"

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
	operations := LoadOperations("bns.json")
	fmt.Println(operations)
}
