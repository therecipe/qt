package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandler_PTR().Pointer()
	}
	return nil
}

func NewQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) *QAbstractMessageHandler {
	var n = new(QAbstractMessageHandler)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractMessageHandler_") {
		n.SetObjectName("QAbstractMessageHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	defer qt.Recovering("QAbstractMessageHandler::~QAbstractMessageHandler")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractMessageHandlerTimerEvent
func callbackQAbstractMessageHandlerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractMessageHandlerChildEvent
func callbackQAbstractMessageHandlerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractMessageHandlerCustomEvent
func callbackQAbstractMessageHandlerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractUriResolver_") {
		n.SetObjectName("QAbstractUriResolver_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractUriResolver::resolve")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QAbstractUriResolver_Resolve(ptr.Pointer(), core.PointerFromQUrl(relative), core.PointerFromQUrl(baseURI)))
	}
	return nil
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	defer qt.Recovering("QAbstractUriResolver::~QAbstractUriResolver")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractUriResolverTimerEvent
func callbackQAbstractUriResolverTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractUriResolverChildEvent
func callbackQAbstractUriResolverChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractUriResolverCustomEvent
func callbackQAbstractUriResolverCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) *QAbstractXmlNodeModel {
	var n = new(QAbstractXmlNodeModel)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractXmlNodeModel_") {
		n.SetObjectNameAbs("QAbstractXmlNodeModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return ptr
}

//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     = QAbstractXmlNodeModel__SimpleAxis(3)
)

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	defer qt.Recovering("QAbstractXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	defer qt.Recovering("QAbstractXmlNodeModel::kind")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QAbstractXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	defer qt.Recovering("QAbstractXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	defer qt.Recovering("QAbstractXmlNodeModel::~QAbstractXmlNodeModel")

	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlNodeModel) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractXmlNodeModel::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractXmlNodeModel::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (p *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractXmlReceiver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) *QAbstractXmlReceiver {
	var n = new(QAbstractXmlReceiver)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractXmlReceiver_") {
		n.SetObjectNameAbs("QAbstractXmlReceiver_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::attribute")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::characters")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	defer qt.Recovering("QAbstractXmlReceiver::comment")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	defer qt.Recovering("QAbstractXmlReceiver::endDocument")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	defer qt.Recovering("QAbstractXmlReceiver::endElement")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	defer qt.Recovering("QAbstractXmlReceiver::endOfSequence")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {
	defer qt.Recovering("QAbstractXmlReceiver::processingInstruction")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(target), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	defer qt.Recovering("QAbstractXmlReceiver::startDocument")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::startElement")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	defer qt.Recovering("QAbstractXmlReceiver::startOfSequence")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	defer qt.Recovering("QAbstractXmlReceiver::~QAbstractXmlReceiver")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractXmlReceiver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlReceiver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractXmlReceiver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractXmlReceiver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) *QSimpleXmlNodeModel {
	var n = new(QSimpleXmlNodeModel)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSimpleXmlNodeModel_") {
		n.SetObjectNameAbs("QSimpleXmlNodeModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	defer qt.Recovering("QSimpleXmlNodeModel::~QSimpleXmlNodeModel")

	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
	}
}

func (ptr *QSimpleXmlNodeModel) ObjectNameAbs() string {
	defer qt.Recovering("QSimpleXmlNodeModel::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSimpleXmlNodeModel::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSourceLocation struct {
	ptr unsafe.Pointer
}

type QSourceLocation_ITF interface {
	QSourceLocation_PTR() *QSourceLocation
}

func (p *QSourceLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSourceLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSourceLocation(ptr QSourceLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocation_PTR().Pointer()
	}
	return nil
}

func NewQSourceLocationFromPointer(ptr unsafe.Pointer) *QSourceLocation {
	var n = new(QSourceLocation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return ptr
}

func NewQSourceLocation() *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation())
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(l), C.int(c)))
}

func (ptr *QSourceLocation) Column() int64 {
	defer qt.Recovering("QSourceLocation::column")

	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) IsNull() bool {
	defer qt.Recovering("QSourceLocation::isNull")

	if ptr.Pointer() != nil {
		return C.QSourceLocation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSourceLocation) Line() int64 {
	defer qt.Recovering("QSourceLocation::line")

	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) SetColumn(newColumn int64) {
	defer qt.Recovering("QSourceLocation::setColumn")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetColumn(ptr.Pointer(), C.longlong(newColumn))
	}
}

