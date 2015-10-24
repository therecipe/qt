#include "qopenglfunctions_2_0.h"
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions_2_0>
#include "_cgo_export.h"

class MyQOpenGLFunctions_2_0: public QOpenGLFunctions_2_0 {
public:
};

QtObjectPtr QOpenGLFunctions_2_0_NewQOpenGLFunctions_2_0(){
	return new QOpenGLFunctions_2_0();
}

int QOpenGLFunctions_2_0_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_2_0*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_2_0_DestroyQOpenGLFunctions_2_0(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->~QOpenGLFunctions_2_0();
}

void QOpenGLFunctions_2_0_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glEnd();
}

void QOpenGLFunctions_2_0_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glEndList();
}

void QOpenGLFunctions_2_0_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glFinish();
}

void QOpenGLFunctions_2_0_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glFlush();
}

void QOpenGLFunctions_2_0_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glInitNames();
}

void QOpenGLFunctions_2_0_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_2_0_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_2_0_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_2_0_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_2_0_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glPopName();
}

void QOpenGLFunctions_2_0_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_2_0*>(ptr)->glPushMatrix();
}

