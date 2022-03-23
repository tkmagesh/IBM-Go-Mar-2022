package main

import (
	"fmt"
	"modularity-demo/calculator"
	calc "modularity-demo/calculator" //package alias
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("modularity demo")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("opCount = ", calculator.OpCount())

	fmt.Println(utils.IsEven(20))
	fmt.Println(calc.Add(100, 2000))

	//using 3rd party packages
	color.Red("This line will be printed in red color")
}
