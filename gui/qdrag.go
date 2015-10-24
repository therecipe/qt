package gui

//#include "qdrag.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDrag struct {
	core.QObject
}

type QDragITF interface {
	core.QObjectITF
	QDragPTR() *QDrag
}

func PointerFromQDrag(ptr QDragITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragPTR().Pointer()
	}
	return nil
}

func QDragFromPointer(ptr unsafe.Pointer) *QDrag {
	var n = new(QDrag)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDrag_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDrag) QDragPTR() *QDrag {
	return ptr
}

func NewQDrag(dragSource core.QObjectITF) *QDrag {
	return QDragFromPointer(unsafe.Pointer(C.QDrag_NewQDrag(C.QtObjectPtr(core.PointerFromQObject(dragSource)))))
}

func (ptr *QDrag) ConnectActionChanged(f func(action core.Qt__DropAction)) {
	if ptr.Pointer() != nil {
		C.QDrag_ConnectActionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actionChanged", f)
	}
}

func (ptr *QDrag) DisconnectActionChanged() {
	if ptr.Pointer() != nil {
		C.QDrag_DisconnectActionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actionChanged")
	}
}

//export callbackQDragActionChanged
func callbackQDragActionChanged(ptrName *C.char, action C.int) {
	qt.GetSignal(C.GoString(ptrName), "actionChanged").(func(core.Qt__DropAction))(core.Qt__DropAction(action))
}

func (ptr *QDrag) Exec(supportedActions core.Qt__DropAction) core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec(C.QtObjectPtr(ptr.Pointer()), C.int(supportedActions)))
	}
	return 0
}

func (ptr *QDrag) Exec2(supportedActions core.Qt__DropAction, defaultDropAction core.Qt__DropAction) core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec2(C.QtObjectPtr(ptr.Pointer()), C.int(supportedActions), C.int(defaultDropAction)))
	}
	return 0
}

func (ptr *QDrag) MimeData() *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.QMimeDataFromPointer(unsafe.Pointer(C.QDrag_MimeData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDrag) SetDragCursor(cursor QPixmapITF, action core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QDrag_SetDragCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(cursor)), C.int(action))
	}
}

func (ptr *QDrag) SetHotSpot(hotspot core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QDrag_SetHotSpot(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(hotspot)))
	}
}

func (ptr *QDrag) SetMimeData(data core.QMimeDataITF) {
	if ptr.Pointer() != nil {
		C.QDrag_SetMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMimeData(data)))
	}
}

func (ptr *QDrag) SetPixmap(pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QDrag_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QDrag) Source() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QDrag_Source(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDrag) SupportedActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_SupportedActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDrag) Target() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QDrag_Target(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDrag) ConnectTargetChanged(f func(newTarget core.QObjectITF)) {
	if ptr.Pointer() != nil {
		C.QDrag_ConnectTargetChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetChanged", f)
	}
}

func (ptr *QDrag) DisconnectTargetChanged() {
	if ptr.Pointer() != nil {
		C.QDrag_DisconnectTargetChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetChanged")
	}
}

//export callbackQDragTargetChanged
func callbackQDragTargetChanged(ptrName *C.char, newTarget unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "targetChanged").(func(*core.QObject))(core.QObjectFromPointer(newTarget))
}

func (ptr *QDrag) DestroyQDrag() {
	if ptr.Pointer() != nil {
		C.QDrag_DestroyQDrag(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
