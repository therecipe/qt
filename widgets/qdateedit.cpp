#include "qdateedit.h"
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QDateEdit>
#include "_cgo_export.h"

class MyQDateEdit: public QDateEdit {
public:
};

void* QDateEdit_NewQDateEdit(void* parent){
	return new QDateEdit(static_cast<QWidget*>(parent));
}

void* QDateEdit_NewQDateEdit2(void* date, void* parent){
	return new QDateEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

void QDateEdit_DestroyQDateEdit(void* ptr){
	static_cast<QDateEdit*>(ptr)->~QDateEdit();
}

