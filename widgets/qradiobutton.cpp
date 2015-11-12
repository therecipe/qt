#include "qradiobutton.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QRadioButton>
#include "_cgo_export.h"

class MyQRadioButton: public QRadioButton {
public:
};

void* QRadioButton_NewQRadioButton(void* parent){
	return new QRadioButton(static_cast<QWidget*>(parent));
}

void* QRadioButton_NewQRadioButton2(char* text, void* parent){
	return new QRadioButton(QString(text), static_cast<QWidget*>(parent));
}

void QRadioButton_DestroyQRadioButton(void* ptr){
	static_cast<QRadioButton*>(ptr)->~QRadioButton();
}

