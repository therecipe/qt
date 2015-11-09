#include "qrotationreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRotationReading>
#include "_cgo_export.h"

class MyQRotationReading: public QRotationReading {
public:
};

double QRotationReading_X(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->x());
}

double QRotationReading_Y(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->y());
}

double QRotationReading_Z(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->z());
}

void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z){
	static_cast<QRotationReading*>(ptr)->setFromEuler(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(z));
}

