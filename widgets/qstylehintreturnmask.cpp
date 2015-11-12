#include "qstylehintreturnmask.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleHintReturn>
#include <QStyle>
#include <QStyleHintReturnMask>
#include "_cgo_export.h"

class MyQStyleHintReturnMask: public QStyleHintReturnMask {
public:
};

void* QStyleHintReturnMask_NewQStyleHintReturnMask(){
	return new QStyleHintReturnMask();
}

void QStyleHintReturnMask_DestroyQStyleHintReturnMask(void* ptr){
	static_cast<QStyleHintReturnMask*>(ptr)->~QStyleHintReturnMask();
}

