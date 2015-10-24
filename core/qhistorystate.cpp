#include "qhistorystate.h"
#include <QState>
#include <QAbstractState>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHistoryState>
#include "_cgo_export.h"

class MyQHistoryState: public QHistoryState {
public:
void Signal_DefaultStateChanged(){callbackQHistoryStateDefaultStateChanged(this->objectName().toUtf8().data());};
void Signal_HistoryTypeChanged(){callbackQHistoryStateHistoryTypeChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QHistoryState_NewQHistoryState2(int ty, QtObjectPtr parent){
	return new QHistoryState(static_cast<QHistoryState::HistoryType>(ty), static_cast<QState*>(parent));
}

QtObjectPtr QHistoryState_NewQHistoryState(QtObjectPtr parent){
	return new QHistoryState(static_cast<QState*>(parent));
}

QtObjectPtr QHistoryState_DefaultState(QtObjectPtr ptr){
	return static_cast<QHistoryState*>(ptr)->defaultState();
}

void QHistoryState_ConnectDefaultStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

void QHistoryState_DisconnectDefaultStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

int QHistoryState_HistoryType(QtObjectPtr ptr){
	return static_cast<QHistoryState*>(ptr)->historyType();
}

void QHistoryState_ConnectHistoryTypeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_DisconnectHistoryTypeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_SetDefaultState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QHistoryState*>(ptr)->setDefaultState(static_cast<QAbstractState*>(state));
}

void QHistoryState_SetHistoryType(QtObjectPtr ptr, int ty){
	static_cast<QHistoryState*>(ptr)->setHistoryType(static_cast<QHistoryState::HistoryType>(ty));
}

void QHistoryState_DestroyQHistoryState(QtObjectPtr ptr){
	static_cast<QHistoryState*>(ptr)->~QHistoryState();
}

