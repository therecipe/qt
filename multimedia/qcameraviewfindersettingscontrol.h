#ifdef __cplusplus
extern "C" {
#endif

int QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(void* ptr, int parameter);
void QCameraViewfinderSettingsControl_SetViewfinderParameter(void* ptr, int parameter, void* value);
void* QCameraViewfinderSettingsControl_ViewfinderParameter(void* ptr, int parameter);
void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(void* ptr);

#ifdef __cplusplus
}
#endif