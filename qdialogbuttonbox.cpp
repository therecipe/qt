#include "qdialogbuttonbox.h"
#include <QDialogButtonBox>
#include "cgoexport.h"

//MyQDialogButtonBox
class MyQDialogButtonBox: public QDialogButtonBox {
Q_OBJECT
public:
void Signal_Accepted() { callbackQDialogButtonBox(this, QString("accepted").toUtf8().data()); };
void Signal_HelpRequested() { callbackQDialogButtonBox(this, QString("helpRequested").toUtf8().data()); };
void Signal_Rejected() { callbackQDialogButtonBox(this, QString("rejected").toUtf8().data()); };

signals:

};
#include "qdialogbuttonbox.moc"

//Public Types
int QDialogButtonBox_WinLayout() { return QDialogButtonBox::WinLayout; }
int QDialogButtonBox_MacLayout() { return QDialogButtonBox::MacLayout; }
int QDialogButtonBox_KdeLayout() { return QDialogButtonBox::KdeLayout; }
int QDialogButtonBox_GnomeLayout() { return QDialogButtonBox::GnomeLayout; }
int QDialogButtonBox_InvalidRole() { return QDialogButtonBox::InvalidRole; }
int QDialogButtonBox_AcceptRole() { return QDialogButtonBox::AcceptRole; }
int QDialogButtonBox_RejectRole() { return QDialogButtonBox::RejectRole; }
int QDialogButtonBox_DestructiveRole() { return QDialogButtonBox::DestructiveRole; }
int QDialogButtonBox_ActionRole() { return QDialogButtonBox::ActionRole; }
int QDialogButtonBox_HelpRole() { return QDialogButtonBox::HelpRole; }
int QDialogButtonBox_YesRole() { return QDialogButtonBox::YesRole; }
int QDialogButtonBox_NoRole() { return QDialogButtonBox::NoRole; }
int QDialogButtonBox_ApplyRole() { return QDialogButtonBox::ApplyRole; }
int QDialogButtonBox_ResetRole() { return QDialogButtonBox::ResetRole; }
int QDialogButtonBox_Ok() { return QDialogButtonBox::Ok; }
int QDialogButtonBox_Open() { return QDialogButtonBox::Open; }
int QDialogButtonBox_Save() { return QDialogButtonBox::Save; }
int QDialogButtonBox_Cancel() { return QDialogButtonBox::Cancel; }
int QDialogButtonBox_Close() { return QDialogButtonBox::Close; }
int QDialogButtonBox_Discard() { return QDialogButtonBox::Discard; }
int QDialogButtonBox_Apply() { return QDialogButtonBox::Apply; }
int QDialogButtonBox_Reset() { return QDialogButtonBox::Reset; }
int QDialogButtonBox_RestoreDefaults() { return QDialogButtonBox::RestoreDefaults; }
int QDialogButtonBox_Help() { return QDialogButtonBox::Help; }
int QDialogButtonBox_SaveAll() { return QDialogButtonBox::SaveAll; }
int QDialogButtonBox_Yes() { return QDialogButtonBox::Yes; }
int QDialogButtonBox_YesToAll() { return QDialogButtonBox::YesToAll; }
int QDialogButtonBox_No() { return QDialogButtonBox::No; }
int QDialogButtonBox_NoToAll() { return QDialogButtonBox::NoToAll; }
int QDialogButtonBox_Abort() { return QDialogButtonBox::Abort; }
int QDialogButtonBox_Retry() { return QDialogButtonBox::Retry; }
int QDialogButtonBox_Ignore() { return QDialogButtonBox::Ignore; }

//Public Functions
QtObjectPtr QDialogButtonBox_New_QWidget(QtObjectPtr parent)
{
	return new QDialogButtonBox(((QWidget*)(parent)));
}

QtObjectPtr QDialogButtonBox_New_Orientation_QWidget(int orientation, QtObjectPtr parent)
{
	return new QDialogButtonBox(((Qt::Orientation)(orientation)), ((QWidget*)(parent)));
}

QtObjectPtr QDialogButtonBox_New_StandardButton_QWidget(int buttons, QtObjectPtr parent)
{
	return new QDialogButtonBox(((QDialogButtonBox::StandardButton)(buttons)), ((QWidget*)(parent)));
}

