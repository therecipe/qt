#include "qirproximityfilter.h"
#include <QModelIndex>
#include <QIRProximityReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QIRProximityFilter>
#include "_cgo_export.h"

class MyQIRProximityFilter: public QIRProximityFilter {
public:
};

int QIRProximityFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

