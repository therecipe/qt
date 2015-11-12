#include "qambienttemperaturefilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientTemperatureReading>
#include <QAmbientTemperatureFilter>
#include "_cgo_export.h"

class MyQAmbientTemperatureFilter: public QAmbientTemperatureFilter {
public:
};

int QAmbientTemperatureFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

