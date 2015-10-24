package gui

//#include "qopenglcontext.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLContext struct {
	core.QObject
}

type QOpenGLContextITF interface {
	core.QObjectITF
	QOpenGLContextPTR() *QOpenGLContext
}

func PointerFromQOpenGLContext(ptr QOpenGLContextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLContextPTR().Pointer()
	}
	return nil
}

func QOpenGLContextFromPointer(ptr unsafe.Pointer) *QOpenGLContext {
	var n = new(QOpenGLContext)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLContext_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLContext) QOpenGLContextPTR() *QOpenGLContext {
	return ptr
}

//QOpenGLContext::OpenGLModuleType
type QOpenGLContext__OpenGLModuleType int

var (
	QOpenGLContext__LibGL   = QOpenGLContext__OpenGLModuleType(0)
	QOpenGLContext__LibGLES = QOpenGLContext__OpenGLModuleType(1)
)

func NewQOpenGLContext(parent core.QObjectITF) *QOpenGLContext {
	return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLContext_NewQOpenGLContext(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLContext) ConnectAboutToBeDestroyed(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_ConnectAboutToBeDestroyed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToBeDestroyed", f)
	}
}

func (ptr *QOpenGLContext) DisconnectAboutToBeDestroyed() {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_DisconnectAboutToBeDestroyed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToBeDestroyed")
	}
}

//export callbackQOpenGLContextAboutToBeDestroyed
func callbackQOpenGLContextAboutToBeDestroyed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToBeDestroyed").(func())()
}

func QOpenGLContext_AreSharing(first QOpenGLContextITF, second QOpenGLContextITF) bool {
	return C.QOpenGLContext_QOpenGLContext_AreSharing(C.QtObjectPtr(PointerFromQOpenGLContext(first)), C.QtObjectPtr(PointerFromQOpenGLContext(second))) != 0
}

func (ptr *QOpenGLContext) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLContext_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QOpenGLContext_CurrentContext() *QOpenGLContext {
	return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLContext_QOpenGLContext_CurrentContext()))
}

func (ptr *QOpenGLContext) DoneCurrent() {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_DoneCurrent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLContext) Functions() *QOpenGLFunctions {
	if ptr.Pointer() != nil {
		return QOpenGLFunctionsFromPointer(unsafe.Pointer(C.QOpenGLContext_Functions(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QOpenGLContext_GlobalShareContext() *QOpenGLContext {
	return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLContext_QOpenGLContext_GlobalShareContext()))
}

func (ptr *QOpenGLContext) HasExtension(extension core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLContext_HasExtension(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(extension))) != 0
	}
	return false
}

func (ptr *QOpenGLContext) IsOpenGLES() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLContext_IsOpenGLES(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLContext) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLContext_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLContext) MakeCurrent(surface QSurfaceITF) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLContext_MakeCurrent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSurface(surface))) != 0
	}
	return false
}

func (ptr *QOpenGLContext) NativeHandle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QOpenGLContext_NativeHandle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QOpenGLContext_OpenGLModuleHandle() {
	C.QOpenGLContext_QOpenGLContext_OpenGLModuleHandle()
}

func QOpenGLContext_OpenGLModuleType() QOpenGLContext__OpenGLModuleType {
	return QOpenGLContext__OpenGLModuleType(C.QOpenGLContext_QOpenGLContext_OpenGLModuleType())
}

func (ptr *QOpenGLContext) SetFormat(format QSurfaceFormatITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSurfaceFormat(format)))
	}
}

func (ptr *QOpenGLContext) SetNativeHandle(handle string) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_SetNativeHandle(C.QtObjectPtr(ptr.Pointer()), C.CString(handle))
	}
}

func (ptr *QOpenGLContext) Screen() *QScreen {
	if ptr.Pointer() != nil {
		return QScreenFromPointer(unsafe.Pointer(C.QOpenGLContext_Screen(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLContext) SetScreen(screen QScreenITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_SetScreen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScreen(screen)))
	}
}

func (ptr *QOpenGLContext) SetShareContext(shareContext QOpenGLContextITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_SetShareContext(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLContext(shareContext)))
	}
}

func (ptr *QOpenGLContext) ShareContext() *QOpenGLContext {
	if ptr.Pointer() != nil {
		return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLContext_ShareContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLContext) ShareGroup() *QOpenGLContextGroup {
	if ptr.Pointer() != nil {
		return QOpenGLContextGroupFromPointer(unsafe.Pointer(C.QOpenGLContext_ShareGroup(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QOpenGLContext_SupportsThreadedOpenGL() bool {
	return C.QOpenGLContext_QOpenGLContext_SupportsThreadedOpenGL() != 0
}

func (ptr *QOpenGLContext) Surface() *QSurface {
	if ptr.Pointer() != nil {
		return QSurfaceFromPointer(unsafe.Pointer(C.QOpenGLContext_Surface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLContext) SwapBuffers(surface QSurfaceITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_SwapBuffers(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSurface(surface)))
	}
}

func (ptr *QOpenGLContext) VersionFunctions(versionProfile QOpenGLVersionProfileITF) *QAbstractOpenGLFunctions {
	if ptr.Pointer() != nil {
		return QAbstractOpenGLFunctionsFromPointer(unsafe.Pointer(C.QOpenGLContext_VersionFunctions(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLVersionProfile(versionProfile)))))
	}
	return nil
}

func (ptr *QOpenGLContext) DestroyQOpenGLContext() {
	if ptr.Pointer() != nil {
		C.QOpenGLContext_DestroyQOpenGLContext(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
