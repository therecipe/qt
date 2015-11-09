#include "qstylehintreturn.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
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

void* QStyleHintReturn_NewQStyleHintReturn(int version, int ty){
	return new QStyleHintReturn(version, ty);
}

