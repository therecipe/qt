#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraCaptureDestinationControl_CaptureDestination(QtObjectPtr ptr);
void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(QtObjectPtr ptr);
void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(QtObjectPtr ptr);
int QCameraCaptureDestinationControl_IsCaptureDestinationSupported(QtObjectPtr ptr, int destination);
void QCameraCaptureDestinationControl_SetCaptureDestination(QtObjectPtr ptr, int destination);
void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif