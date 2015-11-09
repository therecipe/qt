#include "qmagnetometerfilter.h"
#include <QMagnetometer>
#include <QMagnetometerReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMagnetometerFilter>
#include "_cgo_export.h"

class MyQMagnetometerFilter: public QMagnetometerFilter {
public:
};

int QMagnetometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

