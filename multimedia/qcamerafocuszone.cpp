#include "qcamerafocuszone.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraFocus>
#include <QString>
#include <QCameraFocusZone>
#include "_cgo_export.h"

class MyQCameraFocusZone: public QCameraFocusZone {
public:
};

void* QCameraFocusZone_NewQCameraFocusZone(void* other){
	return new QCameraFocusZone(*static_cast<QCameraFocusZone*>(other));
}

int QCameraFocusZone_IsValid(void* ptr){
	return static_cast<QCameraFocusZone*>(ptr)->isValid();
}

int QCameraFocusZone_Status(void* ptr){
	return static_cast<QCameraFocusZone*>(ptr)->status();
}

void QCameraFocusZone_DestroyQCameraFocusZone(void* ptr){
	static_cast<QCameraFocusZone*>(ptr)->~QCameraFocusZone();
}

