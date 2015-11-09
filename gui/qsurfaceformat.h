#ifdef __cplusplus
extern "C" {
#endif

void* QSurfaceFormat_NewQSurfaceFormat();
void* QSurfaceFormat_NewQSurfaceFormat2(int options);
void* QSurfaceFormat_NewQSurfaceFormat3(void* other);
int QSurfaceFormat_AlphaBufferSize(void* ptr);
int QSurfaceFormat_BlueBufferSize(void* ptr);
int QSurfaceFormat_DepthBufferSize(void* ptr);
int QSurfaceFormat_GreenBufferSize(void* ptr);
int QSurfaceFormat_HasAlpha(void* ptr);
int QSurfaceFormat_MajorVersion(void* ptr);
int QSurfaceFormat_MinorVersion(void* ptr);
int QSurfaceFormat_Options(void* ptr);
int QSurfaceFormat_Profile(void* ptr);
int QSurfaceFormat_RedBufferSize(void* ptr);
int QSurfaceFormat_RenderableType(void* ptr);
int QSurfaceFormat_Samples(void* ptr);
void QSurfaceFormat_SetAlphaBufferSize(void* ptr, int size);
void QSurfaceFormat_SetBlueBufferSize(void* ptr, int size);
void QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(void* format);
void QSurfaceFormat_SetDepthBufferSize(void* ptr, int size);
void QSurfaceFormat_SetGreenBufferSize(void* ptr, int size);
void QSurfaceFormat_SetMajorVersion(void* ptr, int major);
void QSurfaceFormat_SetMinorVersion(void* ptr, int minor);
void QSurfaceFormat_SetOption(void* ptr, int option, int on);
void QSurfaceFormat_SetOptions(void* ptr, int options);
void QSurfaceFormat_SetProfile(void* ptr, int profile);
void QSurfaceFormat_SetRedBufferSize(void* ptr, int size);
void QSurfaceFormat_SetRenderableType(void* ptr, int ty);
void QSurfaceFormat_SetSamples(void* ptr, int numSamples);
void QSurfaceFormat_SetStencilBufferSize(void* ptr, int size);
void QSurfaceFormat_SetStereo(void* ptr, int enable);
void QSurfaceFormat_SetSwapBehavior(void* ptr, int behavior);
void QSurfaceFormat_SetSwapInterval(void* ptr, int interval);
void QSurfaceFormat_SetVersion(void* ptr, int major, int minor);
int QSurfaceFormat_StencilBufferSize(void* ptr);
int QSurfaceFormat_Stereo(void* ptr);
int QSurfaceFormat_SwapBehavior(void* ptr);
int QSurfaceFormat_SwapInterval(void* ptr);
int QSurfaceFormat_TestOption(void* ptr, int option);
void QSurfaceFormat_DestroyQSurfaceFormat(void* ptr);

#ifdef __cplusplus
}
#endif