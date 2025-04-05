package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) (err error)
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, i := range dataset {
		err := dp.Parse(i)
		if err != nil {
			fmt.Println("parse error", err)
			continue
		}
		actInfo, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("error in func ActionInfo")
		}
		fmt.Println(actInfo)
	}
}
