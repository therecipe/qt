#include "qcommandlineparser.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCoreApplication>
#include <QCommandLineOption>
#include <QCommandLineParser>
#include "_cgo_export.h"

class MyQCommandLineParser: public QCommandLineParser {
public:
};

void* QCommandLineParser_NewQCommandLineParser(){
	return new QCommandLineParser();
}

void* QCommandLineParser_AddHelpOption(void* ptr){
	return new QCommandLineOption(static_cast<QCommandLineParser*>(ptr)->addHelpOption());
}

int QCommandLineParser_AddOption(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->addOption(*static_cast<QCommandLineOption*>(option));
}

void QCommandLineParser_AddPositionalArgument(void* ptr, char* name, char* description, char* syntax){
	static_cast<QCommandLineParser*>(ptr)->addPositionalArgument(QString(name), QString(description), QString(syntax));
}

void* QCommandLineParser_AddVersionOption(void* ptr){
	return new QCommandLineOption(static_cast<QCommandLineParser*>(ptr)->addVersionOption());
}

char* QCommandLineParser_ApplicationDescription(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->applicationDescription().toUtf8().data();
}

void QCommandLineParser_ClearPositionalArguments(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->clearPositionalArguments();
}

char* QCommandLineParser_ErrorText(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->errorText().toUtf8().data();
}

char* QCommandLineParser_HelpText(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->helpText().toUtf8().data();
}

int QCommandLineParser_IsSet2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->isSet(*static_cast<QCommandLineOption*>(option));
}

int QCommandLineParser_IsSet(void* ptr, char* name){
	return static_cast<QCommandLineParser*>(ptr)->isSet(QString(name));
}

char* QCommandLineParser_OptionNames(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->optionNames().join("|").toUtf8().data();
}

int QCommandLineParser_Parse(void* ptr, char* arguments){
	return static_cast<QCommandLineParser*>(ptr)->parse(QString(arguments).split("|", QString::SkipEmptyParts));
}

char* QCommandLineParser_PositionalArguments(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->positionalArguments().join("|").toUtf8().data();
}

void QCommandLineParser_Process2(void* ptr, void* app){
	static_cast<QCommandLineParser*>(ptr)->process(*static_cast<QCoreApplication*>(app));
}

void QCommandLineParser_Process(void* ptr, char* arguments){
	static_cast<QCommandLineParser*>(ptr)->process(QString(arguments).split("|", QString::SkipEmptyParts));
}

void QCommandLineParser_SetApplicationDescription(void* ptr, char* description){
	static_cast<QCommandLineParser*>(ptr)->setApplicationDescription(QString(description));
}

void QCommandLineParser_SetSingleDashWordOptionMode(void* ptr, int singleDashWordOptionMode){
	static_cast<QCommandLineParser*>(ptr)->setSingleDashWordOptionMode(static_cast<QCommandLineParser::SingleDashWordOptionMode>(singleDashWordOptionMode));
}

void QCommandLineParser_ShowHelp(void* ptr, int exitCode){
	static_cast<QCommandLineParser*>(ptr)->showHelp(exitCode);
}

void QCommandLineParser_ShowVersion(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->showVersion();
}

char* QCommandLineParser_UnknownOptionNames(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->unknownOptionNames().join("|").toUtf8().data();
}

char* QCommandLineParser_Value2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->value(*static_cast<QCommandLineOption*>(option)).toUtf8().data();
}

char* QCommandLineParser_Value(void* ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->value(QString(optionName)).toUtf8().data();
}

char* QCommandLineParser_Values2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->values(*static_cast<QCommandLineOption*>(option)).join("|").toUtf8().data();
}

char* QCommandLineParser_Values(void* ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->values(QString(optionName)).join("|").toUtf8().data();
}

void QCommandLineParser_DestroyQCommandLineParser(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->~QCommandLineParser();
}

