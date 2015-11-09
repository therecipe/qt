#ifdef __cplusplus
extern "C" {
#endif

int QWebChannel_BlockUpdates(void* ptr);
void QWebChannel_SetBlockUpdates(void* ptr, int block);
void* QWebChannel_NewQWebChannel(void* parent);
void QWebChannel_ConnectBlockUpdatesChanged(void* ptr);
void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr);
void QWebChannel_ConnectTo(void* ptr, void* transport);
void QWebChannel_DeregisterObject(void* ptr, void* object);
void QWebChannel_DisconnectFrom(void* ptr, void* transport);
void QWebChannel_RegisterObject(void* ptr, char* id, void* object);
void QWebChannel_DestroyQWebChannel(void* ptr);

#ifdef __cplusplus
}
#endif