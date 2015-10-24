#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QCameraImageCaptureControl_CancelCapture(QtObjectPtr ptr);
int QCameraImageCaptureControl_Capture(QtObjectPtr ptr, char* fileName);
int QCameraImageCaptureControl_DriveMode(QtObjectPtr ptr);
void QCameraImageCaptureControl_ConnectError(QtObjectPtr ptr);
void QCameraImageCaptureControl_DisconnectError(QtObjectPtr ptr);
void QCameraImageCaptureControl_ConnectImageExposed(QtObjectPtr ptr);
void QCameraImageCaptureControl_DisconnectImageExposed(QtObjectPtr ptr);
void QCameraImageCaptureControl_ConnectImageMetadataAvailable(QtObjectPtr ptr);
void QCameraImageCaptureControl_DisconnectImageMetadataAvailable(QtObjectPtr ptr);
void QCameraImageCaptureControl_ConnectImageSaved(QtObjectPtr ptr);
void QCameraImageCaptureControl_DisconnectImageSaved(QtObjectPtr ptr);
int QCameraImageCaptureControl_IsReadyForCapture(QtObjectPtr ptr);
void QCameraImageCaptureControl_ConnectReadyForCaptureChanged(QtObjectPtr ptr);
void QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(QtObjectPtr ptr);
void QCameraImageCaptureControl_SetDriveMode(QtObjectPtr ptr, int mode);
void QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif