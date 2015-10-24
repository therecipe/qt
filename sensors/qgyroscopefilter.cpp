#include "qgyroscopefilter.h"
#include <QGyroscope>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGyroscopeReading>
#include <QGyroscopeFilter>
#include "_cgo_export.h"

class MyQGyroscopeFilter: public QGyroscopeFilter {
public:
};

int QGyroscopeFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

