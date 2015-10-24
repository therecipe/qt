#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLWidget_NewQOpenGLWidget(QtObjectPtr parent, int f);
void QOpenGLWidget_ConnectAboutToCompose(QtObjectPtr ptr);
void QOpenGLWidget_DisconnectAboutToCompose(QtObjectPtr ptr);
void QOpenGLWidget_ConnectAboutToResize(QtObjectPtr ptr);
void QOpenGLWidget_DisconnectAboutToResize(QtObjectPtr ptr);
QtObjectPtr QOpenGLWidget_Context(QtObjectPtr ptr);
void QOpenGLWidget_DoneCurrent(QtObjectPtr ptr);
int QOpenGLWidget_IsValid(QtObjectPtr ptr);
void QOpenGLWidget_MakeCurrent(QtObjectPtr ptr);
void QOpenGLWidget_ConnectFrameSwapped(QtObjectPtr ptr);
void QOpenGLWidget_DisconnectFrameSwapped(QtObjectPtr ptr);
void QOpenGLWidget_ConnectResized(QtObjectPtr ptr);
void QOpenGLWidget_DisconnectResized(QtObjectPtr ptr);
void QOpenGLWidget_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QOpenGLWidget_SetUpdateBehavior(QtObjectPtr ptr, int updateBehavior);
int QOpenGLWidget_UpdateBehavior(QtObjectPtr ptr);
void QOpenGLWidget_DestroyQOpenGLWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif