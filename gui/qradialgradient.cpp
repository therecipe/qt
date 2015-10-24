#include "qradialgradient.h"
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QRadialGradient>
#include "_cgo_export.h"

class MyQRadialGradient: public QRadialGradient {
public:
};

QtObjectPtr QRadialGradient_NewQRadialGradient(){
	return new QRadialGradient();
}

void QRadialGradient_SetCenter(QtObjectPtr ptr, QtObjectPtr center){
	static_cast<QRadialGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

void QRadialGradient_SetFocalPoint(QtObjectPtr ptr, QtObjectPtr focalPoint){
	static_cast<QRadialGradient*>(ptr)->setFocalPoint(*static_cast<QPointF*>(focalPoint));
}

