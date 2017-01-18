package parser

type Module struct {
	Namespace *Namespace `xml:"namespace"`
	Project   string     `xml:"project,attr"`
}

type Namespace struct {
	Classes []*Class `xml:"class"`
	//Functions    []*Function   `xml:"function"` //TODO: uncomment
	SubNamespace *SubNamespace `xml:"namespace"`
}

type SubNamespace struct {
	//Classes   []*Class    `xml:"class"`    //TODO: uncomment
	//Functions []*Function `xml:"function"` //TODO: uncomment
	Enums []*Enum `xml:"enum"`
}

func (m *Module) Prepare() error {

	//register classes from namespace
	for _, c := range m.Namespace.Classes {
		c.register(m.Project)
	}

	//register enums from subnamespace
	if m.Namespace.SubNamespace != nil {
		for _, e := range m.Namespace.SubNamespace.Enums {
			e.register(m.Project)
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

	return nil
}
