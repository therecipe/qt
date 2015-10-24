package macextras

//#include "qmacpasteboardmime.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMacPasteboardMime struct {
	ptr unsafe.Pointer
}

type QMacPasteboardMimeITF interface {
	QMacPasteboardMimePTR() *QMacPasteboardMime
}

func (p *QMacPasteboardMime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMacPasteboardMime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMacPasteboardMime(ptr QMacPasteboardMimeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacPasteboardMimePTR().Pointer()
	}
	return nil
}

func QMacPasteboardMimeFromPointer(ptr unsafe.Pointer) *QMacPasteboardMime {
	var n = new(QMacPasteboardMime)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMacPasteboardMime) QMacPasteboardMimePTR() *QMacPasteboardMime {
	return ptr
}

func (ptr *QMacPasteboardMime) ConvertorName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_ConvertorName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMacPasteboardMime) Count(mimeData core.QMimeDataITF) int {
	if ptr.Pointer() != nil {
		return int(C.QMacPasteboardMime_Count(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_FlavorFor(C.QtObjectPtr(ptr.Pointer()), C.CString(mime)))
	}
	return ""
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_MimeFor(C.QtObjectPtr(ptr.Pointer()), C.CString(flav)))
	}
	return ""
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMime(C.QtObjectPtr(ptr.Pointer()))
	}
}
