#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSurface_SupportsOpenGL(QtObjectPtr ptr);
int QSurface_SurfaceClass(QtObjectPtr ptr);
int QSurface_SurfaceType(QtObjectPtr ptr);
void QSurface_DestroyQSurface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif