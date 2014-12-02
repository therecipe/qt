#include "qthread.h"
#include <QThread>
#include "cgoexport.h"

//MyQThread
class MyQThread: public QThread {
Q_OBJECT
public:
void Signal_Finished() { callbackQThread(this, QString("finished").toUtf8().data()); };
void Signal_Started() { callbackQThread(this, QString("started").toUtf8().data()); };

signals:
void Slot_Quit();
void Slot_Terminate();

};
#include "qthread.moc"

//Public Types
int QThread_IdlePriority() { return QThread::IdlePriority; }
int QThread_LowestPriority() { return QThread::LowestPriority; }
int QThread_LowPriority() { return QThread::LowPriority; }
int QThread_NormalPriority() { return QThread::NormalPriority; }
int QThread_HighPriority() { return QThread::HighPriority; }
int QThread_HighestPriority() { return QThread::HighestPriority; }
int QThread_TimeCriticalPriority() { return QThread::TimeCriticalPriority; }
int QThread_InheritPriority() { return QThread::InheritPriority; }

//Public Functions
QtObjectPtr QThread_New_QObject(QtObjectPtr parent)
{
	return new QThread(((QObject*)(parent)));
}

void QThread_Destroy(QtObjectPtr ptr)
{
	((QThread*)(ptr))->~QThread();
}

void QThread_Exit_Int(QtObjectPtr ptr, int returnCode)
{
	((QThread*)(ptr))->exit(returnCode);
}

int QThread_IsFinished(QtObjectPtr ptr)
{
	return ((QThread*)(ptr))->isFinished();
}

int QThread_IsInterruptionRequested(QtObjectPtr ptr)
{
	return ((QThread*)(ptr))->isInterruptionRequested();
}

int QThread_IsRunning(QtObjectPtr ptr)
{
	return ((QThread*)(ptr))->isRunning();
}

int QThread_Priority(QtObjectPtr ptr)
{
	return ((QThread*)(ptr))->priority();
}

void QThread_RequestInterruption(QtObjectPtr ptr)
{
	((QThread*)(ptr))->requestInterruption();
}

void QThread_SetPriority_Priority(QtObjectPtr ptr, int priority)
{
	((QThread*)(ptr))->setPriority(((QThread::Priority)(priority)));
}

//Public Slots
void QThread_ConnectSlotQuit(QtObjectPtr ptr)
{
	QObject::connect(((MyQThread*)(ptr)), &MyQThread::Slot_Quit, ((QThread*)(ptr)), &QThread::quit, Qt::QueuedConnection);
}

void QThread_DisconnectSlotQuit(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQThread*)(ptr)), &MyQThread::Slot_Quit, ((QThread*)(ptr)), &QThread::quit);
}

void QThread_Quit(QtObjectPtr ptr)
{
	((MyQThread*)(ptr))->Slot_Quit();
}

void QThread_ConnectSlotTerminate(QtObjectPtr ptr)
{
	QObject::connect(((MyQThread*)(ptr)), &MyQThread::Slot_Terminate, ((QThread*)(ptr)), &QThread::terminate, Qt::QueuedConnection);
}

void QThread_DisconnectSlotTerminate(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQThread*)(ptr)), &MyQThread::Slot_Terminate, ((QThread*)(ptr)), &QThread::terminate);
}

void QThread_Terminate(QtObjectPtr ptr)
{
	((MyQThread*)(ptr))->Slot_Terminate();
}

//Signals
void QThread_ConnectSignalFinished(QtObjectPtr ptr)
{
	QObject::connect(((QThread*)(ptr)), &QThread::finished, ((MyQThread*)(ptr)), &MyQThread::Signal_Finished, Qt::QueuedConnection);
}

void QThread_DisconnectSignalFinished(QtObjectPtr ptr)
{
	QObject::disconnect(((QThread*)(ptr)), &QThread::finished, ((MyQThread*)(ptr)), &MyQThread::Signal_Finished);
}

void QThread_ConnectSignalStarted(QtObjectPtr ptr)
{
	QObject::connect(((QThread*)(ptr)), &QThread::started, ((MyQThread*)(ptr)), &MyQThread::Signal_Started, Qt::QueuedConnection);
}

void QThread_DisconnectSignalStarted(QtObjectPtr ptr)
{
	QObject::disconnect(((QThread*)(ptr)), &QThread::started, ((MyQThread*)(ptr)), &MyQThread::Signal_Started);
}

//Static Public Members
QtObjectPtr QThread_CurrentThread()
{
	return QThread::currentThread();
}

int QThread_IdealThreadCount()
{
	return QThread::idealThreadCount();
}

