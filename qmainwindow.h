#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QMainWindow_New_QWidget_WindowType(QtObjectPtr parent, int flags);
void QMainWindow_Destroy(QtObjectPtr ptr);
void QMainWindow_AddToolBar_ToolBarArea_QToolBar(QtObjectPtr ptr, int area, QtObjectPtr toolbar);
void QMainWindow_AddToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar);
QtObjectPtr QMainWindow_AddToolBar_String(QtObjectPtr ptr, char* title);
void QMainWindow_AddToolBarBreak_ToolBarArea(QtObjectPtr ptr, int area);
QtObjectPtr QMainWindow_CentralWidget(QtObjectPtr ptr);
int QMainWindow_Corner_Corner(QtObjectPtr ptr, int corner);
QtObjectPtr QMainWindow_CreatePopupMenu(QtObjectPtr ptr);
int QMainWindow_DocumentMode(QtObjectPtr ptr);
void QMainWindow_InsertToolBar_QToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr toolbar);
void QMainWindow_InsertToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr before);
int QMainWindow_IsAnimated(QtObjectPtr ptr);
int QMainWindow_IsDockNestingEnabled(QtObjectPtr ptr);
QtObjectPtr QMainWindow_MenuBar(QtObjectPtr ptr);
QtObjectPtr QMainWindow_MenuWidget(QtObjectPtr ptr);
void QMainWindow_RemoveToolBar_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar);
void QMainWindow_RemoveToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr before);
void QMainWindow_SetCentralWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QMainWindow_SetCorner_Corner_DockWidgetArea(QtObjectPtr ptr, int corner, int area);
void QMainWindow_SetDocumentMode_Bool(QtObjectPtr ptr, int enabled);
void QMainWindow_SetMenuBar_QMenuBar(QtObjectPtr ptr, QtObjectPtr menuBar);
void QMainWindow_SetMenuWidget_QWidget(QtObjectPtr ptr, QtObjectPtr menuBar);
void QMainWindow_SetStatusBar_QStatusBar(QtObjectPtr ptr, QtObjectPtr statusbar);
void QMainWindow_SetToolButtonStyle_ToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle);
void QMainWindow_SetUnifiedTitleAndToolBarOnMac_Bool(QtObjectPtr ptr, int set);
QtObjectPtr QMainWindow_StatusBar(QtObjectPtr ptr);
QtObjectPtr QMainWindow_TakeCentralWidget(QtObjectPtr ptr);
int QMainWindow_ToolBarArea_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar);
int QMainWindow_ToolBarBreak_QToolBar(QtObjectPtr ptr, QtObjectPtr toolbar);
int QMainWindow_ToolButtonStyle(QtObjectPtr ptr);
int QMainWindow_UnifiedTitleAndToolBarOnMac(QtObjectPtr ptr);
//Public Slots
void QMainWindow_ConnectSlotSetAnimated(QtObjectPtr ptr);
void QMainWindow_DisconnectSlotSetAnimated(QtObjectPtr ptr);
void QMainWindow_SetAnimated_Bool(QtObjectPtr ptr, int enabled);
void QMainWindow_ConnectSlotSetDockNestingEnabled(QtObjectPtr ptr);
void QMainWindow_DisconnectSlotSetDockNestingEnabled(QtObjectPtr ptr);
void QMainWindow_SetDockNestingEnabled_Bool(QtObjectPtr ptr, int enabled);
//Signals
void QMainWindow_ConnectSignalToolButtonStyleChanged(QtObjectPtr ptr);
void QMainWindow_DisconnectSignalToolButtonStyleChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
