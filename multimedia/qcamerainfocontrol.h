#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraInfoControl_CameraOrientation(QtObjectPtr ptr, char* deviceName);
int QCameraInfoControl_CameraPosition(QtObjectPtr ptr, char* deviceName);
void QCameraInfoControl_DestroyQCameraInfoControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif