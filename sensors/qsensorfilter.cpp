#include "qsensorfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QSensorReading>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QSensorFilter>
#include "_cgo_export.h"

class MyQSensorFilter: public QSensorFilter {
public:
};

int QSensorFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

