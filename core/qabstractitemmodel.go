package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QAbstractItemModel struct {
	QObject
}

type QAbstractItemModel_ITF interface {
	QObject_ITF
	QAbstractItemModel_PTR() *QAbstractItemModel
}

func PointerFromQAbstractItemModel(ptr QAbstractItemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemModelFromPointer(ptr unsafe.Pointer) *QAbstractItemModel {
	var n = new(QAbstractItemModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractItemModel_") {
		n.SetObjectName("QAbstractItemModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractItemModel) QAbstractItemModel_PTR() *QAbstractItemModel {
	return ptr
}

//QAbstractItemModel::LayoutChangeHint
type QAbstractItemModel__LayoutChangeHint int64

const (
	QAbstractItemModel__NoLayoutChangeHint = QAbstractItemModel__LayoutChangeHint(0)
	QAbstractItemModel__VerticalSortHint   = QAbstractItemModel__LayoutChangeHint(1)
	QAbstractItemModel__HorizontalSortHint = QAbstractItemModel__LayoutChangeHint(2)
)

func (ptr *QAbstractItemModel) Sibling(row int, column int, index QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractItemModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractItemModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemModel) Buddy(index QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractItemModel::buddy")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractItemModel_Buddy(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemModel) CanDropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_CanDropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) CanFetchMore(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_CanFetchMore(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) ColumnCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QAbstractItemModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QAbstractItemModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeInserted(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeInserted() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeInserted")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeInserted
func callbackQAbstractItemModelColumnsAboutToBeInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsAboutToBeInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeInserted")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeMoved(f func(sourceParent *QModelIndex, sourceStart int, sourceEnd int, destinationParent *QModelIndex, destinationColumn int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsAboutToBeMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeMoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsAboutToBeMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeMoved")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeMoved
func callbackQAbstractItemModelColumnsAboutToBeMoved(ptrName *C.char, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsAboutToBeMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeMoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int, *QModelIndex, int))(NewQModelIndexFromPointer(sourceParent), int(sourceStart), int(sourceEnd), NewQModelIndexFromPointer(destinationParent), int(destinationColumn))
	}

}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeRemoved(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeRemoved")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeRemoved
func callbackQAbstractItemModelColumnsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsAboutToBeRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeRemoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) ConnectColumnsInserted(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsInserted() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsInserted")
	}
}

//export callbackQAbstractItemModelColumnsInserted
func callbackQAbstractItemModelColumnsInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsInserted")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) ConnectColumnsMoved(f func(parent *QModelIndex, start int, end int, destination *QModelIndex, column int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsMoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsMoved")
	}
}

//export callbackQAbstractItemModelColumnsMoved
func callbackQAbstractItemModelColumnsMoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsMoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int, *QModelIndex, int))(NewQModelIndexFromPointer(parent), int(start), int(end), NewQModelIndexFromPointer(destination), int(column))
	}

}

func (ptr *QAbstractItemModel) ConnectColumnsRemoved(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::columnsRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "columnsRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsRemoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::columnsRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "columnsRemoved")
	}
}

//export callbackQAbstractItemModelColumnsRemoved
func callbackQAbstractItemModelColumnsRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::columnsRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "columnsRemoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) Data(index QModelIndex_ITF, role int) *QVariant {
	defer qt.Recovering("QAbstractItemModel::data")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractItemModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractItemModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QAbstractItemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQAbstractItemModelFetchMore
func callbackQAbstractItemModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemModel::fetchMore")

	var signal = qt.GetSignal(C.GoString(ptrName), "fetchMore")
	if signal != nil {
		defer signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
		return true
	}
	return false

}

func (ptr *QAbstractItemModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QAbstractItemModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractItemModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractItemModel) HasChildren(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) HasIndex(row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::hasIndex")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_HasIndex(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer qt.Recovering("QAbstractItemModel::headerData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractItemModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractItemModel) ConnectHeaderDataChanged(f func(orientation Qt__Orientation, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::headerDataChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectHeaderDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "headerDataChanged", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectHeaderDataChanged() {
	defer qt.Recovering("disconnect QAbstractItemModel::headerDataChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectHeaderDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "headerDataChanged")
	}
}

//export callbackQAbstractItemModelHeaderDataChanged
func callbackQAbstractItemModelHeaderDataChanged(ptrName *C.char, orientation C.int, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::headerDataChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "headerDataChanged")
	if signal != nil {
		signal.(func(Qt__Orientation, int, int))(Qt__Orientation(orientation), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractItemModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractItemModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QAbstractItemModel) InsertColumn(column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::insertColumn")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertColumn(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRow(row int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::insertRow")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertRow(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MimeTypes() []string {
	defer qt.Recovering("QAbstractItemModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAbstractItemModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAbstractItemModel) ConnectModelAboutToBeReset(f func()) {
	defer qt.Recovering("connect QAbstractItemModel::modelAboutToBeReset")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectModelAboutToBeReset(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modelAboutToBeReset", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectModelAboutToBeReset() {
	defer qt.Recovering("disconnect QAbstractItemModel::modelAboutToBeReset")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectModelAboutToBeReset(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modelAboutToBeReset")
	}
}

//export callbackQAbstractItemModelModelAboutToBeReset
func callbackQAbstractItemModelModelAboutToBeReset(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractItemModel::modelAboutToBeReset")

	var signal = qt.GetSignal(C.GoString(ptrName), "modelAboutToBeReset")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractItemModel) ConnectModelReset(f func()) {
	defer qt.Recovering("connect QAbstractItemModel::modelReset")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectModelReset(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modelReset", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectModelReset() {
	defer qt.Recovering("disconnect QAbstractItemModel::modelReset")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectModelReset(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modelReset")
	}
}

//export callbackQAbstractItemModelModelReset
func callbackQAbstractItemModelModelReset(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractItemModel::modelReset")

	var signal = qt.GetSignal(C.GoString(ptrName), "modelReset")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractItemModel) MoveColumn(sourceParent QModelIndex_ITF, sourceColumn int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QAbstractItemModel::moveColumn")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveColumn(ptr.Pointer(), PointerFromQModelIndex(sourceParent), C.int(sourceColumn), PointerFromQModelIndex(destinationParent), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveColumns(sourceParent QModelIndex_ITF, sourceColumn int, count int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QAbstractItemModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveColumns(ptr.Pointer(), PointerFromQModelIndex(sourceParent), C.int(sourceColumn), C.int(count), PointerFromQModelIndex(destinationParent), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveRow(sourceParent QModelIndex_ITF, sourceRow int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QAbstractItemModel::moveRow")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveRow(ptr.Pointer(), PointerFromQModelIndex(sourceParent), C.int(sourceRow), PointerFromQModelIndex(destinationParent), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveRows(sourceParent QModelIndex_ITF, sourceRow int, count int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QAbstractItemModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveRows(ptr.Pointer(), PointerFromQModelIndex(sourceParent), C.int(sourceRow), C.int(count), PointerFromQModelIndex(destinationParent), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) Parent(index QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractItemModel::parent")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractItemModel_Parent(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemModel) RemoveColumn(column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::removeColumn")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveColumn(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRow(row int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::removeRow")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveRow(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QAbstractItemModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QAbstractItemModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQAbstractItemModelRevert
func callbackQAbstractItemModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractItemModel::revert")

	var signal = qt.GetSignal(C.GoString(ptrName), "revert")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractItemModel) RowCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QAbstractItemModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QAbstractItemModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeInserted(f func(parent *QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeInserted() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeInserted")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeInserted
func callbackQAbstractItemModelRowsAboutToBeInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsAboutToBeInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeInserted")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(start), int(end))
	}

}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeMoved(f func(sourceParent *QModelIndex, sourceStart int, sourceEnd int, destinationParent *QModelIndex, destinationRow int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsAboutToBeMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeMoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsAboutToBeMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeMoved")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeMoved
func callbackQAbstractItemModelRowsAboutToBeMoved(ptrName *C.char, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsAboutToBeMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeMoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int, *QModelIndex, int))(NewQModelIndexFromPointer(sourceParent), int(sourceStart), int(sourceEnd), NewQModelIndexFromPointer(destinationParent), int(destinationRow))
	}

}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeRemoved(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeRemoved
func callbackQAbstractItemModelRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsAboutToBeRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) ConnectRowsInserted(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsInserted")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQAbstractItemModelRowsInserted
func callbackQAbstractItemModelRowsInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsInserted")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) ConnectRowsMoved(f func(parent *QModelIndex, start int, end int, destination *QModelIndex, row int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsMoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsMoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsMoved")
	}
}

//export callbackQAbstractItemModelRowsMoved
func callbackQAbstractItemModelRowsMoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsMoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int, *QModelIndex, int))(NewQModelIndexFromPointer(parent), int(start), int(end), NewQModelIndexFromPointer(destination), int(row))
	}

}

func (ptr *QAbstractItemModel) ConnectRowsRemoved(f func(parent *QModelIndex, first int, last int)) {
	defer qt.Recovering("connect QAbstractItemModel::rowsRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rowsRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsRemoved() {
	defer qt.Recovering("disconnect QAbstractItemModel::rowsRemoved")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rowsRemoved")
	}
}

//export callbackQAbstractItemModelRowsRemoved
func callbackQAbstractItemModelRowsRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	defer qt.Recovering("callback QAbstractItemModel::rowsRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsRemoved")
	if signal != nil {
		signal.(func(*QModelIndex, int, int))(NewQModelIndexFromPointer(parent), int(first), int(last))
	}

}

func (ptr *QAbstractItemModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QAbstractItemModel::setData")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) SetHeaderData(section int, orientation Qt__Orientation, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QAbstractItemModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QAbstractItemModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectSort() {
	defer qt.Recovering("disconnect QAbstractItemModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQAbstractItemModelSort
func callbackQAbstractItemModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QAbstractItemModel::sort")

	var signal = qt.GetSignal(C.GoString(ptrName), "sort")
	if signal != nil {
		defer signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QAbstractItemModel) Span(index QModelIndex_ITF) *QSize {
	defer qt.Recovering("QAbstractItemModel::span")

	if ptr.Pointer() != nil {
		return NewQSizeFromPointer(C.QAbstractItemModel_Span(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemModel) Submit() bool {
	defer qt.Recovering("QAbstractItemModel::submit")

	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) SupportedDragActions() Qt__DropAction {
	defer qt.Recovering("QAbstractItemModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractItemModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemModel) SupportedDropActions() Qt__DropAction {
	defer qt.Recovering("QAbstractItemModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractItemModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemModel) DestroyQAbstractItemModel() {
	defer qt.Recovering("QAbstractItemModel::~QAbstractItemModel")

	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DestroyQAbstractItemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
