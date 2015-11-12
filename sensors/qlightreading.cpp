#include "qlightreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QLightReading>
#include "_cgo_export.h"

class MyQLightReading: public QLightReading {
public:
};

double QLightReading_Lux(void* ptr){
	return static_cast<double>(static_cast<QLightReading*>(ptr)->lux());
}

void QLightReading_SetLux(void* ptr, double lux){
	static_cast<QLightReading*>(ptr)->setLux(static_cast<qreal>(lux));
}

