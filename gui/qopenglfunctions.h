#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLFunctions_NewQOpenGLFunctions();
QtObjectPtr QOpenGLFunctions_NewQOpenGLFunctions2(QtObjectPtr context);
void QOpenGLFunctions_GlFinish(QtObjectPtr ptr);
void QOpenGLFunctions_GlFlush(QtObjectPtr ptr);
void QOpenGLFunctions_GlReleaseShaderCompiler(QtObjectPtr ptr);
int QOpenGLFunctions_HasOpenGLFeature(QtObjectPtr ptr, int feature);
void QOpenGLFunctions_InitializeOpenGLFunctions(QtObjectPtr ptr);
int QOpenGLFunctions_OpenGLFeatures(QtObjectPtr ptr);
void QOpenGLFunctions_DestroyQOpenGLFunctions(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif