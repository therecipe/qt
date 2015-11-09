#include "qsgmaterial.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSGMaterial>
#include "_cgo_export.h"

class MyQSGMaterial: public QSGMaterial {
public:
};

int QSGMaterial_Compare(void* ptr, void* other){
	return static_cast<QSGMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

void* QSGMaterial_CreateShader(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->createShader();
}

int QSGMaterial_Flags(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->flags();
}

void QSGMaterial_SetFlag(void* ptr, int flags, int on){
	static_cast<QSGMaterial*>(ptr)->setFlag(static_cast<QSGMaterial::Flag>(flags), on != 0);
}

void* QSGMaterial_Type(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->type();
}

