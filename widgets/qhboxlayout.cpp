#include "qhboxlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QHBoxLayout>
#include "_cgo_export.h"

class MyQHBoxLayout: public QHBoxLayout {
public:
};

QtObjectPtr QHBoxLayout_NewQHBoxLayout(){
	return new QHBoxLayout();
}

QtObjectPtr QHBoxLayout_NewQHBoxLayout2(QtObjectPtr parent){
	return new QHBoxLayout(static_cast<QWidget*>(parent));
}

void QHBoxLayout_DestroyQHBoxLayout(QtObjectPtr ptr){
	static_cast<QHBoxLayout*>(ptr)->~QHBoxLayout();
}

