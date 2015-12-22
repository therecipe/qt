package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWindow struct {
	core.QObject
	QSurface
}

type QWindow_ITF interface {
	core.QObject_ITF
	QSurface_ITF
	QWindow_PTR() *QWindow
}

func (p *QWindow) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QWindow) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QSurface_PTR().SetPointer(ptr)
}

func PointerFromQWindow(ptr QWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func NewQWindowFromPointer(ptr unsafe.Pointer) *QWindow {
	var n = new(QWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWindow_") {
		n.SetObjectName("QWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QWindow) QWindow_PTR() *QWindow {
	return ptr
}

//QWindow::AncestorMode
type QWindow__AncestorMode int64

const (
	QWindow__ExcludeTransients = QWindow__AncestorMode(0)
	QWindow__IncludeTransients = QWindow__AncestorMode(1)
)

//QWindow::Visibility
type QWindow__Visibility int64

const (
	QWindow__Hidden              = QWindow__Visibility(0)
	QWindow__AutomaticVisibility = QWindow__Visibility(1)
	QWindow__Windowed            = QWindow__Visibility(2)
	QWindow__Minimized           = QWindow__Visibility(3)
	QWindow__Maximized           = QWindow__Visibility(4)
	QWindow__FullScreen          = QWindow__Visibility(5)
)

func (ptr *QWindow) ContentOrientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QWindow::contentOrientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QWindow_ContentOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) Flags() core.Qt__WindowType {
	defer qt.Recovering("QWindow::flags")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) IsVisible() bool {
	defer qt.Recovering("QWindow::isVisible")

	if ptr.Pointer() != nil {
		return C.QWindow_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) MapFromGlobal(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWindow::mapFromGlobal")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWindow_MapFromGlobal(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWindow) MapToGlobal(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWindow::mapToGlobal")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWindow_MapToGlobal(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWindow) Modality() core.Qt__WindowModality {
	defer qt.Recovering("QWindow::modality")

	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWindow_Modality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) Opacity() float64 {
	defer qt.Recovering("QWindow::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QWindow_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ReportContentOrientationChange(orientation core.Qt__ScreenOrientation) {
	defer qt.Recovering("QWindow::reportContentOrientationChange")

	if ptr.Pointer() != nil {
		C.QWindow_ReportContentOrientationChange(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QWindow) SetFlags(flags core.Qt__WindowType) {
	defer qt.Recovering("QWindow::setFlags")

	if ptr.Pointer() != nil {
		C.QWindow_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QWindow) SetHeight(arg int) {
	defer qt.Recovering("QWindow::setHeight")

	if ptr.Pointer() != nil {
		C.QWindow_SetHeight(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetMaximumHeight(h int) {
	defer qt.Recovering("QWindow::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWindow) SetMaximumWidth(w int) {
	defer qt.Recovering("QWindow::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWindow) SetMinimumHeight(h int) {
	defer qt.Recovering("QWindow::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWindow) SetMinimumWidth(w int) {
	defer qt.Recovering("QWindow::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWindow) SetModality(modality core.Qt__WindowModality) {
	defer qt.Recovering("QWindow::setModality")

	if ptr.Pointer() != nil {
		C.QWindow_SetModality(ptr.Pointer(), C.int(modality))
	}
}

func (ptr *QWindow) SetOpacity(level float64) {
	defer qt.Recovering("QWindow::setOpacity")

	if ptr.Pointer() != nil {
		C.QWindow_SetOpacity(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QWindow) SetTitle(v string) {
	defer qt.Recovering("QWindow::setTitle")

	if ptr.Pointer() != nil {
		C.QWindow_SetTitle(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWindow) SetVisibility(v QWindow__Visibility) {
	defer qt.Recovering("QWindow::setVisibility")

	if ptr.Pointer() != nil {
		C.QWindow_SetVisibility(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QWindow) SetVisible(visible bool) {
	defer qt.Recovering("QWindow::setVisible")

	if ptr.Pointer() != nil {
		C.QWindow_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWindow) SetWidth(arg int) {
	defer qt.Recovering("QWindow::setWidth")

	if ptr.Pointer() != nil {
		C.QWindow_SetWidth(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetX(arg int) {
	defer qt.Recovering("QWindow::setX")

	if ptr.Pointer() != nil {
		C.QWindow_SetX(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetY(arg int) {
	defer qt.Recovering("QWindow::setY")

	if ptr.Pointer() != nil {
		C.QWindow_SetY(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) Title() string {
	defer qt.Recovering("QWindow::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWindow) Visibility() QWindow__Visibility {
	defer qt.Recovering("QWindow::visibility")

	if ptr.Pointer() != nil {
		return QWindow__Visibility(C.QWindow_Visibility(ptr.Pointer()))
	}
	return 0
}

func NewQWindow(targetScreen QScreen_ITF) *QWindow {
	defer qt.Recovering("QWindow::QWindow")

	return NewQWindowFromPointer(C.QWindow_NewQWindow(PointerFromQScreen(targetScreen)))
}

func NewQWindow2(parent QWindow_ITF) *QWindow {
	defer qt.Recovering("QWindow::QWindow")

	return NewQWindowFromPointer(C.QWindow_NewQWindow2(PointerFromQWindow(parent)))
}

func (ptr *QWindow) ConnectActiveChanged(f func()) {
	defer qt.Recovering("connect QWindow::activeChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QWindow) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QWindow::activeChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQWindowActiveChanged
func callbackQWindowActiveChanged(ptrName *C.char) {
	defer qt.Recovering("callback QWindow::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWindow) Alert(msec int) {
	defer qt.Recovering("QWindow::alert")

	if ptr.Pointer() != nil {
		C.QWindow_Alert(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QWindow) BaseSize() *core.QSize {
	defer qt.Recovering("QWindow::baseSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWindow_BaseSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Close() bool {
	defer qt.Recovering("QWindow::close")

	if ptr.Pointer() != nil {
		return C.QWindow_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) ConnectContentOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	defer qt.Recovering("connect QWindow::contentOrientationChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectContentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentOrientationChanged", f)
	}
}

func (ptr *QWindow) DisconnectContentOrientationChanged() {
	defer qt.Recovering("disconnect QWindow::contentOrientationChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectContentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentOrientationChanged")
	}
}

//export callbackQWindowContentOrientationChanged
func callbackQWindowContentOrientationChanged(ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QWindow::contentOrientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentOrientationChanged"); signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

func (ptr *QWindow) Create() {
	defer qt.Recovering("QWindow::create")

	if ptr.Pointer() != nil {
		C.QWindow_Create(ptr.Pointer())
	}
}

func (ptr *QWindow) Destroy() {
	defer qt.Recovering("QWindow::destroy")

	if ptr.Pointer() != nil {
		C.QWindow_Destroy(ptr.Pointer())
	}
}

func (ptr *QWindow) DevicePixelRatio() float64 {
	defer qt.Recovering("QWindow::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QWindow_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectExposeEvent(f func(ev *QExposeEvent)) {
	defer qt.Recovering("connect QWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QWindow) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQWindowExposeEvent
func callbackQWindowExposeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*QExposeEvent))(NewQExposeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) FilePath() string {
	defer qt.Recovering("QWindow::filePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_FilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWindow) ConnectFocusInEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQWindowFocusInEvent
func callbackQWindowFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) FocusObject() *core.QObject {
	defer qt.Recovering("QWindow::focusObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QWindow_FocusObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) ConnectFocusObjectChanged(f func(object *core.QObject)) {
	defer qt.Recovering("connect QWindow::focusObjectChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectFocusObjectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusObjectChanged", f)
	}
}

func (ptr *QWindow) DisconnectFocusObjectChanged() {
	defer qt.Recovering("disconnect QWindow::focusObjectChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectFocusObjectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusObjectChanged")
	}
}

//export callbackQWindowFocusObjectChanged
func callbackQWindowFocusObjectChanged(ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QWindow::focusObjectChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusObjectChanged"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QWindow) ConnectFocusOutEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQWindowFocusOutEvent
func callbackQWindowFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) FrameGeometry() *core.QRect {
	defer qt.Recovering("QWindow::frameGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWindow_FrameGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) FramePosition() *core.QPoint {
	defer qt.Recovering("QWindow::framePosition")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWindow_FramePosition(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Geometry() *core.QRect {
	defer qt.Recovering("QWindow::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWindow_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Height() int {
	defer qt.Recovering("QWindow::height")

	if ptr.Pointer() != nil {
		return int(C.QWindow_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectHeightChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::heightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "heightChanged", f)
	}
}

func (ptr *QWindow) DisconnectHeightChanged() {
	defer qt.Recovering("disconnect QWindow::heightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "heightChanged")
	}
}

//export callbackQWindowHeightChanged
func callbackQWindowHeightChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::heightChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) Hide() {
	defer qt.Recovering("QWindow::hide")

	if ptr.Pointer() != nil {
		C.QWindow_Hide(ptr.Pointer())
	}
}

func (ptr *QWindow) ConnectHideEvent(f func(ev *QHideEvent)) {
	defer qt.Recovering("connect QWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQWindowHideEvent
func callbackQWindowHideEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*QHideEvent))(NewQHideEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) Icon() *QIcon {
	defer qt.Recovering("QWindow::icon")

	if ptr.Pointer() != nil {
		return NewQIconFromPointer(C.QWindow_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) IsActive() bool {
	defer qt.Recovering("QWindow::isActive")

	if ptr.Pointer() != nil {
		return C.QWindow_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsAncestorOf(child QWindow_ITF, mode QWindow__AncestorMode) bool {
	defer qt.Recovering("QWindow::isAncestorOf")

	if ptr.Pointer() != nil {
		return C.QWindow_IsAncestorOf(ptr.Pointer(), PointerFromQWindow(child), C.int(mode)) != 0
	}
	return false
}

func (ptr *QWindow) IsExposed() bool {
	defer qt.Recovering("QWindow::isExposed")

	if ptr.Pointer() != nil {
		return C.QWindow_IsExposed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsModal() bool {
	defer qt.Recovering("QWindow::isModal")

	if ptr.Pointer() != nil {
		return C.QWindow_IsModal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsTopLevel() bool {
	defer qt.Recovering("QWindow::isTopLevel")

	if ptr.Pointer() != nil {
		return C.QWindow_IsTopLevel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) ConnectKeyPressEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQWindowKeyPressEvent
func callbackQWindowKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectKeyReleaseEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQWindowKeyReleaseEvent
func callbackQWindowKeyReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) Lower() {
	defer qt.Recovering("QWindow::lower")

	if ptr.Pointer() != nil {
		C.QWindow_Lower(ptr.Pointer())
	}
}

func (ptr *QWindow) Mask() *QRegion {
	defer qt.Recovering("QWindow::mask")

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QWindow_Mask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) MaximumHeight() int {
	defer qt.Recovering("QWindow::maximumHeight")

	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumHeightChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::maximumHeightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumHeightChanged() {
	defer qt.Recovering("disconnect QWindow::maximumHeightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumHeightChanged")
	}
}

//export callbackQWindowMaximumHeightChanged
func callbackQWindowMaximumHeightChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::maximumHeightChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumHeightChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) MaximumSize() *core.QSize {
	defer qt.Recovering("QWindow::maximumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWindow_MaximumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) MaximumWidth() int {
	defer qt.Recovering("QWindow::maximumWidth")

	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumWidthChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::maximumWidthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumWidthChanged() {
	defer qt.Recovering("disconnect QWindow::maximumWidthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumWidthChanged")
	}
}

//export callbackQWindowMaximumWidthChanged
func callbackQWindowMaximumWidthChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::maximumWidthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumWidthChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) MinimumHeight() int {
	defer qt.Recovering("QWindow::minimumHeight")

	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumHeightChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::minimumHeightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "minimumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumHeightChanged() {
	defer qt.Recovering("disconnect QWindow::minimumHeightChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "minimumHeightChanged")
	}
}

//export callbackQWindowMinimumHeightChanged
func callbackQWindowMinimumHeightChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::minimumHeightChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumHeightChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) MinimumSize() *core.QSize {
	defer qt.Recovering("QWindow::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWindow_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) MinimumWidth() int {
	defer qt.Recovering("QWindow::minimumWidth")

	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumWidthChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::minimumWidthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "minimumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumWidthChanged() {
	defer qt.Recovering("disconnect QWindow::minimumWidthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "minimumWidthChanged")
	}
}

//export callbackQWindowMinimumWidthChanged
func callbackQWindowMinimumWidthChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::minimumWidthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumWidthChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) ConnectModalityChanged(f func(modality core.Qt__WindowModality)) {
	defer qt.Recovering("connect QWindow::modalityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectModalityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modalityChanged", f)
	}
}

func (ptr *QWindow) DisconnectModalityChanged() {
	defer qt.Recovering("disconnect QWindow::modalityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectModalityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modalityChanged")
	}
}

//export callbackQWindowModalityChanged
func callbackQWindowModalityChanged(ptrName *C.char, modality C.int) {
	defer qt.Recovering("callback QWindow::modalityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modalityChanged"); signal != nil {
		signal.(func(core.Qt__WindowModality))(core.Qt__WindowModality(modality))
	}

}

func (ptr *QWindow) ConnectMouseDoubleClickEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQWindowMouseDoubleClickEvent
func callbackQWindowMouseDoubleClickEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectMouseMoveEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQWindowMouseMoveEvent
func callbackQWindowMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectMousePressEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQWindowMousePressEvent
func callbackQWindowMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectMouseReleaseEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQWindowMouseReleaseEvent
func callbackQWindowMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectMoveEvent(f func(ev *QMoveEvent)) {
	defer qt.Recovering("connect QWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQWindowMoveEvent
func callbackQWindowMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QMoveEvent))(NewQMoveEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectOpacityChanged(f func(opacity float64)) {
	defer qt.Recovering("connect QWindow::opacityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityChanged", f)
	}
}

func (ptr *QWindow) DisconnectOpacityChanged() {
	defer qt.Recovering("disconnect QWindow::opacityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityChanged")
	}
}

//export callbackQWindowOpacityChanged
func callbackQWindowOpacityChanged(ptrName *C.char, opacity C.double) {
	defer qt.Recovering("callback QWindow::opacityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "opacityChanged"); signal != nil {
		signal.(func(float64))(float64(opacity))
	}

}

func (ptr *QWindow) Parent() *QWindow {
	defer qt.Recovering("QWindow::parent")

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QWindow_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Position() *core.QPoint {
	defer qt.Recovering("QWindow::position")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWindow_Position(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Raise() {
	defer qt.Recovering("QWindow::raise")

	if ptr.Pointer() != nil {
		C.QWindow_Raise(ptr.Pointer())
	}
}

func (ptr *QWindow) RequestActivate() {
	defer qt.Recovering("QWindow::requestActivate")

	if ptr.Pointer() != nil {
		C.QWindow_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QWindow) RequestUpdate() {
	defer qt.Recovering("QWindow::requestUpdate")

	if ptr.Pointer() != nil {
		C.QWindow_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QWindow) Resize(newSize core.QSize_ITF) {
	defer qt.Recovering("QWindow::resize")

	if ptr.Pointer() != nil {
		C.QWindow_Resize(ptr.Pointer(), core.PointerFromQSize(newSize))
	}
}

func (ptr *QWindow) Resize2(w int, h int) {
	defer qt.Recovering("QWindow::resize")

	if ptr.Pointer() != nil {
		C.QWindow_Resize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) ConnectResizeEvent(f func(ev *QResizeEvent)) {
	defer qt.Recovering("connect QWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQWindowResizeEvent
func callbackQWindowResizeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QResizeEvent))(NewQResizeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) Screen() *QScreen {
	defer qt.Recovering("QWindow::screen")

	if ptr.Pointer() != nil {
		return NewQScreenFromPointer(C.QWindow_Screen(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) ConnectScreenChanged(f func(screen *QScreen)) {
	defer qt.Recovering("connect QWindow::screenChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QWindow) DisconnectScreenChanged() {
	defer qt.Recovering("disconnect QWindow::screenChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQWindowScreenChanged
func callbackQWindowScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	defer qt.Recovering("callback QWindow::screenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenChanged"); signal != nil {
		signal.(func(*QScreen))(NewQScreenFromPointer(screen))
	}

}

func (ptr *QWindow) SetBaseSize(size core.QSize_ITF) {
	defer qt.Recovering("QWindow::setBaseSize")

	if ptr.Pointer() != nil {
		C.QWindow_SetBaseSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetCursor(cursor QCursor_ITF) {
	defer qt.Recovering("QWindow::setCursor")

	if ptr.Pointer() != nil {
		C.QWindow_SetCursor(ptr.Pointer(), PointerFromQCursor(cursor))
	}
}

func (ptr *QWindow) SetFilePath(filePath string) {
	defer qt.Recovering("QWindow::setFilePath")

	if ptr.Pointer() != nil {
		C.QWindow_SetFilePath(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QWindow) SetFormat(format QSurfaceFormat_ITF) {
	defer qt.Recovering("QWindow::setFormat")

	if ptr.Pointer() != nil {
		C.QWindow_SetFormat(ptr.Pointer(), PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QWindow) SetFramePosition(point core.QPoint_ITF) {
	defer qt.Recovering("QWindow::setFramePosition")

	if ptr.Pointer() != nil {
		C.QWindow_SetFramePosition(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *QWindow) SetGeometry2(rect core.QRect_ITF) {
	defer qt.Recovering("QWindow::setGeometry")

	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry2(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	defer qt.Recovering("QWindow::setGeometry")

	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry(ptr.Pointer(), C.int(posx), C.int(posy), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) SetIcon(icon QIcon_ITF) {
	defer qt.Recovering("QWindow::setIcon")

	if ptr.Pointer() != nil {
		C.QWindow_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	defer qt.Recovering("QWindow::setKeyboardGrabEnabled")

	if ptr.Pointer() != nil {
		return C.QWindow_SetKeyboardGrabEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetMask(region QRegion_ITF) {
	defer qt.Recovering("QWindow::setMask")

	if ptr.Pointer() != nil {
		C.QWindow_SetMask(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QWindow) SetMaximumSize(size core.QSize_ITF) {
	defer qt.Recovering("QWindow::setMaximumSize")

	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetMinimumSize(size core.QSize_ITF) {
	defer qt.Recovering("QWindow::setMinimumSize")

	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetMouseGrabEnabled(grab bool) bool {
	defer qt.Recovering("QWindow::setMouseGrabEnabled")

	if ptr.Pointer() != nil {
		return C.QWindow_SetMouseGrabEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetParent(parent QWindow_ITF) {
	defer qt.Recovering("QWindow::setParent")

	if ptr.Pointer() != nil {
		C.QWindow_SetParent(ptr.Pointer(), PointerFromQWindow(parent))
	}
}

func (ptr *QWindow) SetPosition(pt core.QPoint_ITF) {
	defer qt.Recovering("QWindow::setPosition")

	if ptr.Pointer() != nil {
		C.QWindow_SetPosition(ptr.Pointer(), core.PointerFromQPoint(pt))
	}
}

func (ptr *QWindow) SetPosition2(posx int, posy int) {
	defer qt.Recovering("QWindow::setPosition")

	if ptr.Pointer() != nil {
		C.QWindow_SetPosition2(ptr.Pointer(), C.int(posx), C.int(posy))
	}
}

func (ptr *QWindow) SetScreen(newScreen QScreen_ITF) {
	defer qt.Recovering("QWindow::setScreen")

	if ptr.Pointer() != nil {
		C.QWindow_SetScreen(ptr.Pointer(), PointerFromQScreen(newScreen))
	}
}

func (ptr *QWindow) SetSizeIncrement(size core.QSize_ITF) {
	defer qt.Recovering("QWindow::setSizeIncrement")

	if ptr.Pointer() != nil {
		C.QWindow_SetSizeIncrement(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetSurfaceType(surfaceType QSurface__SurfaceType) {
	defer qt.Recovering("QWindow::setSurfaceType")

	if ptr.Pointer() != nil {
		C.QWindow_SetSurfaceType(ptr.Pointer(), C.int(surfaceType))
	}
}

func (ptr *QWindow) SetTransientParent(parent QWindow_ITF) {
	defer qt.Recovering("QWindow::setTransientParent")

	if ptr.Pointer() != nil {
		C.QWindow_SetTransientParent(ptr.Pointer(), PointerFromQWindow(parent))
	}
}

func (ptr *QWindow) SetWindowState(state core.Qt__WindowState) {
	defer qt.Recovering("QWindow::setWindowState")

	if ptr.Pointer() != nil {
		C.QWindow_SetWindowState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QWindow) Show() {
	defer qt.Recovering("QWindow::show")

	if ptr.Pointer() != nil {
		C.QWindow_Show(ptr.Pointer())
	}
}

func (ptr *QWindow) ConnectShowEvent(f func(ev *QShowEvent)) {
	defer qt.Recovering("connect QWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQWindowShowEvent
func callbackQWindowShowEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*QShowEvent))(NewQShowEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ShowFullScreen() {
	defer qt.Recovering("QWindow::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWindow_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowMaximized() {
	defer qt.Recovering("QWindow::showMaximized")

	if ptr.Pointer() != nil {
		C.QWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowMinimized() {
	defer qt.Recovering("QWindow::showMinimized")

	if ptr.Pointer() != nil {
		C.QWindow_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowNormal() {
	defer qt.Recovering("QWindow::showNormal")

	if ptr.Pointer() != nil {
		C.QWindow_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWindow) Size() *core.QSize {
	defer qt.Recovering("QWindow::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWindow_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) SizeIncrement() *core.QSize {
	defer qt.Recovering("QWindow::sizeIncrement")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWindow_SizeIncrement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) SurfaceType() QSurface__SurfaceType {
	defer qt.Recovering("QWindow::surfaceType")

	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QWindow_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectTabletEvent(f func(ev *QTabletEvent)) {
	defer qt.Recovering("connect QWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQWindowTabletEvent
func callbackQWindowTabletEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*QTabletEvent))(NewQTabletEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) ConnectTouchEvent(f func(ev *QTouchEvent)) {
	defer qt.Recovering("connect QWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QWindow) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQWindowTouchEvent
func callbackQWindowTouchEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*QTouchEvent))(NewQTouchEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) TransientParent() *QWindow {
	defer qt.Recovering("QWindow::transientParent")

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QWindow_TransientParent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Type() core.Qt__WindowType {
	defer qt.Recovering("QWindow::type")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) UnsetCursor() {
	defer qt.Recovering("QWindow::unsetCursor")

	if ptr.Pointer() != nil {
		C.QWindow_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QWindow) ConnectVisibilityChanged(f func(visibility QWindow__Visibility)) {
	defer qt.Recovering("connect QWindow::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibilityChanged() {
	defer qt.Recovering("disconnect QWindow::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQWindowVisibilityChanged
func callbackQWindowVisibilityChanged(ptrName *C.char, visibility C.int) {
	defer qt.Recovering("callback QWindow::visibilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibilityChanged"); signal != nil {
		signal.(func(QWindow__Visibility))(QWindow__Visibility(visibility))
	}

}

func (ptr *QWindow) ConnectVisibleChanged(f func(arg bool)) {
	defer qt.Recovering("connect QWindow::visibleChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibleChanged() {
	defer qt.Recovering("disconnect QWindow::visibleChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQWindowVisibleChanged
func callbackQWindowVisibleChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::visibleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibleChanged"); signal != nil {
		signal.(func(bool))(int(arg) != 0)
	}

}

func (ptr *QWindow) ConnectWheelEvent(f func(ev *QWheelEvent)) {
	defer qt.Recovering("connect QWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQWindowWheelEvent
func callbackQWindowWheelEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QWheelEvent))(NewQWheelEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QWindow) Width() int {
	defer qt.Recovering("QWindow::width")

	if ptr.Pointer() != nil {
		return int(C.QWindow_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectWidthChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::widthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widthChanged", f)
	}
}

func (ptr *QWindow) DisconnectWidthChanged() {
	defer qt.Recovering("disconnect QWindow::widthChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widthChanged")
	}
}

//export callbackQWindowWidthChanged
func callbackQWindowWidthChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::widthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "widthChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) WindowState() core.Qt__WindowState {
	defer qt.Recovering("QWindow::windowState")

	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWindow_WindowState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectWindowStateChanged(f func(windowState core.Qt__WindowState)) {
	defer qt.Recovering("connect QWindow::windowStateChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowStateChanged() {
	defer qt.Recovering("disconnect QWindow::windowStateChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQWindowWindowStateChanged
func callbackQWindowWindowStateChanged(ptrName *C.char, windowState C.int) {
	defer qt.Recovering("callback QWindow::windowStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowStateChanged"); signal != nil {
		signal.(func(core.Qt__WindowState))(core.Qt__WindowState(windowState))
	}

}

func (ptr *QWindow) ConnectWindowTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWindow::windowTitleChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowTitleChanged() {
	defer qt.Recovering("disconnect QWindow::windowTitleChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWindowWindowTitleChanged
func callbackQWindowWindowTitleChanged(ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWindow::windowTitleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowTitleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWindow) X() int {
	defer qt.Recovering("QWindow::x")

	if ptr.Pointer() != nil {
		return int(C.QWindow_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectXChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::xChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectXChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QWindow) DisconnectXChanged() {
	defer qt.Recovering("disconnect QWindow::xChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQWindowXChanged
func callbackQWindowXChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::xChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "xChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) Y() int {
	defer qt.Recovering("QWindow::y")

	if ptr.Pointer() != nil {
		return int(C.QWindow_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectYChanged(f func(arg int)) {
	defer qt.Recovering("connect QWindow::yChanged")

	if ptr.Pointer() != nil {
		C.QWindow_ConnectYChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QWindow) DisconnectYChanged() {
	defer qt.Recovering("disconnect QWindow::yChanged")

	if ptr.Pointer() != nil {
		C.QWindow_DisconnectYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQWindowYChanged
func callbackQWindowYChanged(ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QWindow::yChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "yChanged"); signal != nil {
		signal.(func(int))(int(arg))
	}

}

func (ptr *QWindow) DestroyQWindow() {
	defer qt.Recovering("QWindow::~QWindow")

	if ptr.Pointer() != nil {
		C.QWindow_DestroyQWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
