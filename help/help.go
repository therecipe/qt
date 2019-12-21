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
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtHelp_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtHelp_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QCompressedHelpInfo struct {
	ptr unsafe.Pointer
}

type QCompressedHelpInfo_ITF interface {
	QCompressedHelpInfo_PTR() *QCompressedHelpInfo
}

func (ptr *QCompressedHelpInfo) QCompressedHelpInfo_PTR() *QCompressedHelpInfo {
	return ptr
}

func (ptr *QCompressedHelpInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCompressedHelpInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCompressedHelpInfo(ptr QCompressedHelpInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompressedHelpInfo_PTR().Pointer()
	}
	return nil
}

func NewQCompressedHelpInfoFromPointer(ptr unsafe.Pointer) (n *QCompressedHelpInfo) {
	n = new(QCompressedHelpInfo)
	n.SetPointer(ptr)
	return
}
func NewQCompressedHelpInfo() *QCompressedHelpInfo {
	tmpValue := NewQCompressedHelpInfoFromPointer(C.QCompressedHelpInfo_NewQCompressedHelpInfo())
	qt.SetFinalizer(tmpValue, (*QCompressedHelpInfo).DestroyQCompressedHelpInfo)
	return tmpValue
}

func NewQCompressedHelpInfo2(other QCompressedHelpInfo_ITF) *QCompressedHelpInfo {
	tmpValue := NewQCompressedHelpInfoFromPointer(C.QCompressedHelpInfo_NewQCompressedHelpInfo2(PointerFromQCompressedHelpInfo(other)))
	qt.SetFinalizer(tmpValue, (*QCompressedHelpInfo).DestroyQCompressedHelpInfo)
	return tmpValue
}

func NewQCompressedHelpInfo3(other QCompressedHelpInfo_ITF) *QCompressedHelpInfo {
	tmpValue := NewQCompressedHelpInfoFromPointer(C.QCompressedHelpInfo_NewQCompressedHelpInfo3(PointerFromQCompressedHelpInfo(other)))
	qt.SetFinalizer(tmpValue, (*QCompressedHelpInfo).DestroyQCompressedHelpInfo)
	return tmpValue
}

func (ptr *QCompressedHelpInfo) Component() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCompressedHelpInfo_Component(ptr.Pointer()))
	}
	return ""
}

func QCompressedHelpInfo_FromCompressedHelpFile(documentationFileName string) *QCompressedHelpInfo {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	tmpValue := NewQCompressedHelpInfoFromPointer(C.QCompressedHelpInfo_QCompressedHelpInfo_FromCompressedHelpFile(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}))
	qt.SetFinalizer(tmpValue, (*QCompressedHelpInfo).DestroyQCompressedHelpInfo)
	return tmpValue
}

func (ptr *QCompressedHelpInfo) FromCompressedHelpFile(documentationFileName string) *QCompressedHelpInfo {
	var documentationFileNameC *C.char
	if documentationFileName != "" {
		documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
	}
	tmpValue := NewQCompressedHelpInfoFromPointer(C.QCompressedHelpInfo_QCompressedHelpInfo_FromCompressedHelpFile(C.struct_QtHelp_PackedString{data: documentationFileNameC, len: C.longlong(len(documentationFileName))}))
	qt.SetFinalizer(tmpValue, (*QCompressedHelpInfo).DestroyQCompressedHelpInfo)
	return tmpValue
}

func (ptr *QCompressedHelpInfo) NamespaceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCompressedHelpInfo_NamespaceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCompressedHelpInfo) Swap(other QCompressedHelpInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QCompressedHelpInfo_Swap(ptr.Pointer(), PointerFromQCompressedHelpInfo(other))
	}
}

func (ptr *QCompressedHelpInfo) Version() *core.QVersionNumber {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVersionNumberFromPointer(C.QCompressedHelpInfo_Version(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVersionNumber).DestroyQVersionNumber)
		return tmpValue
	}
	return nil
}

func (ptr *QCompressedHelpInfo) DestroyQCompressedHelpInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCompressedHelpInfo_DestroyQCompressedHelpInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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
func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(int32(row))))
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

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_Row(ptr.Pointer())))
	}
	return 0
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

//export callbackQHelpContentModel_ColumnCount
func callbackQHelpContentModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpContentModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpContentModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnCount"); signal != nil {
			f := func(parent *core.QModelIndex) int {
				(*(*func(*core.QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
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

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

//export callbackQHelpContentModel_ContentsCreated
func callbackQHelpContentModel_ContentsCreated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsCreated"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsCreated") {
			C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "contentsCreated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsCreated"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "contentsCreated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreated", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsCreationStarted") {
			C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "contentsCreationStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsCreationStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "contentsCreationStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsCreationStarted", unsafe.Pointer(&f))
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

//export callbackQHelpContentModel_Data
func callbackQHelpContentModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant((*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpContentModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(index *core.QModelIndex, role int) *core.QVariant {
				(*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(index, role)
				return f(index, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
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
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Index
func callbackQHelpContentModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpContentModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "index"); signal != nil {
			f := func(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
				(*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(row, column, parent)
				return f(row, column, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

//export callbackQHelpContentModel_Parent
func callbackQHelpContentModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parent"); signal != nil {
			f := func(index *core.QModelIndex) *core.QModelIndex {
				(*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_RowCount
func callbackQHelpContentModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpContentModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpContentModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rowCount"); signal != nil {
			f := func(parent *core.QModelIndex) int {
				(*(*func(*core.QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
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

//export callbackQHelpContentModel_DestroyQHelpContentModel
func callbackQHelpContentModel_DestroyQHelpContentModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHelpContentModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHelpContentModelFromPointer(ptr).DestroyQHelpContentModelDefault()
	}
}

func (ptr *QHelpContentModel) ConnectDestroyQHelpContentModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpContentModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpContentModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpContentModel", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) DestroyQHelpContentModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpContentModel_DestroyQHelpContentModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpContentModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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

func (ptr *QHelpContentModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QHelpContentModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
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

func (ptr *QHelpContentModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpContentModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpContentModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpContentModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QHelpContentModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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

func (ptr *QHelpContentModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpContentModel_Buddy
func callbackQHelpContentModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_CanDropMimeData
func callbackQHelpContentModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_ColumnsAboutToBeInserted
func callbackQHelpContentModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsAboutToBeMoved
func callbackQHelpContentModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQHelpContentModel_ColumnsAboutToBeRemoved
func callbackQHelpContentModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsInserted
func callbackQHelpContentModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_ColumnsMoved
func callbackQHelpContentModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQHelpContentModel_ColumnsRemoved
func callbackQHelpContentModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_DataChanged
func callbackQHelpContentModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*core.QModelIndex, *core.QModelIndex, []int))(signal))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackQHelpContentModel_DropMimeData
func callbackQHelpContentModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_FetchMore
func callbackQHelpContentModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpContentModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpContentModel_Flags
func callbackQHelpContentModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*core.QModelIndex) core.Qt__ItemFlag)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpContentModel_HasChildren
func callbackQHelpContentModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_HeaderData
func callbackQHelpContentModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant((*(*func(int, core.Qt__Orientation, int) *core.QVariant)(signal))(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpContentModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_HeaderDataChanged
func callbackQHelpContentModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(core.Qt__Orientation, int, int))(signal))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_InsertColumns
func callbackQHelpContentModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_ItemData
func callbackQHelpContentModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*core.QModelIndex) map[int]*core.QVariant)(signal))(core.NewQModelIndexFromPointer(index)) {
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

//export callbackQHelpContentModel_LayoutAboutToBeChanged
func callbackQHelpContentModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
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
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpContentModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpContentModel_Match
func callbackQHelpContentModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
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

//export callbackQHelpContentModel_MimeData
func callbackQHelpContentModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtHelp_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData((*(*func([]*core.QModelIndex) *core.QMimeData)(signal))(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
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

//export callbackQHelpContentModel_MimeTypes
func callbackQHelpContentModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtHelp_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQHelpContentModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QHelpContentModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpContentModel_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQHelpContentModel_ModelAboutToBeReset
func callbackQHelpContentModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQHelpContentModel_ModelReset
func callbackQHelpContentModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQHelpContentModel_MoveColumns
func callbackQHelpContentModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpContentModel_ResetInternalData
func callbackQHelpContentModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_RoleNames
func callbackQHelpContentModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpContentModelFromPointer(NewQHelpContentModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*core.QByteArray)(signal))() {
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

//export callbackQHelpContentModel_RowsAboutToBeInserted
func callbackQHelpContentModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQHelpContentModel_RowsAboutToBeMoved
func callbackQHelpContentModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQHelpContentModel_RowsAboutToBeRemoved
func callbackQHelpContentModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_RowsInserted
func callbackQHelpContentModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_RowsMoved
func callbackQHelpContentModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQHelpContentModel_RowsRemoved
func callbackQHelpContentModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpContentModel_SetData
func callbackQHelpContentModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, *core.QVariant, int) bool)(signal))(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(int, core.Qt__Orientation, *core.QVariant, int) bool)(signal))(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, map[int]*core.QVariant) bool)(signal))(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
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

//export callbackQHelpContentModel_Sibling
func callbackQHelpContentModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Sort
func callbackQHelpContentModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, core.Qt__SortOrder))(signal))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpContentModel_Span
func callbackQHelpContentModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize((*(*func(*core.QModelIndex) *core.QSize)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpContentModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Submit
func callbackQHelpContentModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpContentModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentModel_SupportedDragActions
func callbackQHelpContentModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
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
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpContentModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_ChildEvent
func callbackQHelpContentModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpContentModel_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_Destroyed
func callbackQHelpContentModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpContentModel_DisconnectNotify
func callbackQHelpContentModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_Event
func callbackQHelpContentModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentModel_MetaObject
func callbackQHelpContentModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpContentModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpContentModel_ObjectNameChanged
func callbackQHelpContentModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpContentModel_TimerEvent
func callbackQHelpContentModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_LinkActivated
func callbackQHelpContentWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linkActivated"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkActivated") {
			C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "linkActivated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkActivated"); signal != nil {
			f := func(link *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(link)
				f(link)
			}
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", unsafe.Pointer(&f))
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpContentWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpContentWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpContentWidget_Collapse
func callbackQHelpContentWidget_Collapse(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapse"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func())(signal))()
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ColumnCountChanged
func callbackQHelpContentWidget_ColumnCountChanged(ptr unsafe.Pointer, oldCount C.int, newCount C.int) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(oldCount)), int(int32(newCount)))
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
		(*(*func())(signal))()
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
		(*(*func(int, int, int))(signal))(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
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
		(*(*func(*core.QModelIndex, *core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
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
		(*(*func(*core.QModelIndex, *core.QModelIndex, []int))(signal))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
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
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_DrawBranches
func callbackQHelpContentWidget_DrawBranches(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawBranches"); signal != nil {
		(*(*func(*gui.QPainter, *core.QRect, *core.QModelIndex))(signal))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
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
		(*(*func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(signal))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawRowDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) DrawRowDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRowDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Expand
func callbackQHelpContentWidget_Expand(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expand"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandAllDefault()
	}
}

func (ptr *QHelpContentWidget) ExpandAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ExpandRecursively
func callbackQHelpContentWidget_ExpandRecursively(ptr unsafe.Pointer, index unsafe.Pointer, depth C.int) {
	if signal := qt.GetSignal(ptr, "expandRecursively"); signal != nil {
		(*(*func(*core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(index), int(int32(depth)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandRecursivelyDefault(core.NewQModelIndexFromPointer(index), int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ExpandRecursivelyDefault(index core.QModelIndex_ITF, depth int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandRecursivelyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(depth)))
	}
}

//export callbackQHelpContentWidget_ExpandToDepth
func callbackQHelpContentWidget_ExpandToDepth(ptr unsafe.Pointer, depth C.int) {
	if signal := qt.GetSignal(ptr, "expandToDepth"); signal != nil {
		(*(*func(int))(signal))(int(int32(depth)))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_HideColumn
func callbackQHelpContentWidget_HideColumn(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "hideColumn"); signal != nil {
		(*(*func(int))(signal))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) HideColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_HorizontalOffset
func callbackQHelpContentWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpContentWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_IndexAt
func callbackQHelpContentWidget_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QPoint) *core.QModelIndex)(signal))(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QHelpContentWidget) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_IsIndexHidden
func callbackQHelpContentWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpContentWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_KeyPressEvent
func callbackQHelpContentWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(string))(signal))(cGoUnpackString(search))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MoveCursor
func callbackQHelpContentWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpContentWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_PaintEvent
func callbackQHelpContentWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func(int))(signal))(int(int32(column)))
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
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
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
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
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
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
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
		(*(*func(int, int))(signal))(int(int32(dx)), int(int32(dy)))
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
		(*(*func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(signal))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
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
		(*(*func())(signal))()
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
		(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
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
		(*(*func(*core.QAbstractItemModel))(signal))(core.NewQAbstractItemModelFromPointer(model))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
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
		(*(*func(*core.QItemSelectionModel))(signal))(core.NewQItemSelectionModelFromPointer(selectionModel))
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
		(*(*func(int))(signal))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ShowColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_SizeHintForColumn
func callbackQHelpContentWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForColumn"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(column)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpContentWidget) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpContentWidget_SortByColumn
func callbackQHelpContentWidget_SortByColumn(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sortByColumn"); signal != nil {
		(*(*func(int, core.Qt__SortOrder))(signal))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SortByColumnDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentWidget) SortByColumnDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SortByColumnDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpContentWidget_TimerEvent
func callbackQHelpContentWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_VerticalOffset
func callbackQHelpContentWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpContentWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_ViewportEvent
func callbackQHelpContentWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_ViewportSizeHint
func callbackQHelpContentWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpContentWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRect
func callbackQHelpContentWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect((*(*func(*core.QModelIndex) *core.QRect)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpContentWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHelpContentWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRegionForSelection
func callbackQHelpContentWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion((*(*func(*core.QItemSelection) *gui.QRegion)(signal))(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpContentWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpContentWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_Activated
func callbackQHelpContentWidget_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ClearSelection
func callbackQHelpContentWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_CloseEditor
func callbackQHelpContentWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		(*(*func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(signal))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
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
		(*(*func(*widgets.QWidget))(signal))(widgets.NewQWidgetFromPointer(editor))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_DragEnterEvent
func callbackQHelpContentWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
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
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(event))
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
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(event))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Edit2
func callbackQHelpContentWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(signal))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_EditorDestroyed
func callbackQHelpContentWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editorDestroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(editor))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_Event
func callbackQHelpContentWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_FocusInEvent
func callbackQHelpContentWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_FocusNextPrevChild
func callbackQHelpContentWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpContentWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_FocusOutEvent
func callbackQHelpContentWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
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
		(*(*func(*core.QSize))(signal))(core.NewQSizeFromPointer(size))
	}

}

//export callbackQHelpContentWidget_InputMethodEvent
func callbackQHelpContentWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpContentWidget_InputMethodQuery
func callbackQHelpContentWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpContentWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpContentWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_Pressed
func callbackQHelpContentWidget_Pressed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpContentWidget_ResizeEvent
func callbackQHelpContentWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpContentWidget) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionCommand
func callbackQHelpContentWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "selectionCommand"); signal != nil {
		return C.longlong((*(*func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(signal))(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpContentWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpContentWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpContentWidget_SetCurrentIndex
func callbackQHelpContentWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SizeHintForRow
func callbackQHelpContentWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForRow"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(row)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpContentWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpContentWidget_StartDrag
func callbackQHelpContentWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	if signal := qt.GetSignal(ptr, "startDrag"); signal != nil {
		(*(*func(core.Qt__DropAction))(signal))(core.Qt__DropAction(supportedActions))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ViewOptions
func callbackQHelpContentWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem((*(*func() *widgets.QStyleOptionViewItem)(signal))())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpContentWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpContentWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptionsDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ViewportEntered
func callbackQHelpContentWidget_ViewportEntered(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "viewportEntered"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQHelpContentWidget_ContextMenuEvent
func callbackQHelpContentWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpContentWidget_MinimumSizeHint
func callbackQHelpContentWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpContentWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_SetupViewport
func callbackQHelpContentWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		(*(*func(*widgets.QWidget))(signal))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpContentWidget_SizeHint
func callbackQHelpContentWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpContentWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_WheelEvent
func callbackQHelpContentWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpContentWidget_ChangeEvent
func callbackQHelpContentWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpContentWidget_ActionEvent
func callbackQHelpContentWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpContentWidget_Close
func callbackQHelpContentWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpContentWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpContentWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpContentWidget_CloseEvent
func callbackQHelpContentWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
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
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpContentWidget_EnterEvent
func callbackQHelpContentWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_HasHeightForWidth
func callbackQHelpContentWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
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
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpContentWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpContentWidget_Hide
func callbackQHelpContentWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpContentWidget_InitPainter
func callbackQHelpContentWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpContentWidget_KeyReleaseEvent
func callbackQHelpContentWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpContentWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Metric
func callbackQHelpContentWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpContentWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpContentWidget_MoveEvent
func callbackQHelpContentWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_NativeEvent
func callbackQHelpContentWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QHelpContentWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QHelpContentWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_PaintEngine
func callbackQHelpContentWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQHelpContentWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpContentWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpContentWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpContentWidget_Raise
func callbackQHelpContentWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(disable) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(hidden) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
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
		(*(*func(bool))(signal))(int8(visible) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(vqs))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpContentWidget_WindowTitleChanged
func callbackQHelpContentWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQHelpContentWidget_ChildEvent
func callbackQHelpContentWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpContentWidget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Destroyed
func callbackQHelpContentWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpContentWidget_DisconnectNotify
func callbackQHelpContentWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_MetaObject
func callbackQHelpContentWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpContentWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpContentWidget_ObjectNameChanged
func callbackQHelpContentWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
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
		(*(*func())(signal))()
	} else {
		NewQHelpEngineFromPointer(ptr).DestroyQHelpEngineDefault()
	}
}

func (ptr *QHelpEngine) ConnectDestroyQHelpEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngine", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngine) DestroyQHelpEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpEngine_DestroyQHelpEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
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

func (ptr *QHelpEngineCore) CurrentFilter() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(defaultValue)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
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

func (ptr *QHelpEngineCore) Error() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) Files2(namespaceName string, filterName string, extensionFilter string) []*core.QUrl {
	if ptr.Pointer() != nil {
		var namespaceNameC *C.char
		if namespaceName != "" {
			namespaceNameC = C.CString(namespaceName)
			defer C.free(unsafe.Pointer(namespaceNameC))
		}
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		var extensionFilterC *C.char
		if extensionFilter != "" {
			extensionFilterC = C.CString(extensionFilter)
			defer C.free(unsafe.Pointer(extensionFilterC))
		}
		return func(l C.struct_QtHelp_PackedList) []*core.QUrl {
			out := make([]*core.QUrl, int(l.len))
			tmpList := NewQHelpEngineCoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__files_atList2(i)
			}
			return out
		}(C.QHelpEngineCore_Files2(ptr.Pointer(), C.struct_QtHelp_PackedString{data: namespaceNameC, len: C.longlong(len(namespaceName))}, C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))}, C.struct_QtHelp_PackedString{data: extensionFilterC, len: C.longlong(len(extensionFilter))}))
	}
	return make([]*core.QUrl, 0)
}

func (ptr *QHelpEngineCore) FilterEngine() *QHelpFilterEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpFilterEngineFromPointer(C.QHelpEngineCore_FilterEngine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
	qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
	qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
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

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())))
	}
	return make([]string, 0)
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

func (ptr *QHelpEngineCore) SetUsesFilterEngine(uses bool) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetUsesFilterEngine(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(uses))))
	}
}

func (ptr *QHelpEngineCore) SetupData() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_SetupData(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpEngineCore_SetupFinished
func callbackQHelpEngineCore_SetupFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setupFinished") {
			C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "setupFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setupFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "setupFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setupFinished", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setupStarted") {
			C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "setupStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setupStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "setupStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setupStarted", unsafe.Pointer(&f))
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

func (ptr *QHelpEngineCore) UsesFilterEngine() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_UsesFilterEngine(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpEngineCore_Warning
func callbackQHelpEngineCore_Warning(ptr unsafe.Pointer, msg C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "warning"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(msg))
	}

}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "warning") {
			C.QHelpEngineCore_ConnectWarning(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "warning")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "warning"); signal != nil {
			f := func(msg string) {
				(*(*func(string))(signal))(msg)
				f(msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "warning", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "warning", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DestroyQHelpEngineCoreDefault()
	}
}

func (ptr *QHelpEngineCore) ConnectDestroyQHelpEngineCore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpEngineCore"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngineCore", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpEngineCore", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCoreDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpEngineCore_DestroyQHelpEngineCoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) __files_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore___files_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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

func (ptr *QHelpEngineCore) __files_atList2(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpEngineCore___files_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) __files_setList2(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore___files_setList2(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QHelpEngineCore) __files_newList2() unsafe.Pointer {
	return C.QHelpEngineCore___files_newList2(ptr.Pointer())
}

func (ptr *QHelpEngineCore) __filterAttributeSets_atList(i int) []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpEngineCore___filterAttributeSets_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) __filterAttributeSets_setList(i []string) {
	if ptr.Pointer() != nil {
		iC := C.CString(strings.Join(i, "¡¦!"))
		defer C.free(unsafe.Pointer(iC))
		C.QHelpEngineCore___filterAttributeSets_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(strings.Join(i, "¡¦!")))})
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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

func (ptr *QHelpEngineCore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpEngineCore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpEngineCore_ChildEvent
func callbackQHelpEngineCore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngineCore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpEngineCore_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpEngineCore_Destroyed
func callbackQHelpEngineCore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpEngineCore_DisconnectNotify
func callbackQHelpEngineCore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_Event
func callbackQHelpEngineCore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngineCore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpEngineCore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpEngineCore_MetaObject
func callbackQHelpEngineCore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineCoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngineCore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpEngineCore_ObjectNameChanged
func callbackQHelpEngineCore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpEngineCore_TimerEvent
func callbackQHelpEngineCore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHelpFilterData struct {
	ptr unsafe.Pointer
}

type QHelpFilterData_ITF interface {
	QHelpFilterData_PTR() *QHelpFilterData
}

func (ptr *QHelpFilterData) QHelpFilterData_PTR() *QHelpFilterData {
	return ptr
}

func (ptr *QHelpFilterData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHelpFilterData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHelpFilterData(ptr QHelpFilterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpFilterData_PTR().Pointer()
	}
	return nil
}

func NewQHelpFilterDataFromPointer(ptr unsafe.Pointer) (n *QHelpFilterData) {
	n = new(QHelpFilterData)
	n.SetPointer(ptr)
	return
}
func NewQHelpFilterData() *QHelpFilterData {
	tmpValue := NewQHelpFilterDataFromPointer(C.QHelpFilterData_NewQHelpFilterData())
	qt.SetFinalizer(tmpValue, (*QHelpFilterData).DestroyQHelpFilterData)
	return tmpValue
}

func NewQHelpFilterData2(other QHelpFilterData_ITF) *QHelpFilterData {
	tmpValue := NewQHelpFilterDataFromPointer(C.QHelpFilterData_NewQHelpFilterData2(PointerFromQHelpFilterData(other)))
	qt.SetFinalizer(tmpValue, (*QHelpFilterData).DestroyQHelpFilterData)
	return tmpValue
}

func NewQHelpFilterData3(other QHelpFilterData_ITF) *QHelpFilterData {
	tmpValue := NewQHelpFilterDataFromPointer(C.QHelpFilterData_NewQHelpFilterData3(PointerFromQHelpFilterData(other)))
	qt.SetFinalizer(tmpValue, (*QHelpFilterData).DestroyQHelpFilterData)
	return tmpValue
}

func (ptr *QHelpFilterData) Components() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpFilterData_Components(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QHelpFilterData) SetComponents(components []string) {
	if ptr.Pointer() != nil {
		componentsC := C.CString(strings.Join(components, "¡¦!"))
		defer C.free(unsafe.Pointer(componentsC))
		C.QHelpFilterData_SetComponents(ptr.Pointer(), C.struct_QtHelp_PackedString{data: componentsC, len: C.longlong(len(strings.Join(components, "¡¦!")))})
	}
}

func (ptr *QHelpFilterData) SetVersions(versions []*core.QVersionNumber) {
	if ptr.Pointer() != nil {
		C.QHelpFilterData_SetVersions(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQHelpFilterDataFromPointer(NewQHelpFilterDataFromPointer(nil).__setVersions_versions_newList())
			for _, v := range versions {
				tmpList.__setVersions_versions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QHelpFilterData) Versions() []*core.QVersionNumber {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []*core.QVersionNumber {
			out := make([]*core.QVersionNumber, int(l.len))
			tmpList := NewQHelpFilterDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__versions_atList(i)
			}
			return out
		}(C.QHelpFilterData_Versions(ptr.Pointer()))
	}
	return make([]*core.QVersionNumber, 0)
}

func (ptr *QHelpFilterData) DestroyQHelpFilterData() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpFilterData_DestroyQHelpFilterData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpFilterData) __setVersions_versions_atList(i int) *core.QVersionNumber {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVersionNumberFromPointer(C.QHelpFilterData___setVersions_versions_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVersionNumber).DestroyQVersionNumber)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterData) __setVersions_versions_setList(i core.QVersionNumber_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterData___setVersions_versions_setList(ptr.Pointer(), core.PointerFromQVersionNumber(i))
	}
}

