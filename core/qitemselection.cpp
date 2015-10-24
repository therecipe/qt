#include "qitemselection.h"
#include <QItemSelectionModel>
#include <QItemSelectionRange>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemSelection>
#include "_cgo_export.h"

class MyQItemSelection: public QItemSelection {
public:
};

QtObjectPtr QItemSelection_NewQItemSelection(){
	return new QItemSelection();
}

QtObjectPtr QItemSelection_NewQItemSelection2(QtObjectPtr topLeft, QtObjectPtr bottomRight){
	return new QItemSelection(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelection_Contains(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QItemSelection*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

void QItemSelection_Merge(QtObjectPtr ptr, QtObjectPtr other, int command){
	static_cast<QItemSelection*>(ptr)->merge(*static_cast<QItemSelection*>(other), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QItemSelection_Select(QtObjectPtr ptr, QtObjectPtr topLeft, QtObjectPtr bottomRight){
	static_cast<QItemSelection*>(ptr)->select(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

void QItemSelection_QItemSelection_Split(QtObjectPtr ran, QtObjectPtr other, QtObjectPtr result){
	QItemSelection::split(*static_cast<QItemSelectionRange*>(ran), *static_cast<QItemSelectionRange*>(other), static_cast<QItemSelection*>(result));
}

