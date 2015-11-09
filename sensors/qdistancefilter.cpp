#include "qdistancefilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDistanceReading>
#include <QDistanceFilter>
#include "_cgo_export.h"

class MyQDistanceFilter: public QDistanceFilter {
public:
};

int QDistanceFilter_Filter(void* ptr, void* reading){
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

