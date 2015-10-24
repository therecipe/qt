#include "qopenglfunctions_1_1.h"
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLFunctions_1_1>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_1: public QOpenGLFunctions_1_1 {
public:
};

QtObjectPtr QOpenGLFunctions_1_1_NewQOpenGLFunctions_1_1(){
	return new QOpenGLFunctions_1_1();
}

int QOpenGLFunctions_1_1_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_1*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_1_DestroyQOpenGLFunctions_1_1(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->~QOpenGLFunctions_1_1();
}

void QOpenGLFunctions_1_1_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_1_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_1_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_1_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_1_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_1_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_1_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_1_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_1_1_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_1_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_1_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_1*>(ptr)->glPushMatrix();
}

