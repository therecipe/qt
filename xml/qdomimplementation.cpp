#include "qdomimplementation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomImplementation>
#include "_cgo_export.h"

class MyQDomImplementation: public QDomImplementation {
public:
};

void* QDomImplementation_NewQDomImplementation(){
	return new QDomImplementation();
}

void* QDomImplementation_NewQDomImplementation2(void* x){
	return new QDomImplementation(*static_cast<QDomImplementation*>(x));
}

int QDomImplementation_HasFeature(void* ptr, char* feature, char* version){
	return static_cast<QDomImplementation*>(ptr)->hasFeature(QString(feature), QString(version));
}

int QDomImplementation_QDomImplementation_InvalidDataPolicy(){
	return QDomImplementation::invalidDataPolicy();
}

int QDomImplementation_IsNull(void* ptr){
	return static_cast<QDomImplementation*>(ptr)->isNull();
}

void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(int policy){
	QDomImplementation::setInvalidDataPolicy(static_cast<QDomImplementation::InvalidDataPolicy>(policy));
}

void QDomImplementation_DestroyQDomImplementation(void* ptr){
	static_cast<QDomImplementation*>(ptr)->~QDomImplementation();
}

