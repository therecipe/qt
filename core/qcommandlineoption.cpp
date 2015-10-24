#include "qcommandlineoption.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCommandLineOption>
#include "_cgo_export.h"

class MyQCommandLineOption: public QCommandLineOption {
public:
};

QtObjectPtr QCommandLineOption_NewQCommandLineOption5(QtObjectPtr other){
	return new QCommandLineOption(*static_cast<QCommandLineOption*>(other));
}

QtObjectPtr QCommandLineOption_NewQCommandLineOption(char* name){
	return new QCommandLineOption(QString(name));
}

QtObjectPtr QCommandLineOption_NewQCommandLineOption3(char* name, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(name), QString(description), QString(valueName), QString(defaultValue));
}

QtObjectPtr QCommandLineOption_NewQCommandLineOption2(char* names){
	return new QCommandLineOption(QString(names).split("|", QString::SkipEmptyParts));
}

QtObjectPtr QCommandLineOption_NewQCommandLineOption4(char* names, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(names).split("|", QString::SkipEmptyParts), QString(description), QString(valueName), QString(defaultValue));
}

char* QCommandLineOption_DefaultValues(QtObjectPtr ptr){
	return static_cast<QCommandLineOption*>(ptr)->defaultValues().join("|").toUtf8().data();
}

char* QCommandLineOption_Description(QtObjectPtr ptr){
	return static_cast<QCommandLineOption*>(ptr)->description().toUtf8().data();
}

char* QCommandLineOption_Names(QtObjectPtr ptr){
	return static_cast<QCommandLineOption*>(ptr)->names().join("|").toUtf8().data();
}

void QCommandLineOption_SetDefaultValue(QtObjectPtr ptr, char* defaultValue){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValue(QString(defaultValue));
}

void QCommandLineOption_SetDefaultValues(QtObjectPtr ptr, char* defaultValues){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValues(QString(defaultValues).split("|", QString::SkipEmptyParts));
}

void QCommandLineOption_SetDescription(QtObjectPtr ptr, char* description){
	static_cast<QCommandLineOption*>(ptr)->setDescription(QString(description));
}

void QCommandLineOption_SetValueName(QtObjectPtr ptr, char* valueName){
	static_cast<QCommandLineOption*>(ptr)->setValueName(QString(valueName));
}

void QCommandLineOption_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QCommandLineOption*>(ptr)->swap(*static_cast<QCommandLineOption*>(other));
}

char* QCommandLineOption_ValueName(QtObjectPtr ptr){
	return static_cast<QCommandLineOption*>(ptr)->valueName().toUtf8().data();
}

void QCommandLineOption_DestroyQCommandLineOption(QtObjectPtr ptr){
	static_cast<QCommandLineOption*>(ptr)->~QCommandLineOption();
}

