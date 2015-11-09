#include "qitemselection.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemSelectionRange>
#include <QItemSelectionModel>
#include <QItemSelection>
#include "_cgo_export.h"

class MyQItemSelection: public QItemSelection {
public:
};

void* QItemSelection_NewQItemSelection(){
	return new QItemSelection();
}

void* QItemSelection_NewQItemSelection2(void* topLeft, void* bottomRight){
	return new QItemSelection(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelection_Contains(void* ptr, void* index){
	return static_cast<QItemSelection*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

void QItemSelection_Merge(void* ptr, void* other, int command){
	static_cast<QItemSelection*>(ptr)->merge(*static_cast<QItemSelection*>(other), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QItemSelection_Select(void* ptr, void* topLeft, void* bottomRight){
	static_cast<QItemSelection*>(ptr)->select(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

void QItemSelection_QItemSelection_Split(void* ran, void* other, void* result){
	QItemSelection::split(*static_cast<QItemSelectionRange*>(ran), *static_cast<QItemSelectionRange*>(other), static_cast<QItemSelection*>(result));
}

