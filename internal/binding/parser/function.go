package parser

import (
	"sort"
	"strings"
)

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

//TODO: multipoly [][]string
//TODO: connect/disconnect slot functions + add necessary SIGNAL_* functions (check first if really needed)
//TODO: replace self poly deduction with overridden methods ?

func (f *Function) PossiblePolymorphic(self bool) ([]string, string) {
	var out = make([]string, 0)

	var params = func() []*Parameter {
		if self {
			return []*Parameter{{Name: "ptr", Value: f.ClassName()}}
		}
		return f.Parameters
	}()

	var fc, _ = f.Class()

	for _, p := range params {
		var c, exist = CurrentState.ClassMap[CleanValue(p.Value)]
		if !exist {
			continue
		}

		for _, class := range CurrentState.ClassMap {
			if class.Module != fc.Module {
				continue
			}

			if class.IsPolymorphic() && class.IsSubClassOf(c.Name) {
				out = append(out, class.Name)
			}
		}

		//TODO: multipoly
		if len(out) > 0 {
			sort.Stable(sort.StringSlice(out))
			out = append(out, c.Name)
			return out, CleanName(p.Name, p.Value)
		}
	}

	return out, ""
}
