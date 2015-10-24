#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat();
QtObjectPtr QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat2(QtObjectPtr other);
int QOpenGLFramebufferObjectFormat_Attachment(QtObjectPtr ptr);
int QOpenGLFramebufferObjectFormat_Mipmap(QtObjectPtr ptr);
int QOpenGLFramebufferObjectFormat_Samples(QtObjectPtr ptr);
void QOpenGLFramebufferObjectFormat_SetAttachment(QtObjectPtr ptr, int attachment);
void QOpenGLFramebufferObjectFormat_SetMipmap(QtObjectPtr ptr, int enabled);
void QOpenGLFramebufferObjectFormat_SetSamples(QtObjectPtr ptr, int samples);
void QOpenGLFramebufferObjectFormat_DestroyQOpenGLFramebufferObjectFormat(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif