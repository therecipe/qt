#include "qorientationfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOrientationReading>
#include <QOrientationFilter>
#include "_cgo_export.h"

class MyQOrientationFilter: public QOrientationFilter {
public:
};

int QOrientationFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QOrientationFilter*>(ptr)->filter(static_cast<QOrientationReading*>(reading));
}

