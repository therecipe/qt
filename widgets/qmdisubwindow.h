#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMdiSubWindow_KeyboardPageStep(QtObjectPtr ptr);
int QMdiSubWindow_KeyboardSingleStep(QtObjectPtr ptr);
void QMdiSubWindow_SetKeyboardPageStep(QtObjectPtr ptr, int step);
void QMdiSubWindow_SetKeyboardSingleStep(QtObjectPtr ptr, int step);
QtObjectPtr QMdiSubWindow_NewQMdiSubWindow(QtObjectPtr parent, int flags);
void QMdiSubWindow_ConnectAboutToActivate(QtObjectPtr ptr);
void QMdiSubWindow_DisconnectAboutToActivate(QtObjectPtr ptr);
int QMdiSubWindow_IsShaded(QtObjectPtr ptr);
QtObjectPtr QMdiSubWindow_MdiArea(QtObjectPtr ptr);
void QMdiSubWindow_SetOption(QtObjectPtr ptr, int option, int on);
void QMdiSubWindow_SetSystemMenu(QtObjectPtr ptr, QtObjectPtr systemMenu);
void QMdiSubWindow_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QMdiSubWindow_ShowShaded(QtObjectPtr ptr);
void QMdiSubWindow_ShowSystemMenu(QtObjectPtr ptr);
QtObjectPtr QMdiSubWindow_SystemMenu(QtObjectPtr ptr);
int QMdiSubWindow_TestOption(QtObjectPtr ptr, int option);
QtObjectPtr QMdiSubWindow_Widget(QtObjectPtr ptr);
void QMdiSubWindow_ConnectWindowStateChanged(QtObjectPtr ptr);
void QMdiSubWindow_DisconnectWindowStateChanged(QtObjectPtr ptr);
void QMdiSubWindow_DestroyQMdiSubWindow(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif