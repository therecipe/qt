package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QDir struct {
	ptr unsafe.Pointer
}

type QDir_ITF interface {
	QDir_PTR() *QDir
}

func (p *QDir) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDir) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDir(ptr QDir_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDir_PTR().Pointer()
	}
	return nil
}

func NewQDirFromPointer(ptr unsafe.Pointer) *QDir {
	var n = new(QDir)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDir) QDir_PTR() *QDir {
	return ptr
}

//QDir::Filter
type QDir__Filter int64

const (
	QDir__Dirs           = QDir__Filter(0x001)
	QDir__Files          = QDir__Filter(0x002)
	QDir__Drives         = QDir__Filter(0x004)
	QDir__NoSymLinks     = QDir__Filter(0x008)
	QDir__AllEntries     = QDir__Filter(QDir__Dirs | QDir__Files | QDir__Drives)
	QDir__TypeMask       = QDir__Filter(0x00f)
	QDir__Readable       = QDir__Filter(0x010)
	QDir__Writable       = QDir__Filter(0x020)
	QDir__Executable     = QDir__Filter(0x040)
	QDir__PermissionMask = QDir__Filter(0x070)
	QDir__Modified       = QDir__Filter(0x080)
	QDir__Hidden         = QDir__Filter(0x100)
	QDir__System         = QDir__Filter(0x200)
	QDir__AccessMask     = QDir__Filter(0x3F0)
	QDir__AllDirs        = QDir__Filter(0x400)
	QDir__CaseSensitive  = QDir__Filter(0x800)
	QDir__NoDot          = QDir__Filter(0x2000)
	QDir__NoDotDot       = QDir__Filter(0x4000)
	QDir__NoDotAndDotDot = QDir__Filter(QDir__NoDot | QDir__NoDotDot)
	QDir__NoFilter       = QDir__Filter(-1)
)

//QDir::SortFlag
type QDir__SortFlag int64

const (
	QDir__Name        = QDir__SortFlag(0x00)
	QDir__Time        = QDir__SortFlag(0x01)
	QDir__Size        = QDir__SortFlag(0x02)
	QDir__Unsorted    = QDir__SortFlag(0x03)
	QDir__SortByMask  = QDir__SortFlag(0x03)
	QDir__DirsFirst   = QDir__SortFlag(0x04)
	QDir__Reversed    = QDir__SortFlag(0x08)
	QDir__IgnoreCase  = QDir__SortFlag(0x10)
	QDir__DirsLast    = QDir__SortFlag(0x20)
	QDir__LocaleAware = QDir__SortFlag(0x40)
	QDir__Type        = QDir__SortFlag(0x80)
	QDir__NoSort      = QDir__SortFlag(-1)
)

func NewQDir(dir QDir_ITF) *QDir {
	defer qt.Recovering("QDir::QDir")

	return NewQDirFromPointer(C.QDir_NewQDir(PointerFromQDir(dir)))
}

func NewQDir2(path string) *QDir {
	defer qt.Recovering("QDir::QDir")

	return NewQDirFromPointer(C.QDir_NewQDir2(C.CString(path)))
}

func NewQDir3(path string, nameFilter string, sort QDir__SortFlag, filters QDir__Filter) *QDir {
	defer qt.Recovering("QDir::QDir")

	return NewQDirFromPointer(C.QDir_NewQDir3(C.CString(path), C.CString(nameFilter), C.int(sort), C.int(filters)))
}

func (ptr *QDir) AbsoluteFilePath(fileName string) string {
	defer qt.Recovering("QDir::absoluteFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_AbsoluteFilePath(ptr.Pointer(), C.CString(fileName)))
	}
	return ""
}

func (ptr *QDir) AbsolutePath() string {
	defer qt.Recovering("QDir::absolutePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_AbsolutePath(ptr.Pointer()))
	}
	return ""
}

func QDir_AddSearchPath(prefix string, path string) {
	defer qt.Recovering("QDir::addSearchPath")

	C.QDir_QDir_AddSearchPath(C.CString(prefix), C.CString(path))
}

func (ptr *QDir) CanonicalPath() string {
	defer qt.Recovering("QDir::canonicalPath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_CanonicalPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDir) Cd(dirName string) bool {
	defer qt.Recovering("QDir::cd")

	if ptr.Pointer() != nil {
		return C.QDir_Cd(ptr.Pointer(), C.CString(dirName)) != 0
	}
	return false
}

func (ptr *QDir) CdUp() bool {
	defer qt.Recovering("QDir::cdUp")

	if ptr.Pointer() != nil {
		return C.QDir_CdUp(ptr.Pointer()) != 0
	}
	return false
}

func QDir_CleanPath(path string) string {
	defer qt.Recovering("QDir::cleanPath")

	return C.GoString(C.QDir_QDir_CleanPath(C.CString(path)))
}

func QDir_Current() *QDir {
	defer qt.Recovering("QDir::current")

	return NewQDirFromPointer(C.QDir_QDir_Current())
}

func QDir_CurrentPath() string {
	defer qt.Recovering("QDir::currentPath")

	return C.GoString(C.QDir_QDir_CurrentPath())
}

func (ptr *QDir) DirName() string {
	defer qt.Recovering("QDir::dirName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_DirName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDir) EntryList2(filters QDir__Filter, sort QDir__SortFlag) []string {
	defer qt.Recovering("QDir::entryList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDir_EntryList2(ptr.Pointer(), C.int(filters), C.int(sort))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QDir) EntryList(nameFilters []string, filters QDir__Filter, sort QDir__SortFlag) []string {
	defer qt.Recovering("QDir::entryList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDir_EntryList(ptr.Pointer(), C.CString(strings.Join(nameFilters, ",,,")), C.int(filters), C.int(sort))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QDir) Exists2() bool {
	defer qt.Recovering("QDir::exists")

	if ptr.Pointer() != nil {
		return C.QDir_Exists2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDir) Exists(name string) bool {
	defer qt.Recovering("QDir::exists")

	if ptr.Pointer() != nil {
		return C.QDir_Exists(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDir) FilePath(fileName string) string {
	defer qt.Recovering("QDir::filePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_FilePath(ptr.Pointer(), C.CString(fileName)))
	}
	return ""
}

func (ptr *QDir) Filter() QDir__Filter {
	defer qt.Recovering("QDir::filter")

	if ptr.Pointer() != nil {
		return QDir__Filter(C.QDir_Filter(ptr.Pointer()))
	}
	return 0
}

func QDir_FromNativeSeparators(pathName string) string {
	defer qt.Recovering("QDir::fromNativeSeparators")

	return C.GoString(C.QDir_QDir_FromNativeSeparators(C.CString(pathName)))
}

func QDir_Home() *QDir {
	defer qt.Recovering("QDir::home")

	return NewQDirFromPointer(C.QDir_QDir_Home())
}

func QDir_HomePath() string {
	defer qt.Recovering("QDir::homePath")

	return C.GoString(C.QDir_QDir_HomePath())
}

func (ptr *QDir) IsAbsolute() bool {
	defer qt.Recovering("QDir::isAbsolute")

	if ptr.Pointer() != nil {
		return C.QDir_IsAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func QDir_IsAbsolutePath(path string) bool {
	defer qt.Recovering("QDir::isAbsolutePath")

	return C.QDir_QDir_IsAbsolutePath(C.CString(path)) != 0
}

func (ptr *QDir) IsReadable() bool {
	defer qt.Recovering("QDir::isReadable")

	if ptr.Pointer() != nil {
		return C.QDir_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDir) IsRelative() bool {
	defer qt.Recovering("QDir::isRelative")

	if ptr.Pointer() != nil {
		return C.QDir_IsRelative(ptr.Pointer()) != 0
	}
	return false
}

func QDir_IsRelativePath(path string) bool {
	defer qt.Recovering("QDir::isRelativePath")

	return C.QDir_QDir_IsRelativePath(C.CString(path)) != 0
}

func (ptr *QDir) IsRoot() bool {
	defer qt.Recovering("QDir::isRoot")

	if ptr.Pointer() != nil {
		return C.QDir_IsRoot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDir) MakeAbsolute() bool {
	defer qt.Recovering("QDir::makeAbsolute")

	if ptr.Pointer() != nil {
		return C.QDir_MakeAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func QDir_Match(filter string, fileName string) bool {
	defer qt.Recovering("QDir::match")

	return C.QDir_QDir_Match(C.CString(filter), C.CString(fileName)) != 0
}

func QDir_Match2(filters []string, fileName string) bool {
	defer qt.Recovering("QDir::match")

	return C.QDir_QDir_Match2(C.CString(strings.Join(filters, ",,,")), C.CString(fileName)) != 0
}

func (ptr *QDir) Mkdir(dirName string) bool {
	defer qt.Recovering("QDir::mkdir")

	if ptr.Pointer() != nil {
		return C.QDir_Mkdir(ptr.Pointer(), C.CString(dirName)) != 0
	}
	return false
}

func (ptr *QDir) Mkpath(dirPath string) bool {
	defer qt.Recovering("QDir::mkpath")

	if ptr.Pointer() != nil {
		return C.QDir_Mkpath(ptr.Pointer(), C.CString(dirPath)) != 0
	}
	return false
}

func (ptr *QDir) NameFilters() []string {
	defer qt.Recovering("QDir::nameFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDir_NameFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QDir) Path() string {
	defer qt.Recovering("QDir::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDir) Refresh() {
	defer qt.Recovering("QDir::refresh")

	if ptr.Pointer() != nil {
		C.QDir_Refresh(ptr.Pointer())
	}
}

func (ptr *QDir) RelativeFilePath(fileName string) string {
	defer qt.Recovering("QDir::relativeFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDir_RelativeFilePath(ptr.Pointer(), C.CString(fileName)))
	}
	return ""
}

func (ptr *QDir) RemoveRecursively() bool {
	defer qt.Recovering("QDir::removeRecursively")

	if ptr.Pointer() != nil {
		return C.QDir_RemoveRecursively(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDir) Rename(oldName string, newName string) bool {
	defer qt.Recovering("QDir::rename")

	if ptr.Pointer() != nil {
		return C.QDir_Rename(ptr.Pointer(), C.CString(oldName), C.CString(newName)) != 0
	}
	return false
}

func (ptr *QDir) Rmdir(dirName string) bool {
	defer qt.Recovering("QDir::rmdir")

	if ptr.Pointer() != nil {
		return C.QDir_Rmdir(ptr.Pointer(), C.CString(dirName)) != 0
	}
	return false
}

func (ptr *QDir) Rmpath(dirPath string) bool {
	defer qt.Recovering("QDir::rmpath")

	if ptr.Pointer() != nil {
		return C.QDir_Rmpath(ptr.Pointer(), C.CString(dirPath)) != 0
	}
	return false
}

func QDir_Root() *QDir {
	defer qt.Recovering("QDir::root")

	return NewQDirFromPointer(C.QDir_QDir_Root())
}

func QDir_RootPath() string {
	defer qt.Recovering("QDir::rootPath")

	return C.GoString(C.QDir_QDir_RootPath())
}

func QDir_SearchPaths(prefix string) []string {
	defer qt.Recovering("QDir::searchPaths")

	return strings.Split(C.GoString(C.QDir_QDir_SearchPaths(C.CString(prefix))), ",,,")
}

func QDir_SetCurrent(path string) bool {
	defer qt.Recovering("QDir::setCurrent")

	return C.QDir_QDir_SetCurrent(C.CString(path)) != 0
}

func (ptr *QDir) SetFilter(filters QDir__Filter) {
	defer qt.Recovering("QDir::setFilter")

	if ptr.Pointer() != nil {
		C.QDir_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QDir) SetNameFilters(nameFilters []string) {
	defer qt.Recovering("QDir::setNameFilters")

	if ptr.Pointer() != nil {
		C.QDir_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(nameFilters, ",,,")))
	}
}

func (ptr *QDir) SetPath(path string) {
	defer qt.Recovering("QDir::setPath")

	if ptr.Pointer() != nil {
		C.QDir_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func QDir_SetSearchPaths(prefix string, searchPaths []string) {
	defer qt.Recovering("QDir::setSearchPaths")

	C.QDir_QDir_SetSearchPaths(C.CString(prefix), C.CString(strings.Join(searchPaths, ",,,")))
}

func (ptr *QDir) SetSorting(sort QDir__SortFlag) {
	defer qt.Recovering("QDir::setSorting")

	if ptr.Pointer() != nil {
		C.QDir_SetSorting(ptr.Pointer(), C.int(sort))
	}
}

func (ptr *QDir) Sorting() QDir__SortFlag {
	defer qt.Recovering("QDir::sorting")

	if ptr.Pointer() != nil {
		return QDir__SortFlag(C.QDir_Sorting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDir) Swap(other QDir_ITF) {
	defer qt.Recovering("QDir::swap")

	if ptr.Pointer() != nil {
		C.QDir_Swap(ptr.Pointer(), PointerFromQDir(other))
	}
}

func QDir_Temp() *QDir {
	defer qt.Recovering("QDir::temp")

	return NewQDirFromPointer(C.QDir_QDir_Temp())
}

func QDir_TempPath() string {
	defer qt.Recovering("QDir::tempPath")

	return C.GoString(C.QDir_QDir_TempPath())
}

func QDir_ToNativeSeparators(pathName string) string {
	defer qt.Recovering("QDir::toNativeSeparators")

	return C.GoString(C.QDir_QDir_ToNativeSeparators(C.CString(pathName)))
}

func (ptr *QDir) DestroyQDir() {
	defer qt.Recovering("QDir::~QDir")

	if ptr.Pointer() != nil {
		C.QDir_DestroyQDir(ptr.Pointer())
	}
}
