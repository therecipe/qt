package network

//#include "qnetworkaddressentry.h"
import "C"
import (
	"unsafe"
)

type QNetworkAddressEntry struct {
	ptr unsafe.Pointer
}

type QNetworkAddressEntryITF interface {
	QNetworkAddressEntryPTR() *QNetworkAddressEntry
}

func (p *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkAddressEntry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntryPTR().Pointer()
	}
	return nil
}

func QNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntryPTR() *QNetworkAddressEntry {
	return ptr
}

func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	return QNetworkAddressEntryFromPointer(unsafe.Pointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry()))
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntryITF) *QNetworkAddressEntry {
	return QNetworkAddressEntryFromPointer(unsafe.Pointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(C.QtObjectPtr(PointerFromQNetworkAddressEntry(other)))))
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {
	if ptr.Pointer() != nil {
		return int(C.QNetworkAddressEntry_PrefixLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddressITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(newBroadcast)))
	}
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddressITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetIp(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(newIp)))
	}
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddressITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetNetmask(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(newNetmask)))
	}
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetPrefixLength(C.QtObjectPtr(ptr.Pointer()), C.int(length))
	}
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntryITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkAddressEntry(other)))
	}
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(C.QtObjectPtr(ptr.Pointer()))
	}
}
