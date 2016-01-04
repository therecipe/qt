#define protected public

#include "macextras.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QIcon>
#include <QMacPasteboardMime>
#include <QMacToolBar>
#include <QMacToolBarItem>
#include <QMimeData>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWindow>

class MyQMacPasteboardMime: public QMacPasteboardMime {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

char* QMacPasteboardMime_ConvertorName(void* ptr){
	return static_cast<QMacPasteboardMime*>(ptr)->convertorName().toUtf8().data();
}

int QMacPasteboardMime_Count(void* ptr, void* mimeData){
	return static_cast<QMacPasteboardMime*>(ptr)->count(static_cast<QMimeData*>(mimeData));
}

char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime){
	return static_cast<QMacPasteboardMime*>(ptr)->flavorFor(QString(mime)).toUtf8().data();
}

char* QMacPasteboardMime_MimeFor(void* ptr, char* flav){
	return static_cast<QMacPasteboardMime*>(ptr)->mimeFor(QString(flav)).toUtf8().data();
}

void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr){
	static_cast<QMacPasteboardMime*>(ptr)->~QMacPasteboardMime();
}

char* QMacPasteboardMime_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQMacPasteboardMime*>(static_cast<QMacPasteboardMime*>(ptr))) {
		return static_cast<MyQMacPasteboardMime*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QMacPasteboardMime_BASE").toUtf8().data();
}

void QMacPasteboardMime_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQMacPasteboardMime*>(static_cast<QMacPasteboardMime*>(ptr))) {
		static_cast<MyQMacPasteboardMime*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QMacToolBar_NewQMacToolBar(void* parent){
	return new QMacToolBar(static_cast<QObject*>(parent));
}

void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent){
	return new QMacToolBar(QString(identifier), static_cast<QObject*>(parent));
}

void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text){
	return static_cast<QMacToolBar*>(ptr)->addAllowedItem(*static_cast<QIcon*>(icon), QString(text));
}

void* QMacToolBar_AddItem(void* ptr, void* icon, char* text){
	return static_cast<QMacToolBar*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text));
}

void QMacToolBar_AddSeparator(void* ptr){
	static_cast<QMacToolBar*>(ptr)->addSeparator();
}

void QMacToolBar_AttachToWindow(void* ptr, void* window){
	static_cast<QMacToolBar*>(ptr)->attachToWindow(static_cast<QWindow*>(window));
}

void QMacToolBar_DetachFromWindow(void* ptr){
	static_cast<QMacToolBar*>(ptr)->detachFromWindow();
}

void QMacToolBar_DestroyQMacToolBar(void* ptr){
	static_cast<QMacToolBar*>(ptr)->~QMacToolBar();
}

void QMacToolBar_TimerEvent(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBar_TimerEventDefault(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBar_ChildEvent(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBar_ChildEventDefault(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBar_CustomEvent(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMacToolBar_CustomEventDefault(void* ptr, void* event){
	static_cast<QMacToolBar*>(ptr)->QMacToolBar::customEvent(static_cast<QEvent*>(event));
}

class MyQMacToolBarItem: public QMacToolBarItem {
public:
	MyQMacToolBarItem(QObject *parent) : QMacToolBarItem(parent) {};
	void Signal_Activated() { callbackQMacToolBarItemActivated(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQMacToolBarItemTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQMacToolBarItemChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQMacToolBarItemCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QMacToolBarItem_NewQMacToolBarItem(void* parent){
	return new MyQMacToolBarItem(static_cast<QObject*>(parent));
}

void QMacToolBarItem_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_Activated(void* ptr){
	static_cast<QMacToolBarItem*>(ptr)->activated();
}

void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr){
	static_cast<QMacToolBarItem*>(ptr)->~QMacToolBarItem();
}

void* QMacToolBarItem_Icon(void* ptr){
	return new QIcon(static_cast<QMacToolBarItem*>(ptr)->icon());
}

int QMacToolBarItem_Selectable(void* ptr){
	return static_cast<QMacToolBarItem*>(ptr)->selectable();
}

void QMacToolBarItem_SetIcon(void* ptr, void* icon){
	static_cast<QMacToolBarItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMacToolBarItem_SetSelectable(void* ptr, int selectable){
	static_cast<QMacToolBarItem*>(ptr)->setSelectable(selectable != 0);
}

void QMacToolBarItem_SetStandardItem(void* ptr, int standardItem){
	static_cast<QMacToolBarItem*>(ptr)->setStandardItem(static_cast<QMacToolBarItem::StandardItem>(standardItem));
}

void QMacToolBarItem_SetText(void* ptr, char* text){
	static_cast<QMacToolBarItem*>(ptr)->setText(QString(text));
}

int QMacToolBarItem_StandardItem(void* ptr){
	return static_cast<QMacToolBarItem*>(ptr)->standardItem();
}

char* QMacToolBarItem_Text(void* ptr){
	return static_cast<QMacToolBarItem*>(ptr)->text().toUtf8().data();
}

void QMacToolBarItem_TimerEvent(void* ptr, void* event){
	static_cast<MyQMacToolBarItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBarItem_TimerEventDefault(void* ptr, void* event){
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMacToolBarItem_ChildEvent(void* ptr, void* event){
	static_cast<MyQMacToolBarItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBarItem_ChildEventDefault(void* ptr, void* event){
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::childEvent(static_cast<QChildEvent*>(event));
}

void QMacToolBarItem_CustomEvent(void* ptr, void* event){
	static_cast<MyQMacToolBarItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMacToolBarItem_CustomEventDefault(void* ptr, void* event){
	static_cast<QMacToolBarItem*>(ptr)->QMacToolBarItem::customEvent(static_cast<QEvent*>(event));
}

