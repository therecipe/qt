// +build !minimal

#define protected public
#define private public

#include "purchasing.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QInAppProduct>
#include <QInAppStore>
#include <QInAppTransaction>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

class MyQInAppProduct: public QInAppProduct
{
public:
	void purchase() { callbackQInAppProduct_Purchase(this); };
	bool event(QEvent * e) { return callbackQInAppProduct_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppProduct_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQInAppProduct_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppProduct_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppProduct_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppProduct_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQInAppProduct_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppProduct_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPurchasing_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQInAppProduct_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQInAppProduct_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppProduct_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQInAppProduct*)

int QInAppProduct_QInAppProduct_QRegisterMetaType(){qRegisterMetaType<QInAppProduct*>(); return qRegisterMetaType<MyQInAppProduct*>();}

void QInAppProduct_Purchase(void* ptr)
{
	static_cast<QInAppProduct*>(ptr)->purchase();
}

long long QInAppProduct_ProductType(void* ptr)
{
	return static_cast<QInAppProduct*>(ptr)->productType();
}

struct QtPurchasing_PackedString QInAppProduct_Identifier(void* ptr)
{
	return ({ QByteArray t1115c2 = static_cast<QInAppProduct*>(ptr)->identifier().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t1115c2.prepend("WHITESPACE").constData()+10), t1115c2.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppProduct_Description(void* ptr)
{
	return ({ QByteArray t8de023 = static_cast<QInAppProduct*>(ptr)->description().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t8de023.prepend("WHITESPACE").constData()+10), t8de023.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppProduct_Price(void* ptr)
{
	return ({ QByteArray tb111f7 = static_cast<QInAppProduct*>(ptr)->price().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(tb111f7.prepend("WHITESPACE").constData()+10), tb111f7.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppProduct_Title(void* ptr)
{
	return ({ QByteArray t2b6ee0 = static_cast<QInAppProduct*>(ptr)->title().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t2b6ee0.prepend("WHITESPACE").constData()+10), t2b6ee0.size()-10 }; });
}

void* QInAppProduct___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QInAppProduct___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QInAppProduct___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QInAppProduct___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppProduct___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppProduct___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppProduct___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppProduct___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppProduct___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppProduct___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppProduct___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppProduct___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppProduct___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QInAppProduct___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppProduct___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QInAppProduct_EventDefault(void* ptr, void* e)
{
		return static_cast<QInAppProduct*>(ptr)->QInAppProduct::event(static_cast<QEvent*>(e));
}

char QInAppProduct_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QInAppProduct*>(ptr)->QInAppProduct::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QInAppProduct_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppProduct_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppProduct_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::customEvent(static_cast<QEvent*>(event));
}

void QInAppProduct_DeleteLaterDefault(void* ptr)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::deleteLater();
}

void QInAppProduct_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppProduct_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QInAppProduct*>(ptr)->QInAppProduct::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QInAppProduct_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QInAppProduct*>(ptr)->QInAppProduct::metaObject());
}

class MyQInAppStore: public QInAppStore
{
public:
	MyQInAppStore(QObject *parent = Q_NULLPTR) : QInAppStore(parent) {QInAppStore_QInAppStore_QRegisterMetaType();};
	void Signal_ProductRegistered(QInAppProduct * product) { callbackQInAppStore_ProductRegistered(this, product); };
	void Signal_ProductUnknown(QInAppProduct::ProductType productType, const QString & identifier) { QByteArray tfae9fd = identifier.toUtf8(); QtPurchasing_PackedString identifierPacked = { const_cast<char*>(tfae9fd.prepend("WHITESPACE").constData()+10), tfae9fd.size()-10 };callbackQInAppStore_ProductUnknown(this, productType, identifierPacked); };
	void Signal_TransactionReady(QInAppTransaction * transaction) { callbackQInAppStore_TransactionReady(this, transaction); };
	bool event(QEvent * e) { return callbackQInAppStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppStore_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQInAppStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppStore_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQInAppStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPurchasing_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQInAppStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQInAppStore_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppStore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQInAppStore*)

int QInAppStore_QInAppStore_QRegisterMetaType(){qRegisterMetaType<QInAppStore*>(); return qRegisterMetaType<MyQInAppStore*>();}

void* QInAppStore_NewQInAppStore(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQInAppStore(static_cast<QWindow*>(parent));
	} else {
		return new MyQInAppStore(static_cast<QObject*>(parent));
	}
}

void QInAppStore_ConnectProductRegistered(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct *)>(&QInAppStore::productRegistered), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct *)>(&MyQInAppStore::Signal_ProductRegistered));
}

void QInAppStore_DisconnectProductRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct *)>(&QInAppStore::productRegistered), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct *)>(&MyQInAppStore::Signal_ProductRegistered));
}

void QInAppStore_ProductRegistered(void* ptr, void* product)
{
	static_cast<QInAppStore*>(ptr)->productRegistered(static_cast<QInAppProduct*>(product));
}

void QInAppStore_ConnectProductUnknown(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&QInAppStore::productUnknown), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&MyQInAppStore::Signal_ProductUnknown));
}

void QInAppStore_DisconnectProductUnknown(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&QInAppStore::productUnknown), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&MyQInAppStore::Signal_ProductUnknown));
}

void QInAppStore_ProductUnknown(void* ptr, long long productType, struct QtPurchasing_PackedString identifier)
{
	static_cast<QInAppStore*>(ptr)->productUnknown(static_cast<QInAppProduct::ProductType>(productType), QString::fromUtf8(identifier.data, identifier.len));
}

void QInAppStore_RegisterProduct(void* ptr, long long productType, struct QtPurchasing_PackedString identifier)
{
	static_cast<QInAppStore*>(ptr)->registerProduct(static_cast<QInAppProduct::ProductType>(productType), QString::fromUtf8(identifier.data, identifier.len));
}

void QInAppStore_RestorePurchases(void* ptr)
{
	static_cast<QInAppStore*>(ptr)->restorePurchases();
}

void QInAppStore_SetPlatformProperty(void* ptr, struct QtPurchasing_PackedString propertyName, struct QtPurchasing_PackedString value)
{
	static_cast<QInAppStore*>(ptr)->setPlatformProperty(QString::fromUtf8(propertyName.data, propertyName.len), QString::fromUtf8(value.data, value.len));
}

void QInAppStore_ConnectTransactionReady(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppTransaction *)>(&QInAppStore::transactionReady), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppTransaction *)>(&MyQInAppStore::Signal_TransactionReady));
}

void QInAppStore_DisconnectTransactionReady(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppTransaction *)>(&QInAppStore::transactionReady), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppTransaction *)>(&MyQInAppStore::Signal_TransactionReady));
}

void QInAppStore_TransactionReady(void* ptr, void* transaction)
{
	static_cast<QInAppStore*>(ptr)->transactionReady(static_cast<QInAppTransaction*>(transaction));
}

void QInAppStore_DestroyQInAppStore(void* ptr)
{
	static_cast<QInAppStore*>(ptr)->~QInAppStore();
}

void* QInAppStore_RegisteredProduct(void* ptr, struct QtPurchasing_PackedString identifier)
{
	return static_cast<QInAppStore*>(ptr)->registeredProduct(QString::fromUtf8(identifier.data, identifier.len));
}

void* QInAppStore___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QInAppStore___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QInAppStore___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QInAppStore___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppStore___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppStore___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppStore___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppStore___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppStore___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppStore___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppStore___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppStore___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppStore___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QInAppStore___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppStore___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QInAppStore_EventDefault(void* ptr, void* e)
{
		return static_cast<QInAppStore*>(ptr)->QInAppStore::event(static_cast<QEvent*>(e));
}

char QInAppStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QInAppStore*>(ptr)->QInAppStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QInAppStore_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppStore_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppStore_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::customEvent(static_cast<QEvent*>(event));
}

void QInAppStore_DeleteLaterDefault(void* ptr)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::deleteLater();
}

void QInAppStore_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppStore_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QInAppStore*>(ptr)->QInAppStore::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QInAppStore_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QInAppStore*>(ptr)->QInAppStore::metaObject());
}

class MyQInAppTransaction: public QInAppTransaction
{
public:
	QString errorString() const { return ({ QtPurchasing_PackedString tempVal = callbackQInAppTransaction_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString orderId() const { return ({ QtPurchasing_PackedString tempVal = callbackQInAppTransaction_OrderId(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void finalize() { callbackQInAppTransaction_Finalize(this); };
	FailureReason failureReason() const { return static_cast<QInAppTransaction::FailureReason>(callbackQInAppTransaction_FailureReason(const_cast<void*>(static_cast<const void*>(this)))); };
	QDateTime timestamp() const { return *static_cast<QDateTime*>(callbackQInAppTransaction_Timestamp(const_cast<void*>(static_cast<const void*>(this)))); };
	QString platformProperty(const QString & propertyName) const { QByteArray tdeaeb2 = propertyName.toUtf8(); QtPurchasing_PackedString propertyNamePacked = { const_cast<char*>(tdeaeb2.prepend("WHITESPACE").constData()+10), tdeaeb2.size()-10 };return ({ QtPurchasing_PackedString tempVal = callbackQInAppTransaction_PlatformProperty(const_cast<void*>(static_cast<const void*>(this)), propertyNamePacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool event(QEvent * e) { return callbackQInAppTransaction_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppTransaction_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQInAppTransaction_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppTransaction_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppTransaction_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppTransaction_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQInAppTransaction_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppTransaction_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPurchasing_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQInAppTransaction_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQInAppTransaction_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppTransaction_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQInAppTransaction*)

int QInAppTransaction_QInAppTransaction_QRegisterMetaType(){qRegisterMetaType<QInAppTransaction*>(); return qRegisterMetaType<MyQInAppTransaction*>();}

struct QtPurchasing_PackedString QInAppTransaction_ErrorString(void* ptr)
{
	return ({ QByteArray t910a82 = static_cast<QInAppTransaction*>(ptr)->errorString().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t910a82.prepend("WHITESPACE").constData()+10), t910a82.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppTransaction_ErrorStringDefault(void* ptr)
{
		return ({ QByteArray tb31c41 = static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::errorString().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(tb31c41.prepend("WHITESPACE").constData()+10), tb31c41.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppTransaction_OrderId(void* ptr)
{
	return ({ QByteArray t614537 = static_cast<QInAppTransaction*>(ptr)->orderId().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t614537.prepend("WHITESPACE").constData()+10), t614537.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppTransaction_OrderIdDefault(void* ptr)
{
		return ({ QByteArray t9740f5 = static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::orderId().toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t9740f5.prepend("WHITESPACE").constData()+10), t9740f5.size()-10 }; });
}

void QInAppTransaction_Finalize(void* ptr)
{
	static_cast<QInAppTransaction*>(ptr)->finalize();
}

long long QInAppTransaction_FailureReason(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->failureReason();
}

long long QInAppTransaction_FailureReasonDefault(void* ptr)
{
		return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::failureReason();
}

void* QInAppTransaction_Timestamp(void* ptr)
{
	return new QDateTime(static_cast<QInAppTransaction*>(ptr)->timestamp());
}

void* QInAppTransaction_TimestampDefault(void* ptr)
{
		return new QDateTime(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::timestamp());
}

void* QInAppTransaction_Product(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->product();
}

struct QtPurchasing_PackedString QInAppTransaction_PlatformProperty(void* ptr, struct QtPurchasing_PackedString propertyName)
{
	return ({ QByteArray t2b73cb = static_cast<QInAppTransaction*>(ptr)->platformProperty(QString::fromUtf8(propertyName.data, propertyName.len)).toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t2b73cb.prepend("WHITESPACE").constData()+10), t2b73cb.size()-10 }; });
}

struct QtPurchasing_PackedString QInAppTransaction_PlatformPropertyDefault(void* ptr, struct QtPurchasing_PackedString propertyName)
{
		return ({ QByteArray t2cb180 = static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::platformProperty(QString::fromUtf8(propertyName.data, propertyName.len)).toUtf8(); QtPurchasing_PackedString { const_cast<char*>(t2cb180.prepend("WHITESPACE").constData()+10), t2cb180.size()-10 }; });
}

long long QInAppTransaction_Status(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->status();
}

void* QInAppTransaction___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QInAppTransaction___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QInAppTransaction___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QInAppTransaction___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppTransaction___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppTransaction___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppTransaction___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppTransaction___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppTransaction___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppTransaction___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QInAppTransaction___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppTransaction___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QInAppTransaction___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QInAppTransaction___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QInAppTransaction___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QInAppTransaction_EventDefault(void* ptr, void* e)
{
		return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::event(static_cast<QEvent*>(e));
}

char QInAppTransaction_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QInAppTransaction_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppTransaction_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppTransaction_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::customEvent(static_cast<QEvent*>(event));
}

void QInAppTransaction_DeleteLaterDefault(void* ptr)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::deleteLater();
}

void QInAppTransaction_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppTransaction_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QInAppTransaction_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::metaObject());
}

