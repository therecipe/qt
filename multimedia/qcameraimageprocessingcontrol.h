#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraImageProcessingControl_IsParameterSupported(QtObjectPtr ptr, int parameter);
int QCameraImageProcessingControl_IsParameterValueSupported(QtObjectPtr ptr, int parameter, char* value);
char* QCameraImageProcessingControl_Parameter(QtObjectPtr ptr, int parameter);
void QCameraImageProcessingControl_SetParameter(QtObjectPtr ptr, int parameter, char* value);
void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif