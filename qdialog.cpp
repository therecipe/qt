#include "qdialog.h"
#include <QDialog>
#include "cgoexport.h"

//MyQDialog
class MyQDialog: public QDialog {
Q_OBJECT
public:
void Signal_Accepted() { callbackQDialog(this, QString("accepted").toUtf8().data()); };
void Signal_Finished() { callbackQDialog(this, QString("finished").toUtf8().data()); };
void Signal_Rejected() { callbackQDialog(this, QString("rejected").toUtf8().data()); };

signals:
void Slot_Accept();
void Slot_Done(int r);
void Slot_Exec();
void Slot_Open();
void Slot_Reject();

};
#include "qdialog.moc"


//Public Functions
QtObjectPtr QDialog_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QDialog(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QDialog_Destroy(QtObjectPtr ptr)
{
	((QDialog*)(ptr))->~QDialog();
}

int QDialog_IsSizeGripEnabled(QtObjectPtr ptr)
{
	return ((QDialog*)(ptr))->isSizeGripEnabled();
}

int QDialog_Result(QtObjectPtr ptr)
{
	return ((QDialog*)(ptr))->result();
}

void QDialog_SetModal_Bool(QtObjectPtr ptr, int modal)
{
	((QDialog*)(ptr))->setModal(modal != 0);
}

void QDialog_SetResult_Int(QtObjectPtr ptr, int i)
{
	((QDialog*)(ptr))->setResult(i);
}

void QDialog_SetSizeGripEnabled_Bool(QtObjectPtr ptr, int sizeGripEnabled)
{
	((QDialog*)(ptr))->setSizeGripEnabled(sizeGripEnabled != 0);
}

//Public Slots
void QDialog_ConnectSlotAccept(QtObjectPtr ptr)
{
	QObject::connect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Accept, ((QDialog*)(ptr)), &QDialog::accept, Qt::QueuedConnection);
}

void QDialog_DisconnectSlotAccept(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Accept, ((QDialog*)(ptr)), &QDialog::accept);
}

void QDialog_Accept(QtObjectPtr ptr)
{
	((MyQDialog*)(ptr))->Slot_Accept();
}

void QDialog_ConnectSlotDone(QtObjectPtr ptr)
{
	QObject::connect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Done, ((QDialog*)(ptr)), &QDialog::done, Qt::QueuedConnection);
}

void QDialog_DisconnectSlotDone(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Done, ((QDialog*)(ptr)), &QDialog::done);
}

void QDialog_Done_Int(QtObjectPtr ptr, int r)
{
	((MyQDialog*)(ptr))->Slot_Done(r);
}

void QDialog_ConnectSlotExec(QtObjectPtr ptr)
{
	QObject::connect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Exec, ((QDialog*)(ptr)), &QDialog::exec, Qt::QueuedConnection);
}

void QDialog_DisconnectSlotExec(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Exec, ((QDialog*)(ptr)), &QDialog::exec);
}

void QDialog_Exec(QtObjectPtr ptr)
{
	return ((MyQDialog*)(ptr))->Slot_Exec();
}

void QDialog_ConnectSlotOpen(QtObjectPtr ptr)
{
	QObject::connect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Open, ((QDialog*)(ptr)), &QDialog::open, Qt::QueuedConnection);
}

void QDialog_DisconnectSlotOpen(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Open, ((QDialog*)(ptr)), &QDialog::open);
}

void QDialog_Open(QtObjectPtr ptr)
{
	((MyQDialog*)(ptr))->Slot_Open();
}

void QDialog_ConnectSlotReject(QtObjectPtr ptr)
{
	QObject::connect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Reject, ((QDialog*)(ptr)), &QDialog::reject, Qt::QueuedConnection);
}

void QDialog_DisconnectSlotReject(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQDialog*)(ptr)), &MyQDialog::Slot_Reject, ((QDialog*)(ptr)), &QDialog::reject);
}

void QDialog_Reject(QtObjectPtr ptr)
{
	((MyQDialog*)(ptr))->Slot_Reject();
}

//Signals
void QDialog_ConnectSignalAccepted(QtObjectPtr ptr)
{
	QObject::connect(((QDialog*)(ptr)), &QDialog::accepted, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Accepted, Qt::QueuedConnection);
}

void QDialog_DisconnectSignalAccepted(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialog*)(ptr)), &QDialog::accepted, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Accepted);
}

void QDialog_ConnectSignalFinished(QtObjectPtr ptr)
{
	QObject::connect(((QDialog*)(ptr)), &QDialog::finished, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Finished, Qt::QueuedConnection);
}

void QDialog_DisconnectSignalFinished(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialog*)(ptr)), &QDialog::finished, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Finished);
}

void QDialog_ConnectSignalRejected(QtObjectPtr ptr)
{
	QObject::connect(((QDialog*)(ptr)), &QDialog::rejected, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Rejected, Qt::QueuedConnection);
}

void QDialog_DisconnectSignalRejected(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialog*)(ptr)), &QDialog::rejected, ((MyQDialog*)(ptr)), &MyQDialog::Signal_Rejected);
}

