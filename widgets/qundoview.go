package widgets

//#include "qundoview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QUndoView struct {
	QListView
}

type QUndoViewITF interface {
	QListViewITF
	QUndoViewPTR() *QUndoView
}

func PointerFromQUndoView(ptr QUndoViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoViewPTR().Pointer()
	}
	return nil
}

func QUndoViewFromPointer(ptr unsafe.Pointer) *QUndoView {
	var n = new(QUndoView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QUndoView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUndoView) QUndoViewPTR() *QUndoView {
	return ptr
}

func (ptr *QUndoView) EmptyLabel() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoView_EmptyLabel(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QUndoView) SetCleanIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QUndoView_SetCleanIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QUndoView) SetEmptyLabel(label string) {
	if ptr.Pointer() != nil {
		C.QUndoView_SetEmptyLabel(C.QtObjectPtr(ptr.Pointer()), C.CString(label))
	}
}

func NewQUndoView3(group QUndoGroupITF, parent QWidgetITF) *QUndoView {
	return QUndoViewFromPointer(unsafe.Pointer(C.QUndoView_NewQUndoView3(C.QtObjectPtr(PointerFromQUndoGroup(group)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQUndoView2(stack QUndoStackITF, parent QWidgetITF) *QUndoView {
	return QUndoViewFromPointer(unsafe.Pointer(C.QUndoView_NewQUndoView2(C.QtObjectPtr(PointerFromQUndoStack(stack)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQUndoView(parent QWidgetITF) *QUndoView {
	return QUndoViewFromPointer(unsafe.Pointer(C.QUndoView_NewQUndoView(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QUndoView) Group() *QUndoGroup {
	if ptr.Pointer() != nil {
		return QUndoGroupFromPointer(unsafe.Pointer(C.QUndoView_Group(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QUndoView) SetGroup(group QUndoGroupITF) {
	if ptr.Pointer() != nil {
		C.QUndoView_SetGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQUndoGroup(group)))
	}
}

func (ptr *QUndoView) SetStack(stack QUndoStackITF) {
	if ptr.Pointer() != nil {
		C.QUndoView_SetStack(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQUndoStack(stack)))
	}
}

func (ptr *QUndoView) Stack() *QUndoStack {
	if ptr.Pointer() != nil {
		return QUndoStackFromPointer(unsafe.Pointer(C.QUndoView_Stack(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QUndoView) DestroyQUndoView() {
	if ptr.Pointer() != nil {
		C.QUndoView_DestroyQUndoView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
