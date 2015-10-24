#include "qopenglfunctions_3_2_core.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_3_2_Core>
#include "_cgo_export.h"

class MyQOpenGLFunctions_3_2_Core: public QOpenGLFunctions_3_2_Core {
public:
};

QtObjectPtr QOpenGLFunctions_3_2_Core_NewQOpenGLFunctions_3_2_Core(){
	return new QOpenGLFunctions_3_2_Core();
}

int QOpenGLFunctions_3_2_Core_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_3_2_Core_DestroyQOpenGLFunctions_3_2_Core(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->~QOpenGLFunctions_3_2_Core();
}

void QOpenGLFunctions_3_2_Core_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_3_2_Core_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_3_2_Core_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->glFinish();
}

void QOpenGLFunctions_3_2_Core_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_3_2_Core*>(ptr)->glFlush();
}

