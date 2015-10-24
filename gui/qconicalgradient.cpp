#include "qconicalgradient.h"
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QConicalGradient>
#include "_cgo_export.h"

class MyQConicalGradient: public QConicalGradient {
public:
};

QtObjectPtr QConicalGradient_NewQConicalGradient(){
	return new QConicalGradient();
}

void QConicalGradient_SetCenter(QtObjectPtr ptr, QtObjectPtr center){
	static_cast<QConicalGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

