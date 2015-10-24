#include "qmetatype.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QMetaType>
#include "_cgo_export.h"

class MyQMetaType: public QMetaType {
public:
};

QtObjectPtr QMetaType_NewQMetaType(int typeId){
	return new QMetaType(typeId);
}

int QMetaType_Flags(QtObjectPtr ptr){
	return static_cast<QMetaType*>(ptr)->flags();
}

int QMetaType_QMetaType_IsRegistered(int ty){
	return QMetaType::isRegistered(ty);
}

int QMetaType_IsRegistered2(QtObjectPtr ptr){
	return static_cast<QMetaType*>(ptr)->isRegistered();
}

int QMetaType_IsValid(QtObjectPtr ptr){
	return static_cast<QMetaType*>(ptr)->isValid();
}

QtObjectPtr QMetaType_MetaObject(QtObjectPtr ptr){
	return const_cast<QMetaObject*>(static_cast<QMetaType*>(ptr)->metaObject());
}

QtObjectPtr QMetaType_QMetaType_MetaObjectForType(int ty){
	return const_cast<QMetaObject*>(QMetaType::metaObjectForType(ty));
}

int QMetaType_QMetaType_SizeOf(int ty){
	return QMetaType::sizeOf(ty);
}

int QMetaType_SizeOf2(QtObjectPtr ptr){
	return static_cast<QMetaType*>(ptr)->sizeOf();
}

int QMetaType_QMetaType_Type2(QtObjectPtr typeName){
	return QMetaType::type(*static_cast<QByteArray*>(typeName));
}

int QMetaType_QMetaType_Type(char* typeName){
	return QMetaType::type(const_cast<const char*>(typeName));
}

int QMetaType_QMetaType_TypeFlags(int ty){
	return QMetaType::typeFlags(ty);
}

void QMetaType_DestroyQMetaType(QtObjectPtr ptr){
	static_cast<QMetaType*>(ptr)->~QMetaType();
}

