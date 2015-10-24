#include "qsgmaterial.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSGMaterial>
#include "_cgo_export.h"

class MyQSGMaterial: public QSGMaterial {
public:
};

int QSGMaterial_Compare(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QSGMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

QtObjectPtr QSGMaterial_CreateShader(QtObjectPtr ptr){
	return static_cast<QSGMaterial*>(ptr)->createShader();
}

int QSGMaterial_Flags(QtObjectPtr ptr){
	return static_cast<QSGMaterial*>(ptr)->flags();
}

void QSGMaterial_SetFlag(QtObjectPtr ptr, int flags, int on){
	static_cast<QSGMaterial*>(ptr)->setFlag(static_cast<QSGMaterial::Flag>(flags), on != 0);
}

QtObjectPtr QSGMaterial_Type(QtObjectPtr ptr){
	return static_cast<QSGMaterial*>(ptr)->type();
}

