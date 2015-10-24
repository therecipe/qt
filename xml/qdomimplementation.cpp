#include "qdomimplementation.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDomImplementation>
#include "_cgo_export.h"

class MyQDomImplementation: public QDomImplementation {
public:
};

QtObjectPtr QDomImplementation_NewQDomImplementation(){
	return new QDomImplementation();
}

QtObjectPtr QDomImplementation_NewQDomImplementation2(QtObjectPtr x){
	return new QDomImplementation(*static_cast<QDomImplementation*>(x));
}

int QDomImplementation_HasFeature(QtObjectPtr ptr, char* feature, char* version){
	return static_cast<QDomImplementation*>(ptr)->hasFeature(QString(feature), QString(version));
}

int QDomImplementation_QDomImplementation_InvalidDataPolicy(){
	return QDomImplementation::invalidDataPolicy();
}

int QDomImplementation_IsNull(QtObjectPtr ptr){
	return static_cast<QDomImplementation*>(ptr)->isNull();
}

void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(int policy){
	QDomImplementation::setInvalidDataPolicy(static_cast<QDomImplementation::InvalidDataPolicy>(policy));
}

void QDomImplementation_DestroyQDomImplementation(QtObjectPtr ptr){
	static_cast<QDomImplementation*>(ptr)->~QDomImplementation();
}

