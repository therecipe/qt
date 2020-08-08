package sub

import (
	"github.com/StarAurryon/qt/internal/cmd/moc/test/sub/subsub"

	"github.com/StarAurryon/qt/internal/cmd/moc/test/sub/b"
	abc "github.com/StarAurryon/qt/internal/cmd/moc/test/sub/c"
	. "github.com/StarAurryon/qt/internal/cmd/moc/test/sub/d"

	"github.com/StarAurryon/qt/internal/cmd/moc/test/sub/bgo"
	abcgo "github.com/StarAurryon/qt/internal/cmd/moc/test/sub/cgo"
	. "github.com/StarAurryon/qt/internal/cmd/moc/test/sub/dgo"

	_ "github.com/StarAurryon/qt/internal/cmd/moc/test/sub/conf"
)

//TODO:  duplicate struct name issue

type SubTestStruct struct {
	subsubcustom.SubSubTestStruct

	_ func() `constructor:"init"`

	_ func(string)        `signal:"someSignal"`
	_ func(string) string `slot:"someSlot"`
	_ string              `property:"someProperty"`

	_ func(StructSameGo)  `signal:"sameSignal"`
	_ func(*StructSameGo) `signal:"sameSignalR"`

	_ func(b.StructSubGoB)  `signal:"subSignalb1"`
	_ func(*b.StructSubGoB) `signal:"subSignalb1R"`

	_ func(b.StructSubMocB)  `signal:"subSignalb2"`
	_ func(*b.StructSubMocB) `signal:"subSignalb2R"`

	_ func(bgo.StructSubGoBGo)  `signal:"subSignalb3"`
	_ func(*bgo.StructSubGoBGo) `signal:"subSignalb3R"`

	_ func(abc.StructSubGoC)  `signal:"subSignalc1"`
	_ func(*abc.StructSubGoC) `signal:"subSignalc1R"`

	_ func(abc.StructSubMocC)  `signal:"subSignalc2"`
	_ func(*abc.StructSubMocC) `signal:"subSignalc2R"`

	_ func(abcgo.StructSubGoCGo)  `signal:"subSignalc3"`
	_ func(*abcgo.StructSubGoCGo) `signal:"subSignalc3R"`

	_ func(StructSubGoD)  `signal:"subSignald1"`
	_ func(*StructSubGoD) `signal:"subSignald1R"`

	_ func(StructSubMocD)  `signal:"subSignald2"`
	_ func(*StructSubMocD) `signal:"subSignald2R"`

	_ func(StructSubGoDGo)  `signal:"subSignald3"`
	_ func(*StructSubGoDGo) `signal:"subSignald3R"`
}

func (s *SubTestStruct) init() {
	s.SubSubConstructorProperty++
}

type StructSameGo struct {
	somestring string
}
