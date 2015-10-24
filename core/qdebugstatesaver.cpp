#include "qdebugstatesaver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDebug>
#include <QDebugStateSaver>
#include "_cgo_export.h"

class MyQDebugStateSaver: public QDebugStateSaver {
public:
};

QtObjectPtr QDebugStateSaver_NewQDebugStateSaver(QtObjectPtr dbg){
	return new QDebugStateSaver(*static_cast<QDebug*>(dbg));
}

void QDebugStateSaver_DestroyQDebugStateSaver(QtObjectPtr ptr){
	static_cast<QDebugStateSaver*>(ptr)->~QDebugStateSaver();
}

