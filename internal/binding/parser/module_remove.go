package parser

func (m *Module) remove() {
	m.removeClasses()
}

func (m *Module) removeClasses() {
	for _, c := range CurrentState.ClassMap {
		if c.Module != m.Project {
			continue
		}

		switch {
		case c.Status == "obsolete", c.Status == "compat",
			!(c.Access == "public" || c.Access == "protected"),
			c.Name == "qoutputrange":

			delete(CurrentState.ClassMap, c.Name)
		}
	}
}
