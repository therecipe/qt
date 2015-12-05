package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::QNetworkCacheMetaData")
		}
	}()

	return NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::QNetworkCacheMetaData")
		}
	}()

	return NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::expirationDate")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::lastModified")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::saveToDisk")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::setExpirationDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::setLastModified")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::setSaveToDisk")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(ptr.Pointer(), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::setUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(ptr.Pointer(), PointerFromQNetworkCacheMetaData(other))
	}
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCacheMetaData::~QNetworkCacheMetaData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
	}
}
