// +build !minimal

package help

//#include <stdint.h>
//#include <stdlib.h>
//#include "help.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (p *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return p
}

func (p *QHelpContentItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHelpContentItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	defer qt.Recovering("QHelpContentItem::child")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(int32(row))))
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	defer qt.Recovering("QHelpContentItem::childCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {
	defer qt.Recovering("QHelpContentItem::childPosition")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildPosition(ptr.Pointer(), PointerFromQHelpContentItem(child))))
	}
	return 0
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	defer qt.Recovering("QHelpContentItem::parent")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	defer qt.Recovering("QHelpContentItem::row")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_Row(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) Title() string {
	defer qt.Recovering("QHelpContentItem::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpContentItem) Url() *core.QUrl {
	defer qt.Recovering("QHelpContentItem::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpContentItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	defer qt.Recovering("QHelpContentItem::~QHelpContentItem")

	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
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

func (p *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return p
}

func (p *QHelpContentModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (p *QHelpContentModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractItemModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpContentModel(ptr QHelpContentModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = new(QHelpContentModel)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::columnCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	defer qt.Recovering("QHelpContentModel::contentItemAt")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

//export callbackQHelpContentModel_ContentsCreated
func callbackQHelpContentModel_ContentsCreated(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreated")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "contentsCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "contentsCreated")
	}
}

func (ptr *QHelpContentModel) ContentsCreated() {
	defer qt.Recovering("QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreated(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_ContentsCreationStarted
func callbackQHelpContentModel_ContentsCreationStarted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreationStarted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "contentsCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "contentsCreationStarted")
	}
}

func (ptr *QHelpContentModel) ContentsCreationStarted() {
	defer qt.Recovering("QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	defer qt.Recovering("QHelpContentModel::createContents")

	if ptr.Pointer() != nil {
		var customFilterNameC = C.CString(customFilterName)
		defer C.free(unsafe.Pointer(customFilterNameC))
		C.QHelpContentModel_CreateContents(ptr.Pointer(), customFilterNameC)
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QHelpContentModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	defer qt.Recovering("QHelpContentModel::isCreatingContents")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::rowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	defer qt.Recovering("QHelpContentModel::~QHelpContentModel")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()))
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentModel_Sibling
func callbackQHelpContentModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentModel::sibling")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectSibling(f func(row int, column int, index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpContentModel::sibling")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "sibling", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSibling() {
	defer qt.Recovering("disconnect QHelpContentModel::sibling")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "sibling")
	}
}

func (ptr *QHelpContentModel) Sibling(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Buddy
func callbackQHelpContentModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentModel::buddy")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpContentModel::buddy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "buddy", f)
	}
}

func (ptr *QHelpContentModel) DisconnectBuddy() {
	defer qt.Recovering("disconnect QHelpContentModel::buddy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "buddy")
	}
}

func (ptr *QHelpContentModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_CanDropMimeData
func callbackQHelpContentModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::canDropMimeData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "canDropMimeData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCanDropMimeData() {
	defer qt.Recovering("disconnect QHelpContentModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "canDropMimeData")
	}
}

func (ptr *QHelpContentModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_CanFetchMore
func callbackQHelpContentModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::canFetchMore")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectCanFetchMore(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::canFetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "canFetchMore", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCanFetchMore() {
	defer qt.Recovering("disconnect QHelpContentModel::canFetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "canFetchMore")
	}
}

func (ptr *QHelpContentModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_DropMimeData
func callbackQHelpContentModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::dropMimeData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "dropMimeData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDropMimeData() {
	defer qt.Recovering("disconnect QHelpContentModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "dropMimeData")
	}
}

func (ptr *QHelpContentModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_FetchMore
func callbackQHelpContentModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::fetchMore")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpContentModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpContentModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "fetchMore", f)
	}
}

func (ptr *QHelpContentModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "fetchMore")
	}
}

func (ptr *QHelpContentModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpContentModel_Flags
func callbackQHelpContentModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpContentModel::flags")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectFlags(f func(index *core.QModelIndex) core.Qt__ItemFlag) {
	defer qt.Recovering("connect QHelpContentModel::flags")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "flags", f)
	}
}

func (ptr *QHelpContentModel) DisconnectFlags() {
	defer qt.Recovering("disconnect QHelpContentModel::flags")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "flags")
	}
}

func (ptr *QHelpContentModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QHelpContentModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QHelpContentModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QHelpContentModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpContentModel_HasChildren
func callbackQHelpContentModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::hasChildren")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "hasChildren", f)
	}
}

func (ptr *QHelpContentModel) DisconnectHasChildren() {
	defer qt.Recovering("disconnect QHelpContentModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "hasChildren")
	}
}

func (ptr *QHelpContentModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_HeaderData
func callbackQHelpContentModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentModel::headerData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpContentModel) ConnectHeaderData(f func(section int, orientation core.Qt__Orientation, role int) *core.QVariant) {
	defer qt.Recovering("connect QHelpContentModel::headerData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "headerData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectHeaderData() {
	defer qt.Recovering("disconnect QHelpContentModel::headerData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "headerData")
	}
}

func (ptr *QHelpContentModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QHelpContentModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QHelpContentModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_InsertColumns
func callbackQHelpContentModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::insertColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectInsertColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::insertColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "insertColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectInsertColumns() {
	defer qt.Recovering("disconnect QHelpContentModel::insertColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "insertColumns")
	}
}

func (ptr *QHelpContentModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_InsertRows
func callbackQHelpContentModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::insertRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectInsertRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::insertRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "insertRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectInsertRows() {
	defer qt.Recovering("disconnect QHelpContentModel::insertRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "insertRows")
	}
}

func (ptr *QHelpContentModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_MimeTypes
func callbackQHelpContentModel_MimeTypes(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QHelpContentModel::mimeTypes")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQHelpContentModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QHelpContentModel) ConnectMimeTypes(f func() []string) {
	defer qt.Recovering("connect QHelpContentModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "mimeTypes", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMimeTypes() {
	defer qt.Recovering("disconnect QHelpContentModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "mimeTypes")
	}
}

func (ptr *QHelpContentModel) MimeTypes() []string {
	defer qt.Recovering("QHelpContentModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpContentModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpContentModel) MimeTypesDefault() []string {
	defer qt.Recovering("QHelpContentModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpContentModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpContentModel_MoveColumns
func callbackQHelpContentModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QHelpContentModel::moveColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QHelpContentModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "moveColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMoveColumns() {
	defer qt.Recovering("disconnect QHelpContentModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "moveColumns")
	}
}

func (ptr *QHelpContentModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpContentModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpContentModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpContentModel_MoveRows
func callbackQHelpContentModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QHelpContentModel::moveRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QHelpContentModel::moveRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "moveRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMoveRows() {
	defer qt.Recovering("disconnect QHelpContentModel::moveRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "moveRows")
	}
}

func (ptr *QHelpContentModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpContentModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpContentModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveColumns
func callbackQHelpContentModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::removeColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectRemoveColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::removeColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "removeColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRemoveColumns() {
	defer qt.Recovering("disconnect QHelpContentModel::removeColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "removeColumns")
	}
}

func (ptr *QHelpContentModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveRows
func callbackQHelpContentModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::removeRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectRemoveRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentModel::removeRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "removeRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRemoveRows() {
	defer qt.Recovering("disconnect QHelpContentModel::removeRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "removeRows")
	}
}

func (ptr *QHelpContentModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_ResetInternalData
func callbackQHelpContentModel_ResetInternalData(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::resetInternalData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpContentModel) ConnectResetInternalData(f func()) {
	defer qt.Recovering("connect QHelpContentModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "resetInternalData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectResetInternalData() {
	defer qt.Recovering("disconnect QHelpContentModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "resetInternalData")
	}
}

func (ptr *QHelpContentModel) ResetInternalData() {
	defer qt.Recovering("QHelpContentModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ResetInternalDataDefault() {
	defer qt.Recovering("QHelpContentModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_Revert
func callbackQHelpContentModel_Revert(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::revert")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpContentModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "revert", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "revert")
	}
}

func (ptr *QHelpContentModel) Revert() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_SetData
func callbackQHelpContentModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	defer qt.Recovering("callback QHelpContentModel::setData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) ConnectSetData(f func(index *core.QModelIndex, value *core.QVariant, role int) bool) {
	defer qt.Recovering("connect QHelpContentModel::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "setData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSetData() {
	defer qt.Recovering("disconnect QHelpContentModel::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "setData")
	}
}

func (ptr *QHelpContentModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpContentModel::setData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpContentModel::setData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpContentModel_SetHeaderData
func callbackQHelpContentModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	defer qt.Recovering("callback QHelpContentModel::setHeaderData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) ConnectSetHeaderData(f func(section int, orientation core.Qt__Orientation, value *core.QVariant, role int) bool) {
	defer qt.Recovering("connect QHelpContentModel::setHeaderData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "setHeaderData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSetHeaderData() {
	defer qt.Recovering("disconnect QHelpContentModel::setHeaderData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "setHeaderData")
	}
}

func (ptr *QHelpContentModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpContentModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetHeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpContentModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpContentModel_Sort
func callbackQHelpContentModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	defer qt.Recovering("callback QHelpContentModel::sort")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "sort", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "sort")
	}
}

func (ptr *QHelpContentModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpContentModel_Span
func callbackQHelpContentModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentModel::span")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpContentModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QHelpContentModel::span")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "span", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSpan() {
	defer qt.Recovering("disconnect QHelpContentModel::span")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "span")
	}
}

func (ptr *QHelpContentModel) Span(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QHelpContentModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QHelpContentModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Submit
func callbackQHelpContentModel_Submit(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::submit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpContentModel) ConnectSubmit(f func() bool) {
	defer qt.Recovering("connect QHelpContentModel::submit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "submit", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSubmit() {
	defer qt.Recovering("disconnect QHelpContentModel::submit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "submit")
	}
}

func (ptr *QHelpContentModel) Submit() bool {
	defer qt.Recovering("QHelpContentModel::submit")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SubmitDefault() bool {
	defer qt.Recovering("QHelpContentModel::submit")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentModel_SupportedDragActions
func callbackQHelpContentModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpContentModel::supportedDragActions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpContentModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QHelpContentModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "supportedDragActions", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSupportedDragActions() {
	defer qt.Recovering("disconnect QHelpContentModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "supportedDragActions")
	}
}

func (ptr *QHelpContentModel) SupportedDragActions() core.Qt__DropAction {
	defer qt.Recovering("QHelpContentModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentModel) SupportedDragActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QHelpContentModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_SupportedDropActions
func callbackQHelpContentModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpContentModel::supportedDropActions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpContentModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QHelpContentModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "supportedDropActions", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSupportedDropActions() {
	defer qt.Recovering("disconnect QHelpContentModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "supportedDropActions")
	}
}

func (ptr *QHelpContentModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QHelpContentModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentModel) SupportedDropActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QHelpContentModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_TimerEvent
func callbackQHelpContentModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpContentModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpContentModel_ChildEvent
func callbackQHelpContentModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpContentModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentModel_ConnectNotify
func callbackQHelpContentModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpContentModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpContentModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpContentModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpContentModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_CustomEvent
func callbackQHelpContentModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpContentModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentModel_DeleteLater
func callbackQHelpContentModel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpContentModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpContentModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpContentModel) DeleteLater() {
	defer qt.Recovering("QHelpContentModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()))
		C.QHelpContentModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) DeleteLaterDefault() {
	defer qt.Recovering("QHelpContentModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()))
		C.QHelpContentModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentModel_DisconnectNotify
func callbackQHelpContentModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpContentModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpContentModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpContentModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_Event
func callbackQHelpContentModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::event")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpContentModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpContentModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QHelpContentModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QHelpContentModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentModel::event")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentModel::event")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpContentModel_EventFilter
func callbackQHelpContentModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentModel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpContentModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpContentModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpContentModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpContentModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentModel_MetaObject
func callbackQHelpContentModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentModel::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpContentModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpContentModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentModel(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpContentModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpContentModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpContentModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func (p *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return p
}

func (p *QHelpContentWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QTreeView_PTR().Pointer()
	}
	return nil
}

func (p *QHelpContentWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QTreeView_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpContentWidget) DestroyQHelpContentWidget() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::indexOf")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexOf(ptr.Pointer(), core.PointerFromQUrl(link)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_LinkActivated
func callbackQHelpContentWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::linkActivated")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "linkActivated", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "linkActivated")
	}
}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

//export callbackQHelpContentWidget_Collapse
func callbackQHelpContentWidget_Collapse(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::collapse")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "collapse"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectCollapse(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::collapse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "collapse", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCollapse() {
	defer qt.Recovering("disconnect QHelpContentWidget::collapse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "collapse")
	}
}

func (ptr *QHelpContentWidget) Collapse(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::collapse")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) CollapseDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::collapse")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Expand
func callbackQHelpContentWidget_Expand(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::expand")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "expand"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectExpand(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::expand")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expand", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpand() {
	defer qt.Recovering("disconnect QHelpContentWidget::expand")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expand")
	}
}

func (ptr *QHelpContentWidget) Expand(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::expand")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) ExpandDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::expand")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_CollapseAll
func callbackQHelpContentWidget_CollapseAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::collapseAll")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "collapseAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectCollapseAll(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::collapseAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "collapseAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCollapseAll() {
	defer qt.Recovering("disconnect QHelpContentWidget::collapseAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "collapseAll")
	}
}

func (ptr *QHelpContentWidget) CollapseAll() {
	defer qt.Recovering("QHelpContentWidget::collapseAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) CollapseAllDefault() {
	defer qt.Recovering("QHelpContentWidget::collapseAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ColumnCountChanged
func callbackQHelpContentWidget_ColumnCountChanged(ptr unsafe.Pointer, oldCount C.int, newCount C.int) {
	defer qt.Recovering("callback QHelpContentWidget::columnCountChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "columnCountChanged"); signal != nil {
		signal.(func(int, int))(int(int32(oldCount)), int(int32(newCount)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnCountChangedDefault(int(int32(oldCount)), int(int32(newCount)))
	}
}

func (ptr *QHelpContentWidget) ConnectColumnCountChanged(f func(oldCount int, newCount int)) {
	defer qt.Recovering("connect QHelpContentWidget::columnCountChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnCountChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnCountChanged() {
	defer qt.Recovering("disconnect QHelpContentWidget::columnCountChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnCountChanged")
	}
}

func (ptr *QHelpContentWidget) ColumnCountChanged(oldCount int, newCount int) {
	defer qt.Recovering("QHelpContentWidget::columnCountChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnCountChanged(ptr.Pointer(), C.int(int32(oldCount)), C.int(int32(newCount)))
	}
}

func (ptr *QHelpContentWidget) ColumnCountChangedDefault(oldCount int, newCount int) {
	defer qt.Recovering("QHelpContentWidget::columnCountChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnCountChangedDefault(ptr.Pointer(), C.int(int32(oldCount)), C.int(int32(newCount)))
	}
}

//export callbackQHelpContentWidget_ColumnMoved
func callbackQHelpContentWidget_ColumnMoved(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::columnMoved")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "columnMoved"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnMovedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectColumnMoved(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::columnMoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnMoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnMoved() {
	defer qt.Recovering("disconnect QHelpContentWidget::columnMoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnMoved")
	}
}

func (ptr *QHelpContentWidget) ColumnMoved() {
	defer qt.Recovering("QHelpContentWidget::columnMoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnMoved(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ColumnMovedDefault() {
	defer qt.Recovering("QHelpContentWidget::columnMoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnMovedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ColumnResized
func callbackQHelpContentWidget_ColumnResized(ptr unsafe.Pointer, column C.int, oldSize C.int, newSize C.int) {
	defer qt.Recovering("callback QHelpContentWidget::columnResized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "columnResized"); signal != nil {
		signal.(func(int, int, int))(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnResizedDefault(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	}
}

func (ptr *QHelpContentWidget) ConnectColumnResized(f func(column int, oldSize int, newSize int)) {
	defer qt.Recovering("connect QHelpContentWidget::columnResized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnResized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnResized() {
	defer qt.Recovering("disconnect QHelpContentWidget::columnResized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "columnResized")
	}
}

func (ptr *QHelpContentWidget) ColumnResized(column int, oldSize int, newSize int) {
	defer qt.Recovering("QHelpContentWidget::columnResized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnResized(ptr.Pointer(), C.int(int32(column)), C.int(int32(oldSize)), C.int(int32(newSize)))
	}
}

func (ptr *QHelpContentWidget) ColumnResizedDefault(column int, oldSize int, newSize int) {
	defer qt.Recovering("QHelpContentWidget::columnResized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnResizedDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(oldSize)), C.int(int32(newSize)))
	}
}

//export callbackQHelpContentWidget_CurrentChanged
func callbackQHelpContentWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::currentChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpContentWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "currentChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "currentChanged")
	}
}

func (ptr *QHelpContentWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpContentWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpContentWidget_DragMoveEvent
func callbackQHelpContentWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragMoveEvent")
	}
}

func (ptr *QHelpContentWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_DrawBranches
func callbackQHelpContentWidget_DrawBranches(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::drawBranches")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawBranchesDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectDrawBranches(f func(painter *gui.QPainter, rect *core.QRect, index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "drawBranches", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDrawBranches() {
	defer qt.Recovering("disconnect QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "drawBranches")
	}
}

func (ptr *QHelpContentWidget) DrawBranches(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranches(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranchesDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_DrawRow
func callbackQHelpContentWidget_DrawRow(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::drawRow")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "drawRow"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawRowDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectDrawRow(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::drawRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "drawRow", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDrawRow() {
	defer qt.Recovering("disconnect QHelpContentWidget::drawRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "drawRow")
	}
}

func (ptr *QHelpContentWidget) DrawRow(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawRow")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRow(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) DrawRowDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawRow")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRowDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ExpandAll
func callbackQHelpContentWidget_ExpandAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::expandAll")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "expandAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectExpandAll(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::expandAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expandAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpandAll() {
	defer qt.Recovering("disconnect QHelpContentWidget::expandAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expandAll")
	}
}

func (ptr *QHelpContentWidget) ExpandAll() {
	defer qt.Recovering("QHelpContentWidget::expandAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ExpandAllDefault() {
	defer qt.Recovering("QHelpContentWidget::expandAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ExpandToDepth
func callbackQHelpContentWidget_ExpandToDepth(ptr unsafe.Pointer, depth C.int) {
	defer qt.Recovering("callback QHelpContentWidget::expandToDepth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "expandToDepth"); signal != nil {
		signal.(func(int))(int(int32(depth)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandToDepthDefault(int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ConnectExpandToDepth(f func(depth int)) {
	defer qt.Recovering("connect QHelpContentWidget::expandToDepth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expandToDepth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpandToDepth() {
	defer qt.Recovering("disconnect QHelpContentWidget::expandToDepth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "expandToDepth")
	}
}

func (ptr *QHelpContentWidget) ExpandToDepth(depth int) {
	defer qt.Recovering("QHelpContentWidget::expandToDepth")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandToDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ExpandToDepthDefault(depth int) {
	defer qt.Recovering("QHelpContentWidget::expandToDepth")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandToDepthDefault(ptr.Pointer(), C.int(int32(depth)))
	}
}

//export callbackQHelpContentWidget_HideColumn
func callbackQHelpContentWidget_HideColumn(ptr unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QHelpContentWidget::hideColumn")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "hideColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectHideColumn(f func(column int)) {
	defer qt.Recovering("connect QHelpContentWidget::hideColumn")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hideColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideColumn() {
	defer qt.Recovering("disconnect QHelpContentWidget::hideColumn")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hideColumn")
	}
}

func (ptr *QHelpContentWidget) HideColumn(column int) {
	defer qt.Recovering("QHelpContentWidget::hideColumn")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) HideColumnDefault(column int) {
	defer qt.Recovering("QHelpContentWidget::hideColumn")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_HorizontalOffset
func callbackQHelpContentWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpContentWidget::horizontalOffset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpContentWidget) ConnectHorizontalOffset(f func() int) {
	defer qt.Recovering("connect QHelpContentWidget::horizontalOffset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "horizontalOffset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHorizontalOffset() {
	defer qt.Recovering("disconnect QHelpContentWidget::horizontalOffset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "horizontalOffset")
	}
}

func (ptr *QHelpContentWidget) HorizontalOffset() int {
	defer qt.Recovering("QHelpContentWidget::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentWidget) HorizontalOffsetDefault() int {
	defer qt.Recovering("QHelpContentWidget::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_IndexAt
func callbackQHelpContentWidget_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::indexAt")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QHelpContentWidget) ConnectIndexAt(f func(point *core.QPoint) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpContentWidget::indexAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "indexAt", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectIndexAt() {
	defer qt.Recovering("disconnect QHelpContentWidget::indexAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "indexAt")
	}
}

func (ptr *QHelpContentWidget) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::indexAt")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::indexAt")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_IsIndexHidden
func callbackQHelpContentWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::isIndexHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpContentWidget) ConnectIsIndexHidden(f func(index *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpContentWidget::isIndexHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "isIndexHidden", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectIsIndexHidden() {
	defer qt.Recovering("disconnect QHelpContentWidget::isIndexHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "isIndexHidden")
	}
}

func (ptr *QHelpContentWidget) IsIndexHidden(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_KeyPressEvent
func callbackQHelpContentWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyPressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyPressEvent")
	}
}

func (ptr *QHelpContentWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_KeyboardSearch
func callbackQHelpContentWidget_KeyboardSearch(ptr unsafe.Pointer, search *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::keyboardSearch")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyboardSearch", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyboardSearch")
	}
}

func (ptr *QHelpContentWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpContentWidget_KeyboardSearch(ptr.Pointer(), searchC)
	}
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpContentWidget_KeyboardSearchDefault(ptr.Pointer(), searchC)
	}
}

//export callbackQHelpContentWidget_MouseDoubleClickEvent
func callbackQHelpContentWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent")
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseMoveEvent
func callbackQHelpContentWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *QHelpContentWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MousePressEvent
func callbackQHelpContentWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mousePressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *QHelpContentWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseReleaseEvent
func callbackQHelpContentWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MoveCursor
func callbackQHelpContentWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::moveCursor")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpContentWidget) ConnectMoveCursor(f func(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpContentWidget::moveCursor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "moveCursor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveCursor() {
	defer qt.Recovering("disconnect QHelpContentWidget::moveCursor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "moveCursor")
	}
}

func (ptr *QHelpContentWidget) MoveCursor(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::moveCursor")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::moveCursor")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_PaintEvent
func callbackQHelpContentWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::paintEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "paintEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "paintEvent")
	}
}

func (ptr *QHelpContentWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpContentWidget_Reset
func callbackQHelpContentWidget_Reset(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::reset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "reset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "reset")
	}
}

func (ptr *QHelpContentWidget) Reset() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ResetDefault() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ResizeColumnToContents
func callbackQHelpContentWidget_ResizeColumnToContents(ptr unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QHelpContentWidget::resizeColumnToContents")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "resizeColumnToContents"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeColumnToContentsDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeColumnToContents(f func(column int)) {
	defer qt.Recovering("connect QHelpContentWidget::resizeColumnToContents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "resizeColumnToContents", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeColumnToContents() {
	defer qt.Recovering("disconnect QHelpContentWidget::resizeColumnToContents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "resizeColumnToContents")
	}
}

func (ptr *QHelpContentWidget) ResizeColumnToContents(column int) {
	defer qt.Recovering("QHelpContentWidget::resizeColumnToContents")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeColumnToContents(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ResizeColumnToContentsDefault(column int) {
	defer qt.Recovering("QHelpContentWidget::resizeColumnToContents")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeColumnToContentsDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_RowsAboutToBeRemoved
func callbackQHelpContentWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpContentWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsAboutToBeRemoved")
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsInserted
func callbackQHelpContentWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpContentWidget::rowsInserted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsInserted", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsInserted")
	}
}

func (ptr *QHelpContentWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsRemoved
func callbackQHelpContentWidget_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpContentWidget::rowsRemoved")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpContentWidget::rowsRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsRemoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsRemoved() {
	defer qt.Recovering("disconnect QHelpContentWidget::rowsRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "rowsRemoved")
	}
}

func (ptr *QHelpContentWidget) RowsRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_ScrollContentsBy
func callbackQHelpContentWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpContentWidget::scrollContentsBy")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollContentsBy", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollContentsBy")
	}
}

func (ptr *QHelpContentWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsBy(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpContentWidget_ScrollTo
func callbackQHelpContentWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	defer qt.Recovering("callback QHelpContentWidget::scrollTo")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollTo", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollTo")
	}
}

func (ptr *QHelpContentWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QHelpContentWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_SelectAll
func callbackQHelpContentWidget_SelectAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::selectAll")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectAll")
	}
}

func (ptr *QHelpContentWidget) SelectAll() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionChanged
func callbackQHelpContentWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::selectionChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpContentWidget) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	defer qt.Recovering("connect QHelpContentWidget::selectionChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectionChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QHelpContentWidget::selectionChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectionChanged")
	}
}

func (ptr *QHelpContentWidget) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	defer qt.Recovering("QHelpContentWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QHelpContentWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	defer qt.Recovering("QHelpContentWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpContentWidget_SetModel
func callbackQHelpContentWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setModel")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpContentWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setModel")
	}
}

func (ptr *QHelpContentWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpContentWidget_SetRootIndex
func callbackQHelpContentWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setRootIndex")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setRootIndex", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setRootIndex")
	}
}

func (ptr *QHelpContentWidget) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SetSelection
func callbackQHelpContentWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	defer qt.Recovering("callback QHelpContentWidget::setSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setSelection")
	}
}

func (ptr *QHelpContentWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpContentWidget_SetSelectionModel
func callbackQHelpContentWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setSelectionModel")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setSelectionModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setSelectionModel")
	}
}

func (ptr *QHelpContentWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpContentWidget_ShowColumn
func callbackQHelpContentWidget_ShowColumn(ptr unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QHelpContentWidget::showColumn")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectShowColumn(f func(column int)) {
	defer qt.Recovering("connect QHelpContentWidget::showColumn")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowColumn() {
	defer qt.Recovering("disconnect QHelpContentWidget::showColumn")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showColumn")
	}
}

func (ptr *QHelpContentWidget) ShowColumn(column int) {
	defer qt.Recovering("QHelpContentWidget::showColumn")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ShowColumnDefault(column int) {
	defer qt.Recovering("QHelpContentWidget::showColumn")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_SizeHintForColumn
func callbackQHelpContentWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	defer qt.Recovering("callback QHelpContentWidget::sizeHintForColumn")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpContentWidget) ConnectSizeHintForColumn(f func(column int) int) {
	defer qt.Recovering("connect QHelpContentWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHintForColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHintForColumn() {
	defer qt.Recovering("disconnect QHelpContentWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHintForColumn")
	}
}

func (ptr *QHelpContentWidget) SizeHintForColumn(column int) int {
	defer qt.Recovering("QHelpContentWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumn(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) SizeHintForColumnDefault(column int) int {
	defer qt.Recovering("QHelpContentWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpContentWidget_UpdateGeometries
func callbackQHelpContentWidget_UpdateGeometries(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::updateGeometries")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "updateGeometries", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "updateGeometries")
	}
}

func (ptr *QHelpContentWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_VerticalOffset
func callbackQHelpContentWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpContentWidget::verticalOffset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpContentWidget) ConnectVerticalOffset(f func() int) {
	defer qt.Recovering("connect QHelpContentWidget::verticalOffset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "verticalOffset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVerticalOffset() {
	defer qt.Recovering("disconnect QHelpContentWidget::verticalOffset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "verticalOffset")
	}
}

func (ptr *QHelpContentWidget) VerticalOffset() int {
	defer qt.Recovering("QHelpContentWidget::verticalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentWidget) VerticalOffsetDefault() int {
	defer qt.Recovering("QHelpContentWidget::verticalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_ViewportEvent
func callbackQHelpContentWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::viewportEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectViewportEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpContentWidget::viewportEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewportEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewportEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::viewportEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewportEvent")
	}
}

func (ptr *QHelpContentWidget) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_ViewportSizeHint
func callbackQHelpContentWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::viewportSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectViewportSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpContentWidget::viewportSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewportSizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewportSizeHint() {
	defer qt.Recovering("disconnect QHelpContentWidget::viewportSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewportSizeHint")
	}
}

func (ptr *QHelpContentWidget) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::viewportSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) ViewportSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::viewportSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRect
func callbackQHelpContentWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::visualRect")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpContentWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentWidget) ConnectVisualRect(f func(index *core.QModelIndex) *core.QRect) {
	defer qt.Recovering("connect QHelpContentWidget::visualRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "visualRect", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVisualRect() {
	defer qt.Recovering("disconnect QHelpContentWidget::visualRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "visualRect")
	}
}

func (ptr *QHelpContentWidget) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QHelpContentWidget::visualRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpContentWidget_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QHelpContentWidget::visualRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpContentWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRegionForSelection
func callbackQHelpContentWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::visualRegionForSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpContentWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpContentWidget) ConnectVisualRegionForSelection(f func(selection *core.QItemSelection) *gui.QRegion) {
	defer qt.Recovering("connect QHelpContentWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "visualRegionForSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVisualRegionForSelection() {
	defer qt.Recovering("disconnect QHelpContentWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "visualRegionForSelection")
	}
}

func (ptr *QHelpContentWidget) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QHelpContentWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QHelpContentWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_DragLeaveEvent
func callbackQHelpContentWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragLeaveEvent")
	}
}

func (ptr *QHelpContentWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpContentWidget_ClearSelection
func callbackQHelpContentWidget_ClearSelection(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::clearSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectClearSelection(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::clearSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "clearSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectClearSelection() {
	defer qt.Recovering("disconnect QHelpContentWidget::clearSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "clearSelection")
	}
}

func (ptr *QHelpContentWidget) ClearSelection() {
	defer qt.Recovering("QHelpContentWidget::clearSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ClearSelectionDefault() {
	defer qt.Recovering("QHelpContentWidget::clearSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_CloseEditor
func callbackQHelpContentWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	defer qt.Recovering("callback QHelpContentWidget::closeEditor")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "closeEditor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "closeEditor")
	}
}

func (ptr *QHelpContentWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_CommitData
func callbackQHelpContentWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::commitData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "commitData", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "commitData")
	}
}

func (ptr *QHelpContentWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpContentWidget_DragEnterEvent
func callbackQHelpContentWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragEnterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dragEnterEvent")
	}
}

func (ptr *QHelpContentWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpContentWidget_DropEvent
func callbackQHelpContentWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dropEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "dropEvent")
	}
}

func (ptr *QHelpContentWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpContentWidget_Edit
func callbackQHelpContentWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::edit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectEdit(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::edit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "edit", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEdit() {
	defer qt.Recovering("disconnect QHelpContentWidget::edit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "edit")
	}
}

func (ptr *QHelpContentWidget) Edit(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::edit")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) EditDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::edit")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Edit2
func callbackQHelpContentWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::edit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectEdit2(f func(index *core.QModelIndex, trigger widgets.QAbstractItemView__EditTrigger, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpContentWidget::edit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "edit2", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEdit2() {
	defer qt.Recovering("disconnect QHelpContentWidget::edit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "edit2")
	}
}

func (ptr *QHelpContentWidget) Edit2(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::edit")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Edit2(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::edit")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_EditorDestroyed
func callbackQHelpContentWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::editorDestroyed")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "editorDestroyed", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "editorDestroyed")
	}
}

func (ptr *QHelpContentWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpContentWidget_FocusInEvent
func callbackQHelpContentWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusInEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusInEvent")
	}
}

func (ptr *QHelpContentWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_FocusNextPrevChild
func callbackQHelpContentWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QHelpContentWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpContentWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QHelpContentWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusNextPrevChild", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusNextPrevChild")
	}
}

func (ptr *QHelpContentWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QHelpContentWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QHelpContentWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_FocusOutEvent
func callbackQHelpContentWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusOutEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "focusOutEvent")
	}
}

func (ptr *QHelpContentWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_InputMethodEvent
func callbackQHelpContentWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "inputMethodEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "inputMethodEvent")
	}
}

func (ptr *QHelpContentWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpContentWidget_InputMethodQuery
func callbackQHelpContentWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpContentWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpContentWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QHelpContentWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "inputMethodQuery", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QHelpContentWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "inputMethodQuery")
	}
}

func (ptr *QHelpContentWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpContentWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpContentWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ResizeEvent
func callbackQHelpContentWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "resizeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "resizeEvent")
	}
}

func (ptr *QHelpContentWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpContentWidget_ScrollToBottom
func callbackQHelpContentWidget_ScrollToBottom(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::scrollToBottom")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectScrollToBottom(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::scrollToBottom")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollToBottom", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollToBottom() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollToBottom")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollToBottom")
	}
}

func (ptr *QHelpContentWidget) ScrollToBottom() {
	defer qt.Recovering("QHelpContentWidget::scrollToBottom")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ScrollToBottomDefault() {
	defer qt.Recovering("QHelpContentWidget::scrollToBottom")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ScrollToTop
func callbackQHelpContentWidget_ScrollToTop(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::scrollToTop")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectScrollToTop(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::scrollToTop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollToTop", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollToTop() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollToTop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "scrollToTop")
	}
}

func (ptr *QHelpContentWidget) ScrollToTop() {
	defer qt.Recovering("QHelpContentWidget::scrollToTop")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ScrollToTopDefault() {
	defer qt.Recovering("QHelpContentWidget::scrollToTop")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionCommand
func callbackQHelpContentWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpContentWidget::selectionCommand")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpContentWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpContentWidget) ConnectSelectionCommand(f func(index *core.QModelIndex, event *core.QEvent) core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("connect QHelpContentWidget::selectionCommand")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectionCommand", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectionCommand() {
	defer qt.Recovering("disconnect QHelpContentWidget::selectionCommand")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "selectionCommand")
	}
}

func (ptr *QHelpContentWidget) SelectionCommand(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	defer qt.Recovering("QHelpContentWidget::selectionCommand")

	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommand(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

func (ptr *QHelpContentWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	defer qt.Recovering("QHelpContentWidget::selectionCommand")

	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpContentWidget_SetCurrentIndex
func callbackQHelpContentWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setCurrentIndex")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectSetCurrentIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setCurrentIndex", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetCurrentIndex() {
	defer qt.Recovering("disconnect QHelpContentWidget::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setCurrentIndex")
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SizeHintForRow
func callbackQHelpContentWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	defer qt.Recovering("callback QHelpContentWidget::sizeHintForRow")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpContentWidget) ConnectSizeHintForRow(f func(row int) int) {
	defer qt.Recovering("connect QHelpContentWidget::sizeHintForRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHintForRow", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHintForRow() {
	defer qt.Recovering("disconnect QHelpContentWidget::sizeHintForRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHintForRow")
	}
}

func (ptr *QHelpContentWidget) SizeHintForRow(row int) int {
	defer qt.Recovering("QHelpContentWidget::sizeHintForRow")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRow(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) SizeHintForRowDefault(row int) int {
	defer qt.Recovering("QHelpContentWidget::sizeHintForRow")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpContentWidget_StartDrag
func callbackQHelpContentWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	defer qt.Recovering("callback QHelpContentWidget::startDrag")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpContentWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "startDrag", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "startDrag")
	}
}

func (ptr *QHelpContentWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDrag(ptr.Pointer(), C.longlong(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpContentWidget_Update
func callbackQHelpContentWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::update")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectUpdate(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "update", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QHelpContentWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "update")
	}
}

func (ptr *QHelpContentWidget) Update(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) UpdateDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ViewOptions
func callbackQHelpContentWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::viewOptions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpContentWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpContentWidget) ConnectViewOptions(f func() *widgets.QStyleOptionViewItem) {
	defer qt.Recovering("connect QHelpContentWidget::viewOptions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewOptions", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewOptions() {
	defer qt.Recovering("disconnect QHelpContentWidget::viewOptions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "viewOptions")
	}
}

func (ptr *QHelpContentWidget) ViewOptions() *widgets.QStyleOptionViewItem {
	defer qt.Recovering("QHelpContentWidget::viewOptions")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptions(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	defer qt.Recovering("QHelpContentWidget::viewOptions")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ContextMenuEvent
func callbackQHelpContentWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "contextMenuEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "contextMenuEvent")
	}
}

func (ptr *QHelpContentWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpContentWidget_MinimumSizeHint
func callbackQHelpContentWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpContentWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "minimumSizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QHelpContentWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "minimumSizeHint")
	}
}

func (ptr *QHelpContentWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_SetupViewport
func callbackQHelpContentWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setupViewport")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setupViewport", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setupViewport")
	}
}

func (ptr *QHelpContentWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpContentWidget_SizeHint
func callbackQHelpContentWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::sizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpContentWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QHelpContentWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "sizeHint")
	}
}

func (ptr *QHelpContentWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpContentWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_WheelEvent
func callbackQHelpContentWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "wheelEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *QHelpContentWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpContentWidget_ChangeEvent
func callbackQHelpContentWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "changeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "changeEvent")
	}
}

func (ptr *QHelpContentWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpContentWidget_ActionEvent
func callbackQHelpContentWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "actionEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "actionEvent")
	}
}

func (ptr *QHelpContentWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpContentWidget_EnterEvent
func callbackQHelpContentWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "enterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "enterEvent")
	}
}

func (ptr *QHelpContentWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_HideEvent
func callbackQHelpContentWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hideEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hideEvent")
	}
}

func (ptr *QHelpContentWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpContentWidget_LeaveEvent
func callbackQHelpContentWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "leaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "leaveEvent")
	}
}

func (ptr *QHelpContentWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_MoveEvent
func callbackQHelpContentWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "moveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "moveEvent")
	}
}

func (ptr *QHelpContentWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_SetEnabled
func callbackQHelpContentWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setEnabled", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QHelpContentWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setEnabled")
	}
}

func (ptr *QHelpContentWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QHelpContentWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpContentWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QHelpContentWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetStyleSheet
func callbackQHelpContentWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QHelpContentWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QHelpContentWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setStyleSheet", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QHelpContentWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setStyleSheet")
	}
}

func (ptr *QHelpContentWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QHelpContentWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpContentWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpContentWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QHelpContentWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpContentWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpContentWidget_SetVisible
func callbackQHelpContentWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setVisible")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setVisible", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setVisible")
	}
}

func (ptr *QHelpContentWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpContentWidget_SetWindowModified
func callbackQHelpContentWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setWindowModified", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QHelpContentWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setWindowModified")
	}
}

func (ptr *QHelpContentWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QHelpContentWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpContentWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QHelpContentWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetWindowTitle
func callbackQHelpContentWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QHelpContentWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QHelpContentWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setWindowTitle", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QHelpContentWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setWindowTitle")
	}
}

func (ptr *QHelpContentWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QHelpContentWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpContentWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpContentWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QHelpContentWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpContentWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpContentWidget_ShowEvent
func callbackQHelpContentWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showEvent")
	}
}

func (ptr *QHelpContentWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpContentWidget_Close
func callbackQHelpContentWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::close")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpContentWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QHelpContentWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "close", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QHelpContentWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "close")
	}
}

func (ptr *QHelpContentWidget) Close() bool {
	defer qt.Recovering("QHelpContentWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) CloseDefault() bool {
	defer qt.Recovering("QHelpContentWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentWidget_CloseEvent
func callbackQHelpContentWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "closeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "closeEvent")
	}
}

func (ptr *QHelpContentWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpContentWidget_HasHeightForWidth
func callbackQHelpContentWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpContentWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QHelpContentWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hasHeightForWidth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QHelpContentWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hasHeightForWidth")
	}
}

func (ptr *QHelpContentWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QHelpContentWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QHelpContentWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentWidget_HeightForWidth
func callbackQHelpContentWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QHelpContentWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpContentWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QHelpContentWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "heightForWidth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QHelpContentWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "heightForWidth")
	}
}

func (ptr *QHelpContentWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QHelpContentWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QHelpContentWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpContentWidget_Hide
func callbackQHelpContentWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::hide")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hide", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QHelpContentWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "hide")
	}
}

func (ptr *QHelpContentWidget) Hide() {
	defer qt.Recovering("QHelpContentWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) HideDefault() {
	defer qt.Recovering("QHelpContentWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_KeyReleaseEvent
func callbackQHelpContentWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "keyReleaseEvent")
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_Lower
func callbackQHelpContentWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::lower")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "lower", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QHelpContentWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "lower")
	}
}

func (ptr *QHelpContentWidget) Lower() {
	defer qt.Recovering("QHelpContentWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) LowerDefault() {
	defer qt.Recovering("QHelpContentWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_NativeEvent
func callbackQHelpContentWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QHelpContentWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QHelpContentWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QHelpContentWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "nativeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "nativeEvent")
	}
}

func (ptr *QHelpContentWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpContentWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpContentWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpContentWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpContentWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_Raise
func callbackQHelpContentWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::raise")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "raise", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QHelpContentWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "raise")
	}
}

func (ptr *QHelpContentWidget) Raise() {
	defer qt.Recovering("QHelpContentWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) RaiseDefault() {
	defer qt.Recovering("QHelpContentWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Repaint
func callbackQHelpContentWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "repaint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QHelpContentWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "repaint")
	}
}

func (ptr *QHelpContentWidget) Repaint() {
	defer qt.Recovering("QHelpContentWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) RepaintDefault() {
	defer qt.Recovering("QHelpContentWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetDisabled
func callbackQHelpContentWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setDisabled", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QHelpContentWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setDisabled")
	}
}

func (ptr *QHelpContentWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QHelpContentWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpContentWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QHelpContentWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpContentWidget_SetFocus2
func callbackQHelpContentWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpContentWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setFocus2", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QHelpContentWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setFocus2")
	}
}

func (ptr *QHelpContentWidget) SetFocus2() {
	defer qt.Recovering("QHelpContentWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SetFocus2Default() {
	defer qt.Recovering("QHelpContentWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetHidden
func callbackQHelpContentWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QHelpContentWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setHidden", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QHelpContentWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "setHidden")
	}
}

func (ptr *QHelpContentWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QHelpContentWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpContentWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QHelpContentWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpContentWidget_Show
func callbackQHelpContentWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::show")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "show", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QHelpContentWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "show")
	}
}

func (ptr *QHelpContentWidget) Show() {
	defer qt.Recovering("QHelpContentWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowDefault() {
	defer qt.Recovering("QHelpContentWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowFullScreen
func callbackQHelpContentWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showFullScreen", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QHelpContentWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showFullScreen")
	}
}

func (ptr *QHelpContentWidget) ShowFullScreen() {
	defer qt.Recovering("QHelpContentWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QHelpContentWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMaximized
func callbackQHelpContentWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showMaximized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QHelpContentWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showMaximized")
	}
}

func (ptr *QHelpContentWidget) ShowMaximized() {
	defer qt.Recovering("QHelpContentWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QHelpContentWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMinimized
func callbackQHelpContentWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showMinimized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QHelpContentWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showMinimized")
	}
}

func (ptr *QHelpContentWidget) ShowMinimized() {
	defer qt.Recovering("QHelpContentWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QHelpContentWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowNormal
func callbackQHelpContentWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showNormal", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QHelpContentWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "showNormal")
	}
}

func (ptr *QHelpContentWidget) ShowNormal() {
	defer qt.Recovering("QHelpContentWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowNormalDefault() {
	defer qt.Recovering("QHelpContentWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_TabletEvent
func callbackQHelpContentWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "tabletEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "tabletEvent")
	}
}

func (ptr *QHelpContentWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpContentWidget_UpdateMicroFocus
func callbackQHelpContentWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "updateMicroFocus", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QHelpContentWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "updateMicroFocus")
	}
}

func (ptr *QHelpContentWidget) UpdateMicroFocus() {
	defer qt.Recovering("QHelpContentWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QHelpContentWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ChildEvent
func callbackQHelpContentWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpContentWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentWidget_ConnectNotify
func callbackQHelpContentWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpContentWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpContentWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpContentWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_CustomEvent
func callbackQHelpContentWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpContentWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_DeleteLater
func callbackQHelpContentWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpContentWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpContentWidget) DeleteLater() {
	defer qt.Recovering("QHelpContentWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()))
		C.QHelpContentWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentWidget) DeleteLaterDefault() {
	defer qt.Recovering("QHelpContentWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()))
		C.QHelpContentWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentWidget_DisconnectNotify
func callbackQHelpContentWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpContentWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpContentWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpContentWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpContentWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_EventFilter
func callbackQHelpContentWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpContentWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpContentWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpContentWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpContentWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpContentWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_MetaObject
func callbackQHelpContentWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpContentWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpContentWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpContentWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpContentWidget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpContentWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpContentWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpContentWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func (p *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return p
}

func (p *QHelpEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func (p *QHelpEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QHelpEngineCore_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	return n
}
func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	defer qt.Recovering("QHelpEngine::QHelpEngine")

	var collectionFileC = C.CString(collectionFile)
	defer C.free(unsafe.Pointer(collectionFileC))
	return NewQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(collectionFileC, core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	defer qt.Recovering("QHelpEngine::contentModel")

	if ptr.Pointer() != nil {
		return NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	defer qt.Recovering("QHelpEngine::contentWidget")

	if ptr.Pointer() != nil {
		return NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	defer qt.Recovering("QHelpEngine::indexModel")

	if ptr.Pointer() != nil {
		return NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	defer qt.Recovering("QHelpEngine::indexWidget")

	if ptr.Pointer() != nil {
		return NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	defer qt.Recovering("QHelpEngine::searchEngine")

	if ptr.Pointer() != nil {
		return NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	defer qt.Recovering("QHelpEngine::~QHelpEngine")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngine_TimerEvent
func callbackQHelpEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpEngine_ChildEvent
func callbackQHelpEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpEngine_ConnectNotify
func callbackQHelpEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngine_CustomEvent
func callbackQHelpEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpEngine_DeleteLater
func callbackQHelpEngine_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpEngine) DeleteLater() {
	defer qt.Recovering("QHelpEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngine) DeleteLaterDefault() {
	defer qt.Recovering("QHelpEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngine_DisconnectNotify
func callbackQHelpEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngine_Event
func callbackQHelpEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpEngine::event")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QHelpEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QHelpEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QHelpEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngine::event")

	if ptr.Pointer() != nil {
		return C.QHelpEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngine::event")

	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpEngine_EventFilter
func callbackQHelpEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpEngine::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpEngine_MetaObject
func callbackQHelpEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpEngine::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngine(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngine(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngine_MetaObjectDefault(ptr.Pointer()))
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

func (p *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return p
}

func (p *QHelpEngineCore) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QHelpEngineCore) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	defer qt.Recovering("QHelpEngineCore::autoSaveFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	defer qt.Recovering("QHelpEngineCore::collectionFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	defer qt.Recovering("QHelpEngineCore::currentFilter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	defer qt.Recovering("QHelpEngineCore::setAutoSaveFilter")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(save))))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	defer qt.Recovering("QHelpEngineCore::setCollectionFile")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	defer qt.Recovering("QHelpEngineCore::setCurrentFilter")

	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), filterNameC)
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	defer qt.Recovering("QHelpEngineCore::QHelpEngineCore")

	var collectionFileC = C.CString(collectionFile)
	defer C.free(unsafe.Pointer(collectionFileC))
	return NewQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(collectionFileC, core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	defer qt.Recovering("QHelpEngineCore::addCustomFilter")

	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		var attributesC = C.CString(strings.Join(attributes, "|"))
		defer C.free(unsafe.Pointer(attributesC))
		return C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), filterNameC, attributesC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	defer qt.Recovering("QHelpEngineCore::copyCollectionFile")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		return C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), fileNameC) != 0
	}
	return false
}

//export callbackQHelpEngineCore_CurrentFilterChanged
func callbackQHelpEngineCore_CurrentFilterChanged(ptr unsafe.Pointer, newFilter *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::currentFilterChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "currentFilterChanged"); signal != nil {
		signal.(func(string))(C.GoString(newFilter))
	}

}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	defer qt.Recovering("connect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	defer qt.Recovering("disconnect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "currentFilterChanged")
	}
}

func (ptr *QHelpEngineCore) CurrentFilterChanged(newFilter string) {
	defer qt.Recovering("QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		var newFilterC = C.CString(newFilter)
		defer C.free(unsafe.Pointer(newFilterC))
		C.QHelpEngineCore_CurrentFilterChanged(ptr.Pointer(), newFilterC)
	}
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	defer qt.Recovering("QHelpEngineCore::customFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::customValue")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), keyC, core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	defer qt.Recovering("QHelpEngineCore::documentationFileName")

	if ptr.Pointer() != nil {
		var namespaceNameC = C.CString(namespaceName)
		defer C.free(unsafe.Pointer(namespaceNameC))
		return C.GoString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), namespaceNameC))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	defer qt.Recovering("QHelpEngineCore::error")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) string {
	defer qt.Recovering("QHelpEngineCore::fileData")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url))))
	}
	return ""
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), filterNameC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QHelpEngineCore::findFile")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::metaData")

	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(documentationFileNameC, nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QHelpEngineCore) MetaData(documentationFileName string, name string) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::metaData")

	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(documentationFileNameC, nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	defer qt.Recovering("QHelpEngineCore::namespaceName")

	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(documentationFileNameC))
}

func (ptr *QHelpEngineCore) NamespaceName(documentationFileName string) string {
	defer qt.Recovering("QHelpEngineCore::namespaceName")

	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(documentationFileNameC))
}

//export callbackQHelpEngineCore_ReadersAboutToBeInvalidated
func callbackQHelpEngineCore_ReadersAboutToBeInvalidated(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::readersAboutToBeInvalidated")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "readersAboutToBeInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	defer qt.Recovering("disconnect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "readersAboutToBeInvalidated")
	}
}

func (ptr *QHelpEngineCore) ReadersAboutToBeInvalidated() {
	defer qt.Recovering("QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ReadersAboutToBeInvalidated(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	defer qt.Recovering("QHelpEngineCore::registerDocumentation")

	if ptr.Pointer() != nil {
		var documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
		return C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), documentationFileNameC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	defer qt.Recovering("QHelpEngineCore::registeredDocumentations")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomFilter")

	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		return C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), filterNameC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomValue")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), keyC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::setCustomValue")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), keyC, core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	defer qt.Recovering("QHelpEngineCore::setupData")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpEngineCore_SetupFinished
func callbackQHelpEngineCore_SetupFinished(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::setupFinished")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "setupFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "setupFinished")
	}
}

func (ptr *QHelpEngineCore) SetupFinished() {
	defer qt.Recovering("QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupFinished(ptr.Pointer())
	}
}

//export callbackQHelpEngineCore_SetupStarted
func callbackQHelpEngineCore_SetupStarted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::setupStarted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "setupStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "setupStarted")
	}
}

func (ptr *QHelpEngineCore) SetupStarted() {
	defer qt.Recovering("QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupStarted(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	defer qt.Recovering("QHelpEngineCore::unregisterDocumentation")

	if ptr.Pointer() != nil {
		var namespaceNameC = C.CString(namespaceName)
		defer C.free(unsafe.Pointer(namespaceNameC))
		return C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), namespaceNameC) != 0
	}
	return false
}

//export callbackQHelpEngineCore_Warning
func callbackQHelpEngineCore_Warning(ptr unsafe.Pointer, msg *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::warning")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "warning"); signal != nil {
		signal.(func(string))(C.GoString(msg))
	}

}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	defer qt.Recovering("connect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	defer qt.Recovering("disconnect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "warning")
	}
}

func (ptr *QHelpEngineCore) Warning(msg string) {
	defer qt.Recovering("QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QHelpEngineCore_Warning(ptr.Pointer(), msgC)
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	defer qt.Recovering("QHelpEngineCore::~QHelpEngineCore")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()))
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngineCore_TimerEvent
func callbackQHelpEngineCore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpEngineCore) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpEngineCore_ChildEvent
func callbackQHelpEngineCore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpEngineCore) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngineCore) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpEngineCore_ConnectNotify
func callbackQHelpEngineCore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpEngineCore::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpEngineCore::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpEngineCore) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngineCore::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngineCore::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_CustomEvent
func callbackQHelpEngineCore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpEngineCore) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngineCore) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpEngineCore_DeleteLater
func callbackQHelpEngineCore_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngineCore) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpEngineCore::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpEngineCore) DeleteLater() {
	defer qt.Recovering("QHelpEngineCore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()))
		C.QHelpEngineCore_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) DeleteLaterDefault() {
	defer qt.Recovering("QHelpEngineCore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()))
		C.QHelpEngineCore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngineCore_DisconnectNotify
func callbackQHelpEngineCore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpEngineCore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpEngineCore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpEngineCore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngineCore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngineCore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpEngineCore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_Event
func callbackQHelpEngineCore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpEngineCore::event")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpEngineCore) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpEngineCore::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QHelpEngineCore) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::event")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::event")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpEngineCore_EventFilter
func callbackQHelpEngineCore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpEngineCore::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngineCore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpEngineCore::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpEngineCore::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpEngineCore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpEngineCore_MetaObject
func callbackQHelpEngineCore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpEngineCore::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineCoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngineCore) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpEngineCore::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpEngineCore::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpEngineCore(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpEngineCore) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpEngineCore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngineCore) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpEngineCore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModel_ITF interface {
	core.QStringListModel_ITF
	QHelpIndexModel_PTR() *QHelpIndexModel
}

func (p *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return p
}

func (p *QHelpIndexModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QStringListModel_PTR().Pointer()
	}
	return nil
}

func (p *QHelpIndexModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QStringListModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpIndexModel) DestroyQHelpIndexModel() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	defer qt.Recovering("QHelpIndexModel::createIndex")

	if ptr.Pointer() != nil {
		var customFilterNameC = C.CString(customFilterName)
		defer C.free(unsafe.Pointer(customFilterNameC))
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), customFilterNameC)
	}
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::filter")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		var wildcardC = C.CString(wildcard)
		defer C.free(unsafe.Pointer(wildcardC))
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Filter(ptr.Pointer(), filterC, wildcardC))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_IndexCreated
func callbackQHelpIndexModel_IndexCreated(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreated")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "indexCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "indexCreated")
	}
}

func (ptr *QHelpIndexModel) IndexCreated() {
	defer qt.Recovering("QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreated(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_IndexCreationStarted
func callbackQHelpIndexModel_IndexCreationStarted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreationStarted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "indexCreationStarted")
	}
}

func (ptr *QHelpIndexModel) IndexCreationStarted() {
	defer qt.Recovering("QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	defer qt.Recovering("QHelpIndexModel::isCreatingIndex")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Data
func callbackQHelpIndexModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::data")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*core.QModelIndex, int) *core.QVariant)(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpIndexModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {
	defer qt.Recovering("connect QHelpIndexModel::data")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "data", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectData() {
	defer qt.Recovering("disconnect QHelpIndexModel::data")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "data")
	}
}

func (ptr *QHelpIndexModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QHelpIndexModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QHelpIndexModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Flags
func callbackQHelpIndexModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpIndexModel::flags")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectFlags(f func(index *core.QModelIndex) core.Qt__ItemFlag) {
	defer qt.Recovering("connect QHelpIndexModel::flags")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "flags", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectFlags() {
	defer qt.Recovering("disconnect QHelpIndexModel::flags")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "flags")
	}
}

func (ptr *QHelpIndexModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QHelpIndexModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QHelpIndexModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QHelpIndexModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpIndexModel_InsertRows
func callbackQHelpIndexModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::insertRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectInsertRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::insertRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "insertRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectInsertRows() {
	defer qt.Recovering("disconnect QHelpIndexModel::insertRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "insertRows")
	}
}

func (ptr *QHelpIndexModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RemoveRows
func callbackQHelpIndexModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::removeRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectRemoveRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::removeRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "removeRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRemoveRows() {
	defer qt.Recovering("disconnect QHelpIndexModel::removeRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "removeRows")
	}
}

func (ptr *QHelpIndexModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RowCount
func callbackQHelpIndexModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpIndexModel::rowCount")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	defer qt.Recovering("connect QHelpIndexModel::rowCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "rowCount", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRowCount() {
	defer qt.Recovering("disconnect QHelpIndexModel::rowCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "rowCount")
	}
}

func (ptr *QHelpIndexModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpIndexModel::rowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpIndexModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpIndexModel::rowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_SetData
func callbackQHelpIndexModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	defer qt.Recovering("callback QHelpIndexModel::setData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) ConnectSetData(f func(index *core.QModelIndex, value *core.QVariant, role int) bool) {
	defer qt.Recovering("connect QHelpIndexModel::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "setData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSetData() {
	defer qt.Recovering("disconnect QHelpIndexModel::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "setData")
	}
}

func (ptr *QHelpIndexModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpIndexModel::setData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpIndexModel::setData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Sibling
func callbackQHelpIndexModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::sibling")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QHelpIndexModel) ConnectSibling(f func(row int, column int, idx *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexModel::sibling")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "sibling", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSibling() {
	defer qt.Recovering("disconnect QHelpIndexModel::sibling")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "sibling")
	}
}

func (ptr *QHelpIndexModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Sort
func callbackQHelpIndexModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	defer qt.Recovering("callback QHelpIndexModel::sort")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "sort", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "sort")
	}
}

func (ptr *QHelpIndexModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpIndexModel_SupportedDropActions
func callbackQHelpIndexModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpIndexModel::supportedDropActions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpIndexModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QHelpIndexModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "supportedDropActions", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSupportedDropActions() {
	defer qt.Recovering("disconnect QHelpIndexModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "supportedDropActions")
	}
}

func (ptr *QHelpIndexModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QHelpIndexModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpIndexModel) SupportedDropActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QHelpIndexModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_Index
func callbackQHelpIndexModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::index")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpIndexModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexModel::index")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "index", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndex() {
	defer qt.Recovering("disconnect QHelpIndexModel::index")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "index")
	}
}

func (ptr *QHelpIndexModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_DropMimeData
func callbackQHelpIndexModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::dropMimeData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "dropMimeData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDropMimeData() {
	defer qt.Recovering("disconnect QHelpIndexModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "dropMimeData")
	}
}

func (ptr *QHelpIndexModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Buddy
func callbackQHelpIndexModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::buddy")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexModel::buddy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "buddy", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectBuddy() {
	defer qt.Recovering("disconnect QHelpIndexModel::buddy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "buddy")
	}
}

func (ptr *QHelpIndexModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_CanDropMimeData
func callbackQHelpIndexModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::canDropMimeData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "canDropMimeData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCanDropMimeData() {
	defer qt.Recovering("disconnect QHelpIndexModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "canDropMimeData")
	}
}

func (ptr *QHelpIndexModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_CanFetchMore
func callbackQHelpIndexModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::canFetchMore")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectCanFetchMore(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::canFetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "canFetchMore", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCanFetchMore() {
	defer qt.Recovering("disconnect QHelpIndexModel::canFetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "canFetchMore")
	}
}

func (ptr *QHelpIndexModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ColumnCount
func callbackQHelpIndexModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpIndexModel::columnCount")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {
	defer qt.Recovering("connect QHelpIndexModel::columnCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "columnCount", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectColumnCount() {
	defer qt.Recovering("disconnect QHelpIndexModel::columnCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "columnCount")
	}
}

func (ptr *QHelpIndexModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpIndexModel::columnCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpIndexModel::columnCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_FetchMore
func callbackQHelpIndexModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::fetchMore")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpIndexModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpIndexModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "fetchMore", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "fetchMore")
	}
}

func (ptr *QHelpIndexModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpIndexModel_HasChildren
func callbackQHelpIndexModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::hasChildren")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "hasChildren", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectHasChildren() {
	defer qt.Recovering("disconnect QHelpIndexModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "hasChildren")
	}
}

func (ptr *QHelpIndexModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_HeaderData
func callbackQHelpIndexModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::headerData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpIndexModel) ConnectHeaderData(f func(section int, orientation core.Qt__Orientation, role int) *core.QVariant) {
	defer qt.Recovering("connect QHelpIndexModel::headerData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "headerData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectHeaderData() {
	defer qt.Recovering("disconnect QHelpIndexModel::headerData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "headerData")
	}
}

func (ptr *QHelpIndexModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QHelpIndexModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QHelpIndexModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_InsertColumns
func callbackQHelpIndexModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::insertColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectInsertColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::insertColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "insertColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectInsertColumns() {
	defer qt.Recovering("disconnect QHelpIndexModel::insertColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "insertColumns")
	}
}

func (ptr *QHelpIndexModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MimeTypes
func callbackQHelpIndexModel_MimeTypes(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QHelpIndexModel::mimeTypes")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQHelpIndexModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QHelpIndexModel) ConnectMimeTypes(f func() []string) {
	defer qt.Recovering("connect QHelpIndexModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "mimeTypes", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMimeTypes() {
	defer qt.Recovering("disconnect QHelpIndexModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "mimeTypes")
	}
}

func (ptr *QHelpIndexModel) MimeTypes() []string {
	defer qt.Recovering("QHelpIndexModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpIndexModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpIndexModel) MimeTypesDefault() []string {
	defer qt.Recovering("QHelpIndexModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpIndexModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpIndexModel_MoveColumns
func callbackQHelpIndexModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QHelpIndexModel::moveColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QHelpIndexModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "moveColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMoveColumns() {
	defer qt.Recovering("disconnect QHelpIndexModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "moveColumns")
	}
}

func (ptr *QHelpIndexModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpIndexModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpIndexModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MoveRows
func callbackQHelpIndexModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QHelpIndexModel::moveRows")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QHelpIndexModel::moveRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "moveRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMoveRows() {
	defer qt.Recovering("disconnect QHelpIndexModel::moveRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "moveRows")
	}
}

func (ptr *QHelpIndexModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpIndexModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QHelpIndexModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Parent
func callbackQHelpIndexModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::parent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexModel::parent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "parent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectParent() {
	defer qt.Recovering("disconnect QHelpIndexModel::parent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "parent")
	}
}

func (ptr *QHelpIndexModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_RemoveColumns
func callbackQHelpIndexModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::removeColumns")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectRemoveColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexModel::removeColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "removeColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRemoveColumns() {
	defer qt.Recovering("disconnect QHelpIndexModel::removeColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "removeColumns")
	}
}

func (ptr *QHelpIndexModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ResetInternalData
func callbackQHelpIndexModel_ResetInternalData(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::resetInternalData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectResetInternalData(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "resetInternalData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectResetInternalData() {
	defer qt.Recovering("disconnect QHelpIndexModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "resetInternalData")
	}
}

func (ptr *QHelpIndexModel) ResetInternalData() {
	defer qt.Recovering("QHelpIndexModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ResetInternalDataDefault() {
	defer qt.Recovering("QHelpIndexModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_Revert
func callbackQHelpIndexModel_Revert(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::revert")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "revert", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "revert")
	}
}

func (ptr *QHelpIndexModel) Revert() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_SetHeaderData
func callbackQHelpIndexModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	defer qt.Recovering("callback QHelpIndexModel::setHeaderData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) ConnectSetHeaderData(f func(section int, orientation core.Qt__Orientation, value *core.QVariant, role int) bool) {
	defer qt.Recovering("connect QHelpIndexModel::setHeaderData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "setHeaderData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSetHeaderData() {
	defer qt.Recovering("disconnect QHelpIndexModel::setHeaderData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "setHeaderData")
	}
}

func (ptr *QHelpIndexModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpIndexModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetHeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QHelpIndexModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Span
func callbackQHelpIndexModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::span")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpIndexModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QHelpIndexModel::span")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "span", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSpan() {
	defer qt.Recovering("disconnect QHelpIndexModel::span")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "span")
	}
}

func (ptr *QHelpIndexModel) Span(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QHelpIndexModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QHelpIndexModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Submit
func callbackQHelpIndexModel_Submit(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::submit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpIndexModel) ConnectSubmit(f func() bool) {
	defer qt.Recovering("connect QHelpIndexModel::submit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "submit", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSubmit() {
	defer qt.Recovering("disconnect QHelpIndexModel::submit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "submit")
	}
}

func (ptr *QHelpIndexModel) Submit() bool {
	defer qt.Recovering("QHelpIndexModel::submit")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SubmitDefault() bool {
	defer qt.Recovering("QHelpIndexModel::submit")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SupportedDragActions
func callbackQHelpIndexModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpIndexModel::supportedDragActions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpIndexModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QHelpIndexModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "supportedDragActions", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSupportedDragActions() {
	defer qt.Recovering("disconnect QHelpIndexModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "supportedDragActions")
	}
}

func (ptr *QHelpIndexModel) SupportedDragActions() core.Qt__DropAction {
	defer qt.Recovering("QHelpIndexModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpIndexModel) SupportedDragActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QHelpIndexModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_TimerEvent
func callbackQHelpIndexModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpIndexModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpIndexModel_ChildEvent
func callbackQHelpIndexModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpIndexModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexModel_ConnectNotify
func callbackQHelpIndexModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpIndexModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpIndexModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpIndexModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_CustomEvent
func callbackQHelpIndexModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpIndexModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexModel_DeleteLater
func callbackQHelpIndexModel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpIndexModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpIndexModel) DeleteLater() {
	defer qt.Recovering("QHelpIndexModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()))
		C.QHelpIndexModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpIndexModel) DeleteLaterDefault() {
	defer qt.Recovering("QHelpIndexModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()))
		C.QHelpIndexModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpIndexModel_DisconnectNotify
func callbackQHelpIndexModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpIndexModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpIndexModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpIndexModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_Event
func callbackQHelpIndexModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::event")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpIndexModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpIndexModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QHelpIndexModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::event")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::event")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_EventFilter
func callbackQHelpIndexModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexModel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpIndexModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpIndexModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpIndexModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MetaObject
func callbackQHelpIndexModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexModel::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpIndexModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpIndexModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexModel(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpIndexModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpIndexModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpIndexModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func (p *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return p
}

func (p *QHelpIndexWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QListView_PTR().Pointer()
	}
	return nil
}

func (p *QHelpIndexWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QListView_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpIndexWidget) DestroyQHelpIndexWidget() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQHelpIndexWidget_ActivateCurrentItem
func callbackQHelpIndexWidget_ActivateCurrentItem(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::activateCurrentItem")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "activateCurrentItem"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexWidget) ConnectActivateCurrentItem(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "activateCurrentItem", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActivateCurrentItem() {
	defer qt.Recovering("disconnect QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "activateCurrentItem")
	}
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	defer qt.Recovering("QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_FilterIndices
func callbackQHelpIndexWidget_FilterIndices(ptr unsafe.Pointer, filter *C.char, wildcard *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::filterIndices")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "filterIndices"); signal != nil {
		signal.(func(string, string))(C.GoString(filter), C.GoString(wildcard))
	}

}

func (ptr *QHelpIndexWidget) ConnectFilterIndices(f func(filter string, wildcard string)) {
	defer qt.Recovering("connect QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "filterIndices", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFilterIndices(filter string, wildcard string) {
	defer qt.Recovering("disconnect QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "filterIndices")
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	defer qt.Recovering("QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		var wildcardC = C.CString(wildcard)
		defer C.free(unsafe.Pointer(wildcardC))
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), filterC, wildcardC)
	}
}

//export callbackQHelpIndexWidget_LinkActivated
func callbackQHelpIndexWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer, keyword *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::linkActivated")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(link), C.GoString(keyword))
	}

}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	defer qt.Recovering("connect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "linkActivated", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "linkActivated")
	}
}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {
	defer qt.Recovering("QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		var keywordC = C.CString(keyword)
		defer C.free(unsafe.Pointer(keywordC))
		C.QHelpIndexWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link), keywordC)
	}
}

//export callbackQHelpIndexWidget_CurrentChanged
func callbackQHelpIndexWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::currentChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpIndexWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "currentChanged", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "currentChanged")
	}
}

func (ptr *QHelpIndexWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpIndexWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpIndexWidget_DragLeaveEvent
func callbackQHelpIndexWidget_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragLeaveEvent")
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DragMoveEvent
func callbackQHelpIndexWidget_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragMoveEvent")
	}
}

func (ptr *QHelpIndexWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DropEvent
func callbackQHelpIndexWidget_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dropEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dropEvent")
	}
}

func (ptr *QHelpIndexWidget) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQHelpIndexWidget_HorizontalOffset
func callbackQHelpIndexWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpIndexWidget::horizontalOffset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) ConnectHorizontalOffset(f func() int) {
	defer qt.Recovering("connect QHelpIndexWidget::horizontalOffset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "horizontalOffset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHorizontalOffset() {
	defer qt.Recovering("disconnect QHelpIndexWidget::horizontalOffset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "horizontalOffset")
	}
}

func (ptr *QHelpIndexWidget) HorizontalOffset() int {
	defer qt.Recovering("QHelpIndexWidget::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpIndexWidget) HorizontalOffsetDefault() int {
	defer qt.Recovering("QHelpIndexWidget::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_IndexAt
func callbackQHelpIndexWidget_IndexAt(ptr unsafe.Pointer, p unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::indexAt")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(p)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(p)))
}

func (ptr *QHelpIndexWidget) ConnectIndexAt(f func(p *core.QPoint) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexWidget::indexAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "indexAt", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectIndexAt() {
	defer qt.Recovering("disconnect QHelpIndexWidget::indexAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "indexAt")
	}
}

func (ptr *QHelpIndexWidget) IndexAt(p core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexWidget::indexAt")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAt(ptr.Pointer(), core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) IndexAtDefault(p core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexWidget::indexAt")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_IsIndexHidden
func callbackQHelpIndexWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::isIndexHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpIndexWidget) ConnectIsIndexHidden(f func(index *core.QModelIndex) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::isIndexHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "isIndexHidden", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectIsIndexHidden() {
	defer qt.Recovering("disconnect QHelpIndexWidget::isIndexHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "isIndexHidden")
	}
}

func (ptr *QHelpIndexWidget) IsIndexHidden(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_MouseMoveEvent
func callbackQHelpIndexWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MouseReleaseEvent
func callbackQHelpIndexWidget_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MoveCursor
func callbackQHelpIndexWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::moveCursor")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpIndexWidget) ConnectMoveCursor(f func(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex) {
	defer qt.Recovering("connect QHelpIndexWidget::moveCursor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "moveCursor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveCursor() {
	defer qt.Recovering("disconnect QHelpIndexWidget::moveCursor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "moveCursor")
	}
}

func (ptr *QHelpIndexWidget) MoveCursor(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexWidget::moveCursor")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexWidget::moveCursor")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_PaintEvent
func callbackQHelpIndexWidget_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::paintEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "paintEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "paintEvent")
	}
}

func (ptr *QHelpIndexWidget) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

//export callbackQHelpIndexWidget_ResizeEvent
func callbackQHelpIndexWidget_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "resizeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "resizeEvent")
	}
}

func (ptr *QHelpIndexWidget) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

//export callbackQHelpIndexWidget_RowsAboutToBeRemoved
func callbackQHelpIndexWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "rowsAboutToBeRemoved")
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_RowsInserted
func callbackQHelpIndexWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::rowsInserted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "rowsInserted", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "rowsInserted")
	}
}

func (ptr *QHelpIndexWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_ScrollTo
func callbackQHelpIndexWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollTo")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollTo", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollTo")
	}
}

func (ptr *QHelpIndexWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_SelectionChanged
func callbackQHelpIndexWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::selectionChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	defer qt.Recovering("connect QHelpIndexWidget::selectionChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectionChanged", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QHelpIndexWidget::selectionChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectionChanged")
	}
}

func (ptr *QHelpIndexWidget) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	defer qt.Recovering("QHelpIndexWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QHelpIndexWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	defer qt.Recovering("QHelpIndexWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpIndexWidget_SetSelection
func callbackQHelpIndexWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setSelection")
	}
}

func (ptr *QHelpIndexWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpIndexWidget_StartDrag
func callbackQHelpIndexWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	defer qt.Recovering("callback QHelpIndexWidget::startDrag")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "startDrag", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "startDrag")
	}
}

func (ptr *QHelpIndexWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDrag(ptr.Pointer(), C.longlong(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpIndexWidget_UpdateGeometries
func callbackQHelpIndexWidget_UpdateGeometries(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::updateGeometries")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "updateGeometries", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "updateGeometries")
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_VerticalOffset
func callbackQHelpIndexWidget_VerticalOffset(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QHelpIndexWidget::verticalOffset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) ConnectVerticalOffset(f func() int) {
	defer qt.Recovering("connect QHelpIndexWidget::verticalOffset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "verticalOffset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVerticalOffset() {
	defer qt.Recovering("disconnect QHelpIndexWidget::verticalOffset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "verticalOffset")
	}
}

func (ptr *QHelpIndexWidget) VerticalOffset() int {
	defer qt.Recovering("QHelpIndexWidget::verticalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpIndexWidget) VerticalOffsetDefault() int {
	defer qt.Recovering("QHelpIndexWidget::verticalOffset")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_ViewOptions
func callbackQHelpIndexWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::viewOptions")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpIndexWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpIndexWidget) ConnectViewOptions(f func() *widgets.QStyleOptionViewItem) {
	defer qt.Recovering("connect QHelpIndexWidget::viewOptions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewOptions", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewOptions() {
	defer qt.Recovering("disconnect QHelpIndexWidget::viewOptions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewOptions")
	}
}

func (ptr *QHelpIndexWidget) ViewOptions() *widgets.QStyleOptionViewItem {
	defer qt.Recovering("QHelpIndexWidget::viewOptions")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptions(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	defer qt.Recovering("QHelpIndexWidget::viewOptions")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ViewportSizeHint
func callbackQHelpIndexWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::viewportSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectViewportSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpIndexWidget::viewportSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewportSizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewportSizeHint() {
	defer qt.Recovering("disconnect QHelpIndexWidget::viewportSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewportSizeHint")
	}
}

func (ptr *QHelpIndexWidget) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::viewportSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) ViewportSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::viewportSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRect
func callbackQHelpIndexWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::visualRect")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpIndexWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexWidget) ConnectVisualRect(f func(index *core.QModelIndex) *core.QRect) {
	defer qt.Recovering("connect QHelpIndexWidget::visualRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "visualRect", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVisualRect() {
	defer qt.Recovering("disconnect QHelpIndexWidget::visualRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "visualRect")
	}
}

func (ptr *QHelpIndexWidget) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QHelpIndexWidget::visualRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QHelpIndexWidget::visualRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRegionForSelection
func callbackQHelpIndexWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::visualRegionForSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpIndexWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpIndexWidget) ConnectVisualRegionForSelection(f func(selection *core.QItemSelection) *gui.QRegion) {
	defer qt.Recovering("connect QHelpIndexWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "visualRegionForSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVisualRegionForSelection() {
	defer qt.Recovering("disconnect QHelpIndexWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "visualRegionForSelection")
	}
}

func (ptr *QHelpIndexWidget) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QHelpIndexWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QHelpIndexWidget::visualRegionForSelection")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_WheelEvent
func callbackQHelpIndexWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "wheelEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *QHelpIndexWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpIndexWidget_ViewportEvent
func callbackQHelpIndexWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::viewportEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectViewportEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::viewportEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewportEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewportEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::viewportEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "viewportEvent")
	}
}

func (ptr *QHelpIndexWidget) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_ClearSelection
func callbackQHelpIndexWidget_ClearSelection(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::clearSelection")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectClearSelection(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::clearSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "clearSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectClearSelection() {
	defer qt.Recovering("disconnect QHelpIndexWidget::clearSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "clearSelection")
	}
}

func (ptr *QHelpIndexWidget) ClearSelection() {
	defer qt.Recovering("QHelpIndexWidget::clearSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ClearSelectionDefault() {
	defer qt.Recovering("QHelpIndexWidget::clearSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_CloseEditor
func callbackQHelpIndexWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	defer qt.Recovering("callback QHelpIndexWidget::closeEditor")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "closeEditor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "closeEditor")
	}
}

func (ptr *QHelpIndexWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_CommitData
func callbackQHelpIndexWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::commitData")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "commitData", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "commitData")
	}
}

func (ptr *QHelpIndexWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpIndexWidget_DragEnterEvent
func callbackQHelpIndexWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragEnterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "dragEnterEvent")
	}
}

func (ptr *QHelpIndexWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpIndexWidget_Edit
func callbackQHelpIndexWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::edit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectEdit(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "edit", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEdit() {
	defer qt.Recovering("disconnect QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "edit")
	}
}

func (ptr *QHelpIndexWidget) Edit(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) EditDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_Edit2
func callbackQHelpIndexWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::edit")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectEdit2(f func(index *core.QModelIndex, trigger widgets.QAbstractItemView__EditTrigger, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "edit2", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEdit2() {
	defer qt.Recovering("disconnect QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "edit2")
	}
}

func (ptr *QHelpIndexWidget) Edit2(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Edit2(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::edit")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_EditorDestroyed
func callbackQHelpIndexWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::editorDestroyed")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "editorDestroyed", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "editorDestroyed")
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpIndexWidget_FocusInEvent
func callbackQHelpIndexWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusInEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusInEvent")
	}
}

func (ptr *QHelpIndexWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_FocusNextPrevChild
func callbackQHelpIndexWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpIndexWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusNextPrevChild", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusNextPrevChild")
	}
}

func (ptr *QHelpIndexWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QHelpIndexWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QHelpIndexWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_FocusOutEvent
func callbackQHelpIndexWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusOutEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "focusOutEvent")
	}
}

func (ptr *QHelpIndexWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_InputMethodEvent
func callbackQHelpIndexWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "inputMethodEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "inputMethodEvent")
	}
}

func (ptr *QHelpIndexWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpIndexWidget_InputMethodQuery
func callbackQHelpIndexWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpIndexWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpIndexWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QHelpIndexWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "inputMethodQuery", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QHelpIndexWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "inputMethodQuery")
	}
}

func (ptr *QHelpIndexWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpIndexWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpIndexWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_KeyPressEvent
func callbackQHelpIndexWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyPressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyPressEvent")
	}
}

func (ptr *QHelpIndexWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_KeyboardSearch
func callbackQHelpIndexWidget_KeyboardSearch(ptr unsafe.Pointer, search *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::keyboardSearch")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyboardSearch", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyboardSearch")
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpIndexWidget_KeyboardSearch(ptr.Pointer(), searchC)
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpIndexWidget_KeyboardSearchDefault(ptr.Pointer(), searchC)
	}
}

//export callbackQHelpIndexWidget_MouseDoubleClickEvent
func callbackQHelpIndexWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_MousePressEvent
func callbackQHelpIndexWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mousePressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *QHelpIndexWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_Reset
func callbackQHelpIndexWidget_Reset(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::reset")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "reset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "reset")
	}
}

func (ptr *QHelpIndexWidget) Reset() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ResetDefault() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToBottom
func callbackQHelpIndexWidget_ScrollToBottom(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollToBottom")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollToBottom(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollToBottom")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollToBottom", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollToBottom() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollToBottom")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollToBottom")
	}
}

