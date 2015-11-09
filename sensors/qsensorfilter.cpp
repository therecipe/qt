#include "qsensorfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorReading>
#include <QString>
#include <QVariant>
#include <QSensorFilter>
#include "_cgo_export.h"

class MyQSensorFilter: public QSensorFilter {
public:
};

int QSensorFilter_Filter(void* ptr, void* reading){
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

