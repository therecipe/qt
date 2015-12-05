package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QStandardPaths struct {
	ptr unsafe.Pointer
}

type QStandardPaths_ITF interface {
	QStandardPaths_PTR() *QStandardPaths
}

func (p *QStandardPaths) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStandardPaths) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStandardPaths(ptr QStandardPaths_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardPaths_PTR().Pointer()
	}
	return nil
}

func NewQStandardPathsFromPointer(ptr unsafe.Pointer) *QStandardPaths {
	var n = new(QStandardPaths)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStandardPaths) QStandardPaths_PTR() *QStandardPaths {
	return ptr
}

//QStandardPaths::LocateOption
type QStandardPaths__LocateOption int64

const (
	QStandardPaths__LocateFile      = QStandardPaths__LocateOption(0x0)
	QStandardPaths__LocateDirectory = QStandardPaths__LocateOption(0x1)
)

//QStandardPaths::StandardLocation
type QStandardPaths__StandardLocation int64

const (
	QStandardPaths__DesktopLocation       = QStandardPaths__StandardLocation(0)
	QStandardPaths__DocumentsLocation     = QStandardPaths__StandardLocation(1)
	QStandardPaths__FontsLocation         = QStandardPaths__StandardLocation(2)
	QStandardPaths__ApplicationsLocation  = QStandardPaths__StandardLocation(3)
	QStandardPaths__MusicLocation         = QStandardPaths__StandardLocation(4)
	QStandardPaths__MoviesLocation        = QStandardPaths__StandardLocation(5)
	QStandardPaths__PicturesLocation      = QStandardPaths__StandardLocation(6)
	QStandardPaths__TempLocation          = QStandardPaths__StandardLocation(7)
	QStandardPaths__HomeLocation          = QStandardPaths__StandardLocation(8)
	QStandardPaths__DataLocation          = QStandardPaths__StandardLocation(9)
	QStandardPaths__CacheLocation         = QStandardPaths__StandardLocation(10)
	QStandardPaths__GenericDataLocation   = QStandardPaths__StandardLocation(11)
	QStandardPaths__RuntimeLocation       = QStandardPaths__StandardLocation(12)
	QStandardPaths__ConfigLocation        = QStandardPaths__StandardLocation(13)
	QStandardPaths__DownloadLocation      = QStandardPaths__StandardLocation(14)
	QStandardPaths__GenericCacheLocation  = QStandardPaths__StandardLocation(15)
	QStandardPaths__GenericConfigLocation = QStandardPaths__StandardLocation(16)
	QStandardPaths__AppDataLocation       = QStandardPaths__StandardLocation(17)
	QStandardPaths__AppConfigLocation     = QStandardPaths__StandardLocation(18)
	QStandardPaths__AppLocalDataLocation  = QStandardPaths__StandardLocation(QStandardPaths__DataLocation)
)

func QStandardPaths_SetTestModeEnabled(testMode bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::setTestModeEnabled")
		}
	}()

	C.QStandardPaths_QStandardPaths_SetTestModeEnabled(C.int(qt.GoBoolToInt(testMode)))
}

func QStandardPaths_FindExecutable(executableName string, paths []string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::findExecutable")
		}
	}()

	return C.GoString(C.QStandardPaths_QStandardPaths_FindExecutable(C.CString(executableName), C.CString(strings.Join(paths, ",,,"))))
}

func QStandardPaths_Locate(ty QStandardPaths__StandardLocation, fileName string, options QStandardPaths__LocateOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::locate")
		}
	}()

	return C.GoString(C.QStandardPaths_QStandardPaths_Locate(C.int(ty), C.CString(fileName), C.int(options)))
}

func QStandardPaths_LocateAll(ty QStandardPaths__StandardLocation, fileName string, options QStandardPaths__LocateOption) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::locateAll")
		}
	}()

	return strings.Split(C.GoString(C.QStandardPaths_QStandardPaths_LocateAll(C.int(ty), C.CString(fileName), C.int(options))), ",,,")
}

func QStandardPaths_DisplayName(ty QStandardPaths__StandardLocation) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::displayName")
		}
	}()

	return C.GoString(C.QStandardPaths_QStandardPaths_DisplayName(C.int(ty)))
}

func QStandardPaths_StandardLocations(ty QStandardPaths__StandardLocation) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::standardLocations")
		}
	}()

	return strings.Split(C.GoString(C.QStandardPaths_QStandardPaths_StandardLocations(C.int(ty))), ",,,")
}

func QStandardPaths_WritableLocation(ty QStandardPaths__StandardLocation) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardPaths::writableLocation")
		}
	}()

	return C.GoString(C.QStandardPaths_QStandardPaths_WritableLocation(C.int(ty)))
}
