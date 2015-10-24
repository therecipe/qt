#include "qsensorreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QSensorReading>
#include "_cgo_export.h"

class MyQSensorReading: public QSensorReading {
public:
};

char* QSensorReading_Value(QtObjectPtr ptr, int index){
	return static_cast<QSensorReading*>(ptr)->value(index).toString().toUtf8().data();
}

int QSensorReading_ValueCount(QtObjectPtr ptr){
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

