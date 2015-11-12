#include "qsensorgesturerecognizer.h"
#include <QSensorGesture>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorGestureRecognizer>
#include "_cgo_export.h"

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer {
public:
};

void QSensorGestureRecognizer_CreateBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->createBackend();
}

char* QSensorGestureRecognizer_GestureSignals(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGestureRecognizer_Id(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->id().toUtf8().data();
}

int QSensorGestureRecognizer_IsActive(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->isActive();
}

void QSensorGestureRecognizer_StartBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->startBackend();
}

void QSensorGestureRecognizer_StopBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->stopBackend();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->~QSensorGestureRecognizer();
}

