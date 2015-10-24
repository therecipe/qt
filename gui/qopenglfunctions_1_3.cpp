#include "qopenglfunctions_1_3.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_1_3>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_3: public QOpenGLFunctions_1_3 {
public:
};

QtObjectPtr QOpenGLFunctions_1_3_NewQOpenGLFunctions_1_3(){
	return new QOpenGLFunctions_1_3();
}

int QOpenGLFunctions_1_3_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_3*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_3_DestroyQOpenGLFunctions_1_3(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->~QOpenGLFunctions_1_3();
}

void QOpenGLFunctions_1_3_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_3_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_3_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_3_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_3_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_3_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_3_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_3_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_1_3_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_3_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_3_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_3*>(ptr)->glPushMatrix();
}

