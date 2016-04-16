package parser

import (
	"strconv"
	"strings"
)

type Function struct {
	Name           string       `xml:"name,attr"`
	Fullname       string       `xml:"fullname,attr"`
	Href           string       `xml:"href,attr"`
	Status         string       `xml:"status,attr"`
	Access         string       `xml:"access,attr"`
	Filepath       string       `xml:"filepath,attr"`
	Virtual        string       `xml:"virtual,attr"`
	Meta           string       `xml:"meta,attr"`
	Static         bool         `xml:"static,attr"`
	Overload       bool         `xml:"overload,attr"`
	OverloadNumber string       `xml:"overload-number,attr"`
	Output         string       `xml:"type,attr"`
	Signature      string       `xml:"signature,attr"`
	Parameters     []*Parameter `xml:"parameter"`
	SignalMode     string
	TemplateMode   string
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"left,attr"`
}

func (f *Function) Class() string { return strings.Split(f.Fullname, "::")[0] }

func (f *Function) register(module string) {
	if c, exists := ClassMap[f.Class()]; !exists {
		ClassMap[f.Class()] = &Class{Name: f.Class(), Module: module, Access: "public", Functions: []*Function{f}}
	} else {
		c.Functions = append(c.Functions, f)
	}
}

func (f *Function) fix() {
	if f.Fullname == "QThread::start" {
		f.Parameters = make([]*Parameter, 0)
	}

	if f.Fullname == "QQuickImageProvider::requestImage" || f.Fullname == "QQuickImageProvider::requestPixmap" {
		f.Parameters = f.Parameters[:len(f.Parameters)-1]
	}
}

func (f *Function) fixOverload() {

	if strings.Contains(f.Href, "-") {
		var tmp, err = strconv.Atoi(strings.Split(f.Href, "-")[1])
		if err == nil && tmp > 0 {
			f.Overload = true
			f.OverloadNumber = strconv.Itoa(tmp)
		}
	}

	if f.OverloadNumber == "0" {
		f.Overload = false
		f.OverloadNumber = ""
	}
}
