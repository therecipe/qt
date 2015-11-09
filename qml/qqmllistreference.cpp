#include "qqmllistreference.h"
#include <QList>
#include <QList>
#include <QQmlEngine>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlListReference>
#include "_cgo_export.h"

class MyQQmlListReference: public QQmlListReference {
public:
};

void* QQmlListReference_NewQQmlListReference(){
	return new QQmlListReference();
}

void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine){
	return new QQmlListReference(static_cast<QObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
}

int QQmlListReference_Append(void* ptr, void* object){
	return static_cast<QQmlListReference*>(ptr)->append(static_cast<QObject*>(object));
}

void* QQmlListReference_At(void* ptr, int index){
	return static_cast<QQmlListReference*>(ptr)->at(index);
}

int QQmlListReference_CanAppend(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canAppend();
}

int QQmlListReference_CanAt(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canAt();
}

int QQmlListReference_CanClear(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canClear();
}

int QQmlListReference_CanCount(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canCount();
}

int QQmlListReference_Clear(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->clear();
}

int QQmlListReference_Count(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->count();
}

int QQmlListReference_IsManipulable(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isManipulable();
}

int QQmlListReference_IsReadable(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isReadable();
}

int QQmlListReference_IsValid(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isValid();
}

void* QQmlListReference_ListElementType(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QQmlListReference*>(ptr)->listElementType());
}

void* QQmlListReference_Object(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->object();
}

