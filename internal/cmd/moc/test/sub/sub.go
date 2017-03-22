package sub

import "github.com/therecipe/qt/internal/cmd/moc/test/sub/subsub"

type SubTestStruct struct {
	subsub.SubSubTestStruct

	_ func(string)        `signal:"someSignal"`
	_ func(string) string `slot:"someSlot"`
	_ string              `property:"someProperty"`
}
