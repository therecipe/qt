#ifdef __cplusplus
extern "C" {
#endif

int QQuickWidget_ResizeMode(void* ptr);
void QQuickWidget_SetResizeMode(void* ptr, int v);
int QQuickWidget_Status(void* ptr);
void* QQuickWidget_NewQQuickWidget2(void* engine, void* parent);
void* QQuickWidget_NewQQuickWidget(void* parent);
void* QQuickWidget_NewQQuickWidget3(void* source, void* parent);
void* QQuickWidget_Engine(void* ptr);
void* QQuickWidget_QuickWindow(void* ptr);
void* QQuickWidget_RootContext(void* ptr);
void* QQuickWidget_RootObject(void* ptr);
void QQuickWidget_ConnectSceneGraphError(void* ptr);
void QQuickWidget_DisconnectSceneGraphError(void* ptr);
void QQuickWidget_SetClearColor(void* ptr, void* color);
void QQuickWidget_SetFormat(void* ptr, void* format);
void QQuickWidget_SetSource(void* ptr, void* url);
void QQuickWidget_ConnectStatusChanged(void* ptr);
void QQuickWidget_DisconnectStatusChanged(void* ptr);
void QQuickWidget_DestroyQQuickWidget(void* ptr);

#ifdef __cplusplus
}
#endif