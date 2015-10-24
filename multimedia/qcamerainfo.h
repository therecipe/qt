#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCameraInfo_NewQCameraInfo(QtObjectPtr name);
QtObjectPtr QCameraInfo_NewQCameraInfo2(QtObjectPtr camera);
QtObjectPtr QCameraInfo_NewQCameraInfo3(QtObjectPtr other);
char* QCameraInfo_Description(QtObjectPtr ptr);
char* QCameraInfo_DeviceName(QtObjectPtr ptr);
int QCameraInfo_IsNull(QtObjectPtr ptr);
int QCameraInfo_Orientation(QtObjectPtr ptr);
int QCameraInfo_Position(QtObjectPtr ptr);
void QCameraInfo_DestroyQCameraInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif