package parser

import "strings"

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

	//register namespace
	if m.Namespace != nil {
		for _, c := range m.Namespace.Classes {
			c.register(m.Project)
			for _, v := range c.Variables {
				if !c.hasFunctionWithName(v.Name) {
					c.Functions = append(c.Functions, v.toFunction(GETTER))
					if !strings.Contains(v.Output, "const") {
						c.Functions = append(c.Functions, v.toFunction(SETTER))
					}
				}
			}
		}
	}

	//register subnamespace
	if m.Project == "QtCore" || m.Project == "QtMultimedia" {
		if m.Namespace.SubNamespace != nil {
			for _, e := range m.Namespace.SubNamespace.Enums {
				e.register(m.Project)
			}
		}
	}

	//remove obsolete and private
	m.removeClasses()
	for _, c := range ClassMap {
		if c.Module == m.Project {
			c.fix()
			c.removeFunctions()
			c.removeEnums()
		}
	}

	for _, c := range ClassMap {
		if c.Module == m.Project {
			c.fixBases()

			for _, f := range c.Functions {
				f.fix()
				f.fixOverload()

				if f.Virtual == "virtual" {
					f.Virtual = IMPURE
				}
				if f.Meta == COPY_CONSTRUCTOR || f.Meta == MOVE_CONSTRUCTOR {
					f.Meta = CONSTRUCTOR
				}
			}
		}
	}
}

func (m *Module) removeClasses() {
	for _, c := range ClassMap {
		if (c.Status == "obsolete" || c.Status == "compat") || !(c.Access == "public" || c.Access == "protected") || c.Name == "qoutputrange" {
			delete(ClassMap, c.Name)
		}
	}
}

func sailfishModule() *Module {
	//TODO: should be in Namespace.Functions
	return &Module{Project: "QtSailfish", Namespace: &Namespace{Classes: []*Class{{
		Name:   "SailfishApp",
		Access: "public",
		Module: "QtSailfish",
		Functions: []*Function{
			{
				Name:      "application",
				Fullname:  "SailfishApp::application",
				Access:    "public",
				Virtual:   "non",
				Meta:      PLAIN,
				Static:    true,
				Output:    "QGuiApplication*",
				Signature: "()",
				Parameters: []*Parameter{
					{
						Name:  "argc",
						Value: "int &",
					},
					{
						Name:  "argv",
						Value: "char **",
					},
				}},
			{
				Name:      "main",
				Fullname:  "SailfishApp::main",
				Access:    "public",
				Virtual:   "non",
				Meta:      PLAIN,
				Static:    true,
				Output:    "int",
				Signature: "()",
				Parameters: []*Parameter{
					{
						Name:  "argc",
						Value: "int &",
					},
					{
						Name:  "argv",
						Value: "char **",
					},
				}},
			{
				Name:      "createView",
				Fullname:  "SailfishApp::createView",
				Access:    "public",
				Virtual:   "non",
				Meta:      PLAIN,
				Static:    true,
				Output:    "QQuickView*",
				Signature: "()",
			},
			{
				Name:      "pathTo",
				Fullname:  "SailfishApp::pathTo",
				Access:    "public",
				Virtual:   "non",
				Meta:      PLAIN,
				Static:    true,
				Output:    "QUrl",
				Signature: "pathTo(const QString &filename)",
				Parameters: []*Parameter{
					{
						Name:  "filename",
						Value: "QString &",
					},
				}},
		}}}}}
}
