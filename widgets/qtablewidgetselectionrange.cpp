#include "qtablewidgetselectionrange.h"
#include <QUrl>
#include <QModelIndex>
#include <QTableWidget>
#include <QString>
#include <QVariant>
#include <QTableWidgetSelectionRange>
#include "_cgo_export.h"

class MyQTableWidgetSelectionRange: public QTableWidgetSelectionRange {
public:
};

QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange(){
	return new QTableWidgetSelectionRange();
}

QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(QtObjectPtr other){
	return new QTableWidgetSelectionRange(*static_cast<QTableWidgetSelectionRange*>(other));
}

QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(int top, int left, int bottom, int right){
	return new QTableWidgetSelectionRange(top, left, bottom, right);
}

int QTableWidgetSelectionRange_BottomRow(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->bottomRow();
}

int QTableWidgetSelectionRange_ColumnCount(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->columnCount();
}

int QTableWidgetSelectionRange_LeftColumn(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->leftColumn();
}

int QTableWidgetSelectionRange_RightColumn(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rightColumn();
}

int QTableWidgetSelectionRange_RowCount(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rowCount();
}

int QTableWidgetSelectionRange_TopRow(QtObjectPtr ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->topRow();
}

void QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(QtObjectPtr ptr){
	static_cast<QTableWidgetSelectionRange*>(ptr)->~QTableWidgetSelectionRange();
}

