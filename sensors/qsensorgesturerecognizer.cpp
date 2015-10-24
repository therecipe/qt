#include "qsensorgesturerecognizer.h"
#include <QModelIndex>
#include <QSensor>
#include <QSensorGesture>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorGestureRecognizer>
#include "_cgo_export.h"

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer {
public:
};

void QSensorGestureRecognizer_CreateBackend(QtObjectPtr ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->createBackend();
}

char* QSensorGestureRecognizer_GestureSignals(QtObjectPtr ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGestureRecognizer_Id(QtObjectPtr ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->id().toUtf8().data();
}

int QSensorGestureRecognizer_IsActive(QtObjectPtr ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->isActive();
}

void QSensorGestureRecognizer_StartBackend(QtObjectPtr ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->startBackend();
}

void QSensorGestureRecognizer_StopBackend(QtObjectPtr ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->stopBackend();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(QtObjectPtr ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->~QSensorGestureRecognizer();
}

