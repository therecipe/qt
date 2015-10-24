#include "qtiltfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTiltReading>
#include <QTiltFilter>
#include "_cgo_export.h"

class MyQTiltFilter: public QTiltFilter {
public:
};

int QTiltFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