func (ptr *QHelpFilterData) __setVersions_versions_newList() unsafe.Pointer {
	return C.QHelpFilterData___setVersions_versions_newList(ptr.Pointer())
}

func (ptr *QHelpFilterData) __versions_atList(i int) *core.QVersionNumber {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVersionNumberFromPointer(C.QHelpFilterData___versions_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVersionNumber).DestroyQVersionNumber)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterData) __versions_setList(i core.QVersionNumber_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterData___versions_setList(ptr.Pointer(), core.PointerFromQVersionNumber(i))
	}
}

func (ptr *QHelpFilterData) __versions_newList() unsafe.Pointer {
	return C.QHelpFilterData___versions_newList(ptr.Pointer())
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

func NewQHelpFilterEngineFromPointer(ptr unsafe.Pointer) (n *QHelpFilterEngine) {
	n = new(QHelpFilterEngine)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpFilterEngine) ActiveFilter() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpFilterEngine_ActiveFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpFilterEngine) AvailableComponents() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpFilterEngine_AvailableComponents(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQHelpFilterEngine_FilterActivated
func callbackQHelpFilterEngine_FilterActivated(ptr unsafe.Pointer, newFilter C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "filterActivated"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(newFilter))
	}

}

func (ptr *QHelpFilterEngine) ConnectFilterActivated(f func(newFilter string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "filterActivated") {
			C.QHelpFilterEngine_ConnectFilterActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "filterActivated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "filterActivated"); signal != nil {
			f := func(newFilter string) {
				(*(*func(string))(signal))(newFilter)
				f(newFilter)
			}
			qt.ConnectSignal(ptr.Pointer(), "filterActivated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filterActivated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHelpFilterEngine) DisconnectFilterActivated() {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_DisconnectFilterActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "filterActivated")
	}
}

func (ptr *QHelpFilterEngine) FilterActivated(newFilter string) {
	if ptr.Pointer() != nil {
		var newFilterC *C.char
		if newFilter != "" {
			newFilterC = C.CString(newFilter)
			defer C.free(unsafe.Pointer(newFilterC))
		}
		C.QHelpFilterEngine_FilterActivated(ptr.Pointer(), C.struct_QtHelp_PackedString{data: newFilterC, len: C.longlong(len(newFilter))})
	}
}

func (ptr *QHelpFilterEngine) FilterData(filterName string) *QHelpFilterData {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		tmpValue := NewQHelpFilterDataFromPointer(C.QHelpFilterEngine_FilterData(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))}))
		qt.SetFinalizer(tmpValue, (*QHelpFilterData).DestroyQHelpFilterData)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) Filters() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpFilterEngine_Filters(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QHelpFilterEngine) NamespaceToComponent() map[string]string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[string]string {
			out := make(map[string]string, int(l.len))
			tmpList := NewQHelpFilterEngineFromPointer(l.data)
			for i, v := range tmpList.__namespaceToComponent_keyList() {
				out[v] = tmpList.__namespaceToComponent_atList(v, i)
			}
			return out
		}(C.QHelpFilterEngine_NamespaceToComponent(ptr.Pointer()))
	}
	return make(map[string]string, 0)
}

func (ptr *QHelpFilterEngine) NamespaceToVersion() map[string]*core.QVersionNumber {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) map[string]*core.QVersionNumber {
			out := make(map[string]*core.QVersionNumber, int(l.len))
			tmpList := NewQHelpFilterEngineFromPointer(l.data)
			for i, v := range tmpList.__namespaceToVersion_keyList() {
				out[v] = tmpList.__namespaceToVersion_atList(v, i)
			}
			return out
		}(C.QHelpFilterEngine_NamespaceToVersion(ptr.Pointer()))
	}
	return make(map[string]*core.QVersionNumber, 0)
}

func (ptr *QHelpFilterEngine) NamespacesForFilter(filterName string) []string {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return unpackStringList(cGoUnpackString(C.QHelpFilterEngine_NamespacesForFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})))
	}
	return make([]string, 0)
}

func (ptr *QHelpFilterEngine) RemoveFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return int8(C.QHelpFilterEngine_RemoveFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})) != 0
	}
	return false
}

func (ptr *QHelpFilterEngine) SetActiveFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return int8(C.QHelpFilterEngine_SetActiveFilter(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))})) != 0
	}
	return false
}

func (ptr *QHelpFilterEngine) SetFilterData(filterName string, filterData QHelpFilterData_ITF) bool {
	if ptr.Pointer() != nil {
		var filterNameC *C.char
		if filterName != "" {
			filterNameC = C.CString(filterName)
			defer C.free(unsafe.Pointer(filterNameC))
		}
		return int8(C.QHelpFilterEngine_SetFilterData(ptr.Pointer(), C.struct_QtHelp_PackedString{data: filterNameC, len: C.longlong(len(filterName))}, PointerFromQHelpFilterData(filterData))) != 0
	}
	return false
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_atList(v string, i int) string {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		return cGoUnpackString(C.QHelpFilterEngine___namespaceToComponent_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_setList(key string, i string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpFilterEngine___namespaceToComponent_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_newList() unsafe.Pointer {
	return C.QHelpFilterEngine___namespaceToComponent_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __namespaceToComponent_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpFilterEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____namespaceToComponent_keyList_atList(i)
			}
			return out
		}(C.QHelpFilterEngine___namespaceToComponent_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_atList(v string, i int) *core.QVersionNumber {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVersionNumberFromPointer(C.QHelpFilterEngine___namespaceToVersion_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVersionNumber).DestroyQVersionNumber)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_setList(key string, i core.QVersionNumber_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QHelpFilterEngine___namespaceToVersion_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVersionNumber(i))
	}
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_newList() unsafe.Pointer {
	return C.QHelpFilterEngine___namespaceToVersion_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __namespaceToVersion_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtHelp_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQHelpFilterEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____namespaceToVersion_keyList_atList(i)
			}
			return out
		}(C.QHelpFilterEngine___namespaceToVersion_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpFilterEngine_____namespaceToComponent_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpFilterEngine_____namespaceToComponent_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpFilterEngine) ____namespaceToComponent_keyList_newList() unsafe.Pointer {
	return C.QHelpFilterEngine_____namespaceToComponent_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpFilterEngine_____namespaceToVersion_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QHelpFilterEngine_____namespaceToVersion_keyList_setList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QHelpFilterEngine) ____namespaceToVersion_keyList_newList() unsafe.Pointer {
	return C.QHelpFilterEngine_____namespaceToVersion_keyList_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpFilterEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpFilterEngine) __children_newList() unsafe.Pointer {
	return C.QHelpFilterEngine___children_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpFilterEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHelpFilterEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHelpFilterEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpFilterEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpFilterEngine) __findChildren_newList() unsafe.Pointer {
	return C.QHelpFilterEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QHelpFilterEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHelpFilterEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpFilterEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHelpFilterEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QHelpFilterEngine___findChildren_newList3(ptr.Pointer())
}

//export callbackQHelpFilterEngine_ChildEvent
func callbackQHelpFilterEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpFilterEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpFilterEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpFilterEngine_ConnectNotify
func callbackQHelpFilterEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpFilterEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpFilterEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpFilterEngine_CustomEvent
func callbackQHelpFilterEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpFilterEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpFilterEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpFilterEngine_DeleteLater
func callbackQHelpFilterEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHelpFilterEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpFilterEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpFilterEngine_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpFilterEngine_Destroyed
func callbackQHelpFilterEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpFilterEngine_DisconnectNotify
func callbackQHelpFilterEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpFilterEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpFilterEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpFilterEngine_Event
func callbackQHelpFilterEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpFilterEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpFilterEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpFilterEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpFilterEngine_EventFilter
func callbackQHelpFilterEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpFilterEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpFilterEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpFilterEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpFilterEngine_MetaObject
func callbackQHelpFilterEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpFilterEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpFilterEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpFilterEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpFilterEngine_ObjectNameChanged
func callbackQHelpFilterEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpFilterEngine_TimerEvent
func callbackQHelpFilterEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpFilterEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpFilterEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpFilterEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHelpGlobal struct {
	ptr unsafe.Pointer
}

type QHelpGlobal_ITF interface {
	QHelpGlobal_PTR() *QHelpGlobal
}

func (ptr *QHelpGlobal) QHelpGlobal_PTR() *QHelpGlobal {
	return ptr
}

func (ptr *QHelpGlobal) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHelpGlobal) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHelpGlobal(ptr QHelpGlobal_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpGlobal_PTR().Pointer()
	}
	return nil
}

func NewQHelpGlobalFromPointer(ptr unsafe.Pointer) (n *QHelpGlobal) {
	n = new(QHelpGlobal)
	n.SetPointer(ptr)
	return
}
func (ptr *QHelpGlobal) DestroyQHelpGlobal() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_IndexCreated
func callbackQHelpIndexModel_IndexCreated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "indexCreated"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexCreated") {
			C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "indexCreated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexCreated"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "indexCreated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexCreated", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexCreationStarted") {
			C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "indexCreationStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexCreationStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "indexCreationStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexCreationStarted", unsafe.Pointer(&f))
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

func (ptr *QHelpIndexModel) __linksForKeyword_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpIndexModel___linksForKeyword_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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

func (ptr *QHelpIndexModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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

func (ptr *QHelpIndexModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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

func (ptr *QHelpIndexModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
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

func (ptr *QHelpIndexModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpIndexModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpIndexModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpIndexModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QHelpIndexModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpIndexModel_Data
func callbackQHelpIndexModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant((*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpIndexModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Flags
func callbackQHelpIndexModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*core.QModelIndex) core.Qt__ItemFlag)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpIndexModel_InsertRows
func callbackQHelpIndexModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ItemData
func callbackQHelpIndexModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*core.QModelIndex) map[int]*core.QVariant)(signal))(core.NewQModelIndexFromPointer(index)) {
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

//export callbackQHelpIndexModel_MoveRows
func callbackQHelpIndexModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RemoveRows
func callbackQHelpIndexModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RowCount
func callbackQHelpIndexModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_SetData
func callbackQHelpIndexModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, *core.QVariant, int) bool)(signal))(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SetItemData
func callbackQHelpIndexModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_QtHelp_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, map[int]*core.QVariant) bool)(signal))(core.NewQModelIndexFromPointer(index), func(l C.struct_QtHelp_PackedList) map[int]*core.QVariant {
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

//export callbackQHelpIndexModel_Sibling
func callbackQHelpIndexModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QHelpIndexModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Sort
func callbackQHelpIndexModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, core.Qt__SortOrder))(signal))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpIndexModel_SupportedDropActions
func callbackQHelpIndexModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpIndexModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_DropMimeData
func callbackQHelpIndexModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
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
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpIndexModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Buddy
func callbackQHelpIndexModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_CanDropMimeData
func callbackQHelpIndexModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ColumnCount
func callbackQHelpIndexModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_ColumnsAboutToBeInserted
func callbackQHelpIndexModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsAboutToBeMoved
func callbackQHelpIndexModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQHelpIndexModel_ColumnsAboutToBeRemoved
func callbackQHelpIndexModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsInserted
func callbackQHelpIndexModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_ColumnsMoved
func callbackQHelpIndexModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQHelpIndexModel_ColumnsRemoved
func callbackQHelpIndexModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_DataChanged
func callbackQHelpIndexModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtHelp_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*core.QModelIndex, *core.QModelIndex, []int))(signal))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpIndexModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpIndexModel_HasChildren
func callbackQHelpIndexModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_HeaderData
func callbackQHelpIndexModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant((*(*func(int, core.Qt__Orientation, int) *core.QVariant)(signal))(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpIndexModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_HeaderDataChanged
func callbackQHelpIndexModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(core.Qt__Orientation, int, int))(signal))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_InsertColumns
func callbackQHelpIndexModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_LayoutAboutToBeChanged
func callbackQHelpIndexModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtHelp_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
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
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtHelp_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQHelpIndexModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQHelpIndexModel_Match
func callbackQHelpIndexModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
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

//export callbackQHelpIndexModel_MimeData
func callbackQHelpIndexModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtHelp_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData((*(*func([]*core.QModelIndex) *core.QMimeData)(signal))(func(l C.struct_QtHelp_PackedList) []*core.QModelIndex {
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

//export callbackQHelpIndexModel_MimeTypes
func callbackQHelpIndexModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtHelp_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQHelpIndexModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtHelp_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QHelpIndexModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QHelpIndexModel_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQHelpIndexModel_ModelAboutToBeReset
func callbackQHelpIndexModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQHelpIndexModel_ModelReset
func callbackQHelpIndexModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQHelpIndexModel_MoveColumns
func callbackQHelpIndexModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Parent
func callbackQHelpIndexModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_RemoveColumns
func callbackQHelpIndexModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ResetInternalData
func callbackQHelpIndexModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_RoleNames
func callbackQHelpIndexModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexModelFromPointer(NewQHelpIndexModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*core.QByteArray)(signal))() {
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

//export callbackQHelpIndexModel_RowsAboutToBeInserted
func callbackQHelpIndexModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQHelpIndexModel_RowsAboutToBeMoved
func callbackQHelpIndexModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQHelpIndexModel_RowsAboutToBeRemoved
func callbackQHelpIndexModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_RowsInserted
func callbackQHelpIndexModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_RowsMoved
func callbackQHelpIndexModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQHelpIndexModel_RowsRemoved
func callbackQHelpIndexModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQHelpIndexModel_SetHeaderData
func callbackQHelpIndexModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, core.Qt__Orientation, *core.QVariant, int) bool)(signal))(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Span
func callbackQHelpIndexModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize((*(*func(*core.QModelIndex) *core.QSize)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpIndexModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Submit
func callbackQHelpIndexModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpIndexModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SupportedDragActions
func callbackQHelpIndexModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpIndexModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_ChildEvent
func callbackQHelpIndexModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpIndexModel_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_Destroyed
func callbackQHelpIndexModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpIndexModel_DisconnectNotify
func callbackQHelpIndexModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_Event
func callbackQHelpIndexModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MetaObject
func callbackQHelpIndexModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpIndexModel_ObjectNameChanged
func callbackQHelpIndexModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpIndexModel_TimerEvent
func callbackQHelpIndexModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
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

//export callbackQHelpIndexWidget_ActivateCurrentItem
func callbackQHelpIndexWidget_ActivateCurrentItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activateCurrentItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActivateCurrentItemDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectActivateCurrentItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activateCurrentItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "activateCurrentItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activateCurrentItem", unsafe.Pointer(&f))
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
		(*(*func(string, string))(signal))(cGoUnpackString(filter), cGoUnpackString(wildcard))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FilterIndicesDefault(cGoUnpackString(filter), cGoUnpackString(wildcard))
	}
}

func (ptr *QHelpIndexWidget) ConnectFilterIndices(f func(filter string, wildcard string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filterIndices"); signal != nil {
			f := func(filter string, wildcard string) {
				(*(*func(string, string))(signal))(filter, wildcard)
				f(filter, wildcard)
			}
			qt.ConnectSignal(ptr.Pointer(), "filterIndices", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filterIndices", unsafe.Pointer(&f))
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
		(*(*func(*core.QUrl, string))(signal))(core.NewQUrlFromPointer(link), cGoUnpackString(keyword))
	}

}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkActivated") {
			C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "linkActivated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkActivated"); signal != nil {
			f := func(link *core.QUrl, keyword string) {
				(*(*func(*core.QUrl, string))(signal))(link, keyword)
				f(link, keyword)
			}
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkActivated", unsafe.Pointer(&f))
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
		(*(*func(map[string]*core.QUrl, string))(signal))(func(l C.struct_QtHelp_PackedList) map[string]*core.QUrl {
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
			C.QHelpIndexWidget_ConnectLinksActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "linksActivated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linksActivated"); signal != nil {
			f := func(links map[string]*core.QUrl, keyword string) {
				(*(*func(map[string]*core.QUrl, string))(signal))(links, keyword)
				f(links, keyword)
			}
			qt.ConnectSignal(ptr.Pointer(), "linksActivated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linksActivated", unsafe.Pointer(&f))
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

func (ptr *QHelpIndexWidget) __linksActivated_links_atList(v string, i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQUrlFromPointer(C.QHelpIndexWidget___linksActivated_links_atList(ptr.Pointer(), C.struct_QtHelp_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
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

func (ptr *QHelpIndexWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpIndexWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpIndexWidget_CurrentChanged
func callbackQHelpIndexWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		(*(*func(*core.QModelIndex, *core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
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
		(*(*func(*core.QModelIndex, *core.QModelIndex, []int))(signal))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtHelp_PackedList) []int {
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
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(e))
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
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(e))
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
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQHelpIndexWidget_Event
func callbackQHelpIndexWidget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpIndexWidget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_HorizontalOffset
func callbackQHelpIndexWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_IndexAt
func callbackQHelpIndexWidget_IndexAt(ptr unsafe.Pointer, p unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QPoint) *core.QModelIndex)(signal))(core.NewQPointFromPointer(p)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(p)))
}

func (ptr *QHelpIndexWidget) IndexAtDefault(p core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(p)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_IsIndexHidden
func callbackQHelpIndexWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpIndexWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_MouseMoveEvent
func callbackQHelpIndexWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(e))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MoveCursor
func callbackQHelpIndexWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpIndexWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_PaintEvent
func callbackQHelpIndexWidget_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(e))
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
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(e))
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
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
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
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
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
		(*(*func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(signal))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_SelectedIndexes
func callbackQHelpIndexWidget_SelectedIndexes(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectedIndexes"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQHelpIndexWidgetFromPointer(NewQHelpIndexWidgetFromPointer(nil).__selectedIndexes_newList())
			for _, v := range (*(*func() []*core.QModelIndex)(signal))() {
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

//export callbackQHelpIndexWidget_SelectionChanged
func callbackQHelpIndexWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
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
		(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
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
		(*(*func(core.Qt__DropAction))(signal))(core.Qt__DropAction(supportedActions))
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
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(e))
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_VerticalOffset
func callbackQHelpIndexWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_ViewOptions
func callbackQHelpIndexWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem((*(*func() *widgets.QStyleOptionViewItem)(signal))())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpIndexWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpIndexWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptionsDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ViewportSizeHint
func callbackQHelpIndexWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRect
func callbackQHelpIndexWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect((*(*func(*core.QModelIndex) *core.QRect)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpIndexWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRegionForSelection
func callbackQHelpIndexWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion((*(*func(*core.QItemSelection) *gui.QRegion)(signal))(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpIndexWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpIndexWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_WheelEvent
func callbackQHelpIndexWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpIndexWidget_Activated
func callbackQHelpIndexWidget_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_ClearSelection
func callbackQHelpIndexWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_CloseEditor
func callbackQHelpIndexWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		(*(*func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(signal))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
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
		(*(*func(*widgets.QWidget))(signal))(widgets.NewQWidgetFromPointer(editor))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_DragEnterEvent
func callbackQHelpIndexWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_Edit2
func callbackQHelpIndexWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(signal))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_EditorDestroyed
func callbackQHelpIndexWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editorDestroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(editor))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_EventFilter
func callbackQHelpIndexWidget_EventFilter(ptr unsafe.Pointer, object unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(object), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) EventFilterDefault(object core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_FocusInEvent
func callbackQHelpIndexWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_FocusNextPrevChild
func callbackQHelpIndexWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpIndexWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_FocusOutEvent
func callbackQHelpIndexWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
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
		(*(*func(*core.QSize))(signal))(core.NewQSizeFromPointer(size))
	}

}

//export callbackQHelpIndexWidget_InputMethodEvent
func callbackQHelpIndexWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpIndexWidget_InputMethodQuery
func callbackQHelpIndexWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpIndexWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpIndexWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_KeyPressEvent
func callbackQHelpIndexWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(string))(signal))(cGoUnpackString(search))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQHelpIndexWidget_Reset
func callbackQHelpIndexWidget_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectionCommand
func callbackQHelpIndexWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "selectionCommand"); signal != nil {
		return C.longlong((*(*func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(signal))(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpIndexWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpIndexWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpIndexWidget_SetCurrentIndex
func callbackQHelpIndexWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func(*core.QAbstractItemModel))(signal))(core.NewQAbstractItemModelFromPointer(model))
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
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func(*core.QItemSelectionModel))(signal))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpIndexWidget_SizeHintForColumn
func callbackQHelpIndexWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForColumn"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(column)))))
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
		return C.int(int32((*(*func(int) int)(signal))(int(int32(row)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpIndexWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Update
func callbackQHelpIndexWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
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
		(*(*func())(signal))()
	}

}

//export callbackQHelpIndexWidget_ViewportEvent
func callbackQHelpIndexWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_ContextMenuEvent
func callbackQHelpIndexWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpIndexWidget_MinimumSizeHint
func callbackQHelpIndexWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpIndexWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ScrollContentsBy
func callbackQHelpIndexWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(dx)), int(int32(dy)))
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
		(*(*func(*widgets.QWidget))(signal))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpIndexWidget_SizeHint
func callbackQHelpIndexWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpIndexWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ChangeEvent
func callbackQHelpIndexWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpIndexWidget_ActionEvent
func callbackQHelpIndexWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpIndexWidget_Close
func callbackQHelpIndexWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpIndexWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpIndexWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_CloseEvent
func callbackQHelpIndexWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
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
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpIndexWidget_EnterEvent
func callbackQHelpIndexWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_HasHeightForWidth
func callbackQHelpIndexWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
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
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpIndexWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Hide
func callbackQHelpIndexWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpIndexWidget_InitPainter
func callbackQHelpIndexWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpIndexWidget_KeyReleaseEvent
func callbackQHelpIndexWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpIndexWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Metric
func callbackQHelpIndexWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpIndexWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpIndexWidget_MoveEvent
func callbackQHelpIndexWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpIndexWidget_NativeEvent
func callbackQHelpIndexWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QHelpIndexWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QHelpIndexWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_PaintEngine
func callbackQHelpIndexWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQHelpIndexWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpIndexWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpIndexWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpIndexWidget_Raise
func callbackQHelpIndexWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(disable) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(hidden) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
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
		(*(*func(bool))(signal))(int8(visible) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(vqs))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpIndexWidget_WindowTitleChanged
func callbackQHelpIndexWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQHelpIndexWidget_ChildEvent
func callbackQHelpIndexWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpIndexWidget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Destroyed
func callbackQHelpIndexWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpIndexWidget_DisconnectNotify
func callbackQHelpIndexWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_MetaObject
func callbackQHelpIndexWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpIndexWidget_ObjectNameChanged
func callbackQHelpIndexWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
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

//export callbackQHelpSearchEngine_CancelIndexing
func callbackQHelpSearchEngine_CancelIndexing(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cancelIndexing"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CancelIndexingDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectCancelIndexing(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cancelIndexing"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cancelIndexing", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cancelIndexing", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CancelSearchingDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectCancelSearching(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cancelSearching"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cancelSearching", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cancelSearching", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexingFinished") {
			C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "indexingFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexingFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "indexingFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexingFinished", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "indexingStarted") {
			C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "indexingStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "indexingStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "indexingStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexingStarted", unsafe.Pointer(&f))
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

//export callbackQHelpSearchEngine_ReindexDocumentation
func callbackQHelpSearchEngine_ReindexDocumentation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reindexDocumentation"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ReindexDocumentationDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectReindexDocumentation(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reindexDocumentation"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "reindexDocumentation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reindexDocumentation", unsafe.Pointer(&f))
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

//export callbackQHelpSearchEngine_Search2
func callbackQHelpSearchEngine_Search2(ptr unsafe.Pointer, searchInput C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "search2"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(searchInput))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).Search2Default(cGoUnpackString(searchInput))
	}
}

func (ptr *QHelpSearchEngine) ConnectSearch2(f func(searchInput string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "search2"); signal != nil {
			f := func(searchInput string) {
				(*(*func(string))(signal))(searchInput)
				f(searchInput)
			}
			qt.ConnectSignal(ptr.Pointer(), "search2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "search2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearch2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "search2")
	}
}

func (ptr *QHelpSearchEngine) Search2(searchInput string) {
	if ptr.Pointer() != nil {
		var searchInputC *C.char
		if searchInput != "" {
			searchInputC = C.CString(searchInput)
			defer C.free(unsafe.Pointer(searchInputC))
		}
		C.QHelpSearchEngine_Search2(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchInputC, len: C.longlong(len(searchInput))})
	}
}

func (ptr *QHelpSearchEngine) Search2Default(searchInput string) {
	if ptr.Pointer() != nil {
		var searchInputC *C.char
		if searchInput != "" {
			searchInputC = C.CString(searchInput)
			defer C.free(unsafe.Pointer(searchInputC))
		}
		C.QHelpSearchEngine_Search2Default(ptr.Pointer(), C.struct_QtHelp_PackedString{data: searchInputC, len: C.longlong(len(searchInput))})
	}
}

func (ptr *QHelpSearchEngine) SearchInput() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchEngine_SearchInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpSearchEngine) SearchResultCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchEngine_SearchResultCount(ptr.Pointer())))
	}
	return 0
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

//export callbackQHelpSearchEngine_SearchingFinished
func callbackQHelpSearchEngine_SearchingFinished(ptr unsafe.Pointer, searchResultCount C.int) {
	if signal := qt.GetSignal(ptr, "searchingFinished"); signal != nil {
		(*(*func(int))(signal))(int(int32(searchResultCount)))
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(searchResultCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "searchingFinished") {
			C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "searchingFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "searchingFinished"); signal != nil {
			f := func(searchResultCount int) {
				(*(*func(int))(signal))(searchResultCount)
				f(searchResultCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "searchingFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "searchingFinished", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "searchingStarted") {
			C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "searchingStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "searchingStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "searchingStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "searchingStarted", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DestroyQHelpSearchEngineDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectDestroyQHelpSearchEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchEngine", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchEngine_DestroyQHelpSearchEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) __query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchEngine___query_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
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

func (ptr *QHelpSearchEngine) __search_queryList_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchEngine___search_queryList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
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

func (ptr *QHelpSearchEngine) __searchResults_atList(i int) *QHelpSearchResult {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchEngine___searchResults_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
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

func (ptr *QHelpSearchEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpSearchEngine_ChildEvent
func callbackQHelpSearchEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchEngine_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_Destroyed
func callbackQHelpSearchEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchEngine_DisconnectNotify
func callbackQHelpSearchEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_Event
func callbackQHelpSearchEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_MetaObject
func callbackQHelpSearchEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchEngine_ObjectNameChanged
func callbackQHelpSearchEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchEngine_TimerEvent
func callbackQHelpSearchEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQHelpSearchQuery() *QHelpSearchQuery {
	tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
	qt.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
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
		(*(*func())(signal))()
	}

}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "search") {
			C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "search")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "search"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "search", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "search", unsafe.Pointer(&f))
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

func (ptr *QHelpSearchQueryWidget) SearchInput() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpSearchQueryWidget_SearchInput(ptr.Pointer()))
	}
	return ""
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DestroyQHelpSearchQueryWidgetDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDestroyQHelpSearchQueryWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchQueryWidget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchQueryWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchQueryWidget", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidgetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) __query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQueryWidget___query_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
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

func (ptr *QHelpSearchQueryWidget) __setQuery_queryList_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQHelpSearchQueryFromPointer(C.QHelpSearchQueryWidget___setQuery_queryList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
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

func (ptr *QHelpSearchQueryWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchQueryWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpSearchQueryWidget_ActionEvent
func callbackQHelpSearchQueryWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Close
func callbackQHelpSearchQueryWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchQueryWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_CloseEvent
func callbackQHelpSearchQueryWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
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
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(event))
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
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpSearchQueryWidget_DragEnterEvent
func callbackQHelpSearchQueryWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
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
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(event))
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
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(event))
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
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Event
func callbackQHelpSearchQueryWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_FocusInEvent
func callbackQHelpSearchQueryWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusNextPrevChild
func callbackQHelpSearchQueryWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_FocusOutEvent
func callbackQHelpSearchQueryWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_HasHeightForWidth
func callbackQHelpSearchQueryWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
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
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchQueryWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_Hide
func callbackQHelpSearchQueryWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_InitPainter
func callbackQHelpSearchQueryWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchQueryWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodEvent
func callbackQHelpSearchQueryWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodQuery
func callbackQHelpSearchQueryWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchQueryWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_KeyPressEvent
func callbackQHelpSearchQueryWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Metric
func callbackQHelpSearchQueryWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpSearchQueryWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_MinimumSizeHint
func callbackQHelpSearchQueryWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_MouseDoubleClickEvent
func callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_NativeEvent
func callbackQHelpSearchQueryWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QHelpSearchQueryWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QHelpSearchQueryWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_PaintEngine
func callbackQHelpSearchQueryWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpSearchQueryWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpSearchQueryWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_PaintEvent
func callbackQHelpSearchQueryWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
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
		(*(*func(bool))(signal))(int8(disable) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(hidden) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
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
		(*(*func(bool))(signal))(int8(visible) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(vqs))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_SizeHint
func callbackQHelpSearchQueryWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_TabletEvent
func callbackQHelpSearchQueryWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(event))
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
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpSearchQueryWidget_WindowTitleChanged
func callbackQHelpSearchQueryWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQHelpSearchQueryWidget_ChildEvent
func callbackQHelpSearchQueryWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchQueryWidget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Destroyed
func callbackQHelpSearchQueryWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchQueryWidget_DisconnectNotify
func callbackQHelpSearchQueryWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_EventFilter
func callbackQHelpSearchQueryWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchQueryWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_MetaObject
func callbackQHelpSearchQueryWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchQueryWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchQueryWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_ObjectNameChanged
func callbackQHelpSearchQueryWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchQueryWidget_TimerEvent
func callbackQHelpSearchQueryWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
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
	qt.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
	return tmpValue
}

func NewQHelpSearchResult2(other QHelpSearchResult_ITF) *QHelpSearchResult {
	tmpValue := NewQHelpSearchResultFromPointer(C.QHelpSearchResult_NewQHelpSearchResult2(PointerFromQHelpSearchResult(other)))
	qt.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
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
	qt.SetFinalizer(tmpValue, (*QHelpSearchResult).DestroyQHelpSearchResult)
	return tmpValue
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResult) DestroyQHelpSearchResult() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchResult_DestroyQHelpSearchResult(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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
func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_RequestShowLink
func callbackQHelpSearchResultWidget_RequestShowLink(ptr unsafe.Pointer, link unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestShowLink"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestShowLink") {
			C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "requestShowLink")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestShowLink"); signal != nil {
			f := func(link *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(link)
				f(link)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestShowLink", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestShowLink", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DestroyQHelpSearchResultWidgetDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDestroyQHelpSearchResultWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHelpSearchResultWidget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchResultWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHelpSearchResultWidget", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidgetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func (ptr *QHelpSearchResultWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHelpSearchResultWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHelpSearchResultWidget_ActionEvent
func callbackQHelpSearchResultWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Close
func callbackQHelpSearchResultWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchResultWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_CloseEvent
func callbackQHelpSearchResultWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
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
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(event))
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
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQHelpSearchResultWidget_DragEnterEvent
func callbackQHelpSearchResultWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
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
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(event))
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
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(event))
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
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Event
func callbackQHelpSearchResultWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_FocusInEvent
func callbackQHelpSearchResultWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusNextPrevChild
func callbackQHelpSearchResultWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_FocusOutEvent
func callbackQHelpSearchResultWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_HasHeightForWidth
func callbackQHelpSearchResultWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
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
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchResultWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_Hide
func callbackQHelpSearchResultWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
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
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_InitPainter
func callbackQHelpSearchResultWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQHelpSearchResultWidget_InputMethodEvent
func callbackQHelpSearchResultWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_InputMethodQuery
func callbackQHelpSearchResultWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchResultWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_KeyPressEvent
func callbackQHelpSearchResultWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchResultWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_Metric
func callbackQHelpSearchResultWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QHelpSearchResultWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_MinimumSizeHint
func callbackQHelpSearchResultWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_MouseDoubleClickEvent
func callbackQHelpSearchResultWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
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
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_NativeEvent
func callbackQHelpSearchResultWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QHelpSearchResultWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QHelpSearchResultWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_PaintEngine
func callbackQHelpSearchResultWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQHelpSearchResultWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QHelpSearchResultWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QHelpSearchResultWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchResultWidget_PaintEvent
func callbackQHelpSearchResultWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
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
		(*(*func(bool))(signal))(int8(disable) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func())(signal))()
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
		(*(*func(bool))(signal))(int8(hidden) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
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
		(*(*func(bool))(signal))(int8(visible) != 0)
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
		(*(*func(bool))(signal))(int8(vbo) != 0)
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
		(*(*func(string))(signal))(cGoUnpackString(vqs))
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_SizeHint
func callbackQHelpSearchResultWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_TabletEvent
func callbackQHelpSearchResultWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
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
		(*(*func())(signal))()
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
		(*(*func())(signal))()
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
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(event))
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
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQHelpSearchResultWidget_WindowTitleChanged
func callbackQHelpSearchResultWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQHelpSearchResultWidget_ChildEvent
func callbackQHelpSearchResultWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHelpSearchResultWidget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_Destroyed
func callbackQHelpSearchResultWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHelpSearchResultWidget_DisconnectNotify
func callbackQHelpSearchResultWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_EventFilter
func callbackQHelpSearchResultWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHelpSearchResultWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_MetaObject
func callbackQHelpSearchResultWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchResultWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchResultWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchResultWidget_ObjectNameChanged
func callbackQHelpSearchResultWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHelpSearchResultWidget_TimerEvent
func callbackQHelpSearchResultWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["help.QCompressedHelpInfo_ITF"] = QCompressedHelpInfo{}
	qt.FuncMap["help.NewQCompressedHelpInfo"] = NewQCompressedHelpInfo
	qt.FuncMap["help.NewQCompressedHelpInfo2"] = NewQCompressedHelpInfo2
	qt.FuncMap["help.NewQCompressedHelpInfo3"] = NewQCompressedHelpInfo3
	qt.FuncMap["help.QCompressedHelpInfo_FromCompressedHelpFile"] = QCompressedHelpInfo_FromCompressedHelpFile
	qt.ItfMap["help.QHelpContentItem_ITF"] = QHelpContentItem{}
	qt.ItfMap["help.QHelpContentModel_ITF"] = QHelpContentModel{}
	qt.ItfMap["help.QHelpContentWidget_ITF"] = QHelpContentWidget{}
	qt.ItfMap["help.QHelpEngine_ITF"] = QHelpEngine{}
	qt.FuncMap["help.NewQHelpEngine"] = NewQHelpEngine
	qt.ItfMap["help.QHelpEngineCore_ITF"] = QHelpEngineCore{}
	qt.FuncMap["help.NewQHelpEngineCore"] = NewQHelpEngineCore
	qt.FuncMap["help.QHelpEngineCore_MetaData"] = QHelpEngineCore_MetaData
	qt.FuncMap["help.QHelpEngineCore_NamespaceName"] = QHelpEngineCore_NamespaceName
	qt.ItfMap["help.QHelpFilterData_ITF"] = QHelpFilterData{}
	qt.FuncMap["help.NewQHelpFilterData"] = NewQHelpFilterData
	qt.FuncMap["help.NewQHelpFilterData2"] = NewQHelpFilterData2
	qt.FuncMap["help.NewQHelpFilterData3"] = NewQHelpFilterData3
	qt.ItfMap["help.QHelpFilterEngine_ITF"] = QHelpFilterEngine{}
	qt.ItfMap["help.QHelpGlobal_ITF"] = QHelpGlobal{}
	qt.ItfMap["help.QHelpIndexModel_ITF"] = QHelpIndexModel{}
	qt.ItfMap["help.QHelpIndexWidget_ITF"] = QHelpIndexWidget{}
	qt.ItfMap["help.QHelpSearchEngine_ITF"] = QHelpSearchEngine{}
	qt.FuncMap["help.NewQHelpSearchEngine"] = NewQHelpSearchEngine
	qt.ItfMap["help.QHelpSearchQuery_ITF"] = QHelpSearchQuery{}
	qt.FuncMap["help.NewQHelpSearchQuery"] = NewQHelpSearchQuery
	qt.ItfMap["help.QHelpSearchQueryWidget_ITF"] = QHelpSearchQueryWidget{}
	qt.FuncMap["help.NewQHelpSearchQueryWidget"] = NewQHelpSearchQueryWidget
	qt.ItfMap["help.QHelpSearchResult_ITF"] = QHelpSearchResult{}
	qt.FuncMap["help.NewQHelpSearchResult"] = NewQHelpSearchResult
	qt.FuncMap["help.NewQHelpSearchResult2"] = NewQHelpSearchResult2
	qt.FuncMap["help.NewQHelpSearchResult3"] = NewQHelpSearchResult3
	qt.ItfMap["help.QHelpSearchResultWidget_ITF"] = QHelpSearchResultWidget{}
}
