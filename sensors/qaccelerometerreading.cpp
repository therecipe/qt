#include "qaccelerometerreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometer>
#include <QAccelerometerReading>
#include "_cgo_export.h"

class MyQAccelerometerReading: public QAccelerometerReading {
public:
};

double QAccelerometerReading_X(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->x());
}

double QAccelerometerReading_Y(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->y());
}

double QAccelerometerReading_Z(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->z());
}

void QAccelerometerReading_SetX(void* ptr, double x){
	static_cast<QAccelerometerReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QAccelerometerReading_SetY(void* ptr, double y){
	static_cast<QAccelerometerReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QAccelerometerReading_SetZ(void* ptr, double z){
	static_cast<QAccelerometerReading*>(ptr)->setZ(static_cast<qreal>(z));
}

