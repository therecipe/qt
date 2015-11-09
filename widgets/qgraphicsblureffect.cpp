#include "qgraphicsblureffect.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QGraphicsBlurEffect>
#include "_cgo_export.h"

class MyQGraphicsBlurEffect: public QGraphicsBlurEffect {
public:
void Signal_BlurHintsChanged(QGraphicsBlurEffect::BlurHints hints){callbackQGraphicsBlurEffectBlurHintsChanged(this->objectName().toUtf8().data(), hints);};
};

int QGraphicsBlurEffect_BlurHints(void* ptr){
	return static_cast<QGraphicsBlurEffect*>(ptr)->blurHints();
}

double QGraphicsBlurEffect_BlurRadius(void* ptr){
	return static_cast<double>(static_cast<QGraphicsBlurEffect*>(ptr)->blurRadius());
}

void QGraphicsBlurEffect_SetBlurHints(void* ptr, int hints){
	QMetaObject::invokeMethod(static_cast<QGraphicsBlurEffect*>(ptr), "setBlurHints", Q_ARG(QGraphicsBlurEffect::BlurHint, static_cast<QGraphicsBlurEffect::BlurHint>(hints)));
}

void QGraphicsBlurEffect_SetBlurRadius(void* ptr, double blurRadius){
	QMetaObject::invokeMethod(static_cast<QGraphicsBlurEffect*>(ptr), "setBlurRadius", Q_ARG(qreal, static_cast<qreal>(blurRadius)));
}

void* QGraphicsBlurEffect_NewQGraphicsBlurEffect(void* parent){
	return new QGraphicsBlurEffect(static_cast<QObject*>(parent));
}

void QGraphicsBlurEffect_ConnectBlurHintsChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DisconnectBlurHintsChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(void* ptr){
	static_cast<QGraphicsBlurEffect*>(ptr)->~QGraphicsBlurEffect();
}

