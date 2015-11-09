#include "qirproximityreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIRProximityReading>
#include "_cgo_export.h"

class MyQIRProximityReading: public QIRProximityReading {
public:
};

double QIRProximityReading_Reflectance(void* ptr){
	return static_cast<double>(static_cast<QIRProximityReading*>(ptr)->reflectance());
}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance){
	static_cast<QIRProximityReading*>(ptr)->setReflectance(static_cast<qreal>(reflectance));
}

