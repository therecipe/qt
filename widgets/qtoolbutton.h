#ifdef __cplusplus
extern "C" {
#endif

int QToolButton_ArrowType(void* ptr);
int QToolButton_AutoRaise(void* ptr);
int QToolButton_PopupMode(void* ptr);
void QToolButton_SetArrowType(void* ptr, int ty);
void QToolButton_SetAutoRaise(void* ptr, int enable);
void QToolButton_SetPopupMode(void* ptr, int mode);
void QToolButton_SetToolButtonStyle(void* ptr, int style);
int QToolButton_ToolButtonStyle(void* ptr);
void* QToolButton_NewQToolButton(void* parent);
void* QToolButton_Menu(void* ptr);
void QToolButton_SetMenu(void* ptr, void* menu);
void QToolButton_ShowMenu(void* ptr);
void QToolButton_ConnectTriggered(void* ptr);
void QToolButton_DisconnectTriggered(void* ptr);
void QToolButton_DestroyQToolButton(void* ptr);

#ifdef __cplusplus
}
#endif