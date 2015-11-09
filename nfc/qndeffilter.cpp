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

void* QNdefFilter_NewQNdefFilter(){
	return new QNdefFilter();
}

void* QNdefFilter_NewQNdefFilter2(void* other){
	return new QNdefFilter(*static_cast<QNdefFilter*>(other));
}

void QNdefFilter_Clear(void* ptr){
	static_cast<QNdefFilter*>(ptr)->clear();
}

int QNdefFilter_OrderMatch(void* ptr){
	return static_cast<QNdefFilter*>(ptr)->orderMatch();
}

int QNdefFilter_RecordCount(void* ptr){
	return static_cast<QNdefFilter*>(ptr)->recordCount();
}

void QNdefFilter_SetOrderMatch(void* ptr, int on){
	static_cast<QNdefFilter*>(ptr)->setOrderMatch(on != 0);
}

void QNdefFilter_DestroyQNdefFilter(void* ptr){
	static_cast<QNdefFilter*>(ptr)->~QNdefFilter();
}

