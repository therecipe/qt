package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) *QAbstractScrollArea {
	var n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractScrollArea_") {
		n.SetObjectName("QAbstractScrollArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::horizontalScrollBarPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setHorizontalScrollBarPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setSizeAdjustPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setVerticalScrollBarPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::sizeAdjustPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractScrollArea__SizeAdjustPolicy(C.QAbstractScrollArea_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::verticalScrollBarPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::QAbstractScrollArea")
		}
	}()

	return NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QAbstractScrollArea) AddScrollBarWidget(widget QWidget_ITF, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::addScrollBarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_AddScrollBarWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(alignment))
	}
}

func (ptr *QAbstractScrollArea) CornerWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::cornerWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_CornerWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::horizontalScrollBar")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_HorizontalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setCornerWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBar_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setHorizontalScrollBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBar_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setVerticalScrollBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetViewport(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setViewport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetViewport(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetupViewport(viewport QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::setupViewport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::verticalScrollBar")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_VerticalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) Viewport() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::viewport")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_Viewport(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractScrollArea::~QAbstractScrollArea")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
