#include "qstylehintreturnmask.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleHintReturn>
#include <QStyleHintReturnMask>
#include "_cgo_export.h"

class MyQStyleHintReturnMask: public QStyleHintReturnMask {
public:
};

QtObjectPtr QStyleHintReturnMask_NewQStyleHintReturnMask(){
	return new QStyleHintReturnMask();
}

void QStyleHintReturnMask_DestroyQStyleHintReturnMask(QtObjectPtr ptr){
	static_cast<QStyleHintReturnMask*>(ptr)->~QStyleHintReturnMask();
}

