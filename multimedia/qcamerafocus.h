#ifdef __cplusplus
extern "C" {
#endif

double QCameraFocus_DigitalZoom(void* ptr);
int QCameraFocus_FocusMode(void* ptr);
int QCameraFocus_FocusPointMode(void* ptr);
double QCameraFocus_OpticalZoom(void* ptr);
void QCameraFocus_SetCustomFocusPoint(void* ptr, void* point);
void QCameraFocus_SetFocusMode(void* ptr, int mode);
void QCameraFocus_SetFocusPointMode(void* ptr, int mode);
void QCameraFocus_ConnectFocusZonesChanged(void* ptr);
void QCameraFocus_DisconnectFocusZonesChanged(void* ptr);
int QCameraFocus_IsAvailable(void* ptr);
int QCameraFocus_IsFocusModeSupported(void* ptr, int mode);
int QCameraFocus_IsFocusPointModeSupported(void* ptr, int mode);
double QCameraFocus_MaximumDigitalZoom(void* ptr);
double QCameraFocus_MaximumOpticalZoom(void* ptr);
void QCameraFocus_ZoomTo(void* ptr, double optical, double digital);

#ifdef __cplusplus
}
#endif