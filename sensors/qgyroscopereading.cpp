#include "qgyroscopereading.h"
#include <QModelIndex>
#include <QGyroscope>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGyroscopeReading>
#include "_cgo_export.h"

class MyQGyroscopeReading: public QGyroscopeReading {
public:
};

double QGyroscopeReading_X(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->x());
}

double QGyroscopeReading_Y(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->y());
}

double QGyroscopeReading_Z(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->z());
}

void QGyroscopeReading_SetX(void* ptr, double x){
	static_cast<QGyroscopeReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QGyroscopeReading_SetY(void* ptr, double y){
	static_cast<QGyroscopeReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QGyroscopeReading_SetZ(void* ptr, double z){
	static_cast<QGyroscopeReading*>(ptr)->setZ(static_cast<qreal>(z));
}

