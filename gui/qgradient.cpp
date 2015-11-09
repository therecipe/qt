#include "qgradient.h"
#include <QModelIndex>
#include <QColor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGradient>
#include "_cgo_export.h"

class MyQGradient: public QGradient {
public:
};

void QGradient_SetColorAt(void* ptr, double position, void* color){
	static_cast<QGradient*>(ptr)->setColorAt(static_cast<qreal>(position), *static_cast<QColor*>(color));
}

int QGradient_CoordinateMode(void* ptr){
	return static_cast<QGradient*>(ptr)->coordinateMode();
}

void QGradient_SetCoordinateMode(void* ptr, int mode){
	static_cast<QGradient*>(ptr)->setCoordinateMode(static_cast<QGradient::CoordinateMode>(mode));
}

void QGradient_SetSpread(void* ptr, int method){
	static_cast<QGradient*>(ptr)->setSpread(static_cast<QGradient::Spread>(method));
}

int QGradient_Spread(void* ptr){
	return static_cast<QGradient*>(ptr)->spread();
}

int QGradient_Type(void* ptr){
	return static_cast<QGradient*>(ptr)->type();
}

