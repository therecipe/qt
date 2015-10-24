#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat();
QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat2(QtObjectPtr size, int format, int ty);
QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat3(QtObjectPtr other);
int QVideoSurfaceFormat_FrameHeight(QtObjectPtr ptr);
int QVideoSurfaceFormat_FrameWidth(QtObjectPtr ptr);
int QVideoSurfaceFormat_HandleType(QtObjectPtr ptr);
int QVideoSurfaceFormat_IsValid(QtObjectPtr ptr);
int QVideoSurfaceFormat_PixelFormat(QtObjectPtr ptr);
char* QVideoSurfaceFormat_Property(QtObjectPtr ptr, char* name);
int QVideoSurfaceFormat_ScanLineDirection(QtObjectPtr ptr);
void QVideoSurfaceFormat_SetFrameSize(QtObjectPtr ptr, QtObjectPtr size);
void QVideoSurfaceFormat_SetFrameSize2(QtObjectPtr ptr, int width, int height);
void QVideoSurfaceFormat_SetPixelAspectRatio(QtObjectPtr ptr, QtObjectPtr ratio);
void QVideoSurfaceFormat_SetPixelAspectRatio2(QtObjectPtr ptr, int horizontal, int vertical);
void QVideoSurfaceFormat_SetProperty(QtObjectPtr ptr, char* name, char* value);
void QVideoSurfaceFormat_SetScanLineDirection(QtObjectPtr ptr, int direction);
void QVideoSurfaceFormat_SetViewport(QtObjectPtr ptr, QtObjectPtr viewport);
void QVideoSurfaceFormat_SetYCbCrColorSpace(QtObjectPtr ptr, int space);
int QVideoSurfaceFormat_YCbCrColorSpace(QtObjectPtr ptr);
void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif