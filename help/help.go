// +build !minimal

package help

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QCompressedHelpInfo struct {
	internal.Internal
}

type QCompressedHelpInfo_ITF interface {
	QCompressedHelpInfo_PTR() *QCompressedHelpInfo
}

func (ptr *QCompressedHelpInfo) QCompressedHelpInfo_PTR() *QCompressedHelpInfo {
	return ptr
}

func (ptr *QCompressedHelpInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCompressedHelpInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCompressedHelpInfo(ptr QCompressedHelpInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompressedHelpInfo_PTR().Pointer()
	}
	return nil
}

func (n *QCompressedHelpInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCompressedHelpInfoFromPointer(ptr unsafe.Pointer) (n *QCompressedHelpInfo) {
	n = new(QCompressedHelpInfo)
	n.InitFromInternal(uintptr(ptr), "help.QCompressedHelpInfo")
	return
}
func NewQCompressedHelpInfo() *QCompressedHelpInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQCompressedHelpInfo", ""}).(*QCompressedHelpInfo)
}

func NewQCompressedHelpInfo2(other QCompressedHelpInfo_ITF) *QCompressedHelpInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQCompressedHelpInfo2", "", other}).(*QCompressedHelpInfo)
}

func NewQCompressedHelpInfo3(other QCompressedHelpInfo_ITF) *QCompressedHelpInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQCompressedHelpInfo3", "", other}).(*QCompressedHelpInfo)
}

func (ptr *QCompressedHelpInfo) Component() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Component"}).(string)
}

func QCompressedHelpInfo_FromCompressedHelpFile(documentationFileName string) *QCompressedHelpInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QCompressedHelpInfo_FromCompressedHelpFile", "", documentationFileName}).(*QCompressedHelpInfo)
}

func (ptr *QCompressedHelpInfo) FromCompressedHelpFile(documentationFileName string) *QCompressedHelpInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QCompressedHelpInfo_FromCompressedHelpFile", "", documentationFileName}).(*QCompressedHelpInfo)
}

func (ptr *QCompressedHelpInfo) NamespaceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceName"}).(string)
}

func (ptr *QCompressedHelpInfo) Swap(other QCompressedHelpInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QCompressedHelpInfo) Version() *core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Version"}).(*core.QVersionNumber)
}

func (ptr *QCompressedHelpInfo) DestroyQCompressedHelpInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCompressedHelpInfo"})
}

type QHelpContentItem struct {
	internal.Internal
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (ptr *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return ptr
}

func (ptr *QHelpContentItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHelpContentItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func (n *QHelpContentItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) (n *QHelpContentItem) {
	n = new(QHelpContentItem)
	n.InitFromInternal(uintptr(ptr), "help.QHelpContentItem")
	return
}
func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Child", row}).(*QHelpContentItem)
}

func (ptr *QHelpContentItem) ChildCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildCount"}).(float64))
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildPosition", child}).(float64))
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parent"}).(*QHelpContentItem)
}

func (ptr *QHelpContentItem) Row() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Row"}).(float64))
}

func (ptr *QHelpContentItem) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QHelpContentItem) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpContentItem"})
}

type QHelpContentModel struct {
	core.QAbstractItemModel
}

type QHelpContentModel_ITF interface {
	core.QAbstractItemModel_ITF
	QHelpContentModel_PTR() *QHelpContentModel
}

func (ptr *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return ptr
}

func (ptr *QHelpContentModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpContentModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemModel_PTR().SetPointer(p)
	}
}

func PointerFromQHelpContentModel(ptr QHelpContentModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModel_PTR().Pointer()
	}
	return nil
}

func (n *QHelpContentModel) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractItemModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpContentModel) ClassNameInternalF() string {
	return n.QAbstractItemModel_PTR().ClassNameInternalF()
}

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) (n *QHelpContentModel) {
	n = new(QHelpContentModel)
	n.InitFromInternal(uintptr(ptr), "help.QHelpContentModel")
	return
}
func (ptr *QHelpContentModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectColumnCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCount"})
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount", parent}).(float64))
}

func (ptr *QHelpContentModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountDefault", parent}).(float64))
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentItemAt", index}).(*QHelpContentItem)
}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentsCreated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentsCreated"})
}

func (ptr *QHelpContentModel) ContentsCreated() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsCreated"})
}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentsCreationStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentsCreationStarted"})
}

func (ptr *QHelpContentModel) ContentsCreationStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsCreationStarted"})
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateContents", customFilterName})
}

func (ptr *QHelpContentModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", index, role}).(*core.QVariant)
}

func (ptr *QHelpContentModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", index, role}).(*core.QVariant)
}

func (ptr *QHelpContentModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndex", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectIndex() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndex"})
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDefault", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCreatingContents"}).(bool)
}

func (ptr *QHelpContentModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectParent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParent"})
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parent", index}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentDefault", index}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectRowCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCount"})
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount", parent}).(float64))
}

func (ptr *QHelpContentModel) RowCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountDefault", parent}).(float64))
}

func (ptr *QHelpContentModel) ConnectDestroyQHelpContentModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpContentModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentModel) DisconnectDestroyQHelpContentModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpContentModel"})
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpContentModel"})
}

func (ptr *QHelpContentModel) DestroyQHelpContentModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpContentModelDefault"})
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_setList", i})
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_setList", i})
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QHelpContentModel) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QHelpContentModel) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __itemData_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_atList", v, i}).(*core.QVariant)
}

func (ptr *QHelpContentModel) __itemData_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_setList", key, i})
}

func (ptr *QHelpContentModel) __itemData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __itemData_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_keyList"}).([]int)
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_setList", i})
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QHelpContentModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_setList", i})
}

func (ptr *QHelpContentModel) __layoutChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __match_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) __match_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_setList", i})
}

func (ptr *QHelpContentModel) __match_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __mimeData_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_setList", i})
}

func (ptr *QHelpContentModel) __mimeData_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __persistentIndexList_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_setList", i})
}

func (ptr *QHelpContentModel) __persistentIndexList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __roleNames_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_atList", v, i}).(*core.QByteArray)
}

func (ptr *QHelpContentModel) __roleNames_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_setList", key, i})
}

func (ptr *QHelpContentModel) __roleNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __roleNames_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_keyList"}).([]int)
}

func (ptr *QHelpContentModel) __setItemData_roles_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_atList", v, i}).(*core.QVariant)
}

func (ptr *QHelpContentModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_setList", key, i})
}

func (ptr *QHelpContentModel) __setItemData_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __setItemData_roles_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_keyList"}).([]int)
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) ____itemData_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_atList", i}).(float64))
}

func (ptr *QHelpContentModel) ____itemData_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_setList", i})
}

func (ptr *QHelpContentModel) ____itemData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) ____roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpContentModel) ____roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_setList", i})
}

func (ptr *QHelpContentModel) ____roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_atList", i}).(float64))
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_setList", i})
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpContentModel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpContentModel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpContentModel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpContentModel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpContentModel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpContentModel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpContentModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BuddyDefault", index}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanDropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QHelpContentModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanFetchMoreDefault", parent}).(bool)
}

func (ptr *QHelpContentModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchMoreDefault", parent})
}

func (ptr *QHelpContentModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {

	return core.Qt__ItemFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault", index}).(float64))
}

func (ptr *QHelpContentModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildrenDefault", parent}).(bool)
}

func (ptr *QHelpContentModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeaderDataDefault", section, orientation, role}).(*core.QVariant)
}

func (ptr *QHelpContentModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QHelpContentModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowsDefault", row, count, parent}).(bool)
}

func (ptr *QHelpContentModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemDataDefault", index}).(map[int]*core.QVariant)
}

func (ptr *QHelpContentModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatchDefault", start, role, value, hits, flags}).([]*core.QModelIndex)
}

func (ptr *QHelpContentModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeDataDefault", indexes}).(*core.QMimeData)
}

func (ptr *QHelpContentModel) MimeTypesDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeTypesDefault"}).([]string)
}

func (ptr *QHelpContentModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveColumnsDefault", sourceParent, sourceColumn, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QHelpContentModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveRowsDefault", sourceParent, sourceRow, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QHelpContentModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QHelpContentModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRowsDefault", row, count, parent}).(bool)
}

func (ptr *QHelpContentModel) ResetInternalDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetInternalDataDefault"})
}

func (ptr *QHelpContentModel) RevertDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertDefault"})
}

func (ptr *QHelpContentModel) RoleNamesDefault() map[int]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoleNamesDefault"}).(map[int]*core.QByteArray)
}

func (ptr *QHelpContentModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", index, value, role}).(bool)
}

func (ptr *QHelpContentModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeaderDataDefault", section, orientation, value, role}).(bool)
}

func (ptr *QHelpContentModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemDataDefault", index, roles}).(bool)
}

func (ptr *QHelpContentModel) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SiblingDefault", row, column, index}).(*core.QModelIndex)
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortDefault", column, order})
}

func (ptr *QHelpContentModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpanDefault", index}).(*core.QSize)
}

func (ptr *QHelpContentModel) SubmitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitDefault"}).(bool)
}

func (ptr *QHelpContentModel) SupportedDragActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDragActionsDefault"}).(float64))
}

func (ptr *QHelpContentModel) SupportedDropActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDropActionsDefault"}).(float64))
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpContentModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpContentModel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpContentModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpContentModel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpContentModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpContentModel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func (ptr *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return ptr
}

func (ptr *QHelpContentWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeView_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpContentWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTreeView_PTR().SetPointer(p)
	}
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func (n *QHelpContentWidget) InitFromInternal(ptr uintptr, name string) {
	n.QTreeView_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpContentWidget) ClassNameInternalF() string {
	return n.QTreeView_PTR().ClassNameInternalF()
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpContentWidget) {
	n = new(QHelpContentWidget)
	n.InitFromInternal(uintptr(ptr), "help.QHelpContentWidget")
	return
}

func (ptr *QHelpContentWidget) DestroyQHelpContentWidget() {
}

func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexOf", link}).(*core.QModelIndex)
}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLinkActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLinkActivated"})
}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinkActivated", link})
}

func (ptr *QHelpContentWidget) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QHelpContentWidget) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QHelpContentWidget) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __selectedIndexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpContentWidget) __selectedIndexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_setList", i})
}

func (ptr *QHelpContentWidget) __selectedIndexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_atList(i int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_atList", i}).(*widgets.QWidget)
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_setList", i})
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpContentWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QHelpContentWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpContentWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QHelpContentWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpContentWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QHelpContentWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpContentWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpContentWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpContentWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpContentWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpContentWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpContentWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpContentWidget) CollapseDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollapseDefault", index})
}

func (ptr *QHelpContentWidget) CollapseAllDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollapseAllDefault"})
}

func (ptr *QHelpContentWidget) ColumnCountChangedDefault(oldCount int, newCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChangedDefault", oldCount, newCount})
}

func (ptr *QHelpContentWidget) ColumnMovedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnMovedDefault"})
}

func (ptr *QHelpContentWidget) ColumnResizedDefault(column int, oldSize int, newSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnResizedDefault", column, oldSize, newSize})
}

func (ptr *QHelpContentWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentChangedDefault", current, previous})
}

func (ptr *QHelpContentWidget) DataChangedDefault(topLeft core.QModelIndex_ITF, bottomRight core.QModelIndex_ITF, roles []int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataChangedDefault", topLeft, bottomRight, roles})
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QHelpContentWidget) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawBranchesDefault", painter, rect, index})
}

func (ptr *QHelpContentWidget) DrawRowDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawRowDefault", painter, option, index})
}

func (ptr *QHelpContentWidget) ExpandDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpandDefault", index})
}

func (ptr *QHelpContentWidget) ExpandAllDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpandAllDefault"})
}

func (ptr *QHelpContentWidget) ExpandRecursivelyDefault(index core.QModelIndex_ITF, depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpandRecursivelyDefault", index, depth})
}

func (ptr *QHelpContentWidget) ExpandToDepthDefault(depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpandToDepthDefault", depth})
}

func (ptr *QHelpContentWidget) HideColumnDefault(column int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideColumnDefault", column})
}

func (ptr *QHelpContentWidget) HorizontalOffsetDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HorizontalOffsetDefault"}).(float64))
}

func (ptr *QHelpContentWidget) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexAtDefault", point}).(*core.QModelIndex)
}

func (ptr *QHelpContentWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsIndexHiddenDefault", index}).(bool)
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyboardSearchDefault", search})
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QHelpContentWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveCursorDefault", cursorAction, modifiers}).(*core.QModelIndex)
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QHelpContentWidget) ResetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"})
}

func (ptr *QHelpContentWidget) ResizeColumnToContentsDefault(column int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeColumnToContentsDefault", column})
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsAboutToBeRemovedDefault", parent, start, end})
}

func (ptr *QHelpContentWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsInsertedDefault", parent, start, end})
}

func (ptr *QHelpContentWidget) RowsRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsRemovedDefault", parent, start, end})
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollContentsByDefault", dx, dy})
}

func (ptr *QHelpContentWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToDefault", index, hint})
}

func (ptr *QHelpContentWidget) SelectAllDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectAllDefault"})
}

func (ptr *QHelpContentWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionChangedDefault", selected, deselected})
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModelDefault", model})
}

func (ptr *QHelpContentWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRootIndexDefault", index})
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionDefault", rect, command})
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionModelDefault", selectionModel})
}

func (ptr *QHelpContentWidget) ShowColumnDefault(column int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowColumnDefault", column})
}

func (ptr *QHelpContentWidget) SizeHintForColumnDefault(column int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintForColumnDefault", column}).(float64))
}

func (ptr *QHelpContentWidget) SortByColumnDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortByColumnDefault", column, order})
}

func (ptr *QHelpContentWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateGeometriesDefault"})
}

func (ptr *QHelpContentWidget) VerticalOffsetDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VerticalOffsetDefault"}).(float64))
}

func (ptr *QHelpContentWidget) ViewportEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportEventDefault", event}).(bool)
}

func (ptr *QHelpContentWidget) ViewportSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpContentWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisualRectDefault", index}).(*core.QRect)
}

func (ptr *QHelpContentWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisualRegionForSelectionDefault", selection}).(*gui.QRegion)
}

func (ptr *QHelpContentWidget) ClearSelectionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearSelectionDefault"})
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEditorDefault", editor, hint})
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CommitDataDefault", editor})
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QHelpContentWidget) EditDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditDefault", index})
}

func (ptr *QHelpContentWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Edit2Default", index, trigger, event}).(bool)
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditorDestroyedDefault", editor})
}

func (ptr *QHelpContentWidget) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QHelpContentWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", object, event}).(bool)
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QHelpContentWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QHelpContentWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QHelpContentWidget) ScrollToBottomDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToBottomDefault"})
}

func (ptr *QHelpContentWidget) ScrollToTopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToTopDefault"})
}

func (ptr *QHelpContentWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {

	return core.QItemSelectionModel__SelectionFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionCommandDefault", index, event}).(float64))
}

func (ptr *QHelpContentWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentIndexDefault", index})
}

func (ptr *QHelpContentWidget) SizeHintForRowDefault(row int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintForRowDefault", row}).(float64))
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDragDefault", supportedActions})
}

func (ptr *QHelpContentWidget) UpdateDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault", index})
}

func (ptr *QHelpContentWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewOptionsDefault"}).(*widgets.QStyleOptionViewItem)
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", e})
}

func (ptr *QHelpContentWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupViewportDefault", viewport})
}

func (ptr *QHelpContentWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", e})
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", ev})
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QHelpContentWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QHelpContentWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QHelpContentWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QHelpContentWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QHelpContentWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QHelpContentWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QHelpContentWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QHelpContentWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QHelpContentWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QHelpContentWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QHelpContentWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QHelpContentWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QHelpContentWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QHelpContentWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QHelpContentWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QHelpContentWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QHelpContentWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QHelpContentWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QHelpContentWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QHelpContentWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QHelpContentWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QHelpContentWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QHelpContentWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QHelpContentWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpContentWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpContentWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpContentWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpContentWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func (ptr *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return ptr
}

func (ptr *QHelpEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QHelpEngineCore_PTR().SetPointer(p)
	}
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func (n *QHelpEngine) InitFromInternal(ptr uintptr, name string) {
	n.QHelpEngineCore_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpEngine) ClassNameInternalF() string {
	return n.QHelpEngineCore_PTR().ClassNameInternalF()
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) (n *QHelpEngine) {
	n = new(QHelpEngine)
	n.InitFromInternal(uintptr(ptr), "help.QHelpEngine")
	return
}
func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpEngine", "", collectionFile, parent}).(*QHelpEngine)
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentModel"}).(*QHelpContentModel)
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentWidget"}).(*QHelpContentWidget)
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexModel"}).(*QHelpIndexModel)
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexWidget"}).(*QHelpIndexWidget)
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchEngine"}).(*QHelpSearchEngine)
}

func (ptr *QHelpEngine) ConnectDestroyQHelpEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpEngine) DisconnectDestroyQHelpEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpEngine"})
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpEngine"})
}

func (ptr *QHelpEngine) DestroyQHelpEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpEngineDefault"})
}

type QHelpEngineCore struct {
	core.QObject
}

type QHelpEngineCore_ITF interface {
	core.QObject_ITF
	QHelpEngineCore_PTR() *QHelpEngineCore
}

func (ptr *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return ptr
}

func (ptr *QHelpEngineCore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpEngineCore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func (n *QHelpEngineCore) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpEngineCore) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) (n *QHelpEngineCore) {
	n = new(QHelpEngineCore)
	n.InitFromInternal(uintptr(ptr), "help.QHelpEngineCore")
	return
}
func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpEngineCore", "", collectionFile, parent}).(*QHelpEngineCore)
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoSaveFilter"}).(bool)
}

func (ptr *QHelpEngineCore) CollectionFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollectionFile"}).(string)
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyCollectionFile", fileName}).(bool)
}

func (ptr *QHelpEngineCore) CurrentFilter() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentFilter"}).(string)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomValue", key, defaultValue}).(*core.QVariant)
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentationFileName", namespaceName}).(string)
}

func (ptr *QHelpEngineCore) Error() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(string)
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileData", url}).(*core.QByteArray)
}

func (ptr *QHelpEngineCore) Files2(namespaceName string, filterName string, extensionFilter string) []*core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Files2", namespaceName, filterName, extensionFilter}).([]*core.QUrl)
}

func (ptr *QHelpEngineCore) FilterEngine() *QHelpFilterEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterEngine"}).(*QHelpFilterEngine)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FindFile", url}).(*core.QUrl)
}

func (ptr *QHelpEngineCore) LinksForIdentifier(id string) map[string]*core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinksForIdentifier", id}).(map[string]*core.QUrl)
}

func (ptr *QHelpEngineCore) LinksForKeyword(keyword string) map[string]*core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinksForKeyword", keyword}).(map[string]*core.QUrl)
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QHelpEngineCore_MetaData", "", documentationFileName, name}).(*core.QVariant)
}

func (ptr *QHelpEngineCore) MetaData(documentationFileName string, name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QHelpEngineCore_MetaData", "", documentationFileName, name}).(*core.QVariant)
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QHelpEngineCore_NamespaceName", "", documentationFileName}).(string)
}

func (ptr *QHelpEngineCore) NamespaceName(documentationFileName string) string {

	return internal.CallLocalFunction([]interface{}{"", "", "help.QHelpEngineCore_NamespaceName", "", documentationFileName}).(string)
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterDocumentation", documentationFileName}).(bool)
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisteredDocumentations"}).([]string)
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveCustomValue", key}).(bool)
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoSaveFilter", save})
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCollectionFile", fileName})
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentFilter", filterName})
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCustomValue", key, value}).(bool)
}

func (ptr *QHelpEngineCore) SetUsesFilterEngine(uses bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUsesFilterEngine", uses})
}

func (ptr *QHelpEngineCore) SetupData() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupData"}).(bool)
}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetupFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetupFinished"})
}

func (ptr *QHelpEngineCore) SetupFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupFinished"})
}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetupStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetupStarted"})
}

func (ptr *QHelpEngineCore) SetupStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupStarted"})
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterDocumentation", namespaceName}).(bool)
}

func (ptr *QHelpEngineCore) UsesFilterEngine() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UsesFilterEngine"}).(bool)
}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWarning", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpEngineCore) DisconnectWarning() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWarning"})
}

func (ptr *QHelpEngineCore) Warning(msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Warning", msg})
}

func (ptr *QHelpEngineCore) ConnectDestroyQHelpEngineCore(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpEngineCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpEngineCore) DisconnectDestroyQHelpEngineCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpEngineCore"})
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpEngineCore"})
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCoreDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpEngineCoreDefault"})
}

func (ptr *QHelpEngineCore) __files_atList(i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_atList", i}).(*core.QUrl)
}

func (ptr *QHelpEngineCore) __files_setList(i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_setList", i})
}

func (ptr *QHelpEngineCore) __files_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __files_atList2(i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_atList2", i}).(*core.QUrl)
}

func (ptr *QHelpEngineCore) __files_setList2(i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_setList2", i})
}

func (ptr *QHelpEngineCore) __files_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__files_newList2"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __filterAttributeSets_atList(i int) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filterAttributeSets_atList", i}).([]string)
}

func (ptr *QHelpEngineCore) __filterAttributeSets_setList(i []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filterAttributeSets_setList", i})
}

func (ptr *QHelpEngineCore) __filterAttributeSets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filterAttributeSets_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __linksForIdentifier_atList(v string, i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForIdentifier_atList", v, i}).(*core.QUrl)
}

func (ptr *QHelpEngineCore) __linksForIdentifier_setList(key string, i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForIdentifier_setList", key, i})
}

func (ptr *QHelpEngineCore) __linksForIdentifier_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForIdentifier_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __linksForIdentifier_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForIdentifier_keyList"}).([]string)
}

func (ptr *QHelpEngineCore) __linksForKeyword_atList(v string, i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_atList", v, i}).(*core.QUrl)
}

func (ptr *QHelpEngineCore) __linksForKeyword_setList(key string, i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_setList", key, i})
}

func (ptr *QHelpEngineCore) __linksForKeyword_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __linksForKeyword_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_keyList"}).([]string)
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForIdentifier_keyList_atList", i}).(string)
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForIdentifier_keyList_setList", i})
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForIdentifier_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_atList", i}).(string)
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_setList", i})
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpEngineCore) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpEngineCore) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpEngineCore) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpEngineCore) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpEngineCore) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpEngineCore) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpEngineCore) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpEngineCore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpEngineCore) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpEngineCore) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpEngineCore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpEngineCore) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpEngineCore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpEngineCore) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpFilterData struct {
	internal.Internal
}

type QHelpFilterData_ITF interface {
	QHelpFilterData_PTR() *QHelpFilterData
}

func (ptr *QHelpFilterData) QHelpFilterData_PTR() *QHelpFilterData {
	return ptr
}

func (ptr *QHelpFilterData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHelpFilterData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHelpFilterData(ptr QHelpFilterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpFilterData_PTR().Pointer()
	}
	return nil
}

func (n *QHelpFilterData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHelpFilterDataFromPointer(ptr unsafe.Pointer) (n *QHelpFilterData) {
	n = new(QHelpFilterData)
	n.InitFromInternal(uintptr(ptr), "help.QHelpFilterData")
	return
}
func NewQHelpFilterData() *QHelpFilterData {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpFilterData", ""}).(*QHelpFilterData)
}

func NewQHelpFilterData2(other QHelpFilterData_ITF) *QHelpFilterData {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpFilterData2", "", other}).(*QHelpFilterData)
}

func NewQHelpFilterData3(other QHelpFilterData_ITF) *QHelpFilterData {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpFilterData3", "", other}).(*QHelpFilterData)
}

func (ptr *QHelpFilterData) Components() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Components"}).([]string)
}

func (ptr *QHelpFilterData) SetComponents(components []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetComponents", components})
}

func (ptr *QHelpFilterData) SetVersions(versions []*core.QVersionNumber) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVersions", versions})
}

func (ptr *QHelpFilterData) Versions() []*core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Versions"}).([]*core.QVersionNumber)
}

func (ptr *QHelpFilterData) DestroyQHelpFilterData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpFilterData"})
}

func (ptr *QHelpFilterData) __setVersions_versions_atList(i int) *core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVersions_versions_atList", i}).(*core.QVersionNumber)
}

func (ptr *QHelpFilterData) __setVersions_versions_setList(i core.QVersionNumber_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVersions_versions_setList", i})
}

func (ptr *QHelpFilterData) __setVersions_versions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVersions_versions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterData) __versions_atList(i int) *core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__versions_atList", i}).(*core.QVersionNumber)
}

func (ptr *QHelpFilterData) __versions_setList(i core.QVersionNumber_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__versions_setList", i})
}

func (ptr *QHelpFilterData) __versions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__versions_newList"}).(unsafe.Pointer)
}

type QHelpFilterEngine struct {
	core.QObject
}

type QHelpFilterEngine_ITF interface {
	core.QObject_ITF
	QHelpFilterEngine_PTR() *QHelpFilterEngine
}

func (ptr *QHelpFilterEngine) QHelpFilterEngine_PTR() *QHelpFilterEngine {
	return ptr
}

func (ptr *QHelpFilterEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpFilterEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHelpFilterEngine(ptr QHelpFilterEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpFilterEngine_PTR().Pointer()
	}
	return nil
}

func (n *QHelpFilterEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpFilterEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHelpFilterEngineFromPointer(ptr unsafe.Pointer) (n *QHelpFilterEngine) {
	n = new(QHelpFilterEngine)
	n.InitFromInternal(uintptr(ptr), "help.QHelpFilterEngine")
	return
}

func (ptr *QHelpFilterEngine) DestroyQHelpFilterEngine() {
}

func (ptr *QHelpFilterEngine) ActiveFilter() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveFilter"}).(string)
}

func (ptr *QHelpFilterEngine) AvailableComponents() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableComponents"}).([]string)
}

func (ptr *QHelpFilterEngine) ConnectFilterActivated(f func(newFilter string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilterActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpFilterEngine) DisconnectFilterActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilterActivated"})
}

func (ptr *QHelpFilterEngine) FilterActivated(newFilter string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterActivated", newFilter})
}

func (ptr *QHelpFilterEngine) FilterData(filterName string) *QHelpFilterData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterData", filterName}).(*QHelpFilterData)
}

func (ptr *QHelpFilterEngine) Filters() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filters"}).([]string)
}

func (ptr *QHelpFilterEngine) NamespaceToComponent() map[string]string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceToComponent"}).(map[string]string)
}

func (ptr *QHelpFilterEngine) NamespaceToVersion() map[string]*core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceToVersion"}).(map[string]*core.QVersionNumber)
}

func (ptr *QHelpFilterEngine) NamespacesForFilter(filterName string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespacesForFilter", filterName}).([]string)
}

func (ptr *QHelpFilterEngine) RemoveFilter(filterName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveFilter", filterName}).(bool)
}

func (ptr *QHelpFilterEngine) SetActiveFilter(filterName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveFilter", filterName}).(bool)
}

func (ptr *QHelpFilterEngine) SetFilterData(filterName string, filterData QHelpFilterData_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFilterData", filterName, filterData}).(bool)
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_atList(v string, i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToComponent_atList", v, i}).(string)
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_setList(key string, i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToComponent_setList", key, i})
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToComponent_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToComponent_keyList"}).([]string)
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_atList(v string, i int) *core.QVersionNumber {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToVersion_atList", v, i}).(*core.QVersionNumber)
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_setList(key string, i core.QVersionNumber_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToVersion_setList", key, i})
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToVersion_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceToVersion_keyList"}).([]string)
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToComponent_keyList_atList", i}).(string)
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToComponent_keyList_setList", i})
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToComponent_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToVersion_keyList_atList", i}).(string)
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToVersion_keyList_setList", i})
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____namespaceToVersion_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpFilterEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpFilterEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpFilterEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpFilterEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpFilterEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpFilterEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpFilterEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpFilterEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpFilterEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpFilterEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpFilterEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpFilterEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpFilterEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpFilterEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpFilterEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpGlobal struct {
	internal.Internal
}

type QHelpGlobal_ITF interface {
	QHelpGlobal_PTR() *QHelpGlobal
}

func (ptr *QHelpGlobal) QHelpGlobal_PTR() *QHelpGlobal {
	return ptr
}

func (ptr *QHelpGlobal) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHelpGlobal) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHelpGlobal(ptr QHelpGlobal_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpGlobal_PTR().Pointer()
	}
	return nil
}

func (n *QHelpGlobal) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHelpGlobalFromPointer(ptr unsafe.Pointer) (n *QHelpGlobal) {
	n = new(QHelpGlobal)
	n.InitFromInternal(uintptr(ptr), "help.QHelpGlobal")
	return
}

func (ptr *QHelpGlobal) DestroyQHelpGlobal() {
}

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModel_ITF interface {
	core.QStringListModel_ITF
	QHelpIndexModel_PTR() *QHelpIndexModel
}

func (ptr *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return ptr
}

func (ptr *QHelpIndexModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpIndexModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QStringListModel_PTR().SetPointer(p)
	}
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModel_PTR().Pointer()
	}
	return nil
}

func (n *QHelpIndexModel) InitFromInternal(ptr uintptr, name string) {
	n.QStringListModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpIndexModel) ClassNameInternalF() string {
	return n.QStringListModel_PTR().ClassNameInternalF()
}

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) (n *QHelpIndexModel) {
	n = new(QHelpIndexModel)
	n.InitFromInternal(uintptr(ptr), "help.QHelpIndexModel")
	return
}

func (ptr *QHelpIndexModel) DestroyQHelpIndexModel() {
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateIndex", customFilterName})
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", filter, wildcard}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexCreated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexCreated"})
}

func (ptr *QHelpIndexModel) IndexCreated() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexCreated"})
}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexCreationStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexCreationStarted"})
}

func (ptr *QHelpIndexModel) IndexCreationStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexCreationStarted"})
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCreatingIndex"}).(bool)
}

func (ptr *QHelpIndexModel) __linksForKeyword_atList(v string, i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_atList", v, i}).(*core.QUrl)
}

func (ptr *QHelpIndexModel) __linksForKeyword_setList(key string, i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_setList", key, i})
}

func (ptr *QHelpIndexModel) __linksForKeyword_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __linksForKeyword_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksForKeyword_keyList"}).([]string)
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_atList", i}).(string)
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksForKeyword_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __itemData_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_atList", v, i}).(*core.QVariant)
}

func (ptr *QHelpIndexModel) __itemData_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_setList", key, i})
}

func (ptr *QHelpIndexModel) __itemData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __itemData_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_keyList"}).([]int)
}

func (ptr *QHelpIndexModel) __setItemData_roles_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_atList", v, i}).(*core.QVariant)
}

func (ptr *QHelpIndexModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_setList", key, i})
}

func (ptr *QHelpIndexModel) __setItemData_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __setItemData_roles_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_keyList"}).([]int)
}

func (ptr *QHelpIndexModel) ____itemData_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) ____itemData_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____itemData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_setList", i})
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_setList", i})
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QHelpIndexModel) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_setList", i})
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_setList", i})
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __match_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) __match_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_setList", i})
}

func (ptr *QHelpIndexModel) __match_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __mimeData_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_setList", i})
}

func (ptr *QHelpIndexModel) __mimeData_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __persistentIndexList_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_setList", i})
}

func (ptr *QHelpIndexModel) __persistentIndexList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __roleNames_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_atList", v, i}).(*core.QByteArray)
}

func (ptr *QHelpIndexModel) __roleNames_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_setList", key, i})
}

func (ptr *QHelpIndexModel) __roleNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __roleNames_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_keyList"}).([]int)
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpIndexModel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpIndexModel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpIndexModel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpIndexModel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpIndexModel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpIndexModel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", index, role}).(*core.QVariant)
}

func (ptr *QHelpIndexModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {

	return core.Qt__ItemFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault", index}).(float64))
}

func (ptr *QHelpIndexModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowsDefault", row, count, parent}).(bool)
}

func (ptr *QHelpIndexModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemDataDefault", index}).(map[int]*core.QVariant)
}

func (ptr *QHelpIndexModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveRowsDefault", sourceParent, sourceRow, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QHelpIndexModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRowsDefault", row, count, parent}).(bool)
}

func (ptr *QHelpIndexModel) RowCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountDefault", parent}).(float64))
}

func (ptr *QHelpIndexModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", index, value, role}).(bool)
}

func (ptr *QHelpIndexModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemDataDefault", index, roles}).(bool)
}

func (ptr *QHelpIndexModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SiblingDefault", row, column, idx}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortDefault", column, order})
}

func (ptr *QHelpIndexModel) SupportedDropActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDropActionsDefault"}).(float64))
}

func (ptr *QHelpIndexModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QHelpIndexModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDefault", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BuddyDefault", index}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanDropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QHelpIndexModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanFetchMoreDefault", parent}).(bool)
}

func (ptr *QHelpIndexModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountDefault", parent}).(float64))
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchMoreDefault", parent})
}

func (ptr *QHelpIndexModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildrenDefault", parent}).(bool)
}

func (ptr *QHelpIndexModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeaderDataDefault", section, orientation, role}).(*core.QVariant)
}

func (ptr *QHelpIndexModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QHelpIndexModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatchDefault", start, role, value, hits, flags}).([]*core.QModelIndex)
}

func (ptr *QHelpIndexModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeDataDefault", indexes}).(*core.QMimeData)
}

func (ptr *QHelpIndexModel) MimeTypesDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeTypesDefault"}).([]string)
}

func (ptr *QHelpIndexModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveColumnsDefault", sourceParent, sourceColumn, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QHelpIndexModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentDefault", index}).(*core.QModelIndex)
}

func (ptr *QHelpIndexModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QHelpIndexModel) ResetInternalDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetInternalDataDefault"})
}

func (ptr *QHelpIndexModel) RevertDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertDefault"})
}

func (ptr *QHelpIndexModel) RoleNamesDefault() map[int]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoleNamesDefault"}).(map[int]*core.QByteArray)
}

func (ptr *QHelpIndexModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeaderDataDefault", section, orientation, value, role}).(bool)
}

func (ptr *QHelpIndexModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpanDefault", index}).(*core.QSize)
}

func (ptr *QHelpIndexModel) SubmitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitDefault"}).(bool)
}

func (ptr *QHelpIndexModel) SupportedDragActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDragActionsDefault"}).(float64))
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpIndexModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpIndexModel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpIndexModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpIndexModel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpIndexModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpIndexModel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func (ptr *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return ptr
}

func (ptr *QHelpIndexWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QListView_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpIndexWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QListView_PTR().SetPointer(p)
	}
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func (n *QHelpIndexWidget) InitFromInternal(ptr uintptr, name string) {
	n.QListView_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpIndexWidget) ClassNameInternalF() string {
	return n.QListView_PTR().ClassNameInternalF()
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpIndexWidget) {
	n = new(QHelpIndexWidget)
	n.InitFromInternal(uintptr(ptr), "help.QHelpIndexWidget")
	return
}

func (ptr *QHelpIndexWidget) DestroyQHelpIndexWidget() {
}

func (ptr *QHelpIndexWidget) ConnectActivateCurrentItem(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActivateCurrentItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexWidget) DisconnectActivateCurrentItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActivateCurrentItem"})
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivateCurrentItem"})
}

func (ptr *QHelpIndexWidget) ActivateCurrentItemDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivateCurrentItemDefault"})
}

func (ptr *QHelpIndexWidget) ConnectFilterIndices(f func(filter string, wildcard string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilterIndices", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexWidget) DisconnectFilterIndices() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilterIndices"})
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterIndices", filter, wildcard})
}

func (ptr *QHelpIndexWidget) FilterIndicesDefault(filter string, wildcard string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterIndicesDefault", filter, wildcard})
}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLinkActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLinkActivated"})
}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinkActivated", link, keyword})
}

func (ptr *QHelpIndexWidget) ConnectLinksActivated(f func(links map[string]*core.QUrl, keyword string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLinksActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpIndexWidget) DisconnectLinksActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLinksActivated"})
}

func (ptr *QHelpIndexWidget) LinksActivated(links map[string]*core.QUrl, keyword string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinksActivated", links, keyword})
}

func (ptr *QHelpIndexWidget) __linksActivated_links_atList(v string, i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksActivated_links_atList", v, i}).(*core.QUrl)
}

func (ptr *QHelpIndexWidget) __linksActivated_links_setList(key string, i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksActivated_links_setList", key, i})
}

func (ptr *QHelpIndexWidget) __linksActivated_links_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksActivated_links_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __linksActivated_links_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__linksActivated_links_keyList"}).([]string)
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksActivated_links_keyList_atList", i}).(string)
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksActivated_links_keyList_setList", i})
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____linksActivated_links_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__indexesMoved_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__indexesMoved_indexes_setList", i})
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__indexesMoved_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __selectedIndexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QHelpIndexWidget) __selectedIndexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_setList", i})
}

func (ptr *QHelpIndexWidget) __selectedIndexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectedIndexes_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_atList(i int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_atList", i}).(*widgets.QWidget)
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_setList", i})
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpIndexWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QHelpIndexWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpIndexWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QHelpIndexWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpIndexWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QHelpIndexWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpIndexWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpIndexWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpIndexWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpIndexWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpIndexWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpIndexWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpIndexWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentChangedDefault", current, previous})
}

func (ptr *QHelpIndexWidget) DataChangedDefault(topLeft core.QModelIndex_ITF, bottomRight core.QModelIndex_ITF, roles []int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataChangedDefault", topLeft, bottomRight, roles})
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", e})
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", e})
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", e})
}

func (ptr *QHelpIndexWidget) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpIndexWidget) HorizontalOffsetDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HorizontalOffsetDefault"}).(float64))
}

func (ptr *QHelpIndexWidget) IndexAtDefault(p core.QPoint_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexAtDefault", p}).(*core.QModelIndex)
}

func (ptr *QHelpIndexWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsIndexHiddenDefault", index}).(bool)
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", e})
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", e})
}

func (ptr *QHelpIndexWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveCursorDefault", cursorAction, modifiers}).(*core.QModelIndex)
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", e})
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", e})
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsAboutToBeRemovedDefault", parent, start, end})
}

func (ptr *QHelpIndexWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsInsertedDefault", parent, start, end})
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToDefault", index, hint})
}

func (ptr *QHelpIndexWidget) SelectedIndexesDefault() []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedIndexesDefault"}).([]*core.QModelIndex)
}

func (ptr *QHelpIndexWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionChangedDefault", selected, deselected})
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionDefault", rect, command})
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDragDefault", supportedActions})
}

func (ptr *QHelpIndexWidget) TimerEventDefault(e core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", e})
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateGeometriesDefault"})
}

func (ptr *QHelpIndexWidget) VerticalOffsetDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VerticalOffsetDefault"}).(float64))
}

func (ptr *QHelpIndexWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewOptionsDefault"}).(*widgets.QStyleOptionViewItem)
}

func (ptr *QHelpIndexWidget) ViewportSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpIndexWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisualRectDefault", index}).(*core.QRect)
}

func (ptr *QHelpIndexWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisualRegionForSelectionDefault", selection}).(*gui.QRegion)
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", e})
}

func (ptr *QHelpIndexWidget) ClearSelectionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearSelectionDefault"})
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEditorDefault", editor, hint})
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CommitDataDefault", editor})
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QHelpIndexWidget) EditDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditDefault", index})
}

func (ptr *QHelpIndexWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Edit2Default", index, trigger, event}).(bool)
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditorDestroyedDefault", editor})
}

func (ptr *QHelpIndexWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", object, event}).(bool)
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QHelpIndexWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QHelpIndexWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyboardSearchDefault", search})
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QHelpIndexWidget) ResetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"})
}

func (ptr *QHelpIndexWidget) ScrollToBottomDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToBottomDefault"})
}

func (ptr *QHelpIndexWidget) ScrollToTopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollToTopDefault"})
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectAllDefault"})
}

func (ptr *QHelpIndexWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {

	return core.QItemSelectionModel__SelectionFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionCommandDefault", index, event}).(float64))
}

func (ptr *QHelpIndexWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentIndexDefault", index})
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModelDefault", model})
}

func (ptr *QHelpIndexWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRootIndexDefault", index})
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionModelDefault", selectionModel})
}

func (ptr *QHelpIndexWidget) SizeHintForColumnDefault(column int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintForColumnDefault", column}).(float64))
}

func (ptr *QHelpIndexWidget) SizeHintForRowDefault(row int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintForRowDefault", row}).(float64))
}

func (ptr *QHelpIndexWidget) UpdateDefault(index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault", index})
}

func (ptr *QHelpIndexWidget) ViewportEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportEventDefault", event}).(bool)
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", e})
}

func (ptr *QHelpIndexWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollContentsByDefault", dx, dy})
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupViewportDefault", viewport})
}

func (ptr *QHelpIndexWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", ev})
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QHelpIndexWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QHelpIndexWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QHelpIndexWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QHelpIndexWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QHelpIndexWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QHelpIndexWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QHelpIndexWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QHelpIndexWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QHelpIndexWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QHelpIndexWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QHelpIndexWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QHelpIndexWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QHelpIndexWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QHelpIndexWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QHelpIndexWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QHelpIndexWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QHelpIndexWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QHelpIndexWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QHelpIndexWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QHelpIndexWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QHelpIndexWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QHelpIndexWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QHelpIndexWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QHelpIndexWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpIndexWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpIndexWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpIndexWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpIndexWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngine_ITF interface {
	core.QObject_ITF
	QHelpSearchEngine_PTR() *QHelpSearchEngine
}

func (ptr *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return ptr
}

func (ptr *QHelpSearchEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpSearchEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEngine_PTR().Pointer()
	}
	return nil
}

func (n *QHelpSearchEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpSearchEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) (n *QHelpSearchEngine) {
	n = new(QHelpSearchEngine)
	n.InitFromInternal(uintptr(ptr), "help.QHelpSearchEngine")
	return
}
func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchEngine", "", helpEngine, parent}).(*QHelpSearchEngine)
}

func (ptr *QHelpSearchEngine) ConnectCancelIndexing(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCancelIndexing", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectCancelIndexing() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCancelIndexing"})
}

func (ptr *QHelpSearchEngine) CancelIndexing() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelIndexing"})
}

func (ptr *QHelpSearchEngine) CancelIndexingDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelIndexingDefault"})
}

func (ptr *QHelpSearchEngine) ConnectCancelSearching(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCancelSearching", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectCancelSearching() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCancelSearching"})
}

func (ptr *QHelpSearchEngine) CancelSearching() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelSearching"})
}

func (ptr *QHelpSearchEngine) CancelSearchingDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelSearchingDefault"})
}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexingFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexingFinished"})
}

func (ptr *QHelpSearchEngine) IndexingFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexingFinished"})
}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexingStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexingStarted"})
}

func (ptr *QHelpSearchEngine) IndexingStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexingStarted"})
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryWidget"}).(*QHelpSearchQueryWidget)
}

func (ptr *QHelpSearchEngine) ConnectReindexDocumentation(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReindexDocumentation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectReindexDocumentation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReindexDocumentation"})
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReindexDocumentation"})
}

func (ptr *QHelpSearchEngine) ReindexDocumentationDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReindexDocumentationDefault"})
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResultWidget"}).(*QHelpSearchResultWidget)
}

func (ptr *QHelpSearchEngine) ConnectSearch2(f func(searchInput string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSearch2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectSearch2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSearch2"})
}

func (ptr *QHelpSearchEngine) Search2(searchInput string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Search2", searchInput})
}

func (ptr *QHelpSearchEngine) Search2Default(searchInput string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Search2Default", searchInput})
}

func (ptr *QHelpSearchEngine) SearchInput() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchInput"}).(string)
}

func (ptr *QHelpSearchEngine) SearchResultCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchResultCount"}).(float64))
}

func (ptr *QHelpSearchEngine) SearchResults(start int, end int) []*QHelpSearchResult {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchResults", start, end}).([]*QHelpSearchResult)
}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(searchResultCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSearchingFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSearchingFinished"})
}

func (ptr *QHelpSearchEngine) SearchingFinished(searchResultCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchingFinished", searchResultCount})
}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSearchingStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSearchingStarted"})
}

func (ptr *QHelpSearchEngine) SearchingStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchingStarted"})
}

func (ptr *QHelpSearchEngine) ConnectDestroyQHelpSearchEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpSearchEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchEngine) DisconnectDestroyQHelpSearchEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpSearchEngine"})
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchEngine"})
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchEngineDefault"})
}

func (ptr *QHelpSearchEngine) __query_atList(i int) *QHelpSearchQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_atList", i}).(*QHelpSearchQuery)
}

func (ptr *QHelpSearchEngine) __query_setList(i QHelpSearchQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_setList", i})
}

func (ptr *QHelpSearchEngine) __query_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __search_queryList_atList(i int) *QHelpSearchQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__search_queryList_atList", i}).(*QHelpSearchQuery)
}

func (ptr *QHelpSearchEngine) __search_queryList_setList(i QHelpSearchQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__search_queryList_setList", i})
}

func (ptr *QHelpSearchEngine) __search_queryList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__search_queryList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __searchResults_atList(i int) *QHelpSearchResult {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__searchResults_atList", i}).(*QHelpSearchResult)
}

func (ptr *QHelpSearchEngine) __searchResults_setList(i QHelpSearchResult_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__searchResults_setList", i})
}

func (ptr *QHelpSearchEngine) __searchResults_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__searchResults_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpSearchEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpSearchEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpSearchEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpSearchEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpSearchEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpSearchEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpSearchEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpSearchEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpSearchEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHelpSearchEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpSearchEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpSearchEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpSearchQuery struct {
	internal.Internal
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (ptr *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return ptr
}

func (ptr *QHelpSearchQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHelpSearchQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func (n *QHelpSearchQuery) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) (n *QHelpSearchQuery) {
	n = new(QHelpSearchQuery)
	n.InitFromInternal(uintptr(ptr), "help.QHelpSearchQuery")
	return
}

func (ptr *QHelpSearchQuery) DestroyQHelpSearchQuery() {
}

func NewQHelpSearchQuery() *QHelpSearchQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchQuery", ""}).(*QHelpSearchQuery)
}

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func (n *QHelpSearchQueryWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpSearchQueryWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpSearchQueryWidget) {
	n = new(QHelpSearchQueryWidget)
	n.InitFromInternal(uintptr(ptr), "help.QHelpSearchQueryWidget")
	return
}
func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchQueryWidget", "", parent}).(*QHelpSearchQueryWidget)
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollapseExtendedSearch"})
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpandExtendedSearch"})
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSearch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSearch"})
}

func (ptr *QHelpSearchQueryWidget) Search() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Search"})
}

func (ptr *QHelpSearchQueryWidget) SearchInput() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SearchInput"}).(string)
}

func (ptr *QHelpSearchQueryWidget) SetSearchInput(searchInput string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSearchInput", searchInput})
}

func (ptr *QHelpSearchQueryWidget) ConnectDestroyQHelpSearchQueryWidget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpSearchQueryWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchQueryWidget) DisconnectDestroyQHelpSearchQueryWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpSearchQueryWidget"})
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchQueryWidget"})
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidgetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchQueryWidgetDefault"})
}

func (ptr *QHelpSearchQueryWidget) __query_atList(i int) *QHelpSearchQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_atList", i}).(*QHelpSearchQuery)
}

func (ptr *QHelpSearchQueryWidget) __query_setList(i QHelpSearchQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __query_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__query_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_atList(i int) *QHelpSearchQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setQuery_queryList_atList", i}).(*QHelpSearchQuery)
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_setList(i QHelpSearchQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setQuery_queryList_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setQuery_queryList_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchQueryWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchQueryWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchQueryWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpSearchQueryWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpSearchQueryWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpSearchQueryWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchQueryWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QHelpSearchQueryWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QHelpSearchQueryWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QHelpSearchQueryWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QHelpSearchQueryWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QHelpSearchQueryWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QHelpSearchQueryWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QHelpSearchQueryWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QHelpSearchQueryWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QHelpSearchQueryWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QHelpSearchQueryWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QHelpSearchQueryWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QHelpSearchQueryWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QHelpSearchQueryWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QHelpSearchQueryWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QHelpSearchQueryWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QHelpSearchQueryWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QHelpSearchQueryWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QHelpSearchQueryWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QHelpSearchQueryWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QHelpSearchQueryWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QHelpSearchQueryWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpSearchQueryWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QHelpSearchQueryWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpSearchQueryWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpSearchQueryWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpSearchQueryWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpSearchQueryWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpSearchQueryWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHelpSearchResult struct {
	internal.Internal
}

type QHelpSearchResult_ITF interface {
	QHelpSearchResult_PTR() *QHelpSearchResult
}

func (ptr *QHelpSearchResult) QHelpSearchResult_PTR() *QHelpSearchResult {
	return ptr
}

func (ptr *QHelpSearchResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHelpSearchResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHelpSearchResult(ptr QHelpSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResult_PTR().Pointer()
	}
	return nil
}

func (n *QHelpSearchResult) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHelpSearchResultFromPointer(ptr unsafe.Pointer) (n *QHelpSearchResult) {
	n = new(QHelpSearchResult)
	n.InitFromInternal(uintptr(ptr), "help.QHelpSearchResult")
	return
}
func NewQHelpSearchResult() *QHelpSearchResult {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchResult", ""}).(*QHelpSearchResult)
}

func NewQHelpSearchResult2(other QHelpSearchResult_ITF) *QHelpSearchResult {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchResult2", "", other}).(*QHelpSearchResult)
}

func NewQHelpSearchResult3(url core.QUrl_ITF, title string, snippet string) *QHelpSearchResult {

	return internal.CallLocalFunction([]interface{}{"", "", "help.NewQHelpSearchResult3", "", url, title, snippet}).(*QHelpSearchResult)
}

func (ptr *QHelpSearchResult) Snippet() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Snippet"}).(string)
}

func (ptr *QHelpSearchResult) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QHelpSearchResult) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QHelpSearchResult) DestroyQHelpSearchResult() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchResult"})
}

type QHelpSearchResultWidget struct {
	widgets.QWidget
}

type QHelpSearchResultWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func (n *QHelpSearchResultWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHelpSearchResultWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpSearchResultWidget) {
	n = new(QHelpSearchResultWidget)
	n.InitFromInternal(uintptr(ptr), "help.QHelpSearchResultWidget")
	return
}
func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinkAt", point}).(*core.QUrl)
}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestShowLink", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestShowLink"})
}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestShowLink", link})
}

func (ptr *QHelpSearchResultWidget) ConnectDestroyQHelpSearchResultWidget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHelpSearchResultWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHelpSearchResultWidget) DisconnectDestroyQHelpSearchResultWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHelpSearchResultWidget"})
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchResultWidget"})
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidgetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHelpSearchResultWidgetDefault"})
}

func (ptr *QHelpSearchResultWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchResultWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QHelpSearchResultWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchResultWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHelpSearchResultWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHelpSearchResultWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHelpSearchResultWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHelpSearchResultWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHelpSearchResultWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QHelpSearchResultWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QHelpSearchResultWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QHelpSearchResultWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QHelpSearchResultWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QHelpSearchResultWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QHelpSearchResultWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QHelpSearchResultWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QHelpSearchResultWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QHelpSearchResultWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QHelpSearchResultWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QHelpSearchResultWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QHelpSearchResultWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QHelpSearchResultWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QHelpSearchResultWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QHelpSearchResultWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QHelpSearchResultWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHelpSearchResultWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHelpSearchResultWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHelpSearchResultWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHelpSearchResultWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["help.QCompressedHelpInfo"] = NewQCompressedHelpInfoFromPointer
	internal.ConstructorTable["help.QHelpContentItem"] = NewQHelpContentItemFromPointer
	internal.ConstructorTable["help.QHelpContentModel"] = NewQHelpContentModelFromPointer
	internal.ConstructorTable["help.QHelpContentWidget"] = NewQHelpContentWidgetFromPointer
	internal.ConstructorTable["help.QHelpEngine"] = NewQHelpEngineFromPointer
	internal.ConstructorTable["help.QHelpEngineCore"] = NewQHelpEngineCoreFromPointer
	internal.ConstructorTable["help.QHelpFilterData"] = NewQHelpFilterDataFromPointer
	internal.ConstructorTable["help.QHelpFilterEngine"] = NewQHelpFilterEngineFromPointer
	internal.ConstructorTable["help.QHelpGlobal"] = NewQHelpGlobalFromPointer
	internal.ConstructorTable["help.QHelpIndexModel"] = NewQHelpIndexModelFromPointer
	internal.ConstructorTable["help.QHelpIndexWidget"] = NewQHelpIndexWidgetFromPointer
	internal.ConstructorTable["help.QHelpSearchEngine"] = NewQHelpSearchEngineFromPointer
	internal.ConstructorTable["help.QHelpSearchQuery"] = NewQHelpSearchQueryFromPointer
	internal.ConstructorTable["help.QHelpSearchQueryWidget"] = NewQHelpSearchQueryWidgetFromPointer
	internal.ConstructorTable["help.QHelpSearchResult"] = NewQHelpSearchResultFromPointer
	internal.ConstructorTable["help.QHelpSearchResultWidget"] = NewQHelpSearchResultWidgetFromPointer
}
