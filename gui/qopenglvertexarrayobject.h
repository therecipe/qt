#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLVertexArrayObject_NewQOpenGLVertexArrayObject(QtObjectPtr parent);
void QOpenGLVertexArrayObject_Bind(QtObjectPtr ptr);
int QOpenGLVertexArrayObject_Create(QtObjectPtr ptr);
void QOpenGLVertexArrayObject_Destroy(QtObjectPtr ptr);
int QOpenGLVertexArrayObject_IsCreated(QtObjectPtr ptr);
void QOpenGLVertexArrayObject_Release(QtObjectPtr ptr);
void QOpenGLVertexArrayObject_DestroyQOpenGLVertexArrayObject(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif