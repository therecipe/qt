#include "qgraphicsscale.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMatrix4x4>
#include <QVector>
#include <QVector3D>
#include <QObject>
#include <QGraphicsScale>
#include "_cgo_export.h"

class MyQGraphicsScale: public QGraphicsScale {
public:
void Signal_OriginChanged(){callbackQGraphicsScaleOriginChanged(this->objectName().toUtf8().data());};
void Signal_ScaleChanged(){callbackQGraphicsScaleScaleChanged(this->objectName().toUtf8().data());};
void Signal_XScaleChanged(){callbackQGraphicsScaleXScaleChanged(this->objectName().toUtf8().data());};
void Signal_YScaleChanged(){callbackQGraphicsScaleYScaleChanged(this->objectName().toUtf8().data());};
void Signal_ZScaleChanged(){callbackQGraphicsScaleZScaleChanged(this->objectName().toUtf8().data());};
};

void QGraphicsScale_SetOrigin(void* ptr, void* point){
	static_cast<QGraphicsScale*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

void QGraphicsScale_SetXScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setXScale(static_cast<qreal>(v));
}

void QGraphicsScale_SetYScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setYScale(static_cast<qreal>(v));
}

void QGraphicsScale_SetZScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setZScale(static_cast<qreal>(v));
}

double QGraphicsScale_XScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->xScale());
}

double QGraphicsScale_YScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->yScale());
}

double QGraphicsScale_ZScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->zScale());
}

void* QGraphicsScale_NewQGraphicsScale(void* parent){
	return new QGraphicsScale(static_cast<QObject*>(parent));
}

void QGraphicsScale_ApplyTo(void* ptr, void* matrix){
	static_cast<QGraphicsScale*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsScale_ConnectOriginChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_DisconnectOriginChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_ConnectScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_DisconnectScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_ConnectXScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_DisconnectXScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_ConnectYScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_DisconnectYScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_ConnectZScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DisconnectZScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DestroyQGraphicsScale(void* ptr){
	static_cast<QGraphicsScale*>(ptr)->~QGraphicsScale();
}

