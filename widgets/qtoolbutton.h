#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QToolButton_ArrowType(QtObjectPtr ptr);
int QToolButton_AutoRaise(QtObjectPtr ptr);
int QToolButton_PopupMode(QtObjectPtr ptr);
void QToolButton_SetArrowType(QtObjectPtr ptr, int ty);
void QToolButton_SetAutoRaise(QtObjectPtr ptr, int enable);
void QToolButton_SetPopupMode(QtObjectPtr ptr, int mode);
void QToolButton_SetToolButtonStyle(QtObjectPtr ptr, int style);
int QToolButton_ToolButtonStyle(QtObjectPtr ptr);
QtObjectPtr QToolButton_NewQToolButton(QtObjectPtr parent);
QtObjectPtr QToolButton_Menu(QtObjectPtr ptr);
void QToolButton_SetMenu(QtObjectPtr ptr, QtObjectPtr menu);
void QToolButton_ShowMenu(QtObjectPtr ptr);
void QToolButton_ConnectTriggered(QtObjectPtr ptr);
void QToolButton_DisconnectTriggered(QtObjectPtr ptr);
void QToolButton_DestroyQToolButton(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif