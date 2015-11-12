#include "qdebugstatesaver.h"
#include <QModelIndex>
#include <QDebug>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDebugStateSaver>
#include "_cgo_export.h"

class MyQDebugStateSaver: public QDebugStateSaver {
public:
};

void* QDebugStateSaver_NewQDebugStateSaver(void* dbg){
	return new QDebugStateSaver(*static_cast<QDebug*>(dbg));
}

void QDebugStateSaver_DestroyQDebugStateSaver(void* ptr){
	static_cast<QDebugStateSaver*>(ptr)->~QDebugStateSaver();
}

