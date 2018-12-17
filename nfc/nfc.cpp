// +build !minimal

#define protected public
#define private public

#include "nfc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFileInfo>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNdefFilter>
#include <QNdefMessage>
#include <QNdefNfcSmartPosterRecord>
#include <QNdefNfcTextRecord>
#include <QNdefNfcUriRecord>
#include <QNdefRecord>
#include <QNearFieldManager>
#include <QNearFieldShareManager>
#include <QNearFieldShareTarget>
#include <QNearFieldTarget>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQmlNdefRecord>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QWidget>
#include <QWindow>

void* QNdefFilter_NewQNdefFilter()
{
	return new QNdefFilter();
}

void* QNdefFilter_NewQNdefFilter2(void* other)
{
	return new QNdefFilter(*static_cast<QNdefFilter*>(other));
}

void QNdefFilter_AppendRecord2(void* ptr, long long typeNameFormat, void* ty, unsigned int min, unsigned int max)
{
	static_cast<QNdefFilter*>(ptr)->appendRecord(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), min, max);
}

void QNdefFilter_Clear(void* ptr)
{
	static_cast<QNdefFilter*>(ptr)->clear();
}

void QNdefFilter_SetOrderMatch(void* ptr, char on)
{
	static_cast<QNdefFilter*>(ptr)->setOrderMatch(on != 0);
}

void QNdefFilter_DestroyQNdefFilter(void* ptr)
{
	static_cast<QNdefFilter*>(ptr)->~QNdefFilter();
}

char QNdefFilter_OrderMatch(void* ptr)
{
	return static_cast<QNdefFilter*>(ptr)->orderMatch();
}

int QNdefFilter_RecordCount(void* ptr)
{
	return static_cast<QNdefFilter*>(ptr)->recordCount();
}

void* QNdefMessage_QNdefMessage_FromByteArray(void* message)
{
	return new QNdefMessage(QNdefMessage::fromByteArray(*static_cast<QByteArray*>(message)));
}

void* QNdefMessage_NewQNdefMessage()
{
	return new QNdefMessage();
}

void* QNdefMessage_NewQNdefMessage4(void* records)
{
	return new QNdefMessage(*static_cast<QList<QNdefRecord>*>(records));
}

void* QNdefMessage_NewQNdefMessage3(void* message)
{
	return new QNdefMessage(*static_cast<QNdefMessage*>(message));
}

void* QNdefMessage_NewQNdefMessage2(void* record)
{
	return new QNdefMessage(*static_cast<QNdefRecord*>(record));
}

void* QNdefMessage_ToByteArray(void* ptr)
{
	return new QByteArray(static_cast<QNdefMessage*>(ptr)->toByteArray());
}

void* QNdefMessage___QNdefMessage_records_atList4(void* ptr, int i)
{
	return new QNdefRecord(({QNdefRecord tmp = static_cast<QList<QNdefRecord>*>(ptr)->at(i); if (i == static_cast<QList<QNdefRecord>*>(ptr)->size()-1) { static_cast<QList<QNdefRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNdefMessage___QNdefMessage_records_setList4(void* ptr, void* i)
{
	static_cast<QList<QNdefRecord>*>(ptr)->append(*static_cast<QNdefRecord*>(i));
}

void* QNdefMessage___QNdefMessage_records_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefRecord>();
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord()
{
	return new QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other)
{
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefNfcSmartPosterRecord*>(other));
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other)
{
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefRecord*>(other));
}

char QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

char QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, struct QtNfc_PackedString text, struct QtNfc_PackedString locale, long long encoding)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(QString::fromUtf8(text.data, text.len), QString::fromUtf8(locale.data, locale.len), static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

char QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, void* ty)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeIcon(*static_cast<QByteArray*>(ty));
}

char QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

char QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, struct QtNfc_PackedString locale)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(QString::fromUtf8(locale.data, locale.len));
}

void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, void* ty, void* data)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addIcon(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(data));
}

void QNdefNfcSmartPosterRecord_SetAction(void* ptr, long long act)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setAction(static_cast<QNdefNfcSmartPosterRecord::Action>(act));
}

void QNdefNfcSmartPosterRecord_SetSize(void* ptr, unsigned int size)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setSize(size);
}

void QNdefNfcSmartPosterRecord_SetTitles(void* ptr, void* titles)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTitles(*static_cast<QList<QNdefNfcTextRecord>*>(titles));
}

void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, void* ty)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTypeInfo(*static_cast<QByteArray*>(ty));
}

void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QNdefNfcUriRecord*>(url));
}

void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QUrl*>(url));
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_Icon(void* ptr, void* mimetype)
{
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->icon(*static_cast<QByteArray*>(mimetype)));
}

void* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr)
{
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->typeInfo());
}

long long QNdefNfcSmartPosterRecord_Action(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->action();
}

void* QNdefNfcSmartPosterRecord_TitleRecord(void* ptr, int index)
{
	return new QNdefNfcTextRecord(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleRecord(index));
}

void* QNdefNfcSmartPosterRecord_UriRecord(void* ptr)
{
	return new QNdefNfcUriRecord(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uriRecord());
}

struct QtNfc_PackedString QNdefNfcSmartPosterRecord_Title(void* ptr, struct QtNfc_PackedString locale)
{
	return ({ QByteArray t928b42 = static_cast<QNdefNfcSmartPosterRecord*>(ptr)->title(QString::fromUtf8(locale.data, locale.len)).toUtf8(); QtNfc_PackedString { const_cast<char*>(t928b42.prepend("WHITESPACE").constData()+10), t928b42.size()-10 }; });
}

void* QNdefNfcSmartPosterRecord_Uri(void* ptr)
{
	return new QUrl(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uri());
}

char QNdefNfcSmartPosterRecord_HasAction(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasAction();
}

char QNdefNfcSmartPosterRecord_HasIcon(void* ptr, void* mimetype)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasIcon(*static_cast<QByteArray*>(mimetype));
}

char QNdefNfcSmartPosterRecord_HasSize(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasSize();
}

char QNdefNfcSmartPosterRecord_HasTitle(void* ptr, struct QtNfc_PackedString locale)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTitle(QString::fromUtf8(locale.data, locale.len));
}

char QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTypeInfo();
}

int QNdefNfcSmartPosterRecord_IconCount(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->iconCount();
}

int QNdefNfcSmartPosterRecord_TitleCount(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleCount();
}

unsigned int QNdefNfcSmartPosterRecord_Size(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->size();
}

void* QNdefNfcSmartPosterRecord___setIcons_icons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefNfcIconRecord>();
}

void* QNdefNfcSmartPosterRecord___setTitles_titles_atList(void* ptr, int i)
{
	return new QNdefNfcTextRecord(({QNdefNfcTextRecord tmp = static_cast<QList<QNdefNfcTextRecord>*>(ptr)->at(i); if (i == static_cast<QList<QNdefNfcTextRecord>*>(ptr)->size()-1) { static_cast<QList<QNdefNfcTextRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNdefNfcSmartPosterRecord___setTitles_titles_setList(void* ptr, void* i)
{
	static_cast<QList<QNdefNfcTextRecord>*>(ptr)->append(*static_cast<QNdefNfcTextRecord*>(i));
}

void* QNdefNfcSmartPosterRecord___setTitles_titles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefNfcTextRecord>();
}

void* QNdefNfcSmartPosterRecord___iconRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefNfcIconRecord>();
}

void* QNdefNfcSmartPosterRecord___titleRecords_atList(void* ptr, int i)
{
	return new QNdefNfcTextRecord(({QNdefNfcTextRecord tmp = static_cast<QList<QNdefNfcTextRecord>*>(ptr)->at(i); if (i == static_cast<QList<QNdefNfcTextRecord>*>(ptr)->size()-1) { static_cast<QList<QNdefNfcTextRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNdefNfcSmartPosterRecord___titleRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QNdefNfcTextRecord>*>(ptr)->append(*static_cast<QNdefNfcTextRecord*>(i));
}

void* QNdefNfcSmartPosterRecord___titleRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefNfcTextRecord>();
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord()
{
	return new QNdefNfcTextRecord();
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other)
{
	return new QNdefNfcTextRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcTextRecord_SetEncoding(void* ptr, long long encoding)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setEncoding(static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

void QNdefNfcTextRecord_SetLocale(void* ptr, struct QtNfc_PackedString locale)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setLocale(QString::fromUtf8(locale.data, locale.len));
}

void QNdefNfcTextRecord_SetText(void* ptr, struct QtNfc_PackedString text)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

long long QNdefNfcTextRecord_Encoding(void* ptr)
{
	return static_cast<QNdefNfcTextRecord*>(ptr)->encoding();
}

struct QtNfc_PackedString QNdefNfcTextRecord_Locale(void* ptr)
{
	return ({ QByteArray tb8fb0c = static_cast<QNdefNfcTextRecord*>(ptr)->locale().toUtf8(); QtNfc_PackedString { const_cast<char*>(tb8fb0c.prepend("WHITESPACE").constData()+10), tb8fb0c.size()-10 }; });
}

struct QtNfc_PackedString QNdefNfcTextRecord_Text(void* ptr)
{
	return ({ QByteArray tbe9455 = static_cast<QNdefNfcTextRecord*>(ptr)->text().toUtf8(); QtNfc_PackedString { const_cast<char*>(tbe9455.prepend("WHITESPACE").constData()+10), tbe9455.size()-10 }; });
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord()
{
	return new QNdefNfcUriRecord();
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other)
{
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(void* ptr, void* uri)
{
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(*static_cast<QUrl*>(uri));
}

void* QNdefNfcUriRecord_Uri(void* ptr)
{
	return new QUrl(static_cast<QNdefNfcUriRecord*>(ptr)->uri());
}

void* QNdefRecord_NewQNdefRecord()
{
	return new QNdefRecord();
}

void* QNdefRecord_NewQNdefRecord2(void* other)
{
	return new QNdefRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefRecord_SetId(void* ptr, void* id)
{
	static_cast<QNdefRecord*>(ptr)->setId(*static_cast<QByteArray*>(id));
}

void QNdefRecord_SetPayload(void* ptr, void* payload)
{
	static_cast<QNdefRecord*>(ptr)->setPayload(*static_cast<QByteArray*>(payload));
}

void QNdefRecord_SetType(void* ptr, void* ty)
{
	static_cast<QNdefRecord*>(ptr)->setType(*static_cast<QByteArray*>(ty));
}

void QNdefRecord_SetTypeNameFormat(void* ptr, long long typeNameFormat)
{
	static_cast<QNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat));
}

void QNdefRecord_DestroyQNdefRecord(void* ptr)
{
	static_cast<QNdefRecord*>(ptr)->~QNdefRecord();
}

void* QNdefRecord_Id(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->id());
}

void* QNdefRecord_Payload(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->payload());
}

void* QNdefRecord_Type(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->type());
}

long long QNdefRecord_TypeNameFormat(void* ptr)
{
	return static_cast<QNdefRecord*>(ptr)->typeNameFormat();
}

char QNdefRecord_IsEmpty(void* ptr)
{
	return static_cast<QNdefRecord*>(ptr)->isEmpty();
}

class MyQNearFieldManager: public QNearFieldManager
{
public:
	MyQNearFieldManager(QObject *parent = Q_NULLPTR) : QNearFieldManager(parent) {QNearFieldManager_QNearFieldManager_QRegisterMetaType();};
	void Signal_TargetDetected(QNearFieldTarget * target) { callbackQNearFieldManager_TargetDetected(this, target); };
	void Signal_TargetLost(QNearFieldTarget * target) { callbackQNearFieldManager_TargetLost(this, target); };
	 ~MyQNearFieldManager() { callbackQNearFieldManager_DestroyQNearFieldManager(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNearFieldManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNearFieldManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNearFieldManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNfc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNearFieldManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldManager_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNearFieldManager*)

int QNearFieldManager_QNearFieldManager_QRegisterMetaType(){qRegisterMetaType<QNearFieldManager*>(); return qRegisterMetaType<MyQNearFieldManager*>();}

void* QNearFieldManager_NewQNearFieldManager(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldManager(static_cast<QWindow*>(parent));
	} else {
		return new MyQNearFieldManager(static_cast<QObject*>(parent));
	}
}

struct QtNfc_PackedString QNearFieldManager_QNearFieldManager_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t9661e6 = QNearFieldManager::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t9661e6.prepend("WHITESPACE").constData()+10), t9661e6.size()-10 }; });
}

struct QtNfc_PackedString QNearFieldManager_QNearFieldManager_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t0e7de3 = QNearFieldManager::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t0e7de3.prepend("WHITESPACE").constData()+10), t0e7de3.size()-10 }; });
}

char QNearFieldManager_StartTargetDetection(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->startTargetDetection();
}

char QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId)
{
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, long long typeNameFormat, void* ty, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(*static_cast<QNdefFilter*>(filter), static_cast<QObject*>(object), const_cast<const char*>(method));
}

void QNearFieldManager_SetTargetAccessModes(void* ptr, long long accessModes)
{
	static_cast<QNearFieldManager*>(ptr)->setTargetAccessModes(static_cast<QNearFieldManager::TargetAccessMode>(accessModes));
}

void QNearFieldManager_StopTargetDetection(void* ptr)
{
	static_cast<QNearFieldManager*>(ptr)->stopTargetDetection();
}

void QNearFieldManager_ConnectTargetDetected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));
}

void QNearFieldManager_DisconnectTargetDetected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));
}

void QNearFieldManager_TargetDetected(void* ptr, void* target)
{
	static_cast<QNearFieldManager*>(ptr)->targetDetected(static_cast<QNearFieldTarget*>(target));
}

void QNearFieldManager_ConnectTargetLost(void* ptr)
{
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));
}

void QNearFieldManager_DisconnectTargetLost(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));
}

void QNearFieldManager_TargetLost(void* ptr, void* target)
{
	static_cast<QNearFieldManager*>(ptr)->targetLost(static_cast<QNearFieldTarget*>(target));
}

void QNearFieldManager_DestroyQNearFieldManager(void* ptr)
{
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

void QNearFieldManager_DestroyQNearFieldManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QNearFieldManager_TargetAccessModes(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->targetAccessModes();
}

char QNearFieldManager_IsAvailable(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->isAvailable();
}

char QNearFieldManager_IsSupported(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->isSupported();
}

void* QNearFieldManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::metaObject());
}

void* QNearFieldManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNearFieldManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNearFieldManager___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldManager___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldManager___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldManager___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNearFieldManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::event(static_cast<QEvent*>(e));
}

char QNearFieldManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNearFieldManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::deleteLater();
}

void QNearFieldManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQNearFieldShareManager: public QNearFieldShareManager
{
public:
	MyQNearFieldShareManager(QObject *parent = Q_NULLPTR) : QNearFieldShareManager(parent) {QNearFieldShareManager_QNearFieldShareManager_QRegisterMetaType();};
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareManager_Error(this, error); };
	void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes) { callbackQNearFieldShareManager_ShareModesChanged(this, modes); };
	void Signal_TargetDetected(QNearFieldShareTarget * shareTarget) { callbackQNearFieldShareManager_TargetDetected(this, shareTarget); };
	 ~MyQNearFieldShareManager() { callbackQNearFieldShareManager_DestroyQNearFieldShareManager(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldShareManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNearFieldShareManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldShareManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldShareManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldShareManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldShareManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNearFieldShareManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldShareManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNfc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNearFieldShareManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareManager_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNearFieldShareManager*)

int QNearFieldShareManager_QNearFieldShareManager_QRegisterMetaType(){qRegisterMetaType<QNearFieldShareManager*>(); return qRegisterMetaType<MyQNearFieldShareManager*>();}

void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldShareManager(static_cast<QWindow*>(parent));
	} else {
		return new MyQNearFieldShareManager(static_cast<QObject*>(parent));
	}
}

long long QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes()
{
	return QNearFieldShareManager::supportedShareModes();
}

struct QtNfc_PackedString QNearFieldShareManager_QNearFieldShareManager_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t891e6c = QNearFieldShareManager::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t891e6c.prepend("WHITESPACE").constData()+10), t891e6c.size()-10 }; });
}

struct QtNfc_PackedString QNearFieldShareManager_QNearFieldShareManager_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray taafa97 = QNearFieldShareManager::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(taafa97.prepend("WHITESPACE").constData()+10), taafa97.size()-10 }; });
}

void QNearFieldShareManager_ConnectError(void* ptr)
{
	qRegisterMetaType<QNearFieldShareManager::ShareError>();
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));
}

void QNearFieldShareManager_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));
}

void QNearFieldShareManager_Error(void* ptr, long long error)
{
	static_cast<QNearFieldShareManager*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
}

void QNearFieldShareManager_SetShareModes(void* ptr, long long mode)
{
	static_cast<QNearFieldShareManager*>(ptr)->setShareModes(static_cast<QNearFieldShareManager::ShareMode>(mode));
}

void QNearFieldShareManager_ConnectShareModesChanged(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));
}

void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));
}

void QNearFieldShareManager_ShareModesChanged(void* ptr, long long modes)
{
	static_cast<QNearFieldShareManager*>(ptr)->shareModesChanged(static_cast<QNearFieldShareManager::ShareMode>(modes));
}

void QNearFieldShareManager_ConnectTargetDetected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));
}

void QNearFieldShareManager_DisconnectTargetDetected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));
}

void QNearFieldShareManager_TargetDetected(void* ptr, void* shareTarget)
{
	static_cast<QNearFieldShareManager*>(ptr)->targetDetected(static_cast<QNearFieldShareTarget*>(shareTarget));
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr)
{
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

void QNearFieldShareManager_DestroyQNearFieldShareManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QNearFieldShareManager_ShareError(void* ptr)
{
	return static_cast<QNearFieldShareManager*>(ptr)->shareError();
}

long long QNearFieldShareManager_ShareModes(void* ptr)
{
	return static_cast<QNearFieldShareManager*>(ptr)->shareModes();
}

void* QNearFieldShareManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::metaObject());
}

void* QNearFieldShareManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldShareManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNearFieldShareManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNearFieldShareManager___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareManager___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareManager___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareManager___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNearFieldShareManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::event(static_cast<QEvent*>(e));
}

char QNearFieldShareManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNearFieldShareManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::deleteLater();
}

void QNearFieldShareManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQNearFieldShareTarget: public QNearFieldShareTarget
{
public:
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareTarget_Error(this, error); };
	void Signal_ShareFinished() { callbackQNearFieldShareTarget_ShareFinished(this); };
	 ~MyQNearFieldShareTarget() { callbackQNearFieldShareTarget_DestroyQNearFieldShareTarget(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldShareTarget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNearFieldShareTarget_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldShareTarget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareTarget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldShareTarget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldShareTarget_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldShareTarget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNearFieldShareTarget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldShareTarget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNfc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNearFieldShareTarget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareTarget_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNearFieldShareTarget*)

int QNearFieldShareTarget_QNearFieldShareTarget_QRegisterMetaType(){qRegisterMetaType<QNearFieldShareTarget*>(); return qRegisterMetaType<MyQNearFieldShareTarget*>();}

struct QtNfc_PackedString QNearFieldShareTarget_QNearFieldShareTarget_Tr(char* s, char* c, int n)
{
	return ({ QByteArray te76c4f = QNearFieldShareTarget::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(te76c4f.prepend("WHITESPACE").constData()+10), te76c4f.size()-10 }; });
}

struct QtNfc_PackedString QNearFieldShareTarget_QNearFieldShareTarget_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t939e17 = QNearFieldShareTarget::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t939e17.prepend("WHITESPACE").constData()+10), t939e17.size()-10 }; });
}

char QNearFieldShareTarget_Share2(void* ptr, void* files)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->share(*static_cast<QList<QFileInfo>*>(files));
}

char QNearFieldShareTarget_Share(void* ptr, void* message)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->share(*static_cast<QNdefMessage*>(message));
}

void QNearFieldShareTarget_Cancel(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->cancel();
}

void QNearFieldShareTarget_ConnectError(void* ptr)
{
	qRegisterMetaType<QNearFieldShareManager::ShareError>();
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));
}

void QNearFieldShareTarget_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));
}

void QNearFieldShareTarget_Error(void* ptr, long long error)
{
	static_cast<QNearFieldShareTarget*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
}

void QNearFieldShareTarget_ConnectShareFinished(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));
}

