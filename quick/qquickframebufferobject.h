#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QQuickFramebufferObject_SetTextureFollowsItemSize(QtObjectPtr ptr, int follows);
int QQuickFramebufferObject_TextureFollowsItemSize(QtObjectPtr ptr);
int QQuickFramebufferObject_IsTextureProvider(QtObjectPtr ptr);
void QQuickFramebufferObject_ReleaseResources(QtObjectPtr ptr);
void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(QtObjectPtr ptr);
void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(QtObjectPtr ptr);
QtObjectPtr QQuickFramebufferObject_TextureProvider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif