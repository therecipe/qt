#ifdef __cplusplus
extern "C" {
#endif

void* QOffscreenSurface_NewQOffscreenSurface(void* targetScreen);
void QOffscreenSurface_Create(void* ptr);
void QOffscreenSurface_Destroy(void* ptr);
int QOffscreenSurface_IsValid(void* ptr);
void* QOffscreenSurface_Screen(void* ptr);
void QOffscreenSurface_ConnectScreenChanged(void* ptr);
void QOffscreenSurface_DisconnectScreenChanged(void* ptr);
void QOffscreenSurface_SetFormat(void* ptr, void* format);
void QOffscreenSurface_SetScreen(void* ptr, void* newScreen);
int QOffscreenSurface_SurfaceType(void* ptr);
void QOffscreenSurface_DestroyQOffscreenSurface(void* ptr);

#ifdef __cplusplus
}
#endif