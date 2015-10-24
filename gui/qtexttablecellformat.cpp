#include "qtexttablecellformat.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextTable>
#include <QTextTableCell>
#include <QString>
#include <QTextTableCellFormat>
#include "_cgo_export.h"

class MyQTextTableCellFormat: public QTextTableCellFormat {
public:
};

QtObjectPtr QTextTableCellFormat_NewQTextTableCellFormat(){
	return new QTextTableCellFormat();
}

int QTextTableCellFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextTableCellFormat*>(ptr)->isValid();
}

