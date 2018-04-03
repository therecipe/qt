package subsubcustom

import "github.com/therecipe/qt/gui"

type SubSubTestStruct struct {
	gui.QWindow

	_ func() `constructor:"Init"`

	_ string `property:"subsubProperty"`

	SubSubConstructorProperty int
}

func (s *SubSubTestStruct) Init() {
	s.SubSubConstructorProperty++
}
