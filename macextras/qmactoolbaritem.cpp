#include "qmactoolbaritem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMacToolBar>
#include <QObject>
#include <QIcon>
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

