#ifdef __cplusplus
extern "C" {
#endif

int QSurface_SupportsOpenGL(void* ptr);
int QSurface_SurfaceClass(void* ptr);
int QSurface_SurfaceType(void* ptr);
void QSurface_DestroyQSurface(void* ptr);

#ifdef __cplusplus
}
#endif