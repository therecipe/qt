#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery();
QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery7(QtObjectPtr networkConfiguration, char* hostname, int port, char* protocolTag, int queryType);
QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery6(QtObjectPtr networkConfiguration, char* requestUrl, int queryType);
QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery5(QtObjectPtr other);
QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType);
QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery2(char* requestUrl, int queryType);
int QNetworkProxyQuery_LocalPort(QtObjectPtr ptr);
char* QNetworkProxyQuery_PeerHostName(QtObjectPtr ptr);
int QNetworkProxyQuery_PeerPort(QtObjectPtr ptr);
char* QNetworkProxyQuery_ProtocolTag(QtObjectPtr ptr);
int QNetworkProxyQuery_QueryType(QtObjectPtr ptr);
void QNetworkProxyQuery_SetLocalPort(QtObjectPtr ptr, int port);
void QNetworkProxyQuery_SetNetworkConfiguration(QtObjectPtr ptr, QtObjectPtr networkConfiguration);
void QNetworkProxyQuery_SetPeerHostName(QtObjectPtr ptr, char* hostname);
void QNetworkProxyQuery_SetPeerPort(QtObjectPtr ptr, int port);
void QNetworkProxyQuery_SetProtocolTag(QtObjectPtr ptr, char* protocolTag);
void QNetworkProxyQuery_SetQueryType(QtObjectPtr ptr, int ty);
void QNetworkProxyQuery_SetUrl(QtObjectPtr ptr, char* url);
void QNetworkProxyQuery_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QNetworkProxyQuery_Url(QtObjectPtr ptr);
void QNetworkProxyQuery_DestroyQNetworkProxyQuery(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif