#include "qmagnetometerfilter.h"
#include <QMagnetometerReading>
#include <QMagnetometer>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMagnetometerFilter>
#include "_cgo_export.h"

class MyQMagnetometerFilter: public QMagnetometerFilter {
public:
};

int QMagnetometerFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

