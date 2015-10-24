#include "qmessagelogger.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QMessageLogger>
#include "_cgo_export.h"

class MyQMessageLogger: public QMessageLogger {
public:
};

QtObjectPtr QMessageLogger_NewQMessageLogger(){
	return new QMessageLogger();
}

QtObjectPtr QMessageLogger_NewQMessageLogger2(char* file, int line, char* function){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function));
}

QtObjectPtr QMessageLogger_NewQMessageLogger3(char* file, int line, char* function, char* category){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function), const_cast<const char*>(category));
}

