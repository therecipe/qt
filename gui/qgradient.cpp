#include "qgradient.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGradient>
#include "_cgo_export.h"

class MyQGradient: public QGradient {
public:
};

int QGradient_CoordinateMode(QtObjectPtr ptr){
	return static_cast<QGradient*>(ptr)->coordinateMode();
}

void QGradient_SetCoordinateMode(QtObjectPtr ptr, int mode){
	static_cast<QGradient*>(ptr)->setCoordinateMode(static_cast<QGradient::CoordinateMode>(mode));
}

void QGradient_SetSpread(QtObjectPtr ptr, int method){
	static_cast<QGradient*>(ptr)->setSpread(static_cast<QGradient::Spread>(method));
}

int QGradient_Spread(QtObjectPtr ptr){
	return static_cast<QGradient*>(ptr)->spread();
}

int QGradient_Type(QtObjectPtr ptr){
	return static_cast<QGradient*>(ptr)->type();
}

