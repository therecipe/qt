#include "qmetaproperty.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaProperty>
#include "_cgo_export.h"

class MyQMetaProperty: public QMetaProperty {
public:
};

int QMetaProperty_HasNotifySignal(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->hasNotifySignal();
}

int QMetaProperty_IsConstant(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isConstant();
}

int QMetaProperty_IsDesignable(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isDesignable(static_cast<QObject*>(object));
}

int QMetaProperty_IsEnumType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isEnumType();
}

int QMetaProperty_IsFinal(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isFinal();
}

int QMetaProperty_IsFlagType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isFlagType();
}

int QMetaProperty_IsReadable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isReadable();
}

int QMetaProperty_IsResettable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isResettable();
}

int QMetaProperty_IsScriptable(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isScriptable(static_cast<QObject*>(object));
}

int QMetaProperty_IsStored(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isStored(static_cast<QObject*>(object));
}

int QMetaProperty_IsUser(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isUser(static_cast<QObject*>(object));
}

int QMetaProperty_IsValid(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isValid();
}

int QMetaProperty_IsWritable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isWritable();
}

int QMetaProperty_NotifySignalIndex(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->notifySignalIndex();
}

int QMetaProperty_PropertyIndex(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->propertyIndex();
}

void* QMetaProperty_Read(void* ptr, void* object){
	return new QVariant(static_cast<QMetaProperty*>(ptr)->read(static_cast<QObject*>(object)));
}

int QMetaProperty_Reset(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->reset(static_cast<QObject*>(object));
}

int QMetaProperty_Revision(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->revision();
}

int QMetaProperty_UserType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->userType();
}

int QMetaProperty_Write(void* ptr, void* object, void* value){
	return static_cast<QMetaProperty*>(ptr)->write(static_cast<QObject*>(object), *static_cast<QVariant*>(value));
}

