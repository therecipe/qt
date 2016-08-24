// +build !minimal

package printsupport

//#include <stdint.h>
//#include <stdlib.h>
//#include "printsupport.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

//QAbstractPrintDialog::PrintDialogOption
type QAbstractPrintDialog__PrintDialogOption int64

const (
	QAbstractPrintDialog__None               = QAbstractPrintDialog__PrintDialogOption(0x0000)
	QAbstractPrintDialog__PrintToFile        = QAbstractPrintDialog__PrintDialogOption(0x0001)
	QAbstractPrintDialog__PrintSelection     = QAbstractPrintDialog__PrintDialogOption(0x0002)
	QAbstractPrintDialog__PrintPageRange     = QAbstractPrintDialog__PrintDialogOption(0x0004)
	QAbstractPrintDialog__PrintShowPageSize  = QAbstractPrintDialog__PrintDialogOption(0x0008)
	QAbstractPrintDialog__PrintCollateCopies = QAbstractPrintDialog__PrintDialogOption(0x0010)
	QAbstractPrintDialog__DontUseSheet       = QAbstractPrintDialog__PrintDialogOption(0x0020)
	QAbstractPrintDialog__PrintCurrentPage   = QAbstractPrintDialog__PrintDialogOption(0x0040)
)

//QAbstractPrintDialog::PrintRange
type QAbstractPrintDialog__PrintRange int64

const (
	QAbstractPrintDialog__AllPages    = QAbstractPrintDialog__PrintRange(0)
	QAbstractPrintDialog__Selection   = QAbstractPrintDialog__PrintRange(1)
	QAbstractPrintDialog__PageRange   = QAbstractPrintDialog__PrintRange(2)
	QAbstractPrintDialog__CurrentPage = QAbstractPrintDialog__PrintRange(3)
)

type QAbstractPrintDialog struct {
	widgets.QDialog
}

type QAbstractPrintDialog_ITF interface {
	widgets.QDialog_ITF
	QAbstractPrintDialog_PTR() *QAbstractPrintDialog
}

func (p *QAbstractPrintDialog) QAbstractPrintDialog_PTR() *QAbstractPrintDialog {
	return p
}

func (p *QAbstractPrintDialog) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDialog_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractPrintDialog) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDialog_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractPrintDialog(ptr QAbstractPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func NewQAbstractPrintDialogFromPointer(ptr unsafe.Pointer) *QAbstractPrintDialog {
	var n = new(QAbstractPrintDialog)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractPrintDialog) DestroyQAbstractPrintDialog() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQAbstractPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QAbstractPrintDialog {
	defer qt.Recovering("QAbstractPrintDialog::QAbstractPrintDialog")

	var tmpValue = NewQAbstractPrintDialogFromPointer(C.QAbstractPrintDialog_NewQAbstractPrintDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQAbstractPrintDialog_Exec
func callbackQAbstractPrintDialog_Exec(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractPrintDialog::exec")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractPrintDialog) ConnectExec(f func() int) {
	defer qt.Recovering("connect QAbstractPrintDialog::exec")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::exec", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectExec() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::exec")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::exec")
	}
}

func (ptr *QAbstractPrintDialog) Exec() int {
	defer qt.Recovering("QAbstractPrintDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) FromPage() int {
	defer qt.Recovering("QAbstractPrintDialog::fromPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_FromPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) MaxPage() int {
	defer qt.Recovering("QAbstractPrintDialog::maxPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_MaxPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) MinPage() int {
	defer qt.Recovering("QAbstractPrintDialog::minPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_MinPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) PrintRange() QAbstractPrintDialog__PrintRange {
	defer qt.Recovering("QAbstractPrintDialog::printRange")

	if ptr.Pointer() != nil {
		return QAbstractPrintDialog__PrintRange(C.QAbstractPrintDialog_PrintRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) Printer() *QPrinter {
	defer qt.Recovering("QAbstractPrintDialog::printer")

	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QAbstractPrintDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractPrintDialog) SetFromTo(from int, to int) {
	defer qt.Recovering("QAbstractPrintDialog::setFromTo")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetFromTo(ptr.Pointer(), C.int(int32(from)), C.int(int32(to)))
	}
}

func (ptr *QAbstractPrintDialog) SetMinMax(min int, max int) {
	defer qt.Recovering("QAbstractPrintDialog::setMinMax")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetMinMax(ptr.Pointer(), C.int(int32(min)), C.int(int32(max)))
	}
}

func (ptr *QAbstractPrintDialog) SetPrintRange(ran QAbstractPrintDialog__PrintRange) {
	defer qt.Recovering("QAbstractPrintDialog::setPrintRange")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetPrintRange(ptr.Pointer(), C.longlong(ran))
	}
}

func (ptr *QAbstractPrintDialog) ToPage() int {
	defer qt.Recovering("QAbstractPrintDialog::toPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_ToPage(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractPrintDialog_Accept
func callbackQAbstractPrintDialog_Accept(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::accept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::accept", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::accept")
	}
}

func (ptr *QAbstractPrintDialog) Accept() {
	defer qt.Recovering("QAbstractPrintDialog::accept")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) AcceptDefault() {
	defer qt.Recovering("QAbstractPrintDialog::accept")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_CloseEvent
func callbackQAbstractPrintDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::closeEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::closeEvent")
	}
}

func (ptr *QAbstractPrintDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QAbstractPrintDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQAbstractPrintDialog_ContextMenuEvent
func callbackQAbstractPrintDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::contextMenuEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::contextMenuEvent")
	}
}

func (ptr *QAbstractPrintDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QAbstractPrintDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQAbstractPrintDialog_Done
func callbackQAbstractPrintDialog_Done(ptr unsafe.Pointer, r C.int) {
	defer qt.Recovering("callback QAbstractPrintDialog::done")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::done"); signal != nil {
		signal.(func(int))(int(int32(r)))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DoneDefault(int(int32(r)))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDone(f func(r int)) {
	defer qt.Recovering("connect QAbstractPrintDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::done", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::done")
	}
}

func (ptr *QAbstractPrintDialog) Done(r int) {
	defer qt.Recovering("QAbstractPrintDialog::done")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Done(ptr.Pointer(), C.int(int32(r)))
	}
}

func (ptr *QAbstractPrintDialog) DoneDefault(r int) {
	defer qt.Recovering("QAbstractPrintDialog::done")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DoneDefault(ptr.Pointer(), C.int(int32(r)))
	}
}

//export callbackQAbstractPrintDialog_KeyPressEvent
func callbackQAbstractPrintDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::keyPressEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::keyPressEvent")
	}
}

func (ptr *QAbstractPrintDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QAbstractPrintDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQAbstractPrintDialog_MinimumSizeHint
func callbackQAbstractPrintDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractPrintDialog::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQAbstractPrintDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QAbstractPrintDialog) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QAbstractPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::minimumSizeHint", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::minimumSizeHint")
	}
}

func (ptr *QAbstractPrintDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QAbstractPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QAbstractPrintDialog_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QAbstractPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QAbstractPrintDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_Open
func callbackQAbstractPrintDialog_Open(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::open"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).OpenDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::open", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::open")
	}
}

func (ptr *QAbstractPrintDialog) Open() {
	defer qt.Recovering("QAbstractPrintDialog::open")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Open(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) OpenDefault() {
	defer qt.Recovering("QAbstractPrintDialog::open")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_OpenDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Reject
func callbackQAbstractPrintDialog_Reject(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::reject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::reject"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::reject", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::reject")
	}
}

func (ptr *QAbstractPrintDialog) Reject() {
	defer qt.Recovering("QAbstractPrintDialog::reject")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) RejectDefault() {
	defer qt.Recovering("QAbstractPrintDialog::reject")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ResizeEvent
func callbackQAbstractPrintDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QAbstractPrintDialog) ConnectResizeEvent(f func(vqr *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::resizeEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::resizeEvent")
	}
}

func (ptr *QAbstractPrintDialog) ResizeEvent(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

func (ptr *QAbstractPrintDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQAbstractPrintDialog_SetVisible
func callbackQAbstractPrintDialog_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setVisible", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setVisible")
	}
}

func (ptr *QAbstractPrintDialog) SetVisible(visible bool) {
	defer qt.Recovering("QAbstractPrintDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractPrintDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QAbstractPrintDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractPrintDialog_ShowEvent
func callbackQAbstractPrintDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showEvent")
	}
}

func (ptr *QAbstractPrintDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQAbstractPrintDialog_SizeHint
func callbackQAbstractPrintDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractPrintDialog::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQAbstractPrintDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QAbstractPrintDialog) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QAbstractPrintDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::sizeHint", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::sizeHint")
	}
}

func (ptr *QAbstractPrintDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QAbstractPrintDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QAbstractPrintDialog_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QAbstractPrintDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QAbstractPrintDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_ActionEvent
func callbackQAbstractPrintDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::actionEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::actionEvent")
	}
}

func (ptr *QAbstractPrintDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DragEnterEvent
func callbackQAbstractPrintDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragEnterEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragEnterEvent")
	}
}

func (ptr *QAbstractPrintDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DragLeaveEvent
func callbackQAbstractPrintDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragLeaveEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragLeaveEvent")
	}
}

func (ptr *QAbstractPrintDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DragMoveEvent
func callbackQAbstractPrintDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragMoveEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dragMoveEvent")
	}
}

func (ptr *QAbstractPrintDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DropEvent
func callbackQAbstractPrintDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dropEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::dropEvent")
	}
}

func (ptr *QAbstractPrintDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQAbstractPrintDialog_EnterEvent
func callbackQAbstractPrintDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::enterEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::enterEvent")
	}
}

func (ptr *QAbstractPrintDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_FocusInEvent
func callbackQAbstractPrintDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusInEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusInEvent")
	}
}

func (ptr *QAbstractPrintDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQAbstractPrintDialog_FocusOutEvent
func callbackQAbstractPrintDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusOutEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusOutEvent")
	}
}

func (ptr *QAbstractPrintDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQAbstractPrintDialog_HideEvent
func callbackQAbstractPrintDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hideEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hideEvent")
	}
}

func (ptr *QAbstractPrintDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQAbstractPrintDialog_LeaveEvent
func callbackQAbstractPrintDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::leaveEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::leaveEvent")
	}
}

func (ptr *QAbstractPrintDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MoveEvent
func callbackQAbstractPrintDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::moveEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::moveEvent")
	}
}

func (ptr *QAbstractPrintDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_PaintEvent
func callbackQAbstractPrintDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::paintEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::paintEvent")
	}
}

func (ptr *QAbstractPrintDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQAbstractPrintDialog_SetEnabled
func callbackQAbstractPrintDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setEnabled", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setEnabled")
	}
}

func (ptr *QAbstractPrintDialog) SetEnabled(vbo bool) {
	defer qt.Recovering("QAbstractPrintDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAbstractPrintDialog) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QAbstractPrintDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAbstractPrintDialog_SetStyleSheet
func callbackQAbstractPrintDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setStyleSheet", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setStyleSheet")
	}
}

func (ptr *QAbstractPrintDialog) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QAbstractPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QAbstractPrintDialog_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QAbstractPrintDialog) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QAbstractPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QAbstractPrintDialog_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQAbstractPrintDialog_SetWindowModified
func callbackQAbstractPrintDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setWindowModified", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setWindowModified")
	}
}

func (ptr *QAbstractPrintDialog) SetWindowModified(vbo bool) {
	defer qt.Recovering("QAbstractPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAbstractPrintDialog) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QAbstractPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAbstractPrintDialog_SetWindowTitle
func callbackQAbstractPrintDialog_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setWindowTitle", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setWindowTitle")
	}
}

func (ptr *QAbstractPrintDialog) SetWindowTitle(vqs string) {
	defer qt.Recovering("QAbstractPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QAbstractPrintDialog_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QAbstractPrintDialog) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QAbstractPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QAbstractPrintDialog_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQAbstractPrintDialog_ChangeEvent
func callbackQAbstractPrintDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::changeEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::changeEvent")
	}
}

func (ptr *QAbstractPrintDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Close
func callbackQAbstractPrintDialog_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractPrintDialog::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QAbstractPrintDialog) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QAbstractPrintDialog::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::close", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectClose() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::close")
	}
}

func (ptr *QAbstractPrintDialog) Close() bool {
	defer qt.Recovering("QAbstractPrintDialog::close")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractPrintDialog) CloseDefault() bool {
	defer qt.Recovering("QAbstractPrintDialog::close")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_FocusNextPrevChild
func callbackQAbstractPrintDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QAbstractPrintDialog::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QAbstractPrintDialog) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QAbstractPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusNextPrevChild", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::focusNextPrevChild")
	}
}

func (ptr *QAbstractPrintDialog) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QAbstractPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QAbstractPrintDialog) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QAbstractPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_HasHeightForWidth
func callbackQAbstractPrintDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractPrintDialog::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QAbstractPrintDialog) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QAbstractPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hasHeightForWidth", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hasHeightForWidth")
	}
}

func (ptr *QAbstractPrintDialog) HasHeightForWidth() bool {
	defer qt.Recovering("QAbstractPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractPrintDialog) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QAbstractPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QAbstractPrintDialog_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_HeightForWidth
func callbackQAbstractPrintDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QAbstractPrintDialog::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQAbstractPrintDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QAbstractPrintDialog) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QAbstractPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::heightForWidth", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::heightForWidth")
	}
}

func (ptr *QAbstractPrintDialog) HeightForWidth(w int) int {
	defer qt.Recovering("QAbstractPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QAbstractPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQAbstractPrintDialog_Hide
func callbackQAbstractPrintDialog_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectHide(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hide", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectHide() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::hide")
	}
}

func (ptr *QAbstractPrintDialog) Hide() {
	defer qt.Recovering("QAbstractPrintDialog::hide")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Hide(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) HideDefault() {
	defer qt.Recovering("QAbstractPrintDialog::hide")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_InputMethodEvent
func callbackQAbstractPrintDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::inputMethodEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::inputMethodEvent")
	}
}

func (ptr *QAbstractPrintDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQAbstractPrintDialog_InputMethodQuery
func callbackQAbstractPrintDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractPrintDialog::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQAbstractPrintDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QAbstractPrintDialog) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QAbstractPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::inputMethodQuery", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::inputMethodQuery")
	}
}

func (ptr *QAbstractPrintDialog) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QAbstractPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractPrintDialog_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QAbstractPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractPrintDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_KeyReleaseEvent
func callbackQAbstractPrintDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::keyReleaseEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::keyReleaseEvent")
	}
}

func (ptr *QAbstractPrintDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Lower
func callbackQAbstractPrintDialog_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectLower(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::lower", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectLower() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::lower")
	}
}

func (ptr *QAbstractPrintDialog) Lower() {
	defer qt.Recovering("QAbstractPrintDialog::lower")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Lower(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) LowerDefault() {
	defer qt.Recovering("QAbstractPrintDialog::lower")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_MouseDoubleClickEvent
func callbackQAbstractPrintDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseDoubleClickEvent")
	}
}

func (ptr *QAbstractPrintDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MouseMoveEvent
func callbackQAbstractPrintDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseMoveEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseMoveEvent")
	}
}

func (ptr *QAbstractPrintDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MousePressEvent
func callbackQAbstractPrintDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mousePressEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mousePressEvent")
	}
}

func (ptr *QAbstractPrintDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MouseReleaseEvent
func callbackQAbstractPrintDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::mouseReleaseEvent")
	}
}

func (ptr *QAbstractPrintDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_NativeEvent
func callbackQAbstractPrintDialog_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QAbstractPrintDialog::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QAbstractPrintDialog) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QAbstractPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::nativeEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::nativeEvent")
	}
}

func (ptr *QAbstractPrintDialog) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QAbstractPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QAbstractPrintDialog_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QAbstractPrintDialog) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QAbstractPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QAbstractPrintDialog_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_Raise
func callbackQAbstractPrintDialog_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectRaise(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::raise", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectRaise() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::raise")
	}
}

func (ptr *QAbstractPrintDialog) Raise() {
	defer qt.Recovering("QAbstractPrintDialog::raise")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Raise(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) RaiseDefault() {
	defer qt.Recovering("QAbstractPrintDialog::raise")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Repaint
func callbackQAbstractPrintDialog_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::repaint", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectRepaint() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::repaint")
	}
}

func (ptr *QAbstractPrintDialog) Repaint() {
	defer qt.Recovering("QAbstractPrintDialog::repaint")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Repaint(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) RepaintDefault() {
	defer qt.Recovering("QAbstractPrintDialog::repaint")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_SetDisabled
func callbackQAbstractPrintDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setDisabled", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setDisabled")
	}
}

func (ptr *QAbstractPrintDialog) SetDisabled(disable bool) {
	defer qt.Recovering("QAbstractPrintDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QAbstractPrintDialog) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QAbstractPrintDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQAbstractPrintDialog_SetFocus2
func callbackQAbstractPrintDialog_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setFocus2", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setFocus2")
	}
}

func (ptr *QAbstractPrintDialog) SetFocus2() {
	defer qt.Recovering("QAbstractPrintDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) SetFocus2Default() {
	defer qt.Recovering("QAbstractPrintDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_SetHidden
func callbackQAbstractPrintDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QAbstractPrintDialog::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QAbstractPrintDialog) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QAbstractPrintDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setHidden", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::setHidden")
	}
}

func (ptr *QAbstractPrintDialog) SetHidden(hidden bool) {
	defer qt.Recovering("QAbstractPrintDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QAbstractPrintDialog) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QAbstractPrintDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQAbstractPrintDialog_Show
func callbackQAbstractPrintDialog_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::show"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectShow(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::show", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShow() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::show")
	}
}

func (ptr *QAbstractPrintDialog) Show() {
	defer qt.Recovering("QAbstractPrintDialog::show")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Show(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) ShowDefault() {
	defer qt.Recovering("QAbstractPrintDialog::show")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowFullScreen
func callbackQAbstractPrintDialog_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showFullScreen", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showFullScreen")
	}
}

func (ptr *QAbstractPrintDialog) ShowFullScreen() {
	defer qt.Recovering("QAbstractPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) ShowFullScreenDefault() {
	defer qt.Recovering("QAbstractPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowMaximized
func callbackQAbstractPrintDialog_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showMaximized", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showMaximized")
	}
}

func (ptr *QAbstractPrintDialog) ShowMaximized() {
	defer qt.Recovering("QAbstractPrintDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) ShowMaximizedDefault() {
	defer qt.Recovering("QAbstractPrintDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowMinimized
func callbackQAbstractPrintDialog_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showMinimized", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showMinimized")
	}
}

func (ptr *QAbstractPrintDialog) ShowMinimized() {
	defer qt.Recovering("QAbstractPrintDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) ShowMinimizedDefault() {
	defer qt.Recovering("QAbstractPrintDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowNormal
func callbackQAbstractPrintDialog_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showNormal", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::showNormal")
	}
}

func (ptr *QAbstractPrintDialog) ShowNormal() {
	defer qt.Recovering("QAbstractPrintDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) ShowNormalDefault() {
	defer qt.Recovering("QAbstractPrintDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_TabletEvent
func callbackQAbstractPrintDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::tabletEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::tabletEvent")
	}
}

func (ptr *QAbstractPrintDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Update
func callbackQAbstractPrintDialog_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::update"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::update", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectUpdate() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::update")
	}
}

func (ptr *QAbstractPrintDialog) Update() {
	defer qt.Recovering("QAbstractPrintDialog::update")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_Update(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) UpdateDefault() {
	defer qt.Recovering("QAbstractPrintDialog::update")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_UpdateMicroFocus
func callbackQAbstractPrintDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::updateMicroFocus", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::updateMicroFocus")
	}
}

func (ptr *QAbstractPrintDialog) UpdateMicroFocus() {
	defer qt.Recovering("QAbstractPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QAbstractPrintDialog) UpdateMicroFocusDefault() {
	defer qt.Recovering("QAbstractPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_WheelEvent
func callbackQAbstractPrintDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::wheelEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::wheelEvent")
	}
}

func (ptr *QAbstractPrintDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQAbstractPrintDialog_TimerEvent
func callbackQAbstractPrintDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::timerEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::timerEvent")
	}
}

func (ptr *QAbstractPrintDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractPrintDialog_ChildEvent
func callbackQAbstractPrintDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::childEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::childEvent")
	}
}

func (ptr *QAbstractPrintDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractPrintDialog_ConnectNotify
func callbackQAbstractPrintDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractPrintDialog) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractPrintDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::connectNotify", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::connectNotify")
	}
}

func (ptr *QAbstractPrintDialog) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractPrintDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractPrintDialog_CustomEvent
func callbackQAbstractPrintDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractPrintDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::customEvent", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::customEvent")
	}
}

func (ptr *QAbstractPrintDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractPrintDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DeleteLater
func callbackQAbstractPrintDialog_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractPrintDialog) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QAbstractPrintDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::deleteLater", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::deleteLater")
	}
}

func (ptr *QAbstractPrintDialog) DeleteLater() {
	defer qt.Recovering("QAbstractPrintDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractPrintDialog) DeleteLaterDefault() {
	defer qt.Recovering("QAbstractPrintDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractPrintDialog_DisconnectNotify
func callbackQAbstractPrintDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractPrintDialog::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractPrintDialog) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::disconnectNotify", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::disconnectNotify")
	}
}

func (ptr *QAbstractPrintDialog) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractPrintDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractPrintDialog_MetaObject
func callbackQAbstractPrintDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractPrintDialog::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractPrintDialog::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractPrintDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractPrintDialog) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QAbstractPrintDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::metaObject", f)
	}
}

func (ptr *QAbstractPrintDialog) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QAbstractPrintDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractPrintDialog::metaObject")
	}
}

func (ptr *QAbstractPrintDialog) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QAbstractPrintDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractPrintDialog_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractPrintDialog) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QAbstractPrintDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractPrintDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QPageSetupDialog struct {
	widgets.QDialog
}

type QPageSetupDialog_ITF interface {
	widgets.QDialog_ITF
	QPageSetupDialog_PTR() *QPageSetupDialog
}

func (p *QPageSetupDialog) QPageSetupDialog_PTR() *QPageSetupDialog {
	return p
}

func (p *QPageSetupDialog) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDialog_PTR().Pointer()
	}
	return nil
}

func (p *QPageSetupDialog) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDialog_PTR().SetPointer(ptr)
	}
}

func PointerFromQPageSetupDialog(ptr QPageSetupDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageSetupDialog_PTR().Pointer()
	}
	return nil
}

func NewQPageSetupDialogFromPointer(ptr unsafe.Pointer) *QPageSetupDialog {
	var n = new(QPageSetupDialog)
	n.SetPointer(ptr)
	return n
}
func (ptr *QPageSetupDialog) Printer() *QPrinter {
	defer qt.Recovering("QPageSetupDialog::printer")

	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QPageSetupDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPageSetupDialog) DestroyQPageSetupDialog() {
	defer qt.Recovering("QPageSetupDialog::~QPageSetupDialog")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DestroyQPageSetupDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPageSetupDialog) Done(result int) {
	defer qt.Recovering("QPageSetupDialog::done")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPageSetupDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QPageSetupDialog::open")

	if ptr.Pointer() != nil {
		var memberC = C.CString(member)
		defer C.free(unsafe.Pointer(memberC))
		C.QPageSetupDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

func NewQPageSetupDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPageSetupDialog {
	defer qt.Recovering("QPageSetupDialog::QPageSetupDialog")

	var tmpValue = NewQPageSetupDialogFromPointer(C.QPageSetupDialog_NewQPageSetupDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func NewQPageSetupDialog2(parent widgets.QWidget_ITF) *QPageSetupDialog {
	defer qt.Recovering("QPageSetupDialog::QPageSetupDialog")

	var tmpValue = NewQPageSetupDialogFromPointer(C.QPageSetupDialog_NewQPageSetupDialog2(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QPageSetupDialog) Exec() int {
	defer qt.Recovering("QPageSetupDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageSetupDialog) SetVisible(visible bool) {
	defer qt.Recovering("QPageSetupDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQPageSetupDialog_Accept
func callbackQPageSetupDialog_Accept(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::accept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::accept", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QPageSetupDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::accept")
	}
}

func (ptr *QPageSetupDialog) Accept() {
	defer qt.Recovering("QPageSetupDialog::accept")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) AcceptDefault() {
	defer qt.Recovering("QPageSetupDialog::accept")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_CloseEvent
func callbackQPageSetupDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::closeEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::closeEvent")
	}
}

func (ptr *QPageSetupDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QPageSetupDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQPageSetupDialog_ContextMenuEvent
func callbackQPageSetupDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::contextMenuEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::contextMenuEvent")
	}
}

func (ptr *QPageSetupDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QPageSetupDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQPageSetupDialog_KeyPressEvent
func callbackQPageSetupDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::keyPressEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::keyPressEvent")
	}
}

func (ptr *QPageSetupDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QPageSetupDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQPageSetupDialog_MinimumSizeHint
func callbackQPageSetupDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPageSetupDialog::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPageSetupDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPageSetupDialog) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPageSetupDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::minimumSizeHint", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QPageSetupDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::minimumSizeHint")
	}
}

func (ptr *QPageSetupDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPageSetupDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPageSetupDialog_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QPageSetupDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPageSetupDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_Reject
func callbackQPageSetupDialog_Reject(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::reject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::reject"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::reject", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QPageSetupDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::reject")
	}
}

func (ptr *QPageSetupDialog) Reject() {
	defer qt.Recovering("QPageSetupDialog::reject")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) RejectDefault() {
	defer qt.Recovering("QPageSetupDialog::reject")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ResizeEvent
func callbackQPageSetupDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QPageSetupDialog) ConnectResizeEvent(f func(vqr *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::resizeEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::resizeEvent")
	}
}

func (ptr *QPageSetupDialog) ResizeEvent(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

func (ptr *QPageSetupDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQPageSetupDialog_ShowEvent
func callbackQPageSetupDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showEvent")
	}
}

func (ptr *QPageSetupDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QPageSetupDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPageSetupDialog_SizeHint
func callbackQPageSetupDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPageSetupDialog::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPageSetupDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPageSetupDialog) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPageSetupDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::sizeHint", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QPageSetupDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::sizeHint")
	}
}

func (ptr *QPageSetupDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QPageSetupDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPageSetupDialog_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QPageSetupDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPageSetupDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_ActionEvent
func callbackQPageSetupDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::actionEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::actionEvent")
	}
}

func (ptr *QPageSetupDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPageSetupDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPageSetupDialog_DragEnterEvent
func callbackQPageSetupDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragEnterEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragEnterEvent")
	}
}

func (ptr *QPageSetupDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QPageSetupDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPageSetupDialog_DragLeaveEvent
func callbackQPageSetupDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragLeaveEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragLeaveEvent")
	}
}

func (ptr *QPageSetupDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QPageSetupDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPageSetupDialog_DragMoveEvent
func callbackQPageSetupDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragMoveEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dragMoveEvent")
	}
}

func (ptr *QPageSetupDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QPageSetupDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPageSetupDialog_DropEvent
func callbackQPageSetupDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dropEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::dropEvent")
	}
}

func (ptr *QPageSetupDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QPageSetupDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPageSetupDialog_EnterEvent
func callbackQPageSetupDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::enterEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::enterEvent")
	}
}

func (ptr *QPageSetupDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPageSetupDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_FocusInEvent
func callbackQPageSetupDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusInEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusInEvent")
	}
}

func (ptr *QPageSetupDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPageSetupDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPageSetupDialog_FocusOutEvent
func callbackQPageSetupDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusOutEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusOutEvent")
	}
}

func (ptr *QPageSetupDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPageSetupDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPageSetupDialog_HideEvent
func callbackQPageSetupDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hideEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hideEvent")
	}
}

func (ptr *QPageSetupDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPageSetupDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPageSetupDialog_LeaveEvent
func callbackQPageSetupDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::leaveEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::leaveEvent")
	}
}

func (ptr *QPageSetupDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPageSetupDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_MoveEvent
func callbackQPageSetupDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::moveEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::moveEvent")
	}
}

func (ptr *QPageSetupDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPageSetupDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPageSetupDialog_PaintEvent
func callbackQPageSetupDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::paintEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::paintEvent")
	}
}

func (ptr *QPageSetupDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QPageSetupDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPageSetupDialog_SetEnabled
func callbackQPageSetupDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPageSetupDialog) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QPageSetupDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setEnabled", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QPageSetupDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setEnabled")
	}
}

func (ptr *QPageSetupDialog) SetEnabled(vbo bool) {
	defer qt.Recovering("QPageSetupDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPageSetupDialog) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QPageSetupDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPageSetupDialog_SetStyleSheet
func callbackQPageSetupDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QPageSetupDialog) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QPageSetupDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setStyleSheet", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QPageSetupDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setStyleSheet")
	}
}

func (ptr *QPageSetupDialog) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QPageSetupDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPageSetupDialog_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QPageSetupDialog) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QPageSetupDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPageSetupDialog_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQPageSetupDialog_SetWindowModified
func callbackQPageSetupDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPageSetupDialog) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QPageSetupDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setWindowModified", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QPageSetupDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setWindowModified")
	}
}

func (ptr *QPageSetupDialog) SetWindowModified(vbo bool) {
	defer qt.Recovering("QPageSetupDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPageSetupDialog) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QPageSetupDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPageSetupDialog_SetWindowTitle
func callbackQPageSetupDialog_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QPageSetupDialog) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QPageSetupDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setWindowTitle", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QPageSetupDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setWindowTitle")
	}
}

func (ptr *QPageSetupDialog) SetWindowTitle(vqs string) {
	defer qt.Recovering("QPageSetupDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPageSetupDialog_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QPageSetupDialog) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QPageSetupDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPageSetupDialog_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQPageSetupDialog_ChangeEvent
func callbackQPageSetupDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::changeEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::changeEvent")
	}
}

func (ptr *QPageSetupDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPageSetupDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_Close
func callbackQPageSetupDialog_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPageSetupDialog::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QPageSetupDialog) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QPageSetupDialog::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::close", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectClose() {
	defer qt.Recovering("disconnect QPageSetupDialog::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::close")
	}
}

func (ptr *QPageSetupDialog) Close() bool {
	defer qt.Recovering("QPageSetupDialog::close")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPageSetupDialog) CloseDefault() bool {
	defer qt.Recovering("QPageSetupDialog::close")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPageSetupDialog_FocusNextPrevChild
func callbackQPageSetupDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QPageSetupDialog::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPageSetupDialog) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QPageSetupDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusNextPrevChild", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QPageSetupDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::focusNextPrevChild")
	}
}

func (ptr *QPageSetupDialog) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QPageSetupDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QPageSetupDialog) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QPageSetupDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQPageSetupDialog_HasHeightForWidth
func callbackQPageSetupDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPageSetupDialog::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPageSetupDialog) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QPageSetupDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hasHeightForWidth", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QPageSetupDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hasHeightForWidth")
	}
}

func (ptr *QPageSetupDialog) HasHeightForWidth() bool {
	defer qt.Recovering("QPageSetupDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPageSetupDialog) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QPageSetupDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPageSetupDialog_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPageSetupDialog_HeightForWidth
func callbackQPageSetupDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QPageSetupDialog::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPageSetupDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPageSetupDialog) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QPageSetupDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::heightForWidth", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QPageSetupDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::heightForWidth")
	}
}

func (ptr *QPageSetupDialog) HeightForWidth(w int) int {
	defer qt.Recovering("QPageSetupDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QPageSetupDialog) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QPageSetupDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPageSetupDialog_Hide
func callbackQPageSetupDialog_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectHide(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hide", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectHide() {
	defer qt.Recovering("disconnect QPageSetupDialog::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::hide")
	}
}

func (ptr *QPageSetupDialog) Hide() {
	defer qt.Recovering("QPageSetupDialog::hide")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Hide(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) HideDefault() {
	defer qt.Recovering("QPageSetupDialog::hide")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_InputMethodEvent
func callbackQPageSetupDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::inputMethodEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::inputMethodEvent")
	}
}

func (ptr *QPageSetupDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QPageSetupDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPageSetupDialog_InputMethodQuery
func callbackQPageSetupDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QPageSetupDialog::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPageSetupDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPageSetupDialog) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QPageSetupDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::inputMethodQuery", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QPageSetupDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::inputMethodQuery")
	}
}

func (ptr *QPageSetupDialog) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPageSetupDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPageSetupDialog_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPageSetupDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPageSetupDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_KeyReleaseEvent
func callbackQPageSetupDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::keyReleaseEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::keyReleaseEvent")
	}
}

func (ptr *QPageSetupDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QPageSetupDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPageSetupDialog_Lower
func callbackQPageSetupDialog_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectLower(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::lower", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectLower() {
	defer qt.Recovering("disconnect QPageSetupDialog::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::lower")
	}
}

func (ptr *QPageSetupDialog) Lower() {
	defer qt.Recovering("QPageSetupDialog::lower")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Lower(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) LowerDefault() {
	defer qt.Recovering("QPageSetupDialog::lower")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_MouseDoubleClickEvent
func callbackQPageSetupDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseDoubleClickEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseDoubleClickEvent")
	}
}

func (ptr *QPageSetupDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPageSetupDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MouseMoveEvent
func callbackQPageSetupDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseMoveEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseMoveEvent")
	}
}

func (ptr *QPageSetupDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPageSetupDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MousePressEvent
func callbackQPageSetupDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mousePressEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mousePressEvent")
	}
}

func (ptr *QPageSetupDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPageSetupDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MouseReleaseEvent
func callbackQPageSetupDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseReleaseEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::mouseReleaseEvent")
	}
}

func (ptr *QPageSetupDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPageSetupDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_NativeEvent
func callbackQPageSetupDialog_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QPageSetupDialog::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QPageSetupDialog) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QPageSetupDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::nativeEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::nativeEvent")
	}
}

func (ptr *QPageSetupDialog) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPageSetupDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPageSetupDialog_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QPageSetupDialog) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPageSetupDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPageSetupDialog_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQPageSetupDialog_Raise
func callbackQPageSetupDialog_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectRaise(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::raise", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectRaise() {
	defer qt.Recovering("disconnect QPageSetupDialog::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::raise")
	}
}

func (ptr *QPageSetupDialog) Raise() {
	defer qt.Recovering("QPageSetupDialog::raise")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Raise(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) RaiseDefault() {
	defer qt.Recovering("QPageSetupDialog::raise")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_Repaint
func callbackQPageSetupDialog_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::repaint", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectRepaint() {
	defer qt.Recovering("disconnect QPageSetupDialog::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::repaint")
	}
}

func (ptr *QPageSetupDialog) Repaint() {
	defer qt.Recovering("QPageSetupDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Repaint(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) RepaintDefault() {
	defer qt.Recovering("QPageSetupDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_SetDisabled
func callbackQPageSetupDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPageSetupDialog) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QPageSetupDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setDisabled", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QPageSetupDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setDisabled")
	}
}

func (ptr *QPageSetupDialog) SetDisabled(disable bool) {
	defer qt.Recovering("QPageSetupDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QPageSetupDialog) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QPageSetupDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPageSetupDialog_SetFocus2
func callbackQPageSetupDialog_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPageSetupDialog) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setFocus2", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QPageSetupDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setFocus2")
	}
}

func (ptr *QPageSetupDialog) SetFocus2() {
	defer qt.Recovering("QPageSetupDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) SetFocus2Default() {
	defer qt.Recovering("QPageSetupDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_SetHidden
func callbackQPageSetupDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QPageSetupDialog::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPageSetupDialog) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QPageSetupDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setHidden", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QPageSetupDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::setHidden")
	}
}

func (ptr *QPageSetupDialog) SetHidden(hidden bool) {
	defer qt.Recovering("QPageSetupDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QPageSetupDialog) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QPageSetupDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPageSetupDialog_Show
func callbackQPageSetupDialog_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::show"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectShow(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::show", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShow() {
	defer qt.Recovering("disconnect QPageSetupDialog::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::show")
	}
}

func (ptr *QPageSetupDialog) Show() {
	defer qt.Recovering("QPageSetupDialog::show")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Show(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) ShowDefault() {
	defer qt.Recovering("QPageSetupDialog::show")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowFullScreen
func callbackQPageSetupDialog_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showFullScreen", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QPageSetupDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showFullScreen")
	}
}

func (ptr *QPageSetupDialog) ShowFullScreen() {
	defer qt.Recovering("QPageSetupDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) ShowFullScreenDefault() {
	defer qt.Recovering("QPageSetupDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowMaximized
func callbackQPageSetupDialog_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showMaximized", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QPageSetupDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showMaximized")
	}
}

func (ptr *QPageSetupDialog) ShowMaximized() {
	defer qt.Recovering("QPageSetupDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) ShowMaximizedDefault() {
	defer qt.Recovering("QPageSetupDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowMinimized
func callbackQPageSetupDialog_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showMinimized", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QPageSetupDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showMinimized")
	}
}

func (ptr *QPageSetupDialog) ShowMinimized() {
	defer qt.Recovering("QPageSetupDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) ShowMinimizedDefault() {
	defer qt.Recovering("QPageSetupDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowNormal
func callbackQPageSetupDialog_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showNormal", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QPageSetupDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::showNormal")
	}
}

func (ptr *QPageSetupDialog) ShowNormal() {
	defer qt.Recovering("QPageSetupDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) ShowNormalDefault() {
	defer qt.Recovering("QPageSetupDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_TabletEvent
func callbackQPageSetupDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::tabletEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::tabletEvent")
	}
}

func (ptr *QPageSetupDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPageSetupDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPageSetupDialog_Update
func callbackQPageSetupDialog_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::update"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::update", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectUpdate() {
	defer qt.Recovering("disconnect QPageSetupDialog::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::update")
	}
}

func (ptr *QPageSetupDialog) Update() {
	defer qt.Recovering("QPageSetupDialog::update")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Update(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) UpdateDefault() {
	defer qt.Recovering("QPageSetupDialog::update")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_UpdateMicroFocus
func callbackQPageSetupDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::updateMicroFocus", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QPageSetupDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::updateMicroFocus")
	}
}

func (ptr *QPageSetupDialog) UpdateMicroFocus() {
	defer qt.Recovering("QPageSetupDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QPageSetupDialog) UpdateMicroFocusDefault() {
	defer qt.Recovering("QPageSetupDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_WheelEvent
func callbackQPageSetupDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::wheelEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::wheelEvent")
	}
}

func (ptr *QPageSetupDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QPageSetupDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPageSetupDialog_TimerEvent
func callbackQPageSetupDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::timerEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::timerEvent")
	}
}

func (ptr *QPageSetupDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPageSetupDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPageSetupDialog_ChildEvent
func callbackQPageSetupDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::childEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::childEvent")
	}
}

func (ptr *QPageSetupDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPageSetupDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPageSetupDialog_ConnectNotify
func callbackQPageSetupDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPageSetupDialog) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPageSetupDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::connectNotify", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QPageSetupDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::connectNotify")
	}
}

func (ptr *QPageSetupDialog) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPageSetupDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPageSetupDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPageSetupDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPageSetupDialog_CustomEvent
func callbackQPageSetupDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPageSetupDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::customEvent", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPageSetupDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::customEvent")
	}
}

func (ptr *QPageSetupDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPageSetupDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPageSetupDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_DeleteLater
func callbackQPageSetupDialog_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QPageSetupDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::deleteLater", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QPageSetupDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::deleteLater")
	}
}

func (ptr *QPageSetupDialog) DeleteLater() {
	defer qt.Recovering("QPageSetupDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPageSetupDialog) DeleteLaterDefault() {
	defer qt.Recovering("QPageSetupDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPageSetupDialog_DisconnectNotify
func callbackQPageSetupDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPageSetupDialog::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPageSetupDialog) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPageSetupDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::disconnectNotify", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QPageSetupDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::disconnectNotify")
	}
}

func (ptr *QPageSetupDialog) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPageSetupDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPageSetupDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPageSetupDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPageSetupDialog_MetaObject
func callbackQPageSetupDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPageSetupDialog::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPageSetupDialog::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPageSetupDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPageSetupDialog) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QPageSetupDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::metaObject", f)
	}
}

func (ptr *QPageSetupDialog) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QPageSetupDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPageSetupDialog::metaObject")
	}
}

func (ptr *QPageSetupDialog) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QPageSetupDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPageSetupDialog_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPageSetupDialog) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QPageSetupDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPageSetupDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QPrintDialog struct {
	QAbstractPrintDialog
}

type QPrintDialog_ITF interface {
	QAbstractPrintDialog_ITF
	QPrintDialog_PTR() *QPrintDialog
}

func (p *QPrintDialog) QPrintDialog_PTR() *QPrintDialog {
	return p
}

func (p *QPrintDialog) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func (p *QPrintDialog) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractPrintDialog_PTR().SetPointer(ptr)
	}
}

func PointerFromQPrintDialog(ptr QPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintDialog_PTR().Pointer()
	}
	return nil
}

func NewQPrintDialogFromPointer(ptr unsafe.Pointer) *QPrintDialog {
	var n = new(QPrintDialog)
	n.SetPointer(ptr)
	return n
}
func (ptr *QPrintDialog) Options() QAbstractPrintDialog__PrintDialogOption {
	defer qt.Recovering("QPrintDialog::options")

	if ptr.Pointer() != nil {
		return QAbstractPrintDialog__PrintDialogOption(C.QPrintDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintDialog) SetOptions(options QAbstractPrintDialog__PrintDialogOption) {
	defer qt.Recovering("QPrintDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetOptions(ptr.Pointer(), C.longlong(options))
	}
}

//export callbackQPrintDialog_Accepted
func callbackQPrintDialog_Accepted(ptr unsafe.Pointer, printer unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::accepted")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::accepted"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintDialog) ConnectAccepted(f func(printer *QPrinter)) {
	defer qt.Recovering("connect QPrintDialog::accepted")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::accepted", f)
	}
}

func (ptr *QPrintDialog) DisconnectAccepted() {
	defer qt.Recovering("disconnect QPrintDialog::accepted")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::accepted")
	}
}

func (ptr *QPrintDialog) Accepted(printer QPrinter_ITF) {
	defer qt.Recovering("QPrintDialog::accepted")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Accepted(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

func (ptr *QPrintDialog) Done(result int) {
	defer qt.Recovering("QPrintDialog::done")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPrintDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QPrintDialog::open")

	if ptr.Pointer() != nil {
		var memberC = C.CString(member)
		defer C.free(unsafe.Pointer(memberC))
		C.QPrintDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

func (ptr *QPrintDialog) Printer() *QPrinter {
	defer qt.Recovering("QPrintDialog::printer")

	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QPrintDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintDialog) SetOption(option QAbstractPrintDialog__PrintDialogOption, on bool) {
	defer qt.Recovering("QPrintDialog::setOption")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetOption(ptr.Pointer(), C.longlong(option), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QPrintDialog) TestOption(option QAbstractPrintDialog__PrintDialogOption) bool {
	defer qt.Recovering("QPrintDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_TestOption(ptr.Pointer(), C.longlong(option)) != 0
	}
	return false
}

func NewQPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPrintDialog {
	defer qt.Recovering("QPrintDialog::QPrintDialog")

	var tmpValue = NewQPrintDialogFromPointer(C.QPrintDialog_NewQPrintDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func NewQPrintDialog2(parent widgets.QWidget_ITF) *QPrintDialog {
	defer qt.Recovering("QPrintDialog::QPrintDialog")

	var tmpValue = NewQPrintDialogFromPointer(C.QPrintDialog_NewQPrintDialog2(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQPrintDialog_Exec
func callbackQPrintDialog_Exec(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QPrintDialog::exec")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQPrintDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QPrintDialog) ConnectExec(f func() int) {
	defer qt.Recovering("connect QPrintDialog::exec")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::exec", f)
	}
}

func (ptr *QPrintDialog) DisconnectExec() {
	defer qt.Recovering("disconnect QPrintDialog::exec")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::exec")
	}
}

func (ptr *QPrintDialog) Exec() int {
	defer qt.Recovering("QPrintDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintDialog) ExecDefault() int {
	defer qt.Recovering("QPrintDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintDialog) SetVisible(visible bool) {
	defer qt.Recovering("QPrintDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QPrintDialog) DestroyQPrintDialog() {
	defer qt.Recovering("QPrintDialog::~QPrintDialog")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DestroyQPrintDialog(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintDialog_Accept
func callbackQPrintDialog_Accept(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::accept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPrintDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QPrintDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::accept", f)
	}
}

func (ptr *QPrintDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QPrintDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::accept")
	}
}

func (ptr *QPrintDialog) Accept() {
	defer qt.Recovering("QPrintDialog::accept")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) AcceptDefault() {
	defer qt.Recovering("QPrintDialog::accept")

	if ptr.Pointer() != nil {
		C.QPrintDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_CloseEvent
func callbackQPrintDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQPrintDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QPrintDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPrintDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::closeEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPrintDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::closeEvent")
	}
}

func (ptr *QPrintDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QPrintDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQPrintDialog_ContextMenuEvent
func callbackQPrintDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQPrintDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QPrintDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::contextMenuEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::contextMenuEvent")
	}
}

func (ptr *QPrintDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QPrintDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQPrintDialog_KeyPressEvent
func callbackQPrintDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPrintDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPrintDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::keyPressEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::keyPressEvent")
	}
}

func (ptr *QPrintDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QPrintDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQPrintDialog_MinimumSizeHint
func callbackQPrintDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintDialog::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPrintDialog) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::minimumSizeHint", f)
	}
}

func (ptr *QPrintDialog) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::minimumSizeHint")
	}
}

func (ptr *QPrintDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintDialog_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintDialog) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintDialog_Reject
func callbackQPrintDialog_Reject(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::reject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::reject"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QPrintDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QPrintDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::reject", f)
	}
}

func (ptr *QPrintDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QPrintDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::reject")
	}
}

func (ptr *QPrintDialog) Reject() {
	defer qt.Recovering("QPrintDialog::reject")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) RejectDefault() {
	defer qt.Recovering("QPrintDialog::reject")

	if ptr.Pointer() != nil {
		C.QPrintDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_ResizeEvent
func callbackQPrintDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQPrintDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QPrintDialog) ConnectResizeEvent(f func(vqr *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::resizeEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::resizeEvent")
	}
}

func (ptr *QPrintDialog) ResizeEvent(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

func (ptr *QPrintDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQPrintDialog_ShowEvent
func callbackQPrintDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QPrintDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPrintDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showEvent")
	}
}

func (ptr *QPrintDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QPrintDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPrintDialog_SizeHint
func callbackQPrintDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintDialog::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPrintDialog) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::sizeHint", f)
	}
}

func (ptr *QPrintDialog) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QPrintDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::sizeHint")
	}
}

func (ptr *QPrintDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QPrintDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintDialog_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintDialog) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintDialog_ActionEvent
func callbackQPrintDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPrintDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::actionEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPrintDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::actionEvent")
	}
}

func (ptr *QPrintDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPrintDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPrintDialog_DragEnterEvent
func callbackQPrintDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragEnterEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragEnterEvent")
	}
}

func (ptr *QPrintDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QPrintDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPrintDialog_DragLeaveEvent
func callbackQPrintDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragLeaveEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragLeaveEvent")
	}
}

func (ptr *QPrintDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QPrintDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPrintDialog_DragMoveEvent
func callbackQPrintDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragMoveEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dragMoveEvent")
	}
}

func (ptr *QPrintDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QPrintDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPrintDialog_DropEvent
func callbackQPrintDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QPrintDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dropEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPrintDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::dropEvent")
	}
}

func (ptr *QPrintDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QPrintDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPrintDialog_EnterEvent
func callbackQPrintDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::enterEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPrintDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::enterEvent")
	}
}

func (ptr *QPrintDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintDialog_FocusInEvent
func callbackQPrintDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusInEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusInEvent")
	}
}

func (ptr *QPrintDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintDialog_FocusOutEvent
func callbackQPrintDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusOutEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusOutEvent")
	}
}

func (ptr *QPrintDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintDialog_HideEvent
func callbackQPrintDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPrintDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hideEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPrintDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hideEvent")
	}
}

func (ptr *QPrintDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPrintDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPrintDialog_LeaveEvent
func callbackQPrintDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::leaveEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::leaveEvent")
	}
}

func (ptr *QPrintDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintDialog_MoveEvent
func callbackQPrintDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPrintDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::moveEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPrintDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::moveEvent")
	}
}

func (ptr *QPrintDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPrintDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPrintDialog_PaintEvent
func callbackQPrintDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPrintDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::paintEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPrintDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::paintEvent")
	}
}

func (ptr *QPrintDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QPrintDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPrintDialog_SetEnabled
func callbackQPrintDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintDialog::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintDialog) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setEnabled", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QPrintDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setEnabled")
	}
}

func (ptr *QPrintDialog) SetEnabled(vbo bool) {
	defer qt.Recovering("QPrintDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintDialog) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QPrintDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintDialog_SetStyleSheet
func callbackQPrintDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QPrintDialog::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQPrintDialogFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QPrintDialog) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setStyleSheet", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setStyleSheet")
	}
}

func (ptr *QPrintDialog) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintDialog_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QPrintDialog) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QPrintDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintDialog_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQPrintDialog_SetWindowModified
func callbackQPrintDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintDialog::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintDialog) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setWindowModified", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setWindowModified")
	}
}

func (ptr *QPrintDialog) SetWindowModified(vbo bool) {
	defer qt.Recovering("QPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintDialog) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QPrintDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintDialog_SetWindowTitle
func callbackQPrintDialog_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QPrintDialog::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQPrintDialogFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QPrintDialog) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setWindowTitle", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setWindowTitle")
	}
}

func (ptr *QPrintDialog) SetWindowTitle(vqs string) {
	defer qt.Recovering("QPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintDialog_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QPrintDialog) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QPrintDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintDialog_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQPrintDialog_ChangeEvent
func callbackQPrintDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::changeEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPrintDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::changeEvent")
	}
}

func (ptr *QPrintDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintDialog_Close
func callbackQPrintDialog_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintDialog::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QPrintDialog) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QPrintDialog::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::close", f)
	}
}

func (ptr *QPrintDialog) DisconnectClose() {
	defer qt.Recovering("disconnect QPrintDialog::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::close")
	}
}

func (ptr *QPrintDialog) Close() bool {
	defer qt.Recovering("QPrintDialog::close")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintDialog) CloseDefault() bool {
	defer qt.Recovering("QPrintDialog::close")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintDialog_FocusNextPrevChild
func callbackQPrintDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QPrintDialog::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPrintDialog) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusNextPrevChild", f)
	}
}

func (ptr *QPrintDialog) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::focusNextPrevChild")
	}
}

func (ptr *QPrintDialog) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QPrintDialog) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QPrintDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQPrintDialog_HasHeightForWidth
func callbackQPrintDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintDialog::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPrintDialog) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hasHeightForWidth", f)
	}
}

func (ptr *QPrintDialog) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hasHeightForWidth")
	}
}

func (ptr *QPrintDialog) HasHeightForWidth() bool {
	defer qt.Recovering("QPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintDialog) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QPrintDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintDialog_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintDialog_HeightForWidth
func callbackQPrintDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QPrintDialog::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPrintDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPrintDialog) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::heightForWidth", f)
	}
}

func (ptr *QPrintDialog) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::heightForWidth")
	}
}

func (ptr *QPrintDialog) HeightForWidth(w int) int {
	defer qt.Recovering("QPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QPrintDialog) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QPrintDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPrintDialog_Hide
func callbackQPrintDialog_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPrintDialog) ConnectHide(f func()) {
	defer qt.Recovering("connect QPrintDialog::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hide", f)
	}
}

func (ptr *QPrintDialog) DisconnectHide() {
	defer qt.Recovering("disconnect QPrintDialog::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::hide")
	}
}

func (ptr *QPrintDialog) Hide() {
	defer qt.Recovering("QPrintDialog::hide")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Hide(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) HideDefault() {
	defer qt.Recovering("QPrintDialog::hide")

	if ptr.Pointer() != nil {
		C.QPrintDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_InputMethodEvent
func callbackQPrintDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::inputMethodEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::inputMethodEvent")
	}
}

func (ptr *QPrintDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QPrintDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPrintDialog_InputMethodQuery
func callbackQPrintDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QPrintDialog::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPrintDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPrintDialog) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::inputMethodQuery", f)
	}
}

func (ptr *QPrintDialog) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::inputMethodQuery")
	}
}

func (ptr *QPrintDialog) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintDialog_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintDialog_KeyReleaseEvent
func callbackQPrintDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::keyReleaseEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::keyReleaseEvent")
	}
}

func (ptr *QPrintDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QPrintDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintDialog_Lower
func callbackQPrintDialog_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPrintDialog) ConnectLower(f func()) {
	defer qt.Recovering("connect QPrintDialog::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::lower", f)
	}
}

func (ptr *QPrintDialog) DisconnectLower() {
	defer qt.Recovering("disconnect QPrintDialog::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::lower")
	}
}

func (ptr *QPrintDialog) Lower() {
	defer qt.Recovering("QPrintDialog::lower")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Lower(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) LowerDefault() {
	defer qt.Recovering("QPrintDialog::lower")

	if ptr.Pointer() != nil {
		C.QPrintDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_MouseDoubleClickEvent
func callbackQPrintDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseDoubleClickEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseDoubleClickEvent")
	}
}

func (ptr *QPrintDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintDialog_MouseMoveEvent
func callbackQPrintDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseMoveEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseMoveEvent")
	}
}

func (ptr *QPrintDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintDialog_MousePressEvent
func callbackQPrintDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mousePressEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mousePressEvent")
	}
}

func (ptr *QPrintDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintDialog_MouseReleaseEvent
func callbackQPrintDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseReleaseEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::mouseReleaseEvent")
	}
}

func (ptr *QPrintDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintDialog_NativeEvent
func callbackQPrintDialog_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QPrintDialog::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintDialogFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QPrintDialog) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::nativeEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::nativeEvent")
	}
}

func (ptr *QPrintDialog) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintDialog_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QPrintDialog) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintDialog_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQPrintDialog_Raise
func callbackQPrintDialog_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPrintDialog) ConnectRaise(f func()) {
	defer qt.Recovering("connect QPrintDialog::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::raise", f)
	}
}

func (ptr *QPrintDialog) DisconnectRaise() {
	defer qt.Recovering("disconnect QPrintDialog::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::raise")
	}
}

func (ptr *QPrintDialog) Raise() {
	defer qt.Recovering("QPrintDialog::raise")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Raise(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) RaiseDefault() {
	defer qt.Recovering("QPrintDialog::raise")

	if ptr.Pointer() != nil {
		C.QPrintDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_Repaint
func callbackQPrintDialog_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPrintDialog) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QPrintDialog::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::repaint", f)
	}
}

func (ptr *QPrintDialog) DisconnectRepaint() {
	defer qt.Recovering("disconnect QPrintDialog::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::repaint")
	}
}

func (ptr *QPrintDialog) Repaint() {
	defer qt.Recovering("QPrintDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Repaint(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) RepaintDefault() {
	defer qt.Recovering("QPrintDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPrintDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_SetDisabled
func callbackQPrintDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QPrintDialog::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPrintDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPrintDialog) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QPrintDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setDisabled", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QPrintDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setDisabled")
	}
}

func (ptr *QPrintDialog) SetDisabled(disable bool) {
	defer qt.Recovering("QPrintDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QPrintDialog) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QPrintDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPrintDialog_SetFocus2
func callbackQPrintDialog_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPrintDialog) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QPrintDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setFocus2", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QPrintDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setFocus2")
	}
}

func (ptr *QPrintDialog) SetFocus2() {
	defer qt.Recovering("QPrintDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) SetFocus2Default() {
	defer qt.Recovering("QPrintDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPrintDialog_SetHidden
func callbackQPrintDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QPrintDialog::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPrintDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPrintDialog) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QPrintDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setHidden", f)
	}
}

func (ptr *QPrintDialog) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QPrintDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::setHidden")
	}
}

func (ptr *QPrintDialog) SetHidden(hidden bool) {
	defer qt.Recovering("QPrintDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QPrintDialog) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QPrintDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPrintDialog_Show
func callbackQPrintDialog_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::show"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPrintDialog) ConnectShow(f func()) {
	defer qt.Recovering("connect QPrintDialog::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::show", f)
	}
}

func (ptr *QPrintDialog) DisconnectShow() {
	defer qt.Recovering("disconnect QPrintDialog::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::show")
	}
}

func (ptr *QPrintDialog) Show() {
	defer qt.Recovering("QPrintDialog::show")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Show(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) ShowDefault() {
	defer qt.Recovering("QPrintDialog::show")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_ShowFullScreen
func callbackQPrintDialog_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPrintDialog) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showFullScreen", f)
	}
}

func (ptr *QPrintDialog) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showFullScreen")
	}
}

func (ptr *QPrintDialog) ShowFullScreen() {
	defer qt.Recovering("QPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) ShowFullScreenDefault() {
	defer qt.Recovering("QPrintDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_ShowMaximized
func callbackQPrintDialog_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPrintDialog) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QPrintDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showMaximized", f)
	}
}

func (ptr *QPrintDialog) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QPrintDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showMaximized")
	}
}

func (ptr *QPrintDialog) ShowMaximized() {
	defer qt.Recovering("QPrintDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) ShowMaximizedDefault() {
	defer qt.Recovering("QPrintDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_ShowMinimized
func callbackQPrintDialog_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPrintDialog) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QPrintDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showMinimized", f)
	}
}

func (ptr *QPrintDialog) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QPrintDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showMinimized")
	}
}

func (ptr *QPrintDialog) ShowMinimized() {
	defer qt.Recovering("QPrintDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) ShowMinimizedDefault() {
	defer qt.Recovering("QPrintDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_ShowNormal
func callbackQPrintDialog_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPrintDialog) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QPrintDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showNormal", f)
	}
}

func (ptr *QPrintDialog) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QPrintDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::showNormal")
	}
}

func (ptr *QPrintDialog) ShowNormal() {
	defer qt.Recovering("QPrintDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) ShowNormalDefault() {
	defer qt.Recovering("QPrintDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_TabletEvent
func callbackQPrintDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::tabletEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::tabletEvent")
	}
}

func (ptr *QPrintDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPrintDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPrintDialog_Update
func callbackQPrintDialog_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::update"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPrintDialog) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QPrintDialog::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::update", f)
	}
}

func (ptr *QPrintDialog) DisconnectUpdate() {
	defer qt.Recovering("disconnect QPrintDialog::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::update")
	}
}

func (ptr *QPrintDialog) Update() {
	defer qt.Recovering("QPrintDialog::update")

	if ptr.Pointer() != nil {
		C.QPrintDialog_Update(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) UpdateDefault() {
	defer qt.Recovering("QPrintDialog::update")

	if ptr.Pointer() != nil {
		C.QPrintDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_UpdateMicroFocus
func callbackQPrintDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPrintDialog) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::updateMicroFocus", f)
	}
}

func (ptr *QPrintDialog) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::updateMicroFocus")
	}
}

func (ptr *QPrintDialog) UpdateMicroFocus() {
	defer qt.Recovering("QPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintDialog_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) UpdateMicroFocusDefault() {
	defer qt.Recovering("QPrintDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_WheelEvent
func callbackQPrintDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::wheelEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::wheelEvent")
	}
}

func (ptr *QPrintDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QPrintDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPrintDialog_TimerEvent
func callbackQPrintDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPrintDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::timerEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPrintDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::timerEvent")
	}
}

func (ptr *QPrintDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPrintDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPrintDialog_ChildEvent
func callbackQPrintDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPrintDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::childEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPrintDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::childEvent")
	}
}

func (ptr *QPrintDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPrintDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPrintDialog_ConnectNotify
func callbackQPrintDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintDialog) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::connectNotify", f)
	}
}

func (ptr *QPrintDialog) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QPrintDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::connectNotify")
	}
}

func (ptr *QPrintDialog) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintDialog_CustomEvent
func callbackQPrintDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::customEvent", f)
	}
}

func (ptr *QPrintDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPrintDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::customEvent")
	}
}

func (ptr *QPrintDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintDialog_DeleteLater
func callbackQPrintDialog_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPrintDialog) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QPrintDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::deleteLater", f)
	}
}

func (ptr *QPrintDialog) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QPrintDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::deleteLater")
	}
}

func (ptr *QPrintDialog) DeleteLater() {
	defer qt.Recovering("QPrintDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrintDialog) DeleteLaterDefault() {
	defer qt.Recovering("QPrintDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintDialog_DisconnectNotify
func callbackQPrintDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintDialog::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintDialog) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::disconnectNotify", f)
	}
}

func (ptr *QPrintDialog) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::disconnectNotify")
	}
}

func (ptr *QPrintDialog) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintDialog_MetaObject
func callbackQPrintDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintDialog::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintDialog::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPrintDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPrintDialog) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QPrintDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::metaObject", f)
	}
}

func (ptr *QPrintDialog) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QPrintDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintDialog::metaObject")
	}
}

func (ptr *QPrintDialog) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QPrintDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintDialog_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintDialog) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QPrintDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QPrintEngine::PrintEnginePropertyKey
type QPrintEngine__PrintEnginePropertyKey int64

const (
	QPrintEngine__PPK_CollateCopies          = QPrintEngine__PrintEnginePropertyKey(0)
	QPrintEngine__PPK_ColorMode              = QPrintEngine__PrintEnginePropertyKey(1)
	QPrintEngine__PPK_Creator                = QPrintEngine__PrintEnginePropertyKey(2)
	QPrintEngine__PPK_DocumentName           = QPrintEngine__PrintEnginePropertyKey(3)
	QPrintEngine__PPK_FullPage               = QPrintEngine__PrintEnginePropertyKey(4)
	QPrintEngine__PPK_NumberOfCopies         = QPrintEngine__PrintEnginePropertyKey(5)
	QPrintEngine__PPK_Orientation            = QPrintEngine__PrintEnginePropertyKey(6)
	QPrintEngine__PPK_OutputFileName         = QPrintEngine__PrintEnginePropertyKey(7)
	QPrintEngine__PPK_PageOrder              = QPrintEngine__PrintEnginePropertyKey(8)
	QPrintEngine__PPK_PageRect               = QPrintEngine__PrintEnginePropertyKey(9)
	QPrintEngine__PPK_PageSize               = QPrintEngine__PrintEnginePropertyKey(10)
	QPrintEngine__PPK_PaperRect              = QPrintEngine__PrintEnginePropertyKey(11)
	QPrintEngine__PPK_PaperSource            = QPrintEngine__PrintEnginePropertyKey(12)
	QPrintEngine__PPK_PrinterName            = QPrintEngine__PrintEnginePropertyKey(13)
	QPrintEngine__PPK_PrinterProgram         = QPrintEngine__PrintEnginePropertyKey(14)
	QPrintEngine__PPK_Resolution             = QPrintEngine__PrintEnginePropertyKey(15)
	QPrintEngine__PPK_SelectionOption        = QPrintEngine__PrintEnginePropertyKey(16)
	QPrintEngine__PPK_SupportedResolutions   = QPrintEngine__PrintEnginePropertyKey(17)
	QPrintEngine__PPK_WindowsPageSize        = QPrintEngine__PrintEnginePropertyKey(18)
	QPrintEngine__PPK_FontEmbedding          = QPrintEngine__PrintEnginePropertyKey(19)
	QPrintEngine__PPK_Duplex                 = QPrintEngine__PrintEnginePropertyKey(20)
	QPrintEngine__PPK_PaperSources           = QPrintEngine__PrintEnginePropertyKey(21)
	QPrintEngine__PPK_CustomPaperSize        = QPrintEngine__PrintEnginePropertyKey(22)
	QPrintEngine__PPK_PageMargins            = QPrintEngine__PrintEnginePropertyKey(23)
	QPrintEngine__PPK_CopyCount              = QPrintEngine__PrintEnginePropertyKey(24)
	QPrintEngine__PPK_SupportsMultipleCopies = QPrintEngine__PrintEnginePropertyKey(25)
	QPrintEngine__PPK_PaperName              = QPrintEngine__PrintEnginePropertyKey(26)
	QPrintEngine__PPK_QPageSize              = QPrintEngine__PrintEnginePropertyKey(27)
	QPrintEngine__PPK_QPageMargins           = QPrintEngine__PrintEnginePropertyKey(28)
	QPrintEngine__PPK_QPageLayout            = QPrintEngine__PrintEnginePropertyKey(29)
	QPrintEngine__PPK_PaperSize              = QPrintEngine__PrintEnginePropertyKey(QPrintEngine__PPK_PageSize)
	QPrintEngine__PPK_CustomBase             = QPrintEngine__PrintEnginePropertyKey(0xff00)
)

type QPrintEngine struct {
	ptr unsafe.Pointer
}

type QPrintEngine_ITF interface {
	QPrintEngine_PTR() *QPrintEngine
}

func (p *QPrintEngine) QPrintEngine_PTR() *QPrintEngine {
	return p
}

func (p *QPrintEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QPrintEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQPrintEngine(ptr QPrintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintEngine_PTR().Pointer()
	}
	return nil
}

func NewQPrintEngineFromPointer(ptr unsafe.Pointer) *QPrintEngine {
	var n = new(QPrintEngine)
	n.SetPointer(ptr)
	return n
}

//export callbackQPrintEngine_Abort
func callbackQPrintEngine_Abort(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintEngine::abort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::abort"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPrintEngine) ConnectAbort(f func() bool) {
	defer qt.Recovering("connect QPrintEngine::abort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::abort", f)
	}
}

func (ptr *QPrintEngine) DisconnectAbort() {
	defer qt.Recovering("disconnect QPrintEngine::abort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::abort")
	}
}

func (ptr *QPrintEngine) Abort() bool {
	defer qt.Recovering("QPrintEngine::abort")

	if ptr.Pointer() != nil {
		return C.QPrintEngine_Abort(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintEngine_Metric
func callbackQPrintEngine_Metric(ptr unsafe.Pointer, id C.longlong) C.int {
	defer qt.Recovering("callback QPrintEngine::metric")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(id))))
	}

	return C.int(int32(0))
}

func (ptr *QPrintEngine) ConnectMetric(f func(id gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QPrintEngine::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::metric", f)
	}
}

func (ptr *QPrintEngine) DisconnectMetric(id gui.QPaintDevice__PaintDeviceMetric) {
	defer qt.Recovering("disconnect QPrintEngine::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::metric")
	}
}

func (ptr *QPrintEngine) Metric(id gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QPrintEngine::metric")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintEngine_Metric(ptr.Pointer(), C.longlong(id))))
	}
	return 0
}

//export callbackQPrintEngine_NewPage
func callbackQPrintEngine_NewPage(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintEngine::newPage")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPrintEngine) ConnectNewPage(f func() bool) {
	defer qt.Recovering("connect QPrintEngine::newPage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::newPage", f)
	}
}

func (ptr *QPrintEngine) DisconnectNewPage() {
	defer qt.Recovering("disconnect QPrintEngine::newPage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::newPage")
	}
}

func (ptr *QPrintEngine) NewPage() bool {
	defer qt.Recovering("QPrintEngine::newPage")

	if ptr.Pointer() != nil {
		return C.QPrintEngine_NewPage(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintEngine_PrinterState
func callbackQPrintEngine_PrinterState(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QPrintEngine::printerState")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::printerState"); signal != nil {
		return C.longlong(signal.(func() QPrinter__PrinterState)())
	}

	return C.longlong(0)
}

func (ptr *QPrintEngine) ConnectPrinterState(f func() QPrinter__PrinterState) {
	defer qt.Recovering("connect QPrintEngine::printerState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::printerState", f)
	}
}

func (ptr *QPrintEngine) DisconnectPrinterState() {
	defer qt.Recovering("disconnect QPrintEngine::printerState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::printerState")
	}
}

func (ptr *QPrintEngine) PrinterState() QPrinter__PrinterState {
	defer qt.Recovering("QPrintEngine::printerState")

	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrintEngine_PrinterState(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrintEngine_Property
func callbackQPrintEngine_Property(ptr unsafe.Pointer, key C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QPrintEngine::property")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::property"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QPrintEngine__PrintEnginePropertyKey) *core.QVariant)(QPrintEngine__PrintEnginePropertyKey(key)))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QPrintEngine) ConnectProperty(f func(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant) {
	defer qt.Recovering("connect QPrintEngine::property")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::property", f)
	}
}

func (ptr *QPrintEngine) DisconnectProperty(key QPrintEngine__PrintEnginePropertyKey) {
	defer qt.Recovering("disconnect QPrintEngine::property")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::property")
	}
}

func (ptr *QPrintEngine) Property(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant {
	defer qt.Recovering("QPrintEngine::property")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintEngine_Property(ptr.Pointer(), C.longlong(key)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintEngine_SetProperty
func callbackQPrintEngine_SetProperty(ptr unsafe.Pointer, key C.longlong, value unsafe.Pointer) {
	defer qt.Recovering("callback QPrintEngine::setProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintEngine::setProperty"); signal != nil {
		signal.(func(QPrintEngine__PrintEnginePropertyKey, *core.QVariant))(QPrintEngine__PrintEnginePropertyKey(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QPrintEngine) ConnectSetProperty(f func(key QPrintEngine__PrintEnginePropertyKey, value *core.QVariant)) {
	defer qt.Recovering("connect QPrintEngine::setProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::setProperty", f)
	}
}

func (ptr *QPrintEngine) DisconnectSetProperty(key QPrintEngine__PrintEnginePropertyKey, value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QPrintEngine::setProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintEngine::setProperty")
	}
}

func (ptr *QPrintEngine) SetProperty(key QPrintEngine__PrintEnginePropertyKey, value core.QVariant_ITF) {
	defer qt.Recovering("QPrintEngine::setProperty")

	if ptr.Pointer() != nil {
		C.QPrintEngine_SetProperty(ptr.Pointer(), C.longlong(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QPrintEngine) DestroyQPrintEngine() {
	defer qt.Recovering("QPrintEngine::~QPrintEngine")

	if ptr.Pointer() != nil {
		C.QPrintEngine_DestroyQPrintEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QPrintPreviewDialog struct {
	widgets.QDialog
}

type QPrintPreviewDialog_ITF interface {
	widgets.QDialog_ITF
	QPrintPreviewDialog_PTR() *QPrintPreviewDialog
}

func (p *QPrintPreviewDialog) QPrintPreviewDialog_PTR() *QPrintPreviewDialog {
	return p
}

func (p *QPrintPreviewDialog) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDialog_PTR().Pointer()
	}
	return nil
}

func (p *QPrintPreviewDialog) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDialog_PTR().SetPointer(ptr)
	}
}

func PointerFromQPrintPreviewDialog(ptr QPrintPreviewDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewDialog_PTR().Pointer()
	}
	return nil
}

func NewQPrintPreviewDialogFromPointer(ptr unsafe.Pointer) *QPrintPreviewDialog {
	var n = new(QPrintPreviewDialog)
	n.SetPointer(ptr)
	return n
}
func NewQPrintPreviewDialog2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {
	defer qt.Recovering("QPrintPreviewDialog::QPrintPreviewDialog")

	var tmpValue = NewQPrintPreviewDialogFromPointer(C.QPrintPreviewDialog_NewQPrintPreviewDialog2(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func NewQPrintPreviewDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {
	defer qt.Recovering("QPrintPreviewDialog::QPrintPreviewDialog")

	var tmpValue = NewQPrintPreviewDialogFromPointer(C.QPrintPreviewDialog_NewQPrintPreviewDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QPrintPreviewDialog) Done(result int) {
	defer qt.Recovering("QPrintPreviewDialog::done")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPrintPreviewDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QPrintPreviewDialog::open")

	if ptr.Pointer() != nil {
		var memberC = C.CString(member)
		defer C.free(unsafe.Pointer(memberC))
		C.QPrintPreviewDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

//export callbackQPrintPreviewDialog_PaintRequested
func callbackQPrintPreviewDialog_PaintRequested(ptr unsafe.Pointer, printer unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::paintRequested")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::paintRequested"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintPreviewDialog) ConnectPaintRequested(f func(printer *QPrinter)) {
	defer qt.Recovering("connect QPrintPreviewDialog::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ConnectPaintRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::paintRequested", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectPaintRequested() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DisconnectPaintRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::paintRequested")
	}
}

func (ptr *QPrintPreviewDialog) PaintRequested(printer QPrinter_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_PaintRequested(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

func (ptr *QPrintPreviewDialog) Printer() *QPrinter {
	defer qt.Recovering("QPrintPreviewDialog::printer")

	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QPrintPreviewDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintPreviewDialog) SetVisible(visible bool) {
	defer qt.Recovering("QPrintPreviewDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QPrintPreviewDialog) DestroyQPrintPreviewDialog() {
	defer qt.Recovering("QPrintPreviewDialog::~QPrintPreviewDialog")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DestroyQPrintPreviewDialog(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintPreviewDialog_Accept
func callbackQPrintPreviewDialog_Accept(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::accept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::accept", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::accept")
	}
}

func (ptr *QPrintPreviewDialog) Accept() {
	defer qt.Recovering("QPrintPreviewDialog::accept")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) AcceptDefault() {
	defer qt.Recovering("QPrintPreviewDialog::accept")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_CloseEvent
func callbackQPrintPreviewDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::closeEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::closeEvent")
	}
}

func (ptr *QPrintPreviewDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QPrintPreviewDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQPrintPreviewDialog_ContextMenuEvent
func callbackQPrintPreviewDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::contextMenuEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::contextMenuEvent")
	}
}

func (ptr *QPrintPreviewDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QPrintPreviewDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQPrintPreviewDialog_Exec
func callbackQPrintPreviewDialog_Exec(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QPrintPreviewDialog::exec")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQPrintPreviewDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QPrintPreviewDialog) ConnectExec(f func() int) {
	defer qt.Recovering("connect QPrintPreviewDialog::exec")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::exec", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectExec() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::exec")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::exec")
	}
}

func (ptr *QPrintPreviewDialog) Exec() int {
	defer qt.Recovering("QPrintPreviewDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintPreviewDialog) ExecDefault() int {
	defer qt.Recovering("QPrintPreviewDialog::exec")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQPrintPreviewDialog_KeyPressEvent
func callbackQPrintPreviewDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::keyPressEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::keyPressEvent")
	}
}

func (ptr *QPrintPreviewDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QPrintPreviewDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQPrintPreviewDialog_MinimumSizeHint
func callbackQPrintPreviewDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewDialog::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPrintPreviewDialog) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintPreviewDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::minimumSizeHint", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::minimumSizeHint")
	}
}

func (ptr *QPrintPreviewDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPrintPreviewDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewDialog_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintPreviewDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_Reject
func callbackQPrintPreviewDialog_Reject(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::reject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::reject"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::reject", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::reject")
	}
}

func (ptr *QPrintPreviewDialog) Reject() {
	defer qt.Recovering("QPrintPreviewDialog::reject")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) RejectDefault() {
	defer qt.Recovering("QPrintPreviewDialog::reject")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ResizeEvent
func callbackQPrintPreviewDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QPrintPreviewDialog) ConnectResizeEvent(f func(vqr *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::resizeEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::resizeEvent")
	}
}

func (ptr *QPrintPreviewDialog) ResizeEvent(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

func (ptr *QPrintPreviewDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQPrintPreviewDialog_ShowEvent
func callbackQPrintPreviewDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showEvent")
	}
}

func (ptr *QPrintPreviewDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPrintPreviewDialog_SizeHint
func callbackQPrintPreviewDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewDialog::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPrintPreviewDialog) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintPreviewDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::sizeHint", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::sizeHint")
	}
}

func (ptr *QPrintPreviewDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QPrintPreviewDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewDialog_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintPreviewDialog::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_ActionEvent
func callbackQPrintPreviewDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::actionEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::actionEvent")
	}
}

func (ptr *QPrintPreviewDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DragEnterEvent
func callbackQPrintPreviewDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragEnterEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragEnterEvent")
	}
}

func (ptr *QPrintPreviewDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DragLeaveEvent
func callbackQPrintPreviewDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragLeaveEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragLeaveEvent")
	}
}

func (ptr *QPrintPreviewDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DragMoveEvent
func callbackQPrintPreviewDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragMoveEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dragMoveEvent")
	}
}

func (ptr *QPrintPreviewDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DropEvent
func callbackQPrintPreviewDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dropEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::dropEvent")
	}
}

func (ptr *QPrintPreviewDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPrintPreviewDialog_EnterEvent
func callbackQPrintPreviewDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::enterEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::enterEvent")
	}
}

func (ptr *QPrintPreviewDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_FocusInEvent
func callbackQPrintPreviewDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusInEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusInEvent")
	}
}

func (ptr *QPrintPreviewDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewDialog_FocusOutEvent
func callbackQPrintPreviewDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusOutEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusOutEvent")
	}
}

func (ptr *QPrintPreviewDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewDialog_HideEvent
func callbackQPrintPreviewDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hideEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hideEvent")
	}
}

func (ptr *QPrintPreviewDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPrintPreviewDialog_LeaveEvent
func callbackQPrintPreviewDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::leaveEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::leaveEvent")
	}
}

func (ptr *QPrintPreviewDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MoveEvent
func callbackQPrintPreviewDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::moveEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::moveEvent")
	}
}

func (ptr *QPrintPreviewDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_PaintEvent
func callbackQPrintPreviewDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::paintEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::paintEvent")
	}
}

func (ptr *QPrintPreviewDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPrintPreviewDialog_SetEnabled
func callbackQPrintPreviewDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setEnabled", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setEnabled")
	}
}

func (ptr *QPrintPreviewDialog) SetEnabled(vbo bool) {
	defer qt.Recovering("QPrintPreviewDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintPreviewDialog) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QPrintPreviewDialog::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewDialog_SetStyleSheet
func callbackQPrintPreviewDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setStyleSheet", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setStyleSheet")
	}
}

func (ptr *QPrintPreviewDialog) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QPrintPreviewDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintPreviewDialog_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QPrintPreviewDialog) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QPrintPreviewDialog::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintPreviewDialog_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQPrintPreviewDialog_SetWindowModified
func callbackQPrintPreviewDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setWindowModified", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setWindowModified")
	}
}

func (ptr *QPrintPreviewDialog) SetWindowModified(vbo bool) {
	defer qt.Recovering("QPrintPreviewDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintPreviewDialog) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QPrintPreviewDialog::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewDialog_SetWindowTitle
func callbackQPrintPreviewDialog_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setWindowTitle", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setWindowTitle")
	}
}

func (ptr *QPrintPreviewDialog) SetWindowTitle(vqs string) {
	defer qt.Recovering("QPrintPreviewDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintPreviewDialog_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QPrintPreviewDialog) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QPrintPreviewDialog::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintPreviewDialog_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQPrintPreviewDialog_ChangeEvent
func callbackQPrintPreviewDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::changeEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::changeEvent")
	}
}

func (ptr *QPrintPreviewDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Close
func callbackQPrintPreviewDialog_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintPreviewDialog::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QPrintPreviewDialog) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QPrintPreviewDialog::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::close", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectClose() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::close")
	}
}

func (ptr *QPrintPreviewDialog) Close() bool {
	defer qt.Recovering("QPrintPreviewDialog::close")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintPreviewDialog) CloseDefault() bool {
	defer qt.Recovering("QPrintPreviewDialog::close")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_FocusNextPrevChild
func callbackQPrintPreviewDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QPrintPreviewDialog::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPrintPreviewDialog) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QPrintPreviewDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusNextPrevChild", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::focusNextPrevChild")
	}
}

func (ptr *QPrintPreviewDialog) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QPrintPreviewDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QPrintPreviewDialog) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QPrintPreviewDialog::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_HasHeightForWidth
func callbackQPrintPreviewDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintPreviewDialog::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPrintPreviewDialog) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QPrintPreviewDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hasHeightForWidth", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hasHeightForWidth")
	}
}

func (ptr *QPrintPreviewDialog) HasHeightForWidth() bool {
	defer qt.Recovering("QPrintPreviewDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintPreviewDialog) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QPrintPreviewDialog::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewDialog_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_HeightForWidth
func callbackQPrintPreviewDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QPrintPreviewDialog::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPrintPreviewDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPrintPreviewDialog) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QPrintPreviewDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::heightForWidth", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::heightForWidth")
	}
}

func (ptr *QPrintPreviewDialog) HeightForWidth(w int) int {
	defer qt.Recovering("QPrintPreviewDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QPrintPreviewDialog) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QPrintPreviewDialog::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPrintPreviewDialog_Hide
func callbackQPrintPreviewDialog_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectHide(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hide", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectHide() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::hide")
	}
}

func (ptr *QPrintPreviewDialog) Hide() {
	defer qt.Recovering("QPrintPreviewDialog::hide")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Hide(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) HideDefault() {
	defer qt.Recovering("QPrintPreviewDialog::hide")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_InputMethodEvent
func callbackQPrintPreviewDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::inputMethodEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::inputMethodEvent")
	}
}

func (ptr *QPrintPreviewDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPrintPreviewDialog_InputMethodQuery
func callbackQPrintPreviewDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewDialog::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPrintPreviewDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPrintPreviewDialog) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QPrintPreviewDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::inputMethodQuery", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::inputMethodQuery")
	}
}

func (ptr *QPrintPreviewDialog) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintPreviewDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintPreviewDialog_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintPreviewDialog::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintPreviewDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_KeyReleaseEvent
func callbackQPrintPreviewDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::keyReleaseEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::keyReleaseEvent")
	}
}

func (ptr *QPrintPreviewDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Lower
func callbackQPrintPreviewDialog_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectLower(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::lower", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectLower() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::lower")
	}
}

func (ptr *QPrintPreviewDialog) Lower() {
	defer qt.Recovering("QPrintPreviewDialog::lower")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Lower(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) LowerDefault() {
	defer qt.Recovering("QPrintPreviewDialog::lower")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_MouseDoubleClickEvent
func callbackQPrintPreviewDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseDoubleClickEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseDoubleClickEvent")
	}
}

func (ptr *QPrintPreviewDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MouseMoveEvent
func callbackQPrintPreviewDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseMoveEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseMoveEvent")
	}
}

func (ptr *QPrintPreviewDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MousePressEvent
func callbackQPrintPreviewDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mousePressEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mousePressEvent")
	}
}

func (ptr *QPrintPreviewDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MouseReleaseEvent
func callbackQPrintPreviewDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseReleaseEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::mouseReleaseEvent")
	}
}

func (ptr *QPrintPreviewDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_NativeEvent
func callbackQPrintPreviewDialog_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QPrintPreviewDialog::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QPrintPreviewDialog) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QPrintPreviewDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::nativeEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::nativeEvent")
	}
}

func (ptr *QPrintPreviewDialog) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintPreviewDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintPreviewDialog_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QPrintPreviewDialog) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintPreviewDialog::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintPreviewDialog_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_Raise
func callbackQPrintPreviewDialog_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectRaise(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::raise", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectRaise() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::raise")
	}
}

func (ptr *QPrintPreviewDialog) Raise() {
	defer qt.Recovering("QPrintPreviewDialog::raise")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Raise(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) RaiseDefault() {
	defer qt.Recovering("QPrintPreviewDialog::raise")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_Repaint
func callbackQPrintPreviewDialog_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::repaint", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectRepaint() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::repaint")
	}
}

func (ptr *QPrintPreviewDialog) Repaint() {
	defer qt.Recovering("QPrintPreviewDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Repaint(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) RepaintDefault() {
	defer qt.Recovering("QPrintPreviewDialog::repaint")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_SetDisabled
func callbackQPrintPreviewDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setDisabled", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setDisabled")
	}
}

func (ptr *QPrintPreviewDialog) SetDisabled(disable bool) {
	defer qt.Recovering("QPrintPreviewDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QPrintPreviewDialog) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QPrintPreviewDialog::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPrintPreviewDialog_SetFocus2
func callbackQPrintPreviewDialog_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setFocus2", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setFocus2")
	}
}

func (ptr *QPrintPreviewDialog) SetFocus2() {
	defer qt.Recovering("QPrintPreviewDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) SetFocus2Default() {
	defer qt.Recovering("QPrintPreviewDialog::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_SetHidden
func callbackQPrintPreviewDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QPrintPreviewDialog::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPrintPreviewDialog) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QPrintPreviewDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setHidden", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::setHidden")
	}
}

func (ptr *QPrintPreviewDialog) SetHidden(hidden bool) {
	defer qt.Recovering("QPrintPreviewDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QPrintPreviewDialog) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QPrintPreviewDialog::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPrintPreviewDialog_Show
func callbackQPrintPreviewDialog_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::show"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectShow(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::show", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShow() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::show")
	}
}

func (ptr *QPrintPreviewDialog) Show() {
	defer qt.Recovering("QPrintPreviewDialog::show")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Show(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) ShowDefault() {
	defer qt.Recovering("QPrintPreviewDialog::show")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowFullScreen
func callbackQPrintPreviewDialog_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showFullScreen", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showFullScreen")
	}
}

func (ptr *QPrintPreviewDialog) ShowFullScreen() {
	defer qt.Recovering("QPrintPreviewDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) ShowFullScreenDefault() {
	defer qt.Recovering("QPrintPreviewDialog::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowMaximized
func callbackQPrintPreviewDialog_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showMaximized", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showMaximized")
	}
}

func (ptr *QPrintPreviewDialog) ShowMaximized() {
	defer qt.Recovering("QPrintPreviewDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) ShowMaximizedDefault() {
	defer qt.Recovering("QPrintPreviewDialog::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowMinimized
func callbackQPrintPreviewDialog_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showMinimized", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showMinimized")
	}
}

func (ptr *QPrintPreviewDialog) ShowMinimized() {
	defer qt.Recovering("QPrintPreviewDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) ShowMinimizedDefault() {
	defer qt.Recovering("QPrintPreviewDialog::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowNormal
func callbackQPrintPreviewDialog_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showNormal", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::showNormal")
	}
}

func (ptr *QPrintPreviewDialog) ShowNormal() {
	defer qt.Recovering("QPrintPreviewDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) ShowNormalDefault() {
	defer qt.Recovering("QPrintPreviewDialog::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_TabletEvent
func callbackQPrintPreviewDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::tabletEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::tabletEvent")
	}
}

func (ptr *QPrintPreviewDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Update
func callbackQPrintPreviewDialog_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::update"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::update", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectUpdate() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::update")
	}
}

func (ptr *QPrintPreviewDialog) Update() {
	defer qt.Recovering("QPrintPreviewDialog::update")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Update(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) UpdateDefault() {
	defer qt.Recovering("QPrintPreviewDialog::update")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_UpdateMicroFocus
func callbackQPrintPreviewDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::updateMicroFocus", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::updateMicroFocus")
	}
}

func (ptr *QPrintPreviewDialog) UpdateMicroFocus() {
	defer qt.Recovering("QPrintPreviewDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewDialog) UpdateMicroFocusDefault() {
	defer qt.Recovering("QPrintPreviewDialog::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_WheelEvent
func callbackQPrintPreviewDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::wheelEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::wheelEvent")
	}
}

func (ptr *QPrintPreviewDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPrintPreviewDialog_TimerEvent
func callbackQPrintPreviewDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::timerEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::timerEvent")
	}
}

func (ptr *QPrintPreviewDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPrintPreviewDialog_ChildEvent
func callbackQPrintPreviewDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::childEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::childEvent")
	}
}

func (ptr *QPrintPreviewDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPrintPreviewDialog_ConnectNotify
func callbackQPrintPreviewDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewDialog) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintPreviewDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::connectNotify", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::connectNotify")
	}
}

func (ptr *QPrintPreviewDialog) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintPreviewDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewDialog_CustomEvent
func callbackQPrintPreviewDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::customEvent", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::customEvent")
	}
}

func (ptr *QPrintPreviewDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DeleteLater
func callbackQPrintPreviewDialog_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QPrintPreviewDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::deleteLater", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::deleteLater")
	}
}

func (ptr *QPrintPreviewDialog) DeleteLater() {
	defer qt.Recovering("QPrintPreviewDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrintPreviewDialog) DeleteLaterDefault() {
	defer qt.Recovering("QPrintPreviewDialog::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintPreviewDialog_DisconnectNotify
func callbackQPrintPreviewDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewDialog::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintPreviewDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::disconnectNotify", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::disconnectNotify")
	}
}

func (ptr *QPrintPreviewDialog) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintPreviewDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewDialog::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewDialog_MetaObject
func callbackQPrintPreviewDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewDialog::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewDialog::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPrintPreviewDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPrintPreviewDialog) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QPrintPreviewDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::metaObject", f)
	}
}

func (ptr *QPrintPreviewDialog) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QPrintPreviewDialog::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewDialog::metaObject")
	}
}

func (ptr *QPrintPreviewDialog) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QPrintPreviewDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewDialog_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintPreviewDialog) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QPrintPreviewDialog::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QPrintPreviewWidget::ViewMode
type QPrintPreviewWidget__ViewMode int64

const (
	QPrintPreviewWidget__SinglePageView  = QPrintPreviewWidget__ViewMode(0)
	QPrintPreviewWidget__FacingPagesView = QPrintPreviewWidget__ViewMode(1)
	QPrintPreviewWidget__AllPagesView    = QPrintPreviewWidget__ViewMode(2)
)

//QPrintPreviewWidget::ZoomMode
type QPrintPreviewWidget__ZoomMode int64

const (
	QPrintPreviewWidget__CustomZoom = QPrintPreviewWidget__ZoomMode(0)
	QPrintPreviewWidget__FitToWidth = QPrintPreviewWidget__ZoomMode(1)
	QPrintPreviewWidget__FitInView  = QPrintPreviewWidget__ZoomMode(2)
)

type QPrintPreviewWidget struct {
	widgets.QWidget
}

type QPrintPreviewWidget_ITF interface {
	widgets.QWidget_ITF
	QPrintPreviewWidget_PTR() *QPrintPreviewWidget
}

func (p *QPrintPreviewWidget) QPrintPreviewWidget_PTR() *QPrintPreviewWidget {
	return p
}

func (p *QPrintPreviewWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QPrintPreviewWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQPrintPreviewWidget(ptr QPrintPreviewWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewWidget_PTR().Pointer()
	}
	return nil
}

func NewQPrintPreviewWidgetFromPointer(ptr unsafe.Pointer) *QPrintPreviewWidget {
	var n = new(QPrintPreviewWidget)
	n.SetPointer(ptr)
	return n
}
func NewQPrintPreviewWidget(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {
	defer qt.Recovering("QPrintPreviewWidget::QPrintPreviewWidget")

	var tmpValue = NewQPrintPreviewWidgetFromPointer(C.QPrintPreviewWidget_NewQPrintPreviewWidget(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func NewQPrintPreviewWidget2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {
	defer qt.Recovering("QPrintPreviewWidget::QPrintPreviewWidget")

	var tmpValue = NewQPrintPreviewWidgetFromPointer(C.QPrintPreviewWidget_NewQPrintPreviewWidget2(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QPrintPreviewWidget) CurrentPage() int {
	defer qt.Recovering("QPrintPreviewWidget::currentPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_CurrentPage(ptr.Pointer())))
	}
	return 0
}

//export callbackQPrintPreviewWidget_FitInView
func callbackQPrintPreviewWidget_FitInView(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::fitInView")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::fitInView"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectFitInView(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::fitInView")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::fitInView", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFitInView() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::fitInView")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::fitInView")
	}
}

func (ptr *QPrintPreviewWidget) FitInView() {
	defer qt.Recovering("QPrintPreviewWidget::fitInView")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitInView(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_FitToWidth
func callbackQPrintPreviewWidget_FitToWidth(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::fitToWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::fitToWidth"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectFitToWidth(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::fitToWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::fitToWidth", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFitToWidth() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::fitToWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::fitToWidth")
	}
}

func (ptr *QPrintPreviewWidget) FitToWidth() {
	defer qt.Recovering("QPrintPreviewWidget::fitToWidth")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitToWidth(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) Orientation() QPrinter__Orientation {
	defer qt.Recovering("QPrintPreviewWidget::orientation")

	if ptr.Pointer() != nil {
		return QPrinter__Orientation(C.QPrintPreviewWidget_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) PageCount() int {
	defer qt.Recovering("QPrintPreviewWidget::pageCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_PageCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQPrintPreviewWidget_PaintRequested
func callbackQPrintPreviewWidget_PaintRequested(ptr unsafe.Pointer, printer unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::paintRequested")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::paintRequested"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintPreviewWidget) ConnectPaintRequested(f func(printer *QPrinter)) {
	defer qt.Recovering("connect QPrintPreviewWidget::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ConnectPaintRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::paintRequested", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPaintRequested() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectPaintRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::paintRequested")
	}
}

func (ptr *QPrintPreviewWidget) PaintRequested(printer QPrinter_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::paintRequested")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PaintRequested(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

//export callbackQPrintPreviewWidget_PreviewChanged
func callbackQPrintPreviewWidget_PreviewChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::previewChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::previewChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectPreviewChanged(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::previewChanged")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ConnectPreviewChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::previewChanged", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPreviewChanged() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::previewChanged")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectPreviewChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::previewChanged")
	}
}

func (ptr *QPrintPreviewWidget) PreviewChanged() {
	defer qt.Recovering("QPrintPreviewWidget::previewChanged")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PreviewChanged(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_Print
func callbackQPrintPreviewWidget_Print(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::print")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::print"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectPrint(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::print")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::print", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPrint() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::print")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::print")
	}
}

func (ptr *QPrintPreviewWidget) Print() {
	defer qt.Recovering("QPrintPreviewWidget::print")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Print(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetAllPagesViewMode
func callbackQPrintPreviewWidget_SetAllPagesViewMode(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setAllPagesViewMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setAllPagesViewMode"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetAllPagesViewMode(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setAllPagesViewMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setAllPagesViewMode", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetAllPagesViewMode() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setAllPagesViewMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setAllPagesViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetAllPagesViewMode() {
	defer qt.Recovering("QPrintPreviewWidget::setAllPagesViewMode")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetAllPagesViewMode(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetCurrentPage
func callbackQPrintPreviewWidget_SetCurrentPage(ptr unsafe.Pointer, page C.int) {
	defer qt.Recovering("callback QPrintPreviewWidget::setCurrentPage")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setCurrentPage"); signal != nil {
		signal.(func(int))(int(int32(page)))
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetCurrentPage(f func(page int)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setCurrentPage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setCurrentPage", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetCurrentPage(page int) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setCurrentPage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setCurrentPage")
	}
}

func (ptr *QPrintPreviewWidget) SetCurrentPage(page int) {
	defer qt.Recovering("QPrintPreviewWidget::setCurrentPage")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetCurrentPage(ptr.Pointer(), C.int(int32(page)))
	}
}

//export callbackQPrintPreviewWidget_SetFacingPagesViewMode
func callbackQPrintPreviewWidget_SetFacingPagesViewMode(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setFacingPagesViewMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setFacingPagesViewMode"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetFacingPagesViewMode(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setFacingPagesViewMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setFacingPagesViewMode", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetFacingPagesViewMode() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setFacingPagesViewMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setFacingPagesViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetFacingPagesViewMode() {
	defer qt.Recovering("QPrintPreviewWidget::setFacingPagesViewMode")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFacingPagesViewMode(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetLandscapeOrientation
func callbackQPrintPreviewWidget_SetLandscapeOrientation(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setLandscapeOrientation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setLandscapeOrientation"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetLandscapeOrientation(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setLandscapeOrientation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setLandscapeOrientation", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetLandscapeOrientation() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setLandscapeOrientation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setLandscapeOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetLandscapeOrientation() {
	defer qt.Recovering("QPrintPreviewWidget::setLandscapeOrientation")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetLandscapeOrientation(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetOrientation
func callbackQPrintPreviewWidget_SetOrientation(ptr unsafe.Pointer, orientation C.longlong) {
	defer qt.Recovering("callback QPrintPreviewWidget::setOrientation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setOrientation"); signal != nil {
		signal.(func(QPrinter__Orientation))(QPrinter__Orientation(orientation))
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetOrientation(f func(orientation QPrinter__Orientation)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setOrientation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setOrientation", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetOrientation(orientation QPrinter__Orientation) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setOrientation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetOrientation(orientation QPrinter__Orientation) {
	defer qt.Recovering("QPrintPreviewWidget::setOrientation")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

//export callbackQPrintPreviewWidget_SetPortraitOrientation
func callbackQPrintPreviewWidget_SetPortraitOrientation(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setPortraitOrientation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setPortraitOrientation"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetPortraitOrientation(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setPortraitOrientation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setPortraitOrientation", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetPortraitOrientation() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setPortraitOrientation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setPortraitOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetPortraitOrientation() {
	defer qt.Recovering("QPrintPreviewWidget::setPortraitOrientation")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetPortraitOrientation(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetSinglePageViewMode
func callbackQPrintPreviewWidget_SetSinglePageViewMode(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setSinglePageViewMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setSinglePageViewMode"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetSinglePageViewMode(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setSinglePageViewMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setSinglePageViewMode", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetSinglePageViewMode() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setSinglePageViewMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setSinglePageViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetSinglePageViewMode() {
	defer qt.Recovering("QPrintPreviewWidget::setSinglePageViewMode")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetSinglePageViewMode(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetViewMode
func callbackQPrintPreviewWidget_SetViewMode(ptr unsafe.Pointer, mode C.longlong) {
	defer qt.Recovering("callback QPrintPreviewWidget::setViewMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setViewMode"); signal != nil {
		signal.(func(QPrintPreviewWidget__ViewMode))(QPrintPreviewWidget__ViewMode(mode))
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetViewMode(f func(mode QPrintPreviewWidget__ViewMode)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setViewMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setViewMode", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetViewMode(mode QPrintPreviewWidget__ViewMode) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setViewMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetViewMode(mode QPrintPreviewWidget__ViewMode) {
	defer qt.Recovering("QPrintPreviewWidget::setViewMode")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetViewMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QPrintPreviewWidget) SetVisible(visible bool) {
	defer qt.Recovering("QPrintPreviewWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQPrintPreviewWidget_SetZoomFactor
func callbackQPrintPreviewWidget_SetZoomFactor(ptr unsafe.Pointer, factor C.double) {
	defer qt.Recovering("callback QPrintPreviewWidget::setZoomFactor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setZoomFactor"); signal != nil {
		signal.(func(float64))(float64(factor))
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetZoomFactor(f func(factor float64)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setZoomFactor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setZoomFactor", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomFactor(factor float64) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setZoomFactor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setZoomFactor")
	}
}

func (ptr *QPrintPreviewWidget) SetZoomFactor(factor float64) {
	defer qt.Recovering("QPrintPreviewWidget::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQPrintPreviewWidget_SetZoomMode
func callbackQPrintPreviewWidget_SetZoomMode(ptr unsafe.Pointer, zoomMode C.longlong) {
	defer qt.Recovering("callback QPrintPreviewWidget::setZoomMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setZoomMode"); signal != nil {
		signal.(func(QPrintPreviewWidget__ZoomMode))(QPrintPreviewWidget__ZoomMode(zoomMode))
	}

}

func (ptr *QPrintPreviewWidget) ConnectSetZoomMode(f func(zoomMode QPrintPreviewWidget__ZoomMode)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setZoomMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setZoomMode", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomMode(zoomMode QPrintPreviewWidget__ZoomMode) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setZoomMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setZoomMode")
	}
}

func (ptr *QPrintPreviewWidget) SetZoomMode(zoomMode QPrintPreviewWidget__ZoomMode) {
	defer qt.Recovering("QPrintPreviewWidget::setZoomMode")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomMode(ptr.Pointer(), C.longlong(zoomMode))
	}
}

//export callbackQPrintPreviewWidget_UpdatePreview
func callbackQPrintPreviewWidget_UpdatePreview(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::updatePreview")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::updatePreview"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectUpdatePreview(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::updatePreview")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::updatePreview", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectUpdatePreview() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::updatePreview")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::updatePreview")
	}
}

func (ptr *QPrintPreviewWidget) UpdatePreview() {
	defer qt.Recovering("QPrintPreviewWidget::updatePreview")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdatePreview(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ViewMode() QPrintPreviewWidget__ViewMode {
	defer qt.Recovering("QPrintPreviewWidget::viewMode")

	if ptr.Pointer() != nil {
		return QPrintPreviewWidget__ViewMode(C.QPrintPreviewWidget_ViewMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) ZoomFactor() float64 {
	defer qt.Recovering("QPrintPreviewWidget::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QPrintPreviewWidget_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrintPreviewWidget_ZoomIn
func callbackQPrintPreviewWidget_ZoomIn(ptr unsafe.Pointer, factor C.double) {
	defer qt.Recovering("callback QPrintPreviewWidget::zoomIn")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::zoomIn"); signal != nil {
		signal.(func(float64))(float64(factor))
	}

}

func (ptr *QPrintPreviewWidget) ConnectZoomIn(f func(factor float64)) {
	defer qt.Recovering("connect QPrintPreviewWidget::zoomIn")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::zoomIn", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectZoomIn(factor float64) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::zoomIn")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::zoomIn")
	}
}

func (ptr *QPrintPreviewWidget) ZoomIn(factor float64) {
	defer qt.Recovering("QPrintPreviewWidget::zoomIn")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomIn(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPrintPreviewWidget) ZoomMode() QPrintPreviewWidget__ZoomMode {
	defer qt.Recovering("QPrintPreviewWidget::zoomMode")

	if ptr.Pointer() != nil {
		return QPrintPreviewWidget__ZoomMode(C.QPrintPreviewWidget_ZoomMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrintPreviewWidget_ZoomOut
func callbackQPrintPreviewWidget_ZoomOut(ptr unsafe.Pointer, factor C.double) {
	defer qt.Recovering("callback QPrintPreviewWidget::zoomOut")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::zoomOut"); signal != nil {
		signal.(func(float64))(float64(factor))
	}

}

func (ptr *QPrintPreviewWidget) ConnectZoomOut(f func(factor float64)) {
	defer qt.Recovering("connect QPrintPreviewWidget::zoomOut")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::zoomOut", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectZoomOut(factor float64) {
	defer qt.Recovering("disconnect QPrintPreviewWidget::zoomOut")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::zoomOut")
	}
}

func (ptr *QPrintPreviewWidget) ZoomOut(factor float64) {
	defer qt.Recovering("QPrintPreviewWidget::zoomOut")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomOut(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPrintPreviewWidget) DestroyQPrintPreviewWidget() {
	defer qt.Recovering("QPrintPreviewWidget::~QPrintPreviewWidget")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DestroyQPrintPreviewWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintPreviewWidget_ActionEvent
func callbackQPrintPreviewWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::actionEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::actionEvent")
	}
}

func (ptr *QPrintPreviewWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DragEnterEvent
func callbackQPrintPreviewWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragEnterEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragEnterEvent")
	}
}

func (ptr *QPrintPreviewWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DragLeaveEvent
func callbackQPrintPreviewWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragLeaveEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragLeaveEvent")
	}
}

func (ptr *QPrintPreviewWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DragMoveEvent
func callbackQPrintPreviewWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragMoveEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dragMoveEvent")
	}
}

func (ptr *QPrintPreviewWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DropEvent
func callbackQPrintPreviewWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dropEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::dropEvent")
	}
}

func (ptr *QPrintPreviewWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPrintPreviewWidget_EnterEvent
func callbackQPrintPreviewWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::enterEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::enterEvent")
	}
}

func (ptr *QPrintPreviewWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_FocusInEvent
func callbackQPrintPreviewWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusInEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusInEvent")
	}
}

func (ptr *QPrintPreviewWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewWidget_FocusOutEvent
func callbackQPrintPreviewWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusOutEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusOutEvent")
	}
}

func (ptr *QPrintPreviewWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewWidget_HideEvent
func callbackQPrintPreviewWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hideEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hideEvent")
	}
}

func (ptr *QPrintPreviewWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPrintPreviewWidget_LeaveEvent
func callbackQPrintPreviewWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::leaveEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::leaveEvent")
	}
}

func (ptr *QPrintPreviewWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MinimumSizeHint
func callbackQPrintPreviewWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPrintPreviewWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintPreviewWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::minimumSizeHint", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::minimumSizeHint")
	}
}

func (ptr *QPrintPreviewWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPrintPreviewWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintPreviewWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_MoveEvent
func callbackQPrintPreviewWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::moveEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::moveEvent")
	}
}

func (ptr *QPrintPreviewWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_PaintEvent
func callbackQPrintPreviewWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::paintEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::paintEvent")
	}
}

func (ptr *QPrintPreviewWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPrintPreviewWidget_SetEnabled
func callbackQPrintPreviewWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setEnabled", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setEnabled")
	}
}

func (ptr *QPrintPreviewWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QPrintPreviewWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintPreviewWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QPrintPreviewWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewWidget_SetStyleSheet
func callbackQPrintPreviewWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setStyleSheet", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setStyleSheet")
	}
}

func (ptr *QPrintPreviewWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QPrintPreviewWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintPreviewWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QPrintPreviewWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QPrintPreviewWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QPrintPreviewWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQPrintPreviewWidget_SetWindowModified
func callbackQPrintPreviewWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setWindowModified", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setWindowModified")
	}
}

func (ptr *QPrintPreviewWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QPrintPreviewWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QPrintPreviewWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QPrintPreviewWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewWidget_SetWindowTitle
func callbackQPrintPreviewWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setWindowTitle", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setWindowTitle")
	}
}

func (ptr *QPrintPreviewWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QPrintPreviewWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintPreviewWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QPrintPreviewWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QPrintPreviewWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QPrintPreviewWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQPrintPreviewWidget_ShowEvent
func callbackQPrintPreviewWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showEvent")
	}
}

func (ptr *QPrintPreviewWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPrintPreviewWidget_SizeHint
func callbackQPrintPreviewWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewWidget::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPrintPreviewWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QPrintPreviewWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::sizeHint", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::sizeHint")
	}
}

func (ptr *QPrintPreviewWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QPrintPreviewWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QPrintPreviewWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QPrintPreviewWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_ChangeEvent
func callbackQPrintPreviewWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::changeEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::changeEvent")
	}
}

func (ptr *QPrintPreviewWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Close
func callbackQPrintPreviewWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintPreviewWidget::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QPrintPreviewWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QPrintPreviewWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::close", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::close")
	}
}

func (ptr *QPrintPreviewWidget) Close() bool {
	defer qt.Recovering("QPrintPreviewWidget::close")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintPreviewWidget) CloseDefault() bool {
	defer qt.Recovering("QPrintPreviewWidget::close")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_CloseEvent
func callbackQPrintPreviewWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::closeEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::closeEvent")
	}
}

func (ptr *QPrintPreviewWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ContextMenuEvent
func callbackQPrintPreviewWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::contextMenuEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::contextMenuEvent")
	}
}

func (ptr *QPrintPreviewWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQPrintPreviewWidget_FocusNextPrevChild
func callbackQPrintPreviewWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QPrintPreviewWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPrintPreviewWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QPrintPreviewWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusNextPrevChild", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::focusNextPrevChild")
	}
}

func (ptr *QPrintPreviewWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QPrintPreviewWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QPrintPreviewWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QPrintPreviewWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_HasHeightForWidth
func callbackQPrintPreviewWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintPreviewWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPrintPreviewWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QPrintPreviewWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hasHeightForWidth", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hasHeightForWidth")
	}
}

func (ptr *QPrintPreviewWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QPrintPreviewWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrintPreviewWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QPrintPreviewWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_HeightForWidth
func callbackQPrintPreviewWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QPrintPreviewWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPrintPreviewWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPrintPreviewWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QPrintPreviewWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::heightForWidth", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::heightForWidth")
	}
}

func (ptr *QPrintPreviewWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QPrintPreviewWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QPrintPreviewWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPrintPreviewWidget_Hide
func callbackQPrintPreviewWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hide", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::hide")
	}
}

func (ptr *QPrintPreviewWidget) Hide() {
	defer qt.Recovering("QPrintPreviewWidget::hide")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) HideDefault() {
	defer qt.Recovering("QPrintPreviewWidget::hide")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_InputMethodEvent
func callbackQPrintPreviewWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::inputMethodEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::inputMethodEvent")
	}
}

func (ptr *QPrintPreviewWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPrintPreviewWidget_InputMethodQuery
func callbackQPrintPreviewWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPrintPreviewWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPrintPreviewWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QPrintPreviewWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::inputMethodQuery", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::inputMethodQuery")
	}
}

func (ptr *QPrintPreviewWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintPreviewWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintPreviewWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPrintPreviewWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QPrintPreviewWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_KeyPressEvent
func callbackQPrintPreviewWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::keyPressEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::keyPressEvent")
	}
}

func (ptr *QPrintPreviewWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewWidget_KeyReleaseEvent
func callbackQPrintPreviewWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::keyReleaseEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::keyReleaseEvent")
	}
}

func (ptr *QPrintPreviewWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Lower
func callbackQPrintPreviewWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::lower", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::lower")
	}
}

func (ptr *QPrintPreviewWidget) Lower() {
	defer qt.Recovering("QPrintPreviewWidget::lower")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) LowerDefault() {
	defer qt.Recovering("QPrintPreviewWidget::lower")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_MouseDoubleClickEvent
func callbackQPrintPreviewWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QPrintPreviewWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MouseMoveEvent
func callbackQPrintPreviewWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseMoveEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseMoveEvent")
	}
}

func (ptr *QPrintPreviewWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MousePressEvent
func callbackQPrintPreviewWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mousePressEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mousePressEvent")
	}
}

func (ptr *QPrintPreviewWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MouseReleaseEvent
func callbackQPrintPreviewWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::mouseReleaseEvent")
	}
}

func (ptr *QPrintPreviewWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_NativeEvent
func callbackQPrintPreviewWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QPrintPreviewWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QPrintPreviewWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QPrintPreviewWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::nativeEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::nativeEvent")
	}
}

func (ptr *QPrintPreviewWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintPreviewWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintPreviewWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QPrintPreviewWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QPrintPreviewWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QPrintPreviewWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_Raise
func callbackQPrintPreviewWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::raise", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::raise")
	}
}

func (ptr *QPrintPreviewWidget) Raise() {
	defer qt.Recovering("QPrintPreviewWidget::raise")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) RaiseDefault() {
	defer qt.Recovering("QPrintPreviewWidget::raise")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_Repaint
func callbackQPrintPreviewWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::repaint", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::repaint")
	}
}

func (ptr *QPrintPreviewWidget) Repaint() {
	defer qt.Recovering("QPrintPreviewWidget::repaint")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) RepaintDefault() {
	defer qt.Recovering("QPrintPreviewWidget::repaint")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ResizeEvent
func callbackQPrintPreviewWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::resizeEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::resizeEvent")
	}
}

func (ptr *QPrintPreviewWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQPrintPreviewWidget_SetDisabled
func callbackQPrintPreviewWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setDisabled", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setDisabled")
	}
}

func (ptr *QPrintPreviewWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QPrintPreviewWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QPrintPreviewWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QPrintPreviewWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPrintPreviewWidget_SetFocus2
func callbackQPrintPreviewWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setFocus2", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setFocus2")
	}
}

func (ptr *QPrintPreviewWidget) SetFocus2() {
	defer qt.Recovering("QPrintPreviewWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetFocus2Default() {
	defer qt.Recovering("QPrintPreviewWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetHidden
func callbackQPrintPreviewWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QPrintPreviewWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QPrintPreviewWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setHidden", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::setHidden")
	}
}

func (ptr *QPrintPreviewWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QPrintPreviewWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QPrintPreviewWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QPrintPreviewWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPrintPreviewWidget_Show
func callbackQPrintPreviewWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::show", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::show")
	}
}

func (ptr *QPrintPreviewWidget) Show() {
	defer qt.Recovering("QPrintPreviewWidget::show")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Show(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ShowDefault() {
	defer qt.Recovering("QPrintPreviewWidget::show")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowFullScreen
func callbackQPrintPreviewWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showFullScreen", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showFullScreen")
	}
}

func (ptr *QPrintPreviewWidget) ShowFullScreen() {
	defer qt.Recovering("QPrintPreviewWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QPrintPreviewWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowMaximized
func callbackQPrintPreviewWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showMaximized", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showMaximized")
	}
}

func (ptr *QPrintPreviewWidget) ShowMaximized() {
	defer qt.Recovering("QPrintPreviewWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QPrintPreviewWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowMinimized
func callbackQPrintPreviewWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showMinimized", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showMinimized")
	}
}

func (ptr *QPrintPreviewWidget) ShowMinimized() {
	defer qt.Recovering("QPrintPreviewWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QPrintPreviewWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowNormal
func callbackQPrintPreviewWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showNormal", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::showNormal")
	}
}

func (ptr *QPrintPreviewWidget) ShowNormal() {
	defer qt.Recovering("QPrintPreviewWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) ShowNormalDefault() {
	defer qt.Recovering("QPrintPreviewWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_TabletEvent
func callbackQPrintPreviewWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::tabletEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::tabletEvent")
	}
}

func (ptr *QPrintPreviewWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Update
func callbackQPrintPreviewWidget_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::update"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::update", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::update")
	}
}

func (ptr *QPrintPreviewWidget) Update() {
	defer qt.Recovering("QPrintPreviewWidget::update")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Update(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) UpdateDefault() {
	defer qt.Recovering("QPrintPreviewWidget::update")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_UpdateMicroFocus
func callbackQPrintPreviewWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::updateMicroFocus", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::updateMicroFocus")
	}
}

func (ptr *QPrintPreviewWidget) UpdateMicroFocus() {
	defer qt.Recovering("QPrintPreviewWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QPrintPreviewWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_WheelEvent
func callbackQPrintPreviewWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::wheelEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::wheelEvent")
	}
}

func (ptr *QPrintPreviewWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPrintPreviewWidget_TimerEvent
func callbackQPrintPreviewWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::timerEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::timerEvent")
	}
}

func (ptr *QPrintPreviewWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ChildEvent
func callbackQPrintPreviewWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::childEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::childEvent")
	}
}

func (ptr *QPrintPreviewWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ConnectNotify
func callbackQPrintPreviewWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintPreviewWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::connectNotify", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::connectNotify")
	}
}

func (ptr *QPrintPreviewWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintPreviewWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewWidget_CustomEvent
func callbackQPrintPreviewWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPrintPreviewWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::customEvent", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::customEvent")
	}
}

func (ptr *QPrintPreviewWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPrintPreviewWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DeleteLater
func callbackQPrintPreviewWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QPrintPreviewWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::deleteLater", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::deleteLater")
	}
}

func (ptr *QPrintPreviewWidget) DeleteLater() {
	defer qt.Recovering("QPrintPreviewWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrintPreviewWidget) DeleteLaterDefault() {
	defer qt.Recovering("QPrintPreviewWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPrintPreviewWidget_DisconnectNotify
func callbackQPrintPreviewWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QPrintPreviewWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QPrintPreviewWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::disconnectNotify", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::disconnectNotify")
	}
}

func (ptr *QPrintPreviewWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPrintPreviewWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QPrintPreviewWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewWidget_EventFilter
func callbackQPrintPreviewWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QPrintPreviewWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPrintPreviewWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QPrintPreviewWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::eventFilter", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::eventFilter")
	}
}

func (ptr *QPrintPreviewWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QPrintPreviewWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QPrintPreviewWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QPrintPreviewWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QPrintPreviewWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_MetaObject
func callbackQPrintPreviewWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QPrintPreviewWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrintPreviewWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPrintPreviewWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPrintPreviewWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QPrintPreviewWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::metaObject", f)
	}
}

func (ptr *QPrintPreviewWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QPrintPreviewWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrintPreviewWidget::metaObject")
	}
}

