package gui

//#include "gui.h"
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
	for len(n.ObjectName()) < len("QPdfWriter_") {
		n.SetObjectName("QPdfWriter_" + qt.Identifier())
	}
	return n
}

func (ptr *QPdfWriter) QPdfWriter_PTR() *QPdfWriter {
	return ptr
}

func NewQPdfWriter2(device core.QIODevice_ITF) *QPdfWriter {
	defer qt.Recovering("QPdfWriter::QPdfWriter")

	return NewQPdfWriterFromPointer(C.QPdfWriter_NewQPdfWriter2(core.PointerFromQIODevice(device)))
}

func NewQPdfWriter(filename string) *QPdfWriter {
	defer qt.Recovering("QPdfWriter::QPdfWriter")

	return NewQPdfWriterFromPointer(C.QPdfWriter_NewQPdfWriter(C.CString(filename)))
}

func (ptr *QPdfWriter) Creator() string {
	defer qt.Recovering("QPdfWriter::creator")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Creator(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPdfWriter) NewPage() bool {
	defer qt.Recovering("QPdfWriter::newPage")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_NewPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPdfWriter) PaintEngine() *QPaintEngine {
	defer qt.Recovering("QPdfWriter::paintEngine")

	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPdfWriter_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPdfWriter) Resolution() int {
	defer qt.Recovering("QPdfWriter::resolution")

	if ptr.Pointer() != nil {
		return int(C.QPdfWriter_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPdfWriter) SetCreator(creator string) {
	defer qt.Recovering("QPdfWriter::setCreator")

	if ptr.Pointer() != nil {
		C.QPdfWriter_SetCreator(ptr.Pointer(), C.CString(creator))
	}
}

func (ptr *QPdfWriter) SetPageLayout(newPageLayout QPageLayout_ITF) bool {
	defer qt.Recovering("QPdfWriter::setPageLayout")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageLayout(ptr.Pointer(), PointerFromQPageLayout(newPageLayout)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins(margins core.QMarginsF_ITF) bool {
	defer qt.Recovering("QPdfWriter::setPageMargins")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins(ptr.Pointer(), core.PointerFromQMarginsF(margins)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageMargins2(margins core.QMarginsF_ITF, units QPageLayout__Unit) bool {
	defer qt.Recovering("QPdfWriter::setPageMargins")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageMargins2(ptr.Pointer(), core.PointerFromQMarginsF(margins), C.int(units)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	defer qt.Recovering("QPdfWriter::setPageOrientation")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageOrientation(ptr.Pointer(), C.int(orientation)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetPageSize(pageSize QPageSize_ITF) bool {
	defer qt.Recovering("QPdfWriter::setPageSize")

	if ptr.Pointer() != nil {
		return C.QPdfWriter_SetPageSize(ptr.Pointer(), PointerFromQPageSize(pageSize)) != 0
	}
	return false
}

func (ptr *QPdfWriter) SetResolution(resolution int) {
	defer qt.Recovering("QPdfWriter::setResolution")

	if ptr.Pointer() != nil {
		C.QPdfWriter_SetResolution(ptr.Pointer(), C.int(resolution))
	}
}

func (ptr *QPdfWriter) SetTitle(title string) {
	defer qt.Recovering("QPdfWriter::setTitle")

	if ptr.Pointer() != nil {
		C.QPdfWriter_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QPdfWriter) Title() string {
	defer qt.Recovering("QPdfWriter::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPdfWriter_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPdfWriter) DestroyQPdfWriter() {
	defer qt.Recovering("QPdfWriter::~QPdfWriter")

	if ptr.Pointer() != nil {
		C.QPdfWriter_DestroyQPdfWriter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPdfWriter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPdfWriter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPdfWriter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPdfWriter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPdfWriterTimerEvent
func callbackQPdfWriterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPdfWriter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPdfWriterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPdfWriter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPdfWriter::timerEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPdfWriter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPdfWriter::timerEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPdfWriter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPdfWriter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPdfWriter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPdfWriter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPdfWriterChildEvent
func callbackQPdfWriterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPdfWriter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPdfWriterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPdfWriter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPdfWriter::childEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPdfWriter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPdfWriter::childEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPdfWriter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPdfWriter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPdfWriter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPdfWriter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPdfWriterCustomEvent
func callbackQPdfWriterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPdfWriter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPdfWriterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPdfWriter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPdfWriter::customEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPdfWriter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPdfWriter::customEvent")

	if ptr.Pointer() != nil {
		C.QPdfWriter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
