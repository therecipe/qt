package core

//#include "qsortfilterproxymodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QSortFilterProxyModel struct {
	QAbstractProxyModel
}

type QSortFilterProxyModelITF interface {
	QAbstractProxyModelITF
	QSortFilterProxyModelPTR() *QSortFilterProxyModel
}

func PointerFromQSortFilterProxyModel(ptr QSortFilterProxyModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSortFilterProxyModelPTR().Pointer()
	}
	return nil
}

func QSortFilterProxyModelFromPointer(ptr unsafe.Pointer) *QSortFilterProxyModel {
	var n = new(QSortFilterProxyModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSortFilterProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSortFilterProxyModel) QSortFilterProxyModelPTR() *QSortFilterProxyModel {
	return ptr
}

func (ptr *QSortFilterProxyModel) DynamicSortFilter() bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DynamicSortFilter(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FilterCaseSensitivity() Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_FilterCaseSensitivity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterKeyColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterKeyColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterRole() int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) IsSortLocaleAware() bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_IsSortLocaleAware(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetDynamicSortFilter(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterCaseSensitivity(cs Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterCaseSensitivity(C.QtObjectPtr(ptr.Pointer()), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterKeyColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp(regExp QRegExpITF) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegExp(regExp)))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRole(role int) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SetSortCaseSensitivity(cs Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortCaseSensitivity(C.QtObjectPtr(ptr.Pointer()), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortLocaleAware(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSortFilterProxyModel) SetSortRole(role int) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SortCaseSensitivity() Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_SortCaseSensitivity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortRole() int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQSortFilterProxyModel(parent QObjectITF) *QSortFilterProxyModel {
	return QSortFilterProxyModelFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_NewQSortFilterProxyModel(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QSortFilterProxyModel) Buddy(index QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_Buddy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) CanFetchMore(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_CanFetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) ColumnCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) Data(index QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSortFilterProxyModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QSortFilterProxyModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FetchMore(parent QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_FetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent)))
	}
}

func (ptr *QSortFilterProxyModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QSortFilterProxyModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) HasChildren(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_HasChildren(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSortFilterProxyModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QSortFilterProxyModel) Index(row int, column int, parent QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) InsertColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) InsertRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) Invalidate() {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSortFilterProxyModel) MapFromSource(sourceIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_MapFromSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceIndex)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MapToSource(proxyIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_MapToSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(proxyIndex)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSortFilterProxyModel_MimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSortFilterProxyModel) Parent(child QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(child)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) RemoveColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RemoveRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RowCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SetData(index QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterFixedString(C.QtObjectPtr(ptr.Pointer()), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp2(pattern string) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp2(C.QtObjectPtr(ptr.Pointer()), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterWildcard(C.QtObjectPtr(ptr.Pointer()), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetHeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetSourceModel(sourceModel QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSourceModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemModel(sourceModel)))
	}
}

func (ptr *QSortFilterProxyModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QSortFilterProxyModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) Sort(column int, order Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QSortFilterProxyModel) SortColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortOrder() Qt__SortOrder {
	if ptr.Pointer() != nil {
		return Qt__SortOrder(C.QSortFilterProxyModel_SortOrder(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SupportedDropActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QSortFilterProxyModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) DestroyQSortFilterProxyModel() {
	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_DestroyQSortFilterProxyModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
