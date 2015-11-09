#include "qstyleoptionsizegrip.h"
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionSizeGrip>
#include "_cgo_export.h"

class MyQStyleOptionSizeGrip: public QStyleOptionSizeGrip {
public:
};

void* QStyleOptionSizeGrip_NewQStyleOptionSizeGrip(){
	return new QStyleOptionSizeGrip();
}

void* QStyleOptionSizeGrip_NewQStyleOptionSizeGrip2(void* other){
	return new QStyleOptionSizeGrip(*static_cast<QStyleOptionSizeGrip*>(other));
}

