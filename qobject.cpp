#include "qobject.h"
#include <QObject>
#include "cgoexport.h"

//MyQObject
class MyQObject: public QObject {
Q_OBJECT
public:
void Signal_Destroyed() { callbackQObject(this, QString("destroyed").toUtf8().data()); };
void Signal_ObjectNameChanged() { callbackQObject(this, QString("objectNameChanged").toUtf8().data()); };

signals:
void Slot_DeleteLater();

};
#include "qobject.moc"


//Public Functions
QtObjectPtr QObject_New_QObject(QtObjectPtr parent)
{
	return new QObject(((QObject*)(parent)));
}

void QObject_Destroy(QtObjectPtr ptr)
{
	((QObject*)(ptr))->~QObject();
}

int QObject_BlockSignals_Bool(QtObjectPtr ptr, int block)
{
	return ((QObject*)(ptr))->blockSignals(block != 0);
}

int QObject_Disconnect_String_QObject_String(QtObjectPtr ptr, char* signal, QtObjectPtr receiver, char* method)
{
	return ((QObject*)(ptr))->disconnect(signal, ((QObject*)(receiver)), method);
}

int QObject_Disconnect_QObject_String(QtObjectPtr ptr, QtObjectPtr receiver, char* method)
{
	return ((QObject*)(ptr))->disconnect(((QObject*)(receiver)), method);
}

void QObject_DumpObjectInfo(QtObjectPtr ptr)
{
	((QObject*)(ptr))->dumpObjectInfo();
}

void QObject_DumpObjectTree(QtObjectPtr ptr)
{
	((QObject*)(ptr))->dumpObjectTree();
}

int QObject_Inherits_String(QtObjectPtr ptr, char* className)
{
	return ((QObject*)(ptr))->inherits(className);
}

void QObject_InstallEventFilter_QObject(QtObjectPtr ptr, QtObjectPtr filterObj)
{
	((QObject*)(ptr))->installEventFilter(((QObject*)(filterObj)));
}

int QObject_IsWidgetType(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->isWidgetType();
}

int QObject_IsWindowType(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->isWindowType();
}

void QObject_KillTimer_Int(QtObjectPtr ptr, int id)
{
	((QObject*)(ptr))->killTimer(id);
}

void QObject_MoveToThread_QThread(QtObjectPtr ptr, QtObjectPtr targetThread)
{
	((QObject*)(ptr))->moveToThread(((QThread*)(targetThread)));
}

char* QObject_ObjectName(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->objectName().toUtf8().data();
}

QtObjectPtr QObject_Parent(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->parent();
}

void QObject_RemoveEventFilter_QObject(QtObjectPtr ptr, QtObjectPtr obj)
{
	((QObject*)(ptr))->removeEventFilter(((QObject*)(obj)));
}

void QObject_SetObjectName_String(QtObjectPtr ptr, char* name)
{
	((QObject*)(ptr))->setObjectName(QString(name));
}

void QObject_SetParent_QObject(QtObjectPtr ptr, QtObjectPtr parent)
{
	((QObject*)(ptr))->setParent(((QObject*)(parent)));
}

int QObject_SignalsBlocked(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->signalsBlocked();
}

int QObject_StartTimer_Int_TimerType(QtObjectPtr ptr, int interval, int timerType)
{
	return ((QObject*)(ptr))->startTimer(interval, ((Qt::TimerType)(timerType)));
}

QtObjectPtr QObject_Thread(QtObjectPtr ptr)
{
	return ((QObject*)(ptr))->thread();
}

//Public Slots
void QObject_ConnectSlotDeleteLater(QtObjectPtr ptr)
{
	QObject::connect(((MyQObject*)(ptr)), &MyQObject::Slot_DeleteLater, ((QObject*)(ptr)), &QObject::deleteLater, Qt::QueuedConnection);
}

void QObject_DisconnectSlotDeleteLater(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQObject*)(ptr)), &MyQObject::Slot_DeleteLater, ((QObject*)(ptr)), &QObject::deleteLater);
}

void QObject_DeleteLater(QtObjectPtr ptr)
{
	((MyQObject*)(ptr))->Slot_DeleteLater();
}

//Signals
void QObject_ConnectSignalDestroyed(QtObjectPtr ptr)
{
	QObject::connect(((QObject*)(ptr)), &QObject::destroyed, ((MyQObject*)(ptr)), &MyQObject::Signal_Destroyed, Qt::QueuedConnection);
}

void QObject_DisconnectSignalDestroyed(QtObjectPtr ptr)
{
	QObject::disconnect(((QObject*)(ptr)), &QObject::destroyed, ((MyQObject*)(ptr)), &MyQObject::Signal_Destroyed);
}

void QObject_ConnectSignalObjectNameChanged(QtObjectPtr ptr)
{
	QObject::connect(((QObject*)(ptr)), &QObject::objectNameChanged, ((MyQObject*)(ptr)), &MyQObject::Signal_ObjectNameChanged, Qt::QueuedConnection);
}

void QObject_DisconnectSignalObjectNameChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QObject*)(ptr)), &QObject::objectNameChanged, ((MyQObject*)(ptr)), &MyQObject::Signal_ObjectNameChanged);
}

//Static Public Members
int QObject_Disconnect_QObject_String_QObject_String(QtObjectPtr sender, char* signal, QtObjectPtr receiver, char* method)
{
	return QObject::disconnect(((QObject*)(sender)), signal, ((QObject*)(receiver)), method);
}

char* QObject_Tr_String_String_Int(char* sourceText, char* disambiguation, int n)
{
	return QObject::tr(sourceText, disambiguation, n).toUtf8().data();
}

