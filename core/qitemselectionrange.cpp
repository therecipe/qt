#include "qitemselectionrange.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemSelection>
#include <QAbstractItemModel>
#include <QItemSelectionRange>
#include "_cgo_export.h"

class MyQItemSelectionRange: public QItemSelectionRange {
public:
};

int QItemSelectionRange_Intersects(void* ptr, void* other){
	return static_cast<QItemSelectionRange*>(ptr)->intersects(*static_cast<QItemSelectionRange*>(other));
}

void* QItemSelectionRange_NewQItemSelectionRange(){
	return new QItemSelectionRange();
}

void* QItemSelectionRange_NewQItemSelectionRange2(void* other){
	return new QItemSelectionRange(*static_cast<QItemSelectionRange*>(other));
}

void* QItemSelectionRange_NewQItemSelectionRange4(void* index){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(index));
}

void* QItemSelectionRange_NewQItemSelectionRange3(void* topLeft, void* bottomRight){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelectionRange_Bottom(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->bottom();
}

int QItemSelectionRange_Contains(void* ptr, void* index){
	return static_cast<QItemSelectionRange*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

int QItemSelectionRange_Contains2(void* ptr, int row, int column, void* parentIndex){
	return static_cast<QItemSelectionRange*>(ptr)->contains(row, column, *static_cast<QModelIndex*>(parentIndex));
}

int QItemSelectionRange_Height(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->height();
}

int QItemSelectionRange_IsEmpty(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isEmpty();
}

int QItemSelectionRange_IsValid(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isValid();
}

int QItemSelectionRange_Left(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->left();
}

void* QItemSelectionRange_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionRange*>(ptr)->model());
}

void* QItemSelectionRange_Parent(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->parent().internalPointer();
}

int QItemSelectionRange_Right(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->right();
}

int QItemSelectionRange_Top(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->top();
}

int QItemSelectionRange_Width(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->width();
}

