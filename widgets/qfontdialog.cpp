#include "qfontdialog.h"
#include <QModelIndex>
#include <QObject>
#include <QFont>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFontDialog>
#include "_cgo_export.h"

class MyQFontDialog: public QFontDialog {
public:
};

int QFontDialog_Options(QtObjectPtr ptr){
	return static_cast<QFontDialog*>(ptr)->options();
}

void QFontDialog_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QFontDialog*>(ptr)->setOptions(static_cast<QFontDialog::FontDialogOption>(options));
}

QtObjectPtr QFontDialog_NewQFontDialog(QtObjectPtr parent){
	return new QFontDialog(static_cast<QWidget*>(parent));
}

QtObjectPtr QFontDialog_NewQFontDialog2(QtObjectPtr initial, QtObjectPtr parent){
	return new QFontDialog(*static_cast<QFont*>(initial), static_cast<QWidget*>(parent));
}

void QFontDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QFontDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QFontDialog_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QFontDialog*>(ptr)->setCurrentFont(*static_cast<QFont*>(font));
}

void QFontDialog_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QFontDialog*>(ptr)->setOption(static_cast<QFontDialog::FontDialogOption>(option), on != 0);
}

void QFontDialog_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QFontDialog*>(ptr)->setVisible(visible != 0);
}

int QFontDialog_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QFontDialog*>(ptr)->testOption(static_cast<QFontDialog::FontDialogOption>(option));
}

