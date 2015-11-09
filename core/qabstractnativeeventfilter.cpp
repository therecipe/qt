#include "qabstractnativeeventfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractNativeEventFilter>
#include "_cgo_export.h"

class MyQAbstractNativeEventFilter: public QAbstractNativeEventFilter {
public:
};

void QAbstractNativeEventFilter_DestroyQAbstractNativeEventFilter(void* ptr){
	static_cast<QAbstractNativeEventFilter*>(ptr)->~QAbstractNativeEventFilter();
}

