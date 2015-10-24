#include "qcompassfilter.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompassReading>
#include <QCompass>
#include <QString>
#include <QCompassFilter>
#include "_cgo_export.h"

class MyQCompassFilter: public QCompassFilter {
public:
};

int QCompassFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

