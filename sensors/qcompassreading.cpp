#include "qcompassreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompass>
#include <QCompassReading>
#include "_cgo_export.h"

class MyQCompassReading: public QCompassReading {
public:
};

double QCompassReading_Azimuth(void* ptr){
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->azimuth());
}

double QCompassReading_CalibrationLevel(void* ptr){
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->calibrationLevel());
}

void QCompassReading_SetAzimuth(void* ptr, double azimuth){
	static_cast<QCompassReading*>(ptr)->setAzimuth(static_cast<qreal>(azimuth));
}

void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel){
	static_cast<QCompassReading*>(ptr)->setCalibrationLevel(static_cast<qreal>(calibrationLevel));
}

