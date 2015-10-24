#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLDebugLogger_NewQOpenGLDebugLogger(QtObjectPtr parent);
void QOpenGLDebugLogger_DisableMessages(QtObjectPtr ptr, int sources, int types, int severities);
void QOpenGLDebugLogger_EnableMessages(QtObjectPtr ptr, int sources, int types, int severities);
int QOpenGLDebugLogger_Initialize(QtObjectPtr ptr);
int QOpenGLDebugLogger_IsLogging(QtObjectPtr ptr);
void QOpenGLDebugLogger_LogMessage(QtObjectPtr ptr, QtObjectPtr debugMessage);
int QOpenGLDebugLogger_LoggingMode(QtObjectPtr ptr);
void QOpenGLDebugLogger_PopGroup(QtObjectPtr ptr);
void QOpenGLDebugLogger_StartLogging(QtObjectPtr ptr, int loggingMode);
void QOpenGLDebugLogger_StopLogging(QtObjectPtr ptr);
void QOpenGLDebugLogger_DestroyQOpenGLDebugLogger(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif