#ifdef __cplusplus
extern "C" {
#endif

int QAbstractButton_AutoExclusive(void* ptr);
int QAbstractButton_AutoRepeat(void* ptr);
int QAbstractButton_AutoRepeatDelay(void* ptr);
int QAbstractButton_AutoRepeatInterval(void* ptr);
int QAbstractButton_IsCheckable(void* ptr);
int QAbstractButton_IsChecked(void* ptr);
int QAbstractButton_IsDown(void* ptr);
void QAbstractButton_SetAutoExclusive(void* ptr, int v);
void QAbstractButton_SetAutoRepeat(void* ptr, int v);
void QAbstractButton_SetAutoRepeatDelay(void* ptr, int v);
void QAbstractButton_SetAutoRepeatInterval(void* ptr, int v);
void QAbstractButton_SetCheckable(void* ptr, int v);
void QAbstractButton_SetChecked(void* ptr, int v);
void QAbstractButton_SetDown(void* ptr, int v);
void QAbstractButton_SetIcon(void* ptr, void* icon);
void QAbstractButton_SetIconSize(void* ptr, void* size);
void QAbstractButton_SetShortcut(void* ptr, void* key);
void QAbstractButton_SetText(void* ptr, char* text);
char* QAbstractButton_Text(void* ptr);
void QAbstractButton_Toggle(void* ptr);
void QAbstractButton_AnimateClick(void* ptr, int msec);
void QAbstractButton_Click(void* ptr);
void QAbstractButton_ConnectClicked(void* ptr);
void QAbstractButton_DisconnectClicked(void* ptr);
void* QAbstractButton_Group(void* ptr);
void QAbstractButton_ConnectPressed(void* ptr);
void QAbstractButton_DisconnectPressed(void* ptr);
void QAbstractButton_ConnectReleased(void* ptr);
void QAbstractButton_DisconnectReleased(void* ptr);
void QAbstractButton_ConnectToggled(void* ptr);
void QAbstractButton_DisconnectToggled(void* ptr);
void QAbstractButton_DestroyQAbstractButton(void* ptr);

#ifdef __cplusplus
}
#endif