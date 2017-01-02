package parser

import (
	"fmt"
	"strings"
)

func (c *Class) add() {
	c.addGeneralFuncs()

	c.addVarAndPropFuncs()

	c.addMocFuncs()
}

func (c *Class) addGeneralFuncs() {
	switch c.Name {
	case "QColor", "QFont", "QImage", "QObject":
		{
			c.Functions = append(c.Functions, &Function{
				Name:       "toVariant",
				Fullname:   fmt.Sprintf("%v::toVariant", c.Name),
				Access:     "public",
				Virtual:    "non",
				Meta:       PLAIN,
				Output:     "QVariant",
				Parameters: []*Parameter{},
				Signature:  "()",
			})
		}

	case "QVariant":
		{
			for _, name := range []string{"toColor", "toFont", "toImage", "toObject"} {
				c.Functions = append(c.Functions, &Function{
					Name:       name,
					Fullname:   fmt.Sprintf("%v::%v", c.Name, name),
					Access:     "public",
					Virtual:    "non",
					Meta:       PLAIN,
					Output:     strings.Replace(name, "to", "Q", -1),
					Parameters: []*Parameter{},
					Signature:  "()",
				})
			}
		}
	}

	if c.Name == "QQmlNetworkAccessManagerFactory" && !c.HasConstructor() {
		c.Functions = append(c.Functions, &Function{
			Name:       c.Name,
			Fullname:   fmt.Sprintf("%v::%v", c.Name, c.Name),
			Access:     "public",
			Virtual:    "non",
			Meta:       CONSTRUCTOR,
			Parameters: []*Parameter{},
			Signature:  "()",
		})
	}
}

func (c *Class) addVarAndPropFuncs() {
	for _, v := range c.Variables {
		c.Functions = append(c.Functions, v.varToFunc()...)
	}
	for _, p := range c.Properties {
		c.Functions = append(c.Functions, p.propToFunc(c)...)
	}
}

func (c *Class) addMocFuncs() {
	if !State.Moc {
		return
	}

	//add generic qRegisterMetaType functions
	if c.HasFunctionWithName("qRegisterMetaType") {
		return
	}

	var tmpF = &Function{
		Name:           "qRegisterMetaType",
		Fullname:       fmt.Sprintf("%v::qRegisterMetaType", c.Name),
		Access:         "public",
		Virtual:        "non",
		Meta:           PLAIN,
		NonMember:      true,
		NoMocDeduce:    true,
		Static:         true,
		Output:         fmt.Sprintf("int"),
		Parameters:     []*Parameter{},
		Signature:      "()",
		TemplateModeGo: fmt.Sprintf("%v*", c.Name),
	}
	c.Functions = append(c.Functions, tmpF)

	var tmpF2 = *tmpF
	tmpF2.Overload = true
	tmpF2.OverloadNumber = "2"
	tmpF2.Parameters = []*Parameter{{Name: "typeName", Value: "const char *"}}
	tmpF2.Signature = "(const char *typeName)"
	c.Functions = append(c.Functions, &tmpF2)
}
