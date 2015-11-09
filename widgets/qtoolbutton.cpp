#include "qtoolbutton.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QAction>
#include <QString>
#include <QObject>
#include <QMenu>
#include <QToolButton>
#include "_cgo_export.h"

class MyQToolButton: public QToolButton {
public:
void Signal_Triggered(QAction * action){callbackQToolButtonTriggered(this->objectName().toUtf8().data(), action);};
};

int QToolButton_ArrowType(void* ptr){
	return static_cast<QToolButton*>(ptr)->arrowType();
}

int QToolButton_AutoRaise(void* ptr){
	return static_cast<QToolButton*>(ptr)->autoRaise();
}

int QToolButton_PopupMode(void* ptr){
	return static_cast<QToolButton*>(ptr)->popupMode();
}

void QToolButton_SetArrowType(void* ptr, int ty){
	static_cast<QToolButton*>(ptr)->setArrowType(static_cast<Qt::ArrowType>(ty));
}

void QToolButton_SetAutoRaise(void* ptr, int enable){
	static_cast<QToolButton*>(ptr)->setAutoRaise(enable != 0);
}

void QToolButton_SetPopupMode(void* ptr, int mode){
	static_cast<QToolButton*>(ptr)->setPopupMode(static_cast<QToolButton::ToolButtonPopupMode>(mode));
}

void QToolButton_SetToolButtonStyle(void* ptr, int style){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(style)));
}

int QToolButton_ToolButtonStyle(void* ptr){
	return static_cast<QToolButton*>(ptr)->toolButtonStyle();
}

void* QToolButton_NewQToolButton(void* parent){
	return new QToolButton(static_cast<QWidget*>(parent));
}

void* QToolButton_Menu(void* ptr){
	return static_cast<QToolButton*>(ptr)->menu();
}

void QToolButton_SetMenu(void* ptr, void* menu){
	static_cast<QToolButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QToolButton_ShowMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "showMenu");
}

void QToolButton_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DestroyQToolButton(void* ptr){
	static_cast<QToolButton*>(ptr)->~QToolButton();
}

