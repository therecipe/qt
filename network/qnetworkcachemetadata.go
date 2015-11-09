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

type QNetworkCacheMetaData_ITF interface {
	QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData
}

func (p *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCacheMetaData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkCacheMetaData) QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData {
	return ptr
}

func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	return NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	return NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(ptr.Pointer(), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(ptr.Pointer(), PointerFromQNetworkCacheMetaData(other))
	}
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
	}
}
