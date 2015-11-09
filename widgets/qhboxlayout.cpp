#include "qhboxlayout.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QHBoxLayout>
#include "_cgo_export.h"

class MyQHBoxLayout: public QHBoxLayout {
public:
};

void* QHBoxLayout_NewQHBoxLayout(){
	return new QHBoxLayout();
}

void* QHBoxLayout_NewQHBoxLayout2(void* parent){
	return new QHBoxLayout(static_cast<QWidget*>(parent));
}

void QHBoxLayout_DestroyQHBoxLayout(void* ptr){
	static_cast<QHBoxLayout*>(ptr)->~QHBoxLayout();
}

