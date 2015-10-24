#include "qopenglfunctions_1_4.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_1_4>
#include "_cgo_export.h"

class MyQOpenGLFunctions_1_4: public QOpenGLFunctions_1_4 {
public:
};

QtObjectPtr QOpenGLFunctions_1_4_NewQOpenGLFunctions_1_4(){
	return new QOpenGLFunctions_1_4();
}

int QOpenGLFunctions_1_4_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_1_4*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_1_4_DestroyQOpenGLFunctions_1_4(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->~QOpenGLFunctions_1_4();
}

void QOpenGLFunctions_1_4_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glEnd();
}

void QOpenGLFunctions_1_4_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glEndList();
}

void QOpenGLFunctions_1_4_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glFinish();
}

void QOpenGLFunctions_1_4_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glFlush();
}

void QOpenGLFunctions_1_4_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glInitNames();
}

void QOpenGLFunctions_1_4_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_1_4_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_1_4_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_1_4_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_1_4_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glPopName();
}

void QOpenGLFunctions_1_4_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_1_4*>(ptr)->glPushMatrix();
}

