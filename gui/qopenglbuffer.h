#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer();
QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer2(int ty);
QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer3(QtObjectPtr other);
void QOpenGLBuffer_Allocate2(QtObjectPtr ptr, int count);
int QOpenGLBuffer_Bind(QtObjectPtr ptr);
int QOpenGLBuffer_Create(QtObjectPtr ptr);
void QOpenGLBuffer_Destroy(QtObjectPtr ptr);
int QOpenGLBuffer_IsCreated(QtObjectPtr ptr);
void QOpenGLBuffer_Map(QtObjectPtr ptr, int access);
void QOpenGLBuffer_MapRange(QtObjectPtr ptr, int offset, int count, int access);
void QOpenGLBuffer_Release(QtObjectPtr ptr);
void QOpenGLBuffer_QOpenGLBuffer_Release2(int ty);
void QOpenGLBuffer_SetUsagePattern(QtObjectPtr ptr, int value);
int QOpenGLBuffer_Size(QtObjectPtr ptr);
int QOpenGLBuffer_Type(QtObjectPtr ptr);
int QOpenGLBuffer_Unmap(QtObjectPtr ptr);
int QOpenGLBuffer_UsagePattern(QtObjectPtr ptr);
void QOpenGLBuffer_DestroyQOpenGLBuffer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif