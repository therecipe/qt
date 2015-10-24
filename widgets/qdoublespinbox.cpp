#include "qdoublespinbox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QDoubleSpinBox>
#include "_cgo_export.h"

class MyQDoubleSpinBox: public QDoubleSpinBox {
public:
};

char* QDoubleSpinBox_CleanText(QtObjectPtr ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QDoubleSpinBox_Decimals(QtObjectPtr ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->decimals();
}

char* QDoubleSpinBox_Prefix(QtObjectPtr ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QDoubleSpinBox_SetDecimals(QtObjectPtr ptr, int prec){
	static_cast<QDoubleSpinBox*>(ptr)->setDecimals(prec);
}

void QDoubleSpinBox_SetPrefix(QtObjectPtr ptr, char* prefix){
	static_cast<QDoubleSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QDoubleSpinBox_SetSuffix(QtObjectPtr ptr, char* suffix){
	static_cast<QDoubleSpinBox*>(ptr)->setSuffix(QString(suffix));
}

char* QDoubleSpinBox_Suffix(QtObjectPtr ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->suffix().toUtf8().data();
}

QtObjectPtr QDoubleSpinBox_NewQDoubleSpinBox(QtObjectPtr parent){
	return new QDoubleSpinBox(static_cast<QWidget*>(parent));
}

void QDoubleSpinBox_DestroyQDoubleSpinBox(QtObjectPtr ptr){
	static_cast<QDoubleSpinBox*>(ptr)->~QDoubleSpinBox();
}

