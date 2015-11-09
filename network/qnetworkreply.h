#ifdef __cplusplus
extern "C" {
#endif

void QNetworkReply_Abort(void* ptr);
void* QNetworkReply_Attribute(void* ptr, int code);
void QNetworkReply_Close(void* ptr);
void QNetworkReply_ConnectEncrypted(void* ptr);
void QNetworkReply_DisconnectEncrypted(void* ptr);
int QNetworkReply_Error(void* ptr);
void QNetworkReply_ConnectFinished(void* ptr);
void QNetworkReply_DisconnectFinished(void* ptr);
int QNetworkReply_HasRawHeader(void* ptr, void* headerName);
void* QNetworkReply_Header(void* ptr, int header);
void QNetworkReply_IgnoreSslErrors(void* ptr);
int QNetworkReply_IsFinished(void* ptr);
int QNetworkReply_IsRunning(void* ptr);
void* QNetworkReply_Manager(void* ptr);
void QNetworkReply_ConnectMetaDataChanged(void* ptr);
void QNetworkReply_DisconnectMetaDataChanged(void* ptr);
int QNetworkReply_Operation(void* ptr);
void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(void* ptr);
void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(void* ptr);
void* QNetworkReply_RawHeader(void* ptr, void* headerName);
void QNetworkReply_SetSslConfiguration(void* ptr, void* config);
void QNetworkReply_DestroyQNetworkReply(void* ptr);

#ifdef __cplusplus
}
#endif