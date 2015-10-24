#include "qvariantanimation.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEasingCurve>
#include <QVariantAnimation>
#include "_cgo_export.h"

class MyQVariantAnimation: public QVariantAnimation {
public:
void Signal_ValueChanged(const QVariant & value){callbackQVariantAnimationValueChanged(this->objectName().toUtf8().data(), value.toString().toUtf8().data());};
};

char* QVariantAnimation_CurrentValue(QtObjectPtr ptr){
	return static_cast<QVariantAnimation*>(ptr)->currentValue().toString().toUtf8().data();
}

int QVariantAnimation_Duration(QtObjectPtr ptr){
	return static_cast<QVariantAnimation*>(ptr)->duration();
}

char* QVariantAnimation_EndValue(QtObjectPtr ptr){
	return static_cast<QVariantAnimation*>(ptr)->endValue().toString().toUtf8().data();
}

void QVariantAnimation_SetDuration(QtObjectPtr ptr, int msecs){
	static_cast<QVariantAnimation*>(ptr)->setDuration(msecs);
}

void QVariantAnimation_SetEasingCurve(QtObjectPtr ptr, QtObjectPtr easing){
	static_cast<QVariantAnimation*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(easing));
}

void QVariantAnimation_SetEndValue(QtObjectPtr ptr, char* value){
	static_cast<QVariantAnimation*>(ptr)->setEndValue(QVariant(value));
}

void QVariantAnimation_SetStartValue(QtObjectPtr ptr, char* value){
	static_cast<QVariantAnimation*>(ptr)->setStartValue(QVariant(value));
}

char* QVariantAnimation_StartValue(QtObjectPtr ptr){
	return static_cast<QVariantAnimation*>(ptr)->startValue().toString().toUtf8().data();
}

QtObjectPtr QVariantAnimation_NewQVariantAnimation(QtObjectPtr parent){
	return new QVariantAnimation(static_cast<QObject*>(parent));
}

void QVariantAnimation_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_DestroyQVariantAnimation(QtObjectPtr ptr){
	static_cast<QVariantAnimation*>(ptr)->~QVariantAnimation();
}

