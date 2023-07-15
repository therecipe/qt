//go:build !minimal
// +build !minimal

package remoteobjects

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/internal"
)

type DataEntries struct {
	internal.Internal
}

type DataEntries_ITF interface {
	DataEntries_PTR() *DataEntries
}

func (ptr *DataEntries) DataEntries_PTR() *DataEntries {
	return ptr
}

func (ptr *DataEntries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *DataEntries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromDataEntries(ptr DataEntries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DataEntries_PTR().Pointer()
	}
	return nil
}

func (n *DataEntries) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewDataEntriesFromPointer(ptr unsafe.Pointer) (n *DataEntries) {
	n = new(DataEntries)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.DataEntries")
	return
}

func (ptr *DataEntries) DestroyDataEntries() {
}

type IndexValuePair struct {
	internal.Internal
}

type IndexValuePair_ITF interface {
	IndexValuePair_PTR() *IndexValuePair
}

func (ptr *IndexValuePair) IndexValuePair_PTR() *IndexValuePair {
	return ptr
}

func (ptr *IndexValuePair) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *IndexValuePair) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromIndexValuePair(ptr IndexValuePair_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IndexValuePair_PTR().Pointer()
	}
	return nil
}

func (n *IndexValuePair) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewIndexValuePairFromPointer(ptr unsafe.Pointer) (n *IndexValuePair) {
	n = new(IndexValuePair)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.IndexValuePair")
	return
}

func (ptr *IndexValuePair) DestroyIndexValuePair() {
}

type ModelIndex struct {
	internal.Internal
}

type ModelIndex_ITF interface {
	ModelIndex_PTR() *ModelIndex
}

func (ptr *ModelIndex) ModelIndex_PTR() *ModelIndex {
	return ptr
}

func (ptr *ModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *ModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromModelIndex(ptr ModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ModelIndex_PTR().Pointer()
	}
	return nil
}

func (n *ModelIndex) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewModelIndexFromPointer(ptr unsafe.Pointer) (n *ModelIndex) {
	n = new(ModelIndex)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.ModelIndex")
	return
}

func (ptr *ModelIndex) DestroyModelIndex() {
}

type QAbstractItemModelReplica struct {
	core.QAbstractItemModel
}

type QAbstractItemModelReplica_ITF interface {
	core.QAbstractItemModel_ITF
	QAbstractItemModelReplica_PTR() *QAbstractItemModelReplica
}

func (ptr *QAbstractItemModelReplica) QAbstractItemModelReplica_PTR() *QAbstractItemModelReplica {
	return ptr
}

func (ptr *QAbstractItemModelReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractItemModelReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemModel_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractItemModelReplica(ptr QAbstractItemModelReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModelReplica_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractItemModelReplica) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractItemModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractItemModelReplica) ClassNameInternalF() string {
	return n.QAbstractItemModel_PTR().ClassNameInternalF()
}

func NewQAbstractItemModelReplicaFromPointer(ptr unsafe.Pointer) (n *QAbstractItemModelReplica) {
	n = new(QAbstractItemModelReplica)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QAbstractItemModelReplica")
	return
}

func (ptr *QAbstractItemModelReplica) DestroyQAbstractItemModelReplica() {
}

func (ptr *QAbstractItemModelReplica) __QAbstractItemModelReplica_rolesHint_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QAbstractItemModelReplica_rolesHint_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) __QAbstractItemModelReplica_rolesHint_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QAbstractItemModelReplica_rolesHint_setList", i})
}

func (ptr *QAbstractItemModelReplica) __QAbstractItemModelReplica_rolesHint_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QAbstractItemModelReplica_rolesHint_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __availableRoles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableRoles_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) __availableRoles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableRoles_setList", i})
}

func (ptr *QAbstractItemModelReplica) __availableRoles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableRoles_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __roleNames_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_atList", v, i}).(*core.QByteArray)
}

func (ptr *QAbstractItemModelReplica) __roleNames_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_setList", key, i})
}

func (ptr *QAbstractItemModelReplica) __roleNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __roleNames_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_keyList"}).([]int)
}

func (ptr *QAbstractItemModelReplica) ____roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) ____roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_setList", i})
}

func (ptr *QAbstractItemModelReplica) ____roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_atList", i}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_setList", i})
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_from_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_atList", i}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_setList", i})
}

func (ptr *QAbstractItemModelReplica) __changePersistentIndexList_to_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QAbstractItemModelReplica) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __itemData_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_atList", v, i}).(*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) __itemData_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_setList", key, i})
}

func (ptr *QAbstractItemModelReplica) __itemData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __itemData_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_keyList"}).([]int)
}

func (ptr *QAbstractItemModelReplica) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QAbstractItemModelReplica) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_setList", i})
}

func (ptr *QAbstractItemModelReplica) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QAbstractItemModelReplica) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_setList", i})
}

func (ptr *QAbstractItemModelReplica) __layoutChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __match_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_atList", i}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) __match_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_setList", i})
}

func (ptr *QAbstractItemModelReplica) __match_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __mimeData_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) __mimeData_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_setList", i})
}

func (ptr *QAbstractItemModelReplica) __mimeData_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __persistentIndexList_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_atList", i}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) __persistentIndexList_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_setList", i})
}

func (ptr *QAbstractItemModelReplica) __persistentIndexList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __setItemData_roles_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_atList", v, i}).(*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) __setItemData_roles_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_setList", key, i})
}

func (ptr *QAbstractItemModelReplica) __setItemData_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __setItemData_roles_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_keyList"}).([]int)
}

func (ptr *QAbstractItemModelReplica) ____doSetRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) ____doSetRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QAbstractItemModelReplica) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) ____itemData_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) ____itemData_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_setList", i})
}

func (ptr *QAbstractItemModelReplica) ____itemData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) ____setItemData_roles_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) ____setItemData_roles_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_setList", i})
}

func (ptr *QAbstractItemModelReplica) ____setItemData_roles_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) ____setRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QAbstractItemModelReplica) ____setRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QAbstractItemModelReplica) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractItemModelReplica) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractItemModelReplica) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractItemModelReplica) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractItemModelReplica) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractItemModelReplica) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractItemModelReplica) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractItemModelReplica) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractItemModelReplica) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractItemModelReplica) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BuddyDefault", index}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanDropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanFetchMoreDefault", parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) ColumnCount(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount", parent}).(float64))
}

func (ptr *QAbstractItemModelReplica) ColumnCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountDefault", parent}).(float64))
}

func (ptr *QAbstractItemModelReplica) Data(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", index, role}).(*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", index, role}).(*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) FetchMoreDefault(parent core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchMoreDefault", parent})
}

func (ptr *QAbstractItemModelReplica) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {

	return core.Qt__ItemFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault", index}).(float64))
}

func (ptr *QAbstractItemModelReplica) HasChildrenDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildrenDefault", parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeaderDataDefault", section, orientation, role}).(*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDefault", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowsDefault", row, count, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemDataDefault", index}).(map[int]*core.QVariant)
}

func (ptr *QAbstractItemModelReplica) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatchDefault", start, role, value, hits, flags}).([]*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeDataDefault", indexes}).(*core.QMimeData)
}

func (ptr *QAbstractItemModelReplica) MimeTypesDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeTypesDefault"}).([]string)
}

func (ptr *QAbstractItemModelReplica) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveColumnsDefault", sourceParent, sourceColumn, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QAbstractItemModelReplica) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveRowsDefault", sourceParent, sourceRow, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QAbstractItemModelReplica) Parent(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parent", index}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentDefault", index}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRowsDefault", row, count, parent}).(bool)
}

func (ptr *QAbstractItemModelReplica) ResetInternalDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetInternalDataDefault"})
}

func (ptr *QAbstractItemModelReplica) RevertDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertDefault"})
}

func (ptr *QAbstractItemModelReplica) RoleNamesDefault() map[int]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoleNamesDefault"}).(map[int]*core.QByteArray)
}

func (ptr *QAbstractItemModelReplica) RowCount(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount", parent}).(float64))
}

func (ptr *QAbstractItemModelReplica) RowCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountDefault", parent}).(float64))
}

func (ptr *QAbstractItemModelReplica) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", index, value, role}).(bool)
}

func (ptr *QAbstractItemModelReplica) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeaderDataDefault", section, orientation, value, role}).(bool)
}

func (ptr *QAbstractItemModelReplica) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemDataDefault", index, roles}).(bool)
}

func (ptr *QAbstractItemModelReplica) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SiblingDefault", row, column, index}).(*core.QModelIndex)
}

func (ptr *QAbstractItemModelReplica) SortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortDefault", column, order})
}

func (ptr *QAbstractItemModelReplica) SpanDefault(index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpanDefault", index}).(*core.QSize)
}

func (ptr *QAbstractItemModelReplica) SubmitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitDefault"}).(bool)
}

func (ptr *QAbstractItemModelReplica) SupportedDragActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDragActionsDefault"}).(float64))
}

func (ptr *QAbstractItemModelReplica) SupportedDropActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDropActionsDefault"}).(float64))
}

func (ptr *QAbstractItemModelReplica) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractItemModelReplica) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractItemModelReplica) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractItemModelReplica) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractItemModelReplica) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractItemModelReplica) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractItemModelReplica) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractItemModelReplica) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractItemModelReplica) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QIOQnxSource struct {
	core.QIODevice
}

type QIOQnxSource_ITF interface {
	core.QIODevice_ITF
	QIOQnxSource_PTR() *QIOQnxSource
}

func (ptr *QIOQnxSource) QIOQnxSource_PTR() *QIOQnxSource {
	return ptr
}

func (ptr *QIOQnxSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QIOQnxSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQIOQnxSource(ptr QIOQnxSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIOQnxSource_PTR().Pointer()
	}
	return nil
}

func (n *QIOQnxSource) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QIOQnxSource) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQIOQnxSourceFromPointer(ptr unsafe.Pointer) (n *QIOQnxSource) {
	n = new(QIOQnxSource)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QIOQnxSource")
	return
}

func (ptr *QIOQnxSource) DestroyQIOQnxSource() {
}

type QMetaTypeId struct {
	internal.Internal
}

type QMetaTypeId_ITF interface {
	QMetaTypeId_PTR() *QMetaTypeId
}

func (ptr *QMetaTypeId) QMetaTypeId_PTR() *QMetaTypeId {
	return ptr
}

func (ptr *QMetaTypeId) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QMetaTypeId) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQMetaTypeId(ptr QMetaTypeId_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaTypeId_PTR().Pointer()
	}
	return nil
}

