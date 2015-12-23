package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QDataWidgetMapper_") {
		n.SetObjectName("QDataWidgetMapper_" + qt.Identifier())
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
	defer qt.Recovering("QDataWidgetMapper::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QDataWidgetMapper::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDataWidgetMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataWidgetMapper) ConnectSetCurrentIndex(f func(index int)) {
	defer qt.Recovering("connect QDataWidgetMapper::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setCurrentIndex", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectSetCurrentIndex() {
	defer qt.Recovering("disconnect QDataWidgetMapper::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setCurrentIndex")
	}
}

//export callbackQDataWidgetMapperSetCurrentIndex
func callbackQDataWidgetMapperSetCurrentIndex(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QDataWidgetMapper::setCurrentIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setCurrentIndex"); signal != nil {
		signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QDataWidgetMapper) SetOrientation(aOrientation core.Qt__Orientation) {
	defer qt.Recovering("QDataWidgetMapper::setOrientation")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetOrientation(ptr.Pointer(), C.int(aOrientation))
	}
}

func (ptr *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper__SubmitPolicy) {
	defer qt.Recovering("QDataWidgetMapper::setSubmitPolicy")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetSubmitPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper__SubmitPolicy {
	defer qt.Recovering("QDataWidgetMapper::submitPolicy")

	if ptr.Pointer() != nil {
		return QDataWidgetMapper__SubmitPolicy(C.QDataWidgetMapper_SubmitPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQDataWidgetMapper(parent core.QObject_ITF) *QDataWidgetMapper {
	defer qt.Recovering("QDataWidgetMapper::QDataWidgetMapper")

	return NewQDataWidgetMapperFromPointer(C.QDataWidgetMapper_NewQDataWidgetMapper(core.PointerFromQObject(parent)))
}

func (ptr *QDataWidgetMapper) AddMapping(widget QWidget_ITF, section int) {
	defer qt.Recovering("QDataWidgetMapper::addMapping")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping(ptr.Pointer(), PointerFromQWidget(widget), C.int(section))
	}
}

func (ptr *QDataWidgetMapper) AddMapping2(widget QWidget_ITF, section int, propertyName core.QByteArray_ITF) {
	defer qt.Recovering("QDataWidgetMapper::addMapping")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_AddMapping2(ptr.Pointer(), PointerFromQWidget(widget), C.int(section), core.PointerFromQByteArray(propertyName))
	}
}

func (ptr *QDataWidgetMapper) ClearMapping() {
	defer qt.Recovering("QDataWidgetMapper::clearMapping")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ClearMapping(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ConnectCurrentIndexChanged(f func(index int)) {
	defer qt.Recovering("connect QDataWidgetMapper::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectCurrentIndexChanged() {
	defer qt.Recovering("disconnect QDataWidgetMapper::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQDataWidgetMapperCurrentIndexChanged
func callbackQDataWidgetMapperCurrentIndexChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QDataWidgetMapper::currentIndexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIndexChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	defer qt.Recovering("QDataWidgetMapper::itemDelegate")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QDataWidgetMapper_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedPropertyName(widget QWidget_ITF) *core.QByteArray {
	defer qt.Recovering("QDataWidgetMapper::mappedPropertyName")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QDataWidgetMapper_MappedPropertyName(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) MappedSection(widget QWidget_ITF) int {
	defer qt.Recovering("QDataWidgetMapper::mappedSection")

	if ptr.Pointer() != nil {
		return int(C.QDataWidgetMapper_MappedSection(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	defer qt.Recovering("QDataWidgetMapper::mappedWidgetAt")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDataWidgetMapper_MappedWidgetAt(ptr.Pointer(), C.int(section)))
	}
	return nil
}

func (ptr *QDataWidgetMapper) Model() *core.QAbstractItemModel {
	defer qt.Recovering("QDataWidgetMapper::model")

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QDataWidgetMapper_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) RemoveMapping(widget QWidget_ITF) {
	defer qt.Recovering("QDataWidgetMapper::removeMapping")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_RemoveMapping(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDataWidgetMapper) Revert() {
	defer qt.Recovering("QDataWidgetMapper::revert")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_Revert(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) RootIndex() *core.QModelIndex {
	defer qt.Recovering("QDataWidgetMapper::rootIndex")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QDataWidgetMapper_RootIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataWidgetMapper) SetCurrentModelIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QDataWidgetMapper::setCurrentModelIndex")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetCurrentModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QDataWidgetMapper::setItemDelegate")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QDataWidgetMapper) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QDataWidgetMapper::setModel")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QDataWidgetMapper) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QDataWidgetMapper::setRootIndex")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QDataWidgetMapper) Submit() bool {
	defer qt.Recovering("QDataWidgetMapper::submit")

	if ptr.Pointer() != nil {
		return C.QDataWidgetMapper_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDataWidgetMapper) ToFirst() {
	defer qt.Recovering("QDataWidgetMapper::toFirst")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToFirst(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToLast() {
	defer qt.Recovering("QDataWidgetMapper::toLast")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToLast(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToNext() {
	defer qt.Recovering("QDataWidgetMapper::toNext")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToNext(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) ToPrevious() {
	defer qt.Recovering("QDataWidgetMapper::toPrevious")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_ToPrevious(ptr.Pointer())
	}
}

func (ptr *QDataWidgetMapper) DestroyQDataWidgetMapper() {
	defer qt.Recovering("QDataWidgetMapper::~QDataWidgetMapper")

	if ptr.Pointer() != nil {
		C.QDataWidgetMapper_DestroyQDataWidgetMapper(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDataWidgetMapper) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDataWidgetMapper::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDataWidgetMapper::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDataWidgetMapperTimerEvent
func callbackQDataWidgetMapperTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDataWidgetMapper::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDataWidgetMapper) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDataWidgetMapper::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDataWidgetMapper::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDataWidgetMapperChildEvent
func callbackQDataWidgetMapperChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDataWidgetMapper::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDataWidgetMapper) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDataWidgetMapper::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDataWidgetMapper) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDataWidgetMapper::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDataWidgetMapperCustomEvent
func callbackQDataWidgetMapperCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDataWidgetMapper::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
