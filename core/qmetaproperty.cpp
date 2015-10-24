#include "qmetaproperty.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QMetaProperty>
#include "_cgo_export.h"

class MyQMetaProperty: public QMetaProperty {
public:
};

int QMetaProperty_HasNotifySignal(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->hasNotifySignal();
}

int QMetaProperty_IsConstant(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isConstant();
}

int QMetaProperty_IsDesignable(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->isDesignable(static_cast<QObject*>(object));
}

int QMetaProperty_IsEnumType(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isEnumType();
}

int QMetaProperty_IsFinal(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isFinal();
}

int QMetaProperty_IsFlagType(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isFlagType();
}

int QMetaProperty_IsReadable(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isReadable();
}

int QMetaProperty_IsResettable(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isResettable();
}

int QMetaProperty_IsScriptable(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->isScriptable(static_cast<QObject*>(object));
}

int QMetaProperty_IsStored(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->isStored(static_cast<QObject*>(object));
}

int QMetaProperty_IsUser(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->isUser(static_cast<QObject*>(object));
}

int QMetaProperty_IsValid(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isValid();
}

int QMetaProperty_IsWritable(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->isWritable();
}

int QMetaProperty_NotifySignalIndex(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->notifySignalIndex();
}

int QMetaProperty_PropertyIndex(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->propertyIndex();
}

char* QMetaProperty_Read(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->read(static_cast<QObject*>(object)).toString().toUtf8().data();
}

int QMetaProperty_Reset(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMetaProperty*>(ptr)->reset(static_cast<QObject*>(object));
}

int QMetaProperty_Revision(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->revision();
}

int QMetaProperty_UserType(QtObjectPtr ptr){
	return static_cast<QMetaProperty*>(ptr)->userType();
}

int QMetaProperty_Write(QtObjectPtr ptr, QtObjectPtr object, char* value){
	return static_cast<QMetaProperty*>(ptr)->write(static_cast<QObject*>(object), QVariant(value));
}

