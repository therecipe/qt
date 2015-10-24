#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQuickWindow_ActiveFocusItem(QtObjectPtr ptr);
QtObjectPtr QQuickWindow_ContentItem(QtObjectPtr ptr);
void QQuickWindow_SetColor(QtObjectPtr ptr, QtObjectPtr color);
QtObjectPtr QQuickWindow_NewQQuickWindow(QtObjectPtr parent);
QtObjectPtr QQuickWindow_AccessibleRoot(QtObjectPtr ptr);
void QQuickWindow_ConnectActiveFocusItemChanged(QtObjectPtr ptr);
void QQuickWindow_DisconnectActiveFocusItemChanged(QtObjectPtr ptr);
void QQuickWindow_ConnectAfterAnimating(QtObjectPtr ptr);
void QQuickWindow_DisconnectAfterAnimating(QtObjectPtr ptr);
void QQuickWindow_ConnectAfterRendering(QtObjectPtr ptr);
void QQuickWindow_DisconnectAfterRendering(QtObjectPtr ptr);
void QQuickWindow_ConnectAfterSynchronizing(QtObjectPtr ptr);
void QQuickWindow_DisconnectAfterSynchronizing(QtObjectPtr ptr);
void QQuickWindow_ConnectBeforeRendering(QtObjectPtr ptr);
void QQuickWindow_DisconnectBeforeRendering(QtObjectPtr ptr);
void QQuickWindow_ConnectBeforeSynchronizing(QtObjectPtr ptr);
void QQuickWindow_DisconnectBeforeSynchronizing(QtObjectPtr ptr);
int QQuickWindow_ClearBeforeRendering(QtObjectPtr ptr);
QtObjectPtr QQuickWindow_CreateTextureFromImage2(QtObjectPtr ptr, QtObjectPtr image);
QtObjectPtr QQuickWindow_CreateTextureFromImage(QtObjectPtr ptr, QtObjectPtr image, int options);
void QQuickWindow_ConnectFrameSwapped(QtObjectPtr ptr);
void QQuickWindow_DisconnectFrameSwapped(QtObjectPtr ptr);
int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer();
QtObjectPtr QQuickWindow_IncubationController(QtObjectPtr ptr);
int QQuickWindow_IsPersistentOpenGLContext(QtObjectPtr ptr);
int QQuickWindow_IsPersistentSceneGraph(QtObjectPtr ptr);
int QQuickWindow_IsSceneGraphInitialized(QtObjectPtr ptr);
QtObjectPtr QQuickWindow_MouseGrabberItem(QtObjectPtr ptr);
QtObjectPtr QQuickWindow_OpenglContext(QtObjectPtr ptr);
void QQuickWindow_ConnectOpenglContextCreated(QtObjectPtr ptr);
void QQuickWindow_DisconnectOpenglContextCreated(QtObjectPtr ptr);
void QQuickWindow_ReleaseResources(QtObjectPtr ptr);
QtObjectPtr QQuickWindow_RenderTarget(QtObjectPtr ptr);
void QQuickWindow_ResetOpenGLState(QtObjectPtr ptr);
void QQuickWindow_ConnectSceneGraphAboutToStop(QtObjectPtr ptr);
void QQuickWindow_DisconnectSceneGraphAboutToStop(QtObjectPtr ptr);
void QQuickWindow_ConnectSceneGraphError(QtObjectPtr ptr);
void QQuickWindow_DisconnectSceneGraphError(QtObjectPtr ptr);
void QQuickWindow_ConnectSceneGraphInitialized(QtObjectPtr ptr);
void QQuickWindow_DisconnectSceneGraphInitialized(QtObjectPtr ptr);
void QQuickWindow_ConnectSceneGraphInvalidated(QtObjectPtr ptr);
void QQuickWindow_DisconnectSceneGraphInvalidated(QtObjectPtr ptr);
void QQuickWindow_ScheduleRenderJob(QtObjectPtr ptr, QtObjectPtr job, int stage);
int QQuickWindow_SendEvent(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr e);
void QQuickWindow_SetClearBeforeRendering(QtObjectPtr ptr, int enabled);
void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha);
void QQuickWindow_SetPersistentOpenGLContext(QtObjectPtr ptr, int persistent);
void QQuickWindow_SetPersistentSceneGraph(QtObjectPtr ptr, int persistent);
void QQuickWindow_SetRenderTarget(QtObjectPtr ptr, QtObjectPtr fbo);
void QQuickWindow_Update(QtObjectPtr ptr);
void QQuickWindow_DestroyQQuickWindow(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif