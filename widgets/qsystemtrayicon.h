#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSystemTrayIcon_IsVisible(QtObjectPtr ptr);
void QSystemTrayIcon_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QSystemTrayIcon_SetToolTip(QtObjectPtr ptr, char* tip);
void QSystemTrayIcon_SetVisible(QtObjectPtr ptr, int visible);
void QSystemTrayIcon_ShowMessage(QtObjectPtr ptr, char* title, char* message, int icon, int millisecondsTimeoutHint);
char* QSystemTrayIcon_ToolTip(QtObjectPtr ptr);
QtObjectPtr QSystemTrayIcon_NewQSystemTrayIcon(QtObjectPtr parent);
QtObjectPtr QSystemTrayIcon_NewQSystemTrayIcon2(QtObjectPtr icon, QtObjectPtr parent);
void QSystemTrayIcon_ConnectActivated(QtObjectPtr ptr);
void QSystemTrayIcon_DisconnectActivated(QtObjectPtr ptr);
QtObjectPtr QSystemTrayIcon_ContextMenu(QtObjectPtr ptr);
void QSystemTrayIcon_Hide(QtObjectPtr ptr);
int QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable();
void QSystemTrayIcon_ConnectMessageClicked(QtObjectPtr ptr);
void QSystemTrayIcon_DisconnectMessageClicked(QtObjectPtr ptr);
void QSystemTrayIcon_SetContextMenu(QtObjectPtr ptr, QtObjectPtr menu);
void QSystemTrayIcon_Show(QtObjectPtr ptr);
int QSystemTrayIcon_QSystemTrayIcon_SupportsMessages();
void QSystemTrayIcon_DestroyQSystemTrayIcon(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif