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

type QDataWidgetMapperITF interface {
	core.QObjectITF
	QDataWidgetMapperPTR() *QDataWidgetMapper
}

func PointerFromQDataWidgetMapper(ptr QDataWidgetMapperITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDataWidgetMapperPTR().Pointer()
	}
	return nil
}

func QDataWidgetMapperFromPointer(ptr unsafe.Pointer) *QDataWidgetMapper {
	var n = new(QDataWidgetMapper)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDataWidgetMapper_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDataWidgetMapper) QDataWidgetMapperPTR() *QDataWidgetMapper {
	return ptr
}

//QDataWidgetMapper::SubmitPolicy
type QDataWidgetMapper__SubmitPolicy int

var (
	QDataWidgetMapper__AutoSubmit   = QDataWidgetMapper__SubmitPolicy(0)
	QDataWidgetMapper__ManualSubmit = QDataWidgetMapper__SubmitPolicy(1)
)

func (ptr *QDataWidgetMapper) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataWidgetMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDataWidgetMapper_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataWidgetMapper) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QDataWidgetMapper) SetOrientation(aOrientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(aOrientation))
	}
}

func (ptr *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper__SubmitPolicy) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetSubmitPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper__SubmitPolicy {
	if ptr.Pointer() != nil {
		return QDataWidgetMapper__SubmitPolicy(C.QDataWidgetMapper_SubmitPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQDataWidgetMapper(parent core.QObjectITF) *QDataWidgetMapper {
	return QDataWidgetMapperFromPointer(unsafe.Pointer(C.QDataWidgetMapper_NewQDataWidgetMapper(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDataWidgetMapper) AddMapping(widget QWidgetITF, section int) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(section))
	}
}

func (ptr *QDataWidgetMapper) AddMapping2(widget QWidgetITF, section int, propertyName core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(section), C.QtObjectPtr(core.PointerFromQByteArray(propertyName)))
	}
}

func (ptr *QDataWidgetMapper) ClearMapping() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ClearMapping(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) ConnectCurrentIndexChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ConnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectCurrentIndexChanged() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DisconnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQDataWidgetMapperCurrentIndexChanged
func callbackQDataWidgetMapperCurrentIndexChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QDataWidgetMapper_ItemDelegate(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedSection(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_MappedSection(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QDataWidgetMapper_MappedWidgetAt(C.QtObjectPtr(ptr.Pointer()), C.int(section))))
	}
	return nil
}

func (ptr *QDataWidgetMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.QAbstractItemModelFromPointer(unsafe.Pointer(C.QDataWidgetMapper_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDataWidgetMapper) RemoveMapping(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_RemoveMapping(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QDataWidgetMapper) Revert() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_Revert(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) RootIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QDataWidgetMapper_RootIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDataWidgetMapper) SetCurrentModelIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentModelIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QDataWidgetMapper) SetItemDelegate(delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetItemDelegate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QDataWidgetMapper) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QDataWidgetMapper) SetRootIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetRootIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QDataWidgetMapper) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QDataWidgetMapper_Submit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDataWidgetMapper) ToFirst() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToFirst(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) ToLast() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToLast(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) ToNext() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToNext(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) ToPrevious() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToPrevious(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataWidgetMapper) DestroyQDataWidgetMapper() {
	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DestroyQDataWidgetMapper(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
