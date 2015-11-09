#include "qeventlooplocker.h"
#include <QThread>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEventLoop>
#include <QEvent>
#include <QEventLoopLocker>
#include "_cgo_export.h"

class MyQEventLoopLocker: public QEventLoopLocker {
public:
};

void* QEventLoopLocker_NewQEventLoopLocker(){
	return new QEventLoopLocker();
}

void* QEventLoopLocker_NewQEventLoopLocker2(void* loop){
	return new QEventLoopLocker(static_cast<QEventLoop*>(loop));
}

void* QEventLoopLocker_NewQEventLoopLocker3(void* thread){
	return new QEventLoopLocker(static_cast<QThread*>(thread));
}

void QEventLoopLocker_DestroyQEventLoopLocker(void* ptr){
	static_cast<QEventLoopLocker*>(ptr)->~QEventLoopLocker();
}