QtObjectPtr QDialogButtonBox_New_StandardButton_Orientation_QWidget(int buttons, int orientation, QtObjectPtr parent)
{
	return new QDialogButtonBox(((QDialogButtonBox::StandardButton)(buttons)), ((Qt::Orientation)(orientation)), ((QWidget*)(parent)));
}

void QDialogButtonBox_Destroy(QtObjectPtr ptr)
{
	((QDialogButtonBox*)(ptr))->~QDialogButtonBox();
}

void QDialogButtonBox_AddButton_QAbstractButton_ButtonRole(QtObjectPtr ptr, QtObjectPtr button, int role)
{
	((QDialogButtonBox*)(ptr))->addButton(((QAbstractButton*)(button)), ((QDialogButtonBox::ButtonRole)(role)));
}

QtObjectPtr QDialogButtonBox_AddButton_String_ButtonRole(QtObjectPtr ptr, char* text, int role)
{
	return ((QDialogButtonBox*)(ptr))->addButton(QString(text), ((QDialogButtonBox::ButtonRole)(role)));
}

QtObjectPtr QDialogButtonBox_AddButton_StandardButton(QtObjectPtr ptr, int button)
{
	return ((QDialogButtonBox*)(ptr))->addButton(((QDialogButtonBox::StandardButton)(button)));
}

QtObjectPtr QDialogButtonBox_Button_StandardButton(QtObjectPtr ptr, int which)
{
	return ((QDialogButtonBox*)(ptr))->button(((QDialogButtonBox::StandardButton)(which)));
}

int QDialogButtonBox_ButtonRole_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button)
{
	return ((QDialogButtonBox*)(ptr))->buttonRole(((QAbstractButton*)(button)));
}

int QDialogButtonBox_CenterButtons(QtObjectPtr ptr)
{
	return ((QDialogButtonBox*)(ptr))->centerButtons();
}

void QDialogButtonBox_Clear(QtObjectPtr ptr)
{
	((QDialogButtonBox*)(ptr))->clear();
}

int QDialogButtonBox_Orientation(QtObjectPtr ptr)
{
	return ((QDialogButtonBox*)(ptr))->orientation();
}

void QDialogButtonBox_RemoveButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button)
{
	((QDialogButtonBox*)(ptr))->removeButton(((QAbstractButton*)(button)));
}

void QDialogButtonBox_SetCenterButtons_Bool(QtObjectPtr ptr, int center)
{
	((QDialogButtonBox*)(ptr))->setCenterButtons(center != 0);
}

void QDialogButtonBox_SetOrientation_Orientation(QtObjectPtr ptr, int orientation)
{
	((QDialogButtonBox*)(ptr))->setOrientation(((Qt::Orientation)(orientation)));
}

void QDialogButtonBox_SetStandardButtons_StandardButton(QtObjectPtr ptr, int buttons)
{
	((QDialogButtonBox*)(ptr))->setStandardButtons(((QDialogButtonBox::StandardButton)(buttons)));
}

int QDialogButtonBox_StandardButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button)
{
	return ((QDialogButtonBox*)(ptr))->standardButton(((QAbstractButton*)(button)));
}

int QDialogButtonBox_StandardButtons(QtObjectPtr ptr)
{
	return ((QDialogButtonBox*)(ptr))->standardButtons();
}

//Signals
void QDialogButtonBox_ConnectSignalAccepted(QtObjectPtr ptr)
{
	QObject::connect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::accepted, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_Accepted, Qt::QueuedConnection);
}

void QDialogButtonBox_DisconnectSignalAccepted(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::accepted, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_Accepted);
}

void QDialogButtonBox_ConnectSignalHelpRequested(QtObjectPtr ptr)
{
	QObject::connect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::helpRequested, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_HelpRequested, Qt::QueuedConnection);
}

void QDialogButtonBox_DisconnectSignalHelpRequested(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::helpRequested, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_HelpRequested);
}

void QDialogButtonBox_ConnectSignalRejected(QtObjectPtr ptr)
{
	QObject::connect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::rejected, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_Rejected, Qt::QueuedConnection);
}

void QDialogButtonBox_DisconnectSignalRejected(QtObjectPtr ptr)
{
	QObject::disconnect(((QDialogButtonBox*)(ptr)), &QDialogButtonBox::rejected, ((MyQDialogButtonBox*)(ptr)), &MyQDialogButtonBox::Signal_Rejected);
}

