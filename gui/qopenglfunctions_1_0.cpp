#include "qopenglfunctions_1_0.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QString>
#include <QOpenGLFunctions_1_0>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_0: public QOpenGLFunctions_1_0 {
public:
};

QtObjectPtr QOpenGLFunctions_1_0_NewQOpenGLFunctions_1_0(){
	return new QOpenGLFunctions_1_0();
}

int QOpenGLFunctions_1_0_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_0*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_0_DestroyQOpenGLFunctions_1_0(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->~QOpenGLFunctions_1_0();
}

void QOpenGLFunctions_1_0_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_0_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_0_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_0_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_0_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_0_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_0_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_0_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_0_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_0_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_0*>(ptr)->glPushMatrix();
}

