#ifdef __cplusplus
extern "C" {
#endif

int QMdiSubWindow_KeyboardPageStep(void* ptr);
int QMdiSubWindow_KeyboardSingleStep(void* ptr);
void QMdiSubWindow_SetKeyboardPageStep(void* ptr, int step);
void QMdiSubWindow_SetKeyboardSingleStep(void* ptr, int step);
void* QMdiSubWindow_NewQMdiSubWindow(void* parent, int flags);
void QMdiSubWindow_ConnectAboutToActivate(void* ptr);
void QMdiSubWindow_DisconnectAboutToActivate(void* ptr);
int QMdiSubWindow_IsShaded(void* ptr);
void* QMdiSubWindow_MdiArea(void* ptr);
void QMdiSubWindow_SetOption(void* ptr, int option, int on);
void QMdiSubWindow_SetSystemMenu(void* ptr, void* systemMenu);
void QMdiSubWindow_SetWidget(void* ptr, void* widget);
void QMdiSubWindow_ShowShaded(void* ptr);
void QMdiSubWindow_ShowSystemMenu(void* ptr);
void* QMdiSubWindow_SystemMenu(void* ptr);
int QMdiSubWindow_TestOption(void* ptr, int option);
void* QMdiSubWindow_Widget(void* ptr);
void QMdiSubWindow_ConnectWindowStateChanged(void* ptr);
void QMdiSubWindow_DisconnectWindowStateChanged(void* ptr);
void QMdiSubWindow_DestroyQMdiSubWindow(void* ptr);

#ifdef __cplusplus
}
#endif