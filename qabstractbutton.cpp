#include "qabstractbutton.h"
#include <QAbstractButton>
#include "cgoexport.h"

//MyQAbstractButton
class MyQAbstractButton: public QAbstractButton {
Q_OBJECT
public:
void Signal_Clicked() { callbackQAbstractButton(this, QString("clicked").toUtf8().data()); };
void Signal_Pressed() { callbackQAbstractButton(this, QString("pressed").toUtf8().data()); };
void Signal_Released() { callbackQAbstractButton(this, QString("released").toUtf8().data()); };
void Signal_Toggled() { callbackQAbstractButton(this, QString("toggled").toUtf8().data()); };

signals:
void Slot_AnimateClick(int msec);
void Slot_Click();
void Slot_SetChecked(bool checked);
void Slot_Toggle();

};
#include "qabstractbutton.moc"


//Public Functions
void QAbstractButton_Destroy(QtObjectPtr ptr)
{
	((QAbstractButton*)(ptr))->~QAbstractButton();
}

int QAbstractButton_AutoExclusive(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->autoExclusive();
}

int QAbstractButton_AutoRepeat(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->autoRepeat();
}

int QAbstractButton_AutoRepeatDelay(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->autoRepeatDelay();
}

int QAbstractButton_AutoRepeatInterval(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->autoRepeatInterval();
}

int QAbstractButton_IsCheckable(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->isCheckable();
}

int QAbstractButton_IsChecked(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->isChecked();
}

int QAbstractButton_IsDown(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->isDown();
}

void QAbstractButton_SetAutoExclusive_Bool(QtObjectPtr ptr, int autoExclusive)
{
	((QAbstractButton*)(ptr))->setAutoExclusive(autoExclusive != 0);
}

void QAbstractButton_SetAutoRepeat_Bool(QtObjectPtr ptr, int autoRepeat)
{
	((QAbstractButton*)(ptr))->setAutoRepeat(autoRepeat != 0);
}

void QAbstractButton_SetAutoRepeatDelay_Int(QtObjectPtr ptr, int autoRepeatDelay)
{
	((QAbstractButton*)(ptr))->setAutoRepeatDelay(autoRepeatDelay);
}

void QAbstractButton_SetAutoRepeatInterval_Int(QtObjectPtr ptr, int autoRepeatInterval)
{
	((QAbstractButton*)(ptr))->setAutoRepeatInterval(autoRepeatInterval);
}

void QAbstractButton_SetCheckable_Bool(QtObjectPtr ptr, int checkable)
{
	((QAbstractButton*)(ptr))->setCheckable(checkable != 0);
}

void QAbstractButton_SetDown_Bool(QtObjectPtr ptr, int down)
{
	((QAbstractButton*)(ptr))->setDown(down != 0);
}

void QAbstractButton_SetText_String(QtObjectPtr ptr, char* text)
{
	((QAbstractButton*)(ptr))->setText(QString(text));
}

char* QAbstractButton_Text(QtObjectPtr ptr)
{
	return ((QAbstractButton*)(ptr))->text().toUtf8().data();
}

//Public Slots
void QAbstractButton_ConnectSlotAnimateClick(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_AnimateClick, ((QAbstractButton*)(ptr)), &QAbstractButton::animateClick, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSlotAnimateClick(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_AnimateClick, ((QAbstractButton*)(ptr)), &QAbstractButton::animateClick);
}

void QAbstractButton_AnimateClick_Int(QtObjectPtr ptr, int msec)
{
	((MyQAbstractButton*)(ptr))->Slot_AnimateClick(msec);
}

void QAbstractButton_ConnectSlotClick(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_Click, ((QAbstractButton*)(ptr)), &QAbstractButton::click, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSlotClick(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_Click, ((QAbstractButton*)(ptr)), &QAbstractButton::click);
}

void QAbstractButton_Click(QtObjectPtr ptr)
{
	((MyQAbstractButton*)(ptr))->Slot_Click();
}

void QAbstractButton_ConnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_SetChecked, ((QAbstractButton*)(ptr)), &QAbstractButton::setChecked, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_SetChecked, ((QAbstractButton*)(ptr)), &QAbstractButton::setChecked);
}

void QAbstractButton_SetChecked_Bool(QtObjectPtr ptr, int checked)
{
	((MyQAbstractButton*)(ptr))->Slot_SetChecked(checked != 0);
}

void QAbstractButton_ConnectSlotToggle(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_Toggle, ((QAbstractButton*)(ptr)), &QAbstractButton::toggle, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSlotToggle(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Slot_Toggle, ((QAbstractButton*)(ptr)), &QAbstractButton::toggle);
}

void QAbstractButton_Toggle(QtObjectPtr ptr)
{
	((MyQAbstractButton*)(ptr))->Slot_Toggle();
}

//Signals
void QAbstractButton_ConnectSignalClicked(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractButton*)(ptr)), &QAbstractButton::clicked, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Clicked, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSignalClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractButton*)(ptr)), &QAbstractButton::clicked, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Clicked);
}

void QAbstractButton_ConnectSignalPressed(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractButton*)(ptr)), &QAbstractButton::pressed, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Pressed, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSignalPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractButton*)(ptr)), &QAbstractButton::pressed, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Pressed);
}

void QAbstractButton_ConnectSignalReleased(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractButton*)(ptr)), &QAbstractButton::released, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Released, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSignalReleased(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractButton*)(ptr)), &QAbstractButton::released, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Released);
}

void QAbstractButton_ConnectSignalToggled(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractButton*)(ptr)), &QAbstractButton::toggled, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Toggled, Qt::QueuedConnection);
}

void QAbstractButton_DisconnectSignalToggled(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractButton*)(ptr)), &QAbstractButton::toggled, ((MyQAbstractButton*)(ptr)), &MyQAbstractButton::Signal_Toggled);
}

