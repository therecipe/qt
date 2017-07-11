// +build !minimal

#define protected public
#define private public

#include "macextras.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIcon>
#include <QLayout>
#include <QList>
#include <QMacPasteboardMime>
#include <QMacToolBar>
#include <QMacToolBarItem>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
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
#include <QVariant>
#include <QWidget>
#include <QWindow>

class MyQMacPasteboardMime: public QMacPasteboardMime
{
public:
	MyQMacPasteboardMime(char t) : QMacPasteboardMime(t) {};
	QList<QByteArray> convertFromMime(const QString & mime, QVariant data, QString flav) { QByteArray tc6d51a = mime.toUtf8(); QtMacExtras_PackedString mimePacked = { const_cast<char*>(tc6d51a.prepend("WHITESPACE").constData()+10), tc6d51a.size()-10 };QByteArray t81c607 = flav.toUtf8(); QtMacExtras_PackedString flavPacked = { const_cast<char*>(t81c607.prepend("WHITESPACE").constData()+10), t81c607.size()-10 };return *static_cast<QList<QByteArray>*>(callbackQMacPasteboardMime_ConvertFromMime(this, mimePacked, new QVariant(data), flavPacked)); };
	QString convertorName() { return ({ QtMacExtras_PackedString tempVal = callbackQMacPasteboardMime_ConvertorName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString flavorFor(const QString & mime) { QByteArray tc6d51a = mime.toUtf8(); QtMacExtras_PackedString mimePacked = { const_cast<char*>(tc6d51a.prepend("WHITESPACE").constData()+10), tc6d51a.size()-10 };return ({ QtMacExtras_PackedString tempVal = callbackQMacPasteboardMime_FlavorFor(this, mimePacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString mimeFor(QString flav) { QByteArray t81c607 = flav.toUtf8(); QtMacExtras_PackedString flavPacked = { const_cast<char*>(t81c607.prepend("WHITESPACE").constData()+10), t81c607.size()-10 };return ({ QtMacExtras_PackedString tempVal = callbackQMacPasteboardMime_MimeFor(this, flavPacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QVariant convertToMime(const QString & mime, QList<QByteArray> data, QString flav) { QByteArray tc6d51a = mime.toUtf8(); QtMacExtras_PackedString mimePacked = { const_cast<char*>(tc6d51a.prepend("WHITESPACE").constData()+10), tc6d51a.size()-10 };QByteArray t81c607 = flav.toUtf8(); QtMacExtras_PackedString flavPacked = { const_cast<char*>(t81c607.prepend("WHITESPACE").constData()+10), t81c607.size()-10 };return *static_cast<QVariant*>(callbackQMacPasteboardMime_ConvertToMime(this, mimePacked, ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(data); QtMacExtras_PackedList { tmpValue, tmpValue->size() }; }), flavPacked)); };
	bool canConvert(const QString & mime, QString flav) { QByteArray tc6d51a = mime.toUtf8(); QtMacExtras_PackedString mimePacked = { const_cast<char*>(tc6d51a.prepend("WHITESPACE").constData()+10), tc6d51a.size()-10 };QByteArray t81c607 = flav.toUtf8(); QtMacExtras_PackedString flavPacked = { const_cast<char*>(t81c607.prepend("WHITESPACE").constData()+10), t81c607.size()-10 };return callbackQMacPasteboardMime_CanConvert(this, mimePacked, flavPacked) != 0; };
	int count(QMimeData * mimeData) { return callbackQMacPasteboardMime_Count(this, mimeData); };
	 ~MyQMacPasteboardMime() { callbackQMacPasteboardMime_DestroyQMacPasteboardMime(this); };
};

struct QtMacExtras_PackedList QMacPasteboardMime_ConvertFromMime(void* ptr, struct QtMacExtras_PackedString mime, void* data, struct QtMacExtras_PackedString flav)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QMacPasteboardMime*>(ptr)->convertFromMime(QString::fromUtf8(mime.data, mime.len), *static_cast<QVariant*>(data), QString::fromUtf8(flav.data, flav.len))); QtMacExtras_PackedList { tmpValue, tmpValue->size() }; });
}

void* QMacPasteboardMime_NewQMacPasteboardMime(char* t)
{
	return new MyQMacPasteboardMime(*t);
}

struct QtMacExtras_PackedString QMacPasteboardMime_ConvertorName(void* ptr)
{
	return ({ QByteArray t4ba9d6 = static_cast<QMacPasteboardMime*>(ptr)->convertorName().toUtf8(); QtMacExtras_PackedString { const_cast<char*>(t4ba9d6.prepend("WHITESPACE").constData()+10), t4ba9d6.size()-10 }; });
}

struct QtMacExtras_PackedString QMacPasteboardMime_FlavorFor(void* ptr, struct QtMacExtras_PackedString mime)
{
	return ({ QByteArray t7733db = static_cast<QMacPasteboardMime*>(ptr)->flavorFor(QString::fromUtf8(mime.data, mime.len)).toUtf8(); QtMacExtras_PackedString { const_cast<char*>(t7733db.prepend("WHITESPACE").constData()+10), t7733db.size()-10 }; });
}

struct QtMacExtras_PackedString QMacPasteboardMime_MimeFor(void* ptr, struct QtMacExtras_PackedString flav)
{
	return ({ QByteArray t4660db = static_cast<QMacPasteboardMime*>(ptr)->mimeFor(QString::fromUtf8(flav.data, flav.len)).toUtf8(); QtMacExtras_PackedString { const_cast<char*>(t4660db.prepend("WHITESPACE").constData()+10), t4660db.size()-10 }; });
}

void* QMacPasteboardMime_ConvertToMime(void* ptr, struct QtMacExtras_PackedString mime, void* data, struct QtMacExtras_PackedString flav)
{
	return new QVariant(static_cast<QMacPasteboardMime*>(ptr)->convertToMime(QString::fromUtf8(mime.data, mime.len), *static_cast<QList<QByteArray>*>(data), QString::fromUtf8(flav.data, flav.len)));
}

char QMacPasteboardMime_CanConvert(void* ptr, struct QtMacExtras_PackedString mime, struct QtMacExtras_PackedString flav)
{
	return static_cast<QMacPasteboardMime*>(ptr)->canConvert(QString::fromUtf8(mime.data, mime.len), QString::fromUtf8(flav.data, flav.len));
}

int QMacPasteboardMime_Count(void* ptr, void* mimeData)
{
	return static_cast<QMacPasteboardMime*>(ptr)->count(static_cast<QMimeData*>(mimeData));
}

int QMacPasteboardMime_CountDefault(void* ptr, void* mimeData)
{
		return static_cast<QMacPasteboardMime*>(ptr)->QMacPasteboardMime::count(static_cast<QMimeData*>(mimeData));
}

void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr)
{
	static_cast<QMacPasteboardMime*>(ptr)->~QMacPasteboardMime();
}

void QMacPasteboardMime_DestroyQMacPasteboardMimeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMacPasteboardMime___convertFromMime_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QMacPasteboardMime___convertFromMime_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMacPasteboardMime___convertFromMime_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QMacPasteboardMime___convertToMime_data_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QMacPasteboardMime___convertToMime_data_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMacPasteboardMime___convertToMime_data_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

class MyQMacToolBar: public QMacToolBar
{
public:
	MyQMacToolBar(QObject *parent = Q_NULLPTR) : QMacToolBar(parent) {QMacToolBar_QMacToolBar_QRegisterMetaType();};
	MyQMacToolBar(const QString &identifier, QObject *parent = Q_NULLPTR) : QMacToolBar(identifier, parent) {QMacToolBar_QMacToolBar_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQMacToolBar_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMacToolBar_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQMacToolBar_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMacToolBar_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMacToolBar_CustomEvent(this, event); };
	void deleteLater() { callbackQMacToolBar_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQMacToolBar_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMacToolBar_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtMacExtras_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQMacToolBar_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQMacToolBar_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMacToolBar_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQMacToolBar*)

int QMacToolBar_QMacToolBar_QRegisterMetaType(){qRegisterMetaType<QMacToolBar*>(); return qRegisterMetaType<MyQMacToolBar*>();}

struct QtMacExtras_PackedList QMacToolBar_AllowedItems(void* ptr)
{
	return ({ QList<QMacToolBarItem *>* tmpValue = new QList<QMacToolBarItem *>(static_cast<QMacToolBar*>(ptr)->allowedItems()); QtMacExtras_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtMacExtras_PackedList QMacToolBar_Items(void* ptr)
{
	return ({ QList<QMacToolBarItem *>* tmpValue = new QList<QMacToolBarItem *>(static_cast<QMacToolBar*>(ptr)->items()); QtMacExtras_PackedList { tmpValue, tmpValue->size() }; });
}

void* QMacToolBar_NewQMacToolBar(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(static_cast<QWindow*>(parent));
	} else {
		return new MyQMacToolBar(static_cast<QObject*>(parent));
	}
}

void* QMacToolBar_NewQMacToolBar2(struct QtMacExtras_PackedString identifier, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQMacToolBar(QString::fromUtf8(identifier.data, identifier.len), static_cast<QObject*>(parent));
	}
}

void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, struct QtMacExtras_PackedString text)
{
	return static_cast<QMacToolBar*>(ptr)->addAllowedItem(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len));
}

void* QMacToolBar_AddItem(void* ptr, void* icon, struct QtMacExtras_PackedString text)
{
	return static_cast<QMacToolBar*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len));
}

void QMacToolBar_AddSeparator(void* ptr)
{
	static_cast<QMacToolBar*>(ptr)->addSeparator();
}

void QMacToolBar_AttachToWindow(void* ptr, void* window)
{
	static_cast<QMacToolBar*>(ptr)->attachToWindow(static_cast<QWindow*>(window));
}

void QMacToolBar_DetachFromWindow(void* ptr)
{
	static_cast<QMacToolBar*>(ptr)->detachFromWindow();
}

void QMacToolBar_SetAllowedItems(void* ptr, void* allowedItems)
{
	static_cast<QMacToolBar*>(ptr)->setAllowedItems(*static_cast<QList<QMacToolBarItem *>*>(allowedItems));
}

void QMacToolBar_SetItems(void* ptr, void* items)
{
	static_cast<QMacToolBar*>(ptr)->setItems(*static_cast<QList<QMacToolBarItem *>*>(items));
}

void QMacToolBar_DestroyQMacToolBar(void* ptr)
{
	static_cast<QMacToolBar*>(ptr)->~QMacToolBar();
}

void* QMacToolBar___allowedItems_atList(void* ptr, int i)
{
	return const_cast<QMacToolBarItem*>(static_cast<QList<QMacToolBarItem *>*>(ptr)->at(i));
}

void QMacToolBar___allowedItems_setList(void* ptr, void* i)
{
	static_cast<QList<QMacToolBarItem *>*>(ptr)->append(static_cast<QMacToolBarItem*>(i));
}

void* QMacToolBar___allowedItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QMacToolBarItem *>;
}

void* QMacToolBar___items_atList(void* ptr, int i)
{
	return const_cast<QMacToolBarItem*>(static_cast<QList<QMacToolBarItem *>*>(ptr)->at(i));
}

void QMacToolBar___items_setList(void* ptr, void* i)
{
	static_cast<QList<QMacToolBarItem *>*>(ptr)->append(static_cast<QMacToolBarItem*>(i));
}

void* QMacToolBar___items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QMacToolBarItem *>;
}

