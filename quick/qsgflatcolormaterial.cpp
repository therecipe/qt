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

void QSGFlatColorMaterial_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QSGFlatColorMaterial*>(ptr)->setColor(*static_cast<QColor*>(color));
}

