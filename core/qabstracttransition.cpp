#include "qabstracttransition.h"
#include <QUrl>
#include <QModelIndex>
#include <QAbstractAnimation>
#include <QObject>
#include <QAbstractState>
#include <QString>
#include <QVariant>
#include <QAbstractTransition>
#include "_cgo_export.h"

class MyQAbstractTransition: public QAbstractTransition {
public:
void Signal_TargetStateChanged(){callbackQAbstractTransitionTargetStateChanged(this->objectName().toUtf8().data());};
void Signal_TargetStatesChanged(){callbackQAbstractTransitionTargetStatesChanged(this->objectName().toUtf8().data());};
void Signal_Triggered(){callbackQAbstractTransitionTriggered(this->objectName().toUtf8().data());};
};

void QAbstractTransition_AddAnimation(void* ptr, void* animation){
	static_cast<QAbstractTransition*>(ptr)->addAnimation(static_cast<QAbstractAnimation*>(animation));
}

void* QAbstractTransition_Machine(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->machine();
}

void QAbstractTransition_RemoveAnimation(void* ptr, void* animation){
	static_cast<QAbstractTransition*>(ptr)->removeAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QAbstractTransition_SetTargetState(void* ptr, void* target){
	static_cast<QAbstractTransition*>(ptr)->setTargetState(static_cast<QAbstractState*>(target));
}

void QAbstractTransition_SetTransitionType(void* ptr, int ty){
	static_cast<QAbstractTransition*>(ptr)->setTransitionType(static_cast<QAbstractTransition::TransitionType>(ty));
}

void* QAbstractTransition_SourceState(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->sourceState();
}

void* QAbstractTransition_TargetState(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->targetState();
}

void QAbstractTransition_ConnectTargetStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_DisconnectTargetStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_ConnectTargetStatesChanged(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

void QAbstractTransition_DisconnectTargetStatesChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

int QAbstractTransition_TransitionType(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->transitionType();
}

void QAbstractTransition_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DestroyQAbstractTransition(void* ptr){
	static_cast<QAbstractTransition*>(ptr)->~QAbstractTransition();
}

