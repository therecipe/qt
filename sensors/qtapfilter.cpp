#include "qtapfilter.h"
#include <QModelIndex>
#include <QTapReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTapFilter>
#include "_cgo_export.h"

class MyQTapFilter: public QTapFilter {
public:
};

int QTapFilter_Filter(void* ptr, void* reading){
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

