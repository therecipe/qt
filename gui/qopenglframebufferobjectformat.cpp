#include "qopenglframebufferobjectformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFramebufferObject>
#include <QOpenGLFramebufferObjectFormat>
#include "_cgo_export.h"

class MyQOpenGLFramebufferObjectFormat: public QOpenGLFramebufferObjectFormat {
public:
};

QtObjectPtr QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat(){
	return new QOpenGLFramebufferObjectFormat();
}

QtObjectPtr QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat2(QtObjectPtr other){
	return new QOpenGLFramebufferObjectFormat(*static_cast<QOpenGLFramebufferObjectFormat*>(other));
}

int QOpenGLFramebufferObjectFormat_Attachment(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->attachment();
}

int QOpenGLFramebufferObjectFormat_Mipmap(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->mipmap();
}

int QOpenGLFramebufferObjectFormat_Samples(QtObjectPtr ptr){
	return static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->samples();
}

void QOpenGLFramebufferObjectFormat_SetAttachment(QtObjectPtr ptr, int attachment){
	static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->setAttachment(static_cast<QOpenGLFramebufferObject::Attachment>(attachment));
}

void QOpenGLFramebufferObjectFormat_SetMipmap(QtObjectPtr ptr, int enabled){
	static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->setMipmap(enabled != 0);
}

void QOpenGLFramebufferObjectFormat_SetSamples(QtObjectPtr ptr, int samples){
	static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->setSamples(samples);
}

void QOpenGLFramebufferObjectFormat_DestroyQOpenGLFramebufferObjectFormat(QtObjectPtr ptr){
	static_cast<QOpenGLFramebufferObjectFormat*>(ptr)->~QOpenGLFramebufferObjectFormat();
}

