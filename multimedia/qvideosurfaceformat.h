#ifdef __cplusplus
extern "C" {
#endif

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat();
void* QVideoSurfaceFormat_NewQVideoSurfaceFormat2(void* size, int format, int ty);
void* QVideoSurfaceFormat_NewQVideoSurfaceFormat3(void* other);
int QVideoSurfaceFormat_FrameHeight(void* ptr);
double QVideoSurfaceFormat_FrameRate(void* ptr);
int QVideoSurfaceFormat_FrameWidth(void* ptr);
int QVideoSurfaceFormat_HandleType(void* ptr);
int QVideoSurfaceFormat_IsValid(void* ptr);
int QVideoSurfaceFormat_PixelFormat(void* ptr);
void* QVideoSurfaceFormat_Property(void* ptr, char* name);
int QVideoSurfaceFormat_ScanLineDirection(void* ptr);
void QVideoSurfaceFormat_SetFrameRate(void* ptr, double rate);
void QVideoSurfaceFormat_SetFrameSize(void* ptr, void* size);
void QVideoSurfaceFormat_SetFrameSize2(void* ptr, int width, int height);
void QVideoSurfaceFormat_SetPixelAspectRatio(void* ptr, void* ratio);
void QVideoSurfaceFormat_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical);
void QVideoSurfaceFormat_SetProperty(void* ptr, char* name, void* value);
void QVideoSurfaceFormat_SetScanLineDirection(void* ptr, int direction);
void QVideoSurfaceFormat_SetViewport(void* ptr, void* viewport);
void QVideoSurfaceFormat_SetYCbCrColorSpace(void* ptr, int space);
int QVideoSurfaceFormat_YCbCrColorSpace(void* ptr);
void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(void* ptr);

#ifdef __cplusplus
}
#endif