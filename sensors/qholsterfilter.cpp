#include "qholsterfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QHolsterReading>
#include <QString>
#include <QVariant>
#include <QHolsterFilter>
#include "_cgo_export.h"

class MyQHolsterFilter: public QHolsterFilter {
public:
};

int QHolsterFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

