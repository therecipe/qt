#include "qabstractvideobuffer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractVideoBuffer>
#include "_cgo_export.h"

class MyQAbstractVideoBuffer: public QAbstractVideoBuffer {
public:
};

char* QAbstractVideoBuffer_Handle(QtObjectPtr ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->handle().toString().toUtf8().data();
}

int QAbstractVideoBuffer_HandleType(QtObjectPtr ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->handleType();
}

int QAbstractVideoBuffer_MapMode(QtObjectPtr ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->mapMode();
}

void QAbstractVideoBuffer_Release(QtObjectPtr ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->release();
}

void QAbstractVideoBuffer_Unmap(QtObjectPtr ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->unmap();
}

void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(QtObjectPtr ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->~QAbstractVideoBuffer();
}

