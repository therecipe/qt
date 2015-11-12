#include "qqmlincubator.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QQmlIncubator>
#include "_cgo_export.h"

class MyQQmlIncubator: public QQmlIncubator {
public:
};

void* QQmlIncubator_NewQQmlIncubator(int mode){
	return new QQmlIncubator(static_cast<QQmlIncubator::IncubationMode>(mode));
}

void QQmlIncubator_Clear(void* ptr){
	static_cast<QQmlIncubator*>(ptr)->clear();
}

void QQmlIncubator_ForceCompletion(void* ptr){
	static_cast<QQmlIncubator*>(ptr)->forceCompletion();
}

int QQmlIncubator_IncubationMode(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->incubationMode();
}

int QQmlIncubator_IsError(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isError();
}

int QQmlIncubator_IsLoading(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isLoading();
}

int QQmlIncubator_IsNull(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isNull();
}

int QQmlIncubator_IsReady(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isReady();
}

void* QQmlIncubator_Object(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->object();
}

int QQmlIncubator_Status(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->status();
}

