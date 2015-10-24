#include "qstylehintreturnvariant.h"
#include <QStyleHintReturn>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleHintReturnVariant>
#include "_cgo_export.h"

class MyQStyleHintReturnVariant: public QStyleHintReturnVariant {
public:
};

QtObjectPtr QStyleHintReturnVariant_NewQStyleHintReturnVariant(){
	return new QStyleHintReturnVariant();
}

void QStyleHintReturnVariant_DestroyQStyleHintReturnVariant(QtObjectPtr ptr){
	static_cast<QStyleHintReturnVariant*>(ptr)->~QStyleHintReturnVariant();
}

