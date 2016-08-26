// +build !minimal

#define protected public
#define private public

#include "macextras.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QIcon>
#include <QMacPasteboardMime>
#include <QMacToolBar>
#include <QMacToolBarItem>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWindow>

class MyQMacPasteboardMime: public QMacPasteboardMime
{
public:
	bool canConvert(const QString & mime, QString flav) { return callbackQMacPasteboardMime_CanConvert(this, const_cast<char*>(mime.toUtf8().constData()), const_cast<char*>(flav.toUtf8().constData())) != 0; };
	QString convertorName() { return QString(callbackQMacPasteboardMime_ConvertorName(this)); };
	int count(QMimeData * mimeData) { return callbackQMacPasteboardMime_Count(this, mimeData); };
	QString flavorFor(const QString & mime) { return QString(callbackQMacPasteboardMime_FlavorFor(this, const_cast<char*>(mime.toUtf8().constData()))); };
	QString mimeFor(QString flav) { return QString(callbackQMacPasteboardMime_MimeFor(this, const_cast<char*>(flav.toUtf8().constData()))); };
	 ~MyQMacPasteboardMime() { callbackQMacPasteboardMime_DestroyQMacPasteboardMime(this); };
};

char QMacPasteboardMime_CanConvert(void* ptr, char* mime, char* flav)
{
	return static_cast<QMacPasteboardMime*>(ptr)->canConvert(QString(mime), QString(flav));
}

char* QMacPasteboardMime_ConvertorName(void* ptr)
{
	return const_cast<char*>(static_cast<QMacPasteboardMime*>(ptr)->convertorName().toUtf8().constData());
}

int QMacPasteboardMime_Count(void* ptr, void* mimeData)
{
	return static_cast<QMacPasteboardMime*>(ptr)->count(static_cast<QMimeData*>(mimeData));
}

int QMacPasteboardMime_CountDefault(void* ptr, void* mimeData)
{
#ifdef Q_OS_OSX
	return static_cast<QMacPasteboardMime*>(ptr)->QMacPasteboardMime::count(static_cast<QMimeData*>(mimeData));
#else
	return 0;
#endif
}

char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime)
{
	return const_cast<char*>(static_cast<QMacPasteboardMime*>(ptr)->flavorFor(QString(mime)).toUtf8().constData());
}

char* QMacPasteboardMime_MimeFor(void* ptr, char* flav)
{
	return const_cast<char*>(static_cast<QMacPasteboardMime*>(ptr)->mimeFor(QString(flav)).toUtf8().constData());
}

void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr)
{
	static_cast<QMacPasteboardMime*>(ptr)->~QMacPasteboardMime();
}

void QMacPasteboardMime_DestroyQMacPasteboardMimeDefault(void* ptr)
{
#ifdef Q_OS_OSX

#endif
}

void* QMacToolBar_NewQMacToolBar(void* parent)
{
	return new QMacToolBar(static_cast<QObject*>(parent));
}

void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent)
{
	return new QMacToolBar(QString(identifier), static_cast<QObject*>(parent));
}

void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text)
{
	return static_cast<QMacToolBar*>(ptr)->addAllowedItem(*static_cast<QIcon*>(icon), QString(text));
}

void* QMacToolBar_AddItem(void* ptr, void* icon, char* text)
{
	return static_cast<QMacToolBar*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text));
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

void QMacToolBar_DestroyQMacToolBar(void* ptr)
{
	static_cast<QMacToolBar*>(ptr)->~QMacToolBar();
}

void QMacToolBar_TimerEvent(void* ptr, void* event)
{
	static_cast<QMacToolBar*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBar_TimerEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::timerEvent(static_cast<QTimerEvent*>(event));
#endif
}

void QMacToolBar_ChildEvent(void* ptr, void* event)
{
	static_cast<QMacToolBar*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBar_ChildEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::childEvent(static_cast<QChildEvent*>(event));
#endif
}

void QMacToolBar_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMacToolBar*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBar_ConnectNotifyDefault(void* ptr, void* sign)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::connectNotify(*static_cast<QMetaMethod*>(sign));
#endif
}

void QMacToolBar_CustomEvent(void* ptr, void* event)
{
	static_cast<QMacToolBar*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMacToolBar_CustomEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::customEvent(static_cast<QEvent*>(event));
#endif
}

void QMacToolBar_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMacToolBar*>(ptr), "deleteLater");
}

void QMacToolBar_DeleteLaterDefault(void* ptr)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::deleteLater();
#endif
}

void QMacToolBar_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMacToolBar*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBar_DisconnectNotifyDefault(void* ptr, void* sign)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
#endif
}

char QMacToolBar_Event(void* ptr, void* e)
{
	return static_cast<QMacToolBar*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMacToolBar_EventDefault(void* ptr, void* e)
{
#ifdef Q_OS_OSX
	return static_cast<QMacToolBar*>(ptr)->QMacToolBar::event(static_cast<QEvent*>(e));
#else
	return false;
#endif
}

char QMacToolBar_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMacToolBar*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMacToolBar_EventFilterDefault(void* ptr, void* watched, void* event)
{
#ifdef Q_OS_OSX
	return static_cast<QMacToolBar*>(ptr)->QMacToolBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
#else
	return false;
#endif
}

void* QMacToolBar_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMacToolBar*>(ptr)->metaObject());
}

void* QMacToolBar_MetaObjectDefault(void* ptr)
{
#ifdef Q_OS_OSX
	return const_cast<QMetaObject*>(static_cast<QMacToolBar*>(ptr)->QMacToolBar::metaObject());
#else
	return NULL;
#endif
}

class MyQMacToolBarItem: public QMacToolBarItem
{
public:
	MyQMacToolBarItem(QObject *parent) : QMacToolBarItem(parent) {};
	void Signal_Activated() { callbackQMacToolBarItem_Activated(this); };
	 ~MyQMacToolBarItem() { callbackQMacToolBarItem_DestroyQMacToolBarItem(this); };
	void timerEvent(QTimerEvent * event) { callbackQMacToolBarItem_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMacToolBarItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMacToolBarItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMacToolBarItem_CustomEvent(this, event); };
	void deleteLater() { callbackQMacToolBarItem_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMacToolBarItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMacToolBarItem_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMacToolBarItem_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMacToolBarItem_MetaObject(const_cast<MyQMacToolBarItem*>(this))); };
};

void* QMacToolBarItem_NewQMacToolBarItem(void* parent)
{
	return new MyQMacToolBarItem(static_cast<QObject*>(parent));
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
#ifdef Q_OS_OSX

#endif
}

void* QMacToolBarItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QMacToolBarItem*>(ptr)->icon());
}

char QMacToolBarItem_Selectable(void* ptr)
{
	return static_cast<QMacToolBarItem*>(ptr)->selectable();
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

void QMacToolBarItem_SetText(void* ptr, char* text)
{
	static_cast<QMacToolBarItem*>(ptr)->setText(QString(text));
}

long long QMacToolBarItem_StandardItem(void* ptr)
{
	return static_cast<QMacToolBarItem*>(ptr)->standardItem();
}

char* QMacToolBarItem_Text(void* ptr)
{
	return const_cast<char*>(static_cast<QMacToolBarItem*>(ptr)->text().toUtf8().constData());
}

void QMacToolBarItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QMacToolBarItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBarItem_TimerEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::timerEvent(static_cast<QTimerEvent*>(event));
#endif
}

void QMacToolBarItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QMacToolBarItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBarItem_ChildEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::childEvent(static_cast<QChildEvent*>(event));
#endif
}

void QMacToolBarItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMacToolBarItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBarItem_ConnectNotifyDefault(void* ptr, void* sign)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::connectNotify(*static_cast<QMetaMethod*>(sign));
#endif
}

void QMacToolBarItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QMacToolBarItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMacToolBarItem_CustomEventDefault(void* ptr, void* event)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::customEvent(static_cast<QEvent*>(event));
#endif
}

void QMacToolBarItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMacToolBarItem*>(ptr), "deleteLater");
}

void QMacToolBarItem_DeleteLaterDefault(void* ptr)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::deleteLater();
#endif
}

void QMacToolBarItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMacToolBarItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMacToolBarItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
#ifdef Q_OS_OSX
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
#endif
}

char QMacToolBarItem_Event(void* ptr, void* e)
{
	return static_cast<QMacToolBarItem*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMacToolBarItem_EventDefault(void* ptr, void* e)
{
#ifdef Q_OS_OSX
	return static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::event(static_cast<QEvent*>(e));
#else
	return false;
#endif
}

char QMacToolBarItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMacToolBarItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMacToolBarItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
#ifdef Q_OS_OSX
	return static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
#else
	return false;
#endif
}

void* QMacToolBarItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMacToolBarItem*>(ptr)->metaObject());
}

void* QMacToolBarItem_MetaObjectDefault(void* ptr)
{
#ifdef Q_OS_OSX
	return const_cast<QMetaObject*>(static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::metaObject());
#else
	return NULL;
#endif
}

