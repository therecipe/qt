#ifdef __cplusplus
extern "C" {
#endif

void* QCameraInfo_NewQCameraInfo(void* name);
void* QCameraInfo_NewQCameraInfo2(void* camera);
void* QCameraInfo_NewQCameraInfo3(void* other);
char* QCameraInfo_Description(void* ptr);
char* QCameraInfo_DeviceName(void* ptr);
int QCameraInfo_IsNull(void* ptr);
int QCameraInfo_Orientation(void* ptr);
int QCameraInfo_Position(void* ptr);
void QCameraInfo_DestroyQCameraInfo(void* ptr);

#ifdef __cplusplus
}
#endif