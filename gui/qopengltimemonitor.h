#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLTimeMonitor_NewQOpenGLTimeMonitor(QtObjectPtr parent);
int QOpenGLTimeMonitor_Create(QtObjectPtr ptr);
void QOpenGLTimeMonitor_Destroy(QtObjectPtr ptr);
int QOpenGLTimeMonitor_IsCreated(QtObjectPtr ptr);
int QOpenGLTimeMonitor_IsResultAvailable(QtObjectPtr ptr);
int QOpenGLTimeMonitor_RecordSample(QtObjectPtr ptr);
void QOpenGLTimeMonitor_Reset(QtObjectPtr ptr);
int QOpenGLTimeMonitor_SampleCount(QtObjectPtr ptr);
void QOpenGLTimeMonitor_SetSampleCount(QtObjectPtr ptr, int sampleCount);
void QOpenGLTimeMonitor_DestroyQOpenGLTimeMonitor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif