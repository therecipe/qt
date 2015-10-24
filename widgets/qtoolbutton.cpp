#include "qtoolbutton.h"
#include <QUrl>
#include <QMetaObject>
#include <QAction>
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QToolButton>
#include "_cgo_export.h"

class MyQToolButton: public QToolButton {
public:
void Signal_Triggered(QAction * action){callbackQToolButtonTriggered(this->objectName().toUtf8().data(), action);};
};

int QToolButton_ArrowType(QtObjectPtr ptr){
	return static_cast<QToolButton*>(ptr)->arrowType();
}

int QToolButton_AutoRaise(QtObjectPtr ptr){
	return static_cast<QToolButton*>(ptr)->autoRaise();
}

int QToolButton_PopupMode(QtObjectPtr ptr){
	return static_cast<QToolButton*>(ptr)->popupMode();
}

void QToolButton_SetArrowType(QtObjectPtr ptr, int ty){
	static_cast<QToolButton*>(ptr)->setArrowType(static_cast<Qt::ArrowType>(ty));
}

void QToolButton_SetAutoRaise(QtObjectPtr ptr, int enable){
	static_cast<QToolButton*>(ptr)->setAutoRaise(enable != 0);
}

void QToolButton_SetPopupMode(QtObjectPtr ptr, int mode){
	static_cast<QToolButton*>(ptr)->setPopupMode(static_cast<QToolButton::ToolButtonPopupMode>(mode));
}

void QToolButton_SetToolButtonStyle(QtObjectPtr ptr, int style){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(style)));
}

int QToolButton_ToolButtonStyle(QtObjectPtr ptr){
	return static_cast<QToolButton*>(ptr)->toolButtonStyle();
}

QtObjectPtr QToolButton_NewQToolButton(QtObjectPtr parent){
	return new QToolButton(static_cast<QWidget*>(parent));
}

QtObjectPtr QToolButton_Menu(QtObjectPtr ptr){
	return static_cast<QToolButton*>(ptr)->menu();
}

void QToolButton_SetMenu(QtObjectPtr ptr, QtObjectPtr menu){
	static_cast<QToolButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QToolButton_ShowMenu(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "showMenu");
}

void QToolButton_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DestroyQToolButton(QtObjectPtr ptr){
	static_cast<QToolButton*>(ptr)->~QToolButton();
}

