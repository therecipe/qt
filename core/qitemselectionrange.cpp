#include "qitemselectionrange.h"
#include <QItemSelection>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemSelectionRange>
#include "_cgo_export.h"

class MyQItemSelectionRange: public QItemSelectionRange {
public:
};

int QItemSelectionRange_Intersects(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QItemSelectionRange*>(ptr)->intersects(*static_cast<QItemSelectionRange*>(other));
}

QtObjectPtr QItemSelectionRange_NewQItemSelectionRange(){
	return new QItemSelectionRange();
}

QtObjectPtr QItemSelectionRange_NewQItemSelectionRange2(QtObjectPtr other){
	return new QItemSelectionRange(*static_cast<QItemSelectionRange*>(other));
}

QtObjectPtr QItemSelectionRange_NewQItemSelectionRange4(QtObjectPtr index){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QItemSelectionRange_NewQItemSelectionRange3(QtObjectPtr topLeft, QtObjectPtr bottomRight){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelectionRange_Bottom(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->bottom();
}

int QItemSelectionRange_Contains(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QItemSelectionRange*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

int QItemSelectionRange_Contains2(QtObjectPtr ptr, int row, int column, QtObjectPtr parentIndex){
	return static_cast<QItemSelectionRange*>(ptr)->contains(row, column, *static_cast<QModelIndex*>(parentIndex));
}

int QItemSelectionRange_Height(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->height();
}

int QItemSelectionRange_IsEmpty(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isEmpty();
}

int QItemSelectionRange_IsValid(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isValid();
}

int QItemSelectionRange_Left(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->left();
}

QtObjectPtr QItemSelectionRange_Model(QtObjectPtr ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionRange*>(ptr)->model());
}

QtObjectPtr QItemSelectionRange_Parent(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->parent().internalPointer();
}

int QItemSelectionRange_Right(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->right();
}

int QItemSelectionRange_Top(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->top();
}

int QItemSelectionRange_Width(QtObjectPtr ptr){
	return static_cast<QItemSelectionRange*>(ptr)->width();
}

