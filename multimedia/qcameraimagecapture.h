#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraImageCapture_IsReadyForCapture(QtObjectPtr ptr);
QtObjectPtr QCameraImageCapture_NewQCameraImageCapture(QtObjectPtr mediaObject, QtObjectPtr parent);
int QCameraImageCapture_BufferFormat(QtObjectPtr ptr);
void QCameraImageCapture_ConnectBufferFormatChanged(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectBufferFormatChanged(QtObjectPtr ptr);
void QCameraImageCapture_CancelCapture(QtObjectPtr ptr);
int QCameraImageCapture_Capture(QtObjectPtr ptr, char* file);
int QCameraImageCapture_CaptureDestination(QtObjectPtr ptr);
void QCameraImageCapture_ConnectCaptureDestinationChanged(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectCaptureDestinationChanged(QtObjectPtr ptr);
int QCameraImageCapture_Error(QtObjectPtr ptr);
char* QCameraImageCapture_ErrorString(QtObjectPtr ptr);
char* QCameraImageCapture_ImageCodecDescription(QtObjectPtr ptr, char* codec);
void QCameraImageCapture_ConnectImageExposed(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectImageExposed(QtObjectPtr ptr);
void QCameraImageCapture_ConnectImageMetadataAvailable(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectImageMetadataAvailable(QtObjectPtr ptr);
void QCameraImageCapture_ConnectImageSaved(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectImageSaved(QtObjectPtr ptr);
int QCameraImageCapture_IsAvailable(QtObjectPtr ptr);
int QCameraImageCapture_IsCaptureDestinationSupported(QtObjectPtr ptr, int destination);
QtObjectPtr QCameraImageCapture_MediaObject(QtObjectPtr ptr);
void QCameraImageCapture_ConnectReadyForCaptureChanged(QtObjectPtr ptr);
void QCameraImageCapture_DisconnectReadyForCaptureChanged(QtObjectPtr ptr);
void QCameraImageCapture_SetBufferFormat(QtObjectPtr ptr, int format);
void QCameraImageCapture_SetCaptureDestination(QtObjectPtr ptr, int destination);
void QCameraImageCapture_SetEncodingSettings(QtObjectPtr ptr, QtObjectPtr settings);
char* QCameraImageCapture_SupportedImageCodecs(QtObjectPtr ptr);
void QCameraImageCapture_DestroyQCameraImageCapture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif