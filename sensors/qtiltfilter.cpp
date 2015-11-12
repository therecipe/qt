#include "qtiltfilter.h"
#include <QTiltReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTiltFilter>
#include "_cgo_export.h"

class MyQTiltFilter: public QTiltFilter {
public:
};

int QTiltFilter_Filter(void* ptr, void* reading){
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

