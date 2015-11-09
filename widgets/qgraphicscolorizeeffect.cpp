#include "qgraphicscolorizeeffect.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QColor>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsColorizeEffect>
#include "_cgo_export.h"

class MyQGraphicsColorizeEffect: public QGraphicsColorizeEffect {
public:
void Signal_ColorChanged(const QColor & color){callbackQGraphicsColorizeEffectColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

void* QGraphicsColorizeEffect_Color(void* ptr){
	return new QColor(static_cast<QGraphicsColorizeEffect*>(ptr)->color());
}

void QGraphicsColorizeEffect_SetColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QGraphicsColorizeEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QGraphicsColorizeEffect_SetStrength(void* ptr, double strength){
	QMetaObject::invokeMethod(static_cast<QGraphicsColorizeEffect*>(ptr), "setStrength", Q_ARG(qreal, static_cast<qreal>(strength)));
}

double QGraphicsColorizeEffect_Strength(void* ptr){
	return static_cast<double>(static_cast<QGraphicsColorizeEffect*>(ptr)->strength());
}

void* QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(void* parent){
	return new QGraphicsColorizeEffect(static_cast<QObject*>(parent));
}

void QGraphicsColorizeEffect_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsColorizeEffect*>(ptr), static_cast<void (QGraphicsColorizeEffect::*)(const QColor &)>(&QGraphicsColorizeEffect::colorChanged), static_cast<MyQGraphicsColorizeEffect*>(ptr), static_cast<void (MyQGraphicsColorizeEffect::*)(const QColor &)>(&MyQGraphicsColorizeEffect::Signal_ColorChanged));;
}

void QGraphicsColorizeEffect_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsColorizeEffect*>(ptr), static_cast<void (QGraphicsColorizeEffect::*)(const QColor &)>(&QGraphicsColorizeEffect::colorChanged), static_cast<MyQGraphicsColorizeEffect*>(ptr), static_cast<void (MyQGraphicsColorizeEffect::*)(const QColor &)>(&MyQGraphicsColorizeEffect::Signal_ColorChanged));;
}

void QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(void* ptr){
	static_cast<QGraphicsColorizeEffect*>(ptr)->~QGraphicsColorizeEffect();
}

