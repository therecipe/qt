#include "qlightfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLightReading>
#include <QLightFilter>
#include "_cgo_export.h"

class MyQLightFilter: public QLightFilter {
public:
};

int QLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

