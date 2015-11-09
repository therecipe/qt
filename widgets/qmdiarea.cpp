#include "qmdiarea.h"
#include <QMdiSubWindow>
#include <QObject>
#include <QVariant>
#include <QTabWidget>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QBrush>
#include <QString>
#include <QUrl>
#include <QMdiArea>
#include "_cgo_export.h"

class MyQMdiArea: public QMdiArea {
public:
void Signal_SubWindowActivated(QMdiSubWindow * window){callbackQMdiAreaSubWindowActivated(this->objectName().toUtf8().data(), window);};
};

int QMdiArea_ActivationOrder(void* ptr){
	return static_cast<QMdiArea*>(ptr)->activationOrder();
}

void* QMdiArea_Background(void* ptr){
	return new QBrush(static_cast<QMdiArea*>(ptr)->background());
}

int QMdiArea_DocumentMode(void* ptr){
	return static_cast<QMdiArea*>(ptr)->documentMode();
}

void QMdiArea_SetActivationOrder(void* ptr, int order){
	static_cast<QMdiArea*>(ptr)->setActivationOrder(static_cast<QMdiArea::WindowOrder>(order));
}

void QMdiArea_SetBackground(void* ptr, void* background){
	static_cast<QMdiArea*>(ptr)->setBackground(*static_cast<QBrush*>(background));
}

void QMdiArea_SetDocumentMode(void* ptr, int enabled){
	static_cast<QMdiArea*>(ptr)->setDocumentMode(enabled != 0);
}

void QMdiArea_SetTabPosition(void* ptr, int position){
	static_cast<QMdiArea*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(position));
}

void QMdiArea_SetTabShape(void* ptr, int shape){
	static_cast<QMdiArea*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(shape));
}

void QMdiArea_SetTabsClosable(void* ptr, int closable){
	static_cast<QMdiArea*>(ptr)->setTabsClosable(closable != 0);
}

void QMdiArea_SetTabsMovable(void* ptr, int movable){
	static_cast<QMdiArea*>(ptr)->setTabsMovable(movable != 0);
}

void QMdiArea_SetViewMode(void* ptr, int mode){
	static_cast<QMdiArea*>(ptr)->setViewMode(static_cast<QMdiArea::ViewMode>(mode));
}

int QMdiArea_TabPosition(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabPosition();
}

int QMdiArea_TabShape(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabShape();
}

int QMdiArea_TabsClosable(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabsClosable();
}

int QMdiArea_TabsMovable(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabsMovable();
}

int QMdiArea_ViewMode(void* ptr){
	return static_cast<QMdiArea*>(ptr)->viewMode();
}

void* QMdiArea_NewQMdiArea(void* parent){
	return new QMdiArea(static_cast<QWidget*>(parent));
}

void QMdiArea_ActivateNextSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activateNextSubWindow");
}

void QMdiArea_ActivatePreviousSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activatePreviousSubWindow");
}

void* QMdiArea_ActiveSubWindow(void* ptr){
	return static_cast<QMdiArea*>(ptr)->activeSubWindow();
}

void* QMdiArea_AddSubWindow(void* ptr, void* widget, int windowFlags){
	return static_cast<QMdiArea*>(ptr)->addSubWindow(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(windowFlags));
}

void QMdiArea_CascadeSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "cascadeSubWindows");
}

void QMdiArea_CloseActiveSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeActiveSubWindow");
}

void QMdiArea_CloseAllSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeAllSubWindows");
}

void* QMdiArea_CurrentSubWindow(void* ptr){
	return static_cast<QMdiArea*>(ptr)->currentSubWindow();
}

void QMdiArea_RemoveSubWindow(void* ptr, void* widget){
	static_cast<QMdiArea*>(ptr)->removeSubWindow(static_cast<QWidget*>(widget));
}

void QMdiArea_SetActiveSubWindow(void* ptr, void* window){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "setActiveSubWindow", Q_ARG(QMdiSubWindow*, static_cast<QMdiSubWindow*>(window)));
}

void QMdiArea_SetOption(void* ptr, int option, int on){
	static_cast<QMdiArea*>(ptr)->setOption(static_cast<QMdiArea::AreaOption>(option), on != 0);
}

void QMdiArea_ConnectSubWindowActivated(void* ptr){
	QObject::connect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

void QMdiArea_DisconnectSubWindowActivated(void* ptr){
	QObject::disconnect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

int QMdiArea_TestOption(void* ptr, int option){
	return static_cast<QMdiArea*>(ptr)->testOption(static_cast<QMdiArea::AreaOption>(option));
}

void QMdiArea_TileSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "tileSubWindows");
}

void QMdiArea_DestroyQMdiArea(void* ptr){
	static_cast<QMdiArea*>(ptr)->~QMdiArea();
}

