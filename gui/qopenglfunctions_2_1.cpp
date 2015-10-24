#include "qopenglfunctions_2_1.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_2_1>
#include "_cgo_export.h"

class MyQOpenGLFunctions_2_1: public QOpenGLFunctions_2_1 {
public:
};

QtObjectPtr QOpenGLFunctions_2_1_NewQOpenGLFunctions_2_1(){
	return new QOpenGLFunctions_2_1();
}

int QOpenGLFunctions_2_1_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_2_1*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_2_1_DestroyQOpenGLFunctions_2_1(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->~QOpenGLFunctions_2_1();
}

void QOpenGLFunctions_2_1_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glEnd();
}

void QOpenGLFunctions_2_1_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glEndList();
}

void QOpenGLFunctions_2_1_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glFinish();
}

void QOpenGLFunctions_2_1_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glFlush();
}

void QOpenGLFunctions_2_1_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glInitNames();
}

void QOpenGLFunctions_2_1_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_2_1_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_2_1_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_2_1_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_2_1_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glPopName();
}

void QOpenGLFunctions_2_1_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_1*>(ptr)->glPushMatrix();
}

