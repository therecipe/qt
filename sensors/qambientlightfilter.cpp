#include "qambientlightfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightReading>
#include <QAmbientLightFilter>
#include "_cgo_export.h"

class MyQAmbientLightFilter: public QAmbientLightFilter {
public:
};

int QAmbientLightFilter_Filter(QtObjectPtr ptr, QtObjectPtr reading){
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

