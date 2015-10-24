#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QNetworkReply_Abort(QtObjectPtr ptr);
char* QNetworkReply_Attribute(QtObjectPtr ptr, int code);
void QNetworkReply_Close(QtObjectPtr ptr);
void QNetworkReply_ConnectEncrypted(QtObjectPtr ptr);
void QNetworkReply_DisconnectEncrypted(QtObjectPtr ptr);
int QNetworkReply_Error(QtObjectPtr ptr);
void QNetworkReply_ConnectFinished(QtObjectPtr ptr);
void QNetworkReply_DisconnectFinished(QtObjectPtr ptr);
int QNetworkReply_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName);
char* QNetworkReply_Header(QtObjectPtr ptr, int header);
void QNetworkReply_IgnoreSslErrors(QtObjectPtr ptr);
int QNetworkReply_IsFinished(QtObjectPtr ptr);
int QNetworkReply_IsRunning(QtObjectPtr ptr);
QtObjectPtr QNetworkReply_Manager(QtObjectPtr ptr);
void QNetworkReply_ConnectMetaDataChanged(QtObjectPtr ptr);
void QNetworkReply_DisconnectMetaDataChanged(QtObjectPtr ptr);
int QNetworkReply_Operation(QtObjectPtr ptr);
void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
void QNetworkReply_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr config);
char* QNetworkReply_Url(QtObjectPtr ptr);
void QNetworkReply_DestroyQNetworkReply(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif