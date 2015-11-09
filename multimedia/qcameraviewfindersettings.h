#ifdef __cplusplus
extern "C" {
#endif

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings();
void* QCameraViewfinderSettings_NewQCameraViewfinderSettings2(void* other);
int QCameraViewfinderSettings_IsNull(void* ptr);
double QCameraViewfinderSettings_MaximumFrameRate(void* ptr);
double QCameraViewfinderSettings_MinimumFrameRate(void* ptr);
int QCameraViewfinderSettings_PixelFormat(void* ptr);
void QCameraViewfinderSettings_SetMaximumFrameRate(void* ptr, double rate);
void QCameraViewfinderSettings_SetMinimumFrameRate(void* ptr, double rate);
void QCameraViewfinderSettings_SetPixelAspectRatio(void* ptr, void* ratio);
void QCameraViewfinderSettings_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical);
void QCameraViewfinderSettings_SetPixelFormat(void* ptr, int format);
void QCameraViewfinderSettings_SetResolution(void* ptr, void* resolution);
void QCameraViewfinderSettings_SetResolution2(void* ptr, int width, int height);
void QCameraViewfinderSettings_Swap(void* ptr, void* other);
void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(void* ptr);

#ifdef __cplusplus
}
#endif