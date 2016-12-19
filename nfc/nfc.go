// +build !minimal

package nfc

//#include <stdint.h>
//#include <stdlib.h>
//#include "nfc.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtNfc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QNdefFilter struct {
	ptr unsafe.Pointer
}

type QNdefFilter_ITF interface {
	QNdefFilter_PTR() *QNdefFilter
}

func (p *QNdefFilter) QNdefFilter_PTR() *QNdefFilter {
	return p
}

func (p *QNdefFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNdefFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNdefFilter(ptr QNdefFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefFilter_PTR().Pointer()
	}
	return nil
}

func NewQNdefFilterFromPointer(ptr unsafe.Pointer) *QNdefFilter {
	var n = new(QNdefFilter)
	n.SetPointer(ptr)
	return n
}
func NewQNdefFilter() *QNdefFilter {
	var tmpValue = NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter())
	runtime.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {
	var tmpValue = NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter2(PointerFromQNdefFilter(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func (ptr *QNdefFilter) AppendRecord2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, min uint, max uint) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_AppendRecord2(ptr.Pointer(), C.longlong(typeNameFormat), core.PointerFromQByteArray(ty), C.uint(uint32(min)), C.uint(uint32(max)))
	}
}

func (ptr *QNdefFilter) Clear() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_Clear(ptr.Pointer())
	}
}

func (ptr *QNdefFilter) OrderMatch() bool {
	if ptr.Pointer() != nil {
		return C.QNdefFilter_OrderMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefFilter) RecordCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefFilter_RecordCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_DestroyQNdefFilter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNdefMessage struct {
	core.QList
}

type QNdefMessage_ITF interface {
	core.QList_ITF
	QNdefMessage_PTR() *QNdefMessage
}

func (p *QNdefMessage) QNdefMessage_PTR() *QNdefMessage {
	return p
}

func (p *QNdefMessage) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QList_PTR().Pointer()
	}
	return nil
}

func (p *QNdefMessage) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QList_PTR().SetPointer(ptr)
	}
}

func PointerFromQNdefMessage(ptr QNdefMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefMessage_PTR().Pointer()
	}
	return nil
}

func NewQNdefMessageFromPointer(ptr unsafe.Pointer) *QNdefMessage {
	var n = new(QNdefMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefMessage) DestroyQNdefMessage() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQNdefMessage() *QNdefMessage {
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage())
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage3(PointerFromQNdefMessage(message)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage2(PointerFromQNdefRecord(record)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func QNdefMessage_FromByteArray(message core.QByteArray_ITF) *QNdefMessage {
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(core.PointerFromQByteArray(message)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) FromByteArray(message core.QByteArray_ITF) *QNdefMessage {
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(core.PointerFromQByteArray(message)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) ToByteArray() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefMessage_ToByteArray(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int64

const (
	QNdefNfcSmartPosterRecord__UnspecifiedAction = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        = QNdefNfcSmartPosterRecord__Action(2)
)

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord
}

func (p *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord {
	return p
}

func (p *QNdefNfcSmartPosterRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (p *QNdefNfcSmartPosterRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QNdefRecord_PTR().SetPointer(ptr)
	}
}

func PointerFromQNdefNfcSmartPosterRecord(ptr QNdefNfcSmartPosterRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcSmartPosterRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcSmartPosterRecord {
	var n = new(QNdefNfcSmartPosterRecord)
	n.SetPointer(ptr)
	return n
}
func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {
	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {
	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(PointerFromQNdefNfcSmartPosterRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {
	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArray_ITF, data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty), core.PointerFromQByteArray(data))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.QNdefNfcSmartPosterRecord_AddTitle2(ptr.Pointer(), textC, localeC, C.longlong(encoding)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasAction() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasAction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasIcon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {
	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.QNdefNfcSmartPosterRecord_HasTitle(ptr.Pointer(), localeC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTypeInfo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_Icon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_IconCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {
	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.QNdefNfcSmartPosterRecord_RemoveTitle2(ptr.Pointer(), localeC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(ptr.Pointer(), C.longlong(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetSize(size uint) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetSize(ptr.Pointer(), C.uint(uint32(size)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri(ptr.Pointer(), PointerFromQNdefNfcUriRecord(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri2(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) Size() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QNdefNfcSmartPosterRecord_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return cGoUnpackString(C.QNdefNfcSmartPosterRecord_Title(ptr.Pointer(), localeC))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_TitleCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) TitleRecord(index int) *QNdefNfcTextRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcSmartPosterRecord_TitleRecord(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_TypeInfo(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNdefNfcSmartPosterRecord_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) UriRecord() *QNdefNfcUriRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNdefNfcUriRecordFromPointer(C.QNdefNfcSmartPosterRecord_UriRecord(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNdefNfcSmartPosterRecord) titleRecords_atList(i int) *QNdefNfcTextRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcSmartPosterRecord_titleRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
		return tmpValue
	}
	return nil
}

//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int64

const (
	QNdefNfcTextRecord__Utf8  = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 = QNdefNfcTextRecord__Encoding(1)
)

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord
}

func (p *QNdefNfcTextRecord) QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord {
	return p
}

func (p *QNdefNfcTextRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (p *QNdefNfcTextRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QNdefRecord_PTR().SetPointer(ptr)
	}
}

func PointerFromQNdefNfcTextRecord(ptr QNdefNfcTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcTextRecord {
	var n = new(QNdefNfcTextRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcTextRecord) DestroyQNdefNfcTextRecord() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {
	var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {
	var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {
	if ptr.Pointer() != nil {
		return QNdefNfcTextRecord__Encoding(C.QNdefNfcTextRecord_Encoding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcTextRecord) Locale() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNdefNfcTextRecord_Locale(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {
	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetEncoding(ptr.Pointer(), C.longlong(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		C.QNdefNfcTextRecord_SetLocale(ptr.Pointer(), localeC)
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QNdefNfcTextRecord_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QNdefNfcTextRecord) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNdefNfcTextRecord_Text(ptr.Pointer()))
	}
	return ""
}

type QNdefNfcUriRecord struct {
	QNdefRecord
}

type QNdefNfcUriRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord
}

func (p *QNdefNfcUriRecord) QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord {
	return p
}

func (p *QNdefNfcUriRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (p *QNdefNfcUriRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QNdefRecord_PTR().SetPointer(ptr)
	}
}

func PointerFromQNdefNfcUriRecord(ptr QNdefNfcUriRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcUriRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcUriRecord {
	var n = new(QNdefNfcUriRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcUriRecord) DestroyQNdefNfcUriRecord() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {
	var tmpValue = NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {
	var tmpValue = NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func (ptr *QNdefNfcUriRecord) SetUri(uri core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcUriRecord_SetUri(ptr.Pointer(), core.PointerFromQUrl(uri))
	}
}

func (ptr *QNdefNfcUriRecord) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNdefNfcUriRecord_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//QNdefRecord::TypeNameFormat
type QNdefRecord__TypeNameFormat int64

const (
	QNdefRecord__Empty       = QNdefRecord__TypeNameFormat(0x00)
	QNdefRecord__NfcRtd      = QNdefRecord__TypeNameFormat(0x01)
	QNdefRecord__Mime        = QNdefRecord__TypeNameFormat(0x02)
	QNdefRecord__Uri         = QNdefRecord__TypeNameFormat(0x03)
	QNdefRecord__ExternalRtd = QNdefRecord__TypeNameFormat(0x04)
	QNdefRecord__Unknown     = QNdefRecord__TypeNameFormat(0x05)
)

type QNdefRecord struct {
	ptr unsafe.Pointer
}

type QNdefRecord_ITF interface {
	QNdefRecord_PTR() *QNdefRecord
}

func (p *QNdefRecord) QNdefRecord_PTR() *QNdefRecord {
	return p
}

func (p *QNdefRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNdefRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNdefRecord(ptr QNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefRecordFromPointer(ptr unsafe.Pointer) *QNdefRecord {
	var n = new(QNdefRecord)
	n.SetPointer(ptr)
	return n
}
func NewQNdefRecord() *QNdefRecord {
	var tmpValue = NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {
	var tmpValue = NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func (ptr *QNdefRecord) Id() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefRecord_Id(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QNdefRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefRecord) Payload() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefRecord_Payload(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) SetId(id core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetId(ptr.Pointer(), core.PointerFromQByteArray(id))
	}
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetPayload(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QNdefRecord) SetType(ty core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetType(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(typeNameFormat))
	}
}

func (ptr *QNdefRecord) Type() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNdefRecord_Type(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QNdefRecord__TypeNameFormat(C.QNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {
	if ptr.Pointer() != nil {
		C.QNdefRecord_DestroyQNdefRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess = QNearFieldManager__TargetAccessMode(0x04)
)

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
}

func (p *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return p
}

func (p *QNearFieldManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNearFieldManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNearFieldManager(ptr QNearFieldManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldManagerFromPointer(ptr unsafe.Pointer) *QNearFieldManager {
	var n = new(QNearFieldManager)
	n.SetPointer(ptr)
	return n
}
func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_StartTargetDetection(ptr.Pointer()) != 0
	}
	return false
}

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {
	var tmpValue = NewQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.longlong(typeNameFormat), core.PointerFromQByteArray(ty), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.longlong(accessModes))
	}
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(ptr.Pointer())
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldManager_TargetDetected
func callbackQNearFieldManager_TargetDetected(ptr unsafe.Pointer, target unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::targetDetected"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::targetDetected")
	}
}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

//export callbackQNearFieldManager_TargetLost
func callbackQNearFieldManager_TargetLost(ptr unsafe.Pointer, target unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::targetLost"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::targetLost")
	}
}

func (ptr *QNearFieldManager) TargetLost(target QNearFieldTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetLost(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(int32(handlerId))) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldManager_TimerEvent
func callbackQNearFieldManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::timerEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::timerEvent")
	}
}

func (ptr *QNearFieldManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldManager_ChildEvent
func callbackQNearFieldManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::childEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::childEvent")
	}
}

func (ptr *QNearFieldManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldManager_ConnectNotify
func callbackQNearFieldManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::connectNotify", f)
	}
}

func (ptr *QNearFieldManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::connectNotify")
	}
}

func (ptr *QNearFieldManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_CustomEvent
func callbackQNearFieldManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::customEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::customEvent")
	}
}

func (ptr *QNearFieldManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldManager_DeleteLater
func callbackQNearFieldManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::deleteLater", f)
	}
}

func (ptr *QNearFieldManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::deleteLater")
	}
}

func (ptr *QNearFieldManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldManager_DisconnectNotify
func callbackQNearFieldManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::disconnectNotify", f)
	}
}

func (ptr *QNearFieldManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::disconnectNotify")
	}
}

func (ptr *QNearFieldManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_Event
func callbackQNearFieldManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::event", f)
	}
}

func (ptr *QNearFieldManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::event")
	}
}

func (ptr *QNearFieldManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldManager_EventFilter
func callbackQNearFieldManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::eventFilter", f)
	}
}

func (ptr *QNearFieldManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::eventFilter")
	}
}

func (ptr *QNearFieldManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldManager_MetaObject
func callbackQNearFieldManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::metaObject", f)
	}
}

func (ptr *QNearFieldManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldManager::metaObject")
	}
}

func (ptr *QNearFieldManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
	QNearFieldShareManager__NoError                     = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  = QNearFieldShareManager__ShareError(8)
)

//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare = QNearFieldShareManager__ShareMode(0x02)
)

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func (p *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return p
}

func (p *QNearFieldShareManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNearFieldShareManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	return n
}
func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	var tmpValue = NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNearFieldShareManager_Error
func callbackQNearFieldShareManager_Error(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::error")
	}
}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareManager_ShareModesChanged
func callbackQNearFieldShareManager_ShareModesChanged(ptr unsafe.Pointer, modes C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::shareModesChanged"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
	}

}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::shareModesChanged")
	}
}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ShareModesChanged(ptr.Pointer(), C.longlong(modes))
	}
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) SupportedShareModes() QNearFieldShareManager__ShareMode {
	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

//export callbackQNearFieldShareManager_TargetDetected
func callbackQNearFieldShareManager_TargetDetected(ptr unsafe.Pointer, shareTarget unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::targetDetected"); signal != nil {
		signal.(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
	}

}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::targetDetected")
	}
}

func (ptr *QNearFieldShareManager) TargetDetected(shareTarget QNearFieldShareTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldShareTarget(shareTarget))
	}
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareManager_TimerEvent
func callbackQNearFieldShareManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::timerEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::timerEvent")
	}
}

func (ptr *QNearFieldShareManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldShareManager_ChildEvent
func callbackQNearFieldShareManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::childEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::childEvent")
	}
}

func (ptr *QNearFieldShareManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareManager_ConnectNotify
func callbackQNearFieldShareManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::connectNotify", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::connectNotify")
	}
}

func (ptr *QNearFieldShareManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_CustomEvent
func callbackQNearFieldShareManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::customEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::customEvent")
	}
}

func (ptr *QNearFieldShareManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareManager_DeleteLater
func callbackQNearFieldShareManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::deleteLater", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::deleteLater")
	}
}

func (ptr *QNearFieldShareManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareManager_DisconnectNotify
func callbackQNearFieldShareManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::disconnectNotify", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::disconnectNotify")
	}
}

func (ptr *QNearFieldShareManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_Event
func callbackQNearFieldShareManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::event", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::event")
	}
}

func (ptr *QNearFieldShareManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldShareManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_EventFilter
func callbackQNearFieldShareManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::eventFilter", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::eventFilter")
	}
}

func (ptr *QNearFieldShareManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldShareManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_MetaObject
func callbackQNearFieldShareManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::metaObject", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareManager::metaObject")
	}
}

func (ptr *QNearFieldShareManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldShareManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
}

func (p *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return p
}

func (p *QNearFieldShareTarget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNearFieldShareTarget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) *QNearFieldShareTarget {
	var n = new(QNearFieldShareTarget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QNearFieldShareTarget) Cancel() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

//export callbackQNearFieldShareTarget_Error
func callbackQNearFieldShareTarget_Error(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::error")
	}
}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_IsShareInProgress(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Share(ptr.Pointer(), PointerFromQNdefMessage(message)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareTarget_ShareFinished
func callbackQNearFieldShareTarget_ShareFinished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::shareFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::shareFinished")
	}
}

func (ptr *QNearFieldShareTarget) ShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ShareFinished(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareTarget_TimerEvent
func callbackQNearFieldShareTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::timerEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::timerEvent")
	}
}

func (ptr *QNearFieldShareTarget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldShareTarget_ChildEvent
func callbackQNearFieldShareTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::childEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::childEvent")
	}
}

func (ptr *QNearFieldShareTarget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareTarget_ConnectNotify
func callbackQNearFieldShareTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::connectNotify", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::connectNotify")
	}
}

func (ptr *QNearFieldShareTarget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_CustomEvent
func callbackQNearFieldShareTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::customEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::customEvent")
	}
}

func (ptr *QNearFieldShareTarget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareTarget_DeleteLater
func callbackQNearFieldShareTarget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareTarget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::deleteLater", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::deleteLater")
	}
}

func (ptr *QNearFieldShareTarget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareTarget_DisconnectNotify
func callbackQNearFieldShareTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::disconnectNotify", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::disconnectNotify")
	}
}

func (ptr *QNearFieldShareTarget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_Event
func callbackQNearFieldShareTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareTarget) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::event", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::event")
	}
}

func (ptr *QNearFieldShareTarget) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_EventFilter
func callbackQNearFieldShareTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareTarget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::eventFilter", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::eventFilter")
	}
}

func (ptr *QNearFieldShareTarget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_MetaObject
func callbackQNearFieldShareTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldShareTarget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareTarget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::metaObject", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldShareTarget::metaObject")
	}
}

func (ptr *QNearFieldShareTarget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareTarget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldShareTarget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareTarget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNearFieldTarget::AccessMethod
type QNearFieldTarget__AccessMethod int64

const (
	QNearFieldTarget__UnknownAccess         = QNearFieldTarget__AccessMethod(0x00)
	QNearFieldTarget__NdefAccess            = QNearFieldTarget__AccessMethod(0x01)
	QNearFieldTarget__TagTypeSpecificAccess = QNearFieldTarget__AccessMethod(0x02)
	QNearFieldTarget__LlcpAccess            = QNearFieldTarget__AccessMethod(0x04)
)

//QNearFieldTarget::Error
type QNearFieldTarget__Error int64

const (
	QNearFieldTarget__NoError                = QNearFieldTarget__Error(0)
	QNearFieldTarget__UnknownError           = QNearFieldTarget__Error(1)
	QNearFieldTarget__UnsupportedError       = QNearFieldTarget__Error(2)
	QNearFieldTarget__TargetOutOfRangeError  = QNearFieldTarget__Error(3)
	QNearFieldTarget__NoResponseError        = QNearFieldTarget__Error(4)
	QNearFieldTarget__ChecksumMismatchError  = QNearFieldTarget__Error(5)
	QNearFieldTarget__InvalidParametersError = QNearFieldTarget__Error(6)
	QNearFieldTarget__NdefReadError          = QNearFieldTarget__Error(7)
	QNearFieldTarget__NdefWriteError         = QNearFieldTarget__Error(8)
)

//QNearFieldTarget::Type
type QNearFieldTarget__Type int64

const (
	QNearFieldTarget__ProprietaryTag = QNearFieldTarget__Type(0)
	QNearFieldTarget__NfcTagType1    = QNearFieldTarget__Type(1)
	QNearFieldTarget__NfcTagType2    = QNearFieldTarget__Type(2)
	QNearFieldTarget__NfcTagType3    = QNearFieldTarget__Type(3)
	QNearFieldTarget__NfcTagType4    = QNearFieldTarget__Type(4)
	QNearFieldTarget__MifareTag      = QNearFieldTarget__Type(5)
)

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTarget_ITF interface {
	core.QObject_ITF
	QNearFieldTarget_PTR() *QNearFieldTarget
}

func (p *QNearFieldTarget) QNearFieldTarget_PTR() *QNearFieldTarget {
	return p
}

func (p *QNearFieldTarget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNearFieldTarget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNearFieldTarget(ptr QNearFieldTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldTargetFromPointer(ptr unsafe.Pointer) *QNearFieldTarget {
	var n = new(QNearFieldTarget)
	n.SetPointer(ptr)
	return n
}
func NewQNearFieldTarget(parent core.QObject_ITF) *QNearFieldTarget {
	var tmpValue = NewQNearFieldTargetFromPointer(C.QNearFieldTarget_NewQNearFieldTarget(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNearFieldTarget_AccessMethods
func callbackQNearFieldTarget_AccessMethods(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::accessMethods"); signal != nil {
		return C.longlong(signal.(func() QNearFieldTarget__AccessMethod)())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectAccessMethods(f func() QNearFieldTarget__AccessMethod) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::accessMethods", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectAccessMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::accessMethods")
	}
}

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldTarget_Disconnected
func callbackQNearFieldTarget_Disconnected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::disconnected", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::disconnected")
	}
}

func (ptr *QNearFieldTarget) Disconnected() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_Disconnected(ptr.Pointer())
	}
}

func (ptr *QNearFieldTarget) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::error")
	}
}

func (ptr *QNearFieldTarget) DisconnectHandleResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::handleResponse")
	}
}

//export callbackQNearFieldTarget_HasNdefMessage
func callbackQNearFieldTarget_HasNdefMessage(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::hasNdefMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).HasNdefMessageDefault())))
}

func (ptr *QNearFieldTarget) ConnectHasNdefMessage(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::hasNdefMessage", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectHasNdefMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::hasNdefMessage")
	}
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) HasNdefMessageDefault() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessageDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_IsProcessingCommand(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNearFieldTarget_NdefMessageRead
func callbackQNearFieldTarget_NdefMessageRead(ptr unsafe.Pointer, message unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::ndefMessageRead"); signal != nil {
		signal.(func(*QNdefMessage))(NewQNdefMessageFromPointer(message))
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessageRead(f func(message *QNdefMessage)) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessageRead(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::ndefMessageRead", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessageRead() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessageRead(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::ndefMessageRead")
	}
}

func (ptr *QNearFieldTarget) NdefMessageRead(message QNdefMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessageRead(ptr.Pointer(), PointerFromQNdefMessage(message))
	}
}

//export callbackQNearFieldTarget_NdefMessagesWritten
func callbackQNearFieldTarget_NdefMessagesWritten(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::ndefMessagesWritten"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessagesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::ndefMessagesWritten", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::ndefMessagesWritten")
	}
}

func (ptr *QNearFieldTarget) NdefMessagesWritten() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessagesWritten(ptr.Pointer())
	}
}

func (ptr *QNearFieldTarget) DisconnectReadNdefMessages() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::readNdefMessages")
	}
}

func (ptr *QNearFieldTarget) DisconnectRequestCompleted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::requestCompleted")
	}
}

func (ptr *QNearFieldTarget) DisconnectSendCommand() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::sendCommand")
	}
}

//export callbackQNearFieldTarget_Type
func callbackQNearFieldTarget_Type(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::type"); signal != nil {
		return C.longlong(signal.(func() QNearFieldTarget__Type)())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectType(f func() QNearFieldTarget__Type) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::type", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::type")
	}
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldTarget_Uid
func callbackQNearFieldTarget_Uid(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::uid"); signal != nil {
		return core.PointerFromQByteArray(signal.(func() *core.QByteArray)())
	}

	return core.PointerFromQByteArray(nil)
}

func (ptr *QNearFieldTarget) ConnectUid(f func() *core.QByteArray) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::uid", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectUid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::uid")
	}
}

func (ptr *QNearFieldTarget) Uid() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNearFieldTarget_Uid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQNearFieldTarget_Url
func callbackQNearFieldTarget_Url(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::url"); signal != nil {
		return core.PointerFromQUrl(signal.(func() *core.QUrl)())
	}

	return core.PointerFromQUrl(NewQNearFieldTargetFromPointer(ptr).UrlDefault())
}

func (ptr *QNearFieldTarget) ConnectUrl(f func() *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::url", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::url")
	}
}

func (ptr *QNearFieldTarget) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNearFieldTarget_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) UrlDefault() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNearFieldTarget_UrlDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) DisconnectWaitForRequestCompleted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::waitForRequestCompleted")
	}
}

//export callbackQNearFieldTarget_DestroyQNearFieldTarget
func callbackQNearFieldTarget_DestroyQNearFieldTarget(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::~QNearFieldTarget"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldTargetFromPointer(ptr).DestroyQNearFieldTargetDefault()
	}
}

func (ptr *QNearFieldTarget) ConnectDestroyQNearFieldTarget(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::~QNearFieldTarget", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDestroyQNearFieldTarget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::~QNearFieldTarget")
	}
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DestroyQNearFieldTarget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTargetDefault() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DestroyQNearFieldTargetDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldTarget_TimerEvent
func callbackQNearFieldTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::timerEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::timerEvent")
	}
}

func (ptr *QNearFieldTarget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldTarget_ChildEvent
func callbackQNearFieldTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::childEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::childEvent")
	}
}

func (ptr *QNearFieldTarget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldTarget_ConnectNotify
func callbackQNearFieldTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::connectNotify", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::connectNotify")
	}
}

func (ptr *QNearFieldTarget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_CustomEvent
func callbackQNearFieldTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::customEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::customEvent")
	}
}

func (ptr *QNearFieldTarget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldTarget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldTarget_DeleteLater
func callbackQNearFieldTarget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldTarget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::deleteLater", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::deleteLater")
	}
}

func (ptr *QNearFieldTarget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldTarget_DisconnectNotify
func callbackQNearFieldTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::disconnectNotify", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::disconnectNotify")
	}
}

func (ptr *QNearFieldTarget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_Event
func callbackQNearFieldTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldTarget) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::event", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::event")
	}
}

func (ptr *QNearFieldTarget) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldTarget_EventFilter
func callbackQNearFieldTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldTarget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::eventFilter", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::eventFilter")
	}
}

func (ptr *QNearFieldTarget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldTarget_MetaObject
func callbackQNearFieldTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNearFieldTarget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldTarget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::metaObject", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNearFieldTarget::metaObject")
	}
}

func (ptr *QNearFieldTarget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldTarget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldTarget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldTarget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QQmlNdefRecord::TypeNameFormat
type QQmlNdefRecord__TypeNameFormat int64

const (
	QQmlNdefRecord__Empty       = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Empty)
	QQmlNdefRecord__NfcRtd      = QQmlNdefRecord__TypeNameFormat(QNdefRecord__NfcRtd)
	QQmlNdefRecord__Mime        = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Mime)
	QQmlNdefRecord__Uri         = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Uri)
	QQmlNdefRecord__ExternalRtd = QQmlNdefRecord__TypeNameFormat(QNdefRecord__ExternalRtd)
	QQmlNdefRecord__Unknown     = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Unknown)
)

type QQmlNdefRecord struct {
	core.QObject
}

type QQmlNdefRecord_ITF interface {
	core.QObject_ITF
	QQmlNdefRecord_PTR() *QQmlNdefRecord
}

func (p *QQmlNdefRecord) QQmlNdefRecord_PTR() *QQmlNdefRecord {
	return p
}

func (p *QQmlNdefRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlNdefRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQmlNdefRecord(ptr QQmlNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQQmlNdefRecordFromPointer(ptr unsafe.Pointer) *QQmlNdefRecord {
	var n = new(QQmlNdefRecord)
	n.SetPointer(ptr)
	return n
}
func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	var tmpValue = NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	var tmpValue = NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlNdefRecord) Record() *QNdefRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNdefRecordFromPointer(C.QQmlNdefRecord_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
		return tmpValue
	}
	return nil
}

//export callbackQQmlNdefRecord_RecordChanged
func callbackQQmlNdefRecord_RecordChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::recordChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::recordChanged")
	}
}

func (ptr *QQmlNdefRecord) RecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_RecordChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(ptr.Pointer(), PointerFromQNdefRecord(record))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	if ptr.Pointer() != nil {
		var newtypeC = C.CString(newtype)
		defer C.free(unsafe.Pointer(newtypeC))
		C.QQmlNdefRecord_SetType(ptr.Pointer(), newtypeC)
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

//export callbackQQmlNdefRecord_TypeChanged
func callbackQQmlNdefRecord_TypeChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::typeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::typeChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeChanged(ptr.Pointer())
	}
}

//export callbackQQmlNdefRecord_TypeNameFormatChanged
func callbackQQmlNdefRecord_TypeNameFormatChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::typeNameFormatChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::typeNameFormatChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeNameFormatChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecord() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DestroyQQmlNdefRecord(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlNdefRecord_TimerEvent
func callbackQQmlNdefRecord_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::timerEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::timerEvent")
	}
}

func (ptr *QQmlNdefRecord) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlNdefRecord) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlNdefRecord_ChildEvent
func callbackQQmlNdefRecord_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::childEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::childEvent")
	}
}

func (ptr *QQmlNdefRecord) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlNdefRecord) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlNdefRecord_ConnectNotify
func callbackQQmlNdefRecord_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::connectNotify", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::connectNotify")
	}
}

func (ptr *QQmlNdefRecord) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_CustomEvent
func callbackQQmlNdefRecord_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::customEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::customEvent")
	}
}

func (ptr *QQmlNdefRecord) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlNdefRecord) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlNdefRecord_DeleteLater
func callbackQQmlNdefRecord_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlNdefRecord) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::deleteLater", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::deleteLater")
	}
}

func (ptr *QQmlNdefRecord) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNdefRecord) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlNdefRecord_DisconnectNotify
func callbackQQmlNdefRecord_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::disconnectNotify", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::disconnectNotify")
	}
}

func (ptr *QQmlNdefRecord) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlNdefRecord) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_Event
func callbackQQmlNdefRecord_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlNdefRecord) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::event", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::event")
	}
}

func (ptr *QQmlNdefRecord) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlNdefRecord) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_EventFilter
func callbackQQmlNdefRecord_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlNdefRecord) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::eventFilter", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::eventFilter")
	}
}

func (ptr *QQmlNdefRecord) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlNdefRecord) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_MetaObject
func callbackQQmlNdefRecord_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNdefRecord::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlNdefRecordFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlNdefRecord) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::metaObject", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNdefRecord::metaObject")
	}
}

func (ptr *QQmlNdefRecord) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlNdefRecord_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlNdefRecord) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlNdefRecord_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
