#include "qambienttemperaturefilter.h"
#include <QModelIndex>
#include <QAmbientTemperatureReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAmbientTemperatureFilter>
#include "_cgo_export.h"

class MyQAmbientTemperatureFilter: public QAmbientTemperatureFilter {
public:
};

int QAmbientTemperatureFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

