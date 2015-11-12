#include "qrunnable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRunnable>
#include "_cgo_export.h"

class MyQRunnable: public QRunnable {
public:
};

int QRunnable_AutoDelete(void* ptr){
	return static_cast<QRunnable*>(ptr)->autoDelete();
}

void QRunnable_Run(void* ptr){
	static_cast<QRunnable*>(ptr)->run();
}

void QRunnable_SetAutoDelete(void* ptr, int autoDelete){
	static_cast<QRunnable*>(ptr)->setAutoDelete(autoDelete != 0);
}

void QRunnable_DestroyQRunnable(void* ptr){
	static_cast<QRunnable*>(ptr)->~QRunnable();
}

