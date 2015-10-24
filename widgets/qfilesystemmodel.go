package widgets

//#include "qfilesystemmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QFileSystemModel struct {
	core.QAbstractItemModel
}

type QFileSystemModelITF interface {
	core.QAbstractItemModelITF
	QFileSystemModelPTR() *QFileSystemModel
}

func PointerFromQFileSystemModel(ptr QFileSystemModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemModelPTR().Pointer()
	}
	return nil
}

func QFileSystemModelFromPointer(ptr unsafe.Pointer) *QFileSystemModel {
	var n = new(QFileSystemModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileSystemModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSystemModel) QFileSystemModelPTR() *QFileSystemModel {
	return ptr
}

//QFileSystemModel::Roles
type QFileSystemModel__Roles int

var (
	QFileSystemModel__FileIconRole    = QFileSystemModel__Roles(core.Qt__DecorationRole)
	QFileSystemModel__FilePathRole    = QFileSystemModel__Roles(C.QFileSystemModel_FilePathRole_Type())
	QFileSystemModel__FileNameRole    = QFileSystemModel__Roles(C.QFileSystemModel_FileNameRole_Type())
	QFileSystemModel__FilePermissions = QFileSystemModel__Roles(C.QFileSystemModel_FilePermissions_Type())
)

func (ptr *QFileSystemModel) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileSystemModel) NameFilterDisables() bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_NameFilterDisables(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ResolveSymlinks() bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_ResolveSymlinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileSystemModel) Rmdir(index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_Rmdir(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetNameFilterDisables(enable bool) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilterDisables(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetReadOnly(enable bool) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetResolveSymlinks(enable bool) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetResolveSymlinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQFileSystemModel(parent core.QObjectITF) *QFileSystemModel {
	return QFileSystemModelFromPointer(unsafe.Pointer(C.QFileSystemModel_NewQFileSystemModel(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QFileSystemModel) CanFetchMore(parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_CanFetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ColumnCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QFileSystemModel) Data(index core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectDirectoryLoaded(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectDirectoryLoaded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "directoryLoaded", f)
	}
}

func (ptr *QFileSystemModel) DisconnectDirectoryLoaded() {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectDirectoryLoaded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "directoryLoaded")
	}
}

//export callbackQFileSystemModelDirectoryLoaded
func callbackQFileSystemModelDirectoryLoaded(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryLoaded").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemModel) DropMimeData(data core.QMimeDataITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QFileSystemModel) FetchMore(parent core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_FetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent)))
	}
}

func (ptr *QFileSystemModel) FileName(index core.QModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FileName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return ""
}

func (ptr *QFileSystemModel) FilePath(index core.QModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FilePath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectFileRenamed(f func(path string, oldName string, newName string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectFileRenamed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fileRenamed", f)
	}
}

func (ptr *QFileSystemModel) DisconnectFileRenamed() {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectFileRenamed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fileRenamed")
	}
}

//export callbackQFileSystemModelFileRenamed
func callbackQFileSystemModelFileRenamed(ptrName *C.char, path *C.char, oldName *C.char, newName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fileRenamed").(func(string, string, string))(C.GoString(path), C.GoString(oldName), C.GoString(newName))
}

func (ptr *QFileSystemModel) Filter() core.QDir__Filter {
	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileSystemModel_Filter(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileSystemModel) Flags(index core.QModelIndexITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QFileSystemModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QFileSystemModel) HasChildren(parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_HasChildren(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QFileSystemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QFileSystemModel) IconProvider() *QFileIconProvider {
	if ptr.Pointer() != nil {
		return QFileIconProviderFromPointer(unsafe.Pointer(C.QFileSystemModel_IconProvider(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QFileSystemModel) Index2(path string, column int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QFileSystemModel_Index2(C.QtObjectPtr(ptr.Pointer()), C.CString(path), C.int(column))))
	}
	return nil
}

func (ptr *QFileSystemModel) Index(row int, column int, parent core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QFileSystemModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QFileSystemModel) IsDir(index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsDir(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QFileSystemModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_MimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Mkdir(parent core.QModelIndexITF, name string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QFileSystemModel_Mkdir(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent)), C.CString(name))))
	}
	return nil
}

func (ptr *QFileSystemModel) MyComputer(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_MyComputer(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QFileSystemModel) NameFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_NameFilters(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Parent(index core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QFileSystemModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QFileSystemModel) RootPath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_RootPath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectRootPathChanged(f func(newPath string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectRootPathChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rootPathChanged", f)
	}
}

func (ptr *QFileSystemModel) DisconnectRootPathChanged() {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectRootPathChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rootPathChanged")
	}
}

//export callbackQFileSystemModelRootPathChanged
func callbackQFileSystemModelRootPathChanged(ptrName *C.char, newPath *C.char) {
	qt.GetSignal(C.GoString(ptrName), "rootPathChanged").(func(string))(C.GoString(newPath))
}

func (ptr *QFileSystemModel) RowCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QFileSystemModel) SetData(idx core.QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(idx)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetFilter(filters core.QDir__Filter) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetFilter(C.QtObjectPtr(ptr.Pointer()), C.int(filters))
	}
}

func (ptr *QFileSystemModel) SetIconProvider(provider QFileIconProviderITF) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetIconProvider(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFileIconProvider(provider)))
	}
}

func (ptr *QFileSystemModel) SetNameFilters(filters []string) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilters(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(filters, "|")))
	}
}

func (ptr *QFileSystemModel) SetRootPath(newPath string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QFileSystemModel_SetRootPath(C.QtObjectPtr(ptr.Pointer()), C.CString(newPath))))
	}
	return nil
}

func (ptr *QFileSystemModel) Sort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QFileSystemModel) SupportedDropActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QFileSystemModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileSystemModel) Type(index core.QModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_Type(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return ""
}

func (ptr *QFileSystemModel) DestroyQFileSystemModel() {
	if ptr.Pointer() != nil {
		C.QFileSystemModel_DestroyQFileSystemModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
