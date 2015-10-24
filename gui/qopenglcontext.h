#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLContext_NewQOpenGLContext(QtObjectPtr parent);
void QOpenGLContext_ConnectAboutToBeDestroyed(QtObjectPtr ptr);
void QOpenGLContext_DisconnectAboutToBeDestroyed(QtObjectPtr ptr);
int QOpenGLContext_QOpenGLContext_AreSharing(QtObjectPtr first, QtObjectPtr second);
int QOpenGLContext_Create(QtObjectPtr ptr);
QtObjectPtr QOpenGLContext_QOpenGLContext_CurrentContext();
void QOpenGLContext_DoneCurrent(QtObjectPtr ptr);
QtObjectPtr QOpenGLContext_Functions(QtObjectPtr ptr);
QtObjectPtr QOpenGLContext_QOpenGLContext_GlobalShareContext();
int QOpenGLContext_HasExtension(QtObjectPtr ptr, QtObjectPtr extension);
int QOpenGLContext_IsOpenGLES(QtObjectPtr ptr);
int QOpenGLContext_IsValid(QtObjectPtr ptr);
int QOpenGLContext_MakeCurrent(QtObjectPtr ptr, QtObjectPtr surface);
char* QOpenGLContext_NativeHandle(QtObjectPtr ptr);
void QOpenGLContext_QOpenGLContext_OpenGLModuleHandle();
int QOpenGLContext_QOpenGLContext_OpenGLModuleType();
void QOpenGLContext_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QOpenGLContext_SetNativeHandle(QtObjectPtr ptr, char* handle);
QtObjectPtr QOpenGLContext_Screen(QtObjectPtr ptr);
void QOpenGLContext_SetScreen(QtObjectPtr ptr, QtObjectPtr screen);
void QOpenGLContext_SetShareContext(QtObjectPtr ptr, QtObjectPtr shareContext);
QtObjectPtr QOpenGLContext_ShareContext(QtObjectPtr ptr);
QtObjectPtr QOpenGLContext_ShareGroup(QtObjectPtr ptr);
int QOpenGLContext_QOpenGLContext_SupportsThreadedOpenGL();
QtObjectPtr QOpenGLContext_Surface(QtObjectPtr ptr);
void QOpenGLContext_SwapBuffers(QtObjectPtr ptr, QtObjectPtr surface);
QtObjectPtr QOpenGLContext_VersionFunctions(QtObjectPtr ptr, QtObjectPtr versionProfile);
void QOpenGLContext_DestroyQOpenGLContext(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif