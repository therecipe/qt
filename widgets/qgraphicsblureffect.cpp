#include "qgraphicsblureffect.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QGraphicsBlurEffect>
#include "_cgo_export.h"

class MyQGraphicsBlurEffect: public QGraphicsBlurEffect {
public:
void Signal_BlurHintsChanged(QGraphicsBlurEffect::BlurHints hints){callbackQGraphicsBlurEffectBlurHintsChanged(this->objectName().toUtf8().data(), hints);};
};

int QGraphicsBlurEffect_BlurHints(QtObjectPtr ptr){
	return static_cast<QGraphicsBlurEffect*>(ptr)->blurHints();
}

void QGraphicsBlurEffect_SetBlurHints(QtObjectPtr ptr, int hints){
	QMetaObject::invokeMethod(static_cast<QGraphicsBlurEffect*>(ptr), "setBlurHints", Q_ARG(QGraphicsBlurEffect::BlurHint, static_cast<QGraphicsBlurEffect::BlurHint>(hints)));
}

QtObjectPtr QGraphicsBlurEffect_NewQGraphicsBlurEffect(QtObjectPtr parent){
	return new QGraphicsBlurEffect(static_cast<QObject*>(parent));
}

void QGraphicsBlurEffect_ConnectBlurHintsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DisconnectBlurHintsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(QtObjectPtr ptr){
	static_cast<QGraphicsBlurEffect*>(ptr)->~QGraphicsBlurEffect();
}

