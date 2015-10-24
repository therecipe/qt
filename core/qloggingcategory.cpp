#include "qloggingcategory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLoggingCategory>
#include "_cgo_export.h"

class MyQLoggingCategory: public QLoggingCategory {
public:
};

QtObjectPtr QLoggingCategory_NewQLoggingCategory(char* category){
	return new QLoggingCategory(const_cast<const char*>(category));
}

QtObjectPtr QLoggingCategory_QLoggingCategory_DefaultCategory(){
	return QLoggingCategory::defaultCategory();
}

int QLoggingCategory_IsCriticalEnabled(QtObjectPtr ptr){
	return static_cast<QLoggingCategory*>(ptr)->isCriticalEnabled();
}

int QLoggingCategory_IsDebugEnabled(QtObjectPtr ptr){
	return static_cast<QLoggingCategory*>(ptr)->isDebugEnabled();
}

int QLoggingCategory_IsInfoEnabled(QtObjectPtr ptr){
	return static_cast<QLoggingCategory*>(ptr)->isInfoEnabled();
}

int QLoggingCategory_IsWarningEnabled(QtObjectPtr ptr){
	return static_cast<QLoggingCategory*>(ptr)->isWarningEnabled();
}

void QLoggingCategory_QLoggingCategory_SetFilterRules(char* rules){
	QLoggingCategory::setFilterRules(QString(rules));
}

void QLoggingCategory_DestroyQLoggingCategory(QtObjectPtr ptr){
	static_cast<QLoggingCategory*>(ptr)->~QLoggingCategory();
}

