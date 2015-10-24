#include "qcolordialog.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QColor>
#include <QObject>
#include <QString>
#include <QColorDialog>
#include "_cgo_export.h"

class MyQColorDialog: public QColorDialog {
public:
};

int QColorDialog_Options(QtObjectPtr ptr){
	return static_cast<QColorDialog*>(ptr)->options();
}

void QColorDialog_SetCurrentColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QColorDialog*>(ptr)->setCurrentColor(*static_cast<QColor*>(color));
}

void QColorDialog_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QColorDialog*>(ptr)->setOptions(static_cast<QColorDialog::ColorDialogOption>(options));
}

QtObjectPtr QColorDialog_NewQColorDialog(QtObjectPtr parent){
	return new QColorDialog(static_cast<QWidget*>(parent));
}

QtObjectPtr QColorDialog_NewQColorDialog2(QtObjectPtr initial, QtObjectPtr parent){
	return new QColorDialog(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent));
}

int QColorDialog_QColorDialog_CustomCount(){
	return QColorDialog::customCount();
}

void QColorDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QColorDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QColorDialog_QColorDialog_SetCustomColor(int index, QtObjectPtr color){
	QColorDialog::setCustomColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QColorDialog*>(ptr)->setOption(static_cast<QColorDialog::ColorDialogOption>(option), on != 0);
}

void QColorDialog_QColorDialog_SetStandardColor(int index, QtObjectPtr color){
	QColorDialog::setStandardColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QColorDialog*>(ptr)->setVisible(visible != 0);
}

int QColorDialog_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QColorDialog*>(ptr)->testOption(static_cast<QColorDialog::ColorDialogOption>(option));
}

void QColorDialog_DestroyQColorDialog(QtObjectPtr ptr){
	static_cast<QColorDialog*>(ptr)->~QColorDialog();
}

