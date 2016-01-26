package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptor_ITF interface {
	core.QObject_ITF
	QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) *QDBusAbstractAdaptor {
	var n = new(QDBusAbstractAdaptor)
	n.SetPointer(ptr)
	return n
}

func newQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) *QDBusAbstractAdaptor {
	var n = NewQDBusAbstractAdaptorFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusAbstractAdaptor_") {
		n.SetObjectName("QDBusAbstractAdaptor_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusAbstractAdaptor) QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor {
	return ptr
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	defer qt.Recovering("QDBusAbstractAdaptor::~QDBusAbstractAdaptor")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusAbstractAdaptorTimerEvent
func callbackQDBusAbstractAdaptorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusAbstractAdaptorChildEvent
func callbackQDBusAbstractAdaptorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusAbstractAdaptorCustomEvent
func callbackQDBusAbstractAdaptorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	return n
}

func newQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = NewQDBusAbstractInterfaceFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusAbstractInterface_") {
		n.SetObjectName("QDBusAbstractInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Interface() string {
	defer qt.Recovering("QDBusAbstractInterface::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	defer qt.Recovering("QDBusAbstractInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) Path() string {
	defer qt.Recovering("QDBusAbstractInterface::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	defer qt.Recovering("QDBusAbstractInterface::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	defer qt.Recovering("QDBusAbstractInterface::setTimeout")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	defer qt.Recovering("QDBusAbstractInterface::timeout")

	if ptr.Pointer() != nil {
		return int(C.QDBusAbstractInterface_Timeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	defer qt.Recovering("QDBusAbstractInterface::~QDBusAbstractInterface")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusAbstractInterfaceTimerEvent
func callbackQDBusAbstractInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusAbstractInterfaceChildEvent
func callbackQDBusAbstractInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusAbstractInterfaceCustomEvent
func callbackQDBusAbstractInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusArgument struct {
	ptr unsafe.Pointer
}

type QDBusArgument_ITF interface {
	QDBusArgument_PTR() *QDBusArgument
}

func (p *QDBusArgument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusArgument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusArgument(ptr QDBusArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgument_PTR().Pointer()
	}
	return nil
}

func NewQDBusArgumentFromPointer(ptr unsafe.Pointer) *QDBusArgument {
	var n = new(QDBusArgument)
	n.SetPointer(ptr)
	return n
}

func newQDBusArgumentFromPointer(ptr unsafe.Pointer) *QDBusArgument {
	var n = NewQDBusArgumentFromPointer(ptr)
	return n
}

func (ptr *QDBusArgument) QDBusArgument_PTR() *QDBusArgument {
	return ptr
}

//QDBusArgument::ElementType
type QDBusArgument__ElementType int64

const (
	QDBusArgument__BasicType     = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   = QDBusArgument__ElementType(-1)
)

func NewQDBusArgument() *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	return newQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument())
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	return newQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument2(PointerFromQDBusArgument(other)))
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {
	defer qt.Recovering("QDBusArgument::asVariant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QDBusArgument_AsVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusArgument) AtEnd() bool {
	defer qt.Recovering("QDBusArgument::atEnd")

	if ptr.Pointer() != nil {
		return C.QDBusArgument_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray(id int) {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QDBusArgument) BeginArray2() {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(ptr.Pointer(), C.int(kid), C.int(vid))
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	defer qt.Recovering("QDBusArgument::currentType")

	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusArgument) EndArray() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray2() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap2() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
	defer qt.Recovering("QDBusArgument::~QDBusArgument")

	if ptr.Pointer() != nil {
		C.QDBusArgument_DestroyQDBusArgument(ptr.Pointer())
	}
}

type QDBusConnection struct {
	ptr unsafe.Pointer
}

type QDBusConnection_ITF interface {
	QDBusConnection_PTR() *QDBusConnection
}

func (p *QDBusConnection) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusConnection) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusConnection(ptr QDBusConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnection_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionFromPointer(ptr unsafe.Pointer) *QDBusConnection {
	var n = new(QDBusConnection)
	n.SetPointer(ptr)
	return n
}

func newQDBusConnectionFromPointer(ptr unsafe.Pointer) *QDBusConnection {
	var n = NewQDBusConnectionFromPointer(ptr)
	return n
}

func (ptr *QDBusConnection) QDBusConnection_PTR() *QDBusConnection {
	return ptr
}

//QDBusConnection::BusType
type QDBusConnection__BusType int64

const (
	QDBusConnection__SessionBus    = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus = QDBusConnection__BusType(2)
)

//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int64

const (
	QDBusConnection__UnixFileDescriptorPassing = QDBusConnection__ConnectionCapability(0x0001)
)

//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int64

const (
	QDBusConnection__ExportAdaptors                = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportAllSlots                = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSlots | QDBusConnection__ExportNonScriptableSlots)
	QDBusConnection__ExportAllSignals              = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSignals | QDBusConnection__ExportNonScriptableSignals)
	QDBusConnection__ExportAllProperties           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableProperties | QDBusConnection__ExportNonScriptableProperties)
	QDBusConnection__ExportAllInvokables           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableInvokables | QDBusConnection__ExportNonScriptableInvokables)
	QDBusConnection__ExportAllContents             = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableContents | QDBusConnection__ExportNonScriptableContents)
	QDBusConnection__ExportChildObjects            = QDBusConnection__RegisterOption(0x1000)
)

//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int64

const (
	QDBusConnection__UnregisterNode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree = QDBusConnection__UnregisterMode(1)
)

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	return newQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection2(PointerFromQDBusConnection(other)))
}

func NewQDBusConnection(name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	return newQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection(C.CString(name)))
}

func (ptr *QDBusConnection) BaseService() string {
	defer qt.Recovering("QDBusConnection::baseService")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_BaseService(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {
	defer qt.Recovering("QDBusConnection::callWithCallback")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_CallWithCallback(ptr.Pointer(), PointerFromQDBusMessage(message), core.PointerFromQObject(receiver), C.CString(returnMethod), C.CString(errorMethod), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect2(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(signature), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect3(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(strings.Join(argumentMatch, "|")), C.CString(signature), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	defer qt.Recovering("QDBusConnection::connectionCapabilities")

	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(ptr.Pointer()))
	}
	return 0
}

func QDBusConnection_DisconnectFromBus(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromBus")

	C.QDBusConnection_QDBusConnection_DisconnectFromBus(C.CString(name))
}

func QDBusConnection_DisconnectFromPeer(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromPeer")

	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(C.CString(name))
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	defer qt.Recovering("QDBusConnection::interface")

	if ptr.Pointer() != nil {
		return NewQDBusConnectionInterfaceFromPointer(C.QDBusConnection_Interface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusConnection) IsConnected() bool {
	defer qt.Recovering("QDBusConnection::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func QDBusConnection_LocalMachineId() *core.QByteArray {
	defer qt.Recovering("QDBusConnection::localMachineId")

	return core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
}

func (ptr *QDBusConnection) Name() string {
	defer qt.Recovering("QDBusConnection::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	defer qt.Recovering("QDBusConnection::objectRegisteredAt")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDBusConnection_ObjectRegisteredAt(ptr.Pointer(), C.CString(path)))
	}
	return nil
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject(ptr.Pointer(), C.CString(path), core.PointerFromQObject(object), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject2(ptr.Pointer(), C.CString(path), C.CString(interfa), core.PointerFromQObject(object), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::registerService")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterService(ptr.Pointer(), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {
	defer qt.Recovering("QDBusConnection::send")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Send(ptr.Pointer(), PointerFromQDBusMessage(message)) != 0
	}
	return false
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	defer qt.Recovering("QDBusConnection::unregisterObject")

	if ptr.Pointer() != nil {
		C.QDBusConnection_UnregisterObject(ptr.Pointer(), C.CString(path), C.int(mode))
	}
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::unregisterService")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_UnregisterService(ptr.Pointer(), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
	defer qt.Recovering("QDBusConnection::~QDBusConnection")

	if ptr.Pointer() != nil {
		C.QDBusConnection_DestroyQDBusConnection(ptr.Pointer())
	}
}

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) *QDBusConnectionInterface {
	var n = new(QDBusConnectionInterface)
	n.SetPointer(ptr)
	return n
}

func newQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) *QDBusConnectionInterface {
	var n = NewQDBusConnectionInterfaceFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusConnectionInterface_") {
		n.SetObjectName("QDBusConnectionInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return ptr
}

//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        = QDBusConnectionInterface__RegisterServiceReply(2)
)

//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceRegistered
func callbackQDBusConnectionInterfaceServiceRegistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceRegistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ServiceRegistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceUnregistered
func callbackQDBusConnectionInterfaceServiceUnregistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceUnregistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ServiceUnregistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusConnectionInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusConnectionInterfaceTimerEvent
func callbackQDBusConnectionInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusConnectionInterfaceChildEvent
func callbackQDBusConnectionInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusConnectionInterfaceCustomEvent
func callbackQDBusConnectionInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusContext struct {
	ptr unsafe.Pointer
}

type QDBusContext_ITF interface {
	QDBusContext_PTR() *QDBusContext
}

func (p *QDBusContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusContext(ptr QDBusContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContext_PTR().Pointer()
	}
	return nil
}

func NewQDBusContextFromPointer(ptr unsafe.Pointer) *QDBusContext {
	var n = new(QDBusContext)
	n.SetPointer(ptr)
	return n
}

func newQDBusContextFromPointer(ptr unsafe.Pointer) *QDBusContext {
	var n = NewQDBusContextFromPointer(ptr)
	return n
}

func (ptr *QDBusContext) QDBusContext_PTR() *QDBusContext {
	return ptr
}

func NewQDBusContext() *QDBusContext {
	defer qt.Recovering("QDBusContext::QDBusContext")

	return newQDBusContextFromPointer(C.QDBusContext_NewQDBusContext())
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	defer qt.Recovering("QDBusContext::calledFromDBus")

	if ptr.Pointer() != nil {
		return C.QDBusContext_CalledFromDBus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	defer qt.Recovering("QDBusContext::isDelayedReply")

	if ptr.Pointer() != nil {
		return C.QDBusContext_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply2(ptr.Pointer(), C.int(ty), C.CString(msg))
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply(ptr.Pointer(), C.CString(name), C.CString(msg))
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	defer qt.Recovering("QDBusContext::setDelayedReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusContext) DestroyQDBusContext() {
	defer qt.Recovering("QDBusContext::~QDBusContext")

	if ptr.Pointer() != nil {
		C.QDBusContext_DestroyQDBusContext(ptr.Pointer())
	}
}

type QDBusError struct {
	ptr unsafe.Pointer
}

type QDBusError_ITF interface {
	QDBusError_PTR() *QDBusError
}

func (p *QDBusError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusError(ptr QDBusError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusError_PTR().Pointer()
	}
	return nil
}

func NewQDBusErrorFromPointer(ptr unsafe.Pointer) *QDBusError {
	var n = new(QDBusError)
	n.SetPointer(ptr)
	return n
}

func newQDBusErrorFromPointer(ptr unsafe.Pointer) *QDBusError {
	var n = NewQDBusErrorFromPointer(ptr)
	return n
}

func (ptr *QDBusError) QDBusError_PTR() *QDBusError {
	return ptr
}

//QDBusError::ErrorType
type QDBusError__ErrorType int64

const (
	QDBusError__NoError           = QDBusError__ErrorType(0)
	QDBusError__Other             = QDBusError__ErrorType(1)
	QDBusError__Failed            = QDBusError__ErrorType(2)
	QDBusError__NoMemory          = QDBusError__ErrorType(3)
	QDBusError__ServiceUnknown    = QDBusError__ErrorType(4)
	QDBusError__NoReply           = QDBusError__ErrorType(5)
	QDBusError__BadAddress        = QDBusError__ErrorType(6)
	QDBusError__NotSupported      = QDBusError__ErrorType(7)
	QDBusError__LimitsExceeded    = QDBusError__ErrorType(8)
	QDBusError__AccessDenied      = QDBusError__ErrorType(9)
	QDBusError__NoServer          = QDBusError__ErrorType(10)
	QDBusError__Timeout           = QDBusError__ErrorType(11)
	QDBusError__NoNetwork         = QDBusError__ErrorType(12)
	QDBusError__AddressInUse      = QDBusError__ErrorType(13)
	QDBusError__Disconnected      = QDBusError__ErrorType(14)
	QDBusError__InvalidArgs       = QDBusError__ErrorType(15)
	QDBusError__UnknownMethod     = QDBusError__ErrorType(16)
	QDBusError__TimedOut          = QDBusError__ErrorType(17)
	QDBusError__InvalidSignature  = QDBusError__ErrorType(18)
	QDBusError__UnknownInterface  = QDBusError__ErrorType(19)
	QDBusError__UnknownObject     = QDBusError__ErrorType(20)
	QDBusError__UnknownProperty   = QDBusError__ErrorType(21)
	QDBusError__PropertyReadOnly  = QDBusError__ErrorType(22)
	QDBusError__InternalError     = QDBusError__ErrorType(23)
	QDBusError__InvalidService    = QDBusError__ErrorType(24)
	QDBusError__InvalidObjectPath = QDBusError__ErrorType(25)
	QDBusError__InvalidInterface  = QDBusError__ErrorType(26)
	QDBusError__InvalidMember     = QDBusError__ErrorType(27)
)

func QDBusError_ErrorString(error QDBusError__ErrorType) string {
	defer qt.Recovering("QDBusError::errorString")

	return C.GoString(C.QDBusError_QDBusError_ErrorString(C.int(error)))
}

func (ptr *QDBusError) IsValid() bool {
	defer qt.Recovering("QDBusError::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusError) Message() string {
	defer qt.Recovering("QDBusError::message")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Name() string {
	defer qt.Recovering("QDBusError::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {
	defer qt.Recovering("QDBusError::type")

	if ptr.Pointer() != nil {
		return QDBusError__ErrorType(C.QDBusError_Type(ptr.Pointer()))
	}
	return 0
}

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = new(QDBusInterface)
	n.SetPointer(ptr)
	return n
}

func newQDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = NewQDBusInterfaceFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusInterface_") {
		n.SetObjectName("QDBusInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return ptr
}

func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {
	defer qt.Recovering("QDBusInterface::QDBusInterface")

	return newQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(C.CString(service), C.CString(path), C.CString(interfa), PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	defer qt.Recovering("QDBusInterface::~QDBusInterface")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusInterfaceTimerEvent
func callbackQDBusInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusInterfaceChildEvent
func callbackQDBusInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusInterfaceCustomEvent
func callbackQDBusInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusMessage struct {
	ptr unsafe.Pointer
}

type QDBusMessage_ITF interface {
	QDBusMessage_PTR() *QDBusMessage
}

func (p *QDBusMessage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusMessage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusMessage(ptr QDBusMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessage_PTR().Pointer()
	}
	return nil
}

func NewQDBusMessageFromPointer(ptr unsafe.Pointer) *QDBusMessage {
	var n = new(QDBusMessage)
	n.SetPointer(ptr)
	return n
}

func newQDBusMessageFromPointer(ptr unsafe.Pointer) *QDBusMessage {
	var n = NewQDBusMessageFromPointer(ptr)
	return n
}

func (ptr *QDBusMessage) QDBusMessage_PTR() *QDBusMessage {
	return ptr
}

//QDBusMessage::MessageType
type QDBusMessage__MessageType int64

const (
	QDBusMessage__InvalidMessage    = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     = QDBusMessage__MessageType(4)
)

func NewQDBusMessage() *QDBusMessage {
	defer qt.Recovering("QDBusMessage::QDBusMessage")

	return newQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage())
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::QDBusMessage")

	return newQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage2(PointerFromQDBusMessage(other)))
}

func (ptr *QDBusMessage) AutoStartService() bool {
	defer qt.Recovering("QDBusMessage::autoStartService")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_AutoStartService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) ErrorMessage() string {
	defer qt.Recovering("QDBusMessage::errorMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	defer qt.Recovering("QDBusMessage::errorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	defer qt.Recovering("QDBusMessage::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	defer qt.Recovering("QDBusMessage::isDelayedReply")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	defer qt.Recovering("QDBusMessage::isReplyRequired")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsReplyRequired(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) Member() string {
	defer qt.Recovering("QDBusMessage::member")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Member(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	defer qt.Recovering("QDBusMessage::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	defer qt.Recovering("QDBusMessage::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	defer qt.Recovering("QDBusMessage::setAutoStartService")

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	defer qt.Recovering("QDBusMessage::setDelayedReply")

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) Signature() string {
	defer qt.Recovering("QDBusMessage::signature")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	defer qt.Recovering("QDBusMessage::type")

	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	defer qt.Recovering("QDBusMessage::~QDBusMessage")

	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(ptr.Pointer())
	}
}

type QDBusObjectPath struct {
	ptr unsafe.Pointer
}

type QDBusObjectPath_ITF interface {
	QDBusObjectPath_PTR() *QDBusObjectPath
}

func (p *QDBusObjectPath) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusObjectPath) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPath_PTR().Pointer()
	}
	return nil
}

func NewQDBusObjectPathFromPointer(ptr unsafe.Pointer) *QDBusObjectPath {
	var n = new(QDBusObjectPath)
	n.SetPointer(ptr)
	return n
}

func newQDBusObjectPathFromPointer(ptr unsafe.Pointer) *QDBusObjectPath {
	var n = NewQDBusObjectPathFromPointer(ptr)
	return n
}

func (ptr *QDBusObjectPath) QDBusObjectPath_PTR() *QDBusObjectPath {
	return ptr
}

func NewQDBusObjectPath() *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	return newQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath())
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	return newQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath3(core.PointerFromQLatin1String(path)))
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	return newQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath4(C.CString(path)))
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	return newQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath2(C.CString(path)))
}

func (ptr *QDBusObjectPath) Path() string {
	defer qt.Recovering("QDBusObjectPath::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusObjectPath_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	defer qt.Recovering("QDBusObjectPath::setPath")

	if ptr.Pointer() != nil {
		C.QDBusObjectPath_SetPath(ptr.Pointer(), C.CString(path))
	}
}

type QDBusPendingCall struct {
	ptr unsafe.Pointer
}

type QDBusPendingCall_ITF interface {
	QDBusPendingCall_PTR() *QDBusPendingCall
}

func (p *QDBusPendingCall) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusPendingCall) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallFromPointer(ptr unsafe.Pointer) *QDBusPendingCall {
	var n = new(QDBusPendingCall)
	n.SetPointer(ptr)
	return n
}

func newQDBusPendingCallFromPointer(ptr unsafe.Pointer) *QDBusPendingCall {
	var n = NewQDBusPendingCallFromPointer(ptr)
	return n
}

func (ptr *QDBusPendingCall) QDBusPendingCall_PTR() *QDBusPendingCall {
	return ptr
}

func NewQDBusPendingCall(other QDBusPendingCall_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::QDBusPendingCall")

	return newQDBusPendingCallFromPointer(C.QDBusPendingCall_NewQDBusPendingCall(PointerFromQDBusPendingCall(other)))
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {
	defer qt.Recovering("QDBusPendingCall::swap")

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(ptr.Pointer(), PointerFromQDBusPendingCall(other))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
	defer qt.Recovering("QDBusPendingCall::~QDBusPendingCall")

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_DestroyQDBusPendingCall(ptr.Pointer())
	}
}

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QDBusPendingCall_ITF
	QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher
}

func (p *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QDBusPendingCallWatcher) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QDBusPendingCall_PTR().SetPointer(ptr)
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	return n
}

func newQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = NewQDBusPendingCallWatcherFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusPendingCallWatcher_") {
		n.SetObjectName("QDBusPendingCallWatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	defer qt.Recovering("QDBusPendingCallWatcher::waitForFinished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	defer qt.Recovering("QDBusPendingCallWatcher::QDBusPendingCallWatcher")

	return newQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDBusPendingCallWatcherFinished
func callbackQDBusPendingCallWatcherFinished(ptr unsafe.Pointer, ptrName *C.char, self unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
	}

}

func (ptr *QDBusPendingCallWatcher) Finished(self QDBusPendingCallWatcher_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQDBusPendingCallWatcher(self))
	}
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	defer qt.Recovering("QDBusPendingCallWatcher::isFinished")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	defer qt.Recovering("QDBusPendingCallWatcher::~QDBusPendingCallWatcher")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusPendingCallWatcherTimerEvent
func callbackQDBusPendingCallWatcherTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusPendingCallWatcherChildEvent
func callbackQDBusPendingCallWatcherChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusPendingCallWatcherCustomEvent
func callbackQDBusPendingCallWatcherCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReply_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReply_PTR() *QDBusPendingReply
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingReplyFromPointer(ptr unsafe.Pointer) *QDBusPendingReply {
	var n = new(QDBusPendingReply)
	n.SetPointer(ptr)
	return n
}

func newQDBusPendingReplyFromPointer(ptr unsafe.Pointer) *QDBusPendingReply {
	var n = NewQDBusPendingReplyFromPointer(ptr)
	return n
}

func (ptr *QDBusPendingReply) QDBusPendingReply_PTR() *QDBusPendingReply {
	return ptr
}

type QDBusReply struct {
	ptr unsafe.Pointer
}

type QDBusReply_ITF interface {
	QDBusReply_PTR() *QDBusReply
}

func (p *QDBusReply) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusReply) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusReply(ptr QDBusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusReplyFromPointer(ptr unsafe.Pointer) *QDBusReply {
	var n = new(QDBusReply)
	n.SetPointer(ptr)
	return n
}

func newQDBusReplyFromPointer(ptr unsafe.Pointer) *QDBusReply {
	var n = NewQDBusReplyFromPointer(ptr)
	return n
}

func (ptr *QDBusReply) QDBusReply_PTR() *QDBusReply {
	return ptr
}

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = new(QDBusServer)
	n.SetPointer(ptr)
	return n
}

func newQDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = NewQDBusServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServer_") {
		n.SetObjectName("QDBusServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return ptr
}

func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return newQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return newQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(C.CString(address), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServer) Address() string {
	defer qt.Recovering("QDBusServer::address")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	defer qt.Recovering("QDBusServer::isAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	defer qt.Recovering("QDBusServer::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	defer qt.Recovering("QDBusServer::setAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	defer qt.Recovering("QDBusServer::~QDBusServer")

	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusServerTimerEvent
func callbackQDBusServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusServerChildEvent
func callbackQDBusServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusServerCustomEvent
func callbackQDBusServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	return n
}

func newQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = NewQDBusServiceWatcherFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServiceWatcher_") {
		n.SetObjectName("QDBusServiceWatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return ptr
}

//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchMode")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	defer qt.Recovering("QDBusServiceWatcher::watchMode")

	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	return newQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	return newQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(C.CString(service), PointerFromQDBusConnection(connection), C.int(watchMode), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	defer qt.Recovering("QDBusServiceWatcher::addWatchedService")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), C.CString(newService))
	}
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	defer qt.Recovering("QDBusServiceWatcher::removeWatchedService")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), C.CString(service)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceOwnerChanged")
	}
}

//export callbackQDBusServiceWatcherServiceOwnerChanged
func callbackQDBusServiceWatcherServiceOwnerChanged(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char, oldOwner *C.char, newOwner *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceOwnerChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceOwnerChanged"); signal != nil {
		signal.(func(string, string, string))(C.GoString(serviceName), C.GoString(oldOwner), C.GoString(newOwner))
	}

}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceOwnerChanged(ptr.Pointer(), C.CString(serviceName), C.CString(oldOwner), C.CString(newOwner))
	}
}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusServiceWatcherServiceRegistered
func callbackQDBusServiceWatcherServiceRegistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceRegistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceRegistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusServiceWatcherServiceUnregistered
func callbackQDBusServiceWatcherServiceUnregistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceUnregistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceUnregistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::setConnection")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchedServices")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), C.CString(strings.Join(services, "|")))
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	defer qt.Recovering("QDBusServiceWatcher::watchedServices")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	defer qt.Recovering("QDBusServiceWatcher::~QDBusServiceWatcher")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServiceWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusServiceWatcherTimerEvent
func callbackQDBusServiceWatcherTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusServiceWatcherChildEvent
func callbackQDBusServiceWatcherChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusServiceWatcherCustomEvent
func callbackQDBusServiceWatcherCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDBusSignature struct {
	ptr unsafe.Pointer
}

type QDBusSignature_ITF interface {
	QDBusSignature_PTR() *QDBusSignature
}

func (p *QDBusSignature) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusSignature) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusSignature(ptr QDBusSignature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignature_PTR().Pointer()
	}
	return nil
}

func NewQDBusSignatureFromPointer(ptr unsafe.Pointer) *QDBusSignature {
	var n = new(QDBusSignature)
	n.SetPointer(ptr)
	return n
}

func newQDBusSignatureFromPointer(ptr unsafe.Pointer) *QDBusSignature {
	var n = NewQDBusSignatureFromPointer(ptr)
	return n
}

func (ptr *QDBusSignature) QDBusSignature_PTR() *QDBusSignature {
	return ptr
}

func NewQDBusSignature() *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	return newQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature())
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	return newQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature3(core.PointerFromQLatin1String(signature)))
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	return newQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature4(C.CString(signature)))
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	return newQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature2(C.CString(signature)))
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	defer qt.Recovering("QDBusSignature::setSignature")

	if ptr.Pointer() != nil {
		C.QDBusSignature_SetSignature(ptr.Pointer(), C.CString(signature))
	}
}

func (ptr *QDBusSignature) Signature() string {
	defer qt.Recovering("QDBusSignature::signature")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusSignature_Signature(ptr.Pointer()))
	}
	return ""
}

type QDBusUnixFileDescriptor struct {
	ptr unsafe.Pointer
}

type QDBusUnixFileDescriptor_ITF interface {
	QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor
}

func (p *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusUnixFileDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) *QDBusUnixFileDescriptor {
	var n = new(QDBusUnixFileDescriptor)
	n.SetPointer(ptr)
	return n
}

func newQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) *QDBusUnixFileDescriptor {
	var n = NewQDBusUnixFileDescriptorFromPointer(ptr)
	return n
}

func (ptr *QDBusUnixFileDescriptor) QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor {
	return ptr
}

func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	return newQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	return newQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	return newQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(fileDescriptor)))
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	defer qt.Recovering("QDBusUnixFileDescriptor::fileDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer()))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	defer qt.Recovering("QDBusUnixFileDescriptor::isSupported")

	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	defer qt.Recovering("QDBusUnixFileDescriptor::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	defer qt.Recovering("QDBusUnixFileDescriptor::setFileDescriptor")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(fileDescriptor))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	defer qt.Recovering("QDBusUnixFileDescriptor::swap")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	defer qt.Recovering("QDBusUnixFileDescriptor::~QDBusUnixFileDescriptor")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(ptr.Pointer())
	}
}

type QDBusVariant struct {
	ptr unsafe.Pointer
}

type QDBusVariant_ITF interface {
	QDBusVariant_PTR() *QDBusVariant
}

func (p *QDBusVariant) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusVariant) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusVariant(ptr QDBusVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariant_PTR().Pointer()
	}
	return nil
}

func NewQDBusVariantFromPointer(ptr unsafe.Pointer) *QDBusVariant {
	var n = new(QDBusVariant)
	n.SetPointer(ptr)
	return n
}

func newQDBusVariantFromPointer(ptr unsafe.Pointer) *QDBusVariant {
	var n = NewQDBusVariantFromPointer(ptr)
	return n
}

func (ptr *QDBusVariant) QDBusVariant_PTR() *QDBusVariant {
	return ptr
}

func NewQDBusVariant() *QDBusVariant {
	defer qt.Recovering("QDBusVariant::QDBusVariant")

	return newQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant())
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {
	defer qt.Recovering("QDBusVariant::QDBusVariant")

	return newQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant2(core.PointerFromQVariant(variant)))
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {
	defer qt.Recovering("QDBusVariant::setVariant")

	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(ptr.Pointer(), core.PointerFromQVariant(variant))
	}
}

func (ptr *QDBusVariant) Variant() *core.QVariant {
	defer qt.Recovering("QDBusVariant::variant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QDBusVariant_Variant(ptr.Pointer()))
	}
	return nil
}

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	return n
}

func newQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = NewQDBusVirtualObjectFromPointer(ptr)
	for len(n.ObjectName()) < len("QDBusVirtualObject_") {
		n.SetObjectName("QDBusVirtualObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	defer qt.Recovering("QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), C.CString(path)))
	}
	return ""
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	defer qt.Recovering("QDBusVirtualObject::~QDBusVirtualObject")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusVirtualObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusVirtualObjectTimerEvent
func callbackQDBusVirtualObjectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusVirtualObjectChildEvent
func callbackQDBusVirtualObjectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusVirtualObjectCustomEvent
func callbackQDBusVirtualObjectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
