package widgets

//#include "qdatawidgetmapper.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
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
	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDataWidgetMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QDataWidgetMapper) SetOrientation(aOrientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetOrientation(ptr.Pointer(), C.int(aOrientation))
	}
}

func (ptr *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper__SubmitPolicy) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetSubmitPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper__SubmitPolicy {
	if ptr.Pointer() != nil {
		return QDataWidgetMapper__SubmitPolicy(C.QDataWidgetMapper_SubmitPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQDataWidgetMapper(parent core.QObject_ITF) *QDataWidgetMapper {
	return NewQDataWidgetMapperFromPointer(C.QDataWidgetMapper_NewQDataWidgetMapper(core.PointerFromQObject(parent)))
}

func (ptr *QDataWidgetMapper) AddMapping(widget QWidget_ITF, section int) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping(ptr.Pointer(), PointerFromQWidget(widget), C.int(section))
	}
}

func (ptr *QDataWidgetMapper) AddMapping2(widget QWidget_ITF, section int, propertyName core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping2(ptr.Pointer(), PointerFromQWidget(widget), C.int(section), core.PointerFromQByteArray(propertyName))
	}
}

func (ptr *QDataWidgetMapper) ClearMapping() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ClearMapping(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ConnectCurrentIndexChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectCurrentIndexChanged() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQDataWidgetMapperCurrentIndexChanged
func callbackQDataWidgetMapperCurrentIndexChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QDataWidgetMapper_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedPropertyName(widget QWidget_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QDataWidgetMapper_MappedPropertyName(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedSection(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_MappedSection(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDataWidgetMapper_MappedWidgetAt(ptr.Pointer(), C.int(section)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QDataWidgetMapper_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) RemoveMapping(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_RemoveMapping(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDataWidgetMapper) Revert() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_Revert(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) RootIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QDataWidgetMapper_RootIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) SetCurrentModelIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QDataWidgetMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QDataWidgetMapper) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QDataWidgetMapper_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDataWidgetMapper) ToFirst() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToFirst(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToLast() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToLast(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToNext() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToNext(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToPrevious() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToPrevious(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) DestroyQDataWidgetMapper() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DestroyQDataWidgetMapper(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
