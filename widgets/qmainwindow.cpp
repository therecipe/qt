#include "qmainwindow.h"
#include <QVariant>
#include <QWidget>
#include <QByteArray>
#include <QToolBar>
#include <QSize>
#include <QString>
#include <QDockWidget>
#include <QStatusBar>
#include <QMenuBar>
#include <QTabWidget>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QMenu>
#include <QObject>
#include <QMainWindow>
#include "_cgo_export.h"

class MyQMainWindow: public QMainWindow {
public:
void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle){callbackQMainWindowToolButtonStyleChanged(this->objectName().toUtf8().data(), toolButtonStyle);};
};

int QMainWindow_DockOptions(void* ptr){
	return static_cast<QMainWindow*>(ptr)->dockOptions();
}

int QMainWindow_DocumentMode(void* ptr){
	return static_cast<QMainWindow*>(ptr)->documentMode();
}

int QMainWindow_IsAnimated(void* ptr){
	return static_cast<QMainWindow*>(ptr)->isAnimated();
}

int QMainWindow_IsDockNestingEnabled(void* ptr){
	return static_cast<QMainWindow*>(ptr)->isDockNestingEnabled();
}

void QMainWindow_SetAnimated(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setAnimated", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockNestingEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setDockNestingEnabled", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockOptions(void* ptr, int options){
	static_cast<QMainWindow*>(ptr)->setDockOptions(static_cast<QMainWindow::DockOption>(options));
}

void QMainWindow_SetDocumentMode(void* ptr, int enabled){
	static_cast<QMainWindow*>(ptr)->setDocumentMode(enabled != 0);
}

void QMainWindow_SetIconSize(void* ptr, void* iconSize){
	static_cast<QMainWindow*>(ptr)->setIconSize(*static_cast<QSize*>(iconSize));
}

void QMainWindow_SetTabShape(void* ptr, int tabShape){
	static_cast<QMainWindow*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(tabShape));
}

void QMainWindow_SetToolButtonStyle(void* ptr, int toolButtonStyle){
	static_cast<QMainWindow*>(ptr)->setToolButtonStyle(static_cast<Qt::ToolButtonStyle>(toolButtonStyle));
}

void QMainWindow_SetUnifiedTitleAndToolBarOnMac(void* ptr, int set){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setUnifiedTitleAndToolBarOnMac", Q_ARG(bool, set != 0));
}

void QMainWindow_SplitDockWidget(void* ptr, void* first, void* second, int orientation){
	static_cast<QMainWindow*>(ptr)->splitDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second), static_cast<Qt::Orientation>(orientation));
}

int QMainWindow_TabShape(void* ptr){
	return static_cast<QMainWindow*>(ptr)->tabShape();
}

void QMainWindow_TabifyDockWidget(void* ptr, void* first, void* second){
	static_cast<QMainWindow*>(ptr)->tabifyDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second));
}

int QMainWindow_ToolButtonStyle(void* ptr){
	return static_cast<QMainWindow*>(ptr)->toolButtonStyle();
}

int QMainWindow_UnifiedTitleAndToolBarOnMac(void* ptr){
	return static_cast<QMainWindow*>(ptr)->unifiedTitleAndToolBarOnMac();
}

void* QMainWindow_NewQMainWindow(void* parent, int flags){
	return new QMainWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMainWindow_AddDockWidget(void* ptr, int area, void* dockwidget){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_AddDockWidget2(void* ptr, int area, void* dockwidget, int orientation){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget), static_cast<Qt::Orientation>(orientation));
}

void* QMainWindow_AddToolBar3(void* ptr, char* title){
	return static_cast<QMainWindow*>(ptr)->addToolBar(QString(title));
}

void QMainWindow_AddToolBar2(void* ptr, void* toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBar(void* ptr, int area, void* toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<Qt::ToolBarArea>(area), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBarBreak(void* ptr, int area){
	static_cast<QMainWindow*>(ptr)->addToolBarBreak(static_cast<Qt::ToolBarArea>(area));
}

void* QMainWindow_CentralWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->centralWidget();
}

int QMainWindow_Corner(void* ptr, int corner){
	return static_cast<QMainWindow*>(ptr)->corner(static_cast<Qt::Corner>(corner));
}

void* QMainWindow_CreatePopupMenu(void* ptr){
	return static_cast<QMainWindow*>(ptr)->createPopupMenu();
}

int QMainWindow_DockWidgetArea(void* ptr, void* dockwidget){
	return static_cast<QMainWindow*>(ptr)->dockWidgetArea(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_InsertToolBar(void* ptr, void* before, void* toolbar){
	static_cast<QMainWindow*>(ptr)->insertToolBar(static_cast<QToolBar*>(before), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_InsertToolBarBreak(void* ptr, void* before){
	static_cast<QMainWindow*>(ptr)->insertToolBarBreak(static_cast<QToolBar*>(before));
}

void* QMainWindow_MenuBar(void* ptr){
	return static_cast<QMainWindow*>(ptr)->menuBar();
}

void* QMainWindow_MenuWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->menuWidget();
}

void QMainWindow_RemoveDockWidget(void* ptr, void* dockwidget){
	static_cast<QMainWindow*>(ptr)->removeDockWidget(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_RemoveToolBar(void* ptr, void* toolbar){
	static_cast<QMainWindow*>(ptr)->removeToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_RemoveToolBarBreak(void* ptr, void* before){
	static_cast<QMainWindow*>(ptr)->removeToolBarBreak(static_cast<QToolBar*>(before));
}

int QMainWindow_RestoreDockWidget(void* ptr, void* dockwidget){
	return static_cast<QMainWindow*>(ptr)->restoreDockWidget(static_cast<QDockWidget*>(dockwidget));
}

int QMainWindow_RestoreState(void* ptr, void* state, int version){
	return static_cast<QMainWindow*>(ptr)->restoreState(*static_cast<QByteArray*>(state), version);
}

void* QMainWindow_SaveState(void* ptr, int version){
	return new QByteArray(static_cast<QMainWindow*>(ptr)->saveState(version));
}

void QMainWindow_SetCentralWidget(void* ptr, void* widget){
	static_cast<QMainWindow*>(ptr)->setCentralWidget(static_cast<QWidget*>(widget));
}

void QMainWindow_SetCorner(void* ptr, int corner, int area){
	static_cast<QMainWindow*>(ptr)->setCorner(static_cast<Qt::Corner>(corner), static_cast<Qt::DockWidgetArea>(area));
}

void QMainWindow_SetMenuBar(void* ptr, void* menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuBar(static_cast<QMenuBar*>(menuBar));
}

void QMainWindow_SetMenuWidget(void* ptr, void* menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuWidget(static_cast<QWidget*>(menuBar));
}

void QMainWindow_SetStatusBar(void* ptr, void* statusbar){
	static_cast<QMainWindow*>(ptr)->setStatusBar(static_cast<QStatusBar*>(statusbar));
}

void QMainWindow_SetTabPosition(void* ptr, int areas, int tabPosition){
	static_cast<QMainWindow*>(ptr)->setTabPosition(static_cast<Qt::DockWidgetArea>(areas), static_cast<QTabWidget::TabPosition>(tabPosition));
}

void* QMainWindow_StatusBar(void* ptr){
	return static_cast<QMainWindow*>(ptr)->statusBar();
}

int QMainWindow_TabPosition(void* ptr, int area){
	return static_cast<QMainWindow*>(ptr)->tabPosition(static_cast<Qt::DockWidgetArea>(area));
}

void* QMainWindow_TakeCentralWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->takeCentralWidget();
}

int QMainWindow_ToolBarArea(void* ptr, void* toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarArea(static_cast<QToolBar*>(toolbar));
}

int QMainWindow_ToolBarBreak(void* ptr, void* toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarBreak(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_ConnectToolButtonStyleChanged(void* ptr){
	QObject::connect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DisconnectToolButtonStyleChanged(void* ptr){
	QObject::disconnect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DestroyQMainWindow(void* ptr){
	static_cast<QMainWindow*>(ptr)->~QMainWindow();
}

