#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraFlashControl_FlashMode(QtObjectPtr ptr);
void QCameraFlashControl_ConnectFlashReady(QtObjectPtr ptr);
void QCameraFlashControl_DisconnectFlashReady(QtObjectPtr ptr);
int QCameraFlashControl_IsFlashModeSupported(QtObjectPtr ptr, int mode);
int QCameraFlashControl_IsFlashReady(QtObjectPtr ptr);
void QCameraFlashControl_SetFlashMode(QtObjectPtr ptr, int mode);
void QCameraFlashControl_DestroyQCameraFlashControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif