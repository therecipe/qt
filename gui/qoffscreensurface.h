#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOffscreenSurface_NewQOffscreenSurface(QtObjectPtr targetScreen);
void QOffscreenSurface_Create(QtObjectPtr ptr);
void QOffscreenSurface_Destroy(QtObjectPtr ptr);
int QOffscreenSurface_IsValid(QtObjectPtr ptr);
QtObjectPtr QOffscreenSurface_Screen(QtObjectPtr ptr);
void QOffscreenSurface_ConnectScreenChanged(QtObjectPtr ptr);
void QOffscreenSurface_DisconnectScreenChanged(QtObjectPtr ptr);
void QOffscreenSurface_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QOffscreenSurface_SetScreen(QtObjectPtr ptr, QtObjectPtr newScreen);
int QOffscreenSurface_SurfaceType(QtObjectPtr ptr);
void QOffscreenSurface_DestroyQOffscreenSurface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif