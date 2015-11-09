package core

//#include "qunicodetools.h"
import "C"
import (
	"unsafe"
)

type QUnicodeTools struct {
	ptr unsafe.Pointer
}

type QUnicodeTools_ITF interface {
	QUnicodeTools_PTR() *QUnicodeTools
}

func (p *QUnicodeTools) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUnicodeTools) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUnicodeTools(ptr QUnicodeTools_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUnicodeTools_PTR().Pointer()
	}
	return nil
}

func NewQUnicodeToolsFromPointer(ptr unsafe.Pointer) *QUnicodeTools {
	var n = new(QUnicodeTools)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUnicodeTools) QUnicodeTools_PTR() *QUnicodeTools {
	return ptr
}

//QUnicodeTools::CharAttributeOption
type QUnicodeTools__CharAttributeOption int64

const (
	QUnicodeTools__GraphemeBreaks       = QUnicodeTools__CharAttributeOption(0x01)
	QUnicodeTools__WordBreaks           = QUnicodeTools__CharAttributeOption(0x02)
	QUnicodeTools__SentenceBreaks       = QUnicodeTools__CharAttributeOption(0x04)
	QUnicodeTools__LineBreaks           = QUnicodeTools__CharAttributeOption(0x08)
	QUnicodeTools__WhiteSpaces          = QUnicodeTools__CharAttributeOption(0x10)
	QUnicodeTools__DefaultOptionsCompat = QUnicodeTools__CharAttributeOption(QUnicodeTools__GraphemeBreaks | QUnicodeTools__LineBreaks | QUnicodeTools__WhiteSpaces)
	QUnicodeTools__DontClearAttributes  = QUnicodeTools__CharAttributeOption(0x1000)
)
