#include "qopenglfunctions_4_0_core.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_4_0_Core>
#include "_cgo_export.h"

class MyQOpenGLFunctions_4_0_Core: public QOpenGLFunctions_4_0_Core {
public:
};

QtObjectPtr QOpenGLFunctions_4_0_Core_NewQOpenGLFunctions_4_0_Core(){
	return new QOpenGLFunctions_4_0_Core();
}

int QOpenGLFunctions_4_0_Core_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_4_0_Core_DestroyQOpenGLFunctions_4_0_Core(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->~QOpenGLFunctions_4_0_Core();
}

void QOpenGLFunctions_4_0_Core_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_4_0_Core_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_4_0_Core_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glFinish();
}

void QOpenGLFunctions_4_0_Core_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glFlush();
}

void QOpenGLFunctions_4_0_Core_GlPauseTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glPauseTransformFeedback();
}

void QOpenGLFunctions_4_0_Core_GlResumeTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_0_Core*>(ptr)->glResumeTransformFeedback();
}

