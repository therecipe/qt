#include "qmacnativewidget.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMacNativeWidget>
#include "_cgo_export.h"

class MyQMacNativeWidget: public QMacNativeWidget {
public:
};

void QMacNativeWidget_DestroyQMacNativeWidget(void* ptr){
	static_cast<QMacNativeWidget*>(ptr)->~QMacNativeWidget();
}

