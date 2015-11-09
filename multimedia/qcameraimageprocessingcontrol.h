#ifdef __cplusplus
extern "C" {
#endif

int QCameraImageProcessingControl_IsParameterSupported(void* ptr, int parameter);
int QCameraImageProcessingControl_IsParameterValueSupported(void* ptr, int parameter, void* value);
void* QCameraImageProcessingControl_Parameter(void* ptr, int parameter);
void QCameraImageProcessingControl_SetParameter(void* ptr, int parameter, void* value);
void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(void* ptr);

#ifdef __cplusplus
}
#endif