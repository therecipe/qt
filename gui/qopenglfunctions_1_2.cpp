#include "qopenglfunctions_1_2.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_1_2>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_2: public QOpenGLFunctions_1_2 {
public:
};

QtObjectPtr QOpenGLFunctions_1_2_NewQOpenGLFunctions_1_2(){
	return new QOpenGLFunctions_1_2();
}

int QOpenGLFunctions_1_2_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_2*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_2_DestroyQOpenGLFunctions_1_2(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->~QOpenGLFunctions_1_2();
}

void QOpenGLFunctions_1_2_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_2_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_2_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_2_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_2_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_2_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_2_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_2_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_1_2_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_2_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_2_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_2*>(ptr)->glPushMatrix();
}

