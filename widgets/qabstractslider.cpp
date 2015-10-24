#include "qabstractslider.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QAbstractSlider>
#include "_cgo_export.h"

class MyQAbstractSlider: public QAbstractSlider {
public:
void Signal_ActionTriggered(int action){callbackQAbstractSliderActionTriggered(this->objectName().toUtf8().data(), action);};
void Signal_RangeChanged(int min, int max){callbackQAbstractSliderRangeChanged(this->objectName().toUtf8().data(), min, max);};
void Signal_SliderMoved(int value){callbackQAbstractSliderSliderMoved(this->objectName().toUtf8().data(), value);};
void Signal_SliderPressed(){callbackQAbstractSliderSliderPressed(this->objectName().toUtf8().data());};
void Signal_SliderReleased(){callbackQAbstractSliderSliderReleased(this->objectName().toUtf8().data());};
void Signal_ValueChanged(int value){callbackQAbstractSliderValueChanged(this->objectName().toUtf8().data(), value);};
};

int QAbstractSlider_HasTracking(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->hasTracking();
}

int QAbstractSlider_InvertedAppearance(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->invertedAppearance();
}

int QAbstractSlider_InvertedControls(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->invertedControls();
}

int QAbstractSlider_IsSliderDown(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->isSliderDown();
}

int QAbstractSlider_Maximum(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->maximum();
}

int QAbstractSlider_Minimum(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->minimum();
}

int QAbstractSlider_Orientation(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->orientation();
}

int QAbstractSlider_PageStep(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->pageStep();
}

void QAbstractSlider_SetInvertedAppearance(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setInvertedAppearance(v != 0);
}

void QAbstractSlider_SetInvertedControls(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setInvertedControls(v != 0);
}

void QAbstractSlider_SetMaximum(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setMaximum(v);
}

void QAbstractSlider_SetMinimum(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setMinimum(v);
}

void QAbstractSlider_SetOrientation(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(v)));
}

void QAbstractSlider_SetPageStep(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setPageStep(v);
}

void QAbstractSlider_SetSingleStep(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSingleStep(v);
}

void QAbstractSlider_SetSliderDown(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSliderDown(v != 0);
}

void QAbstractSlider_SetSliderPosition(QtObjectPtr ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSliderPosition(v);
}

void QAbstractSlider_SetTracking(QtObjectPtr ptr, int enable){
	static_cast<QAbstractSlider*>(ptr)->setTracking(enable != 0);
}

void QAbstractSlider_SetValue(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setValue", Q_ARG(int, v));
}

int QAbstractSlider_SingleStep(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->singleStep();
}

int QAbstractSlider_SliderPosition(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->sliderPosition();
}

int QAbstractSlider_Value(QtObjectPtr ptr){
	return static_cast<QAbstractSlider*>(ptr)->value();
}

QtObjectPtr QAbstractSlider_NewQAbstractSlider(QtObjectPtr parent){
	return new QAbstractSlider(static_cast<QWidget*>(parent));
}

void QAbstractSlider_ConnectActionTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::actionTriggered), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ActionTriggered));;
}

void QAbstractSlider_DisconnectActionTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::actionTriggered), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ActionTriggered));;
}

void QAbstractSlider_ConnectRangeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int, int)>(&QAbstractSlider::rangeChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int, int)>(&MyQAbstractSlider::Signal_RangeChanged));;
}

void QAbstractSlider_DisconnectRangeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int, int)>(&QAbstractSlider::rangeChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int, int)>(&MyQAbstractSlider::Signal_RangeChanged));;
}

void QAbstractSlider_SetRange(QtObjectPtr ptr, int min, int max){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setRange", Q_ARG(int, min), Q_ARG(int, max));
}

void QAbstractSlider_ConnectSliderMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::sliderMoved), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_SliderMoved));;
}

void QAbstractSlider_DisconnectSliderMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::sliderMoved), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_SliderMoved));;
}

void QAbstractSlider_ConnectSliderPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderPressed), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderPressed));;
}

void QAbstractSlider_DisconnectSliderPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderPressed), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderPressed));;
}

void QAbstractSlider_ConnectSliderReleased(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderReleased), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderReleased));;
}

void QAbstractSlider_DisconnectSliderReleased(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderReleased), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderReleased));;
}

void QAbstractSlider_TriggerAction(QtObjectPtr ptr, int action){
	static_cast<QAbstractSlider*>(ptr)->triggerAction(static_cast<QAbstractSlider::SliderAction>(action));
}

void QAbstractSlider_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::valueChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ValueChanged));;
}

void QAbstractSlider_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::valueChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ValueChanged));;
}

void QAbstractSlider_DestroyQAbstractSlider(QtObjectPtr ptr){
	static_cast<QAbstractSlider*>(ptr)->~QAbstractSlider();
}

