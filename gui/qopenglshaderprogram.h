#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLShaderProgram_NewQOpenGLShaderProgram(QtObjectPtr parent);
int QOpenGLShaderProgram_AddShader(QtObjectPtr ptr, QtObjectPtr shader);
int QOpenGLShaderProgram_AddShaderFromSourceCode2(QtObjectPtr ptr, int ty, QtObjectPtr source);
int QOpenGLShaderProgram_AddShaderFromSourceCode3(QtObjectPtr ptr, int ty, char* source);
int QOpenGLShaderProgram_AddShaderFromSourceCode(QtObjectPtr ptr, int ty, char* source);
int QOpenGLShaderProgram_AddShaderFromSourceFile(QtObjectPtr ptr, int ty, char* fileName);
int QOpenGLShaderProgram_AttributeLocation2(QtObjectPtr ptr, QtObjectPtr name);
int QOpenGLShaderProgram_AttributeLocation3(QtObjectPtr ptr, char* name);
int QOpenGLShaderProgram_AttributeLocation(QtObjectPtr ptr, char* name);
int QOpenGLShaderProgram_Bind(QtObjectPtr ptr);
void QOpenGLShaderProgram_BindAttributeLocation2(QtObjectPtr ptr, QtObjectPtr name, int location);
void QOpenGLShaderProgram_BindAttributeLocation3(QtObjectPtr ptr, char* name, int location);
void QOpenGLShaderProgram_BindAttributeLocation(QtObjectPtr ptr, char* name, int location);
int QOpenGLShaderProgram_Create(QtObjectPtr ptr);
void QOpenGLShaderProgram_DisableAttributeArray2(QtObjectPtr ptr, char* name);
void QOpenGLShaderProgram_DisableAttributeArray(QtObjectPtr ptr, int location);
void QOpenGLShaderProgram_EnableAttributeArray2(QtObjectPtr ptr, char* name);
void QOpenGLShaderProgram_EnableAttributeArray(QtObjectPtr ptr, int location);
int QOpenGLShaderProgram_QOpenGLShaderProgram_HasOpenGLShaderPrograms(QtObjectPtr context);
int QOpenGLShaderProgram_IsLinked(QtObjectPtr ptr);
int QOpenGLShaderProgram_Link(QtObjectPtr ptr);
char* QOpenGLShaderProgram_Log(QtObjectPtr ptr);
int QOpenGLShaderProgram_MaxGeometryOutputVertices(QtObjectPtr ptr);
int QOpenGLShaderProgram_PatchVertexCount(QtObjectPtr ptr);
void QOpenGLShaderProgram_Release(QtObjectPtr ptr);
void QOpenGLShaderProgram_RemoveAllShaders(QtObjectPtr ptr);
void QOpenGLShaderProgram_RemoveShader(QtObjectPtr ptr, QtObjectPtr shader);
void QOpenGLShaderProgram_SetAttributeArray7(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeArray8(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeArray9(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeArray2(QtObjectPtr ptr, int location, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeArray3(QtObjectPtr ptr, int location, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeArray4(QtObjectPtr ptr, int location, QtObjectPtr values, int stride);
void QOpenGLShaderProgram_SetAttributeValue17(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue14(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue15(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue16(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue8(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue5(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue6(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetAttributeValue7(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetPatchVertexCount(QtObjectPtr ptr, int count);
void QOpenGLShaderProgram_SetUniformValue34(QtObjectPtr ptr, char* name, QtObjectPtr color);
void QOpenGLShaderProgram_SetUniformValue47(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue35(QtObjectPtr ptr, char* name, QtObjectPtr point);
void QOpenGLShaderProgram_SetUniformValue36(QtObjectPtr ptr, char* name, QtObjectPtr point);
void QOpenGLShaderProgram_SetUniformValue37(QtObjectPtr ptr, char* name, QtObjectPtr size);
void QOpenGLShaderProgram_SetUniformValue38(QtObjectPtr ptr, char* name, QtObjectPtr size);
void QOpenGLShaderProgram_SetUniformValue54(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue31(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue32(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue33(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue10(QtObjectPtr ptr, int location, QtObjectPtr color);
void QOpenGLShaderProgram_SetUniformValue23(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue11(QtObjectPtr ptr, int location, QtObjectPtr point);
void QOpenGLShaderProgram_SetUniformValue12(QtObjectPtr ptr, int location, QtObjectPtr point);
void QOpenGLShaderProgram_SetUniformValue13(QtObjectPtr ptr, int location, QtObjectPtr size);
void QOpenGLShaderProgram_SetUniformValue14(QtObjectPtr ptr, int location, QtObjectPtr size);
void QOpenGLShaderProgram_SetUniformValue24(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue7(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue8(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValue9(QtObjectPtr ptr, int location, QtObjectPtr value);
void QOpenGLShaderProgram_SetUniformValueArray30(QtObjectPtr ptr, char* name, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray19(QtObjectPtr ptr, char* name, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray20(QtObjectPtr ptr, char* name, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray21(QtObjectPtr ptr, char* name, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray15(QtObjectPtr ptr, int location, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray4(QtObjectPtr ptr, int location, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray5(QtObjectPtr ptr, int location, QtObjectPtr values, int count);
void QOpenGLShaderProgram_SetUniformValueArray6(QtObjectPtr ptr, int location, QtObjectPtr values, int count);
int QOpenGLShaderProgram_UniformLocation2(QtObjectPtr ptr, QtObjectPtr name);
int QOpenGLShaderProgram_UniformLocation3(QtObjectPtr ptr, char* name);
int QOpenGLShaderProgram_UniformLocation(QtObjectPtr ptr, char* name);
void QOpenGLShaderProgram_DestroyQOpenGLShaderProgram(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif