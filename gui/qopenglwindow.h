#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLWindow_NewQOpenGLWindow2(QtObjectPtr shareContext, int updateBehavior, QtObjectPtr parent);
QtObjectPtr QOpenGLWindow_NewQOpenGLWindow(int updateBehavior, QtObjectPtr parent);
QtObjectPtr QOpenGLWindow_Context(QtObjectPtr ptr);
void QOpenGLWindow_DoneCurrent(QtObjectPtr ptr);
void QOpenGLWindow_ConnectFrameSwapped(QtObjectPtr ptr);
void QOpenGLWindow_DisconnectFrameSwapped(QtObjectPtr ptr);
int QOpenGLWindow_IsValid(QtObjectPtr ptr);
void QOpenGLWindow_MakeCurrent(QtObjectPtr ptr);
QtObjectPtr QOpenGLWindow_ShareContext(QtObjectPtr ptr);
int QOpenGLWindow_UpdateBehavior(QtObjectPtr ptr);
void QOpenGLWindow_DestroyQOpenGLWindow(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif