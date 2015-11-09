#include "qaltimeterfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAltimeterReading>
#include <QAltimeter>
#include <QAltimeterFilter>
#include "_cgo_export.h"

class MyQAltimeterFilter: public QAltimeterFilter {
public:
};

int QAltimeterFilter_Filter(void* ptr, void* reading){
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