void QNearFieldShareTarget_DisconnectShareFinished(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));
}

void QNearFieldShareTarget_ShareFinished(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->shareFinished();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->~QNearFieldShareTarget();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTargetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QNearFieldShareTarget_ShareError(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->shareError();
}

long long QNearFieldShareTarget_ShareModes(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->shareModes();
}

char QNearFieldShareTarget_IsShareInProgress(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->isShareInProgress();
}

void* QNearFieldShareTarget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::metaObject());
}

void* QNearFieldShareTarget___share_files_atList2(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldShareTarget___share_files_setList2(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QNearFieldShareTarget___share_files_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QNearFieldShareTarget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldShareTarget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNearFieldShareTarget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNearFieldShareTarget___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareTarget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareTarget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareTarget___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareTarget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareTarget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareTarget___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareTarget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareTarget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldShareTarget___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldShareTarget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldShareTarget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNearFieldShareTarget_EventDefault(void* ptr, void* e)
{
		return static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::event(static_cast<QEvent*>(e));
}

char QNearFieldShareTarget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNearFieldShareTarget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareTarget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareTarget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareTarget_DeleteLaterDefault(void* ptr)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::deleteLater();
}

void QNearFieldShareTarget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareTarget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQNearFieldTarget: public QNearFieldTarget
{
public:
	MyQNearFieldTarget(QObject *parent = Q_NULLPTR) : QNearFieldTarget(parent) {QNearFieldTarget_QNearFieldTarget_QRegisterMetaType();};
	bool hasNdefMessage() { return callbackQNearFieldTarget_HasNdefMessage(this) != 0; };
	void Signal_Disconnected() { callbackQNearFieldTarget_Disconnected(this); };
	void Signal_NdefMessageRead(const QNdefMessage & message) { callbackQNearFieldTarget_NdefMessageRead(this, const_cast<QNdefMessage*>(&message)); };
	void Signal_NdefMessagesWritten() { callbackQNearFieldTarget_NdefMessagesWritten(this); };
	 ~MyQNearFieldTarget() { callbackQNearFieldTarget_DestroyQNearFieldTarget(this); };
	QByteArray uid() const { return *static_cast<QByteArray*>(callbackQNearFieldTarget_Uid(const_cast<void*>(static_cast<const void*>(this)))); };
	QNearFieldTarget::AccessMethods accessMethods() const { return static_cast<QNearFieldTarget::AccessMethod>(callbackQNearFieldTarget_AccessMethods(const_cast<void*>(static_cast<const void*>(this)))); };
	QNearFieldTarget::Type type() const { return static_cast<QNearFieldTarget::Type>(callbackQNearFieldTarget_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	QUrl url() const { return *static_cast<QUrl*>(callbackQNearFieldTarget_Url(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldTarget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNearFieldTarget_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldTarget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNearFieldTarget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldTarget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldTarget_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldTarget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNearFieldTarget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldTarget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNfc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNearFieldTarget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldTarget_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNearFieldTarget*)

int QNearFieldTarget_QNearFieldTarget_QRegisterMetaType(){qRegisterMetaType<QNearFieldTarget*>(); return qRegisterMetaType<MyQNearFieldTarget*>();}

void* QNearFieldTarget_NewQNearFieldTarget(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNearFieldTarget(static_cast<QWindow*>(parent));
	} else {
		return new MyQNearFieldTarget(static_cast<QObject*>(parent));
	}
}

struct QtNfc_PackedString QNearFieldTarget_QNearFieldTarget_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t6b5b27 = QNearFieldTarget::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t6b5b27.prepend("WHITESPACE").constData()+10), t6b5b27.size()-10 }; });
}

struct QtNfc_PackedString QNearFieldTarget_QNearFieldTarget_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray td71f90 = QNearFieldTarget::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(td71f90.prepend("WHITESPACE").constData()+10), td71f90.size()-10 }; });
}

char QNearFieldTarget_Disconnect(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->disconnect();
}

char QNearFieldTarget_HasNdefMessage(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->hasNdefMessage();
}

char QNearFieldTarget_HasNdefMessageDefault(void* ptr)
{
		return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::hasNdefMessage();
}

char QNearFieldTarget_SetKeepConnection(void* ptr, char isPersistent)
{
	return static_cast<QNearFieldTarget*>(ptr)->setKeepConnection(isPersistent != 0);
}

void QNearFieldTarget_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));
}

void QNearFieldTarget_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));
}

void QNearFieldTarget_Disconnected(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->disconnected();
}

void QNearFieldTarget_ConnectNdefMessageRead(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)(const QNdefMessage &)>(&QNearFieldTarget::ndefMessageRead), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)(const QNdefMessage &)>(&MyQNearFieldTarget::Signal_NdefMessageRead));
}

void QNearFieldTarget_DisconnectNdefMessageRead(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)(const QNdefMessage &)>(&QNearFieldTarget::ndefMessageRead), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)(const QNdefMessage &)>(&MyQNearFieldTarget::Signal_NdefMessageRead));
}

void QNearFieldTarget_NdefMessageRead(void* ptr, void* message)
{
	static_cast<QNearFieldTarget*>(ptr)->ndefMessageRead(*static_cast<QNdefMessage*>(message));
}

void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));
}

void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));
}

void QNearFieldTarget_NdefMessagesWritten(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->ndefMessagesWritten();
}

void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

void QNearFieldTarget_DestroyQNearFieldTargetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNearFieldTarget_Uid(void* ptr)
{
	return new QByteArray(static_cast<QNearFieldTarget*>(ptr)->uid());
}

long long QNearFieldTarget_AccessMethods(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->accessMethods();
}

long long QNearFieldTarget_Type(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

void* QNearFieldTarget_Url(void* ptr)
{
	return new QUrl(static_cast<QNearFieldTarget*>(ptr)->url());
}

void* QNearFieldTarget_UrlDefault(void* ptr)
{
		return new QUrl(static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::url());
}

char QNearFieldTarget_IsProcessingCommand(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->isProcessingCommand();
}

char QNearFieldTarget_KeepConnection(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->keepConnection();
}

void* QNearFieldTarget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::metaObject());
}

int QNearFieldTarget_MaxCommandLength(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->maxCommandLength();
}

void* QNearFieldTarget___sendCommands_commands_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldTarget___sendCommands_commands_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNearFieldTarget___sendCommands_commands_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNearFieldTarget___writeNdefMessages_messages_atList(void* ptr, int i)
{
	return new QNdefMessage(({QNdefMessage tmp = static_cast<QList<QNdefMessage>*>(ptr)->at(i); if (i == static_cast<QList<QNdefMessage>*>(ptr)->size()-1) { static_cast<QList<QNdefMessage>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldTarget___writeNdefMessages_messages_setList(void* ptr, void* i)
{
	static_cast<QList<QNdefMessage>*>(ptr)->append(*static_cast<QNdefMessage*>(i));
}

void* QNearFieldTarget___writeNdefMessages_messages_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNdefMessage>();
}

void* QNearFieldTarget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNearFieldTarget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNearFieldTarget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNearFieldTarget___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldTarget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldTarget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldTarget___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldTarget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldTarget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldTarget___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldTarget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldTarget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNearFieldTarget___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNearFieldTarget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNearFieldTarget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNearFieldTarget_EventDefault(void* ptr, void* e)
{
		return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::event(static_cast<QEvent*>(e));
}

char QNearFieldTarget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNearFieldTarget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldTarget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldTarget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldTarget_DeleteLaterDefault(void* ptr)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::deleteLater();
}

void QNearFieldTarget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldTarget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlNdefRecord: public QQmlNdefRecord
{
public:
	MyQQmlNdefRecord(QObject *parent = Q_NULLPTR) : QQmlNdefRecord(parent) {QQmlNdefRecord_QQmlNdefRecord_QRegisterMetaType();};
	MyQQmlNdefRecord(const QNdefRecord &record, QObject *parent = Q_NULLPTR) : QQmlNdefRecord(record, parent) {QQmlNdefRecord_QQmlNdefRecord_QRegisterMetaType();};
	void Signal_RecordChanged() { callbackQQmlNdefRecord_RecordChanged(this); };
	void Signal_TypeChanged() { callbackQQmlNdefRecord_TypeChanged(this); };
	void Signal_TypeNameFormatChanged() { callbackQQmlNdefRecord_TypeNameFormatChanged(this); };
	 ~MyQQmlNdefRecord() { callbackQQmlNdefRecord_DestroyQQmlNdefRecord(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlNdefRecord_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQQmlNdefRecord_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlNdefRecord_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQmlNdefRecord_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlNdefRecord_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlNdefRecord_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlNdefRecord_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlNdefRecord_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlNdefRecord_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNfc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQmlNdefRecord_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlNdefRecord_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQmlNdefRecord*)

int QQmlNdefRecord_QQmlNdefRecord_QRegisterMetaType(){qRegisterMetaType<QQmlNdefRecord*>(); return qRegisterMetaType<MyQQmlNdefRecord*>();}

void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlNdefRecord(static_cast<QObject*>(parent));
	}
}

void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QObject*>(parent));
	}
}

