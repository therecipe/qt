#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMdiArea_ActivationOrder(QtObjectPtr ptr);
int QMdiArea_DocumentMode(QtObjectPtr ptr);
void QMdiArea_SetActivationOrder(QtObjectPtr ptr, int order);
void QMdiArea_SetBackground(QtObjectPtr ptr, QtObjectPtr background);
void QMdiArea_SetDocumentMode(QtObjectPtr ptr, int enabled);
void QMdiArea_SetTabPosition(QtObjectPtr ptr, int position);
void QMdiArea_SetTabShape(QtObjectPtr ptr, int shape);
void QMdiArea_SetTabsClosable(QtObjectPtr ptr, int closable);
void QMdiArea_SetTabsMovable(QtObjectPtr ptr, int movable);
void QMdiArea_SetViewMode(QtObjectPtr ptr, int mode);
int QMdiArea_TabPosition(QtObjectPtr ptr);
int QMdiArea_TabShape(QtObjectPtr ptr);
int QMdiArea_TabsClosable(QtObjectPtr ptr);
int QMdiArea_TabsMovable(QtObjectPtr ptr);
int QMdiArea_ViewMode(QtObjectPtr ptr);
QtObjectPtr QMdiArea_NewQMdiArea(QtObjectPtr parent);
void QMdiArea_ActivateNextSubWindow(QtObjectPtr ptr);
void QMdiArea_ActivatePreviousSubWindow(QtObjectPtr ptr);
QtObjectPtr QMdiArea_ActiveSubWindow(QtObjectPtr ptr);
QtObjectPtr QMdiArea_AddSubWindow(QtObjectPtr ptr, QtObjectPtr widget, int windowFlags);
void QMdiArea_CascadeSubWindows(QtObjectPtr ptr);
void QMdiArea_CloseActiveSubWindow(QtObjectPtr ptr);
void QMdiArea_CloseAllSubWindows(QtObjectPtr ptr);
QtObjectPtr QMdiArea_CurrentSubWindow(QtObjectPtr ptr);
void QMdiArea_RemoveSubWindow(QtObjectPtr ptr, QtObjectPtr widget);
void QMdiArea_SetActiveSubWindow(QtObjectPtr ptr, QtObjectPtr window);
void QMdiArea_SetOption(QtObjectPtr ptr, int option, int on);
void QMdiArea_ConnectSubWindowActivated(QtObjectPtr ptr);
void QMdiArea_DisconnectSubWindowActivated(QtObjectPtr ptr);
int QMdiArea_TestOption(QtObjectPtr ptr, int option);
void QMdiArea_TileSubWindows(QtObjectPtr ptr);
void QMdiArea_DestroyQMdiArea(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif