#include "qaccelerometerfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometerReading>
#include <QAccelerometer>
#include <QAccelerometerFilter>
#include "_cgo_export.h"

class MyQAccelerometerFilter: public QAccelerometerFilter {
public:
};

int QAccelerometerFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

