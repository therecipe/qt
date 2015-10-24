#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCameraViewfinderSettings_NewQCameraViewfinderSettings();
QtObjectPtr QCameraViewfinderSettings_NewQCameraViewfinderSettings2(QtObjectPtr other);
int QCameraViewfinderSettings_IsNull(QtObjectPtr ptr);
int QCameraViewfinderSettings_PixelFormat(QtObjectPtr ptr);
void QCameraViewfinderSettings_SetPixelAspectRatio(QtObjectPtr ptr, QtObjectPtr ratio);
void QCameraViewfinderSettings_SetPixelAspectRatio2(QtObjectPtr ptr, int horizontal, int vertical);
void QCameraViewfinderSettings_SetPixelFormat(QtObjectPtr ptr, int format);
void QCameraViewfinderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution);
void QCameraViewfinderSettings_SetResolution2(QtObjectPtr ptr, int width, int height);
void QCameraViewfinderSettings_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif