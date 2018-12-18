// +build !minimal

package printsupport

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "printsupport.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtPrintSupport_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtPrintSupport_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QAbstractPrintDialog struct {
	widgets.QDialog
}

type QAbstractPrintDialog_ITF interface {
	widgets.QDialog_ITF
	QAbstractPrintDialog_PTR() *QAbstractPrintDialog
}

func (ptr *QAbstractPrintDialog) QAbstractPrintDialog_PTR() *QAbstractPrintDialog {
	return ptr
}

func (ptr *QAbstractPrintDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractPrintDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractPrintDialog(ptr QAbstractPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func NewQAbstractPrintDialogFromPointer(ptr unsafe.Pointer) (n *QAbstractPrintDialog) {
	n = new(QAbstractPrintDialog)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractPrintDialog__PrintDialogOption
//QAbstractPrintDialog::PrintDialogOption
type QAbstractPrintDialog__PrintDialogOption int64

const (
	QAbstractPrintDialog__None               QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0000)
	QAbstractPrintDialog__PrintToFile        QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0001)
	QAbstractPrintDialog__PrintSelection     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0002)
	QAbstractPrintDialog__PrintPageRange     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0004)
	QAbstractPrintDialog__PrintShowPageSize  QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0008)
	QAbstractPrintDialog__PrintCollateCopies QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0010)
	QAbstractPrintDialog__DontUseSheet       QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0020)
	QAbstractPrintDialog__PrintCurrentPage   QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0040)
)

//go:generate stringer -type=QAbstractPrintDialog__PrintRange
//QAbstractPrintDialog::PrintRange
type QAbstractPrintDialog__PrintRange int64

const (
	QAbstractPrintDialog__AllPages    QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(0)
	QAbstractPrintDialog__Selection   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(1)
	QAbstractPrintDialog__PageRange   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(2)
	QAbstractPrintDialog__CurrentPage QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(3)
)

func NewQAbstractPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QAbstractPrintDialog {
	tmpValue := NewQAbstractPrintDialogFromPointer(C.QAbstractPrintDialog_NewQAbstractPrintDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QAbstractPrintDialog_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractPrintDialog_QAbstractPrintDialog_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractPrintDialog) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractPrintDialog_QAbstractPrintDialog_Tr(sC, cC, C.int(int32(n))))
}

func QAbstractPrintDialog_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractPrintDialog_QAbstractPrintDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractPrintDialog) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractPrintDialog_QAbstractPrintDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQAbstractPrintDialog_Exec
func callbackQAbstractPrintDialog_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractPrintDialog) ConnectExec(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exec"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exec", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exec", f)
		}
	}
}

func (ptr *QAbstractPrintDialog) DisconnectExec() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exec")
	}
}

func (ptr *QAbstractPrintDialog) Exec() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) SetFromTo(from int, to int) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetFromTo(ptr.Pointer(), C.int(int32(from)), C.int(int32(to)))
	}
}

func (ptr *QAbstractPrintDialog) SetMinMax(min int, max int) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetMinMax(ptr.Pointer(), C.int(int32(min)), C.int(int32(max)))
	}
}

func (ptr *QAbstractPrintDialog) SetOptionTabs(tabs []*widgets.QWidget) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetOptionTabs(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQAbstractPrintDialogFromPointer(NewQAbstractPrintDialogFromPointer(nil).__setOptionTabs_tabs_newList())
			for _, v := range tabs {
				tmpList.__setOptionTabs_tabs_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QAbstractPrintDialog) SetPrintRange(ran QAbstractPrintDialog__PrintRange) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetPrintRange(ptr.Pointer(), C.longlong(ran))
	}
}

func (ptr *QAbstractPrintDialog) PrintRange() QAbstractPrintDialog__PrintRange {
	if ptr.Pointer() != nil {
		return QAbstractPrintDialog__PrintRange(C.QAbstractPrintDialog_PrintRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) Printer() *QPrinter {
	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QAbstractPrintDialog_Printer(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractPrintDialog_MetaObject
func callbackQAbstractPrintDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractPrintDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractPrintDialog) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractPrintDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractPrintDialog) FromPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_FromPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) MaxPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_MaxPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) MinPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_MinPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) ToPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_ToPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QAbstractPrintDialog___setOptionTabs_tabs_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___setOptionTabs_tabs_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___setOptionTabs_tabs_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QAbstractPrintDialog___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QAbstractPrintDialog) __addActions_actions_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QAbstractPrintDialog___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QAbstractPrintDialog___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QAbstractPrintDialog) __actions_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___actions_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractPrintDialog___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractPrintDialog___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractPrintDialog) __findChildren_newList2() unsafe.Pointer {
	return C.QAbstractPrintDialog___findChildren_newList2(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractPrintDialog___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractPrintDialog) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractPrintDialog___findChildren_newList3(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractPrintDialog___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractPrintDialog) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractPrintDialog) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractPrintDialog___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractPrintDialog) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractPrintDialog) __children_newList() unsafe.Pointer {
	return C.QAbstractPrintDialog___children_newList(ptr.Pointer())
}

//export callbackQAbstractPrintDialog_EventFilter
func callbackQAbstractPrintDialog_EventFilter(ptr unsafe.Pointer, o unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractPrintDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractPrintDialog_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_Accept
func callbackQAbstractPrintDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QAbstractPrintDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Accepted
func callbackQAbstractPrintDialog_Accepted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		signal.(func())()
	}

}

//export callbackQAbstractPrintDialog_CloseEvent
func callbackQAbstractPrintDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQAbstractPrintDialog_ContextMenuEvent
func callbackQAbstractPrintDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQAbstractPrintDialog_Done
func callbackQAbstractPrintDialog_Done(ptr unsafe.Pointer, r C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		signal.(func(int))(int(int32(r)))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DoneDefault(int(int32(r)))
	}
}

func (ptr *QAbstractPrintDialog) DoneDefault(r int) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DoneDefault(ptr.Pointer(), C.int(int32(r)))
	}
}

//export callbackQAbstractPrintDialog_Finished
func callbackQAbstractPrintDialog_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(int))(int(int32(result)))
	}

}

//export callbackQAbstractPrintDialog_KeyPressEvent
func callbackQAbstractPrintDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QAbstractPrintDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQAbstractPrintDialog_Open
func callbackQAbstractPrintDialog_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).OpenDefault()
	}
}

func (ptr *QAbstractPrintDialog) OpenDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_OpenDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Reject
func callbackQAbstractPrintDialog_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QAbstractPrintDialog) RejectDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Rejected
func callbackQAbstractPrintDialog_Rejected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rejected"); signal != nil {
		signal.(func())()
	}

}

//export callbackQAbstractPrintDialog_ResizeEvent
func callbackQAbstractPrintDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QAbstractPrintDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQAbstractPrintDialog_SetVisible
func callbackQAbstractPrintDialog_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QAbstractPrintDialog) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractPrintDialog_ShowEvent
func callbackQAbstractPrintDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MinimumSizeHint
func callbackQAbstractPrintDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQAbstractPrintDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QAbstractPrintDialog) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QAbstractPrintDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_SizeHint
func callbackQAbstractPrintDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQAbstractPrintDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QAbstractPrintDialog) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QAbstractPrintDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_Close
func callbackQAbstractPrintDialog_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QAbstractPrintDialog) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractPrintDialog_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_Event
func callbackQAbstractPrintDialog_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractPrintDialog) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractPrintDialog_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_FocusNextPrevChild
func callbackQAbstractPrintDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QAbstractPrintDialog) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractPrintDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_ActionEvent
func callbackQAbstractPrintDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQAbstractPrintDialog_ChangeEvent
func callbackQAbstractPrintDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_CustomContextMenuRequested
func callbackQAbstractPrintDialog_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQAbstractPrintDialog_DragEnterEvent
func callbackQAbstractPrintDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DragLeaveEvent
func callbackQAbstractPrintDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DragMoveEvent
func callbackQAbstractPrintDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DropEvent
func callbackQAbstractPrintDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQAbstractPrintDialog_EnterEvent
func callbackQAbstractPrintDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_FocusInEvent
func callbackQAbstractPrintDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQAbstractPrintDialog_FocusOutEvent
func callbackQAbstractPrintDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Hide
func callbackQAbstractPrintDialog_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QAbstractPrintDialog) HideDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_HideEvent
func callbackQAbstractPrintDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQAbstractPrintDialog_InputMethodEvent
func callbackQAbstractPrintDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQAbstractPrintDialog_KeyReleaseEvent
func callbackQAbstractPrintDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQAbstractPrintDialog_LeaveEvent
func callbackQAbstractPrintDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Lower
func callbackQAbstractPrintDialog_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QAbstractPrintDialog) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_MouseDoubleClickEvent
func callbackQAbstractPrintDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MouseMoveEvent
func callbackQAbstractPrintDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MousePressEvent
func callbackQAbstractPrintDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MouseReleaseEvent
func callbackQAbstractPrintDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstractPrintDialog_MoveEvent
func callbackQAbstractPrintDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQAbstractPrintDialog_PaintEvent
func callbackQAbstractPrintDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Raise
func callbackQAbstractPrintDialog_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QAbstractPrintDialog) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_Repaint
func callbackQAbstractPrintDialog_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QAbstractPrintDialog) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_SetDisabled
func callbackQAbstractPrintDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QAbstractPrintDialog) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQAbstractPrintDialog_SetEnabled
func callbackQAbstractPrintDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QAbstractPrintDialog) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAbstractPrintDialog_SetFocus2
func callbackQAbstractPrintDialog_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QAbstractPrintDialog) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_SetHidden
func callbackQAbstractPrintDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QAbstractPrintDialog) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQAbstractPrintDialog_SetStyleSheet
func callbackQAbstractPrintDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QAbstractPrintDialog) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QAbstractPrintDialog_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQAbstractPrintDialog_SetWindowModified
func callbackQAbstractPrintDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QAbstractPrintDialog) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAbstractPrintDialog_SetWindowTitle
func callbackQAbstractPrintDialog_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QAbstractPrintDialog) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QAbstractPrintDialog_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQAbstractPrintDialog_Show
func callbackQAbstractPrintDialog_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QAbstractPrintDialog) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowFullScreen
func callbackQAbstractPrintDialog_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QAbstractPrintDialog) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowMaximized
func callbackQAbstractPrintDialog_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QAbstractPrintDialog) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowMinimized
func callbackQAbstractPrintDialog_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QAbstractPrintDialog) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_ShowNormal
func callbackQAbstractPrintDialog_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QAbstractPrintDialog) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_TabletEvent
func callbackQAbstractPrintDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQAbstractPrintDialog_Update
func callbackQAbstractPrintDialog_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QAbstractPrintDialog) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_UpdateMicroFocus
func callbackQAbstractPrintDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QAbstractPrintDialog) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQAbstractPrintDialog_WheelEvent
func callbackQAbstractPrintDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQAbstractPrintDialog_WindowIconChanged
func callbackQAbstractPrintDialog_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQAbstractPrintDialog_WindowTitleChanged
func callbackQAbstractPrintDialog_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQAbstractPrintDialog_PaintEngine
func callbackQAbstractPrintDialog_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQAbstractPrintDialogFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QAbstractPrintDialog) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QAbstractPrintDialog_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractPrintDialog_InputMethodQuery
func callbackQAbstractPrintDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQAbstractPrintDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QAbstractPrintDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAbstractPrintDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractPrintDialog_HasHeightForWidth
func callbackQAbstractPrintDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractPrintDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QAbstractPrintDialog) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractPrintDialog_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractPrintDialog_HeightForWidth
func callbackQAbstractPrintDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQAbstractPrintDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QAbstractPrintDialog) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQAbstractPrintDialog_Metric
func callbackQAbstractPrintDialog_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQAbstractPrintDialogFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QAbstractPrintDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractPrintDialog_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQAbstractPrintDialog_InitPainter
func callbackQAbstractPrintDialog_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QAbstractPrintDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQAbstractPrintDialog_ChildEvent
func callbackQAbstractPrintDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractPrintDialog_ConnectNotify
func callbackQAbstractPrintDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractPrintDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractPrintDialog_CustomEvent
func callbackQAbstractPrintDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractPrintDialog_DeleteLater
func callbackQAbstractPrintDialog_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractPrintDialog) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAbstractPrintDialog_Destroyed
func callbackQAbstractPrintDialog_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractPrintDialog_DisconnectNotify
func callbackQAbstractPrintDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractPrintDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractPrintDialog_ObjectNameChanged
func callbackQAbstractPrintDialog_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractPrintDialog_TimerEvent
func callbackQAbstractPrintDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractPrintDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractPrintDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractPrintDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPageSetupDialog struct {
	widgets.QDialog
}

type QPageSetupDialog_ITF interface {
	widgets.QDialog_ITF
	QPageSetupDialog_PTR() *QPageSetupDialog
}

func (ptr *QPageSetupDialog) QPageSetupDialog_PTR() *QPageSetupDialog {
	return ptr
}

func (ptr *QPageSetupDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPageSetupDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPageSetupDialog(ptr QPageSetupDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageSetupDialog_PTR().Pointer()
	}
	return nil
}

func NewQPageSetupDialogFromPointer(ptr unsafe.Pointer) (n *QPageSetupDialog) {
	n = new(QPageSetupDialog)
	n.SetPointer(ptr)
	return
}
func NewQPageSetupDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPageSetupDialog {
	tmpValue := NewQPageSetupDialogFromPointer(C.QPageSetupDialog_NewQPageSetupDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPageSetupDialog2(parent widgets.QWidget_ITF) *QPageSetupDialog {
	tmpValue := NewQPageSetupDialogFromPointer(C.QPageSetupDialog_NewQPageSetupDialog2(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPageSetupDialog) Printer() *QPrinter {
	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QPageSetupDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func QPageSetupDialog_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPageSetupDialog_QPageSetupDialog_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QPageSetupDialog) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPageSetupDialog_QPageSetupDialog_Tr(sC, cC, C.int(int32(n))))
}

func QPageSetupDialog_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPageSetupDialog_QPageSetupDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QPageSetupDialog) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPageSetupDialog_QPageSetupDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQPageSetupDialog_Exec
func callbackQPageSetupDialog_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQPageSetupDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QPageSetupDialog) ConnectExec(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exec"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exec", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exec", f)
		}
	}
}

func (ptr *QPageSetupDialog) DisconnectExec() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exec")
	}
}

func (ptr *QPageSetupDialog) Exec() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageSetupDialog) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQPageSetupDialog_Done
func callbackQPageSetupDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		signal.(func(int))(int(int32(result)))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QPageSetupDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "done", func(result int) {
				signal.(func(int))(result)
				f(result)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", f)
		}
	}
}

func (ptr *QPageSetupDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QPageSetupDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPageSetupDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

//export callbackQPageSetupDialog_DestroyQPageSetupDialog
func callbackQPageSetupDialog_DestroyQPageSetupDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPageSetupDialog"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).DestroyQPageSetupDialogDefault()
	}
}

func (ptr *QPageSetupDialog) ConnectDestroyQPageSetupDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPageSetupDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPageSetupDialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPageSetupDialog", f)
		}
	}
}

func (ptr *QPageSetupDialog) DisconnectDestroyQPageSetupDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPageSetupDialog")
	}
}

func (ptr *QPageSetupDialog) DestroyQPageSetupDialog() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DestroyQPageSetupDialog(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPageSetupDialog) DestroyQPageSetupDialogDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DestroyQPageSetupDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPageSetupDialog_MetaObject
func callbackQPageSetupDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPageSetupDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPageSetupDialog) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPageSetupDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPageSetupDialog) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPageSetupDialog___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPageSetupDialog) __addActions_actions_newList() unsafe.Pointer {
	return C.QPageSetupDialog___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPageSetupDialog___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPageSetupDialog) __insertActions_actions_newList() unsafe.Pointer {
	return C.QPageSetupDialog___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPageSetupDialog___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPageSetupDialog) __actions_newList() unsafe.Pointer {
	return C.QPageSetupDialog___actions_newList(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QPageSetupDialog___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QPageSetupDialog___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPageSetupDialog___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPageSetupDialog) __findChildren_newList2() unsafe.Pointer {
	return C.QPageSetupDialog___findChildren_newList2(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPageSetupDialog___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPageSetupDialog) __findChildren_newList3() unsafe.Pointer {
	return C.QPageSetupDialog___findChildren_newList3(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPageSetupDialog___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPageSetupDialog) __findChildren_newList() unsafe.Pointer {
	return C.QPageSetupDialog___findChildren_newList(ptr.Pointer())
}

func (ptr *QPageSetupDialog) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPageSetupDialog___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPageSetupDialog) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPageSetupDialog) __children_newList() unsafe.Pointer {
	return C.QPageSetupDialog___children_newList(ptr.Pointer())
}

//export callbackQPageSetupDialog_EventFilter
func callbackQPageSetupDialog_EventFilter(ptr unsafe.Pointer, o unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
}

func (ptr *QPageSetupDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSetupDialog_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQPageSetupDialog_Accept
func callbackQPageSetupDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPageSetupDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_Accepted
func callbackQPageSetupDialog_Accepted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		signal.(func())()
	}

}

//export callbackQPageSetupDialog_CloseEvent
func callbackQPageSetupDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQPageSetupDialog_ContextMenuEvent
func callbackQPageSetupDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQPageSetupDialog_Finished
func callbackQPageSetupDialog_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(int))(int(int32(result)))
	}

}

//export callbackQPageSetupDialog_KeyPressEvent
func callbackQPageSetupDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPageSetupDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPageSetupDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQPageSetupDialog_Reject
func callbackQPageSetupDialog_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QPageSetupDialog) RejectDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_Rejected
func callbackQPageSetupDialog_Rejected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rejected"); signal != nil {
		signal.(func())()
	}

}

//export callbackQPageSetupDialog_ResizeEvent
func callbackQPageSetupDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QPageSetupDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQPageSetupDialog_SetVisible
func callbackQPageSetupDialog_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QPageSetupDialog) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQPageSetupDialog_ShowEvent
func callbackQPageSetupDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPageSetupDialog_MinimumSizeHint
func callbackQPageSetupDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPageSetupDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPageSetupDialog) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPageSetupDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_SizeHint
func callbackQPageSetupDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPageSetupDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPageSetupDialog) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPageSetupDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_Close
func callbackQPageSetupDialog_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QPageSetupDialog) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSetupDialog_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPageSetupDialog_Event
func callbackQPageSetupDialog_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QPageSetupDialog) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSetupDialog_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPageSetupDialog_FocusNextPrevChild
func callbackQPageSetupDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPageSetupDialog) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSetupDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQPageSetupDialog_ActionEvent
func callbackQPageSetupDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPageSetupDialog_ChangeEvent
func callbackQPageSetupDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_CustomContextMenuRequested
func callbackQPageSetupDialog_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQPageSetupDialog_DragEnterEvent
func callbackQPageSetupDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPageSetupDialog_DragLeaveEvent
func callbackQPageSetupDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPageSetupDialog_DragMoveEvent
func callbackQPageSetupDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPageSetupDialog_DropEvent
func callbackQPageSetupDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPageSetupDialog_EnterEvent
func callbackQPageSetupDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_FocusInEvent
func callbackQPageSetupDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPageSetupDialog_FocusOutEvent
func callbackQPageSetupDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPageSetupDialog_Hide
func callbackQPageSetupDialog_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPageSetupDialog) HideDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_HideEvent
func callbackQPageSetupDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPageSetupDialog_InputMethodEvent
func callbackQPageSetupDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPageSetupDialog_KeyReleaseEvent
func callbackQPageSetupDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPageSetupDialog_LeaveEvent
func callbackQPageSetupDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_Lower
func callbackQPageSetupDialog_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPageSetupDialog) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_MouseDoubleClickEvent
func callbackQPageSetupDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MouseMoveEvent
func callbackQPageSetupDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MousePressEvent
func callbackQPageSetupDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MouseReleaseEvent
func callbackQPageSetupDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPageSetupDialog_MoveEvent
func callbackQPageSetupDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPageSetupDialog_PaintEvent
func callbackQPageSetupDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPageSetupDialog_Raise
func callbackQPageSetupDialog_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPageSetupDialog) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_Repaint
func callbackQPageSetupDialog_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPageSetupDialog) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_SetDisabled
func callbackQPageSetupDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPageSetupDialog) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPageSetupDialog_SetEnabled
func callbackQPageSetupDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPageSetupDialog) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPageSetupDialog_SetFocus2
func callbackQPageSetupDialog_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPageSetupDialog) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_SetHidden
func callbackQPageSetupDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPageSetupDialog) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPageSetupDialog_SetStyleSheet
func callbackQPageSetupDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QPageSetupDialog) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QPageSetupDialog_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQPageSetupDialog_SetWindowModified
func callbackQPageSetupDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPageSetupDialog) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPageSetupDialog_SetWindowTitle
func callbackQPageSetupDialog_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQPageSetupDialogFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QPageSetupDialog) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QPageSetupDialog_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQPageSetupDialog_Show
func callbackQPageSetupDialog_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPageSetupDialog) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowFullScreen
func callbackQPageSetupDialog_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPageSetupDialog) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowMaximized
func callbackQPageSetupDialog_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPageSetupDialog) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowMinimized
func callbackQPageSetupDialog_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPageSetupDialog) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_ShowNormal
func callbackQPageSetupDialog_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPageSetupDialog) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_TabletEvent
func callbackQPageSetupDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPageSetupDialog_Update
func callbackQPageSetupDialog_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPageSetupDialog) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_UpdateMicroFocus
func callbackQPageSetupDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPageSetupDialog) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPageSetupDialog_WheelEvent
func callbackQPageSetupDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPageSetupDialog_WindowIconChanged
func callbackQPageSetupDialog_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQPageSetupDialog_WindowTitleChanged
func callbackQPageSetupDialog_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQPageSetupDialog_PaintEngine
func callbackQPageSetupDialog_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQPageSetupDialogFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPageSetupDialog) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPageSetupDialog_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQPageSetupDialog_InputMethodQuery
func callbackQPageSetupDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPageSetupDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPageSetupDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPageSetupDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPageSetupDialog_HasHeightForWidth
func callbackQPageSetupDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPageSetupDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPageSetupDialog) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSetupDialog_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPageSetupDialog_HeightForWidth
func callbackQPageSetupDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPageSetupDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPageSetupDialog) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPageSetupDialog_Metric
func callbackQPageSetupDialog_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQPageSetupDialogFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QPageSetupDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPageSetupDialog_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQPageSetupDialog_InitPainter
func callbackQPageSetupDialog_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQPageSetupDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QPageSetupDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQPageSetupDialog_ChildEvent
func callbackQPageSetupDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPageSetupDialog_ConnectNotify
func callbackQPageSetupDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPageSetupDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPageSetupDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPageSetupDialog_CustomEvent
func callbackQPageSetupDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPageSetupDialog_DeleteLater
func callbackQPageSetupDialog_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPageSetupDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPageSetupDialog) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPageSetupDialog_Destroyed
func callbackQPageSetupDialog_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQPageSetupDialog_DisconnectNotify
func callbackQPageSetupDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPageSetupDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPageSetupDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPageSetupDialog_ObjectNameChanged
func callbackQPageSetupDialog_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQPageSetupDialog_TimerEvent
func callbackQPageSetupDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPageSetupDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPageSetupDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPageSetupDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPrint struct {
	ptr unsafe.Pointer
}

