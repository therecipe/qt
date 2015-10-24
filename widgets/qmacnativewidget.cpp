#include "qmacnativewidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMacNativeWidget>
#include "_cgo_export.h"

class MyQMacNativeWidget: public QMacNativeWidget {
public:
};

void QMacNativeWidget_DestroyQMacNativeWidget(QtObjectPtr ptr){
	static_cast<QMacNativeWidget*>(ptr)->~QMacNativeWidget();
}

