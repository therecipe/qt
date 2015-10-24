package gui

//#include "qpdfwriter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPdfWriter struct {
	core.QObject
	QPagedPaintDevice
}

type QPdfWriterITF interface {
	core.QObjectITF
	QPagedPaintDeviceITF
	QPdfWriterPTR() *QPdfWriter
}

func (p *QPdfWriter) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QPdfWriter) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QPagedPaintDevicePTR().SetPointer(ptr)
}

func PointerFromQPdfWriter(ptr QPdfWriterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPdfWriterPTR().Pointer()
	}
	return nil
}

func QPdfWriterFromPointer(ptr unsafe.Pointer) *QPdfWriter {
	var n = new(QPdfWriter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPdfWriter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPdfWriter) QPdfWriterPTR() *QPdfWriter {
	return ptr
}

func NewQPdfWriter2(device core.QIODeviceITF) *QPdfWriter {
	return QPdfWriterFromPointer(unsafe.Pointer(C.QPdfWriter_NewQPdfWriter2(C.QtObjectPtr(core.PointerFromQIODevice(device)))))
}

func NewQPdfWriter(filename string) *QPdfWriter {
	return QPdfWriterFromPointer(unsafe.Pointer(C.QPdfWriter_NewQPdfWriter(C.CString(filename))))
}

func (ptr *QPdfWriter) Creator() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Creator(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPdfWriter) NewPage() bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_NewPage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPdfWriter) Resolution() int {
	if ptr.Pointer() != nil {
		return int(C.QPdfWriter_Resolution(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPdfWriter) SetCreator(creator string) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetCreator(C.QtObjectPtr(ptr.Pointer()), C.CString(creator))
	}
}

func (ptr *QPdfWriter) SetPageLayout(newPageLayout QPageLayoutITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageLayout(newPageLayout))) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins(margins core.QMarginsFITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMarginsF(margins))) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins2(margins core.QMarginsFITF, units QPageLayout__Unit) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMarginsF(margins)), C.int(units)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageSize(pageSize QPageSizeITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageSize(pageSize))) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetResolution(resolution int) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.int(resolution))
	}
}

func (ptr *QPdfWriter) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QPdfWriter) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPdfWriter) DestroyQPdfWriter() {
	if ptr.Pointer() != nil {
		C.QPdfWriter_DestroyQPdfWriter(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
