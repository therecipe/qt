#include "qtoolbar.h"
#include <QVariant>
#include <QModelIndex>
#include <QSize>
#include <QObject>
#include <QPoint>
#include <QIcon>
#include <QString>
#include <QUrl>
#include <QMetaObject>
#include <QAction>
#include <QWidget>
#include <QToolBar>
#include "_cgo_export.h"

class MyQToolBar: public QToolBar {
public:
void Signal_ActionTriggered(QAction * action){callbackQToolBarActionTriggered(this->objectName().toUtf8().data(), action);};
void Signal_AllowedAreasChanged(Qt::ToolBarAreas allowedAreas){callbackQToolBarAllowedAreasChanged(this->objectName().toUtf8().data(), allowedAreas);};
void Signal_MovableChanged(bool movable){callbackQToolBarMovableChanged(this->objectName().toUtf8().data(), movable);};
void Signal_OrientationChanged(Qt::Orientation orientation){callbackQToolBarOrientationChanged(this->objectName().toUtf8().data(), orientation);};
void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle){callbackQToolBarToolButtonStyleChanged(this->objectName().toUtf8().data(), toolButtonStyle);};
void Signal_TopLevelChanged(bool topLevel){callbackQToolBarTopLevelChanged(this->objectName().toUtf8().data(), topLevel);};
void Signal_VisibilityChanged(bool visible){callbackQToolBarVisibilityChanged(this->objectName().toUtf8().data(), visible);};
};

int QToolBar_AllowedAreas(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->allowedAreas();
}

int QToolBar_IsFloatable(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->isFloatable();
}

int QToolBar_IsFloating(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->isFloating();
}

int QToolBar_IsMovable(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->isMovable();
}

int QToolBar_Orientation(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->orientation();
}

void QToolBar_SetAllowedAreas(QtObjectPtr ptr, int areas){
	static_cast<QToolBar*>(ptr)->setAllowedAreas(static_cast<Qt::ToolBarArea>(areas));
}

void QToolBar_SetFloatable(QtObjectPtr ptr, int floatable){
	static_cast<QToolBar*>(ptr)->setFloatable(floatable != 0);
}

void QToolBar_SetIconSize(QtObjectPtr ptr, QtObjectPtr iconSize){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(iconSize)));
}

void QToolBar_SetMovable(QtObjectPtr ptr, int movable){
	static_cast<QToolBar*>(ptr)->setMovable(movable != 0);
}

void QToolBar_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QToolBar*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QToolBar_SetToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(toolButtonStyle)));
}

int QToolBar_ToolButtonStyle(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->toolButtonStyle();
}

QtObjectPtr QToolBar_NewQToolBar2(QtObjectPtr parent){
	return new QToolBar(static_cast<QWidget*>(parent));
}

QtObjectPtr QToolBar_NewQToolBar(char* title, QtObjectPtr parent){
	return new QToolBar(QString(title), static_cast<QWidget*>(parent));
}

QtObjectPtr QToolBar_ActionAt(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QToolBar*>(ptr)->actionAt(*static_cast<QPoint*>(p));
}

QtObjectPtr QToolBar_ActionAt2(QtObjectPtr ptr, int x, int y){
	return static_cast<QToolBar*>(ptr)->actionAt(x, y);
}

void QToolBar_ConnectActionTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

void QToolBar_DisconnectActionTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

QtObjectPtr QToolBar_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QToolBar_AddAction4(QtObjectPtr ptr, QtObjectPtr icon, char* text, QtObjectPtr receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

QtObjectPtr QToolBar_AddAction(QtObjectPtr ptr, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text));
}

QtObjectPtr QToolBar_AddAction3(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

QtObjectPtr QToolBar_AddSeparator(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->addSeparator();
}

QtObjectPtr QToolBar_AddWidget(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QToolBar*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QToolBar_ConnectAllowedAreasChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_DisconnectAllowedAreasChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_Clear(QtObjectPtr ptr){
	static_cast<QToolBar*>(ptr)->clear();
}

QtObjectPtr QToolBar_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before){
	return static_cast<QToolBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

QtObjectPtr QToolBar_InsertWidget(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr widget){
	return static_cast<QToolBar*>(ptr)->insertWidget(static_cast<QAction*>(before), static_cast<QWidget*>(widget));
}

int QToolBar_IsAreaAllowed(QtObjectPtr ptr, int area){
	return static_cast<QToolBar*>(ptr)->isAreaAllowed(static_cast<Qt::ToolBarArea>(area));
}

void QToolBar_ConnectMovableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_DisconnectMovableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_ConnectOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

void QToolBar_DisconnectOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

QtObjectPtr QToolBar_ToggleViewAction(QtObjectPtr ptr){
	return static_cast<QToolBar*>(ptr)->toggleViewAction();
}

void QToolBar_ConnectToolButtonStyleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_DisconnectToolButtonStyleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_ConnectTopLevelChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_DisconnectTopLevelChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_ConnectVisibilityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

void QToolBar_DisconnectVisibilityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

QtObjectPtr QToolBar_WidgetForAction(QtObjectPtr ptr, QtObjectPtr action){
	return static_cast<QToolBar*>(ptr)->widgetForAction(static_cast<QAction*>(action));
}

void QToolBar_DestroyQToolBar(QtObjectPtr ptr){
	static_cast<QToolBar*>(ptr)->~QToolBar();
}

