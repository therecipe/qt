package core

//#include "qlibrary.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QLibrary struct {
	QObject
}

type QLibrary_ITF interface {
	QObject_ITF
	QLibrary_PTR() *QLibrary
}

func PointerFromQLibrary(ptr QLibrary_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLibrary_PTR().Pointer()
	}
	return nil
}

func NewQLibraryFromPointer(ptr unsafe.Pointer) *QLibrary {
	var n = new(QLibrary)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QLibrary_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLibrary) QLibrary_PTR() *QLibrary {
	return ptr
}

//QLibrary::LoadHint
type QLibrary__LoadHint int64

const (
	QLibrary__ResolveAllSymbolsHint     = QLibrary__LoadHint(0x01)
	QLibrary__ExportExternalSymbolsHint = QLibrary__LoadHint(0x02)
	QLibrary__LoadArchiveMemberHint     = QLibrary__LoadHint(0x04)
	QLibrary__PreventUnloadHint         = QLibrary__LoadHint(0x08)
	QLibrary__DeepBindHint              = QLibrary__LoadHint(0x10)
)

func (ptr *QLibrary) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLibrary_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLibrary) LoadHints() QLibrary__LoadHint {
	if ptr.Pointer() != nil {
		return QLibrary__LoadHint(C.QLibrary_LoadHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLibrary) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QLibrary_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QLibrary) SetFileNameAndVersion(fileName string, versionNumber int) {
	if ptr.Pointer() != nil {
		C.QLibrary_SetFileNameAndVersion(ptr.Pointer(), C.CString(fileName), C.int(versionNumber))
	}
}

func (ptr *QLibrary) SetLoadHints(hints QLibrary__LoadHint) {
	if ptr.Pointer() != nil {
		C.QLibrary_SetLoadHints(ptr.Pointer(), C.int(hints))
	}
}

func NewQLibrary(parent QObject_ITF) *QLibrary {
	return NewQLibraryFromPointer(C.QLibrary_NewQLibrary(PointerFromQObject(parent)))
}

func NewQLibrary2(fileName string, parent QObject_ITF) *QLibrary {
	return NewQLibraryFromPointer(C.QLibrary_NewQLibrary2(C.CString(fileName), PointerFromQObject(parent)))
}

func NewQLibrary4(fileName string, version string, parent QObject_ITF) *QLibrary {
	return NewQLibraryFromPointer(C.QLibrary_NewQLibrary4(C.CString(fileName), C.CString(version), PointerFromQObject(parent)))
}

func NewQLibrary3(fileName string, verNum int, parent QObject_ITF) *QLibrary {
	return NewQLibraryFromPointer(C.QLibrary_NewQLibrary3(C.CString(fileName), C.int(verNum), PointerFromQObject(parent)))
}

func (ptr *QLibrary) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLibrary_ErrorString(ptr.Pointer()))
	}
	return ""
}

func QLibrary_IsLibrary(fileName string) bool {
	return C.QLibrary_QLibrary_IsLibrary(C.CString(fileName)) != 0
}

func (ptr *QLibrary) IsLoaded() bool {
	if ptr.Pointer() != nil {
		return C.QLibrary_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLibrary) Load() bool {
	if ptr.Pointer() != nil {
		return C.QLibrary_Load(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLibrary) SetFileNameAndVersion2(fileName string, version string) {
	if ptr.Pointer() != nil {
		C.QLibrary_SetFileNameAndVersion2(ptr.Pointer(), C.CString(fileName), C.CString(version))
	}
}

func (ptr *QLibrary) Unload() bool {
	if ptr.Pointer() != nil {
		return C.QLibrary_Unload(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLibrary) DestroyQLibrary() {
	if ptr.Pointer() != nil {
		C.QLibrary_DestroyQLibrary(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
