// +build !minimal

package macextras

//#include <stdint.h>
//#include <stdlib.h>
//#include "macextras.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"runtime"
	"unsafe"
)

type QMacPasteboardMime struct {
	ptr unsafe.Pointer
}

type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (p *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime {
	return p
}

func (p *QMacPasteboardMime) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QMacPasteboardMime) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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
	return n
}

//export callbackQMacPasteboardMime_CanConvert
func callbackQMacPasteboardMime_CanConvert(ptr unsafe.Pointer, mime *C.char, flav *C.char) C.char {
	defer qt.Recovering("callback QMacPasteboardMime::canConvert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::canConvert"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(mime), C.GoString(flav)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMacPasteboardMime) ConnectCanConvert(f func(mime string, flav string) bool) {
	defer qt.Recovering("connect QMacPasteboardMime::canConvert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::canConvert", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectCanConvert(mime string, flav string) {
	defer qt.Recovering("disconnect QMacPasteboardMime::canConvert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::canConvert")
	}
}

func (ptr *QMacPasteboardMime) CanConvert(mime string, flav string) bool {
	defer qt.Recovering("QMacPasteboardMime::canConvert")

	if ptr.Pointer() != nil {
		var mimeC = C.CString(mime)
		defer C.free(unsafe.Pointer(mimeC))
		var flavC = C.CString(flav)
		defer C.free(unsafe.Pointer(flavC))
		return C.QMacPasteboardMime_CanConvert(ptr.Pointer(), mimeC, flavC) != 0
	}
	return false
}

//export callbackQMacPasteboardMime_ConvertorName
func callbackQMacPasteboardMime_ConvertorName(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QMacPasteboardMime::convertorName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::convertorName"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectConvertorName(f func() string) {
	defer qt.Recovering("connect QMacPasteboardMime::convertorName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::convertorName", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectConvertorName() {
	defer qt.Recovering("disconnect QMacPasteboardMime::convertorName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::convertorName")
	}
}

func (ptr *QMacPasteboardMime) ConvertorName() string {
	defer qt.Recovering("QMacPasteboardMime::convertorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacPasteboardMime_ConvertorName(ptr.Pointer()))
	}
	return ""
}

//export callbackQMacPasteboardMime_Count
func callbackQMacPasteboardMime_Count(ptr unsafe.Pointer, mimeData unsafe.Pointer) C.int {
	defer qt.Recovering("callback QMacPasteboardMime::count")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::count"); signal != nil {
		return C.int(int32(signal.(func(*core.QMimeData) int)(core.NewQMimeDataFromPointer(mimeData))))
	}

	return C.int(int32(NewQMacPasteboardMimeFromPointer(ptr).CountDefault(core.NewQMimeDataFromPointer(mimeData))))
}

func (ptr *QMacPasteboardMime) ConnectCount(f func(mimeData *core.QMimeData) int) {
	defer qt.Recovering("connect QMacPasteboardMime::count")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::count", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectCount() {
	defer qt.Recovering("disconnect QMacPasteboardMime::count")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::count")
	}
}

func (ptr *QMacPasteboardMime) Count(mimeData core.QMimeData_ITF) int {
	defer qt.Recovering("QMacPasteboardMime::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_Count(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

func (ptr *QMacPasteboardMime) CountDefault(mimeData core.QMimeData_ITF) int {
	defer qt.Recovering("QMacPasteboardMime::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_CountDefault(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

//export callbackQMacPasteboardMime_FlavorFor
func callbackQMacPasteboardMime_FlavorFor(ptr unsafe.Pointer, mime *C.char) *C.char {
	defer qt.Recovering("callback QMacPasteboardMime::flavorFor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::flavorFor"); signal != nil {
		return C.CString(signal.(func(string) string)(C.GoString(mime)))
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectFlavorFor(f func(mime string) string) {
	defer qt.Recovering("connect QMacPasteboardMime::flavorFor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::flavorFor", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectFlavorFor(mime string) {
	defer qt.Recovering("disconnect QMacPasteboardMime::flavorFor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::flavorFor")
	}
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {
	defer qt.Recovering("QMacPasteboardMime::flavorFor")

	if ptr.Pointer() != nil {
		var mimeC = C.CString(mime)
		defer C.free(unsafe.Pointer(mimeC))
		return C.GoString(C.QMacPasteboardMime_FlavorFor(ptr.Pointer(), mimeC))
	}
	return ""
}

//export callbackQMacPasteboardMime_MimeFor
func callbackQMacPasteboardMime_MimeFor(ptr unsafe.Pointer, flav *C.char) *C.char {
	defer qt.Recovering("callback QMacPasteboardMime::mimeFor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::mimeFor"); signal != nil {
		return C.CString(signal.(func(string) string)(C.GoString(flav)))
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectMimeFor(f func(flav string) string) {
	defer qt.Recovering("connect QMacPasteboardMime::mimeFor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::mimeFor", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectMimeFor(flav string) {
	defer qt.Recovering("disconnect QMacPasteboardMime::mimeFor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::mimeFor")
	}
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {
	defer qt.Recovering("QMacPasteboardMime::mimeFor")

	if ptr.Pointer() != nil {
		var flavC = C.CString(flav)
		defer C.free(unsafe.Pointer(flavC))
		return C.GoString(C.QMacPasteboardMime_MimeFor(ptr.Pointer(), flavC))
	}
	return ""
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
	defer qt.Recovering("QMacPasteboardMime::~QMacPasteboardMime")

	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMime(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QMacToolBar struct {
	core.QObject
}

type QMacToolBar_ITF interface {
	core.QObject_ITF
	QMacToolBar_PTR() *QMacToolBar
}

func (p *QMacToolBar) QMacToolBar_PTR() *QMacToolBar {
	return p
}

func (p *QMacToolBar) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QMacToolBar) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQMacToolBar(ptr QMacToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBar_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) *QMacToolBar {
	var n = new(QMacToolBar)
	n.SetPointer(ptr)
	return n
}
func NewQMacToolBar(parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	var tmpValue = NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQMacToolBar2(identifier string, parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	var identifierC = C.CString(identifier)
	defer C.free(unsafe.Pointer(identifierC))
	var tmpValue = NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar2(identifierC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addAllowedItem")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedItem(ptr.Pointer(), gui.PointerFromQIcon(icon), textC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addItem")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_AddItem(ptr.Pointer(), gui.PointerFromQIcon(icon), textC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddSeparator() {
	defer qt.Recovering("QMacToolBar::addSeparator")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindow_ITF) {
	defer qt.Recovering("QMacToolBar::attachToWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	defer qt.Recovering("QMacToolBar::detachFromWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	defer qt.Recovering("QMacToolBar::~QMacToolBar")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQMacToolBar_TimerEvent
func callbackQMacToolBar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::timerEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::timerEvent")
	}
}

func (ptr *QMacToolBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMacToolBar_ChildEvent
func callbackQMacToolBar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::childEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::childEvent")
	}
}

func (ptr *QMacToolBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBar_ConnectNotify
func callbackQMacToolBar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMacToolBar::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::connectNotify", f)
	}
}

func (ptr *QMacToolBar) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QMacToolBar::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::connectNotify")
	}
}

func (ptr *QMacToolBar) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBar::connectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBar::connectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_CustomEvent
func callbackQMacToolBar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::customEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::customEvent")
	}
}

func (ptr *QMacToolBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBar_DeleteLater
func callbackQMacToolBar_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBar) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QMacToolBar::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::deleteLater", f)
	}
}

func (ptr *QMacToolBar) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QMacToolBar::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::deleteLater")
	}
}

func (ptr *QMacToolBar) DeleteLater() {
	defer qt.Recovering("QMacToolBar::deleteLater")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBar) DeleteLaterDefault() {
	defer qt.Recovering("QMacToolBar::deleteLater")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQMacToolBar_DisconnectNotify
func callbackQMacToolBar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMacToolBar::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::disconnectNotify", f)
	}
}

func (ptr *QMacToolBar) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QMacToolBar::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::disconnectNotify")
	}
}

func (ptr *QMacToolBar) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBar::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBar::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_Event
func callbackQMacToolBar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QMacToolBar::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBar) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QMacToolBar::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::event", f)
	}
}

func (ptr *QMacToolBar) DisconnectEvent() {
	defer qt.Recovering("disconnect QMacToolBar::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::event")
	}
}

