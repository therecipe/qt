#include "qvboxlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QVBoxLayout>
#include "_cgo_export.h"

class MyQVBoxLayout: public QVBoxLayout {
public:
};

QtObjectPtr QVBoxLayout_NewQVBoxLayout(){
	return new QVBoxLayout();
}

QtObjectPtr QVBoxLayout_NewQVBoxLayout2(QtObjectPtr parent){
	return new QVBoxLayout(static_cast<QWidget*>(parent));
}

void QVBoxLayout_DestroyQVBoxLayout(QtObjectPtr ptr){
	static_cast<QVBoxLayout*>(ptr)->~QVBoxLayout();
}

