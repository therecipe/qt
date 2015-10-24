#include "qqmlincubator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlIncubator>
#include "_cgo_export.h"

class MyQQmlIncubator: public QQmlIncubator {
public:
};

QtObjectPtr QQmlIncubator_NewQQmlIncubator(int mode){
	return new QQmlIncubator(static_cast<QQmlIncubator::IncubationMode>(mode));
}

void QQmlIncubator_Clear(QtObjectPtr ptr){
	static_cast<QQmlIncubator*>(ptr)->clear();
}

void QQmlIncubator_ForceCompletion(QtObjectPtr ptr){
	static_cast<QQmlIncubator*>(ptr)->forceCompletion();
}

int QQmlIncubator_IncubationMode(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->incubationMode();
}

int QQmlIncubator_IsError(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->isError();
}

int QQmlIncubator_IsLoading(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->isLoading();
}

int QQmlIncubator_IsNull(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->isNull();
}

int QQmlIncubator_IsReady(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->isReady();
}

QtObjectPtr QQmlIncubator_Object(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->object();
}

int QQmlIncubator_Status(QtObjectPtr ptr){
	return static_cast<QQmlIncubator*>(ptr)->status();
}

