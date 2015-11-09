#ifdef __cplusplus
extern "C" {
#endif

int QQuickView_ResizeMode(void* ptr);
void QQuickView_SetResizeMode(void* ptr, int v);
int QQuickView_Status(void* ptr);
void* QQuickView_NewQQuickView2(void* engine, void* parent);
void* QQuickView_NewQQuickView(void* parent);
void* QQuickView_NewQQuickView3(void* source, void* parent);
void* QQuickView_Engine(void* ptr);
void* QQuickView_RootContext(void* ptr);
void* QQuickView_RootObject(void* ptr);
void QQuickView_SetSource(void* ptr, void* url);
void QQuickView_ConnectStatusChanged(void* ptr);
void QQuickView_DisconnectStatusChanged(void* ptr);
void QQuickView_DestroyQQuickView(void* ptr);

#ifdef __cplusplus
}
#endif