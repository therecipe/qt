#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent);
void QNetworkSession_Accept(void* ptr);
void QNetworkSession_Close(void* ptr);
void QNetworkSession_ConnectClosed(void* ptr);
void QNetworkSession_DisconnectClosed(void* ptr);
int QNetworkSession_Error(void* ptr);
char* QNetworkSession_ErrorString(void* ptr);
void QNetworkSession_Ignore(void* ptr);
int QNetworkSession_IsOpen(void* ptr);
void QNetworkSession_Migrate(void* ptr);
void QNetworkSession_ConnectNewConfigurationActivated(void* ptr);
void QNetworkSession_DisconnectNewConfigurationActivated(void* ptr);
void QNetworkSession_Open(void* ptr);
void QNetworkSession_ConnectOpened(void* ptr);
void QNetworkSession_DisconnectOpened(void* ptr);
void QNetworkSession_Reject(void* ptr);
void* QNetworkSession_SessionProperty(void* ptr, char* key);
void QNetworkSession_SetSessionProperty(void* ptr, char* key, void* value);
void QNetworkSession_ConnectStateChanged(void* ptr);
void QNetworkSession_DisconnectStateChanged(void* ptr);
void QNetworkSession_Stop(void* ptr);
int QNetworkSession_UsagePolicies(void* ptr);
void QNetworkSession_ConnectUsagePoliciesChanged(void* ptr);
void QNetworkSession_DisconnectUsagePoliciesChanged(void* ptr);
int QNetworkSession_WaitForOpened(void* ptr, int msecs);
void QNetworkSession_DestroyQNetworkSession(void* ptr);

#ifdef __cplusplus
}
#endif