func (ptr *QSourceLocation) SetLine(newLine int64) {
	defer qt.Recovering("QSourceLocation::setLine")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetLine(ptr.Pointer(), C.longlong(newLine))
	}
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {
	defer qt.Recovering("QSourceLocation::setUri")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(ptr.Pointer(), core.PointerFromQUrl(newUri))
	}
}

func (ptr *QSourceLocation) Uri() *core.QUrl {
	defer qt.Recovering("QSourceLocation::uri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QSourceLocation_Uri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	defer qt.Recovering("QSourceLocation::~QSourceLocation")

	if ptr.Pointer() != nil {
		C.QSourceLocation_DestroyQSourceLocation(ptr.Pointer())
	}
}

type QXmlFormatter struct {
	QXmlSerializer
}

type QXmlFormatter_ITF interface {
	QXmlSerializer_ITF
	QXmlFormatter_PTR() *QXmlFormatter
}

func PointerFromQXmlFormatter(ptr QXmlFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatter_PTR().Pointer()
	}
	return nil
}

func NewQXmlFormatterFromPointer(ptr unsafe.Pointer) *QXmlFormatter {
	var n = new(QXmlFormatter)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlFormatter_") {
		n.SetObjectNameAbs("QXmlFormatter_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	defer qt.Recovering("QXmlFormatter::QXmlFormatter")

	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlFormatter) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

//export callbackQXmlFormatterCharacters
func callbackQXmlFormatterCharacters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::characters")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) CharactersDefault(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::characters")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlFormatter) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

//export callbackQXmlFormatterComment
func callbackQXmlFormatterComment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QXmlFormatter::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CommentDefault(C.GoString(value))
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	defer qt.Recovering("QXmlFormatter::comment")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlFormatter) CommentDefault(value string) {
	defer qt.Recovering("QXmlFormatter::comment")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_CommentDefault(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlFormatter) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

//export callbackQXmlFormatterEndDocument
func callbackQXmlFormatterEndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	defer qt.Recovering("QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndDocumentDefault() {
	defer qt.Recovering("QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

//export callbackQXmlFormatterEndElement
func callbackQXmlFormatterEndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlFormatter) EndElement() {
	defer qt.Recovering("QXmlFormatter::endElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndElementDefault() {
	defer qt.Recovering("QXmlFormatter::endElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElementDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

//export callbackQXmlFormatterEndOfSequence
func callbackQXmlFormatterEndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	defer qt.Recovering("QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndOfSequenceDefault() {
	defer qt.Recovering("QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	defer qt.Recovering("QXmlFormatter::indentationDepth")

	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	defer qt.Recovering("QXmlFormatter::setIndentationDepth")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QXmlFormatter) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

//export callbackQXmlFormatterStartDocument
func callbackQXmlFormatterStartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	defer qt.Recovering("QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartDocumentDefault() {
	defer qt.Recovering("QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

//export callbackQXmlFormatterStartOfSequence
func callbackQXmlFormatterStartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	defer qt.Recovering("QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartOfSequenceDefault() {
	defer qt.Recovering("QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) ObjectNameAbs() string {
	defer qt.Recovering("QXmlFormatter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlFormatter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlFormatter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlFormatter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlItem struct {
	ptr unsafe.Pointer
}

type QXmlItem_ITF interface {
	QXmlItem_PTR() *QXmlItem
}

func (p *QXmlItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlItem(ptr QXmlItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItem_PTR().Pointer()
	}
	return nil
}

func NewQXmlItemFromPointer(ptr unsafe.Pointer) *QXmlItem {
	var n = new(QXmlItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return ptr
}

func NewQXmlItem() *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
}

func (ptr *QXmlItem) IsNode() bool {
	defer qt.Recovering("QXmlItem::isNode")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	defer qt.Recovering("QXmlItem::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	defer qt.Recovering("QXmlItem::~QXmlItem")

	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
	}
}

type QXmlName struct {
	ptr unsafe.Pointer
}

type QXmlName_ITF interface {
	QXmlName_PTR() *QXmlName
}

func (p *QXmlName) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlName) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlName(ptr QXmlName_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlName_PTR().Pointer()
	}
	return nil
}

func NewQXmlNameFromPointer(ptr unsafe.Pointer) *QXmlName {
	var n = new(QXmlName)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlName) QXmlName_PTR() *QXmlName {
	return ptr
}

func NewQXmlName() *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	return NewQXmlNameFromPointer(C.QXmlName_NewQXmlName())
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	return NewQXmlNameFromPointer(C.QXmlName_NewQXmlName2(PointerFromQXmlNamePool(namePool), C.CString(localName), C.CString(namespaceURI), C.CString(prefix)))
}

func QXmlName_IsNCName(candidate string) bool {
	defer qt.Recovering("QXmlName::isNCName")

	return C.QXmlName_QXmlName_IsNCName(C.CString(candidate)) != 0
}

func (ptr *QXmlName) IsNull() bool {
	defer qt.Recovering("QXmlName::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlName_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_LocalName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::namespaceUri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_NamespaceUri(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_Prefix(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::toClarkName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_ToClarkName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (p *QXmlNamePool) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamePool) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamePool(ptr QXmlNamePool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePool_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamePoolFromPointer(ptr unsafe.Pointer) *QXmlNamePool {
	var n = new(QXmlNamePool)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return ptr
}

func NewQXmlNamePool() *QXmlNamePool {
	defer qt.Recovering("QXmlNamePool::QXmlNamePool")

	return NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool())
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	defer qt.Recovering("QXmlNamePool::QXmlNamePool")

	return NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	defer qt.Recovering("QXmlNamePool::~QXmlNamePool")

	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
	}
}

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (p *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNodeModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) *QXmlNodeModelIndex {
	var n = new(QXmlNodeModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return ptr
}

//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  = QXmlNodeModelIndex__DocumentOrder(1)
)

//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  = QXmlNodeModelIndex__NodeKind(64)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	defer qt.Recovering("QXmlNodeModelIndex::QXmlNodeModelIndex")

	return NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QXmlNodeModelIndex::QXmlNodeModelIndex")

	return NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
}

func (ptr *QXmlNodeModelIndex) AdditionalData() int64 {
	defer qt.Recovering("QXmlNodeModelIndex::additionalData")

	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_AdditionalData(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) Data() int64 {
	defer qt.Recovering("QXmlNodeModelIndex::data")

	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {
	defer qt.Recovering("QXmlNodeModelIndex::internalPointer")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlNodeModelIndex_InternalPointer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	defer qt.Recovering("QXmlNodeModelIndex::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	defer qt.Recovering("QXmlNodeModelIndex::model")

	if ptr.Pointer() != nil {
		return NewQAbstractXmlNodeModelFromPointer(C.QXmlNodeModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

type QXmlQuery struct {
	ptr unsafe.Pointer
}

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (p *QXmlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlQuery(ptr QXmlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQuery_PTR().Pointer()
	}
	return nil
}

func NewQXmlQueryFromPointer(ptr unsafe.Pointer) *QXmlQuery {
	var n = new(QXmlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return ptr
}

//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery())
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.int(queryLanguage), PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
}

func (ptr *QXmlQuery) BindVariable5(localName string, device core.QIODevice_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), C.CString(localName), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable4(localName string, value QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable4(ptr.Pointer(), C.CString(localName), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable6(ptr.Pointer(), C.CString(localName), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) BindVariable2(name QXmlName_ITF, device core.QIODevice_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable2(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, query QXmlQuery_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) IsValid() bool {
	defer qt.Recovering("QXmlQuery::isValid")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlQuery::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlQuery_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlQuery::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlQuery_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	defer qt.Recovering("QXmlQuery::queryLanguage")

	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus3(ptr.Pointer(), core.PointerFromQIODevice(document)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus4(ptr.Pointer(), C.CString(focus)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	defer qt.Recovering("QXmlQuery::setInitialTemplateName")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName2(ptr.Pointer(), C.CString(localName))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	defer qt.Recovering("QXmlQuery::setInitialTemplateName")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlQuery::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(aMessageHandler))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlQuery::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(newManager))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(ptr.Pointer(), core.PointerFromQIODevice(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(sourceCode string, documentURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(ptr.Pointer(), C.CString(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery2(ptr.Pointer(), core.PointerFromQUrl(queryURI), core.PointerFromQUrl(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlQuery::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlQuery::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlQuery_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	defer qt.Recovering("QXmlQuery::~QXmlQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_DestroyQXmlQuery(ptr.Pointer())
	}
}

type QXmlResultItems struct {
	ptr unsafe.Pointer
}

type QXmlResultItems_ITF interface {
	QXmlResultItems_PTR() *QXmlResultItems
}

func (p *QXmlResultItems) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlResultItems) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlResultItems(ptr QXmlResultItems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItems_PTR().Pointer()
	}
	return nil
}

func NewQXmlResultItemsFromPointer(ptr unsafe.Pointer) *QXmlResultItems {
	var n = new(QXmlResultItems)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlResultItems_") {
		n.SetObjectNameAbs("QXmlResultItems_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func NewQXmlResultItems() *QXmlResultItems {
	defer qt.Recovering("QXmlResultItems::QXmlResultItems")

	return NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

func (ptr *QXmlResultItems) HasError() bool {
	defer qt.Recovering("QXmlResultItems::hasError")

	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	defer qt.Recovering("QXmlResultItems::~QXmlResultItems")

	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
	}
}

func (ptr *QXmlResultItems) ObjectNameAbs() string {
	defer qt.Recovering("QXmlResultItems::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlResultItems_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlResultItems) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlResultItems::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlResultItems_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (p *QXmlSchema) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchema) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchema(ptr QXmlSchema_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchema_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaFromPointer(ptr unsafe.Pointer) *QXmlSchema {
	var n = new(QXmlSchema)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return ptr
}

func NewQXmlSchema() *QXmlSchema {
	defer qt.Recovering("QXmlSchema::QXmlSchema")

	return NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema())
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	defer qt.Recovering("QXmlSchema::QXmlSchema")

	return NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
}

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {
	defer qt.Recovering("QXmlSchema::documentUri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QXmlSchema_DocumentUri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) IsValid() bool {
	defer qt.Recovering("QXmlSchema::isValid")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlSchema::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchema_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlSchema::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchema_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlSchema::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlSchema::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlSchema::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlSchema::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchema_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	defer qt.Recovering("QXmlSchema::~QXmlSchema")

	if ptr.Pointer() != nil {
		C.QXmlSchema_DestroyQXmlSchema(ptr.Pointer())
	}
}

type QXmlSchemaValidator struct {
	ptr unsafe.Pointer
}

type QXmlSchemaValidator_ITF interface {
	QXmlSchemaValidator_PTR() *QXmlSchemaValidator
}

func (p *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchemaValidator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidator_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) *QXmlSchemaValidator {
	var n = new(QXmlSchemaValidator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return ptr
}

func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	defer qt.Recovering("QXmlSchemaValidator::QXmlSchemaValidator")

	return NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator())
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	defer qt.Recovering("QXmlSchemaValidator::QXmlSchemaValidator")

	return NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlSchemaValidator::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchemaValidator_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlSchemaValidator::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchemaValidator_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setSchema")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(ptr.Pointer(), PointerFromQXmlSchema(schema))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlSchemaValidator::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchemaValidator_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	defer qt.Recovering("QXmlSchemaValidator::~QXmlSchemaValidator")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(ptr.Pointer())
	}
}

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializer_ITF interface {
	QAbstractXmlReceiver_ITF
	QXmlSerializer_PTR() *QXmlSerializer
}

func PointerFromQXmlSerializer(ptr QXmlSerializer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func NewQXmlSerializerFromPointer(ptr unsafe.Pointer) *QXmlSerializer {
	var n = new(QXmlSerializer)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlSerializer_") {
		n.SetObjectNameAbs("QXmlSerializer_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	defer qt.Recovering("QXmlSerializer::QXmlSerializer")

	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlSerializer) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

//export callbackQXmlSerializerCharacters
func callbackQXmlSerializerCharacters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlSerializer) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

//export callbackQXmlSerializerComment
func callbackQXmlSerializerComment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QXmlSerializer::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	}

}

func (ptr *QXmlSerializer) Comment(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlSerializer) CommentDefault(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_CommentDefault(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

//export callbackQXmlSerializerEndDocument
func callbackQXmlSerializerEndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndDocument() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

//export callbackQXmlSerializerEndElement
func callbackQXmlSerializerEndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndElement() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElementDefault() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElementDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	defer qt.Recovering("QXmlSerializer::codec")

	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

//export callbackQXmlSerializerEndOfSequence
func callbackQXmlSerializerEndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndOfSequence() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QXmlSerializer::outputDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	defer qt.Recovering("QXmlSerializer::setCodec")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

//export callbackQXmlSerializerStartDocument
func callbackQXmlSerializerStartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) StartDocument() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

//export callbackQXmlSerializerStartOfSequence
func callbackQXmlSerializerStartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) StartOfSequence() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSerializer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSerializer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSerializer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSerializer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
