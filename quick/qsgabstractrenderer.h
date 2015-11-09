#ifdef __cplusplus
extern "C" {
#endif

void* QSGAbstractRenderer_ClearColor(void* ptr);
int QSGAbstractRenderer_ClearMode(void* ptr);
void QSGAbstractRenderer_ConnectSceneGraphChanged(void* ptr);
void QSGAbstractRenderer_DisconnectSceneGraphChanged(void* ptr);
void QSGAbstractRenderer_SetClearColor(void* ptr, void* color);
void QSGAbstractRenderer_SetClearMode(void* ptr, int mode);
void QSGAbstractRenderer_SetDeviceRect(void* ptr, void* rect);
void QSGAbstractRenderer_SetDeviceRect2(void* ptr, void* size);
void QSGAbstractRenderer_SetProjectionMatrix(void* ptr, void* matrix);
void QSGAbstractRenderer_SetProjectionMatrixToRect(void* ptr, void* rect);
void QSGAbstractRenderer_SetViewportRect(void* ptr, void* rect);
void QSGAbstractRenderer_SetViewportRect2(void* ptr, void* size);

#ifdef __cplusplus
}
#endif