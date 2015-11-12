#include "qloggingcategory.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QLoggingCategory>
#include "_cgo_export.h"

class MyQLoggingCategory: public QLoggingCategory {
public:
};

void* QLoggingCategory_NewQLoggingCategory(char* category){
	return new QLoggingCategory(const_cast<const char*>(category));
}

void* QLoggingCategory_QLoggingCategory_DefaultCategory(){
	return QLoggingCategory::defaultCategory();
}

int QLoggingCategory_IsCriticalEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isCriticalEnabled();
}

int QLoggingCategory_IsDebugEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isDebugEnabled();
}

int QLoggingCategory_IsInfoEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isInfoEnabled();
}

int QLoggingCategory_IsWarningEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isWarningEnabled();
}

void QLoggingCategory_QLoggingCategory_SetFilterRules(char* rules){
	QLoggingCategory::setFilterRules(QString(rules));
}

void QLoggingCategory_DestroyQLoggingCategory(void* ptr){
	static_cast<QLoggingCategory*>(ptr)->~QLoggingCategory();
}