type QPrint_ITF interface {
	QPrint_PTR() *QPrint
}

func (ptr *QPrint) QPrint_PTR() *QPrint {
	return ptr
}

func (ptr *QPrint) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPrint) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPrint(ptr QPrint_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrint_PTR().Pointer()
	}
	return nil
}

func NewQPrintFromPointer(ptr unsafe.Pointer) (n *QPrint) {
	n = new(QPrint)
	n.SetPointer(ptr)
	return
}

func (ptr *QPrint) DestroyQPrint() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QPrint__ColorMode
//QPrint::ColorMode
type QPrint__ColorMode int64

const (
	QPrint__GrayScale QPrint__ColorMode = QPrint__ColorMode(0)
	QPrint__Color     QPrint__ColorMode = QPrint__ColorMode(1)
)

//go:generate stringer -type=QPrint__DeviceState
//QPrint::DeviceState
type QPrint__DeviceState int64

const (
	QPrint__Idle    QPrint__DeviceState = QPrint__DeviceState(0)
	QPrint__Active  QPrint__DeviceState = QPrint__DeviceState(1)
	QPrint__Aborted QPrint__DeviceState = QPrint__DeviceState(2)
	QPrint__Error   QPrint__DeviceState = QPrint__DeviceState(3)
)

//go:generate stringer -type=QPrint__DuplexMode
//QPrint::DuplexMode
type QPrint__DuplexMode int64

const (
	QPrint__DuplexNone      QPrint__DuplexMode = QPrint__DuplexMode(0)
	QPrint__DuplexAuto      QPrint__DuplexMode = QPrint__DuplexMode(1)
	QPrint__DuplexLongSide  QPrint__DuplexMode = QPrint__DuplexMode(2)
	QPrint__DuplexShortSide QPrint__DuplexMode = QPrint__DuplexMode(3)
)

//go:generate stringer -type=QPrint__InputSlotId
//QPrint::InputSlotId
type QPrint__InputSlotId int64

const (
	QPrint__Upper           QPrint__InputSlotId = QPrint__InputSlotId(0)
	QPrint__Lower           QPrint__InputSlotId = QPrint__InputSlotId(1)
	QPrint__Middle          QPrint__InputSlotId = QPrint__InputSlotId(2)
	QPrint__Manual          QPrint__InputSlotId = QPrint__InputSlotId(3)
	QPrint__Envelope        QPrint__InputSlotId = QPrint__InputSlotId(4)
	QPrint__EnvelopeManual  QPrint__InputSlotId = QPrint__InputSlotId(5)
	QPrint__Auto            QPrint__InputSlotId = QPrint__InputSlotId(6)
	QPrint__Tractor         QPrint__InputSlotId = QPrint__InputSlotId(7)
	QPrint__SmallFormat     QPrint__InputSlotId = QPrint__InputSlotId(8)
	QPrint__LargeFormat     QPrint__InputSlotId = QPrint__InputSlotId(9)
	QPrint__LargeCapacity   QPrint__InputSlotId = QPrint__InputSlotId(10)
	QPrint__Cassette        QPrint__InputSlotId = QPrint__InputSlotId(11)
	QPrint__FormSource      QPrint__InputSlotId = QPrint__InputSlotId(12)
	QPrint__MaxPageSource   QPrint__InputSlotId = QPrint__InputSlotId(13)
	QPrint__CustomInputSlot QPrint__InputSlotId = QPrint__InputSlotId(14)
	QPrint__LastInputSlot   QPrint__InputSlotId = QPrint__InputSlotId(QPrint__CustomInputSlot)
	QPrint__OnlyOne         QPrint__InputSlotId = QPrint__InputSlotId(QPrint__Upper)
)

//go:generate stringer -type=QPrint__OutputBinId
//QPrint::OutputBinId
type QPrint__OutputBinId int64

const (
	QPrint__AutoOutputBin   QPrint__OutputBinId = QPrint__OutputBinId(0)
	QPrint__UpperBin        QPrint__OutputBinId = QPrint__OutputBinId(1)
	QPrint__LowerBin        QPrint__OutputBinId = QPrint__OutputBinId(2)
	QPrint__RearBin         QPrint__OutputBinId = QPrint__OutputBinId(3)
	QPrint__CustomOutputBin QPrint__OutputBinId = QPrint__OutputBinId(4)
	QPrint__LastOutputBin   QPrint__OutputBinId = QPrint__OutputBinId(QPrint__CustomOutputBin)
)

type QPrintDialog struct {
	QAbstractPrintDialog
}

type QPrintDialog_ITF interface {
	QAbstractPrintDialog_ITF
	QPrintDialog_PTR() *QPrintDialog
}

func (ptr *QPrintDialog) QPrintDialog_PTR() *QPrintDialog {
	return ptr
}

func (ptr *QPrintDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractPrintDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPrintDialog(ptr QPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintDialog_PTR().Pointer()
	}
	return nil
}

func NewQPrintDialogFromPointer(ptr unsafe.Pointer) (n *QPrintDialog) {
	n = new(QPrintDialog)
	n.SetPointer(ptr)
	return
}
func NewQPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPrintDialog {
	tmpValue := NewQPrintDialogFromPointer(C.QPrintDialog_NewQPrintDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPrintDialog2(parent widgets.QWidget_ITF) *QPrintDialog {
	tmpValue := NewQPrintDialogFromPointer(C.QPrintDialog_NewQPrintDialog2(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQPrintDialog_Exec
func callbackQPrintDialog_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQPrintDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QPrintDialog) ConnectExec(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exec"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exec", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exec", f)
		}
	}
}

func (ptr *QPrintDialog) DisconnectExec() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exec")
	}
}

func (ptr *QPrintDialog) Exec() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintDialog) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQPrintDialog_Accept
func callbackQPrintDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPrintDialog) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accept"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "accept", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accept", f)
		}
	}
}

func (ptr *QPrintDialog) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accept")
	}
}

func (ptr *QPrintDialog) Accept() {
	if ptr.Pointer() != nil {
		C.QPrintDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QPrintDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QPrintDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPrintDialog_Accepted
func callbackQPrintDialog_Accepted(ptr unsafe.Pointer, printer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintDialog) ConnectAccepted(f func(printer *QPrinter)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "accepted") {
			C.QPrintDialog_ConnectAccepted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "accepted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "accepted", func(printer *QPrinter) {
				signal.(func(*QPrinter))(printer)
				f(printer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accepted", f)
		}
	}
}

func (ptr *QPrintDialog) DisconnectAccepted() {
	if ptr.Pointer() != nil {
		C.QPrintDialog_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "accepted")
	}
}

func (ptr *QPrintDialog) Accepted(printer QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintDialog_Accepted(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

//export callbackQPrintDialog_Done
func callbackQPrintDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		signal.(func(int))(int(int32(result)))
	} else {
		NewQPrintDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QPrintDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "done", func(result int) {
				signal.(func(int))(result)
				f(result)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", f)
		}
	}
}

func (ptr *QPrintDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QPrintDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QPrintDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPrintDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QPrintDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPrintDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QPrintDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

func (ptr *QPrintDialog) SetOption(option QAbstractPrintDialog__PrintDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QPrintDialog_SetOption(ptr.Pointer(), C.longlong(option), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QPrintDialog) SetOptions(options QAbstractPrintDialog__PrintDialogOption) {
	if ptr.Pointer() != nil {
		C.QPrintDialog_SetOptions(ptr.Pointer(), C.longlong(options))
	}
}

//export callbackQPrintDialog_DestroyQPrintDialog
func callbackQPrintDialog_DestroyQPrintDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrintDialog"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintDialogFromPointer(ptr).DestroyQPrintDialogDefault()
	}
}

func (ptr *QPrintDialog) ConnectDestroyQPrintDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrintDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintDialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintDialog", f)
		}
	}
}

func (ptr *QPrintDialog) DisconnectDestroyQPrintDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrintDialog")
	}
}

func (ptr *QPrintDialog) DestroyQPrintDialog() {
	if ptr.Pointer() != nil {
		C.QPrintDialog_DestroyQPrintDialog(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintDialog) DestroyQPrintDialogDefault() {
	if ptr.Pointer() != nil {
		C.QPrintDialog_DestroyQPrintDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintDialog) Options() QAbstractPrintDialog__PrintDialogOption {
	if ptr.Pointer() != nil {
		return QAbstractPrintDialog__PrintDialogOption(C.QPrintDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintDialog) TestOption(option QAbstractPrintDialog__PrintDialogOption) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintDialog_TestOption(ptr.Pointer(), C.longlong(option))) != 0
	}
	return false
}

type QPrintEngine struct {
	ptr unsafe.Pointer
}

type QPrintEngine_ITF interface {
	QPrintEngine_PTR() *QPrintEngine
}

func (ptr *QPrintEngine) QPrintEngine_PTR() *QPrintEngine {
	return ptr
}

func (ptr *QPrintEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPrintEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPrintEngine(ptr QPrintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintEngine_PTR().Pointer()
	}
	return nil
}

func NewQPrintEngineFromPointer(ptr unsafe.Pointer) (n *QPrintEngine) {
	n = new(QPrintEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPrintEngine__PrintEnginePropertyKey
//QPrintEngine::PrintEnginePropertyKey
type QPrintEngine__PrintEnginePropertyKey int64

const (
	QPrintEngine__PPK_CollateCopies          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0)
	QPrintEngine__PPK_ColorMode              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(1)
	QPrintEngine__PPK_Creator                QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(2)
	QPrintEngine__PPK_DocumentName           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(3)
	QPrintEngine__PPK_FullPage               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(4)
	QPrintEngine__PPK_NumberOfCopies         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(5)
	QPrintEngine__PPK_Orientation            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(6)
	QPrintEngine__PPK_OutputFileName         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(7)
	QPrintEngine__PPK_PageOrder              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(8)
	QPrintEngine__PPK_PageRect               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(9)
	QPrintEngine__PPK_PageSize               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(10)
	QPrintEngine__PPK_PaperRect              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(11)
	QPrintEngine__PPK_PaperSource            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(12)
	QPrintEngine__PPK_PrinterName            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(13)
	QPrintEngine__PPK_PrinterProgram         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(14)
	QPrintEngine__PPK_Resolution             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(15)
	QPrintEngine__PPK_SelectionOption        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(16)
	QPrintEngine__PPK_SupportedResolutions   QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(17)
	QPrintEngine__PPK_WindowsPageSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(18)
	QPrintEngine__PPK_FontEmbedding          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(19)
	QPrintEngine__PPK_Duplex                 QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(20)
	QPrintEngine__PPK_PaperSources           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(21)
	QPrintEngine__PPK_CustomPaperSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(22)
	QPrintEngine__PPK_PageMargins            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(23)
	QPrintEngine__PPK_CopyCount              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(24)
	QPrintEngine__PPK_SupportsMultipleCopies QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(25)
	QPrintEngine__PPK_PaperName              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(26)
	QPrintEngine__PPK_QPageSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(27)
	QPrintEngine__PPK_QPageMargins           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(28)
	QPrintEngine__PPK_QPageLayout            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(29)
	QPrintEngine__PPK_PaperSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(QPrintEngine__PPK_PageSize)
	QPrintEngine__PPK_CustomBase             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0xff00)
)

//export callbackQPrintEngine_Abort
func callbackQPrintEngine_Abort(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPrintEngine) ConnectAbort(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "abort", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "abort")
	}
}

func (ptr *QPrintEngine) Abort() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintEngine_Abort(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintEngine_NewPage
func callbackQPrintEngine_NewPage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPrintEngine) ConnectNewPage(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newPage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newPage", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newPage", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectNewPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newPage")
	}
}

func (ptr *QPrintEngine) NewPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintEngine_NewPage(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintEngine_SetProperty
func callbackQPrintEngine_SetProperty(ptr unsafe.Pointer, key C.longlong, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		signal.(func(QPrintEngine__PrintEnginePropertyKey, *core.QVariant))(QPrintEngine__PrintEnginePropertyKey(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QPrintEngine) ConnectSetProperty(f func(key QPrintEngine__PrintEnginePropertyKey, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(key QPrintEngine__PrintEnginePropertyKey, value *core.QVariant) {
				signal.(func(QPrintEngine__PrintEnginePropertyKey, *core.QVariant))(key, value)
				f(key, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QPrintEngine) SetProperty(key QPrintEngine__PrintEnginePropertyKey, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintEngine_SetProperty(ptr.Pointer(), C.longlong(key), core.PointerFromQVariant(value))
	}
}

//export callbackQPrintEngine_DestroyQPrintEngine
func callbackQPrintEngine_DestroyQPrintEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrintEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintEngineFromPointer(ptr).DestroyQPrintEngineDefault()
	}
}

func (ptr *QPrintEngine) ConnectDestroyQPrintEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrintEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintEngine", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectDestroyQPrintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrintEngine")
	}
}

func (ptr *QPrintEngine) DestroyQPrintEngine() {
	if ptr.Pointer() != nil {
		C.QPrintEngine_DestroyQPrintEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintEngine) DestroyQPrintEngineDefault() {
	if ptr.Pointer() != nil {
		C.QPrintEngine_DestroyQPrintEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPrintEngine_PrinterState
func callbackQPrintEngine_PrinterState(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "printerState"); signal != nil {
		return C.longlong(signal.(func() QPrinter__PrinterState)())
	}

	return C.longlong(0)
}

func (ptr *QPrintEngine) ConnectPrinterState(f func() QPrinter__PrinterState) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "printerState"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "printerState", func() QPrinter__PrinterState {
				signal.(func() QPrinter__PrinterState)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "printerState", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectPrinterState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "printerState")
	}
}

func (ptr *QPrintEngine) PrinterState() QPrinter__PrinterState {
	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrintEngine_PrinterState(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrintEngine_Property
func callbackQPrintEngine_Property(ptr unsafe.Pointer, key C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "property"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QPrintEngine__PrintEnginePropertyKey) *core.QVariant)(QPrintEngine__PrintEnginePropertyKey(key)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QPrintEngine) ConnectProperty(f func(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "property", func(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant {
				signal.(func(QPrintEngine__PrintEnginePropertyKey) *core.QVariant)(key)
				return f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "property")
	}
}

func (ptr *QPrintEngine) Property(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPrintEngine_Property(ptr.Pointer(), C.longlong(key)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintEngine_Metric
func callbackQPrintEngine_Metric(ptr unsafe.Pointer, id C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(id))))
	}

	return C.int(int32(0))
}

func (ptr *QPrintEngine) ConnectMetric(f func(id gui.QPaintDevice__PaintDeviceMetric) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metric"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "metric", func(id gui.QPaintDevice__PaintDeviceMetric) int {
				signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(id)
				return f(id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metric", f)
		}
	}
}

func (ptr *QPrintEngine) DisconnectMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metric")
	}
}

func (ptr *QPrintEngine) Metric(id gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintEngine_Metric(ptr.Pointer(), C.longlong(id))))
	}
	return 0
}

type QPrintPreviewDialog struct {
	widgets.QDialog
}

type QPrintPreviewDialog_ITF interface {
	widgets.QDialog_ITF
	QPrintPreviewDialog_PTR() *QPrintPreviewDialog
}

func (ptr *QPrintPreviewDialog) QPrintPreviewDialog_PTR() *QPrintPreviewDialog {
	return ptr
}

func (ptr *QPrintPreviewDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintPreviewDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPrintPreviewDialog(ptr QPrintPreviewDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewDialog_PTR().Pointer()
	}
	return nil
}

func NewQPrintPreviewDialogFromPointer(ptr unsafe.Pointer) (n *QPrintPreviewDialog) {
	n = new(QPrintPreviewDialog)
	n.SetPointer(ptr)
	return
}
func NewQPrintPreviewDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {
	tmpValue := NewQPrintPreviewDialogFromPointer(C.QPrintPreviewDialog_NewQPrintPreviewDialog(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPrintPreviewDialog2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {
	tmpValue := NewQPrintPreviewDialogFromPointer(C.QPrintPreviewDialog_NewQPrintPreviewDialog2(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPrintPreviewDialog) Printer() *QPrinter {
	if ptr.Pointer() != nil {
		return NewQPrinterFromPointer(C.QPrintPreviewDialog_Printer(ptr.Pointer()))
	}
	return nil
}

func QPrintPreviewDialog_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewDialog_QPrintPreviewDialog_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QPrintPreviewDialog) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewDialog_QPrintPreviewDialog_Tr(sC, cC, C.int(int32(n))))
}

func QPrintPreviewDialog_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewDialog_QPrintPreviewDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QPrintPreviewDialog) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewDialog_QPrintPreviewDialog_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQPrintPreviewDialog_Done
func callbackQPrintPreviewDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		signal.(func(int))(int(int32(result)))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QPrintPreviewDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "done", func(result int) {
				signal.(func(int))(result)
				f(result)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", f)
		}
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QPrintPreviewDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QPrintPreviewDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

//export callbackQPrintPreviewDialog_PaintRequested
func callbackQPrintPreviewDialog_PaintRequested(ptr unsafe.Pointer, printer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintRequested"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintPreviewDialog) ConnectPaintRequested(f func(printer *QPrinter)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "paintRequested") {
			C.QPrintPreviewDialog_ConnectPaintRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "paintRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paintRequested", func(printer *QPrinter) {
				signal.(func(*QPrinter))(printer)
				f(printer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintRequested", f)
		}
	}
}

func (ptr *QPrintPreviewDialog) DisconnectPaintRequested() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DisconnectPaintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "paintRequested")
	}
}

func (ptr *QPrintPreviewDialog) PaintRequested(printer QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_PaintRequested(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

//export callbackQPrintPreviewDialog_SetVisible
func callbackQPrintPreviewDialog_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QPrintPreviewDialog) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQPrintPreviewDialog_DestroyQPrintPreviewDialog
func callbackQPrintPreviewDialog_DestroyQPrintPreviewDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrintPreviewDialog"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DestroyQPrintPreviewDialogDefault()
	}
}

func (ptr *QPrintPreviewDialog) ConnectDestroyQPrintPreviewDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrintPreviewDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintPreviewDialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintPreviewDialog", f)
		}
	}
}

func (ptr *QPrintPreviewDialog) DisconnectDestroyQPrintPreviewDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrintPreviewDialog")
	}
}

func (ptr *QPrintPreviewDialog) DestroyQPrintPreviewDialog() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DestroyQPrintPreviewDialog(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintPreviewDialog) DestroyQPrintPreviewDialogDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DestroyQPrintPreviewDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPrintPreviewDialog_MetaObject
func callbackQPrintPreviewDialog_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPrintPreviewDialogFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPrintPreviewDialog) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewDialog_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewDialog___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewDialog) __addActions_actions_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewDialog___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewDialog___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewDialog) __actions_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QPrintPreviewDialog___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewDialog___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewDialog) __findChildren_newList2() unsafe.Pointer {
	return C.QPrintPreviewDialog___findChildren_newList2(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewDialog___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewDialog) __findChildren_newList3() unsafe.Pointer {
	return C.QPrintPreviewDialog___findChildren_newList3(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewDialog___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewDialog) __findChildren_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___findChildren_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewDialog) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewDialog___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewDialog) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewDialog) __children_newList() unsafe.Pointer {
	return C.QPrintPreviewDialog___children_newList(ptr.Pointer())
}

//export callbackQPrintPreviewDialog_EventFilter
func callbackQPrintPreviewDialog_EventFilter(ptr unsafe.Pointer, o unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(o), core.NewQEventFromPointer(e)))))
}

func (ptr *QPrintPreviewDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewDialog_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_Exec
func callbackQPrintPreviewDialog_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQPrintPreviewDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QPrintPreviewDialog) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQPrintPreviewDialog_Accept
func callbackQPrintPreviewDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QPrintPreviewDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_Accepted
func callbackQPrintPreviewDialog_Accepted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		signal.(func())()
	}

}

//export callbackQPrintPreviewDialog_CloseEvent
func callbackQPrintPreviewDialog_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

//export callbackQPrintPreviewDialog_ContextMenuEvent
func callbackQPrintPreviewDialog_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQPrintPreviewDialog_Finished
func callbackQPrintPreviewDialog_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(int))(int(int32(result)))
	}

}

//export callbackQPrintPreviewDialog_KeyPressEvent
func callbackQPrintPreviewDialog_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPrintPreviewDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQPrintPreviewDialog_Reject
func callbackQPrintPreviewDialog_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QPrintPreviewDialog) RejectDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RejectDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_Rejected
func callbackQPrintPreviewDialog_Rejected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rejected"); signal != nil {
		signal.(func())()
	}

}

