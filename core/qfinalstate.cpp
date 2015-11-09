#include "qfinalstate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QState>
#include <QFinalState>
#include "_cgo_export.h"

class MyQFinalState: public QFinalState {
public:
};

void* QFinalState_NewQFinalState(void* parent){
	return new QFinalState(static_cast<QState*>(parent));
}

void QFinalState_DestroyQFinalState(void* ptr){
	static_cast<QFinalState*>(ptr)->~QFinalState();
}

