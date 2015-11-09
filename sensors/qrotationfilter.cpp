#include "qrotationfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRotationReading>
#include <QRotationFilter>
#include "_cgo_export.h"

class MyQRotationFilter: public QRotationFilter {
public:
};

int QRotationFilter_Filter(void* ptr, void* reading){
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

