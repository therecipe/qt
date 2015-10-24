#include "qcamerafocuszone.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraFocus>
#include <QString>
#include <QVariant>
#include <QCameraFocusZone>
#include "_cgo_export.h"

class MyQCameraFocusZone: public QCameraFocusZone {
public:
};

QtObjectPtr QCameraFocusZone_NewQCameraFocusZone(QtObjectPtr other){
	return new QCameraFocusZone(*static_cast<QCameraFocusZone*>(other));
}

int QCameraFocusZone_IsValid(QtObjectPtr ptr){
	return static_cast<QCameraFocusZone*>(ptr)->isValid();
}

int QCameraFocusZone_Status(QtObjectPtr ptr){
	return static_cast<QCameraFocusZone*>(ptr)->status();
}

void QCameraFocusZone_DestroyQCameraFocusZone(QtObjectPtr ptr){
	static_cast<QCameraFocusZone*>(ptr)->~QCameraFocusZone();
}

