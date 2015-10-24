#include "qmetaenum.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QMetaEnum>
#include "_cgo_export.h"

class MyQMetaEnum: public QMetaEnum {
public:
};

int QMetaEnum_IsFlag(QtObjectPtr ptr){
	return static_cast<QMetaEnum*>(ptr)->isFlag();
}

int QMetaEnum_IsValid(QtObjectPtr ptr){
	return static_cast<QMetaEnum*>(ptr)->isValid();
}

int QMetaEnum_KeyCount(QtObjectPtr ptr){
	return static_cast<QMetaEnum*>(ptr)->keyCount();
}

int QMetaEnum_KeyToValue(QtObjectPtr ptr, char* key, int ok){
	return static_cast<QMetaEnum*>(ptr)->keyToValue(const_cast<const char*>(key), NULL);
}

int QMetaEnum_KeysToValue(QtObjectPtr ptr, char* keys, int ok){
	return static_cast<QMetaEnum*>(ptr)->keysToValue(const_cast<const char*>(keys), NULL);
}

int QMetaEnum_Value(QtObjectPtr ptr, int index){
	return static_cast<QMetaEnum*>(ptr)->value(index);
}

