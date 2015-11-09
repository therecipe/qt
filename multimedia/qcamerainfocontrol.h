#ifdef __cplusplus
extern "C" {
#endif

int QCameraInfoControl_CameraOrientation(void* ptr, char* deviceName);
int QCameraInfoControl_CameraPosition(void* ptr, char* deviceName);
void QCameraInfoControl_DestroyQCameraInfoControl(void* ptr);

#ifdef __cplusplus
}
#endif