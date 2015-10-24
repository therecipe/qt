#include "qabstractstate.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QAbstractState>
#include "_cgo_export.h"

class MyQAbstractState: public QAbstractState {
public:
void Signal_ActiveChanged(bool active){callbackQAbstractStateActiveChanged(this->objectName().toUtf8().data(), active);};
void Signal_Entered(){callbackQAbstractStateEntered(this->objectName().toUtf8().data());};
void Signal_Exited(){callbackQAbstractStateExited(this->objectName().toUtf8().data());};
};

int QAbstractState_Active(QtObjectPtr ptr){
	return static_cast<QAbstractState*>(ptr)->active();
}

void QAbstractState_ConnectActiveChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_DisconnectActiveChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_ConnectEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

void QAbstractState_DisconnectEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

void QAbstractState_ConnectExited(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

void QAbstractState_DisconnectExited(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

QtObjectPtr QAbstractState_Machine(QtObjectPtr ptr){
	return static_cast<QAbstractState*>(ptr)->machine();
}

QtObjectPtr QAbstractState_ParentState(QtObjectPtr ptr){
	return static_cast<QAbstractState*>(ptr)->parentState();
}

void QAbstractState_DestroyQAbstractState(QtObjectPtr ptr){
	static_cast<QAbstractState*>(ptr)->~QAbstractState();
}

