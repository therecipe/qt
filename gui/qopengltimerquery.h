#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLTimerQuery_NewQOpenGLTimerQuery(QtObjectPtr parent);
void QOpenGLTimerQuery_Begin(QtObjectPtr ptr);
int QOpenGLTimerQuery_Create(QtObjectPtr ptr);
void QOpenGLTimerQuery_Destroy(QtObjectPtr ptr);
void QOpenGLTimerQuery_End(QtObjectPtr ptr);
int QOpenGLTimerQuery_IsCreated(QtObjectPtr ptr);
int QOpenGLTimerQuery_IsResultAvailable(QtObjectPtr ptr);
void QOpenGLTimerQuery_RecordTimestamp(QtObjectPtr ptr);
void QOpenGLTimerQuery_DestroyQOpenGLTimerQuery(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif