package quick

//#include "qquickwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = new(QQuickWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return ptr
}

//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     = QQuickWindow__CreateTextureOption(0x0008)
)

//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           = QQuickWindow__RenderStage(4)
)

//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable = QQuickWindow__SceneGraphError(1)
)

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	return NewQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeFocusItemChanged")
	}
}

//export callbackQQuickWindowActiveFocusItemChanged
func callbackQQuickWindowActiveFocusItemChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged").(func())()
}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterAnimating")
	}
}

//export callbackQQuickWindowAfterAnimating
func callbackQQuickWindowAfterAnimating(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterAnimating").(func())()
}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterRendering")
	}
}

//export callbackQQuickWindowAfterRendering
func callbackQQuickWindowAfterRendering(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterRendering").(func())()
}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterSynchronizing")
	}
}

//export callbackQQuickWindowAfterSynchronizing
func callbackQQuickWindowAfterSynchronizing(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterSynchronizing").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeRendering")
	}
}

//export callbackQQuickWindowBeforeRendering
func callbackQQuickWindowBeforeRendering(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "beforeRendering").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeSynchronizing")
	}
}

//export callbackQQuickWindowBeforeSynchronizing
func callbackQQuickWindowBeforeSynchronizing(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing").(func())()
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectColorChanged(f func(v *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQQuickWindowColorChanged
func callbackQQuickWindowColorChanged(ptrName *C.char, v unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "colorChanged").(func(*gui.QColor))(gui.NewQColorFromPointer(v))
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQQuickWindowFrameSwapped
func callbackQQuickWindowFrameSwapped(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "frameSwapped").(func())()
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	if ptr.Pointer() != nil {
		return gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "openglContextCreated")
	}
}

//export callbackQQuickWindowOpenglContextCreated
func callbackQQuickWindowOpenglContextCreated(ptrName *C.char, context unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "openglContextCreated").(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
}

func (ptr *QQuickWindow) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop")
	}
}

//export callbackQQuickWindowSceneGraphAboutToStop
func callbackQQuickWindowSceneGraphAboutToStop(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWindowSceneGraphError
func callbackQQuickWindowSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphError").(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInitialized")
	}
}

//export callbackQQuickWindowSceneGraphInitialized
func callbackQQuickWindowSceneGraphInitialized(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInvalidated")
	}
}

//export callbackQQuickWindowSceneGraphInvalidated
func callbackQQuickWindowSceneGraphInvalidated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated").(func())()
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(ptr.Pointer(), core.PointerFromQRunnable(job), C.int(stage))
	}
}

func (ptr *QQuickWindow) SendEvent(item QQuickItem_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.int(qt.GoBoolToInt(useAlpha)))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(ptr.Pointer(), gui.PointerFromQOpenGLFramebufferObject(fbo))
	}
}

func (ptr *QQuickWindow) Update() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
