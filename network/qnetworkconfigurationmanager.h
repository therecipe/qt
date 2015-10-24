#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkConfigurationManager_NewQNetworkConfigurationManager(QtObjectPtr parent);
int QNetworkConfigurationManager_Capabilities(QtObjectPtr ptr);
int QNetworkConfigurationManager_IsOnline(QtObjectPtr ptr);
void QNetworkConfigurationManager_ConnectOnlineStateChanged(QtObjectPtr ptr);
void QNetworkConfigurationManager_DisconnectOnlineStateChanged(QtObjectPtr ptr);
void QNetworkConfigurationManager_ConnectUpdateCompleted(QtObjectPtr ptr);
void QNetworkConfigurationManager_DisconnectUpdateCompleted(QtObjectPtr ptr);
void QNetworkConfigurationManager_UpdateConfigurations(QtObjectPtr ptr);
void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif