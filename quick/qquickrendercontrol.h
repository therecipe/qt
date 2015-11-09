#ifdef __cplusplus
extern "C" {
#endif

void* QQuickRenderControl_NewQQuickRenderControl(void* parent);
void QQuickRenderControl_Initialize(void* ptr, void* gl);
void QQuickRenderControl_Invalidate(void* ptr);
void QQuickRenderControl_PolishItems(void* ptr);
void QQuickRenderControl_PrepareThread(void* ptr, void* targetThread);
void QQuickRenderControl_Render(void* ptr);
void QQuickRenderControl_ConnectRenderRequested(void* ptr);
void QQuickRenderControl_DisconnectRenderRequested(void* ptr);
void* QQuickRenderControl_RenderWindow(void* ptr, void* offset);
void* QQuickRenderControl_QQuickRenderControl_RenderWindowFor(void* win, void* offset);
void QQuickRenderControl_ConnectSceneChanged(void* ptr);
void QQuickRenderControl_DisconnectSceneChanged(void* ptr);
int QQuickRenderControl_Sync(void* ptr);
void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr);

#ifdef __cplusplus
}
#endif