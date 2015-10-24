#include "qmetamethod.h"
#include <QGenericArgument>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGenericReturnArgument>
#include <QMetaMethod>
#include "_cgo_export.h"

class MyQMetaMethod: public QMetaMethod {
public:
};

int QMetaMethod_Access(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->access();
}

int QMetaMethod_Invoke4(QtObjectPtr ptr, QtObjectPtr object, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke2(QtObjectPtr ptr, QtObjectPtr object, QtObjectPtr returnValue, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke3(QtObjectPtr ptr, QtObjectPtr object, int connectionType, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke(QtObjectPtr ptr, QtObjectPtr object, int connectionType, QtObjectPtr returnValue, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_IsValid(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->isValid();
}

int QMetaMethod_MethodIndex(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->methodIndex();
}

int QMetaMethod_MethodType(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->methodType();
}

int QMetaMethod_ParameterCount(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->parameterCount();
}

int QMetaMethod_ParameterType(QtObjectPtr ptr, int index){
	return static_cast<QMetaMethod*>(ptr)->parameterType(index);
}

int QMetaMethod_ReturnType(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->returnType();
}

int QMetaMethod_Revision(QtObjectPtr ptr){
	return static_cast<QMetaMethod*>(ptr)->revision();
}

