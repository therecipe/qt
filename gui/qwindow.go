package gui

//#include "qwindow.h"
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

type QWindowITF interface {
	core.QObjectITF
	QSurfaceITF
	QWindowPTR() *QWindow
}

func (p *QWindow) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QWindow) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QSurfacePTR().SetPointer(ptr)
}

func PointerFromQWindow(ptr QWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindowPTR().Pointer()
	}
	return nil
}

func QWindowFromPointer(ptr unsafe.Pointer) *QWindow {
	var n = new(QWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWindow) QWindowPTR() *QWindow {
	return ptr
}

//QWindow::AncestorMode
type QWindow__AncestorMode int

var (
	QWindow__ExcludeTransients = QWindow__AncestorMode(0)
	QWindow__IncludeTransients = QWindow__AncestorMode(1)
)

//QWindow::Visibility
type QWindow__Visibility int

var (
	QWindow__Hidden              = QWindow__Visibility(0)
	QWindow__AutomaticVisibility = QWindow__Visibility(1)
	QWindow__Windowed            = QWindow__Visibility(2)
	QWindow__Minimized           = QWindow__Visibility(3)
	QWindow__Maximized           = QWindow__Visibility(4)
	QWindow__FullScreen          = QWindow__Visibility(5)
)

func (ptr *QWindow) ContentOrientation() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QWindow_ContentOrientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) Flags() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) Modality() core.Qt__WindowModality {
	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWindow_Modality(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ReportContentOrientationChange(orientation core.Qt__ScreenOrientation) {
	if ptr.Pointer() != nil {
		C.QWindow_ReportContentOrientationChange(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QWindow) SetFlags(flags core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QWindow) SetHeight(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetHeight(C.QtObjectPtr(ptr.Pointer()), C.int(arg))
	}
}

func (ptr *QWindow) SetMaximumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(h))
	}
}

func (ptr *QWindow) SetMaximumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w))
	}
}

func (ptr *QWindow) SetMinimumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(h))
	}
}

func (ptr *QWindow) SetMinimumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w))
	}
}

func (ptr *QWindow) SetModality(modality core.Qt__WindowModality) {
	if ptr.Pointer() != nil {
		C.QWindow_SetModality(C.QtObjectPtr(ptr.Pointer()), C.int(modality))
	}
}

func (ptr *QWindow) SetTitle(v string) {
	if ptr.Pointer() != nil {
		C.QWindow_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWindow) SetVisibility(v QWindow__Visibility) {
	if ptr.Pointer() != nil {
		C.QWindow_SetVisibility(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QWindow) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWindow_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWindow) SetWidth(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetWidth(C.QtObjectPtr(ptr.Pointer()), C.int(arg))
	}
}

func (ptr *QWindow) SetX(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetX(C.QtObjectPtr(ptr.Pointer()), C.int(arg))
	}
}

func (ptr *QWindow) SetY(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetY(C.QtObjectPtr(ptr.Pointer()), C.int(arg))
	}
}

func (ptr *QWindow) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWindow) Visibility() QWindow__Visibility {
	if ptr.Pointer() != nil {
		return QWindow__Visibility(C.QWindow_Visibility(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQWindow(targetScreen QScreenITF) *QWindow {
	return QWindowFromPointer(unsafe.Pointer(C.QWindow_NewQWindow(C.QtObjectPtr(PointerFromQScreen(targetScreen)))))
}

func NewQWindow2(parent QWindowITF) *QWindow {
	return QWindowFromPointer(unsafe.Pointer(C.QWindow_NewQWindow2(C.QtObjectPtr(PointerFromQWindow(parent)))))
}

func (ptr *QWindow) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QWindow) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQWindowActiveChanged
func callbackQWindowActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QWindow) Alert(msec int) {
	if ptr.Pointer() != nil {
		C.QWindow_Alert(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QWindow) Close() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_Close(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) ConnectContentOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectContentOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentOrientationChanged", f)
	}
}

func (ptr *QWindow) DisconnectContentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectContentOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentOrientationChanged")
	}
}

//export callbackQWindowContentOrientationChanged
func callbackQWindowContentOrientationChanged(ptrName *C.char, orientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "contentOrientationChanged").(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
}

func (ptr *QWindow) Create() {
	if ptr.Pointer() != nil {
		C.QWindow_Create(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) Destroy() {
	if ptr.Pointer() != nil {
		C.QWindow_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) FilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_FilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWindow) FocusObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QWindow_FocusObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWindow) ConnectFocusObjectChanged(f func(object core.QObjectITF)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectFocusObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusObjectChanged", f)
	}
}

func (ptr *QWindow) DisconnectFocusObjectChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectFocusObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusObjectChanged")
	}
}

//export callbackQWindowFocusObjectChanged
func callbackQWindowFocusObjectChanged(ptrName *C.char, object unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusObjectChanged").(func(*core.QObject))(core.QObjectFromPointer(object))
}

func (ptr *QWindow) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "heightChanged", f)
	}
}

func (ptr *QWindow) DisconnectHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "heightChanged")
	}
}

//export callbackQWindowHeightChanged
func callbackQWindowHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "heightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) Hide() {
	if ptr.Pointer() != nil {
		C.QWindow_Hide(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) IsAncestorOf(child QWindowITF, mode QWindow__AncestorMode) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsAncestorOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWindow(child)), C.int(mode)) != 0
	}
	return false
}

func (ptr *QWindow) IsExposed() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsExposed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) IsModal() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsModal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) IsTopLevel() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsTopLevel(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) Lower() {
	if ptr.Pointer() != nil {
		C.QWindow_Lower(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) MaximumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "maximumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "maximumHeightChanged")
	}
}

//export callbackQWindowMaximumHeightChanged
func callbackQWindowMaximumHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "maximumHeightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "maximumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "maximumWidthChanged")
	}
}

//export callbackQWindowMaximumWidthChanged
func callbackQWindowMaximumWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "maximumWidthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "minimumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumHeightChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "minimumHeightChanged")
	}
}

//export callbackQWindowMinimumHeightChanged
func callbackQWindowMinimumHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "minimumHeightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "minimumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "minimumWidthChanged")
	}
}

//export callbackQWindowMinimumWidthChanged
func callbackQWindowMinimumWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "minimumWidthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) ConnectModalityChanged(f func(modality core.Qt__WindowModality)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectModalityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modalityChanged", f)
	}
}

func (ptr *QWindow) DisconnectModalityChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectModalityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modalityChanged")
	}
}

//export callbackQWindowModalityChanged
func callbackQWindowModalityChanged(ptrName *C.char, modality C.int) {
	qt.GetSignal(C.GoString(ptrName), "modalityChanged").(func(core.Qt__WindowModality))(core.Qt__WindowModality(modality))
}

func (ptr *QWindow) Parent() *QWindow {
	if ptr.Pointer() != nil {
		return QWindowFromPointer(unsafe.Pointer(C.QWindow_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWindow) Raise() {
	if ptr.Pointer() != nil {
		C.QWindow_Raise(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) RequestActivate() {
	if ptr.Pointer() != nil {
		C.QWindow_RequestActivate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QWindow_RequestUpdate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) Resize(newSize core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(newSize)))
	}
}

func (ptr *QWindow) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize2(C.QtObjectPtr(ptr.Pointer()), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) Screen() *QScreen {
	if ptr.Pointer() != nil {
		return QScreenFromPointer(unsafe.Pointer(C.QWindow_Screen(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWindow) ConnectScreenChanged(f func(screen QScreenITF)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QWindow) DisconnectScreenChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQWindowScreenChanged
func callbackQWindowScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenChanged").(func(*QScreen))(QScreenFromPointer(screen))
}

func (ptr *QWindow) SetBaseSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetBaseSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QWindow) SetCursor(cursor QCursorITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCursor(cursor)))
	}
}

func (ptr *QWindow) SetFilePath(filePath string) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFilePath(C.QtObjectPtr(ptr.Pointer()), C.CString(filePath))
	}
}

func (ptr *QWindow) SetFormat(format QSurfaceFormatITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSurfaceFormat(format)))
	}
}

func (ptr *QWindow) SetFramePosition(point core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFramePosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))
	}
}

func (ptr *QWindow) SetGeometry2(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.int(posx), C.int(posy), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) SetIcon(icon QIconITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIcon(icon)))
	}
}

func (ptr *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_SetKeyboardGrabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetMask(region QRegionITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMask(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)))
	}
}

func (ptr *QWindow) SetMaximumSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QWindow) SetMinimumSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QWindow) SetMouseGrabEnabled(grab bool) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_SetMouseGrabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetParent(parent QWindowITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetParent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWindow(parent)))
	}
}

func (ptr *QWindow) SetPosition(pt core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pt)))
	}
}

func (ptr *QWindow) SetPosition2(posx int, posy int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition2(C.QtObjectPtr(ptr.Pointer()), C.int(posx), C.int(posy))
	}
}

func (ptr *QWindow) SetScreen(newScreen QScreenITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetScreen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScreen(newScreen)))
	}
}

func (ptr *QWindow) SetSizeIncrement(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetSizeIncrement(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QWindow) SetSurfaceType(surfaceType QSurface__SurfaceType) {
	if ptr.Pointer() != nil {
		C.QWindow_SetSurfaceType(C.QtObjectPtr(ptr.Pointer()), C.int(surfaceType))
	}
}

func (ptr *QWindow) SetTransientParent(parent QWindowITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetTransientParent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWindow(parent)))
	}
}

func (ptr *QWindow) SetWindowState(state core.Qt__WindowState) {
	if ptr.Pointer() != nil {
		C.QWindow_SetWindowState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QWindow) Show() {
	if ptr.Pointer() != nil {
		C.QWindow_Show(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowFullScreen(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMaximized(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMinimized(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowNormal(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QWindow_SurfaceType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) TransientParent() *QWindow {
	if ptr.Pointer() != nil {
		return QWindowFromPointer(unsafe.Pointer(C.QWindow_TransientParent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWindow) Type() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QWindow_UnsetCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWindow) ConnectVisibilityChanged(f func(visibility QWindow__Visibility)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQWindowVisibilityChanged
func callbackQWindowVisibilityChanged(ptrName *C.char, visibility C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(QWindow__Visibility))(QWindow__Visibility(visibility))
}

func (ptr *QWindow) ConnectVisibleChanged(f func(arg bool)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQWindowVisibleChanged
func callbackQWindowVisibleChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func(bool))(int(arg) != 0)
}

func (ptr *QWindow) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "widthChanged", f)
	}
}

func (ptr *QWindow) DisconnectWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWidthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "widthChanged")
	}
}

//export callbackQWindowWidthChanged
func callbackQWindowWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "widthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) WindowState() core.Qt__WindowState {
	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWindow_WindowState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectWindowStateChanged(f func(windowState core.Qt__WindowState)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowStateChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQWindowWindowStateChanged
func callbackQWindowWindowStateChanged(ptrName *C.char, windowState C.int) {
	qt.GetSignal(C.GoString(ptrName), "windowStateChanged").(func(core.Qt__WindowState))(core.Qt__WindowState(windowState))
}

func (ptr *QWindow) ConnectWindowTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowTitleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowTitleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWindowWindowTitleChanged
func callbackQWindowWindowTitleChanged(ptrName *C.char, title *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowTitleChanged").(func(string))(C.GoString(title))
}

func (ptr *QWindow) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectXChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectXChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QWindow) DisconnectXChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectXChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQWindowXChanged
func callbackQWindowXChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "xChanged").(func(int))(int(arg))
}

func (ptr *QWindow) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) ConnectYChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectYChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QWindow) DisconnectYChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectYChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQWindowYChanged
func callbackQWindowYChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "yChanged").(func(int))(int(arg))
}

func (ptr *QWindow) DestroyQWindow() {
	if ptr.Pointer() != nil {
		C.QWindow_DestroyQWindow(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
