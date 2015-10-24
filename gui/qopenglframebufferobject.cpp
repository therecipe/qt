#include "qopenglframebufferobject.h"
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFramebufferObjectFormat>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QOpenGLFramebufferObject>
#include "_cgo_export.h"

class MyQOpenGLFramebufferObject: public QOpenGLFramebufferObject {
public:
};

int QOpenGLFramebufferObject_Bind(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->bind();
}

int QOpenGLFramebufferObject_QOpenGLFramebufferObject_BindDefault(){
	return QOpenGLFramebufferObject::bindDefault();
}

int QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferBlit(){
	return QOpenGLFramebufferObject::hasOpenGLFramebufferBlit();
}

int QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferObjects(){
	return QOpenGLFramebufferObject::hasOpenGLFramebufferObjects();
}

int QOpenGLFramebufferObject_IsValid(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->isValid();
}

int QOpenGLFramebufferObject_Release(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->release();
}

void QOpenGLFramebufferObject_DestroyQOpenGLFramebufferObject(QtObjectPtr ptr){
	static_cast<QOpenGLFramebufferObject*>(ptr)->~QOpenGLFramebufferObject();
}

QtObjectPtr QOpenGLFramebufferObject_NewQOpenGLFramebufferObject3(QtObjectPtr size, QtObjectPtr format){
	return new QOpenGLFramebufferObject(*static_cast<QSize*>(size), *static_cast<QOpenGLFramebufferObjectFormat*>(format));
}

QtObjectPtr QOpenGLFramebufferObject_NewQOpenGLFramebufferObject4(int width, int height, QtObjectPtr format){
	return new QOpenGLFramebufferObject(width, height, *static_cast<QOpenGLFramebufferObjectFormat*>(format));
}

int QOpenGLFramebufferObject_Attachment(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->attachment();
}

int QOpenGLFramebufferObject_Height(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->height();
}

int QOpenGLFramebufferObject_IsBound(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->isBound();
}

void QOpenGLFramebufferObject_SetAttachment(QtObjectPtr ptr, int attachment){
	static_cast<QOpenGLFramebufferObject*>(ptr)->setAttachment(static_cast<QOpenGLFramebufferObject::Attachment>(attachment));
}

int QOpenGLFramebufferObject_Width(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObject*>(ptr)->width();
}

