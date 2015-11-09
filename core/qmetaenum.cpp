#include "qmetaenum.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QMetaEnum>
#include "_cgo_export.h"

class MyQMetaEnum: public QMetaEnum {
public:
};

int QMetaEnum_IsFlag(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->isFlag();
}

int QMetaEnum_IsValid(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->isValid();
}

int QMetaEnum_KeyCount(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->keyCount();
}

int QMetaEnum_KeyToValue(void* ptr, char* key, int ok){
	return static_cast<QMetaEnum*>(ptr)->keyToValue(const_cast<const char*>(key), NULL);
}

int QMetaEnum_KeysToValue(void* ptr, char* keys, int ok){
	return static_cast<QMetaEnum*>(ptr)->keysToValue(const_cast<const char*>(keys), NULL);
}

int QMetaEnum_Value(void* ptr, int index){
	return static_cast<QMetaEnum*>(ptr)->value(index);
}

void* QMetaEnum_ValueToKeys(void* ptr, int value){
	return new QByteArray(static_cast<QMetaEnum*>(ptr)->valueToKeys(value));
}

