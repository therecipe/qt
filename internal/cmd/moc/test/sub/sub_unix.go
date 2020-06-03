// +build !windows

package sub

import (
	"github.com/therecipe/qt/internal/cmd/moc/test/sub/subsub"

	"a"   //vendor
	"ago" //vendor
)

//TODO: seems like there is a bug? in the Go 1.14 module mode that makes it impossible to pull (and use?) a project that uses vendor dirs on windows ...

type SubTestStructNotUsed struct {
	subsubcustom.SubSubTestStruct

	_ func(a.StructSubGoA)  `signal:"vendorSignal1"`
	_ func(*a.StructSubGoA) `signal:"vendorSignal1R"`

	_ func(a.StructSubMocA)  `signal:"vendorSignal2"`
	_ func(*a.StructSubMocA) `signal:"vendorSignal2R"`

	_ func(ago.StructSubGoAGo)  `signal:"vendorSignal3"`
	_ func(*ago.StructSubGoAGo) `signal:"vendorSignal3R"`
}