func (n *QMetaTypeId) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQMetaTypeIdFromPointer(ptr unsafe.Pointer) (n *QMetaTypeId) {
	n = new(QMetaTypeId)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QMetaTypeId")
	return
}

func (ptr *QMetaTypeId) DestroyQMetaTypeId() {
}

type QQnxNativeIo struct {
	core.QIODevice
}

type QQnxNativeIo_ITF interface {
	core.QIODevice_ITF
	QQnxNativeIo_PTR() *QQnxNativeIo
}

func (ptr *QQnxNativeIo) QQnxNativeIo_PTR() *QQnxNativeIo {
	return ptr
}

func (ptr *QQnxNativeIo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QQnxNativeIo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQQnxNativeIo(ptr QQnxNativeIo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQnxNativeIo_PTR().Pointer()
	}
	return nil
}

func (n *QQnxNativeIo) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQnxNativeIo) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQQnxNativeIoFromPointer(ptr unsafe.Pointer) (n *QQnxNativeIo) {
	n = new(QQnxNativeIo)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QQnxNativeIo")
	return
}

func (ptr *QQnxNativeIo) DestroyQQnxNativeIo() {
}

type QQnxNativeServer struct {
	core.QObject
}

type QQnxNativeServer_ITF interface {
	core.QObject_ITF
	QQnxNativeServer_PTR() *QQnxNativeServer
}

func (ptr *QQnxNativeServer) QQnxNativeServer_PTR() *QQnxNativeServer {
	return ptr
}

func (ptr *QQnxNativeServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQnxNativeServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQnxNativeServer(ptr QQnxNativeServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQnxNativeServer_PTR().Pointer()
	}
	return nil
}

func (n *QQnxNativeServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQnxNativeServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQnxNativeServerFromPointer(ptr unsafe.Pointer) (n *QQnxNativeServer) {
	n = new(QQnxNativeServer)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QQnxNativeServer")
	return
}

func (ptr *QQnxNativeServer) DestroyQQnxNativeServer() {
}

type QRemoteObjectAbstractPersistedStore struct {
	core.QObject
}

type QRemoteObjectAbstractPersistedStore_ITF interface {
	core.QObject_ITF
	QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore
}

func (ptr *QRemoteObjectAbstractPersistedStore) QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore {
	return ptr
}

func (ptr *QRemoteObjectAbstractPersistedStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectAbstractPersistedStore(ptr QRemoteObjectAbstractPersistedStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectAbstractPersistedStore_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectAbstractPersistedStore) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectAbstractPersistedStore) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectAbstractPersistedStore) {
	n = new(QRemoteObjectAbstractPersistedStore)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectAbstractPersistedStore")
	return
}

func (ptr *QRemoteObjectAbstractPersistedStore) DestroyQRemoteObjectAbstractPersistedStore() {
}

func NewQRemoteObjectAbstractPersistedStore(parent core.QObject_ITF) *QRemoteObjectAbstractPersistedStore {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectAbstractPersistedStore", "", parent}).(*QRemoteObjectAbstractPersistedStore)
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectRestoreProperties(f func(repName string, repSig *core.QByteArray) []*core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRestoreProperties", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectRestoreProperties() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRestoreProperties"})
}

func (ptr *QRemoteObjectAbstractPersistedStore) RestoreProperties(repName string, repSig core.QByteArray_ITF) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RestoreProperties", repName, repSig}).([]*core.QVariant)
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectSaveProperties(f func(repName string, repSig *core.QByteArray, values []*core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSaveProperties", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectSaveProperties() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSaveProperties"})
}

func (ptr *QRemoteObjectAbstractPersistedStore) SaveProperties(repName string, repSig core.QByteArray_ITF, values []*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SaveProperties", repName, repSig, values})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__restoreProperties_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__restoreProperties_setList", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__restoreProperties_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__saveProperties_values_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__saveProperties_values_setList", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__saveProperties_values_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectAbstractPersistedStore) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectAbstractPersistedStore) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QRemoteObjectAbstractPersistedStore) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QRemoteObjectAbstractPersistedStore) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QRemoteObjectAbstractPersistedStore) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QRemoteObjectDynamicReplica struct {
	QRemoteObjectReplica
}

type QRemoteObjectDynamicReplica_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica
}

func (ptr *QRemoteObjectDynamicReplica) QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica {
	return ptr
}

func (ptr *QRemoteObjectDynamicReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectDynamicReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectDynamicReplica(ptr QRemoteObjectDynamicReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectDynamicReplica_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectDynamicReplica) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectReplica_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectDynamicReplica) ClassNameInternalF() string {
	return n.QRemoteObjectReplica_PTR().ClassNameInternalF()
}

func NewQRemoteObjectDynamicReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectDynamicReplica) {
	n = new(QRemoteObjectDynamicReplica)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectDynamicReplica")
	return
}
func (ptr *QRemoteObjectDynamicReplica) ConnectDestroyQRemoteObjectDynamicReplica(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQRemoteObjectDynamicReplica", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectDynamicReplica) DisconnectDestroyQRemoteObjectDynamicReplica() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQRemoteObjectDynamicReplica"})
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplica() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRemoteObjectDynamicReplica"})
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplicaDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRemoteObjectDynamicReplicaDefault"})
}

type QRemoteObjectHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectHost_PTR() *QRemoteObjectHost
}

func (ptr *QRemoteObjectHost) QRemoteObjectHost_PTR() *QRemoteObjectHost {
	return ptr
}

func (ptr *QRemoteObjectHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHost(ptr QRemoteObjectHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHost_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectHost) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectHostBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectHost) ClassNameInternalF() string {
	return n.QRemoteObjectHostBase_PTR().ClassNameInternalF()
}

func NewQRemoteObjectHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHost) {
	n = new(QRemoteObjectHost)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectHost")
	return
}

func (ptr *QRemoteObjectHost) DestroyQRemoteObjectHost() {
}

func NewQRemoteObjectHost(parent core.QObject_ITF) *QRemoteObjectHost {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectHost", "", parent}).(*QRemoteObjectHost)
}

func NewQRemoteObjectHost2(address core.QUrl_ITF, registryAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas, parent core.QObject_ITF) *QRemoteObjectHost {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectHost2", "", address, registryAddress, allowedSchemas, parent}).(*QRemoteObjectHost)
}

func NewQRemoteObjectHost3(address core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectHost {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectHost3", "", address, parent}).(*QRemoteObjectHost)
}

func (ptr *QRemoteObjectHost) ConnectHostUrl(f func() *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHostUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectHost) DisconnectHostUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHostUrl"})
}

func (ptr *QRemoteObjectHost) HostUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostUrl"}).(*core.QUrl)
}

func (ptr *QRemoteObjectHost) HostUrlDefault() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostUrlDefault"}).(*core.QUrl)
}

func (ptr *QRemoteObjectHost) ConnectSetHostUrl(f func(hostAddress *core.QUrl, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetHostUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectHost) DisconnectSetHostUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetHostUrl"})
}

func (ptr *QRemoteObjectHost) SetHostUrl(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostUrl", hostAddress, allowedSchemas}).(bool)
}

func (ptr *QRemoteObjectHost) SetHostUrlDefault(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostUrlDefault", hostAddress, allowedSchemas}).(bool)
}

type QRemoteObjectHostBase struct {
	QRemoteObjectNode
}

type QRemoteObjectHostBase_ITF interface {
	QRemoteObjectNode_ITF
	QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase
}

func (ptr *QRemoteObjectHostBase) QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase {
	return ptr
}

func (ptr *QRemoteObjectHostBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHostBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectNode_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHostBase(ptr QRemoteObjectHostBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectHostBase) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectHostBase) ClassNameInternalF() string {
	return n.QRemoteObjectNode_PTR().ClassNameInternalF()
}

func NewQRemoteObjectHostBaseFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHostBase) {
	n = new(QRemoteObjectHostBase)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectHostBase")
	return
}

func (ptr *QRemoteObjectHostBase) DestroyQRemoteObjectHostBase() {
}

//go:generate stringer -type=QRemoteObjectHostBase__AllowedSchemas
//QRemoteObjectHostBase::AllowedSchemas
type QRemoteObjectHostBase__AllowedSchemas int64

const (
	QRemoteObjectHostBase__BuiltInSchemasOnly        QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(0)
	QRemoteObjectHostBase__AllowExternalRegistration QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(1)
)

func (ptr *QRemoteObjectHostBase) AddHostSideConnection(ioDevice core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddHostSideConnection", ioDevice})
}

func (ptr *QRemoteObjectHostBase) DisableRemoting(remoteObject core.QObject_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisableRemoting", remoteObject}).(bool)
}

func (ptr *QRemoteObjectHostBase) EnableRemoting2(object core.QObject_ITF, name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnableRemoting2", object, name}).(bool)
}

func (ptr *QRemoteObjectHostBase) EnableRemoting3(model core.QAbstractItemModel_ITF, name string, roles []int, selectionModel core.QItemSelectionModel_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnableRemoting3", model, name, roles, selectionModel}).(bool)
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_atList3(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__enableRemoting_roles_atList3", i}).(float64))
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_setList3(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__enableRemoting_roles_setList3", i})
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__enableRemoting_roles_newList3"}).(unsafe.Pointer)
}

type QRemoteObjectNode struct {
	core.QObject
}

type QRemoteObjectNode_ITF interface {
	core.QObject_ITF
	QRemoteObjectNode_PTR() *QRemoteObjectNode
}

func (ptr *QRemoteObjectNode) QRemoteObjectNode_PTR() *QRemoteObjectNode {
	return ptr
}

func (ptr *QRemoteObjectNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectNode(ptr QRemoteObjectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectNode) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectNode) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQRemoteObjectNodeFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectNode) {
	n = new(QRemoteObjectNode)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectNode")
	return
}

func (ptr *QRemoteObjectNode) DestroyQRemoteObjectNode() {
}

//go:generate stringer -type=QRemoteObjectNode__ErrorCode
//QRemoteObjectNode::ErrorCode
type QRemoteObjectNode__ErrorCode int64

const (
	QRemoteObjectNode__NoError                       QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(0)
	QRemoteObjectNode__RegistryNotAcquired           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(1)
	QRemoteObjectNode__RegistryAlreadyHosted         QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(2)
	QRemoteObjectNode__NodeIsNoServer                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(3)
	QRemoteObjectNode__ServerAlreadyCreated          QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(4)
	QRemoteObjectNode__UnintendedRegistryHosting     QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(5)
	QRemoteObjectNode__OperationNotValidOnClientNode QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(6)
	QRemoteObjectNode__SourceNotRegistered           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(7)
	QRemoteObjectNode__MissingObjectName             QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(8)
	QRemoteObjectNode__HostUrlInvalid                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(9)
	QRemoteObjectNode__ProtocolMismatch              QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(10)
	QRemoteObjectNode__ListenFailed                  QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(11)
)

func NewQRemoteObjectNode(parent core.QObject_ITF) *QRemoteObjectNode {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectNode", "", parent}).(*QRemoteObjectNode)
}

func NewQRemoteObjectNode2(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectNode {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectNode2", "", registryAddress, parent}).(*QRemoteObjectNode)
}

func (ptr *QRemoteObjectNode) AcquireDynamic(name string) *QRemoteObjectDynamicReplica {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcquireDynamic", name}).(*QRemoteObjectDynamicReplica)
}

func (ptr *QRemoteObjectNode) AddClientSideConnection(ioDevice core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddClientSideConnection", ioDevice})
}

func (ptr *QRemoteObjectNode) ConnectToNode(address core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToNode", address}).(bool)
}

func (ptr *QRemoteObjectNode) HeartbeatInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeartbeatInterval"}).(float64))
}

func (ptr *QRemoteObjectNode) ConnectHeartbeatIntervalChanged(f func(heartbeatInterval int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHeartbeatIntervalChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectNode) DisconnectHeartbeatIntervalChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHeartbeatIntervalChanged"})
}

func (ptr *QRemoteObjectNode) HeartbeatIntervalChanged(heartbeatInterval int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeartbeatIntervalChanged", heartbeatInterval})
}

func (ptr *QRemoteObjectNode) Instances2(typeName string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Instances2", typeName}).([]string)
}

func (ptr *QRemoteObjectNode) LastError() QRemoteObjectNode__ErrorCode {

	return QRemoteObjectNode__ErrorCode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(float64))
}

func (ptr *QRemoteObjectNode) PersistedStore() *QRemoteObjectAbstractPersistedStore {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistedStore"}).(*QRemoteObjectAbstractPersistedStore)
}

func (ptr *QRemoteObjectNode) Registry() *QRemoteObjectRegistry {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Registry"}).(*QRemoteObjectRegistry)
}

func (ptr *QRemoteObjectNode) RegistryUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegistryUrl"}).(*core.QUrl)
}

func (ptr *QRemoteObjectNode) DisconnectRemoteObjectAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoteObjectAdded"})
}

func (ptr *QRemoteObjectNode) SetHeartbeatInterval(interval int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeartbeatInterval", interval})
}

func (ptr *QRemoteObjectNode) ConnectSetName(f func(name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectNode) DisconnectSetName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetName"})
}

func (ptr *QRemoteObjectNode) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QRemoteObjectNode) SetNameDefault(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNameDefault", name})
}

func (ptr *QRemoteObjectNode) SetPersistedStore(persistedStore QRemoteObjectAbstractPersistedStore_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistedStore", persistedStore})
}

func (ptr *QRemoteObjectNode) ConnectSetRegistryUrl(f func(registryAddress *core.QUrl) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRegistryUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectNode) DisconnectSetRegistryUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRegistryUrl"})
}

func (ptr *QRemoteObjectNode) SetRegistryUrl(registryAddress core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRegistryUrl", registryAddress}).(bool)
}

func (ptr *QRemoteObjectNode) SetRegistryUrlDefault(registryAddress core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRegistryUrlDefault", registryAddress}).(bool)
}

func (ptr *QRemoteObjectNode) TimerEventDefault(vqt core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", vqt})
}

func (ptr *QRemoteObjectNode) WaitForRegistry(timeout int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForRegistry", timeout}).(bool)
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__acquireModel_rolesHint_atList", i}).(float64))
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__acquireModel_rolesHint_setList", i})
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__acquireModel_rolesHint_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __persistProperties_props_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectNode) __persistProperties_props_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_setList", i})
}

func (ptr *QRemoteObjectNode) __persistProperties_props_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __retrieveProperties_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectNode) __retrieveProperties_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_setList", i})
}

func (ptr *QRemoteObjectNode) __retrieveProperties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectNode) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QRemoteObjectNode) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectNode) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QRemoteObjectNode) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QRemoteObjectNode) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QRemoteObjectNode) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectNode) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QRemoteObjectNode) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectNode) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QRemoteObjectNode) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QRemoteObjectNode) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectNode) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QRemoteObjectNode) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QRemoteObjectNode) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

type QRemoteObjectPendingCall struct {
	internal.Internal
}

type QRemoteObjectPendingCall_ITF interface {
	QRemoteObjectPendingCall_PTR() *QRemoteObjectPendingCall
}

func (ptr *QRemoteObjectPendingCall) QRemoteObjectPendingCall_PTR() *QRemoteObjectPendingCall {
	return ptr
}

func (ptr *QRemoteObjectPendingCall) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QRemoteObjectPendingCall) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQRemoteObjectPendingCall(ptr QRemoteObjectPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCall_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectPendingCall) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQRemoteObjectPendingCallFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingCall) {
	n = new(QRemoteObjectPendingCall)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectPendingCall")
	return
}

func (ptr *QRemoteObjectPendingCall) DestroyQRemoteObjectPendingCall() {
}

type QRemoteObjectPendingCallWatcher struct {
	core.QObject
	QRemoteObjectPendingCall
}

type QRemoteObjectPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QRemoteObjectPendingCall_ITF
	QRemoteObjectPendingCallWatcher_PTR() *QRemoteObjectPendingCallWatcher
}

func (ptr *QRemoteObjectPendingCallWatcher) QRemoteObjectPendingCallWatcher_PTR() *QRemoteObjectPendingCallWatcher {
	return ptr
}

func (ptr *QRemoteObjectPendingCallWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QRemoteObjectPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectPendingCallWatcher(ptr QRemoteObjectPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectPendingCallWatcher) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QRemoteObjectPendingCall_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectPendingCallWatcher) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQRemoteObjectPendingCallWatcherFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingCallWatcher) {
	n = new(QRemoteObjectPendingCallWatcher)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectPendingCallWatcher")
	return
}

func (ptr *QRemoteObjectPendingCallWatcher) DestroyQRemoteObjectPendingCallWatcher() {
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QRemoteObjectPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QRemoteObjectPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QRemoteObjectPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QRemoteObjectPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QRemoteObjectPendingCallWatcher) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QRemoteObjectPendingCallWatcher) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QRemoteObjectPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QRemoteObjectPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectPendingCallWatcher) Event(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", e}).(bool)
}

func (ptr *QRemoteObjectPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QRemoteObjectPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QRemoteObjectPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QRemoteObjectPendingCallWatcher) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QRemoteObjectPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QRemoteObjectPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QRemoteObjectPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QRemoteObjectPendingReply struct {
	QRemoteObjectPendingCall
}

type QRemoteObjectPendingReply_ITF interface {
	QRemoteObjectPendingCall_ITF
	QRemoteObjectPendingReply_PTR() *QRemoteObjectPendingReply
}

func (ptr *QRemoteObjectPendingReply) QRemoteObjectPendingReply_PTR() *QRemoteObjectPendingReply {
	return ptr
}

func (ptr *QRemoteObjectPendingReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCall_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectPendingReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectPendingReply(ptr QRemoteObjectPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingReply_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectPendingReply) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectPendingCall_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectPendingReply) ClassNameInternalF() string {
	return n.QRemoteObjectPendingCall_PTR().ClassNameInternalF()
}

func NewQRemoteObjectPendingReplyFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingReply) {
	n = new(QRemoteObjectPendingReply)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectPendingReply")
	return
}

func (ptr *QRemoteObjectPendingReply) DestroyQRemoteObjectPendingReply() {
}

type QRemoteObjectRegistry struct {
	QRemoteObjectReplica
}

type QRemoteObjectRegistry_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry
}

func (ptr *QRemoteObjectRegistry) QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry {
	return ptr
}

func (ptr *QRemoteObjectRegistry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistry(ptr QRemoteObjectRegistry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistry_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectRegistry) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectReplica_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectRegistry) ClassNameInternalF() string {
	return n.QRemoteObjectReplica_PTR().ClassNameInternalF()
}

func NewQRemoteObjectRegistryFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistry) {
	n = new(QRemoteObjectRegistry)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectRegistry")
	return
}
func (ptr *QRemoteObjectRegistry) DisconnectRemoteObjectAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoteObjectAdded"})
}

func (ptr *QRemoteObjectRegistry) DisconnectRemoteObjectRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoteObjectRemoved"})
}

func (ptr *QRemoteObjectRegistry) ConnectDestroyQRemoteObjectRegistry(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQRemoteObjectRegistry", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectRegistry) DisconnectDestroyQRemoteObjectRegistry() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQRemoteObjectRegistry"})
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistry() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRemoteObjectRegistry"})
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRemoteObjectRegistryDefault"})
}

type QRemoteObjectRegistryHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectRegistryHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost
}

func (ptr *QRemoteObjectRegistryHost) QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost {
	return ptr
}

func (ptr *QRemoteObjectRegistryHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistryHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistryHost(ptr QRemoteObjectRegistryHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistryHost_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectRegistryHost) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectHostBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectRegistryHost) ClassNameInternalF() string {
	return n.QRemoteObjectHostBase_PTR().ClassNameInternalF()
}

func NewQRemoteObjectRegistryHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistryHost) {
	n = new(QRemoteObjectRegistryHost)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectRegistryHost")
	return
}

func (ptr *QRemoteObjectRegistryHost) DestroyQRemoteObjectRegistryHost() {
}

func NewQRemoteObjectRegistryHost(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectRegistryHost {

	return internal.CallLocalFunction([]interface{}{"", "", "remoteobjects.NewQRemoteObjectRegistryHost", "", registryAddress, parent}).(*QRemoteObjectRegistryHost)
}

type QRemoteObjectReplica struct {
	core.QObject
}

type QRemoteObjectReplica_ITF interface {
	core.QObject_ITF
	QRemoteObjectReplica_PTR() *QRemoteObjectReplica
}

func (ptr *QRemoteObjectReplica) QRemoteObjectReplica_PTR() *QRemoteObjectReplica {
	return ptr
}

func (ptr *QRemoteObjectReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectReplica(ptr QRemoteObjectReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectReplica) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectReplica) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQRemoteObjectReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectReplica) {
	n = new(QRemoteObjectReplica)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectReplica")
	return
}

func (ptr *QRemoteObjectReplica) DestroyQRemoteObjectReplica() {
}

//go:generate stringer -type=QRemoteObjectReplica__State
//QRemoteObjectReplica::State
type QRemoteObjectReplica__State int64

const (
	QRemoteObjectReplica__Uninitialized     QRemoteObjectReplica__State = QRemoteObjectReplica__State(0)
	QRemoteObjectReplica__Default           QRemoteObjectReplica__State = QRemoteObjectReplica__State(1)
	QRemoteObjectReplica__Valid             QRemoteObjectReplica__State = QRemoteObjectReplica__State(2)
	QRemoteObjectReplica__Suspect           QRemoteObjectReplica__State = QRemoteObjectReplica__State(3)
	QRemoteObjectReplica__SignatureMismatch QRemoteObjectReplica__State = QRemoteObjectReplica__State(4)
)

func (ptr *QRemoteObjectReplica) ConnectInitialized(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInitialized", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectReplica) DisconnectInitialized() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInitialized"})
}

func (ptr *QRemoteObjectReplica) Initialized() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initialized"})
}

func (ptr *QRemoteObjectReplica) IsInitialized() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInitialized"}).(bool)
}

func (ptr *QRemoteObjectReplica) IsReplicaValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReplicaValid"}).(bool)
}

func (ptr *QRemoteObjectReplica) Node() *QRemoteObjectNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Node"}).(*QRemoteObjectNode)
}

func (ptr *QRemoteObjectReplica) ConnectSetNode(f func(node *QRemoteObjectNode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetNode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectReplica) DisconnectSetNode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetNode"})
}

func (ptr *QRemoteObjectReplica) SetNode(node QRemoteObjectNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNode", node})
}

func (ptr *QRemoteObjectReplica) SetNodeDefault(node QRemoteObjectNode_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNodeDefault", node})
}

func (ptr *QRemoteObjectReplica) State() QRemoteObjectReplica__State {

	return QRemoteObjectReplica__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QRemoteObjectReplica) ConnectStateChanged(f func(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRemoteObjectReplica) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QRemoteObjectReplica) StateChanged(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state, oldState})
}

func (ptr *QRemoteObjectReplica) WaitForSource(timeout int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForSource", timeout}).(bool)
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_setList", i})
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistProperties_props_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_setList", i})
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__retrieveProperties_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __send_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__send_args_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectReplica) __send_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__send_args_setList", i})
}

func (ptr *QRemoteObjectReplica) __send_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__send_args_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendWithReply_args_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendWithReply_args_setList", i})
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendWithReply_args_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setProperties_properties_atList", i}).(*core.QVariant)
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setProperties_properties_setList", i})
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setProperties_properties_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectReplica) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QRemoteObjectReplica) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QRemoteObjectReplica) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QRemoteObjectReplica) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QRemoteObjectReplica) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QRemoteObjectReplica) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QRemoteObjectReplica) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QRemoteObjectReplica) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectReplica) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QRemoteObjectReplica) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QRemoteObjectReplica) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QRemoteObjectReplica) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QRemoteObjectReplica) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QRemoteObjectReplica) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QRemoteObjectReplica) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QRemoteObjectSettingsStore struct {
	QRemoteObjectAbstractPersistedStore
}

type QRemoteObjectSettingsStore_ITF interface {
	QRemoteObjectAbstractPersistedStore_ITF
	QRemoteObjectSettingsStore_PTR() *QRemoteObjectSettingsStore
}

func (ptr *QRemoteObjectSettingsStore) QRemoteObjectSettingsStore_PTR() *QRemoteObjectSettingsStore {
	return ptr
}

func (ptr *QRemoteObjectSettingsStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectAbstractPersistedStore_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectSettingsStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectAbstractPersistedStore_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectSettingsStore(ptr QRemoteObjectSettingsStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectSettingsStore_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectSettingsStore) InitFromInternal(ptr uintptr, name string) {
	n.QRemoteObjectAbstractPersistedStore_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRemoteObjectSettingsStore) ClassNameInternalF() string {
	return n.QRemoteObjectAbstractPersistedStore_PTR().ClassNameInternalF()
}

func NewQRemoteObjectSettingsStoreFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectSettingsStore) {
	n = new(QRemoteObjectSettingsStore)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectSettingsStore")
	return
}

func (ptr *QRemoteObjectSettingsStore) DestroyQRemoteObjectSettingsStore() {
}

func (ptr *QRemoteObjectSettingsStore) RestoreProperties(repName string, repSig core.QByteArray_ITF) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RestoreProperties", repName, repSig}).([]*core.QVariant)
}

func (ptr *QRemoteObjectSettingsStore) RestorePropertiesDefault(repName string, repSig core.QByteArray_ITF) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RestorePropertiesDefault", repName, repSig}).([]*core.QVariant)
}

func (ptr *QRemoteObjectSettingsStore) SaveProperties(repName string, repSig core.QByteArray_ITF, values []*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SaveProperties", repName, repSig, values})
}

func (ptr *QRemoteObjectSettingsStore) SavePropertiesDefault(repName string, repSig core.QByteArray_ITF, values []*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SavePropertiesDefault", repName, repSig, values})
}

type QRemoteObjectSourceLocationInfo struct {
	internal.Internal
}

type QRemoteObjectSourceLocationInfo_ITF interface {
	QRemoteObjectSourceLocationInfo_PTR() *QRemoteObjectSourceLocationInfo
}

func (ptr *QRemoteObjectSourceLocationInfo) QRemoteObjectSourceLocationInfo_PTR() *QRemoteObjectSourceLocationInfo {
	return ptr
}

func (ptr *QRemoteObjectSourceLocationInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QRemoteObjectSourceLocationInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQRemoteObjectSourceLocationInfo(ptr QRemoteObjectSourceLocationInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectSourceLocationInfo_PTR().Pointer()
	}
	return nil
}

func (n *QRemoteObjectSourceLocationInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQRemoteObjectSourceLocationInfoFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectSourceLocationInfo) {
	n = new(QRemoteObjectSourceLocationInfo)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QRemoteObjectSourceLocationInfo")
	return
}

func (ptr *QRemoteObjectSourceLocationInfo) DestroyQRemoteObjectSourceLocationInfo() {
}

type QTypeInfo struct {
	internal.Internal
}

type QTypeInfo_ITF interface {
	QTypeInfo_PTR() *QTypeInfo
}

func (ptr *QTypeInfo) QTypeInfo_PTR() *QTypeInfo {
	return ptr
}

func (ptr *QTypeInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTypeInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTypeInfo(ptr QTypeInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTypeInfo_PTR().Pointer()
	}
	return nil
}

func (n *QTypeInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTypeInfoFromPointer(ptr unsafe.Pointer) (n *QTypeInfo) {
	n = new(QTypeInfo)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QTypeInfo")
	return
}

func (ptr *QTypeInfo) DestroyQTypeInfo() {
}

type QtROClientFactory struct {
	internal.Internal
}

type QtROClientFactory_ITF interface {
	QtROClientFactory_PTR() *QtROClientFactory
}

func (ptr *QtROClientFactory) QtROClientFactory_PTR() *QtROClientFactory {
	return ptr
}

func (ptr *QtROClientFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QtROClientFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQtROClientFactory(ptr QtROClientFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROClientFactory_PTR().Pointer()
	}
	return nil
}

func (n *QtROClientFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQtROClientFactoryFromPointer(ptr unsafe.Pointer) (n *QtROClientFactory) {
	n = new(QtROClientFactory)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QtROClientFactory")
	return
}

func (ptr *QtROClientFactory) DestroyQtROClientFactory() {
}

type QtROServerFactory struct {
	internal.Internal
}

type QtROServerFactory_ITF interface {
	QtROServerFactory_PTR() *QtROServerFactory
}

func (ptr *QtROServerFactory) QtROServerFactory_PTR() *QtROServerFactory {
	return ptr
}

func (ptr *QtROServerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QtROServerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQtROServerFactory(ptr QtROServerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROServerFactory_PTR().Pointer()
	}
	return nil
}

func (n *QtROServerFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQtROServerFactoryFromPointer(ptr unsafe.Pointer) (n *QtROServerFactory) {
	n = new(QtROServerFactory)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.QtROServerFactory")
	return
}

func (ptr *QtROServerFactory) DestroyQtROServerFactory() {
}

type SourceApiMap struct {
	internal.Internal
}

type SourceApiMap_ITF interface {
	SourceApiMap_PTR() *SourceApiMap
}

func (ptr *SourceApiMap) SourceApiMap_PTR() *SourceApiMap {
	return ptr
}

func (ptr *SourceApiMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *SourceApiMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromSourceApiMap(ptr SourceApiMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SourceApiMap_PTR().Pointer()
	}
	return nil
}

func (n *SourceApiMap) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewSourceApiMapFromPointer(ptr unsafe.Pointer) (n *SourceApiMap) {
	n = new(SourceApiMap)
	n.InitFromInternal(uintptr(ptr), "remoteobjects.SourceApiMap")
	return
}

func (ptr *SourceApiMap) DestroySourceApiMap() {
}

func init() {
	internal.ConstructorTable["remoteobjects.QAbstractItemModelReplica"] = NewQAbstractItemModelReplicaFromPointer
	internal.ConstructorTable["remoteobjects.QMetaTypeId"] = NewQMetaTypeIdFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectAbstractPersistedStore"] = NewQRemoteObjectAbstractPersistedStoreFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectDynamicReplica"] = NewQRemoteObjectDynamicReplicaFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectHost"] = NewQRemoteObjectHostFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectHostBase"] = NewQRemoteObjectHostBaseFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectNode"] = NewQRemoteObjectNodeFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectPendingCall"] = NewQRemoteObjectPendingCallFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectPendingCallWatcher"] = NewQRemoteObjectPendingCallWatcherFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectPendingReply"] = NewQRemoteObjectPendingReplyFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectRegistry"] = NewQRemoteObjectRegistryFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectRegistryHost"] = NewQRemoteObjectRegistryHostFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectReplica"] = NewQRemoteObjectReplicaFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectSettingsStore"] = NewQRemoteObjectSettingsStoreFromPointer
	internal.ConstructorTable["remoteobjects.QRemoteObjectSourceLocationInfo"] = NewQRemoteObjectSourceLocationInfoFromPointer
	internal.ConstructorTable["remoteobjects.QTypeInfo"] = NewQTypeInfoFromPointer
}
