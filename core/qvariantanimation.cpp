#include "qvariantanimation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QEasingCurve>
#include <QVariantAnimation>
#include "_cgo_export.h"

class MyQVariantAnimation: public QVariantAnimation {
public:
void Signal_ValueChanged(const QVariant & value){callbackQVariantAnimationValueChanged(this->objectName().toUtf8().data(), new QVariant(value));};
};

void* QVariantAnimation_CurrentValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->currentValue());
}

int QVariantAnimation_Duration(void* ptr){
	return static_cast<QVariantAnimation*>(ptr)->duration();
}

void* QVariantAnimation_EasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QVariantAnimation*>(ptr)->easingCurve());
}

void* QVariantAnimation_EndValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->endValue());
}

void QVariantAnimation_SetDuration(void* ptr, int msecs){
	static_cast<QVariantAnimation*>(ptr)->setDuration(msecs);
}

void QVariantAnimation_SetEasingCurve(void* ptr, void* easing){
	static_cast<QVariantAnimation*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(easing));
}

void QVariantAnimation_SetEndValue(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->setEndValue(*static_cast<QVariant*>(value));
}

void QVariantAnimation_SetStartValue(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->setStartValue(*static_cast<QVariant*>(value));
}

void* QVariantAnimation_StartValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->startValue());
}

void* QVariantAnimation_NewQVariantAnimation(void* parent){
	return new QVariantAnimation(static_cast<QObject*>(parent));
}

void* QVariantAnimation_KeyValueAt(void* ptr, double step){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->keyValueAt(static_cast<qreal>(step)));
}

void QVariantAnimation_SetKeyValueAt(void* ptr, double step, void* value){
	static_cast<QVariantAnimation*>(ptr)->setKeyValueAt(static_cast<qreal>(step), *static_cast<QVariant*>(value));
}

void QVariantAnimation_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_DestroyQVariantAnimation(void* ptr){
	static_cast<QVariantAnimation*>(ptr)->~QVariantAnimation();
}

