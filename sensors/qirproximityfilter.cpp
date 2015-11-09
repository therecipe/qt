#include "qirproximityfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIRProximityReading>
#include <QIRProximityFilter>
#include "_cgo_export.h"

class MyQIRProximityFilter: public QIRProximityFilter {
public:
};

int QIRProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

