// +build !minimal

package nfc

//#include <stdint.h>
//#include <stdlib.h>
//#include "nfc.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

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
	defer qt.Recovering("QNdefFilter::QNdefFilter")

	var tmpValue = NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter())
	runtime.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {
	defer qt.Recovering("QNdefFilter::QNdefFilter")

	var tmpValue = NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter2(PointerFromQNdefFilter(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func (ptr *QNdefFilter) AppendRecord2(typeNameFormat QNdefRecord__TypeNameFormat, ty string, min uint, max uint) {
	defer qt.Recovering("QNdefFilter::appendRecord")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		C.QNdefFilter_AppendRecord2(ptr.Pointer(), C.longlong(typeNameFormat), tyC, C.uint(uint32(min)), C.uint(uint32(max)))
	}
}

func (ptr *QNdefFilter) Clear() {
	defer qt.Recovering("QNdefFilter::clear")

	if ptr.Pointer() != nil {
		C.QNdefFilter_Clear(ptr.Pointer())
	}
}

func (ptr *QNdefFilter) OrderMatch() bool {
	defer qt.Recovering("QNdefFilter::orderMatch")

	if ptr.Pointer() != nil {
		return C.QNdefFilter_OrderMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefFilter) RecordCount() int {
	defer qt.Recovering("QNdefFilter::recordCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QNdefFilter_RecordCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	defer qt.Recovering("QNdefFilter::setOrderMatch")

	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	defer qt.Recovering("QNdefFilter::~QNdefFilter")

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
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage())
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage3(PointerFromQNdefMessage(message)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage2(PointerFromQNdefRecord(record)))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func QNdefMessage_FromByteArray(message string) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::fromByteArray")

	var messageC = C.CString(hex.EncodeToString([]byte(message)))
	defer C.free(unsafe.Pointer(messageC))
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(messageC))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) FromByteArray(message string) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::fromByteArray")

	var messageC = C.CString(hex.EncodeToString([]byte(message)))
	defer C.free(unsafe.Pointer(messageC))
	var tmpValue = NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(messageC))
	runtime.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) ToByteArray() string {
	defer qt.Recovering("QNdefMessage::toByteArray")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNdefMessage_ToByteArray(ptr.Pointer())))
	}
	return ""
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
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(PointerFromQNdefNfcSmartPosterRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	var tmpValue = NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::action")

	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty string, data string) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::addIcon")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		C.QNdefNfcSmartPosterRecord_AddIcon2(ptr.Pointer(), tyC, dataC)
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecord_ITF) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::addTitle")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::addTitle")

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
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasAction")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasAction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype string) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasIcon")

	if ptr.Pointer() != nil {
		var mimetypeC = C.CString(hex.EncodeToString([]byte(mimetype)))
		defer C.free(unsafe.Pointer(mimetypeC))
		return C.QNdefNfcSmartPosterRecord_HasIcon(ptr.Pointer(), mimetypeC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasSize")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasTitle")

	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.QNdefNfcSmartPosterRecord_HasTitle(ptr.Pointer(), localeC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasTypeInfo")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTypeInfo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype string) string {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::icon")

	if ptr.Pointer() != nil {
		var mimetypeC = C.CString(hex.EncodeToString([]byte(mimetype)))
		defer C.free(unsafe.Pointer(mimetypeC))
		return qt.HexDecodeToString(C.GoString(C.QNdefNfcSmartPosterRecord_Icon(ptr.Pointer(), mimetypeC)))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::iconCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_IconCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty string) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::removeIcon")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		return C.QNdefNfcSmartPosterRecord_RemoveIcon2(ptr.Pointer(), tyC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecord_ITF) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::removeTitle")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::removeTitle")

	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.QNdefNfcSmartPosterRecord_RemoveTitle2(ptr.Pointer(), localeC) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setAction")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(ptr.Pointer(), C.longlong(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetSize(size uint) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setSize")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetSize(ptr.Pointer(), C.uint(uint32(size)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty string) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setTypeInfo")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(ptr.Pointer(), tyC)
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecord_ITF) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setUri")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri(ptr.Pointer(), PointerFromQNdefNfcUriRecord(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url core.QUrl_ITF) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setUri")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri2(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) Size() uint {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::size")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QNdefNfcSmartPosterRecord_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::title")

	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		return C.GoString(C.QNdefNfcSmartPosterRecord_Title(ptr.Pointer(), localeC))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::titleCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_TitleCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() string {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::typeInfo")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNdefNfcSmartPosterRecord_TypeInfo(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() *core.QUrl {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::uri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNdefNfcSmartPosterRecord_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::~QNdefNfcSmartPosterRecord")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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
	defer qt.Recovering("QNdefNfcTextRecord::QNdefNfcTextRecord")

	var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {
	defer qt.Recovering("QNdefNfcTextRecord::QNdefNfcTextRecord")

	var tmpValue = NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {
	defer qt.Recovering("QNdefNfcTextRecord::encoding")

	if ptr.Pointer() != nil {
		return QNdefNfcTextRecord__Encoding(C.QNdefNfcTextRecord_Encoding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcTextRecord) Locale() string {
	defer qt.Recovering("QNdefNfcTextRecord::locale")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Locale(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {
	defer qt.Recovering("QNdefNfcTextRecord::setEncoding")

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetEncoding(ptr.Pointer(), C.longlong(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	defer qt.Recovering("QNdefNfcTextRecord::setLocale")

	if ptr.Pointer() != nil {
		var localeC = C.CString(locale)
		defer C.free(unsafe.Pointer(localeC))
		C.QNdefNfcTextRecord_SetLocale(ptr.Pointer(), localeC)
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	defer qt.Recovering("QNdefNfcTextRecord::setText")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QNdefNfcTextRecord_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QNdefNfcTextRecord) Text() string {
	defer qt.Recovering("QNdefNfcTextRecord::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Text(ptr.Pointer()))
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
	defer qt.Recovering("QNdefNfcUriRecord::QNdefNfcUriRecord")

	var tmpValue = NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {
	defer qt.Recovering("QNdefNfcUriRecord::QNdefNfcUriRecord")

	var tmpValue = NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func (ptr *QNdefNfcUriRecord) SetUri(uri core.QUrl_ITF) {
	defer qt.Recovering("QNdefNfcUriRecord::setUri")

	if ptr.Pointer() != nil {
		C.QNdefNfcUriRecord_SetUri(ptr.Pointer(), core.PointerFromQUrl(uri))
	}
}

func (ptr *QNdefNfcUriRecord) Uri() *core.QUrl {
	defer qt.Recovering("QNdefNfcUriRecord::uri")

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
	defer qt.Recovering("QNdefRecord::QNdefRecord")

	var tmpValue = NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord())
	runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {
	defer qt.Recovering("QNdefRecord::QNdefRecord")

	var tmpValue = NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord2(PointerFromQNdefRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func (ptr *QNdefRecord) Id() string {
	defer qt.Recovering("QNdefRecord::id")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNdefRecord_Id(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefRecord) IsEmpty() bool {
	defer qt.Recovering("QNdefRecord::isEmpty")

	if ptr.Pointer() != nil {
		return C.QNdefRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefRecord) Payload() string {
	defer qt.Recovering("QNdefRecord::payload")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNdefRecord_Payload(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefRecord) SetId(id string) {
	defer qt.Recovering("QNdefRecord::setId")

	if ptr.Pointer() != nil {
		var idC = C.CString(hex.EncodeToString([]byte(id)))
		defer C.free(unsafe.Pointer(idC))
		C.QNdefRecord_SetId(ptr.Pointer(), idC)
	}
}

func (ptr *QNdefRecord) SetPayload(payload string) {
	defer qt.Recovering("QNdefRecord::setPayload")

	if ptr.Pointer() != nil {
		var payloadC = C.CString(hex.EncodeToString([]byte(payload)))
		defer C.free(unsafe.Pointer(payloadC))
		C.QNdefRecord_SetPayload(ptr.Pointer(), payloadC)
	}
}

func (ptr *QNdefRecord) SetType(ty string) {
	defer qt.Recovering("QNdefRecord::setType")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		C.QNdefRecord_SetType(ptr.Pointer(), tyC)
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	defer qt.Recovering("QNdefRecord::setTypeNameFormat")

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(typeNameFormat))
	}
}

func (ptr *QNdefRecord) Type() string {
	defer qt.Recovering("QNdefRecord::type")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNdefRecord_Type(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {
	defer qt.Recovering("QNdefRecord::typeNameFormat")

	if ptr.Pointer() != nil {
		return QNdefRecord__TypeNameFormat(C.QNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {
	defer qt.Recovering("QNdefRecord::~QNdefRecord")

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
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	defer qt.Recovering("QNearFieldManager::startTargetDetection")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_StartTargetDetection(ptr.Pointer()) != 0
	}
	return false
}

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {
	defer qt.Recovering("QNearFieldManager::QNearFieldManager")

	return NewQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	defer qt.Recovering("QNearFieldManager::isAvailable")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty string, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		var tyC = C.CString(hex.EncodeToString([]byte(ty)))
		defer C.free(unsafe.Pointer(tyC))
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.longlong(typeNameFormat), tyC, core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	defer qt.Recovering("QNearFieldManager::setTargetAccessModes")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.longlong(accessModes))
	}
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	defer qt.Recovering("QNearFieldManager::stopTargetDetection")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(ptr.Pointer())
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	defer qt.Recovering("QNearFieldManager::targetAccessModes")

	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldManager_TargetDetected
func callbackQNearFieldManager_TargetDetected(ptr unsafe.Pointer, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetDetected")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "targetDetected")
	}
}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {
	defer qt.Recovering("QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

//export callbackQNearFieldManager_TargetLost
func callbackQNearFieldManager_TargetLost(ptr unsafe.Pointer, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetLost")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "targetLost"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	defer qt.Recovering("disconnect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "targetLost")
	}
}

func (ptr *QNearFieldManager) TargetLost(target QNearFieldTarget_ITF) {
	defer qt.Recovering("QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetLost(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	defer qt.Recovering("QNearFieldManager::unregisterNdefMessageHandler")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(int32(handlerId))) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	defer qt.Recovering("QNearFieldManager::~QNearFieldManager")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()))
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldManager_TimerEvent
func callbackQNearFieldManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QNearFieldManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldManager_ChildEvent
func callbackQNearFieldManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QNearFieldManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldManager_ConnectNotify
func callbackQNearFieldManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QNearFieldManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNearFieldManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QNearFieldManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_CustomEvent
func callbackQNearFieldManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QNearFieldManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldManager_DeleteLater
func callbackQNearFieldManager_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNearFieldManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QNearFieldManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNearFieldManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QNearFieldManager) DeleteLater() {
	defer qt.Recovering("QNearFieldManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()))
		C.QNearFieldManager_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) DeleteLaterDefault() {
	defer qt.Recovering("QNearFieldManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()))
		C.QNearFieldManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldManager_DisconnectNotify
func callbackQNearFieldManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QNearFieldManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNearFieldManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QNearFieldManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_Event
func callbackQNearFieldManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldManager::event")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QNearFieldManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QNearFieldManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldManager::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldManager::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldManager_EventFilter
func callbackQNearFieldManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldManager::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QNearFieldManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNearFieldManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QNearFieldManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldManager_MetaObject
func callbackQNearFieldManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNearFieldManager::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNearFieldManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QNearFieldManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNearFieldManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldManager(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QNearFieldManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNearFieldManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNearFieldManager::metaObject")

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
	defer qt.Recovering("QNearFieldShareManager::QNearFieldShareManager")

	return NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
}

//export callbackQNearFieldShareManager_Error
func callbackQNearFieldShareManager_Error(ptr unsafe.Pointer, error C.longlong) {
	defer qt.Recovering("callback QNearFieldShareManager::error")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "error")
	}
}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::setShareModes")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	defer qt.Recovering("QNearFieldShareManager::shareError")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::shareModes")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareManager_ShareModesChanged
func callbackQNearFieldShareManager_ShareModesChanged(ptr unsafe.Pointer, modes C.longlong) {
	defer qt.Recovering("callback QNearFieldShareManager::shareModesChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "shareModesChanged"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
	}

}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	defer qt.Recovering("connect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	defer qt.Recovering("disconnect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "shareModesChanged")
	}
}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ShareModesChanged(ptr.Pointer(), C.longlong(modes))
	}
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::supportedShareModes")

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::supportedShareModes")

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

//export callbackQNearFieldShareManager_TargetDetected
func callbackQNearFieldShareManager_TargetDetected(ptr unsafe.Pointer, shareTarget unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::targetDetected")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
	}

}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	defer qt.Recovering("connect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "targetDetected")
	}
}

func (ptr *QNearFieldShareManager) TargetDetected(shareTarget QNearFieldShareTarget_ITF) {
	defer qt.Recovering("QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldShareTarget(shareTarget))
	}
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	defer qt.Recovering("QNearFieldShareManager::~QNearFieldShareManager")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()))
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareManager_TimerEvent
func callbackQNearFieldShareManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QNearFieldShareManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldShareManager_ChildEvent
func callbackQNearFieldShareManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QNearFieldShareManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareManager_ConnectNotify
func callbackQNearFieldShareManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldShareManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNearFieldShareManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QNearFieldShareManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_CustomEvent
func callbackQNearFieldShareManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QNearFieldShareManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareManager_DeleteLater
func callbackQNearFieldShareManager_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNearFieldShareManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNearFieldShareManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QNearFieldShareManager) DeleteLater() {
	defer qt.Recovering("QNearFieldShareManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()))
		C.QNearFieldShareManager_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) DeleteLaterDefault() {
	defer qt.Recovering("QNearFieldShareManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()))
		C.QNearFieldShareManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareManager_DisconnectNotify
func callbackQNearFieldShareManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldShareManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNearFieldShareManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QNearFieldShareManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_Event
func callbackQNearFieldShareManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldShareManager::event")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldShareManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QNearFieldShareManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareManager::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldShareManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareManager::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_EventFilter
func callbackQNearFieldShareManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldShareManager::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldShareManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNearFieldShareManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QNearFieldShareManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldShareManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_MetaObject
func callbackQNearFieldShareManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNearFieldShareManager::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNearFieldShareManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNearFieldShareManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareManager(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QNearFieldShareManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNearFieldShareManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldShareManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNearFieldShareManager::metaObject")

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
	defer qt.Recovering("QNearFieldShareTarget::cancel")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

//export callbackQNearFieldShareTarget_Error
func callbackQNearFieldShareTarget_Error(ptr unsafe.Pointer, error C.longlong) {
	defer qt.Recovering("callback QNearFieldShareTarget::error")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "error")
	}
}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	defer qt.Recovering("QNearFieldShareTarget::isShareInProgress")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_IsShareInProgress(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::share")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Share(ptr.Pointer(), PointerFromQNdefMessage(message)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	defer qt.Recovering("QNearFieldShareTarget::shareError")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareTarget_ShareFinished
func callbackQNearFieldShareTarget_ShareFinished(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::shareFinished")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "shareFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	defer qt.Recovering("connect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "shareFinished")
	}
}

func (ptr *QNearFieldShareTarget) ShareFinished() {
	defer qt.Recovering("QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ShareFinished(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareTarget::shareModes")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	defer qt.Recovering("QNearFieldShareTarget::~QNearFieldShareTarget")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()))
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareTarget_TimerEvent
func callbackQNearFieldShareTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QNearFieldShareTarget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldShareTarget_ChildEvent
func callbackQNearFieldShareTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QNearFieldShareTarget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareTarget_ConnectNotify
func callbackQNearFieldShareTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldShareTarget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QNearFieldShareTarget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_CustomEvent
func callbackQNearFieldShareTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QNearFieldShareTarget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareTarget_DeleteLater
func callbackQNearFieldShareTarget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareTarget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNearFieldShareTarget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QNearFieldShareTarget) DeleteLater() {
	defer qt.Recovering("QNearFieldShareTarget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()))
		C.QNearFieldShareTarget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) DeleteLaterDefault() {
	defer qt.Recovering("QNearFieldShareTarget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()))
		C.QNearFieldShareTarget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldShareTarget_DisconnectNotify
func callbackQNearFieldShareTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldShareTarget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QNearFieldShareTarget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldShareTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_Event
func callbackQNearFieldShareTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldShareTarget::event")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareTarget) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldShareTarget::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QNearFieldShareTarget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_EventFilter
func callbackQNearFieldShareTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldShareTarget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareTarget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldShareTarget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QNearFieldShareTarget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_MetaObject
func callbackQNearFieldShareTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNearFieldShareTarget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareTarget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNearFieldShareTarget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldShareTarget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QNearFieldShareTarget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNearFieldShareTarget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareTarget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldShareTarget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNearFieldShareTarget::metaObject")

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
	defer qt.Recovering("QNearFieldTarget::QNearFieldTarget")

	return NewQNearFieldTargetFromPointer(C.QNearFieldTarget_NewQNearFieldTarget(core.PointerFromQObject(parent)))
}

//export callbackQNearFieldTarget_AccessMethods
func callbackQNearFieldTarget_AccessMethods(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNearFieldTarget::accessMethods")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "accessMethods"); signal != nil {
		return C.longlong(signal.(func() QNearFieldTarget__AccessMethod)())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectAccessMethods(f func() QNearFieldTarget__AccessMethod) {
	defer qt.Recovering("connect QNearFieldTarget::accessMethods")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "accessMethods", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectAccessMethods() {
	defer qt.Recovering("disconnect QNearFieldTarget::accessMethods")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "accessMethods")
	}
}

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	defer qt.Recovering("QNearFieldTarget::accessMethods")

	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldTarget_Disconnected
func callbackQNearFieldTarget_Disconnected(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::disconnected")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "disconnected", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "disconnected")
	}
}

func (ptr *QNearFieldTarget) Disconnected() {
	defer qt.Recovering("QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_Disconnected(ptr.Pointer())
	}
}

//export callbackQNearFieldTarget_HasNdefMessage
func callbackQNearFieldTarget_HasNdefMessage(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldTarget::hasNdefMessage")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "hasNdefMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).HasNdefMessageDefault())))
}

func (ptr *QNearFieldTarget) ConnectHasNdefMessage(f func() bool) {
	defer qt.Recovering("connect QNearFieldTarget::hasNdefMessage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "hasNdefMessage", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectHasNdefMessage() {
	defer qt.Recovering("disconnect QNearFieldTarget::hasNdefMessage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "hasNdefMessage")
	}
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	defer qt.Recovering("QNearFieldTarget::hasNdefMessage")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) HasNdefMessageDefault() bool {
	defer qt.Recovering("QNearFieldTarget::hasNdefMessage")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessageDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {
	defer qt.Recovering("QNearFieldTarget::isProcessingCommand")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_IsProcessingCommand(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNearFieldTarget_NdefMessageRead
func callbackQNearFieldTarget_NdefMessageRead(ptr unsafe.Pointer, message unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::ndefMessageRead")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "ndefMessageRead"); signal != nil {
		signal.(func(*QNdefMessage))(NewQNdefMessageFromPointer(message))
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessageRead(f func(message *QNdefMessage)) {
	defer qt.Recovering("connect QNearFieldTarget::ndefMessageRead")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessageRead(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "ndefMessageRead", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessageRead() {
	defer qt.Recovering("disconnect QNearFieldTarget::ndefMessageRead")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessageRead(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "ndefMessageRead")
	}
}

func (ptr *QNearFieldTarget) NdefMessageRead(message QNdefMessage_ITF) {
	defer qt.Recovering("QNearFieldTarget::ndefMessageRead")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessageRead(ptr.Pointer(), PointerFromQNdefMessage(message))
	}
}

//export callbackQNearFieldTarget_NdefMessagesWritten
func callbackQNearFieldTarget_NdefMessagesWritten(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::ndefMessagesWritten")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "ndefMessagesWritten"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	defer qt.Recovering("connect QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessagesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "ndefMessagesWritten", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	defer qt.Recovering("disconnect QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "ndefMessagesWritten")
	}
}

func (ptr *QNearFieldTarget) NdefMessagesWritten() {
	defer qt.Recovering("QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessagesWritten(ptr.Pointer())
	}
}

//export callbackQNearFieldTarget_Type
func callbackQNearFieldTarget_Type(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNearFieldTarget::type")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "type"); signal != nil {
		return C.longlong(signal.(func() QNearFieldTarget__Type)())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectType(f func() QNearFieldTarget__Type) {
	defer qt.Recovering("connect QNearFieldTarget::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "type", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectType() {
	defer qt.Recovering("disconnect QNearFieldTarget::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "type")
	}
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	defer qt.Recovering("QNearFieldTarget::type")

	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldTarget_Uid
func callbackQNearFieldTarget_Uid(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QNearFieldTarget::uid")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "uid"); signal != nil {
		return C.CString(hex.EncodeToString([]byte(signal.(func() string)())))
	}

	return C.CString(hex.EncodeToString([]byte("")))
}

func (ptr *QNearFieldTarget) ConnectUid(f func() string) {
	defer qt.Recovering("connect QNearFieldTarget::uid")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "uid", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectUid() {
	defer qt.Recovering("disconnect QNearFieldTarget::uid")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "uid")
	}
}

func (ptr *QNearFieldTarget) Uid() string {
	defer qt.Recovering("QNearFieldTarget::uid")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNearFieldTarget_Uid(ptr.Pointer())))
	}
	return ""
}

//export callbackQNearFieldTarget_Url
func callbackQNearFieldTarget_Url(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNearFieldTarget::url")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "url"); signal != nil {
		return core.PointerFromQUrl(signal.(func() *core.QUrl)())
	}

	return core.PointerFromQUrl(NewQNearFieldTargetFromPointer(ptr).UrlDefault())
}

func (ptr *QNearFieldTarget) ConnectUrl(f func() *core.QUrl) {
	defer qt.Recovering("connect QNearFieldTarget::url")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "url", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectUrl() {
	defer qt.Recovering("disconnect QNearFieldTarget::url")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "url")
	}
}

func (ptr *QNearFieldTarget) Url() *core.QUrl {
	defer qt.Recovering("QNearFieldTarget::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNearFieldTarget_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) UrlDefault() *core.QUrl {
	defer qt.Recovering("QNearFieldTarget::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNearFieldTarget_UrlDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	defer qt.Recovering("QNearFieldTarget::~QNearFieldTarget")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()))
		C.QNearFieldTarget_DestroyQNearFieldTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldTarget_TimerEvent
func callbackQNearFieldTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QNearFieldTarget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNearFieldTarget_ChildEvent
func callbackQNearFieldTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QNearFieldTarget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldTarget_ConnectNotify
func callbackQNearFieldTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldTarget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNearFieldTarget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QNearFieldTarget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldTarget::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldTarget::connectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_CustomEvent
func callbackQNearFieldTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QNearFieldTarget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldTarget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldTarget_DeleteLater
func callbackQNearFieldTarget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNearFieldTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldTarget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNearFieldTarget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNearFieldTarget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QNearFieldTarget) DeleteLater() {
	defer qt.Recovering("QNearFieldTarget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()))
		C.QNearFieldTarget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) DeleteLaterDefault() {
	defer qt.Recovering("QNearFieldTarget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()))
		C.QNearFieldTarget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNearFieldTarget_DisconnectNotify
func callbackQNearFieldTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNearFieldTarget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNearFieldTarget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QNearFieldTarget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldTarget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNearFieldTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNearFieldTarget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_Event
func callbackQNearFieldTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldTarget::event")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldTarget) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldTarget::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QNearFieldTarget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldTarget::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldTarget::event")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNearFieldTarget_EventFilter
func callbackQNearFieldTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNearFieldTarget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldTarget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNearFieldTarget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNearFieldTarget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QNearFieldTarget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldTarget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNearFieldTarget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNearFieldTarget_MetaObject
func callbackQNearFieldTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNearFieldTarget::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNearFieldTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldTarget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNearFieldTarget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNearFieldTarget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QNearFieldTarget(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QNearFieldTarget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNearFieldTarget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldTarget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldTarget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNearFieldTarget::metaObject")

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
	defer qt.Recovering("QQmlNdefRecord::typeNameFormat")

	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
}

func (ptr *QQmlNdefRecord) Record() *QNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNdefRecordFromPointer(C.QQmlNdefRecord_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
		return tmpValue
	}
	return nil
}

//export callbackQQmlNdefRecord_RecordChanged
func callbackQQmlNdefRecord_RecordChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::recordChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "recordChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "recordChanged")
	}
}

func (ptr *QQmlNdefRecord) RecordChanged() {
	defer qt.Recovering("QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_RecordChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {
	defer qt.Recovering("QQmlNdefRecord::setRecord")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(ptr.Pointer(), PointerFromQNdefRecord(record))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	defer qt.Recovering("QQmlNdefRecord::setType")

	if ptr.Pointer() != nil {
		var newtypeC = C.CString(newtype)
		defer C.free(unsafe.Pointer(newtypeC))
		C.QQmlNdefRecord_SetType(ptr.Pointer(), newtypeC)
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	defer qt.Recovering("QQmlNdefRecord::setTypeNameFormat")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	defer qt.Recovering("QQmlNdefRecord::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

//export callbackQQmlNdefRecord_TypeChanged
func callbackQQmlNdefRecord_TypeChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::typeChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "typeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "typeChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeChanged() {
	defer qt.Recovering("QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeChanged(ptr.Pointer())
	}
}

//export callbackQQmlNdefRecord_TypeNameFormatChanged
func callbackQQmlNdefRecord_TypeNameFormatChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::typeNameFormatChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "typeNameFormatChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "typeNameFormatChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeNameFormatChanged() {
	defer qt.Recovering("QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeNameFormatChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecord() {
	defer qt.Recovering("QQmlNdefRecord::~QQmlNdefRecord")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()))
		C.QQmlNdefRecord_DestroyQQmlNdefRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlNdefRecord_TimerEvent
func callbackQQmlNdefRecord_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QQmlNdefRecord) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlNdefRecord) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlNdefRecord_ChildEvent
func callbackQQmlNdefRecord_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QQmlNdefRecord) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlNdefRecord) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlNdefRecord_ConnectNotify
func callbackQQmlNdefRecord_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlNdefRecord::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlNdefRecord::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QQmlNdefRecord) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlNdefRecord::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlNdefRecord::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_CustomEvent
func callbackQQmlNdefRecord_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QQmlNdefRecord) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlNdefRecord) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlNdefRecord_DeleteLater
func callbackQQmlNdefRecord_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlNdefRecord) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlNdefRecord::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QQmlNdefRecord) DeleteLater() {
	defer qt.Recovering("QQmlNdefRecord::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()))
		C.QQmlNdefRecord_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNdefRecord) DeleteLaterDefault() {
	defer qt.Recovering("QQmlNdefRecord::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()))
		C.QQmlNdefRecord_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlNdefRecord_DisconnectNotify
func callbackQQmlNdefRecord_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlNdefRecord::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlNdefRecord::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QQmlNdefRecord) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlNdefRecord::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlNdefRecord) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlNdefRecord::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_Event
func callbackQQmlNdefRecord_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlNdefRecord::event")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlNdefRecord) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlNdefRecord::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QQmlNdefRecord) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlNdefRecord::event")

	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlNdefRecord) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlNdefRecord::event")

	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_EventFilter
func callbackQQmlNdefRecord_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlNdefRecord::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlNdefRecord) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlNdefRecord::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlNdefRecord::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QQmlNdefRecord) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlNdefRecord::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlNdefRecord) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlNdefRecord::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlNdefRecord_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_MetaObject
func callbackQQmlNdefRecord_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlNdefRecord::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlNdefRecordFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlNdefRecord) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlNdefRecord::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlNdefRecord::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QQmlNdefRecord(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QQmlNdefRecord) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlNdefRecord::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlNdefRecord_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlNdefRecord) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlNdefRecord::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlNdefRecord_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
