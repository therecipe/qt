#include "qgyroscopefilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGyroscope>
#include <QGyroscopeReading>
#include <QGyroscopeFilter>
#include "_cgo_export.h"

class MyQGyroscopeFilter: public QGyroscopeFilter {
public:
};

int QGyroscopeFilter_Filter(void* ptr, void* reading){
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

