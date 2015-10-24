#include "qstate.h"
#include <QAbstractState>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

QtObjectPtr QState_NewQState2(int childMode, QtObjectPtr parent){
	return new QState(static_cast<QState::ChildMode>(childMode), static_cast<QState*>(parent));
}

QtObjectPtr QState_NewQState(QtObjectPtr parent){
	return new QState(static_cast<QState*>(parent));
}

QtObjectPtr QState_AddTransition3(QtObjectPtr ptr, QtObjectPtr target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractState*>(target));
}

QtObjectPtr QState_AddTransition2(QtObjectPtr ptr, QtObjectPtr sender, char* signal, QtObjectPtr target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QAbstractState*>(target));
}

void QState_AddTransition(QtObjectPtr ptr, QtObjectPtr transition){
	static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_AssignProperty(QtObjectPtr ptr, QtObjectPtr object, char* name, char* value){
	static_cast<QState*>(ptr)->assignProperty(static_cast<QObject*>(object), const_cast<const char*>(name), QVariant(value));
}

int QState_ChildMode(QtObjectPtr ptr){
	return static_cast<QState*>(ptr)->childMode();
}

void QState_ConnectChildModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

void QState_DisconnectChildModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

QtObjectPtr QState_ErrorState(QtObjectPtr ptr){
	return static_cast<QState*>(ptr)->errorState();
}

void QState_ConnectErrorStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

void QState_DisconnectErrorStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

void QState_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

void QState_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

QtObjectPtr QState_InitialState(QtObjectPtr ptr){
	return static_cast<QState*>(ptr)->initialState();
}

void QState_ConnectInitialStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_DisconnectInitialStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_ConnectPropertiesAssigned(QtObjectPtr ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_DisconnectPropertiesAssigned(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_RemoveTransition(QtObjectPtr ptr, QtObjectPtr transition){
	static_cast<QState*>(ptr)->removeTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_SetChildMode(QtObjectPtr ptr, int mode){
	static_cast<QState*>(ptr)->setChildMode(static_cast<QState::ChildMode>(mode));
}

void QState_SetErrorState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QState*>(ptr)->setErrorState(static_cast<QAbstractState*>(state));
}

void QState_SetInitialState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QState*>(ptr)->setInitialState(static_cast<QAbstractState*>(state));
}

void QState_DestroyQState(QtObjectPtr ptr){
	static_cast<QState*>(ptr)->~QState();
}

