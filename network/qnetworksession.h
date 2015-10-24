#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkSession_NewQNetworkSession(QtObjectPtr connectionConfig, QtObjectPtr parent);
void QNetworkSession_Accept(QtObjectPtr ptr);
void QNetworkSession_Close(QtObjectPtr ptr);
void QNetworkSession_ConnectClosed(QtObjectPtr ptr);
void QNetworkSession_DisconnectClosed(QtObjectPtr ptr);
int QNetworkSession_Error(QtObjectPtr ptr);
char* QNetworkSession_ErrorString(QtObjectPtr ptr);
void QNetworkSession_Ignore(QtObjectPtr ptr);
int QNetworkSession_IsOpen(QtObjectPtr ptr);
void QNetworkSession_Migrate(QtObjectPtr ptr);
void QNetworkSession_ConnectNewConfigurationActivated(QtObjectPtr ptr);
void QNetworkSession_DisconnectNewConfigurationActivated(QtObjectPtr ptr);
void QNetworkSession_Open(QtObjectPtr ptr);
void QNetworkSession_ConnectOpened(QtObjectPtr ptr);
void QNetworkSession_DisconnectOpened(QtObjectPtr ptr);
void QNetworkSession_Reject(QtObjectPtr ptr);
char* QNetworkSession_SessionProperty(QtObjectPtr ptr, char* key);
void QNetworkSession_SetSessionProperty(QtObjectPtr ptr, char* key, char* value);
void QNetworkSession_ConnectStateChanged(QtObjectPtr ptr);
void QNetworkSession_DisconnectStateChanged(QtObjectPtr ptr);
void QNetworkSession_Stop(QtObjectPtr ptr);
int QNetworkSession_UsagePolicies(QtObjectPtr ptr);
void QNetworkSession_ConnectUsagePoliciesChanged(QtObjectPtr ptr);
void QNetworkSession_DisconnectUsagePoliciesChanged(QtObjectPtr ptr);
int QNetworkSession_WaitForOpened(QtObjectPtr ptr, int msecs);
void QNetworkSession_DestroyQNetworkSession(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif