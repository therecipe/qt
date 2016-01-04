package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QFileSystemModel struct {
	core.QAbstractItemModel
}

type QFileSystemModel_ITF interface {
	core.QAbstractItemModel_ITF
	QFileSystemModel_PTR() *QFileSystemModel
}

func PointerFromQFileSystemModel(ptr QFileSystemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemModel_PTR().Pointer()
	}
	return nil
}

func NewQFileSystemModelFromPointer(ptr unsafe.Pointer) *QFileSystemModel {
	var n = new(QFileSystemModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileSystemModel_") {
		n.SetObjectName("QFileSystemModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QFileSystemModel) QFileSystemModel_PTR() *QFileSystemModel {
	return ptr
}

//QFileSystemModel::Roles
type QFileSystemModel__Roles int64

var (
	QFileSystemModel__FileIconRole    = QFileSystemModel__Roles(core.Qt__DecorationRole)
	QFileSystemModel__FilePathRole    = QFileSystemModel__Roles(C.QFileSystemModel_FilePathRole_Type())
	QFileSystemModel__FileNameRole    = QFileSystemModel__Roles(C.QFileSystemModel_FileNameRole_Type())
	QFileSystemModel__FilePermissions = QFileSystemModel__Roles(C.QFileSystemModel_FilePermissions_Type())
)

func (ptr *QFileSystemModel) IsReadOnly() bool {
	defer qt.Recovering("QFileSystemModel::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) NameFilterDisables() bool {
	defer qt.Recovering("QFileSystemModel::nameFilterDisables")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_NameFilterDisables(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ResolveSymlinks() bool {
	defer qt.Recovering("QFileSystemModel::resolveSymlinks")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_ResolveSymlinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) Rmdir(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QFileSystemModel::rmdir")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_Rmdir(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetNameFilterDisables(enable bool) {
	defer qt.Recovering("QFileSystemModel::setNameFilterDisables")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilterDisables(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetReadOnly(enable bool) {
	defer qt.Recovering("QFileSystemModel::setReadOnly")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetResolveSymlinks(enable bool) {
	defer qt.Recovering("QFileSystemModel::setResolveSymlinks")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetResolveSymlinks(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQFileSystemModel(parent core.QObject_ITF) *QFileSystemModel {
	defer qt.Recovering("QFileSystemModel::QFileSystemModel")

	return NewQFileSystemModelFromPointer(C.QFileSystemModel_NewQFileSystemModel(core.PointerFromQObject(parent)))
}

func (ptr *QFileSystemModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QFileSystemModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QFileSystemModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QFileSystemModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QFileSystemModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) ConnectDirectoryLoaded(f func(path string)) {
	defer qt.Recovering("connect QFileSystemModel::directoryLoaded")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectDirectoryLoaded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryLoaded", f)
	}
}

func (ptr *QFileSystemModel) DisconnectDirectoryLoaded() {
	defer qt.Recovering("disconnect QFileSystemModel::directoryLoaded")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectDirectoryLoaded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryLoaded")
	}
}

//export callbackQFileSystemModelDirectoryLoaded
func callbackQFileSystemModelDirectoryLoaded(ptr unsafe.Pointer, ptrName *C.char, path *C.char) {
	defer qt.Recovering("callback QFileSystemModel::directoryLoaded")

	if signal := qt.GetSignal(C.GoString(ptrName), "directoryLoaded"); signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QFileSystemModel) DirectoryLoaded(path string) {
	defer qt.Recovering("QFileSystemModel::directoryLoaded")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DirectoryLoaded(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QFileSystemModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QFileSystemModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QFileSystemModel::event")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QFileSystemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QFileSystemModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QFileSystemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQFileSystemModelFetchMore
func callbackQFileSystemModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQFileSystemModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QFileSystemModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QFileSystemModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QFileSystemModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QFileSystemModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QFileSystemModel) FileIcon(index core.QModelIndex_ITF) *gui.QIcon {
	defer qt.Recovering("QFileSystemModel::fileIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QFileSystemModel_FileIcon(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QFileSystemModel) FileName(index core.QModelIndex_ITF) string {
	defer qt.Recovering("QFileSystemModel::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FileName(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) FilePath(index core.QModelIndex_ITF) string {
	defer qt.Recovering("QFileSystemModel::filePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FilePath(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectFileRenamed(f func(path string, oldName string, newName string)) {
	defer qt.Recovering("connect QFileSystemModel::fileRenamed")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectFileRenamed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileRenamed", f)
	}
}

func (ptr *QFileSystemModel) DisconnectFileRenamed() {
	defer qt.Recovering("disconnect QFileSystemModel::fileRenamed")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectFileRenamed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileRenamed")
	}
}

//export callbackQFileSystemModelFileRenamed
func callbackQFileSystemModelFileRenamed(ptr unsafe.Pointer, ptrName *C.char, path *C.char, oldName *C.char, newName *C.char) {
	defer qt.Recovering("callback QFileSystemModel::fileRenamed")

	if signal := qt.GetSignal(C.GoString(ptrName), "fileRenamed"); signal != nil {
		signal.(func(string, string, string))(C.GoString(path), C.GoString(oldName), C.GoString(newName))
	}

}

func (ptr *QFileSystemModel) FileRenamed(path string, oldName string, newName string) {
	defer qt.Recovering("QFileSystemModel::fileRenamed")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_FileRenamed(ptr.Pointer(), C.CString(path), C.CString(oldName), C.CString(newName))
	}
}

func (ptr *QFileSystemModel) Filter() core.QDir__Filter {
	defer qt.Recovering("QFileSystemModel::filter")

	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileSystemModel_Filter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileSystemModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QFileSystemModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QFileSystemModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QFileSystemModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QFileSystemModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QFileSystemModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) IconProvider() *QFileIconProvider {
	defer qt.Recovering("QFileSystemModel::iconProvider")

	if ptr.Pointer() != nil {
		return NewQFileIconProviderFromPointer(C.QFileSystemModel_IconProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileSystemModel) Index2(path string, column int) *core.QModelIndex {
	defer qt.Recovering("QFileSystemModel::index")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Index2(ptr.Pointer(), C.CString(path), C.int(column)))
	}
	return nil
}

func (ptr *QFileSystemModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QFileSystemModel::index")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QFileSystemModel) IsDir(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QFileSystemModel::isDir")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsDir(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) LastModified(index core.QModelIndex_ITF) *core.QDateTime {
	defer qt.Recovering("QFileSystemModel::lastModified")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QFileSystemModel_LastModified(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QFileSystemModel) MimeTypes() []string {
	defer qt.Recovering("QFileSystemModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Mkdir(parent core.QModelIndex_ITF, name string) *core.QModelIndex {
	defer qt.Recovering("QFileSystemModel::mkdir")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Mkdir(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QFileSystemModel) MyComputer(role int) *core.QVariant {
	defer qt.Recovering("QFileSystemModel::myComputer")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_MyComputer(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) NameFilters() []string {
	defer qt.Recovering("QFileSystemModel::nameFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_NameFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QFileSystemModel::parent")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QFileSystemModel) RootDirectory() *core.QDir {
	defer qt.Recovering("QFileSystemModel::rootDirectory")

	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QFileSystemModel_RootDirectory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileSystemModel) RootPath() string {
	defer qt.Recovering("QFileSystemModel::rootPath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_RootPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectRootPathChanged(f func(newPath string)) {
	defer qt.Recovering("connect QFileSystemModel::rootPathChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectRootPathChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rootPathChanged", f)
	}
}

func (ptr *QFileSystemModel) DisconnectRootPathChanged() {
	defer qt.Recovering("disconnect QFileSystemModel::rootPathChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectRootPathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rootPathChanged")
	}
}

//export callbackQFileSystemModelRootPathChanged
func callbackQFileSystemModelRootPathChanged(ptr unsafe.Pointer, ptrName *C.char, newPath *C.char) {
	defer qt.Recovering("callback QFileSystemModel::rootPathChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "rootPathChanged"); signal != nil {
		signal.(func(string))(C.GoString(newPath))
	}

}

func (ptr *QFileSystemModel) RootPathChanged(newPath string) {
	defer qt.Recovering("QFileSystemModel::rootPathChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_RootPathChanged(ptr.Pointer(), C.CString(newPath))
	}
}

func (ptr *QFileSystemModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QFileSystemModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QFileSystemModel) SetData(idx core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QFileSystemModel::setData")

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(idx), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetFilter(filters core.QDir__Filter) {
	defer qt.Recovering("QFileSystemModel::setFilter")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFileSystemModel) SetIconProvider(provider QFileIconProvider_ITF) {
	defer qt.Recovering("QFileSystemModel::setIconProvider")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetIconProvider(ptr.Pointer(), PointerFromQFileIconProvider(provider))
	}
}

func (ptr *QFileSystemModel) SetNameFilters(filters []string) {
	defer qt.Recovering("QFileSystemModel::setNameFilters")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileSystemModel) SetRootPath(newPath string) *core.QModelIndex {
	defer qt.Recovering("QFileSystemModel::setRootPath")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_SetRootPath(ptr.Pointer(), C.CString(newPath)))
	}
	return nil
}

func (ptr *QFileSystemModel) Size(index core.QModelIndex_ITF) int64 {
	defer qt.Recovering("QFileSystemModel::size")

	if ptr.Pointer() != nil {
		return int64(C.QFileSystemModel_Size(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QFileSystemModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QFileSystemModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QFileSystemModel) DisconnectSort() {
	defer qt.Recovering("disconnect QFileSystemModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQFileSystemModelSort
func callbackQFileSystemModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QFileSystemModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQFileSystemModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QFileSystemModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QFileSystemModel::sort")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QFileSystemModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QFileSystemModel::sort")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QFileSystemModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QFileSystemModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QFileSystemModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileSystemModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFileSystemModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFileSystemModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFileSystemModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFileSystemModelTimerEvent
func callbackQFileSystemModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQFileSystemModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFileSystemModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSystemModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSystemModel) Type(index core.QModelIndex_ITF) string {
	defer qt.Recovering("QFileSystemModel::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_Type(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) DestroyQFileSystemModel() {
	defer qt.Recovering("QFileSystemModel::~QFileSystemModel")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DestroyQFileSystemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileSystemModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QFileSystemModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QFileSystemModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QFileSystemModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQFileSystemModelRevert
func callbackQFileSystemModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QFileSystemModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFileSystemModel) Revert() {
	defer qt.Recovering("QFileSystemModel::revert")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_Revert(ptr.Pointer())
	}
}

func (ptr *QFileSystemModel) RevertDefault() {
	defer qt.Recovering("QFileSystemModel::revert")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QFileSystemModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFileSystemModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFileSystemModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFileSystemModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFileSystemModelChildEvent
func callbackQFileSystemModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQFileSystemModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QFileSystemModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSystemModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSystemModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFileSystemModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFileSystemModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFileSystemModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFileSystemModelCustomEvent
func callbackQFileSystemModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFileSystemModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFileSystemModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFileSystemModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFileSystemModel::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
