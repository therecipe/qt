// +build !minimal

package quick

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type BacktraceJob struct {
	CollectJob
}

type BacktraceJob_ITF interface {
	CollectJob_ITF
	BacktraceJob_PTR() *BacktraceJob
}

func (ptr *BacktraceJob) BacktraceJob_PTR() *BacktraceJob {
	return ptr
}

func (ptr *BacktraceJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.CollectJob_PTR().Pointer()
	}
	return nil
}

func (ptr *BacktraceJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.CollectJob_PTR().SetPointer(p)
	}
}

func PointerFromBacktraceJob(ptr BacktraceJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.BacktraceJob_PTR().Pointer()
	}
	return nil
}

func (n *BacktraceJob) InitFromInternal(ptr uintptr, name string) {
	n.CollectJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *BacktraceJob) ClassNameInternalF() string {
	return n.CollectJob_PTR().ClassNameInternalF()
}

func NewBacktraceJobFromPointer(ptr unsafe.Pointer) (n *BacktraceJob) {
	n = new(BacktraceJob)
	n.InitFromInternal(uintptr(ptr), "quick.BacktraceJob")
	return
}

func (ptr *BacktraceJob) DestroyBacktraceJob() {
}

type CollectJob struct {
	QV4DebugJob
}

type CollectJob_ITF interface {
	QV4DebugJob_ITF
	CollectJob_PTR() *CollectJob
}

func (ptr *CollectJob) CollectJob_PTR() *CollectJob {
	return ptr
}

func (ptr *CollectJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebugJob_PTR().Pointer()
	}
	return nil
}

func (ptr *CollectJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QV4DebugJob_PTR().SetPointer(p)
	}
}

func PointerFromCollectJob(ptr CollectJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CollectJob_PTR().Pointer()
	}
	return nil
}

func (n *CollectJob) InitFromInternal(ptr uintptr, name string) {
	n.QV4DebugJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *CollectJob) ClassNameInternalF() string {
	return n.QV4DebugJob_PTR().ClassNameInternalF()
}

func NewCollectJobFromPointer(ptr unsafe.Pointer) (n *CollectJob) {
	n = new(CollectJob)
	n.InitFromInternal(uintptr(ptr), "quick.CollectJob")
	return
}

func (ptr *CollectJob) DestroyCollectJob() {
}

type EvalJob struct {
	JavaScriptJob
}

type EvalJob_ITF interface {
	JavaScriptJob_ITF
	EvalJob_PTR() *EvalJob
}

func (ptr *EvalJob) EvalJob_PTR() *EvalJob {
	return ptr
}

func (ptr *EvalJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.JavaScriptJob_PTR().Pointer()
	}
	return nil
}

func (ptr *EvalJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.JavaScriptJob_PTR().SetPointer(p)
	}
}

func PointerFromEvalJob(ptr EvalJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.EvalJob_PTR().Pointer()
	}
	return nil
}

func (n *EvalJob) InitFromInternal(ptr uintptr, name string) {
	n.JavaScriptJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *EvalJob) ClassNameInternalF() string {
	return n.JavaScriptJob_PTR().ClassNameInternalF()
}

func NewEvalJobFromPointer(ptr unsafe.Pointer) (n *EvalJob) {
	n = new(EvalJob)
	n.InitFromInternal(uintptr(ptr), "quick.EvalJob")
	return
}

func (ptr *EvalJob) DestroyEvalJob() {
}

type ExpressionEvalJob struct {
	JavaScriptJob
}

type ExpressionEvalJob_ITF interface {
	JavaScriptJob_ITF
	ExpressionEvalJob_PTR() *ExpressionEvalJob
}

func (ptr *ExpressionEvalJob) ExpressionEvalJob_PTR() *ExpressionEvalJob {
	return ptr
}

func (ptr *ExpressionEvalJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.JavaScriptJob_PTR().Pointer()
	}
	return nil
}

func (ptr *ExpressionEvalJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.JavaScriptJob_PTR().SetPointer(p)
	}
}

func PointerFromExpressionEvalJob(ptr ExpressionEvalJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ExpressionEvalJob_PTR().Pointer()
	}
	return nil
}

func (n *ExpressionEvalJob) InitFromInternal(ptr uintptr, name string) {
	n.JavaScriptJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *ExpressionEvalJob) ClassNameInternalF() string {
	return n.JavaScriptJob_PTR().ClassNameInternalF()
}

func NewExpressionEvalJobFromPointer(ptr unsafe.Pointer) (n *ExpressionEvalJob) {
	n = new(ExpressionEvalJob)
	n.InitFromInternal(uintptr(ptr), "quick.ExpressionEvalJob")
	return
}

func (ptr *ExpressionEvalJob) DestroyExpressionEvalJob() {
}

type FrameJob struct {
	CollectJob
}

type FrameJob_ITF interface {
	CollectJob_ITF
	FrameJob_PTR() *FrameJob
}

func (ptr *FrameJob) FrameJob_PTR() *FrameJob {
	return ptr
}

func (ptr *FrameJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.CollectJob_PTR().Pointer()
	}
	return nil
}

func (ptr *FrameJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.CollectJob_PTR().SetPointer(p)
	}
}

func PointerFromFrameJob(ptr FrameJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FrameJob_PTR().Pointer()
	}
	return nil
}

func (n *FrameJob) InitFromInternal(ptr uintptr, name string) {
	n.CollectJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *FrameJob) ClassNameInternalF() string {
	return n.CollectJob_PTR().ClassNameInternalF()
}

func NewFrameJobFromPointer(ptr unsafe.Pointer) (n *FrameJob) {
	n = new(FrameJob)
	n.InitFromInternal(uintptr(ptr), "quick.FrameJob")
	return
}

func (ptr *FrameJob) DestroyFrameJob() {
}

type GatherSourcesJob struct {
	QV4DebugJob
}

type GatherSourcesJob_ITF interface {
	QV4DebugJob_ITF
	GatherSourcesJob_PTR() *GatherSourcesJob
}

func (ptr *GatherSourcesJob) GatherSourcesJob_PTR() *GatherSourcesJob {
	return ptr
}

func (ptr *GatherSourcesJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebugJob_PTR().Pointer()
	}
	return nil
}

func (ptr *GatherSourcesJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QV4DebugJob_PTR().SetPointer(p)
	}
}

func PointerFromGatherSourcesJob(ptr GatherSourcesJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.GatherSourcesJob_PTR().Pointer()
	}
	return nil
}

func (n *GatherSourcesJob) InitFromInternal(ptr uintptr, name string) {
	n.QV4DebugJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *GatherSourcesJob) ClassNameInternalF() string {
	return n.QV4DebugJob_PTR().ClassNameInternalF()
}

func NewGatherSourcesJobFromPointer(ptr unsafe.Pointer) (n *GatherSourcesJob) {
	n = new(GatherSourcesJob)
	n.InitFromInternal(uintptr(ptr), "quick.GatherSourcesJob")
	return
}

func (ptr *GatherSourcesJob) DestroyGatherSourcesJob() {
}

type JavaScriptJob struct {
	QV4DebugJob
}

type JavaScriptJob_ITF interface {
	QV4DebugJob_ITF
	JavaScriptJob_PTR() *JavaScriptJob
}

func (ptr *JavaScriptJob) JavaScriptJob_PTR() *JavaScriptJob {
	return ptr
}

func (ptr *JavaScriptJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebugJob_PTR().Pointer()
	}
	return nil
}

func (ptr *JavaScriptJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QV4DebugJob_PTR().SetPointer(p)
	}
}

func PointerFromJavaScriptJob(ptr JavaScriptJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.JavaScriptJob_PTR().Pointer()
	}
	return nil
}

func (n *JavaScriptJob) InitFromInternal(ptr uintptr, name string) {
	n.QV4DebugJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *JavaScriptJob) ClassNameInternalF() string {
	return n.QV4DebugJob_PTR().ClassNameInternalF()
}

func NewJavaScriptJobFromPointer(ptr unsafe.Pointer) (n *JavaScriptJob) {
	n = new(JavaScriptJob)
	n.InitFromInternal(uintptr(ptr), "quick.JavaScriptJob")
	return
}

func (ptr *JavaScriptJob) DestroyJavaScriptJob() {
}

type QDebugMessageServiceFactory struct {
	internal.Internal
}

type QDebugMessageServiceFactory_ITF interface {
	QDebugMessageServiceFactory_PTR() *QDebugMessageServiceFactory
}

func (ptr *QDebugMessageServiceFactory) QDebugMessageServiceFactory_PTR() *QDebugMessageServiceFactory {
	return ptr
}

func (ptr *QDebugMessageServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDebugMessageServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDebugMessageServiceFactory(ptr QDebugMessageServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebugMessageServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QDebugMessageServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDebugMessageServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QDebugMessageServiceFactory) {
	n = new(QDebugMessageServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QDebugMessageServiceFactory")
	return
}

func (ptr *QDebugMessageServiceFactory) DestroyQDebugMessageServiceFactory() {
}

type QDebugMessageServiceImpl struct {
	internal.Internal
}

type QDebugMessageServiceImpl_ITF interface {
	QDebugMessageServiceImpl_PTR() *QDebugMessageServiceImpl
}

func (ptr *QDebugMessageServiceImpl) QDebugMessageServiceImpl_PTR() *QDebugMessageServiceImpl {
	return ptr
}

func (ptr *QDebugMessageServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDebugMessageServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDebugMessageServiceImpl(ptr QDebugMessageServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebugMessageServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QDebugMessageServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDebugMessageServiceImplFromPointer(ptr unsafe.Pointer) (n *QDebugMessageServiceImpl) {
	n = new(QDebugMessageServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QDebugMessageServiceImpl")
	return
}

func (ptr *QDebugMessageServiceImpl) DestroyQDebugMessageServiceImpl() {
}

type QLocalClientConnectionFactory struct {
	internal.Internal
}

type QLocalClientConnectionFactory_ITF interface {
	QLocalClientConnectionFactory_PTR() *QLocalClientConnectionFactory
}

func (ptr *QLocalClientConnectionFactory) QLocalClientConnectionFactory_PTR() *QLocalClientConnectionFactory {
	return ptr
}

func (ptr *QLocalClientConnectionFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLocalClientConnectionFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLocalClientConnectionFactory(ptr QLocalClientConnectionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalClientConnectionFactory_PTR().Pointer()
	}
	return nil
}

func (n *QLocalClientConnectionFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLocalClientConnectionFactoryFromPointer(ptr unsafe.Pointer) (n *QLocalClientConnectionFactory) {
	n = new(QLocalClientConnectionFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QLocalClientConnectionFactory")
	return
}

func (ptr *QLocalClientConnectionFactory) DestroyQLocalClientConnectionFactory() {
}

type QOpenVGMatrix struct {
	internal.Internal
}

type QOpenVGMatrix_ITF interface {
	QOpenVGMatrix_PTR() *QOpenVGMatrix
}

func (ptr *QOpenVGMatrix) QOpenVGMatrix_PTR() *QOpenVGMatrix {
	return ptr
}

func (ptr *QOpenVGMatrix) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QOpenVGMatrix) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQOpenVGMatrix(ptr QOpenVGMatrix_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenVGMatrix_PTR().Pointer()
	}
	return nil
}

func (n *QOpenVGMatrix) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQOpenVGMatrixFromPointer(ptr unsafe.Pointer) (n *QOpenVGMatrix) {
	n = new(QOpenVGMatrix)
	n.InitFromInternal(uintptr(ptr), "quick.QOpenVGMatrix")
	return
}

func (ptr *QOpenVGMatrix) DestroyQOpenVGMatrix() {
}

type QOpenVGOffscreenSurface struct {
	internal.Internal
}

type QOpenVGOffscreenSurface_ITF interface {
	QOpenVGOffscreenSurface_PTR() *QOpenVGOffscreenSurface
}

func (ptr *QOpenVGOffscreenSurface) QOpenVGOffscreenSurface_PTR() *QOpenVGOffscreenSurface {
	return ptr
}

func (ptr *QOpenVGOffscreenSurface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QOpenVGOffscreenSurface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQOpenVGOffscreenSurface(ptr QOpenVGOffscreenSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenVGOffscreenSurface_PTR().Pointer()
	}
	return nil
}

func (n *QOpenVGOffscreenSurface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQOpenVGOffscreenSurfaceFromPointer(ptr unsafe.Pointer) (n *QOpenVGOffscreenSurface) {
	n = new(QOpenVGOffscreenSurface)
	n.InitFromInternal(uintptr(ptr), "quick.QOpenVGOffscreenSurface")
	return
}

func (ptr *QOpenVGOffscreenSurface) DestroyQOpenVGOffscreenSurface() {
}

type QQmlDebugServerFactory struct {
	internal.Internal
}

type QQmlDebugServerFactory_ITF interface {
	QQmlDebugServerFactory_PTR() *QQmlDebugServerFactory
}

func (ptr *QQmlDebugServerFactory) QQmlDebugServerFactory_PTR() *QQmlDebugServerFactory {
	return ptr
}

func (ptr *QQmlDebugServerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlDebugServerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlDebugServerFactory(ptr QQmlDebugServerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlDebugServerFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlDebugServerFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlDebugServerFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlDebugServerFactory) {
	n = new(QQmlDebugServerFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlDebugServerFactory")
	return
}

func (ptr *QQmlDebugServerFactory) DestroyQQmlDebugServerFactory() {
}

type QQmlDebuggerServiceFactory struct {
	internal.Internal
}

type QQmlDebuggerServiceFactory_ITF interface {
	QQmlDebuggerServiceFactory_PTR() *QQmlDebuggerServiceFactory
}

func (ptr *QQmlDebuggerServiceFactory) QQmlDebuggerServiceFactory_PTR() *QQmlDebuggerServiceFactory {
	return ptr
}

func (ptr *QQmlDebuggerServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlDebuggerServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlDebuggerServiceFactory(ptr QQmlDebuggerServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlDebuggerServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlDebuggerServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlDebuggerServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlDebuggerServiceFactory) {
	n = new(QQmlDebuggerServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlDebuggerServiceFactory")
	return
}

func (ptr *QQmlDebuggerServiceFactory) DestroyQQmlDebuggerServiceFactory() {
}

type QQmlEngineControlServiceImpl struct {
	internal.Internal
}

type QQmlEngineControlServiceImpl_ITF interface {
	QQmlEngineControlServiceImpl_PTR() *QQmlEngineControlServiceImpl
}

func (ptr *QQmlEngineControlServiceImpl) QQmlEngineControlServiceImpl_PTR() *QQmlEngineControlServiceImpl {
	return ptr
}

func (ptr *QQmlEngineControlServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlEngineControlServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlEngineControlServiceImpl(ptr QQmlEngineControlServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngineControlServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QQmlEngineControlServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlEngineControlServiceImplFromPointer(ptr unsafe.Pointer) (n *QQmlEngineControlServiceImpl) {
	n = new(QQmlEngineControlServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlEngineControlServiceImpl")
	return
}

func (ptr *QQmlEngineControlServiceImpl) DestroyQQmlEngineControlServiceImpl() {
}

type QQmlEngineDebugServiceImpl struct {
	internal.Internal
}

type QQmlEngineDebugServiceImpl_ITF interface {
	QQmlEngineDebugServiceImpl_PTR() *QQmlEngineDebugServiceImpl
}

func (ptr *QQmlEngineDebugServiceImpl) QQmlEngineDebugServiceImpl_PTR() *QQmlEngineDebugServiceImpl {
	return ptr
}

func (ptr *QQmlEngineDebugServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlEngineDebugServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlEngineDebugServiceImpl(ptr QQmlEngineDebugServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngineDebugServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QQmlEngineDebugServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlEngineDebugServiceImplFromPointer(ptr unsafe.Pointer) (n *QQmlEngineDebugServiceImpl) {
	n = new(QQmlEngineDebugServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlEngineDebugServiceImpl")
	return
}

func (ptr *QQmlEngineDebugServiceImpl) DestroyQQmlEngineDebugServiceImpl() {
}

type QQmlInspectorServiceFactory struct {
	internal.Internal
}

type QQmlInspectorServiceFactory_ITF interface {
	QQmlInspectorServiceFactory_PTR() *QQmlInspectorServiceFactory
}

func (ptr *QQmlInspectorServiceFactory) QQmlInspectorServiceFactory_PTR() *QQmlInspectorServiceFactory {
	return ptr
}

func (ptr *QQmlInspectorServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlInspectorServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlInspectorServiceFactory(ptr QQmlInspectorServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlInspectorServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlInspectorServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlInspectorServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlInspectorServiceFactory) {
	n = new(QQmlInspectorServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlInspectorServiceFactory")
	return
}

func (ptr *QQmlInspectorServiceFactory) DestroyQQmlInspectorServiceFactory() {
}

type QQmlNativeDebugConnector struct {
	internal.Internal
}

type QQmlNativeDebugConnector_ITF interface {
	QQmlNativeDebugConnector_PTR() *QQmlNativeDebugConnector
}

func (ptr *QQmlNativeDebugConnector) QQmlNativeDebugConnector_PTR() *QQmlNativeDebugConnector {
	return ptr
}

func (ptr *QQmlNativeDebugConnector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlNativeDebugConnector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlNativeDebugConnector(ptr QQmlNativeDebugConnector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNativeDebugConnector_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNativeDebugConnector) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlNativeDebugConnectorFromPointer(ptr unsafe.Pointer) (n *QQmlNativeDebugConnector) {
	n = new(QQmlNativeDebugConnector)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlNativeDebugConnector")
	return
}

func (ptr *QQmlNativeDebugConnector) DestroyQQmlNativeDebugConnector() {
}

type QQmlNativeDebugConnectorFactory struct {
	internal.Internal
}

type QQmlNativeDebugConnectorFactory_ITF interface {
	QQmlNativeDebugConnectorFactory_PTR() *QQmlNativeDebugConnectorFactory
}

func (ptr *QQmlNativeDebugConnectorFactory) QQmlNativeDebugConnectorFactory_PTR() *QQmlNativeDebugConnectorFactory {
	return ptr
}

func (ptr *QQmlNativeDebugConnectorFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlNativeDebugConnectorFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlNativeDebugConnectorFactory(ptr QQmlNativeDebugConnectorFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNativeDebugConnectorFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNativeDebugConnectorFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlNativeDebugConnectorFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlNativeDebugConnectorFactory) {
	n = new(QQmlNativeDebugConnectorFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlNativeDebugConnectorFactory")
	return
}

func (ptr *QQmlNativeDebugConnectorFactory) DestroyQQmlNativeDebugConnectorFactory() {
}

type QQmlNativeDebugServiceFactory struct {
	internal.Internal
}

type QQmlNativeDebugServiceFactory_ITF interface {
	QQmlNativeDebugServiceFactory_PTR() *QQmlNativeDebugServiceFactory
}

func (ptr *QQmlNativeDebugServiceFactory) QQmlNativeDebugServiceFactory_PTR() *QQmlNativeDebugServiceFactory {
	return ptr
}

func (ptr *QQmlNativeDebugServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlNativeDebugServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlNativeDebugServiceFactory(ptr QQmlNativeDebugServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNativeDebugServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNativeDebugServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlNativeDebugServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlNativeDebugServiceFactory) {
	n = new(QQmlNativeDebugServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlNativeDebugServiceFactory")
	return
}

func (ptr *QQmlNativeDebugServiceFactory) DestroyQQmlNativeDebugServiceFactory() {
}

type QQmlNativeDebugServiceImpl struct {
	internal.Internal
}

type QQmlNativeDebugServiceImpl_ITF interface {
	QQmlNativeDebugServiceImpl_PTR() *QQmlNativeDebugServiceImpl
}

func (ptr *QQmlNativeDebugServiceImpl) QQmlNativeDebugServiceImpl_PTR() *QQmlNativeDebugServiceImpl {
	return ptr
}

func (ptr *QQmlNativeDebugServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlNativeDebugServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlNativeDebugServiceImpl(ptr QQmlNativeDebugServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNativeDebugServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNativeDebugServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlNativeDebugServiceImplFromPointer(ptr unsafe.Pointer) (n *QQmlNativeDebugServiceImpl) {
	n = new(QQmlNativeDebugServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlNativeDebugServiceImpl")
	return
}

func (ptr *QQmlNativeDebugServiceImpl) DestroyQQmlNativeDebugServiceImpl() {
}

type QQmlPreviewBlacklist struct {
	internal.Internal
}

type QQmlPreviewBlacklist_ITF interface {
	QQmlPreviewBlacklist_PTR() *QQmlPreviewBlacklist
}

func (ptr *QQmlPreviewBlacklist) QQmlPreviewBlacklist_PTR() *QQmlPreviewBlacklist {
	return ptr
}

func (ptr *QQmlPreviewBlacklist) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewBlacklist) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewBlacklist(ptr QQmlPreviewBlacklist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewBlacklist_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewBlacklist) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewBlacklistFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewBlacklist) {
	n = new(QQmlPreviewBlacklist)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewBlacklist")
	return
}

func (ptr *QQmlPreviewBlacklist) DestroyQQmlPreviewBlacklist() {
}

type QQmlPreviewFileEngine struct {
	internal.Internal
}

type QQmlPreviewFileEngine_ITF interface {
	QQmlPreviewFileEngine_PTR() *QQmlPreviewFileEngine
}

func (ptr *QQmlPreviewFileEngine) QQmlPreviewFileEngine_PTR() *QQmlPreviewFileEngine {
	return ptr
}

func (ptr *QQmlPreviewFileEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewFileEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewFileEngine(ptr QQmlPreviewFileEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewFileEngine_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewFileEngine) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewFileEngineFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewFileEngine) {
	n = new(QQmlPreviewFileEngine)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewFileEngine")
	return
}

func (ptr *QQmlPreviewFileEngine) DestroyQQmlPreviewFileEngine() {
}

type QQmlPreviewFileEngineHandler struct {
	internal.Internal
}

type QQmlPreviewFileEngineHandler_ITF interface {
	QQmlPreviewFileEngineHandler_PTR() *QQmlPreviewFileEngineHandler
}

func (ptr *QQmlPreviewFileEngineHandler) QQmlPreviewFileEngineHandler_PTR() *QQmlPreviewFileEngineHandler {
	return ptr
}

func (ptr *QQmlPreviewFileEngineHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewFileEngineHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewFileEngineHandler(ptr QQmlPreviewFileEngineHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewFileEngineHandler_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewFileEngineHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewFileEngineHandlerFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewFileEngineHandler) {
	n = new(QQmlPreviewFileEngineHandler)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewFileEngineHandler")
	return
}

func (ptr *QQmlPreviewFileEngineHandler) DestroyQQmlPreviewFileEngineHandler() {
}

type QQmlPreviewFileLoader struct {
	core.QObject
}

type QQmlPreviewFileLoader_ITF interface {
	core.QObject_ITF
	QQmlPreviewFileLoader_PTR() *QQmlPreviewFileLoader
}

func (ptr *QQmlPreviewFileLoader) QQmlPreviewFileLoader_PTR() *QQmlPreviewFileLoader {
	return ptr
}

func (ptr *QQmlPreviewFileLoader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlPreviewFileLoader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlPreviewFileLoader(ptr QQmlPreviewFileLoader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewFileLoader_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewFileLoader) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlPreviewFileLoader) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlPreviewFileLoaderFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewFileLoader) {
	n = new(QQmlPreviewFileLoader)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewFileLoader")
	return
}

func (ptr *QQmlPreviewFileLoader) DestroyQQmlPreviewFileLoader() {
}

type QQmlPreviewHandler struct {
	core.QObject
}

type QQmlPreviewHandler_ITF interface {
	core.QObject_ITF
	QQmlPreviewHandler_PTR() *QQmlPreviewHandler
}

func (ptr *QQmlPreviewHandler) QQmlPreviewHandler_PTR() *QQmlPreviewHandler {
	return ptr
}

func (ptr *QQmlPreviewHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlPreviewHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlPreviewHandler(ptr QQmlPreviewHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewHandler_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewHandler) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlPreviewHandler) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlPreviewHandlerFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewHandler) {
	n = new(QQmlPreviewHandler)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewHandler")
	return
}

func (ptr *QQmlPreviewHandler) DestroyQQmlPreviewHandler() {
}

type QQmlPreviewPosition struct {
	internal.Internal
}

type QQmlPreviewPosition_ITF interface {
	QQmlPreviewPosition_PTR() *QQmlPreviewPosition
}

func (ptr *QQmlPreviewPosition) QQmlPreviewPosition_PTR() *QQmlPreviewPosition {
	return ptr
}

func (ptr *QQmlPreviewPosition) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewPosition) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewPosition(ptr QQmlPreviewPosition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewPosition_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewPosition) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewPositionFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewPosition) {
	n = new(QQmlPreviewPosition)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewPosition")
	return
}

func (ptr *QQmlPreviewPosition) DestroyQQmlPreviewPosition() {
}

type QQmlPreviewServiceFactory struct {
	internal.Internal
}

type QQmlPreviewServiceFactory_ITF interface {
	QQmlPreviewServiceFactory_PTR() *QQmlPreviewServiceFactory
}

func (ptr *QQmlPreviewServiceFactory) QQmlPreviewServiceFactory_PTR() *QQmlPreviewServiceFactory {
	return ptr
}

func (ptr *QQmlPreviewServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewServiceFactory(ptr QQmlPreviewServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewServiceFactory) {
	n = new(QQmlPreviewServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewServiceFactory")
	return
}

func (ptr *QQmlPreviewServiceFactory) DestroyQQmlPreviewServiceFactory() {
}

type QQmlPreviewServiceImpl struct {
	internal.Internal
}

type QQmlPreviewServiceImpl_ITF interface {
	QQmlPreviewServiceImpl_PTR() *QQmlPreviewServiceImpl
}

func (ptr *QQmlPreviewServiceImpl) QQmlPreviewServiceImpl_PTR() *QQmlPreviewServiceImpl {
	return ptr
}

func (ptr *QQmlPreviewServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPreviewServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPreviewServiceImpl(ptr QQmlPreviewServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPreviewServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPreviewServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPreviewServiceImplFromPointer(ptr unsafe.Pointer) (n *QQmlPreviewServiceImpl) {
	n = new(QQmlPreviewServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlPreviewServiceImpl")
	return
}

func (ptr *QQmlPreviewServiceImpl) DestroyQQmlPreviewServiceImpl() {
}

type QQmlProfilerAdapter struct {
	internal.Internal
}

type QQmlProfilerAdapter_ITF interface {
	QQmlProfilerAdapter_PTR() *QQmlProfilerAdapter
}

func (ptr *QQmlProfilerAdapter) QQmlProfilerAdapter_PTR() *QQmlProfilerAdapter {
	return ptr
}

func (ptr *QQmlProfilerAdapter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlProfilerAdapter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlProfilerAdapter(ptr QQmlProfilerAdapter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProfilerAdapter_PTR().Pointer()
	}
	return nil
}

func (n *QQmlProfilerAdapter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlProfilerAdapterFromPointer(ptr unsafe.Pointer) (n *QQmlProfilerAdapter) {
	n = new(QQmlProfilerAdapter)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlProfilerAdapter")
	return
}

func (ptr *QQmlProfilerAdapter) DestroyQQmlProfilerAdapter() {
}

type QQmlProfilerServiceFactory struct {
	internal.Internal
}

type QQmlProfilerServiceFactory_ITF interface {
	QQmlProfilerServiceFactory_PTR() *QQmlProfilerServiceFactory
}

func (ptr *QQmlProfilerServiceFactory) QQmlProfilerServiceFactory_PTR() *QQmlProfilerServiceFactory {
	return ptr
}

func (ptr *QQmlProfilerServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlProfilerServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlProfilerServiceFactory(ptr QQmlProfilerServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProfilerServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlProfilerServiceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlProfilerServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlProfilerServiceFactory) {
	n = new(QQmlProfilerServiceFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlProfilerServiceFactory")
	return
}

func (ptr *QQmlProfilerServiceFactory) DestroyQQmlProfilerServiceFactory() {
}

type QQmlProfilerServiceImpl struct {
	internal.Internal
}

type QQmlProfilerServiceImpl_ITF interface {
	QQmlProfilerServiceImpl_PTR() *QQmlProfilerServiceImpl
}

func (ptr *QQmlProfilerServiceImpl) QQmlProfilerServiceImpl_PTR() *QQmlProfilerServiceImpl {
	return ptr
}

func (ptr *QQmlProfilerServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlProfilerServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlProfilerServiceImpl(ptr QQmlProfilerServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProfilerServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QQmlProfilerServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlProfilerServiceImplFromPointer(ptr unsafe.Pointer) (n *QQmlProfilerServiceImpl) {
	n = new(QQmlProfilerServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlProfilerServiceImpl")
	return
}

func (ptr *QQmlProfilerServiceImpl) DestroyQQmlProfilerServiceImpl() {
}

type QQmlWatcher struct {
	core.QObject
}

type QQmlWatcher_ITF interface {
	core.QObject_ITF
	QQmlWatcher_PTR() *QQmlWatcher
}

func (ptr *QQmlWatcher) QQmlWatcher_PTR() *QQmlWatcher {
	return ptr
}

func (ptr *QQmlWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWatcher(ptr QQmlWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWatcher_PTR().Pointer()
	}
	return nil
}

func (n *QQmlWatcher) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlWatcher) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlWatcherFromPointer(ptr unsafe.Pointer) (n *QQmlWatcher) {
	n = new(QQmlWatcher)
	n.InitFromInternal(uintptr(ptr), "quick.QQmlWatcher")
	return
}

func (ptr *QQmlWatcher) DestroyQQmlWatcher() {
}

type QQuickAsyncImageProvider struct {
	QQuickImageProvider
}

type QQuickAsyncImageProvider_ITF interface {
	QQuickImageProvider_ITF
	QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider
}

func (ptr *QQuickAsyncImageProvider) QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider {
	return ptr
}

func (ptr *QQuickAsyncImageProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickImageProvider_PTR().SetPointer(p)
	}
}

func PointerFromQQuickAsyncImageProvider(ptr QQuickAsyncImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickAsyncImageProvider_PTR().Pointer()
	}
	return nil
}

func (n *QQuickAsyncImageProvider) InitFromInternal(ptr uintptr, name string) {
	n.QQuickImageProvider_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickAsyncImageProvider) ClassNameInternalF() string {
	return n.QQuickImageProvider_PTR().ClassNameInternalF()
}

func NewQQuickAsyncImageProviderFromPointer(ptr unsafe.Pointer) (n *QQuickAsyncImageProvider) {
	n = new(QQuickAsyncImageProvider)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickAsyncImageProvider")
	return
}
func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickAsyncImageProvider", ""}).(*QQuickAsyncImageProvider)
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestImageResponse(f func(id string, requestedSize *core.QSize) *QQuickImageResponse) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestImageResponse", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestImageResponse() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestImageResponse"})
}

func (ptr *QQuickAsyncImageProvider) RequestImageResponse(id string, requestedSize core.QSize_ITF) *QQuickImageResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestImageResponse", id, requestedSize}).(*QQuickImageResponse)
}

func (ptr *QQuickAsyncImageProvider) ConnectDestroyQQuickAsyncImageProvider(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickAsyncImageProvider", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickAsyncImageProvider) DisconnectDestroyQQuickAsyncImageProvider() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickAsyncImageProvider"})
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProvider() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickAsyncImageProvider"})
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProviderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickAsyncImageProviderDefault"})
}

type QQuickFolderListModel struct {
	core.QAbstractListModel
	qml.QQmlParserStatus
}

type QQuickFolderListModel_ITF interface {
	core.QAbstractListModel_ITF
	qml.QQmlParserStatus_ITF
	QQuickFolderListModel_PTR() *QQuickFolderListModel
}

func (ptr *QQuickFolderListModel) QQuickFolderListModel_PTR() *QQuickFolderListModel {
	return ptr
}

func (ptr *QQuickFolderListModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickFolderListModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
		ptr.QQmlParserStatus_PTR().SetPointer(p)
	}
}

func PointerFromQQuickFolderListModel(ptr QQuickFolderListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFolderListModel_PTR().Pointer()
	}
	return nil
}

func (n *QQuickFolderListModel) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractListModel_PTR().InitFromInternal(uintptr(ptr), name)
	n.QQmlParserStatus_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickFolderListModel) ClassNameInternalF() string {
	return n.QAbstractListModel_PTR().ClassNameInternalF()
}

func NewQQuickFolderListModelFromPointer(ptr unsafe.Pointer) (n *QQuickFolderListModel) {
	n = new(QQuickFolderListModel)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickFolderListModel")
	return
}

func (ptr *QQuickFolderListModel) DestroyQQuickFolderListModel() {
}

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickFramebufferObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromQQuickFramebufferObject(ptr QQuickFramebufferObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFramebufferObject_PTR().Pointer()
	}
	return nil
}

func (n *QQuickFramebufferObject) InitFromInternal(ptr uintptr, name string) {
	n.QQuickItem_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickFramebufferObject) ClassNameInternalF() string {
	return n.QQuickItem_PTR().ClassNameInternalF()
}

func NewQQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) (n *QQuickFramebufferObject) {
	n = new(QQuickFramebufferObject)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickFramebufferObject")
	return
}

func (ptr *QQuickFramebufferObject) DestroyQQuickFramebufferObject() {
}

func (ptr *QQuickFramebufferObject) MirrorVertically() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MirrorVertically"}).(bool)
}

func (ptr *QQuickFramebufferObject) ConnectMirrorVerticallyChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMirrorVerticallyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickFramebufferObject) DisconnectMirrorVerticallyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMirrorVerticallyChanged"})
}

func (ptr *QQuickFramebufferObject) MirrorVerticallyChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MirrorVerticallyChanged", vbo})
}

func (ptr *QQuickFramebufferObject) SetMirrorVertically(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMirrorVertically", enable})
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureFollowsItemSize", follows})
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFollowsItemSize"}).(bool)
}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureFollowsItemSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureFollowsItemSizeChanged"})
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSizeChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFollowsItemSizeChanged", vbo})
}

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func (ptr *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return ptr
}

func (ptr *QQuickImageProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickImageProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQmlImageProviderBase_PTR().SetPointer(p)
	}
}

func PointerFromQQuickImageProvider(ptr QQuickImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func (n *QQuickImageProvider) InitFromInternal(ptr uintptr, name string) {
	n.QQmlImageProviderBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickImageProvider) ClassNameInternalF() string {
	return n.QQmlImageProviderBase_PTR().ClassNameInternalF()
}

func NewQQuickImageProviderFromPointer(ptr unsafe.Pointer) (n *QQuickImageProvider) {
	n = new(QQuickImageProvider)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickImageProvider")
	return
}
func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickImageProvider", "", ty, flags}).(*QQuickImageProvider)
}

func (ptr *QQuickImageProvider) ConnectFlags(f func() qml.QQmlImageProviderBase__Flag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlags", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectFlags() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlags"})
}

func (ptr *QQuickImageProvider) Flags() qml.QQmlImageProviderBase__Flag {

	return qml.QQmlImageProviderBase__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QQuickImageProvider) FlagsDefault() qml.QQmlImageProviderBase__Flag {

	return qml.QQmlImageProviderBase__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault"}).(float64))
}

func (ptr *QQuickImageProvider) ConnectImageType(f func() qml.QQmlImageProviderBase__ImageType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectImageType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectImageType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectImageType"})
}

func (ptr *QQuickImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {

	return qml.QQmlImageProviderBase__ImageType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImageType"}).(float64))
}

func (ptr *QQuickImageProvider) ImageTypeDefault() qml.QQmlImageProviderBase__ImageType {

	return qml.QQmlImageProviderBase__ImageType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImageTypeDefault"}).(float64))
}

func (ptr *QQuickImageProvider) ConnectRequestImage(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QImage) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestImage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectRequestImage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestImage"})
}

func (ptr *QQuickImageProvider) RequestImage(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestImage", id, size, requestedSize}).(*gui.QImage)
}

func (ptr *QQuickImageProvider) RequestImageDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestImageDefault", id, size, requestedSize}).(*gui.QImage)
}

func (ptr *QQuickImageProvider) ConnectRequestPixmap(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QPixmap) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestPixmap", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectRequestPixmap() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestPixmap"})
}

func (ptr *QQuickImageProvider) RequestPixmap(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestPixmap", id, size, requestedSize}).(*gui.QPixmap)
}

func (ptr *QQuickImageProvider) RequestPixmapDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestPixmapDefault", id, size, requestedSize}).(*gui.QPixmap)
}

func (ptr *QQuickImageProvider) ConnectRequestTexture(f func(id string, size *core.QSize, requestedSize *core.QSize) *QQuickTextureFactory) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectRequestTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestTexture"})
}

func (ptr *QQuickImageProvider) RequestTexture(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestTexture", id, size, requestedSize}).(*QQuickTextureFactory)
}

func (ptr *QQuickImageProvider) RequestTextureDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestTextureDefault", id, size, requestedSize}).(*QQuickTextureFactory)
}

func (ptr *QQuickImageProvider) ConnectDestroyQQuickImageProvider(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickImageProvider", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageProvider) DisconnectDestroyQQuickImageProvider() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickImageProvider"})
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickImageProvider"})
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProviderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickImageProviderDefault"})
}

type QQuickImageResponse struct {
	core.QObject
}

type QQuickImageResponse_ITF interface {
	core.QObject_ITF
	QQuickImageResponse_PTR() *QQuickImageResponse
}

func (ptr *QQuickImageResponse) QQuickImageResponse_PTR() *QQuickImageResponse {
	return ptr
}

func (ptr *QQuickImageResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickImageResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickImageResponse(ptr QQuickImageResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageResponse_PTR().Pointer()
	}
	return nil
}

func (n *QQuickImageResponse) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickImageResponse) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickImageResponseFromPointer(ptr unsafe.Pointer) (n *QQuickImageResponse) {
	n = new(QQuickImageResponse)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickImageResponse")
	return
}
func NewQQuickImageResponse() *QQuickImageResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickImageResponse", ""}).(*QQuickImageResponse)
}

func (ptr *QQuickImageResponse) ConnectCancel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCancel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageResponse) DisconnectCancel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCancel"})
}

func (ptr *QQuickImageResponse) Cancel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Cancel"})
}

func (ptr *QQuickImageResponse) CancelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelDefault"})
}

func (ptr *QQuickImageResponse) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageResponse) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QQuickImageResponse) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QQuickImageResponse) ErrorStringDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorStringDefault"}).(string)
}

func (ptr *QQuickImageResponse) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageResponse) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QQuickImageResponse) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QQuickImageResponse) ConnectTextureFactory(f func() *QQuickTextureFactory) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageResponse) DisconnectTextureFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureFactory"})
}

func (ptr *QQuickImageResponse) TextureFactory() *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFactory"}).(*QQuickTextureFactory)
}

func (ptr *QQuickImageResponse) ConnectDestroyQQuickImageResponse(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickImageResponse", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickImageResponse) DisconnectDestroyQQuickImageResponse() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickImageResponse"})
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponse() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickImageResponse"})
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickImageResponseDefault"})
}

func (ptr *QQuickImageResponse) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickImageResponse) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickImageResponse) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickImageResponse) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickImageResponse) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickImageResponse) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickImageResponse) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickImageResponse) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickImageResponse) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickImageResponse) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickImageResponse) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickImageResponse) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickImageResponse) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickImageResponse) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickImageResponse) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickImageResponse) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickImageResponse) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickImageResponse) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (ptr *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return ptr
}

func (ptr *QQuickItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QQmlParserStatus_PTR().SetPointer(p)
	}
}

func PointerFromQQuickItem(ptr QQuickItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (n *QQuickItem) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QQmlParserStatus_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickItem) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickItemFromPointer(ptr unsafe.Pointer) (n *QQuickItem) {
	n = new(QQuickItem)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickItem")
	return
}

//go:generate stringer -type=QQuickItem__Flag
//QQuickItem::Flag
type QQuickItem__Flag int64

const (
	QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = QQuickItem__Flag(0x01)
	QQuickItem__ItemAcceptsInputMethod   QQuickItem__Flag = QQuickItem__Flag(0x02)
	QQuickItem__ItemIsFocusScope         QQuickItem__Flag = QQuickItem__Flag(0x04)
	QQuickItem__ItemHasContents          QQuickItem__Flag = QQuickItem__Flag(0x08)
	QQuickItem__ItemAcceptsDrops         QQuickItem__Flag = QQuickItem__Flag(0x10)
)

//go:generate stringer -type=QQuickItem__ItemChange
//QQuickItem::ItemChange
type QQuickItem__ItemChange int64

const (
	QQuickItem__ItemChildAddedChange           QQuickItem__ItemChange = QQuickItem__ItemChange(0)
	QQuickItem__ItemChildRemovedChange         QQuickItem__ItemChange = QQuickItem__ItemChange(1)
	QQuickItem__ItemSceneChange                QQuickItem__ItemChange = QQuickItem__ItemChange(2)
	QQuickItem__ItemVisibleHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(3)
	QQuickItem__ItemParentHasChanged           QQuickItem__ItemChange = QQuickItem__ItemChange(4)
	QQuickItem__ItemOpacityHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(5)
	QQuickItem__ItemActiveFocusHasChanged      QQuickItem__ItemChange = QQuickItem__ItemChange(6)
	QQuickItem__ItemRotationHasChanged         QQuickItem__ItemChange = QQuickItem__ItemChange(7)
	QQuickItem__ItemAntialiasingHasChanged     QQuickItem__ItemChange = QQuickItem__ItemChange(8)
	QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = QQuickItem__ItemChange(9)
	QQuickItem__ItemEnabledHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(10)
)

//go:generate stringer -type=QQuickItem__TransformOrigin
//QQuickItem::TransformOrigin
type QQuickItem__TransformOrigin int64

const (
	QQuickItem__TopLeft     QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(0)
	QQuickItem__Top         QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(1)
	QQuickItem__TopRight    QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(2)
	QQuickItem__Left        QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(3)
	QQuickItem__Center      QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(4)
	QQuickItem__Right       QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(5)
	QQuickItem__BottomLeft  QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(6)
	QQuickItem__Bottom      QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(7)
	QQuickItem__BottomRight QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(8)
)

func NewQQuickItem(parent QQuickItem_ITF) *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickItem", "", parent}).(*QQuickItem)
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "AcceptHoverEvents"}).(bool)
}

func (ptr *QQuickItem) AcceptTouchEvents() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "AcceptTouchEvents"}).(bool)
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {

	return core.Qt__MouseButton(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "AcceptedMouseButtons"}).(float64))
}

func (ptr *QQuickItem) ConnectActiveFocusChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectActiveFocusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectActiveFocusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectActiveFocusChanged"})
}

func (ptr *QQuickItem) ActiveFocusChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ActiveFocusChanged", vbo})
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ActiveFocusOnTab"}).(bool)
}

func (ptr *QQuickItem) ConnectActiveFocusOnTabChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectActiveFocusOnTabChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectActiveFocusOnTabChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectActiveFocusOnTabChanged"})
}

func (ptr *QQuickItem) ActiveFocusOnTabChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ActiveFocusOnTabChanged", vbo})
}

func (ptr *QQuickItem) Antialiasing() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Antialiasing"}).(bool)
}

func (ptr *QQuickItem) ConnectAntialiasingChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectAntialiasingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectAntialiasingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectAntialiasingChanged"})
}

func (ptr *QQuickItem) AntialiasingChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "AntialiasingChanged", vbo})
}

func (ptr *QQuickItem) BaselineOffset() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "BaselineOffset"}).(float64)
}

func (ptr *QQuickItem) ConnectBaselineOffsetChanged(f func(vqr float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectBaselineOffsetChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectBaselineOffsetChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectBaselineOffsetChanged"})
}

func (ptr *QQuickItem) BaselineOffsetChanged(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "BaselineOffsetChanged", vqr})
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildAt", x, y}).(*QQuickItem)
}

func (ptr *QQuickItem) ChildItems() []*QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildItems"}).([]*QQuickItem)
}

func (ptr *QQuickItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectChildMouseEventFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectChildMouseEventFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectChildMouseEventFilter"})
}

func (ptr *QQuickItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildMouseEventFilter", item, event}).(bool)
}

func (ptr *QQuickItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildMouseEventFilterDefault", item, event}).(bool)
}

func (ptr *QQuickItem) ChildrenRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildrenRect"}).(*core.QRectF)
}

func (ptr *QQuickItem) ConnectChildrenRectChanged(f func(vqr *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectChildrenRectChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectChildrenRectChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectChildrenRectChanged"})
}

func (ptr *QQuickItem) ChildrenRectChanged(vqr core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildrenRectChanged", vqr})
}

func (ptr *QQuickItem) ConnectClassBegin(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectClassBegin", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectClassBegin() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectClassBegin"})
}

func (ptr *QQuickItem) ClassBegin() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ClassBegin"})
}

func (ptr *QQuickItem) ClassBeginDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ClassBeginDefault"})
}

func (ptr *QQuickItem) Clip() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Clip"}).(bool)
}

func (ptr *QQuickItem) ConnectClipChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectClipChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectClipChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectClipChanged"})
}

func (ptr *QQuickItem) ClipChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ClipChanged", vbo})
}

func (ptr *QQuickItem) ConnectComponentComplete(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectComponentComplete", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectComponentComplete() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectComponentComplete"})
}

func (ptr *QQuickItem) ComponentComplete() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ComponentComplete"})
}

func (ptr *QQuickItem) ComponentCompleteDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ComponentCompleteDefault"})
}

func (ptr *QQuickItem) ContainmentMask() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ContainmentMask"}).(*core.QObject)
}

func (ptr *QQuickItem) ConnectContainmentMaskChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectContainmentMaskChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectContainmentMaskChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectContainmentMaskChanged"})
}

func (ptr *QQuickItem) ContainmentMaskChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ContainmentMaskChanged"})
}

func (ptr *QQuickItem) ConnectContains(f func(point *core.QPointF) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectContains", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectContains() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectContains"})
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Contains", point}).(bool)
}

func (ptr *QQuickItem) ContainsDefault(point core.QPointF_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ContainsDefault", point}).(bool)
}

func (ptr *QQuickItem) Cursor() *gui.QCursor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Cursor"}).(*gui.QCursor)
}

func (ptr *QQuickItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDragEnterEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectDragEnterEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDragEnterEvent"})
}

func (ptr *QQuickItem) DragEnterEvent(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragEnterEvent", event})
}

func (ptr *QQuickItem) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QQuickItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDragLeaveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectDragLeaveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDragLeaveEvent"})
}

func (ptr *QQuickItem) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragLeaveEvent", event})
}

func (ptr *QQuickItem) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QQuickItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDragMoveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectDragMoveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDragMoveEvent"})
}

func (ptr *QQuickItem) DragMoveEvent(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragMoveEvent", event})
}

func (ptr *QQuickItem) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QQuickItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDropEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectDropEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDropEvent"})
}

func (ptr *QQuickItem) DropEvent(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DropEvent", event})
}

func (ptr *QQuickItem) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QQuickItem) ConnectEnabledChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectEnabledChanged"})
}

func (ptr *QQuickItem) EnabledChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EnabledChanged"})
}

func (ptr *QQuickItem) ConnectEvent(f func(ev *core.QEvent) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectEvent"})
}

func (ptr *QQuickItem) Event(ev core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", ev}).(bool)
}

func (ptr *QQuickItem) EventDefault(ev core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", ev}).(bool)
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FiltersChildMouseEvents"}).(bool)
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {

	return QQuickItem__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QQuickItem) ConnectFocusChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectFocusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectFocusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectFocusChanged"})
}

func (ptr *QQuickItem) FocusChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FocusChanged", vbo})
}

func (ptr *QQuickItem) ConnectFocusInEvent(f func(vqf *gui.QFocusEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectFocusInEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectFocusInEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectFocusInEvent"})
}

func (ptr *QQuickItem) FocusInEvent(vqf gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FocusInEvent", vqf})
}

func (ptr *QQuickItem) FocusInEventDefault(vqf gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FocusInEventDefault", vqf})
}

func (ptr *QQuickItem) ConnectFocusOutEvent(f func(vqf *gui.QFocusEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectFocusOutEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectFocusOutEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectFocusOutEvent"})
}

func (ptr *QQuickItem) FocusOutEvent(vqf gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FocusOutEvent", vqf})
}

func (ptr *QQuickItem) FocusOutEventDefault(vqf gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "FocusOutEventDefault", vqf})
}

func (ptr *QQuickItem) ForceActiveFocus() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ForceActiveFocus"})
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ForceActiveFocus2", reason})
}

func (ptr *QQuickItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectGeometryChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectGeometryChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectGeometryChanged"})
}

func (ptr *QQuickItem) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "GeometryChanged", newGeometry, oldGeometry})
}

func (ptr *QQuickItem) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "GeometryChangedDefault", newGeometry, oldGeometry})
}

func (ptr *QQuickItem) GrabMouse() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "GrabMouse"})
}

func (ptr *QQuickItem) GrabTouchPoints(ids []int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "GrabTouchPoints", ids})
}

func (ptr *QQuickItem) HasActiveFocus() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HasActiveFocus"}).(bool)
}

func (ptr *QQuickItem) HasFocus() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HasFocus"}).(bool)
}

func (ptr *QQuickItem) Height() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Height"}).(float64)
}

func (ptr *QQuickItem) ConnectHeightChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectHeightChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectHeightChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectHeightChanged"})
}

func (ptr *QQuickItem) HeightChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HeightChanged"})
}

func (ptr *QQuickItem) HeightValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HeightValid"}).(bool)
}

func (ptr *QQuickItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectHoverEnterEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectHoverEnterEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectHoverEnterEvent"})
}

func (ptr *QQuickItem) HoverEnterEvent(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverEnterEvent", event})
}

func (ptr *QQuickItem) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverEnterEventDefault", event})
}

func (ptr *QQuickItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectHoverLeaveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectHoverLeaveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectHoverLeaveEvent"})
}

func (ptr *QQuickItem) HoverLeaveEvent(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverLeaveEvent", event})
}

func (ptr *QQuickItem) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverLeaveEventDefault", event})
}

func (ptr *QQuickItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectHoverMoveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectHoverMoveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectHoverMoveEvent"})
}

func (ptr *QQuickItem) HoverMoveEvent(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverMoveEvent", event})
}

func (ptr *QQuickItem) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "HoverMoveEventDefault", event})
}

func (ptr *QQuickItem) ImplicitHeight() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ImplicitHeight"}).(float64)
}

func (ptr *QQuickItem) ConnectImplicitHeightChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectImplicitHeightChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectImplicitHeightChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectImplicitHeightChanged"})
}

func (ptr *QQuickItem) ImplicitHeightChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ImplicitHeightChanged"})
}

func (ptr *QQuickItem) ImplicitWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ImplicitWidth"}).(float64)
}

func (ptr *QQuickItem) ConnectImplicitWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectImplicitWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectImplicitWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectImplicitWidthChanged"})
}

func (ptr *QQuickItem) ImplicitWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ImplicitWidthChanged"})
}

func (ptr *QQuickItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectInputMethodEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectInputMethodEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectInputMethodEvent"})
}

func (ptr *QQuickItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "InputMethodEvent", event})
}

func (ptr *QQuickItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QQuickItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectInputMethodQuery", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectInputMethodQuery() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectInputMethodQuery"})
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "InputMethodQuery", query}).(*core.QVariant)
}

func (ptr *QQuickItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QQuickItem) IsAncestorOf(child QQuickItem_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsAncestorOf", child}).(bool)
}

func (ptr *QQuickItem) IsComponentComplete() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsComponentComplete"}).(bool)
}

func (ptr *QQuickItem) IsEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsEnabled"}).(bool)
}

func (ptr *QQuickItem) IsFocusScope() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsFocusScope"}).(bool)
}

func (ptr *QQuickItem) ConnectIsTextureProvider(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectIsTextureProvider", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectIsTextureProvider() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectIsTextureProvider"})
}

func (ptr *QQuickItem) IsTextureProvider() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsTextureProvider"}).(bool)
}

func (ptr *QQuickItem) IsTextureProviderDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsTextureProviderDefault"}).(bool)
}

func (ptr *QQuickItem) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QQuickItem) KeepMouseGrab() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeepMouseGrab"}).(bool)
}

func (ptr *QQuickItem) KeepTouchGrab() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeepTouchGrab"}).(bool)
}

func (ptr *QQuickItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectKeyPressEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectKeyPressEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectKeyPressEvent"})
}

func (ptr *QQuickItem) KeyPressEvent(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeyPressEvent", event})
}

func (ptr *QQuickItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QQuickItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectKeyReleaseEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectKeyReleaseEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectKeyReleaseEvent"})
}

func (ptr *QQuickItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeyReleaseEvent", event})
}

func (ptr *QQuickItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QQuickItem) MapFromGlobal(point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapFromGlobal", point}).(*core.QPointF)
}

func (ptr *QQuickItem) MapFromItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapFromItem", item, point}).(*core.QPointF)
}

func (ptr *QQuickItem) MapFromScene(point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapFromScene", point}).(*core.QPointF)
}

func (ptr *QQuickItem) MapRectFromItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapRectFromItem", item, rect}).(*core.QRectF)
}

func (ptr *QQuickItem) MapRectFromScene(rect core.QRectF_ITF) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapRectFromScene", rect}).(*core.QRectF)
}

func (ptr *QQuickItem) MapRectToItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapRectToItem", item, rect}).(*core.QRectF)
}

func (ptr *QQuickItem) MapRectToScene(rect core.QRectF_ITF) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapRectToScene", rect}).(*core.QRectF)
}

func (ptr *QQuickItem) MapToGlobal(point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapToGlobal", point}).(*core.QPointF)
}

func (ptr *QQuickItem) MapToItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapToItem", item, point}).(*core.QPointF)
}

func (ptr *QQuickItem) MapToScene(point core.QPointF_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MapToScene", point}).(*core.QPointF)
}

func (ptr *QQuickItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectMouseDoubleClickEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectMouseDoubleClickEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectMouseDoubleClickEvent"})
}

func (ptr *QQuickItem) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseDoubleClickEvent", event})
}

func (ptr *QQuickItem) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QQuickItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectMouseMoveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectMouseMoveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectMouseMoveEvent"})
}

func (ptr *QQuickItem) MouseMoveEvent(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseMoveEvent", event})
}

func (ptr *QQuickItem) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QQuickItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectMousePressEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectMousePressEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectMousePressEvent"})
}

func (ptr *QQuickItem) MousePressEvent(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MousePressEvent", event})
}

func (ptr *QQuickItem) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QQuickItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectMouseReleaseEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectMouseReleaseEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectMouseReleaseEvent"})
}

func (ptr *QQuickItem) MouseReleaseEvent(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseReleaseEvent", event})
}

func (ptr *QQuickItem) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QQuickItem) ConnectMouseUngrabEvent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectMouseUngrabEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectMouseUngrabEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectMouseUngrabEvent"})
}

func (ptr *QQuickItem) MouseUngrabEvent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseUngrabEvent"})
}

func (ptr *QQuickItem) MouseUngrabEventDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MouseUngrabEventDefault"})
}

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "NextItemInFocusChain", forward}).(*QQuickItem)
}

func (ptr *QQuickItem) Opacity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Opacity"}).(float64)
}

func (ptr *QQuickItem) ConnectOpacityChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectOpacityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectOpacityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectOpacityChanged"})
}

func (ptr *QQuickItem) OpacityChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "OpacityChanged"})
}

func (ptr *QQuickItem) ConnectParentChanged(f func(vqq *QQuickItem)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectParentChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectParentChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectParentChanged"})
}

func (ptr *QQuickItem) ParentChanged(vqq QQuickItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ParentChanged", vqq})
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ParentItem"}).(*QQuickItem)
}

func (ptr *QQuickItem) Polish() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Polish"})
}

func (ptr *QQuickItem) ConnectReleaseResources(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectReleaseResources", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectReleaseResources() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectReleaseResources"})
}

func (ptr *QQuickItem) ReleaseResources() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ReleaseResources"})
}

func (ptr *QQuickItem) ReleaseResourcesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ReleaseResourcesDefault"})
}

func (ptr *QQuickItem) ResetAntialiasing() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ResetAntialiasing"})
}

func (ptr *QQuickItem) ResetHeight() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ResetHeight"})
}

func (ptr *QQuickItem) ResetWidth() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ResetWidth"})
}

func (ptr *QQuickItem) Rotation() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Rotation"}).(float64)
}

func (ptr *QQuickItem) ConnectRotationChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectRotationChanged"})
}

func (ptr *QQuickItem) RotationChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "RotationChanged"})
}

func (ptr *QQuickItem) Scale() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Scale"}).(float64)
}

func (ptr *QQuickItem) ConnectScaleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectScaleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectScaleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectScaleChanged"})
}

func (ptr *QQuickItem) ScaleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ScaleChanged"})
}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ScopedFocusItem"}).(*QQuickItem)
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetAcceptHoverEvents", enabled})
}

func (ptr *QQuickItem) SetAcceptTouchEvents(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetAcceptTouchEvents", enabled})
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetAcceptedMouseButtons", buttons})
}

func (ptr *QQuickItem) SetActiveFocusOnTab(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetActiveFocusOnTab", vbo})
}

func (ptr *QQuickItem) SetAntialiasing(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetAntialiasing", vbo})
}

func (ptr *QQuickItem) SetBaselineOffset(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetBaselineOffset", vqr})
}

func (ptr *QQuickItem) SetClip(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetClip", vbo})
}

func (ptr *QQuickItem) SetContainmentMask(mask core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetContainmentMask", mask})
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetCursor", cursor})
}

func (ptr *QQuickItem) SetEnabled(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetEnabled", vbo})
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetFiltersChildMouseEvents", filter})
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetFlag", flag, enabled})
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetFlags", flags})
}

func (ptr *QQuickItem) SetFocus(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetFocus", vbo})
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetFocus2", focus, reason})
}

func (ptr *QQuickItem) SetHeight(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetHeight", vqr})
}

func (ptr *QQuickItem) SetImplicitHeight(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetImplicitHeight", vqr})
}

func (ptr *QQuickItem) SetImplicitWidth(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetImplicitWidth", vqr})
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetKeepMouseGrab", keep})
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetKeepTouchGrab", keep})
}

func (ptr *QQuickItem) SetOpacity(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetOpacity", vqr})
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetParentItem", parent})
}

func (ptr *QQuickItem) SetRotation(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetRotation", vqr})
}

func (ptr *QQuickItem) SetScale(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetScale", vqr})
}

func (ptr *QQuickItem) SetSize(size core.QSizeF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetSize", size})
}

func (ptr *QQuickItem) SetSmooth(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetSmooth", vbo})
}

func (ptr *QQuickItem) SetState(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetState", vqs})
}

func (ptr *QQuickItem) SetTransformOrigin(vqq QQuickItem__TransformOrigin) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetTransformOrigin", vqq})
}

func (ptr *QQuickItem) SetVisible(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetVisible", vbo})
}

func (ptr *QQuickItem) SetWidth(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetWidth", vqr})
}

func (ptr *QQuickItem) SetX(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetX", vqr})
}

func (ptr *QQuickItem) SetY(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetY", vqr})
}

func (ptr *QQuickItem) SetZ(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetZ", vqr})
}

func (ptr *QQuickItem) Size() *core.QSizeF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Size"}).(*core.QSizeF)
}

func (ptr *QQuickItem) Smooth() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Smooth"}).(bool)
}

func (ptr *QQuickItem) ConnectSmoothChanged(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectSmoothChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectSmoothChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectSmoothChanged"})
}

func (ptr *QQuickItem) SmoothChanged(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SmoothChanged", vbo})
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "StackAfter", sibling})
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "StackBefore", sibling})
}

func (ptr *QQuickItem) State() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "State"}).(string)
}

func (ptr *QQuickItem) ConnectStateChanged(f func(vqs string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QQuickItem) StateChanged(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "StateChanged", vqs})
}

func (ptr *QQuickItem) ConnectTextureProvider(f func() *QSGTextureProvider) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectTextureProvider", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectTextureProvider() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectTextureProvider"})
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TextureProvider"}).(*QSGTextureProvider)
}

func (ptr *QQuickItem) TextureProviderDefault() *QSGTextureProvider {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TextureProviderDefault"}).(*QSGTextureProvider)
}

func (ptr *QQuickItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectTouchEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectTouchEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectTouchEvent"})
}

func (ptr *QQuickItem) TouchEvent(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TouchEvent", event})
}

func (ptr *QQuickItem) TouchEventDefault(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TouchEventDefault", event})
}

func (ptr *QQuickItem) ConnectTouchUngrabEvent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectTouchUngrabEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectTouchUngrabEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectTouchUngrabEvent"})
}

func (ptr *QQuickItem) TouchUngrabEvent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TouchUngrabEvent"})
}

func (ptr *QQuickItem) TouchUngrabEventDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TouchUngrabEventDefault"})
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {

	return QQuickItem__TransformOrigin(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TransformOrigin"}).(float64))
}

func (ptr *QQuickItem) ConnectTransformOriginChanged(f func(vqq QQuickItem__TransformOrigin)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectTransformOriginChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectTransformOriginChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectTransformOriginChanged"})
}

func (ptr *QQuickItem) TransformOriginChanged(vqq QQuickItem__TransformOrigin) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TransformOriginChanged", vqq})
}

func (ptr *QQuickItem) UngrabMouse() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UngrabMouse"})
}

func (ptr *QQuickItem) UngrabTouchPoints() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UngrabTouchPoints"})
}

func (ptr *QQuickItem) UnsetCursor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UnsetCursor"})
}

func (ptr *QQuickItem) ConnectUpdate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectUpdate"})
}

func (ptr *QQuickItem) Update() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Update"})
}

func (ptr *QQuickItem) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QQuickItem) UpdateInputMethod(queries core.Qt__InputMethodQuery) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UpdateInputMethod", queries})
}

func (ptr *QQuickItem) ConnectUpdatePolish(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectUpdatePolish", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectUpdatePolish() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectUpdatePolish"})
}

func (ptr *QQuickItem) UpdatePolish() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UpdatePolish"})
}

func (ptr *QQuickItem) UpdatePolishDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UpdatePolishDefault"})
}

func (ptr *QQuickItem) ConnectVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectVisibleChanged"})
}

func (ptr *QQuickItem) VisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "VisibleChanged"})
}

func (ptr *QQuickItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectWheelEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectWheelEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectWheelEvent"})
}

func (ptr *QQuickItem) WheelEvent(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WheelEvent", event})
}

func (ptr *QQuickItem) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QQuickItem) Width() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Width"}).(float64)
}

func (ptr *QQuickItem) ConnectWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectWidthChanged"})
}

func (ptr *QQuickItem) WidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WidthChanged"})
}

func (ptr *QQuickItem) WidthValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WidthValid"}).(bool)
}

func (ptr *QQuickItem) Window() *QQuickWindow {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Window"}).(*QQuickWindow)
}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectWindowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectWindowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectWindowChanged"})
}

func (ptr *QQuickItem) WindowChanged(window QQuickWindow_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WindowChanged", window})
}

func (ptr *QQuickItem) X() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "X"}).(float64)
}

func (ptr *QQuickItem) ConnectXChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectXChanged"})
}

func (ptr *QQuickItem) XChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "XChanged"})
}

func (ptr *QQuickItem) Y() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Y"}).(float64)
}

func (ptr *QQuickItem) ConnectYChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectYChanged"})
}

func (ptr *QQuickItem) YChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "YChanged"})
}

func (ptr *QQuickItem) Z() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Z"}).(float64)
}

func (ptr *QQuickItem) ConnectZChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectZChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectZChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectZChanged"})
}

func (ptr *QQuickItem) ZChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ZChanged"})
}

func (ptr *QQuickItem) ConnectDestroyQQuickItem(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDestroyQQuickItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItem) DisconnectDestroyQQuickItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDestroyQQuickItem"})
}

func (ptr *QQuickItem) DestroyQQuickItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQQuickItem"})
}

func (ptr *QQuickItem) DestroyQQuickItemDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQQuickItemDefault"})
}

func (ptr *QQuickItem) __childItems_atList(i int) *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__childItems_atList", i}).(*QQuickItem)
}

func (ptr *QQuickItem) __childItems_setList(i QQuickItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__childItems_setList", i})
}

func (ptr *QQuickItem) __childItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__childItems_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) __grabTouchPoints_ids_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__grabTouchPoints_ids_atList", i}).(float64))
}

func (ptr *QQuickItem) __grabTouchPoints_ids_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__grabTouchPoints_ids_setList", i})
}

func (ptr *QQuickItem) __grabTouchPoints_ids_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__grabTouchPoints_ids_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickItem) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickItem) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickItem) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickItem) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickItem) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickItem) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickItem) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickItem) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QQuickItem) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickItem) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QQuickItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickItem) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QQuickItem) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickItem) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QQuickItem) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickItem) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QQuickItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QQuickItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickItem) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QQuickItem) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickItem) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QQuickItem) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickItemGrabResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func (n *QQuickItemGrabResult) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickItemGrabResult) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) (n *QQuickItemGrabResult) {
	n = new(QQuickItemGrabResult)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickItemGrabResult")
	return
}

func (ptr *QQuickItemGrabResult) DestroyQQuickItemGrabResult() {
}

func (ptr *QQuickItemGrabResult) Image() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Image"}).(*gui.QImage)
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReady", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReady"})
}

func (ptr *QQuickItemGrabResult) Ready() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Ready"})
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SaveToFile", fileName}).(bool)
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QQuickItemGrabResult) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickItemGrabResult) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickItemGrabResult) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItemGrabResult) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickItemGrabResult) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickItemGrabResult) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickItemGrabResult) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickItemGrabResult) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickItemGrabResult) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickItemGrabResult) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickItemGrabResult) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickItemGrabResult) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickItemGrabResult) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickItemGrabResult) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickItemGrabResult) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickItemGrabResult) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickItemGrabResult) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickItemGrabResult) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickPaintedItem struct {
	QQuickItem
}

type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func (ptr *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem {
	return ptr
}

func (ptr *QQuickPaintedItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickPaintedItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromQQuickPaintedItem(ptr QQuickPaintedItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickPaintedItem_PTR().Pointer()
	}
	return nil
}

func (n *QQuickPaintedItem) InitFromInternal(ptr uintptr, name string) {
	n.QQuickItem_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickPaintedItem) ClassNameInternalF() string {
	return n.QQuickItem_PTR().ClassNameInternalF()
}

func NewQQuickPaintedItemFromPointer(ptr unsafe.Pointer) (n *QQuickPaintedItem) {
	n = new(QQuickPaintedItem)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickPaintedItem")
	return
}

//go:generate stringer -type=QQuickPaintedItem__RenderTarget
//QQuickPaintedItem::RenderTarget
type QQuickPaintedItem__RenderTarget int64

const (
	QQuickPaintedItem__Image                      QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(0)
	QQuickPaintedItem__FramebufferObject          QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(1)
	QQuickPaintedItem__InvertedYFramebufferObject QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(2)
)

//go:generate stringer -type=QQuickPaintedItem__PerformanceHint
//QQuickPaintedItem::PerformanceHint
type QQuickPaintedItem__PerformanceHint int64

const (
	QQuickPaintedItem__FastFBOResizing QQuickPaintedItem__PerformanceHint = QQuickPaintedItem__PerformanceHint(0x1)
)

func NewQQuickPaintedItem(parent QQuickItem_ITF) *QQuickPaintedItem {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickPaintedItem", "", parent}).(*QQuickPaintedItem)
}

func (ptr *QQuickPaintedItem) ContentsScale() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsScale"}).(float64)
}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentsScaleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentsScaleChanged"})
}

func (ptr *QQuickPaintedItem) ContentsScaleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsScaleChanged"})
}

func (ptr *QQuickPaintedItem) ContentsSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsSize"}).(*core.QSize)
}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentsSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentsSizeChanged"})
}

func (ptr *QQuickPaintedItem) ContentsSizeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsSizeChanged"})
}

func (ptr *QQuickPaintedItem) FillColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FillColor"}).(*gui.QColor)
}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFillColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFillColorChanged"})
}

func (ptr *QQuickPaintedItem) FillColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FillColorChanged"})
}

func (ptr *QQuickPaintedItem) Mipmap() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Mipmap"}).(bool)
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpaquePainting"}).(bool)
}

func (ptr *QQuickPaintedItem) ConnectPaint(f func(painter *gui.QPainter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaint", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectPaint() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaint"})
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Paint", painter})
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {

	return QQuickPaintedItem__PerformanceHint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PerformanceHints"}).(float64))
}

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {

	return QQuickPaintedItem__RenderTarget(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderTarget"}).(float64))
}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRenderTargetChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRenderTargetChanged"})
}

func (ptr *QQuickPaintedItem) RenderTargetChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderTargetChanged"})
}

func (ptr *QQuickPaintedItem) SetContentsScale(vqr float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentsScale", vqr})
}

func (ptr *QQuickPaintedItem) SetContentsSize(vqs core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentsSize", vqs})
}

func (ptr *QQuickPaintedItem) SetFillColor(vqc gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFillColor", vqc})
}

func (ptr *QQuickPaintedItem) SetMipmap(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMipmap", enable})
}

func (ptr *QQuickPaintedItem) SetOpaquePainting(opaque bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpaquePainting", opaque})
}

func (ptr *QQuickPaintedItem) SetPerformanceHint(hint QQuickPaintedItem__PerformanceHint, enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPerformanceHint", hint, enabled})
}

func (ptr *QQuickPaintedItem) SetPerformanceHints(hints QQuickPaintedItem__PerformanceHint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPerformanceHints", hints})
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRenderTarget", target})
}

func (ptr *QQuickPaintedItem) SetTextureSize(size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureSize", size})
}

func (ptr *QQuickPaintedItem) TextureSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSize"}).(*core.QSize)
}

func (ptr *QQuickPaintedItem) ConnectTextureSizeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectTextureSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureSizeChanged"})
}

func (ptr *QQuickPaintedItem) TextureSizeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSizeChanged"})
}

func (ptr *QQuickPaintedItem) Update(rect core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Update", rect})
}

func (ptr *QQuickPaintedItem) ConnectDestroyQQuickPaintedItem(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickPaintedItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickPaintedItem) DisconnectDestroyQQuickPaintedItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickPaintedItem"})
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickPaintedItem"})
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItemDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickPaintedItemDefault"})
}

type QQuickProfilerAdapter struct {
	internal.Internal
}

type QQuickProfilerAdapter_ITF interface {
	QQuickProfilerAdapter_PTR() *QQuickProfilerAdapter
}

func (ptr *QQuickProfilerAdapter) QQuickProfilerAdapter_PTR() *QQuickProfilerAdapter {
	return ptr
}

func (ptr *QQuickProfilerAdapter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQuickProfilerAdapter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQuickProfilerAdapter(ptr QQuickProfilerAdapter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickProfilerAdapter_PTR().Pointer()
	}
	return nil
}

func (n *QQuickProfilerAdapter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQuickProfilerAdapterFromPointer(ptr unsafe.Pointer) (n *QQuickProfilerAdapter) {
	n = new(QQuickProfilerAdapter)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickProfilerAdapter")
	return
}

func (ptr *QQuickProfilerAdapter) DestroyQQuickProfilerAdapter() {
}

type QQuickProfilerAdapterFactory struct {
	internal.Internal
}

type QQuickProfilerAdapterFactory_ITF interface {
	QQuickProfilerAdapterFactory_PTR() *QQuickProfilerAdapterFactory
}

func (ptr *QQuickProfilerAdapterFactory) QQuickProfilerAdapterFactory_PTR() *QQuickProfilerAdapterFactory {
	return ptr
}

func (ptr *QQuickProfilerAdapterFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQuickProfilerAdapterFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQuickProfilerAdapterFactory(ptr QQuickProfilerAdapterFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickProfilerAdapterFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQuickProfilerAdapterFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQuickProfilerAdapterFactoryFromPointer(ptr unsafe.Pointer) (n *QQuickProfilerAdapterFactory) {
	n = new(QQuickProfilerAdapterFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickProfilerAdapterFactory")
	return
}

func (ptr *QQuickProfilerAdapterFactory) DestroyQQuickProfilerAdapterFactory() {
}

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func (ptr *QQuickRenderControl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickRenderControl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func (n *QQuickRenderControl) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickRenderControl) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) (n *QQuickRenderControl) {
	n = new(QQuickRenderControl)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickRenderControl")
	return
}
func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickRenderControl", "", parent}).(*QQuickRenderControl)
}

func (ptr *QQuickRenderControl) Grab() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Grab"}).(*gui.QImage)
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initialize", gl})
}

func (ptr *QQuickRenderControl) Invalidate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Invalidate"})
}

func (ptr *QQuickRenderControl) PolishItems() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PolishItems"})
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrepareThread", targetThread})
}

func (ptr *QQuickRenderControl) Render() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render"})
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRenderRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRenderRequested"})
}

func (ptr *QQuickRenderControl) RenderRequested() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderRequested"})
}

func (ptr *QQuickRenderControl) ConnectRenderWindow(f func(offset *core.QPoint) *gui.QWindow) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRenderWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickRenderControl) DisconnectRenderWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRenderWindow"})
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderWindow", offset}).(*gui.QWindow)
}

func (ptr *QQuickRenderControl) RenderWindowDefault(offset core.QPoint_ITF) *gui.QWindow {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderWindowDefault", offset}).(*gui.QWindow)
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickRenderControl_RenderWindowFor", "", win, offset}).(*gui.QWindow)
}

func (ptr *QQuickRenderControl) RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickRenderControl_RenderWindowFor", "", win, offset}).(*gui.QWindow)
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneChanged"})
}

func (ptr *QQuickRenderControl) SceneChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneChanged"})
}

func (ptr *QQuickRenderControl) Sync() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Sync"}).(bool)
}

func (ptr *QQuickRenderControl) ConnectDestroyQQuickRenderControl(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickRenderControl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickRenderControl) DisconnectDestroyQQuickRenderControl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickRenderControl"})
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickRenderControl"})
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControlDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickRenderControlDefault"})
}

func (ptr *QQuickRenderControl) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickRenderControl) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickRenderControl) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickRenderControl) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickRenderControl) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickRenderControl) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickRenderControl) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickRenderControl) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickRenderControl) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickRenderControl) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickRenderControl) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickRenderControl) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickRenderControl) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickRenderControl) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickRenderControl) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickRenderControl) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickRenderControl) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickRenderControl) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return ptr
}

func (ptr *QQuickTextDocument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickTextDocument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func (n *QQuickTextDocument) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickTextDocument) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) (n *QQuickTextDocument) {
	n = new(QQuickTextDocument)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickTextDocument")
	return
}

func (ptr *QQuickTextDocument) DestroyQQuickTextDocument() {
}

func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickTextDocument", "", parent}).(*QQuickTextDocument)
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextDocument"}).(*gui.QTextDocument)
}

func (ptr *QQuickTextDocument) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickTextDocument) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickTextDocument) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextDocument) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickTextDocument) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickTextDocument) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextDocument) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickTextDocument) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickTextDocument) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickTextDocument) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickTextDocument) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickTextDocument) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickTextDocument) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickTextDocument) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickTextDocument) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickTextDocument) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickTextureFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQuickTextureFactory) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickTextureFactory) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) (n *QQuickTextureFactory) {
	n = new(QQuickTextureFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickTextureFactory")
	return
}
func NewQQuickTextureFactory() *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickTextureFactory", ""}).(*QQuickTextureFactory)
}

func (ptr *QQuickTextureFactory) ConnectCreateTexture(f func(window *QQuickWindow) *QSGTexture) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickTextureFactory) DisconnectCreateTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateTexture"})
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTexture", window}).(*QSGTexture)
}

func (ptr *QQuickTextureFactory) ConnectImage(f func() *gui.QImage) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectImage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickTextureFactory) DisconnectImage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectImage"})
}

func (ptr *QQuickTextureFactory) Image() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Image"}).(*gui.QImage)
}

func (ptr *QQuickTextureFactory) ImageDefault() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImageDefault"}).(*gui.QImage)
}

func (ptr *QQuickTextureFactory) ConnectTextureByteCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureByteCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickTextureFactory) DisconnectTextureByteCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureByteCount"})
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureByteCount"}).(float64))
}

func QQuickTextureFactory_TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickTextureFactory_TextureFactoryForImage", "", image}).(*QQuickTextureFactory)
}

func (ptr *QQuickTextureFactory) TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickTextureFactory_TextureFactoryForImage", "", image}).(*QQuickTextureFactory)
}

func (ptr *QQuickTextureFactory) ConnectTextureSize(f func() *core.QSize) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickTextureFactory) DisconnectTextureSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureSize"})
}

func (ptr *QQuickTextureFactory) TextureSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSize"}).(*core.QSize)
}

func (ptr *QQuickTextureFactory) ConnectDestroyQQuickTextureFactory(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickTextureFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickTextureFactory) DisconnectDestroyQQuickTextureFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickTextureFactory"})
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickTextureFactory"})
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactoryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickTextureFactoryDefault"})
}

func (ptr *QQuickTextureFactory) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickTextureFactory) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickTextureFactory) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextureFactory) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickTextureFactory) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickTextureFactory) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTextureFactory) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickTextureFactory) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickTextureFactory) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickTextureFactory) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickTextureFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickTextureFactory) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickTextureFactory) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickTextureFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickTextureFactory) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickTextureFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickTextureFactory) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickTextureFactory) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickTransform struct {
	core.QObject
}

type QQuickTransform_ITF interface {
	core.QObject_ITF
	QQuickTransform_PTR() *QQuickTransform
}

func (ptr *QQuickTransform) QQuickTransform_PTR() *QQuickTransform {
	return ptr
}

func (ptr *QQuickTransform) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickTransform) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickTransform(ptr QQuickTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTransform_PTR().Pointer()
	}
	return nil
}

func (n *QQuickTransform) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickTransform) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickTransformFromPointer(ptr unsafe.Pointer) (n *QQuickTransform) {
	n = new(QQuickTransform)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickTransform")
	return
}

func (ptr *QQuickTransform) DestroyQQuickTransform() {
}

func (ptr *QQuickTransform) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickTransform) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickTransform) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTransform) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickTransform) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickTransform) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTransform) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickTransform) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickTransform) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickTransform) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickTransform) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickTransform) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickTransform) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickTransform) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickTransform) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickTransform) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickTransform) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickTransform) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickTransform) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickTransform) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickTransform) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func (ptr *QQuickView) QQuickView_PTR() *QQuickView {
	return ptr
}

func (ptr *QQuickView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickWindow_PTR().SetPointer(p)
	}
}

func PointerFromQQuickView(ptr QQuickView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickView_PTR().Pointer()
	}
	return nil
}

func (n *QQuickView) InitFromInternal(ptr uintptr, name string) {
	n.QQuickWindow_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickView) ClassNameInternalF() string {
	return n.QQuickWindow_PTR().ClassNameInternalF()
}

func NewQQuickViewFromPointer(ptr unsafe.Pointer) (n *QQuickView) {
	n = new(QQuickView)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickView")
	return
}

//go:generate stringer -type=QQuickView__ResizeMode
//QQuickView::ResizeMode
type QQuickView__ResizeMode int64

const (
	QQuickView__SizeViewToRootObject QQuickView__ResizeMode = QQuickView__ResizeMode(0)
	QQuickView__SizeRootObjectToView QQuickView__ResizeMode = QQuickView__ResizeMode(1)
)

//go:generate stringer -type=QQuickView__Status
//QQuickView::Status
type QQuickView__Status int64

const (
	QQuickView__Null    QQuickView__Status = QQuickView__Status(0)
	QQuickView__Ready   QQuickView__Status = QQuickView__Status(1)
	QQuickView__Loading QQuickView__Status = QQuickView__Status(2)
	QQuickView__Error   QQuickView__Status = QQuickView__Status(3)
)

func NewQQuickView(parent gui.QWindow_ITF) *QQuickView {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickView", "", parent}).(*QQuickView)
}

func NewQQuickView2(engine qml.QQmlEngine_ITF, parent gui.QWindow_ITF) *QQuickView {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickView2", "", engine, parent}).(*QQuickView)
}

func NewQQuickView3(source core.QUrl_ITF, parent gui.QWindow_ITF) *QQuickView {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickView3", "", source, parent}).(*QQuickView)
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*qml.QQmlEngine)
}

func (ptr *QQuickView) Errors() []*qml.QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Errors"}).([]*qml.QQmlError)
}

func (ptr *QQuickView) InitialSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitialSize"}).(*core.QSize)
}

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {

	return QQuickView__ResizeMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeMode"}).(float64))
}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootContext"}).(*qml.QQmlContext)
}

func (ptr *QQuickView) RootObject() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootObject"}).(*QQuickItem)
}

func (ptr *QQuickView) SetResizeMode(vqq QQuickView__ResizeMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetResizeMode", vqq})
}

func (ptr *QQuickView) ConnectSetSource(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickView) DisconnectSetSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSource"})
}

func (ptr *QQuickView) SetSource(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSource", url})
}

func (ptr *QQuickView) SetSourceDefault(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceDefault", url})
}

func (ptr *QQuickView) Source() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Source"}).(*core.QUrl)
}

func (ptr *QQuickView) Status() QQuickView__Status {

	return QQuickView__Status(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Status"}).(float64))
}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStatusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickView) DisconnectStatusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStatusChanged"})
}

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StatusChanged", status})
}

func (ptr *QQuickView) ConnectDestroyQQuickView(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickView", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickView) DisconnectDestroyQQuickView() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickView"})
}

func (ptr *QQuickView) DestroyQQuickView() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickView"})
}

func (ptr *QQuickView) DestroyQQuickViewDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickViewDefault"})
}

func (ptr *QQuickView) __errors_atList(i int) *qml.QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_atList", i}).(*qml.QQmlError)
}

func (ptr *QQuickView) __errors_setList(i qml.QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_setList", i})
}

func (ptr *QQuickView) __errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_newList"}).(unsafe.Pointer)
}

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return ptr
}

func (ptr *QQuickWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func (n *QQuickWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) (n *QQuickWidget) {
	n = new(QQuickWidget)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickWidget")
	return
}

//go:generate stringer -type=QQuickWidget__ResizeMode
//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int64

const (
	QQuickWidget__SizeViewToRootObject QQuickWidget__ResizeMode = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView QQuickWidget__ResizeMode = QQuickWidget__ResizeMode(1)
)

//go:generate stringer -type=QQuickWidget__Status
//QQuickWidget::Status
type QQuickWidget__Status int64

const (
	QQuickWidget__Null    QQuickWidget__Status = QQuickWidget__Status(0)
	QQuickWidget__Ready   QQuickWidget__Status = QQuickWidget__Status(1)
	QQuickWidget__Loading QQuickWidget__Status = QQuickWidget__Status(2)
	QQuickWidget__Error   QQuickWidget__Status = QQuickWidget__Status(3)
)

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickWidget", "", parent}).(*QQuickWidget)
}

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickWidget2", "", engine, parent}).(*QQuickWidget)
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickWidget3", "", source, parent}).(*QQuickWidget)
}

func (ptr *QQuickWidget) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", e})
}

func (ptr *QQuickWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", e})
}

func (ptr *QQuickWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", e})
}

func (ptr *QQuickWidget) DropEventDefault(e gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", e})
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*qml.QQmlEngine)
}

func (ptr *QQuickWidget) Errors() []*qml.QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Errors"}).([]*qml.QQmlError)
}

func (ptr *QQuickWidget) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QQuickWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QQuickWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QQuickWidget) Format() *gui.QSurfaceFormat {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Format"}).(*gui.QSurfaceFormat)
}

func (ptr *QQuickWidget) GrabFramebuffer() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabFramebuffer"}).(*gui.QImage)
}

func (ptr *QQuickWidget) HideEventDefault(vqh gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", vqh})
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitialSize"}).(*core.QSize)
}

func (ptr *QQuickWidget) KeyPressEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", e})
}

func (ptr *QQuickWidget) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", e})
}

func (ptr *QQuickWidget) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", e})
}

func (ptr *QQuickWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", e})
}

func (ptr *QQuickWidget) MousePressEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", e})
}

func (ptr *QQuickWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", e})
}

func (ptr *QQuickWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QuickWindow"}).(*QQuickWindow)
}

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {

	return QQuickWidget__ResizeMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeMode"}).(float64))
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootContext"}).(*qml.QQmlContext)
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootObject"}).(*QQuickItem)
}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphError"})
}

func (ptr *QQuickWidget) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphError", error, message})
}

func (ptr *QQuickWidget) SetClearColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClearColor", color})
}

func (ptr *QQuickWidget) SetFormat(format gui.QSurfaceFormat_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFormat", format})
}

func (ptr *QQuickWidget) SetResizeMode(vqq QQuickWidget__ResizeMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetResizeMode", vqq})
}

func (ptr *QQuickWidget) ConnectSetSource(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWidget) DisconnectSetSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSource"})
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSource", url})
}

func (ptr *QQuickWidget) SetSourceDefault(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceDefault", url})
}

func (ptr *QQuickWidget) ShowEventDefault(vqs gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", vqs})
}

func (ptr *QQuickWidget) Source() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Source"}).(*core.QUrl)
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {

	return QQuickWidget__Status(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Status"}).(float64))
}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStatusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStatusChanged"})
}

func (ptr *QQuickWidget) StatusChanged(status QQuickWidget__Status) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StatusChanged", status})
}

func (ptr *QQuickWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", e})
}

func (ptr *QQuickWidget) ConnectDestroyQQuickWidget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWidget) DisconnectDestroyQQuickWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickWidget"})
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickWidget"})
}

func (ptr *QQuickWidget) DestroyQQuickWidgetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickWidgetDefault"})
}

func (ptr *QQuickWidget) __errors_atList(i int) *qml.QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_atList", i}).(*qml.QQmlError)
}

func (ptr *QQuickWidget) __errors_setList(i qml.QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_setList", i})
}

func (ptr *QQuickWidget) __errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QQuickWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QQuickWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QQuickWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QQuickWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QQuickWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QQuickWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QQuickWidget) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QQuickWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QQuickWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QQuickWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QQuickWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QQuickWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QQuickWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QQuickWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QQuickWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QQuickWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QQuickWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QQuickWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QQuickWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QQuickWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QQuickWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QQuickWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QQuickWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QQuickWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QQuickWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QQuickWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QQuickWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QQuickWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QQuickWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QQuickWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QQuickWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QQuickWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QQuickWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QQuickWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QQuickWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QQuickWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QQuickWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QQuickWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QQuickWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QQuickWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QQuickWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QQuickWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QQuickWidget) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QQuickWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QQuickWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return ptr
}

func (ptr *QQuickWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWindow_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func (n *QQuickWindow) InitFromInternal(ptr uintptr, name string) {
	n.QWindow_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickWindow) ClassNameInternalF() string {
	return n.QWindow_PTR().ClassNameInternalF()
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) (n *QQuickWindow) {
	n = new(QQuickWindow)
	n.InitFromInternal(uintptr(ptr), "quick.QQuickWindow")
	return
}

//go:generate stringer -type=QQuickWindow__CreateTextureOption
//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0008)
	QQuickWindow__TextureIsOpaque        QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0010)
)

//go:generate stringer -type=QQuickWindow__RenderStage
//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage QQuickWindow__RenderStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  QQuickWindow__RenderStage = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     QQuickWindow__RenderStage = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      QQuickWindow__RenderStage = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           QQuickWindow__RenderStage = QQuickWindow__RenderStage(4)
	QQuickWindow__NoStage                  QQuickWindow__RenderStage = QQuickWindow__RenderStage(5)
)

//go:generate stringer -type=QQuickWindow__SceneGraphError
//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable QQuickWindow__SceneGraphError = QQuickWindow__SceneGraphError(1)
)

//go:generate stringer -type=QQuickWindow__TextRenderType
//QQuickWindow::TextRenderType
type QQuickWindow__TextRenderType int64

const (
	QQuickWindow__QtTextRendering     QQuickWindow__TextRenderType = QQuickWindow__TextRenderType(0)
	QQuickWindow__NativeTextRendering QQuickWindow__TextRenderType = QQuickWindow__TextRenderType(1)
)

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQQuickWindow", "", parent}).(*QQuickWindow)
}

func (ptr *QQuickWindow) ConnectAccessibleRoot(f func() *gui.QAccessibleInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAccessibleRoot", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectAccessibleRoot() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAccessibleRoot"})
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AccessibleRoot"}).(*gui.QAccessibleInterface)
}

func (ptr *QQuickWindow) AccessibleRootDefault() *gui.QAccessibleInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AccessibleRootDefault"}).(*gui.QAccessibleInterface)
}

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveFocusItem"}).(*QQuickItem)
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveFocusItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveFocusItemChanged"})
}

func (ptr *QQuickWindow) ActiveFocusItemChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveFocusItemChanged"})
}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAfterAnimating", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAfterAnimating"})
}

func (ptr *QQuickWindow) AfterAnimating() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AfterAnimating"})
}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAfterRendering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAfterRendering"})
}

func (ptr *QQuickWindow) AfterRendering() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AfterRendering"})
}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAfterSynchronizing", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAfterSynchronizing"})
}

func (ptr *QQuickWindow) AfterSynchronizing() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AfterSynchronizing"})
}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeforeRendering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeforeRendering"})
}

func (ptr *QQuickWindow) BeforeRendering() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeforeRendering"})
}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeforeSynchronizing", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeforeSynchronizing"})
}

func (ptr *QQuickWindow) BeforeSynchronizing() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeforeSynchronizing"})
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearBeforeRendering"}).(bool)
}

func (ptr *QQuickWindow) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QQuickWindow) ConnectColorChanged(f func(vqc *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QQuickWindow) ColorChanged(vqc gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", vqc})
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentItem"}).(*QQuickItem)
}

func (ptr *QQuickWindow) CreateImageNode() *QSGImageNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateImageNode"}).(*QSGImageNode)
}

func (ptr *QQuickWindow) CreateRectangleNode() *QSGRectangleNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRectangleNode"}).(*QSGRectangleNode)
}

func (ptr *QQuickWindow) CreateTextureFromId(id uint, size core.QSize_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextureFromId", id, size, options}).(*QSGTexture)
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextureFromImage", image, options}).(*QSGTexture)
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextureFromImage2", image}).(*QSGTexture)
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EffectiveDevicePixelRatio"}).(float64)
}

func (ptr *QQuickWindow) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickWindow) ExposeEventDefault(vqe gui.QExposeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExposeEventDefault", vqe})
}

func (ptr *QQuickWindow) FocusInEventDefault(ev gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", ev})
}

func (ptr *QQuickWindow) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", ev})
}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFrameSwapped", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFrameSwapped"})
}

func (ptr *QQuickWindow) FrameSwapped() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FrameSwapped"})
}

func (ptr *QQuickWindow) GrabWindow() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabWindow"}).(*gui.QImage)
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_HasDefaultAlphaBuffer", ""}).(bool)
}

func (ptr *QQuickWindow) HasDefaultAlphaBuffer() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_HasDefaultAlphaBuffer", ""}).(bool)
}

func (ptr *QQuickWindow) HideEventDefault(vqh gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", vqh})
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubationController"}).(*qml.QQmlIncubationController)
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPersistentOpenGLContext"}).(bool)
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPersistentSceneGraph"}).(bool)
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSceneGraphInitialized"}).(bool)
}

func (ptr *QQuickWindow) KeyPressEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", e})
}

func (ptr *QQuickWindow) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", e})
}

func (ptr *QQuickWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseGrabberItem"}).(*QQuickItem)
}

func (ptr *QQuickWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QQuickWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QQuickWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenglContext"}).(*gui.QOpenGLContext)
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpenglContextCreated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpenglContextCreated"})
}

func (ptr *QQuickWindow) OpenglContextCreated(context gui.QOpenGLContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenglContextCreated", context})
}

func (ptr *QQuickWindow) ConnectReleaseResources(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleaseResources", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectReleaseResources() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleaseResources"})
}

func (ptr *QQuickWindow) ReleaseResources() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseResources"})
}

func (ptr *QQuickWindow) ReleaseResourcesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseResourcesDefault"})
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderTarget"}).(*gui.QOpenGLFramebufferObject)
}

func (ptr *QQuickWindow) RenderTargetId() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderTargetId"}).(float64))
}

func (ptr *QQuickWindow) RenderTargetSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderTargetSize"}).(*core.QSize)
}

func (ptr *QQuickWindow) RendererInterface() *QSGRendererInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RendererInterface"}).(*QSGRendererInterface)
}

func (ptr *QQuickWindow) ResetOpenGLState() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetOpenGLState"})
}

func (ptr *QQuickWindow) ResizeEventDefault(ev gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", ev})
}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphAboutToStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphAboutToStop"})
}

func (ptr *QQuickWindow) SceneGraphAboutToStop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphAboutToStop"})
}

func QQuickWindow_SceneGraphBackend() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SceneGraphBackend", ""}).(string)
}

func (ptr *QQuickWindow) SceneGraphBackend() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SceneGraphBackend", ""}).(string)
}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphError"})
}

func (ptr *QQuickWindow) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphError", error, message})
}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphInitialized", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphInitialized"})
}

func (ptr *QQuickWindow) SceneGraphInitialized() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphInitialized"})
}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphInvalidated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphInvalidated"})
}

func (ptr *QQuickWindow) SceneGraphInvalidated() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphInvalidated"})
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScheduleRenderJob", job, stage})
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClearBeforeRendering", enabled})
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetDefaultAlphaBuffer", "", useAlpha})
}

func (ptr *QQuickWindow) SetDefaultAlphaBuffer(useAlpha bool) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetDefaultAlphaBuffer", "", useAlpha})
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentOpenGLContext", persistent})
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentSceneGraph", persistent})
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRenderTarget", fbo})
}

func (ptr *QQuickWindow) SetRenderTarget2(fboId uint, size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRenderTarget2", fboId, size})
}

func QQuickWindow_SetSceneGraphBackend(api QSGRendererInterface__GraphicsApi) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetSceneGraphBackend", "", api})
}

func (ptr *QQuickWindow) SetSceneGraphBackend(api QSGRendererInterface__GraphicsApi) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetSceneGraphBackend", "", api})
}

func QQuickWindow_SetSceneGraphBackend2(backend string) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetSceneGraphBackend2", "", backend})
}

func (ptr *QQuickWindow) SetSceneGraphBackend2(backend string) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetSceneGraphBackend2", "", backend})
}

func QQuickWindow_SetTextRenderType(renderType QQuickWindow__TextRenderType) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetTextRenderType", "", renderType})
}

func (ptr *QQuickWindow) SetTextRenderType(renderType QQuickWindow__TextRenderType) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_SetTextRenderType", "", renderType})
}

func (ptr *QQuickWindow) ShowEventDefault(vqs gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", vqs})
}

func QQuickWindow_TextRenderType() QQuickWindow__TextRenderType {

	return QQuickWindow__TextRenderType(internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_TextRenderType", ""}).(float64))
}

func (ptr *QQuickWindow) TextRenderType() QQuickWindow__TextRenderType {

	return QQuickWindow__TextRenderType(internal.CallLocalFunction([]interface{}{"", "", "quick.QQuickWindow_TextRenderType", ""}).(float64))
}

func (ptr *QQuickWindow) ConnectUpdate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdate"})
}

func (ptr *QQuickWindow) Update() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Update"})
}

func (ptr *QQuickWindow) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QQuickWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QQuickWindow) ConnectDestroyQQuickWindow(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQuickWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWindow) DisconnectDestroyQQuickWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQuickWindow"})
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickWindow"})
}

func (ptr *QQuickWindow) DestroyQQuickWindowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQuickWindowDefault"})
}

func (ptr *QQuickWindow) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickWindow) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickWindow) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWindow) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickWindow) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickWindow) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWindow) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickWindow) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickWindow) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWindow) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickWindow) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickWindow) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickWindow) AlertDefault(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AlertDefault", msec})
}

func (ptr *QQuickWindow) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QQuickWindow) FocusObjectDefault() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusObjectDefault"}).(*core.QObject)
}

func (ptr *QQuickWindow) FormatDefault() *gui.QSurfaceFormat {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormatDefault"}).(*gui.QSurfaceFormat)
}

func (ptr *QQuickWindow) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QQuickWindow) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QQuickWindow) MoveEventDefault(ev gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", ev})
}

func (ptr *QQuickWindow) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QQuickWindow) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QQuickWindow) RequestActivateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestActivateDefault"})
}

func (ptr *QQuickWindow) RequestUpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdateDefault"})
}

func (ptr *QQuickWindow) SetGeometryDefault(posx int, posy int, w int, h int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGeometryDefault", posx, posy, w, h})
}

func (ptr *QQuickWindow) SetGeometry2Default(rect core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGeometry2Default", rect})
}

func (ptr *QQuickWindow) SetHeightDefault(arg int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeightDefault", arg})
}

func (ptr *QQuickWindow) SetMaximumHeightDefault(h int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumHeightDefault", h})
}

func (ptr *QQuickWindow) SetMaximumWidthDefault(w int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumWidthDefault", w})
}

func (ptr *QQuickWindow) SetMinimumHeightDefault(h int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinimumHeightDefault", h})
}

func (ptr *QQuickWindow) SetMinimumWidthDefault(w int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinimumWidthDefault", w})
}

func (ptr *QQuickWindow) SetTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleDefault", vqs})
}

func (ptr *QQuickWindow) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QQuickWindow) SetWidthDefault(arg int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWidthDefault", arg})
}

func (ptr *QQuickWindow) SetXDefault(arg int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXDefault", arg})
}

func (ptr *QQuickWindow) SetYDefault(arg int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYDefault", arg})
}

func (ptr *QQuickWindow) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QQuickWindow) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QQuickWindow) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QQuickWindow) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QQuickWindow) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QQuickWindow) SizeDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(*core.QSize)
}

func (ptr *QQuickWindow) SurfaceTypeDefault() gui.QSurface__SurfaceType {

	return gui.QSurface__SurfaceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SurfaceTypeDefault"}).(float64))
}

func (ptr *QQuickWindow) TabletEventDefault(ev gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", ev})
}

func (ptr *QQuickWindow) TouchEventDefault(ev gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TouchEventDefault", ev})
}

func (ptr *QQuickWindow) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickWindow) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickWindow) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickWindow) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickWindow) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func (ptr *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return ptr
}

func (ptr *QSGAbstractRenderer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGAbstractRenderer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRenderer_PTR().Pointer()
	}
	return nil
}

func (n *QSGAbstractRenderer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGAbstractRenderer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSGAbstractRendererFromPointer(ptr unsafe.Pointer) (n *QSGAbstractRenderer) {
	n = new(QSGAbstractRenderer)
	n.InitFromInternal(uintptr(ptr), "quick.QSGAbstractRenderer")
	return
}

func (ptr *QSGAbstractRenderer) DestroyQSGAbstractRenderer() {
}

//go:generate stringer -type=QSGAbstractRenderer__ClearModeBit
//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0004)
)

func (ptr *QSGAbstractRenderer) ClearColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearColor"}).(*gui.QColor)
}

func (ptr *QSGAbstractRenderer) ClearMode() QSGAbstractRenderer__ClearModeBit {

	return QSGAbstractRenderer__ClearModeBit(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearMode"}).(float64))
}

func (ptr *QSGAbstractRenderer) DeviceRect() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceRect"}).(*core.QRect)
}

func (ptr *QSGAbstractRenderer) ProjectionMatrix() *gui.QMatrix4x4 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProjectionMatrix"}).(*gui.QMatrix4x4)
}

func (ptr *QSGAbstractRenderer) ConnectRenderScene(f func(fboId uint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRenderScene", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGAbstractRenderer) DisconnectRenderScene() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRenderScene"})
}

func (ptr *QSGAbstractRenderer) RenderScene(fboId uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderScene", fboId})
}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneGraphChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneGraphChanged"})
}

func (ptr *QSGAbstractRenderer) SceneGraphChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneGraphChanged"})
}

func (ptr *QSGAbstractRenderer) SetClearColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClearColor", color})
}

func (ptr *QSGAbstractRenderer) SetClearMode(mode QSGAbstractRenderer__ClearModeBit) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClearMode", mode})
}

func (ptr *QSGAbstractRenderer) SetDeviceRect(rect core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeviceRect", rect})
}

func (ptr *QSGAbstractRenderer) SetDeviceRect2(size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeviceRect2", size})
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrix(matrix gui.QMatrix4x4_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProjectionMatrix", matrix})
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrixToRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProjectionMatrixToRect", rect})
}

func (ptr *QSGAbstractRenderer) SetViewportRect(rect core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewportRect", rect})
}

func (ptr *QSGAbstractRenderer) SetViewportRect2(size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewportRect2", size})
}

func (ptr *QSGAbstractRenderer) ViewportRect() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportRect"}).(*core.QRect)
}

func (ptr *QSGAbstractRenderer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSGAbstractRenderer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSGAbstractRenderer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSGAbstractRenderer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSGAbstractRenderer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSGAbstractRenderer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSGAbstractRenderer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSGAbstractRenderer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSGAbstractRenderer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSGAbstractRenderer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSGAbstractRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSGAbstractRenderer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSGAbstractRenderer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSGAbstractRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSGAbstractRenderer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSGAbstractRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSGAbstractRenderer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSGAbstractRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return ptr
}

func (ptr *QSGBasicGeometryNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGBasicGeometryNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGBasicGeometryNode) ClassNameInternalF() string {
	return n.QSGNode_PTR().ClassNameInternalF()
}

func NewQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) (n *QSGBasicGeometryNode) {
	n = new(QSGBasicGeometryNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGBasicGeometryNode")
	return
}
func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Geometry"}).(*QSGGeometry)
}

func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Geometry2"}).(*QSGGeometry)
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGeometry", geometry})
}

func (ptr *QSGBasicGeometryNode) ConnectDestroyQSGBasicGeometryNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGBasicGeometryNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGBasicGeometryNode) DisconnectDestroyQSGBasicGeometryNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGBasicGeometryNode"})
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGBasicGeometryNode"})
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGBasicGeometryNodeDefault"})
}

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func (ptr *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return ptr
}

func (ptr *QSGClipNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGClipNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGBasicGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGClipNode(ptr QSGClipNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGClipNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGBasicGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGClipNode) ClassNameInternalF() string {
	return n.QSGBasicGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGClipNodeFromPointer(ptr unsafe.Pointer) (n *QSGClipNode) {
	n = new(QSGClipNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGClipNode")
	return
}
func NewQSGClipNode() *QSGClipNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGClipNode", ""}).(*QSGClipNode)
}

func (ptr *QSGClipNode) ClipRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClipRect"}).(*core.QRectF)
}

func (ptr *QSGClipNode) IsRectangular() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRectangular"}).(bool)
}

func (ptr *QSGClipNode) SetClipRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClipRect", rect})
}

func (ptr *QSGClipNode) SetIsRectangular(rectHint bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIsRectangular", rectHint})
}

func (ptr *QSGClipNode) ConnectDestroyQSGClipNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGClipNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGClipNode) DisconnectDestroyQSGClipNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGClipNode"})
}

func (ptr *QSGClipNode) DestroyQSGClipNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGClipNode"})
}

func (ptr *QSGClipNode) DestroyQSGClipNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGClipNodeDefault"})
}

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGDynamicTexture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGTexture_PTR().SetPointer(p)
	}
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func (n *QSGDynamicTexture) InitFromInternal(ptr uintptr, name string) {
	n.QSGTexture_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGDynamicTexture) ClassNameInternalF() string {
	return n.QSGTexture_PTR().ClassNameInternalF()
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) (n *QSGDynamicTexture) {
	n = new(QSGDynamicTexture)
	n.InitFromInternal(uintptr(ptr), "quick.QSGDynamicTexture")
	return
}

func (ptr *QSGDynamicTexture) DestroyQSGDynamicTexture() {
}

func (ptr *QSGDynamicTexture) ConnectUpdateTexture(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGDynamicTexture) DisconnectUpdateTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateTexture"})
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateTexture"}).(bool)
}

func (ptr *QSGDynamicTexture) Bind() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Bind"})
}

func (ptr *QSGDynamicTexture) BindDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindDefault"})
}

func (ptr *QSGDynamicTexture) HasAlphaChannel() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAlphaChannel"}).(bool)
}

func (ptr *QSGDynamicTexture) HasAlphaChannelDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAlphaChannelDefault"}).(bool)
}

func (ptr *QSGDynamicTexture) HasMipmaps() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasMipmaps"}).(bool)
}

func (ptr *QSGDynamicTexture) HasMipmapsDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasMipmapsDefault"}).(bool)
}

func (ptr *QSGDynamicTexture) TextureId() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureId"}).(float64))
}

func (ptr *QSGDynamicTexture) TextureIdDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureIdDefault"}).(float64))
}

func (ptr *QSGDynamicTexture) TextureSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSize"}).(*core.QSize)
}

func (ptr *QSGDynamicTexture) TextureSizeDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSizeDefault"}).(*core.QSize)
}

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return ptr
}

func (ptr *QSGEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func (n *QSGEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) (n *QSGEngine) {
	n = new(QSGEngine)
	n.InitFromInternal(uintptr(ptr), "quick.QSGEngine")
	return
}

//go:generate stringer -type=QSGEngine__CreateTextureOption
//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0008)
	QSGEngine__TextureIsOpaque        QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0010)
)

func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGEngine", "", parent}).(*QSGEngine)
}

func (ptr *QSGEngine) CreateImageNode() *QSGImageNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateImageNode"}).(*QSGImageNode)
}

func (ptr *QSGEngine) CreateRectangleNode() *QSGRectangleNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRectangleNode"}).(*QSGRectangleNode)
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRenderer"}).(*QSGAbstractRenderer)
}

func (ptr *QSGEngine) CreateTextureFromId(id uint, size core.QSize_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextureFromId", id, size, options}).(*QSGTexture)
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextureFromImage", image, options}).(*QSGTexture)
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initialize", context})
}

func (ptr *QSGEngine) Invalidate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Invalidate"})
}

func (ptr *QSGEngine) RendererInterface() *QSGRendererInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RendererInterface"}).(*QSGRendererInterface)
}

func (ptr *QSGEngine) ConnectDestroyQSGEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGEngine) DisconnectDestroyQSGEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGEngine"})
}

func (ptr *QSGEngine) DestroyQSGEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGEngine"})
}

func (ptr *QSGEngine) DestroyQSGEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGEngineDefault"})
}

func (ptr *QSGEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSGEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSGEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSGEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSGEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSGEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSGEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSGEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSGEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSGEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSGEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSGEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSGEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSGEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSGEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSGEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSGEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSGEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSGEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSGEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSGEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return ptr
}

func (ptr *QSGFlatColorMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGFlatColorMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGFlatColorMaterial) ClassNameInternalF() string {
	return n.QSGMaterial_PTR().ClassNameInternalF()
}

func NewQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) (n *QSGFlatColorMaterial) {
	n = new(QSGFlatColorMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGFlatColorMaterial")
	return
}

func (ptr *QSGFlatColorMaterial) DestroyQSGFlatColorMaterial() {
}

func NewQSGFlatColorMaterial() *QSGFlatColorMaterial {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGFlatColorMaterial", ""}).(*QSGFlatColorMaterial)
}

func (ptr *QSGFlatColorMaterial) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QSGFlatColorMaterial) CreateShader() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShader"}).(*QSGMaterialShader)
}

func (ptr *QSGFlatColorMaterial) CreateShaderDefault() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShaderDefault"}).(*QSGMaterialShader)
}

func (ptr *QSGFlatColorMaterial) Type() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*QSGMaterialType)
}

func (ptr *QSGFlatColorMaterial) TypeDefault() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(*QSGMaterialType)
}

type QSGGeometry struct {
	internal.Internal
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (ptr *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return ptr
}

func (ptr *QSGGeometry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGGeometry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGGeometry(ptr QSGGeometry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometry_PTR().Pointer()
	}
	return nil
}

func (n *QSGGeometry) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGGeometryFromPointer(ptr unsafe.Pointer) (n *QSGGeometry) {
	n = new(QSGGeometry)
	n.InitFromInternal(uintptr(ptr), "quick.QSGGeometry")
	return
}

//go:generate stringer -type=QSGGeometry__AttributeType
//QSGGeometry::AttributeType
type QSGGeometry__AttributeType int64

const (
	QSGGeometry__UnknownAttribute   QSGGeometry__AttributeType = QSGGeometry__AttributeType(0)
	QSGGeometry__PositionAttribute  QSGGeometry__AttributeType = QSGGeometry__AttributeType(1)
	QSGGeometry__ColorAttribute     QSGGeometry__AttributeType = QSGGeometry__AttributeType(2)
	QSGGeometry__TexCoordAttribute  QSGGeometry__AttributeType = QSGGeometry__AttributeType(3)
	QSGGeometry__TexCoord1Attribute QSGGeometry__AttributeType = QSGGeometry__AttributeType(4)
	QSGGeometry__TexCoord2Attribute QSGGeometry__AttributeType = QSGGeometry__AttributeType(5)
)

//go:generate stringer -type=QSGGeometry__DataPattern
//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      QSGGeometry__DataPattern = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(3)
)

//go:generate stringer -type=QSGGeometry__DrawingMode
//QSGGeometry::DrawingMode
type QSGGeometry__DrawingMode int64

const (
	QSGGeometry__DrawPoints        QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0000)
	QSGGeometry__DrawLines         QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0001)
	QSGGeometry__DrawLineLoop      QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0002)
	QSGGeometry__DrawLineStrip     QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0003)
	QSGGeometry__DrawTriangles     QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0004)
	QSGGeometry__DrawTriangleStrip QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0005)
	QSGGeometry__DrawTriangleFan   QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0006)
)

//go:generate stringer -type=QSGGeometry__Type
//QSGGeometry::Type
type QSGGeometry__Type int64

const (
	QSGGeometry__ByteType          QSGGeometry__Type = QSGGeometry__Type(0x1400)
	QSGGeometry__UnsignedByteType  QSGGeometry__Type = QSGGeometry__Type(0x1401)
	QSGGeometry__ShortType         QSGGeometry__Type = QSGGeometry__Type(0x1402)
	QSGGeometry__UnsignedShortType QSGGeometry__Type = QSGGeometry__Type(0x1403)
	QSGGeometry__IntType           QSGGeometry__Type = QSGGeometry__Type(0x1404)
	QSGGeometry__UnsignedIntType   QSGGeometry__Type = QSGGeometry__Type(0x1405)
	QSGGeometry__FloatType         QSGGeometry__Type = QSGGeometry__Type(0x1406)
)

func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Allocate", vertexCount, indexCount})
}

func (ptr *QSGGeometry) AttributeCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeCount"}).(float64))
}

func (ptr *QSGGeometry) DrawingMode() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawingMode"}).(float64))
}

func (ptr *QSGGeometry) IndexCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexCount"}).(float64))
}

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexData"}).(unsafe.Pointer)
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexData2"}).(unsafe.Pointer)
}

func (ptr *QSGGeometry) IndexDataAsUInt() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDataAsUInt"}).(float64))
}

func (ptr *QSGGeometry) IndexDataAsUInt2() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDataAsUInt2"}).(float64))
}

func (ptr *QSGGeometry) IndexDataAsUShort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDataAsUShort"}).(float64))
}

func (ptr *QSGGeometry) IndexDataAsUShort2() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDataAsUShort2"}).(float64))
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {

	return QSGGeometry__DataPattern(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDataPattern"}).(float64))
}

func (ptr *QSGGeometry) IndexType() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexType"}).(float64))
}

func (ptr *QSGGeometry) LineWidth() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineWidth"}).(float32)
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkIndexDataDirty"})
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkVertexDataDirty"})
}

func (ptr *QSGGeometry) SetDrawingMode(mode uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDrawingMode", mode})
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIndexDataPattern", p})
}

func (ptr *QSGGeometry) SetLineWidth(width float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLineWidth", width})
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVertexDataPattern", p})
}

func (ptr *QSGGeometry) SizeOfIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeOfIndex"}).(float64))
}

func (ptr *QSGGeometry) SizeOfVertex() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeOfVertex"}).(float64))
}

func QSGGeometry_UpdateColoredRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateColoredRectGeometry", "", g, rect})
}

func (ptr *QSGGeometry) UpdateColoredRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateColoredRectGeometry", "", g, rect})
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateRectGeometry", "", g, rect})
}

func (ptr *QSGGeometry) UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateRectGeometry", "", g, rect})
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateTexturedRectGeometry", "", g, rect, textureRect})
}

func (ptr *QSGGeometry) UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGGeometry_UpdateTexturedRectGeometry", "", g, rect, textureRect})
}

func (ptr *QSGGeometry) VertexCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexCount"}).(float64))
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexData"}).(unsafe.Pointer)
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexData2"}).(unsafe.Pointer)
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {

	return QSGGeometry__DataPattern(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexDataPattern"}).(float64))
}

func (ptr *QSGGeometry) ConnectDestroyQSGGeometry(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGGeometry", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGGeometry) DisconnectDestroyQSGGeometry() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGGeometry"})
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGGeometry"})
}

func (ptr *QSGGeometry) DestroyQSGGeometryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGGeometryDefault"})
}

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return ptr
}

func (ptr *QSGGeometryNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGGeometryNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGBasicGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGGeometryNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGBasicGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGGeometryNode) ClassNameInternalF() string {
	return n.QSGBasicGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) (n *QSGGeometryNode) {
	n = new(QSGGeometryNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGGeometryNode")
	return
}
func NewQSGGeometryNode() *QSGGeometryNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGGeometryNode", ""}).(*QSGGeometryNode)
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Material"}).(*QSGMaterial)
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpaqueMaterial"}).(*QSGMaterial)
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaterial", material})
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpaqueMaterial", material})
}

func (ptr *QSGGeometryNode) ConnectDestroyQSGGeometryNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGGeometryNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGGeometryNode) DisconnectDestroyQSGGeometryNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGGeometryNode"})
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGGeometryNode"})
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGGeometryNodeDefault"})
}

type QSGImageNode struct {
	QSGGeometryNode
}

type QSGImageNode_ITF interface {
	QSGGeometryNode_ITF
	QSGImageNode_PTR() *QSGImageNode
}

func (ptr *QSGImageNode) QSGImageNode_PTR() *QSGImageNode {
	return ptr
}

func (ptr *QSGImageNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGImageNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGImageNode(ptr QSGImageNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGImageNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGImageNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGImageNode) ClassNameInternalF() string {
	return n.QSGGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGImageNodeFromPointer(ptr unsafe.Pointer) (n *QSGImageNode) {
	n = new(QSGImageNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGImageNode")
	return
}

func (ptr *QSGImageNode) DestroyQSGImageNode() {
}

//go:generate stringer -type=QSGImageNode__TextureCoordinatesTransformFlag
//QSGImageNode::TextureCoordinatesTransformFlag
type QSGImageNode__TextureCoordinatesTransformFlag int64

const (
	QSGImageNode__NoTransform        QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x00)
	QSGImageNode__MirrorHorizontally QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x01)
	QSGImageNode__MirrorVertically   QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x02)
)

func (ptr *QSGImageNode) ConnectFiltering(f func() QSGTexture__Filtering) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFiltering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectFiltering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFiltering"})
}

func (ptr *QSGImageNode) Filtering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filtering"}).(float64))
}

func (ptr *QSGImageNode) ConnectMipmapFiltering(f func() QSGTexture__Filtering) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMipmapFiltering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectMipmapFiltering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMipmapFiltering"})
}

func (ptr *QSGImageNode) MipmapFiltering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MipmapFiltering"}).(float64))
}

func (ptr *QSGImageNode) ConnectOwnsTexture(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOwnsTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectOwnsTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOwnsTexture"})
}

func (ptr *QSGImageNode) OwnsTexture() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OwnsTexture"}).(bool)
}

func QSGImageNode_RebuildGeometry(g QSGGeometry_ITF, texture QSGTexture_ITF, rect core.QRectF_ITF, sourceRect core.QRectF_ITF, texCoordMode QSGImageNode__TextureCoordinatesTransformFlag) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGImageNode_RebuildGeometry", "", g, texture, rect, sourceRect, texCoordMode})
}

func (ptr *QSGImageNode) RebuildGeometry(g QSGGeometry_ITF, texture QSGTexture_ITF, rect core.QRectF_ITF, sourceRect core.QRectF_ITF, texCoordMode QSGImageNode__TextureCoordinatesTransformFlag) {

	internal.CallLocalFunction([]interface{}{"", "", "quick.QSGImageNode_RebuildGeometry", "", g, texture, rect, sourceRect, texCoordMode})
}

func (ptr *QSGImageNode) ConnectRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRect"})
}

func (ptr *QSGImageNode) Rect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rect"}).(*core.QRectF)
}

func (ptr *QSGImageNode) ConnectSetFiltering(f func(filtering QSGTexture__Filtering)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFiltering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetFiltering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFiltering"})
}

func (ptr *QSGImageNode) SetFiltering(filtering QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFiltering", filtering})
}

func (ptr *QSGImageNode) ConnectSetMipmapFiltering(f func(filtering QSGTexture__Filtering)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetMipmapFiltering", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetMipmapFiltering() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetMipmapFiltering"})
}

func (ptr *QSGImageNode) SetMipmapFiltering(filtering QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMipmapFiltering", filtering})
}

func (ptr *QSGImageNode) ConnectSetOwnsTexture(f func(owns bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetOwnsTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetOwnsTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetOwnsTexture"})
}

func (ptr *QSGImageNode) SetOwnsTexture(owns bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOwnsTexture", owns})
}

func (ptr *QSGImageNode) ConnectSetRect(f func(rect *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRect"})
}

func (ptr *QSGImageNode) SetRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect", rect})
}

func (ptr *QSGImageNode) SetRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect2", x, y, w, h})
}

func (ptr *QSGImageNode) ConnectSetSourceRect(f func(rect *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSourceRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetSourceRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSourceRect"})
}

func (ptr *QSGImageNode) SetSourceRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceRect", rect})
}

func (ptr *QSGImageNode) SetSourceRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceRect2", x, y, w, h})
}

func (ptr *QSGImageNode) ConnectSetTexture(f func(texture *QSGTexture)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetTexture"})
}

func (ptr *QSGImageNode) SetTexture(texture QSGTexture_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTexture", texture})
}

func (ptr *QSGImageNode) ConnectSetTextureCoordinatesTransform(f func(mode QSGImageNode__TextureCoordinatesTransformFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetTextureCoordinatesTransform", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSetTextureCoordinatesTransform() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetTextureCoordinatesTransform"})
}

func (ptr *QSGImageNode) SetTextureCoordinatesTransform(mode QSGImageNode__TextureCoordinatesTransformFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureCoordinatesTransform", mode})
}

func (ptr *QSGImageNode) ConnectSourceRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSourceRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectSourceRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSourceRect"})
}

func (ptr *QSGImageNode) SourceRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceRect"}).(*core.QRectF)
}

func (ptr *QSGImageNode) ConnectTexture(f func() *QSGTexture) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTexture"})
}

func (ptr *QSGImageNode) Texture() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Texture"}).(*QSGTexture)
}

func (ptr *QSGImageNode) ConnectTextureCoordinatesTransform(f func() QSGImageNode__TextureCoordinatesTransformFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureCoordinatesTransform", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGImageNode) DisconnectTextureCoordinatesTransform() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureCoordinatesTransform"})
}

func (ptr *QSGImageNode) TextureCoordinatesTransform() QSGImageNode__TextureCoordinatesTransformFlag {

	return QSGImageNode__TextureCoordinatesTransformFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureCoordinatesTransform"}).(float64))
}

type QSGMaterial struct {
	internal.Internal
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return ptr
}

func (ptr *QSGMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGMaterial(ptr QSGMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGMaterial) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGMaterialFromPointer(ptr unsafe.Pointer) (n *QSGMaterial) {
	n = new(QSGMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGMaterial")
	return
}

func (ptr *QSGMaterial) DestroyQSGMaterial() {
}

//go:generate stringer -type=QSGMaterial__Flag
//QSGMaterial::Flag
type QSGMaterial__Flag int64

const (
	QSGMaterial__Blending                          QSGMaterial__Flag = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               QSGMaterial__Flag = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate QSGMaterial__Flag = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                QSGMaterial__Flag = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 QSGMaterial__Flag = QSGMaterial__Flag(0x0010)
)

func (ptr *QSGMaterial) ConnectCompare(f func(other *QSGMaterial) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCompare", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterial) DisconnectCompare() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCompare"})
}

func (ptr *QSGMaterial) Compare(other QSGMaterial_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Compare", other}).(float64))
}

func (ptr *QSGMaterial) CompareDefault(other QSGMaterial_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompareDefault", other}).(float64))
}

func (ptr *QSGMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateShader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterial) DisconnectCreateShader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateShader"})
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShader"}).(*QSGMaterialShader)
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {

	return QSGMaterial__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlag", flags, on})
}

func (ptr *QSGMaterial) ConnectType(f func() *QSGMaterialType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterial) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*QSGMaterialType)
}

type QSGMaterialShader struct {
	internal.Internal
}

type QSGMaterialShader_ITF interface {
	QSGMaterialShader_PTR() *QSGMaterialShader
}

func (ptr *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGMaterialShader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func (n *QSGMaterialShader) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGMaterialShaderFromPointer(ptr unsafe.Pointer) (n *QSGMaterialShader) {
	n = new(QSGMaterialShader)
	n.InitFromInternal(uintptr(ptr), "quick.QSGMaterialShader")
	return
}

func (ptr *QSGMaterialShader) DestroyQSGMaterialShader() {
}

func (ptr *QSGMaterialShader) ConnectActivate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActivate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectActivate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActivate"})
}

func (ptr *QSGMaterialShader) Activate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Activate"})
}

func (ptr *QSGMaterialShader) ActivateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivateDefault"})
}

func (ptr *QSGMaterialShader) ConnectCompile(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCompile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectCompile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCompile"})
}

func (ptr *QSGMaterialShader) Compile() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Compile"})
}

func (ptr *QSGMaterialShader) CompileDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompileDefault"})
}

func (ptr *QSGMaterialShader) ConnectDeactivate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeactivate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectDeactivate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeactivate"})
}

func (ptr *QSGMaterialShader) Deactivate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Deactivate"})
}

func (ptr *QSGMaterialShader) DeactivateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeactivateDefault"})
}

func (ptr *QSGMaterialShader) ConnectFragmentShader(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFragmentShader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectFragmentShader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFragmentShader"})
}

func (ptr *QSGMaterialShader) FragmentShader() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FragmentShader"}).(string)
}

func (ptr *QSGMaterialShader) FragmentShaderDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FragmentShaderDefault"}).(string)
}

func (ptr *QSGMaterialShader) ConnectInitialize(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInitialize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectInitialize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInitialize"})
}

func (ptr *QSGMaterialShader) Initialize() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initialize"})
}

func (ptr *QSGMaterialShader) InitializeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitializeDefault"})
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Program"}).(*gui.QOpenGLShaderProgram)
}

func (ptr *QSGMaterialShader) SetShaderSourceFile(ty gui.QOpenGLShader__ShaderTypeBit, sourceFile string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShaderSourceFile", ty, sourceFile})
}

func (ptr *QSGMaterialShader) SetShaderSourceFiles(ty gui.QOpenGLShader__ShaderTypeBit, sourceFiles []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShaderSourceFiles", ty, sourceFiles})
}

func (ptr *QSGMaterialShader) ConnectVertexShader(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVertexShader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGMaterialShader) DisconnectVertexShader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVertexShader"})
}

func (ptr *QSGMaterialShader) VertexShader() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexShader"}).(string)
}

func (ptr *QSGMaterialShader) VertexShaderDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VertexShaderDefault"}).(string)
}

type QSGMaterialType struct {
	internal.Internal
}

type QSGMaterialType_ITF interface {
	QSGMaterialType_PTR() *QSGMaterialType
}

func (ptr *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return ptr
}

func (ptr *QSGMaterialType) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGMaterialType) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGMaterialType(ptr QSGMaterialType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialType_PTR().Pointer()
	}
	return nil
}

func (n *QSGMaterialType) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGMaterialTypeFromPointer(ptr unsafe.Pointer) (n *QSGMaterialType) {
	n = new(QSGMaterialType)
	n.InitFromInternal(uintptr(ptr), "quick.QSGMaterialType")
	return
}

func (ptr *QSGMaterialType) DestroyQSGMaterialType() {
}

type QSGNode struct {
	internal.Internal
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (ptr *QSGNode) QSGNode_PTR() *QSGNode {
	return ptr
}

func (ptr *QSGNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGNode(ptr QSGNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGNode) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGNodeFromPointer(ptr unsafe.Pointer) (n *QSGNode) {
	n = new(QSGNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGNode")
	return
}

//go:generate stringer -type=QSGNode__NodeType
//QSGNode::NodeType
type QSGNode__NodeType int64

const (
	QSGNode__BasicNodeType     QSGNode__NodeType = QSGNode__NodeType(0)
	QSGNode__GeometryNodeType  QSGNode__NodeType = QSGNode__NodeType(1)
	QSGNode__TransformNodeType QSGNode__NodeType = QSGNode__NodeType(2)
	QSGNode__ClipNodeType      QSGNode__NodeType = QSGNode__NodeType(3)
	QSGNode__OpacityNodeType   QSGNode__NodeType = QSGNode__NodeType(4)
	QSGNode__RootNodeType      QSGNode__NodeType = QSGNode__NodeType(5)
	QSGNode__RenderNodeType    QSGNode__NodeType = QSGNode__NodeType(6)
)

//go:generate stringer -type=QSGNode__Flag
//QSGNode::Flag
type QSGNode__Flag int64

const (
	QSGNode__OwnedByParent      QSGNode__Flag = QSGNode__Flag(0x0001)
	QSGNode__UsePreprocess      QSGNode__Flag = QSGNode__Flag(0x0002)
	QSGNode__OwnsGeometry       QSGNode__Flag = QSGNode__Flag(0x00010000)
	QSGNode__OwnsMaterial       QSGNode__Flag = QSGNode__Flag(0x00020000)
	QSGNode__OwnsOpaqueMaterial QSGNode__Flag = QSGNode__Flag(0x00040000)
	QSGNode__IsVisitableNode    QSGNode__Flag = QSGNode__Flag(0x01000000)
	QSGNode__InternalReserved   QSGNode__Flag = QSGNode__Flag(0x01000000)
)

//go:generate stringer -type=QSGNode__DirtyStateBit
//QSGNode::DirtyStateBit
type QSGNode__DirtyStateBit int64

const (
	QSGNode__DirtySubtreeBlocked QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0080)
	QSGNode__DirtyMatrix         QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0100)
	QSGNode__DirtyNodeAdded      QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0400)
	QSGNode__DirtyNodeRemoved    QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0800)
	QSGNode__DirtyGeometry       QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x1000)
	QSGNode__DirtyMaterial       QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x2000)
	QSGNode__DirtyOpacity        QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x4000)
	QSGNode__DirtyForceUpdate    QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x8000)
	QSGNode__DirtyUsePreprocess  QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(QSGNode__UsePreprocess)
)

func NewQSGNode() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGNode", ""}).(*QSGNode)
}

func (ptr *QSGNode) AppendChildNode(node QSGNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AppendChildNode", node})
}

func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildAtIndex", i}).(*QSGNode)
}

func (ptr *QSGNode) ChildCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildCount"}).(float64))
}

func (ptr *QSGNode) FirstChild() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstChild"}).(*QSGNode)
}

func (ptr *QSGNode) Flags() QSGNode__Flag {

	return QSGNode__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QSGNode) InsertChildNodeAfter(node QSGNode_ITF, after QSGNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertChildNodeAfter", node, after})
}

func (ptr *QSGNode) InsertChildNodeBefore(node QSGNode_ITF, before QSGNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertChildNodeBefore", node, before})
}

func (ptr *QSGNode) ConnectIsSubtreeBlocked(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsSubtreeBlocked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGNode) DisconnectIsSubtreeBlocked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsSubtreeBlocked"})
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSubtreeBlocked"}).(bool)
}

func (ptr *QSGNode) IsSubtreeBlockedDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSubtreeBlockedDefault"}).(bool)
}

func (ptr *QSGNode) LastChild() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastChild"}).(*QSGNode)
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkDirty", bits})
}

func (ptr *QSGNode) NextSibling() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextSibling"}).(*QSGNode)
}

func (ptr *QSGNode) Parent() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parent"}).(*QSGNode)
}

func (ptr *QSGNode) PrependChildNode(node QSGNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrependChildNode", node})
}

func (ptr *QSGNode) ConnectPreprocess(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreprocess", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGNode) DisconnectPreprocess() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreprocess"})
}

func (ptr *QSGNode) Preprocess() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Preprocess"})
}

func (ptr *QSGNode) PreprocessDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreprocessDefault"})
}

func (ptr *QSGNode) PreviousSibling() *QSGNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousSibling"}).(*QSGNode)
}

func (ptr *QSGNode) RemoveAllChildNodes() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAllChildNodes"})
}

func (ptr *QSGNode) RemoveChildNode(node QSGNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveChildNode", node})
}

func (ptr *QSGNode) SetFlag(ff QSGNode__Flag, enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlag", ff, enabled})
}

func (ptr *QSGNode) SetFlags(ff QSGNode__Flag, enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlags", ff, enabled})
}

func (ptr *QSGNode) Type() QSGNode__NodeType {

	return QSGNode__NodeType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QSGNode) ConnectDestroyQSGNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGNode) DisconnectDestroyQSGNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGNode"})
}

func (ptr *QSGNode) DestroyQSGNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGNode"})
}

func (ptr *QSGNode) DestroyQSGNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGNodeDefault"})
}

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func (ptr *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return ptr
}

func (ptr *QSGOpacityNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpacityNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpacityNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpacityNode) ClassNameInternalF() string {
	return n.QSGNode_PTR().ClassNameInternalF()
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpacityNode) {
	n = new(QSGOpacityNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpacityNode")
	return
}
func NewQSGOpacityNode() *QSGOpacityNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGOpacityNode", ""}).(*QSGOpacityNode)
}

func (ptr *QSGOpacityNode) Opacity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Opacity"}).(float64)
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpacity", opacity})
}

func (ptr *QSGOpacityNode) ConnectDestroyQSGOpacityNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGOpacityNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGOpacityNode) DisconnectDestroyQSGOpacityNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGOpacityNode"})
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGOpacityNode"})
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGOpacityNodeDefault"})
}

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterial_ITF interface {
	QSGMaterial_ITF
	QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial
}

func (ptr *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return ptr
}

func (ptr *QSGOpaqueTextureMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpaqueTextureMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpaqueTextureMaterial) ClassNameInternalF() string {
	return n.QSGMaterial_PTR().ClassNameInternalF()
}

func NewQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) (n *QSGOpaqueTextureMaterial) {
	n = new(QSGOpaqueTextureMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpaqueTextureMaterial")
	return
}

func (ptr *QSGOpaqueTextureMaterial) DestroyQSGOpaqueTextureMaterial() {
}

func NewQSGOpaqueTextureMaterial() *QSGOpaqueTextureMaterial {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGOpaqueTextureMaterial", ""}).(*QSGOpaqueTextureMaterial)
}

func (ptr *QSGOpaqueTextureMaterial) AnisotropyLevel() QSGTexture__AnisotropyLevel {

	return QSGTexture__AnisotropyLevel(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnisotropyLevel"}).(float64))
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filtering"}).(float64))
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {

	return QSGTexture__WrapMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HorizontalWrapMode"}).(float64))
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MipmapFiltering"}).(float64))
}

func (ptr *QSGOpaqueTextureMaterial) SetAnisotropyLevel(level QSGTexture__AnisotropyLevel) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnisotropyLevel", level})
}

func (ptr *QSGOpaqueTextureMaterial) SetFiltering(filtering QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFiltering", filtering})
}

func (ptr *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode QSGTexture__WrapMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHorizontalWrapMode", mode})
}

func (ptr *QSGOpaqueTextureMaterial) SetMipmapFiltering(filtering QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMipmapFiltering", filtering})
}

func (ptr *QSGOpaqueTextureMaterial) SetTexture(texture QSGTexture_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTexture", texture})
}

func (ptr *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode QSGTexture__WrapMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVerticalWrapMode", mode})
}

func (ptr *QSGOpaqueTextureMaterial) Texture() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Texture"}).(*QSGTexture)
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {

	return QSGTexture__WrapMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VerticalWrapMode"}).(float64))
}

func (ptr *QSGOpaqueTextureMaterial) CreateShader() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShader"}).(*QSGMaterialShader)
}

func (ptr *QSGOpaqueTextureMaterial) CreateShaderDefault() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShaderDefault"}).(*QSGMaterialShader)
}

func (ptr *QSGOpaqueTextureMaterial) Type() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*QSGMaterialType)
}

func (ptr *QSGOpaqueTextureMaterial) TypeDefault() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(*QSGMaterialType)
}

type QSGOpenVGFontGlyphCache struct {
	internal.Internal
}

type QSGOpenVGFontGlyphCache_ITF interface {
	QSGOpenVGFontGlyphCache_PTR() *QSGOpenVGFontGlyphCache
}

func (ptr *QSGOpenVGFontGlyphCache) QSGOpenVGFontGlyphCache_PTR() *QSGOpenVGFontGlyphCache {
	return ptr
}

func (ptr *QSGOpenVGFontGlyphCache) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGOpenVGFontGlyphCache) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGOpenVGFontGlyphCache(ptr QSGOpenVGFontGlyphCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGFontGlyphCache_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGFontGlyphCache) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGOpenVGFontGlyphCacheFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGFontGlyphCache) {
	n = new(QSGOpenVGFontGlyphCache)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGFontGlyphCache")
	return
}

func (ptr *QSGOpenVGFontGlyphCache) DestroyQSGOpenVGFontGlyphCache() {
}

type QSGOpenVGFontGlyphCacheManager struct {
	internal.Internal
}

type QSGOpenVGFontGlyphCacheManager_ITF interface {
	QSGOpenVGFontGlyphCacheManager_PTR() *QSGOpenVGFontGlyphCacheManager
}

func (ptr *QSGOpenVGFontGlyphCacheManager) QSGOpenVGFontGlyphCacheManager_PTR() *QSGOpenVGFontGlyphCacheManager {
	return ptr
}

func (ptr *QSGOpenVGFontGlyphCacheManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGOpenVGFontGlyphCacheManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGOpenVGFontGlyphCacheManager(ptr QSGOpenVGFontGlyphCacheManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGFontGlyphCacheManager_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGFontGlyphCacheManager) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGOpenVGFontGlyphCacheManagerFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGFontGlyphCacheManager) {
	n = new(QSGOpenVGFontGlyphCacheManager)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGFontGlyphCacheManager")
	return
}

func (ptr *QSGOpenVGFontGlyphCacheManager) DestroyQSGOpenVGFontGlyphCacheManager() {
}

type QSGOpenVGImageNode struct {
	QSGImageNode
	QSGOpenVGRenderable
}

type QSGOpenVGImageNode_ITF interface {
	QSGImageNode_ITF
	QSGOpenVGRenderable_ITF
	QSGOpenVGImageNode_PTR() *QSGOpenVGImageNode
}

func (ptr *QSGOpenVGImageNode) QSGOpenVGImageNode_PTR() *QSGOpenVGImageNode {
	return ptr
}

func (ptr *QSGOpenVGImageNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGImageNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGImageNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGImageNode_PTR().SetPointer(p)
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGImageNode(ptr QSGOpenVGImageNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGImageNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGImageNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGImageNode_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGImageNode) ClassNameInternalF() string {
	return n.QSGImageNode_PTR().ClassNameInternalF()
}

func NewQSGOpenVGImageNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGImageNode) {
	n = new(QSGOpenVGImageNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGImageNode")
	return
}

func (ptr *QSGOpenVGImageNode) DestroyQSGOpenVGImageNode() {
}

type QSGOpenVGInternalImageNode struct {
	QSGOpenVGRenderable
}

type QSGOpenVGInternalImageNode_ITF interface {
	QSGOpenVGRenderable_ITF
	QSGOpenVGInternalImageNode_PTR() *QSGOpenVGInternalImageNode
}

func (ptr *QSGOpenVGInternalImageNode) QSGOpenVGInternalImageNode_PTR() *QSGOpenVGInternalImageNode {
	return ptr
}

func (ptr *QSGOpenVGInternalImageNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGInternalImageNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGInternalImageNode(ptr QSGOpenVGInternalImageNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGInternalImageNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGInternalImageNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGInternalImageNode) ClassNameInternalF() string {
	return n.QSGOpenVGRenderable_PTR().ClassNameInternalF()
}

func NewQSGOpenVGInternalImageNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGInternalImageNode) {
	n = new(QSGOpenVGInternalImageNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGInternalImageNode")
	return
}

func (ptr *QSGOpenVGInternalImageNode) DestroyQSGOpenVGInternalImageNode() {
}

type QSGOpenVGInternalRectangleNode struct {
	QSGOpenVGRenderable
}

type QSGOpenVGInternalRectangleNode_ITF interface {
	QSGOpenVGRenderable_ITF
	QSGOpenVGInternalRectangleNode_PTR() *QSGOpenVGInternalRectangleNode
}

func (ptr *QSGOpenVGInternalRectangleNode) QSGOpenVGInternalRectangleNode_PTR() *QSGOpenVGInternalRectangleNode {
	return ptr
}

func (ptr *QSGOpenVGInternalRectangleNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGInternalRectangleNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGInternalRectangleNode(ptr QSGOpenVGInternalRectangleNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGInternalRectangleNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGInternalRectangleNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGInternalRectangleNode) ClassNameInternalF() string {
	return n.QSGOpenVGRenderable_PTR().ClassNameInternalF()
}

func NewQSGOpenVGInternalRectangleNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGInternalRectangleNode) {
	n = new(QSGOpenVGInternalRectangleNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGInternalRectangleNode")
	return
}

func (ptr *QSGOpenVGInternalRectangleNode) DestroyQSGOpenVGInternalRectangleNode() {
}

type QSGOpenVGLayer struct {
	internal.Internal
}

type QSGOpenVGLayer_ITF interface {
	QSGOpenVGLayer_PTR() *QSGOpenVGLayer
}

func (ptr *QSGOpenVGLayer) QSGOpenVGLayer_PTR() *QSGOpenVGLayer {
	return ptr
}

func (ptr *QSGOpenVGLayer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGOpenVGLayer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGOpenVGLayer(ptr QSGOpenVGLayer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGLayer_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGLayer) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGOpenVGLayerFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGLayer) {
	n = new(QSGOpenVGLayer)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGLayer")
	return
}

func (ptr *QSGOpenVGLayer) DestroyQSGOpenVGLayer() {
}

type QSGOpenVGNinePatchNode struct {
	QSGGeometryNode
	QSGOpenVGRenderable
}

type QSGOpenVGNinePatchNode_ITF interface {
	QSGGeometryNode_ITF
	QSGOpenVGRenderable_ITF
	QSGOpenVGNinePatchNode_PTR() *QSGOpenVGNinePatchNode
}

func (ptr *QSGOpenVGNinePatchNode) QSGOpenVGNinePatchNode_PTR() *QSGOpenVGNinePatchNode {
	return ptr
}

func (ptr *QSGOpenVGNinePatchNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGNinePatchNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGNinePatchNode(ptr QSGOpenVGNinePatchNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGNinePatchNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGNinePatchNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGNinePatchNode) ClassNameInternalF() string {
	return n.QSGGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGOpenVGNinePatchNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGNinePatchNode) {
	n = new(QSGOpenVGNinePatchNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGNinePatchNode")
	return
}

func (ptr *QSGOpenVGNinePatchNode) DestroyQSGOpenVGNinePatchNode() {
}

type QSGOpenVGNodeVisitor struct {
	internal.Internal
}

type QSGOpenVGNodeVisitor_ITF interface {
	QSGOpenVGNodeVisitor_PTR() *QSGOpenVGNodeVisitor
}

func (ptr *QSGOpenVGNodeVisitor) QSGOpenVGNodeVisitor_PTR() *QSGOpenVGNodeVisitor {
	return ptr
}

func (ptr *QSGOpenVGNodeVisitor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGOpenVGNodeVisitor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGOpenVGNodeVisitor(ptr QSGOpenVGNodeVisitor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGNodeVisitor_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGNodeVisitor) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGOpenVGNodeVisitorFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGNodeVisitor) {
	n = new(QSGOpenVGNodeVisitor)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGNodeVisitor")
	return
}

func (ptr *QSGOpenVGNodeVisitor) DestroyQSGOpenVGNodeVisitor() {
}

type QSGOpenVGPainterNode struct {
	QSGOpenVGRenderable
}

type QSGOpenVGPainterNode_ITF interface {
	QSGOpenVGRenderable_ITF
	QSGOpenVGPainterNode_PTR() *QSGOpenVGPainterNode
}

func (ptr *QSGOpenVGPainterNode) QSGOpenVGPainterNode_PTR() *QSGOpenVGPainterNode {
	return ptr
}

func (ptr *QSGOpenVGPainterNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGPainterNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGPainterNode(ptr QSGOpenVGPainterNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGPainterNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGPainterNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGPainterNode) ClassNameInternalF() string {
	return n.QSGOpenVGRenderable_PTR().ClassNameInternalF()
}

func NewQSGOpenVGPainterNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGPainterNode) {
	n = new(QSGOpenVGPainterNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGPainterNode")
	return
}

func (ptr *QSGOpenVGPainterNode) DestroyQSGOpenVGPainterNode() {
}

type QSGOpenVGRectangleNode struct {
	QSGOpenVGRenderable
	QSGRectangleNode
}

type QSGOpenVGRectangleNode_ITF interface {
	QSGOpenVGRenderable_ITF
	QSGRectangleNode_ITF
	QSGOpenVGRectangleNode_PTR() *QSGOpenVGRectangleNode
}

func (ptr *QSGOpenVGRectangleNode) QSGOpenVGRectangleNode_PTR() *QSGOpenVGRectangleNode {
	return ptr
}

func (ptr *QSGOpenVGRectangleNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGRectangleNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
		ptr.QSGRectangleNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGRectangleNode(ptr QSGOpenVGRectangleNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRectangleNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGRectangleNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSGRectangleNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGRectangleNode) ClassNameInternalF() string {
	return n.QSGOpenVGRenderable_PTR().ClassNameInternalF()
}

func NewQSGOpenVGRectangleNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGRectangleNode) {
	n = new(QSGOpenVGRectangleNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGRectangleNode")
	return
}

func (ptr *QSGOpenVGRectangleNode) DestroyQSGOpenVGRectangleNode() {
}

type QSGOpenVGRenderable struct {
	internal.Internal
}

type QSGOpenVGRenderable_ITF interface {
	QSGOpenVGRenderable_PTR() *QSGOpenVGRenderable
}

func (ptr *QSGOpenVGRenderable) QSGOpenVGRenderable_PTR() *QSGOpenVGRenderable {
	return ptr
}

func (ptr *QSGOpenVGRenderable) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGOpenVGRenderable) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGOpenVGRenderable(ptr QSGOpenVGRenderable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGRenderable) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGOpenVGRenderableFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGRenderable) {
	n = new(QSGOpenVGRenderable)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGRenderable")
	return
}

func (ptr *QSGOpenVGRenderable) DestroyQSGOpenVGRenderable() {
}

type QSGOpenVGSpriteNode struct {
	QSGOpenVGRenderable
}

type QSGOpenVGSpriteNode_ITF interface {
	QSGOpenVGRenderable_ITF
	QSGOpenVGSpriteNode_PTR() *QSGOpenVGSpriteNode
}

func (ptr *QSGOpenVGSpriteNode) QSGOpenVGSpriteNode_PTR() *QSGOpenVGSpriteNode {
	return ptr
}

func (ptr *QSGOpenVGSpriteNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGRenderable_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGSpriteNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpenVGRenderable_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGSpriteNode(ptr QSGOpenVGSpriteNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGSpriteNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGSpriteNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpenVGRenderable_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGSpriteNode) ClassNameInternalF() string {
	return n.QSGOpenVGRenderable_PTR().ClassNameInternalF()
}

func NewQSGOpenVGSpriteNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGSpriteNode) {
	n = new(QSGOpenVGSpriteNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGSpriteNode")
	return
}

func (ptr *QSGOpenVGSpriteNode) DestroyQSGOpenVGSpriteNode() {
}

type QSGOpenVGTexture struct {
	QSGTexture
}

type QSGOpenVGTexture_ITF interface {
	QSGTexture_ITF
	QSGOpenVGTexture_PTR() *QSGOpenVGTexture
}

func (ptr *QSGOpenVGTexture) QSGOpenVGTexture_PTR() *QSGOpenVGTexture {
	return ptr
}

func (ptr *QSGOpenVGTexture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpenVGTexture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGTexture_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpenVGTexture(ptr QSGOpenVGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpenVGTexture_PTR().Pointer()
	}
	return nil
}

func (n *QSGOpenVGTexture) InitFromInternal(ptr uintptr, name string) {
	n.QSGTexture_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGOpenVGTexture) ClassNameInternalF() string {
	return n.QSGTexture_PTR().ClassNameInternalF()
}

func NewQSGOpenVGTextureFromPointer(ptr unsafe.Pointer) (n *QSGOpenVGTexture) {
	n = new(QSGOpenVGTexture)
	n.InitFromInternal(uintptr(ptr), "quick.QSGOpenVGTexture")
	return
}

func (ptr *QSGOpenVGTexture) DestroyQSGOpenVGTexture() {
}

type QSGRectangleNode struct {
	QSGGeometryNode
}

type QSGRectangleNode_ITF interface {
	QSGGeometryNode_ITF
	QSGRectangleNode_PTR() *QSGRectangleNode
}

func (ptr *QSGRectangleNode) QSGRectangleNode_PTR() *QSGRectangleNode {
	return ptr
}

func (ptr *QSGRectangleNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGRectangleNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGRectangleNode(ptr QSGRectangleNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRectangleNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGRectangleNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGRectangleNode) ClassNameInternalF() string {
	return n.QSGGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGRectangleNodeFromPointer(ptr unsafe.Pointer) (n *QSGRectangleNode) {
	n = new(QSGRectangleNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGRectangleNode")
	return
}

func (ptr *QSGRectangleNode) DestroyQSGRectangleNode() {
}

func (ptr *QSGRectangleNode) ConnectColor(f func() *gui.QColor) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRectangleNode) DisconnectColor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColor"})
}

func (ptr *QSGRectangleNode) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QSGRectangleNode) ConnectRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRectangleNode) DisconnectRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRect"})
}

func (ptr *QSGRectangleNode) Rect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rect"}).(*core.QRectF)
}

func (ptr *QSGRectangleNode) ConnectSetColor(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetColor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRectangleNode) DisconnectSetColor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetColor"})
}

func (ptr *QSGRectangleNode) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QSGRectangleNode) ConnectSetRect(f func(rect *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRectangleNode) DisconnectSetRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRect"})
}

func (ptr *QSGRectangleNode) SetRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect", rect})
}

func (ptr *QSGRectangleNode) SetRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect2", x, y, w, h})
}

type QSGRenderNode struct {
	QSGNode
}

type QSGRenderNode_ITF interface {
	QSGNode_ITF
	QSGRenderNode_PTR() *QSGRenderNode
}

func (ptr *QSGRenderNode) QSGRenderNode_PTR() *QSGRenderNode {
	return ptr
}

func (ptr *QSGRenderNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGRenderNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGRenderNode(ptr QSGRenderNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRenderNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGRenderNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGRenderNode) ClassNameInternalF() string {
	return n.QSGNode_PTR().ClassNameInternalF()
}

func NewQSGRenderNodeFromPointer(ptr unsafe.Pointer) (n *QSGRenderNode) {
	n = new(QSGRenderNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGRenderNode")
	return
}

//go:generate stringer -type=QSGRenderNode__StateFlag
//QSGRenderNode::StateFlag
type QSGRenderNode__StateFlag int64

const (
	QSGRenderNode__DepthState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x01)
	QSGRenderNode__StencilState      QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x02)
	QSGRenderNode__ScissorState      QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x04)
	QSGRenderNode__ColorState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x08)
	QSGRenderNode__BlendState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x10)
	QSGRenderNode__CullState         QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x20)
	QSGRenderNode__ViewportState     QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x40)
	QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x80)
)

//go:generate stringer -type=QSGRenderNode__RenderingFlag
//QSGRenderNode::RenderingFlag
type QSGRenderNode__RenderingFlag int64

const (
	QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x01)
	QSGRenderNode__DepthAwareRendering  QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x02)
	QSGRenderNode__OpaqueRendering      QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x04)
)

func (ptr *QSGRenderNode) ConnectChangedStates(f func() QSGRenderNode__StateFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectChangedStates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRenderNode) DisconnectChangedStates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectChangedStates"})
}

func (ptr *QSGRenderNode) ChangedStates() QSGRenderNode__StateFlag {

	return QSGRenderNode__StateFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangedStates"}).(float64))
}

func (ptr *QSGRenderNode) ChangedStatesDefault() QSGRenderNode__StateFlag {

	return QSGRenderNode__StateFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangedStatesDefault"}).(float64))
}

func (ptr *QSGRenderNode) ClipList() *QSGClipNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClipList"}).(*QSGClipNode)
}

func (ptr *QSGRenderNode) ConnectFlags(f func() QSGRenderNode__RenderingFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlags", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRenderNode) DisconnectFlags() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlags"})
}

func (ptr *QSGRenderNode) Flags() QSGRenderNode__RenderingFlag {

	return QSGRenderNode__RenderingFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QSGRenderNode) FlagsDefault() QSGRenderNode__RenderingFlag {

	return QSGRenderNode__RenderingFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault"}).(float64))
}

func (ptr *QSGRenderNode) InheritedOpacity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InheritedOpacity"}).(float64)
}

func (ptr *QSGRenderNode) Matrix() *gui.QMatrix4x4 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Matrix"}).(*gui.QMatrix4x4)
}

func (ptr *QSGRenderNode) ConnectRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRenderNode) DisconnectRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRect"})
}

func (ptr *QSGRenderNode) Rect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rect"}).(*core.QRectF)
}

func (ptr *QSGRenderNode) RectDefault() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RectDefault"}).(*core.QRectF)
}

func (ptr *QSGRenderNode) ConnectReleaseResources(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleaseResources", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRenderNode) DisconnectReleaseResources() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleaseResources"})
}

func (ptr *QSGRenderNode) ReleaseResources() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseResources"})
}

func (ptr *QSGRenderNode) ReleaseResourcesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseResourcesDefault"})
}

func (ptr *QSGRenderNode) ConnectDestroyQSGRenderNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGRenderNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRenderNode) DisconnectDestroyQSGRenderNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGRenderNode"})
}

func (ptr *QSGRenderNode) DestroyQSGRenderNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGRenderNode"})
}

func (ptr *QSGRenderNode) DestroyQSGRenderNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGRenderNodeDefault"})
}

type QSGRendererInterface struct {
	internal.Internal
}

type QSGRendererInterface_ITF interface {
	QSGRendererInterface_PTR() *QSGRendererInterface
}

func (ptr *QSGRendererInterface) QSGRendererInterface_PTR() *QSGRendererInterface {
	return ptr
}

func (ptr *QSGRendererInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSGRendererInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSGRendererInterface(ptr QSGRendererInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRendererInterface_PTR().Pointer()
	}
	return nil
}

func (n *QSGRendererInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSGRendererInterfaceFromPointer(ptr unsafe.Pointer) (n *QSGRendererInterface) {
	n = new(QSGRendererInterface)
	n.InitFromInternal(uintptr(ptr), "quick.QSGRendererInterface")
	return
}

func (ptr *QSGRendererInterface) DestroyQSGRendererInterface() {
}

//go:generate stringer -type=QSGRendererInterface__GraphicsApi
//QSGRendererInterface::GraphicsApi
type QSGRendererInterface__GraphicsApi int64

const (
	QSGRendererInterface__Unknown    QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(0)
	QSGRendererInterface__Software   QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(1)
	QSGRendererInterface__OpenGL     QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(2)
	QSGRendererInterface__Direct3D12 QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(3)
	QSGRendererInterface__OpenVG     QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(4)
)

//go:generate stringer -type=QSGRendererInterface__Resource
//QSGRendererInterface::Resource
type QSGRendererInterface__Resource int64

const (
	QSGRendererInterface__DeviceResource       QSGRendererInterface__Resource = QSGRendererInterface__Resource(0)
	QSGRendererInterface__CommandQueueResource QSGRendererInterface__Resource = QSGRendererInterface__Resource(1)
	QSGRendererInterface__CommandListResource  QSGRendererInterface__Resource = QSGRendererInterface__Resource(2)
	QSGRendererInterface__PainterResource      QSGRendererInterface__Resource = QSGRendererInterface__Resource(3)
)

//go:generate stringer -type=QSGRendererInterface__ShaderType
//QSGRendererInterface::ShaderType
type QSGRendererInterface__ShaderType int64

const (
	QSGRendererInterface__UnknownShadingLanguage QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(0)
	QSGRendererInterface__GLSL                   QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(1)
	QSGRendererInterface__HLSL                   QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(2)
)

//go:generate stringer -type=QSGRendererInterface__ShaderCompilationType
//QSGRendererInterface::ShaderCompilationType
type QSGRendererInterface__ShaderCompilationType int64

const (
	QSGRendererInterface__RuntimeCompilation QSGRendererInterface__ShaderCompilationType = QSGRendererInterface__ShaderCompilationType(0x01)
	QSGRendererInterface__OfflineCompilation QSGRendererInterface__ShaderCompilationType = QSGRendererInterface__ShaderCompilationType(0x02)
)

//go:generate stringer -type=QSGRendererInterface__ShaderSourceType
//QSGRendererInterface::ShaderSourceType
type QSGRendererInterface__ShaderSourceType int64

const (
	QSGRendererInterface__ShaderSourceString QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x01)
	QSGRendererInterface__ShaderSourceFile   QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x02)
	QSGRendererInterface__ShaderByteCode     QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x04)
)

func (ptr *QSGRendererInterface) ConnectGetResource(f func(window *QQuickWindow, resource QSGRendererInterface__Resource) unsafe.Pointer) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGetResource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectGetResource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGetResource"})
}

func (ptr *QSGRendererInterface) GetResource(window QQuickWindow_ITF, resource QSGRendererInterface__Resource) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetResource", window, resource}).(unsafe.Pointer)
}

func (ptr *QSGRendererInterface) GetResourceDefault(window QQuickWindow_ITF, resource QSGRendererInterface__Resource) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetResourceDefault", window, resource}).(unsafe.Pointer)
}

func (ptr *QSGRendererInterface) ConnectGetResource2(f func(window *QQuickWindow, resource string) unsafe.Pointer) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGetResource2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectGetResource2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGetResource2"})
}

func (ptr *QSGRendererInterface) GetResource2(window QQuickWindow_ITF, resource string) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetResource2", window, resource}).(unsafe.Pointer)
}

func (ptr *QSGRendererInterface) GetResource2Default(window QQuickWindow_ITF, resource string) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetResource2Default", window, resource}).(unsafe.Pointer)
}

func (ptr *QSGRendererInterface) ConnectGraphicsApi(f func() QSGRendererInterface__GraphicsApi) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGraphicsApi", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectGraphicsApi() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGraphicsApi"})
}

func (ptr *QSGRendererInterface) GraphicsApi() QSGRendererInterface__GraphicsApi {

	return QSGRendererInterface__GraphicsApi(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GraphicsApi"}).(float64))
}

func (ptr *QSGRendererInterface) ConnectShaderCompilationType(f func() QSGRendererInterface__ShaderCompilationType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShaderCompilationType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectShaderCompilationType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShaderCompilationType"})
}

func (ptr *QSGRendererInterface) ShaderCompilationType() QSGRendererInterface__ShaderCompilationType {

	return QSGRendererInterface__ShaderCompilationType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShaderCompilationType"}).(float64))
}

func (ptr *QSGRendererInterface) ConnectShaderSourceType(f func() QSGRendererInterface__ShaderSourceType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShaderSourceType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectShaderSourceType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShaderSourceType"})
}

func (ptr *QSGRendererInterface) ShaderSourceType() QSGRendererInterface__ShaderSourceType {

	return QSGRendererInterface__ShaderSourceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShaderSourceType"}).(float64))
}

func (ptr *QSGRendererInterface) ConnectShaderType(f func() QSGRendererInterface__ShaderType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShaderType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGRendererInterface) DisconnectShaderType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShaderType"})
}

func (ptr *QSGRendererInterface) ShaderType() QSGRendererInterface__ShaderType {

	return QSGRendererInterface__ShaderType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShaderType"}).(float64))
}

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func (ptr *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return ptr
}

func (ptr *QSGSimpleMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGSimpleMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGSimpleMaterial) ClassNameInternalF() string {
	return n.QSGMaterial_PTR().ClassNameInternalF()
}

func NewQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) (n *QSGSimpleMaterial) {
	n = new(QSGSimpleMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGSimpleMaterial")
	return
}

func (ptr *QSGSimpleMaterial) DestroyQSGSimpleMaterial() {
}

type QSGSimpleMaterialComparableMaterial struct {
	QSGSimpleMaterial
}

type QSGSimpleMaterialComparableMaterial_ITF interface {
	QSGSimpleMaterial_ITF
	QSGSimpleMaterialComparableMaterial_PTR() *QSGSimpleMaterialComparableMaterial
}

func (ptr *QSGSimpleMaterialComparableMaterial) QSGSimpleMaterialComparableMaterial_PTR() *QSGSimpleMaterialComparableMaterial {
	return ptr
}

func (ptr *QSGSimpleMaterialComparableMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleMaterialComparableMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGSimpleMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleMaterialComparableMaterial(ptr QSGSimpleMaterialComparableMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialComparableMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGSimpleMaterialComparableMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGSimpleMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGSimpleMaterialComparableMaterial) ClassNameInternalF() string {
	return n.QSGSimpleMaterial_PTR().ClassNameInternalF()
}

func NewQSGSimpleMaterialComparableMaterialFromPointer(ptr unsafe.Pointer) (n *QSGSimpleMaterialComparableMaterial) {
	n = new(QSGSimpleMaterialComparableMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGSimpleMaterialComparableMaterial")
	return
}

func (ptr *QSGSimpleMaterialComparableMaterial) DestroyQSGSimpleMaterialComparableMaterial() {
}

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func (ptr *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return ptr
}

func (ptr *QSGSimpleMaterialShader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleMaterialShader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterialShader_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShader_PTR().Pointer()
	}
	return nil
}

func (n *QSGSimpleMaterialShader) InitFromInternal(ptr uintptr, name string) {
	n.QSGMaterialShader_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGSimpleMaterialShader) ClassNameInternalF() string {
	return n.QSGMaterialShader_PTR().ClassNameInternalF()
}

func NewQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) (n *QSGSimpleMaterialShader) {
	n = new(QSGSimpleMaterialShader)
	n.InitFromInternal(uintptr(ptr), "quick.QSGSimpleMaterialShader")
	return
}

func (ptr *QSGSimpleMaterialShader) DestroyQSGSimpleMaterialShader() {
}

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func (ptr *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return ptr
}

func (ptr *QSGSimpleRectNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleRectNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGSimpleRectNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGSimpleRectNode) ClassNameInternalF() string {
	return n.QSGGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) (n *QSGSimpleRectNode) {
	n = new(QSGSimpleRectNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGSimpleRectNode")
	return
}

func (ptr *QSGSimpleRectNode) DestroyQSGSimpleRectNode() {
}

func NewQSGSimpleRectNode(rect core.QRectF_ITF, color gui.QColor_ITF) *QSGSimpleRectNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGSimpleRectNode", "", rect, color}).(*QSGSimpleRectNode)
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGSimpleRectNode2", ""}).(*QSGSimpleRectNode)
}

func (ptr *QSGSimpleRectNode) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QSGSimpleRectNode) Rect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rect"}).(*core.QRectF)
}

func (ptr *QSGSimpleRectNode) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QSGSimpleRectNode) SetRect(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect", rect})
}

func (ptr *QSGSimpleRectNode) SetRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect2", x, y, w, h})
}

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return ptr
}

func (ptr *QSGSimpleTextureNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGSimpleTextureNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGGeometryNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGSimpleTextureNode) ClassNameInternalF() string {
	return n.QSGGeometryNode_PTR().ClassNameInternalF()
}

func NewQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) (n *QSGSimpleTextureNode) {
	n = new(QSGSimpleTextureNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGSimpleTextureNode")
	return
}

//go:generate stringer -type=QSGSimpleTextureNode__TextureCoordinatesTransformFlag
//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGSimpleTextureNode", ""}).(*QSGSimpleTextureNode)
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filtering"}).(float64))
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OwnsTexture"}).(bool)
}

func (ptr *QSGSimpleTextureNode) Rect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rect"}).(*core.QRectF)
}

func (ptr *QSGSimpleTextureNode) SetFiltering(filtering QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFiltering", filtering})
}

func (ptr *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOwnsTexture", owns})
}

func (ptr *QSGSimpleTextureNode) SetRect(r core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect", r})
}

func (ptr *QSGSimpleTextureNode) SetRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRect2", x, y, w, h})
}

func (ptr *QSGSimpleTextureNode) SetSourceRect(r core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceRect", r})
}

func (ptr *QSGSimpleTextureNode) SetSourceRect2(x float64, y float64, w float64, h float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceRect2", x, y, w, h})
}

func (ptr *QSGSimpleTextureNode) SetTexture(texture QSGTexture_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTexture", texture})
}

func (ptr *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode QSGSimpleTextureNode__TextureCoordinatesTransformFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureCoordinatesTransform", mode})
}

func (ptr *QSGSimpleTextureNode) SourceRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceRect"}).(*core.QRectF)
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Texture"}).(*QSGTexture)
}

func (ptr *QSGSimpleTextureNode) TextureCoordinatesTransform() QSGSimpleTextureNode__TextureCoordinatesTransformFlag {

	return QSGSimpleTextureNode__TextureCoordinatesTransformFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureCoordinatesTransform"}).(float64))
}

func (ptr *QSGSimpleTextureNode) ConnectDestroyQSGSimpleTextureNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGSimpleTextureNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGSimpleTextureNode) DisconnectDestroyQSGSimpleTextureNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGSimpleTextureNode"})
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGSimpleTextureNode"})
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGSimpleTextureNodeDefault"})
}

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func (ptr *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return ptr
}

func (ptr *QSGTexture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTexture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGTexture(ptr QSGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (n *QSGTexture) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGTexture) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSGTextureFromPointer(ptr unsafe.Pointer) (n *QSGTexture) {
	n = new(QSGTexture)
	n.InitFromInternal(uintptr(ptr), "quick.QSGTexture")
	return
}

//go:generate stringer -type=QSGTexture__WrapMode
//QSGTexture::WrapMode
type QSGTexture__WrapMode int64

const (
	QSGTexture__Repeat         QSGTexture__WrapMode = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge    QSGTexture__WrapMode = QSGTexture__WrapMode(1)
	QSGTexture__MirroredRepeat QSGTexture__WrapMode = QSGTexture__WrapMode(2)
)

//go:generate stringer -type=QSGTexture__Filtering
//QSGTexture::Filtering
type QSGTexture__Filtering int64

const (
	QSGTexture__None    QSGTexture__Filtering = QSGTexture__Filtering(0)
	QSGTexture__Nearest QSGTexture__Filtering = QSGTexture__Filtering(1)
	QSGTexture__Linear  QSGTexture__Filtering = QSGTexture__Filtering(2)
)

//go:generate stringer -type=QSGTexture__AnisotropyLevel
//QSGTexture::AnisotropyLevel
type QSGTexture__AnisotropyLevel int64

const (
	QSGTexture__AnisotropyNone QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(0)
	QSGTexture__Anisotropy2x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(1)
	QSGTexture__Anisotropy4x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(2)
	QSGTexture__Anisotropy8x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(3)
	QSGTexture__Anisotropy16x  QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(4)
)

func NewQSGTexture() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGTexture", ""}).(*QSGTexture)
}

func (ptr *QSGTexture) AnisotropyLevel() QSGTexture__AnisotropyLevel {

	return QSGTexture__AnisotropyLevel(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnisotropyLevel"}).(float64))
}

func (ptr *QSGTexture) ConnectBind(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBind", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectBind() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBind"})
}

func (ptr *QSGTexture) Bind() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Bind"})
}

func (ptr *QSGTexture) ConvertToNormalizedSourceRect(rect core.QRectF_ITF) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConvertToNormalizedSourceRect", rect}).(*core.QRectF)
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filtering"}).(float64))
}

func (ptr *QSGTexture) ConnectHasAlphaChannel(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasAlphaChannel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectHasAlphaChannel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasAlphaChannel"})
}

func (ptr *QSGTexture) HasAlphaChannel() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAlphaChannel"}).(bool)
}

func (ptr *QSGTexture) ConnectHasMipmaps(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasMipmaps", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectHasMipmaps() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasMipmaps"})
}

func (ptr *QSGTexture) HasMipmaps() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasMipmaps"}).(bool)
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {

	return QSGTexture__WrapMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HorizontalWrapMode"}).(float64))
}

func (ptr *QSGTexture) ConnectIsAtlasTexture(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsAtlasTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectIsAtlasTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsAtlasTexture"})
}

func (ptr *QSGTexture) IsAtlasTexture() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAtlasTexture"}).(bool)
}

func (ptr *QSGTexture) IsAtlasTextureDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAtlasTextureDefault"}).(bool)
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {

	return QSGTexture__Filtering(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MipmapFiltering"}).(float64))
}

func (ptr *QSGTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNormalizedTextureSubRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectNormalizedTextureSubRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNormalizedTextureSubRect"})
}

func (ptr *QSGTexture) NormalizedTextureSubRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NormalizedTextureSubRect"}).(*core.QRectF)
}

func (ptr *QSGTexture) NormalizedTextureSubRectDefault() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NormalizedTextureSubRectDefault"}).(*core.QRectF)
}

func (ptr *QSGTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemovedFromAtlas", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectRemovedFromAtlas() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemovedFromAtlas"})
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemovedFromAtlas"}).(*QSGTexture)
}

func (ptr *QSGTexture) RemovedFromAtlasDefault() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemovedFromAtlasDefault"}).(*QSGTexture)
}

func (ptr *QSGTexture) SetAnisotropyLevel(level QSGTexture__AnisotropyLevel) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnisotropyLevel", level})
}

func (ptr *QSGTexture) SetFiltering(filter QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFiltering", filter})
}

func (ptr *QSGTexture) SetHorizontalWrapMode(hwrap QSGTexture__WrapMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHorizontalWrapMode", hwrap})
}

func (ptr *QSGTexture) SetMipmapFiltering(filter QSGTexture__Filtering) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMipmapFiltering", filter})
}

func (ptr *QSGTexture) SetVerticalWrapMode(vwrap QSGTexture__WrapMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVerticalWrapMode", vwrap})
}

func (ptr *QSGTexture) ConnectTextureId(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectTextureId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureId"})
}

func (ptr *QSGTexture) TextureId() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureId"}).(float64))
}

func (ptr *QSGTexture) ConnectTextureSize(f func() *core.QSize) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectTextureSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureSize"})
}

func (ptr *QSGTexture) TextureSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureSize"}).(*core.QSize)
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateBindOptions", force})
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {

	return QSGTexture__WrapMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VerticalWrapMode"}).(float64))
}

func (ptr *QSGTexture) ConnectDestroyQSGTexture(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTexture) DisconnectDestroyQSGTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGTexture"})
}

func (ptr *QSGTexture) DestroyQSGTexture() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGTexture"})
}

func (ptr *QSGTexture) DestroyQSGTextureDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGTextureDefault"})
}

func (ptr *QSGTexture) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSGTexture) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSGTexture) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTexture) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSGTexture) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSGTexture) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTexture) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSGTexture) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSGTexture) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTexture) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSGTexture) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSGTexture) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSGTexture) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSGTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSGTexture) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSGTexture) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSGTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSGTexture) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSGTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSGTexture) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSGTexture) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func (ptr *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return ptr
}

func (ptr *QSGTextureMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTextureMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpaqueTextureMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGTextureMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGOpaqueTextureMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGTextureMaterial) ClassNameInternalF() string {
	return n.QSGOpaqueTextureMaterial_PTR().ClassNameInternalF()
}

func NewQSGTextureMaterialFromPointer(ptr unsafe.Pointer) (n *QSGTextureMaterial) {
	n = new(QSGTextureMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGTextureMaterial")
	return
}

func (ptr *QSGTextureMaterial) DestroyQSGTextureMaterial() {
}

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTextureProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func (n *QSGTextureProvider) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGTextureProvider) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) (n *QSGTextureProvider) {
	n = new(QSGTextureProvider)
	n.InitFromInternal(uintptr(ptr), "quick.QSGTextureProvider")
	return
}

func (ptr *QSGTextureProvider) DestroyQSGTextureProvider() {
}

func (ptr *QSGTextureProvider) ConnectTexture(f func() *QSGTexture) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTexture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTextureProvider) DisconnectTexture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTexture"})
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Texture"}).(*QSGTexture)
}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureChanged"})
}

func (ptr *QSGTextureProvider) TextureChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureChanged"})
}

func (ptr *QSGTextureProvider) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSGTextureProvider) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSGTextureProvider) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTextureProvider) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSGTextureProvider) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSGTextureProvider) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSGTextureProvider) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSGTextureProvider) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSGTextureProvider) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSGTextureProvider) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSGTextureProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSGTextureProvider) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSGTextureProvider) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSGTextureProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSGTextureProvider) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSGTextureProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSGTextureProvider) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSGTextureProvider) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func (ptr *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return ptr
}

func (ptr *QSGTransformNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTransformNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGTransformNode(ptr QSGTransformNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNode_PTR().Pointer()
	}
	return nil
}

func (n *QSGTransformNode) InitFromInternal(ptr uintptr, name string) {
	n.QSGNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGTransformNode) ClassNameInternalF() string {
	return n.QSGNode_PTR().ClassNameInternalF()
}

func NewQSGTransformNodeFromPointer(ptr unsafe.Pointer) (n *QSGTransformNode) {
	n = new(QSGTransformNode)
	n.InitFromInternal(uintptr(ptr), "quick.QSGTransformNode")
	return
}
func NewQSGTransformNode() *QSGTransformNode {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGTransformNode", ""}).(*QSGTransformNode)
}

func (ptr *QSGTransformNode) Matrix() *gui.QMatrix4x4 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Matrix"}).(*gui.QMatrix4x4)
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMatrix", matrix})
}

func (ptr *QSGTransformNode) ConnectDestroyQSGTransformNode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSGTransformNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSGTransformNode) DisconnectDestroyQSGTransformNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSGTransformNode"})
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGTransformNode"})
}

func (ptr *QSGTransformNode) DestroyQSGTransformNodeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSGTransformNodeDefault"})
}

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return ptr
}

func (ptr *QSGVertexColorMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterial_PTR().Pointer()
	}
	return nil
}

func (n *QSGVertexColorMaterial) InitFromInternal(ptr uintptr, name string) {
	n.QSGMaterial_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSGVertexColorMaterial) ClassNameInternalF() string {
	return n.QSGMaterial_PTR().ClassNameInternalF()
}

func NewQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) (n *QSGVertexColorMaterial) {
	n = new(QSGVertexColorMaterial)
	n.InitFromInternal(uintptr(ptr), "quick.QSGVertexColorMaterial")
	return
}

func (ptr *QSGVertexColorMaterial) DestroyQSGVertexColorMaterial() {
}

func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {

	return internal.CallLocalFunction([]interface{}{"", "", "quick.NewQSGVertexColorMaterial", ""}).(*QSGVertexColorMaterial)
}

func (ptr *QSGVertexColorMaterial) CreateShader() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShader"}).(*QSGMaterialShader)
}

func (ptr *QSGVertexColorMaterial) CreateShaderDefault() *QSGMaterialShader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateShaderDefault"}).(*QSGMaterialShader)
}

func (ptr *QSGVertexColorMaterial) Type() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*QSGMaterialType)
}

func (ptr *QSGVertexColorMaterial) TypeDefault() *QSGMaterialType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(*QSGMaterialType)
}

type QTcpServerConnectionFactory struct {
	internal.Internal
}

type QTcpServerConnectionFactory_ITF interface {
	QTcpServerConnectionFactory_PTR() *QTcpServerConnectionFactory
}

func (ptr *QTcpServerConnectionFactory) QTcpServerConnectionFactory_PTR() *QTcpServerConnectionFactory {
	return ptr
}

func (ptr *QTcpServerConnectionFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTcpServerConnectionFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTcpServerConnectionFactory(ptr QTcpServerConnectionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServerConnectionFactory_PTR().Pointer()
	}
	return nil
}

func (n *QTcpServerConnectionFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTcpServerConnectionFactoryFromPointer(ptr unsafe.Pointer) (n *QTcpServerConnectionFactory) {
	n = new(QTcpServerConnectionFactory)
	n.InitFromInternal(uintptr(ptr), "quick.QTcpServerConnectionFactory")
	return
}

func (ptr *QTcpServerConnectionFactory) DestroyQTcpServerConnectionFactory() {
}

type QV4DataCollector struct {
	internal.Internal
}

type QV4DataCollector_ITF interface {
	QV4DataCollector_PTR() *QV4DataCollector
}

func (ptr *QV4DataCollector) QV4DataCollector_PTR() *QV4DataCollector {
	return ptr
}

func (ptr *QV4DataCollector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QV4DataCollector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQV4DataCollector(ptr QV4DataCollector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DataCollector_PTR().Pointer()
	}
	return nil
}

func (n *QV4DataCollector) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQV4DataCollectorFromPointer(ptr unsafe.Pointer) (n *QV4DataCollector) {
	n = new(QV4DataCollector)
	n.InitFromInternal(uintptr(ptr), "quick.QV4DataCollector")
	return
}

func (ptr *QV4DataCollector) DestroyQV4DataCollector() {
}

type QV4DebugJob struct {
	internal.Internal
}

type QV4DebugJob_ITF interface {
	QV4DebugJob_PTR() *QV4DebugJob
}

func (ptr *QV4DebugJob) QV4DebugJob_PTR() *QV4DebugJob {
	return ptr
}

func (ptr *QV4DebugJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QV4DebugJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQV4DebugJob(ptr QV4DebugJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebugJob_PTR().Pointer()
	}
	return nil
}

func (n *QV4DebugJob) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQV4DebugJobFromPointer(ptr unsafe.Pointer) (n *QV4DebugJob) {
	n = new(QV4DebugJob)
	n.InitFromInternal(uintptr(ptr), "quick.QV4DebugJob")
	return
}

func (ptr *QV4DebugJob) DestroyQV4DebugJob() {
}

type QV4DebugServiceImpl struct {
	internal.Internal
}

type QV4DebugServiceImpl_ITF interface {
	QV4DebugServiceImpl_PTR() *QV4DebugServiceImpl
}

func (ptr *QV4DebugServiceImpl) QV4DebugServiceImpl_PTR() *QV4DebugServiceImpl {
	return ptr
}

func (ptr *QV4DebugServiceImpl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QV4DebugServiceImpl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQV4DebugServiceImpl(ptr QV4DebugServiceImpl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebugServiceImpl_PTR().Pointer()
	}
	return nil
}

func (n *QV4DebugServiceImpl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQV4DebugServiceImplFromPointer(ptr unsafe.Pointer) (n *QV4DebugServiceImpl) {
	n = new(QV4DebugServiceImpl)
	n.InitFromInternal(uintptr(ptr), "quick.QV4DebugServiceImpl")
	return
}

func (ptr *QV4DebugServiceImpl) DestroyQV4DebugServiceImpl() {
}

type QV4Debugger struct {
	internal.Internal
}

type QV4Debugger_ITF interface {
	QV4Debugger_PTR() *QV4Debugger
}

func (ptr *QV4Debugger) QV4Debugger_PTR() *QV4Debugger {
	return ptr
}

func (ptr *QV4Debugger) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QV4Debugger) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQV4Debugger(ptr QV4Debugger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4Debugger_PTR().Pointer()
	}
	return nil
}

func (n *QV4Debugger) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQV4DebuggerFromPointer(ptr unsafe.Pointer) (n *QV4Debugger) {
	n = new(QV4Debugger)
	n.InitFromInternal(uintptr(ptr), "quick.QV4Debugger")
	return
}

func (ptr *QV4Debugger) DestroyQV4Debugger() {
}

type QV4DebuggerAgent struct {
	core.QObject
}

type QV4DebuggerAgent_ITF interface {
	core.QObject_ITF
	QV4DebuggerAgent_PTR() *QV4DebuggerAgent
}

func (ptr *QV4DebuggerAgent) QV4DebuggerAgent_PTR() *QV4DebuggerAgent {
	return ptr
}

func (ptr *QV4DebuggerAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QV4DebuggerAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQV4DebuggerAgent(ptr QV4DebuggerAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4DebuggerAgent_PTR().Pointer()
	}
	return nil
}

func (n *QV4DebuggerAgent) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QV4DebuggerAgent) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQV4DebuggerAgentFromPointer(ptr unsafe.Pointer) (n *QV4DebuggerAgent) {
	n = new(QV4DebuggerAgent)
	n.InitFromInternal(uintptr(ptr), "quick.QV4DebuggerAgent")
	return
}

func (ptr *QV4DebuggerAgent) DestroyQV4DebuggerAgent() {
}

type QV4ProfilerAdapter struct {
	internal.Internal
}

type QV4ProfilerAdapter_ITF interface {
	QV4ProfilerAdapter_PTR() *QV4ProfilerAdapter
}

func (ptr *QV4ProfilerAdapter) QV4ProfilerAdapter_PTR() *QV4ProfilerAdapter {
	return ptr
}

func (ptr *QV4ProfilerAdapter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QV4ProfilerAdapter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQV4ProfilerAdapter(ptr QV4ProfilerAdapter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4ProfilerAdapter_PTR().Pointer()
	}
	return nil
}

func (n *QV4ProfilerAdapter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQV4ProfilerAdapterFromPointer(ptr unsafe.Pointer) (n *QV4ProfilerAdapter) {
	n = new(QV4ProfilerAdapter)
	n.InitFromInternal(uintptr(ptr), "quick.QV4ProfilerAdapter")
	return
}

func (ptr *QV4ProfilerAdapter) DestroyQV4ProfilerAdapter() {
}

type QWavefrontMesh struct {
	internal.Internal
}

type QWavefrontMesh_ITF interface {
	QWavefrontMesh_PTR() *QWavefrontMesh
}

func (ptr *QWavefrontMesh) QWavefrontMesh_PTR() *QWavefrontMesh {
	return ptr
}

func (ptr *QWavefrontMesh) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWavefrontMesh) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWavefrontMesh(ptr QWavefrontMesh_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWavefrontMesh_PTR().Pointer()
	}
	return nil
}

func (n *QWavefrontMesh) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWavefrontMeshFromPointer(ptr unsafe.Pointer) (n *QWavefrontMesh) {
	n = new(QWavefrontMesh)
	n.InitFromInternal(uintptr(ptr), "quick.QWavefrontMesh")
	return
}

func (ptr *QWavefrontMesh) DestroyQWavefrontMesh() {
}

type ScopeJob struct {
	CollectJob
}

type ScopeJob_ITF interface {
	CollectJob_ITF
	ScopeJob_PTR() *ScopeJob
}

func (ptr *ScopeJob) ScopeJob_PTR() *ScopeJob {
	return ptr
}

func (ptr *ScopeJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.CollectJob_PTR().Pointer()
	}
	return nil
}

func (ptr *ScopeJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.CollectJob_PTR().SetPointer(p)
	}
}

func PointerFromScopeJob(ptr ScopeJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ScopeJob_PTR().Pointer()
	}
	return nil
}

func (n *ScopeJob) InitFromInternal(ptr uintptr, name string) {
	n.CollectJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *ScopeJob) ClassNameInternalF() string {
	return n.CollectJob_PTR().ClassNameInternalF()
}

func NewScopeJobFromPointer(ptr unsafe.Pointer) (n *ScopeJob) {
	n = new(ScopeJob)
	n.InitFromInternal(uintptr(ptr), "quick.ScopeJob")
	return
}

func (ptr *ScopeJob) DestroyScopeJob() {
}

type SharedImageProvider struct {
	internal.Internal
}

type SharedImageProvider_ITF interface {
	SharedImageProvider_PTR() *SharedImageProvider
}

func (ptr *SharedImageProvider) SharedImageProvider_PTR() *SharedImageProvider {
	return ptr
}

func (ptr *SharedImageProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *SharedImageProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromSharedImageProvider(ptr SharedImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SharedImageProvider_PTR().Pointer()
	}
	return nil
}

func (n *SharedImageProvider) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewSharedImageProviderFromPointer(ptr unsafe.Pointer) (n *SharedImageProvider) {
	n = new(SharedImageProvider)
	n.InitFromInternal(uintptr(ptr), "quick.SharedImageProvider")
	return
}

func (ptr *SharedImageProvider) DestroySharedImageProvider() {
}

type ValueLookupJob struct {
	CollectJob
}

type ValueLookupJob_ITF interface {
	CollectJob_ITF
	ValueLookupJob_PTR() *ValueLookupJob
}

func (ptr *ValueLookupJob) ValueLookupJob_PTR() *ValueLookupJob {
	return ptr
}

func (ptr *ValueLookupJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.CollectJob_PTR().Pointer()
	}
	return nil
}

func (ptr *ValueLookupJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.CollectJob_PTR().SetPointer(p)
	}
}

func PointerFromValueLookupJob(ptr ValueLookupJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ValueLookupJob_PTR().Pointer()
	}
	return nil
}

func (n *ValueLookupJob) InitFromInternal(ptr uintptr, name string) {
	n.CollectJob_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *ValueLookupJob) ClassNameInternalF() string {
	return n.CollectJob_PTR().ClassNameInternalF()
}

func NewValueLookupJobFromPointer(ptr unsafe.Pointer) (n *ValueLookupJob) {
	n = new(ValueLookupJob)
	n.InitFromInternal(uintptr(ptr), "quick.ValueLookupJob")
	return
}

func (ptr *ValueLookupJob) DestroyValueLookupJob() {
}

func init() {
	internal.ConstructorTable["quick.QDebugMessageServiceFactory"] = NewQDebugMessageServiceFactoryFromPointer
	internal.ConstructorTable["quick.QDebugMessageServiceImpl"] = NewQDebugMessageServiceImplFromPointer
	internal.ConstructorTable["quick.QLocalClientConnectionFactory"] = NewQLocalClientConnectionFactoryFromPointer
	internal.ConstructorTable["quick.QOpenVGOffscreenSurface"] = NewQOpenVGOffscreenSurfaceFromPointer
	internal.ConstructorTable["quick.QQuickAsyncImageProvider"] = NewQQuickAsyncImageProviderFromPointer
	internal.ConstructorTable["quick.QQuickFramebufferObject"] = NewQQuickFramebufferObjectFromPointer
	internal.ConstructorTable["quick.QQuickImageProvider"] = NewQQuickImageProviderFromPointer
	internal.ConstructorTable["quick.QQuickImageResponse"] = NewQQuickImageResponseFromPointer
	internal.ConstructorTable["quick.QQuickItem"] = NewQQuickItemFromPointer
	internal.ConstructorTable["quick.QQuickItemGrabResult"] = NewQQuickItemGrabResultFromPointer
	internal.ConstructorTable["quick.QQuickPaintedItem"] = NewQQuickPaintedItemFromPointer
	internal.ConstructorTable["quick.QQuickProfilerAdapterFactory"] = NewQQuickProfilerAdapterFactoryFromPointer
	internal.ConstructorTable["quick.QQuickRenderControl"] = NewQQuickRenderControlFromPointer
	internal.ConstructorTable["quick.QQuickTextDocument"] = NewQQuickTextDocumentFromPointer
	internal.ConstructorTable["quick.QQuickTextureFactory"] = NewQQuickTextureFactoryFromPointer
	internal.ConstructorTable["quick.QQuickTransform"] = NewQQuickTransformFromPointer
	internal.ConstructorTable["quick.QQuickView"] = NewQQuickViewFromPointer
	internal.ConstructorTable["quick.QQuickWidget"] = NewQQuickWidgetFromPointer
	internal.ConstructorTable["quick.QQuickWindow"] = NewQQuickWindowFromPointer
	internal.ConstructorTable["quick.QSGAbstractRenderer"] = NewQSGAbstractRendererFromPointer
	internal.ConstructorTable["quick.QSGBasicGeometryNode"] = NewQSGBasicGeometryNodeFromPointer
	internal.ConstructorTable["quick.QSGClipNode"] = NewQSGClipNodeFromPointer
	internal.ConstructorTable["quick.QSGDynamicTexture"] = NewQSGDynamicTextureFromPointer
	internal.ConstructorTable["quick.QSGEngine"] = NewQSGEngineFromPointer
	internal.ConstructorTable["quick.QSGFlatColorMaterial"] = NewQSGFlatColorMaterialFromPointer
	internal.ConstructorTable["quick.QSGGeometry"] = NewQSGGeometryFromPointer
	internal.ConstructorTable["quick.QSGGeometryNode"] = NewQSGGeometryNodeFromPointer
	internal.ConstructorTable["quick.QSGImageNode"] = NewQSGImageNodeFromPointer
	internal.ConstructorTable["quick.QSGMaterial"] = NewQSGMaterialFromPointer
	internal.ConstructorTable["quick.QSGMaterialShader"] = NewQSGMaterialShaderFromPointer
	internal.ConstructorTable["quick.QSGMaterialType"] = NewQSGMaterialTypeFromPointer
	internal.ConstructorTable["quick.QSGNode"] = NewQSGNodeFromPointer
	internal.ConstructorTable["quick.QSGOpacityNode"] = NewQSGOpacityNodeFromPointer
	internal.ConstructorTable["quick.QSGOpaqueTextureMaterial"] = NewQSGOpaqueTextureMaterialFromPointer
	internal.ConstructorTable["quick.QSGRectangleNode"] = NewQSGRectangleNodeFromPointer
	internal.ConstructorTable["quick.QSGRenderNode"] = NewQSGRenderNodeFromPointer
	internal.ConstructorTable["quick.QSGRendererInterface"] = NewQSGRendererInterfaceFromPointer
	internal.ConstructorTable["quick.QSGSimpleRectNode"] = NewQSGSimpleRectNodeFromPointer
	internal.ConstructorTable["quick.QSGSimpleTextureNode"] = NewQSGSimpleTextureNodeFromPointer
	internal.ConstructorTable["quick.QSGTexture"] = NewQSGTextureFromPointer
	internal.ConstructorTable["quick.QSGTextureMaterial"] = NewQSGTextureMaterialFromPointer
	internal.ConstructorTable["quick.QSGTextureProvider"] = NewQSGTextureProviderFromPointer
	internal.ConstructorTable["quick.QSGTransformNode"] = NewQSGTransformNodeFromPointer
	internal.ConstructorTable["quick.QSGVertexColorMaterial"] = NewQSGVertexColorMaterialFromPointer
	internal.ConstructorTable["quick.QTcpServerConnectionFactory"] = NewQTcpServerConnectionFactoryFromPointer
}
