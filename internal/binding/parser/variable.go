package parser

import (
	"fmt"
	"strings"
)

type Variable struct {
	Name     string `xml:"name,attr"`
	Fullname string `xml:"fullname,attr"`
	Href     string `xml:"href,attr"`
	Status   string `xml:"status,attr"`
	Access   string `xml:"access,attr"`
	Filepath string `xml:"filepath,attr"`
	Static   bool   `xml:"static,attr"`
	Output   string `xml:"type,attr"`
}

func (v *Variable) Class() string { return strings.Split(v.Fullname, "::")[0] }

func (v *Variable) toFunction(meta string) *Function {
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
		Parameters: []*Parameter{&Parameter{Value: v.Output}},
		TmpName:    v.Name,
	}
}
