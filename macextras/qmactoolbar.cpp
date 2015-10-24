#include "qmactoolbar.h"
#include <QModelIndex>
#include <QObject>
#include <QIcon>
#include <QWindow>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMacToolBar>
#include "_cgo_export.h"

class MyQMacToolBar: public QMacToolBar {
public:
};

QtObjectPtr QMacToolBar_NewQMacToolBar(QtObjectPtr parent){
	return new QMacToolBar(static_cast<QObject*>(parent));
}

QtObjectPtr QMacToolBar_NewQMacToolBar2(char* identifier, QtObjectPtr parent){
	return new QMacToolBar(QString(identifier), static_cast<QObject*>(parent));
}

QtObjectPtr QMacToolBar_AddAllowedItem(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QMacToolBar*>(ptr)->addAllowedItem(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QMacToolBar_AddItem(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QMacToolBar*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text));
}

void QMacToolBar_AddSeparator(QtObjectPtr ptr){
	static_cast<QMacToolBar*>(ptr)->addSeparator();
}

void QMacToolBar_AttachToWindow(QtObjectPtr ptr, QtObjectPtr window){
	static_cast<QMacToolBar*>(ptr)->attachToWindow(static_cast<QWindow*>(window));
}

void QMacToolBar_DetachFromWindow(QtObjectPtr ptr){
	static_cast<QMacToolBar*>(ptr)->detachFromWindow();
}

void QMacToolBar_DestroyQMacToolBar(QtObjectPtr ptr){
	static_cast<QMacToolBar*>(ptr)->~QMacToolBar();
}

