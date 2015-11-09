#include "qpressurefilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPressureReading>
#include <QPressureFilter>
#include "_cgo_export.h"

class MyQPressureFilter: public QPressureFilter {
public:
};

int QPressureFilter_Filter(void* ptr, void* reading){
	return static_cast<QPressureFilter*>(ptr)->filter(static_cast<QPressureReading*>(reading));
}

