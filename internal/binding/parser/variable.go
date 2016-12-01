package parser

import (
	"fmt"
	"strings"
)

type Variable struct {
	Name     string    `xml:"name,attr"`
	Fullname string    `xml:"fullname,attr"`
	Href     string    `xml:"href,attr"`
	Status   string    `xml:"status,attr"`
	Access   string    `xml:"access,attr"`
	Filepath string    `xml:"filepath,attr"`
	Static   bool      `xml:"static,attr"`
	Output   string    `xml:"type,attr"`
	Brief    string    `xml:"brief,attr"`
	Getter   []*getter `xml:"getter"`
	Setter   []*setter `xml:"setter"`
}
type getter struct{}
type setter struct{}

func (v *Variable) Class() string {
	var s = strings.Split(v.Fullname, "::")
	if len(s) == 3 {
		return s[1]
	}
	return s[0]
}

func (v *Variable) variableToFunction(meta string) *Function {
	if meta == GETTER {
		return &Function{
			Name:     v.Name,
			Fullname: v.Fullname,
			Href:     v.Href,
			Status:   v.Status,
			Access:   v.Access,
			Filepath: v.Filepath,
			Static:   v.Static,
			Output:   v.Output,
			Meta:     meta,
			Brief:    v.Brief,
		}
	}
	return &Function{
		Name:       fmt.Sprintf("set%v", strings.Title(v.Name)),
		Fullname:   fmt.Sprintf("%v::set%v", v.Class(), strings.Title(v.Name)),
		Href:       v.Href,
		Status:     v.Status,
		Access:     v.Access,
		Filepath:   v.Filepath,
		Static:     v.Static,
		Output:     "void",
		Meta:       meta,
		Parameters: []*Parameter{{Value: v.Output}},
		TmpName:    v.Name,
		Brief:      v.Brief,
	}
}

func (v *Variable) propertyToFunctions(c *Class) []*Function {
	var funcs []*Function

	if v.Fullname == "QTextBrowser::modified" {
		return funcs
	}

	if !(c.HasFunctionWithName(v.Name) || c.HasFunctionWithName(fmt.Sprintf("is%v", strings.Title(v.Name))) || c.HasFunctionWithName(fmt.Sprintf("has%v", strings.Title(v.Name)))) && len(v.Getter) == 0 {

		var tmpV = *v

		switch tmpV.Fullname {
		case "QGraphicsObject::z":
			{
				tmpV.Name = "zValue"
				tmpV.Fullname = fmt.Sprintf("%v::%v", tmpV.Class(), tmpV.Name)
			}

		case "QGraphicsObject::effect":
			{
				tmpV.Name = "graphicsEffect"
				tmpV.Fullname = fmt.Sprintf("%v::%v", tmpV.Class(), tmpV.Name)
			}

		default:
			{
				if tmpV.Output == "bool" {
					tmpV.Name = fmt.Sprintf("is%v", strings.Title(tmpV.Name))
					tmpV.Fullname = fmt.Sprintf("%v::%v", tmpV.Class(), tmpV.Name)
				}
			}
		}

		funcs = append(funcs, &Function{
			Name:     tmpV.Name,
			Fullname: tmpV.Fullname,
			Href:     tmpV.Href,
			Status:   tmpV.Status,
			Access:   tmpV.Access,
			Filepath: tmpV.Filepath,
			Static:   tmpV.Static,
			Output:   tmpV.Output,
			Meta:     PLAIN,
			Brief:    tmpV.Brief,
		})
	}

	if !c.HasFunctionWithName(fmt.Sprintf("set%v", strings.Title(v.Name))) && len(v.Getter) == 0 && len(v.Setter) == 0 {

		switch v.Fullname {
		case "QGraphicsObject::z":
			{
				v.Name = "zValue"
				v.Fullname = fmt.Sprintf("%v::%v", v.Class(), v.Name)
			}

		case "QGraphicsObject::effect":
			{
				v.Name = "graphicsEffect"
				v.Fullname = fmt.Sprintf("%v::%v", v.Class(), v.Name)
			}
		}

		funcs = append(funcs,
			&Function{
				Name:       fmt.Sprintf("set%v", strings.Title(v.Name)),
				Fullname:   fmt.Sprintf("%v::set%v", v.Class(), strings.Title(v.Name)),
				Href:       v.Href,
				Status:     v.Status,
				Access:     v.Access,
				Filepath:   v.Filepath,
				Static:     v.Static,
				Output:     "void",
				Meta:       PLAIN,
				Parameters: []*Parameter{{Value: v.Output}},
				Brief:      v.Brief,
			})

		for _, bc := range c.GetAllBases() {
			for _, bcf := range ClassMap[bc].Functions {
				if bcf.Name == fmt.Sprintf("set%v", strings.Title(v.Name)) && bcf.Overload {
					funcs = append(funcs,
						&Function{
							Name:           fmt.Sprintf("set%v", strings.Title(v.Name)),
							Fullname:       fmt.Sprintf("%v::set%v", v.Class(), strings.Title(v.Name)),
							Href:           bcf.Href,
							Status:         bcf.Status,
							Access:         bcf.Access,
							Filepath:       bcf.Filepath,
							Static:         bcf.Static,
							Output:         "void",
							Meta:           PLAIN,
							Parameters:     bcf.Parameters,
							Brief:          bcf.Brief,
							Overload:       true,
							OverloadNumber: bcf.OverloadNumber,
						})
				}
			}
		}
	}

	return funcs
}
