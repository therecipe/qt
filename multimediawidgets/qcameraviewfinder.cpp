#include "qcameraviewfinder.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QCameraViewfinder>
#include "_cgo_export.h"

class MyQCameraViewfinder: public QCameraViewfinder {
public:
};

void* QCameraViewfinder_NewQCameraViewfinder(void* parent){
	return new QCameraViewfinder(static_cast<QWidget*>(parent));
}

void* QCameraViewfinder_MediaObject(void* ptr){
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr){
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

