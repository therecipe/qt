#include "qhistorystate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QState>
#include <QAbstractState>
#include <QObject>
#include <QHistoryState>
#include "_cgo_export.h"

class MyQHistoryState: public QHistoryState {
public:
void Signal_DefaultStateChanged(){callbackQHistoryStateDefaultStateChanged(this->objectName().toUtf8().data());};
void Signal_HistoryTypeChanged(){callbackQHistoryStateHistoryTypeChanged(this->objectName().toUtf8().data());};
};

void* QHistoryState_NewQHistoryState2(int ty, void* parent){
	return new QHistoryState(static_cast<QHistoryState::HistoryType>(ty), static_cast<QState*>(parent));
}

void* QHistoryState_NewQHistoryState(void* parent){
	return new QHistoryState(static_cast<QState*>(parent));
}

void* QHistoryState_DefaultState(void* ptr){
	return static_cast<QHistoryState*>(ptr)->defaultState();
}

void QHistoryState_ConnectDefaultStateChanged(void* ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

void QHistoryState_DisconnectDefaultStateChanged(void* ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

int QHistoryState_HistoryType(void* ptr){
	return static_cast<QHistoryState*>(ptr)->historyType();
}

void QHistoryState_ConnectHistoryTypeChanged(void* ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_DisconnectHistoryTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_SetDefaultState(void* ptr, void* state){
	static_cast<QHistoryState*>(ptr)->setDefaultState(static_cast<QAbstractState*>(state));
}

void QHistoryState_SetHistoryType(void* ptr, int ty){
	static_cast<QHistoryState*>(ptr)->setHistoryType(static_cast<QHistoryState::HistoryType>(ty));
}

void QHistoryState_DestroyQHistoryState(void* ptr){
	static_cast<QHistoryState*>(ptr)->~QHistoryState();
}

