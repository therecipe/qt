#include "qcameraviewfinder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QCamera>
#include <QCameraViewfinder>
#include "_cgo_export.h"

class MyQCameraViewfinder: public QCameraViewfinder {
public:
};

QtObjectPtr QCameraViewfinder_NewQCameraViewfinder(QtObjectPtr parent){
	return new QCameraViewfinder(static_cast<QWidget*>(parent));
}

QtObjectPtr QCameraViewfinder_MediaObject(QtObjectPtr ptr){
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

void QCameraViewfinder_DestroyQCameraViewfinder(QtObjectPtr ptr){
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

