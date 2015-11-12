#include "qtablewidgetselectionrange.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTableWidget>
#include <QString>
#include <QTableWidgetSelectionRange>
#include "_cgo_export.h"

class MyQTableWidgetSelectionRange: public QTableWidgetSelectionRange {
public:
};

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange(){
	return new QTableWidgetSelectionRange();
}

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(void* other){
	return new QTableWidgetSelectionRange(*static_cast<QTableWidgetSelectionRange*>(other));
}

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(int top, int left, int bottom, int right){
	return new QTableWidgetSelectionRange(top, left, bottom, right);
}

int QTableWidgetSelectionRange_BottomRow(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->bottomRow();
}

int QTableWidgetSelectionRange_ColumnCount(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->columnCount();
}

int QTableWidgetSelectionRange_LeftColumn(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->leftColumn();
}

int QTableWidgetSelectionRange_RightColumn(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rightColumn();
}

int QTableWidgetSelectionRange_RowCount(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rowCount();
}

int QTableWidgetSelectionRange_TopRow(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->topRow();
}

void QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(void* ptr){
	static_cast<QTableWidgetSelectionRange*>(ptr)->~QTableWidgetSelectionRange();
}

