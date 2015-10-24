package parser

import "strings"

type Enum struct {
	Name     string   `xml:"name,attr"`
	Fullname string   `xml:"fullname,attr"`
	Status   string   `xml:"status,attr"`
	Access   string   `xml:"access,attr"`
	Typedef  string   `xml:"typedef,attr"`
	Values   []*Value `xml:"value"`
}

type Value struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (e *Enum) Class() string { return strings.Split(e.Fullname, "::")[0] }

func (e *Enum) register(module string) {
	SubnamespaceMap[e.Class()] = true
	if c, exists := ClassMap[e.Class()]; !exists {
		ClassMap[e.Class()] = &Class{Name: e.Class(), Status: "commendable", Module: module, Access: "public", Enums: []*Enum{e}}
	} else {
		c.Enums = append(c.Enums, e)
	}
}
