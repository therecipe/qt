#include "qfinalstate.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QState>
#include <QString>
#include <QFinalState>
#include "_cgo_export.h"

class MyQFinalState: public QFinalState {
public:
};

QtObjectPtr QFinalState_NewQFinalState(QtObjectPtr parent){
	return new QFinalState(static_cast<QState*>(parent));
}

void QFinalState_DestroyQFinalState(QtObjectPtr ptr){
	static_cast<QFinalState*>(ptr)->~QFinalState();
}

