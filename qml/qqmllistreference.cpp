#include "qqmllistreference.h"
#include <QList>
#include <QQmlEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QList>
#include <QQmlListReference>
#include "_cgo_export.h"

class MyQQmlListReference: public QQmlListReference {
public:
};

QtObjectPtr QQmlListReference_NewQQmlListReference(){
	return new QQmlListReference();
}

QtObjectPtr QQmlListReference_NewQQmlListReference2(QtObjectPtr object, char* property, QtObjectPtr engine){
	return new QQmlListReference(static_cast<QObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
}

int QQmlListReference_Append(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QQmlListReference*>(ptr)->append(static_cast<QObject*>(object));
}

QtObjectPtr QQmlListReference_At(QtObjectPtr ptr, int index){
	return static_cast<QQmlListReference*>(ptr)->at(index);
}

int QQmlListReference_CanAppend(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->canAppend();
}

int QQmlListReference_CanAt(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->canAt();
}

int QQmlListReference_CanClear(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->canClear();
}

int QQmlListReference_CanCount(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->canCount();
}

int QQmlListReference_Clear(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->clear();
}

int QQmlListReference_Count(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->count();
}

int QQmlListReference_IsManipulable(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->isManipulable();
}

int QQmlListReference_IsReadable(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->isReadable();
}

int QQmlListReference_IsValid(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->isValid();
}

QtObjectPtr QQmlListReference_ListElementType(QtObjectPtr ptr){
	return const_cast<QMetaObject*>(static_cast<QQmlListReference*>(ptr)->listElementType());
}

QtObjectPtr QQmlListReference_Object(QtObjectPtr ptr){
	return static_cast<QQmlListReference*>(ptr)->object();
}

