#include "qtapfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTapReading>
#include <QTapFilter>
#include "_cgo_export.h"

class MyQTapFilter: public QTapFilter {
public:
};

int QTapFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

