#include "qstylehintreturn.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleHintReturn>
#include "_cgo_export.h"

class MyQStyleHintReturn: public QStyleHintReturn {
public:
};

int QStyleHintReturn_SH_Mask_Type(){
	return QStyleHintReturn::SH_Mask;
}

int QStyleHintReturn_SH_Variant_Type(){
	return QStyleHintReturn::SH_Variant;
}

QtObjectPtr QStyleHintReturn_NewQStyleHintReturn(int version, int ty){
	return new QStyleHintReturn(version, ty);
}

