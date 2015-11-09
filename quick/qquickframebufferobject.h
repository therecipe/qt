#ifdef __cplusplus
extern "C" {
#endif

void QQuickFramebufferObject_SetTextureFollowsItemSize(void* ptr, int follows);
int QQuickFramebufferObject_TextureFollowsItemSize(void* ptr);
int QQuickFramebufferObject_IsTextureProvider(void* ptr);
void QQuickFramebufferObject_ReleaseResources(void* ptr);
void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr);
void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr);
void* QQuickFramebufferObject_TextureProvider(void* ptr);

#ifdef __cplusplus
}
#endif