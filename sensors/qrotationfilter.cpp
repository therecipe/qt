#include "qrotationfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QRotationReading>
#include <QString>
#include <QVariant>
#include <QRotationFilter>
#include "_cgo_export.h"

class MyQRotationFilter: public QRotationFilter {
public:
};

int QRotationFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