struct QtNfc_PackedString QQmlNdefRecord_QQmlNdefRecord_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t6484f5 = QQmlNdefRecord::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(t6484f5.prepend("WHITESPACE").constData()+10), t6484f5.size()-10 }; });
}

struct QtNfc_PackedString QQmlNdefRecord_QQmlNdefRecord_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tf336ee = QQmlNdefRecord::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNfc_PackedString { const_cast<char*>(tf336ee.prepend("WHITESPACE").constData()+10), tf336ee.size()-10 }; });
}

void QQmlNdefRecord_ConnectRecordChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));
}

void QQmlNdefRecord_DisconnectRecordChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));
}

void QQmlNdefRecord_RecordChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->recordChanged();
}

void QQmlNdefRecord_SetRecord(void* ptr, void* record)
{
	static_cast<QQmlNdefRecord*>(ptr)->setRecord(*static_cast<QNdefRecord*>(record));
}

void QQmlNdefRecord_SetType(void* ptr, struct QtNfc_PackedString newtype)
{
	static_cast<QQmlNdefRecord*>(ptr)->setType(QString::fromUtf8(newtype.data, newtype.len));
}

void QQmlNdefRecord_SetTypeNameFormat(void* ptr, long long newTypeNameFormat)
{
	static_cast<QQmlNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QQmlNdefRecord::TypeNameFormat>(newTypeNameFormat));
}

void QQmlNdefRecord_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));
}

void QQmlNdefRecord_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));
}

void QQmlNdefRecord_TypeChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->typeChanged();
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));
}

void QQmlNdefRecord_TypeNameFormatChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->typeNameFormatChanged();
}

void QQmlNdefRecord_DestroyQQmlNdefRecord(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->~QQmlNdefRecord();
}

void QQmlNdefRecord_DestroyQQmlNdefRecordDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlNdefRecord_Record(void* ptr)
{
	return new QNdefRecord(static_cast<QQmlNdefRecord*>(ptr)->record());
}

long long QQmlNdefRecord_TypeNameFormat(void* ptr)
{
	return static_cast<QQmlNdefRecord*>(ptr)->typeNameFormat();
}

struct QtNfc_PackedString QQmlNdefRecord_Type(void* ptr)
{
	return ({ QByteArray t9a31a1 = static_cast<QQmlNdefRecord*>(ptr)->type().toUtf8(); QtNfc_PackedString { const_cast<char*>(t9a31a1.prepend("WHITESPACE").constData()+10), t9a31a1.size()-10 }; });
}

void* QQmlNdefRecord_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::metaObject());
}

void* QQmlNdefRecord___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlNdefRecord___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlNdefRecord___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlNdefRecord___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlNdefRecord___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlNdefRecord___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlNdefRecord___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlNdefRecord___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlNdefRecord___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlNdefRecord___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlNdefRecord___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlNdefRecord___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlNdefRecord___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlNdefRecord___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlNdefRecord___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QQmlNdefRecord_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::event(static_cast<QEvent*>(e));
}

char QQmlNdefRecord_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QQmlNdefRecord_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlNdefRecord_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlNdefRecord_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::customEvent(static_cast<QEvent*>(event));
}

void QQmlNdefRecord_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::deleteLater();
}

void QQmlNdefRecord_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlNdefRecord_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::timerEvent(static_cast<QTimerEvent*>(event));
}

