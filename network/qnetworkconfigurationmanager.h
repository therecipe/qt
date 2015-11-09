#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent);
int QNetworkConfigurationManager_Capabilities(void* ptr);
int QNetworkConfigurationManager_IsOnline(void* ptr);
void QNetworkConfigurationManager_ConnectOnlineStateChanged(void* ptr);
void QNetworkConfigurationManager_DisconnectOnlineStateChanged(void* ptr);
void QNetworkConfigurationManager_ConnectUpdateCompleted(void* ptr);
void QNetworkConfigurationManager_DisconnectUpdateCompleted(void* ptr);
void QNetworkConfigurationManager_UpdateConfigurations(void* ptr);
void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr);

#ifdef __cplusplus
}
#endif