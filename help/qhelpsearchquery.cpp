#include "qhelpsearchquery.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpSearchQuery>
#include "_cgo_export.h"

class MyQHelpSearchQuery: public QHelpSearchQuery {
public:
};

void* QHelpSearchQuery_NewQHelpSearchQuery(){
	return new QHelpSearchQuery();
}

void* QHelpSearchQuery_NewQHelpSearchQuery2(int field, char* wordList){
	return new QHelpSearchQuery(static_cast<QHelpSearchQuery::FieldName>(field), QString(wordList).split("|", QString::SkipEmptyParts));
}

