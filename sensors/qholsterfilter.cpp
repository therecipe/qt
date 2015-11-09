#include "qholsterfilter.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHolsterReading>
#include <QString>
#include <QHolsterFilter>
#include "_cgo_export.h"

class MyQHolsterFilter: public QHolsterFilter {
public:
};

int QHolsterFilter_Filter(void* ptr, void* reading){
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

