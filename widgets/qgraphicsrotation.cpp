#include "qgraphicsrotation.h"
#include <QVector3D>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMatrix4x4>
#include <QVector>
#include <QGraphicsRotation>
#include "_cgo_export.h"

class MyQGraphicsRotation: public QGraphicsRotation {
public:
void Signal_AngleChanged(){callbackQGraphicsRotationAngleChanged(this->objectName().toUtf8().data());};
void Signal_AxisChanged(){callbackQGraphicsRotationAxisChanged(this->objectName().toUtf8().data());};
void Signal_OriginChanged(){callbackQGraphicsRotationOriginChanged(this->objectName().toUtf8().data());};
};

void QGraphicsRotation_SetAxis2(QtObjectPtr ptr, int axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(static_cast<Qt::Axis>(axis));
}

void QGraphicsRotation_SetAxis(QtObjectPtr ptr, QtObjectPtr axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(*static_cast<QVector3D*>(axis));
}

void QGraphicsRotation_SetOrigin(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QGraphicsRotation*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

QtObjectPtr QGraphicsRotation_NewQGraphicsRotation(QtObjectPtr parent){
	return new QGraphicsRotation(static_cast<QObject*>(parent));
}

void QGraphicsRotation_ConnectAngleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_DisconnectAngleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_ApplyTo(QtObjectPtr ptr, QtObjectPtr matrix){
	static_cast<QGraphicsRotation*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsRotation_ConnectAxisChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_DisconnectAxisChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_ConnectOriginChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DisconnectOriginChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DestroyQGraphicsRotation(QtObjectPtr ptr){
	static_cast<QGraphicsRotation*>(ptr)->~QGraphicsRotation();
}

