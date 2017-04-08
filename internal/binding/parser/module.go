package parser

import (
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

type Module struct {
	Namespace *Namespace `xml:"namespace"`
	Project   string     `xml:"project,attr"`
	Pkg       string
}

type Namespace struct {
	Classes []*Class `xml:"class"`
	//Functions     []*Function     `xml:"function"` //TODO: uncomment
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
	utils.Log.WithField("module", strings.TrimPrefix(m.Project, "Qt")).Debug("prepare")

	//register classes from namespace
	for _, c := range m.Namespace.Classes {
		c.register(m)
	}

	//register enums and functions from subnamespaces
	for _, sns := range m.Namespace.SubNamespaces {
		for _, e := range sns.Enums {
			if !(e.Status == "active") || !(e.Access == "public" || e.Access == "protected") ||
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

				if !(f.Status == "active") || !(f.Access == "public" || f.Access == "protected") ||
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
			}
		}
	}

	//mutate classmap
	m.remove()

	//mutate classes
	for _, c := range SortedClassesForModule(m.Project, false) {
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
