#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(QtObjectPtr ptr, int parameter);
void QCameraViewfinderSettingsControl_SetViewfinderParameter(QtObjectPtr ptr, int parameter, char* value);
char* QCameraViewfinderSettingsControl_ViewfinderParameter(QtObjectPtr ptr, int parameter);
void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif