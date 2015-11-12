#include "qfontdialog.h"
#include <QFont>
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFontDialog>
#include "_cgo_export.h"

class MyQFontDialog: public QFontDialog {
public:
};

int QFontDialog_Options(void* ptr){
	return static_cast<QFontDialog*>(ptr)->options();
}

void QFontDialog_SetOptions(void* ptr, int options){
	static_cast<QFontDialog*>(ptr)->setOptions(static_cast<QFontDialog::FontDialogOption>(options));
}

void* QFontDialog_NewQFontDialog(void* parent){
	return new QFontDialog(static_cast<QWidget*>(parent));
}

void* QFontDialog_NewQFontDialog2(void* initial, void* parent){
	return new QFontDialog(*static_cast<QFont*>(initial), static_cast<QWidget*>(parent));
}

void QFontDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QFontDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QFontDialog_SetCurrentFont(void* ptr, void* font){
	static_cast<QFontDialog*>(ptr)->setCurrentFont(*static_cast<QFont*>(font));
}

void QFontDialog_SetOption(void* ptr, int option, int on){
	static_cast<QFontDialog*>(ptr)->setOption(static_cast<QFontDialog::FontDialogOption>(option), on != 0);
}

void QFontDialog_SetVisible(void* ptr, int visible){
	static_cast<QFontDialog*>(ptr)->setVisible(visible != 0);
}

int QFontDialog_TestOption(void* ptr, int option){
	return static_cast<QFontDialog*>(ptr)->testOption(static_cast<QFontDialog::FontDialogOption>(option));
}

