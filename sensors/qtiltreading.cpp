#include "qtiltreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTiltReading>
#include "_cgo_export.h"

class MyQTiltReading: public QTiltReading {
public:
};

double QTiltReading_XRotation(void* ptr){
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->xRotation());
}

double QTiltReading_YRotation(void* ptr){
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->yRotation());
}

void QTiltReading_SetXRotation(void* ptr, double x){
	static_cast<QTiltReading*>(ptr)->setXRotation(static_cast<qreal>(x));
}

void QTiltReading_SetYRotation(void* ptr, double y){
	static_cast<QTiltReading*>(ptr)->setYRotation(static_cast<qreal>(y));
}

