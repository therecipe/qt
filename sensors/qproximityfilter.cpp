#include "qproximityfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProximityReading>
#include <QProximityFilter>
#include "_cgo_export.h"

class MyQProximityFilter: public QProximityFilter {
public:
};

int QProximityFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

