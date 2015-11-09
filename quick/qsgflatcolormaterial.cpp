#include "qsgflatcolormaterial.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QString>
#include <QSGFlatColorMaterial>
#include "_cgo_export.h"

class MyQSGFlatColorMaterial: public QSGFlatColorMaterial {
public:
};

void* QSGFlatColorMaterial_Color(void* ptr){
	return new QColor(static_cast<QSGFlatColorMaterial*>(ptr)->color());
}

void QSGFlatColorMaterial_SetColor(void* ptr, void* color){
	static_cast<QSGFlatColorMaterial*>(ptr)->setColor(*static_cast<QColor*>(color));
}