//export callbackQPrintPreviewDialog_ResizeEvent
func callbackQPrintPreviewDialog_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *QPrintPreviewDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackQPrintPreviewDialog_ShowEvent
func callbackQPrintPreviewDialog_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MinimumSizeHint
func callbackQPrintPreviewDialog_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPrintPreviewDialog) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPrintPreviewDialog_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_SizeHint
func callbackQPrintPreviewDialog_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPrintPreviewDialog) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPrintPreviewDialog_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_Close
func callbackQPrintPreviewDialog_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *QPrintPreviewDialog) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewDialog_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_Event
func callbackQPrintPreviewDialog_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QPrintPreviewDialog) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewDialog_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_FocusNextPrevChild
func callbackQPrintPreviewDialog_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPrintPreviewDialog) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewDialog_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_ActionEvent
func callbackQPrintPreviewDialog_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPrintPreviewDialog_ChangeEvent
func callbackQPrintPreviewDialog_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_CustomContextMenuRequested
func callbackQPrintPreviewDialog_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQPrintPreviewDialog_DragEnterEvent
func callbackQPrintPreviewDialog_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DragLeaveEvent
func callbackQPrintPreviewDialog_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DragMoveEvent
func callbackQPrintPreviewDialog_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DropEvent
func callbackQPrintPreviewDialog_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPrintPreviewDialog_EnterEvent
func callbackQPrintPreviewDialog_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_FocusInEvent
func callbackQPrintPreviewDialog_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewDialog_FocusOutEvent
func callbackQPrintPreviewDialog_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Hide
func callbackQPrintPreviewDialog_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPrintPreviewDialog) HideDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_HideDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_HideEvent
func callbackQPrintPreviewDialog_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPrintPreviewDialog_InputMethodEvent
func callbackQPrintPreviewDialog_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPrintPreviewDialog_KeyReleaseEvent
func callbackQPrintPreviewDialog_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewDialog_LeaveEvent
func callbackQPrintPreviewDialog_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Lower
func callbackQPrintPreviewDialog_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPrintPreviewDialog) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_MouseDoubleClickEvent
func callbackQPrintPreviewDialog_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MouseMoveEvent
func callbackQPrintPreviewDialog_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MousePressEvent
func callbackQPrintPreviewDialog_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MouseReleaseEvent
func callbackQPrintPreviewDialog_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewDialog_MoveEvent
func callbackQPrintPreviewDialog_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPrintPreviewDialog_PaintEvent
func callbackQPrintPreviewDialog_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Raise
func callbackQPrintPreviewDialog_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPrintPreviewDialog) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_Repaint
func callbackQPrintPreviewDialog_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPrintPreviewDialog) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_SetDisabled
func callbackQPrintPreviewDialog_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPrintPreviewDialog) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPrintPreviewDialog_SetEnabled
func callbackQPrintPreviewDialog_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewDialog) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewDialog_SetFocus2
func callbackQPrintPreviewDialog_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPrintPreviewDialog) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_SetHidden
func callbackQPrintPreviewDialog_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPrintPreviewDialog) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPrintPreviewDialog_SetStyleSheet
func callbackQPrintPreviewDialog_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QPrintPreviewDialog) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QPrintPreviewDialog_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQPrintPreviewDialog_SetWindowModified
func callbackQPrintPreviewDialog_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewDialog) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewDialog_SetWindowTitle
func callbackQPrintPreviewDialog_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QPrintPreviewDialog) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QPrintPreviewDialog_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQPrintPreviewDialog_Show
func callbackQPrintPreviewDialog_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPrintPreviewDialog) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowFullScreen
func callbackQPrintPreviewDialog_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPrintPreviewDialog) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowMaximized
func callbackQPrintPreviewDialog_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPrintPreviewDialog) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowMinimized
func callbackQPrintPreviewDialog_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPrintPreviewDialog) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_ShowNormal
func callbackQPrintPreviewDialog_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPrintPreviewDialog) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_TabletEvent
func callbackQPrintPreviewDialog_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPrintPreviewDialog_Update
func callbackQPrintPreviewDialog_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPrintPreviewDialog) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_UpdateMicroFocus
func callbackQPrintPreviewDialog_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPrintPreviewDialog) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewDialog_WheelEvent
func callbackQPrintPreviewDialog_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPrintPreviewDialog_WindowIconChanged
func callbackQPrintPreviewDialog_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQPrintPreviewDialog_WindowTitleChanged
func callbackQPrintPreviewDialog_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQPrintPreviewDialog_PaintEngine
func callbackQPrintPreviewDialog_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQPrintPreviewDialogFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPrintPreviewDialog) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrintPreviewDialog_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQPrintPreviewDialog_InputMethodQuery
func callbackQPrintPreviewDialog_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPrintPreviewDialogFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPrintPreviewDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPrintPreviewDialog_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewDialog_HasHeightForWidth
func callbackQPrintPreviewDialog_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPrintPreviewDialog) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewDialog_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintPreviewDialog_HeightForWidth
func callbackQPrintPreviewDialog_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPrintPreviewDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPrintPreviewDialog) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPrintPreviewDialog_Metric
func callbackQPrintPreviewDialog_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQPrintPreviewDialogFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QPrintPreviewDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewDialog_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQPrintPreviewDialog_InitPainter
func callbackQPrintPreviewDialog_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QPrintPreviewDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQPrintPreviewDialog_ChildEvent
func callbackQPrintPreviewDialog_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPrintPreviewDialog_ConnectNotify
func callbackQPrintPreviewDialog_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewDialog_CustomEvent
func callbackQPrintPreviewDialog_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewDialog_DeleteLater
func callbackQPrintPreviewDialog_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPrintPreviewDialog) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPrintPreviewDialog_Destroyed
func callbackQPrintPreviewDialog_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQPrintPreviewDialog_DisconnectNotify
func callbackQPrintPreviewDialog_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewDialog_ObjectNameChanged
func callbackQPrintPreviewDialog_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQPrintPreviewDialog_TimerEvent
func callbackQPrintPreviewDialog_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPrintPreviewDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPrintPreviewWidget struct {
	widgets.QWidget
}

type QPrintPreviewWidget_ITF interface {
	widgets.QWidget_ITF
	QPrintPreviewWidget_PTR() *QPrintPreviewWidget
}

func (ptr *QPrintPreviewWidget) QPrintPreviewWidget_PTR() *QPrintPreviewWidget {
	return ptr
}

func (ptr *QPrintPreviewWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintPreviewWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQPrintPreviewWidget(ptr QPrintPreviewWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewWidget_PTR().Pointer()
	}
	return nil
}

func NewQPrintPreviewWidgetFromPointer(ptr unsafe.Pointer) (n *QPrintPreviewWidget) {
	n = new(QPrintPreviewWidget)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPrintPreviewWidget__ViewMode
//QPrintPreviewWidget::ViewMode
type QPrintPreviewWidget__ViewMode int64

const (
	QPrintPreviewWidget__SinglePageView  QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(0)
	QPrintPreviewWidget__FacingPagesView QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(1)
	QPrintPreviewWidget__AllPagesView    QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(2)
)

//go:generate stringer -type=QPrintPreviewWidget__ZoomMode
//QPrintPreviewWidget::ZoomMode
type QPrintPreviewWidget__ZoomMode int64

const (
	QPrintPreviewWidget__CustomZoom QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(0)
	QPrintPreviewWidget__FitToWidth QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(1)
	QPrintPreviewWidget__FitInView  QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(2)
)

func NewQPrintPreviewWidget(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {
	tmpValue := NewQPrintPreviewWidgetFromPointer(C.QPrintPreviewWidget_NewQPrintPreviewWidget(PointerFromQPrinter(printer), widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPrintPreviewWidget2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {
	tmpValue := NewQPrintPreviewWidgetFromPointer(C.QPrintPreviewWidget_NewQPrintPreviewWidget2(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QPrintPreviewWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewWidget_QPrintPreviewWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QPrintPreviewWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewWidget_QPrintPreviewWidget_Tr(sC, cC, C.int(int32(n))))
}

func QPrintPreviewWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewWidget_QPrintPreviewWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QPrintPreviewWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QPrintPreviewWidget_QPrintPreviewWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQPrintPreviewWidget_FitInView
func callbackQPrintPreviewWidget_FitInView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fitInView"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FitInViewDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectFitInView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fitInView"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fitInView", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fitInView", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFitInView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fitInView")
	}
}

func (ptr *QPrintPreviewWidget) FitInView() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitInView(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) FitInViewDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitInViewDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_FitToWidth
func callbackQPrintPreviewWidget_FitToWidth(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fitToWidth"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FitToWidthDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectFitToWidth(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fitToWidth"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fitToWidth", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fitToWidth", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectFitToWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fitToWidth")
	}
}

func (ptr *QPrintPreviewWidget) FitToWidth() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitToWidth(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) FitToWidthDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FitToWidthDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_PaintRequested
func callbackQPrintPreviewWidget_PaintRequested(ptr unsafe.Pointer, printer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintRequested"); signal != nil {
		signal.(func(*QPrinter))(NewQPrinterFromPointer(printer))
	}

}

func (ptr *QPrintPreviewWidget) ConnectPaintRequested(f func(printer *QPrinter)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "paintRequested") {
			C.QPrintPreviewWidget_ConnectPaintRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "paintRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paintRequested", func(printer *QPrinter) {
				signal.(func(*QPrinter))(printer)
				f(printer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintRequested", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPaintRequested() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectPaintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "paintRequested")
	}
}

func (ptr *QPrintPreviewWidget) PaintRequested(printer QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PaintRequested(ptr.Pointer(), PointerFromQPrinter(printer))
	}
}

//export callbackQPrintPreviewWidget_PreviewChanged
func callbackQPrintPreviewWidget_PreviewChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "previewChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPrintPreviewWidget) ConnectPreviewChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "previewChanged") {
			C.QPrintPreviewWidget_ConnectPreviewChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "previewChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "previewChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "previewChanged", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPreviewChanged() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectPreviewChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "previewChanged")
	}
}

func (ptr *QPrintPreviewWidget) PreviewChanged() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PreviewChanged(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_Print
func callbackQPrintPreviewWidget_Print(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "print"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).PrintDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectPrint(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "print"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "print", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "print", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectPrint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "print")
	}
}

func (ptr *QPrintPreviewWidget) Print() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_Print(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) PrintDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PrintDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetAllPagesViewMode
func callbackQPrintPreviewWidget_SetAllPagesViewMode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setAllPagesViewMode"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetAllPagesViewModeDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetAllPagesViewMode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAllPagesViewMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setAllPagesViewMode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAllPagesViewMode", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetAllPagesViewMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAllPagesViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetAllPagesViewMode() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetAllPagesViewMode(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetAllPagesViewModeDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetAllPagesViewModeDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetCurrentPage
func callbackQPrintPreviewWidget_SetCurrentPage(ptr unsafe.Pointer, page C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentPage"); signal != nil {
		signal.(func(int))(int(int32(page)))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetCurrentPageDefault(int(int32(page)))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetCurrentPage(f func(page int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentPage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentPage", func(page int) {
				signal.(func(int))(page)
				f(page)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentPage", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetCurrentPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentPage")
	}
}

func (ptr *QPrintPreviewWidget) SetCurrentPage(page int) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetCurrentPage(ptr.Pointer(), C.int(int32(page)))
	}
}

func (ptr *QPrintPreviewWidget) SetCurrentPageDefault(page int) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetCurrentPageDefault(ptr.Pointer(), C.int(int32(page)))
	}
}

//export callbackQPrintPreviewWidget_SetFacingPagesViewMode
func callbackQPrintPreviewWidget_SetFacingPagesViewMode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFacingPagesViewMode"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetFacingPagesViewModeDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetFacingPagesViewMode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFacingPagesViewMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFacingPagesViewMode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFacingPagesViewMode", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetFacingPagesViewMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFacingPagesViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetFacingPagesViewMode() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFacingPagesViewMode(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetFacingPagesViewModeDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFacingPagesViewModeDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetLandscapeOrientation
func callbackQPrintPreviewWidget_SetLandscapeOrientation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLandscapeOrientation"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetLandscapeOrientationDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetLandscapeOrientation(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLandscapeOrientation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLandscapeOrientation", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLandscapeOrientation", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetLandscapeOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLandscapeOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetLandscapeOrientation() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetLandscapeOrientation(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetLandscapeOrientationDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetLandscapeOrientationDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetOrientation
func callbackQPrintPreviewWidget_SetOrientation(ptr unsafe.Pointer, orientation C.longlong) {
	if signal := qt.GetSignal(ptr, "setOrientation"); signal != nil {
		signal.(func(QPrinter__Orientation))(QPrinter__Orientation(orientation))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetOrientationDefault(QPrinter__Orientation(orientation))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetOrientation(f func(orientation QPrinter__Orientation)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOrientation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setOrientation", func(orientation QPrinter__Orientation) {
				signal.(func(QPrinter__Orientation))(orientation)
				f(orientation)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOrientation", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetOrientation(orientation QPrinter__Orientation) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

func (ptr *QPrintPreviewWidget) SetOrientationDefault(orientation QPrinter__Orientation) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetOrientationDefault(ptr.Pointer(), C.longlong(orientation))
	}
}

//export callbackQPrintPreviewWidget_SetPortraitOrientation
func callbackQPrintPreviewWidget_SetPortraitOrientation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPortraitOrientation"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetPortraitOrientationDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetPortraitOrientation(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPortraitOrientation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPortraitOrientation", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPortraitOrientation", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetPortraitOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPortraitOrientation")
	}
}

