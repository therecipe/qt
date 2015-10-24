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

QtObjectPtr QWaitCondition_NewQWaitCondition(){
	return new QWaitCondition();
}

void QWaitCondition_WakeAll(QtObjectPtr ptr){
	static_cast<QWaitCondition*>(ptr)->wakeAll();
}

void QWaitCondition_WakeOne(QtObjectPtr ptr){
	static_cast<QWaitCondition*>(ptr)->wakeOne();
}

void QWaitCondition_DestroyQWaitCondition(QtObjectPtr ptr){
	static_cast<QWaitCondition*>(ptr)->~QWaitCondition();
}

