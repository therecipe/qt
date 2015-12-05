package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QSortFilterProxyModel struct {
	QAbstractProxyModel
}

type QSortFilterProxyModel_ITF interface {
	QAbstractProxyModel_ITF
	QSortFilterProxyModel_PTR() *QSortFilterProxyModel
}

func PointerFromQSortFilterProxyModel(ptr QSortFilterProxyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSortFilterProxyModel_PTR().Pointer()
	}
	return nil
}

func NewQSortFilterProxyModelFromPointer(ptr unsafe.Pointer) *QSortFilterProxyModel {
	var n = new(QSortFilterProxyModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSortFilterProxyModel_") {
		n.SetObjectName("QSortFilterProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSortFilterProxyModel) QSortFilterProxyModel_PTR() *QSortFilterProxyModel {
	return ptr
}

func (ptr *QSortFilterProxyModel) DynamicSortFilter() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::dynamicSortFilter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DynamicSortFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FilterCaseSensitivity() Qt__CaseSensitivity {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::filterCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_FilterCaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterKeyColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::filterKeyColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterKeyColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterRegExp() *QRegExp {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::filterRegExp")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegExpFromPointer(C.QSortFilterProxyModel_FilterRegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) FilterRole() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::filterRole")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) IsSortLocaleAware() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::isSortLocaleAware")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_IsSortLocaleAware(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setDynamicSortFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetDynamicSortFilter(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterCaseSensitivity(cs Qt__CaseSensitivity) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterKeyColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterKeyColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp(regExp QRegExp_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterRegExp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp(ptr.Pointer(), PointerFromQRegExp(regExp))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRole(role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterRole")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SetSortCaseSensitivity(cs Qt__CaseSensitivity) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setSortCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setSortLocaleAware")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortLocaleAware(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSortFilterProxyModel) SetSortRole(role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setSortRole")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SortCaseSensitivity() Qt__CaseSensitivity {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sortCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_SortCaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortRole() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sortRole")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortRole(ptr.Pointer()))
	}
	return 0
}

func NewQSortFilterProxyModel(parent QObject_ITF) *QSortFilterProxyModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::QSortFilterProxyModel")
		}
	}()

	return NewQSortFilterProxyModelFromPointer(C.QSortFilterProxyModel_NewQSortFilterProxyModel(PointerFromQObject(parent)))
}

func (ptr *QSortFilterProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::buddy")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Buddy(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::canFetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_CanFetchMore(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) Data(index QModelIndex_ITF, role int) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QSortFilterProxyModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::dropMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FetchMore(parent QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::fetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QSortFilterProxyModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QSortFilterProxyModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::hasChildren")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QSortFilterProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::index")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::insertColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::insertRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSortFilterProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::mapFromSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::mapToSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MimeTypes() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::mimeTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSortFilterProxyModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSortFilterProxyModel) Parent(child QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Parent(ptr.Pointer(), PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::removeRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RowCount(parent QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterFixedString")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterFixedString(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp2(pattern string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterRegExp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp2(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setFilterWildcard")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterWildcard(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setHeaderData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::setSourceModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QSortFilterProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sibling")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) Sort(column int, order Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSortFilterProxyModel) SortColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sortColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortOrder() Qt__SortOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::sortOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__SortOrder(C.QSortFilterProxyModel_SortOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SupportedDropActions() Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::supportedDropActions")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QSortFilterProxyModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) DestroyQSortFilterProxyModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSortFilterProxyModel::~QSortFilterProxyModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_DestroyQSortFilterProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
