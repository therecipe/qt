#include "qmactoolbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWindow>
#include <QObject>
#include <QIcon>
#include <QMacToolBar>
#include "_cgo_export.h"

class MyQMacToolBar: public QMacToolBar {
public:
};

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

#include "qmactoolbaritem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMacToolBar>
#include <QObject>
#include <QIcon>
#include <QString>
#include <QMacToolBarItem>
#include "_cgo_export.h"

class MyQMacToolBarItem: public QMacToolBarItem {
public:
void Signal_Activated(){callbackQMacToolBarItemActivated(this->objectName().toUtf8().data());};
};

void* QMacToolBarItem_NewQMacToolBarItem(void* parent){
	return new QMacToolBarItem(static_cast<QObject*>(parent));
}

void QMacToolBarItem_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr){
	static_cast<QMacToolBarItem*>(ptr)->~QMacToolBarItem();
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

#include "qmacpasteboardmime.h"
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QMacPasteboardMime>
#include "_cgo_export.h"

class MyQMacPasteboardMime: public QMacPasteboardMime {
public:
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

