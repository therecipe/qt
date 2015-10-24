#include "qstyleoptionsizegrip.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionSizeGrip>
#include "_cgo_export.h"

class MyQStyleOptionSizeGrip: public QStyleOptionSizeGrip {
public:
};

QtObjectPtr QStyleOptionSizeGrip_NewQStyleOptionSizeGrip(){
	return new QStyleOptionSizeGrip();
}

QtObjectPtr QStyleOptionSizeGrip_NewQStyleOptionSizeGrip2(QtObjectPtr other){
	return new QStyleOptionSizeGrip(*static_cast<QStyleOptionSizeGrip*>(other));
}

