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

type QQuickWindowITF interface {
	gui.QWindowITF
	QQuickWindowPTR() *QQuickWindow
}

func PointerFromQQuickWindow(ptr QQuickWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindowPTR().Pointer()
	}
	return nil
}

func QQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = new(QQuickWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickWindow) QQuickWindowPTR() *QQuickWindow {
	return ptr
}

//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int

var (
	QQuickWindow__TextureHasAlphaChannel = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     = QQuickWindow__CreateTextureOption(0x0008)
)

//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int

var (
	QQuickWindow__BeforeSynchronizingStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           = QQuickWindow__RenderStage(4)
)

//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int

var (
	QQuickWindow__ContextNotAvailable = QQuickWindow__SceneGraphError(1)
)

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickWindow_ActiveFocusItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickWindow_ContentItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func NewQQuickWindow(parent gui.QWindowITF) *QQuickWindow {
	return QQuickWindowFromPointer(unsafe.Pointer(C.QQuickWindow_NewQQuickWindow(C.QtObjectPtr(gui.PointerFromQWindow(parent)))))
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QQuickWindow_AccessibleRoot(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeFocusItemChanged")
	}
}

//export callbackQQuickWindowActiveFocusItemChanged
func callbackQQuickWindowActiveFocusItemChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged").(func())()
}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "afterAnimating")
	}
}

//export callbackQQuickWindowAfterAnimating
func callbackQQuickWindowAfterAnimating(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterAnimating").(func())()
}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "afterRendering")
	}
}

//export callbackQQuickWindowAfterRendering
func callbackQQuickWindowAfterRendering(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterRendering").(func())()
}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "afterSynchronizing")
	}
}

//export callbackQQuickWindowAfterSynchronizing
func callbackQQuickWindowAfterSynchronizing(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "afterSynchronizing").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "beforeRendering")
	}
}

//export callbackQQuickWindowBeforeRendering
func callbackQQuickWindowBeforeRendering(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "beforeRendering").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "beforeSynchronizing")
	}
}

//export callbackQQuickWindowBeforeSynchronizing
func callbackQQuickWindowBeforeSynchronizing(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing").(func())()
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImageITF) *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QQuickWindow_CreateTextureFromImage2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQImage(image)))))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImageITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QQuickWindow_CreateTextureFromImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQImage(image)), C.int(options))))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
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
		return qml.QQmlIncubationControllerFromPointer(unsafe.Pointer(C.QQuickWindow_IncubationController(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickWindow_MouseGrabberItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	if ptr.Pointer() != nil {
		return gui.QOpenGLContextFromPointer(unsafe.Pointer(C.QQuickWindow_OpenglContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context gui.QOpenGLContextITF)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "openglContextCreated")
	}
}

//export callbackQQuickWindowOpenglContextCreated
func callbackQQuickWindowOpenglContextCreated(ptrName *C.char, context unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "openglContextCreated").(func(*gui.QOpenGLContext))(gui.QOpenGLContextFromPointer(context))
}

func (ptr *QQuickWindow) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	if ptr.Pointer() != nil {
		return gui.QOpenGLFramebufferObjectFromPointer(unsafe.Pointer(C.QQuickWindow_RenderTarget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop")
	}
}

//export callbackQQuickWindowSceneGraphAboutToStop
func callbackQQuickWindowSceneGraphAboutToStop(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWindowSceneGraphError
func callbackQQuickWindowSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphError").(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInitialized")
	}
}

//export callbackQQuickWindowSceneGraphInitialized
func callbackQQuickWindowSceneGraphInitialized(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInvalidated")
	}
}

//export callbackQQuickWindowSceneGraphInvalidated
func callbackQQuickWindowSceneGraphInvalidated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated").(func())()
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnableITF, stage QQuickWindow__RenderStage) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRunnable(job)), C.int(stage))
	}
}

func (ptr *QQuickWindow) SendEvent(item QQuickItemITF, e core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuickItem(item)), C.QtObjectPtr(core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.int(qt.GoBoolToInt(useAlpha)))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObjectITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQOpenGLFramebufferObject(fbo)))
	}
}

func (ptr *QQuickWindow) Update() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
