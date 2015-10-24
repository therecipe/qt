#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMainWindow_DockOptions(QtObjectPtr ptr);
int QMainWindow_DocumentMode(QtObjectPtr ptr);
int QMainWindow_IsAnimated(QtObjectPtr ptr);
int QMainWindow_IsDockNestingEnabled(QtObjectPtr ptr);
void QMainWindow_SetAnimated(QtObjectPtr ptr, int enabled);
void QMainWindow_SetDockNestingEnabled(QtObjectPtr ptr, int enabled);
void QMainWindow_SetDockOptions(QtObjectPtr ptr, int options);
void QMainWindow_SetDocumentMode(QtObjectPtr ptr, int enabled);
void QMainWindow_SetIconSize(QtObjectPtr ptr, QtObjectPtr iconSize);
void QMainWindow_SetTabShape(QtObjectPtr ptr, int tabShape);
void QMainWindow_SetToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle);
void QMainWindow_SetUnifiedTitleAndToolBarOnMac(QtObjectPtr ptr, int set);
void QMainWindow_SplitDockWidget(QtObjectPtr ptr, QtObjectPtr first, QtObjectPtr second, int orientation);
int QMainWindow_TabShape(QtObjectPtr ptr);
void QMainWindow_TabifyDockWidget(QtObjectPtr ptr, QtObjectPtr first, QtObjectPtr second);
int QMainWindow_ToolButtonStyle(QtObjectPtr ptr);
int QMainWindow_UnifiedTitleAndToolBarOnMac(QtObjectPtr ptr);
QtObjectPtr QMainWindow_NewQMainWindow(QtObjectPtr parent, int flags);
void QMainWindow_AddDockWidget(QtObjectPtr ptr, int area, QtObjectPtr dockwidget);
void QMainWindow_AddDockWidget2(QtObjectPtr ptr, int area, QtObjectPtr dockwidget, int orientation);
QtObjectPtr QMainWindow_AddToolBar3(QtObjectPtr ptr, char* title);
void QMainWindow_AddToolBar2(QtObjectPtr ptr, QtObjectPtr toolbar);
void QMainWindow_AddToolBar(QtObjectPtr ptr, int area, QtObjectPtr toolbar);
void QMainWindow_AddToolBarBreak(QtObjectPtr ptr, int area);
QtObjectPtr QMainWindow_CentralWidget(QtObjectPtr ptr);
int QMainWindow_Corner(QtObjectPtr ptr, int corner);
QtObjectPtr QMainWindow_CreatePopupMenu(QtObjectPtr ptr);
int QMainWindow_DockWidgetArea(QtObjectPtr ptr, QtObjectPtr dockwidget);
void QMainWindow_InsertToolBar(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr toolbar);
void QMainWindow_InsertToolBarBreak(QtObjectPtr ptr, QtObjectPtr before);
QtObjectPtr QMainWindow_MenuBar(QtObjectPtr ptr);
QtObjectPtr QMainWindow_MenuWidget(QtObjectPtr ptr);
void QMainWindow_RemoveDockWidget(QtObjectPtr ptr, QtObjectPtr dockwidget);
void QMainWindow_RemoveToolBar(QtObjectPtr ptr, QtObjectPtr toolbar);
void QMainWindow_RemoveToolBarBreak(QtObjectPtr ptr, QtObjectPtr before);
int QMainWindow_RestoreDockWidget(QtObjectPtr ptr, QtObjectPtr dockwidget);
int QMainWindow_RestoreState(QtObjectPtr ptr, QtObjectPtr state, int version);
void QMainWindow_SetCentralWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QMainWindow_SetCorner(QtObjectPtr ptr, int corner, int area);
void QMainWindow_SetMenuBar(QtObjectPtr ptr, QtObjectPtr menuBar);
void QMainWindow_SetMenuWidget(QtObjectPtr ptr, QtObjectPtr menuBar);
void QMainWindow_SetStatusBar(QtObjectPtr ptr, QtObjectPtr statusbar);
void QMainWindow_SetTabPosition(QtObjectPtr ptr, int areas, int tabPosition);
QtObjectPtr QMainWindow_StatusBar(QtObjectPtr ptr);
int QMainWindow_TabPosition(QtObjectPtr ptr, int area);
QtObjectPtr QMainWindow_TakeCentralWidget(QtObjectPtr ptr);
int QMainWindow_ToolBarArea(QtObjectPtr ptr, QtObjectPtr toolbar);
int QMainWindow_ToolBarBreak(QtObjectPtr ptr, QtObjectPtr toolbar);
void QMainWindow_ConnectToolButtonStyleChanged(QtObjectPtr ptr);
void QMainWindow_DisconnectToolButtonStyleChanged(QtObjectPtr ptr);
void QMainWindow_DestroyQMainWindow(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif