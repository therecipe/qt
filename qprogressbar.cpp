#include "qprogressbar.h"
#include <QProgressBar>
#include "cgoexport.h"

//MyQProgressBar
class MyQProgressBar: public QProgressBar {
Q_OBJECT
public:
void Signal_ValueChanged() { callbackQProgressBar(this, QString("valueChanged").toUtf8().data()); };

signals:
void Slot_Reset();
void Slot_SetMaximum(int maximum);
void Slot_SetMinimum(int minimum);
void Slot_SetOrientation(Qt::Orientation orientation);
void Slot_SetRange(int minimum, int maximum);
void Slot_SetValue(int value);

};
#include "qprogressbar.moc"


//Public Functions
QtObjectPtr QProgressBar_New_QWidget(QtObjectPtr parent)
{
	return new QProgressBar(((QWidget*)(parent)));
}

void QProgressBar_Destroy(QtObjectPtr ptr)
{
	((QProgressBar*)(ptr))->~QProgressBar();
}

int QProgressBar_Alignment(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->alignment();
}

char* QProgressBar_Format(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->format().toUtf8().data();
}

int QProgressBar_InvertedAppearance(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->invertedAppearance();
}

int QProgressBar_IsTextVisible(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->isTextVisible();
}

int QProgressBar_Maximum(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->maximum();
}

int QProgressBar_Minimum(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->minimum();
}

int QProgressBar_Orientation(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->orientation();
}

void QProgressBar_ResetFormat(QtObjectPtr ptr)
{
	((QProgressBar*)(ptr))->resetFormat();
}

void QProgressBar_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QProgressBar*)(ptr))->setAlignment(((Qt::AlignmentFlag)(alignment)));
}

void QProgressBar_SetFormat_String(QtObjectPtr ptr, char* format)
{
	((QProgressBar*)(ptr))->setFormat(QString(format));
}

void QProgressBar_SetInvertedAppearance_Bool(QtObjectPtr ptr, int invert)
{
	((QProgressBar*)(ptr))->setInvertedAppearance(invert != 0);
}

void QProgressBar_SetTextVisible_Bool(QtObjectPtr ptr, int visible)
{
	((QProgressBar*)(ptr))->setTextVisible(visible != 0);
}

char* QProgressBar_Text(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->text().toUtf8().data();
}

int QProgressBar_Value(QtObjectPtr ptr)
{
	return ((QProgressBar*)(ptr))->value();
}

//Public Slots
void QProgressBar_ConnectSlotReset(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_Reset, ((QProgressBar*)(ptr)), &QProgressBar::reset, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotReset(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_Reset, ((QProgressBar*)(ptr)), &QProgressBar::reset);
}

void QProgressBar_Reset(QtObjectPtr ptr)
{
	((MyQProgressBar*)(ptr))->Slot_Reset();
}

void QProgressBar_ConnectSlotSetMaximum(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetMaximum, ((QProgressBar*)(ptr)), &QProgressBar::setMaximum, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotSetMaximum(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetMaximum, ((QProgressBar*)(ptr)), &QProgressBar::setMaximum);
}

void QProgressBar_SetMaximum_Int(QtObjectPtr ptr, int maximum)
{
	((MyQProgressBar*)(ptr))->Slot_SetMaximum(maximum);
}

void QProgressBar_ConnectSlotSetMinimum(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetMinimum, ((QProgressBar*)(ptr)), &QProgressBar::setMinimum, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotSetMinimum(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetMinimum, ((QProgressBar*)(ptr)), &QProgressBar::setMinimum);
}

void QProgressBar_SetMinimum_Int(QtObjectPtr ptr, int minimum)
{
	((MyQProgressBar*)(ptr))->Slot_SetMinimum(minimum);
}

void QProgressBar_ConnectSlotSetOrientation(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetOrientation, ((QProgressBar*)(ptr)), &QProgressBar::setOrientation, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotSetOrientation(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetOrientation, ((QProgressBar*)(ptr)), &QProgressBar::setOrientation);
}

void QProgressBar_SetOrientation_Orientation(QtObjectPtr ptr, int orientation)
{
	((MyQProgressBar*)(ptr))->Slot_SetOrientation(((Qt::Orientation)(orientation)));
}

void QProgressBar_ConnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetRange, ((QProgressBar*)(ptr)), &QProgressBar::setRange, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetRange, ((QProgressBar*)(ptr)), &QProgressBar::setRange);
}

void QProgressBar_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum)
{
	((MyQProgressBar*)(ptr))->Slot_SetRange(minimum, maximum);
}

void QProgressBar_ConnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetValue, ((QProgressBar*)(ptr)), &QProgressBar::setValue, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressBar*)(ptr)), &MyQProgressBar::Slot_SetValue, ((QProgressBar*)(ptr)), &QProgressBar::setValue);
}

void QProgressBar_SetValue_Int(QtObjectPtr ptr, int value)
{
	((MyQProgressBar*)(ptr))->Slot_SetValue(value);
}

//Signals
void QProgressBar_ConnectSignalValueChanged(QtObjectPtr ptr)
{
	QObject::connect(((QProgressBar*)(ptr)), &QProgressBar::valueChanged, ((MyQProgressBar*)(ptr)), &MyQProgressBar::Signal_ValueChanged, Qt::QueuedConnection);
}

void QProgressBar_DisconnectSignalValueChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QProgressBar*)(ptr)), &QProgressBar::valueChanged, ((MyQProgressBar*)(ptr)), &MyQProgressBar::Signal_ValueChanged);
}

