#include "qopenglfunctions_3_1.h"
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QOpenGLFunctions_3_1>
#include "_cgo_export.h"

class MyQOpenGLFunctions_3_1: public QOpenGLFunctions_3_1 {
public:
};

QtObjectPtr QOpenGLFunctions_3_1_NewQOpenGLFunctions_3_1(){
	return new QOpenGLFunctions_3_1();
}

int QOpenGLFunctions_3_1_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_3_1*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_3_1_DestroyQOpenGLFunctions_3_1(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_1*>(ptr)->~QOpenGLFunctions_3_1();
}

void QOpenGLFunctions_3_1_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_1*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_3_1_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_1*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_3_1_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_1*>(ptr)->glFinish();
}

void QOpenGLFunctions_3_1_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_1*>(ptr)->glFlush();
}

