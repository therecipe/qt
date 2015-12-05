package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBoxLayout struct {
	QLayout
}

type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func PointerFromQBoxLayout(ptr QBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQBoxLayoutFromPointer(ptr unsafe.Pointer) *QBoxLayout {
	var n = new(QBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBoxLayout_") {
		n.SetObjectName("QBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout {
	return ptr
}

//QBoxLayout::Direction
type QBoxLayout__Direction int64

const (
	QBoxLayout__LeftToRight = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop = QBoxLayout__Direction(3)
	QBoxLayout__Down        = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func (ptr *QBoxLayout) Direction() QBoxLayout__Direction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::direction")
		}
	}()

	if ptr.Pointer() != nil {
		return QBoxLayout__Direction(C.QBoxLayout_Direction(ptr.Pointer()))
	}
	return 0
}

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidget_ITF) *QBoxLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::QBoxLayout")
		}
	}()

	return NewQBoxLayoutFromPointer(C.QBoxLayout_NewQBoxLayout(C.int(dir), PointerFromQWidget(parent)))
}

func (ptr *QBoxLayout) AddItem(item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddLayout(layout QLayout_ITF, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddSpacerItem(spacerItem QSpacerItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addSpacerItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacerItem(ptr.Pointer(), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) AddSpacing(size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacing(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddStretch(stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addStretch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStretch(ptr.Pointer(), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddStrut(size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addStrut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStrut(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddWidget(widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) ExpandingDirections() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::expandingDirections")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QBoxLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) HasHeightForWidth() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::hasHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBoxLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBoxLayout) HeightForWidth(w int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) InsertItem(index int, item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertItem(ptr.Pointer(), C.int(index), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) InsertLayout(index int, layout QLayout_ITF, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertLayout(ptr.Pointer(), C.int(index), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertSpacerItem(index int, spacerItem QSpacerItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertSpacerItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacerItem(ptr.Pointer(), C.int(index), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) InsertSpacing(index int, size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacing(ptr.Pointer(), C.int(index), C.int(size))
	}
}

func (ptr *QBoxLayout) InsertStretch(index int, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertStretch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertWidget(index int, widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) MinimumHeightForWidth(w int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::minimumHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) SetDirection(direction QBoxLayout__Direction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QBoxLayout) SetGeometry(r core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QBoxLayout) SetSpacing(spacing int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QBoxLayout) SetStretch(index int, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setStretch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) SetStretchFactor2(layout QLayout_ITF, stretch int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor2(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) SetStretchFactor(widget QWidget_ITF, stretch int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::setStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) Spacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::spacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) Stretch(index int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::stretch")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Stretch(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::takeAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBoxLayout::~QBoxLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBoxLayout_DestroyQBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
