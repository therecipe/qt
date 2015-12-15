package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QNetworkAddressEntry struct {
	ptr unsafe.Pointer
}

type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (p *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkAddressEntry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntry_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry {
	return ptr
}

func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	return NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry())
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	return NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(PointerFromQNetworkAddressEntry(other)))
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {
	defer qt.Recovering("QNetworkAddressEntry::prefixLength")

	if ptr.Pointer() != nil {
		return int(C.QNetworkAddressEntry_PrefixLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setBroadcast")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(ptr.Pointer(), PointerFromQHostAddress(newBroadcast))
	}
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setIp")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetIp(ptr.Pointer(), PointerFromQHostAddress(newIp))
	}
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setNetmask")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetNetmask(ptr.Pointer(), PointerFromQHostAddress(newNetmask))
	}
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {
	defer qt.Recovering("QNetworkAddressEntry::setPrefixLength")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetPrefixLength(ptr.Pointer(), C.int(length))
	}
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::swap")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_Swap(ptr.Pointer(), PointerFromQNetworkAddressEntry(other))
	}
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {
	defer qt.Recovering("QNetworkAddressEntry::~QNetworkAddressEntry")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(ptr.Pointer())
	}
}
