#include "qwaitcondition.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWaitCondition>
#include "_cgo_export.h"

class MyQWaitCondition: public QWaitCondition {
public:
};

void* QWaitCondition_NewQWaitCondition(){
	return new QWaitCondition();
}

void QWaitCondition_WakeAll(void* ptr){
	static_cast<QWaitCondition*>(ptr)->wakeAll();
}

void QWaitCondition_WakeOne(void* ptr){
	static_cast<QWaitCondition*>(ptr)->wakeOne();
}

void QWaitCondition_DestroyQWaitCondition(void* ptr){
	static_cast<QWaitCondition*>(ptr)->~QWaitCondition();
}

