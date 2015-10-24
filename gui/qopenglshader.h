#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLShader_NewQOpenGLShader(int ty, QtObjectPtr parent);
int QOpenGLShader_CompileSourceCode2(QtObjectPtr ptr, QtObjectPtr source);
int QOpenGLShader_CompileSourceCode3(QtObjectPtr ptr, char* source);
int QOpenGLShader_CompileSourceCode(QtObjectPtr ptr, char* source);
int QOpenGLShader_CompileSourceFile(QtObjectPtr ptr, char* fileName);
int QOpenGLShader_QOpenGLShader_HasOpenGLShaders(int ty, QtObjectPtr context);
int QOpenGLShader_IsCompiled(QtObjectPtr ptr);
char* QOpenGLShader_Log(QtObjectPtr ptr);
int QOpenGLShader_ShaderType(QtObjectPtr ptr);
void QOpenGLShader_DestroyQOpenGLShader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif