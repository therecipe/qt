#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSurfaceFormat_NewQSurfaceFormat();
QtObjectPtr QSurfaceFormat_NewQSurfaceFormat2(int options);
QtObjectPtr QSurfaceFormat_NewQSurfaceFormat3(QtObjectPtr other);
int QSurfaceFormat_AlphaBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_BlueBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_DepthBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_GreenBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_HasAlpha(QtObjectPtr ptr);
int QSurfaceFormat_MajorVersion(QtObjectPtr ptr);
int QSurfaceFormat_MinorVersion(QtObjectPtr ptr);
int QSurfaceFormat_Options(QtObjectPtr ptr);
int QSurfaceFormat_Profile(QtObjectPtr ptr);
int QSurfaceFormat_RedBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_RenderableType(QtObjectPtr ptr);
int QSurfaceFormat_Samples(QtObjectPtr ptr);
void QSurfaceFormat_SetAlphaBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_SetBlueBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(QtObjectPtr format);
void QSurfaceFormat_SetDepthBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_SetGreenBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_SetMajorVersion(QtObjectPtr ptr, int major);
void QSurfaceFormat_SetMinorVersion(QtObjectPtr ptr, int minor);
void QSurfaceFormat_SetOption(QtObjectPtr ptr, int option, int on);
void QSurfaceFormat_SetOptions(QtObjectPtr ptr, int options);
void QSurfaceFormat_SetProfile(QtObjectPtr ptr, int profile);
void QSurfaceFormat_SetRedBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_SetRenderableType(QtObjectPtr ptr, int ty);
void QSurfaceFormat_SetSamples(QtObjectPtr ptr, int numSamples);
void QSurfaceFormat_SetStencilBufferSize(QtObjectPtr ptr, int size);
void QSurfaceFormat_SetStereo(QtObjectPtr ptr, int enable);
void QSurfaceFormat_SetSwapBehavior(QtObjectPtr ptr, int behavior);
void QSurfaceFormat_SetSwapInterval(QtObjectPtr ptr, int interval);
void QSurfaceFormat_SetVersion(QtObjectPtr ptr, int major, int minor);
int QSurfaceFormat_StencilBufferSize(QtObjectPtr ptr);
int QSurfaceFormat_Stereo(QtObjectPtr ptr);
int QSurfaceFormat_SwapBehavior(QtObjectPtr ptr);
int QSurfaceFormat_SwapInterval(QtObjectPtr ptr);
int QSurfaceFormat_TestOption(QtObjectPtr ptr, int option);
void QSurfaceFormat_DestroyQSurfaceFormat(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif