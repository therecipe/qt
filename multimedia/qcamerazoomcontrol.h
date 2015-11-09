#ifdef __cplusplus
extern "C" {
#endif

double QCameraZoomControl_CurrentDigitalZoom(void* ptr);
double QCameraZoomControl_CurrentOpticalZoom(void* ptr);
double QCameraZoomControl_MaximumDigitalZoom(void* ptr);
double QCameraZoomControl_MaximumOpticalZoom(void* ptr);
double QCameraZoomControl_RequestedDigitalZoom(void* ptr);
double QCameraZoomControl_RequestedOpticalZoom(void* ptr);
void QCameraZoomControl_ZoomTo(void* ptr, double optical, double digital);
void QCameraZoomControl_DestroyQCameraZoomControl(void* ptr);

#ifdef __cplusplus
}
#endif