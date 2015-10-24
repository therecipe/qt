#include "qopenglvertexarrayobject.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QOpenGLVertexArrayObject>
#include "_cgo_export.h"

class MyQOpenGLVertexArrayObject: public QOpenGLVertexArrayObject {
public:
};

QtObjectPtr QOpenGLVertexArrayObject_NewQOpenGLVertexArrayObject(QtObjectPtr parent){
	return new QOpenGLVertexArrayObject(static_cast<QObject*>(parent));
}

void QOpenGLVertexArrayObject_Bind(QtObjectPtr ptr){
	static_cast<QOpenGLVertexArrayObject*>(ptr)->bind();
}

int QOpenGLVertexArrayObject_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLVertexArrayObject*>(ptr)->create();
}

void QOpenGLVertexArrayObject_Destroy(QtObjectPtr ptr){
	static_cast<QOpenGLVertexArrayObject*>(ptr)->destroy();
}

int QOpenGLVertexArrayObject_IsCreated(QtObjectPtr ptr){
	return static_cast<QOpenGLVertexArrayObject*>(ptr)->isCreated();
}

void QOpenGLVertexArrayObject_Release(QtObjectPtr ptr){
	static_cast<QOpenGLVertexArrayObject*>(ptr)->release();
}

void QOpenGLVertexArrayObject_DestroyQOpenGLVertexArrayObject(QtObjectPtr ptr){
	static_cast<QOpenGLVertexArrayObject*>(ptr)->~QOpenGLVertexArrayObject();
}

