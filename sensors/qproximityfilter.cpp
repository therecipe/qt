#include "qproximityfilter.h"
#include <QModelIndex>
#include <QProximityReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QProximityFilter>
#include "_cgo_export.h"

class MyQProximityFilter: public QProximityFilter {
public:
};

int QProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

