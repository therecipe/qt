package parser

import (
	"fmt"
	"strings"
)

type Class struct {
	Name       string      `xml:"name,attr"`
	Status     string      `xml:"status,attr"`
	Access     string      `xml:"access,attr"`
	Abstract   bool        `xml:"abstract,attr"`
	Bases      string      `xml:"bases,attr"`
	Module     string      `xml:"module,attr"`
	Brief      string      `xml:"brief,attr"`
	Functions  []*Function `xml:"function"`
	Enums      []*Enum     `xml:"enum"`
	Variables  []*Variable `xml:"variable"`
	Properties []*Variable `xml:"property"`
	Classes    []*Class    `xml:"class"`

	DocModule string
	Stub      bool
	WeakLink  map[string]struct{}
	Export    bool
	Fullname  string
}

func (c *Class) register(m string) {

	c.DocModule = c.Module
	c.Module = m
	CurrentState.ClassMap[c.Name] = c

	for _, sc := range c.Classes {
		if sc.Name != "PaintContext" { //TODO: remove
			continue
		}

		sc.Fullname = fmt.Sprintf("%v::%v", c.Name, sc.Name)
		sc.register(m)
	}
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

func (c *Class) GetAllBases() []string {

	var input = make([]string, 0)

	for _, b := range c.GetBases() {
		var bc, exists = CurrentState.ClassMap[b]
		if !exists {
			continue
		}

		input = append(input, b)
		for _, sbc := range bc.GetAllBases() {
			input = append(input, sbc)
		}
	}

	return input
}

func (c *Class) GetAllBasesRecursiveCheckFailed(i int) ([]string, bool) {
	var input = make([]string, 0)

	i++
	if i > 100 {
		return input, true
	}

	for _, b := range c.GetBases() {
		var bc, exists = CurrentState.ClassMap[b]
		if !exists {
			continue
		}

		input = append(input, b)
		var bs, isRecursive = bc.GetAllBasesRecursiveCheckFailed(i)
		if isRecursive {
			return input, true
		}
		for _, sbc := range bs {
			input = append(input, sbc)
		}
	}

	return input, false
}

func (c *Class) IsSubClassOfQObject() bool {
	return c.IsSubClassOf("QObject")
}

func (c *Class) IsSubClassOf(class string) bool {
	if c == nil {
		return false
	}

	for _, b := range append([]string{c.Name}, c.GetAllBases()...) {
		if b != class {
			continue
		}

		return true
	}

	return false
}

func (c *Class) isSubClass() bool { return c.Fullname != "" }

func (c *Class) HasFunctionWithName(n string) bool {
	return c.HasFunctionWithNameAndOverloadNumber(n, "")
}

func (c *Class) HasFunctionWithNameAndOverloadNumber(n string, num string) bool {
	for _, f := range c.Functions {
		if strings.ToLower(f.Name) == strings.ToLower(n) && f.OverloadNumber == num {
			return true
		}
	}

	return false
}

func (c *Class) IsPolymorphic() bool { return len(c.GetBases()) > 1 }

func (c *Class) HasConstructor() bool {
	for _, f := range c.Functions {
		if f.Meta == CONSTRUCTOR || f.Meta == COPY_CONSTRUCTOR || f.Meta == MOVE_CONSTRUCTOR {
			return true
		}
	}
	return false
}

func (c *Class) NeedsDestructor() bool {
	for _, f := range c.Functions {
		if f.Meta == DESTRUCTOR {
			return false
		}
	}
	return true
}
