#include "qstatusbar.h"
#include <QStatusBar>
#include "cgoexport.h"

//MyQStatusBar
class MyQStatusBar: public QStatusBar {
Q_OBJECT
public:
void Signal_MessageChanged() { callbackQStatusBar(this, QString("messageChanged").toUtf8().data()); };

signals:
void Slot_ClearMessage();

};
#include "qstatusbar.moc"


//Public Functions
QtObjectPtr QStatusBar_New_QWidget(QtObjectPtr parent)
{
	return new QStatusBar(((QWidget*)(parent)));
}

void QStatusBar_Destroy(QtObjectPtr ptr)
{
	((QStatusBar*)(ptr))->~QStatusBar();
}

void QStatusBar_AddPermanentWidget_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch)
{
	((QStatusBar*)(ptr))->addPermanentWidget(((QWidget*)(widget)), stretch);
}

void QStatusBar_AddWidget_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch)
{
	((QStatusBar*)(ptr))->addWidget(((QWidget*)(widget)), stretch);
}

char* QStatusBar_CurrentMessage(QtObjectPtr ptr)
{
	return ((QStatusBar*)(ptr))->currentMessage().toUtf8().data();
}

int QStatusBar_InsertPermanentWidget_Int_QWidget_Int(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch)
{
	return ((QStatusBar*)(ptr))->insertPermanentWidget(index, ((QWidget*)(widget)), stretch);
}

int QStatusBar_InsertWidget_Int_QWidget_Int(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch)
{
	return ((QStatusBar*)(ptr))->insertWidget(index, ((QWidget*)(widget)), stretch);
}

int QStatusBar_IsSizeGripEnabled(QtObjectPtr ptr)
{
	return ((QStatusBar*)(ptr))->isSizeGripEnabled();
}

void QStatusBar_RemoveWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QStatusBar*)(ptr))->removeWidget(((QWidget*)(widget)));
}

void QStatusBar_SetSizeGripEnabled_Bool(QtObjectPtr ptr, int enabled)
{
	((QStatusBar*)(ptr))->setSizeGripEnabled(enabled != 0);
}

//Public Slots
void QStatusBar_ConnectSlotClearMessage(QtObjectPtr ptr)
{
	QObject::connect(((MyQStatusBar*)(ptr)), &MyQStatusBar::Slot_ClearMessage, ((QStatusBar*)(ptr)), &QStatusBar::clearMessage, Qt::QueuedConnection);
}

void QStatusBar_DisconnectSlotClearMessage(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQStatusBar*)(ptr)), &MyQStatusBar::Slot_ClearMessage, ((QStatusBar*)(ptr)), &QStatusBar::clearMessage);
}

void QStatusBar_ClearMessage(QtObjectPtr ptr)
{
	((MyQStatusBar*)(ptr))->Slot_ClearMessage();
}

//Signals
void QStatusBar_ConnectSignalMessageChanged(QtObjectPtr ptr)
{
	QObject::connect(((QStatusBar*)(ptr)), &QStatusBar::messageChanged, ((MyQStatusBar*)(ptr)), &MyQStatusBar::Signal_MessageChanged, Qt::QueuedConnection);
}

void QStatusBar_DisconnectSignalMessageChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QStatusBar*)(ptr)), &QStatusBar::messageChanged, ((MyQStatusBar*)(ptr)), &MyQStatusBar::Signal_MessageChanged);
}

