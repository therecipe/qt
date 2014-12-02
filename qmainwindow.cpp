#include "qmainwindow.h"
#include <QMainWindow>
#include "cgoexport.h"

//MyQMainWindow
class MyQMainWindow: public QMainWindow {
Q_OBJECT
public:
void Signal_ToolButtonStyleChanged() { callbackQMainWindow(this, QString("toolButtonStyleChanged").toUtf8().data()); };

signals:
void Slot_SetAnimated(bool enabled);
void Slot_SetDockNestingEnabled(bool enabled);

};
#include "qmainwindow.moc"


//Public Functions
QtObjectPtr QMainWindow_New_QWidget_WindowType(QtObjectPtr parent, int flags)
{
	return new QMainWindow(((QWidget*)(parent)), ((Qt::WindowType)(flags)));
}

void QMainWindow_Destroy(QtObjectPtr ptr)
{
	((QMainWindow*)(ptr))->~QMainWindow();
}

void QMainWindow_AddToolBar_ToolBarArea_QToolBar(QtObjectPtr ptr, int area, QtObjectPtr toolbar)
{
	((QMainWindow*)(ptr))->addToolBar(((Qt::ToolBarArea)(area)), ((QToolBar*)(toolbar)));
}

void QMainWindow_AddToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar)
{
	((QMainWindow*)(ptr))->addToolBar(((QToolBar*)(toolbar)));
}

QtObjectPtr QMainWindow_AddToolBar_String(QtObjectPtr ptr, char* title)
{
	return ((QMainWindow*)(ptr))->addToolBar(QString(title));
}

void QMainWindow_AddToolBarBreak_ToolBarArea(QtObjectPtr ptr, int area)
{
	((QMainWindow*)(ptr))->addToolBarBreak(((Qt::ToolBarArea)(area)));
}

QtObjectPtr QMainWindow_CentralWidget(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->centralWidget();
}

int QMainWindow_Corner_Corner(QtObjectPtr ptr, int corner)
{
	return ((QMainWindow*)(ptr))->corner(((Qt::Corner)(corner)));
}

QtObjectPtr QMainWindow_CreatePopupMenu(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->createPopupMenu();
}

int QMainWindow_DocumentMode(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->documentMode();
}

void QMainWindow_InsertToolBar_QToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr toolbar)
{
	((QMainWindow*)(ptr))->insertToolBar(((QToolBar*)(before)), ((QToolBar*)(toolbar)));
}

void QMainWindow_InsertToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr before)
{
	((QMainWindow*)(ptr))->insertToolBarBreak(((QToolBar*)(before)));
}

int QMainWindow_IsAnimated(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->isAnimated();
}

int QMainWindow_IsDockNestingEnabled(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->isDockNestingEnabled();
}

QtObjectPtr QMainWindow_MenuBar(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->menuBar();
}

QtObjectPtr QMainWindow_MenuWidget(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->menuWidget();
}

void QMainWindow_RemoveToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar)
{
	((QMainWindow*)(ptr))->removeToolBar(((QToolBar*)(toolbar)));
}

void QMainWindow_RemoveToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr before)
{
	((QMainWindow*)(ptr))->removeToolBarBreak(((QToolBar*)(before)));
}

void QMainWindow_SetCentralWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QMainWindow*)(ptr))->setCentralWidget(((QWidget*)(widget)));
}

void QMainWindow_SetCorner_Corner_DockWidgetArea(QtObjectPtr ptr, int corner, int area)
{
	((QMainWindow*)(ptr))->setCorner(((Qt::Corner)(corner)), ((Qt::DockWidgetArea)(area)));
}

void QMainWindow_SetDocumentMode_Bool(QtObjectPtr ptr, int enabled)
{
	((QMainWindow*)(ptr))->setDocumentMode(enabled != 0);
}

void QMainWindow_SetMenuBar_QMenuBar(QtObjectPtr ptr, QtObjectPtr menuBar)
{
	((QMainWindow*)(ptr))->setMenuBar(((QMenuBar*)(menuBar)));
}

void QMainWindow_SetMenuWidget_QWidget(QtObjectPtr ptr, QtObjectPtr menuBar)
{
	((QMainWindow*)(ptr))->setMenuWidget(((QWidget*)(menuBar)));
}

void QMainWindow_SetStatusBar_QStatusBar(QtObjectPtr ptr, QtObjectPtr statusbar)
{
	((QMainWindow*)(ptr))->setStatusBar(((QStatusBar*)(statusbar)));
}

void QMainWindow_SetToolButtonStyle_ToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle)
{
	((QMainWindow*)(ptr))->setToolButtonStyle(((Qt::ToolButtonStyle)(toolButtonStyle)));
}

void QMainWindow_SetUnifiedTitleAndToolBarOnMac_Bool(QtObjectPtr ptr, int set)
{
	((QMainWindow*)(ptr))->setUnifiedTitleAndToolBarOnMac(set != 0);
}

QtObjectPtr QMainWindow_StatusBar(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->statusBar();
}

QtObjectPtr QMainWindow_TakeCentralWidget(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->takeCentralWidget();
}

int QMainWindow_ToolBarArea_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar)
{
	return ((QMainWindow*)(ptr))->toolBarArea(((QToolBar*)(toolbar)));
}

int QMainWindow_ToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar)
{
	return ((QMainWindow*)(ptr))->toolBarBreak(((QToolBar*)(toolbar)));
}

int QMainWindow_ToolButtonStyle(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->toolButtonStyle();
}

int QMainWindow_UnifiedTitleAndToolBarOnMac(QtObjectPtr ptr)
{
	return ((QMainWindow*)(ptr))->unifiedTitleAndToolBarOnMac();
}

//Public Slots
void QMainWindow_ConnectSlotSetAnimated(QtObjectPtr ptr)
{
	QObject::connect(((MyQMainWindow*)(ptr)), &MyQMainWindow::Slot_SetAnimated, ((QMainWindow*)(ptr)), &QMainWindow::setAnimated, Qt::QueuedConnection);
}

void QMainWindow_DisconnectSlotSetAnimated(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQMainWindow*)(ptr)), &MyQMainWindow::Slot_SetAnimated, ((QMainWindow*)(ptr)), &QMainWindow::setAnimated);
}

void QMainWindow_SetAnimated_Bool(QtObjectPtr ptr, int enabled)
{
	((MyQMainWindow*)(ptr))->Slot_SetAnimated(enabled != 0);
}

void QMainWindow_ConnectSlotSetDockNestingEnabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQMainWindow*)(ptr)), &MyQMainWindow::Slot_SetDockNestingEnabled, ((QMainWindow*)(ptr)), &QMainWindow::setDockNestingEnabled, Qt::QueuedConnection);
}

void QMainWindow_DisconnectSlotSetDockNestingEnabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQMainWindow*)(ptr)), &MyQMainWindow::Slot_SetDockNestingEnabled, ((QMainWindow*)(ptr)), &QMainWindow::setDockNestingEnabled);
}

void QMainWindow_SetDockNestingEnabled_Bool(QtObjectPtr ptr, int enabled)
{
	((MyQMainWindow*)(ptr))->Slot_SetDockNestingEnabled(enabled != 0);
}

//Signals
void QMainWindow_ConnectSignalToolButtonStyleChanged(QtObjectPtr ptr)
{
	QObject::connect(((QMainWindow*)(ptr)), &QMainWindow::toolButtonStyleChanged, ((MyQMainWindow*)(ptr)), &MyQMainWindow::Signal_ToolButtonStyleChanged, Qt::QueuedConnection);
}

void QMainWindow_DisconnectSignalToolButtonStyleChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QMainWindow*)(ptr)), &QMainWindow::toolButtonStyleChanged, ((MyQMainWindow*)(ptr)), &MyQMainWindow::Signal_ToolButtonStyleChanged);
}

