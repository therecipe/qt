#include "qgraphicsrotation.h"
#include <QUrl>
#include <QModelIndex>
#include <QVector3D>
#include <QObject>
#include <QMatrix4x4>
#include <QVector>
#include <QString>
#include <QVariant>
#include <QGraphicsRotation>
#include "_cgo_export.h"

class MyQGraphicsRotation: public QGraphicsRotation {
public:
void Signal_AngleChanged(){callbackQGraphicsRotationAngleChanged(this->objectName().toUtf8().data());};
void Signal_AxisChanged(){callbackQGraphicsRotationAxisChanged(this->objectName().toUtf8().data());};
void Signal_OriginChanged(){callbackQGraphicsRotationOriginChanged(this->objectName().toUtf8().data());};
};

double QGraphicsRotation_Angle(void* ptr){
	return static_cast<double>(static_cast<QGraphicsRotation*>(ptr)->angle());
}

void QGraphicsRotation_SetAngle(void* ptr, double v){
	static_cast<QGraphicsRotation*>(ptr)->setAngle(static_cast<qreal>(v));
}

void QGraphicsRotation_SetAxis2(void* ptr, int axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(static_cast<Qt::Axis>(axis));
}

void QGraphicsRotation_SetAxis(void* ptr, void* axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(*static_cast<QVector3D*>(axis));
}

void QGraphicsRotation_SetOrigin(void* ptr, void* point){
	static_cast<QGraphicsRotation*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

void* QGraphicsRotation_NewQGraphicsRotation(void* parent){
	return new QGraphicsRotation(static_cast<QObject*>(parent));
}

void QGraphicsRotation_ConnectAngleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_DisconnectAngleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_ApplyTo(void* ptr, void* matrix){
	static_cast<QGraphicsRotation*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsRotation_ConnectAxisChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_DisconnectAxisChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_ConnectOriginChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DisconnectOriginChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DestroyQGraphicsRotation(void* ptr){
	static_cast<QGraphicsRotation*>(ptr)->~QGraphicsRotation();
}

