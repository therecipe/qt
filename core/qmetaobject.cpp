#include "qmetaobject.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QVariant>
#include <QGenericReturnArgument>
#include <QMetaMethod>
#include <QByteArray>
#include <QGenericArgument>
#include <QMetaObject>
#include "_cgo_export.h"

class MyQMetaObject: public QMetaObject {
public:
};

void QMetaObject_QMetaObject_ConnectSlotsByName(void* object){
	QMetaObject::connectSlotsByName(static_cast<QObject*>(object));
}

int QMetaObject_QMetaObject_CheckConnectArgs2(void* signal, void* method){
	return QMetaObject::checkConnectArgs(*static_cast<QMetaMethod*>(signal), *static_cast<QMetaMethod*>(method));
}

int QMetaObject_QMetaObject_CheckConnectArgs(char* signal, char* method){
	return QMetaObject::checkConnectArgs(const_cast<const char*>(signal), const_cast<const char*>(method));
}

int QMetaObject_ClassInfoCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoCount();
}

int QMetaObject_ClassInfoOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoOffset();
}

int QMetaObject_ConstructorCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->constructorCount();
}

int QMetaObject_EnumeratorCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorCount();
}

int QMetaObject_EnumeratorOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorOffset();
}

int QMetaObject_IndexOfClassInfo(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfClassInfo(const_cast<const char*>(name));
}

int QMetaObject_IndexOfConstructor(void* ptr, char* constructor){
	return static_cast<QMetaObject*>(ptr)->indexOfConstructor(const_cast<const char*>(constructor));
}

int QMetaObject_IndexOfEnumerator(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfEnumerator(const_cast<const char*>(name));
}

int QMetaObject_IndexOfMethod(void* ptr, char* method){
	return static_cast<QMetaObject*>(ptr)->indexOfMethod(const_cast<const char*>(method));
}

int QMetaObject_IndexOfProperty(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfProperty(const_cast<const char*>(name));
}

int QMetaObject_IndexOfSignal(void* ptr, char* signal){
	return static_cast<QMetaObject*>(ptr)->indexOfSignal(const_cast<const char*>(signal));
}

int QMetaObject_IndexOfSlot(void* ptr, char* slot){
	return static_cast<QMetaObject*>(ptr)->indexOfSlot(const_cast<const char*>(slot));
}

int QMetaObject_QMetaObject_InvokeMethod4(void* obj, char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod2(void* obj, char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod3(void* obj, char* member, int ty, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod(void* obj, char* member, int ty, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_MethodCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->methodCount();
}

int QMetaObject_MethodOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->methodOffset();
}

void* QMetaObject_NewInstance(void* ptr, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaObject*>(ptr)->newInstance(*static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

void* QMetaObject_QMetaObject_NormalizedSignature(char* method){
	return new QByteArray(QMetaObject::normalizedSignature(const_cast<const char*>(method)));
}

void* QMetaObject_QMetaObject_NormalizedType(char* ty){
	return new QByteArray(QMetaObject::normalizedType(const_cast<const char*>(ty)));
}

int QMetaObject_PropertyCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->propertyCount();
}

int QMetaObject_PropertyOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->propertyOffset();
}

void* QMetaObject_SuperClass(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QMetaObject*>(ptr)->superClass());
}

