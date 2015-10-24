#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QOpenGLFramebufferObject_Bind(QtObjectPtr ptr);
int QOpenGLFramebufferObject_QOpenGLFramebufferObject_BindDefault();
int QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferBlit();
int QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferObjects();
int QOpenGLFramebufferObject_IsValid(QtObjectPtr ptr);
int QOpenGLFramebufferObject_Release(QtObjectPtr ptr);
void QOpenGLFramebufferObject_DestroyQOpenGLFramebufferObject(QtObjectPtr ptr);
QtObjectPtr QOpenGLFramebufferObject_NewQOpenGLFramebufferObject3(QtObjectPtr size, QtObjectPtr format);
QtObjectPtr QOpenGLFramebufferObject_NewQOpenGLFramebufferObject4(int width, int height, QtObjectPtr format);
int QOpenGLFramebufferObject_Attachment(QtObjectPtr ptr);
int QOpenGLFramebufferObject_Height(QtObjectPtr ptr);
int QOpenGLFramebufferObject_IsBound(QtObjectPtr ptr);
void QOpenGLFramebufferObject_SetAttachment(QtObjectPtr ptr, int attachment);
int QOpenGLFramebufferObject_Width(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif