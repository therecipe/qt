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

void* QVBoxLayout_NewQVBoxLayout(){
	return new QVBoxLayout();
}

void* QVBoxLayout_NewQVBoxLayout2(void* parent){
	return new QVBoxLayout(static_cast<QWidget*>(parent));
}

void QVBoxLayout_DestroyQVBoxLayout(void* ptr){
	static_cast<QVBoxLayout*>(ptr)->~QVBoxLayout();
}

