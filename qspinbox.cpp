#include "qspinbox.h"
#include <QSpinBox>
#include "cgoexport.h"

//MyQSpinBox
class MyQSpinBox: public QSpinBox {
Q_OBJECT
public:

signals:
void Slot_SetValue(int val);

};
#include "qspinbox.moc"


//Public Functions
QtObjectPtr QSpinBox_New_QWidget(QtObjectPtr parent)
{
	return new QSpinBox(((QWidget*)(parent)));
}

void QSpinBox_Destroy(QtObjectPtr ptr)
{
	((QSpinBox*)(ptr))->~QSpinBox();
}

char* QSpinBox_CleanText(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->cleanText().toUtf8().data();
}

int QSpinBox_DisplayIntegerBase(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->displayIntegerBase();
}

int QSpinBox_Maximum(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->maximum();
}

int QSpinBox_Minimum(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->minimum();
}

char* QSpinBox_Prefix(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->prefix().toUtf8().data();
}

void QSpinBox_SetDisplayIntegerBase_Int(QtObjectPtr ptr, int base)
{
	((QSpinBox*)(ptr))->setDisplayIntegerBase(base);
}

void QSpinBox_SetMaximum_Int(QtObjectPtr ptr, int max)
{
	((QSpinBox*)(ptr))->setMaximum(max);
}

void QSpinBox_SetMinimum_Int(QtObjectPtr ptr, int min)
{
	((QSpinBox*)(ptr))->setMinimum(min);
}

void QSpinBox_SetPrefix_String(QtObjectPtr ptr, char* prefix)
{
	((QSpinBox*)(ptr))->setPrefix(QString(prefix));
}

void QSpinBox_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum)
{
	((QSpinBox*)(ptr))->setRange(minimum, maximum);
}

void QSpinBox_SetSingleStep_Int(QtObjectPtr ptr, int val)
{
	((QSpinBox*)(ptr))->setSingleStep(val);
}

void QSpinBox_SetSuffix_String(QtObjectPtr ptr, char* suffix)
{
	((QSpinBox*)(ptr))->setSuffix(QString(suffix));
}

int QSpinBox_SingleStep(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->singleStep();
}

char* QSpinBox_Suffix(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->suffix().toUtf8().data();
}

int QSpinBox_Value(QtObjectPtr ptr)
{
	return ((QSpinBox*)(ptr))->value();
}

//Public Slots
void QSpinBox_ConnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::connect(((MyQSpinBox*)(ptr)), &MyQSpinBox::Slot_SetValue, ((QSpinBox*)(ptr)), &QSpinBox::setValue, Qt::QueuedConnection);
}

void QSpinBox_DisconnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQSpinBox*)(ptr)), &MyQSpinBox::Slot_SetValue, ((QSpinBox*)(ptr)), &QSpinBox::setValue);
}

void QSpinBox_SetValue_Int(QtObjectPtr ptr, int val)
{
	((MyQSpinBox*)(ptr))->Slot_SetValue(val);
}

