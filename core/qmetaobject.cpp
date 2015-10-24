#include "qmetaobject.h"
#include <QGenericArgument>
#include <QMetaMethod>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGenericReturnArgument>
#include <QMetaObject>
#include "_cgo_export.h"

class MyQMetaObject: public QMetaObject {
public:
};

void QMetaObject_QMetaObject_ConnectSlotsByName(QtObjectPtr object){
	QMetaObject::connectSlotsByName(static_cast<QObject*>(object));
}

int QMetaObject_QMetaObject_CheckConnectArgs2(QtObjectPtr signal, QtObjectPtr method){
	return QMetaObject::checkConnectArgs(*static_cast<QMetaMethod*>(signal), *static_cast<QMetaMethod*>(method));
}

int QMetaObject_QMetaObject_CheckConnectArgs(char* signal, char* method){
	return QMetaObject::checkConnectArgs(const_cast<const char*>(signal), const_cast<const char*>(method));
}

int QMetaObject_ClassInfoCount(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoCount();
}

int QMetaObject_ClassInfoOffset(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoOffset();
}

int QMetaObject_ConstructorCount(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->constructorCount();
}

int QMetaObject_EnumeratorCount(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorCount();
}

int QMetaObject_EnumeratorOffset(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorOffset();
}

int QMetaObject_IndexOfClassInfo(QtObjectPtr ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfClassInfo(const_cast<const char*>(name));
}

int QMetaObject_IndexOfConstructor(QtObjectPtr ptr, char* constructor){
	return static_cast<QMetaObject*>(ptr)->indexOfConstructor(const_cast<const char*>(constructor));
}

int QMetaObject_IndexOfEnumerator(QtObjectPtr ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfEnumerator(const_cast<const char*>(name));
}

int QMetaObject_IndexOfMethod(QtObjectPtr ptr, char* method){
	return static_cast<QMetaObject*>(ptr)->indexOfMethod(const_cast<const char*>(method));
}

int QMetaObject_IndexOfProperty(QtObjectPtr ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfProperty(const_cast<const char*>(name));
}

int QMetaObject_IndexOfSignal(QtObjectPtr ptr, char* signal){
	return static_cast<QMetaObject*>(ptr)->indexOfSignal(const_cast<const char*>(signal));
}

int QMetaObject_IndexOfSlot(QtObjectPtr ptr, char* slot){
	return static_cast<QMetaObject*>(ptr)->indexOfSlot(const_cast<const char*>(slot));
}

int QMetaObject_QMetaObject_InvokeMethod4(QtObjectPtr obj, char* member, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod2(QtObjectPtr obj, char* member, QtObjectPtr ret, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod3(QtObjectPtr obj, char* member, int ty, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod(QtObjectPtr obj, char* member, int ty, QtObjectPtr ret, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_MethodCount(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->methodCount();
}

int QMetaObject_MethodOffset(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->methodOffset();
}

QtObjectPtr QMetaObject_NewInstance(QtObjectPtr ptr, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return static_cast<QMetaObject*>(ptr)->newInstance(*static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_PropertyCount(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->propertyCount();
}

int QMetaObject_PropertyOffset(QtObjectPtr ptr){
	return static_cast<QMetaObject*>(ptr)->propertyOffset();
}

QtObjectPtr QMetaObject_SuperClass(QtObjectPtr ptr){
	return const_cast<QMetaObject*>(static_cast<QMetaObject*>(ptr)->superClass());
}

