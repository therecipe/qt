#include "qambientlightreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightReading>
#include "_cgo_export.h"

class MyQAmbientLightReading: public QAmbientLightReading {
public:
};

int QAmbientLightReading_LightLevel(QtObjectPtr ptr){
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

void QAmbientLightReading_SetLightLevel(QtObjectPtr ptr, int lightLevel){
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

