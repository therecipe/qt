#include "qmainwindow.h"
#include <QString>
#include <QVariant>
#include <QObject>
#include <QMenuBar>
#include <QModelIndex>
#include <QWidget>
#include <QByteArray>
#include <QTabWidget>
#include <QStatusBar>
#include <QMetaObject>
#include <QSize>
#include <QDockWidget>
#include <QMenu>
#include <QUrl>
#include <QToolBar>
#include <QMainWindow>
#include "_cgo_export.h"

class MyQMainWindow: public QMainWindow {
public:
void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle){callbackQMainWindowToolButtonStyleChanged(this->objectName().toUtf8().data(), toolButtonStyle);};
};

int QMainWindow_DockOptions(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->dockOptions();
}

int QMainWindow_DocumentMode(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->documentMode();
}

int QMainWindow_IsAnimated(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->isAnimated();
}

int QMainWindow_IsDockNestingEnabled(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->isDockNestingEnabled();
}

void QMainWindow_SetAnimated(QtObjectPtr ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setAnimated", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockNestingEnabled(QtObjectPtr ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setDockNestingEnabled", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockOptions(QtObjectPtr ptr, int options){
	static_cast<QMainWindow*>(ptr)->setDockOptions(static_cast<QMainWindow::DockOption>(options));
}

void QMainWindow_SetDocumentMode(QtObjectPtr ptr, int enabled){
	static_cast<QMainWindow*>(ptr)->setDocumentMode(enabled != 0);
}

void QMainWindow_SetIconSize(QtObjectPtr ptr, QtObjectPtr iconSize){
	static_cast<QMainWindow*>(ptr)->setIconSize(*static_cast<QSize*>(iconSize));
}

void QMainWindow_SetTabShape(QtObjectPtr ptr, int tabShape){
	static_cast<QMainWindow*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(tabShape));
}

void QMainWindow_SetToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle){
	static_cast<QMainWindow*>(ptr)->setToolButtonStyle(static_cast<Qt::ToolButtonStyle>(toolButtonStyle));
}

void QMainWindow_SetUnifiedTitleAndToolBarOnMac(QtObjectPtr ptr, int set){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setUnifiedTitleAndToolBarOnMac", Q_ARG(bool, set != 0));
}

void QMainWindow_SplitDockWidget(QtObjectPtr ptr, QtObjectPtr first, QtObjectPtr second, int orientation){
	static_cast<QMainWindow*>(ptr)->splitDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second), static_cast<Qt::Orientation>(orientation));
}

int QMainWindow_TabShape(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->tabShape();
}

void QMainWindow_TabifyDockWidget(QtObjectPtr ptr, QtObjectPtr first, QtObjectPtr second){
	static_cast<QMainWindow*>(ptr)->tabifyDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second));
}

int QMainWindow_ToolButtonStyle(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->toolButtonStyle();
}

int QMainWindow_UnifiedTitleAndToolBarOnMac(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->unifiedTitleAndToolBarOnMac();
}

QtObjectPtr QMainWindow_NewQMainWindow(QtObjectPtr parent, int flags){
	return new QMainWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMainWindow_AddDockWidget(QtObjectPtr ptr, int area, QtObjectPtr dockwidget){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_AddDockWidget2(QtObjectPtr ptr, int area, QtObjectPtr dockwidget, int orientation){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget), static_cast<Qt::Orientation>(orientation));
}

QtObjectPtr QMainWindow_AddToolBar3(QtObjectPtr ptr, char* title){
	return static_cast<QMainWindow*>(ptr)->addToolBar(QString(title));
}

void QMainWindow_AddToolBar2(QtObjectPtr ptr, QtObjectPtr toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBar(QtObjectPtr ptr, int area, QtObjectPtr toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<Qt::ToolBarArea>(area), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBarBreak(QtObjectPtr ptr, int area){
	static_cast<QMainWindow*>(ptr)->addToolBarBreak(static_cast<Qt::ToolBarArea>(area));
}

QtObjectPtr QMainWindow_CentralWidget(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->centralWidget();
}

int QMainWindow_Corner(QtObjectPtr ptr, int corner){
	return static_cast<QMainWindow*>(ptr)->corner(static_cast<Qt::Corner>(corner));
}

QtObjectPtr QMainWindow_CreatePopupMenu(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->createPopupMenu();
}

int QMainWindow_DockWidgetArea(QtObjectPtr ptr, QtObjectPtr dockwidget){
	return static_cast<QMainWindow*>(ptr)->dockWidgetArea(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_InsertToolBar(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr toolbar){
	static_cast<QMainWindow*>(ptr)->insertToolBar(static_cast<QToolBar*>(before), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_InsertToolBarBreak(QtObjectPtr ptr, QtObjectPtr before){
	static_cast<QMainWindow*>(ptr)->insertToolBarBreak(static_cast<QToolBar*>(before));
}

QtObjectPtr QMainWindow_MenuBar(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->menuBar();
}

QtObjectPtr QMainWindow_MenuWidget(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->menuWidget();
}

void QMainWindow_RemoveDockWidget(QtObjectPtr ptr, QtObjectPtr dockwidget){
	static_cast<QMainWindow*>(ptr)->removeDockWidget(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_RemoveToolBar(QtObjectPtr ptr, QtObjectPtr toolbar){
	static_cast<QMainWindow*>(ptr)->removeToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_RemoveToolBarBreak(QtObjectPtr ptr, QtObjectPtr before){
	static_cast<QMainWindow*>(ptr)->removeToolBarBreak(static_cast<QToolBar*>(before));
}

int QMainWindow_RestoreDockWidget(QtObjectPtr ptr, QtObjectPtr dockwidget){
	return static_cast<QMainWindow*>(ptr)->restoreDockWidget(static_cast<QDockWidget*>(dockwidget));
}

int QMainWindow_RestoreState(QtObjectPtr ptr, QtObjectPtr state, int version){
	return static_cast<QMainWindow*>(ptr)->restoreState(*static_cast<QByteArray*>(state), version);
}

void QMainWindow_SetCentralWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QMainWindow*>(ptr)->setCentralWidget(static_cast<QWidget*>(widget));
}

void QMainWindow_SetCorner(QtObjectPtr ptr, int corner, int area){
	static_cast<QMainWindow*>(ptr)->setCorner(static_cast<Qt::Corner>(corner), static_cast<Qt::DockWidgetArea>(area));
}

void QMainWindow_SetMenuBar(QtObjectPtr ptr, QtObjectPtr menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuBar(static_cast<QMenuBar*>(menuBar));
}

void QMainWindow_SetMenuWidget(QtObjectPtr ptr, QtObjectPtr menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuWidget(static_cast<QWidget*>(menuBar));
}

void QMainWindow_SetStatusBar(QtObjectPtr ptr, QtObjectPtr statusbar){
	static_cast<QMainWindow*>(ptr)->setStatusBar(static_cast<QStatusBar*>(statusbar));
}

void QMainWindow_SetTabPosition(QtObjectPtr ptr, int areas, int tabPosition){
	static_cast<QMainWindow*>(ptr)->setTabPosition(static_cast<Qt::DockWidgetArea>(areas), static_cast<QTabWidget::TabPosition>(tabPosition));
}

QtObjectPtr QMainWindow_StatusBar(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->statusBar();
}

int QMainWindow_TabPosition(QtObjectPtr ptr, int area){
	return static_cast<QMainWindow*>(ptr)->tabPosition(static_cast<Qt::DockWidgetArea>(area));
}

QtObjectPtr QMainWindow_TakeCentralWidget(QtObjectPtr ptr){
	return static_cast<QMainWindow*>(ptr)->takeCentralWidget();
}

int QMainWindow_ToolBarArea(QtObjectPtr ptr, QtObjectPtr toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarArea(static_cast<QToolBar*>(toolbar));
}

int QMainWindow_ToolBarBreak(QtObjectPtr ptr, QtObjectPtr toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarBreak(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_ConnectToolButtonStyleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DisconnectToolButtonStyleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DestroyQMainWindow(QtObjectPtr ptr){
	static_cast<QMainWindow*>(ptr)->~QMainWindow();
}