func (ptr *QPrintPreviewWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QPrintPreviewWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintPreviewWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QPrintPreviewWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QPrinter::ColorMode
type QPrinter__ColorMode int64

const (
	QPrinter__GrayScale = QPrinter__ColorMode(0)
	QPrinter__Color     = QPrinter__ColorMode(1)
)

//QPrinter::DuplexMode
type QPrinter__DuplexMode int64

const (
	QPrinter__DuplexNone      = QPrinter__DuplexMode(0)
	QPrinter__DuplexAuto      = QPrinter__DuplexMode(1)
	QPrinter__DuplexLongSide  = QPrinter__DuplexMode(2)
	QPrinter__DuplexShortSide = QPrinter__DuplexMode(3)
)

//QPrinter::Orientation
type QPrinter__Orientation int64

const (
	QPrinter__Portrait  = QPrinter__Orientation(0)
	QPrinter__Landscape = QPrinter__Orientation(1)
)

//QPrinter::OutputFormat
type QPrinter__OutputFormat int64

const (
	QPrinter__NativeFormat = QPrinter__OutputFormat(0)
	QPrinter__PdfFormat    = QPrinter__OutputFormat(1)
)

//QPrinter::PageOrder
type QPrinter__PageOrder int64

const (
	QPrinter__FirstPageFirst = QPrinter__PageOrder(0)
	QPrinter__LastPageFirst  = QPrinter__PageOrder(1)
)

//QPrinter::PaperSize
type QPrinter__PaperSize int64

const (
	QPrinter__A4                 = QPrinter__PaperSize(0)
	QPrinter__B5                 = QPrinter__PaperSize(1)
	QPrinter__Letter             = QPrinter__PaperSize(2)
	QPrinter__Legal              = QPrinter__PaperSize(3)
	QPrinter__Executive          = QPrinter__PaperSize(4)
	QPrinter__A0                 = QPrinter__PaperSize(5)
	QPrinter__A1                 = QPrinter__PaperSize(6)
	QPrinter__A2                 = QPrinter__PaperSize(7)
	QPrinter__A3                 = QPrinter__PaperSize(8)
	QPrinter__A5                 = QPrinter__PaperSize(9)
	QPrinter__A6                 = QPrinter__PaperSize(10)
	QPrinter__A7                 = QPrinter__PaperSize(11)
	QPrinter__A8                 = QPrinter__PaperSize(12)
	QPrinter__A9                 = QPrinter__PaperSize(13)
	QPrinter__B0                 = QPrinter__PaperSize(14)
	QPrinter__B1                 = QPrinter__PaperSize(15)
	QPrinter__B10                = QPrinter__PaperSize(16)
	QPrinter__B2                 = QPrinter__PaperSize(17)
	QPrinter__B3                 = QPrinter__PaperSize(18)
	QPrinter__B4                 = QPrinter__PaperSize(19)
	QPrinter__B6                 = QPrinter__PaperSize(20)
	QPrinter__B7                 = QPrinter__PaperSize(21)
	QPrinter__B8                 = QPrinter__PaperSize(22)
	QPrinter__B9                 = QPrinter__PaperSize(23)
	QPrinter__C5E                = QPrinter__PaperSize(24)
	QPrinter__Comm10E            = QPrinter__PaperSize(25)
	QPrinter__DLE                = QPrinter__PaperSize(26)
	QPrinter__Folio              = QPrinter__PaperSize(27)
	QPrinter__Ledger             = QPrinter__PaperSize(28)
	QPrinter__Tabloid            = QPrinter__PaperSize(29)
	QPrinter__Custom             = QPrinter__PaperSize(30)
	QPrinter__A10                = QPrinter__PaperSize(31)
	QPrinter__A3Extra            = QPrinter__PaperSize(32)
	QPrinter__A4Extra            = QPrinter__PaperSize(33)
	QPrinter__A4Plus             = QPrinter__PaperSize(34)
	QPrinter__A4Small            = QPrinter__PaperSize(35)
	QPrinter__A5Extra            = QPrinter__PaperSize(36)
	QPrinter__B5Extra            = QPrinter__PaperSize(37)
	QPrinter__JisB0              = QPrinter__PaperSize(38)
	QPrinter__JisB1              = QPrinter__PaperSize(39)
	QPrinter__JisB2              = QPrinter__PaperSize(40)
	QPrinter__JisB3              = QPrinter__PaperSize(41)
	QPrinter__JisB4              = QPrinter__PaperSize(42)
	QPrinter__JisB5              = QPrinter__PaperSize(43)
	QPrinter__JisB6              = QPrinter__PaperSize(44)
	QPrinter__JisB7              = QPrinter__PaperSize(45)
	QPrinter__JisB8              = QPrinter__PaperSize(46)
	QPrinter__JisB9              = QPrinter__PaperSize(47)
	QPrinter__JisB10             = QPrinter__PaperSize(48)
	QPrinter__AnsiC              = QPrinter__PaperSize(49)
	QPrinter__AnsiD              = QPrinter__PaperSize(50)
	QPrinter__AnsiE              = QPrinter__PaperSize(51)
	QPrinter__LegalExtra         = QPrinter__PaperSize(52)
	QPrinter__LetterExtra        = QPrinter__PaperSize(53)
	QPrinter__LetterPlus         = QPrinter__PaperSize(54)
	QPrinter__LetterSmall        = QPrinter__PaperSize(55)
	QPrinter__TabloidExtra       = QPrinter__PaperSize(56)
	QPrinter__ArchA              = QPrinter__PaperSize(57)
	QPrinter__ArchB              = QPrinter__PaperSize(58)
	QPrinter__ArchC              = QPrinter__PaperSize(59)
	QPrinter__ArchD              = QPrinter__PaperSize(60)
	QPrinter__ArchE              = QPrinter__PaperSize(61)
	QPrinter__Imperial7x9        = QPrinter__PaperSize(62)
	QPrinter__Imperial8x10       = QPrinter__PaperSize(63)
	QPrinter__Imperial9x11       = QPrinter__PaperSize(64)
	QPrinter__Imperial9x12       = QPrinter__PaperSize(65)
	QPrinter__Imperial10x11      = QPrinter__PaperSize(66)
	QPrinter__Imperial10x13      = QPrinter__PaperSize(67)
	QPrinter__Imperial10x14      = QPrinter__PaperSize(68)
	QPrinter__Imperial12x11      = QPrinter__PaperSize(69)
	QPrinter__Imperial15x11      = QPrinter__PaperSize(70)
	QPrinter__ExecutiveStandard  = QPrinter__PaperSize(71)
	QPrinter__Note               = QPrinter__PaperSize(72)
	QPrinter__Quarto             = QPrinter__PaperSize(73)
	QPrinter__Statement          = QPrinter__PaperSize(74)
	QPrinter__SuperA             = QPrinter__PaperSize(75)
	QPrinter__SuperB             = QPrinter__PaperSize(76)
	QPrinter__Postcard           = QPrinter__PaperSize(77)
	QPrinter__DoublePostcard     = QPrinter__PaperSize(78)
	QPrinter__Prc16K             = QPrinter__PaperSize(79)
	QPrinter__Prc32K             = QPrinter__PaperSize(80)
	QPrinter__Prc32KBig          = QPrinter__PaperSize(81)
	QPrinter__FanFoldUS          = QPrinter__PaperSize(82)
	QPrinter__FanFoldGerman      = QPrinter__PaperSize(83)
	QPrinter__FanFoldGermanLegal = QPrinter__PaperSize(84)
	QPrinter__EnvelopeB4         = QPrinter__PaperSize(85)
	QPrinter__EnvelopeB5         = QPrinter__PaperSize(86)
	QPrinter__EnvelopeB6         = QPrinter__PaperSize(87)
	QPrinter__EnvelopeC0         = QPrinter__PaperSize(88)
	QPrinter__EnvelopeC1         = QPrinter__PaperSize(89)
	QPrinter__EnvelopeC2         = QPrinter__PaperSize(90)
	QPrinter__EnvelopeC3         = QPrinter__PaperSize(91)
	QPrinter__EnvelopeC4         = QPrinter__PaperSize(92)
	QPrinter__EnvelopeC6         = QPrinter__PaperSize(93)
	QPrinter__EnvelopeC65        = QPrinter__PaperSize(94)
	QPrinter__EnvelopeC7         = QPrinter__PaperSize(95)
	QPrinter__Envelope9          = QPrinter__PaperSize(96)
	QPrinter__Envelope11         = QPrinter__PaperSize(97)
	QPrinter__Envelope12         = QPrinter__PaperSize(98)
	QPrinter__Envelope14         = QPrinter__PaperSize(99)
	QPrinter__EnvelopeMonarch    = QPrinter__PaperSize(100)
	QPrinter__EnvelopePersonal   = QPrinter__PaperSize(101)
	QPrinter__EnvelopeChou3      = QPrinter__PaperSize(102)
	QPrinter__EnvelopeChou4      = QPrinter__PaperSize(103)
	QPrinter__EnvelopeInvite     = QPrinter__PaperSize(104)
	QPrinter__EnvelopeItalian    = QPrinter__PaperSize(105)
	QPrinter__EnvelopeKaku2      = QPrinter__PaperSize(106)
	QPrinter__EnvelopeKaku3      = QPrinter__PaperSize(107)
	QPrinter__EnvelopePrc1       = QPrinter__PaperSize(108)
	QPrinter__EnvelopePrc2       = QPrinter__PaperSize(109)
	QPrinter__EnvelopePrc3       = QPrinter__PaperSize(110)
	QPrinter__EnvelopePrc4       = QPrinter__PaperSize(111)
	QPrinter__EnvelopePrc5       = QPrinter__PaperSize(112)
	QPrinter__EnvelopePrc6       = QPrinter__PaperSize(113)
	QPrinter__EnvelopePrc7       = QPrinter__PaperSize(114)
	QPrinter__EnvelopePrc8       = QPrinter__PaperSize(115)
	QPrinter__EnvelopePrc9       = QPrinter__PaperSize(116)
	QPrinter__EnvelopePrc10      = QPrinter__PaperSize(117)
	QPrinter__EnvelopeYou4       = QPrinter__PaperSize(118)
	QPrinter__LastPageSize       = QPrinter__PaperSize(QPrinter__EnvelopeYou4)
	QPrinter__NPageSize          = QPrinter__PaperSize(QPrinter__LastPageSize)
	QPrinter__NPaperSize         = QPrinter__PaperSize(QPrinter__LastPageSize)
	QPrinter__AnsiA              = QPrinter__PaperSize(QPrinter__Letter)
	QPrinter__AnsiB              = QPrinter__PaperSize(QPrinter__Ledger)
	QPrinter__EnvelopeC5         = QPrinter__PaperSize(QPrinter__C5E)
	QPrinter__EnvelopeDL         = QPrinter__PaperSize(QPrinter__DLE)
	QPrinter__Envelope10         = QPrinter__PaperSize(QPrinter__Comm10E)
)

//QPrinter::PaperSource
type QPrinter__PaperSource int64

const (
	QPrinter__OnlyOne         = QPrinter__PaperSource(0)
	QPrinter__Lower           = QPrinter__PaperSource(1)
	QPrinter__Middle          = QPrinter__PaperSource(2)
	QPrinter__Manual          = QPrinter__PaperSource(3)
	QPrinter__Envelope        = QPrinter__PaperSource(4)
	QPrinter__EnvelopeManual  = QPrinter__PaperSource(5)
	QPrinter__Auto            = QPrinter__PaperSource(6)
	QPrinter__Tractor         = QPrinter__PaperSource(7)
	QPrinter__SmallFormat     = QPrinter__PaperSource(8)
	QPrinter__LargeFormat     = QPrinter__PaperSource(9)
	QPrinter__LargeCapacity   = QPrinter__PaperSource(10)
	QPrinter__Cassette        = QPrinter__PaperSource(11)
	QPrinter__FormSource      = QPrinter__PaperSource(12)
	QPrinter__MaxPageSource   = QPrinter__PaperSource(13)
	QPrinter__CustomSource    = QPrinter__PaperSource(14)
	QPrinter__LastPaperSource = QPrinter__PaperSource(QPrinter__CustomSource)
	QPrinter__Upper           = QPrinter__PaperSource(QPrinter__OnlyOne)
)

//QPrinter::PrintRange
type QPrinter__PrintRange int64

const (
	QPrinter__AllPages    = QPrinter__PrintRange(0)
	QPrinter__Selection   = QPrinter__PrintRange(1)
	QPrinter__PageRange   = QPrinter__PrintRange(2)
	QPrinter__CurrentPage = QPrinter__PrintRange(3)
)

//QPrinter::PrinterMode
type QPrinter__PrinterMode int64

const (
	QPrinter__ScreenResolution  = QPrinter__PrinterMode(0)
	QPrinter__PrinterResolution = QPrinter__PrinterMode(1)
	QPrinter__HighResolution    = QPrinter__PrinterMode(2)
)

//QPrinter::PrinterState
type QPrinter__PrinterState int64

const (
	QPrinter__Idle    = QPrinter__PrinterState(0)
	QPrinter__Active  = QPrinter__PrinterState(1)
	QPrinter__Aborted = QPrinter__PrinterState(2)
	QPrinter__Error   = QPrinter__PrinterState(3)
)

//QPrinter::Unit
type QPrinter__Unit int64

const (
	QPrinter__Millimeter  = QPrinter__Unit(0)
	QPrinter__Point       = QPrinter__Unit(1)
	QPrinter__Inch        = QPrinter__Unit(2)
	QPrinter__Pica        = QPrinter__Unit(3)
	QPrinter__Didot       = QPrinter__Unit(4)
	QPrinter__Cicero      = QPrinter__Unit(5)
	QPrinter__DevicePixel = QPrinter__Unit(6)
)

type QPrinter struct {
	gui.QPagedPaintDevice
}

type QPrinter_ITF interface {
	gui.QPagedPaintDevice_ITF
	QPrinter_PTR() *QPrinter
}

func (p *QPrinter) QPrinter_PTR() *QPrinter {
	return p
}

func (p *QPrinter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func (p *QPrinter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QPagedPaintDevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQPrinter(ptr QPrinter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinter_PTR().Pointer()
	}
	return nil
}

func NewQPrinterFromPointer(ptr unsafe.Pointer) *QPrinter {
	var n = new(QPrinter)
	n.SetPointer(ptr)
	return n
}
func (ptr *QPrinter) FromPage() int {
	defer qt.Recovering("QPrinter::fromPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_FromPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) OutputFileName() string {
	defer qt.Recovering("QPrinter::outputFileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_OutputFileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrinterSelectionOption() string {
	defer qt.Recovering("QPrinter::printerSelectionOption")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_PrinterSelectionOption(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) SetPrinterSelectionOption(option string) {
	defer qt.Recovering("QPrinter::setPrinterSelectionOption")

	if ptr.Pointer() != nil {
		var optionC = C.CString(option)
		defer C.free(unsafe.Pointer(optionC))
		C.QPrinter_SetPrinterSelectionOption(ptr.Pointer(), optionC)
	}
}

func NewQPrinter(mode QPrinter__PrinterMode) *QPrinter {
	defer qt.Recovering("QPrinter::QPrinter")

	var tmpValue = NewQPrinterFromPointer(C.QPrinter_NewQPrinter(C.longlong(mode)))
	runtime.SetFinalizer(tmpValue, (*QPrinter).DestroyQPrinter)
	return tmpValue
}

func NewQPrinter2(printer QPrinterInfo_ITF, mode QPrinter__PrinterMode) *QPrinter {
	defer qt.Recovering("QPrinter::QPrinter")

	var tmpValue = NewQPrinterFromPointer(C.QPrinter_NewQPrinter2(PointerFromQPrinterInfo(printer), C.longlong(mode)))
	runtime.SetFinalizer(tmpValue, (*QPrinter).DestroyQPrinter)
	return tmpValue
}

func (ptr *QPrinter) Abort() bool {
	defer qt.Recovering("QPrinter::abort")

	if ptr.Pointer() != nil {
		return C.QPrinter_Abort(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) CollateCopies() bool {
	defer qt.Recovering("QPrinter::collateCopies")

	if ptr.Pointer() != nil {
		return C.QPrinter_CollateCopies(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) ColorMode() QPrinter__ColorMode {
	defer qt.Recovering("QPrinter::colorMode")

	if ptr.Pointer() != nil {
		return QPrinter__ColorMode(C.QPrinter_ColorMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) CopyCount() int {
	defer qt.Recovering("QPrinter::copyCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_CopyCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) Creator() string {
	defer qt.Recovering("QPrinter::creator")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_Creator(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) DocName() string {
	defer qt.Recovering("QPrinter::docName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_DocName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) Duplex() QPrinter__DuplexMode {
	defer qt.Recovering("QPrinter::duplex")

	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinter_Duplex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) FontEmbeddingEnabled() bool {
	defer qt.Recovering("QPrinter::fontEmbeddingEnabled")

	if ptr.Pointer() != nil {
		return C.QPrinter_FontEmbeddingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) FullPage() bool {
	defer qt.Recovering("QPrinter::fullPage")

	if ptr.Pointer() != nil {
		return C.QPrinter_FullPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) IsValid() bool {
	defer qt.Recovering("QPrinter::isValid")

	if ptr.Pointer() != nil {
		return C.QPrinter_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) NewPage() bool {
	defer qt.Recovering("QPrinter::newPage")

	if ptr.Pointer() != nil {
		return C.QPrinter_NewPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) OutputFormat() QPrinter__OutputFormat {
	defer qt.Recovering("QPrinter::outputFormat")

	if ptr.Pointer() != nil {
		return QPrinter__OutputFormat(C.QPrinter_OutputFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PageLayout() *gui.QPageLayout {
	defer qt.Recovering("QPrinter::pageLayout")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPageLayoutFromPointer(C.QPrinter_PageLayout(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageLayout).DestroyQPageLayout)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinter) PageOrder() QPrinter__PageOrder {
	defer qt.Recovering("QPrinter::pageOrder")

	if ptr.Pointer() != nil {
		return QPrinter__PageOrder(C.QPrinter_PageOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PageRect(unit QPrinter__Unit) *core.QRectF {
	defer qt.Recovering("QPrinter::pageRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QPrinter_PageRect(ptr.Pointer(), C.longlong(unit)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinter) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QPrinter::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrinter_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrinter) PaperRect(unit QPrinter__Unit) *core.QRectF {
	defer qt.Recovering("QPrinter::paperRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QPrinter_PaperRect(ptr.Pointer(), C.longlong(unit)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinter) PaperSource() QPrinter__PaperSource {
	defer qt.Recovering("QPrinter::paperSource")

	if ptr.Pointer() != nil {
		return QPrinter__PaperSource(C.QPrinter_PaperSource(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PrintEngine() *QPrintEngine {
	defer qt.Recovering("QPrinter::printEngine")

	if ptr.Pointer() != nil {
		return NewQPrintEngineFromPointer(C.QPrinter_PrintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrinter) PrintProgram() string {
	defer qt.Recovering("QPrinter::printProgram")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_PrintProgram(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrintRange() QPrinter__PrintRange {
	defer qt.Recovering("QPrinter::printRange")

	if ptr.Pointer() != nil {
		return QPrinter__PrintRange(C.QPrinter_PrintRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PrinterName() string {
	defer qt.Recovering("QPrinter::printerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinter_PrinterName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrinterState() QPrinter__PrinterState {
	defer qt.Recovering("QPrinter::printerState")

	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrinter_PrinterState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) Resolution() int {
	defer qt.Recovering("QPrinter::resolution")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_Resolution(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) SetCollateCopies(collate bool) {
	defer qt.Recovering("QPrinter::setCollateCopies")

	if ptr.Pointer() != nil {
		C.QPrinter_SetCollateCopies(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(collate))))
	}
}

func (ptr *QPrinter) SetColorMode(newColorMode QPrinter__ColorMode) {
	defer qt.Recovering("QPrinter::setColorMode")

	if ptr.Pointer() != nil {
		C.QPrinter_SetColorMode(ptr.Pointer(), C.longlong(newColorMode))
	}
}

func (ptr *QPrinter) SetCopyCount(count int) {
	defer qt.Recovering("QPrinter::setCopyCount")

	if ptr.Pointer() != nil {
		C.QPrinter_SetCopyCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QPrinter) SetCreator(creator string) {
	defer qt.Recovering("QPrinter::setCreator")

	if ptr.Pointer() != nil {
		var creatorC = C.CString(creator)
		defer C.free(unsafe.Pointer(creatorC))
		C.QPrinter_SetCreator(ptr.Pointer(), creatorC)
	}
}

func (ptr *QPrinter) SetDocName(name string) {
	defer qt.Recovering("QPrinter::setDocName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QPrinter_SetDocName(ptr.Pointer(), nameC)
	}
}

func (ptr *QPrinter) SetDuplex(duplex QPrinter__DuplexMode) {
	defer qt.Recovering("QPrinter::setDuplex")

	if ptr.Pointer() != nil {
		C.QPrinter_SetDuplex(ptr.Pointer(), C.longlong(duplex))
	}
}

func (ptr *QPrinter) SetEngines(printEngine QPrintEngine_ITF, paintEngine gui.QPaintEngine_ITF) {
	defer qt.Recovering("QPrinter::setEngines")

	if ptr.Pointer() != nil {
		C.QPrinter_SetEngines(ptr.Pointer(), PointerFromQPrintEngine(printEngine), gui.PointerFromQPaintEngine(paintEngine))
	}
}

func (ptr *QPrinter) SetFontEmbeddingEnabled(enable bool) {
	defer qt.Recovering("QPrinter::setFontEmbeddingEnabled")

	if ptr.Pointer() != nil {
		C.QPrinter_SetFontEmbeddingEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QPrinter) SetFromTo(from int, to int) {
	defer qt.Recovering("QPrinter::setFromTo")

	if ptr.Pointer() != nil {
		C.QPrinter_SetFromTo(ptr.Pointer(), C.int(int32(from)), C.int(int32(to)))
	}
}

func (ptr *QPrinter) SetFullPage(fp bool) {
	defer qt.Recovering("QPrinter::setFullPage")

	if ptr.Pointer() != nil {
		C.QPrinter_SetFullPage(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fp))))
	}
}

func (ptr *QPrinter) SetOutputFileName(fileName string) {
	defer qt.Recovering("QPrinter::setOutputFileName")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QPrinter_SetOutputFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QPrinter) SetOutputFormat(format QPrinter__OutputFormat) {
	defer qt.Recovering("QPrinter::setOutputFormat")

	if ptr.Pointer() != nil {
		C.QPrinter_SetOutputFormat(ptr.Pointer(), C.longlong(format))
	}
}

func (ptr *QPrinter) SetPageLayout(newLayout gui.QPageLayout_ITF) bool {
	defer qt.Recovering("QPrinter::setPageLayout")

	if ptr.Pointer() != nil {
		return C.QPrinter_SetPageLayout(ptr.Pointer(), gui.PointerFromQPageLayout(newLayout)) != 0
	}
	return false
}

func (ptr *QPrinter) SetPageMargins(margins core.QMarginsF_ITF) bool {
	defer qt.Recovering("QPrinter::setPageMargins")

	if ptr.Pointer() != nil {
		return C.QPrinter_SetPageMargins(ptr.Pointer(), core.PointerFromQMarginsF(margins)) != 0
	}
	return false
}

func (ptr *QPrinter) SetPageMargins2(margins core.QMarginsF_ITF, units gui.QPageLayout__Unit) bool {
	defer qt.Recovering("QPrinter::setPageMargins")

	if ptr.Pointer() != nil {
		return C.QPrinter_SetPageMargins2(ptr.Pointer(), core.PointerFromQMarginsF(margins), C.longlong(units)) != 0
	}
	return false
}

func (ptr *QPrinter) SetPageOrder(pageOrder QPrinter__PageOrder) {
	defer qt.Recovering("QPrinter::setPageOrder")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPageOrder(ptr.Pointer(), C.longlong(pageOrder))
	}
}

func (ptr *QPrinter) SetPageOrientation(orientation gui.QPageLayout__Orientation) bool {
	defer qt.Recovering("QPrinter::setPageOrientation")

	if ptr.Pointer() != nil {
		return C.QPrinter_SetPageOrientation(ptr.Pointer(), C.longlong(orientation)) != 0
	}
	return false
}

func (ptr *QPrinter) SetPageSize(pageSize gui.QPageSize_ITF) bool {
	defer qt.Recovering("QPrinter::setPageSize")

	if ptr.Pointer() != nil {
		return C.QPrinter_SetPageSize(ptr.Pointer(), gui.PointerFromQPageSize(pageSize)) != 0
	}
	return false
}

func (ptr *QPrinter) SetPaperSource(source QPrinter__PaperSource) {
	defer qt.Recovering("QPrinter::setPaperSource")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPaperSource(ptr.Pointer(), C.longlong(source))
	}
}

func (ptr *QPrinter) SetPrintProgram(printProg string) {
	defer qt.Recovering("QPrinter::setPrintProgram")

	if ptr.Pointer() != nil {
		var printProgC = C.CString(printProg)
		defer C.free(unsafe.Pointer(printProgC))
		C.QPrinter_SetPrintProgram(ptr.Pointer(), printProgC)
	}
}

func (ptr *QPrinter) SetPrintRange(ran QPrinter__PrintRange) {
	defer qt.Recovering("QPrinter::setPrintRange")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPrintRange(ptr.Pointer(), C.longlong(ran))
	}
}

func (ptr *QPrinter) SetPrinterName(name string) {
	defer qt.Recovering("QPrinter::setPrinterName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QPrinter_SetPrinterName(ptr.Pointer(), nameC)
	}
}

func (ptr *QPrinter) SetResolution(dpi int) {
	defer qt.Recovering("QPrinter::setResolution")

	if ptr.Pointer() != nil {
		C.QPrinter_SetResolution(ptr.Pointer(), C.int(int32(dpi)))
	}
}

func (ptr *QPrinter) SupportsMultipleCopies() bool {
	defer qt.Recovering("QPrinter::supportsMultipleCopies")

	if ptr.Pointer() != nil {
		return C.QPrinter_SupportsMultipleCopies(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinter) ToPage() int {
	defer qt.Recovering("QPrinter::toPage")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_ToPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) DestroyQPrinter() {
	defer qt.Recovering("QPrinter::~QPrinter")

	if ptr.Pointer() != nil {
		C.QPrinter_DestroyQPrinter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPrinter_SetPageSize2
func callbackQPrinter_SetPageSize2(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QPrinter::setPageSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrinter::setPageSize2"); signal != nil {
		signal.(func(gui.QPagedPaintDevice__PageSize))(gui.QPagedPaintDevice__PageSize(size))
	} else {
		NewQPrinterFromPointer(ptr).SetPageSize2Default(gui.QPagedPaintDevice__PageSize(size))
	}
}

func (ptr *QPrinter) ConnectSetPageSize2(f func(size gui.QPagedPaintDevice__PageSize)) {
	defer qt.Recovering("connect QPrinter::setPageSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::setPageSize2", f)
	}
}

func (ptr *QPrinter) DisconnectSetPageSize2() {
	defer qt.Recovering("disconnect QPrinter::setPageSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::setPageSize2")
	}
}

func (ptr *QPrinter) SetPageSize2(size gui.QPagedPaintDevice__PageSize) {
	defer qt.Recovering("QPrinter::setPageSize")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSize2(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QPrinter) SetPageSize2Default(size gui.QPagedPaintDevice__PageSize) {
	defer qt.Recovering("QPrinter::setPageSize")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSize2Default(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQPrinter_SetPageSizeMM
func callbackQPrinter_SetPageSizeMM(ptr unsafe.Pointer, size unsafe.Pointer) {
	defer qt.Recovering("callback QPrinter::setPageSizeMM")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrinter::setPageSizeMM"); signal != nil {
		signal.(func(*core.QSizeF))(core.NewQSizeFFromPointer(size))
	} else {
		NewQPrinterFromPointer(ptr).SetPageSizeMMDefault(core.NewQSizeFFromPointer(size))
	}
}

func (ptr *QPrinter) ConnectSetPageSizeMM(f func(size *core.QSizeF)) {
	defer qt.Recovering("connect QPrinter::setPageSizeMM")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::setPageSizeMM", f)
	}
}

func (ptr *QPrinter) DisconnectSetPageSizeMM() {
	defer qt.Recovering("disconnect QPrinter::setPageSizeMM")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::setPageSizeMM")
	}
}

func (ptr *QPrinter) SetPageSizeMM(size core.QSizeF_ITF) {
	defer qt.Recovering("QPrinter::setPageSizeMM")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSizeMM(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QPrinter) SetPageSizeMMDefault(size core.QSizeF_ITF) {
	defer qt.Recovering("QPrinter::setPageSizeMM")

	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSizeMMDefault(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

//export callbackQPrinter_Metric
func callbackQPrinter_Metric(ptr unsafe.Pointer, metric C.longlong) C.int {
	defer qt.Recovering("callback QPrinter::metric")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPrinter::metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(metric))))
	}

	return C.int(int32(NewQPrinterFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(metric))))
}

func (ptr *QPrinter) ConnectMetric(f func(metric gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QPrinter::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::metric", f)
	}
}

func (ptr *QPrinter) DisconnectMetric() {
	defer qt.Recovering("disconnect QPrinter::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPrinter::metric")
	}
}

func (ptr *QPrinter) Metric(metric gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QPrinter::metric")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_Metric(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

func (ptr *QPrinter) MetricDefault(metric gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QPrinter::metric")

	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_MetricDefault(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

type QPrinterInfo struct {
	ptr unsafe.Pointer
}

type QPrinterInfo_ITF interface {
	QPrinterInfo_PTR() *QPrinterInfo
}

func (p *QPrinterInfo) QPrinterInfo_PTR() *QPrinterInfo {
	return p
}

func (p *QPrinterInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QPrinterInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQPrinterInfo(ptr QPrinterInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinterInfo_PTR().Pointer()
	}
	return nil
}

func NewQPrinterInfoFromPointer(ptr unsafe.Pointer) *QPrinterInfo {
	var n = new(QPrinterInfo)
	n.SetPointer(ptr)
	return n
}
func NewQPrinterInfo() *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::QPrinterInfo")

	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo3(printer QPrinter_ITF) *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::QPrinterInfo")

	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo3(PointerFromQPrinter(printer)))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo2(other QPrinterInfo_ITF) *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::QPrinterInfo")

	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo2(PointerFromQPrinterInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func QPrinterInfo_AvailablePrinterNames() []string {
	defer qt.Recovering("QPrinterInfo::availablePrinterNames")

	return strings.Split(C.GoString(C.QPrinterInfo_QPrinterInfo_AvailablePrinterNames()), "|")
}

func (ptr *QPrinterInfo) AvailablePrinterNames() []string {
	defer qt.Recovering("QPrinterInfo::availablePrinterNames")

	return strings.Split(C.GoString(C.QPrinterInfo_QPrinterInfo_AvailablePrinterNames()), "|")
}

func (ptr *QPrinterInfo) DefaultDuplexMode() QPrinter__DuplexMode {
	defer qt.Recovering("QPrinterInfo::defaultDuplexMode")

	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinterInfo_DefaultDuplexMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinterInfo) DefaultPageSize() *gui.QPageSize {
	defer qt.Recovering("QPrinterInfo::defaultPageSize")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPageSizeFromPointer(C.QPrinterInfo_DefaultPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func QPrinterInfo_DefaultPrinter() *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::defaultPrinter")

	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_DefaultPrinter())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) DefaultPrinter() *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::defaultPrinter")

	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_DefaultPrinter())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func QPrinterInfo_DefaultPrinterName() string {
	defer qt.Recovering("QPrinterInfo::defaultPrinterName")

	return C.GoString(C.QPrinterInfo_QPrinterInfo_DefaultPrinterName())
}

func (ptr *QPrinterInfo) DefaultPrinterName() string {
	defer qt.Recovering("QPrinterInfo::defaultPrinterName")

	return C.GoString(C.QPrinterInfo_QPrinterInfo_DefaultPrinterName())
}

func (ptr *QPrinterInfo) Description() string {
	defer qt.Recovering("QPrinterInfo::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinterInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) IsDefault() bool {
	defer qt.Recovering("QPrinterInfo::isDefault")

	if ptr.Pointer() != nil {
		return C.QPrinterInfo_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinterInfo) IsNull() bool {
	defer qt.Recovering("QPrinterInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QPrinterInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinterInfo) IsRemote() bool {
	defer qt.Recovering("QPrinterInfo::isRemote")

	if ptr.Pointer() != nil {
		return C.QPrinterInfo_IsRemote(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinterInfo) Location() string {
	defer qt.Recovering("QPrinterInfo::location")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinterInfo_Location(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) MakeAndModel() string {
	defer qt.Recovering("QPrinterInfo::makeAndModel")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinterInfo_MakeAndModel(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) MaximumPhysicalPageSize() *gui.QPageSize {
	defer qt.Recovering("QPrinterInfo::maximumPhysicalPageSize")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPageSizeFromPointer(C.QPrinterInfo_MaximumPhysicalPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) MinimumPhysicalPageSize() *gui.QPageSize {
	defer qt.Recovering("QPrinterInfo::minimumPhysicalPageSize")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPageSizeFromPointer(C.QPrinterInfo_MinimumPhysicalPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func QPrinterInfo_PrinterInfo(printerName string) *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::printerInfo")

	var printerNameC = C.CString(printerName)
	defer C.free(unsafe.Pointer(printerNameC))
	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_PrinterInfo(printerNameC))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) PrinterInfo(printerName string) *QPrinterInfo {
	defer qt.Recovering("QPrinterInfo::printerInfo")

	var printerNameC = C.CString(printerName)
	defer C.free(unsafe.Pointer(printerNameC))
	var tmpValue = NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_PrinterInfo(printerNameC))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) PrinterName() string {
	defer qt.Recovering("QPrinterInfo::printerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPrinterInfo_PrinterName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) State() QPrinter__PrinterState {
	defer qt.Recovering("QPrinterInfo::state")

	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrinterInfo_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinterInfo) SupportsCustomPageSizes() bool {
	defer qt.Recovering("QPrinterInfo::supportsCustomPageSizes")

	if ptr.Pointer() != nil {
		return C.QPrinterInfo_SupportsCustomPageSizes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPrinterInfo) DestroyQPrinterInfo() {
	defer qt.Recovering("QPrinterInfo::~QPrinterInfo")

	if ptr.Pointer() != nil {
		C.QPrinterInfo_DestroyQPrinterInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
