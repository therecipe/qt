#include "qambienttemperaturereading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QAmbientTemperatureReading>
#include "_cgo_export.h"

class MyQAmbientTemperatureReading: public QAmbientTemperatureReading {
public:
};

double QAmbientTemperatureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QAmbientTemperatureReading*>(ptr)->temperature());
}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(static_cast<qreal>(temperature));
}

