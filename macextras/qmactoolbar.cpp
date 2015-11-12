#include "qmactoolbar.h"
#include <QIcon>
#include <QWindow>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

