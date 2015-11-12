#include "qsizegrip.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QSizeGrip>
#include "_cgo_export.h"

class MyQSizeGrip: public QSizeGrip {
public:
};

void* QSizeGrip_NewQSizeGrip(void* parent){
	return new QSizeGrip(static_cast<QWidget*>(parent));
}

void QSizeGrip_SetVisible(void* ptr, int visible){
	static_cast<QSizeGrip*>(ptr)->setVisible(visible != 0);
}

void QSizeGrip_DestroyQSizeGrip(void* ptr){
	static_cast<QSizeGrip*>(ptr)->~QSizeGrip();
}

