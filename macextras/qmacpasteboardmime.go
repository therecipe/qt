package macextras

//#include "macextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMacPasteboardMime struct {
	ptr unsafe.Pointer
}

type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (p *QMacPasteboardMime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMacPasteboardMime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMacPasteboardMime(ptr QMacPasteboardMime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacPasteboardMime_PTR().Pointer()
	}
	return nil
}

func NewQMacPasteboardMimeFromPointer(ptr unsafe.Pointer) *QMacPasteboardMime {
	var n = new(QMacPasteboardMime)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMacPasteboardMime_") {
		n.SetObjectNameAbs("QMacPasteboardMime_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime {
	return ptr
}

func (ptr *QMacPasteboardMime) ConvertorName() string {
	defer qt.Recovering("QMacPasteboardMime::convertorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_ConvertorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMacPasteboardMime) Count(mimeData core.QMimeData_ITF) int {
	defer qt.Recovering("QMacPasteboardMime::count")

	if ptr.Pointer() != nil {
		return int(C.QMacPasteboardMime_Count(ptr.Pointer(), core.PointerFromQMimeData(mimeData)))
	}
	return 0
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {
	defer qt.Recovering("QMacPasteboardMime::flavorFor")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_FlavorFor(ptr.Pointer(), C.CString(mime)))
	}
	return ""
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {
	defer qt.Recovering("QMacPasteboardMime::mimeFor")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_MimeFor(ptr.Pointer(), C.CString(flav)))
	}
	return ""
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
	defer qt.Recovering("QMacPasteboardMime::~QMacPasteboardMime")

	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMime(ptr.Pointer())
	}
}

func (ptr *QMacPasteboardMime) ObjectNameAbs() string {
	defer qt.Recovering("QMacPasteboardMime::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMacPasteboardMime) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMacPasteboardMime::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
