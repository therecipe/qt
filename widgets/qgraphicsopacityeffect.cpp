#include "qgraphicsopacityeffect.h"
#include <QModelIndex>
#include <QBrush>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsOpacityEffect>
#include "_cgo_export.h"

class MyQGraphicsOpacityEffect: public QGraphicsOpacityEffect {
public:
void Signal_OpacityMaskChanged(const QBrush & mask){callbackQGraphicsOpacityEffectOpacityMaskChanged(this->objectName().toUtf8().data(), new QBrush(mask));};
};

double QGraphicsOpacityEffect_Opacity(void* ptr){
	return static_cast<double>(static_cast<QGraphicsOpacityEffect*>(ptr)->opacity());
}

void* QGraphicsOpacityEffect_OpacityMask(void* ptr){
	return new QBrush(static_cast<QGraphicsOpacityEffect*>(ptr)->opacityMask());
}

void QGraphicsOpacityEffect_SetOpacity(void* ptr, double opacity){
	QMetaObject::invokeMethod(static_cast<QGraphicsOpacityEffect*>(ptr), "setOpacity", Q_ARG(qreal, static_cast<qreal>(opacity)));
}

void QGraphicsOpacityEffect_SetOpacityMask(void* ptr, void* mask){
	QMetaObject::invokeMethod(static_cast<QGraphicsOpacityEffect*>(ptr), "setOpacityMask", Q_ARG(QBrush, *static_cast<QBrush*>(mask)));
}

void* QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(void* parent){
	return new QGraphicsOpacityEffect(static_cast<QObject*>(parent));
}

void QGraphicsOpacityEffect_ConnectOpacityMaskChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsOpacityEffect*>(ptr), static_cast<void (QGraphicsOpacityEffect::*)(const QBrush &)>(&QGraphicsOpacityEffect::opacityMaskChanged), static_cast<MyQGraphicsOpacityEffect*>(ptr), static_cast<void (MyQGraphicsOpacityEffect::*)(const QBrush &)>(&MyQGraphicsOpacityEffect::Signal_OpacityMaskChanged));;
}

void QGraphicsOpacityEffect_DisconnectOpacityMaskChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsOpacityEffect*>(ptr), static_cast<void (QGraphicsOpacityEffect::*)(const QBrush &)>(&QGraphicsOpacityEffect::opacityMaskChanged), static_cast<MyQGraphicsOpacityEffect*>(ptr), static_cast<void (MyQGraphicsOpacityEffect::*)(const QBrush &)>(&MyQGraphicsOpacityEffect::Signal_OpacityMaskChanged));;
}

void QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(void* ptr){
	static_cast<QGraphicsOpacityEffect*>(ptr)->~QGraphicsOpacityEffect();
}

