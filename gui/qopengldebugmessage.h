#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLDebugMessage_NewQOpenGLDebugMessage();
QtObjectPtr QOpenGLDebugMessage_NewQOpenGLDebugMessage2(QtObjectPtr debugMessage);
char* QOpenGLDebugMessage_Message(QtObjectPtr ptr);
int QOpenGLDebugMessage_Severity(QtObjectPtr ptr);
int QOpenGLDebugMessage_Source(QtObjectPtr ptr);
void QOpenGLDebugMessage_Swap(QtObjectPtr ptr, QtObjectPtr debugMessage);
int QOpenGLDebugMessage_Type(QtObjectPtr ptr);
void QOpenGLDebugMessage_DestroyQOpenGLDebugMessage(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif