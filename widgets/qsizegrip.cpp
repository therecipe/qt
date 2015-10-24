#include "qsizegrip.h"
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QSizeGrip>
#include "_cgo_export.h"

class MyQSizeGrip: public QSizeGrip {
public:
};

QtObjectPtr QSizeGrip_NewQSizeGrip(QtObjectPtr parent){
	return new QSizeGrip(static_cast<QWidget*>(parent));
}

void QSizeGrip_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QSizeGrip*>(ptr)->setVisible(visible != 0);
}

void QSizeGrip_DestroyQSizeGrip(QtObjectPtr ptr){
	static_cast<QSizeGrip*>(ptr)->~QSizeGrip();
}

