#include "qopenglfunctions_1_5.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_1_5>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_5: public QOpenGLFunctions_1_5 {
public:
};

QtObjectPtr QOpenGLFunctions_1_5_NewQOpenGLFunctions_1_5(){
	return new QOpenGLFunctions_1_5();
}

int QOpenGLFunctions_1_5_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_5*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_5_DestroyQOpenGLFunctions_1_5(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->~QOpenGLFunctions_1_5();
}

void QOpenGLFunctions_1_5_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_5_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_5_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_5_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_5_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_5_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_5_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_5_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_1_5_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_5_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_5_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_5*>(ptr)->glPushMatrix();
}

