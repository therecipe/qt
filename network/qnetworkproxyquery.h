#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkProxyQuery_NewQNetworkProxyQuery();
void* QNetworkProxyQuery_NewQNetworkProxyQuery7(void* networkConfiguration, char* hostname, int port, char* protocolTag, int queryType);
void* QNetworkProxyQuery_NewQNetworkProxyQuery6(void* networkConfiguration, void* requestUrl, int queryType);
void* QNetworkProxyQuery_NewQNetworkProxyQuery5(void* other);
void* QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType);
void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, int queryType);
int QNetworkProxyQuery_LocalPort(void* ptr);
char* QNetworkProxyQuery_PeerHostName(void* ptr);
int QNetworkProxyQuery_PeerPort(void* ptr);
char* QNetworkProxyQuery_ProtocolTag(void* ptr);
int QNetworkProxyQuery_QueryType(void* ptr);
void QNetworkProxyQuery_SetLocalPort(void* ptr, int port);
void QNetworkProxyQuery_SetNetworkConfiguration(void* ptr, void* networkConfiguration);
void QNetworkProxyQuery_SetPeerHostName(void* ptr, char* hostname);
void QNetworkProxyQuery_SetPeerPort(void* ptr, int port);
void QNetworkProxyQuery_SetProtocolTag(void* ptr, char* protocolTag);
void QNetworkProxyQuery_SetQueryType(void* ptr, int ty);
void QNetworkProxyQuery_SetUrl(void* ptr, void* url);
void QNetworkProxyQuery_Swap(void* ptr, void* other);
void QNetworkProxyQuery_DestroyQNetworkProxyQuery(void* ptr);

#ifdef __cplusplus
}
#endif