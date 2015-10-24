#include "qlightfilter.h"
#include <QLightReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLightFilter>
#include "_cgo_export.h"

class MyQLightFilter: public QLightFilter {
public:
};

int QLightFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

