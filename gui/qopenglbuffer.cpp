#include "qopenglbuffer.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLBuffer>
#include "_cgo_export.h"

class MyQOpenGLBuffer: public QOpenGLBuffer {
public:
};

QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer(){
	return new QOpenGLBuffer();
}

QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer2(int ty){
	return new QOpenGLBuffer(static_cast<QOpenGLBuffer::Type>(ty));
}

QtObjectPtr QOpenGLBuffer_NewQOpenGLBuffer3(QtObjectPtr other){
	return new QOpenGLBuffer(*static_cast<QOpenGLBuffer*>(other));
}

void QOpenGLBuffer_Allocate2(QtObjectPtr ptr, int count){
	static_cast<QOpenGLBuffer*>(ptr)->allocate(count);
}

int QOpenGLBuffer_Bind(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->bind();
}

int QOpenGLBuffer_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->create();
}

void QOpenGLBuffer_Destroy(QtObjectPtr ptr){
	static_cast<QOpenGLBuffer*>(ptr)->destroy();
}

int QOpenGLBuffer_IsCreated(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->isCreated();
}

void QOpenGLBuffer_Map(QtObjectPtr ptr, int access){
	static_cast<QOpenGLBuffer*>(ptr)->map(static_cast<QOpenGLBuffer::Access>(access));
}

void QOpenGLBuffer_MapRange(QtObjectPtr ptr, int offset, int count, int access){
	static_cast<QOpenGLBuffer*>(ptr)->mapRange(offset, count, static_cast<QOpenGLBuffer::RangeAccessFlag>(access));
}

void QOpenGLBuffer_Release(QtObjectPtr ptr){
	static_cast<QOpenGLBuffer*>(ptr)->release();
}

void QOpenGLBuffer_QOpenGLBuffer_Release2(int ty){
	QOpenGLBuffer::release(static_cast<QOpenGLBuffer::Type>(ty));
}

void QOpenGLBuffer_SetUsagePattern(QtObjectPtr ptr, int value){
	static_cast<QOpenGLBuffer*>(ptr)->setUsagePattern(static_cast<QOpenGLBuffer::UsagePattern>(value));
}

int QOpenGLBuffer_Size(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->size();
}

int QOpenGLBuffer_Type(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->type();
}

int QOpenGLBuffer_Unmap(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->unmap();
}

int QOpenGLBuffer_UsagePattern(QtObjectPtr ptr){
	return static_cast<QOpenGLBuffer*>(ptr)->usagePattern();
}

void QOpenGLBuffer_DestroyQOpenGLBuffer(QtObjectPtr ptr){
	static_cast<QOpenGLBuffer*>(ptr)->~QOpenGLBuffer();
}

