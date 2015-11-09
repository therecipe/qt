#ifdef __cplusplus
extern "C" {
#endif

void QAbstractVideoSurface_ConnectActiveChanged(void* ptr);
void QAbstractVideoSurface_DisconnectActiveChanged(void* ptr);
int QAbstractVideoSurface_Error(void* ptr);
int QAbstractVideoSurface_IsActive(void* ptr);
int QAbstractVideoSurface_IsFormatSupported(void* ptr, void* format);
int QAbstractVideoSurface_Present(void* ptr, void* frame);
int QAbstractVideoSurface_Start(void* ptr, void* format);
void QAbstractVideoSurface_Stop(void* ptr);
void QAbstractVideoSurface_ConnectSupportedFormatsChanged(void* ptr);
void QAbstractVideoSurface_DisconnectSupportedFormatsChanged(void* ptr);
void QAbstractVideoSurface_DestroyQAbstractVideoSurface(void* ptr);

#ifdef __cplusplus
}
#endif