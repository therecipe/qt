#include "qtoolbar.h"
#include <QVariant>
#include <QUrl>
#include <QSize>
#include <QAction>
#include <QPoint>
#include <QString>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QIcon>
#include <QModelIndex>
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

int QToolBar_AllowedAreas(void* ptr){
	return static_cast<QToolBar*>(ptr)->allowedAreas();
}

int QToolBar_IsFloatable(void* ptr){
	return static_cast<QToolBar*>(ptr)->isFloatable();
}

int QToolBar_IsFloating(void* ptr){
	return static_cast<QToolBar*>(ptr)->isFloating();
}

int QToolBar_IsMovable(void* ptr){
	return static_cast<QToolBar*>(ptr)->isMovable();
}

int QToolBar_Orientation(void* ptr){
	return static_cast<QToolBar*>(ptr)->orientation();
}

void QToolBar_SetAllowedAreas(void* ptr, int areas){
	static_cast<QToolBar*>(ptr)->setAllowedAreas(static_cast<Qt::ToolBarArea>(areas));
}

void QToolBar_SetFloatable(void* ptr, int floatable){
	static_cast<QToolBar*>(ptr)->setFloatable(floatable != 0);
}

void QToolBar_SetIconSize(void* ptr, void* iconSize){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(iconSize)));
}

void QToolBar_SetMovable(void* ptr, int movable){
	static_cast<QToolBar*>(ptr)->setMovable(movable != 0);
}

void QToolBar_SetOrientation(void* ptr, int orientation){
	static_cast<QToolBar*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QToolBar_SetToolButtonStyle(void* ptr, int toolButtonStyle){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(toolButtonStyle)));
}

int QToolBar_ToolButtonStyle(void* ptr){
	return static_cast<QToolBar*>(ptr)->toolButtonStyle();
}

void* QToolBar_NewQToolBar2(void* parent){
	return new QToolBar(static_cast<QWidget*>(parent));
}

void* QToolBar_NewQToolBar(char* title, void* parent){
	return new QToolBar(QString(title), static_cast<QWidget*>(parent));
}

void* QToolBar_ActionAt(void* ptr, void* p){
	return static_cast<QToolBar*>(ptr)->actionAt(*static_cast<QPoint*>(p));
}

void* QToolBar_ActionAt2(void* ptr, int x, int y){
	return static_cast<QToolBar*>(ptr)->actionAt(x, y);
}

void QToolBar_ConnectActionTriggered(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

void QToolBar_DisconnectActionTriggered(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

void* QToolBar_AddAction2(void* ptr, void* icon, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QToolBar_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QToolBar_AddAction(void* ptr, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text));
}

void* QToolBar_AddAction3(void* ptr, char* text, void* receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QToolBar_AddSeparator(void* ptr){
	return static_cast<QToolBar*>(ptr)->addSeparator();
}

void* QToolBar_AddWidget(void* ptr, void* widget){
	return static_cast<QToolBar*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QToolBar_ConnectAllowedAreasChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_DisconnectAllowedAreasChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_Clear(void* ptr){
	static_cast<QToolBar*>(ptr)->clear();
}

void* QToolBar_InsertSeparator(void* ptr, void* before){
	return static_cast<QToolBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

void* QToolBar_InsertWidget(void* ptr, void* before, void* widget){
	return static_cast<QToolBar*>(ptr)->insertWidget(static_cast<QAction*>(before), static_cast<QWidget*>(widget));
}

int QToolBar_IsAreaAllowed(void* ptr, int area){
	return static_cast<QToolBar*>(ptr)->isAreaAllowed(static_cast<Qt::ToolBarArea>(area));
}

void QToolBar_ConnectMovableChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_DisconnectMovableChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_ConnectOrientationChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

void QToolBar_DisconnectOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

void* QToolBar_ToggleViewAction(void* ptr){
	return static_cast<QToolBar*>(ptr)->toggleViewAction();
}

void QToolBar_ConnectToolButtonStyleChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_DisconnectToolButtonStyleChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_ConnectTopLevelChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_DisconnectTopLevelChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_ConnectVisibilityChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

void QToolBar_DisconnectVisibilityChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

void* QToolBar_WidgetForAction(void* ptr, void* action){
	return static_cast<QToolBar*>(ptr)->widgetForAction(static_cast<QAction*>(action));
}

void QToolBar_DestroyQToolBar(void* ptr){
	static_cast<QToolBar*>(ptr)->~QToolBar();
}

