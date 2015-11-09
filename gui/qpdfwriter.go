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

type QPdfWriter_ITF interface {
	core.QObject_ITF
	QPagedPaintDevice_ITF
	QPdfWriter_PTR() *QPdfWriter
}

func (p *QPdfWriter) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QPdfWriter) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QPagedPaintDevice_PTR().SetPointer(ptr)
}

func PointerFromQPdfWriter(ptr QPdfWriter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPdfWriter_PTR().Pointer()
	}
	return nil
}

func NewQPdfWriterFromPointer(ptr unsafe.Pointer) *QPdfWriter {
	var n = new(QPdfWriter)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPdfWriter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPdfWriter) QPdfWriter_PTR() *QPdfWriter {
	return ptr
}

func NewQPdfWriter2(device core.QIODevice_ITF) *QPdfWriter {
	return NewQPdfWriterFromPointer(C.QPdfWriter_NewQPdfWriter2(core.PointerFromQIODevice(device)))
}

func NewQPdfWriter(filename string) *QPdfWriter {
	return NewQPdfWriterFromPointer(C.QPdfWriter_NewQPdfWriter(C.CString(filename)))
}

func (ptr *QPdfWriter) Creator() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Creator(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPdfWriter) NewPage() bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_NewPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPdfWriter) Resolution() int {
	if ptr.Pointer() != nil {
		return int(C.QPdfWriter_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPdfWriter) SetCreator(creator string) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetCreator(ptr.Pointer(), C.CString(creator))
	}
}

func (ptr *QPdfWriter) SetPageLayout(newPageLayout QPageLayout_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageLayout(ptr.Pointer(), PointerFromQPageLayout(newPageLayout)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins(margins core.QMarginsF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins(ptr.Pointer(), core.PointerFromQMarginsF(margins)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins2(margins core.QMarginsF_ITF, units QPageLayout__Unit) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins2(ptr.Pointer(), core.PointerFromQMarginsF(margins), C.int(units)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageOrientation(ptr.Pointer(), C.int(orientation)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageSize(pageSize QPageSize_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageSize(ptr.Pointer(), PointerFromQPageSize(pageSize)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetResolution(resolution int) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetResolution(ptr.Pointer(), C.int(resolution))
	}
}

func (ptr *QPdfWriter) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QPdfWriter_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QPdfWriter) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPdfWriter) DestroyQPdfWriter() {
	if ptr.Pointer() != nil {
		C.QPdfWriter_DestroyQPdfWriter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
