#include "qopenglshaderprogram.h"
#include <QString>
#include <QMatrix4x4>
#include <QVector2D>
#include <QTransform>
#include <QSize>
#include <QVariant>
#include <QModelIndex>
#include <QPoint>
#include <QByteArray>
#include <QVector3D>
#include <QColor>
#include <QOpenGLShader>
#include <QVector4D>
#include <QObject>
#include <QUrl>
#include <QOpenGLContext>
#include <QVector>
#include <QPointF>
#include <QSizeF>
#include <QOpenGLShaderProgram>
#include "_cgo_export.h"

class MyQOpenGLShaderProgram: public QOpenGLShaderProgram {
public:
};

QtObjectPtr QOpenGLShaderProgram_NewQOpenGLShaderProgram(QtObjectPtr parent){
	return new QOpenGLShaderProgram(static_cast<QObject*>(parent));
}

int QOpenGLShaderProgram_AddShader(QtObjectPtr ptr, QtObjectPtr shader){
	return static_cast<QOpenGLShaderProgram*>(ptr)->addShader(static_cast<QOpenGLShader*>(shader));
}

int QOpenGLShaderProgram_AddShaderFromSourceCode2(QtObjectPtr ptr, int ty, QtObjectPtr source){
	return static_cast<QOpenGLShaderProgram*>(ptr)->addShaderFromSourceCode(static_cast<QOpenGLShader::ShaderTypeBit>(ty), *static_cast<QByteArray*>(source));
}

int QOpenGLShaderProgram_AddShaderFromSourceCode3(QtObjectPtr ptr, int ty, char* source){
	return static_cast<QOpenGLShaderProgram*>(ptr)->addShaderFromSourceCode(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString(source));
}

int QOpenGLShaderProgram_AddShaderFromSourceCode(QtObjectPtr ptr, int ty, char* source){
	return static_cast<QOpenGLShaderProgram*>(ptr)->addShaderFromSourceCode(static_cast<QOpenGLShader::ShaderTypeBit>(ty), const_cast<const char*>(source));
}

int QOpenGLShaderProgram_AddShaderFromSourceFile(QtObjectPtr ptr, int ty, char* fileName){
	return static_cast<QOpenGLShaderProgram*>(ptr)->addShaderFromSourceFile(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString(fileName));
}

int QOpenGLShaderProgram_AttributeLocation2(QtObjectPtr ptr, QtObjectPtr name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->attributeLocation(*static_cast<QByteArray*>(name));
}

int QOpenGLShaderProgram_AttributeLocation3(QtObjectPtr ptr, char* name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->attributeLocation(QString(name));
}

int QOpenGLShaderProgram_AttributeLocation(QtObjectPtr ptr, char* name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->attributeLocation(const_cast<const char*>(name));
}

int QOpenGLShaderProgram_Bind(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->bind();
}

void QOpenGLShaderProgram_BindAttributeLocation2(QtObjectPtr ptr, QtObjectPtr name, int location){
	static_cast<QOpenGLShaderProgram*>(ptr)->bindAttributeLocation(*static_cast<QByteArray*>(name), location);
}

void QOpenGLShaderProgram_BindAttributeLocation3(QtObjectPtr ptr, char* name, int location){
	static_cast<QOpenGLShaderProgram*>(ptr)->bindAttributeLocation(QString(name), location);
}

void QOpenGLShaderProgram_BindAttributeLocation(QtObjectPtr ptr, char* name, int location){
	static_cast<QOpenGLShaderProgram*>(ptr)->bindAttributeLocation(const_cast<const char*>(name), location);
}

int QOpenGLShaderProgram_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->create();
}

void QOpenGLShaderProgram_DisableAttributeArray2(QtObjectPtr ptr, char* name){
	static_cast<QOpenGLShaderProgram*>(ptr)->disableAttributeArray(const_cast<const char*>(name));
}

void QOpenGLShaderProgram_DisableAttributeArray(QtObjectPtr ptr, int location){
	static_cast<QOpenGLShaderProgram*>(ptr)->disableAttributeArray(location);
}

void QOpenGLShaderProgram_EnableAttributeArray2(QtObjectPtr ptr, char* name){
	static_cast<QOpenGLShaderProgram*>(ptr)->enableAttributeArray(const_cast<const char*>(name));
}

void QOpenGLShaderProgram_EnableAttributeArray(QtObjectPtr ptr, int location){
	static_cast<QOpenGLShaderProgram*>(ptr)->enableAttributeArray(location);
}

int QOpenGLShaderProgram_QOpenGLShaderProgram_HasOpenGLShaderPrograms(QtObjectPtr context){
	return QOpenGLShaderProgram::hasOpenGLShaderPrograms(static_cast<QOpenGLContext*>(context));
}

int QOpenGLShaderProgram_IsLinked(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->isLinked();
}

int QOpenGLShaderProgram_Link(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->link();
}

char* QOpenGLShaderProgram_Log(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->log().toUtf8().data();
}

int QOpenGLShaderProgram_MaxGeometryOutputVertices(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->maxGeometryOutputVertices();
}

int QOpenGLShaderProgram_PatchVertexCount(QtObjectPtr ptr){
	return static_cast<QOpenGLShaderProgram*>(ptr)->patchVertexCount();
}

void QOpenGLShaderProgram_Release(QtObjectPtr ptr){
	static_cast<QOpenGLShaderProgram*>(ptr)->release();
}

void QOpenGLShaderProgram_RemoveAllShaders(QtObjectPtr ptr){
	static_cast<QOpenGLShaderProgram*>(ptr)->removeAllShaders();
}

void QOpenGLShaderProgram_RemoveShader(QtObjectPtr ptr, QtObjectPtr shader){
	static_cast<QOpenGLShaderProgram*>(ptr)->removeShader(static_cast<QOpenGLShader*>(shader));
}

void QOpenGLShaderProgram_SetAttributeArray7(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(const_cast<const char*>(name), static_cast<QVector2D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeArray8(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(const_cast<const char*>(name), static_cast<QVector3D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeArray9(QtObjectPtr ptr, char* name, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(const_cast<const char*>(name), static_cast<QVector4D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeArray2(QtObjectPtr ptr, int location, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(location, static_cast<QVector2D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeArray3(QtObjectPtr ptr, int location, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(location, static_cast<QVector3D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeArray4(QtObjectPtr ptr, int location, QtObjectPtr values, int stride){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeArray(location, static_cast<QVector4D*>(values), stride);
}

void QOpenGLShaderProgram_SetAttributeValue17(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(const_cast<const char*>(name), *static_cast<QColor*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue14(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(const_cast<const char*>(name), *static_cast<QVector2D*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue15(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(const_cast<const char*>(name), *static_cast<QVector3D*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue16(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(const_cast<const char*>(name), *static_cast<QVector4D*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue8(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(location, *static_cast<QColor*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue5(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(location, *static_cast<QVector2D*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue6(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(location, *static_cast<QVector3D*>(value));
}

void QOpenGLShaderProgram_SetAttributeValue7(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setAttributeValue(location, *static_cast<QVector4D*>(value));
}

void QOpenGLShaderProgram_SetPatchVertexCount(QtObjectPtr ptr, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setPatchVertexCount(count);
}

void QOpenGLShaderProgram_SetUniformValue34(QtObjectPtr ptr, char* name, QtObjectPtr color){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QColor*>(color));
}

void QOpenGLShaderProgram_SetUniformValue47(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QMatrix4x4*>(value));
}

void QOpenGLShaderProgram_SetUniformValue35(QtObjectPtr ptr, char* name, QtObjectPtr point){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QPoint*>(point));
}

void QOpenGLShaderProgram_SetUniformValue36(QtObjectPtr ptr, char* name, QtObjectPtr point){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QPointF*>(point));
}

void QOpenGLShaderProgram_SetUniformValue37(QtObjectPtr ptr, char* name, QtObjectPtr size){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QSize*>(size));
}

void QOpenGLShaderProgram_SetUniformValue38(QtObjectPtr ptr, char* name, QtObjectPtr size){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QSizeF*>(size));
}

void QOpenGLShaderProgram_SetUniformValue54(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QTransform*>(value));
}

void QOpenGLShaderProgram_SetUniformValue31(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QVector2D*>(value));
}

void QOpenGLShaderProgram_SetUniformValue32(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QVector3D*>(value));
}

void QOpenGLShaderProgram_SetUniformValue33(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(const_cast<const char*>(name), *static_cast<QVector4D*>(value));
}

void QOpenGLShaderProgram_SetUniformValue10(QtObjectPtr ptr, int location, QtObjectPtr color){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QColor*>(color));
}

void QOpenGLShaderProgram_SetUniformValue23(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QMatrix4x4*>(value));
}

void QOpenGLShaderProgram_SetUniformValue11(QtObjectPtr ptr, int location, QtObjectPtr point){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QPoint*>(point));
}

void QOpenGLShaderProgram_SetUniformValue12(QtObjectPtr ptr, int location, QtObjectPtr point){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QPointF*>(point));
}

void QOpenGLShaderProgram_SetUniformValue13(QtObjectPtr ptr, int location, QtObjectPtr size){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QSize*>(size));
}

void QOpenGLShaderProgram_SetUniformValue14(QtObjectPtr ptr, int location, QtObjectPtr size){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QSizeF*>(size));
}

void QOpenGLShaderProgram_SetUniformValue24(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QTransform*>(value));
}

void QOpenGLShaderProgram_SetUniformValue7(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QVector2D*>(value));
}

void QOpenGLShaderProgram_SetUniformValue8(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QVector3D*>(value));
}

void QOpenGLShaderProgram_SetUniformValue9(QtObjectPtr ptr, int location, QtObjectPtr value){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValue(location, *static_cast<QVector4D*>(value));
}

void QOpenGLShaderProgram_SetUniformValueArray30(QtObjectPtr ptr, char* name, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(const_cast<const char*>(name), static_cast<QMatrix4x4*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray19(QtObjectPtr ptr, char* name, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(const_cast<const char*>(name), static_cast<QVector2D*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray20(QtObjectPtr ptr, char* name, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(const_cast<const char*>(name), static_cast<QVector3D*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray21(QtObjectPtr ptr, char* name, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(const_cast<const char*>(name), static_cast<QVector4D*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray15(QtObjectPtr ptr, int location, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(location, static_cast<QMatrix4x4*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray4(QtObjectPtr ptr, int location, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(location, static_cast<QVector2D*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray5(QtObjectPtr ptr, int location, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(location, static_cast<QVector3D*>(values), count);
}

void QOpenGLShaderProgram_SetUniformValueArray6(QtObjectPtr ptr, int location, QtObjectPtr values, int count){
	static_cast<QOpenGLShaderProgram*>(ptr)->setUniformValueArray(location, static_cast<QVector4D*>(values), count);
}

int QOpenGLShaderProgram_UniformLocation2(QtObjectPtr ptr, QtObjectPtr name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->uniformLocation(*static_cast<QByteArray*>(name));
}

int QOpenGLShaderProgram_UniformLocation3(QtObjectPtr ptr, char* name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->uniformLocation(QString(name));
}

int QOpenGLShaderProgram_UniformLocation(QtObjectPtr ptr, char* name){
	return static_cast<QOpenGLShaderProgram*>(ptr)->uniformLocation(const_cast<const char*>(name));
}

void QOpenGLShaderProgram_DestroyQOpenGLShaderProgram(QtObjectPtr ptr){
	static_cast<QOpenGLShaderProgram*>(ptr)->~QOpenGLShaderProgram();
}

