#include "qsensorgesture.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QObject>
#include <QString>
#include <QSensorGesture>
#include "_cgo_export.h"

class MyQSensorGesture: public QSensorGesture {
public:
};

QtObjectPtr QSensorGesture_NewQSensorGesture(char* ids, QtObjectPtr parent){
	return new QSensorGesture(QString(ids).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

char* QSensorGesture_GestureSignals(QtObjectPtr ptr){
	return static_cast<QSensorGesture*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGesture_InvalidIds(QtObjectPtr ptr){
	return static_cast<QSensorGesture*>(ptr)->invalidIds().join("|").toUtf8().data();
}

int QSensorGesture_IsActive(QtObjectPtr ptr){
	return static_cast<QSensorGesture*>(ptr)->isActive();
}

void QSensorGesture_StartDetection(QtObjectPtr ptr){
	static_cast<QSensorGesture*>(ptr)->startDetection();
}

void QSensorGesture_StopDetection(QtObjectPtr ptr){
	static_cast<QSensorGesture*>(ptr)->stopDetection();
}

char* QSensorGesture_ValidIds(QtObjectPtr ptr){
	return static_cast<QSensorGesture*>(ptr)->validIds().join("|").toUtf8().data();
}

void QSensorGesture_DestroyQSensorGesture(QtObjectPtr ptr){
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

