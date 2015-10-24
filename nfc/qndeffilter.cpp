#include "qndeffilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefFilter>
#include "_cgo_export.h"

class MyQNdefFilter: public QNdefFilter {
public:
};

QtObjectPtr QNdefFilter_NewQNdefFilter(){
	return new QNdefFilter();
}

QtObjectPtr QNdefFilter_NewQNdefFilter2(QtObjectPtr other){
	return new QNdefFilter(*static_cast<QNdefFilter*>(other));
}

void QNdefFilter_Clear(QtObjectPtr ptr){
	static_cast<QNdefFilter*>(ptr)->clear();
}

int QNdefFilter_OrderMatch(QtObjectPtr ptr){
	return static_cast<QNdefFilter*>(ptr)->orderMatch();
}

int QNdefFilter_RecordCount(QtObjectPtr ptr){
	return static_cast<QNdefFilter*>(ptr)->recordCount();
}

void QNdefFilter_SetOrderMatch(QtObjectPtr ptr, int on){
	static_cast<QNdefFilter*>(ptr)->setOrderMatch(on != 0);
}

void QNdefFilter_DestroyQNdefFilter(QtObjectPtr ptr){
	static_cast<QNdefFilter*>(ptr)->~QNdefFilter();
}

