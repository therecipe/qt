#include "qcompassfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompassReading>
#include <QCompass>
#include <QCompassFilter>
#include "_cgo_export.h"

class MyQCompassFilter: public QCompassFilter {
public:
};

int QCompassFilter_Filter(void* ptr, void* reading){
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

