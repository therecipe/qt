#include "qopenglfunctions_3_3_core.h"
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLFunctions_3_3_Core>
#include "_cgo_export.h"

class MyQOpenGLFunctions_3_3_Core: public QOpenGLFunctions_3_3_Core {
public:
};

QtObjectPtr QOpenGLFunctions_3_3_Core_NewQOpenGLFunctions_3_3_Core(){
	return new QOpenGLFunctions_3_3_Core();
}

int QOpenGLFunctions_3_3_Core_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_3_3_Core_DestroyQOpenGLFunctions_3_3_Core(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->~QOpenGLFunctions_3_3_Core();
}

void QOpenGLFunctions_3_3_Core_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_3_3_Core_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_3_3_Core_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->glFinish();
}

void QOpenGLFunctions_3_3_Core_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_3_Core*>(ptr)->glFlush();
}

