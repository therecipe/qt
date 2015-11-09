#include "qmagnetometerreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMagnetometer>
#include <QMagnetometerReading>
#include "_cgo_export.h"

class MyQMagnetometerReading: public QMagnetometerReading {
public:
};

double QMagnetometerReading_CalibrationLevel(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->calibrationLevel());
}

double QMagnetometerReading_X(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->x());
}

double QMagnetometerReading_Y(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->y());
}

double QMagnetometerReading_Z(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->z());
}

void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel){
	static_cast<QMagnetometerReading*>(ptr)->setCalibrationLevel(static_cast<qreal>(calibrationLevel));
}

void QMagnetometerReading_SetX(void* ptr, double x){
	static_cast<QMagnetometerReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QMagnetometerReading_SetY(void* ptr, double y){
	static_cast<QMagnetometerReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QMagnetometerReading_SetZ(void* ptr, double z){
	static_cast<QMagnetometerReading*>(ptr)->setZ(static_cast<qreal>(z));
}

