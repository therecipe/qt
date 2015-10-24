#include "qcommandlineparser.h"
#include <QModelIndex>
#include <QCoreApplication>
#include <QCommandLineOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QCommandLineParser>
#include "_cgo_export.h"

class MyQCommandLineParser: public QCommandLineParser {
public:
};

QtObjectPtr QCommandLineParser_NewQCommandLineParser(){
	return new QCommandLineParser();
}

int QCommandLineParser_AddOption(QtObjectPtr ptr, QtObjectPtr option){
	return static_cast<QCommandLineParser*>(ptr)->addOption(*static_cast<QCommandLineOption*>(option));
}

void QCommandLineParser_AddPositionalArgument(QtObjectPtr ptr, char* name, char* description, char* syntax){
	static_cast<QCommandLineParser*>(ptr)->addPositionalArgument(QString(name), QString(description), QString(syntax));
}

char* QCommandLineParser_ApplicationDescription(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->applicationDescription().toUtf8().data();
}

void QCommandLineParser_ClearPositionalArguments(QtObjectPtr ptr){
	static_cast<QCommandLineParser*>(ptr)->clearPositionalArguments();
}

char* QCommandLineParser_ErrorText(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->errorText().toUtf8().data();
}

char* QCommandLineParser_HelpText(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->helpText().toUtf8().data();
}

int QCommandLineParser_IsSet2(QtObjectPtr ptr, QtObjectPtr option){
	return static_cast<QCommandLineParser*>(ptr)->isSet(*static_cast<QCommandLineOption*>(option));
}

int QCommandLineParser_IsSet(QtObjectPtr ptr, char* name){
	return static_cast<QCommandLineParser*>(ptr)->isSet(QString(name));
}

char* QCommandLineParser_OptionNames(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->optionNames().join("|").toUtf8().data();
}

int QCommandLineParser_Parse(QtObjectPtr ptr, char* arguments){
	return static_cast<QCommandLineParser*>(ptr)->parse(QString(arguments).split("|", QString::SkipEmptyParts));
}

char* QCommandLineParser_PositionalArguments(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->positionalArguments().join("|").toUtf8().data();
}

void QCommandLineParser_Process2(QtObjectPtr ptr, QtObjectPtr app){
	static_cast<QCommandLineParser*>(ptr)->process(*static_cast<QCoreApplication*>(app));
}

void QCommandLineParser_Process(QtObjectPtr ptr, char* arguments){
	static_cast<QCommandLineParser*>(ptr)->process(QString(arguments).split("|", QString::SkipEmptyParts));
}

void QCommandLineParser_SetApplicationDescription(QtObjectPtr ptr, char* description){
	static_cast<QCommandLineParser*>(ptr)->setApplicationDescription(QString(description));
}

void QCommandLineParser_SetSingleDashWordOptionMode(QtObjectPtr ptr, int singleDashWordOptionMode){
	static_cast<QCommandLineParser*>(ptr)->setSingleDashWordOptionMode(static_cast<QCommandLineParser::SingleDashWordOptionMode>(singleDashWordOptionMode));
}

void QCommandLineParser_ShowHelp(QtObjectPtr ptr, int exitCode){
	static_cast<QCommandLineParser*>(ptr)->showHelp(exitCode);
}

void QCommandLineParser_ShowVersion(QtObjectPtr ptr){
	static_cast<QCommandLineParser*>(ptr)->showVersion();
}

char* QCommandLineParser_UnknownOptionNames(QtObjectPtr ptr){
	return static_cast<QCommandLineParser*>(ptr)->unknownOptionNames().join("|").toUtf8().data();
}

char* QCommandLineParser_Value2(QtObjectPtr ptr, QtObjectPtr option){
	return static_cast<QCommandLineParser*>(ptr)->value(*static_cast<QCommandLineOption*>(option)).toUtf8().data();
}

char* QCommandLineParser_Value(QtObjectPtr ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->value(QString(optionName)).toUtf8().data();
}

char* QCommandLineParser_Values2(QtObjectPtr ptr, QtObjectPtr option){
	return static_cast<QCommandLineParser*>(ptr)->values(*static_cast<QCommandLineOption*>(option)).join("|").toUtf8().data();
}

char* QCommandLineParser_Values(QtObjectPtr ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->values(QString(optionName)).join("|").toUtf8().data();
}

void QCommandLineParser_DestroyQCommandLineParser(QtObjectPtr ptr){
	static_cast<QCommandLineParser*>(ptr)->~QCommandLineParser();
}

