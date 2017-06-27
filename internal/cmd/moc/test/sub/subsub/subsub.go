package subsub

import "github.com/therecipe/qt/core"

type SubSubTestStruct struct {
	core.QObject //TODO: gui.QWindow

	_ func() `constructor:"Init"`

	_ string `property:"subsubProperty"`

	SubSubConstructorProperty int
}

func (s *SubSubTestStruct) Init() {
	s.SubSubConstructorProperty++
}
