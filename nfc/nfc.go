package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNdefFilter struct {
	ptr unsafe.Pointer
}

type QNdefFilter_ITF interface {
	QNdefFilter_PTR() *QNdefFilter
}

func (p *QNdefFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func newQNdefFilterFromPointer(ptr unsafe.Pointer) *QNdefFilter {
	var n = NewQNdefFilterFromPointer(ptr)
	return n
}

func (ptr *QNdefFilter) QNdefFilter_PTR() *QNdefFilter {
	return ptr
}

func NewQNdefFilter() *QNdefFilter {
	defer qt.Recovering("QNdefFilter::QNdefFilter")

	return newQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter())
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {
	defer qt.Recovering("QNdefFilter::QNdefFilter")

	return newQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter2(PointerFromQNdefFilter(other)))
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
		return int(C.QNdefFilter_RecordCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	defer qt.Recovering("QNdefFilter::setOrderMatch")

	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	defer qt.Recovering("QNdefFilter::~QNdefFilter")

	if ptr.Pointer() != nil {
		C.QNdefFilter_DestroyQNdefFilter(ptr.Pointer())
	}
}

type QNdefMessage struct {
	core.QList
}

type QNdefMessage_ITF interface {
	core.QList_ITF
	QNdefMessage_PTR() *QNdefMessage
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

func newQNdefMessageFromPointer(ptr unsafe.Pointer) *QNdefMessage {
	var n = NewQNdefMessageFromPointer(ptr)
	return n
}

func (ptr *QNdefMessage) QNdefMessage_PTR() *QNdefMessage {
	return ptr
}

func NewQNdefMessage() *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return newQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage())
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return newQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage3(PointerFromQNdefMessage(message)))
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return newQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage2(PointerFromQNdefRecord(record)))
}

func (ptr *QNdefMessage) ToByteArray() *core.QByteArray {
	defer qt.Recovering("QNdefMessage::toByteArray")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefMessage_ToByteArray(ptr.Pointer()))
	}
	return nil
}

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord
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

func newQNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcSmartPosterRecord {
	var n = NewQNdefNfcSmartPosterRecordFromPointer(ptr)
	return n
}

func (ptr *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord {
	return ptr
}

//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int64

const (
	QNdefNfcSmartPosterRecord__UnspecifiedAction = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        = QNdefNfcSmartPosterRecord__Action(2)
)

func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	return newQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord())
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	return newQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(PointerFromQNdefNfcSmartPosterRecord(other)))
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")

	return newQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::action")

	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArray_ITF, data core.QByteArray_ITF) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::addIcon")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty), core.PointerFromQByteArray(data))
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
		return C.QNdefNfcSmartPosterRecord_AddTitle2(ptr.Pointer(), C.CString(text), C.CString(locale), C.int(encoding)) != 0
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

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArray_ITF) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::hasIcon")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasIcon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)) != 0
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
		return C.QNdefNfcSmartPosterRecord_HasTitle(ptr.Pointer(), C.CString(locale)) != 0
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

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::icon")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_Icon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)))
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::iconCount")

	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_IconCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArray_ITF) bool {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::removeIcon")

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty)) != 0
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
		return C.QNdefNfcSmartPosterRecord_RemoveTitle2(ptr.Pointer(), C.CString(locale)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setAction")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(ptr.Pointer(), C.int(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArray_ITF) {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::setTypeInfo")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(ptr.Pointer(), core.PointerFromQByteArray(ty))
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

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcSmartPosterRecord_Title(ptr.Pointer(), C.CString(locale)))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::titleCount")

	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_TitleCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() *core.QByteArray {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::typeInfo")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_TypeInfo(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() *core.QUrl {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::uri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNdefNfcSmartPosterRecord_Uri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	defer qt.Recovering("QNdefNfcSmartPosterRecord::~QNdefNfcSmartPosterRecord")

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(ptr.Pointer())
	}
}

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord
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

func newQNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcTextRecord {
	var n = NewQNdefNfcTextRecordFromPointer(ptr)
	return n
}

func (ptr *QNdefNfcTextRecord) QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord {
	return ptr
}

//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int64

const (
	QNdefNfcTextRecord__Utf8  = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 = QNdefNfcTextRecord__Encoding(1)
)

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {
	defer qt.Recovering("QNdefNfcTextRecord::QNdefNfcTextRecord")

	return newQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord())
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {
	defer qt.Recovering("QNdefNfcTextRecord::QNdefNfcTextRecord")

	return newQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(PointerFromQNdefRecord(other)))
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
		C.QNdefNfcTextRecord_SetEncoding(ptr.Pointer(), C.int(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	defer qt.Recovering("QNdefNfcTextRecord::setLocale")

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetLocale(ptr.Pointer(), C.CString(locale))
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	defer qt.Recovering("QNdefNfcTextRecord::setText")

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetText(ptr.Pointer(), C.CString(text))
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

func newQNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcUriRecord {
	var n = NewQNdefNfcUriRecordFromPointer(ptr)
	return n
}

func (ptr *QNdefNfcUriRecord) QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord {
	return ptr
}

func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {
	defer qt.Recovering("QNdefNfcUriRecord::QNdefNfcUriRecord")

	return newQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord())
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {
	defer qt.Recovering("QNdefNfcUriRecord::QNdefNfcUriRecord")

	return newQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(PointerFromQNdefRecord(other)))
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
		return core.NewQUrlFromPointer(C.QNdefNfcUriRecord_Uri(ptr.Pointer()))
	}
	return nil
}

type QNdefRecord struct {
	ptr unsafe.Pointer
}

type QNdefRecord_ITF interface {
	QNdefRecord_PTR() *QNdefRecord
}

func (p *QNdefRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func newQNdefRecordFromPointer(ptr unsafe.Pointer) *QNdefRecord {
	var n = NewQNdefRecordFromPointer(ptr)
	return n
}

func (ptr *QNdefRecord) QNdefRecord_PTR() *QNdefRecord {
	return ptr
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

func NewQNdefRecord() *QNdefRecord {
	defer qt.Recovering("QNdefRecord::QNdefRecord")

	return newQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord())
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {
	defer qt.Recovering("QNdefRecord::QNdefRecord")

	return newQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefRecord) Id() *core.QByteArray {
	defer qt.Recovering("QNdefRecord::id")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Id(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefRecord) IsEmpty() bool {
	defer qt.Recovering("QNdefRecord::isEmpty")

	if ptr.Pointer() != nil {
		return C.QNdefRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefRecord) Payload() *core.QByteArray {
	defer qt.Recovering("QNdefRecord::payload")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Payload(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefRecord) SetId(id core.QByteArray_ITF) {
	defer qt.Recovering("QNdefRecord::setId")

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetId(ptr.Pointer(), core.PointerFromQByteArray(id))
	}
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArray_ITF) {
	defer qt.Recovering("QNdefRecord::setPayload")

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetPayload(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QNdefRecord) SetType(ty core.QByteArray_ITF) {
	defer qt.Recovering("QNdefRecord::setType")

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetType(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	defer qt.Recovering("QNdefRecord::setTypeNameFormat")

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.int(typeNameFormat))
	}
}

func (ptr *QNdefRecord) Type() *core.QByteArray {
	defer qt.Recovering("QNdefRecord::type")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Type(ptr.Pointer()))
	}
	return nil
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
	}
}

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
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

func newQNearFieldManagerFromPointer(ptr unsafe.Pointer) *QNearFieldManager {
	var n = NewQNearFieldManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldManager_") {
		n.SetObjectName("QNearFieldManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return ptr
}

//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess = QNearFieldManager__TargetAccessMode(0x04)
)

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), C.CString(method)))
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

	return newQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	defer qt.Recovering("QNearFieldManager::isAvailable")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.int(typeNameFormat), core.PointerFromQByteArray(ty), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	defer qt.Recovering("QNearFieldManager::setTargetAccessModes")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.int(accessModes))
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

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldManagerTargetDetected
func callbackQNearFieldManagerTargetDetected(ptr unsafe.Pointer, ptrName *C.char, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetDetected")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {
	defer qt.Recovering("QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	defer qt.Recovering("disconnect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetLost")
	}
}

//export callbackQNearFieldManagerTargetLost
func callbackQNearFieldManagerTargetLost(ptr unsafe.Pointer, ptrName *C.char, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetLost")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetLost"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
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
		return C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(handlerId)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	defer qt.Recovering("QNearFieldManager::~QNearFieldManager")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldManagerTimerEvent
func callbackQNearFieldManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

func (ptr *QNearFieldManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldManagerChildEvent
func callbackQNearFieldManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

func (ptr *QNearFieldManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldManagerCustomEvent
func callbackQNearFieldManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
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

func newQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = NewQNearFieldShareManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareManager_") {
		n.SetObjectName("QNearFieldShareManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
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

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	defer qt.Recovering("QNearFieldShareManager::QNearFieldShareManager")

	return newQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareManagerError
func callbackQNearFieldShareManagerError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::setShareModes")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.int(mode))
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

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	defer qt.Recovering("connect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	defer qt.Recovering("disconnect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareModesChanged")
	}
}

//export callbackQNearFieldShareManagerShareModesChanged
func callbackQNearFieldShareManagerShareModesChanged(ptr unsafe.Pointer, ptrName *C.char, modes C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::shareModesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shareModesChanged"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
	}

}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ShareModesChanged(ptr.Pointer(), C.int(modes))
	}
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::supportedShareModes")

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	defer qt.Recovering("connect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldShareManagerTargetDetected
func callbackQNearFieldShareManagerTargetDetected(ptr unsafe.Pointer, ptrName *C.char, shareTarget unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::targetDetected")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
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
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldShareManagerTimerEvent
func callbackQNearFieldShareManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

func (ptr *QNearFieldShareManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldShareManagerChildEvent
func callbackQNearFieldShareManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

func (ptr *QNearFieldShareManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldShareManagerCustomEvent
func callbackQNearFieldShareManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
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

func newQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) *QNearFieldShareTarget {
	var n = NewQNearFieldShareTargetFromPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareTarget_") {
		n.SetObjectName("QNearFieldShareTarget_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Cancel() {
	defer qt.Recovering("QNearFieldShareTarget::cancel")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareTargetError
func callbackQNearFieldShareTargetError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNearFieldShareTarget::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Error(ptr.Pointer(), C.int(error))
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

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	defer qt.Recovering("connect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareFinished")
	}
}

//export callbackQNearFieldShareTargetShareFinished
func callbackQNearFieldShareTargetShareFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNearFieldShareTarget::shareFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "shareFinished"); signal != nil {
		signal.(func())()
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
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldShareTargetTimerEvent
func callbackQNearFieldShareTargetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

func (ptr *QNearFieldShareTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldShareTargetChildEvent
func callbackQNearFieldShareTargetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

func (ptr *QNearFieldShareTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldShareTargetCustomEvent
func callbackQNearFieldShareTargetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTarget_ITF interface {
	core.QObject_ITF
	QNearFieldTarget_PTR() *QNearFieldTarget
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

func newQNearFieldTargetFromPointer(ptr unsafe.Pointer) *QNearFieldTarget {
	var n = NewQNearFieldTargetFromPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldTarget_") {
		n.SetObjectName("QNearFieldTarget_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldTarget) QNearFieldTarget_PTR() *QNearFieldTarget {
	return ptr
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

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	defer qt.Recovering("QNearFieldTarget::accessMethods")

	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQNearFieldTargetDisconnected
func callbackQNearFieldTargetDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNearFieldTarget::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) Disconnected() {
	defer qt.Recovering("QNearFieldTarget::disconnected")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_Disconnected(ptr.Pointer())
	}
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	defer qt.Recovering("QNearFieldTarget::hasNdefMessage")

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessage(ptr.Pointer()) != 0
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

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	defer qt.Recovering("connect QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessagesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ndefMessagesWritten", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	defer qt.Recovering("disconnect QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ndefMessagesWritten")
	}
}

//export callbackQNearFieldTargetNdefMessagesWritten
func callbackQNearFieldTargetNdefMessagesWritten(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNearFieldTarget::ndefMessagesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "ndefMessagesWritten"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldTarget) NdefMessagesWritten() {
	defer qt.Recovering("QNearFieldTarget::ndefMessagesWritten")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessagesWritten(ptr.Pointer())
	}
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	defer qt.Recovering("QNearFieldTarget::type")

	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldTarget) Uid() *core.QByteArray {
	defer qt.Recovering("QNearFieldTarget::uid")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNearFieldTarget_Uid(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldTarget) Url() *core.QUrl {
	defer qt.Recovering("QNearFieldTarget::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNearFieldTarget_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	defer qt.Recovering("QNearFieldTarget::~QNearFieldTarget")

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DestroyQNearFieldTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldTargetTimerEvent
func callbackQNearFieldTargetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

func (ptr *QNearFieldTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldTargetChildEvent
func callbackQNearFieldTargetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

func (ptr *QNearFieldTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldTargetCustomEvent
func callbackQNearFieldTargetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldTarget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QQmlNdefRecord struct {
	core.QObject
}

type QQmlNdefRecord_ITF interface {
	core.QObject_ITF
	QQmlNdefRecord_PTR() *QQmlNdefRecord
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

func newQQmlNdefRecordFromPointer(ptr unsafe.Pointer) *QQmlNdefRecord {
	var n = NewQQmlNdefRecordFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlNdefRecord_") {
		n.SetObjectName("QQmlNdefRecord_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlNdefRecord) QQmlNdefRecord_PTR() *QQmlNdefRecord {
	return ptr
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

func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {
	defer qt.Recovering("QQmlNdefRecord::typeNameFormat")

	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return newQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return newQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "recordChanged")
	}
}

//export callbackQQmlNdefRecordRecordChanged
func callbackQQmlNdefRecordRecordChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::recordChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "recordChanged"); signal != nil {
		signal.(func())()
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
		C.QQmlNdefRecord_SetType(ptr.Pointer(), C.CString(newtype))
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	defer qt.Recovering("QQmlNdefRecord::setTypeNameFormat")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.int(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	defer qt.Recovering("QQmlNdefRecord::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQQmlNdefRecordTypeChanged
func callbackQQmlNdefRecordTypeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::typeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) TypeChanged() {
	defer qt.Recovering("QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeNameFormatChanged")
	}
}

//export callbackQQmlNdefRecordTypeNameFormatChanged
func callbackQQmlNdefRecordTypeNameFormatChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::typeNameFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeNameFormatChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) TypeNameFormatChanged() {
	defer qt.Recovering("QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeNameFormatChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlNdefRecordTimerEvent
func callbackQQmlNdefRecordTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

func (ptr *QQmlNdefRecord) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlNdefRecordChildEvent
func callbackQQmlNdefRecordChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

func (ptr *QQmlNdefRecord) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlNdefRecordCustomEvent
func callbackQQmlNdefRecordCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlNdefRecord::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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
