#include "qsensorgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QSensor>
#include <QSensorGesture>
#include "_cgo_export.h"

class MyQSensorGesture: public QSensorGesture {
public:
};

void* QSensorGesture_NewQSensorGesture(char* ids, void* parent){
	return new QSensorGesture(QString(ids).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

char* QSensorGesture_GestureSignals(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGesture_InvalidIds(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->invalidIds().join("|").toUtf8().data();
}

int QSensorGesture_IsActive(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->isActive();
}

void QSensorGesture_StartDetection(void* ptr){
	static_cast<QSensorGesture*>(ptr)->startDetection();
}

void QSensorGesture_StopDetection(void* ptr){
	static_cast<QSensorGesture*>(ptr)->stopDetection();
}

char* QSensorGesture_ValidIds(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->validIds().join("|").toUtf8().data();
}

void QSensorGesture_DestroyQSensorGesture(void* ptr){
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

