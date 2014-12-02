#include "qabstractslider.h"
#include <QAbstractSlider>
#include "cgoexport.h"

//MyQAbstractSlider
class MyQAbstractSlider: public QAbstractSlider {
Q_OBJECT
public:
void Signal_ActionTriggered() { callbackQAbstractSlider(this, QString("actionTriggered").toUtf8().data()); };
void Signal_RangeChanged() { callbackQAbstractSlider(this, QString("rangeChanged").toUtf8().data()); };
void Signal_SliderMoved() { callbackQAbstractSlider(this, QString("sliderMoved").toUtf8().data()); };
void Signal_SliderPressed() { callbackQAbstractSlider(this, QString("sliderPressed").toUtf8().data()); };
void Signal_SliderReleased() { callbackQAbstractSlider(this, QString("sliderReleased").toUtf8().data()); };
void Signal_ValueChanged() { callbackQAbstractSlider(this, QString("valueChanged").toUtf8().data()); };

signals:
void Slot_SetOrientation(Qt::Orientation orientation);
void Slot_SetRange(int min, int max);
void Slot_SetValue(int value);

};
#include "qabstractslider.moc"


//Public Functions
QtObjectPtr QAbstractSlider_New_QWidget(QtObjectPtr parent)
{
	return new QAbstractSlider(((QWidget*)(parent)));
}

void QAbstractSlider_Destroy(QtObjectPtr ptr)
{
	((QAbstractSlider*)(ptr))->~QAbstractSlider();
}

int QAbstractSlider_HasTracking(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->hasTracking();
}

int QAbstractSlider_InvertedAppearance(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->invertedAppearance();
}

int QAbstractSlider_InvertedControls(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->invertedControls();
}

int QAbstractSlider_IsSliderDown(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->isSliderDown();
}

int QAbstractSlider_Maximum(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->maximum();
}

int QAbstractSlider_Minimum(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->minimum();
}

int QAbstractSlider_Orientation(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->orientation();
}

int QAbstractSlider_PageStep(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->pageStep();
}

void QAbstractSlider_SetInvertedAppearance_Bool(QtObjectPtr ptr, int invertedAppearance)
{
	((QAbstractSlider*)(ptr))->setInvertedAppearance(invertedAppearance != 0);
}

void QAbstractSlider_SetInvertedControls_Bool(QtObjectPtr ptr, int invertedControls)
{
	((QAbstractSlider*)(ptr))->setInvertedControls(invertedControls != 0);
}

void QAbstractSlider_SetMaximum_Int(QtObjectPtr ptr, int maximum)
{
	((QAbstractSlider*)(ptr))->setMaximum(maximum);
}

void QAbstractSlider_SetMinimum_Int(QtObjectPtr ptr, int minimum)
{
	((QAbstractSlider*)(ptr))->setMinimum(minimum);
}

void QAbstractSlider_SetPageStep_Int(QtObjectPtr ptr, int pageStep)
{
	((QAbstractSlider*)(ptr))->setPageStep(pageStep);
}

void QAbstractSlider_SetSingleStep_Int(QtObjectPtr ptr, int singleStep)
{
	((QAbstractSlider*)(ptr))->setSingleStep(singleStep);
}

void QAbstractSlider_SetSliderDown_Bool(QtObjectPtr ptr, int sliderDown)
{
	((QAbstractSlider*)(ptr))->setSliderDown(sliderDown != 0);
}

void QAbstractSlider_SetSliderPosition_Int(QtObjectPtr ptr, int sliderPosition)
{
	((QAbstractSlider*)(ptr))->setSliderPosition(sliderPosition);
}

void QAbstractSlider_SetTracking_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractSlider*)(ptr))->setTracking(enable != 0);
}

int QAbstractSlider_SingleStep(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->singleStep();
}

int QAbstractSlider_SliderPosition(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->sliderPosition();
}

int QAbstractSlider_Value(QtObjectPtr ptr)
{
	return ((QAbstractSlider*)(ptr))->value();
}

//Public Slots
void QAbstractSlider_ConnectSlotSetOrientation(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetOrientation, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setOrientation, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSlotSetOrientation(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetOrientation, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setOrientation);
}

void QAbstractSlider_SetOrientation_Orientation(QtObjectPtr ptr, int orientation)
{
	((MyQAbstractSlider*)(ptr))->Slot_SetOrientation(((Qt::Orientation)(orientation)));
}

void QAbstractSlider_ConnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetRange, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setRange, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetRange, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setRange);
}

void QAbstractSlider_SetRange_Int_Int(QtObjectPtr ptr, int min, int max)
{
	((MyQAbstractSlider*)(ptr))->Slot_SetRange(min, max);
}

void QAbstractSlider_ConnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetValue, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setValue, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Slot_SetValue, ((QAbstractSlider*)(ptr)), &QAbstractSlider::setValue);
}

void QAbstractSlider_SetValue_Int(QtObjectPtr ptr, int value)
{
	((MyQAbstractSlider*)(ptr))->Slot_SetValue(value);
}

//Signals
void QAbstractSlider_ConnectSignalActionTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::actionTriggered, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_ActionTriggered, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalActionTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::actionTriggered, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_ActionTriggered);
}

void QAbstractSlider_ConnectSignalRangeChanged(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::rangeChanged, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_RangeChanged, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalRangeChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::rangeChanged, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_RangeChanged);
}

void QAbstractSlider_ConnectSignalSliderMoved(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderMoved, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderMoved, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalSliderMoved(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderMoved, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderMoved);
}

void QAbstractSlider_ConnectSignalSliderPressed(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderPressed, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderPressed, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalSliderPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderPressed, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderPressed);
}

void QAbstractSlider_ConnectSignalSliderReleased(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderReleased, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderReleased, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalSliderReleased(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::sliderReleased, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_SliderReleased);
}

void QAbstractSlider_ConnectSignalValueChanged(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSlider*)(ptr)), &QAbstractSlider::valueChanged, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_ValueChanged, Qt::QueuedConnection);
}

void QAbstractSlider_DisconnectSignalValueChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSlider*)(ptr)), &QAbstractSlider::valueChanged, ((MyQAbstractSlider*)(ptr)), &MyQAbstractSlider::Signal_ValueChanged);
}

