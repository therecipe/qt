package parser

import (
	"strings"
)

func (c *Class) remove() {
	c.removeFunctions()
	c.removeFunctions_Version()

	c.removeEnums()
	c.removeEnums_Version()

	c.removeBases()
}

func (c *Class) removeFunctions() {
	for i := len(c.Functions) - 1; i >= 0; i-- {
		f := c.Functions[i]

		switch {
		case (f.Status == "obsolete" || f.Status == "compat") ||
			!(f.Access == "public" || f.Access == "protected") ||
			strings.ContainsAny(f.Name, "&<>=/!()[]{}^|*+-") ||
			strings.Contains(f.Name, "Operator"):
			{
				c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
			}

		case (f.Virtual == IMPURE || f.Virtual == PURE) && f.Meta == CONSTRUCTOR:
			{
				c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
			}
		}
	}
}

func (c *Class) removeFunctions_Version() {
	for i := len(c.Functions) - 1; i >= 0; i-- {
		switch c.Functions[i].Fullname {
		case "QTextBrowser::isModified", "QTextBrowser::setModified":
			{
				c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
			}
		}
	}
}

func (c *Class) removeEnums() {
	for i := len(c.Enums) - 1; i >= 0; i-- {
		if e := c.Enums[i]; (e.Status == "obsolete" || e.Status == "compat") ||
			!(e.Access == "public" || e.Access == "protected") {

			c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
		}
	}
}

func (c *Class) removeEnums_Version() {
	for i := len(c.Enums) - 1; i >= 0; i-- {
		switch c.Enums[i].ClassName() {
		case "QCss", "QScript", "Http2":
			{
				c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
			}
		}
	}
}

func (c *Class) removeBases() {
	var bases = c.GetBases()
	for i := len(bases) - 1; i >= 0; i-- {
		if _, ok := State.ClassMap[bases[i]]; !ok {
			bases = append(bases[:i], bases[i+1:]...)
		}
	}
	c.Bases = strings.Join(bases, ",")
}
