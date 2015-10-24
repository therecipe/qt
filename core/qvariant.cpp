#include "qvariant.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QVariant>
#include "_cgo_export.h"

class MyQVariant: public QVariant {
public:
};

char* QVariant_ToStringList(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->toStringList().join("|").toUtf8().data();
}

char* QVariant_ToUrl(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->toUrl().toString().toUtf8().data();
}

void QVariant_DestroyQVariant(QtObjectPtr ptr){
	static_cast<QVariant*>(ptr)->~QVariant();
}

void QVariant_Clear(QtObjectPtr ptr){
	static_cast<QVariant*>(ptr)->clear();
}

int QVariant_Convert(QtObjectPtr ptr, int targetTypeId){
	return static_cast<QVariant*>(ptr)->convert(targetTypeId);
}

int QVariant_IsNull(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->isNull();
}

int QVariant_IsValid(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->isValid();
}

int QVariant_ToBool(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->toBool();
}

int QVariant_ToInt(QtObjectPtr ptr, int ok){
	return static_cast<QVariant*>(ptr)->toInt(NULL);
}

QtObjectPtr QVariant_ToModelIndex(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->toModelIndex().internalPointer();
}

char* QVariant_ToString(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->toString().toUtf8().data();
}

int QVariant_UserType(QtObjectPtr ptr){
	return static_cast<QVariant*>(ptr)->userType();
}

