#ifdef __cplusplus
extern "C" {
#endif

int QSystemTrayIcon_IsVisible(void* ptr);
void QSystemTrayIcon_SetIcon(void* ptr, void* icon);
void QSystemTrayIcon_SetToolTip(void* ptr, char* tip);
void QSystemTrayIcon_SetVisible(void* ptr, int visible);
void QSystemTrayIcon_ShowMessage(void* ptr, char* title, char* message, int icon, int millisecondsTimeoutHint);
char* QSystemTrayIcon_ToolTip(void* ptr);
void* QSystemTrayIcon_NewQSystemTrayIcon(void* parent);
void* QSystemTrayIcon_NewQSystemTrayIcon2(void* icon, void* parent);
void QSystemTrayIcon_ConnectActivated(void* ptr);
void QSystemTrayIcon_DisconnectActivated(void* ptr);
void* QSystemTrayIcon_ContextMenu(void* ptr);
void QSystemTrayIcon_Hide(void* ptr);
int QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable();
void QSystemTrayIcon_ConnectMessageClicked(void* ptr);
void QSystemTrayIcon_DisconnectMessageClicked(void* ptr);
void QSystemTrayIcon_SetContextMenu(void* ptr, void* menu);
void QSystemTrayIcon_Show(void* ptr);
int QSystemTrayIcon_QSystemTrayIcon_SupportsMessages();
void QSystemTrayIcon_DestroyQSystemTrayIcon(void* ptr);

#ifdef __cplusplus
}
#endif