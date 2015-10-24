#include "qsignalblocker.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSignalBlocker>
#include "_cgo_export.h"

class MyQSignalBlocker: public QSignalBlocker {
public:
};

void QSignalBlocker_Reblock(QtObjectPtr ptr){
	static_cast<QSignalBlocker*>(ptr)->reblock();
}

void QSignalBlocker_Unblock(QtObjectPtr ptr){
	static_cast<QSignalBlocker*>(ptr)->unblock();
}

void QSignalBlocker_DestroyQSignalBlocker(QtObjectPtr ptr){
	static_cast<QSignalBlocker*>(ptr)->~QSignalBlocker();
}

