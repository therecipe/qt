#include "qqmlproperty.h"
#include <QModelIndex>
#include <QObject>
#include <QQmlContext>
#include <QQmlEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QQmlProperty>
#include "_cgo_export.h"

class MyQQmlProperty: public QQmlProperty {
public:
};

QtObjectPtr QQmlProperty_NewQQmlProperty(){
	return new QQmlProperty();
}

QtObjectPtr QQmlProperty_NewQQmlProperty2(QtObjectPtr obj){
	return new QQmlProperty(static_cast<QObject*>(obj));
}

QtObjectPtr QQmlProperty_NewQQmlProperty3(QtObjectPtr obj, QtObjectPtr ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlContext*>(ctxt));
}

QtObjectPtr QQmlProperty_NewQQmlProperty4(QtObjectPtr obj, QtObjectPtr engine){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlEngine*>(engine));
}

QtObjectPtr QQmlProperty_NewQQmlProperty5(QtObjectPtr obj, char* name){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name));
}

QtObjectPtr QQmlProperty_NewQQmlProperty6(QtObjectPtr obj, char* name, QtObjectPtr ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlContext*>(ctxt));
}

QtObjectPtr QQmlProperty_NewQQmlProperty7(QtObjectPtr obj, char* name, QtObjectPtr engine){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlEngine*>(engine));
}

QtObjectPtr QQmlProperty_NewQQmlProperty8(QtObjectPtr other){
	return new QQmlProperty(*static_cast<QQmlProperty*>(other));
}

int QQmlProperty_ConnectNotifySignal(QtObjectPtr ptr, QtObjectPtr dest, char* slot){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), const_cast<const char*>(slot));
}

int QQmlProperty_ConnectNotifySignal2(QtObjectPtr ptr, QtObjectPtr dest, int method){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), method);
}

int QQmlProperty_HasNotifySignal(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->hasNotifySignal();
}

int QQmlProperty_Index(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->index();
}

int QQmlProperty_IsDesignable(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isDesignable();
}

int QQmlProperty_IsProperty(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isProperty();
}

int QQmlProperty_IsResettable(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isResettable();
}

int QQmlProperty_IsSignalProperty(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isSignalProperty();
}

int QQmlProperty_IsValid(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isValid();
}

int QQmlProperty_IsWritable(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->isWritable();
}

char* QQmlProperty_Name(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->name().toUtf8().data();
}

int QQmlProperty_NeedsNotifySignal(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->needsNotifySignal();
}

QtObjectPtr QQmlProperty_Object(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->object();
}

int QQmlProperty_PropertyType(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyType();
}

int QQmlProperty_PropertyTypeCategory(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyTypeCategory();
}

char* QQmlProperty_QQmlProperty_Read2(QtObjectPtr object, char* name){
	return QQmlProperty::read(static_cast<QObject*>(object), QString(name)).toString().toUtf8().data();
}

char* QQmlProperty_QQmlProperty_Read3(QtObjectPtr object, char* name, QtObjectPtr ctxt){
	return QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlContext*>(ctxt)).toString().toUtf8().data();
}

char* QQmlProperty_QQmlProperty_Read4(QtObjectPtr object, char* name, QtObjectPtr engine){
	return QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlEngine*>(engine)).toString().toUtf8().data();
}

char* QQmlProperty_Read(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->read().toString().toUtf8().data();
}

int QQmlProperty_Reset(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->reset();
}

int QQmlProperty_Type(QtObjectPtr ptr){
	return static_cast<QQmlProperty*>(ptr)->type();
}

int QQmlProperty_QQmlProperty_Write2(QtObjectPtr object, char* name, char* value){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), QVariant(value));
}

int QQmlProperty_QQmlProperty_Write3(QtObjectPtr object, char* name, char* value, QtObjectPtr ctxt){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), QVariant(value), static_cast<QQmlContext*>(ctxt));
}

int QQmlProperty_QQmlProperty_Write4(QtObjectPtr object, char* name, char* value, QtObjectPtr engine){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), QVariant(value), static_cast<QQmlEngine*>(engine));
}

int QQmlProperty_Write(QtObjectPtr ptr, char* value){
	return static_cast<QQmlProperty*>(ptr)->write(QVariant(value));
}

