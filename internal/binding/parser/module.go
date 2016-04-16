package parser

type Module struct {
	Namespace *Namespace `xml:"namespace"`
	Project   string     `xml:"project,attr"`
}

type Namespace struct {
	Classes []*Class `xml:"class"`
	//Functions    []*Function   `xml:"function"`
	SubNamespace *SubNamespace `xml:"namespace"`
}

type SubNamespace struct {
	//Classes   []*Class    `xml:"class"`
	//Functions []*Function `xml:"function"`
	Enums []*Enum `xml:"enum"`
}

func (m *Module) Prepare() {

	//Register Namespace
	for _, c := range m.Namespace.Classes {
		c.register(m.Project)
	}

	//Register Subnamespace
	if m.Project == "QtCore" || m.Project == "QtMultimedia" {
		if m.Namespace.SubNamespace != nil {
			for _, e := range m.Namespace.SubNamespace.Enums {
				e.register(m.Project)
			}
		}
	}

	//Remove obsolete and private
	m.removeClasses()
	for _, c := range ClassMap {
		if c.Module == m.Project {
			c.removeFunctions()
			c.removeEnums()
			c.registerAbstact()
			//TODO: register *Class in *Function
		}
	}

	//Fixes for 5.6
	for _, c := range ClassMap {
		if c.Module == m.Project {
			//c.fixBases()

			for _, f := range c.Functions {
				f.fix()
				f.fixOverload()

				if f.Virtual == "virtual" {
					f.Virtual = "impure"
				}
			}
		}
	}
}

func (m *Module) removeClasses() {
	for _, c := range ClassMap {
		if c.Status == "obsolete" || !(c.Access == "public" || c.Access == "protected") || c.Name == "qoutputrange" {
			delete(ClassMap, c.Name)
		}
	}
}
