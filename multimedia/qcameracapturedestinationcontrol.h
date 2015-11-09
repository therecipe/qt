#ifdef __cplusplus
extern "C" {
#endif

int QCameraCaptureDestinationControl_CaptureDestination(void* ptr);
void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(void* ptr);
void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(void* ptr);
int QCameraCaptureDestinationControl_IsCaptureDestinationSupported(void* ptr, int destination);
void QCameraCaptureDestinationControl_SetCaptureDestination(void* ptr, int destination);
void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(void* ptr);

#ifdef __cplusplus
}
#endif