#include "qcommandlineoption.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QCommandLineOption>
#include "_cgo_export.h"

class MyQCommandLineOption: public QCommandLineOption {
public:
};

void* QCommandLineOption_NewQCommandLineOption5(void* other){
	return new QCommandLineOption(*static_cast<QCommandLineOption*>(other));
}

void* QCommandLineOption_NewQCommandLineOption(char* name){
	return new QCommandLineOption(QString(name));
}

void* QCommandLineOption_NewQCommandLineOption3(char* name, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(name), QString(description), QString(valueName), QString(defaultValue));
}

void* QCommandLineOption_NewQCommandLineOption2(char* names){
	return new QCommandLineOption(QString(names).split("|", QString::SkipEmptyParts));
}

void* QCommandLineOption_NewQCommandLineOption4(char* names, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(names).split("|", QString::SkipEmptyParts), QString(description), QString(valueName), QString(defaultValue));
}

char* QCommandLineOption_DefaultValues(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->defaultValues().join("|").toUtf8().data();
}

char* QCommandLineOption_Description(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->description().toUtf8().data();
}

char* QCommandLineOption_Names(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->names().join("|").toUtf8().data();
}

void QCommandLineOption_SetDefaultValue(void* ptr, char* defaultValue){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValue(QString(defaultValue));
}

void QCommandLineOption_SetDefaultValues(void* ptr, char* defaultValues){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValues(QString(defaultValues).split("|", QString::SkipEmptyParts));
}

void QCommandLineOption_SetDescription(void* ptr, char* description){
	static_cast<QCommandLineOption*>(ptr)->setDescription(QString(description));
}

void QCommandLineOption_SetValueName(void* ptr, char* valueName){
	static_cast<QCommandLineOption*>(ptr)->setValueName(QString(valueName));
}

void QCommandLineOption_Swap(void* ptr, void* other){
	static_cast<QCommandLineOption*>(ptr)->swap(*static_cast<QCommandLineOption*>(other));
}

char* QCommandLineOption_ValueName(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->valueName().toUtf8().data();
}

void QCommandLineOption_DestroyQCommandLineOption(void* ptr){
	static_cast<QCommandLineOption*>(ptr)->~QCommandLineOption();
}

