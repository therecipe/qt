#include "qopenglfunctions_3_0.h"
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLFunctions_3_0>
#include "_cgo_export.h"

class MyQOpenGLFunctions_3_0: public QOpenGLFunctions_3_0 {
public:
};

QtObjectPtr QOpenGLFunctions_3_0_NewQOpenGLFunctions_3_0(){
	return new QOpenGLFunctions_3_0();
}

int QOpenGLFunctions_3_0_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_3_0*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_3_0_DestroyQOpenGLFunctions_3_0(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->~QOpenGLFunctions_3_0();
}

void QOpenGLFunctions_3_0_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glEnd();
}

void QOpenGLFunctions_3_0_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_3_0_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glEndList();
}

void QOpenGLFunctions_3_0_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_3_0_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glFinish();
}

void QOpenGLFunctions_3_0_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glFlush();
}

void QOpenGLFunctions_3_0_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glInitNames();
}

void QOpenGLFunctions_3_0_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_3_0_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_3_0_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_3_0_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_3_0_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glPopName();
}

void QOpenGLFunctions_3_0_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_0*>(ptr)->glPushMatrix();
}