void* QMacToolBar___setAllowedItems_allowedItems_atList(void* ptr, int i)
{
	return const_cast<QMacToolBarItem*>(static_cast<QList<QMacToolBarItem *>*>(ptr)->at(i));
}

void QMacToolBar___setAllowedItems_allowedItems_setList(void* ptr, void* i)
{
	static_cast<QList<QMacToolBarItem *>*>(ptr)->append(static_cast<QMacToolBarItem*>(i));
}

void* QMacToolBar___setAllowedItems_allowedItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QMacToolBarItem *>;
}

void* QMacToolBar___setItems_items_atList(void* ptr, int i)
{
	return const_cast<QMacToolBarItem*>(static_cast<QList<QMacToolBarItem *>*>(ptr)->at(i));
}

void QMacToolBar___setItems_items_setList(void* ptr, void* i)
{
	static_cast<QList<QMacToolBarItem *>*>(ptr)->append(static_cast<QMacToolBarItem*>(i));
}

void* QMacToolBar___setItems_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QMacToolBarItem *>;
}

void* QMacToolBar___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QMacToolBar___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMacToolBar___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QMacToolBar___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBar___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBar___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBar___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBar___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBar___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBar___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBar___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBar___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBar___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QMacToolBar___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBar___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QMacToolBar_EventDefault(void* ptr, void* e)
{
		return static_cast<QMacToolBar*>(ptr)->QMacToolBar::event(static_cast<QEvent*>(e));
}

char QMacToolBar_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QMacToolBar*>(ptr)->QMacToolBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QMacToolBar_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBar_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBar_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::customEvent(static_cast<QEvent*>(event));
}

void QMacToolBar_DeleteLaterDefault(void* ptr)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::deleteLater();
}

void QMacToolBar_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBar_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBar*>(ptr)->QMacToolBar::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QMacToolBar_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QMacToolBar*>(ptr)->QMacToolBar::metaObject());
}

class MyQMacToolBarItem: public QMacToolBarItem
{
public:
	MyQMacToolBarItem(QObject *parent = Q_NULLPTR) : QMacToolBarItem(parent) {QMacToolBarItem_QMacToolBarItem_QRegisterMetaType();};
	void Signal_Activated() { callbackQMacToolBarItem_Activated(this); };
	 ~MyQMacToolBarItem() { callbackQMacToolBarItem_DestroyQMacToolBarItem(this); };
	bool event(QEvent * e) { return callbackQMacToolBarItem_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMacToolBarItem_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQMacToolBarItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMacToolBarItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMacToolBarItem_CustomEvent(this, event); };
	void deleteLater() { callbackQMacToolBarItem_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQMacToolBarItem_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMacToolBarItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtMacExtras_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQMacToolBarItem_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQMacToolBarItem_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMacToolBarItem_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQMacToolBarItem*)

int QMacToolBarItem_QMacToolBarItem_QRegisterMetaType(){qRegisterMetaType<QMacToolBarItem*>(); return qRegisterMetaType<MyQMacToolBarItem*>();}

void* QMacToolBarItem_NewQMacToolBarItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMacToolBarItem(static_cast<QWindow*>(parent));
	} else {
		return new MyQMacToolBarItem(static_cast<QObject*>(parent));
	}
}

void QMacToolBarItem_ConnectActivated(void* ptr)
{
	QObject::connect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));
}

void QMacToolBarItem_DisconnectActivated(void* ptr)
{
	QObject::disconnect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));
}

void QMacToolBarItem_Activated(void* ptr)
{
	static_cast<QMacToolBarItem*>(ptr)->activated();
}

void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr)
{
	static_cast<QMacToolBarItem*>(ptr)->~QMacToolBarItem();
}

void QMacToolBarItem_DestroyQMacToolBarItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QMacToolBarItem_SetIcon(void* ptr, void* icon)
{
	static_cast<QMacToolBarItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMacToolBarItem_SetSelectable(void* ptr, char selectable)
{
	static_cast<QMacToolBarItem*>(ptr)->setSelectable(selectable != 0);
}

void QMacToolBarItem_SetStandardItem(void* ptr, long long standardItem)
{
	static_cast<QMacToolBarItem*>(ptr)->setStandardItem(static_cast<QMacToolBarItem::StandardItem>(standardItem));
}

void QMacToolBarItem_SetText(void* ptr, struct QtMacExtras_PackedString text)
{
	static_cast<QMacToolBarItem*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void* QMacToolBarItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QMacToolBarItem*>(ptr)->icon());
}

struct QtMacExtras_PackedString QMacToolBarItem_Text(void* ptr)
{
	return ({ QByteArray t8c9d50 = static_cast<QMacToolBarItem*>(ptr)->text().toUtf8(); QtMacExtras_PackedString { const_cast<char*>(t8c9d50.prepend("WHITESPACE").constData()+10), t8c9d50.size()-10 }; });
}

long long QMacToolBarItem_StandardItem(void* ptr)
{
	return static_cast<QMacToolBarItem*>(ptr)->standardItem();
}

char QMacToolBarItem_Selectable(void* ptr)
{
	return static_cast<QMacToolBarItem*>(ptr)->selectable();
}

void* QMacToolBarItem___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QMacToolBarItem___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMacToolBarItem___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QMacToolBarItem___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBarItem___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBarItem___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBarItem___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBarItem___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBarItem___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBarItem___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QMacToolBarItem___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBarItem___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QMacToolBarItem___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QMacToolBarItem___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMacToolBarItem___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QMacToolBarItem_EventDefault(void* ptr, void* e)
{
		return static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::event(static_cast<QEvent*>(e));
}

char QMacToolBarItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QMacToolBarItem_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBarItem_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBarItem_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::customEvent(static_cast<QEvent*>(event));
}

void QMacToolBarItem_DeleteLaterDefault(void* ptr)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::deleteLater();
}

void QMacToolBarItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBarItem_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QMacToolBarItem_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::metaObject());
}

