#include "qradiobutton.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRadioButton>
#include "_cgo_export.h"

class MyQRadioButton: public QRadioButton {
public:
};

QtObjectPtr QRadioButton_NewQRadioButton(QtObjectPtr parent){
	return new QRadioButton(static_cast<QWidget*>(parent));
}

QtObjectPtr QRadioButton_NewQRadioButton2(char* text, QtObjectPtr parent){
	return new QRadioButton(QString(text), static_cast<QWidget*>(parent));
}

void QRadioButton_DestroyQRadioButton(QtObjectPtr ptr){
	static_cast<QRadioButton*>(ptr)->~QRadioButton();
}

