#ifdef __cplusplus
extern "C" {
#endif

int QCameraFlashControl_FlashMode(void* ptr);
void QCameraFlashControl_ConnectFlashReady(void* ptr);
void QCameraFlashControl_DisconnectFlashReady(void* ptr);
int QCameraFlashControl_IsFlashModeSupported(void* ptr, int mode);
int QCameraFlashControl_IsFlashReady(void* ptr);
void QCameraFlashControl_SetFlashMode(void* ptr, int mode);
void QCameraFlashControl_DestroyQCameraFlashControl(void* ptr);

#ifdef __cplusplus
}
#endif