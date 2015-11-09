#ifdef __cplusplus
extern "C" {
#endif

void* QQuickWindow_ActiveFocusItem(void* ptr);
void* QQuickWindow_Color(void* ptr);
void* QQuickWindow_ContentItem(void* ptr);
void QQuickWindow_SetColor(void* ptr, void* color);
void* QQuickWindow_NewQQuickWindow(void* parent);
void* QQuickWindow_AccessibleRoot(void* ptr);
void QQuickWindow_ConnectActiveFocusItemChanged(void* ptr);
void QQuickWindow_DisconnectActiveFocusItemChanged(void* ptr);
void QQuickWindow_ConnectAfterAnimating(void* ptr);
void QQuickWindow_DisconnectAfterAnimating(void* ptr);
void QQuickWindow_ConnectAfterRendering(void* ptr);
void QQuickWindow_DisconnectAfterRendering(void* ptr);
void QQuickWindow_ConnectAfterSynchronizing(void* ptr);
void QQuickWindow_DisconnectAfterSynchronizing(void* ptr);
void QQuickWindow_ConnectBeforeRendering(void* ptr);
void QQuickWindow_DisconnectBeforeRendering(void* ptr);
void QQuickWindow_ConnectBeforeSynchronizing(void* ptr);
void QQuickWindow_DisconnectBeforeSynchronizing(void* ptr);
int QQuickWindow_ClearBeforeRendering(void* ptr);
void QQuickWindow_ConnectColorChanged(void* ptr);
void QQuickWindow_DisconnectColorChanged(void* ptr);
void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image);
void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, int options);
double QQuickWindow_EffectiveDevicePixelRatio(void* ptr);
void QQuickWindow_ConnectFrameSwapped(void* ptr);
void QQuickWindow_DisconnectFrameSwapped(void* ptr);
int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer();
void* QQuickWindow_IncubationController(void* ptr);
int QQuickWindow_IsPersistentOpenGLContext(void* ptr);
int QQuickWindow_IsPersistentSceneGraph(void* ptr);
int QQuickWindow_IsSceneGraphInitialized(void* ptr);
void* QQuickWindow_MouseGrabberItem(void* ptr);
void* QQuickWindow_OpenglContext(void* ptr);
void QQuickWindow_ConnectOpenglContextCreated(void* ptr);
void QQuickWindow_DisconnectOpenglContextCreated(void* ptr);
void QQuickWindow_ReleaseResources(void* ptr);
void* QQuickWindow_RenderTarget(void* ptr);
void QQuickWindow_ResetOpenGLState(void* ptr);
void QQuickWindow_ConnectSceneGraphAboutToStop(void* ptr);
void QQuickWindow_DisconnectSceneGraphAboutToStop(void* ptr);
void QQuickWindow_ConnectSceneGraphError(void* ptr);
void QQuickWindow_DisconnectSceneGraphError(void* ptr);
void QQuickWindow_ConnectSceneGraphInitialized(void* ptr);
void QQuickWindow_DisconnectSceneGraphInitialized(void* ptr);
void QQuickWindow_ConnectSceneGraphInvalidated(void* ptr);
void QQuickWindow_DisconnectSceneGraphInvalidated(void* ptr);
void QQuickWindow_ScheduleRenderJob(void* ptr, void* job, int stage);
int QQuickWindow_SendEvent(void* ptr, void* item, void* e);
void QQuickWindow_SetClearBeforeRendering(void* ptr, int enabled);
void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha);
void QQuickWindow_SetPersistentOpenGLContext(void* ptr, int persistent);
void QQuickWindow_SetPersistentSceneGraph(void* ptr, int persistent);
void QQuickWindow_SetRenderTarget(void* ptr, void* fbo);
void QQuickWindow_Update(void* ptr);
void QQuickWindow_DestroyQQuickWindow(void* ptr);

#ifdef __cplusplus
}
#endif