func (ptr *QPrintPreviewWidget) SetPortraitOrientation() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetPortraitOrientation(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetPortraitOrientationDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetPortraitOrientationDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetSinglePageViewMode
func callbackQPrintPreviewWidget_SetSinglePageViewMode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSinglePageViewMode"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetSinglePageViewModeDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetSinglePageViewMode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSinglePageViewMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSinglePageViewMode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSinglePageViewMode", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetSinglePageViewMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSinglePageViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetSinglePageViewMode() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetSinglePageViewMode(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) SetSinglePageViewModeDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetSinglePageViewModeDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetViewMode
func callbackQPrintPreviewWidget_SetViewMode(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "setViewMode"); signal != nil {
		signal.(func(QPrintPreviewWidget__ViewMode))(QPrintPreviewWidget__ViewMode(mode))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetViewModeDefault(QPrintPreviewWidget__ViewMode(mode))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetViewMode(f func(mode QPrintPreviewWidget__ViewMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setViewMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setViewMode", func(mode QPrintPreviewWidget__ViewMode) {
				signal.(func(QPrintPreviewWidget__ViewMode))(mode)
				f(mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setViewMode", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetViewMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setViewMode")
	}
}

func (ptr *QPrintPreviewWidget) SetViewMode(mode QPrintPreviewWidget__ViewMode) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetViewMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QPrintPreviewWidget) SetViewModeDefault(mode QPrintPreviewWidget__ViewMode) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetViewModeDefault(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQPrintPreviewWidget_SetVisible
func callbackQPrintPreviewWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVisible")
	}
}

func (ptr *QPrintPreviewWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QPrintPreviewWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQPrintPreviewWidget_SetZoomFactor
func callbackQPrintPreviewWidget_SetZoomFactor(ptr unsafe.Pointer, factor C.double) {
	if signal := qt.GetSignal(ptr, "setZoomFactor"); signal != nil {
		signal.(func(float64))(float64(factor))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetZoomFactorDefault(float64(factor))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetZoomFactor(f func(factor float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setZoomFactor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setZoomFactor", func(factor float64) {
				signal.(func(float64))(factor)
				f(factor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setZoomFactor", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomFactor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setZoomFactor")
	}
}

func (ptr *QPrintPreviewWidget) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPrintPreviewWidget) SetZoomFactorDefault(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomFactorDefault(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQPrintPreviewWidget_SetZoomMode
func callbackQPrintPreviewWidget_SetZoomMode(ptr unsafe.Pointer, zoomMode C.longlong) {
	if signal := qt.GetSignal(ptr, "setZoomMode"); signal != nil {
		signal.(func(QPrintPreviewWidget__ZoomMode))(QPrintPreviewWidget__ZoomMode(zoomMode))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetZoomModeDefault(QPrintPreviewWidget__ZoomMode(zoomMode))
	}
}

func (ptr *QPrintPreviewWidget) ConnectSetZoomMode(f func(zoomMode QPrintPreviewWidget__ZoomMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setZoomMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setZoomMode", func(zoomMode QPrintPreviewWidget__ZoomMode) {
				signal.(func(QPrintPreviewWidget__ZoomMode))(zoomMode)
				f(zoomMode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setZoomMode", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setZoomMode")
	}
}

func (ptr *QPrintPreviewWidget) SetZoomMode(zoomMode QPrintPreviewWidget__ZoomMode) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomMode(ptr.Pointer(), C.longlong(zoomMode))
	}
}

func (ptr *QPrintPreviewWidget) SetZoomModeDefault(zoomMode QPrintPreviewWidget__ZoomMode) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetZoomModeDefault(ptr.Pointer(), C.longlong(zoomMode))
	}
}

//export callbackQPrintPreviewWidget_UpdatePreview
func callbackQPrintPreviewWidget_UpdatePreview(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updatePreview"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).UpdatePreviewDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectUpdatePreview(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updatePreview"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updatePreview", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updatePreview", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectUpdatePreview() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updatePreview")
	}
}

func (ptr *QPrintPreviewWidget) UpdatePreview() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdatePreview(ptr.Pointer())
	}
}

func (ptr *QPrintPreviewWidget) UpdatePreviewDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdatePreviewDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ZoomIn
func callbackQPrintPreviewWidget_ZoomIn(ptr unsafe.Pointer, factor C.double) {
	if signal := qt.GetSignal(ptr, "zoomIn"); signal != nil {
		signal.(func(float64))(float64(factor))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ZoomInDefault(float64(factor))
	}
}

func (ptr *QPrintPreviewWidget) ConnectZoomIn(f func(factor float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "zoomIn"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "zoomIn", func(factor float64) {
				signal.(func(float64))(factor)
				f(factor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "zoomIn", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectZoomIn() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "zoomIn")
	}
}

func (ptr *QPrintPreviewWidget) ZoomIn(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomIn(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPrintPreviewWidget) ZoomInDefault(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomInDefault(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQPrintPreviewWidget_ZoomOut
func callbackQPrintPreviewWidget_ZoomOut(ptr unsafe.Pointer, factor C.double) {
	if signal := qt.GetSignal(ptr, "zoomOut"); signal != nil {
		signal.(func(float64))(float64(factor))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ZoomOutDefault(float64(factor))
	}
}

func (ptr *QPrintPreviewWidget) ConnectZoomOut(f func(factor float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "zoomOut"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "zoomOut", func(factor float64) {
				signal.(func(float64))(factor)
				f(factor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "zoomOut", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectZoomOut() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "zoomOut")
	}
}

func (ptr *QPrintPreviewWidget) ZoomOut(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomOut(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPrintPreviewWidget) ZoomOutDefault(factor float64) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ZoomOutDefault(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQPrintPreviewWidget_DestroyQPrintPreviewWidget
func callbackQPrintPreviewWidget_DestroyQPrintPreviewWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrintPreviewWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DestroyQPrintPreviewWidgetDefault()
	}
}

func (ptr *QPrintPreviewWidget) ConnectDestroyQPrintPreviewWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrintPreviewWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintPreviewWidget", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrintPreviewWidget", f)
		}
	}
}

func (ptr *QPrintPreviewWidget) DisconnectDestroyQPrintPreviewWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrintPreviewWidget")
	}
}

func (ptr *QPrintPreviewWidget) DestroyQPrintPreviewWidget() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DestroyQPrintPreviewWidget(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintPreviewWidget) DestroyQPrintPreviewWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DestroyQPrintPreviewWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrintPreviewWidget) ViewMode() QPrintPreviewWidget__ViewMode {
	if ptr.Pointer() != nil {
		return QPrintPreviewWidget__ViewMode(C.QPrintPreviewWidget_ViewMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) ZoomMode() QPrintPreviewWidget__ZoomMode {
	if ptr.Pointer() != nil {
		return QPrintPreviewWidget__ZoomMode(C.QPrintPreviewWidget_ZoomMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) Orientation() QPrinter__Orientation {
	if ptr.Pointer() != nil {
		return QPrinter__Orientation(C.QPrintPreviewWidget_Orientation(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrintPreviewWidget_MetaObject
func callbackQPrintPreviewWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPrintPreviewWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPrintPreviewWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPrintPreviewWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrintPreviewWidget) CurrentPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_CurrentPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) PageCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_PageCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPrintPreviewWidget_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrintPreviewWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QPrintPreviewWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QPrintPreviewWidget) __actions_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___actions_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QPrintPreviewWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QPrintPreviewWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QPrintPreviewWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewWidget) __findChildren_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QPrintPreviewWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPrintPreviewWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPrintPreviewWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPrintPreviewWidget) __children_newList() unsafe.Pointer {
	return C.QPrintPreviewWidget___children_newList(ptr.Pointer())
}

//export callbackQPrintPreviewWidget_Close
func callbackQPrintPreviewWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QPrintPreviewWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_Event
func callbackQPrintPreviewWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QPrintPreviewWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_FocusNextPrevChild
func callbackQPrintPreviewWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QPrintPreviewWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_ActionEvent
func callbackQPrintPreviewWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ChangeEvent
func callbackQPrintPreviewWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_CloseEvent
func callbackQPrintPreviewWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ContextMenuEvent
func callbackQPrintPreviewWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQPrintPreviewWidget_CustomContextMenuRequested
func callbackQPrintPreviewWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQPrintPreviewWidget_DragEnterEvent
func callbackQPrintPreviewWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DragLeaveEvent
func callbackQPrintPreviewWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DragMoveEvent
func callbackQPrintPreviewWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DropEvent
func callbackQPrintPreviewWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQPrintPreviewWidget_EnterEvent
func callbackQPrintPreviewWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_FocusInEvent
func callbackQPrintPreviewWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewWidget_FocusOutEvent
func callbackQPrintPreviewWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Hide
func callbackQPrintPreviewWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QPrintPreviewWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_HideEvent
func callbackQPrintPreviewWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQPrintPreviewWidget_InputMethodEvent
func callbackQPrintPreviewWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQPrintPreviewWidget_KeyPressEvent
func callbackQPrintPreviewWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewWidget_KeyReleaseEvent
func callbackQPrintPreviewWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQPrintPreviewWidget_LeaveEvent
func callbackQPrintPreviewWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Lower
func callbackQPrintPreviewWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QPrintPreviewWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_MouseDoubleClickEvent
func callbackQPrintPreviewWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MouseMoveEvent
func callbackQPrintPreviewWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MousePressEvent
func callbackQPrintPreviewWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MouseReleaseEvent
func callbackQPrintPreviewWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQPrintPreviewWidget_MoveEvent
func callbackQPrintPreviewWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQPrintPreviewWidget_PaintEvent
func callbackQPrintPreviewWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Raise
func callbackQPrintPreviewWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QPrintPreviewWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_Repaint
func callbackQPrintPreviewWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QPrintPreviewWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ResizeEvent
func callbackQPrintPreviewWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQPrintPreviewWidget_SetDisabled
func callbackQPrintPreviewWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QPrintPreviewWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQPrintPreviewWidget_SetEnabled
func callbackQPrintPreviewWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewWidget_SetFocus2
func callbackQPrintPreviewWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QPrintPreviewWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_SetHidden
func callbackQPrintPreviewWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QPrintPreviewWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQPrintPreviewWidget_SetStyleSheet
func callbackQPrintPreviewWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QPrintPreviewWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QPrintPreviewWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQPrintPreviewWidget_SetWindowModified
func callbackQPrintPreviewWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QPrintPreviewWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQPrintPreviewWidget_SetWindowTitle
func callbackQPrintPreviewWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QPrintPreviewWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QPrintPreviewWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQPrintPreviewWidget_Show
func callbackQPrintPreviewWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QPrintPreviewWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowEvent
func callbackQPrintPreviewWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ShowFullScreen
func callbackQPrintPreviewWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QPrintPreviewWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowMaximized
func callbackQPrintPreviewWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QPrintPreviewWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowMinimized
func callbackQPrintPreviewWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QPrintPreviewWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_ShowNormal
func callbackQPrintPreviewWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QPrintPreviewWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_TabletEvent
func callbackQPrintPreviewWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQPrintPreviewWidget_Update
func callbackQPrintPreviewWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QPrintPreviewWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_UpdateMicroFocus
func callbackQPrintPreviewWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QPrintPreviewWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQPrintPreviewWidget_WheelEvent
func callbackQPrintPreviewWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQPrintPreviewWidget_WindowIconChanged
func callbackQPrintPreviewWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQPrintPreviewWidget_WindowTitleChanged
func callbackQPrintPreviewWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQPrintPreviewWidget_PaintEngine
func callbackQPrintPreviewWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQPrintPreviewWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPrintPreviewWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrintPreviewWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQPrintPreviewWidget_MinimumSizeHint
func callbackQPrintPreviewWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QPrintPreviewWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPrintPreviewWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_SizeHint
func callbackQPrintPreviewWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQPrintPreviewWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QPrintPreviewWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPrintPreviewWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_InputMethodQuery
func callbackQPrintPreviewWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQPrintPreviewWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QPrintPreviewWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPrintPreviewWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPrintPreviewWidget_HasHeightForWidth
func callbackQPrintPreviewWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QPrintPreviewWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_HeightForWidth
func callbackQPrintPreviewWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQPrintPreviewWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QPrintPreviewWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQPrintPreviewWidget_Metric
func callbackQPrintPreviewWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQPrintPreviewWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QPrintPreviewWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrintPreviewWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQPrintPreviewWidget_InitPainter
func callbackQPrintPreviewWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QPrintPreviewWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQPrintPreviewWidget_EventFilter
func callbackQPrintPreviewWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrintPreviewWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPrintPreviewWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrintPreviewWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPrintPreviewWidget_ChildEvent
func callbackQPrintPreviewWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPrintPreviewWidget_ConnectNotify
func callbackQPrintPreviewWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewWidget_CustomEvent
func callbackQPrintPreviewWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPrintPreviewWidget_DeleteLater
func callbackQPrintPreviewWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPrintPreviewWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQPrintPreviewWidget_Destroyed
func callbackQPrintPreviewWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQPrintPreviewWidget_DisconnectNotify
func callbackQPrintPreviewWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPrintPreviewWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPrintPreviewWidget_ObjectNameChanged
func callbackQPrintPreviewWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPrintSupport_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQPrintPreviewWidget_TimerEvent
func callbackQPrintPreviewWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPrintPreviewWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPrintPreviewWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPrintPreviewWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPrinter struct {
	gui.QPagedPaintDevice
}

type QPrinter_ITF interface {
	gui.QPagedPaintDevice_ITF
	QPrinter_PTR() *QPrinter
}

func (ptr *QPrinter) QPrinter_PTR() *QPrinter {
	return ptr
}

func (ptr *QPrinter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrinter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPagedPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPrinter(ptr QPrinter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinter_PTR().Pointer()
	}
	return nil
}

func NewQPrinterFromPointer(ptr unsafe.Pointer) (n *QPrinter) {
	n = new(QPrinter)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPrinter__ColorMode
//QPrinter::ColorMode
type QPrinter__ColorMode int64

const (
	QPrinter__GrayScale QPrinter__ColorMode = QPrinter__ColorMode(0)
	QPrinter__Color     QPrinter__ColorMode = QPrinter__ColorMode(1)
)

//go:generate stringer -type=QPrinter__DuplexMode
//QPrinter::DuplexMode
type QPrinter__DuplexMode int64

const (
	QPrinter__DuplexNone      QPrinter__DuplexMode = QPrinter__DuplexMode(0)
	QPrinter__DuplexAuto      QPrinter__DuplexMode = QPrinter__DuplexMode(1)
	QPrinter__DuplexLongSide  QPrinter__DuplexMode = QPrinter__DuplexMode(2)
	QPrinter__DuplexShortSide QPrinter__DuplexMode = QPrinter__DuplexMode(3)
)

//go:generate stringer -type=QPrinter__Orientation
//QPrinter::Orientation
type QPrinter__Orientation int64

const (
	QPrinter__Portrait  QPrinter__Orientation = QPrinter__Orientation(0)
	QPrinter__Landscape QPrinter__Orientation = QPrinter__Orientation(1)
)

//go:generate stringer -type=QPrinter__OutputFormat
//QPrinter::OutputFormat
type QPrinter__OutputFormat int64

const (
	QPrinter__NativeFormat QPrinter__OutputFormat = QPrinter__OutputFormat(0)
	QPrinter__PdfFormat    QPrinter__OutputFormat = QPrinter__OutputFormat(1)
)

//go:generate stringer -type=QPrinter__PageOrder
//QPrinter::PageOrder
type QPrinter__PageOrder int64

const (
	QPrinter__FirstPageFirst QPrinter__PageOrder = QPrinter__PageOrder(0)
	QPrinter__LastPageFirst  QPrinter__PageOrder = QPrinter__PageOrder(1)
)

//go:generate stringer -type=QPrinter__PaperSource
//QPrinter::PaperSource
type QPrinter__PaperSource int64

const (
	QPrinter__OnlyOne         QPrinter__PaperSource = QPrinter__PaperSource(0)
	QPrinter__Lower           QPrinter__PaperSource = QPrinter__PaperSource(1)
	QPrinter__Middle          QPrinter__PaperSource = QPrinter__PaperSource(2)
	QPrinter__Manual          QPrinter__PaperSource = QPrinter__PaperSource(3)
	QPrinter__Envelope        QPrinter__PaperSource = QPrinter__PaperSource(4)
	QPrinter__EnvelopeManual  QPrinter__PaperSource = QPrinter__PaperSource(5)
	QPrinter__Auto            QPrinter__PaperSource = QPrinter__PaperSource(6)
	QPrinter__Tractor         QPrinter__PaperSource = QPrinter__PaperSource(7)
	QPrinter__SmallFormat     QPrinter__PaperSource = QPrinter__PaperSource(8)
	QPrinter__LargeFormat     QPrinter__PaperSource = QPrinter__PaperSource(9)
	QPrinter__LargeCapacity   QPrinter__PaperSource = QPrinter__PaperSource(10)
	QPrinter__Cassette        QPrinter__PaperSource = QPrinter__PaperSource(11)
	QPrinter__FormSource      QPrinter__PaperSource = QPrinter__PaperSource(12)
	QPrinter__MaxPageSource   QPrinter__PaperSource = QPrinter__PaperSource(13)
	QPrinter__CustomSource    QPrinter__PaperSource = QPrinter__PaperSource(14)
	QPrinter__LastPaperSource QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__CustomSource)
	QPrinter__Upper           QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__OnlyOne)
)

//go:generate stringer -type=QPrinter__PrintRange
//QPrinter::PrintRange
type QPrinter__PrintRange int64

const (
	QPrinter__AllPages    QPrinter__PrintRange = QPrinter__PrintRange(0)
	QPrinter__Selection   QPrinter__PrintRange = QPrinter__PrintRange(1)
	QPrinter__PageRange   QPrinter__PrintRange = QPrinter__PrintRange(2)
	QPrinter__CurrentPage QPrinter__PrintRange = QPrinter__PrintRange(3)
)

//go:generate stringer -type=QPrinter__PrinterMode
//QPrinter::PrinterMode
type QPrinter__PrinterMode int64

const (
	QPrinter__ScreenResolution  QPrinter__PrinterMode = QPrinter__PrinterMode(0)
	QPrinter__PrinterResolution QPrinter__PrinterMode = QPrinter__PrinterMode(1)
	QPrinter__HighResolution    QPrinter__PrinterMode = QPrinter__PrinterMode(2)
)

//go:generate stringer -type=QPrinter__PrinterState
//QPrinter::PrinterState
type QPrinter__PrinterState int64

const (
	QPrinter__Idle    QPrinter__PrinterState = QPrinter__PrinterState(0)
	QPrinter__Active  QPrinter__PrinterState = QPrinter__PrinterState(1)
	QPrinter__Aborted QPrinter__PrinterState = QPrinter__PrinterState(2)
	QPrinter__Error   QPrinter__PrinterState = QPrinter__PrinterState(3)
)

//go:generate stringer -type=QPrinter__Unit
//QPrinter::Unit
type QPrinter__Unit int64

const (
	QPrinter__Millimeter  QPrinter__Unit = QPrinter__Unit(0)
	QPrinter__Point       QPrinter__Unit = QPrinter__Unit(1)
	QPrinter__Inch        QPrinter__Unit = QPrinter__Unit(2)
	QPrinter__Pica        QPrinter__Unit = QPrinter__Unit(3)
	QPrinter__Didot       QPrinter__Unit = QPrinter__Unit(4)
	QPrinter__Cicero      QPrinter__Unit = QPrinter__Unit(5)
	QPrinter__DevicePixel QPrinter__Unit = QPrinter__Unit(6)
)

func NewQPrinter(mode QPrinter__PrinterMode) *QPrinter {
	return NewQPrinterFromPointer(C.QPrinter_NewQPrinter(C.longlong(mode)))
}

func NewQPrinter2(printer QPrinterInfo_ITF, mode QPrinter__PrinterMode) *QPrinter {
	return NewQPrinterFromPointer(C.QPrinter_NewQPrinter2(PointerFromQPrinterInfo(printer), C.longlong(mode)))
}

func (ptr *QPrinter) Abort() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_Abort(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrinter_NewPage
func callbackQPrinter_NewPage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrinterFromPointer(ptr).NewPageDefault())))
}

func (ptr *QPrinter) ConnectNewPage(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newPage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newPage", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newPage", f)
		}
	}
}

func (ptr *QPrinter) DisconnectNewPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newPage")
	}
}

func (ptr *QPrinter) NewPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_NewPage(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) NewPageDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_NewPageDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) SetCollateCopies(collate bool) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetCollateCopies(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(collate))))
	}
}

func (ptr *QPrinter) SetColorMode(newColorMode QPrinter__ColorMode) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetColorMode(ptr.Pointer(), C.longlong(newColorMode))
	}
}

func (ptr *QPrinter) SetCopyCount(count int) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetCopyCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QPrinter) SetCreator(creator string) {
	if ptr.Pointer() != nil {
		var creatorC *C.char
		if creator != "" {
			creatorC = C.CString(creator)
			defer C.free(unsafe.Pointer(creatorC))
		}
		C.QPrinter_SetCreator(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: creatorC, len: C.longlong(len(creator))})
	}
}

func (ptr *QPrinter) SetDocName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QPrinter_SetDocName(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QPrinter) SetDuplex(duplex QPrinter__DuplexMode) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetDuplex(ptr.Pointer(), C.longlong(duplex))
	}
}

func (ptr *QPrinter) SetEngines(printEngine QPrintEngine_ITF, paintEngine gui.QPaintEngine_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetEngines(ptr.Pointer(), PointerFromQPrintEngine(printEngine), gui.PointerFromQPaintEngine(paintEngine))
	}
}

func (ptr *QPrinter) SetFontEmbeddingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetFontEmbeddingEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QPrinter) SetFromTo(from int, to int) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetFromTo(ptr.Pointer(), C.int(int32(from)), C.int(int32(to)))
	}
}

func (ptr *QPrinter) SetFullPage(fp bool) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetFullPage(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fp))))
	}
}

func (ptr *QPrinter) SetOutputFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QPrinter_SetOutputFileName(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QPrinter) SetOutputFormat(format QPrinter__OutputFormat) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetOutputFormat(ptr.Pointer(), C.longlong(format))
	}
}

func (ptr *QPrinter) SetPageOrder(pageOrder QPrinter__PageOrder) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPageOrder(ptr.Pointer(), C.longlong(pageOrder))
	}
}

func (ptr *QPrinter) SetPaperSource(source QPrinter__PaperSource) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPaperSource(ptr.Pointer(), C.longlong(source))
	}
}

func (ptr *QPrinter) SetPdfVersion(version gui.QPagedPaintDevice__PdfVersion) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPdfVersion(ptr.Pointer(), C.longlong(version))
	}
}

func (ptr *QPrinter) SetPrintProgram(printProg string) {
	if ptr.Pointer() != nil {
		var printProgC *C.char
		if printProg != "" {
			printProgC = C.CString(printProg)
			defer C.free(unsafe.Pointer(printProgC))
		}
		C.QPrinter_SetPrintProgram(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: printProgC, len: C.longlong(len(printProg))})
	}
}

func (ptr *QPrinter) SetPrintRange(ran QPrinter__PrintRange) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPrintRange(ptr.Pointer(), C.longlong(ran))
	}
}

func (ptr *QPrinter) SetPrinterName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QPrinter_SetPrinterName(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QPrinter) SetPrinterSelectionOption(option string) {
	if ptr.Pointer() != nil {
		var optionC *C.char
		if option != "" {
			optionC = C.CString(option)
			defer C.free(unsafe.Pointer(optionC))
		}
		C.QPrinter_SetPrinterSelectionOption(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: optionC, len: C.longlong(len(option))})
	}
}

func (ptr *QPrinter) SetResolution(dpi int) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetResolution(ptr.Pointer(), C.int(int32(dpi)))
	}
}

//export callbackQPrinter_DestroyQPrinter
func callbackQPrinter_DestroyQPrinter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrinter"); signal != nil {
		signal.(func())()
	} else {
		NewQPrinterFromPointer(ptr).DestroyQPrinterDefault()
	}
}

func (ptr *QPrinter) ConnectDestroyQPrinter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrinter"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QPrinter", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrinter", f)
		}
	}
}

func (ptr *QPrinter) DisconnectDestroyQPrinter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrinter")
	}
}

func (ptr *QPrinter) DestroyQPrinter() {
	if ptr.Pointer() != nil {
		C.QPrinter_DestroyQPrinter(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrinter) DestroyQPrinterDefault() {
	if ptr.Pointer() != nil {
		C.QPrinter_DestroyQPrinterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrinter) SupportedResolutions() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPrintSupport_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQPrinterFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__supportedResolutions_atList(i)
			}
			return out
		}(C.QPrinter_SupportedResolutions(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QPrinter) PdfVersion() gui.QPagedPaintDevice__PdfVersion {
	if ptr.Pointer() != nil {
		return gui.QPagedPaintDevice__PdfVersion(C.QPrinter_PdfVersion(ptr.Pointer()))
	}
	return 0
}

//export callbackQPrinter_PaintEngine
func callbackQPrinter_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQPrinterFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPrinter) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", func() *gui.QPaintEngine {
				signal.(func() *gui.QPaintEngine)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", f)
		}
	}
}

func (ptr *QPrinter) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEngine")
	}
}

func (ptr *QPrinter) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrinter_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrinter) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrinter_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrinter) PrintEngine() *QPrintEngine {
	if ptr.Pointer() != nil {
		return NewQPrintEngineFromPointer(C.QPrinter_PrintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPrinter) ColorMode() QPrinter__ColorMode {
	if ptr.Pointer() != nil {
		return QPrinter__ColorMode(C.QPrinter_ColorMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) Duplex() QPrinter__DuplexMode {
	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinter_Duplex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) OutputFormat() QPrinter__OutputFormat {
	if ptr.Pointer() != nil {
		return QPrinter__OutputFormat(C.QPrinter_OutputFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PageOrder() QPrinter__PageOrder {
	if ptr.Pointer() != nil {
		return QPrinter__PageOrder(C.QPrinter_PageOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PaperSource() QPrinter__PaperSource {
	if ptr.Pointer() != nil {
		return QPrinter__PaperSource(C.QPrinter_PaperSource(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PrintRange() QPrinter__PrintRange {
	if ptr.Pointer() != nil {
		return QPrinter__PrintRange(C.QPrinter_PrintRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PrinterState() QPrinter__PrinterState {
	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrinter_PrinterState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) PageRect(unit QPrinter__Unit) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPrinter_PageRect(ptr.Pointer(), C.longlong(unit)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinter) PaperRect(unit QPrinter__Unit) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPrinter_PaperRect(ptr.Pointer(), C.longlong(unit)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinter) Creator() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_Creator(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) DocName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_DocName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) OutputFileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_OutputFileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrintProgram() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_PrintProgram(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrinterName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_PrinterName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) PrinterSelectionOption() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_PrinterSelectionOption(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) CollateCopies() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_CollateCopies(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) FontEmbeddingEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_FontEmbeddingEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) FullPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_FullPage(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) SupportsMultipleCopies() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_SupportsMultipleCopies(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) CopyCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_CopyCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) FromPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_FromPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) Resolution() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_Resolution(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) ToPage() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_ToPage(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) __supportedResolutions_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter___supportedResolutions_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QPrinter) __supportedResolutions_setList(i int) {
	if ptr.Pointer() != nil {
		C.QPrinter___supportedResolutions_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QPrinter) __supportedResolutions_newList() unsafe.Pointer {
	return C.QPrinter___supportedResolutions_newList(ptr.Pointer())
}

//export callbackQPrinter_SetPageSize2
func callbackQPrinter_SetPageSize2(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(ptr, "setPageSize2"); signal != nil {
		signal.(func(gui.QPagedPaintDevice__PageSize))(gui.QPagedPaintDevice__PageSize(size))
	} else {
		NewQPrinterFromPointer(ptr).SetPageSize2Default(gui.QPagedPaintDevice__PageSize(size))
	}
}

func (ptr *QPrinter) SetPageSize2Default(size gui.QPagedPaintDevice__PageSize) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSize2Default(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQPrinter_SetPageSizeMM
func callbackQPrinter_SetPageSizeMM(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPageSizeMM"); signal != nil {
		signal.(func(*core.QSizeF))(core.NewQSizeFFromPointer(size))
	} else {
		NewQPrinterFromPointer(ptr).SetPageSizeMMDefault(core.NewQSizeFFromPointer(size))
	}
}

func (ptr *QPrinter) SetPageSizeMMDefault(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetPageSizeMMDefault(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

//export callbackQPrinter_Metric
func callbackQPrinter_Metric(ptr unsafe.Pointer, metric C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(metric))))
	}

	return C.int(int32(NewQPrinterFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(metric))))
}

func (ptr *QPrinter) MetricDefault(metric gui.QPaintDevice__PaintDeviceMetric) int {
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

func (ptr *QPrinterInfo) QPrinterInfo_PTR() *QPrinterInfo {
	return ptr
}

func (ptr *QPrinterInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPrinterInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPrinterInfo(ptr QPrinterInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinterInfo_PTR().Pointer()
	}
	return nil
}

func NewQPrinterInfoFromPointer(ptr unsafe.Pointer) (n *QPrinterInfo) {
	n = new(QPrinterInfo)
	n.SetPointer(ptr)
	return
}
func QPrinterInfo_AvailablePrinters() []*QPrinterInfo {
	return func(l C.struct_QtPrintSupport_PackedList) []*QPrinterInfo {
		out := make([]*QPrinterInfo, int(l.len))
		tmpList := NewQPrinterInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__availablePrinters_atList(i)
		}
		return out
	}(C.QPrinterInfo_QPrinterInfo_AvailablePrinters())
}

func (ptr *QPrinterInfo) AvailablePrinters() []*QPrinterInfo {
	return func(l C.struct_QtPrintSupport_PackedList) []*QPrinterInfo {
		out := make([]*QPrinterInfo, int(l.len))
		tmpList := NewQPrinterInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__availablePrinters_atList(i)
		}
		return out
	}(C.QPrinterInfo_QPrinterInfo_AvailablePrinters())
}

func QPrinterInfo_DefaultPrinter() *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_DefaultPrinter())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) DefaultPrinter() *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_DefaultPrinter())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func QPrinterInfo_PrinterInfo(printerName string) *QPrinterInfo {
	var printerNameC *C.char
	if printerName != "" {
		printerNameC = C.CString(printerName)
		defer C.free(unsafe.Pointer(printerNameC))
	}
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_PrinterInfo(C.struct_QtPrintSupport_PackedString{data: printerNameC, len: C.longlong(len(printerName))}))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) PrinterInfo(printerName string) *QPrinterInfo {
	var printerNameC *C.char
	if printerName != "" {
		printerNameC = C.CString(printerName)
		defer C.free(unsafe.Pointer(printerNameC))
	}
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_QPrinterInfo_PrinterInfo(C.struct_QtPrintSupport_PackedString{data: printerNameC, len: C.longlong(len(printerName))}))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo() *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo())
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo3(printer QPrinter_ITF) *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo3(PointerFromQPrinter(printer)))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo2(other QPrinterInfo_ITF) *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo2(PointerFromQPrinterInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func QPrinterInfo_DefaultPrinterName() string {
	return cGoUnpackString(C.QPrinterInfo_QPrinterInfo_DefaultPrinterName())
}

func (ptr *QPrinterInfo) DefaultPrinterName() string {
	return cGoUnpackString(C.QPrinterInfo_QPrinterInfo_DefaultPrinterName())
}

func QPrinterInfo_AvailablePrinterNames() []string {
	return strings.Split(cGoUnpackString(C.QPrinterInfo_QPrinterInfo_AvailablePrinterNames()), "|")
}

func (ptr *QPrinterInfo) AvailablePrinterNames() []string {
	return strings.Split(cGoUnpackString(C.QPrinterInfo_QPrinterInfo_AvailablePrinterNames()), "|")
}

func (ptr *QPrinterInfo) DestroyQPrinterInfo() {
	if ptr.Pointer() != nil {
		C.QPrinterInfo_DestroyQPrinterInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QPrinterInfo) SupportedPageSizes() []*gui.QPageSize {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPrintSupport_PackedList) []*gui.QPageSize {
			out := make([]*gui.QPageSize, int(l.len))
			tmpList := NewQPrinterInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__supportedPageSizes_atList(i)
			}
			return out
		}(C.QPrinterInfo_SupportedPageSizes(ptr.Pointer()))
	}
	return make([]*gui.QPageSize, 0)
}

func (ptr *QPrinterInfo) SupportedDuplexModes() []QPrinter__DuplexMode {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPrintSupport_PackedList) []QPrinter__DuplexMode {
			out := make([]QPrinter__DuplexMode, int(l.len))
			tmpList := NewQPrinterInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__supportedDuplexModes_atList(i)
			}
			return out
		}(C.QPrinterInfo_SupportedDuplexModes(ptr.Pointer()))
	}
	return make([]QPrinter__DuplexMode, 0)
}

func (ptr *QPrinterInfo) SupportedResolutions() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPrintSupport_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQPrinterInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__supportedResolutions_atList(i)
			}
			return out
		}(C.QPrinterInfo_SupportedResolutions(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QPrinterInfo) DefaultPageSize() *gui.QPageSize {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPageSizeFromPointer(C.QPrinterInfo_DefaultPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) MaximumPhysicalPageSize() *gui.QPageSize {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPageSizeFromPointer(C.QPrinterInfo_MaximumPhysicalPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) MinimumPhysicalPageSize() *gui.QPageSize {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPageSizeFromPointer(C.QPrinterInfo_MinimumPhysicalPageSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) DefaultDuplexMode() QPrinter__DuplexMode {
	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinterInfo_DefaultDuplexMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinterInfo) State() QPrinter__PrinterState {
	if ptr.Pointer() != nil {
		return QPrinter__PrinterState(C.QPrinterInfo_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinterInfo) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinterInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) Location() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinterInfo_Location(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) MakeAndModel() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinterInfo_MakeAndModel(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) PrinterName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinterInfo_PrinterName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) IsDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinterInfo_IsDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinterInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinterInfo_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinterInfo) IsRemote() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinterInfo_IsRemote(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinterInfo) SupportsCustomPageSizes() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinterInfo_SupportsCustomPageSizes(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinterInfo) __availablePrinters_atList(i int) *QPrinterInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo___availablePrinters_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) __availablePrinters_setList(i QPrinterInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___availablePrinters_setList(ptr.Pointer(), PointerFromQPrinterInfo(i))
	}
}

func (ptr *QPrinterInfo) __availablePrinters_newList() unsafe.Pointer {
	return C.QPrinterInfo___availablePrinters_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedPageSizes_atList(i int) *gui.QPageSize {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPageSizeFromPointer(C.QPrinterInfo___supportedPageSizes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) __supportedPageSizes_setList(i gui.QPageSize_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedPageSizes_setList(ptr.Pointer(), gui.PointerFromQPageSize(i))
	}
}

func (ptr *QPrinterInfo) __supportedPageSizes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedPageSizes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedDuplexModes_atList(i int) QPrinter__DuplexMode {
	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinterInfo___supportedDuplexModes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QPrinterInfo) __supportedDuplexModes_setList(i QPrinter__DuplexMode) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedDuplexModes_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QPrinterInfo) __supportedDuplexModes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedDuplexModes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedPaperSizes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedPaperSizes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedResolutions_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinterInfo___supportedResolutions_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QPrinterInfo) __supportedResolutions_setList(i int) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedResolutions_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QPrinterInfo) __supportedResolutions_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedResolutions_newList(ptr.Pointer())
}
