// +build !minimal

package nfc

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"unsafe"
)

type QNdefFilter struct {
	internal.Internal
}

type QNdefFilter_ITF interface {
	QNdefFilter_PTR() *QNdefFilter
}

func (ptr *QNdefFilter) QNdefFilter_PTR() *QNdefFilter {
	return ptr
}

func (ptr *QNdefFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNdefFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNdefFilter(ptr QNdefFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefFilter_PTR().Pointer()
	}
	return nil
}

func (n *QNdefFilter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNdefFilterFromPointer(ptr unsafe.Pointer) (n *QNdefFilter) {
	n = new(QNdefFilter)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefFilter")
	return
}
func NewQNdefFilter() *QNdefFilter {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefFilter", ""}).(*QNdefFilter)
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefFilter2", "", other}).(*QNdefFilter)
}

func (ptr *QNdefFilter) AppendRecord2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, min uint, max uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AppendRecord2", typeNameFormat, ty, min, max})
}

func (ptr *QNdefFilter) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QNdefFilter) OrderMatch() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrderMatch"}).(bool)
}

func (ptr *QNdefFilter) RecordCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecordCount"}).(float64))
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrderMatch", on})
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNdefFilter"})
}

type QNdefMessage struct {
	internal.Internal
}

type QNdefMessage_ITF interface {
	QNdefMessage_PTR() *QNdefMessage
}

func (ptr *QNdefMessage) QNdefMessage_PTR() *QNdefMessage {
	return ptr
}

func (ptr *QNdefMessage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNdefMessage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNdefMessage(ptr QNdefMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefMessage_PTR().Pointer()
	}
	return nil
}

func (n *QNdefMessage) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNdefMessageFromPointer(ptr unsafe.Pointer) (n *QNdefMessage) {
	n = new(QNdefMessage)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefMessage")
	return
}

func (ptr *QNdefMessage) DestroyQNdefMessage() {
}

func NewQNdefMessage() *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefMessage", ""}).(*QNdefMessage)
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefMessage2", "", record}).(*QNdefMessage)
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefMessage3", "", message}).(*QNdefMessage)
}

func NewQNdefMessage4(records []*QNdefRecord) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefMessage4", "", records}).(*QNdefMessage)
}

func QNdefMessage_FromByteArray(message core.QByteArray_ITF) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.QNdefMessage_FromByteArray", "", message}).(*QNdefMessage)
}

func (ptr *QNdefMessage) FromByteArray(message core.QByteArray_ITF) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.QNdefMessage_FromByteArray", "", message}).(*QNdefMessage)
}

func (ptr *QNdefMessage) ToByteArray() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToByteArray"}).(*core.QByteArray)
}

func (ptr *QNdefMessage) __QNdefMessage_records_atList4(i int) *QNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QNdefMessage_records_atList4", i}).(*QNdefRecord)
}

func (ptr *QNdefMessage) __QNdefMessage_records_setList4(i QNdefRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QNdefMessage_records_setList4", i})
}

func (ptr *QNdefMessage) __QNdefMessage_records_newList4() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QNdefMessage_records_newList4"}).(unsafe.Pointer)
}

type QNdefNfcIconRecord struct {
	QNdefRecord
}

type QNdefNfcIconRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcIconRecord_PTR() *QNdefNfcIconRecord
}

func (ptr *QNdefNfcIconRecord) QNdefNfcIconRecord_PTR() *QNdefNfcIconRecord {
	return ptr
}

func (ptr *QNdefNfcIconRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcIconRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcIconRecord(ptr QNdefNfcIconRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcIconRecord_PTR().Pointer()
	}
	return nil
}

func (n *QNdefNfcIconRecord) InitFromInternal(ptr uintptr, name string) {
	n.QNdefRecord_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNdefNfcIconRecord) ClassNameInternalF() string {
	return n.QNdefRecord_PTR().ClassNameInternalF()
}

func NewQNdefNfcIconRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcIconRecord) {
	n = new(QNdefNfcIconRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefNfcIconRecord")
	return
}

func (ptr *QNdefNfcIconRecord) DestroyQNdefNfcIconRecord() {
}

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord
}

