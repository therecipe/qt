#include "qgraphicsscale.h"
#include <QVector>
#include <QVector3D>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMatrix4x4>
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

void QGraphicsScale_SetOrigin(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QGraphicsScale*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

QtObjectPtr QGraphicsScale_NewQGraphicsScale(QtObjectPtr parent){
	return new QGraphicsScale(static_cast<QObject*>(parent));
}

void QGraphicsScale_ApplyTo(QtObjectPtr ptr, QtObjectPtr matrix){
	static_cast<QGraphicsScale*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsScale_ConnectOriginChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_DisconnectOriginChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_ConnectScaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_DisconnectScaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_ConnectXScaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_DisconnectXScaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_ConnectYScaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_DisconnectYScaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_ConnectZScaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DisconnectZScaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DestroyQGraphicsScale(QtObjectPtr ptr){
	static_cast<QGraphicsScale*>(ptr)->~QGraphicsScale();
}

