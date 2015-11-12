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

int QAmbientLightReading_LightLevel(void* ptr){
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

void QAmbientLightReading_SetLightLevel(void* ptr, int lightLevel){
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

