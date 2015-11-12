#include "qstate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractState>
#include <QObject>
#include <QAbstractTransition>
#include <QState>
#include "_cgo_export.h"

class MyQState: public QState {
public:
void Signal_ChildModeChanged(){callbackQStateChildModeChanged(this->objectName().toUtf8().data());};
void Signal_ErrorStateChanged(){callbackQStateErrorStateChanged(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQStateFinished(this->objectName().toUtf8().data());};
void Signal_InitialStateChanged(){callbackQStateInitialStateChanged(this->objectName().toUtf8().data());};
void Signal_PropertiesAssigned(){callbackQStatePropertiesAssigned(this->objectName().toUtf8().data());};
};

void* QState_NewQState2(int childMode, void* parent){
	return new QState(static_cast<QState::ChildMode>(childMode), static_cast<QState*>(parent));
}

void* QState_NewQState(void* parent){
	return new QState(static_cast<QState*>(parent));
}

void* QState_AddTransition3(void* ptr, void* target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractState*>(target));
}

void* QState_AddTransition2(void* ptr, void* sender, char* signal, void* target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QAbstractState*>(target));
}

void QState_AddTransition(void* ptr, void* transition){
	static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_AssignProperty(void* ptr, void* object, char* name, void* value){
	static_cast<QState*>(ptr)->assignProperty(static_cast<QObject*>(object), const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

int QState_ChildMode(void* ptr){
	return static_cast<QState*>(ptr)->childMode();
}

void QState_ConnectChildModeChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

void QState_DisconnectChildModeChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

void* QState_ErrorState(void* ptr){
	return static_cast<QState*>(ptr)->errorState();
}

void QState_ConnectErrorStateChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

void QState_DisconnectErrorStateChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

void QState_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

void QState_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

void* QState_InitialState(void* ptr){
	return static_cast<QState*>(ptr)->initialState();
}

void QState_ConnectInitialStateChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_DisconnectInitialStateChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_ConnectPropertiesAssigned(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_DisconnectPropertiesAssigned(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_RemoveTransition(void* ptr, void* transition){
	static_cast<QState*>(ptr)->removeTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_SetChildMode(void* ptr, int mode){
	static_cast<QState*>(ptr)->setChildMode(static_cast<QState::ChildMode>(mode));
}

void QState_SetErrorState(void* ptr, void* state){
	static_cast<QState*>(ptr)->setErrorState(static_cast<QAbstractState*>(state));
}

void QState_SetInitialState(void* ptr, void* state){
	static_cast<QState*>(ptr)->setInitialState(static_cast<QAbstractState*>(state));
}

void QState_DestroyQState(void* ptr){
	static_cast<QState*>(ptr)->~QState();
}

