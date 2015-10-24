#include "qtextoption.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextOption>
#include "_cgo_export.h"

class MyQTextOption: public QTextOption {
public:
};

QtObjectPtr QTextOption_NewQTextOption3(QtObjectPtr other){
	return new QTextOption(*static_cast<QTextOption*>(other));
}

QtObjectPtr QTextOption_NewQTextOption(){
	return new QTextOption();
}

QtObjectPtr QTextOption_NewQTextOption2(int alignment){
	return new QTextOption(static_cast<Qt::AlignmentFlag>(alignment));
}

int QTextOption_Alignment(QtObjectPtr ptr){
	return static_cast<QTextOption*>(ptr)->alignment();
}

int QTextOption_Flags(QtObjectPtr ptr){
	return static_cast<QTextOption*>(ptr)->flags();
}

void QTextOption_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QTextOption*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextOption_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QTextOption*>(ptr)->setFlags(static_cast<QTextOption::Flag>(flags));
}

void QTextOption_SetTextDirection(QtObjectPtr ptr, int direction){
	static_cast<QTextOption*>(ptr)->setTextDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextOption_SetUseDesignMetrics(QtObjectPtr ptr, int enable){
	static_cast<QTextOption*>(ptr)->setUseDesignMetrics(enable != 0);
}

void QTextOption_SetWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QTextOption*>(ptr)->setWrapMode(static_cast<QTextOption::WrapMode>(mode));
}

int QTextOption_TextDirection(QtObjectPtr ptr){
	return static_cast<QTextOption*>(ptr)->textDirection();
}

int QTextOption_UseDesignMetrics(QtObjectPtr ptr){
	return static_cast<QTextOption*>(ptr)->useDesignMetrics();
}

int QTextOption_WrapMode(QtObjectPtr ptr){
	return static_cast<QTextOption*>(ptr)->wrapMode();
}

void QTextOption_DestroyQTextOption(QtObjectPtr ptr){
	static_cast<QTextOption*>(ptr)->~QTextOption();
}

