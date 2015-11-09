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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QWindow_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QWindow_ContentOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) Flags() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) Modality() core.Qt__WindowModality {
	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWindow_Modality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWindow_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ReportContentOrientationChange(orientation core.Qt__ScreenOrientation) {
	if ptr.Pointer() != nil {
		C.QWindow_ReportContentOrientationChange(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QWindow) SetFlags(flags core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QWindow) SetHeight(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetHeight(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetMaximumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWindow) SetMaximumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWindow) SetMinimumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWindow) SetMinimumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWindow) SetModality(modality core.Qt__WindowModality) {
	if ptr.Pointer() != nil {
		C.QWindow_SetModality(ptr.Pointer(), C.int(modality))
	}
}

func (ptr *QWindow) SetOpacity(level float64) {
	if ptr.Pointer() != nil {
		C.QWindow_SetOpacity(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QWindow) SetTitle(v string) {
	if ptr.Pointer() != nil {
		C.QWindow_SetTitle(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWindow) SetVisibility(v QWindow__Visibility) {
	if ptr.Pointer() != nil {
		C.QWindow_SetVisibility(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QWindow) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWindow_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWindow) SetWidth(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetWidth(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetX(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetX(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) SetY(arg int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetY(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QWindow) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWindow) Visibility() QWindow__Visibility {
	if ptr.Pointer() != nil {
		return QWindow__Visibility(C.QWindow_Visibility(ptr.Pointer()))
	}
	return 0
}

func NewQWindow(targetScreen QScreen_ITF) *QWindow {
	return NewQWindowFromPointer(C.QWindow_NewQWindow(PointerFromQScreen(targetScreen)))
}

func NewQWindow2(parent QWindow_ITF) *QWindow {
	return NewQWindowFromPointer(C.QWindow_NewQWindow2(PointerFromQWindow(parent)))
}

func (ptr *QWindow) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QWindow) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQWindowActiveChanged
func callbackQWindowActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QWindow) Alert(msec int) {
	if ptr.Pointer() != nil {
		C.QWindow_Alert(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QWindow) Close() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) ConnectContentOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectContentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentOrientationChanged", f)
	}
}

func (ptr *QWindow) DisconnectContentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectContentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentOrientationChanged")
	}
}

//export callbackQWindowContentOrientationChanged
func callbackQWindowContentOrientationChanged(ptrName *C.char, orientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "contentOrientationChanged").(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
}

func (ptr *QWindow) Create() {
	if ptr.Pointer() != nil {
		C.QWindow_Create(ptr.Pointer())
	}
}

func (ptr *QWindow) Destroy() {
	if ptr.Pointer() != nil {
		C.QWindow_Destroy(ptr.Pointer())
	}
}

func (ptr *QWindow) DevicePixelRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWindow_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) FilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWindow_FilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWindow) FocusObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QWindow_FocusObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) ConnectFocusObjectChanged(f func(object *core.QObject)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectFocusObjectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusObjectChanged", f)
	}
}

func (ptr *QWindow) DisconnectFocusObjectChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectFocusObjectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusObjectChanged")
	}
}

//export callbackQWindowFocusObjectChanged
func callbackQWindowFocusObjectChanged(ptrName *C.char, object unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusObjectChanged").(func(*core.QObject))(core.NewQObjectFromPointer(object))
}

func (ptr *QWindow) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "heightChanged", f)
	}
}

func (ptr *QWindow) DisconnectHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "heightChanged")
	}
}

//export callbackQWindowHeightChanged
func callbackQWindowHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "heightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) Hide() {
	if ptr.Pointer() != nil {
		C.QWindow_Hide(ptr.Pointer())
	}
}

func (ptr *QWindow) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsAncestorOf(child QWindow_ITF, mode QWindow__AncestorMode) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsAncestorOf(ptr.Pointer(), PointerFromQWindow(child), C.int(mode)) != 0
	}
	return false
}

func (ptr *QWindow) IsExposed() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsExposed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsModal() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsModal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) IsTopLevel() bool {
	if ptr.Pointer() != nil {
		return C.QWindow_IsTopLevel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWindow) Lower() {
	if ptr.Pointer() != nil {
		C.QWindow_Lower(ptr.Pointer())
	}
}

func (ptr *QWindow) Mask() *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QWindow_Mask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) MaximumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumHeightChanged")
	}
}

//export callbackQWindowMaximumHeightChanged
func callbackQWindowMaximumHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "maximumHeightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMaximumWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMaximumWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMaximumWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMaximumWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumWidthChanged")
	}
}

//export callbackQWindowMaximumWidthChanged
func callbackQWindowMaximumWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "maximumWidthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumHeightChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumHeightChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "minimumHeightChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumHeightChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "minimumHeightChanged")
	}
}

//export callbackQWindowMinimumHeightChanged
func callbackQWindowMinimumHeightChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "minimumHeightChanged").(func(int))(int(arg))
}

func (ptr *QWindow) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectMinimumWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectMinimumWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "minimumWidthChanged", f)
	}
}

func (ptr *QWindow) DisconnectMinimumWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectMinimumWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "minimumWidthChanged")
	}
}

//export callbackQWindowMinimumWidthChanged
func callbackQWindowMinimumWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "minimumWidthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) ConnectModalityChanged(f func(modality core.Qt__WindowModality)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectModalityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modalityChanged", f)
	}
}

func (ptr *QWindow) DisconnectModalityChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectModalityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modalityChanged")
	}
}

//export callbackQWindowModalityChanged
func callbackQWindowModalityChanged(ptrName *C.char, modality C.int) {
	qt.GetSignal(C.GoString(ptrName), "modalityChanged").(func(core.Qt__WindowModality))(core.Qt__WindowModality(modality))
}

func (ptr *QWindow) Parent() *QWindow {
	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QWindow_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Raise() {
	if ptr.Pointer() != nil {
		C.QWindow_Raise(ptr.Pointer())
	}
}

func (ptr *QWindow) RequestActivate() {
	if ptr.Pointer() != nil {
		C.QWindow_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QWindow) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QWindow_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QWindow) Resize(newSize core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize(ptr.Pointer(), core.PointerFromQSize(newSize))
	}
}

func (ptr *QWindow) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) Screen() *QScreen {
	if ptr.Pointer() != nil {
		return NewQScreenFromPointer(C.QWindow_Screen(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) ConnectScreenChanged(f func(screen *QScreen)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QWindow) DisconnectScreenChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQWindowScreenChanged
func callbackQWindowScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenChanged").(func(*QScreen))(NewQScreenFromPointer(screen))
}

func (ptr *QWindow) SetBaseSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetBaseSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetCursor(cursor QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetCursor(ptr.Pointer(), PointerFromQCursor(cursor))
	}
}

func (ptr *QWindow) SetFilePath(filePath string) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFilePath(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QWindow) SetFormat(format QSurfaceFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFormat(ptr.Pointer(), PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QWindow) SetFramePosition(point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetFramePosition(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *QWindow) SetGeometry2(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry2(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry(ptr.Pointer(), C.int(posx), C.int(posy), C.int(w), C.int(h))
	}
}

func (ptr *QWindow) SetIcon(icon QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_SetKeyboardGrabEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetMask(region QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMask(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QWindow) SetMaximumSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetMinimumSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetMouseGrabEnabled(grab bool) bool {
	if ptr.Pointer() != nil {
		return C.QWindow_SetMouseGrabEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(grab))) != 0
	}
	return false
}

func (ptr *QWindow) SetParent(parent QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetParent(ptr.Pointer(), PointerFromQWindow(parent))
	}
}

func (ptr *QWindow) SetPosition(pt core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition(ptr.Pointer(), core.PointerFromQPoint(pt))
	}
}

func (ptr *QWindow) SetPosition2(posx int, posy int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition2(ptr.Pointer(), C.int(posx), C.int(posy))
	}
}

func (ptr *QWindow) SetScreen(newScreen QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetScreen(ptr.Pointer(), PointerFromQScreen(newScreen))
	}
}

func (ptr *QWindow) SetSizeIncrement(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetSizeIncrement(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWindow) SetSurfaceType(surfaceType QSurface__SurfaceType) {
	if ptr.Pointer() != nil {
		C.QWindow_SetSurfaceType(ptr.Pointer(), C.int(surfaceType))
	}
}

func (ptr *QWindow) SetTransientParent(parent QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetTransientParent(ptr.Pointer(), PointerFromQWindow(parent))
	}
}

func (ptr *QWindow) SetWindowState(state core.Qt__WindowState) {
	if ptr.Pointer() != nil {
		C.QWindow_SetWindowState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QWindow) Show() {
	if ptr.Pointer() != nil {
		C.QWindow_Show(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWindow) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QWindow_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) TransientParent() *QWindow {
	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QWindow_TransientParent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWindow) Type() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QWindow_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QWindow) ConnectVisibilityChanged(f func(visibility QWindow__Visibility)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQWindowVisibilityChanged
func callbackQWindowVisibilityChanged(ptrName *C.char, visibility C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(QWindow__Visibility))(QWindow__Visibility(visibility))
}

func (ptr *QWindow) ConnectVisibleChanged(f func(arg bool)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QWindow) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQWindowVisibleChanged
func callbackQWindowVisibleChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func(bool))(int(arg) != 0)
}

func (ptr *QWindow) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectWidthChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWidthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widthChanged", f)
	}
}

func (ptr *QWindow) DisconnectWidthChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widthChanged")
	}
}

//export callbackQWindowWidthChanged
func callbackQWindowWidthChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "widthChanged").(func(int))(int(arg))
}

func (ptr *QWindow) WindowState() core.Qt__WindowState {
	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWindow_WindowState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectWindowStateChanged(f func(windowState core.Qt__WindowState)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowStateChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQWindowWindowStateChanged
func callbackQWindowWindowStateChanged(ptrName *C.char, windowState C.int) {
	qt.GetSignal(C.GoString(ptrName), "windowStateChanged").(func(core.Qt__WindowState))(core.Qt__WindowState(windowState))
}

func (ptr *QWindow) ConnectWindowTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectWindowTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWindow) DisconnectWindowTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectWindowTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWindowWindowTitleChanged
func callbackQWindowWindowTitleChanged(ptrName *C.char, title *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowTitleChanged").(func(string))(C.GoString(title))
}

func (ptr *QWindow) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectXChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectXChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QWindow) DisconnectXChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQWindowXChanged
func callbackQWindowXChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "xChanged").(func(int))(int(arg))
}

func (ptr *QWindow) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWindow_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) ConnectYChanged(f func(arg int)) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectYChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QWindow) DisconnectYChanged() {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQWindowYChanged
func callbackQWindowYChanged(ptrName *C.char, arg C.int) {
	qt.GetSignal(C.GoString(ptrName), "yChanged").(func(int))(int(arg))
}

func (ptr *QWindow) DestroyQWindow() {
	if ptr.Pointer() != nil {
		C.QWindow_DestroyQWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
