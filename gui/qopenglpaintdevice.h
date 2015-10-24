#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice();
QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice2(QtObjectPtr size);
QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice3(int width, int height);
QtObjectPtr QOpenGLPaintDevice_Context(QtObjectPtr ptr);
void QOpenGLPaintDevice_EnsureActiveTarget(QtObjectPtr ptr);
QtObjectPtr QOpenGLPaintDevice_PaintEngine(QtObjectPtr ptr);
int QOpenGLPaintDevice_PaintFlipped(QtObjectPtr ptr);
void QOpenGLPaintDevice_SetPaintFlipped(QtObjectPtr ptr, int flipped);
void QOpenGLPaintDevice_SetSize(QtObjectPtr ptr, QtObjectPtr size);
void QOpenGLPaintDevice_DestroyQOpenGLPaintDevice(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif