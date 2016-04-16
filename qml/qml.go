package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = new(QJSEngine)
	n.SetPointer(ptr)
	return n
}

func newQJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = NewQJSEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QJSEngine_") {
		n.SetObjectName("QJSEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func NewQJSEngine() *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	return newQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	return newQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
}

func (ptr *QJSEngine) CollectGarbage() {
	defer qt.Recovering("QJSEngine::collectGarbage")

	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	defer qt.Recovering("QJSEngine::evaluate")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_Evaluate(ptr.Pointer(), C.CString(program), C.CString(fileName), C.int(lineNumber)))
	}
	return nil
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {
	defer qt.Recovering("QJSEngine::globalObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) InstallTranslatorFunctions(object QJSValue_ITF) {
	defer qt.Recovering("QJSEngine::installTranslatorFunctions")

	if ptr.Pointer() != nil {
		C.QJSEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	defer qt.Recovering("QJSEngine::newObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	defer qt.Recovering("QJSEngine::newQObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	defer qt.Recovering("QJSEngine::~QJSEngine")

	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQJSEngineTimerEvent
func callbackQJSEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQJSEngineChildEvent
func callbackQJSEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQJSEngineCustomEvent
func callbackQJSEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (p *QJSValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValue(ptr QJSValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValue_PTR().Pointer()
	}
	return nil
}

func NewQJSValueFromPointer(ptr unsafe.Pointer) *QJSValue {
	var n = new(QJSValue)
	n.SetPointer(ptr)
	return n
}

func newQJSValueFromPointer(ptr unsafe.Pointer) *QJSValue {
	var n = NewQJSValueFromPointer(ptr)
	return n
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue = QJSValue__SpecialValue(1)
)

func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue(C.int(value)))
}

func NewQJSValue4(value bool) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.int(qt.GoBoolToInt(value))))
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
}

func NewQJSValue8(value string) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue8(C.CString(value)))
}

func NewQJSValue10(value string) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue10(C.CString(value)))
}

func NewQJSValue5(value int) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	return newQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(value)))
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	defer qt.Recovering("QJSValue::deleteProperty")

	if ptr.Pointer() != nil {
		return C.QJSValue_DeleteProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) Equals(other QJSValue_ITF) bool {
	defer qt.Recovering("QJSValue::equals")

	if ptr.Pointer() != nil {
		return C.QJSValue_Equals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {
	defer qt.Recovering("QJSValue::hasOwnProperty")

	if ptr.Pointer() != nil {
		return C.QJSValue_HasOwnProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	defer qt.Recovering("QJSValue::hasProperty")

	if ptr.Pointer() != nil {
		return C.QJSValue_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) IsArray() bool {
	defer qt.Recovering("QJSValue::isArray")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsBool() bool {
	defer qt.Recovering("QJSValue::isBool")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsCallable() bool {
	defer qt.Recovering("QJSValue::isCallable")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsCallable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsDate() bool {
	defer qt.Recovering("QJSValue::isDate")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsError() bool {
	defer qt.Recovering("QJSValue::isError")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNull() bool {
	defer qt.Recovering("QJSValue::isNull")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNumber() bool {
	defer qt.Recovering("QJSValue::isNumber")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsObject() bool {
	defer qt.Recovering("QJSValue::isObject")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsQObject() bool {
	defer qt.Recovering("QJSValue::isQObject")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsRegExp() bool {
	defer qt.Recovering("QJSValue::isRegExp")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsString() bool {
	defer qt.Recovering("QJSValue::isString")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsUndefined() bool {
	defer qt.Recovering("QJSValue::isUndefined")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsVariant() bool {
	defer qt.Recovering("QJSValue::isVariant")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) Property(name string) *QJSValue {
	defer qt.Recovering("QJSValue::property")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSValue_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QJSValue) Prototype() *QJSValue {
	defer qt.Recovering("QJSValue::prototype")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSValue_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	defer qt.Recovering("QJSValue::setProperty")

	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	defer qt.Recovering("QJSValue::setPrototype")

	if ptr.Pointer() != nil {
		C.QJSValue_SetPrototype(ptr.Pointer(), PointerFromQJSValue(prototype))
	}
}

func (ptr *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	defer qt.Recovering("QJSValue::strictlyEquals")

	if ptr.Pointer() != nil {
		return C.QJSValue_StrictlyEquals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) ToBool() bool {
	defer qt.Recovering("QJSValue::toBool")

	if ptr.Pointer() != nil {
		return C.QJSValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) ToDateTime() *core.QDateTime {
	defer qt.Recovering("QJSValue::toDateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QJSValue_ToDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	defer qt.Recovering("QJSValue::toQObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QJSValue_ToQObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) ToString() string {
	defer qt.Recovering("QJSValue::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QJSValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	defer qt.Recovering("QJSValue::toVariant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) DestroyQJSValue() {
	defer qt.Recovering("QJSValue::~QJSValue")

	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(ptr.Pointer())
	}
}

type QJSValueIterator struct {
	ptr unsafe.Pointer
}

type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (p *QJSValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValueIterator(ptr QJSValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQJSValueIteratorFromPointer(ptr unsafe.Pointer) *QJSValueIterator {
	var n = new(QJSValueIterator)
	n.SetPointer(ptr)
	return n
}

func newQJSValueIteratorFromPointer(ptr unsafe.Pointer) *QJSValueIterator {
	var n = NewQJSValueIteratorFromPointer(ptr)
	return n
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return ptr
}

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (p *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlAbstractUrlInterceptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptor_PTR().Pointer()
	}
	return nil
}

func NewQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	var n = new(QQmlAbstractUrlInterceptor)
	n.SetPointer(ptr)
	return n
}

func newQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	var n = NewQQmlAbstractUrlInterceptorFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlAbstractUrlInterceptor_") {
		n.SetObjectNameAbs("QQmlAbstractUrlInterceptor_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::intercept")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlAbstractUrlInterceptor_Intercept(ptr.Pointer(), core.PointerFromQUrl(url), C.int(ty)))
	}
	return nil
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::~QQmlAbstractUrlInterceptor")

	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr.Pointer())
	}
}

func (ptr *QQmlAbstractUrlInterceptor) ObjectNameAbs() string {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlAbstractUrlInterceptor_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlAbstractUrlInterceptor) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) *QQmlApplicationEngine {
	var n = new(QQmlApplicationEngine)
	n.SetPointer(ptr)
	return n
}

func newQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) *QQmlApplicationEngine {
	var n = NewQQmlApplicationEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlApplicationEngine_") {
		n.SetObjectName("QQmlApplicationEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return ptr
}

func NewQQmlApplicationEngine(parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return newQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(core.PointerFromQObject(parent)))
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return newQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(C.CString(filePath), core.PointerFromQObject(parent)))
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return newQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load2(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::loadData")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadData(ptr.Pointer(), C.CString(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {
	defer qt.Recovering("connect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectObjectCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "objectCreated", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "objectCreated")
	}
}

//export callbackQQmlApplicationEngineObjectCreated
func callbackQQmlApplicationEngineObjectCreated(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::objectCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "objectCreated"); signal != nil {
		signal.(func(*core.QObject, *core.QUrl))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ObjectCreated(object core.QObject_ITF, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ObjectCreated(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	defer qt.Recovering("QQmlApplicationEngine::~QQmlApplicationEngine")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlApplicationEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlApplicationEngineTimerEvent
func callbackQQmlApplicationEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlApplicationEngineChildEvent
func callbackQQmlApplicationEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlApplicationEngineCustomEvent
func callbackQQmlApplicationEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlComponent struct {
	core.QObject
}

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func PointerFromQQmlComponent(ptr QQmlComponent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponent_PTR().Pointer()
	}
	return nil
}

func NewQQmlComponentFromPointer(ptr unsafe.Pointer) *QQmlComponent {
	var n = new(QQmlComponent)
	n.SetPointer(ptr)
	return n
}

func newQQmlComponentFromPointer(ptr unsafe.Pointer) *QQmlComponent {
	var n = NewQQmlComponentFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlComponent_") {
		n.SetObjectName("QQmlComponent_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return ptr
}

//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      = QQmlComponent__CompilationMode(1)
)

//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    = QQmlComponent__Status(0)
	QQmlComponent__Ready   = QQmlComponent__Status(1)
	QQmlComponent__Loading = QQmlComponent__Status(2)
	QQmlComponent__Error   = QQmlComponent__Status(3)
)

func (ptr *QQmlComponent) Progress() float64 {
	defer qt.Recovering("QQmlComponent::progress")

	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	defer qt.Recovering("QQmlComponent::status")

	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Url() *core.QUrl {
	defer qt.Recovering("QQmlComponent::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlComponent_Url(ptr.Pointer()))
	}
	return nil
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return newQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return newQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), C.CString(fileName), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return newQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), C.CString(fileName), core.PointerFromQObject(parent)))
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return newQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return newQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::beginCreate")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_BeginCreate(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
	}
	return nil
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {
	defer qt.Recovering("connect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "completeCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {
	defer qt.Recovering("disconnect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "completeCreate")
	}
}

//export callbackQQmlComponentCompleteCreate
func callbackQQmlComponentCompleteCreate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlComponent::completeCreate")

	if signal := qt.GetSignal(C.GoString(ptrName), "completeCreate"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).CompleteCreateDefault()
	}
}

func (ptr *QQmlComponent) CompleteCreate() {
	defer qt.Recovering("QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) CompleteCreateDefault() {
	defer qt.Recovering("QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreateDefault(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::create")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_Create(ptr.Pointer(), PointerFromQQmlContext(context)))
	}
	return nil
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {
	defer qt.Recovering("QQmlComponent::create")

	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	defer qt.Recovering("QQmlComponent::creationContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	defer qt.Recovering("QQmlComponent::isError")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	defer qt.Recovering("QQmlComponent::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	defer qt.Recovering("QQmlComponent::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	defer qt.Recovering("QQmlComponent::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.int(mode))
	}
}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {
	defer qt.Recovering("connect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectProgressChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "progressChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {
	defer qt.Recovering("disconnect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectProgressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "progressChanged")
	}
}

//export callbackQQmlComponentProgressChanged
func callbackQQmlComponentProgressChanged(ptr unsafe.Pointer, ptrName *C.char, progress C.double) {
	defer qt.Recovering("callback QQmlComponent::progressChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "progressChanged"); signal != nil {
		signal.(func(float64))(float64(progress))
	}

}

func (ptr *QQmlComponent) ProgressChanged(progress float64) {
	defer qt.Recovering("QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ProgressChanged(ptr.Pointer(), C.double(progress))
	}
}

func (ptr *QQmlComponent) SetData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::setData")

	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(ptr.Pointer(), C.CString(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	defer qt.Recovering("connect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQmlComponentStatusChanged
func callbackQQmlComponentStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQmlComponent::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQmlComponent__Status))(QQmlComponent__Status(status))
	}

}

func (ptr *QQmlComponent) StatusChanged(status QQmlComponent__Status) {
	defer qt.Recovering("QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	defer qt.Recovering("QQmlComponent::~QQmlComponent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlComponentTimerEvent
func callbackQQmlComponentTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlComponent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlComponent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlComponentChildEvent
func callbackQQmlComponentChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlComponent::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlComponent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlComponent::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlComponent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlComponentCustomEvent
func callbackQQmlComponentCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlComponent::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlComponent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlComponent::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = new(QQmlContext)
	n.SetPointer(ptr)
	return n
}

func newQQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = NewQQmlContextFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlContext_") {
		n.SetObjectName("QQmlContext_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return newQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return newQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlContext::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlContext_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	defer qt.Recovering("QQmlContext::contextObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlContext_ContextObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {
	defer qt.Recovering("QQmlContext::contextProperty")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlContext_ContextProperty(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlContext::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlContext_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	defer qt.Recovering("QQmlContext::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	defer qt.Recovering("QQmlContext::nameForObject")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	defer qt.Recovering("QQmlContext::parentContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ResolvedUrl(src core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QQmlContext::resolvedUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlContext_ResolvedUrl(ptr.Pointer(), core.PointerFromQUrl(src)))
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QQmlContext::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextObject")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty(ptr.Pointer(), C.CString(name), core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	defer qt.Recovering("QQmlContext::~QQmlContext")

	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlContextTimerEvent
func callbackQQmlContextTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlContext) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlContext::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlContext) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlContext::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlContext) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlContextChildEvent
func callbackQQmlContextChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlContext::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlContext) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlContext::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlContext) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlContextCustomEvent
func callbackQQmlContextCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlContext) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlContext::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlContext) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlContext::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = new(QQmlEngine)
	n.SetPointer(ptr)
	return n
}

func newQQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = NewQQmlEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlEngine_") {
		n.SetObjectName("QQmlEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership = QQmlEngine__ObjectOwnership(1)
)

func (ptr *QQmlEngine) OfflineStoragePath() string {
	defer qt.Recovering("QQmlEngine::offlineStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	defer qt.Recovering("QQmlEngine::setOfflineStoragePath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), C.CString(dir))
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	defer qt.Recovering("QQmlEngine::QQmlEngine")

	return newQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	defer qt.Recovering("QQmlEngine::addImageProvider")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), C.CString(providerId), PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	defer qt.Recovering("QQmlEngine::addImportPath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	defer qt.Recovering("QQmlEngine::addPluginPath")

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlEngine::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlEngine_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) ClearComponentCache() {
	defer qt.Recovering("QQmlEngine::clearComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlEngine::contextForObject")

	return NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlEngine::event")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	defer qt.Recovering("QQmlEngine::imageProvider")

	if ptr.Pointer() != nil {
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), C.CString(providerId)))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	defer qt.Recovering("QQmlEngine::importPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlEngine::incubationController")

	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlEngine::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	defer qt.Recovering("QQmlEngine::networkAccessManagerFactory")

	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	defer qt.Recovering("QQmlEngine::objectOwnership")

	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	defer qt.Recovering("QQmlEngine::outputWarningsToStandardError")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	defer qt.Recovering("QQmlEngine::pluginPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	defer qt.Recovering("connect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	defer qt.Recovering("disconnect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "quit")
	}
}

//export callbackQQmlEngineQuit
func callbackQQmlEngineQuit(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlEngine::quit")

	if signal := qt.GetSignal(C.GoString(ptrName), "quit"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlEngine) Quit() {
	defer qt.Recovering("QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_Quit(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	defer qt.Recovering("QQmlEngine::removeImageProvider")

	if ptr.Pointer() != nil {
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), C.CString(providerId))
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	defer qt.Recovering("QQmlEngine::rootContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlEngine::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	defer qt.Recovering("QQmlEngine::setContextForObject")

	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setImportPathList")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	defer qt.Recovering("QQmlEngine::setIncubationController")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	defer qt.Recovering("QQmlEngine::setNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	defer qt.Recovering("QQmlEngine::setObjectOwnership")

	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.int(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	defer qt.Recovering("QQmlEngine::setOutputWarningsToStandardError")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setPluginPathList")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	defer qt.Recovering("QQmlEngine::trimComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	defer qt.Recovering("QQmlEngine::~QQmlEngine")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlEngineTimerEvent
func callbackQQmlEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlEngineChildEvent
func callbackQQmlEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlEngineCustomEvent
func callbackQQmlEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (p *QQmlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = new(QQmlError)
	n.SetPointer(ptr)
	return n
}

func newQQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = NewQQmlErrorFromPointer(ptr)
	return n
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func NewQQmlError() *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	return newQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	return newQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
}

func (ptr *QQmlError) Column() int {
	defer qt.Recovering("QQmlError::column")

	if ptr.Pointer() != nil {
		return int(C.QQmlError_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	defer qt.Recovering("QQmlError::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	defer qt.Recovering("QQmlError::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	defer qt.Recovering("QQmlError::line")

	if ptr.Pointer() != nil {
		return int(C.QQmlError_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	defer qt.Recovering("QQmlError::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlError) SetColumn(column int) {
	defer qt.Recovering("QQmlError::setColumn")

	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	defer qt.Recovering("QQmlError::setDescription")

	if ptr.Pointer() != nil {
		C.QQmlError_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QQmlError) SetLine(line int) {
	defer qt.Recovering("QQmlError::setLine")

	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(line))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlError::setObject")

	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlError::setUrl")

	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) ToString() string {
	defer qt.Recovering("QQmlError::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) Url() *core.QUrl {
	defer qt.Recovering("QQmlError::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlError_Url(ptr.Pointer()))
	}
	return nil
}

type QQmlExpression struct {
	core.QObject
}

type QQmlExpression_ITF interface {
	core.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func PointerFromQQmlExpression(ptr QQmlExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpression_PTR().Pointer()
	}
	return nil
}

func NewQQmlExpressionFromPointer(ptr unsafe.Pointer) *QQmlExpression {
	var n = new(QQmlExpression)
	n.SetPointer(ptr)
	return n
}

func newQQmlExpressionFromPointer(ptr unsafe.Pointer) *QQmlExpression {
	var n = NewQQmlExpressionFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlExpression_") {
		n.SetObjectName("QQmlExpression_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return ptr
}

func NewQQmlExpression() *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return newQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression())
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return newQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression2(PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), C.CString(expression), core.PointerFromQObject(parent)))
}

func NewQQmlExpression3(script QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return newQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression3(PointerFromQQmlScriptString(script), PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), core.PointerFromQObject(parent)))
}

func (ptr *QQmlExpression) ClearError() {
	defer qt.Recovering("QQmlExpression::clearError")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) ColumnNumber() int {
	defer qt.Recovering("QQmlExpression::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	defer qt.Recovering("QQmlExpression::context")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlExpression_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlExpression::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlExpression_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) Error() *QQmlError {
	defer qt.Recovering("QQmlExpression::error")

	if ptr.Pointer() != nil {
		return NewQQmlErrorFromPointer(C.QQmlExpression_Error(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined bool) *core.QVariant {
	defer qt.Recovering("QQmlExpression::evaluate")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlExpression_Evaluate(ptr.Pointer(), C.int(qt.GoBoolToInt(valueIsUndefined))))
	}
	return nil
}

func (ptr *QQmlExpression) Expression() string {
	defer qt.Recovering("QQmlExpression::expression")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_Expression(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	defer qt.Recovering("QQmlExpression::hasError")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) LineNumber() int {
	defer qt.Recovering("QQmlExpression::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	defer qt.Recovering("QQmlExpression::notifyOnValueChanged")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_NotifyOnValueChanged(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	defer qt.Recovering("QQmlExpression::scopeObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlExpression_ScopeObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) SetExpression(expression string) {
	defer qt.Recovering("QQmlExpression::setExpression")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetExpression(ptr.Pointer(), C.CString(expression))
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	defer qt.Recovering("QQmlExpression::setNotifyOnValueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(notifyOnChange)))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	defer qt.Recovering("QQmlExpression::setSourceLocation")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetSourceLocation(ptr.Pointer(), C.CString(url), C.int(line), C.int(column))
	}
}

func (ptr *QQmlExpression) SourceFile() string {
	defer qt.Recovering("QQmlExpression::sourceFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_SourceFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	defer qt.Recovering("connect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlExpressionValueChanged
func callbackQQmlExpressionValueChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQmlExpression::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlExpression) ValueChanged() {
	defer qt.Recovering("QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ValueChanged(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	defer qt.Recovering("QQmlExpression::~QQmlExpression")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExpression) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlExpressionTimerEvent
func callbackQQmlExpressionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExpression) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExpression) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlExpression::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlExpression::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlExpressionChildEvent
func callbackQQmlExpressionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExpression::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExpression) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExpression::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExpression) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlExpression::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlExpression::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlExpressionCustomEvent
func callbackQQmlExpressionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExpression::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExpression) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExpression::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	return n
}

func newQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = NewQQmlExtensionPluginFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlExtensionPlugin_") {
		n.SetObjectName("QQmlExtensionPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlExtensionPlugin::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlExtensionPlugin_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	defer qt.Recovering("QQmlExtensionPlugin::registerTypes")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), C.CString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlExtensionPluginTimerEvent
func callbackQQmlExtensionPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlExtensionPluginChildEvent
func callbackQQmlExtensionPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlExtensionPluginCustomEvent
func callbackQQmlExtensionPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	return n
}

func newQQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = NewQQmlFileSelectorFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlFileSelector_") {
		n.SetObjectName("QQmlFileSelector_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::QQmlFileSelector")

	return newQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::get")

	return NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), C.CString(strings.Join(strin, "|")))
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), C.CString(strings.Join(strin, "|")))
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	defer qt.Recovering("QQmlFileSelector::setSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	defer qt.Recovering("QQmlFileSelector::~QQmlFileSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlFileSelector) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlFileSelectorTimerEvent
func callbackQQmlFileSelectorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlFileSelector) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlFileSelector) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlFileSelectorChildEvent
func callbackQQmlFileSelectorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlFileSelector) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlFileSelector) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlFileSelectorCustomEvent
func callbackQQmlFileSelectorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlFileSelector) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (p *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlImageProviderBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func NewQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) *QQmlImageProviderBase {
	var n = new(QQmlImageProviderBase)
	n.SetPointer(ptr)
	return n
}

func newQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) *QQmlImageProviderBase {
	var n = NewQQmlImageProviderBaseFromPointer(ptr)
	return n
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return ptr
}

//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading = QQmlImageProviderBase__Flag(0x01)
)

//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image   = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap  = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid = QQmlImageProviderBase__ImageType(3)
)

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	defer qt.Recovering("QQmlImageProviderBase::flags")

	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	defer qt.Recovering("QQmlImageProviderBase::imageType")

	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__ImageType(C.QQmlImageProviderBase_ImageType(ptr.Pointer()))
	}
	return 0
}

type QQmlIncubationController struct {
	ptr unsafe.Pointer
}

type QQmlIncubationController_ITF interface {
	QQmlIncubationController_PTR() *QQmlIncubationController
}

func (p *QQmlIncubationController) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubationController) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationController_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) *QQmlIncubationController {
	var n = new(QQmlIncubationController)
	n.SetPointer(ptr)
	return n
}

func newQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) *QQmlIncubationController {
	var n = NewQQmlIncubationControllerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlIncubationController_") {
		n.SetObjectNameAbs("QQmlIncubationController_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func NewQQmlIncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlIncubationController::QQmlIncubationController")

	return newQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlIncubationController::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlIncubationController_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	defer qt.Recovering("QQmlIncubationController::incubateFor")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCount")

	if ptr.Pointer() != nil {
		return int(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {
	defer qt.Recovering("connect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "incubatingObjectCountChanged", f)
	}
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {
	defer qt.Recovering("disconnect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "incubatingObjectCountChanged")
	}
}

//export callbackQQmlIncubationControllerIncubatingObjectCountChanged
func callbackQQmlIncubationControllerIncubatingObjectCountChanged(ptr unsafe.Pointer, ptrName *C.char, incubatingObjectCount C.int) {
	defer qt.Recovering("callback QQmlIncubationController::incubatingObjectCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "incubatingObjectCountChanged"); signal != nil {
		signal.(func(int))(int(incubatingObjectCount))
	} else {
		NewQQmlIncubationControllerFromPointer(ptr).IncubatingObjectCountChangedDefault(int(incubatingObjectCount))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChanged(incubatingObjectCount int) {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChanged(ptr.Pointer(), C.int(incubatingObjectCount))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChangedDefault(incubatingObjectCount int) {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChangedDefault(ptr.Pointer(), C.int(incubatingObjectCount))
	}
}

func (ptr *QQmlIncubationController) ObjectNameAbs() string {
	defer qt.Recovering("QQmlIncubationController::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlIncubationController_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlIncubationController) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlIncubationController::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (p *QQmlIncubator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubator(ptr QQmlIncubator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubator_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubatorFromPointer(ptr unsafe.Pointer) *QQmlIncubator {
	var n = new(QQmlIncubator)
	n.SetPointer(ptr)
	return n
}

func newQQmlIncubatorFromPointer(ptr unsafe.Pointer) *QQmlIncubator {
	var n = NewQQmlIncubatorFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlIncubator_") {
		n.SetObjectNameAbs("QQmlIncubator_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return ptr
}

//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          = QQmlIncubator__IncubationMode(2)
)

//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   = QQmlIncubator__Status(1)
	QQmlIncubator__Loading = QQmlIncubator__Status(2)
	QQmlIncubator__Error   = QQmlIncubator__Status(3)
)

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	defer qt.Recovering("QQmlIncubator::QQmlIncubator")

	return newQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.int(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	defer qt.Recovering("QQmlIncubator::clear")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	defer qt.Recovering("QQmlIncubator::forceCompletion")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	defer qt.Recovering("QQmlIncubator::incubationMode")

	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	defer qt.Recovering("QQmlIncubator::isError")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	defer qt.Recovering("QQmlIncubator::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	defer qt.Recovering("QQmlIncubator::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	defer qt.Recovering("QQmlIncubator::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	defer qt.Recovering("QQmlIncubator::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlIncubator_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {
	defer qt.Recovering("connect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setInitialState", f)
	}
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {
	defer qt.Recovering("disconnect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setInitialState")
	}
}

//export callbackQQmlIncubatorSetInitialState
func callbackQQmlIncubatorSetInitialState(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QQmlIncubator::setInitialState")

	if signal := qt.GetSignal(C.GoString(ptrName), "setInitialState"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQQmlIncubatorFromPointer(ptr).SetInitialStateDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QQmlIncubator) SetInitialState(object core.QObject_ITF) {
	defer qt.Recovering("QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialState(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) SetInitialStateDefault(object core.QObject_ITF) {
	defer qt.Recovering("QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialStateDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	defer qt.Recovering("QQmlIncubator::status")

	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {
	defer qt.Recovering("connect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "statusChanged", f)
	}
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "statusChanged")
	}
}

//export callbackQQmlIncubatorStatusChanged
func callbackQQmlIncubatorStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQmlIncubator::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQmlIncubator__Status))(QQmlIncubator__Status(status))
	} else {
		NewQQmlIncubatorFromPointer(ptr).StatusChangedDefault(QQmlIncubator__Status(status))
	}
}

func (ptr *QQmlIncubator) StatusChanged(status QQmlIncubator__Status) {
	defer qt.Recovering("QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQmlIncubator) StatusChangedDefault(status QQmlIncubator__Status) {
	defer qt.Recovering("QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChangedDefault(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQmlIncubator) ObjectNameAbs() string {
	defer qt.Recovering("QQmlIncubator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlIncubator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlIncubator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlIncubator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (p *QQmlListProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListProperty(ptr QQmlListProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlListPropertyFromPointer(ptr unsafe.Pointer) *QQmlListProperty {
	var n = new(QQmlListProperty)
	n.SetPointer(ptr)
	return n
}

func newQQmlListPropertyFromPointer(ptr unsafe.Pointer) *QQmlListProperty {
	var n = NewQQmlListPropertyFromPointer(ptr)
	return n
}

func (ptr *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return ptr
}

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (p *QQmlListReference) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListReference) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListReference(ptr QQmlListReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReference_PTR().Pointer()
	}
	return nil
}

func NewQQmlListReferenceFromPointer(ptr unsafe.Pointer) *QQmlListReference {
	var n = new(QQmlListReference)
	n.SetPointer(ptr)
	return n
}

func newQQmlListReferenceFromPointer(ptr unsafe.Pointer) *QQmlListReference {
	var n = NewQQmlListReferenceFromPointer(ptr)
	return n
}

func (ptr *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return ptr
}

func NewQQmlListReference() *QQmlListReference {
	defer qt.Recovering("QQmlListReference::QQmlListReference")

	return newQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference())
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {
	defer qt.Recovering("QQmlListReference::QQmlListReference")

	return newQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference2(core.PointerFromQObject(object), C.CString(property), PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {
	defer qt.Recovering("QQmlListReference::append")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Append(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	defer qt.Recovering("QQmlListReference::at")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlListReference_At(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QQmlListReference) CanAppend() bool {
	defer qt.Recovering("QQmlListReference::canAppend")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAppend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	defer qt.Recovering("QQmlListReference::canAt")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAt(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	defer qt.Recovering("QQmlListReference::canClear")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanClear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	defer qt.Recovering("QQmlListReference::canCount")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanCount(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	defer qt.Recovering("QQmlListReference::clear")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Count() int {
	defer qt.Recovering("QQmlListReference::count")

	if ptr.Pointer() != nil {
		return int(C.QQmlListReference_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlListReference) IsManipulable() bool {
	defer qt.Recovering("QQmlListReference::isManipulable")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsManipulable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	defer qt.Recovering("QQmlListReference::isReadable")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	defer qt.Recovering("QQmlListReference::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	defer qt.Recovering("QQmlListReference::listElementType")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlListReference_ListElementType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	defer qt.Recovering("QQmlListReference::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlListReference_Object(ptr.Pointer()))
	}
	return nil
}

type QQmlNetworkAccessManagerFactory struct {
	ptr unsafe.Pointer
}

type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (p *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlNetworkAccessManagerFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactory_PTR().Pointer()
	}
	return nil
}

func NewQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	var n = new(QQmlNetworkAccessManagerFactory)
	n.SetPointer(ptr)
	return n
}

func newQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	var n = NewQQmlNetworkAccessManagerFactoryFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlNetworkAccessManagerFactory_") {
		n.SetObjectNameAbs("QQmlNetworkAccessManagerFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::create")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) ObjectNameAbs() string {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNetworkAccessManagerFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlNetworkAccessManagerFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (p *QQmlParserStatus) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlParserStatus) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) *QQmlParserStatus {
	var n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return n
}

func newQQmlParserStatusFromPointer(ptr unsafe.Pointer) *QQmlParserStatus {
	var n = NewQQmlParserStatusFromPointer(ptr)
	return n
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) ClassBegin() {
	defer qt.Recovering("QQmlParserStatus::classBegin")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	defer qt.Recovering("QQmlParserStatus::componentComplete")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (p *QQmlProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlProperty(ptr QQmlProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyFromPointer(ptr unsafe.Pointer) *QQmlProperty {
	var n = new(QQmlProperty)
	n.SetPointer(ptr)
	return n
}

func newQQmlPropertyFromPointer(ptr unsafe.Pointer) *QQmlProperty {
	var n = NewQQmlPropertyFromPointer(ptr)
	return n
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return ptr
}

//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          = QQmlProperty__PropertyTypeCategory(3)
)

//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        = QQmlProperty__Type(0)
	QQmlProperty__Property       = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty())
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty2(core.PointerFromQObject(obj)))
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty3(core.PointerFromQObject(obj), PointerFromQQmlContext(ctxt)))
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty4(core.PointerFromQObject(obj), PointerFromQQmlEngine(engine)))
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty5(core.PointerFromQObject(obj), C.CString(name)))
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty6(core.PointerFromQObject(obj), C.CString(name), PointerFromQQmlContext(ctxt)))
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty7(core.PointerFromQObject(obj), C.CString(name), PointerFromQQmlEngine(engine)))
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	return newQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty8(PointerFromQQmlProperty(other)))
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QQmlProperty::connectNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal(ptr.Pointer(), core.PointerFromQObject(dest), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {
	defer qt.Recovering("QQmlProperty::connectNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal2(ptr.Pointer(), core.PointerFromQObject(dest), C.int(method)) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	defer qt.Recovering("QQmlProperty::hasNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_HasNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Index() int {
	defer qt.Recovering("QQmlProperty::index")

	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) IsDesignable() bool {
	defer qt.Recovering("QQmlProperty::isDesignable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsDesignable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	defer qt.Recovering("QQmlProperty::isProperty")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	defer qt.Recovering("QQmlProperty::isResettable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsResettable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	defer qt.Recovering("QQmlProperty::isSignalProperty")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsSignalProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	defer qt.Recovering("QQmlProperty::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	defer qt.Recovering("QQmlProperty::isWritable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Name() string {
	defer qt.Recovering("QQmlProperty::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	defer qt.Recovering("QQmlProperty::needsNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_NeedsNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Object() *core.QObject {
	defer qt.Recovering("QQmlProperty::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlProperty_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlProperty) PropertyType() int {
	defer qt.Recovering("QQmlProperty::propertyType")

	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_PropertyType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	defer qt.Recovering("QQmlProperty::propertyTypeCategory")

	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), C.CString(name)))
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), C.CString(name), PointerFromQQmlContext(ctxt)))
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), C.CString(name), PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlProperty) Read() *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlProperty_Read(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlProperty) Reset() bool {
	defer qt.Recovering("QQmlProperty::reset")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	defer qt.Recovering("QQmlProperty::type")

	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value)) != 0
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_Write(ptr.Pointer(), core.PointerFromQVariant(value)) != 0
	}
	return false
}

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	return n
}

func newQQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = NewQQmlPropertyMapFromPointer(ptr)
	for len(n.ObjectName()) < len("QQmlPropertyMap_") {
		n.SetObjectName("QQmlPropertyMap_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {
	defer qt.Recovering("QQmlPropertyMap::QQmlPropertyMap")

	return newQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	defer qt.Recovering("QQmlPropertyMap::clear")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_Clear(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	defer qt.Recovering("QQmlPropertyMap::contains")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Contains(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	defer qt.Recovering("QQmlPropertyMap::count")

	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	defer qt.Recovering("QQmlPropertyMap::isEmpty")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	defer qt.Recovering("QQmlPropertyMap::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlPropertyMap_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	defer qt.Recovering("QQmlPropertyMap::size")

	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QQmlPropertyMap::updateValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(input)))
	}
	return nil
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {
	defer qt.Recovering("QQmlPropertyMap::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlPropertyMap_Value(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlPropertyMapValueChanged
func callbackQQmlPropertyMapValueChanged(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ValueChanged(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	defer qt.Recovering("QQmlPropertyMap::~QQmlPropertyMap")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlPropertyMapTimerEvent
func callbackQQmlPropertyMapTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlPropertyMapChildEvent
func callbackQQmlPropertyMapChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlPropertyMapCustomEvent
func callbackQQmlPropertyMapCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQmlPropertyValueSource struct {
	ptr unsafe.Pointer
}

type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (p *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlPropertyValueSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSource_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) *QQmlPropertyValueSource {
	var n = new(QQmlPropertyValueSource)
	n.SetPointer(ptr)
	return n
}

func newQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) *QQmlPropertyValueSource {
	var n = NewQQmlPropertyValueSourceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlPropertyValueSource_") {
		n.SetObjectNameAbs("QQmlPropertyValueSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	defer qt.Recovering("QQmlPropertyValueSource::setTarget")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	defer qt.Recovering("QQmlPropertyValueSource::~QQmlPropertyValueSource")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
	}
}

func (ptr *QQmlPropertyValueSource) ObjectNameAbs() string {
	defer qt.Recovering("QQmlPropertyValueSource::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlPropertyValueSource_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlPropertyValueSource) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlPropertyValueSource::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQmlScriptString struct {
	ptr unsafe.Pointer
}

type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (p *QQmlScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlScriptString(ptr QQmlScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptString_PTR().Pointer()
	}
	return nil
}

func NewQQmlScriptStringFromPointer(ptr unsafe.Pointer) *QQmlScriptString {
	var n = new(QQmlScriptString)
	n.SetPointer(ptr)
	return n
}

func newQQmlScriptStringFromPointer(ptr unsafe.Pointer) *QQmlScriptString {
	var n = NewQQmlScriptStringFromPointer(ptr)
	return n
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return ptr
}

func NewQQmlScriptString() *QQmlScriptString {
	defer qt.Recovering("QQmlScriptString::QQmlScriptString")

	return newQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString())
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {
	defer qt.Recovering("QQmlScriptString::QQmlScriptString")

	return newQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString2(PointerFromQQmlScriptString(other)))
}

func (ptr *QQmlScriptString) BooleanLiteral(ok bool) bool {
	defer qt.Recovering("QQmlScriptString::booleanLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_BooleanLiteral(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	defer qt.Recovering("QQmlScriptString::isEmpty")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	defer qt.Recovering("QQmlScriptString::isNullLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsNullLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	defer qt.Recovering("QQmlScriptString::isUndefinedLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsUndefinedLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) NumberLiteral(ok bool) float64 {
	defer qt.Recovering("QQmlScriptString::numberLiteral")

	if ptr.Pointer() != nil {
		return float64(C.QQmlScriptString_NumberLiteral(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QQmlScriptString) StringLiteral() string {
	defer qt.Recovering("QQmlScriptString::stringLiteral")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlScriptString_StringLiteral(ptr.Pointer()))
	}
	return ""
}
