package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QUndoView struct {
	QListView
}

type QUndoView_ITF interface {
	QListView_ITF
	QUndoView_PTR() *QUndoView
}

func PointerFromQUndoView(ptr QUndoView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoView_PTR().Pointer()
	}
	return nil
}

func NewQUndoViewFromPointer(ptr unsafe.Pointer) *QUndoView {
	var n = new(QUndoView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUndoView_") {
		n.SetObjectName("QUndoView_" + qt.Identifier())
	}
	return n
}

func (ptr *QUndoView) QUndoView_PTR() *QUndoView {
	return ptr
}

func (ptr *QUndoView) EmptyLabel() string {
	defer qt.Recovering("QUndoView::emptyLabel")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoView_EmptyLabel(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoView) SetCleanIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QUndoView::setCleanIcon")

	if ptr.Pointer() != nil {
		C.QUndoView_SetCleanIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QUndoView) SetEmptyLabel(label string) {
	defer qt.Recovering("QUndoView::setEmptyLabel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetEmptyLabel(ptr.Pointer(), C.CString(label))
	}
}

func NewQUndoView3(group QUndoGroup_ITF, parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView3(PointerFromQUndoGroup(group), PointerFromQWidget(parent)))
}

func NewQUndoView2(stack QUndoStack_ITF, parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView2(PointerFromQUndoStack(stack), PointerFromQWidget(parent)))
}

func NewQUndoView(parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView(PointerFromQWidget(parent)))
}

func (ptr *QUndoView) Group() *QUndoGroup {
	defer qt.Recovering("QUndoView::group")

	if ptr.Pointer() != nil {
		return NewQUndoGroupFromPointer(C.QUndoView_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoView) SetGroup(group QUndoGroup_ITF) {
	defer qt.Recovering("QUndoView::setGroup")

	if ptr.Pointer() != nil {
		C.QUndoView_SetGroup(ptr.Pointer(), PointerFromQUndoGroup(group))
	}
}

func (ptr *QUndoView) SetStack(stack QUndoStack_ITF) {
	defer qt.Recovering("QUndoView::setStack")

	if ptr.Pointer() != nil {
		C.QUndoView_SetStack(ptr.Pointer(), PointerFromQUndoStack(stack))
	}
}

func (ptr *QUndoView) Stack() *QUndoStack {
	defer qt.Recovering("QUndoView::stack")

	if ptr.Pointer() != nil {
		return NewQUndoStackFromPointer(C.QUndoView_Stack(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoView) DestroyQUndoView() {
	defer qt.Recovering("QUndoView::~QUndoView")

	if ptr.Pointer() != nil {
		C.QUndoView_DestroyQUndoView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
