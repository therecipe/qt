#ifdef __cplusplus
extern "C" {
#endif

void QCameraImageCaptureControl_CancelCapture(void* ptr);
int QCameraImageCaptureControl_Capture(void* ptr, char* fileName);
int QCameraImageCaptureControl_DriveMode(void* ptr);
void QCameraImageCaptureControl_ConnectError(void* ptr);
void QCameraImageCaptureControl_DisconnectError(void* ptr);
void QCameraImageCaptureControl_ConnectImageExposed(void* ptr);
void QCameraImageCaptureControl_DisconnectImageExposed(void* ptr);
void QCameraImageCaptureControl_ConnectImageMetadataAvailable(void* ptr);
void QCameraImageCaptureControl_DisconnectImageMetadataAvailable(void* ptr);
void QCameraImageCaptureControl_ConnectImageSaved(void* ptr);
void QCameraImageCaptureControl_DisconnectImageSaved(void* ptr);
int QCameraImageCaptureControl_IsReadyForCapture(void* ptr);
void QCameraImageCaptureControl_ConnectReadyForCaptureChanged(void* ptr);
void QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(void* ptr);
void QCameraImageCaptureControl_SetDriveMode(void* ptr, int mode);
void QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(void* ptr);

#ifdef __cplusplus
}
#endif