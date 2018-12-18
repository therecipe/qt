// +build !minimal

package help

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtHelp_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtHelp_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (ptr *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return ptr
}

func (ptr *QHelpContentItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHelpContentItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) (n *QHelpContentItem) {
	n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(int32(row))))
	}
	return nil
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpContentItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpContentItem) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpContentItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildPosition(ptr.Pointer(), PointerFromQHelpContentItem(child))))
	}
	return 0
}

func (ptr *QHelpContentItem) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_Row(ptr.Pointer())))
	}
	return 0
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

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) (n *QHelpContentModel) {
	n = new(QHelpContentModel)
	n.SetPointer(ptr)
	return
}
func QHelpContentModel_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentModel_QHelpContentModel_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpContentModel) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentModel_QHelpContentModel_Tr(sC, cC, C.int(int32(n))))
}

func QHelpContentModel_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentModel_QHelpContentModel_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpContentModel) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentModel_QHelpContentModel_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQHelpContentModel_ContentsCreated
func callbackQHelpContentModel_ContentsCreated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsCreated") {
			C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsCreated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreated", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsCreated")
	}
}

func (ptr *QHelpContentModel) ContentsCreated() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreated(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_ContentsCreationStarted
func callbackQHelpContentModel_ContentsCreationStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsCreationStarted") {
			C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsCreationStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreationStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreationStarted", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsCreationStarted")
	}
}

func (ptr *QHelpContentModel) ContentsCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	if ptr.Pointer() != nil {
		var customFilterNameC *C.char
		if customFilterName != "" {
			customFilterNameC = C.CString(customFilterName)
			defer C.free(unsafe.Pointer(customFilterNameC))
		}
		C.QHelpContentModel_CreateContents(ptr.Pointer(), C.struct_QtHelp_PackedString{data: customFilterNameC, len: C.longlong(len(customFilterName))})
	}
}

//export callbackQHelpContentModel_DestroyQHelpContentModel
func callbackQHelpContentModel_DestroyQHelpContentModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpContentModel"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).DestroyQHelpContentModelDefault()
	}
}

func (ptr *QHelpContentModel) ConnectDestroyQHelpContentModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpContentModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpContentModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpContentModel", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectDestroyQHelpContentModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpContentModel")
	}
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpContentModel) DestroyQHelpContentModelDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

//export callbackQHelpContentModel_Index
func callbackQHelpContentModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpContentModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "index"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "index", func(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
				signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(row, column, parent)
				return f(row, column, parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "index", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "index")
	}
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Parent
func callbackQHelpContentModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parent", func(index *core.QModelIndex) *core.QModelIndex {
				signal.(func(*core.QModelIndex) *core.QModelIndex)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parent", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectParent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parent")
	}
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Data
func callbackQHelpContentModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*core.QModelIndex, int) *core.QVariant)(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpContentModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func(index *core.QModelIndex, role int) *core.QVariant {
				signal.(func(*core.QModelIndex, int) *core.QVariant)(index, role)
				return f(index, role)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_IsCreatingContents(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentModel_MetaObject
func callbackQHelpContentModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpContentModel_ColumnCount
func callbackQHelpContentModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpContentModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpContentModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", func(parent *core.QModelIndex) int {
				signal.(func(*core.QModelIndex) int)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectColumnCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "columnCount")
	}
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpContentModel_RowCount
func callbackQHelpContentModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpContentModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpContentModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rowCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", func(parent *core.QModelIndex) int {
				signal.(func(*core.QModelIndex) int)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", f)
		}
	}
}

func (ptr *QHelpContentModel) DisconnectRowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rowCount")
	}
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QHelpContentModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QHelpContentModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QHelpContentModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___changePersistentIndexList_from_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QHelpContentModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___changePersistentIndexList_to_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QHelpContentModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QHelpContentModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QHelpContentModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QHelpContentModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QHelpContentModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___layoutChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QHelpContentModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpContentModel) __roleNames_newList() unsafe.Pointer {
	return C.QHelpContentModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QHelpContentModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpContentModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __itemData_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QHelpContentModel) __itemData_newList() unsafe.Pointer {
	return C.QHelpContentModel___itemData_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QHelpContentModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpContentModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___mimeData_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QHelpContentModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __match_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___match_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __match_newList() unsafe.Pointer {
	return C.QHelpContentModel___match_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___persistentIndexList_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QHelpContentModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QHelpContentModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpContentModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpContentModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpContentModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QHelpContentModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpContentModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpContentModel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentModel) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpContentModel___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpContentModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentModel) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpContentModel___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpContentModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentModel) __findChildren_newList() unsafe.Pointer {
	return C.QHelpContentModel___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpContentModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentModel) __children_newList() unsafe.Pointer {
	return C.QHelpContentModel___children_newList(ptr.Pointer())
}

//export callbackQHelpContentModel_DropMimeData
func callbackQHelpContentModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_InsertColumns
func callbackQHelpContentModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_InsertRows
func callbackQHelpContentModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_MoveColumns
func callbackQHelpContentModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpContentModel_MoveRows
func callbackQHelpContentModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveColumns
func callbackQHelpContentModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveRows
func callbackQHelpContentModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_SetData
func callbackQHelpContentModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpContentModel_SetHeaderData
func callbackQHelpContentModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpContentModel_SetItemData
func callbackQHelpContentModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_QtHelp_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, map[int]*core.QVariant) bool)(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetItemDataDefault(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
		out := make(map[int]*core.QVariant, int(l.len))
		tmpList := NewQHelpContentModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *QHelpContentModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_SetItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQHelpContentModel_Submit
func callbackQHelpContentModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpContentModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentModel_ColumnsAboutToBeInserted
func callbackQHelpContentModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsAboutToBeMoved
func callbackQHelpContentModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQHelpContentModel_ColumnsAboutToBeRemoved
func callbackQHelpContentModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsInserted
func callbackQHelpContentModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsMoved
func callbackQHelpContentModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQHelpContentModel_ColumnsRemoved
func callbackQHelpContentModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_DataChanged
func callbackQHelpContentModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex, []int))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackQHelpContentModel_FetchMore
func callbackQHelpContentModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpContentModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpContentModel_HeaderDataChanged
func callbackQHelpContentModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(core.Qt__Orientation, int, int))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_LayoutAboutToBeChanged
func callbackQHelpContentModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpContentModel_LayoutChanged
func callbackQHelpContentModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpContentModel_ModelAboutToBeReset
func callbackQHelpContentModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpContentModel_ModelReset
func callbackQHelpContentModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpContentModel_ResetInternalData
func callbackQHelpContentModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpContentModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_Revert
func callbackQHelpContentModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_RowsAboutToBeInserted
func callbackQHelpContentModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQHelpContentModel_RowsAboutToBeMoved
func callbackQHelpContentModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQHelpContentModel_RowsAboutToBeRemoved
func callbackQHelpContentModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_RowsInserted
func callbackQHelpContentModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_RowsMoved
func callbackQHelpContentModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQHelpContentModel_RowsRemoved
func callbackQHelpContentModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_Sort
func callbackQHelpContentModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpContentModel_RoleNames
func callbackQHelpContentModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewQHelpContentModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpContentModel) RoleNamesDefault() map[int]*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[int]*core.QByteArray {
			out := make(map[int]*core.QByteArray, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.QHelpContentModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*core.QByteArray, 0)
}

//export callbackQHelpContentModel_ItemData
func callbackQHelpContentModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*core.QModelIndex) map[int]*core.QVariant)(core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__itemData_newList())
		for k, v := range NewQHelpContentModelFromPointer(ptr).ItemDataDefault(core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpContentModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.QHelpContentModel_ItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*core.QVariant, 0)
}

//export callbackQHelpContentModel_MimeData
func callbackQHelpContentModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtHelp_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData(signal.(func([]*core.QModelIndex) *core.QMimeData)(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return core.PointerFromQMimeData(NewQHelpContentModelFromPointer(ptr).MimeDataDefault(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
		out := make([]*core.QModelIndex, int(l.len))
		tmpList := NewQHelpContentModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *QHelpContentModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMimeDataFromPointer(C.QHelpContentModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Buddy
func callbackQHelpContentModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Sibling
func callbackQHelpContentModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Match
func callbackQHelpContentModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__match_newList())
		for _, v := range NewQHelpContentModelFromPointer(ptr).MatchDefault(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpContentModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QHelpContentModel_MatchDefault(ptr.Pointer(), core.PointerFromQModelIndex(start), C.int(int32(role)), core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQHelpContentModel_Span
func callbackQHelpContentModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpContentModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_MimeTypes
func callbackQHelpContentModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtHelp_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewQHelpContentModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QHelpContentModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpContentModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpContentModel_HeaderData
func callbackQHelpContentModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpContentModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_SupportedDragActions
func callbackQHelpContentModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpContentModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_SupportedDropActions
func callbackQHelpContentModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpContentModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_Flags
func callbackQHelpContentModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpContentModel_CanDropMimeData
func callbackQHelpContentModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_CanFetchMore
func callbackQHelpContentModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_HasChildren
func callbackQHelpContentModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_Event
func callbackQHelpContentModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpContentModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpContentModel_EventFilter
func callbackQHelpContentModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentModel_ChildEvent
func callbackQHelpContentModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentModel_ConnectNotify
func callbackQHelpContentModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_CustomEvent
func callbackQHelpContentModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentModel_DeleteLater
func callbackQHelpContentModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpContentModel_Destroyed
func callbackQHelpContentModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpContentModel_DisconnectNotify
func callbackQHelpContentModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_ObjectNameChanged
func callbackQHelpContentModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpContentModel_TimerEvent
func callbackQHelpContentModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpContentWidget) {
	n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexOf(ptr.Pointer(), core.PointerFromQUrl(link)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func QHelpContentWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentWidget_QHelpContentWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpContentWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentWidget_QHelpContentWidget_Tr(sC, cC, C.int(int32(n))))
}

func QHelpContentWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentWidget_QHelpContentWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpContentWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpContentWidget_QHelpContentWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQHelpContentWidget_LinkActivated
func callbackQHelpContentWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linkActivated"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkActivated") {
			C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkActivated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", func(link *core.QUrl) {
				signal.(func(*core.QUrl))(link)
				f(link)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", f)
		}
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkActivated")
	}
}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

//export callbackQHelpContentWidget_MetaObject
func callbackQHelpContentWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentWidget) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpContentWidget) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QHelpContentWidget___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __selectedIndexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget___selectedIndexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __selectedIndexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___selectedIndexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpContentWidget) __selectedIndexes_newList() unsafe.Pointer {
	return C.QHelpContentWidget___selectedIndexes_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QHelpContentWidget___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___scrollBarWidgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *QHelpContentWidget) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.QHelpContentWidget___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpContentWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpContentWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QHelpContentWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpContentWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpContentWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QHelpContentWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpContentWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpContentWidget) __actions_newList() unsafe.Pointer {
	return C.QHelpContentWidget___actions_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpContentWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpContentWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpContentWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpContentWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentWidget) __findChildren_newList() unsafe.Pointer {
	return C.QHelpContentWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpContentWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpContentWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpContentWidget) __children_newList() unsafe.Pointer {
	return C.QHelpContentWidget___children_newList(ptr.Pointer())
}

//export callbackQHelpContentWidget_MoveCursor
func callbackQHelpContentWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpContentWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ViewportEvent
func callbackQHelpContentWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_Collapse
func callbackQHelpContentWidget_Collapse(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapse"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) CollapseDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_CollapseAll
func callbackQHelpContentWidget_CollapseAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapseAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseAllDefault()
	}
}

func (ptr *QHelpContentWidget) CollapseAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Collapsed
func callbackQHelpContentWidget_Collapsed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapsed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ColumnCountChanged
func callbackQHelpContentWidget_ColumnCountChanged(ptr unsafe.Pointer, oldCount C.int, newCount C.int) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		signal.(func(int, int))(int(int32(oldCount)), int(int32(newCount)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnCountChangedDefault(int(int32(oldCount)), int(int32(newCount)))
	}
}

func (ptr *QHelpContentWidget) ColumnCountChangedDefault(oldCount int, newCount int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnCountChangedDefault(ptr.Pointer(), C.int(int32(oldCount)), C.int(int32(newCount)))
	}
}

//export callbackQHelpContentWidget_ColumnMoved
func callbackQHelpContentWidget_ColumnMoved(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "columnMoved"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnMovedDefault()
	}
}

func (ptr *QHelpContentWidget) ColumnMovedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnMovedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ColumnResized
func callbackQHelpContentWidget_ColumnResized(ptr unsafe.Pointer, column C.int, oldSize C.int, newSize C.int) {
	if signal := qt.GetSignal(ptr, "columnResized"); signal != nil {
		signal.(func(int, int, int))(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnResizedDefault(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	}
}

func (ptr *QHelpContentWidget) ColumnResizedDefault(column int, oldSize int, newSize int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnResizedDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(oldSize)), C.int(int32(newSize)))
	}
}

//export callbackQHelpContentWidget_CurrentChanged
func callbackQHelpContentWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpContentWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpContentWidget_DataChanged
func callbackQHelpContentWidget_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex, []int))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DataChangedDefault(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *QHelpContentWidget) DataChangedDefault(topLeft core.QModelIndex_ITF, bottomRight core.QModelIndex_ITF, roles []int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DataChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(topLeft), core.PointerFromQModelIndex(bottomRight), func() unsafe.Pointer {
			tmpList := NewQHelpContentWidgetFromPointer(NewQHelpContentWidgetFromPointer(nil).__dataChanged_roles_newList())
			for _, v := range roles {
				tmpList.__dataChanged_roles_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQHelpContentWidget_DragMoveEvent
func callbackQHelpContentWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_Expand
func callbackQHelpContentWidget_Expand(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expand"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ExpandDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ExpandAll
func callbackQHelpContentWidget_ExpandAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expandAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandAllDefault()
	}
}

func (ptr *QHelpContentWidget) ExpandAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ExpandToDepth
func callbackQHelpContentWidget_ExpandToDepth(ptr unsafe.Pointer, depth C.int) {
	if signal := qt.GetSignal(ptr, "expandToDepth"); signal != nil {
		signal.(func(int))(int(int32(depth)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandToDepthDefault(int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ExpandToDepthDefault(depth int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandToDepthDefault(ptr.Pointer(), C.int(int32(depth)))
	}
}

//export callbackQHelpContentWidget_Expanded
func callbackQHelpContentWidget_Expanded(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expanded"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_HideColumn
func callbackQHelpContentWidget_HideColumn(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "hideColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) HideColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_KeyPressEvent
func callbackQHelpContentWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_KeyboardSearch
func callbackQHelpContentWidget_KeyboardSearch(ptr unsafe.Pointer, search C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "keyboardSearch"); signal != nil {
		signal.(func(string))(cGoUnpackString(search))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyboardSearchDefault(cGoUnpackString(search))
	}
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {
	if ptr.Pointer() != nil {
		var searchC *C.char
		if search != "" {
			searchC = C.CString(search)
			defer C.free(unsafe.Pointer(searchC))
		}
		C.QHelpContentWidget_KeyboardSearchDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchC, len: C.longlong(len(search))})
	}
}

//export callbackQHelpContentWidget_MouseDoubleClickEvent
func callbackQHelpContentWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseMoveEvent
func callbackQHelpContentWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MousePressEvent
func callbackQHelpContentWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseReleaseEvent
func callbackQHelpContentWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_PaintEvent
func callbackQHelpContentWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpContentWidget_Reset
func callbackQHelpContentWidget_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpContentWidget) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ResizeColumnToContents
func callbackQHelpContentWidget_ResizeColumnToContents(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "resizeColumnToContents"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeColumnToContentsDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ResizeColumnToContentsDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeColumnToContentsDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_RowsAboutToBeRemoved
func callbackQHelpContentWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsInserted
func callbackQHelpContentWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsRemoved
func callbackQHelpContentWidget_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_ScrollContentsBy
func callbackQHelpContentWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpContentWidget_ScrollTo
func callbackQHelpContentWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpContentWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_SelectAll
func callbackQHelpContentWidget_SelectAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpContentWidget) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionChanged
func callbackQHelpContentWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpContentWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpContentWidget_SetModel
func callbackQHelpContentWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpContentWidget_SetRootIndex
func callbackQHelpContentWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SetSelection
func callbackQHelpContentWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpContentWidget_SetSelectionModel
func callbackQHelpContentWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpContentWidget_ShowColumn
func callbackQHelpContentWidget_ShowColumn(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "showColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ShowColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_TimerEvent
func callbackQHelpContentWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpContentWidget_UpdateGeometries
func callbackQHelpContentWidget_UpdateGeometries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_IndexAt
func callbackQHelpContentWidget_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QHelpContentWidget) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRect
func callbackQHelpContentWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpContentWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHelpContentWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRegionForSelection
func callbackQHelpContentWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpContentWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpContentWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ViewportSizeHint
func callbackQHelpContentWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpContentWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_IsIndexHidden
func callbackQHelpContentWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpContentWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_HorizontalOffset
func callbackQHelpContentWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpContentWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_SizeHintForColumn
func callbackQHelpContentWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpContentWidget) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpContentWidget_VerticalOffset
func callbackQHelpContentWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpContentWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_DrawBranches
func callbackQHelpContentWidget_DrawBranches(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawBranchesDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranchesDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_DrawRow
func callbackQHelpContentWidget_DrawRow(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawRow"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawRowDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) DrawRowDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRowDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Edit2
func callbackQHelpContentWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_Event
func callbackQHelpContentWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_EventFilter
func callbackQHelpContentWidget_EventFilter(ptr unsafe.Pointer, object unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_FocusNextPrevChild
func callbackQHelpContentWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpContentWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_Activated
func callbackQHelpContentWidget_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ClearSelection
func callbackQHelpContentWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpContentWidget) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Clicked
func callbackQHelpContentWidget_Clicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_CloseEditor
func callbackQHelpContentWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_CommitData
func callbackQHelpContentWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpContentWidget_DoubleClicked
func callbackQHelpContentWidget_DoubleClicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_DragEnterEvent
func callbackQHelpContentWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpContentWidget_DragLeaveEvent
func callbackQHelpContentWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpContentWidget_DropEvent
func callbackQHelpContentWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpContentWidget_Edit
func callbackQHelpContentWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_EditorDestroyed
func callbackQHelpContentWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpContentWidget_Entered
func callbackQHelpContentWidget_Entered(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "entered"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_FocusInEvent
func callbackQHelpContentWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_FocusOutEvent
func callbackQHelpContentWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_IconSizeChanged
func callbackQHelpContentWidget_IconSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

//export callbackQHelpContentWidget_InputMethodEvent
func callbackQHelpContentWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpContentWidget_Pressed
func callbackQHelpContentWidget_Pressed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ResizeEvent
func callbackQHelpContentWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpContentWidget_ScrollToBottom
func callbackQHelpContentWidget_ScrollToBottom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpContentWidget) ScrollToBottomDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ScrollToTop
func callbackQHelpContentWidget_ScrollToTop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpContentWidget) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetCurrentIndex
func callbackQHelpContentWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_StartDrag
func callbackQHelpContentWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	if signal := qt.GetSignal(ptr, "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpContentWidget_Update
func callbackQHelpContentWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ViewportEntered
func callbackQHelpContentWidget_ViewportEntered(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "viewportEntered"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpContentWidget_SelectionCommand
func callbackQHelpContentWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpContentWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpContentWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpContentWidget_ViewOptions
func callbackQHelpContentWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpContentWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpContentWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_InputMethodQuery
func callbackQHelpContentWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpContentWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpContentWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_SizeHintForRow
func callbackQHelpContentWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpContentWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpContentWidget_ContextMenuEvent
func callbackQHelpContentWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpContentWidget_SetupViewport
func callbackQHelpContentWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpContentWidget_WheelEvent
func callbackQHelpContentWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpContentWidget_MinimumSizeHint
func callbackQHelpContentWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpContentWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_SizeHint
func callbackQHelpContentWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpContentWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ChangeEvent
func callbackQHelpContentWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpContentWidget_Close
func callbackQHelpContentWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpContentWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentWidget_ActionEvent
func callbackQHelpContentWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpContentWidget_CloseEvent
func callbackQHelpContentWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpContentWidget_CustomContextMenuRequested
func callbackQHelpContentWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpContentWidget_EnterEvent
func callbackQHelpContentWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_Hide
func callbackQHelpContentWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpContentWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_HideEvent
func callbackQHelpContentWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpContentWidget_KeyReleaseEvent
func callbackQHelpContentWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_LeaveEvent
func callbackQHelpContentWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_Lower
func callbackQHelpContentWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpContentWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_MoveEvent
func callbackQHelpContentWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_Raise
func callbackQHelpContentWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpContentWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Repaint
func callbackQHelpContentWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpContentWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetDisabled
func callbackQHelpContentWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpContentWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpContentWidget_SetEnabled
func callbackQHelpContentWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetFocus2
func callbackQHelpContentWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpContentWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetHidden
func callbackQHelpContentWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpContentWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpContentWidget_SetStyleSheet
func callbackQHelpContentWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpContentWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QHelpContentWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQHelpContentWidget_SetVisible
func callbackQHelpContentWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpContentWidget_SetWindowModified
func callbackQHelpContentWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetWindowTitle
func callbackQHelpContentWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpContentWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QHelpContentWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQHelpContentWidget_Show
func callbackQHelpContentWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpContentWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowEvent
func callbackQHelpContentWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpContentWidget_ShowFullScreen
func callbackQHelpContentWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpContentWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMaximized
func callbackQHelpContentWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpContentWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMinimized
func callbackQHelpContentWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpContentWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowNormal
func callbackQHelpContentWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpContentWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_TabletEvent
func callbackQHelpContentWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpContentWidget_UpdateMicroFocus
func callbackQHelpContentWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpContentWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_WindowIconChanged
func callbackQHelpContentWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpContentWidget_WindowTitleChanged
func callbackQHelpContentWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQHelpContentWidget_PaintEngine
func callbackQHelpContentWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQHelpContentWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpContentWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpContentWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpContentWidget_HasHeightForWidth
func callbackQHelpContentWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpContentWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentWidget_HeightForWidth
func callbackQHelpContentWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpContentWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpContentWidget_Metric
func callbackQHelpContentWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpContentWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpContentWidget_InitPainter
func callbackQHelpContentWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpContentWidget_ChildEvent
func callbackQHelpContentWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentWidget_ConnectNotify
func callbackQHelpContentWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_CustomEvent
func callbackQHelpContentWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_DeleteLater
func callbackQHelpContentWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpContentWidget_Destroyed
func callbackQHelpContentWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpContentWidget_DisconnectNotify
func callbackQHelpContentWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_ObjectNameChanged
func callbackQHelpContentWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

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

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) (n *QHelpEngine) {
	n = new(QHelpEngine)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	var collectionFileC *C.char
	if collectionFile != "" {
		collectionFileC = C.CString(collectionFile)
		defer C.free(unsafe.Pointer(collectionFileC))
	}
	tmpValue := NewQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(C.struct_QtHelp_PackedString{data: collectionFileC, len: C.longlong(len(collectionFile))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHelpEngine_DestroyQHelpEngine
func callbackQHelpEngine_DestroyQHelpEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineFromPointer(ptr).DestroyQHelpEngineDefault()
	}
}

func (ptr *QHelpEngine) ConnectDestroyQHelpEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngine", f)
		}
	}
}

func (ptr *QHelpEngine) DisconnectDestroyQHelpEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpEngine")
	}
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpEngine) DestroyQHelpEngineDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) (n *QHelpEngineCore) {
	n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	return
}
func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	var collectionFileC *C.char
	if collectionFile != "" {
		collectionFileC = C.CString(collectionFile)
		defer C.free(unsafe.Pointer(collectionFileC))
	}
	tmpValue := NewQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(C.struct_QtHelp_PackedString{data: collectionFileC, len: C.longlong(len(collectionFile))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpEngineCore) Files(namespaceName string, filterAttributes []string, extensionFilter string) []*core.QUrl {
	if ptr.Pointer() != nil {
		var namespaceNameC *C.char
		if namespaceName != "" {
			namespaceNameC = C.CString(namespaceName)
			defer C.free(unsafe.Pointer(namespaceNameC))
		}
		filterAttributesC := C.CString(strings.Join(filterAttributes, "|"))
		defer C.free(unsafe.Pointer(filterAttributesC))
		var extensionFilterC *C.char
		if extensionFilter != "" {
			extensionFilterC = C.CString(extensionFilter)
			defer C.free(unsafe.Pointer(extensionFilterC))
		}
		return func(l C.struct_QtHelp_PackedList) []*core.QUrl {
			out := make([]*core.QUrl, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__files_atList(i)
			}
			return out
		}(C.QHelpEngineCore_Files(ptr.Pointer(), C.struct_QtHelp_PackedString{data: namespaceNameC, len: C.longlong(len(namespaceName))}, C.struct_QtHelp_PackedString{data: filterAttributesC, len: C.longlong(len(strings.Join(filterAttributes, "|")))}, C.struct_QtHelp_PackedString{data: extensionFilterC, len: C.longlong(len(extensionFilter))}))
	}
	return make([]*core.QUrl, 0)
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	if ptr.Pointer() != nil {
		var namespaceNameC *C.char
		if namespaceName != "" {
			namespaceNameC = C.CString(namespaceName)
			defer C.free(unsafe.Pointer(namespaceNameC))
		}
		return cGoUnpackString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), C.struct_QtHelp_PackedString{data: namespaceNameC, len: C.longlong(len(namespaceName))}))
	}
	return ""
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}))
}

func (ptr *QHelpEngineCore) NamespaceName(documentationFileName string) string {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}))
}

func QHelpEngineCore_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpEngineCore) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_Tr(sC, cC, C.int(int32(n))))
}

func QHelpEngineCore_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpEngineCore) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_TrUtf8(sC, cC, C.int(int32(n))))
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}, C.struct_QtHelp_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QHelpEngineCore) MetaData(documentationFileName string, name string) *core.QVariant {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}, C.struct_QtHelp_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		attributesC := C.CString(strings.Join(attributes, "|"))
		defer C.free(unsafe.Pointer(attributesC))
		return int8(C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))}, C.struct_QtHelp_PackedString{data: attributesC, len: C.longlong(len(strings.Join(attributes, "|")))})) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		return int8(C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), C.struct_QtHelp_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	if ptr.Pointer() != nil {
		var documentationFileNameC *C.char
		if documentationFileName != "" {
			documentationFileNameC = C.CString(documentationFileName)
			defer C.free(unsafe.Pointer(documentationFileNameC))
		}
		return int8(C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))})) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return int8(C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		return int8(C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))})) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		return int8(C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(value))) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_SetupData(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	if ptr.Pointer() != nil {
		var namespaceNameC *C.char
		if namespaceName != "" {
			namespaceNameC = C.CString(namespaceName)
			defer C.free(unsafe.Pointer(namespaceNameC))
		}
		return int8(C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), C.struct_QtHelp_PackedString{data: namespaceNameC, len: C.longlong(len(namespaceName))})) != 0
	}
	return false
}

//export callbackQHelpEngineCore_CurrentFilterChanged
func callbackQHelpEngineCore_CurrentFilterChanged(ptr unsafe.Pointer, newFilter C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "currentFilterChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(newFilter))
	}

}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentFilterChanged") {
			C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentFilterChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "currentFilterChanged", func(newFilter string) {
				signal.(func(string))(newFilter)
				f(newFilter)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentFilterChanged", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentFilterChanged")
	}
}

func (ptr *QHelpEngineCore) CurrentFilterChanged(newFilter string) {
	if ptr.Pointer() != nil {
		var newFilterC *C.char
		if newFilter != "" {
			newFilterC = C.CString(newFilter)
			defer C.free(unsafe.Pointer(newFilterC))
		}
		C.QHelpEngineCore_CurrentFilterChanged(ptr.Pointer(), C.struct_QtHelp_PackedString{data: newFilterC, len: C.longlong(len(newFilter))})
	}
}

//export callbackQHelpEngineCore_ReadersAboutToBeInvalidated
func callbackQHelpEngineCore_ReadersAboutToBeInvalidated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readersAboutToBeInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "readersAboutToBeInvalidated") {
			C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "readersAboutToBeInvalidated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readersAboutToBeInvalidated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readersAboutToBeInvalidated", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "readersAboutToBeInvalidated")
	}
}

func (ptr *QHelpEngineCore) ReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ReadersAboutToBeInvalidated(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(save))))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), C.struct_QtHelp_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})
	}
}

//export callbackQHelpEngineCore_SetupFinished
func callbackQHelpEngineCore_SetupFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setupFinished") {
			C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setupFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setupFinished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setupFinished", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "setupFinished")
	}
}

func (ptr *QHelpEngineCore) SetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupFinished(ptr.Pointer())
	}
}

//export callbackQHelpEngineCore_SetupStarted
func callbackQHelpEngineCore_SetupStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setupStarted") {
			C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setupStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setupStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setupStarted", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "setupStarted")
	}
}

func (ptr *QHelpEngineCore) SetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupStarted(ptr.Pointer())
	}
}

//export callbackQHelpEngineCore_Warning
func callbackQHelpEngineCore_Warning(ptr unsafe.Pointer, msg C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "warning"); signal != nil {
		signal.(func(string))(cGoUnpackString(msg))
	}

}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "warning") {
			C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "warning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "warning", func(msg string) {
				signal.(func(string))(msg)
				f(msg)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "warning", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "warning")
	}
}

func (ptr *QHelpEngineCore) Warning(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QHelpEngineCore_Warning(ptr.Pointer(), C.struct_QtHelp_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQHelpEngineCore_DestroyQHelpEngineCore
func callbackQHelpEngineCore_DestroyQHelpEngineCore(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpEngineCore"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DestroyQHelpEngineCoreDefault()
	}
}

func (ptr *QHelpEngineCore) ConnectDestroyQHelpEngineCore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpEngineCore"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngineCore", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngineCore", f)
		}
	}
}

func (ptr *QHelpEngineCore) DisconnectDestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpEngineCore")
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCoreDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) LinksForIdentifier(id string) map[string]*core.QUrl {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		return func(l C.struct_QtHelp_PackedList) map[string]*core.QUrl {
			out := make(map[string]*core.QUrl, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i, v := range tmpList.__linksForIdentifier_keyList() {
				out[v] = tmpList.__linksForIdentifier_atList(v, i)
			}
			return out
		}(C.QHelpEngineCore_LinksForIdentifier(ptr.Pointer(), C.struct_QtHelp_PackedString{data: idC, len: C.longlong(len(id))}))
	}
	return make(map[string]*core.QUrl, 0)
}

func (ptr *QHelpEngineCore) LinksForKeyword(keyword string) map[string]*core.QUrl {
	if ptr.Pointer() != nil {
		var keywordC *C.char
		if keyword != "" {
			keywordC = C.CString(keyword)
			defer C.free(unsafe.Pointer(keywordC))
		}
		return func(l C.struct_QtHelp_PackedList) map[string]*core.QUrl {
			out := make(map[string]*core.QUrl, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i, v := range tmpList.__linksForKeyword_keyList() {
				out[v] = tmpList.__linksForKeyword_atList(v, i)
			}
			return out
		}(C.QHelpEngineCore_LinksForKeyword(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keywordC, len: C.longlong(len(keyword))}))
	}
	return make(map[string]*core.QUrl, 0)
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpEngineCore_MetaObject
func callbackQHelpEngineCore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineCoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngineCore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngineCore) __files_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore___files_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __files_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___files_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpEngineCore) __files_newList() unsafe.Pointer {
	return C.QHelpEngineCore___files_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __filterAttributeSets_atList(i int) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore___filterAttributeSets_atList(ptr.Pointer(), C.int(int32(i)))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) __filterAttributeSets_setList(i []string) {
	if ptr.Pointer() != nil {
		iC := C.CString(strings.Join(i, "|"))
		defer C.free(unsafe.Pointer(iC))
		C.QHelpEngineCore___filterAttributeSets_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(strings.Join(i, "|")))})
	}
}

func (ptr *QHelpEngineCore) __filterAttributeSets_newList() unsafe.Pointer {
	return C.QHelpEngineCore___filterAttributeSets_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __linksForIdentifier_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore___linksForIdentifier_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __linksForIdentifier_setList(key string, i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QHelpEngineCore___linksForIdentifier_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpEngineCore) __linksForIdentifier_newList() unsafe.Pointer {
	return C.QHelpEngineCore___linksForIdentifier_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __linksForIdentifier_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____linksForIdentifier_keyList_atList(i)
			}
			return out
		}(C.QHelpEngineCore___linksForIdentifier_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) __linksForKeyword_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore___linksForKeyword_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __linksForKeyword_setList(key string, i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QHelpEngineCore___linksForKeyword_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpEngineCore) __linksForKeyword_newList() unsafe.Pointer {
	return C.QHelpEngineCore___linksForKeyword_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __linksForKeyword_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____linksForKeyword_keyList_atList(i)
			}
			return out
		}(C.QHelpEngineCore___linksForKeyword_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_____linksForIdentifier_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpEngineCore_____linksForIdentifier_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpEngineCore) ____linksForIdentifier_keyList_newList() unsafe.Pointer {
	return C.QHelpEngineCore_____linksForIdentifier_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_____linksForKeyword_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpEngineCore_____linksForKeyword_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpEngineCore) ____linksForKeyword_keyList_newList() unsafe.Pointer {
	return C.QHelpEngineCore_____linksForKeyword_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpEngineCore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpEngineCore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpEngineCore___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpEngineCore___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpEngineCore) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpEngineCore___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpEngineCore___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpEngineCore) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpEngineCore___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpEngineCore___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpEngineCore) __findChildren_newList() unsafe.Pointer {
	return C.QHelpEngineCore___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpEngineCore___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpEngineCore) __children_newList() unsafe.Pointer {
	return C.QHelpEngineCore___children_newList(ptr.Pointer())
}

//export callbackQHelpEngineCore_Event
func callbackQHelpEngineCore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpEngineCore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpEngineCore_EventFilter
func callbackQHelpEngineCore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngineCore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpEngineCore_ChildEvent
func callbackQHelpEngineCore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpEngineCore_ConnectNotify
func callbackQHelpEngineCore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_CustomEvent
func callbackQHelpEngineCore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpEngineCore_DeleteLater
func callbackQHelpEngineCore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngineCore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpEngineCore_Destroyed
func callbackQHelpEngineCore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpEngineCore_DisconnectNotify
func callbackQHelpEngineCore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_ObjectNameChanged
func callbackQHelpEngineCore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpEngineCore_TimerEvent
func callbackQHelpEngineCore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) (n *QHelpIndexModel) {
	n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var filterC *C.char
		if filter != "" {
			filterC = C.CString(filter)
			defer C.free(unsafe.Pointer(filterC))
		}
		var wildcardC *C.char
		if wildcard != "" {
			wildcardC = C.CString(wildcard)
			defer C.free(unsafe.Pointer(wildcardC))
		}
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_Filter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtHelp_PackedString{data: wildcardC, len: C.longlong(len(wildcard))}))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func QHelpIndexModel_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexModel_QHelpIndexModel_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpIndexModel) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexModel_QHelpIndexModel_Tr(sC, cC, C.int(int32(n))))
}

func QHelpIndexModel_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexModel_QHelpIndexModel_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpIndexModel) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexModel_QHelpIndexModel_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	if ptr.Pointer() != nil {
		var customFilterNameC *C.char
		if customFilterName != "" {
			customFilterNameC = C.CString(customFilterName)
			defer C.free(unsafe.Pointer(customFilterNameC))
		}
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), C.struct_QtHelp_PackedString{data: customFilterNameC, len: C.longlong(len(customFilterName))})
	}
}

//export callbackQHelpIndexModel_IndexCreated
func callbackQHelpIndexModel_IndexCreated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "indexCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexCreated") {
			C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexCreated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexCreated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexCreated", f)
		}
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "indexCreated")
	}
}

func (ptr *QHelpIndexModel) IndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreated(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_IndexCreationStarted
func callbackQHelpIndexModel_IndexCreationStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexCreationStarted") {
			C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexCreationStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexCreationStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexCreationStarted", f)
		}
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "indexCreationStarted")
	}
}

func (ptr *QHelpIndexModel) IndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MetaObject
func callbackQHelpIndexModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexModel) __linksForKeyword_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpIndexModel___linksForKeyword_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __linksForKeyword_setList(key string, i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QHelpIndexModel___linksForKeyword_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpIndexModel) __linksForKeyword_newList() unsafe.Pointer {
	return C.QHelpIndexModel___linksForKeyword_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __linksForKeyword_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____linksForKeyword_keyList_atList(i)
			}
			return out
		}(C.QHelpIndexModel___linksForKeyword_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpIndexModel_____linksForKeyword_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpIndexModel_____linksForKeyword_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpIndexModel) ____linksForKeyword_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____linksForKeyword_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QHelpIndexModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QHelpIndexModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QHelpIndexModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___changePersistentIndexList_from_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QHelpIndexModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___changePersistentIndexList_to_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QHelpIndexModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QHelpIndexModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QHelpIndexModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QHelpIndexModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QHelpIndexModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___layoutChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QHelpIndexModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpIndexModel) __roleNames_newList() unsafe.Pointer {
	return C.QHelpIndexModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QHelpIndexModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpIndexModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __itemData_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QHelpIndexModel) __itemData_newList() unsafe.Pointer {
	return C.QHelpIndexModel___itemData_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QHelpIndexModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QHelpIndexModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___mimeData_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QHelpIndexModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __match_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___match_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __match_newList() unsafe.Pointer {
	return C.QHelpIndexModel___match_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___persistentIndexList_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QHelpIndexModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpIndexModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpIndexModel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexModel) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpIndexModel___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexModel) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpIndexModel___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexModel) __findChildren_newList() unsafe.Pointer {
	return C.QHelpIndexModel___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpIndexModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexModel) __children_newList() unsafe.Pointer {
	return C.QHelpIndexModel___children_newList(ptr.Pointer())
}

//export callbackQHelpIndexModel_InsertRows
func callbackQHelpIndexModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RemoveRows
func callbackQHelpIndexModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SetData
func callbackQHelpIndexModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Sort
func callbackQHelpIndexModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpIndexModel_Sibling
func callbackQHelpIndexModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QHelpIndexModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Data
func callbackQHelpIndexModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*core.QModelIndex, int) *core.QVariant)(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpIndexModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_SupportedDropActions
func callbackQHelpIndexModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpIndexModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_Flags
func callbackQHelpIndexModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpIndexModel_RowCount
func callbackQHelpIndexModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_DropMimeData
func callbackQHelpIndexModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Index
func callbackQHelpIndexModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpIndexModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_InsertColumns
func callbackQHelpIndexModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MoveColumns
func callbackQHelpIndexModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MoveRows
func callbackQHelpIndexModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RemoveColumns
func callbackQHelpIndexModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SetHeaderData
func callbackQHelpIndexModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SetItemData
func callbackQHelpIndexModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_QtHelp_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, map[int]*core.QVariant) bool)(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetItemDataDefault(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
		out := make(map[int]*core.QVariant, int(l.len))
		tmpList := NewQHelpIndexModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *QHelpIndexModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SetItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Submit
func callbackQHelpIndexModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpIndexModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ColumnsAboutToBeInserted
func callbackQHelpIndexModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsAboutToBeMoved
func callbackQHelpIndexModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQHelpIndexModel_ColumnsAboutToBeRemoved
func callbackQHelpIndexModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsInserted
func callbackQHelpIndexModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsMoved
func callbackQHelpIndexModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQHelpIndexModel_ColumnsRemoved
func callbackQHelpIndexModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_DataChanged
func callbackQHelpIndexModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex, []int))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackQHelpIndexModel_FetchMore
func callbackQHelpIndexModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpIndexModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpIndexModel_HeaderDataChanged
func callbackQHelpIndexModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(core.Qt__Orientation, int, int))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_LayoutAboutToBeChanged
func callbackQHelpIndexModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpIndexModel_LayoutChanged
func callbackQHelpIndexModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpIndexModel_ModelAboutToBeReset
func callbackQHelpIndexModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpIndexModel_ModelReset
func callbackQHelpIndexModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpIndexModel_ResetInternalData
func callbackQHelpIndexModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpIndexModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_Revert
func callbackQHelpIndexModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_RowsAboutToBeInserted
func callbackQHelpIndexModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQHelpIndexModel_RowsAboutToBeMoved
func callbackQHelpIndexModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQHelpIndexModel_RowsAboutToBeRemoved
func callbackQHelpIndexModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_RowsInserted
func callbackQHelpIndexModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_RowsMoved
func callbackQHelpIndexModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQHelpIndexModel_RowsRemoved
func callbackQHelpIndexModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_RoleNames
func callbackQHelpIndexModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewQHelpIndexModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpIndexModel) RoleNamesDefault() map[int]*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[int]*core.QByteArray {
			out := make(map[int]*core.QByteArray, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.QHelpIndexModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*core.QByteArray, 0)
}

//export callbackQHelpIndexModel_ItemData
func callbackQHelpIndexModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*core.QModelIndex) map[int]*core.QVariant)(core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__itemData_newList())
		for k, v := range NewQHelpIndexModelFromPointer(ptr).ItemDataDefault(core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpIndexModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.QHelpIndexModel_ItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*core.QVariant, 0)
}

//export callbackQHelpIndexModel_MimeData
func callbackQHelpIndexModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtHelp_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData(signal.(func([]*core.QModelIndex) *core.QMimeData)(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return core.PointerFromQMimeData(NewQHelpIndexModelFromPointer(ptr).MimeDataDefault(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
		out := make([]*core.QModelIndex, int(l.len))
		tmpList := NewQHelpIndexModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *QHelpIndexModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMimeDataFromPointer(C.QHelpIndexModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Buddy
func callbackQHelpIndexModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Parent
func callbackQHelpIndexModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Match
func callbackQHelpIndexModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__match_newList())
		for _, v := range NewQHelpIndexModelFromPointer(ptr).MatchDefault(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpIndexModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QHelpIndexModel_MatchDefault(ptr.Pointer(), core.PointerFromQModelIndex(start), C.int(int32(role)), core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQHelpIndexModel_Span
func callbackQHelpIndexModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpIndexModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_MimeTypes
func callbackQHelpIndexModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtHelp_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewQHelpIndexModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QHelpIndexModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpIndexModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpIndexModel_HeaderData
func callbackQHelpIndexModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpIndexModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_SupportedDragActions
func callbackQHelpIndexModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpIndexModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_CanDropMimeData
func callbackQHelpIndexModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_CanFetchMore
func callbackQHelpIndexModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_HasChildren
func callbackQHelpIndexModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ColumnCount
func callbackQHelpIndexModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_Event
func callbackQHelpIndexModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpIndexModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_EventFilter
func callbackQHelpIndexModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ChildEvent
func callbackQHelpIndexModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexModel_ConnectNotify
func callbackQHelpIndexModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_CustomEvent
func callbackQHelpIndexModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexModel_DeleteLater
func callbackQHelpIndexModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpIndexModel_Destroyed
func callbackQHelpIndexModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpIndexModel_DisconnectNotify
func callbackQHelpIndexModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_ObjectNameChanged
func callbackQHelpIndexModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpIndexModel_TimerEvent
func callbackQHelpIndexModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpIndexWidget) {
	n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	return
}
func QHelpIndexWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexWidget_QHelpIndexWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpIndexWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexWidget_QHelpIndexWidget_Tr(sC, cC, C.int(int32(n))))
}

func QHelpIndexWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexWidget_QHelpIndexWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpIndexWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpIndexWidget_QHelpIndexWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQHelpIndexWidget_ActivateCurrentItem
func callbackQHelpIndexWidget_ActivateCurrentItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activateCurrentItem"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActivateCurrentItemDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectActivateCurrentItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activateCurrentItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activateCurrentItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activateCurrentItem", f)
		}
	}
}

func (ptr *QHelpIndexWidget) DisconnectActivateCurrentItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activateCurrentItem")
	}
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ActivateCurrentItemDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItemDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_FilterIndices
func callbackQHelpIndexWidget_FilterIndices(ptr unsafe.Pointer, filter C.struct_QtHelp_PackedString, wildcard C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "filterIndices"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(filter), cGoUnpackString(wildcard))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FilterIndicesDefault(cGoUnpackString(filter), cGoUnpackString(wildcard))
	}
}

func (ptr *QHelpIndexWidget) ConnectFilterIndices(f func(filter string, wildcard string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filterIndices"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "filterIndices", func(filter string, wildcard string) {
				signal.(func(string, string))(filter, wildcard)
				f(filter, wildcard)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filterIndices", f)
		}
	}
}

func (ptr *QHelpIndexWidget) DisconnectFilterIndices() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filterIndices")
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	if ptr.Pointer() != nil {
		var filterC *C.char
		if filter != "" {
			filterC = C.CString(filter)
			defer C.free(unsafe.Pointer(filterC))
		}
		var wildcardC *C.char
		if wildcard != "" {
			wildcardC = C.CString(wildcard)
			defer C.free(unsafe.Pointer(wildcardC))
		}
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtHelp_PackedString{data: wildcardC, len: C.longlong(len(wildcard))})
	}
}

func (ptr *QHelpIndexWidget) FilterIndicesDefault(filter string, wildcard string) {
	if ptr.Pointer() != nil {
		var filterC *C.char
		if filter != "" {
			filterC = C.CString(filter)
			defer C.free(unsafe.Pointer(filterC))
		}
		var wildcardC *C.char
		if wildcard != "" {
			wildcardC = C.CString(wildcard)
			defer C.free(unsafe.Pointer(wildcardC))
		}
		C.QHelpIndexWidget_FilterIndicesDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtHelp_PackedString{data: wildcardC, len: C.longlong(len(wildcard))})
	}
}

//export callbackQHelpIndexWidget_LinkActivated
func callbackQHelpIndexWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer, keyword C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "linkActivated"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(link), cGoUnpackString(keyword))
	}

}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkActivated") {
			C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkActivated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", func(link *core.QUrl, keyword string) {
				signal.(func(*core.QUrl, string))(link, keyword)
				f(link, keyword)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", f)
		}
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkActivated")
	}
}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {
	if ptr.Pointer() != nil {
		var keywordC *C.char
		if keyword != "" {
			keywordC = C.CString(keyword)
			defer C.free(unsafe.Pointer(keywordC))
		}
		C.QHelpIndexWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link), C.struct_QtHelp_PackedString{data: keywordC, len: C.longlong(len(keyword))})
	}
}

//export callbackQHelpIndexWidget_LinksActivated
func callbackQHelpIndexWidget_LinksActivated(ptr unsafe.Pointer, links C.struct_QtHelp_PackedList, keyword C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "linksActivated"); signal != nil {
		signal.(func(map[string]*core.QUrl, string))(func(l C.struct_QtHelp_PackedList) map[string]*core.QUrl {
			out := make(map[string]*core.QUrl, int(l.len))
			tmpList := NewQHelpIndexWidgetFromPointer(l.data)
			for i, v := range tmpList.__linksActivated_links_keyList() {
				out[v] = tmpList.__linksActivated_links_atList(v, i)
			}
			return out
		}(links), cGoUnpackString(keyword))
	}

}

func (ptr *QHelpIndexWidget) ConnectLinksActivated(f func(links map[string]*core.QUrl, keyword string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linksActivated") {
			C.QHelpIndexWidget_ConnectLinksActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linksActivated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linksActivated", func(links map[string]*core.QUrl, keyword string) {
				signal.(func(map[string]*core.QUrl, string))(links, keyword)
				f(links, keyword)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linksActivated", f)
		}
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinksActivated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinksActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linksActivated")
	}
}

func (ptr *QHelpIndexWidget) LinksActivated(links map[string]*core.QUrl, keyword string) {
	if ptr.Pointer() != nil {
		var keywordC *C.char
		if keyword != "" {
			keywordC = C.CString(keyword)
			defer C.free(unsafe.Pointer(keywordC))
		}
		C.QHelpIndexWidget_LinksActivated(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQHelpIndexWidgetFromPointer(NewQHelpIndexWidgetFromPointer(nil).__linksActivated_links_newList())
			for k, v := range links {
				tmpList.__linksActivated_links_setList(k, v)
			}
			return tmpList.Pointer()
		}(), C.struct_QtHelp_PackedString{data: keywordC, len: C.longlong(len(keyword))})
	}
}

//export callbackQHelpIndexWidget_MetaObject
func callbackQHelpIndexWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexWidget) __linksActivated_links_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpIndexWidget___linksActivated_links_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __linksActivated_links_setList(key string, i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QHelpIndexWidget___linksActivated_links_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpIndexWidget) __linksActivated_links_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___linksActivated_links_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __linksActivated_links_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpIndexWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____linksActivated_links_keyList_atList(i)
			}
			return out
		}(C.QHelpIndexWidget___linksActivated_links_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpIndexWidget_____linksActivated_links_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpIndexWidget_____linksActivated_links_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpIndexWidget) ____linksActivated_links_keyList_newList() unsafe.Pointer {
	return C.QHelpIndexWidget_____linksActivated_links_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QHelpIndexWidget) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget___indexesMoved_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___indexesMoved_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexWidget) __indexesMoved_indexes_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___indexesMoved_indexes_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __selectedIndexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget___selectedIndexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __selectedIndexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___selectedIndexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QHelpIndexWidget) __selectedIndexes_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___selectedIndexes_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QHelpIndexWidget___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___scrollBarWidgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *QHelpIndexWidget) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpIndexWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpIndexWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpIndexWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpIndexWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpIndexWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpIndexWidget) __actions_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___actions_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpIndexWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpIndexWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexWidget) __findChildren_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpIndexWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpIndexWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpIndexWidget) __children_newList() unsafe.Pointer {
	return C.QHelpIndexWidget___children_newList(ptr.Pointer())
}

//export callbackQHelpIndexWidget_MoveCursor
func callbackQHelpIndexWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpIndexWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_Event
func callbackQHelpIndexWidget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpIndexWidget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_CurrentChanged
func callbackQHelpIndexWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpIndexWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpIndexWidget_DataChanged
func callbackQHelpIndexWidget_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex, []int))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DataChangedDefault(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpIndexWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *QHelpIndexWidget) DataChangedDefault(topLeft core.QModelIndex_ITF, bottomRight core.QModelIndex_ITF, roles []int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DataChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(topLeft), core.PointerFromQModelIndex(bottomRight), func() unsafe.Pointer {
			tmpList := NewQHelpIndexWidgetFromPointer(NewQHelpIndexWidgetFromPointer(nil).__dataChanged_roles_newList())
			for _, v := range roles {
				tmpList.__dataChanged_roles_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQHelpIndexWidget_DragLeaveEvent
func callbackQHelpIndexWidget_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DragMoveEvent
func callbackQHelpIndexWidget_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DropEvent
func callbackQHelpIndexWidget_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQHelpIndexWidget_MouseMoveEvent
func callbackQHelpIndexWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MouseReleaseEvent
func callbackQHelpIndexWidget_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_PaintEvent
func callbackQHelpIndexWidget_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

//export callbackQHelpIndexWidget_ResizeEvent
func callbackQHelpIndexWidget_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

//export callbackQHelpIndexWidget_RowsAboutToBeRemoved
func callbackQHelpIndexWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_RowsInserted
func callbackQHelpIndexWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_ScrollTo
func callbackQHelpIndexWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_SelectionChanged
func callbackQHelpIndexWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpIndexWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpIndexWidget_SetSelection
func callbackQHelpIndexWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpIndexWidget_StartDrag
func callbackQHelpIndexWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	if signal := qt.GetSignal(ptr, "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpIndexWidget_TimerEvent
func callbackQHelpIndexWidget_TimerEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) TimerEventDefault(e core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

//export callbackQHelpIndexWidget_UpdateGeometries
func callbackQHelpIndexWidget_UpdateGeometries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_WheelEvent
func callbackQHelpIndexWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpIndexWidget_IndexAt
func callbackQHelpIndexWidget_IndexAt(ptr unsafe.Pointer, p unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(p)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(p)))
}

func (ptr *QHelpIndexWidget) IndexAtDefault(p core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_SelectedIndexes
func callbackQHelpIndexWidget_SelectedIndexes(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectedIndexes"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexWidgetFromPointer(NewQHelpIndexWidgetFromPointer(nil).__selectedIndexes_newList())
			for _, v := range signal.(func() []*core.QModelIndex)() {
				tmpList.__selectedIndexes_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQHelpIndexWidgetFromPointer(NewQHelpIndexWidgetFromPointer(nil).__selectedIndexes_newList())
		for _, v := range NewQHelpIndexWidgetFromPointer(ptr).SelectedIndexesDefault() {
			tmpList.__selectedIndexes_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QHelpIndexWidget) SelectedIndexesDefault() []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQHelpIndexWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__selectedIndexes_atList(i)
			}
			return out
		}(C.QHelpIndexWidget_SelectedIndexesDefault(ptr.Pointer()))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQHelpIndexWidget_VisualRect
func callbackQHelpIndexWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpIndexWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRegionForSelection
func callbackQHelpIndexWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpIndexWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpIndexWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ViewportSizeHint
func callbackQHelpIndexWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ViewOptions
func callbackQHelpIndexWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpIndexWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpIndexWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_IsIndexHidden
func callbackQHelpIndexWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpIndexWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_HorizontalOffset
func callbackQHelpIndexWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_VerticalOffset
func callbackQHelpIndexWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_Edit2
func callbackQHelpIndexWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_EventFilter
func callbackQHelpIndexWidget_EventFilter(ptr unsafe.Pointer, object unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_FocusNextPrevChild
func callbackQHelpIndexWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpIndexWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_ViewportEvent
func callbackQHelpIndexWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_Activated
func callbackQHelpIndexWidget_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_ClearSelection
func callbackQHelpIndexWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpIndexWidget) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Clicked
func callbackQHelpIndexWidget_Clicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_CloseEditor
func callbackQHelpIndexWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_CommitData
func callbackQHelpIndexWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpIndexWidget_DoubleClicked
func callbackQHelpIndexWidget_DoubleClicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_DragEnterEvent
func callbackQHelpIndexWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpIndexWidget_Edit
func callbackQHelpIndexWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_EditorDestroyed
func callbackQHelpIndexWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpIndexWidget_Entered
func callbackQHelpIndexWidget_Entered(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "entered"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_FocusInEvent
func callbackQHelpIndexWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_FocusOutEvent
func callbackQHelpIndexWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_IconSizeChanged
func callbackQHelpIndexWidget_IconSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

//export callbackQHelpIndexWidget_InputMethodEvent
func callbackQHelpIndexWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpIndexWidget_KeyPressEvent
func callbackQHelpIndexWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_KeyboardSearch
func callbackQHelpIndexWidget_KeyboardSearch(ptr unsafe.Pointer, search C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "keyboardSearch"); signal != nil {
		signal.(func(string))(cGoUnpackString(search))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyboardSearchDefault(cGoUnpackString(search))
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {
	if ptr.Pointer() != nil {
		var searchC *C.char
		if search != "" {
			searchC = C.CString(search)
			defer C.free(unsafe.Pointer(searchC))
		}
		C.QHelpIndexWidget_KeyboardSearchDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchC, len: C.longlong(len(search))})
	}
}

//export callbackQHelpIndexWidget_MouseDoubleClickEvent
func callbackQHelpIndexWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_MousePressEvent
func callbackQHelpIndexWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_Pressed
func callbackQHelpIndexWidget_Pressed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_Reset
func callbackQHelpIndexWidget_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpIndexWidget) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToBottom
func callbackQHelpIndexWidget_ScrollToBottom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpIndexWidget) ScrollToBottomDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToTop
func callbackQHelpIndexWidget_ScrollToTop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpIndexWidget) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectAll
func callbackQHelpIndexWidget_SelectAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetCurrentIndex
func callbackQHelpIndexWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetModel
func callbackQHelpIndexWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpIndexWidget_SetRootIndex
func callbackQHelpIndexWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetSelectionModel
func callbackQHelpIndexWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpIndexWidget_Update
func callbackQHelpIndexWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_ViewportEntered
func callbackQHelpIndexWidget_ViewportEntered(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "viewportEntered"); signal != nil {
		signal.(func())()
	}

}

//export callbackQHelpIndexWidget_SelectionCommand
func callbackQHelpIndexWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpIndexWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpIndexWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpIndexWidget_InputMethodQuery
func callbackQHelpIndexWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpIndexWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpIndexWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_SizeHintForColumn
func callbackQHelpIndexWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpIndexWidget) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_SizeHintForRow
func callbackQHelpIndexWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpIndexWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_ContextMenuEvent
func callbackQHelpIndexWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpIndexWidget_ScrollContentsBy
func callbackQHelpIndexWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpIndexWidget_SetupViewport
func callbackQHelpIndexWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpIndexWidget_MinimumSizeHint
func callbackQHelpIndexWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpIndexWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_SizeHint
func callbackQHelpIndexWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpIndexWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ChangeEvent
func callbackQHelpIndexWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpIndexWidget_Close
func callbackQHelpIndexWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpIndexWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_ActionEvent
func callbackQHelpIndexWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpIndexWidget_CloseEvent
func callbackQHelpIndexWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpIndexWidget_CustomContextMenuRequested
func callbackQHelpIndexWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpIndexWidget_EnterEvent
func callbackQHelpIndexWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_Hide
func callbackQHelpIndexWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpIndexWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_HideEvent
func callbackQHelpIndexWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpIndexWidget_KeyReleaseEvent
func callbackQHelpIndexWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_LeaveEvent
func callbackQHelpIndexWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_Lower
func callbackQHelpIndexWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpIndexWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_MoveEvent
func callbackQHelpIndexWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpIndexWidget_Raise
func callbackQHelpIndexWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpIndexWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Repaint
func callbackQHelpIndexWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpIndexWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetDisabled
func callbackQHelpIndexWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpIndexWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpIndexWidget_SetEnabled
func callbackQHelpIndexWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetFocus2
func callbackQHelpIndexWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpIndexWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetHidden
func callbackQHelpIndexWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpIndexWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpIndexWidget_SetStyleSheet
func callbackQHelpIndexWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpIndexWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QHelpIndexWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQHelpIndexWidget_SetVisible
func callbackQHelpIndexWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpIndexWidget_SetWindowModified
func callbackQHelpIndexWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetWindowTitle
func callbackQHelpIndexWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpIndexWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QHelpIndexWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQHelpIndexWidget_Show
func callbackQHelpIndexWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpIndexWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowEvent
func callbackQHelpIndexWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpIndexWidget_ShowFullScreen
func callbackQHelpIndexWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpIndexWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMaximized
func callbackQHelpIndexWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMinimized
func callbackQHelpIndexWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowNormal
func callbackQHelpIndexWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpIndexWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_TabletEvent
func callbackQHelpIndexWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpIndexWidget_UpdateMicroFocus
func callbackQHelpIndexWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpIndexWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_WindowIconChanged
func callbackQHelpIndexWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpIndexWidget_WindowTitleChanged
func callbackQHelpIndexWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQHelpIndexWidget_PaintEngine
func callbackQHelpIndexWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQHelpIndexWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpIndexWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpIndexWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpIndexWidget_HasHeightForWidth
func callbackQHelpIndexWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpIndexWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_HeightForWidth
func callbackQHelpIndexWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpIndexWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Metric
func callbackQHelpIndexWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpIndexWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpIndexWidget_InitPainter
func callbackQHelpIndexWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpIndexWidget_ChildEvent
func callbackQHelpIndexWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexWidget_ConnectNotify
func callbackQHelpIndexWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_CustomEvent
func callbackQHelpIndexWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_DeleteLater
func callbackQHelpIndexWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpIndexWidget_Destroyed
func callbackQHelpIndexWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpIndexWidget_DisconnectNotify
func callbackQHelpIndexWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_ObjectNameChanged
func callbackQHelpIndexWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

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

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) (n *QHelpSearchEngine) {
	n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	return
}
func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	tmpValue := NewQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QHelpSearchEngine_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchEngine_QHelpSearchEngine_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchEngine) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchEngine_QHelpSearchEngine_Tr(sC, cC, C.int(int32(n))))
}

func QHelpSearchEngine_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchEngine_QHelpSearchEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchEngine) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchEngine_QHelpSearchEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQHelpSearchEngine_CancelIndexing
func callbackQHelpSearchEngine_CancelIndexing(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cancelIndexing"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CancelIndexingDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectCancelIndexing(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cancelIndexing"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cancelIndexing", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cancelIndexing", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelIndexing() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cancelIndexing")
	}
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) CancelIndexingDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexingDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_CancelSearching
func callbackQHelpSearchEngine_CancelSearching(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cancelSearching"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CancelSearchingDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectCancelSearching(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cancelSearching"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cancelSearching", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cancelSearching", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelSearching() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cancelSearching")
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) CancelSearchingDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearchingDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_IndexingFinished
func callbackQHelpSearchEngine_IndexingFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "indexingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexingFinished") {
			C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexingFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexingFinished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexingFinished", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "indexingFinished")
	}
}

func (ptr *QHelpSearchEngine) IndexingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingFinished(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_IndexingStarted
func callbackQHelpSearchEngine_IndexingStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "indexingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexingStarted") {
			C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexingStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexingStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexingStarted", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "indexingStarted")
	}
}

func (ptr *QHelpSearchEngine) IndexingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingStarted(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_ReindexDocumentation
func callbackQHelpSearchEngine_ReindexDocumentation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reindexDocumentation"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ReindexDocumentationDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectReindexDocumentation(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reindexDocumentation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reindexDocumentation", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reindexDocumentation", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectReindexDocumentation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reindexDocumentation")
	}
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ReindexDocumentationDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentationDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_Search
func callbackQHelpSearchEngine_Search(ptr unsafe.Pointer, searchInput C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "search"); signal != nil {
		signal.(func(string))(cGoUnpackString(searchInput))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).SearchDefault(cGoUnpackString(searchInput))
	}
}

func (ptr *QHelpSearchEngine) ConnectSearch(f func(searchInput string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "search"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "search", func(searchInput string) {
				signal.(func(string))(searchInput)
				f(searchInput)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "search", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "search")
	}
}

func (ptr *QHelpSearchEngine) Search(searchInput string) {
	if ptr.Pointer() != nil {
		var searchInputC *C.char
		if searchInput != "" {
			searchInputC = C.CString(searchInput)
			defer C.free(unsafe.Pointer(searchInputC))
		}
		C.QHelpSearchEngine_Search(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchInputC, len: C.longlong(len(searchInput))})
	}
}

func (ptr *QHelpSearchEngine) SearchDefault(searchInput string) {
	if ptr.Pointer() != nil {
		var searchInputC *C.char
		if searchInput != "" {
			searchInputC = C.CString(searchInput)
			defer C.free(unsafe.Pointer(searchInputC))
		}
		C.QHelpSearchEngine_SearchDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchInputC, len: C.longlong(len(searchInput))})
	}
}

//export callbackQHelpSearchEngine_SearchingFinished
func callbackQHelpSearchEngine_SearchingFinished(ptr unsafe.Pointer, searchResultCount C.int) {
	if signal := qt.GetSignal(ptr, "searchingFinished"); signal != nil {
		signal.(func(int))(int(int32(searchResultCount)))
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(searchResultCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "searchingFinished") {
			C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "searchingFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "searchingFinished", func(searchResultCount int) {
				signal.(func(int))(searchResultCount)
				f(searchResultCount)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "searchingFinished", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "searchingFinished")
	}
}

func (ptr *QHelpSearchEngine) SearchingFinished(searchResultCount int) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingFinished(ptr.Pointer(), C.int(int32(searchResultCount)))
	}
}

//export callbackQHelpSearchEngine_SearchingStarted
func callbackQHelpSearchEngine_SearchingStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "searchingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "searchingStarted") {
			C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "searchingStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "searchingStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "searchingStarted", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "searchingStarted")
	}
}

func (ptr *QHelpSearchEngine) SearchingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingStarted(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_DestroyQHelpSearchEngine
func callbackQHelpSearchEngine_DestroyQHelpSearchEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpSearchEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DestroyQHelpSearchEngineDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectDestroyQHelpSearchEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchEngine", f)
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectDestroyQHelpSearchEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpSearchEngine")
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngineDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchEngine) SearchInput() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchEngine_SearchInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpSearchEngine) SearchResults(start int, end int) []*QHelpSearchResult {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []*QHelpSearchResult {
			out := make([]*QHelpSearchResult, int(l.len))
			tmpList := NewQHelpSearchEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__searchResults_atList(i)
			}
			return out
		}(C.QHelpSearchEngine_SearchResults(ptr.Pointer(), C.int(int32(start)), C.int(int32(end))))
	}
	return make([]*QHelpSearchResult, 0)
}

//export callbackQHelpSearchEngine_MetaObject
func callbackQHelpSearchEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) SearchResultCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchEngine_SearchResultCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpSearchEngine) __search_queryList_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchEngine___search_queryList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __search_queryList_setList(i QHelpSearchQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___search_queryList_setList(ptr.Pointer(), PointerFromQHelpSearchQuery(i))
	}
}

func (ptr *QHelpSearchEngine) __search_queryList_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___search_queryList_newList(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchEngine___query_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __query_setList(i QHelpSearchQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___query_setList(ptr.Pointer(), PointerFromQHelpSearchQuery(i))
	}
}

func (ptr *QHelpSearchEngine) __query_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___query_newList(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __searchResults_atList(i int) *QHelpSearchResult {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchEngine___searchResults_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __searchResults_setList(i QHelpSearchResult_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___searchResults_setList(ptr.Pointer(), PointerFromQHelpSearchResult(i))
	}
}

func (ptr *QHelpSearchEngine) __searchResults_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___searchResults_newList(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchEngine) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpSearchEngine___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpSearchEngine___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchEngine) __findChildren_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpSearchEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchEngine) __children_newList() unsafe.Pointer {
	return C.QHelpSearchEngine___children_newList(ptr.Pointer())
}

//export callbackQHelpSearchEngine_Event
func callbackQHelpSearchEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpSearchEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_EventFilter
func callbackQHelpSearchEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_ChildEvent
func callbackQHelpSearchEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchEngine_ConnectNotify
func callbackQHelpSearchEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_CustomEvent
func callbackQHelpSearchEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchEngine_DeleteLater
func callbackQHelpSearchEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpSearchEngine_Destroyed
func callbackQHelpSearchEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchEngine_DisconnectNotify
func callbackQHelpSearchEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_ObjectNameChanged
func callbackQHelpSearchEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchEngine_TimerEvent
func callbackQHelpSearchEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (ptr *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return ptr
}

func (ptr *QHelpSearchQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHelpSearchQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) (n *QHelpSearchQuery) {
	n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return
}

func (ptr *QHelpSearchQuery) DestroyQHelpSearchQuery() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQHelpSearchQuery() *QHelpSearchQuery {
	tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
	runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
}

func (ptr *QHelpSearchQuery) WordList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpSearchQuery_WordList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpSearchQuery) SetWordList(vqs []string) {
	if ptr.Pointer() != nil {
		vqsC := C.CString(strings.Join(vqs, "|"))
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQuery_SetWordList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vqsC, len: C.longlong(len(strings.Join(vqs, "|")))})
	}
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

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpSearchQueryWidget) {
	n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	return
}
func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	tmpValue := NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QHelpSearchQueryWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchQueryWidget_QHelpSearchQueryWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchQueryWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchQueryWidget_QHelpSearchQueryWidget_Tr(sC, cC, C.int(int32(n))))
}

func QHelpSearchQueryWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchQueryWidget_QHelpSearchQueryWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchQueryWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchQueryWidget_QHelpSearchQueryWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Search
func callbackQHelpSearchQueryWidget_Search(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "search") {
			C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "search"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "search", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "search", f)
		}
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "search")
	}
}

func (ptr *QHelpSearchQueryWidget) Search() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Search(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_SetCompactMode
func callbackQHelpSearchQueryWidget_SetCompactMode(ptr unsafe.Pointer, on C.char) {
	if signal := qt.GetSignal(ptr, "setCompactMode"); signal != nil {
		signal.(func(bool))(int8(on) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetCompactModeDefault(int8(on) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetCompactMode(f func(on bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCompactMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setCompactMode", func(on bool) {
				signal.(func(bool))(on)
				f(on)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCompactMode", f)
		}
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetCompactMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCompactMode")
	}
}

func (ptr *QHelpSearchQueryWidget) SetCompactMode(on bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetCompactMode(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetCompactModeDefault(on bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetCompactModeDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetSearchInput(searchInput string) {
	if ptr.Pointer() != nil {
		var searchInputC *C.char
		if searchInput != "" {
			searchInputC = C.CString(searchInput)
			defer C.free(unsafe.Pointer(searchInputC))
		}
		C.QHelpSearchQueryWidget_SetSearchInput(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchInputC, len: C.longlong(len(searchInput))})
	}
}

//export callbackQHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget
func callbackQHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpSearchQueryWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DestroyQHelpSearchQueryWidgetDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDestroyQHelpSearchQueryWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchQueryWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchQueryWidget", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchQueryWidget", f)
		}
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDestroyQHelpSearchQueryWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpSearchQueryWidget")
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchQueryWidget) SearchInput() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchQueryWidget_SearchInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_MetaObject
func callbackQHelpSearchQueryWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchQueryWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchQueryWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQueryWidget___setQuery_queryList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_setList(i QHelpSearchQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___setQuery_queryList_setList(ptr.Pointer(), PointerFromQHelpSearchQuery(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___setQuery_queryList_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQueryWidget___query_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __query_setList(i QHelpSearchQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___query_setList(ptr.Pointer(), PointerFromQHelpSearchQuery(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __query_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___query_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchQueryWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchQueryWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchQueryWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __actions_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchQueryWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchQueryWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchQueryWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchQueryWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __findChildren_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpSearchQueryWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchQueryWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchQueryWidget) __children_newList() unsafe.Pointer {
	return C.QHelpSearchQueryWidget___children_newList(ptr.Pointer())
}

//export callbackQHelpSearchQueryWidget_Close
func callbackQHelpSearchQueryWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchQueryWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_Event
func callbackQHelpSearchQueryWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_FocusNextPrevChild
func callbackQHelpSearchQueryWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_ActionEvent
func callbackQHelpSearchQueryWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ChangeEvent
func callbackQHelpSearchQueryWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_CloseEvent
func callbackQHelpSearchQueryWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ContextMenuEvent
func callbackQHelpSearchQueryWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_CustomContextMenuRequested
func callbackQHelpSearchQueryWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpSearchQueryWidget_DragEnterEvent
func callbackQHelpSearchQueryWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragLeaveEvent
func callbackQHelpSearchQueryWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragMoveEvent
func callbackQHelpSearchQueryWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DropEvent
func callbackQHelpSearchQueryWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_EnterEvent
func callbackQHelpSearchQueryWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusInEvent
func callbackQHelpSearchQueryWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusOutEvent
func callbackQHelpSearchQueryWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Hide
func callbackQHelpSearchQueryWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_HideEvent
func callbackQHelpSearchQueryWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodEvent
func callbackQHelpSearchQueryWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_KeyPressEvent
func callbackQHelpSearchQueryWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_KeyReleaseEvent
func callbackQHelpSearchQueryWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_LeaveEvent
func callbackQHelpSearchQueryWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Lower
func callbackQHelpSearchQueryWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_MouseDoubleClickEvent
func callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseMoveEvent
func callbackQHelpSearchQueryWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MousePressEvent
func callbackQHelpSearchQueryWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseReleaseEvent
func callbackQHelpSearchQueryWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MoveEvent
func callbackQHelpSearchQueryWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_PaintEvent
func callbackQHelpSearchQueryWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Raise
func callbackQHelpSearchQueryWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Repaint
func callbackQHelpSearchQueryWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ResizeEvent
func callbackQHelpSearchQueryWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SetDisabled
func callbackQHelpSearchQueryWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchQueryWidget_SetEnabled
func callbackQHelpSearchQueryWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetFocus2
func callbackQHelpSearchQueryWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchQueryWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_SetHidden
func callbackQHelpSearchQueryWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchQueryWidget_SetStyleSheet
func callbackQHelpSearchQueryWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QHelpSearchQueryWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQHelpSearchQueryWidget_SetVisible
func callbackQHelpSearchQueryWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowModified
func callbackQHelpSearchQueryWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowTitle
func callbackQHelpSearchQueryWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QHelpSearchQueryWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQHelpSearchQueryWidget_Show
func callbackQHelpSearchQueryWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowEvent
func callbackQHelpSearchQueryWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ShowFullScreen
func callbackQHelpSearchQueryWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMaximized
func callbackQHelpSearchQueryWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMinimized
func callbackQHelpSearchQueryWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowNormal
func callbackQHelpSearchQueryWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_TabletEvent
func callbackQHelpSearchQueryWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Update
func callbackQHelpSearchQueryWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_UpdateMicroFocus
func callbackQHelpSearchQueryWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_WheelEvent
func callbackQHelpSearchQueryWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_WindowIconChanged
func callbackQHelpSearchQueryWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpSearchQueryWidget_WindowTitleChanged
func callbackQHelpSearchQueryWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQHelpSearchQueryWidget_PaintEngine
func callbackQHelpSearchQueryWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpSearchQueryWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpSearchQueryWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_MinimumSizeHint
func callbackQHelpSearchQueryWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_SizeHint
func callbackQHelpSearchQueryWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_InputMethodQuery
func callbackQHelpSearchQueryWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchQueryWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_HasHeightForWidth
func callbackQHelpSearchQueryWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_HeightForWidth
func callbackQHelpSearchQueryWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchQueryWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_Metric
func callbackQHelpSearchQueryWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpSearchQueryWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_InitPainter
func callbackQHelpSearchQueryWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchQueryWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpSearchQueryWidget_EventFilter
func callbackQHelpSearchQueryWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_ChildEvent
func callbackQHelpSearchQueryWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ConnectNotify
func callbackQHelpSearchQueryWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_CustomEvent
func callbackQHelpSearchQueryWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DeleteLater
func callbackQHelpSearchQueryWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpSearchQueryWidget_Destroyed
func callbackQHelpSearchQueryWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchQueryWidget_DisconnectNotify
func callbackQHelpSearchQueryWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_ObjectNameChanged
func callbackQHelpSearchQueryWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchQueryWidget_TimerEvent
func callbackQHelpSearchQueryWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHelpSearchResult struct {
	ptr unsafe.Pointer
}

type QHelpSearchResult_ITF interface {
	QHelpSearchResult_PTR() *QHelpSearchResult
}

func (ptr *QHelpSearchResult) QHelpSearchResult_PTR() *QHelpSearchResult {
	return ptr
}

func (ptr *QHelpSearchResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHelpSearchResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHelpSearchResult(ptr QHelpSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResult_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultFromPointer(ptr unsafe.Pointer) (n *QHelpSearchResult) {
	n = new(QHelpSearchResult)
	n.SetPointer(ptr)
	return
}
func NewQHelpSearchResult() *QHelpSearchResult {
	tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchResult_NewQHelpSearchResult())
	runtime.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
	return tmpValue
}

func NewQHelpSearchResult2(other QHelpSearchResult_ITF) *QHelpSearchResult {
	tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchResult_NewQHelpSearchResult2(PointerFromQHelpSearchResult(other)))
	runtime.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
	return tmpValue
}

func NewQHelpSearchResult3(url core.QUrl_ITF, title string, snippet string) *QHelpSearchResult {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var snippetC *C.char
	if snippet != "" {
		snippetC = C.CString(snippet)
		defer C.free(unsafe.Pointer(snippetC))
	}
	tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchResult_NewQHelpSearchResult3(core.PointerFromQUrl(url), C.struct_QtHelp_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtHelp_PackedString{data: snippetC, len: C.longlong(len(snippet))}))
	runtime.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
	return tmpValue
}

func (ptr *QHelpSearchResult) DestroyQHelpSearchResult() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResult_DestroyQHelpSearchResult(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchResult) Snippet() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchResult_Snippet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpSearchResult) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchResult_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpSearchResult) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpSearchResult_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
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

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) (n *QHelpSearchResultWidget) {
	n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	return
}
func QHelpSearchResultWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchResultWidget_QHelpSearchResultWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchResultWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchResultWidget_QHelpSearchResultWidget_Tr(sC, cC, C.int(int32(n))))
}

func QHelpSearchResultWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchResultWidget_QHelpSearchResultWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchResultWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QHelpSearchResultWidget_QHelpSearchResultWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_RequestShowLink
func callbackQHelpSearchResultWidget_RequestShowLink(ptr unsafe.Pointer, link unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestShowLink") {
			C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestShowLink"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestShowLink", func(link *core.QUrl) {
				signal.(func(*core.QUrl))(link)
				f(link)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestShowLink", f)
		}
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "requestShowLink")
	}
}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RequestShowLink(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

//export callbackQHelpSearchResultWidget_DestroyQHelpSearchResultWidget
func callbackQHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpSearchResultWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DestroyQHelpSearchResultWidgetDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDestroyQHelpSearchResultWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchResultWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchResultWidget", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchResultWidget", f)
		}
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDestroyQHelpSearchResultWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHelpSearchResultWidget")
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpSearchResultWidget_MetaObject
func callbackQHelpSearchResultWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchResultWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchResultWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchResultWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchResultWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchResultWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchResultWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QHelpSearchResultWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QHelpSearchResultWidget) __actions_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___actions_newList(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchResultWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchResultWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchResultWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QHelpSearchResultWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchResultWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchResultWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpSearchResultWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchResultWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchResultWidget) __findChildren_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpSearchResultWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpSearchResultWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpSearchResultWidget) __children_newList() unsafe.Pointer {
	return C.QHelpSearchResultWidget___children_newList(ptr.Pointer())
}

//export callbackQHelpSearchResultWidget_Close
func callbackQHelpSearchResultWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchResultWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_Event
func callbackQHelpSearchResultWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_FocusNextPrevChild
func callbackQHelpSearchResultWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_ActionEvent
func callbackQHelpSearchResultWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ChangeEvent
func callbackQHelpSearchResultWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_CloseEvent
func callbackQHelpSearchResultWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ContextMenuEvent
func callbackQHelpSearchResultWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_CustomContextMenuRequested
func callbackQHelpSearchResultWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpSearchResultWidget_DragEnterEvent
func callbackQHelpSearchResultWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragLeaveEvent
func callbackQHelpSearchResultWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragMoveEvent
func callbackQHelpSearchResultWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DropEvent
func callbackQHelpSearchResultWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_EnterEvent
func callbackQHelpSearchResultWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusInEvent
func callbackQHelpSearchResultWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusOutEvent
func callbackQHelpSearchResultWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Hide
func callbackQHelpSearchResultWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchResultWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_HideEvent
func callbackQHelpSearchResultWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_InputMethodEvent
func callbackQHelpSearchResultWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_KeyPressEvent
func callbackQHelpSearchResultWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_KeyReleaseEvent
func callbackQHelpSearchResultWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_LeaveEvent
func callbackQHelpSearchResultWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Lower
func callbackQHelpSearchResultWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchResultWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_MouseDoubleClickEvent
func callbackQHelpSearchResultWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseMoveEvent
func callbackQHelpSearchResultWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MousePressEvent
func callbackQHelpSearchResultWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseReleaseEvent
func callbackQHelpSearchResultWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MoveEvent
func callbackQHelpSearchResultWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_PaintEvent
func callbackQHelpSearchResultWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Raise
func callbackQHelpSearchResultWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchResultWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_Repaint
func callbackQHelpSearchResultWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchResultWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ResizeEvent
func callbackQHelpSearchResultWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SetDisabled
func callbackQHelpSearchResultWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchResultWidget_SetEnabled
func callbackQHelpSearchResultWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetFocus2
func callbackQHelpSearchResultWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchResultWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_SetHidden
func callbackQHelpSearchResultWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchResultWidget_SetStyleSheet
func callbackQHelpSearchResultWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpSearchResultWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QHelpSearchResultWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQHelpSearchResultWidget_SetVisible
func callbackQHelpSearchResultWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowModified
func callbackQHelpSearchResultWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowTitle
func callbackQHelpSearchResultWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QHelpSearchResultWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQHelpSearchResultWidget_Show
func callbackQHelpSearchResultWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowEvent
func callbackQHelpSearchResultWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ShowFullScreen
func callbackQHelpSearchResultWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMaximized
func callbackQHelpSearchResultWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMinimized
func callbackQHelpSearchResultWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowNormal
func callbackQHelpSearchResultWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_TabletEvent
func callbackQHelpSearchResultWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Update
func callbackQHelpSearchResultWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchResultWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_UpdateMicroFocus
func callbackQHelpSearchResultWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_WheelEvent
func callbackQHelpSearchResultWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_WindowIconChanged
func callbackQHelpSearchResultWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpSearchResultWidget_WindowTitleChanged
func callbackQHelpSearchResultWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQHelpSearchResultWidget_PaintEngine
func callbackQHelpSearchResultWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQHelpSearchResultWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpSearchResultWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpSearchResultWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchResultWidget_MinimumSizeHint
func callbackQHelpSearchResultWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_SizeHint
func callbackQHelpSearchResultWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_InputMethodQuery
func callbackQHelpSearchResultWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchResultWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_HasHeightForWidth
func callbackQHelpSearchResultWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_HeightForWidth
func callbackQHelpSearchResultWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchResultWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_Metric
func callbackQHelpSearchResultWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpSearchResultWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_InitPainter
func callbackQHelpSearchResultWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpSearchResultWidget_EventFilter
func callbackQHelpSearchResultWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_ChildEvent
func callbackQHelpSearchResultWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ConnectNotify
func callbackQHelpSearchResultWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_CustomEvent
func callbackQHelpSearchResultWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DeleteLater
func callbackQHelpSearchResultWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQHelpSearchResultWidget_Destroyed
func callbackQHelpSearchResultWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchResultWidget_DisconnectNotify
func callbackQHelpSearchResultWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_ObjectNameChanged
func callbackQHelpSearchResultWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchResultWidget_TimerEvent
func callbackQHelpSearchResultWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
