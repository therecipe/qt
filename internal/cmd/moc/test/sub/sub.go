package sub

import "github.com/therecipe/qt/internal/cmd/moc/test/sub/subsub"

type SubTestStruct struct {
	subsub.SubSubTestStruct

	_ func() `constructor:"Init"`

	_ func(string)        `signal:"someSignal"`
	_ func(string) string `slot:"someSlot"`
	_ string              `property:"someProperty"`
}

func (s *SubTestStruct) Init() {
	s.SubSubConstructorProperty++
}
