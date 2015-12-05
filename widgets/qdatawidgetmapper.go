package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDataWidgetMapper struct {
	core.QObject
}

type QDataWidgetMapper_ITF interface {
	core.QObject_ITF
	QDataWidgetMapper_PTR() *QDataWidgetMapper
}

func PointerFromQDataWidgetMapper(ptr QDataWidgetMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDataWidgetMapper_PTR().Pointer()
	}
	return nil
}

func NewQDataWidgetMapperFromPointer(ptr unsafe.Pointer) *QDataWidgetMapper {
	var n = new(QDataWidgetMapper)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDataWidgetMapper_") {
		n.SetObjectName("QDataWidgetMapper_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDataWidgetMapper) QDataWidgetMapper_PTR() *QDataWidgetMapper {
	return ptr
}

//QDataWidgetMapper::SubmitPolicy
type QDataWidgetMapper__SubmitPolicy int64

const (
	QDataWidgetMapper__AutoSubmit   = QDataWidgetMapper__SubmitPolicy(0)
	QDataWidgetMapper__ManualSubmit = QDataWidgetMapper__SubmitPolicy(1)
)

func (ptr *QDataWidgetMapper) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDataWidgetMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QDataWidgetMapper) SetOrientation(aOrientation core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetOrientation(ptr.Pointer(), C.int(aOrientation))
	}
}

func (ptr *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper__SubmitPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setSubmitPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetSubmitPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper__SubmitPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::submitPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QDataWidgetMapper__SubmitPolicy(C.QDataWidgetMapper_SubmitPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQDataWidgetMapper(parent core.QObject_ITF) *QDataWidgetMapper {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::QDataWidgetMapper")
		}
	}()

	return NewQDataWidgetMapperFromPointer(C.QDataWidgetMapper_NewQDataWidgetMapper(core.PointerFromQObject(parent)))
}

func (ptr *QDataWidgetMapper) AddMapping(widget QWidget_ITF, section int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::addMapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping(ptr.Pointer(), PointerFromQWidget(widget), C.int(section))
	}
}

func (ptr *QDataWidgetMapper) AddMapping2(widget QWidget_ITF, section int, propertyName core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::addMapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping2(ptr.Pointer(), PointerFromQWidget(widget), C.int(section), core.PointerFromQByteArray(propertyName))
	}
}

func (ptr *QDataWidgetMapper) ClearMapping() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::clearMapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ClearMapping(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ConnectCurrentIndexChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectCurrentIndexChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQDataWidgetMapperCurrentIndexChanged
func callbackQDataWidgetMapperCurrentIndexChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::currentIndexChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::itemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QDataWidgetMapper_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedPropertyName(widget QWidget_ITF) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::mappedPropertyName")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QDataWidgetMapper_MappedPropertyName(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedSection(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::mappedSection")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_MappedSection(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::mappedWidgetAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDataWidgetMapper_MappedWidgetAt(ptr.Pointer(), C.int(section)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) Model() *core.QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::model")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QDataWidgetMapper_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) RemoveMapping(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::removeMapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_RemoveMapping(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDataWidgetMapper) Revert() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::revert")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_Revert(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) RootIndex() *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::rootIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QDataWidgetMapper_RootIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) SetCurrentModelIndex(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setCurrentModelIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setItemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QDataWidgetMapper) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QDataWidgetMapper) SetRootIndex(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::setRootIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) Submit() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::submit")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDataWidgetMapper_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDataWidgetMapper) ToFirst() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::toFirst")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToFirst(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToLast() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::toLast")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToLast(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToNext() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::toNext")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToNext(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToPrevious() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::toPrevious")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToPrevious(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) DestroyQDataWidgetMapper() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataWidgetMapper::~QDataWidgetMapper")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DestroyQDataWidgetMapper(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
