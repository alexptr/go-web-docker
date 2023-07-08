package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
	//comment

}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	operation := c.Ctx.Input.Param((":operation"))
	num1, _ := strconv.Atoi((c.Ctx.Input.Param(":num1")))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	//Set the values for use in the template
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		c.Data["result"] = add(num1, num2)
	case "product":
		c.Data["result"] = multiply(num1, num2)
	default:
		c.TplName = "invalid-route.html"
	}

}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
