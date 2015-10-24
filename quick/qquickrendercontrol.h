#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQuickRenderControl_NewQQuickRenderControl(QtObjectPtr parent);
void QQuickRenderControl_Initialize(QtObjectPtr ptr, QtObjectPtr gl);
void QQuickRenderControl_Invalidate(QtObjectPtr ptr);
void QQuickRenderControl_PolishItems(QtObjectPtr ptr);
void QQuickRenderControl_PrepareThread(QtObjectPtr ptr, QtObjectPtr targetThread);
void QQuickRenderControl_Render(QtObjectPtr ptr);
void QQuickRenderControl_ConnectRenderRequested(QtObjectPtr ptr);
void QQuickRenderControl_DisconnectRenderRequested(QtObjectPtr ptr);
QtObjectPtr QQuickRenderControl_RenderWindow(QtObjectPtr ptr, QtObjectPtr offset);
QtObjectPtr QQuickRenderControl_QQuickRenderControl_RenderWindowFor(QtObjectPtr win, QtObjectPtr offset);
void QQuickRenderControl_ConnectSceneChanged(QtObjectPtr ptr);
void QQuickRenderControl_DisconnectSceneChanged(QtObjectPtr ptr);
int QQuickRenderControl_Sync(QtObjectPtr ptr);
void QQuickRenderControl_DestroyQQuickRenderControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif