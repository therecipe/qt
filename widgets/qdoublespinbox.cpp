#include "qdoublespinbox.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDoubleSpinBox>
#include "_cgo_export.h"

class MyQDoubleSpinBox: public QDoubleSpinBox {
public:
};

char* QDoubleSpinBox_CleanText(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QDoubleSpinBox_Decimals(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->decimals();
}

char* QDoubleSpinBox_Prefix(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QDoubleSpinBox_SetDecimals(void* ptr, int prec){
	static_cast<QDoubleSpinBox*>(ptr)->setDecimals(prec);
}

void QDoubleSpinBox_SetPrefix(void* ptr, char* prefix){
	static_cast<QDoubleSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QDoubleSpinBox_SetSuffix(void* ptr, char* suffix){
	static_cast<QDoubleSpinBox*>(ptr)->setSuffix(QString(suffix));
}

char* QDoubleSpinBox_Suffix(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->suffix().toUtf8().data();
}

void* QDoubleSpinBox_NewQDoubleSpinBox(void* parent){
	return new QDoubleSpinBox(static_cast<QWidget*>(parent));
}

void QDoubleSpinBox_DestroyQDoubleSpinBox(void* ptr){
	static_cast<QDoubleSpinBox*>(ptr)->~QDoubleSpinBox();
}

