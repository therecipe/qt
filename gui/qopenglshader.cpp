#include "qopenglshader.h"
#include <QModelIndex>
#include <QObject>
#include <QOpenGLContext>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLShader>
#include "_cgo_export.h"

class MyQOpenGLShader: public QOpenGLShader {
public:
};

QtObjectPtr QOpenGLShader_NewQOpenGLShader(int ty, QtObjectPtr parent){
	return new QOpenGLShader(static_cast<QOpenGLShader::ShaderTypeBit>(ty), static_cast<QObject*>(parent));
}

int QOpenGLShader_CompileSourceCode2(QtObjectPtr ptr, QtObjectPtr source){
	return static_cast<QOpenGLShader*>(ptr)->compileSourceCode(*static_cast<QByteArray*>(source));
}

int QOpenGLShader_CompileSourceCode3(QtObjectPtr ptr, char* source){
	return static_cast<QOpenGLShader*>(ptr)->compileSourceCode(QString(source));
}

int QOpenGLShader_CompileSourceCode(QtObjectPtr ptr, char* source){
	return static_cast<QOpenGLShader*>(ptr)->compileSourceCode(const_cast<const char*>(source));
}

int QOpenGLShader_CompileSourceFile(QtObjectPtr ptr, char* fileName){
	return static_cast<QOpenGLShader*>(ptr)->compileSourceFile(QString(fileName));
}

int QOpenGLShader_QOpenGLShader_HasOpenGLShaders(int ty, QtObjectPtr context){
	return QOpenGLShader::hasOpenGLShaders(static_cast<QOpenGLShader::ShaderTypeBit>(ty), static_cast<QOpenGLContext*>(context));
}

int QOpenGLShader_IsCompiled(QtObjectPtr ptr){
	return static_cast<QOpenGLShader*>(ptr)->isCompiled();
}

char* QOpenGLShader_Log(QtObjectPtr ptr){
	return static_cast<QOpenGLShader*>(ptr)->log().toUtf8().data();
}

int QOpenGLShader_ShaderType(QtObjectPtr ptr){
	return static_cast<QOpenGLShader*>(ptr)->shaderType();
}

void QOpenGLShader_DestroyQOpenGLShader(QtObjectPtr ptr){
	static_cast<QOpenGLShader*>(ptr)->~QOpenGLShader();
}

