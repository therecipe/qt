package parser

import "strings"

type Function struct {
	Name             string       `xml:"name,attr"`
	Fullname         string       `xml:"fullname,attr"`
	Href             string       `xml:"href,attr"`
	Status           string       `xml:"status,attr"`
	Access           string       `xml:"access,attr"`
	Filepath         string       `xml:"filepath,attr"`
	Virtual          string       `xml:"virtual,attr"`
	Meta             string       `xml:"meta,attr"`
	Static           bool         `xml:"static,attr"`
	Overload         bool         `xml:"overload,attr"`
	OverloadNumber   string       `xml:"overload-number,attr"`
	Output           string       `xml:"type,attr"`
	Signature        string       `xml:"signature,attr"`
	Parameters       []*Parameter `xml:"parameter"`
	Brief            string       `xml:"brief,attr"`
	SignalMode       string
	TemplateModeJNI  string
	Default          bool
	TmpName          string
	Export           bool
	NeedsFinalizer   bool
	Container        string
	TemplateModeGo   string
	Child            *Function
	NonMember        bool
	NoMocDeduce      bool
	PureBaseFunction bool
	AsError          bool
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"left,attr"`
}

func (f *Function) Class() (*Class, bool) {
	var class, exist = CurrentState.ClassMap[f.ClassName()]
	return class, exist
}
func (f *Function) ClassName() string {
	var s = strings.Split(f.Fullname, "::")
	if len(s) == 3 {
		return s[1]
	}
	return s[0]
}
