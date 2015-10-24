#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QWebChannel_BlockUpdates(QtObjectPtr ptr);
void QWebChannel_SetBlockUpdates(QtObjectPtr ptr, int block);
QtObjectPtr QWebChannel_NewQWebChannel(QtObjectPtr parent);
void QWebChannel_ConnectBlockUpdatesChanged(QtObjectPtr ptr);
void QWebChannel_DisconnectBlockUpdatesChanged(QtObjectPtr ptr);
void QWebChannel_ConnectTo(QtObjectPtr ptr, QtObjectPtr transport);
void QWebChannel_DeregisterObject(QtObjectPtr ptr, QtObjectPtr object);
void QWebChannel_DisconnectFrom(QtObjectPtr ptr, QtObjectPtr transport);
void QWebChannel_RegisterObject(QtObjectPtr ptr, char* id, QtObjectPtr object);
void QWebChannel_DestroyQWebChannel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif