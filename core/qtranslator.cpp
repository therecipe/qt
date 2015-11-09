#include "qtranslator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLocale>
#include <QString>
#include <QTranslator>
#include "_cgo_export.h"

class MyQTranslator: public QTranslator {
public:
};

void* QTranslator_NewQTranslator(void* parent){
	return new QTranslator(static_cast<QObject*>(parent));
}

int QTranslator_IsEmpty(void* ptr){
	return static_cast<QTranslator*>(ptr)->isEmpty();
}

int QTranslator_Load2(void* ptr, void* locale, char* filename, char* prefix, char* directory, char* suffix){
	return static_cast<QTranslator*>(ptr)->load(*static_cast<QLocale*>(locale), QString(filename), QString(prefix), QString(directory), QString(suffix));
}

int QTranslator_Load(void* ptr, char* filename, char* directory, char* search_delimiters, char* suffix){
	return static_cast<QTranslator*>(ptr)->load(QString(filename), QString(directory), QString(search_delimiters), QString(suffix));
}

char* QTranslator_Translate(void* ptr, char* context, char* sourceText, char* disambiguation, int n){
	return static_cast<QTranslator*>(ptr)->translate(const_cast<const char*>(context), const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QTranslator_DestroyQTranslator(void* ptr){
	static_cast<QTranslator*>(ptr)->~QTranslator();
}

