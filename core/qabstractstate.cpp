#include "qabstractstate.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractState>
#include "_cgo_export.h"

class MyQAbstractState: public QAbstractState {
public:
void Signal_ActiveChanged(bool active){callbackQAbstractStateActiveChanged(this->objectName().toUtf8().data(), active);};
void Signal_Entered(){callbackQAbstractStateEntered(this->objectName().toUtf8().data());};
void Signal_Exited(){callbackQAbstractStateExited(this->objectName().toUtf8().data());};
};

int QAbstractState_Active(void* ptr){
	return static_cast<QAbstractState*>(ptr)->active();
}

void QAbstractState_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_ConnectEntered(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

void QAbstractState_DisconnectEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

void QAbstractState_ConnectExited(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

void QAbstractState_DisconnectExited(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

void* QAbstractState_Machine(void* ptr){
	return static_cast<QAbstractState*>(ptr)->machine();
}

void* QAbstractState_ParentState(void* ptr){
	return static_cast<QAbstractState*>(ptr)->parentState();
}

void QAbstractState_DestroyQAbstractState(void* ptr){
	static_cast<QAbstractState*>(ptr)->~QAbstractState();
}

