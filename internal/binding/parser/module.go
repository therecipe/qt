package parser

import (
	"fmt"
	"strings"

	"github.com/StarAurryon/qt/internal/utils"
)

type Module struct {
	Namespace *Namespace `xml:"namespace"`
	Project   string     `xml:"project,attr"`
	Pkg       string
}

type Namespace struct {
	Classes   []*Class    `xml:"class"`
	Structs   []*Class    `xml:"struct"`
	Functions []*Function `xml:"function"`
	//Enums         []*Enum         `xml:"enum"`     //TODO: uncomment
	SubNamespaces []*SubNamespace `xml:"namespace"`
}

type SubNamespace struct {
	//Classes   []*Class    `xml:"class"`    //TODO: uncomment
	Functions []*Function `xml:"function"`
	Enums     []*Enum     `xml:"enum"`
	Status    string      `xml:"status,attr"`
	Access    string      `xml:"access,attr"`
}

func (m *Module) Prepare() error {
	if utils.QT_API_NUM(utils.QT_VERSION()) >= 5120 && m.Project == "QtQuickControls" {
		m.Project = "QtQuickControls2"
	}

	utils.Log.WithField("module", strings.TrimPrefix(m.Project, "Qt")).Debug("prepare")

	m.Namespace.Classes = append(m.Namespace.Classes, m.Namespace.Structs...)

	//register classes from namespace
	for _, c := range m.Namespace.Classes {
		c.register(m)
	}
	m.add()

	for _, f := range m.Namespace.Functions {
		if m.Project == "QtQml" && f.Name == "qjsEngine" {
			f.Fullname = "QJSEngine::qjsEngine"
			f.register(m.Project)
		}
		if m.Project != "QtCore" || !(strings.Contains(f.Name, "Environment") || strings.HasSuffix(f.Name, "env") || f.Name == "qVersion") {
			continue
		}
		f.Fullname = strings.Replace(strings.Replace(f.Fullname, "<", "", -1), ">", "", -1)
		f.Static = true
		if f.Name == "qEnvironmentVariable" && len(f.Parameters) == 2 {
			f.Overload = true
			f.OverloadNumber = "2"
		}
		if !strings.HasPrefix(f.Fullname, "QtGlobal") {
			f.Fullname = fmt.Sprintf("%v::%v", "QtGlobal", f.Name)
		}
		f.register(m.Project)
	}

	//register enums and functions from subnamespaces
	var snsExtraClasses []*Class
	for _, sns := range m.Namespace.SubNamespaces {
		for _, e := range sns.Enums {
			if !(e.Status == "active" || e.Status == "commendable") || !(e.Access == "public" || e.Access == "protected") ||
				strings.Contains(e.Fullname, "Private") || strings.Contains(e.Fullname, "Util") ||
				strings.Contains(e.Fullname, "nternal") || strings.ToLower(e.Name) == e.Name {
				continue
			}
			e.register(m.Project)
		}
		if m.Project != "QtSensors" && m.Project != "QtXmlPatterns" &&
			m.Project != "QtQml" && m.Project != "QtWidgets" && m.Project != "QtMacExtras" &&
			m.Project != "QtTestLib" && m.Project != "QtScript" && m.Project != "QtQuick" {
			for _, f := range sns.Functions {

				if !(f.Status == "active" || f.Status == "commendable") || !(f.Access == "public" || f.Access == "protected") ||
					strings.Contains(f.Fullname, "Private") || strings.Contains(f.Fullname, "Util") ||
					strings.Contains(f.Fullname, "nternal") || f.Name == "qDefaultSurfaceFormat" ||
					f.ClassName() == "QUnicodeTables" ||
					f.ClassName() == "QUtf8Functions" ||
					f.ClassName() == "QUnicodeTools" ||
					f.ClassName() == "HPack" ||
					f.ClassName() == "QHighDpi" ||
					f.ClassName() == "QPdf" ||
					f.ClassName() == "QPlatformGraphicsBufferHelper" ||
					f.ClassName() == "QIcu" ||
					strings.ToLower(f.ClassName()) == f.ClassName() {
					continue
				}

				if m.Project == "QtWebEngine" && f.Name != "initialize" {
					continue
				}

				f.Static = true
				f.register(m.Project)
				if c, ok := f.Class(); ok {
					if l := len(snsExtraClasses); l > 0 && snsExtraClasses[l-1].Name == c.Name {
						continue
					}
					snsExtraClasses = append(snsExtraClasses, c)
				}
			}
		}
	}

	//mutate classmap
	if utils.QT_GEN_QUICK_EXTRAS() {
		if m.Project == "QtQuickControls2" {
			for _, c := range append(SortedClassesForModule(m.Project, false), snsExtraClasses...) {
				if strings.HasPrefix(c.Name, "QQuick") && !strings.HasSuffix(c.Name, "Private") {
					c.Status = "active"
					c.Access = "public"
				}
			}
		}
	}
	m.remove()

	//mutate classes
	for _, c := range append(SortedClassesForModule(m.Project, false), snsExtraClasses...) {
		c.add()
		c.fix()
		c.remove()
	}

	//register derivations
	for _, c := range m.Namespace.Classes {
		c.derivation()
	}

	return nil
}
