#include "qstatemachine.h"
#include <QString>
#include <QEvent>
#include <QAbstractState>
#include <QAbstractAnimation>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QState>
#include <QMetaObject>
#include <QStateMachine>
#include "_cgo_export.h"

class MyQStateMachine: public QStateMachine {
public:
void Signal_RunningChanged(bool running){callbackQStateMachineRunningChanged(this->objectName().toUtf8().data(), running);};
void Signal_Started(){callbackQStateMachineStarted(this->objectName().toUtf8().data());};
void Signal_Stopped(){callbackQStateMachineStopped(this->objectName().toUtf8().data());};
};

void* QStateMachine_NewQStateMachine(void* parent){
	return new QStateMachine(static_cast<QObject*>(parent));
}

void* QStateMachine_NewQStateMachine2(int childMode, void* parent){
	return new QStateMachine(static_cast<QState::ChildMode>(childMode), static_cast<QObject*>(parent));
}

void QStateMachine_AddDefaultAnimation(void* ptr, void* animation){
	static_cast<QStateMachine*>(ptr)->addDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_AddState(void* ptr, void* state){
	static_cast<QStateMachine*>(ptr)->addState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ClearError(void* ptr){
	static_cast<QStateMachine*>(ptr)->clearError();
}

int QStateMachine_CancelDelayedEvent(void* ptr, int id){
	return static_cast<QStateMachine*>(ptr)->cancelDelayedEvent(id);
}

int QStateMachine_Error(void* ptr){
	return static_cast<QStateMachine*>(ptr)->error();
}

char* QStateMachine_ErrorString(void* ptr){
	return static_cast<QStateMachine*>(ptr)->errorString().toUtf8().data();
}

int QStateMachine_EventFilter(void* ptr, void* watched, void* event){
	return static_cast<QStateMachine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QStateMachine_GlobalRestorePolicy(void* ptr){
	return static_cast<QStateMachine*>(ptr)->globalRestorePolicy();
}

int QStateMachine_IsAnimated(void* ptr){
	return static_cast<QStateMachine*>(ptr)->isAnimated();
}

int QStateMachine_IsRunning(void* ptr){
	return static_cast<QStateMachine*>(ptr)->isRunning();
}

int QStateMachine_PostDelayedEvent(void* ptr, void* event, int delay){
	return static_cast<QStateMachine*>(ptr)->postDelayedEvent(static_cast<QEvent*>(event), delay);
}

void QStateMachine_PostEvent(void* ptr, void* event, int priority){
	static_cast<QStateMachine*>(ptr)->postEvent(static_cast<QEvent*>(event), static_cast<QStateMachine::EventPriority>(priority));
}

void QStateMachine_RemoveDefaultAnimation(void* ptr, void* animation){
	static_cast<QStateMachine*>(ptr)->removeDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_RemoveState(void* ptr, void* state){
	static_cast<QStateMachine*>(ptr)->removeState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ConnectRunningChanged(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_DisconnectRunningChanged(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_SetAnimated(void* ptr, int enabled){
	static_cast<QStateMachine*>(ptr)->setAnimated(enabled != 0);
}

void QStateMachine_SetGlobalRestorePolicy(void* ptr, int restorePolicy){
	static_cast<QStateMachine*>(ptr)->setGlobalRestorePolicy(static_cast<QState::RestorePolicy>(restorePolicy));
}

void QStateMachine_SetRunning(void* ptr, int running){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "setRunning", Q_ARG(bool, running != 0));
}

void QStateMachine_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "start");
}

void QStateMachine_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "stop");
}

void QStateMachine_ConnectStopped(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DisconnectStopped(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DestroyQStateMachine(void* ptr){
	static_cast<QStateMachine*>(ptr)->~QStateMachine();
}

