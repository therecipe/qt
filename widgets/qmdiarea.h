#ifdef __cplusplus
extern "C" {
#endif

int QMdiArea_ActivationOrder(void* ptr);
void* QMdiArea_Background(void* ptr);
int QMdiArea_DocumentMode(void* ptr);
void QMdiArea_SetActivationOrder(void* ptr, int order);
void QMdiArea_SetBackground(void* ptr, void* background);
void QMdiArea_SetDocumentMode(void* ptr, int enabled);
void QMdiArea_SetTabPosition(void* ptr, int position);
void QMdiArea_SetTabShape(void* ptr, int shape);
void QMdiArea_SetTabsClosable(void* ptr, int closable);
void QMdiArea_SetTabsMovable(void* ptr, int movable);
void QMdiArea_SetViewMode(void* ptr, int mode);
int QMdiArea_TabPosition(void* ptr);
int QMdiArea_TabShape(void* ptr);
int QMdiArea_TabsClosable(void* ptr);
int QMdiArea_TabsMovable(void* ptr);
int QMdiArea_ViewMode(void* ptr);
void* QMdiArea_NewQMdiArea(void* parent);
void QMdiArea_ActivateNextSubWindow(void* ptr);
void QMdiArea_ActivatePreviousSubWindow(void* ptr);
void* QMdiArea_ActiveSubWindow(void* ptr);
void* QMdiArea_AddSubWindow(void* ptr, void* widget, int windowFlags);
void QMdiArea_CascadeSubWindows(void* ptr);
void QMdiArea_CloseActiveSubWindow(void* ptr);
void QMdiArea_CloseAllSubWindows(void* ptr);
void* QMdiArea_CurrentSubWindow(void* ptr);
void QMdiArea_RemoveSubWindow(void* ptr, void* widget);
void QMdiArea_SetActiveSubWindow(void* ptr, void* window);
void QMdiArea_SetOption(void* ptr, int option, int on);
void QMdiArea_ConnectSubWindowActivated(void* ptr);
void QMdiArea_DisconnectSubWindowActivated(void* ptr);
int QMdiArea_TestOption(void* ptr, int option);
void QMdiArea_TileSubWindows(void* ptr);
void QMdiArea_DestroyQMdiArea(void* ptr);

#ifdef __cplusplus
}
#endif