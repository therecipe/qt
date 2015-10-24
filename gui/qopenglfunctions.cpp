#include "qopenglfunctions.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLContext>
#include <QOpenGLFunctions>
#include "_cgo_export.h"

class MyQOpenGLFunctions: public QOpenGLFunctions {
public:
};

QtObjectPtr QOpenGLFunctions_NewQOpenGLFunctions(){
	return new QOpenGLFunctions();
}

QtObjectPtr QOpenGLFunctions_NewQOpenGLFunctions2(QtObjectPtr context){
	return new QOpenGLFunctions(static_cast<QOpenGLContext*>(context));
}

void QOpenGLFunctions_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions*>(ptr)->glFinish();
}

void QOpenGLFunctions_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions*>(ptr)->glFlush();
}

void QOpenGLFunctions_GlReleaseShaderCompiler(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions*>(ptr)->glReleaseShaderCompiler();
}

int QOpenGLFunctions_HasOpenGLFeature(QtObjectPtr ptr, int feature){
	return static_cast<QOpenGLFunctions*>(ptr)->hasOpenGLFeature(static_cast<QOpenGLFunctions::OpenGLFeature>(feature));
}

void QOpenGLFunctions_InitializeOpenGLFunctions(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions*>(ptr)->initializeOpenGLFunctions();
}

int QOpenGLFunctions_OpenGLFeatures(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions*>(ptr)->openGLFeatures();
}

void QOpenGLFunctions_DestroyQOpenGLFunctions(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions*>(ptr)->~QOpenGLFunctions();
}