func (ptr *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord {
	return ptr
}

func (ptr *QNdefNfcSmartPosterRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcSmartPosterRecord(ptr QNdefNfcSmartPosterRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcSmartPosterRecord_PTR().Pointer()
	}
	return nil
}

func (n *QNdefNfcSmartPosterRecord) InitFromInternal(ptr uintptr, name string) {
	n.QNdefRecord_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNdefNfcSmartPosterRecord) ClassNameInternalF() string {
	return n.QNdefRecord_PTR().ClassNameInternalF()
}

func NewQNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcSmartPosterRecord) {
	n = new(QNdefNfcSmartPosterRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefNfcSmartPosterRecord")
	return
}

//go:generate stringer -type=QNdefNfcSmartPosterRecord__Action
//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int64

const (
	QNdefNfcSmartPosterRecord__UnspecifiedAction QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(2)
)

func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcSmartPosterRecord", ""}).(*QNdefNfcSmartPosterRecord)
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcSmartPosterRecord2", "", other}).(*QNdefNfcSmartPosterRecord)
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcSmartPosterRecord3", "", other}).(*QNdefNfcSmartPosterRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {

	return QNdefNfcSmartPosterRecord__Action(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Action"}).(float64))
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon(icon QNdefNfcIconRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddIcon", icon})
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArray_ITF, data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddIcon2", ty, data})
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddTitle", text}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddTitle2", text, locale, encoding}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) HasAction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAction"}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasIcon", mimetype}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasSize"}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasTitle", locale}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasTypeInfo"}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon", mimetype}).(*core.QByteArray)
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconCount"}).(float64))
}

func (ptr *QNdefNfcSmartPosterRecord) IconRecord(index int) *QNdefNfcIconRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconRecord", index}).(*QNdefNfcIconRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) IconRecords() []*QNdefNfcIconRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconRecords"}).([]*QNdefNfcIconRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon(icon QNdefNfcIconRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveIcon", icon}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveIcon2", ty}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveTitle", text}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveTitle2", locale}).(bool)
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAction", act})
}

func (ptr *QNdefNfcSmartPosterRecord) SetIcons(icons []*QNdefNfcIconRecord) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIcons", icons})
}

func (ptr *QNdefNfcSmartPosterRecord) SetSize(size uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSize", size})
}

func (ptr *QNdefNfcSmartPosterRecord) SetTitles(titles []*QNdefNfcTextRecord) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitles", titles})
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTypeInfo", ty})
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUri", url})
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUri2", url})
}

func (ptr *QNdefNfcSmartPosterRecord) Size() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title", locale}).(string)
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleCount"}).(float64))
}

func (ptr *QNdefNfcSmartPosterRecord) TitleRecord(index int) *QNdefNfcTextRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleRecord", index}).(*QNdefNfcTextRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeInfo"}).(*core.QByteArray)
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uri"}).(*core.QUrl)
}

func (ptr *QNdefNfcSmartPosterRecord) UriRecord() *QNdefNfcUriRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UriRecord"}).(*QNdefNfcUriRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNdefNfcSmartPosterRecord"})
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_atList(i int) *QNdefNfcIconRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__iconRecords_atList", i}).(*QNdefNfcIconRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_setList(i QNdefNfcIconRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__iconRecords_setList", i})
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__iconRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_atList(i int) *QNdefNfcIconRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIcons_icons_atList", i}).(*QNdefNfcIconRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_setList(i QNdefNfcIconRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIcons_icons_setList", i})
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIcons_icons_newList"}).(unsafe.Pointer)
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_atList(i int) *QNdefNfcTextRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTitles_titles_atList", i}).(*QNdefNfcTextRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_setList(i QNdefNfcTextRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTitles_titles_setList", i})
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTitles_titles_newList"}).(unsafe.Pointer)
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_atList(i int) *QNdefNfcTextRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__titleRecords_atList", i}).(*QNdefNfcTextRecord)
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_setList(i QNdefNfcTextRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__titleRecords_setList", i})
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__titleRecords_newList"}).(unsafe.Pointer)
}

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord
}

func (ptr *QNdefNfcTextRecord) QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord {
	return ptr
}

func (ptr *QNdefNfcTextRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcTextRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcTextRecord(ptr QNdefNfcTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcTextRecord_PTR().Pointer()
	}
	return nil
}

func (n *QNdefNfcTextRecord) InitFromInternal(ptr uintptr, name string) {
	n.QNdefRecord_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNdefNfcTextRecord) ClassNameInternalF() string {
	return n.QNdefRecord_PTR().ClassNameInternalF()
}

func NewQNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcTextRecord) {
	n = new(QNdefNfcTextRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefNfcTextRecord")
	return
}

func (ptr *QNdefNfcTextRecord) DestroyQNdefNfcTextRecord() {
}

//go:generate stringer -type=QNdefNfcTextRecord__Encoding
//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int64

const (
	QNdefNfcTextRecord__Utf8  QNdefNfcTextRecord__Encoding = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 QNdefNfcTextRecord__Encoding = QNdefNfcTextRecord__Encoding(1)
)

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcTextRecord", ""}).(*QNdefNfcTextRecord)
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcTextRecord2", "", other}).(*QNdefNfcTextRecord)
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {

	return QNdefNfcTextRecord__Encoding(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Encoding"}).(float64))
}

func (ptr *QNdefNfcTextRecord) Locale() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(string)
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEncoding", encoding})
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetText", text})
}

func (ptr *QNdefNfcTextRecord) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

type QNdefNfcUriRecord struct {
	QNdefRecord
}

type QNdefNfcUriRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord
}

func (ptr *QNdefNfcUriRecord) QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord {
	return ptr
}

func (ptr *QNdefNfcUriRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcUriRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcUriRecord(ptr QNdefNfcUriRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcUriRecord_PTR().Pointer()
	}
	return nil
}

func (n *QNdefNfcUriRecord) InitFromInternal(ptr uintptr, name string) {
	n.QNdefRecord_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNdefNfcUriRecord) ClassNameInternalF() string {
	return n.QNdefRecord_PTR().ClassNameInternalF()
}

func NewQNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcUriRecord) {
	n = new(QNdefNfcUriRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefNfcUriRecord")
	return
}

func (ptr *QNdefNfcUriRecord) DestroyQNdefNfcUriRecord() {
}

func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcUriRecord", ""}).(*QNdefNfcUriRecord)
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefNfcUriRecord2", "", other}).(*QNdefNfcUriRecord)
}

func (ptr *QNdefNfcUriRecord) SetUri(uri core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUri", uri})
}

func (ptr *QNdefNfcUriRecord) Uri() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uri"}).(*core.QUrl)
}

type QNdefRecord struct {
	internal.Internal
}

type QNdefRecord_ITF interface {
	QNdefRecord_PTR() *QNdefRecord
}

func (ptr *QNdefRecord) QNdefRecord_PTR() *QNdefRecord {
	return ptr
}

func (ptr *QNdefRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNdefRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNdefRecord(ptr QNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (n *QNdefRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNdefRecordFromPointer(ptr unsafe.Pointer) (n *QNdefRecord) {
	n = new(QNdefRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QNdefRecord")
	return
}

//go:generate stringer -type=QNdefRecord__TypeNameFormat
//QNdefRecord::TypeNameFormat
type QNdefRecord__TypeNameFormat int64

const (
	QNdefRecord__Empty       QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x00)
	QNdefRecord__NfcRtd      QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x01)
	QNdefRecord__Mime        QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x02)
	QNdefRecord__Uri         QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x03)
	QNdefRecord__ExternalRtd QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x04)
	QNdefRecord__Unknown     QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x05)
)

func NewQNdefRecord() *QNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefRecord", ""}).(*QNdefRecord)
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNdefRecord2", "", other}).(*QNdefRecord)
}

func (ptr *QNdefRecord) Id() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Id"}).(*core.QByteArray)
}

func (ptr *QNdefRecord) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QNdefRecord) Payload() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Payload"}).(*core.QByteArray)
}

func (ptr *QNdefRecord) SetId(id core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetId", id})
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPayload", payload})
}

func (ptr *QNdefRecord) SetType(ty core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", ty})
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTypeNameFormat", typeNameFormat})
}

func (ptr *QNdefRecord) Type() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*core.QByteArray)
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {

	return QNdefRecord__TypeNameFormat(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeNameFormat"}).(float64))
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNdefRecord"})
}

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
}

func (ptr *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return ptr
}

func (ptr *QNearFieldManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldManager(ptr QNearFieldManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManager_PTR().Pointer()
	}
	return nil
}

func (n *QNearFieldManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNearFieldManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNearFieldManagerFromPointer(ptr unsafe.Pointer) (n *QNearFieldManager) {
	n = new(QNearFieldManager)
	n.InitFromInternal(uintptr(ptr), "nfc.QNearFieldManager")
	return
}

//go:generate stringer -type=QNearFieldManager__AdapterState
//QNearFieldManager::AdapterState
type QNearFieldManager__AdapterState int64

const (
	QNearFieldManager__Offline    QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(1)
	QNearFieldManager__TurningOn  QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(2)
	QNearFieldManager__Online     QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(3)
	QNearFieldManager__TurningOff QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(4)
)

//go:generate stringer -type=QNearFieldManager__TargetAccessMode
//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x04)
)

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNearFieldManager", "", parent}).(*QNearFieldManager)
}

func (ptr *QNearFieldManager) IsAvailable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAvailable"}).(bool)
}

func (ptr *QNearFieldManager) IsSupported() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSupported"}).(bool)
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterNdefMessageHandler", object, method}).(float64))
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterNdefMessageHandler2", typeNameFormat, ty, object, method}).(float64))
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterNdefMessageHandler3", filter, object, method}).(float64))
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTargetAccessModes", accessModes})
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartTargetDetection"}).(bool)
}

func (ptr *QNearFieldManager) StopTargetDetection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopTargetDetection"})
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {

	return QNearFieldManager__TargetAccessMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TargetAccessModes"}).(float64))
}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTargetDetected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTargetDetected"})
}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TargetDetected", target})
}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTargetLost", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTargetLost"})
}

func (ptr *QNearFieldManager) TargetLost(target QNearFieldTarget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TargetLost", target})
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterNdefMessageHandler", handlerId}).(bool)
}

func (ptr *QNearFieldManager) ConnectDestroyQNearFieldManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNearFieldManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldManager) DisconnectDestroyQNearFieldManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNearFieldManager"})
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldManager"})
}

func (ptr *QNearFieldManager) DestroyQNearFieldManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldManagerDefault"})
}

func (ptr *QNearFieldManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNearFieldManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNearFieldManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNearFieldManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNearFieldManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNearFieldManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNearFieldManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNearFieldManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNearFieldManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNearFieldManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNearFieldManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNearFieldManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNearFieldManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNearFieldManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
}

func (ptr *QNearFieldShareManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldShareManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func (n *QNearFieldShareManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNearFieldShareManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) (n *QNearFieldShareManager) {
	n = new(QNearFieldShareManager)
	n.InitFromInternal(uintptr(ptr), "nfc.QNearFieldShareManager")
	return
}

//go:generate stringer -type=QNearFieldShareManager__ShareMode
//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x02)
)

//go:generate stringer -type=QNearFieldShareManager__ShareError
//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
	QNearFieldShareManager__NoError                     QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(8)
)

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNearFieldShareManager", "", parent}).(*QNearFieldShareManager)
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareManager) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", error})
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShareModes", mode})
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {

	return QNearFieldShareManager__ShareError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareError"}).(float64))
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {

	return QNearFieldShareManager__ShareMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareModes"}).(float64))
}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShareModesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShareModesChanged"})
}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareModesChanged", modes})
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {

	return QNearFieldShareManager__ShareMode(internal.CallLocalFunction([]interface{}{"", "", "nfc.QNearFieldShareManager_SupportedShareModes", ""}).(float64))
}

func (ptr *QNearFieldShareManager) SupportedShareModes() QNearFieldShareManager__ShareMode {

	return QNearFieldShareManager__ShareMode(internal.CallLocalFunction([]interface{}{"", "", "nfc.QNearFieldShareManager_SupportedShareModes", ""}).(float64))
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTargetDetected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTargetDetected"})
}

func (ptr *QNearFieldShareManager) TargetDetected(shareTarget QNearFieldShareTarget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TargetDetected", shareTarget})
}

func (ptr *QNearFieldShareManager) ConnectDestroyQNearFieldShareManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNearFieldShareManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareManager) DisconnectDestroyQNearFieldShareManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNearFieldShareManager"})
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldShareManager"})
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldShareManagerDefault"})
}

func (ptr *QNearFieldShareManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldShareManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNearFieldShareManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldShareManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNearFieldShareManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNearFieldShareManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNearFieldShareManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNearFieldShareManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNearFieldShareManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNearFieldShareManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNearFieldShareManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNearFieldShareManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNearFieldShareManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNearFieldShareManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNearFieldShareManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldShareTarget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTarget_PTR().Pointer()
	}
	return nil
}

func (n *QNearFieldShareTarget) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNearFieldShareTarget) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) (n *QNearFieldShareTarget) {
	n = new(QNearFieldShareTarget)
	n.InitFromInternal(uintptr(ptr), "nfc.QNearFieldShareTarget")
	return
}
func (ptr *QNearFieldShareTarget) Cancel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Cancel"})
}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareTarget) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", error})
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsShareInProgress"}).(bool)
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Share", message}).(bool)
}

func (ptr *QNearFieldShareTarget) Share2(files []*core.QFileInfo) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Share2", files}).(bool)
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {

	return QNearFieldShareManager__ShareError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareError"}).(float64))
}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShareFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShareFinished"})
}

func (ptr *QNearFieldShareTarget) ShareFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareFinished"})
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {

	return QNearFieldShareManager__ShareMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShareModes"}).(float64))
}

func (ptr *QNearFieldShareTarget) ConnectDestroyQNearFieldShareTarget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNearFieldShareTarget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldShareTarget) DisconnectDestroyQNearFieldShareTarget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNearFieldShareTarget"})
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldShareTarget"})
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTargetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldShareTargetDefault"})
}

func (ptr *QNearFieldShareTarget) __share_files_atList2(i int) *core.QFileInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__share_files_atList2", i}).(*core.QFileInfo)
}

func (ptr *QNearFieldShareTarget) __share_files_setList2(i core.QFileInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__share_files_setList2", i})
}

func (ptr *QNearFieldShareTarget) __share_files_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__share_files_newList2"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareTarget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldShareTarget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNearFieldShareTarget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareTarget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldShareTarget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNearFieldShareTarget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareTarget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNearFieldShareTarget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNearFieldShareTarget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNearFieldShareTarget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNearFieldShareTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNearFieldShareTarget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNearFieldShareTarget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNearFieldShareTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNearFieldShareTarget) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNearFieldShareTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNearFieldShareTarget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNearFieldShareTarget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTarget_ITF interface {
	core.QObject_ITF
	QNearFieldTarget_PTR() *QNearFieldTarget
}

func (ptr *QNearFieldTarget) QNearFieldTarget_PTR() *QNearFieldTarget {
	return ptr
}

func (ptr *QNearFieldTarget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldTarget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldTarget(ptr QNearFieldTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldTarget_PTR().Pointer()
	}
	return nil
}

func (n *QNearFieldTarget) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNearFieldTarget) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNearFieldTargetFromPointer(ptr unsafe.Pointer) (n *QNearFieldTarget) {
	n = new(QNearFieldTarget)
	n.InitFromInternal(uintptr(ptr), "nfc.QNearFieldTarget")
	return
}

//go:generate stringer -type=QNearFieldTarget__Type
//QNearFieldTarget::Type
type QNearFieldTarget__Type int64

const (
	QNearFieldTarget__ProprietaryTag QNearFieldTarget__Type = QNearFieldTarget__Type(0)
	QNearFieldTarget__NfcTagType1    QNearFieldTarget__Type = QNearFieldTarget__Type(1)
	QNearFieldTarget__NfcTagType2    QNearFieldTarget__Type = QNearFieldTarget__Type(2)
	QNearFieldTarget__NfcTagType3    QNearFieldTarget__Type = QNearFieldTarget__Type(3)
	QNearFieldTarget__NfcTagType4    QNearFieldTarget__Type = QNearFieldTarget__Type(4)
	QNearFieldTarget__MifareTag      QNearFieldTarget__Type = QNearFieldTarget__Type(5)
)

//go:generate stringer -type=QNearFieldTarget__AccessMethod
//QNearFieldTarget::AccessMethod
type QNearFieldTarget__AccessMethod int64

const (
	QNearFieldTarget__UnknownAccess         QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x00)
	QNearFieldTarget__NdefAccess            QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x01)
	QNearFieldTarget__TagTypeSpecificAccess QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x02)
	QNearFieldTarget__LlcpAccess            QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x04)
)

//go:generate stringer -type=QNearFieldTarget__Error
//QNearFieldTarget::Error
type QNearFieldTarget__Error int64

const (
	QNearFieldTarget__NoError                QNearFieldTarget__Error = QNearFieldTarget__Error(0)
	QNearFieldTarget__UnknownError           QNearFieldTarget__Error = QNearFieldTarget__Error(1)
	QNearFieldTarget__UnsupportedError       QNearFieldTarget__Error = QNearFieldTarget__Error(2)
	QNearFieldTarget__TargetOutOfRangeError  QNearFieldTarget__Error = QNearFieldTarget__Error(3)
	QNearFieldTarget__NoResponseError        QNearFieldTarget__Error = QNearFieldTarget__Error(4)
	QNearFieldTarget__ChecksumMismatchError  QNearFieldTarget__Error = QNearFieldTarget__Error(5)
	QNearFieldTarget__InvalidParametersError QNearFieldTarget__Error = QNearFieldTarget__Error(6)
	QNearFieldTarget__NdefReadError          QNearFieldTarget__Error = QNearFieldTarget__Error(7)
	QNearFieldTarget__NdefWriteError         QNearFieldTarget__Error = QNearFieldTarget__Error(8)
	QNearFieldTarget__CommandError           QNearFieldTarget__Error = QNearFieldTarget__Error(9)
)

func NewQNearFieldTarget(parent core.QObject_ITF) *QNearFieldTarget {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQNearFieldTarget", "", parent}).(*QNearFieldTarget)
}

func (ptr *QNearFieldTarget) ConnectAccessMethods(f func() QNearFieldTarget__AccessMethod) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAccessMethods", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectAccessMethods() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAccessMethods"})
}

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {

	return QNearFieldTarget__AccessMethod(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AccessMethods"}).(float64))
}

func (ptr *QNearFieldTarget) Disconnect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnect"}).(bool)
}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QNearFieldTarget) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QNearFieldTarget) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QNearFieldTarget) ConnectHasNdefMessage(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasNdefMessage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectHasNdefMessage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasNdefMessage"})
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasNdefMessage"}).(bool)
}

func (ptr *QNearFieldTarget) HasNdefMessageDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasNdefMessageDefault"}).(bool)
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsProcessingCommand"}).(bool)
}

func (ptr *QNearFieldTarget) KeepConnection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeepConnection"}).(bool)
}

func (ptr *QNearFieldTarget) MaxCommandLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxCommandLength"}).(float64))
}

func (ptr *QNearFieldTarget) ConnectNdefMessageRead(f func(message *QNdefMessage)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNdefMessageRead", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectNdefMessageRead() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNdefMessageRead"})
}

func (ptr *QNearFieldTarget) NdefMessageRead(message QNdefMessage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NdefMessageRead", message})
}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNdefMessagesWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNdefMessagesWritten"})
}

func (ptr *QNearFieldTarget) NdefMessagesWritten() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NdefMessagesWritten"})
}

func (ptr *QNearFieldTarget) DisconnectRequestCompleted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestCompleted"})
}

func (ptr *QNearFieldTarget) SetKeepConnection(isPersistent bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetKeepConnection", isPersistent}).(bool)
}

func (ptr *QNearFieldTarget) ConnectType(f func() QNearFieldTarget__Type) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {

	return QNearFieldTarget__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QNearFieldTarget) ConnectUid(f func() *core.QByteArray) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUid", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectUid() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUid"})
}

func (ptr *QNearFieldTarget) Uid() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uid"}).(*core.QByteArray)
}

func (ptr *QNearFieldTarget) ConnectUrl(f func() *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUrl"})
}

func (ptr *QNearFieldTarget) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QNearFieldTarget) UrlDefault() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UrlDefault"}).(*core.QUrl)
}

func (ptr *QNearFieldTarget) ConnectDestroyQNearFieldTarget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNearFieldTarget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNearFieldTarget) DisconnectDestroyQNearFieldTarget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNearFieldTarget"})
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldTarget"})
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTargetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNearFieldTargetDefault"})
}

func (ptr *QNearFieldTarget) __sendCommands_commands_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendCommands_commands_atList", i}).(*core.QByteArray)
}

func (ptr *QNearFieldTarget) __sendCommands_commands_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendCommands_commands_setList", i})
}

func (ptr *QNearFieldTarget) __sendCommands_commands_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sendCommands_commands_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_atList(i int) *QNdefMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__writeNdefMessages_messages_atList", i}).(*QNdefMessage)
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_setList(i QNdefMessage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__writeNdefMessages_messages_setList", i})
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__writeNdefMessages_messages_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldTarget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNearFieldTarget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNearFieldTarget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNearFieldTarget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNearFieldTarget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNearFieldTarget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNearFieldTarget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNearFieldTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNearFieldTarget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNearFieldTarget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNearFieldTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNearFieldTarget) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNearFieldTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNearFieldTarget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNearFieldTarget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlNdefRecord struct {
	core.QObject
}

type QQmlNdefRecord_ITF interface {
	core.QObject_ITF
	QQmlNdefRecord_PTR() *QQmlNdefRecord
}

func (ptr *QQmlNdefRecord) QQmlNdefRecord_PTR() *QQmlNdefRecord {
	return ptr
}

func (ptr *QQmlNdefRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlNdefRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlNdefRecord(ptr QQmlNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNdefRecord_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNdefRecord) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlNdefRecord) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlNdefRecordFromPointer(ptr unsafe.Pointer) (n *QQmlNdefRecord) {
	n = new(QQmlNdefRecord)
	n.InitFromInternal(uintptr(ptr), "nfc.QQmlNdefRecord")
	return
}

//go:generate stringer -type=QQmlNdefRecord__TypeNameFormat
//QQmlNdefRecord::TypeNameFormat
type QQmlNdefRecord__TypeNameFormat int64

const (
	QQmlNdefRecord__Empty       QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Empty)
	QQmlNdefRecord__NfcRtd      QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__NfcRtd)
	QQmlNdefRecord__Mime        QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Mime)
	QQmlNdefRecord__Uri         QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Uri)
	QQmlNdefRecord__ExternalRtd QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__ExternalRtd)
	QQmlNdefRecord__Unknown     QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Unknown)
)

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQQmlNdefRecord", "", parent}).(*QQmlNdefRecord)
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "nfc.NewQQmlNdefRecord2", "", record, parent}).(*QQmlNdefRecord)
}

func (ptr *QQmlNdefRecord) Record() *QNdefRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record"}).(*QNdefRecord)
}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRecordChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRecordChanged"})
}

func (ptr *QQmlNdefRecord) RecordChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecordChanged"})
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRecord", record})
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", newtype})
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTypeNameFormat", newTypeNameFormat})
}

func (ptr *QQmlNdefRecord) Type() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(string)
}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTypeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTypeChanged"})
}

func (ptr *QQmlNdefRecord) TypeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeChanged"})
}

func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {

	return QQmlNdefRecord__TypeNameFormat(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeNameFormat"}).(float64))
}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTypeNameFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTypeNameFormatChanged"})
}

func (ptr *QQmlNdefRecord) TypeNameFormatChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeNameFormatChanged"})
}

func (ptr *QQmlNdefRecord) ConnectDestroyQQmlNdefRecord(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlNdefRecord", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNdefRecord) DisconnectDestroyQQmlNdefRecord() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlNdefRecord"})
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlNdefRecord"})
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecordDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlNdefRecordDefault"})
}

func (ptr *QQmlNdefRecord) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlNdefRecord) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlNdefRecord) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlNdefRecord) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlNdefRecord) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlNdefRecord) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlNdefRecord) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlNdefRecord) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlNdefRecord) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlNdefRecord) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlNdefRecord) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlNdefRecord) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlNdefRecord) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlNdefRecord) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlNdefRecord) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlNdefRecord) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlNdefRecord) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlNdefRecord) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["nfc.QNdefFilter"] = NewQNdefFilterFromPointer
	internal.ConstructorTable["nfc.QNdefMessage"] = NewQNdefMessageFromPointer
	internal.ConstructorTable["nfc.QNdefNfcIconRecord"] = NewQNdefNfcIconRecordFromPointer
	internal.ConstructorTable["nfc.QNdefNfcSmartPosterRecord"] = NewQNdefNfcSmartPosterRecordFromPointer
	internal.ConstructorTable["nfc.QNdefNfcTextRecord"] = NewQNdefNfcTextRecordFromPointer
	internal.ConstructorTable["nfc.QNdefNfcUriRecord"] = NewQNdefNfcUriRecordFromPointer
	internal.ConstructorTable["nfc.QNdefRecord"] = NewQNdefRecordFromPointer
	internal.ConstructorTable["nfc.QNearFieldManager"] = NewQNearFieldManagerFromPointer
	internal.ConstructorTable["nfc.QNearFieldShareManager"] = NewQNearFieldShareManagerFromPointer
	internal.ConstructorTable["nfc.QNearFieldShareTarget"] = NewQNearFieldShareTargetFromPointer
	internal.ConstructorTable["nfc.QNearFieldTarget"] = NewQNearFieldTargetFromPointer
	internal.ConstructorTable["nfc.QQmlNdefRecord"] = NewQQmlNdefRecordFromPointer
}
