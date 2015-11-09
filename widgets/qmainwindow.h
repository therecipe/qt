#ifdef __cplusplus
extern "C" {
#endif

int QMainWindow_DockOptions(void* ptr);
int QMainWindow_DocumentMode(void* ptr);
int QMainWindow_IsAnimated(void* ptr);
int QMainWindow_IsDockNestingEnabled(void* ptr);
void QMainWindow_SetAnimated(void* ptr, int enabled);
void QMainWindow_SetDockNestingEnabled(void* ptr, int enabled);
void QMainWindow_SetDockOptions(void* ptr, int options);
void QMainWindow_SetDocumentMode(void* ptr, int enabled);
void QMainWindow_SetIconSize(void* ptr, void* iconSize);
void QMainWindow_SetTabShape(void* ptr, int tabShape);
void QMainWindow_SetToolButtonStyle(void* ptr, int toolButtonStyle);
void QMainWindow_SetUnifiedTitleAndToolBarOnMac(void* ptr, int set);
void QMainWindow_SplitDockWidget(void* ptr, void* first, void* second, int orientation);
int QMainWindow_TabShape(void* ptr);
void QMainWindow_TabifyDockWidget(void* ptr, void* first, void* second);
int QMainWindow_ToolButtonStyle(void* ptr);
int QMainWindow_UnifiedTitleAndToolBarOnMac(void* ptr);
void* QMainWindow_NewQMainWindow(void* parent, int flags);
void QMainWindow_AddDockWidget(void* ptr, int area, void* dockwidget);
void QMainWindow_AddDockWidget2(void* ptr, int area, void* dockwidget, int orientation);
void* QMainWindow_AddToolBar3(void* ptr, char* title);
void QMainWindow_AddToolBar2(void* ptr, void* toolbar);
void QMainWindow_AddToolBar(void* ptr, int area, void* toolbar);
void QMainWindow_AddToolBarBreak(void* ptr, int area);
void* QMainWindow_CentralWidget(void* ptr);
int QMainWindow_Corner(void* ptr, int corner);
void* QMainWindow_CreatePopupMenu(void* ptr);
int QMainWindow_DockWidgetArea(void* ptr, void* dockwidget);
void QMainWindow_InsertToolBar(void* ptr, void* before, void* toolbar);
void QMainWindow_InsertToolBarBreak(void* ptr, void* before);
void* QMainWindow_MenuBar(void* ptr);
void* QMainWindow_MenuWidget(void* ptr);
void QMainWindow_RemoveDockWidget(void* ptr, void* dockwidget);
void QMainWindow_RemoveToolBar(void* ptr, void* toolbar);
void QMainWindow_RemoveToolBarBreak(void* ptr, void* before);
int QMainWindow_RestoreDockWidget(void* ptr, void* dockwidget);
int QMainWindow_RestoreState(void* ptr, void* state, int version);
void* QMainWindow_SaveState(void* ptr, int version);
void QMainWindow_SetCentralWidget(void* ptr, void* widget);
void QMainWindow_SetCorner(void* ptr, int corner, int area);
void QMainWindow_SetMenuBar(void* ptr, void* menuBar);
void QMainWindow_SetMenuWidget(void* ptr, void* menuBar);
void QMainWindow_SetStatusBar(void* ptr, void* statusbar);
void QMainWindow_SetTabPosition(void* ptr, int areas, int tabPosition);
void* QMainWindow_StatusBar(void* ptr);
int QMainWindow_TabPosition(void* ptr, int area);
void* QMainWindow_TakeCentralWidget(void* ptr);
int QMainWindow_ToolBarArea(void* ptr, void* toolbar);
int QMainWindow_ToolBarBreak(void* ptr, void* toolbar);
void QMainWindow_ConnectToolButtonStyleChanged(void* ptr);
void QMainWindow_DisconnectToolButtonStyleChanged(void* ptr);
void QMainWindow_DestroyQMainWindow(void* ptr);

#ifdef __cplusplus
}
#endif