#include "qaltimeterfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QAltimeter>
#include <QAltimeterReading>
#include <QString>
#include <QVariant>
#include <QAltimeterFilter>
#include "_cgo_export.h"

class MyQAltimeterFilter: public QAltimeterFilter {
public:
};

int QAltimeterFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

