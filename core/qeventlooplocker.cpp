#include "qeventlooplocker.h"
#include <QThread>
#include <QEventLoop>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QEventLoopLocker>
#include "_cgo_export.h"

class MyQEventLoopLocker: public QEventLoopLocker {
public:
};

QtObjectPtr QEventLoopLocker_NewQEventLoopLocker(){
	return new QEventLoopLocker();
}

QtObjectPtr QEventLoopLocker_NewQEventLoopLocker2(QtObjectPtr loop){
	return new QEventLoopLocker(static_cast<QEventLoop*>(loop));
}

QtObjectPtr QEventLoopLocker_NewQEventLoopLocker3(QtObjectPtr thread){
	return new QEventLoopLocker(static_cast<QThread*>(thread));
}

void QEventLoopLocker_DestroyQEventLoopLocker(QtObjectPtr ptr){
	static_cast<QEventLoopLocker*>(ptr)->~QEventLoopLocker();
}

