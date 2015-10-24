#include "qfontcombobox.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QFont>
#include <QFontDatabase>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QFontComboBox>
#include "_cgo_export.h"

class MyQFontComboBox: public QFontComboBox {
public:
};

int QFontComboBox_FontFilters(QtObjectPtr ptr){
	return static_cast<QFontComboBox*>(ptr)->fontFilters();
}

void QFontComboBox_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr font){
	QMetaObject::invokeMethod(static_cast<QFontComboBox*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(font)));
}

void QFontComboBox_SetFontFilters(QtObjectPtr ptr, int filters){
	static_cast<QFontComboBox*>(ptr)->setFontFilters(static_cast<QFontComboBox::FontFilter>(filters));
}

void QFontComboBox_SetWritingSystem(QtObjectPtr ptr, int script){
	static_cast<QFontComboBox*>(ptr)->setWritingSystem(static_cast<QFontDatabase::WritingSystem>(script));
}

int QFontComboBox_WritingSystem(QtObjectPtr ptr){
	return static_cast<QFontComboBox*>(ptr)->writingSystem();
}

QtObjectPtr QFontComboBox_NewQFontComboBox(QtObjectPtr parent){
	return new QFontComboBox(static_cast<QWidget*>(parent));
}

void QFontComboBox_DestroyQFontComboBox(QtObjectPtr ptr){
	static_cast<QFontComboBox*>(ptr)->~QFontComboBox();
}

