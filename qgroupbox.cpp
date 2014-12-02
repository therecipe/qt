#include "qgroupbox.h"
#include <QGroupBox>
#include "cgoexport.h"

//MyQGroupBox
class MyQGroupBox: public QGroupBox {
Q_OBJECT
public:
void Signal_Clicked() { callbackQGroupBox(this, QString("clicked").toUtf8().data()); };
void Signal_Toggled() { callbackQGroupBox(this, QString("toggled").toUtf8().data()); };

signals:
void Slot_SetChecked(bool checked);

};
#include "qgroupbox.moc"


//Public Functions
QtObjectPtr QGroupBox_New_QWidget(QtObjectPtr parent)
{
	return new QGroupBox(((QWidget*)(parent)));
}

QtObjectPtr QGroupBox_New_String_QWidget(char* title, QtObjectPtr parent)
{
	return new QGroupBox(QString(title), ((QWidget*)(parent)));
}

void QGroupBox_Destroy(QtObjectPtr ptr)
{
	((QGroupBox*)(ptr))->~QGroupBox();
}

int QGroupBox_Alignment(QtObjectPtr ptr)
{
	return ((QGroupBox*)(ptr))->alignment();
}

int QGroupBox_IsCheckable(QtObjectPtr ptr)
{
	return ((QGroupBox*)(ptr))->isCheckable();
}

int QGroupBox_IsChecked(QtObjectPtr ptr)
{
	return ((QGroupBox*)(ptr))->isChecked();
}

int QGroupBox_IsFlat(QtObjectPtr ptr)
{
	return ((QGroupBox*)(ptr))->isFlat();
}

void QGroupBox_SetAlignment_Int(QtObjectPtr ptr, int alignment)
{
	((QGroupBox*)(ptr))->setAlignment(alignment);
}

void QGroupBox_SetCheckable_Bool(QtObjectPtr ptr, int checkable)
{
	((QGroupBox*)(ptr))->setCheckable(checkable != 0);
}

void QGroupBox_SetFlat_Bool(QtObjectPtr ptr, int flat)
{
	((QGroupBox*)(ptr))->setFlat(flat != 0);
}

void QGroupBox_SetTitle_String(QtObjectPtr ptr, char* title)
{
	((QGroupBox*)(ptr))->setTitle(QString(title));
}

char* QGroupBox_Title(QtObjectPtr ptr)
{
	return ((QGroupBox*)(ptr))->title().toUtf8().data();
}

//Public Slots
void QGroupBox_ConnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::connect(((MyQGroupBox*)(ptr)), &MyQGroupBox::Slot_SetChecked, ((QGroupBox*)(ptr)), &QGroupBox::setChecked, Qt::QueuedConnection);
}

void QGroupBox_DisconnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQGroupBox*)(ptr)), &MyQGroupBox::Slot_SetChecked, ((QGroupBox*)(ptr)), &QGroupBox::setChecked);
}

void QGroupBox_SetChecked_Bool(QtObjectPtr ptr, int checked)
{
	((MyQGroupBox*)(ptr))->Slot_SetChecked(checked != 0);
}

//Signals
void QGroupBox_ConnectSignalClicked(QtObjectPtr ptr)
{
	QObject::connect(((QGroupBox*)(ptr)), &QGroupBox::clicked, ((MyQGroupBox*)(ptr)), &MyQGroupBox::Signal_Clicked, Qt::QueuedConnection);
}

void QGroupBox_DisconnectSignalClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QGroupBox*)(ptr)), &QGroupBox::clicked, ((MyQGroupBox*)(ptr)), &MyQGroupBox::Signal_Clicked);
}

void QGroupBox_ConnectSignalToggled(QtObjectPtr ptr)
{
	QObject::connect(((QGroupBox*)(ptr)), &QGroupBox::toggled, ((MyQGroupBox*)(ptr)), &MyQGroupBox::Signal_Toggled, Qt::QueuedConnection);
}

void QGroupBox_DisconnectSignalToggled(QtObjectPtr ptr)
{
	QObject::disconnect(((QGroupBox*)(ptr)), &QGroupBox::toggled, ((MyQGroupBox*)(ptr)), &MyQGroupBox::Signal_Toggled);
}

