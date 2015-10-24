#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QWebSocket_Abort(QtObjectPtr ptr);
void QWebSocket_ConnectAboutToClose(QtObjectPtr ptr);
void QWebSocket_DisconnectAboutToClose(QtObjectPtr ptr);
char* QWebSocket_CloseReason(QtObjectPtr ptr);
void QWebSocket_ConnectConnected(QtObjectPtr ptr);
void QWebSocket_DisconnectConnected(QtObjectPtr ptr);
void QWebSocket_ConnectDisconnected(QtObjectPtr ptr);
void QWebSocket_DisconnectDisconnected(QtObjectPtr ptr);
int QWebSocket_Error(QtObjectPtr ptr);
char* QWebSocket_ErrorString(QtObjectPtr ptr);
int QWebSocket_Flush(QtObjectPtr ptr);
void QWebSocket_IgnoreSslErrors(QtObjectPtr ptr);
int QWebSocket_IsValid(QtObjectPtr ptr);
QtObjectPtr QWebSocket_MaskGenerator(QtObjectPtr ptr);
void QWebSocket_Open(QtObjectPtr ptr, char* url);
char* QWebSocket_Origin(QtObjectPtr ptr);
int QWebSocket_PauseMode(QtObjectPtr ptr);
char* QWebSocket_PeerName(QtObjectPtr ptr);
void QWebSocket_Ping(QtObjectPtr ptr, QtObjectPtr payload);
void QWebSocket_ConnectReadChannelFinished(QtObjectPtr ptr);
void QWebSocket_DisconnectReadChannelFinished(QtObjectPtr ptr);
char* QWebSocket_RequestUrl(QtObjectPtr ptr);
char* QWebSocket_ResourceName(QtObjectPtr ptr);
void QWebSocket_Resume(QtObjectPtr ptr);
void QWebSocket_SetMaskGenerator(QtObjectPtr ptr, QtObjectPtr maskGenerator);
void QWebSocket_SetPauseMode(QtObjectPtr ptr, int pauseMode);
void QWebSocket_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy);
void QWebSocket_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr sslConfiguration);
void QWebSocket_ConnectStateChanged(QtObjectPtr ptr);
void QWebSocket_DisconnectStateChanged(QtObjectPtr ptr);
void QWebSocket_ConnectTextFrameReceived(QtObjectPtr ptr);
void QWebSocket_DisconnectTextFrameReceived(QtObjectPtr ptr);
void QWebSocket_ConnectTextMessageReceived(QtObjectPtr ptr);
void QWebSocket_DisconnectTextMessageReceived(QtObjectPtr ptr);
void QWebSocket_DestroyQWebSocket(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif