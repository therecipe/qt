#include "qpressurereading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPressureReading>
#include "_cgo_export.h"

class MyQPressureReading: public QPressureReading {
public:
};

double QPressureReading_Pressure(void* ptr){
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->pressure());
}

double QPressureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->temperature());
}

void QPressureReading_SetPressure(void* ptr, double pressure){
	static_cast<QPressureReading*>(ptr)->setPressure(static_cast<qreal>(pressure));
}

void QPressureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QPressureReading*>(ptr)->setTemperature(static_cast<qreal>(temperature));
}

