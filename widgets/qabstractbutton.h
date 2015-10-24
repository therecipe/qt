#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractButton_AutoExclusive(QtObjectPtr ptr);
int QAbstractButton_AutoRepeat(QtObjectPtr ptr);
int QAbstractButton_AutoRepeatDelay(QtObjectPtr ptr);
int QAbstractButton_AutoRepeatInterval(QtObjectPtr ptr);
int QAbstractButton_IsCheckable(QtObjectPtr ptr);
int QAbstractButton_IsChecked(QtObjectPtr ptr);
int QAbstractButton_IsDown(QtObjectPtr ptr);
void QAbstractButton_SetAutoExclusive(QtObjectPtr ptr, int v);
void QAbstractButton_SetAutoRepeat(QtObjectPtr ptr, int v);
void QAbstractButton_SetAutoRepeatDelay(QtObjectPtr ptr, int v);
void QAbstractButton_SetAutoRepeatInterval(QtObjectPtr ptr, int v);
void QAbstractButton_SetCheckable(QtObjectPtr ptr, int v);
void QAbstractButton_SetChecked(QtObjectPtr ptr, int v);
void QAbstractButton_SetDown(QtObjectPtr ptr, int v);
void QAbstractButton_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QAbstractButton_SetIconSize(QtObjectPtr ptr, QtObjectPtr size);
void QAbstractButton_SetShortcut(QtObjectPtr ptr, QtObjectPtr key);
void QAbstractButton_SetText(QtObjectPtr ptr, char* text);
char* QAbstractButton_Text(QtObjectPtr ptr);
void QAbstractButton_Toggle(QtObjectPtr ptr);
void QAbstractButton_AnimateClick(QtObjectPtr ptr, int msec);
void QAbstractButton_Click(QtObjectPtr ptr);
void QAbstractButton_ConnectClicked(QtObjectPtr ptr);
void QAbstractButton_DisconnectClicked(QtObjectPtr ptr);
QtObjectPtr QAbstractButton_Group(QtObjectPtr ptr);
void QAbstractButton_ConnectPressed(QtObjectPtr ptr);
void QAbstractButton_DisconnectPressed(QtObjectPtr ptr);
void QAbstractButton_ConnectReleased(QtObjectPtr ptr);
void QAbstractButton_DisconnectReleased(QtObjectPtr ptr);
void QAbstractButton_ConnectToggled(QtObjectPtr ptr);
void QAbstractButton_DisconnectToggled(QtObjectPtr ptr);
void QAbstractButton_DestroyQAbstractButton(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif