package core

//#include "qabstractitemmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QAbstractItemModel struct {
	QObject
}

type QAbstractItemModelITF interface {
	QObjectITF
	QAbstractItemModelPTR() *QAbstractItemModel
}

func PointerFromQAbstractItemModel(ptr QAbstractItemModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModelPTR().Pointer()
	}
	return nil
}

func QAbstractItemModelFromPointer(ptr unsafe.Pointer) *QAbstractItemModel {
	var n = new(QAbstractItemModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractItemModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractItemModel) QAbstractItemModelPTR() *QAbstractItemModel {
	return ptr
}

//QAbstractItemModel::LayoutChangeHint
type QAbstractItemModel__LayoutChangeHint int

var (
	QAbstractItemModel__NoLayoutChangeHint = QAbstractItemModel__LayoutChangeHint(0)
	QAbstractItemModel__VerticalSortHint   = QAbstractItemModel__LayoutChangeHint(1)
	QAbstractItemModel__HorizontalSortHint = QAbstractItemModel__LayoutChangeHint(2)
)

func (ptr *QAbstractItemModel) Sibling(row int, column int, index QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemModel) Buddy(index QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemModel_Buddy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemModel) CanDropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_CanDropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) CanFetchMore(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_CanFetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) ColumnCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeInserted(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeInserted() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeInserted")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeInserted
func callbackQAbstractItemModelColumnsAboutToBeInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeInserted").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeMoved(f func(sourceParent QModelIndexITF, sourceStart int, sourceEnd int, destinationParent QModelIndexITF, destinationColumn int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeMoved")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeMoved
func callbackQAbstractItemModelColumnsAboutToBeMoved(ptrName *C.char, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeMoved").(func(*QModelIndex, int, int, *QModelIndex, int))(QModelIndexFromPointer(sourceParent), int(sourceStart), int(sourceEnd), QModelIndexFromPointer(destinationParent), int(destinationColumn))
}

func (ptr *QAbstractItemModel) ConnectColumnsAboutToBeRemoved(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsAboutToBeRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsAboutToBeRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsAboutToBeRemoved")
	}
}

//export callbackQAbstractItemModelColumnsAboutToBeRemoved
func callbackQAbstractItemModelColumnsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsAboutToBeRemoved").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) ConnectColumnsInserted(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsInserted() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsInserted")
	}
}

//export callbackQAbstractItemModelColumnsInserted
func callbackQAbstractItemModelColumnsInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsInserted").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) ConnectColumnsMoved(f func(parent QModelIndexITF, start int, end int, destination QModelIndexITF, column int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsMoved")
	}
}

//export callbackQAbstractItemModelColumnsMoved
func callbackQAbstractItemModelColumnsMoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsMoved").(func(*QModelIndex, int, int, *QModelIndex, int))(QModelIndexFromPointer(parent), int(start), int(end), QModelIndexFromPointer(destination), int(column))
}

func (ptr *QAbstractItemModel) ConnectColumnsRemoved(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectColumnsRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "columnsRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnsRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectColumnsRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "columnsRemoved")
	}
}

//export callbackQAbstractItemModelColumnsRemoved
func callbackQAbstractItemModelColumnsRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "columnsRemoved").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) Data(index QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractItemModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QAbstractItemModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) FetchMore(parent QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_FetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent)))
	}
}

func (ptr *QAbstractItemModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractItemModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QAbstractItemModel) HasChildren(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_HasChildren(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) HasIndex(row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_HasIndex(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) HeaderData(section int, orientation Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractItemModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QAbstractItemModel) ConnectHeaderDataChanged(f func(orientation Qt__Orientation, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectHeaderDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "headerDataChanged", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectHeaderDataChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectHeaderDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "headerDataChanged")
	}
}

//export callbackQAbstractItemModelHeaderDataChanged
func callbackQAbstractItemModelHeaderDataChanged(ptrName *C.char, orientation C.int, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "headerDataChanged").(func(Qt__Orientation, int, int))(Qt__Orientation(orientation), int(first), int(last))
}

func (ptr *QAbstractItemModel) Index(row int, column int, parent QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QAbstractItemModel) InsertColumn(column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRow(row int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertRow(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAbstractItemModel_MimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAbstractItemModel) ConnectModelAboutToBeReset(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectModelAboutToBeReset(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modelAboutToBeReset", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectModelAboutToBeReset() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectModelAboutToBeReset(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modelAboutToBeReset")
	}
}

//export callbackQAbstractItemModelModelAboutToBeReset
func callbackQAbstractItemModelModelAboutToBeReset(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "modelAboutToBeReset").(func())()
}

func (ptr *QAbstractItemModel) ConnectModelReset(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectModelReset(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modelReset", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectModelReset() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectModelReset(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modelReset")
	}
}

//export callbackQAbstractItemModelModelReset
func callbackQAbstractItemModelModelReset(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "modelReset").(func())()
}

func (ptr *QAbstractItemModel) MoveColumn(sourceParent QModelIndexITF, sourceColumn int, destinationParent QModelIndexITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveColumn(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceParent)), C.int(sourceColumn), C.QtObjectPtr(PointerFromQModelIndex(destinationParent)), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveColumns(sourceParent QModelIndexITF, sourceColumn int, count int, destinationParent QModelIndexITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveColumns(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceParent)), C.int(sourceColumn), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(destinationParent)), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveRow(sourceParent QModelIndexITF, sourceRow int, destinationParent QModelIndexITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveRow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceParent)), C.int(sourceRow), C.QtObjectPtr(PointerFromQModelIndex(destinationParent)), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) MoveRows(sourceParent QModelIndexITF, sourceRow int, count int, destinationParent QModelIndexITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_MoveRows(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceParent)), C.int(sourceRow), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(destinationParent)), C.int(destinationChild)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) Parent(index QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemModel) RemoveColumn(column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRow(row int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveRow(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) Revert() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_Revert(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemModel) RowCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeInserted(f func(parent QModelIndexITF, start int, end int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeInserted() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeInserted")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeInserted
func callbackQAbstractItemModelRowsAboutToBeInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeInserted").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(start), int(end))
}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeMoved(f func(sourceParent QModelIndexITF, sourceStart int, sourceEnd int, destinationParent QModelIndexITF, destinationRow int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeMoved")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeMoved
func callbackQAbstractItemModelRowsAboutToBeMoved(ptrName *C.char, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeMoved").(func(*QModelIndex, int, int, *QModelIndex, int))(QModelIndexFromPointer(sourceParent), int(sourceStart), int(sourceEnd), QModelIndexFromPointer(destinationParent), int(destinationRow))
}

func (ptr *QAbstractItemModel) ConnectRowsAboutToBeRemoved(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsAboutToBeRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQAbstractItemModelRowsAboutToBeRemoved
func callbackQAbstractItemModelRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) ConnectRowsInserted(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsInserted() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQAbstractItemModelRowsInserted
func callbackQAbstractItemModelRowsInserted(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsInserted").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) ConnectRowsMoved(f func(parent QModelIndexITF, start int, end int, destination QModelIndexITF, row int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsMoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsMoved")
	}
}

//export callbackQAbstractItemModelRowsMoved
func callbackQAbstractItemModelRowsMoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsMoved").(func(*QModelIndex, int, int, *QModelIndex, int))(QModelIndexFromPointer(parent), int(start), int(end), QModelIndexFromPointer(destination), int(row))
}

func (ptr *QAbstractItemModel) ConnectRowsRemoved(f func(parent QModelIndexITF, first int, last int)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_ConnectRowsRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rowsRemoved", f)
	}
}

func (ptr *QAbstractItemModel) DisconnectRowsRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DisconnectRowsRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rowsRemoved")
	}
}

//export callbackQAbstractItemModelRowsRemoved
func callbackQAbstractItemModelRowsRemoved(ptrName *C.char, parent unsafe.Pointer, first C.int, last C.int) {
	qt.GetSignal(C.GoString(ptrName), "rowsRemoved").(func(*QModelIndex, int, int))(QModelIndexFromPointer(parent), int(first), int(last))
}

func (ptr *QAbstractItemModel) SetData(index QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) SetHeaderData(section int, orientation Qt__Orientation, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_SetHeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) Sort(column int, order Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QAbstractItemModel) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemModel_Submit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) SupportedDragActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractItemModel_SupportedDragActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemModel) SupportedDropActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractItemModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemModel) DestroyQAbstractItemModel() {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_DestroyQAbstractItemModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
