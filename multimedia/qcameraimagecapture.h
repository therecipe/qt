#ifdef __cplusplus
extern "C" {
#endif

int QCameraImageCapture_IsReadyForCapture(void* ptr);
void* QCameraImageCapture_NewQCameraImageCapture(void* mediaObject, void* parent);
int QCameraImageCapture_BufferFormat(void* ptr);
void QCameraImageCapture_ConnectBufferFormatChanged(void* ptr);
void QCameraImageCapture_DisconnectBufferFormatChanged(void* ptr);
void QCameraImageCapture_CancelCapture(void* ptr);
int QCameraImageCapture_Capture(void* ptr, char* file);
int QCameraImageCapture_CaptureDestination(void* ptr);
void QCameraImageCapture_ConnectCaptureDestinationChanged(void* ptr);
void QCameraImageCapture_DisconnectCaptureDestinationChanged(void* ptr);
int QCameraImageCapture_Error(void* ptr);
char* QCameraImageCapture_ErrorString(void* ptr);
char* QCameraImageCapture_ImageCodecDescription(void* ptr, char* codec);
void QCameraImageCapture_ConnectImageExposed(void* ptr);
void QCameraImageCapture_DisconnectImageExposed(void* ptr);
void QCameraImageCapture_ConnectImageMetadataAvailable(void* ptr);
void QCameraImageCapture_DisconnectImageMetadataAvailable(void* ptr);
void QCameraImageCapture_ConnectImageSaved(void* ptr);
void QCameraImageCapture_DisconnectImageSaved(void* ptr);
int QCameraImageCapture_IsAvailable(void* ptr);
int QCameraImageCapture_IsCaptureDestinationSupported(void* ptr, int destination);
void* QCameraImageCapture_MediaObject(void* ptr);
void QCameraImageCapture_ConnectReadyForCaptureChanged(void* ptr);
void QCameraImageCapture_DisconnectReadyForCaptureChanged(void* ptr);
void QCameraImageCapture_SetBufferFormat(void* ptr, int format);
void QCameraImageCapture_SetCaptureDestination(void* ptr, int destination);
void QCameraImageCapture_SetEncodingSettings(void* ptr, void* settings);
char* QCameraImageCapture_SupportedImageCodecs(void* ptr);
void QCameraImageCapture_DestroyQCameraImageCapture(void* ptr);

#ifdef __cplusplus
}
#endif