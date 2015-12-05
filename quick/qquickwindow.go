package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"log"
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
	for len(n.ObjectName()) < len("QQuickWindow_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::activeFocusItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::color")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::contentItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::QQuickWindow")
		}
	}()

	return NewQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::accessibleRoot")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::activeFocusItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::activeFocusItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeFocusItemChanged")
	}
}

//export callbackQQuickWindowActiveFocusItemChanged
func callbackQQuickWindowActiveFocusItemChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::activeFocusItemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged").(func())()
}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterAnimating")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterAnimating")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterAnimating")
	}
}

//export callbackQQuickWindowAfterAnimating
func callbackQQuickWindowAfterAnimating(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterAnimating")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "afterAnimating").(func())()
}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterRendering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterRendering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterRendering")
	}
}

//export callbackQQuickWindowAfterRendering
func callbackQQuickWindowAfterRendering(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterRendering")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "afterRendering").(func())()
}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterSynchronizing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterSynchronizing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterSynchronizing")
	}
}

//export callbackQQuickWindowAfterSynchronizing
func callbackQQuickWindowAfterSynchronizing(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::afterSynchronizing")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "afterSynchronizing").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeRendering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeRendering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeRendering")
	}
}

//export callbackQQuickWindowBeforeRendering
func callbackQQuickWindowBeforeRendering(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeRendering")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "beforeRendering").(func())()
}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeSynchronizing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeSynchronizing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeSynchronizing")
	}
}

//export callbackQQuickWindowBeforeSynchronizing
func callbackQQuickWindowBeforeSynchronizing(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::beforeSynchronizing")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing").(func())()
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::clearBeforeRendering")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectColorChanged(f func(v *gui.QColor)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::colorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::colorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQQuickWindowColorChanged
func callbackQQuickWindowColorChanged(ptrName *C.char, v unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::colorChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "colorChanged").(func(*gui.QColor))(gui.NewQColorFromPointer(v))
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::createTextureFromImage")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::createTextureFromImage")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::effectiveDevicePixelRatio")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::frameSwapped")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::frameSwapped")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQQuickWindowFrameSwapped
func callbackQQuickWindowFrameSwapped(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::frameSwapped")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "frameSwapped").(func())()
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::hasDefaultAlphaBuffer")
		}
	}()

	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::incubationController")
		}
	}()

	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::isPersistentOpenGLContext")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::isPersistentSceneGraph")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::isSceneGraphInitialized")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::mouseGrabberItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::openglContext")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::openglContextCreated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::openglContextCreated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "openglContextCreated")
	}
}

//export callbackQQuickWindowOpenglContextCreated
func callbackQQuickWindowOpenglContextCreated(ptrName *C.char, context unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::openglContextCreated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "openglContextCreated").(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
}

func (ptr *QQuickWindow) ReleaseResources() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::releaseResources")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::renderTarget")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::resetOpenGLState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphAboutToStop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphAboutToStop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop")
	}
}

//export callbackQQuickWindowSceneGraphAboutToStop
func callbackQQuickWindowSceneGraphAboutToStop(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphAboutToStop")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWindowSceneGraphError
func callbackQQuickWindowSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphError")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneGraphError").(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInitialized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInitialized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInitialized")
	}
}

//export callbackQQuickWindowSceneGraphInitialized
func callbackQQuickWindowSceneGraphInitialized(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInitialized")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized").(func())()
}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInvalidated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInvalidated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInvalidated")
	}
}

//export callbackQQuickWindowSceneGraphInvalidated
func callbackQQuickWindowSceneGraphInvalidated(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sceneGraphInvalidated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated").(func())()
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::scheduleRenderJob")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(ptr.Pointer(), core.PointerFromQRunnable(job), C.int(stage))
	}
}

func (ptr *QQuickWindow) SendEvent(item QQuickItem_ITF, e core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::sendEvent")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setClearBeforeRendering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setDefaultAlphaBuffer")
		}
	}()

	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.int(qt.GoBoolToInt(useAlpha)))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setPersistentOpenGLContext")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setPersistentSceneGraph")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::setRenderTarget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(ptr.Pointer(), gui.PointerFromQOpenGLFramebufferObject(fbo))
	}
}

func (ptr *QQuickWindow) Update() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::update")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWindow::~QQuickWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