func (ptr *QMacToolBar) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBar::event")

	if ptr.Pointer() != nil {
		return C.QMacToolBar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMacToolBar) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBar::event")

	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMacToolBar_EventFilter
func callbackQMacToolBar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QMacToolBar::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBar) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QMacToolBar::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::eventFilter", f)
	}
}

func (ptr *QMacToolBar) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QMacToolBar::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::eventFilter")
	}
}

func (ptr *QMacToolBar) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBar::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMacToolBar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBar::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMacToolBar_MetaObject
func callbackQMacToolBar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QMacToolBar::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBar) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QMacToolBar::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::metaObject", f)
	}
}

func (ptr *QMacToolBar) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QMacToolBar::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::metaObject")
	}
}

func (ptr *QMacToolBar) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QMacToolBar::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBar_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBar) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QMacToolBar::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QMacToolBarItem::StandardItem
type QMacToolBarItem__StandardItem int64

const (
	QMacToolBarItem__NoStandardItem = QMacToolBarItem__StandardItem(0)
	QMacToolBarItem__Space          = QMacToolBarItem__StandardItem(1)
	QMacToolBarItem__FlexibleSpace  = QMacToolBarItem__StandardItem(2)
)

type QMacToolBarItem struct {
	core.QObject
}

type QMacToolBarItem_ITF interface {
	core.QObject_ITF
	QMacToolBarItem_PTR() *QMacToolBarItem
}

func (p *QMacToolBarItem) QMacToolBarItem_PTR() *QMacToolBarItem {
	return p
}

func (p *QMacToolBarItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QMacToolBarItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQMacToolBarItem(ptr QMacToolBarItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarItem_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarItemFromPointer(ptr unsafe.Pointer) *QMacToolBarItem {
	var n = new(QMacToolBarItem)
	n.SetPointer(ptr)
	return n
}
func NewQMacToolBarItem(parent core.QObject_ITF) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBarItem::QMacToolBarItem")

	var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBarItem_NewQMacToolBarItem(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMacToolBarItem_Activated
func callbackQMacToolBarItem_Activated(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::activated")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::activated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	defer qt.Recovering("connect QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::activated", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	defer qt.Recovering("disconnect QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::activated")
	}
}

func (ptr *QMacToolBarItem) Activated() {
	defer qt.Recovering("QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_Activated(ptr.Pointer())
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	defer qt.Recovering("QMacToolBarItem::~QMacToolBarItem")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) Icon() *gui.QIcon {
	defer qt.Recovering("QMacToolBarItem::icon")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QMacToolBarItem_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) Selectable() bool {
	defer qt.Recovering("QMacToolBarItem::selectable")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Selectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QMacToolBarItem::setIcon")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	defer qt.Recovering("QMacToolBarItem::setSelectable")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(selectable))))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	defer qt.Recovering("QMacToolBarItem::setStandardItem")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(ptr.Pointer(), C.longlong(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	defer qt.Recovering("QMacToolBarItem::setText")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QMacToolBarItem_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	defer qt.Recovering("QMacToolBarItem::standardItem")

	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	defer qt.Recovering("QMacToolBarItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacToolBarItem_Text(ptr.Pointer()))
	}
	return ""
}

//export callbackQMacToolBarItem_TimerEvent
func callbackQMacToolBarItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::timerEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::timerEvent")
	}
}

func (ptr *QMacToolBarItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBarItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMacToolBarItem_ChildEvent
func callbackQMacToolBarItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::childEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::childEvent")
	}
}

func (ptr *QMacToolBarItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBarItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBarItem_ConnectNotify
func callbackQMacToolBarItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMacToolBarItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::connectNotify", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QMacToolBarItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::connectNotify")
	}
}

func (ptr *QMacToolBarItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBarItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBarItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_CustomEvent
func callbackQMacToolBarItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::customEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::customEvent")
	}
}

func (ptr *QMacToolBarItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBarItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBarItem_DeleteLater
func callbackQMacToolBarItem_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBarItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QMacToolBarItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::deleteLater", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QMacToolBarItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::deleteLater")
	}
}

func (ptr *QMacToolBarItem) DeleteLater() {
	defer qt.Recovering("QMacToolBarItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) DeleteLaterDefault() {
	defer qt.Recovering("QMacToolBarItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMacToolBarItem_DisconnectNotify
func callbackQMacToolBarItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMacToolBarItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::disconnectNotify", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QMacToolBarItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::disconnectNotify")
	}
}

func (ptr *QMacToolBarItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBarItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBarItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMacToolBarItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_Event
func callbackQMacToolBarItem_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QMacToolBarItem::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBarItem) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QMacToolBarItem::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::event", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::event")
	}
}

func (ptr *QMacToolBarItem) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBarItem::event")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBarItem::event")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMacToolBarItem_EventFilter
func callbackQMacToolBarItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QMacToolBarItem::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBarItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QMacToolBarItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::eventFilter", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QMacToolBarItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::eventFilter")
	}
}

func (ptr *QMacToolBarItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBarItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMacToolBarItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMacToolBarItem_MetaObject
func callbackQMacToolBarItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QMacToolBarItem::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBarItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QMacToolBarItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::metaObject", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QMacToolBarItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::metaObject")
	}
}

func (ptr *QMacToolBarItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QMacToolBarItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBarItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBarItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QMacToolBarItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBarItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
