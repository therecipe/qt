#include "qprogressdialog.h"
#include <QProgressDialog>
#include "cgoexport.h"

//MyQProgressDialog
class MyQProgressDialog: public QProgressDialog {
Q_OBJECT
public:
void Signal_Canceled() { callbackQProgressDialog(this, QString("canceled").toUtf8().data()); };

signals:
void Slot_Cancel();
void Slot_Reset();
void Slot_SetMaximum(int maximum);
void Slot_SetMinimum(int minimum);
void Slot_SetMinimumDuration(int ms);
void Slot_SetRange(int minimum, int maximum);
void Slot_SetValue(int progress);

};
#include "qprogressdialog.moc"


//Public Functions
QtObjectPtr QProgressDialog_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QProgressDialog(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

QtObjectPtr QProgressDialog_New_String_String_Int_Int_QWidget_WindowType(char* labelText, char* cancelButtonText, int minimum, int maximum, QtObjectPtr parent, int f)
{
	return new QProgressDialog(QString(labelText), QString(cancelButtonText), minimum, maximum, ((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QProgressDialog_Destroy(QtObjectPtr ptr)
{
	((QProgressDialog*)(ptr))->~QProgressDialog();
}

int QProgressDialog_AutoClose(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->autoClose();
}

int QProgressDialog_AutoReset(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->autoReset();
}

char* QProgressDialog_LabelText(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->labelText().toUtf8().data();
}

int QProgressDialog_Maximum(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->maximum();
}

int QProgressDialog_Minimum(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->minimum();
}

int QProgressDialog_MinimumDuration(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->minimumDuration();
}

void QProgressDialog_SetAutoClose_Bool(QtObjectPtr ptr, int close)
{
	((QProgressDialog*)(ptr))->setAutoClose(close != 0);
}

void QProgressDialog_SetAutoReset_Bool(QtObjectPtr ptr, int reset)
{
	((QProgressDialog*)(ptr))->setAutoReset(reset != 0);
}

void QProgressDialog_SetBar_QProgressBar(QtObjectPtr ptr, QtObjectPtr bar)
{
	((QProgressDialog*)(ptr))->setBar(((QProgressBar*)(bar)));
}

void QProgressDialog_SetCancelButton_QPushButton(QtObjectPtr ptr, QtObjectPtr cancelButton)
{
	((QProgressDialog*)(ptr))->setCancelButton(((QPushButton*)(cancelButton)));
}

void QProgressDialog_SetLabel_QLabel(QtObjectPtr ptr, QtObjectPtr label)
{
	((QProgressDialog*)(ptr))->setLabel(((QLabel*)(label)));
}

int QProgressDialog_Value(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->value();
}

int QProgressDialog_WasCanceled(QtObjectPtr ptr)
{
	return ((QProgressDialog*)(ptr))->wasCanceled();
}

//Public Slots
void QProgressDialog_ConnectSlotCancel(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_Cancel, ((QProgressDialog*)(ptr)), &QProgressDialog::cancel, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotCancel(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_Cancel, ((QProgressDialog*)(ptr)), &QProgressDialog::cancel);
}

void QProgressDialog_Cancel(QtObjectPtr ptr)
{
	((MyQProgressDialog*)(ptr))->Slot_Cancel();
}

void QProgressDialog_ConnectSlotReset(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_Reset, ((QProgressDialog*)(ptr)), &QProgressDialog::reset, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotReset(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_Reset, ((QProgressDialog*)(ptr)), &QProgressDialog::reset);
}

void QProgressDialog_Reset(QtObjectPtr ptr)
{
	((MyQProgressDialog*)(ptr))->Slot_Reset();
}

void QProgressDialog_ConnectSlotSetMaximum(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMaximum, ((QProgressDialog*)(ptr)), &QProgressDialog::setMaximum, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotSetMaximum(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMaximum, ((QProgressDialog*)(ptr)), &QProgressDialog::setMaximum);
}

void QProgressDialog_SetMaximum_Int(QtObjectPtr ptr, int maximum)
{
	((MyQProgressDialog*)(ptr))->Slot_SetMaximum(maximum);
}

void QProgressDialog_ConnectSlotSetMinimum(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMinimum, ((QProgressDialog*)(ptr)), &QProgressDialog::setMinimum, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotSetMinimum(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMinimum, ((QProgressDialog*)(ptr)), &QProgressDialog::setMinimum);
}

void QProgressDialog_SetMinimum_Int(QtObjectPtr ptr, int minimum)
{
	((MyQProgressDialog*)(ptr))->Slot_SetMinimum(minimum);
}

void QProgressDialog_ConnectSlotSetMinimumDuration(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMinimumDuration, ((QProgressDialog*)(ptr)), &QProgressDialog::setMinimumDuration, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotSetMinimumDuration(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetMinimumDuration, ((QProgressDialog*)(ptr)), &QProgressDialog::setMinimumDuration);
}

void QProgressDialog_SetMinimumDuration_Int(QtObjectPtr ptr, int ms)
{
	((MyQProgressDialog*)(ptr))->Slot_SetMinimumDuration(ms);
}

void QProgressDialog_ConnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetRange, ((QProgressDialog*)(ptr)), &QProgressDialog::setRange, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotSetRange(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetRange, ((QProgressDialog*)(ptr)), &QProgressDialog::setRange);
}

void QProgressDialog_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum)
{
	((MyQProgressDialog*)(ptr))->Slot_SetRange(minimum, maximum);
}

void QProgressDialog_ConnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::connect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetValue, ((QProgressDialog*)(ptr)), &QProgressDialog::setValue, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSlotSetValue(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Slot_SetValue, ((QProgressDialog*)(ptr)), &QProgressDialog::setValue);
}

void QProgressDialog_SetValue_Int(QtObjectPtr ptr, int progress)
{
	((MyQProgressDialog*)(ptr))->Slot_SetValue(progress);
}

//Signals
void QProgressDialog_ConnectSignalCanceled(QtObjectPtr ptr)
{
	QObject::connect(((QProgressDialog*)(ptr)), &QProgressDialog::canceled, ((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Signal_Canceled, Qt::QueuedConnection);
}

void QProgressDialog_DisconnectSignalCanceled(QtObjectPtr ptr)
{
	QObject::disconnect(((QProgressDialog*)(ptr)), &QProgressDialog::canceled, ((MyQProgressDialog*)(ptr)), &MyQProgressDialog::Signal_Canceled);
}

