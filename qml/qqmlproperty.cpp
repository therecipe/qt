#include "qqmlproperty.h"
#include <QModelIndex>
#include <QQmlEngine>
#include <QObject>
#include <QQmlContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QQmlProperty>
#include "_cgo_export.h"

class MyQQmlProperty: public QQmlProperty {
public:
};

void* QQmlProperty_NewQQmlProperty(){
	return new QQmlProperty();
}

void* QQmlProperty_NewQQmlProperty2(void* obj){
	return new QQmlProperty(static_cast<QObject*>(obj));
}

void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty5(void* obj, char* name){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name));
}

void* QQmlProperty_NewQQmlProperty6(void* obj, char* name, void* ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty7(void* obj, char* name, void* engine){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty8(void* other){
	return new QQmlProperty(*static_cast<QQmlProperty*>(other));
}

int QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), const_cast<const char*>(slot));
}

int QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), method);
}

int QQmlProperty_HasNotifySignal(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->hasNotifySignal();
}

int QQmlProperty_Index(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->index();
}

int QQmlProperty_IsDesignable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isDesignable();
}

int QQmlProperty_IsProperty(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isProperty();
}

int QQmlProperty_IsResettable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isResettable();
}

int QQmlProperty_IsSignalProperty(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isSignalProperty();
}

int QQmlProperty_IsValid(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isValid();
}

int QQmlProperty_IsWritable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isWritable();
}

char* QQmlProperty_Name(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->name().toUtf8().data();
}

int QQmlProperty_NeedsNotifySignal(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->needsNotifySignal();
}

void* QQmlProperty_Object(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->object();
}

int QQmlProperty_PropertyType(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyType();
}

int QQmlProperty_PropertyTypeCategory(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyTypeCategory();
}

void* QQmlProperty_QQmlProperty_Read2(void* object, char* name){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name)));
}

void* QQmlProperty_QQmlProperty_Read3(void* object, char* name, void* ctxt){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlContext*>(ctxt)));
}

void* QQmlProperty_QQmlProperty_Read4(void* object, char* name, void* engine){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlEngine*>(engine)));
}

void* QQmlProperty_Read(void* ptr){
	return new QVariant(static_cast<QQmlProperty*>(ptr)->read());
}

int QQmlProperty_Reset(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->reset();
}

int QQmlProperty_Type(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->type();
}

int QQmlProperty_QQmlProperty_Write2(void* object, char* name, void* value){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value));
}

int QQmlProperty_QQmlProperty_Write3(void* object, char* name, void* value, void* ctxt){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlContext*>(ctxt));
}

int QQmlProperty_QQmlProperty_Write4(void* object, char* name, void* value, void* engine){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlEngine*>(engine));
}

int QQmlProperty_Write(void* ptr, void* value){
	return static_cast<QQmlProperty*>(ptr)->write(*static_cast<QVariant*>(value));
}

