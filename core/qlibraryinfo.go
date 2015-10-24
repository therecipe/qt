package core

//#include "qlibraryinfo.h"
import "C"
import (
	"unsafe"
)

type QLibraryInfo struct {
	ptr unsafe.Pointer
}

type QLibraryInfoITF interface {
	QLibraryInfoPTR() *QLibraryInfo
}

func (p *QLibraryInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLibraryInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLibraryInfo(ptr QLibraryInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLibraryInfoPTR().Pointer()
	}
	return nil
}

func QLibraryInfoFromPointer(ptr unsafe.Pointer) *QLibraryInfo {
	var n = new(QLibraryInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLibraryInfo) QLibraryInfoPTR() *QLibraryInfo {
	return ptr
}

//QLibraryInfo::LibraryLocation
type QLibraryInfo__LibraryLocation int

var (
	QLibraryInfo__PrefixPath             = QLibraryInfo__LibraryLocation(0)
	QLibraryInfo__DocumentationPath      = QLibraryInfo__LibraryLocation(1)
	QLibraryInfo__HeadersPath            = QLibraryInfo__LibraryLocation(2)
	QLibraryInfo__LibrariesPath          = QLibraryInfo__LibraryLocation(3)
	QLibraryInfo__LibraryExecutablesPath = QLibraryInfo__LibraryLocation(4)
	QLibraryInfo__BinariesPath           = QLibraryInfo__LibraryLocation(5)
	QLibraryInfo__PluginsPath            = QLibraryInfo__LibraryLocation(6)
	QLibraryInfo__ImportsPath            = QLibraryInfo__LibraryLocation(7)
	QLibraryInfo__Qml2ImportsPath        = QLibraryInfo__LibraryLocation(8)
	QLibraryInfo__ArchDataPath           = QLibraryInfo__LibraryLocation(9)
	QLibraryInfo__DataPath               = QLibraryInfo__LibraryLocation(10)
	QLibraryInfo__TranslationsPath       = QLibraryInfo__LibraryLocation(11)
	QLibraryInfo__ExamplesPath           = QLibraryInfo__LibraryLocation(12)
	QLibraryInfo__TestsPath              = QLibraryInfo__LibraryLocation(13)
	QLibraryInfo__SettingsPath           = QLibraryInfo__LibraryLocation(100)
)

func QLibraryInfo_IsDebugBuild() bool {
	return C.QLibraryInfo_QLibraryInfo_IsDebugBuild() != 0
}

func QLibraryInfo_LicensedProducts() string {
	return C.GoString(C.QLibraryInfo_QLibraryInfo_LicensedProducts())
}

func QLibraryInfo_Licensee() string {
	return C.GoString(C.QLibraryInfo_QLibraryInfo_Licensee())
}

func QLibraryInfo_Location(loc QLibraryInfo__LibraryLocation) string {
	return C.GoString(C.QLibraryInfo_QLibraryInfo_Location(C.int(loc)))
}
