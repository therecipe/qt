#ifdef __cplusplus
extern "C" {
#endif

double QQuickPaintedItem_ContentsScale(void* ptr);
void* QQuickPaintedItem_FillColor(void* ptr);
int QQuickPaintedItem_RenderTarget(void* ptr);
void QQuickPaintedItem_SetContentsScale(void* ptr, double v);
void QQuickPaintedItem_SetContentsSize(void* ptr, void* v);
void QQuickPaintedItem_SetFillColor(void* ptr, void* v);
void QQuickPaintedItem_SetRenderTarget(void* ptr, int target);
int QQuickPaintedItem_Antialiasing(void* ptr);
void QQuickPaintedItem_ConnectContentsScaleChanged(void* ptr);
void QQuickPaintedItem_DisconnectContentsScaleChanged(void* ptr);
void QQuickPaintedItem_ConnectContentsSizeChanged(void* ptr);
void QQuickPaintedItem_DisconnectContentsSizeChanged(void* ptr);
void QQuickPaintedItem_ConnectFillColorChanged(void* ptr);
void QQuickPaintedItem_DisconnectFillColorChanged(void* ptr);
int QQuickPaintedItem_IsTextureProvider(void* ptr);
int QQuickPaintedItem_Mipmap(void* ptr);
int QQuickPaintedItem_OpaquePainting(void* ptr);
void QQuickPaintedItem_Paint(void* ptr, void* painter);
int QQuickPaintedItem_PerformanceHints(void* ptr);
void QQuickPaintedItem_ConnectRenderTargetChanged(void* ptr);
void QQuickPaintedItem_DisconnectRenderTargetChanged(void* ptr);
void QQuickPaintedItem_ResetContentsSize(void* ptr);
void QQuickPaintedItem_SetAntialiasing(void* ptr, int enable);
void QQuickPaintedItem_SetMipmap(void* ptr, int enable);
void QQuickPaintedItem_SetOpaquePainting(void* ptr, int opaque);
void QQuickPaintedItem_SetPerformanceHint(void* ptr, int hint, int enabled);
void QQuickPaintedItem_SetPerformanceHints(void* ptr, int hints);
void* QQuickPaintedItem_TextureProvider(void* ptr);
void QQuickPaintedItem_Update(void* ptr, void* rect);
void QQuickPaintedItem_DestroyQQuickPaintedItem(void* ptr);

#ifdef __cplusplus
}
#endif