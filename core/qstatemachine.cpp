#include "qstatemachine.h"
#include <QState>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QAbstractAnimation>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QAbstractState>
#include <QStateMachine>
#include "_cgo_export.h"

class MyQStateMachine: public QStateMachine {
public:
void Signal_RunningChanged(bool running){callbackQStateMachineRunningChanged(this->objectName().toUtf8().data(), running);};
void Signal_Started(){callbackQStateMachineStarted(this->objectName().toUtf8().data());};
void Signal_Stopped(){callbackQStateMachineStopped(this->objectName().toUtf8().data());};
};

QtObjectPtr QStateMachine_NewQStateMachine(QtObjectPtr parent){
	return new QStateMachine(static_cast<QObject*>(parent));
}

QtObjectPtr QStateMachine_NewQStateMachine2(int childMode, QtObjectPtr parent){
	return new QStateMachine(static_cast<QState::ChildMode>(childMode), static_cast<QObject*>(parent));
}

void QStateMachine_AddDefaultAnimation(QtObjectPtr ptr, QtObjectPtr animation){
	static_cast<QStateMachine*>(ptr)->addDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_AddState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QStateMachine*>(ptr)->addState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ClearError(QtObjectPtr ptr){
	static_cast<QStateMachine*>(ptr)->clearError();
}

int QStateMachine_CancelDelayedEvent(QtObjectPtr ptr, int id){
	return static_cast<QStateMachine*>(ptr)->cancelDelayedEvent(id);
}

int QStateMachine_Error(QtObjectPtr ptr){
	return static_cast<QStateMachine*>(ptr)->error();
}

char* QStateMachine_ErrorString(QtObjectPtr ptr){
	return static_cast<QStateMachine*>(ptr)->errorString().toUtf8().data();
}

int QStateMachine_EventFilter(QtObjectPtr ptr, QtObjectPtr watched, QtObjectPtr event){
	return static_cast<QStateMachine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QStateMachine_GlobalRestorePolicy(QtObjectPtr ptr){
	return static_cast<QStateMachine*>(ptr)->globalRestorePolicy();
}

int QStateMachine_IsAnimated(QtObjectPtr ptr){
	return static_cast<QStateMachine*>(ptr)->isAnimated();
}

int QStateMachine_IsRunning(QtObjectPtr ptr){
	return static_cast<QStateMachine*>(ptr)->isRunning();
}

int QStateMachine_PostDelayedEvent(QtObjectPtr ptr, QtObjectPtr event, int delay){
	return static_cast<QStateMachine*>(ptr)->postDelayedEvent(static_cast<QEvent*>(event), delay);
}

void QStateMachine_PostEvent(QtObjectPtr ptr, QtObjectPtr event, int priority){
	static_cast<QStateMachine*>(ptr)->postEvent(static_cast<QEvent*>(event), static_cast<QStateMachine::EventPriority>(priority));
}

void QStateMachine_RemoveDefaultAnimation(QtObjectPtr ptr, QtObjectPtr animation){
	static_cast<QStateMachine*>(ptr)->removeDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_RemoveState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QStateMachine*>(ptr)->removeState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ConnectRunningChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_DisconnectRunningChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_SetAnimated(QtObjectPtr ptr, int enabled){
	static_cast<QStateMachine*>(ptr)->setAnimated(enabled != 0);
}

void QStateMachine_SetGlobalRestorePolicy(QtObjectPtr ptr, int restorePolicy){
	static_cast<QStateMachine*>(ptr)->setGlobalRestorePolicy(static_cast<QState::RestorePolicy>(restorePolicy));
}

void QStateMachine_SetRunning(QtObjectPtr ptr, int running){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "setRunning", Q_ARG(bool, running != 0));
}

void QStateMachine_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "start");
}

void QStateMachine_ConnectStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_DisconnectStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "stop");
}

void QStateMachine_ConnectStopped(QtObjectPtr ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DisconnectStopped(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DestroyQStateMachine(QtObjectPtr ptr){
	static_cast<QStateMachine*>(ptr)->~QStateMachine();
}