func (ptr *QHelpIndexWidget) ScrollToBottom() {
	defer qt.Recovering("QHelpIndexWidget::scrollToBottom")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ScrollToBottomDefault() {
	defer qt.Recovering("QHelpIndexWidget::scrollToBottom")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToTop
func callbackQHelpIndexWidget_ScrollToTop(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollToTop")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollToTop(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollToTop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollToTop", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollToTop() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollToTop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollToTop")
	}
}

func (ptr *QHelpIndexWidget) ScrollToTop() {
	defer qt.Recovering("QHelpIndexWidget::scrollToTop")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ScrollToTopDefault() {
	defer qt.Recovering("QHelpIndexWidget::scrollToTop")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectAll
func callbackQHelpIndexWidget_SelectAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::selectAll")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectAll", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectAll")
	}
}

func (ptr *QHelpIndexWidget) SelectAll() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectionCommand
func callbackQHelpIndexWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QHelpIndexWidget::selectionCommand")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpIndexWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpIndexWidget) ConnectSelectionCommand(f func(index *core.QModelIndex, event *core.QEvent) core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("connect QHelpIndexWidget::selectionCommand")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectionCommand", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectionCommand() {
	defer qt.Recovering("disconnect QHelpIndexWidget::selectionCommand")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "selectionCommand")
	}
}

func (ptr *QHelpIndexWidget) SelectionCommand(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	defer qt.Recovering("QHelpIndexWidget::selectionCommand")

	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommand(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	defer qt.Recovering("QHelpIndexWidget::selectionCommand")

	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpIndexWidget_SetCurrentIndex
func callbackQHelpIndexWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setCurrentIndex")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetCurrentIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setCurrentIndex", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetCurrentIndex() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setCurrentIndex")
	}
}

func (ptr *QHelpIndexWidget) SetCurrentIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetModel
func callbackQHelpIndexWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setModel")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setModel")
	}
}

func (ptr *QHelpIndexWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpIndexWidget_SetRootIndex
func callbackQHelpIndexWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setRootIndex")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setRootIndex", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setRootIndex")
	}
}

func (ptr *QHelpIndexWidget) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetSelectionModel
func callbackQHelpIndexWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelectionModel")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setSelectionModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setSelectionModel")
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpIndexWidget_SizeHintForColumn
func callbackQHelpIndexWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	defer qt.Recovering("callback QHelpIndexWidget::sizeHintForColumn")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpIndexWidget) ConnectSizeHintForColumn(f func(column int) int) {
	defer qt.Recovering("connect QHelpIndexWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHintForColumn", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHintForColumn() {
	defer qt.Recovering("disconnect QHelpIndexWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHintForColumn")
	}
}

func (ptr *QHelpIndexWidget) SizeHintForColumn(column int) int {
	defer qt.Recovering("QHelpIndexWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForColumn(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SizeHintForColumnDefault(column int) int {
	defer qt.Recovering("QHelpIndexWidget::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_SizeHintForRow
func callbackQHelpIndexWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	defer qt.Recovering("callback QHelpIndexWidget::sizeHintForRow")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpIndexWidget) ConnectSizeHintForRow(f func(row int) int) {
	defer qt.Recovering("connect QHelpIndexWidget::sizeHintForRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHintForRow", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHintForRow() {
	defer qt.Recovering("disconnect QHelpIndexWidget::sizeHintForRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHintForRow")
	}
}

func (ptr *QHelpIndexWidget) SizeHintForRow(row int) int {
	defer qt.Recovering("QHelpIndexWidget::sizeHintForRow")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRow(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SizeHintForRowDefault(row int) int {
	defer qt.Recovering("QHelpIndexWidget::sizeHintForRow")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Update
func callbackQHelpIndexWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::update")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdate(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "update", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QHelpIndexWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "update")
	}
}

func (ptr *QHelpIndexWidget) Update(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) UpdateDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_ContextMenuEvent
func callbackQHelpIndexWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "contextMenuEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "contextMenuEvent")
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpIndexWidget_MinimumSizeHint
func callbackQHelpIndexWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpIndexWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "minimumSizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QHelpIndexWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "minimumSizeHint")
	}
}

func (ptr *QHelpIndexWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ScrollContentsBy
func callbackQHelpIndexWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollContentsBy")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollContentsBy", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "scrollContentsBy")
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsBy(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpIndexWidget_SetupViewport
func callbackQHelpIndexWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setupViewport")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setupViewport", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setupViewport")
	}
}

func (ptr *QHelpIndexWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpIndexWidget_SizeHint
func callbackQHelpIndexWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::sizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpIndexWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QHelpIndexWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "sizeHint")
	}
}

func (ptr *QHelpIndexWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpIndexWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ChangeEvent
func callbackQHelpIndexWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "changeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "changeEvent")
	}
}

func (ptr *QHelpIndexWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpIndexWidget_ActionEvent
func callbackQHelpIndexWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "actionEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "actionEvent")
	}
}

func (ptr *QHelpIndexWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpIndexWidget_EnterEvent
func callbackQHelpIndexWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "enterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "enterEvent")
	}
}

func (ptr *QHelpIndexWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_HideEvent
func callbackQHelpIndexWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hideEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hideEvent")
	}
}

func (ptr *QHelpIndexWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpIndexWidget_LeaveEvent
func callbackQHelpIndexWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "leaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "leaveEvent")
	}
}

func (ptr *QHelpIndexWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_MoveEvent
func callbackQHelpIndexWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "moveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "moveEvent")
	}
}

func (ptr *QHelpIndexWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpIndexWidget_SetEnabled
func callbackQHelpIndexWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setEnabled", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setEnabled")
	}
}

func (ptr *QHelpIndexWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QHelpIndexWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpIndexWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QHelpIndexWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetStyleSheet
func callbackQHelpIndexWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QHelpIndexWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setStyleSheet", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setStyleSheet")
	}
}

func (ptr *QHelpIndexWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QHelpIndexWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpIndexWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpIndexWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QHelpIndexWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpIndexWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpIndexWidget_SetVisible
func callbackQHelpIndexWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setVisible")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setVisible", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setVisible")
	}
}

func (ptr *QHelpIndexWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpIndexWidget_SetWindowModified
func callbackQHelpIndexWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setWindowModified", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setWindowModified")
	}
}

func (ptr *QHelpIndexWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QHelpIndexWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpIndexWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QHelpIndexWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetWindowTitle
func callbackQHelpIndexWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QHelpIndexWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setWindowTitle", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setWindowTitle")
	}
}

func (ptr *QHelpIndexWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QHelpIndexWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpIndexWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpIndexWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QHelpIndexWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpIndexWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpIndexWidget_ShowEvent
func callbackQHelpIndexWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showEvent")
	}
}

func (ptr *QHelpIndexWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpIndexWidget_Close
func callbackQHelpIndexWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::close")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpIndexWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QHelpIndexWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "close", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QHelpIndexWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "close")
	}
}

func (ptr *QHelpIndexWidget) Close() bool {
	defer qt.Recovering("QHelpIndexWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) CloseDefault() bool {
	defer qt.Recovering("QHelpIndexWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_CloseEvent
func callbackQHelpIndexWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "closeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "closeEvent")
	}
}

func (ptr *QHelpIndexWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpIndexWidget_HasHeightForWidth
func callbackQHelpIndexWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpIndexWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QHelpIndexWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hasHeightForWidth", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QHelpIndexWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hasHeightForWidth")
	}
}

func (ptr *QHelpIndexWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QHelpIndexWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QHelpIndexWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_HeightForWidth
func callbackQHelpIndexWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QHelpIndexWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpIndexWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QHelpIndexWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "heightForWidth", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QHelpIndexWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "heightForWidth")
	}
}

func (ptr *QHelpIndexWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QHelpIndexWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QHelpIndexWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Hide
func callbackQHelpIndexWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::hide")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hide", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QHelpIndexWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "hide")
	}
}

func (ptr *QHelpIndexWidget) Hide() {
	defer qt.Recovering("QHelpIndexWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) HideDefault() {
	defer qt.Recovering("QHelpIndexWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_KeyReleaseEvent
func callbackQHelpIndexWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "keyReleaseEvent")
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_Lower
func callbackQHelpIndexWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::lower")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "lower", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QHelpIndexWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "lower")
	}
}

func (ptr *QHelpIndexWidget) Lower() {
	defer qt.Recovering("QHelpIndexWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) LowerDefault() {
	defer qt.Recovering("QHelpIndexWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_NativeEvent
func callbackQHelpIndexWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QHelpIndexWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "nativeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "nativeEvent")
	}
}

func (ptr *QHelpIndexWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpIndexWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpIndexWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpIndexWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpIndexWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_Raise
func callbackQHelpIndexWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::raise")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "raise", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QHelpIndexWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "raise")
	}
}

func (ptr *QHelpIndexWidget) Raise() {
	defer qt.Recovering("QHelpIndexWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) RaiseDefault() {
	defer qt.Recovering("QHelpIndexWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Repaint
func callbackQHelpIndexWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "repaint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QHelpIndexWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "repaint")
	}
}

func (ptr *QHelpIndexWidget) Repaint() {
	defer qt.Recovering("QHelpIndexWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) RepaintDefault() {
	defer qt.Recovering("QHelpIndexWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetDisabled
func callbackQHelpIndexWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setDisabled", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setDisabled")
	}
}

func (ptr *QHelpIndexWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QHelpIndexWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpIndexWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QHelpIndexWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpIndexWidget_SetFocus2
func callbackQHelpIndexWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpIndexWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setFocus2", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setFocus2")
	}
}

func (ptr *QHelpIndexWidget) SetFocus2() {
	defer qt.Recovering("QHelpIndexWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SetFocus2Default() {
	defer qt.Recovering("QHelpIndexWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetHidden
func callbackQHelpIndexWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setHidden", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "setHidden")
	}
}

func (ptr *QHelpIndexWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QHelpIndexWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpIndexWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QHelpIndexWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpIndexWidget_Show
func callbackQHelpIndexWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::show")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "show", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QHelpIndexWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "show")
	}
}

func (ptr *QHelpIndexWidget) Show() {
	defer qt.Recovering("QHelpIndexWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowDefault() {
	defer qt.Recovering("QHelpIndexWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowFullScreen
func callbackQHelpIndexWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showFullScreen", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showFullScreen")
	}
}

func (ptr *QHelpIndexWidget) ShowFullScreen() {
	defer qt.Recovering("QHelpIndexWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QHelpIndexWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMaximized
func callbackQHelpIndexWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showMaximized", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showMaximized")
	}
}

func (ptr *QHelpIndexWidget) ShowMaximized() {
	defer qt.Recovering("QHelpIndexWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QHelpIndexWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMinimized
func callbackQHelpIndexWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showMinimized", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showMinimized")
	}
}

func (ptr *QHelpIndexWidget) ShowMinimized() {
	defer qt.Recovering("QHelpIndexWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QHelpIndexWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowNormal
func callbackQHelpIndexWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showNormal", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "showNormal")
	}
}

func (ptr *QHelpIndexWidget) ShowNormal() {
	defer qt.Recovering("QHelpIndexWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowNormalDefault() {
	defer qt.Recovering("QHelpIndexWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_TabletEvent
func callbackQHelpIndexWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "tabletEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "tabletEvent")
	}
}

func (ptr *QHelpIndexWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpIndexWidget_UpdateMicroFocus
func callbackQHelpIndexWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "updateMicroFocus", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QHelpIndexWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "updateMicroFocus")
	}
}

func (ptr *QHelpIndexWidget) UpdateMicroFocus() {
	defer qt.Recovering("QHelpIndexWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QHelpIndexWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ChildEvent
func callbackQHelpIndexWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpIndexWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexWidget_ConnectNotify
func callbackQHelpIndexWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpIndexWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpIndexWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpIndexWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_CustomEvent
func callbackQHelpIndexWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpIndexWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_DeleteLater
func callbackQHelpIndexWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpIndexWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpIndexWidget) DeleteLater() {
	defer qt.Recovering("QHelpIndexWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()))
		C.QHelpIndexWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpIndexWidget) DeleteLaterDefault() {
	defer qt.Recovering("QHelpIndexWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()))
		C.QHelpIndexWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpIndexWidget_DisconnectNotify
func callbackQHelpIndexWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpIndexWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpIndexWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpIndexWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_EventFilter
func callbackQHelpIndexWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpIndexWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpIndexWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpIndexWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpIndexWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpIndexWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_MetaObject
func callbackQHelpIndexWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpIndexWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpIndexWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpIndexWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpIndexWidget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpIndexWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpIndexWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpIndexWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngine_ITF interface {
	core.QObject_ITF
	QHelpSearchEngine_PTR() *QHelpSearchEngine
}

func (p *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return p
}

func (p *QHelpSearchEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	return n
}
func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	defer qt.Recovering("QHelpSearchEngine::QHelpSearchEngine")

	return NewQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
}

//export callbackQHelpSearchEngine_CancelIndexing
func callbackQHelpSearchEngine_CancelIndexing(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::cancelIndexing")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "cancelIndexing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectCancelIndexing(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::cancelIndexing")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "cancelIndexing", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelIndexing() {
	defer qt.Recovering("disconnect QHelpSearchEngine::cancelIndexing")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "cancelIndexing")
	}
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	defer qt.Recovering("QHelpSearchEngine::cancelIndexing")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_CancelSearching
func callbackQHelpSearchEngine_CancelSearching(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::cancelSearching")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "cancelSearching"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectCancelSearching(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::cancelSearching")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "cancelSearching", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelSearching() {
	defer qt.Recovering("disconnect QHelpSearchEngine::cancelSearching")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "cancelSearching")
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	defer qt.Recovering("QHelpSearchEngine::cancelSearching")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	defer qt.Recovering("QHelpSearchEngine::hitCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchEngine_HitCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpSearchEngine_IndexingFinished
func callbackQHelpSearchEngine_IndexingFinished(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingFinished")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "indexingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "indexingFinished")
	}
}

func (ptr *QHelpSearchEngine) IndexingFinished() {
	defer qt.Recovering("QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingFinished(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_IndexingStarted
func callbackQHelpSearchEngine_IndexingStarted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingStarted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "indexingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "indexingStarted")
	}
}

func (ptr *QHelpSearchEngine) IndexingStarted() {
	defer qt.Recovering("QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchEngine::queryWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchEngine_ReindexDocumentation
func callbackQHelpSearchEngine_ReindexDocumentation(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::reindexDocumentation")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "reindexDocumentation"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectReindexDocumentation(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::reindexDocumentation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "reindexDocumentation", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectReindexDocumentation() {
	defer qt.Recovering("disconnect QHelpSearchEngine::reindexDocumentation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "reindexDocumentation")
	}
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	defer qt.Recovering("QHelpSearchEngine::reindexDocumentation")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	defer qt.Recovering("QHelpSearchEngine::resultWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
	}
	return nil
}

//export callbackQHelpSearchEngine_SearchingFinished
func callbackQHelpSearchEngine_SearchingFinished(ptr unsafe.Pointer, hits C.int) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingFinished")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "searchingFinished"); signal != nil {
		signal.(func(int))(int(int32(hits)))
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "searchingFinished")
	}
}

func (ptr *QHelpSearchEngine) SearchingFinished(hits int) {
	defer qt.Recovering("QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingFinished(ptr.Pointer(), C.int(int32(hits)))
	}
}

//export callbackQHelpSearchEngine_SearchingStarted
func callbackQHelpSearchEngine_SearchingStarted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingStarted")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "searchingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "searchingStarted")
	}
}

func (ptr *QHelpSearchEngine) SearchingStarted() {
	defer qt.Recovering("QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	defer qt.Recovering("QHelpSearchEngine::~QHelpSearchEngine")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()))
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchEngine_TimerEvent
func callbackQHelpSearchEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpSearchEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchEngine_ChildEvent
func callbackQHelpSearchEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpSearchEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchEngine_ConnectNotify
func callbackQHelpSearchEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpSearchEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_CustomEvent
func callbackQHelpSearchEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpSearchEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchEngine_DeleteLater
func callbackQHelpSearchEngine_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpSearchEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpSearchEngine) DeleteLater() {
	defer qt.Recovering("QHelpSearchEngine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()))
		C.QHelpSearchEngine_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) DeleteLaterDefault() {
	defer qt.Recovering("QHelpSearchEngine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()))
		C.QHelpSearchEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchEngine_DisconnectNotify
func callbackQHelpSearchEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_Event
func callbackQHelpSearchEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchEngine::event")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpSearchEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpSearchEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QHelpSearchEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchEngine::event")

	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpSearchEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchEngine::event")

	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_EventFilter
func callbackQHelpSearchEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchEngine::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpSearchEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpSearchEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpSearchEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_MetaObject
func callbackQHelpSearchEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchEngine::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpSearchEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpSearchEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchEngine(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpSearchEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QHelpSearchQuery::FieldName
type QHelpSearchQuery__FieldName int64

const (
	QHelpSearchQuery__DEFAULT = QHelpSearchQuery__FieldName(0)
	QHelpSearchQuery__FUZZY   = QHelpSearchQuery__FieldName(1)
	QHelpSearchQuery__WITHOUT = QHelpSearchQuery__FieldName(2)
	QHelpSearchQuery__PHRASE  = QHelpSearchQuery__FieldName(3)
	QHelpSearchQuery__ALL     = QHelpSearchQuery__FieldName(4)
	QHelpSearchQuery__ATLEAST = QHelpSearchQuery__FieldName(5)
)

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (p *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return p
}

func (p *QHelpSearchQuery) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHelpSearchQuery) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpSearchQuery) DestroyQHelpSearchQuery() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQHelpSearchQuery() *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
	runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
}

func NewQHelpSearchQuery2(field QHelpSearchQuery__FieldName, wordList []string) *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	var wordListC = C.CString(strings.Join(wordList, "|"))
	defer C.free(unsafe.Pointer(wordListC))
	var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery2(C.longlong(field), wordListC))
	runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
}

func (ptr *QHelpSearchQuery) FieldName() QHelpSearchQuery__FieldName {
	defer qt.Recovering("QHelpSearchQuery::fieldName")

	if ptr.Pointer() != nil {
		return QHelpSearchQuery__FieldName(C.QHelpSearchQuery_FieldName(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpSearchQuery) SetFieldName(vfi QHelpSearchQuery__FieldName) {
	defer qt.Recovering("QHelpSearchQuery::setFieldName")

	if ptr.Pointer() != nil {
		C.QHelpSearchQuery_SetFieldName(ptr.Pointer(), C.longlong(vfi))
	}
}

func (ptr *QHelpSearchQuery) WordList() []string {
	defer qt.Recovering("QHelpSearchQuery::wordList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpSearchQuery_WordList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpSearchQuery) SetWordList(vqs []string) {
	defer qt.Recovering("QHelpSearchQuery::setWordList")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(strings.Join(vqs, "|"))
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQuery_SetWordList(ptr.Pointer(), vqsC)
	}
}

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func (p *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return p
}

func (p *QHelpSearchQueryWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchQueryWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::isCompactMode")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchQueryWidget::QHelpSearchQueryWidget")

	return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::collapseExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::expandExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Search
func callbackQHelpSearchQueryWidget_Search(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::search")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "search")
	}
}

func (ptr *QHelpSearchQueryWidget) Search() {
	defer qt.Recovering("QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Search(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	defer qt.Recovering("QHelpSearchQueryWidget::~QHelpSearchQueryWidget")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()))
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchQueryWidget_ActionEvent
func callbackQHelpSearchQueryWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "actionEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "actionEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragEnterEvent
func callbackQHelpSearchQueryWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragEnterEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragLeaveEvent
func callbackQHelpSearchQueryWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragLeaveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragMoveEvent
func callbackQHelpSearchQueryWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dragMoveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DropEvent
func callbackQHelpSearchQueryWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dropEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "dropEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_EnterEvent
func callbackQHelpSearchQueryWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "enterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "enterEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusInEvent
func callbackQHelpSearchQueryWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusInEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusInEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusOutEvent
func callbackQHelpSearchQueryWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusOutEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_HideEvent
func callbackQHelpSearchQueryWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hideEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hideEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_LeaveEvent
func callbackQHelpSearchQueryWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "leaveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MinimumSizeHint
func callbackQHelpSearchQueryWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchQueryWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "minimumSizeHint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "minimumSizeHint")
	}
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QHelpSearchQueryWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpSearchQueryWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_MoveEvent
func callbackQHelpSearchQueryWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "moveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "moveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_PaintEvent
func callbackQHelpSearchQueryWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::paintEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "paintEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "paintEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SetEnabled
func callbackQHelpSearchQueryWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setEnabled", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setEnabled")
	}
}

func (ptr *QHelpSearchQueryWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetStyleSheet
func callbackQHelpSearchQueryWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setStyleSheet", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setStyleSheet")
	}
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QHelpSearchQueryWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchQueryWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QHelpSearchQueryWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchQueryWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpSearchQueryWidget_SetVisible
func callbackQHelpSearchQueryWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setVisible")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setVisible", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setVisible")
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowModified
func callbackQHelpSearchQueryWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setWindowModified", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setWindowModified")
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowTitle
func callbackQHelpSearchQueryWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setWindowTitle", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setWindowTitle")
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QHelpSearchQueryWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQueryWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QHelpSearchQueryWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQueryWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpSearchQueryWidget_ShowEvent
func callbackQHelpSearchQueryWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SizeHint
func callbackQHelpSearchQueryWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchQueryWidget::sizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "sizeHint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "sizeHint")
	}
}

func (ptr *QHelpSearchQueryWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QHelpSearchQueryWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpSearchQueryWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_ChangeEvent
func callbackQHelpSearchQueryWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "changeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "changeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Close
func callbackQHelpSearchQueryWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchQueryWidget::close")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchQueryWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "close", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "close")
	}
}

func (ptr *QHelpSearchQueryWidget) Close() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) CloseDefault() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_CloseEvent
func callbackQHelpSearchQueryWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "closeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "closeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ContextMenuEvent
func callbackQHelpSearchQueryWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "contextMenuEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusNextPrevChild
func callbackQHelpSearchQueryWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QHelpSearchQueryWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusNextPrevChild", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "focusNextPrevChild")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_HasHeightForWidth
func callbackQHelpSearchQueryWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchQueryWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchQueryWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hasHeightForWidth", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hasHeightForWidth")
	}
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_HeightForWidth
func callbackQHelpSearchQueryWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QHelpSearchQueryWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchQueryWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "heightForWidth", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "heightForWidth")
	}
}

func (ptr *QHelpSearchQueryWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QHelpSearchQueryWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpSearchQueryWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QHelpSearchQueryWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_Hide
func callbackQHelpSearchQueryWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::hide")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hide", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "hide")
	}
}

func (ptr *QHelpSearchQueryWidget) Hide() {
	defer qt.Recovering("QHelpSearchQueryWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) HideDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodEvent
func callbackQHelpSearchQueryWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "inputMethodEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodQuery
func callbackQHelpSearchQueryWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchQueryWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "inputMethodQuery", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "inputMethodQuery")
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_KeyPressEvent
func callbackQHelpSearchQueryWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "keyPressEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_KeyReleaseEvent
func callbackQHelpSearchQueryWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "keyReleaseEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Lower
func callbackQHelpSearchQueryWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::lower")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "lower", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "lower")
	}
}

func (ptr *QHelpSearchQueryWidget) Lower() {
	defer qt.Recovering("QHelpSearchQueryWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) LowerDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_MouseDoubleClickEvent
func callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseMoveEvent
func callbackQHelpSearchQueryWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MousePressEvent
func callbackQHelpSearchQueryWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseReleaseEvent
func callbackQHelpSearchQueryWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_NativeEvent
func callbackQHelpSearchQueryWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QHelpSearchQueryWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QHelpSearchQueryWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "nativeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "nativeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpSearchQueryWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpSearchQueryWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_Raise
func callbackQHelpSearchQueryWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::raise")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "raise", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "raise")
	}
}

func (ptr *QHelpSearchQueryWidget) Raise() {
	defer qt.Recovering("QHelpSearchQueryWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) RaiseDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Repaint
func callbackQHelpSearchQueryWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "repaint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "repaint")
	}
}

func (ptr *QHelpSearchQueryWidget) Repaint() {
	defer qt.Recovering("QHelpSearchQueryWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) RepaintDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ResizeEvent
func callbackQHelpSearchQueryWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "resizeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SetDisabled
func callbackQHelpSearchQueryWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setDisabled", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setDisabled")
	}
}

func (ptr *QHelpSearchQueryWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchQueryWidget_SetFocus2
func callbackQHelpSearchQueryWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setFocus2", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setFocus2")
	}
}

func (ptr *QHelpSearchQueryWidget) SetFocus2() {
	defer qt.Recovering("QHelpSearchQueryWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) SetFocus2Default() {
	defer qt.Recovering("QHelpSearchQueryWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_SetHidden
func callbackQHelpSearchQueryWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setHidden", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "setHidden")
	}
}

func (ptr *QHelpSearchQueryWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchQueryWidget_Show
func callbackQHelpSearchQueryWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::show")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "show", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "show")
	}
}

func (ptr *QHelpSearchQueryWidget) Show() {
	defer qt.Recovering("QHelpSearchQueryWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowFullScreen
func callbackQHelpSearchQueryWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showFullScreen", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showFullScreen")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreen() {
	defer qt.Recovering("QHelpSearchQueryWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMaximized
func callbackQHelpSearchQueryWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showMaximized", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showMaximized")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMaximized() {
	defer qt.Recovering("QHelpSearchQueryWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMinimized
func callbackQHelpSearchQueryWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showMinimized", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showMinimized")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMinimized() {
	defer qt.Recovering("QHelpSearchQueryWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowNormal
func callbackQHelpSearchQueryWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showNormal", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "showNormal")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormal() {
	defer qt.Recovering("QHelpSearchQueryWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormalDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_TabletEvent
func callbackQHelpSearchQueryWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "tabletEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Update
func callbackQHelpSearchQueryWidget_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::update")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "update", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "update")
	}
}

func (ptr *QHelpSearchQueryWidget) Update() {
	defer qt.Recovering("QHelpSearchQueryWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Update(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_UpdateMicroFocus
func callbackQHelpSearchQueryWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "updateMicroFocus", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "updateMicroFocus")
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocus() {
	defer qt.Recovering("QHelpSearchQueryWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_WheelEvent
func callbackQHelpSearchQueryWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_TimerEvent
func callbackQHelpSearchQueryWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ChildEvent
func callbackQHelpSearchQueryWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ConnectNotify
func callbackQHelpSearchQueryWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_CustomEvent
func callbackQHelpSearchQueryWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DeleteLater
func callbackQHelpSearchQueryWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLater() {
	defer qt.Recovering("QHelpSearchQueryWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()))
		C.QHelpSearchQueryWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLaterDefault() {
	defer qt.Recovering("QHelpSearchQueryWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()))
		C.QHelpSearchQueryWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchQueryWidget_DisconnectNotify
func callbackQHelpSearchQueryWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_EventFilter
func callbackQHelpSearchQueryWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchQueryWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpSearchQueryWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchQueryWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_MetaObject
func callbackQHelpSearchQueryWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchQueryWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchQueryWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchQueryWidget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpSearchQueryWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchQueryWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchQueryWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObjectDefault(ptr.Pointer()))
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

func (p *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return p
}

func (p *QHelpSearchResultWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchResultWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	defer qt.Recovering("QHelpSearchResultWidget::linkAt")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_RequestShowLink
func callbackQHelpSearchResultWidget_RequestShowLink(ptr unsafe.Pointer, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::requestShowLink")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "requestShowLink")
	}
}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RequestShowLink(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	defer qt.Recovering("QHelpSearchResultWidget::~QHelpSearchResultWidget")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()))
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchResultWidget_ActionEvent
func callbackQHelpSearchResultWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "actionEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "actionEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragEnterEvent
func callbackQHelpSearchResultWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragEnterEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragLeaveEvent
func callbackQHelpSearchResultWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragLeaveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragMoveEvent
func callbackQHelpSearchResultWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dragMoveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DropEvent
func callbackQHelpSearchResultWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dropEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "dropEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_EnterEvent
func callbackQHelpSearchResultWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "enterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "enterEvent")
	}
}

func (ptr *QHelpSearchResultWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusInEvent
func callbackQHelpSearchResultWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusInEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusInEvent")
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusOutEvent
func callbackQHelpSearchResultWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusOutEvent")
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_HideEvent
func callbackQHelpSearchResultWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hideEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hideEvent")
	}
}

func (ptr *QHelpSearchResultWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_LeaveEvent
func callbackQHelpSearchResultWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "leaveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MinimumSizeHint
func callbackQHelpSearchResultWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchResultWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpSearchResultWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "minimumSizeHint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "minimumSizeHint")
	}
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QHelpSearchResultWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpSearchResultWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_MoveEvent
func callbackQHelpSearchResultWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "moveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "moveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_PaintEvent
func callbackQHelpSearchResultWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::paintEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "paintEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "paintEvent")
	}
}

func (ptr *QHelpSearchResultWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SetEnabled
func callbackQHelpSearchResultWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setEnabled", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setEnabled")
	}
}

func (ptr *QHelpSearchResultWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchResultWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetStyleSheet
func callbackQHelpSearchResultWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setStyleSheet", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setStyleSheet")
	}
}

func (ptr *QHelpSearchResultWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QHelpSearchResultWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchResultWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpSearchResultWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QHelpSearchResultWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchResultWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpSearchResultWidget_SetVisible
func callbackQHelpSearchResultWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setVisible")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setVisible", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setVisible")
	}
}

func (ptr *QHelpSearchResultWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowModified
func callbackQHelpSearchResultWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setWindowModified", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setWindowModified")
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowTitle
func callbackQHelpSearchResultWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setWindowTitle", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setWindowTitle")
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QHelpSearchResultWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchResultWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QHelpSearchResultWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchResultWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpSearchResultWidget_ShowEvent
func callbackQHelpSearchResultWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SizeHint
func callbackQHelpSearchResultWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchResultWidget::sizeHint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QHelpSearchResultWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "sizeHint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "sizeHint")
	}
}

func (ptr *QHelpSearchResultWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QHelpSearchResultWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QHelpSearchResultWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_ChangeEvent
func callbackQHelpSearchResultWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "changeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "changeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Close
func callbackQHelpSearchResultWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchResultWidget::close")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchResultWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QHelpSearchResultWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "close", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "close")
	}
}

func (ptr *QHelpSearchResultWidget) Close() bool {
	defer qt.Recovering("QHelpSearchResultWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) CloseDefault() bool {
	defer qt.Recovering("QHelpSearchResultWidget::close")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_CloseEvent
func callbackQHelpSearchResultWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "closeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "closeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ContextMenuEvent
func callbackQHelpSearchResultWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "contextMenuEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusNextPrevChild
func callbackQHelpSearchResultWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchResultWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusNextPrevChild", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "focusNextPrevChild")
	}
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QHelpSearchResultWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QHelpSearchResultWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_HasHeightForWidth
func callbackQHelpSearchResultWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchResultWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchResultWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QHelpSearchResultWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hasHeightForWidth", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hasHeightForWidth")
	}
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QHelpSearchResultWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QHelpSearchResultWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_HeightForWidth
func callbackQHelpSearchResultWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QHelpSearchResultWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchResultWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QHelpSearchResultWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "heightForWidth", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "heightForWidth")
	}
}

func (ptr *QHelpSearchResultWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QHelpSearchResultWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpSearchResultWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QHelpSearchResultWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_Hide
func callbackQHelpSearchResultWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::hide")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hide", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "hide")
	}
}

func (ptr *QHelpSearchResultWidget) Hide() {
	defer qt.Recovering("QHelpSearchResultWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) HideDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::hide")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_InputMethodEvent
func callbackQHelpSearchResultWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "inputMethodEvent")
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_InputMethodQuery
func callbackQHelpSearchResultWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchResultWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QHelpSearchResultWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "inputMethodQuery", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "inputMethodQuery")
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_KeyPressEvent
func callbackQHelpSearchResultWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "keyPressEvent")
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_KeyReleaseEvent
func callbackQHelpSearchResultWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "keyReleaseEvent")
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Lower
func callbackQHelpSearchResultWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::lower")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "lower", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "lower")
	}
}

func (ptr *QHelpSearchResultWidget) Lower() {
	defer qt.Recovering("QHelpSearchResultWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) LowerDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::lower")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_MouseDoubleClickEvent
func callbackQHelpSearchResultWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseDoubleClickEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseMoveEvent
func callbackQHelpSearchResultWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MousePressEvent
func callbackQHelpSearchResultWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseReleaseEvent
func callbackQHelpSearchResultWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_NativeEvent
func callbackQHelpSearchResultWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QHelpSearchResultWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QHelpSearchResultWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QHelpSearchResultWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "nativeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "nativeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpSearchResultWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpSearchResultWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QHelpSearchResultWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QHelpSearchResultWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_Raise
func callbackQHelpSearchResultWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::raise")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "raise", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "raise")
	}
}

func (ptr *QHelpSearchResultWidget) Raise() {
	defer qt.Recovering("QHelpSearchResultWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) RaiseDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::raise")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_Repaint
func callbackQHelpSearchResultWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "repaint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "repaint")
	}
}

func (ptr *QHelpSearchResultWidget) Repaint() {
	defer qt.Recovering("QHelpSearchResultWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) RepaintDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::repaint")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ResizeEvent
func callbackQHelpSearchResultWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "resizeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SetDisabled
func callbackQHelpSearchResultWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setDisabled", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setDisabled")
	}
}

func (ptr *QHelpSearchResultWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpSearchResultWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchResultWidget_SetFocus2
func callbackQHelpSearchResultWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setFocus2", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setFocus2")
	}
}

func (ptr *QHelpSearchResultWidget) SetFocus2() {
	defer qt.Recovering("QHelpSearchResultWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) SetFocus2Default() {
	defer qt.Recovering("QHelpSearchResultWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_SetHidden
func callbackQHelpSearchResultWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QHelpSearchResultWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setHidden", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "setHidden")
	}
}

func (ptr *QHelpSearchResultWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpSearchResultWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchResultWidget_Show
func callbackQHelpSearchResultWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::show")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "show", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "show")
	}
}

func (ptr *QHelpSearchResultWidget) Show() {
	defer qt.Recovering("QHelpSearchResultWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::show")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowFullScreen
func callbackQHelpSearchResultWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showFullScreen", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showFullScreen")
	}
}

func (ptr *QHelpSearchResultWidget) ShowFullScreen() {
	defer qt.Recovering("QHelpSearchResultWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMaximized
func callbackQHelpSearchResultWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showMaximized", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showMaximized")
	}
}

func (ptr *QHelpSearchResultWidget) ShowMaximized() {
	defer qt.Recovering("QHelpSearchResultWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMinimized
func callbackQHelpSearchResultWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showMinimized", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showMinimized")
	}
}

func (ptr *QHelpSearchResultWidget) ShowMinimized() {
	defer qt.Recovering("QHelpSearchResultWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowNormal
func callbackQHelpSearchResultWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showNormal", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "showNormal")
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormal() {
	defer qt.Recovering("QHelpSearchResultWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormalDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_TabletEvent
func callbackQHelpSearchResultWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "tabletEvent")
	}
}

func (ptr *QHelpSearchResultWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Update
func callbackQHelpSearchResultWidget_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::update")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "update", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "update")
	}
}

func (ptr *QHelpSearchResultWidget) Update() {
	defer qt.Recovering("QHelpSearchResultWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Update(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) UpdateDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::update")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_UpdateMicroFocus
func callbackQHelpSearchResultWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "updateMicroFocus", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "updateMicroFocus")
	}
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocus() {
	defer qt.Recovering("QHelpSearchResultWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_WheelEvent
func callbackQHelpSearchResultWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *QHelpSearchResultWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_TimerEvent
func callbackQHelpSearchResultWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QHelpSearchResultWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ChildEvent
func callbackQHelpSearchResultWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ConnectNotify
func callbackQHelpSearchResultWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QHelpSearchResultWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_CustomEvent
func callbackQHelpSearchResultWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QHelpSearchResultWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DeleteLater
func callbackQHelpSearchResultWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHelpSearchResultWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLater() {
	defer qt.Recovering("QHelpSearchResultWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()))
		C.QHelpSearchResultWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLaterDefault() {
	defer qt.Recovering("QHelpSearchResultWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()))
		C.QHelpSearchResultWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchResultWidget_DisconnectNotify
func callbackQHelpSearchResultWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_EventFilter
func callbackQHelpSearchResultWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHelpSearchResultWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHelpSearchResultWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QHelpSearchResultWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchResultWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHelpSearchResultWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_MetaObject
func callbackQHelpSearchResultWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHelpSearchResultWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchResultWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHelpSearchResultWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QHelpSearchResultWidget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QHelpSearchResultWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchResultWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHelpSearchResultWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
