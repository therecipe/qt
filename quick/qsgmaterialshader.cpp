#include "qsgmaterialshader.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGMaterial>
#include <QSGMaterialShader>
#include "_cgo_export.h"

class MyQSGMaterialShader: public QSGMaterialShader {
public:
};

void QSGMaterialShader_Activate(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->activate();
}

void QSGMaterialShader_Deactivate(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->deactivate();
}

void* QSGMaterialShader_Program(void* ptr){
	return static_cast<QSGMaterialShader*>(ptr)->program();
}

