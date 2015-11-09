#include "qaccelerometerfilter.h"
#include <QAccelerometer>
#include <QAccelerometerReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometerFilter>
#include "_cgo_export.h"

class MyQAccelerometerFilter: public QAccelerometerFilter {
public:
};

int QAccelerometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

