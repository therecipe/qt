#include "qmactoolbaritem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QIcon>
#include <QMacToolBar>
#include <QString>
#include <QMacToolBarItem>
#include "_cgo_export.h"

class MyQMacToolBarItem: public QMacToolBarItem {
public:
void Signal_Activated(){callbackQMacToolBarItemActivated(this->objectName().toUtf8().data());};
};

QtObjectPtr QMacToolBarItem_NewQMacToolBarItem(QtObjectPtr parent){
	return new QMacToolBarItem(static_cast<QObject*>(parent));
}

void QMacToolBarItem_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMacToolBarItem*>(ptr), static_cast<void (QMacToolBarItem::*)()>(&QMacToolBarItem::activated), static_cast<MyQMacToolBarItem*>(ptr), static_cast<void (MyQMacToolBarItem::*)()>(&MyQMacToolBarItem::Signal_Activated));;
}

void QMacToolBarItem_DestroyQMacToolBarItem(QtObjectPtr ptr){
	static_cast<QMacToolBarItem*>(ptr)->~QMacToolBarItem();
}

int QMacToolBarItem_Selectable(QtObjectPtr ptr){
	return static_cast<QMacToolBarItem*>(ptr)->selectable();
}

void QMacToolBarItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QMacToolBarItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMacToolBarItem_SetSelectable(QtObjectPtr ptr, int selectable){
	static_cast<QMacToolBarItem*>(ptr)->setSelectable(selectable != 0);
}

void QMacToolBarItem_SetStandardItem(QtObjectPtr ptr, int standardItem){
	static_cast<QMacToolBarItem*>(ptr)->setStandardItem(static_cast<QMacToolBarItem::StandardItem>(standardItem));
}

void QMacToolBarItem_SetText(QtObjectPtr ptr, char* text){
	static_cast<QMacToolBarItem*>(ptr)->setText(QString(text));
}

int QMacToolBarItem_StandardItem(QtObjectPtr ptr){
	return static_cast<QMacToolBarItem*>(ptr)->standardItem();
}

char* QMacToolBarItem_Text(QtObjectPtr ptr){
	return static_cast<QMacToolBarItem*>(ptr)->text().toUtf8().data();
}

