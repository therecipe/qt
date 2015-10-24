#include "qabstracttransition.h"
#include <QAbstractAnimation>
#include <QAbstractState>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractTransition>
#include "_cgo_export.h"

class MyQAbstractTransition: public QAbstractTransition {
public:
void Signal_TargetStateChanged(){callbackQAbstractTransitionTargetStateChanged(this->objectName().toUtf8().data());};
void Signal_TargetStatesChanged(){callbackQAbstractTransitionTargetStatesChanged(this->objectName().toUtf8().data());};
void Signal_Triggered(){callbackQAbstractTransitionTriggered(this->objectName().toUtf8().data());};
};

void QAbstractTransition_AddAnimation(QtObjectPtr ptr, QtObjectPtr animation){
	static_cast<QAbstractTransition*>(ptr)->addAnimation(static_cast<QAbstractAnimation*>(animation));
}

QtObjectPtr QAbstractTransition_Machine(QtObjectPtr ptr){
	return static_cast<QAbstractTransition*>(ptr)->machine();
}

void QAbstractTransition_RemoveAnimation(QtObjectPtr ptr, QtObjectPtr animation){
	static_cast<QAbstractTransition*>(ptr)->removeAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QAbstractTransition_SetTargetState(QtObjectPtr ptr, QtObjectPtr target){
	static_cast<QAbstractTransition*>(ptr)->setTargetState(static_cast<QAbstractState*>(target));
}

void QAbstractTransition_SetTransitionType(QtObjectPtr ptr, int ty){
	static_cast<QAbstractTransition*>(ptr)->setTransitionType(static_cast<QAbstractTransition::TransitionType>(ty));
}

QtObjectPtr QAbstractTransition_SourceState(QtObjectPtr ptr){
	return static_cast<QAbstractTransition*>(ptr)->sourceState();
}

QtObjectPtr QAbstractTransition_TargetState(QtObjectPtr ptr){
	return static_cast<QAbstractTransition*>(ptr)->targetState();
}

void QAbstractTransition_ConnectTargetStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_DisconnectTargetStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_ConnectTargetStatesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

void QAbstractTransition_DisconnectTargetStatesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

int QAbstractTransition_TransitionType(QtObjectPtr ptr){
	return static_cast<QAbstractTransition*>(ptr)->transitionType();
}

void QAbstractTransition_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DestroyQAbstractTransition(QtObjectPtr ptr){
	static_cast<QAbstractTransition*>(ptr)->~QAbstractTransition();
}

