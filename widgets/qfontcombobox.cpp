#include "qfontcombobox.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QMetaObject>
#include <QWidget>
#include <QFontDatabase>
#include <QString>
#include <QFontComboBox>
#include "_cgo_export.h"

class MyQFontComboBox: public QFontComboBox {
public:
};

int QFontComboBox_FontFilters(void* ptr){
	return static_cast<QFontComboBox*>(ptr)->fontFilters();
}

void QFontComboBox_SetCurrentFont(void* ptr, void* font){
	QMetaObject::invokeMethod(static_cast<QFontComboBox*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(font)));
}

void QFontComboBox_SetFontFilters(void* ptr, int filters){
	static_cast<QFontComboBox*>(ptr)->setFontFilters(static_cast<QFontComboBox::FontFilter>(filters));
}

void QFontComboBox_SetWritingSystem(void* ptr, int script){
	static_cast<QFontComboBox*>(ptr)->setWritingSystem(static_cast<QFontDatabase::WritingSystem>(script));
}

int QFontComboBox_WritingSystem(void* ptr){
	return static_cast<QFontComboBox*>(ptr)->writingSystem();
}

void* QFontComboBox_NewQFontComboBox(void* parent){
	return new QFontComboBox(static_cast<QWidget*>(parent));
}

void QFontComboBox_DestroyQFontComboBox(void* ptr){
	static_cast<QFontComboBox*>(ptr)->~QFontComboBox();
}

