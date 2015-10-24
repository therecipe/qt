#include "qtimeedit.h"
#include <QTime>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTimeEdit>
#include "_cgo_export.h"

class MyQTimeEdit: public QTimeEdit {
public:
};

QtObjectPtr QTimeEdit_NewQTimeEdit(QtObjectPtr parent){
	return new QTimeEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QTimeEdit_NewQTimeEdit2(QtObjectPtr time, QtObjectPtr parent){
	return new QTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

void QTimeEdit_DestroyQTimeEdit(QtObjectPtr ptr){
	static_cast<QTimeEdit*>(ptr)->~QTimeEdit();
}

