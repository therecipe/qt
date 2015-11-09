#include "qmetamethod.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGenericArgument>
#include <QByteArray>
#include <QGenericReturnArgument>
#include <QString>
#include <QVariant>
#include <QMetaMethod>
#include "_cgo_export.h"

class MyQMetaMethod: public QMetaMethod {
public:
};

int QMetaMethod_Access(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->access();
}

int QMetaMethod_Invoke4(void* ptr, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke2(void* ptr, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke3(void* ptr, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke(void* ptr, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_IsValid(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->isValid();
}

int QMetaMethod_MethodIndex(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->methodIndex();
}

void* QMetaMethod_MethodSignature(void* ptr){
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->methodSignature());
}

int QMetaMethod_MethodType(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->methodType();
}

void* QMetaMethod_Name(void* ptr){
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->name());
}

int QMetaMethod_ParameterCount(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->parameterCount();
}

int QMetaMethod_ParameterType(void* ptr, int index){
	return static_cast<QMetaMethod*>(ptr)->parameterType(index);
}

int QMetaMethod_ReturnType(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->returnType();
}

int QMetaMethod_Revision(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->revision();
}

