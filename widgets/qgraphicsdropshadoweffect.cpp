#include "qgraphicsdropshadoweffect.h"
#include <QString>
#include <QPointF>
#include <QColor>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QGraphicsDropShadowEffect>
#include "_cgo_export.h"

class MyQGraphicsDropShadowEffect: public QGraphicsDropShadowEffect {
public:
void Signal_ColorChanged(const QColor & color){callbackQGraphicsDropShadowEffectColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

double QGraphicsDropShadowEffect_BlurRadius(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->blurRadius());
}

void* QGraphicsDropShadowEffect_Color(void* ptr){
	return new QColor(static_cast<QGraphicsDropShadowEffect*>(ptr)->color());
}

void QGraphicsDropShadowEffect_SetBlurRadius(void* ptr, double blurRadius){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setBlurRadius", Q_ARG(qreal, static_cast<qreal>(blurRadius)));
}

void QGraphicsDropShadowEffect_SetColor(void* ptr, void* color){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QGraphicsDropShadowEffect_SetOffset(void* ptr, void* ofs){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(QPointF, *static_cast<QPointF*>(ofs)));
}

void* QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(void* parent){
	return new QGraphicsDropShadowEffect(static_cast<QObject*>(parent));
}

void QGraphicsDropShadowEffect_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsDropShadowEffect*>(ptr), static_cast<void (QGraphicsDropShadowEffect::*)(const QColor &)>(&QGraphicsDropShadowEffect::colorChanged), static_cast<MyQGraphicsDropShadowEffect*>(ptr), static_cast<void (MyQGraphicsDropShadowEffect::*)(const QColor &)>(&MyQGraphicsDropShadowEffect::Signal_ColorChanged));;
}

void QGraphicsDropShadowEffect_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsDropShadowEffect*>(ptr), static_cast<void (QGraphicsDropShadowEffect::*)(const QColor &)>(&QGraphicsDropShadowEffect::colorChanged), static_cast<MyQGraphicsDropShadowEffect*>(ptr), static_cast<void (MyQGraphicsDropShadowEffect::*)(const QColor &)>(&MyQGraphicsDropShadowEffect::Signal_ColorChanged));;
}

void QGraphicsDropShadowEffect_SetOffset3(void* ptr, double d){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(qreal, static_cast<qreal>(d)));
}

void QGraphicsDropShadowEffect_SetOffset2(void* ptr, double dx, double dy){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(qreal, static_cast<qreal>(dx)), Q_ARG(qreal, static_cast<qreal>(dy)));
}

void QGraphicsDropShadowEffect_SetXOffset(void* ptr, double dx){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setXOffset", Q_ARG(qreal, static_cast<qreal>(dx)));
}

void QGraphicsDropShadowEffect_SetYOffset(void* ptr, double dy){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setYOffset", Q_ARG(qreal, static_cast<qreal>(dy)));
}

double QGraphicsDropShadowEffect_XOffset(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->xOffset());
}

double QGraphicsDropShadowEffect_YOffset(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->yOffset());
}

void QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(void* ptr){
	static_cast<QGraphicsDropShadowEffect*>(ptr)->~QGraphicsDropShadowEffect();
}

