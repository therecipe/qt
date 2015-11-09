#include "qtimeedit.h"
#include <QModelIndex>
#include <QWidget>
#include <QTime>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTimeEdit>
#include "_cgo_export.h"

class MyQTimeEdit: public QTimeEdit {
public:
};

void* QTimeEdit_NewQTimeEdit(void* parent){
	return new QTimeEdit(static_cast<QWidget*>(parent));
}

void* QTimeEdit_NewQTimeEdit2(void* time, void* parent){
	return new QTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

void QTimeEdit_DestroyQTimeEdit(void* ptr){
	static_cast<QTimeEdit*>(ptr)->~QTimeEdit();
}

