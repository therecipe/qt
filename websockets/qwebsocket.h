#ifdef __cplusplus
extern "C" {
#endif

void QWebSocket_Abort(void* ptr);
void QWebSocket_ConnectAboutToClose(void* ptr);
void QWebSocket_DisconnectAboutToClose(void* ptr);
void QWebSocket_ConnectBinaryFrameReceived(void* ptr);
void QWebSocket_DisconnectBinaryFrameReceived(void* ptr);
void QWebSocket_ConnectBinaryMessageReceived(void* ptr);
void QWebSocket_DisconnectBinaryMessageReceived(void* ptr);
char* QWebSocket_CloseReason(void* ptr);
void QWebSocket_ConnectConnected(void* ptr);
void QWebSocket_DisconnectConnected(void* ptr);
void QWebSocket_ConnectDisconnected(void* ptr);
void QWebSocket_DisconnectDisconnected(void* ptr);
int QWebSocket_Error(void* ptr);
char* QWebSocket_ErrorString(void* ptr);
int QWebSocket_Flush(void* ptr);
void QWebSocket_IgnoreSslErrors(void* ptr);
int QWebSocket_IsValid(void* ptr);
void* QWebSocket_MaskGenerator(void* ptr);
void QWebSocket_Open(void* ptr, void* url);
char* QWebSocket_Origin(void* ptr);
int QWebSocket_PauseMode(void* ptr);
char* QWebSocket_PeerName(void* ptr);
void QWebSocket_Ping(void* ptr, void* payload);
void QWebSocket_ConnectReadChannelFinished(void* ptr);
void QWebSocket_DisconnectReadChannelFinished(void* ptr);
char* QWebSocket_ResourceName(void* ptr);
void QWebSocket_Resume(void* ptr);
void QWebSocket_SetMaskGenerator(void* ptr, void* maskGenerator);
void QWebSocket_SetPauseMode(void* ptr, int pauseMode);
void QWebSocket_SetProxy(void* ptr, void* networkProxy);
void QWebSocket_SetSslConfiguration(void* ptr, void* sslConfiguration);
void QWebSocket_ConnectStateChanged(void* ptr);
void QWebSocket_DisconnectStateChanged(void* ptr);
void QWebSocket_ConnectTextFrameReceived(void* ptr);
void QWebSocket_DisconnectTextFrameReceived(void* ptr);
void QWebSocket_ConnectTextMessageReceived(void* ptr);
void QWebSocket_DisconnectTextMessageReceived(void* ptr);
void QWebSocket_DestroyQWebSocket(void* ptr);

#ifdef __cplusplus
}
#endif