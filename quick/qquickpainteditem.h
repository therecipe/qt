#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QQuickPaintedItem_RenderTarget(QtObjectPtr ptr);
void QQuickPaintedItem_SetContentsSize(QtObjectPtr ptr, QtObjectPtr v);
void QQuickPaintedItem_SetFillColor(QtObjectPtr ptr, QtObjectPtr v);
void QQuickPaintedItem_SetRenderTarget(QtObjectPtr ptr, int target);
int QQuickPaintedItem_Antialiasing(QtObjectPtr ptr);
void QQuickPaintedItem_ConnectContentsScaleChanged(QtObjectPtr ptr);
void QQuickPaintedItem_DisconnectContentsScaleChanged(QtObjectPtr ptr);
void QQuickPaintedItem_ConnectContentsSizeChanged(QtObjectPtr ptr);
void QQuickPaintedItem_DisconnectContentsSizeChanged(QtObjectPtr ptr);
void QQuickPaintedItem_ConnectFillColorChanged(QtObjectPtr ptr);
void QQuickPaintedItem_DisconnectFillColorChanged(QtObjectPtr ptr);
int QQuickPaintedItem_IsTextureProvider(QtObjectPtr ptr);
int QQuickPaintedItem_Mipmap(QtObjectPtr ptr);
int QQuickPaintedItem_OpaquePainting(QtObjectPtr ptr);
void QQuickPaintedItem_Paint(QtObjectPtr ptr, QtObjectPtr painter);
int QQuickPaintedItem_PerformanceHints(QtObjectPtr ptr);
void QQuickPaintedItem_ConnectRenderTargetChanged(QtObjectPtr ptr);
void QQuickPaintedItem_DisconnectRenderTargetChanged(QtObjectPtr ptr);
void QQuickPaintedItem_ResetContentsSize(QtObjectPtr ptr);
void QQuickPaintedItem_SetAntialiasing(QtObjectPtr ptr, int enable);
void QQuickPaintedItem_SetMipmap(QtObjectPtr ptr, int enable);
void QQuickPaintedItem_SetOpaquePainting(QtObjectPtr ptr, int opaque);
void QQuickPaintedItem_SetPerformanceHint(QtObjectPtr ptr, int hint, int enabled);
void QQuickPaintedItem_SetPerformanceHints(QtObjectPtr ptr, int hints);
QtObjectPtr QQuickPaintedItem_TextureProvider(QtObjectPtr ptr);
void QQuickPaintedItem_Update(QtObjectPtr ptr, QtObjectPtr rect);
void QQuickPaintedItem_DestroyQQuickPaintedItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif