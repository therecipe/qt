#include "qmessagelogger.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMessageLogger>
#include "_cgo_export.h"

class MyQMessageLogger: public QMessageLogger {
public:
};

void* QMessageLogger_NewQMessageLogger(){
	return new QMessageLogger();
}

void* QMessageLogger_NewQMessageLogger2(char* file, int line, char* function){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function));
}

void* QMessageLogger_NewQMessageLogger3(char* file, int line, char* function, char* category){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function), const_cast<const char*>(category));
}

