package core

//#include "qunicodetools.h"
import "C"
import (
	"unsafe"
)

type QUnicodeTools struct {
	ptr unsafe.Pointer
}

type QUnicodeToolsITF interface {
	QUnicodeToolsPTR() *QUnicodeTools
}

func (p *QUnicodeTools) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUnicodeTools) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUnicodeTools(ptr QUnicodeToolsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUnicodeToolsPTR().Pointer()
	}
	return nil
}

func QUnicodeToolsFromPointer(ptr unsafe.Pointer) *QUnicodeTools {
	var n = new(QUnicodeTools)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUnicodeTools) QUnicodeToolsPTR() *QUnicodeTools {
	return ptr
}

//QUnicodeTools::CharAttributeOption
type QUnicodeTools__CharAttributeOption int

var (
	QUnicodeTools__GraphemeBreaks       = QUnicodeTools__CharAttributeOption(0x01)
	QUnicodeTools__WordBreaks           = QUnicodeTools__CharAttributeOption(0x02)
	QUnicodeTools__SentenceBreaks       = QUnicodeTools__CharAttributeOption(0x04)
	QUnicodeTools__LineBreaks           = QUnicodeTools__CharAttributeOption(0x08)
	QUnicodeTools__WhiteSpaces          = QUnicodeTools__CharAttributeOption(0x10)
	QUnicodeTools__DefaultOptionsCompat = QUnicodeTools__CharAttributeOption(QUnicodeTools__GraphemeBreaks | QUnicodeTools__LineBreaks | QUnicodeTools__WhiteSpaces)
	QUnicodeTools__DontClearAttributes  = QUnicodeTools__CharAttributeOption(0x1000)
)
