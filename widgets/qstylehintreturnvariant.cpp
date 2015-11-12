#include "qstylehintreturnvariant.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleHintReturn>
#include <QStyle>
#include <QStyleHintReturnVariant>
#include "_cgo_export.h"

class MyQStyleHintReturnVariant: public QStyleHintReturnVariant {
public:
};

void* QStyleHintReturnVariant_NewQStyleHintReturnVariant(){
	return new QStyleHintReturnVariant();
}

void QStyleHintReturnVariant_DestroyQStyleHintReturnVariant(void* ptr){
	static_cast<QStyleHintReturnVariant*>(ptr)->~QStyleHintReturnVariant();
}

