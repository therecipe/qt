#include "qstylehintreturnvariant.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleHintReturn>
#include <QString>
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

