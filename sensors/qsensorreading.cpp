#include "qsensorreading.h"
#include <QModelIndex>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorReading>
#include "_cgo_export.h"

class MyQSensorReading: public QSensorReading {
public:
};

void* QSensorReading_Value(void* ptr, int index){
	return new QVariant(static_cast<QSensorReading*>(ptr)->value(index));
}

int QSensorReading_ValueCount(void* ptr){
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

