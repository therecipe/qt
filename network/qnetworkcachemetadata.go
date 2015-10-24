package network

//#include "qnetworkcachemetadata.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkCacheMetaData struct {
	ptr unsafe.Pointer
}

type QNetworkCacheMetaDataITF interface {
	QNetworkCacheMetaDataPTR() *QNetworkCacheMetaData
}

func (p *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCacheMetaData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaDataITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaDataPTR().Pointer()
	}
	return nil
}

func QNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkCacheMetaData) QNetworkCacheMetaDataPTR() *QNetworkCacheMetaData {
	return ptr
}

func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	return QNetworkCacheMetaDataFromPointer(unsafe.Pointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData()))
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaDataITF) *QNetworkCacheMetaData {
	return QNetworkCacheMetaDataFromPointer(unsafe.Pointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(C.QtObjectPtr(PointerFromQNetworkCacheMetaData(other)))))
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(dateTime)))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(dateTime)))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url string) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaDataITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCacheMetaData(other)))
	}
}

func (ptr *QNetworkCacheMetaData) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCacheMetaData_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(C.QtObjectPtr(ptr.Pointer()))
	}
}
