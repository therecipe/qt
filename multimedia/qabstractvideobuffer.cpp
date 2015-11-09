#include "qabstractvideobuffer.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractVideoBuffer>
#include "_cgo_export.h"

class MyQAbstractVideoBuffer: public QAbstractVideoBuffer {
public:
};

void* QAbstractVideoBuffer_Handle(void* ptr){
	return new QVariant(static_cast<QAbstractVideoBuffer*>(ptr)->handle());
}

int QAbstractVideoBuffer_HandleType(void* ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->handleType();
}

int QAbstractVideoBuffer_MapMode(void* ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->mapMode();
}

void QAbstractVideoBuffer_Release(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->release();
}

void QAbstractVideoBuffer_Unmap(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->unmap();
}

void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->~QAbstractVideoBuffer();
}

