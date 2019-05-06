package ui

import "fmt"

func (c *CalculatorForm) init() {
	c.InputSpinBox1.ConnectValueChanged(func(value int) { c.OutputWidget.SetText(fmt.Sprint(value + c.InputSpinBox2.Value())) })
	c.InputSpinBox2.ConnectValueChanged(func(value int) { c.OutputWidget.SetText(fmt.Sprint(value + c.InputSpinBox1.Value())) })
}
