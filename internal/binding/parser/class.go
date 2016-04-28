package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

type Class struct {
	Name      string      `xml:"name,attr"`
	Status    string      `xml:"status,attr"`
	Access    string      `xml:"access,attr"`
	Abstract  bool        `xml:"abstract,attr"`
	Bases     string      `xml:"bases,attr"`
	Module    string      `xml:"module,attr"`
	Brief     string      `xml:"brief,attr"`
	Functions []*Function `xml:"function"`
	Enums     []*Enum     `xml:"enum"`
	DocModule string
	Stub      bool
	WeakLink  map[string]bool
}

func (c *Class) register(module string) {
	c.DocModule = c.Module
	c.Module = module
	ClassMap[c.Name] = c
}

func (c *Class) removeFunctions() {
	for i := len(c.Functions) - 1; i >= 0; i-- {
		var f = c.Functions[i]
		if (f.Status == "obsolete" || f.Status == "compat") || !(f.Access == "public" || f.Access == "protected") || strings.ContainsAny(f.Name, "&<>=/!()[]{}^|*+-") || strings.Contains(f.Name, "Operator") {
			c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
		}
	}
}

func (c *Class) removeEnums() {
	for i := len(c.Enums) - 1; i >= 0; i-- {
		var e = c.Enums[i]
		if (e.Status == "obsolete" || e.Status == "compat") || !(e.Access == "public" || e.Access == "protected") {
			c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
		}
	}
}

func (c *Class) Dump() {
	var bb = new(bytes.Buffer)

	fmt.Fprintln(bb, "########################################\t\t\tFUNCTIONS\t\t\t########################################")
	for _, f := range c.Functions {
		fmt.Fprintln(bb, f)
	}

	fmt.Fprintln(bb, "########################################\t\t\tENUMS\t\t\t########################################")
	for _, e := range c.Enums {
		fmt.Fprintln(bb, e)
	}

	utils.MakeFolder(utils.GetQtPkgPath("internal", "binding", "dump", c.Module))
	utils.SaveBytes(utils.GetQtPkgPath("internal", "binding", "dump", c.Module, fmt.Sprintf("%v.txt", c.Name)), bb.Bytes())
}

func (c *Class) GetBases() []string {

	if c.Bases == "" {
		return make([]string, 0)
	}

	if strings.Contains(c.Bases, ",") {
		return strings.Split(c.Bases, ",")
	}

	return []string{c.Bases}
}

func (c *Class) GetAllBases() []string { return c.getAllBases(make([]string, 0)) }

func (c *Class) getAllBases(input []string) []string {

	var bases = c.GetBases()

	switch len(bases) {
	case 0:
		{
			return input
		}

	case 1:
		{
			if sb, exists := ClassMap[bases[0]]; exists {
				input = append(input, bases[0])
				return sb.getAllBases(input)
			}
			return input
		}

	}

	for _, b := range bases {
		input = append(input, b)
		if bs, exists := ClassMap[b]; exists {
			for _, sb := range bs.GetAllBases() {
				input = append(input, sb)
			}
		}
	}
	return input
}

func (c *Class) IsQObjectSubClass() bool {

	if c != nil {

		if c.Name == "QObject" {
			return true
		}

		for _, b := range c.GetAllBases() {
			if b == "QObject" {
				return true
			}
		}

	}

	return false
}

func (c *Class) fix() {
	if c.Name == "QStyle" {

		var defFunction Function

		for _, f := range c.Functions {
			if f.Name == "standardIcon" {
				defFunction = *f
				break
			}
		}

		defFunction.Name = "standardPixmap"
		defFunction.Output = "QPixmap"
		defFunction.Fullname = c.Name + "::" + defFunction.Name

		c.Functions = append(c.Functions, &defFunction)
	}
}
