#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QQuickView_ResizeMode(QtObjectPtr ptr);
void QQuickView_SetResizeMode(QtObjectPtr ptr, int v);
int QQuickView_Status(QtObjectPtr ptr);
QtObjectPtr QQuickView_NewQQuickView2(QtObjectPtr engine, QtObjectPtr parent);
QtObjectPtr QQuickView_NewQQuickView(QtObjectPtr parent);
QtObjectPtr QQuickView_NewQQuickView3(char* source, QtObjectPtr parent);
QtObjectPtr QQuickView_Engine(QtObjectPtr ptr);
QtObjectPtr QQuickView_RootContext(QtObjectPtr ptr);
QtObjectPtr QQuickView_RootObject(QtObjectPtr ptr);
void QQuickView_SetSource(QtObjectPtr ptr, char* url);
char* QQuickView_Source(QtObjectPtr ptr);
void QQuickView_ConnectStatusChanged(QtObjectPtr ptr);
void QQuickView_DisconnectStatusChanged(QtObjectPtr ptr);
void QQuickView_DestroyQQuickView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif