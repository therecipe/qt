#include "qdateedit.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QDate>
#include <QDateEdit>
#include "_cgo_export.h"

class MyQDateEdit: public QDateEdit {
public:
};

QtObjectPtr QDateEdit_NewQDateEdit(QtObjectPtr parent){
	return new QDateEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QDateEdit_NewQDateEdit2(QtObjectPtr date, QtObjectPtr parent){
	return new QDateEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

void QDateEdit_DestroyQDateEdit(QtObjectPtr ptr){
	static_cast<QDateEdit*>(ptr)->~QDateEdit();
}

