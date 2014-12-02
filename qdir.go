package qt

//#include "qdir.h"
import "C"

import "strings"

type qdir struct {
	ptr C.QtObjectPtr
}

type QDir interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
}

func (p *qdir) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qdir) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//Filter
type Filter int

var (
	DIRS           = Filter(C.QDir_Dirs())
	ALLDIRS        = Filter(C.QDir_AllDirs())
	FILES          = Filter(C.QDir_Files())
	DRIVES         = Filter(C.QDir_Drives())
	NOSYMLINKS     = Filter(C.QDir_NoSymLinks())
	NODOTANDDOTDOT = Filter(C.QDir_NoDotAndDotDot())
	NODOT          = Filter(C.QDir_NoDot())
	NODOTDOT       = Filter(C.QDir_NoDotDot())
	ALLENTRIES     = Filter(C.QDir_AllEntries())
	READABLE       = Filter(C.QDir_Readable())
	WRITABLE       = Filter(C.QDir_Writable())
	EXECUTABLE     = Filter(C.QDir_Executable())
	MODIFIED       = Filter(C.QDir_Modified())
	HIDDEN         = Filter(C.QDir_Hidden())
	SYSTEM         = Filter(C.QDir_System())
)

//SortFlag
type SortFlag int

var (
	NAME        = SortFlag(C.QDir_Name())
	TIME        = SortFlag(C.QDir_Time())
	SIZE        = SortFlag(C.QDir_Size())
	TYPE        = SortFlag(C.QDir_Type())
	UNSORTED    = SortFlag(C.QDir_Unsorted())
	NOSORT      = SortFlag(C.QDir_NoSort())
	DIRSFIRST   = SortFlag(C.QDir_DirsFirst())
	DIRSLAST    = SortFlag(C.QDir_DirsLast())
	REVERSED    = SortFlag(C.QDir_Reversed())
	IGNORECASE  = SortFlag(C.QDir_IgnoreCase())
	LOCALEAWARE = SortFlag(C.QDir_LocaleAware())
)

func QDir_AddSearchPath_String_String(prefix string, path string) {
	C.QDir_AddSearchPath_String_String(C.CString(prefix), C.CString(path))
}

func QDir_CleanPath_String(path string) string {
	return C.GoString(C.QDir_CleanPath_String(C.CString(path)))
}

func QDir_CurrentPath() string {
	return C.GoString(C.QDir_CurrentPath())
}

func QDir_FromNativeSeparators_String(pathName string) string {
	return C.GoString(C.QDir_FromNativeSeparators_String(C.CString(pathName)))
}

func QDir_HomePath() string {
	return C.GoString(C.QDir_HomePath())
}

func QDir_IsAbsolutePath_String(path string) bool {
	return C.QDir_IsAbsolutePath_String(C.CString(path)) != 0
}

func QDir_IsRelativePath_String(path string) bool {
	return C.QDir_IsRelativePath_String(C.CString(path)) != 0
}

func QDir_Match_String_String(filter string, fileName string) bool {
	return C.QDir_Match_String_String(C.CString(filter), C.CString(fileName)) != 0
}

func QDir_Match_QStringList_String(filters []string, fileName string) bool {
	return C.QDir_Match_QStringList_String(C.CString(strings.Join(filters, "|")),
		C.CString(fileName)) != 0
}

func QDir_RootPath() string {
	return C.GoString(C.QDir_RootPath())
}

func QDir_SearchPaths_String(prefix string) []string {
	return strings.Split(C.GoString(C.QDir_SearchPaths_String(C.CString(prefix))), "|")
}

func QDir_SetCurrent_String(path string) bool {
	return C.QDir_SetCurrent_String(C.CString(path)) != 0
}

func QDir_SetSearchPaths_String_QStringList(prefix string, searchPaths []string) {
	C.QDir_SetSearchPaths_String_QStringList(C.CString(prefix), C.CString(strings.Join(searchPaths, "|")))
}

func QDir_TempPath() string {
	return C.GoString(C.QDir_TempPath())
}

func QDir_ToNativeSeparators_String(pathName string) string {
	return C.GoString(C.QDir_ToNativeSeparators_String(C.CString(pathName)))
}
