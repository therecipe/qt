#include "qmdiarea.h"
#include <QVariant>
#include <QModelIndex>
#include <QTabWidget>
#include <QMdiSubWindow>
#include <QMetaObject>
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QBrush>
#include <QObject>
#include <QMdiArea>
#include "_cgo_export.h"

class MyQMdiArea: public QMdiArea {
public:
void Signal_SubWindowActivated(QMdiSubWindow * window){callbackQMdiAreaSubWindowActivated(this->objectName().toUtf8().data(), window);};
};

int QMdiArea_ActivationOrder(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->activationOrder();
}

int QMdiArea_DocumentMode(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->documentMode();
}

void QMdiArea_SetActivationOrder(QtObjectPtr ptr, int order){
	static_cast<QMdiArea*>(ptr)->setActivationOrder(static_cast<QMdiArea::WindowOrder>(order));
}

void QMdiArea_SetBackground(QtObjectPtr ptr, QtObjectPtr background){
	static_cast<QMdiArea*>(ptr)->setBackground(*static_cast<QBrush*>(background));
}

void QMdiArea_SetDocumentMode(QtObjectPtr ptr, int enabled){
	static_cast<QMdiArea*>(ptr)->setDocumentMode(enabled != 0);
}

void QMdiArea_SetTabPosition(QtObjectPtr ptr, int position){
	static_cast<QMdiArea*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(position));
}

void QMdiArea_SetTabShape(QtObjectPtr ptr, int shape){
	static_cast<QMdiArea*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(shape));
}

void QMdiArea_SetTabsClosable(QtObjectPtr ptr, int closable){
	static_cast<QMdiArea*>(ptr)->setTabsClosable(closable != 0);
}

void QMdiArea_SetTabsMovable(QtObjectPtr ptr, int movable){
	static_cast<QMdiArea*>(ptr)->setTabsMovable(movable != 0);
}

void QMdiArea_SetViewMode(QtObjectPtr ptr, int mode){
	static_cast<QMdiArea*>(ptr)->setViewMode(static_cast<QMdiArea::ViewMode>(mode));
}

int QMdiArea_TabPosition(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->tabPosition();
}

int QMdiArea_TabShape(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->tabShape();
}

int QMdiArea_TabsClosable(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->tabsClosable();
}

int QMdiArea_TabsMovable(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->tabsMovable();
}

int QMdiArea_ViewMode(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->viewMode();
}

QtObjectPtr QMdiArea_NewQMdiArea(QtObjectPtr parent){
	return new QMdiArea(static_cast<QWidget*>(parent));
}

void QMdiArea_ActivateNextSubWindow(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activateNextSubWindow");
}

void QMdiArea_ActivatePreviousSubWindow(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activatePreviousSubWindow");
}

QtObjectPtr QMdiArea_ActiveSubWindow(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->activeSubWindow();
}

QtObjectPtr QMdiArea_AddSubWindow(QtObjectPtr ptr, QtObjectPtr widget, int windowFlags){
	return static_cast<QMdiArea*>(ptr)->addSubWindow(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(windowFlags));
}

void QMdiArea_CascadeSubWindows(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "cascadeSubWindows");
}

void QMdiArea_CloseActiveSubWindow(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeActiveSubWindow");
}

void QMdiArea_CloseAllSubWindows(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeAllSubWindows");
}

QtObjectPtr QMdiArea_CurrentSubWindow(QtObjectPtr ptr){
	return static_cast<QMdiArea*>(ptr)->currentSubWindow();
}

void QMdiArea_RemoveSubWindow(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QMdiArea*>(ptr)->removeSubWindow(static_cast<QWidget*>(widget));
}

void QMdiArea_SetActiveSubWindow(QtObjectPtr ptr, QtObjectPtr window){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "setActiveSubWindow", Q_ARG(QMdiSubWindow*, static_cast<QMdiSubWindow*>(window)));
}

void QMdiArea_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QMdiArea*>(ptr)->setOption(static_cast<QMdiArea::AreaOption>(option), on != 0);
}

void QMdiArea_ConnectSubWindowActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

void QMdiArea_DisconnectSubWindowActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

int QMdiArea_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QMdiArea*>(ptr)->testOption(static_cast<QMdiArea::AreaOption>(option));
}

void QMdiArea_TileSubWindows(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "tileSubWindows");
}

void QMdiArea_DestroyQMdiArea(QtObjectPtr ptr){
	static_cast<QMdiArea*>(ptr)->~QMdiArea();
}

