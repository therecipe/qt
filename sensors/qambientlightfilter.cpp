#include "qambientlightfilter.h"
#include <QAmbientLightReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightFilter>
#include "_cgo_export.h"

class MyQAmbientLightFilter: public QAmbientLightFilter {
public:
};

int QAmbientLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

