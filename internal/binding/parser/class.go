package parser

import (
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
}

func (c *Class) register(module string) {
	c.DocModule = c.Module
	c.Module = module
	ClassMap[c.Name] = c
}

func (c *Class) registerAbstact() {
	for _, f := range c.Functions {
		if f.Virtual == "pure" || c.Abstract {
			AbstractMap[c.Name] = true
			return
		}
	}
}

func (c *Class) removeFunctions() {
	for i := len(c.Functions) - 1; i >= 0; i-- {
		var f = c.Functions[i]
		if f.Status == "obsolete" || !(f.Access == "public" || f.Access == "protected") {
			c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
		}
	}
}

func (c *Class) removeEnums() {
	for i := len(c.Enums) - 1; i >= 0; i-- {
		var e = c.Enums[i]
		if e.Status == "obsolete" || !(e.Access == "public" || e.Access == "protected") {
			c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
		}
	}
}

func (c *Class) Dump() {

	var tmp = fmt.Sprintln("########################################\t\t\tFUNCTIONS\t\t\t########################################")

	for _, f := range c.Functions {
		tmp += fmt.Sprintln(f)
	}

	tmp += fmt.Sprintln("########################################\t\t\tENUMS\t\t\t########################################")

	for _, e := range c.Enums {
		tmp += fmt.Sprintln(e)
	}

	utils.MakeFolder(utils.GetQtPkgPath("internal", "binding", "dump", c.Module))
	utils.Save(utils.GetQtPkgPath("internal", "binding", "dump", c.Module, fmt.Sprintf("%v.txt", c.Name)), tmp)
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

func (c *Class) GetAllBases(input []string) []string {

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
				return sb.GetAllBases(input)
			}
			return input
		}

	}

	for _, b := range bases {
		input = append(input, b)
		if bs, exists := ClassMap[b]; exists {
			for _, sb := range bs.GetAllBases([]string{}) {
				input = append(input, sb)
			}
		}
	}
	return input
}

func (c *Class) IsQObjectSubClass() bool {
	if c.Name == "QObject" {
		return true
	}

	for _, b := range c.GetAllBases([]string{}) {
		if b == "QObject" {
			return true
		}
	}

	return false
}
