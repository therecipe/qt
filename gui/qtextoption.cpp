#include "qtextoption.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextOption>
#include "_cgo_export.h"

class MyQTextOption: public QTextOption {
public:
};

void* QTextOption_NewQTextOption3(void* other){
	return new QTextOption(*static_cast<QTextOption*>(other));
}

void* QTextOption_NewQTextOption(){
	return new QTextOption();
}

void* QTextOption_NewQTextOption2(int alignment){
	return new QTextOption(static_cast<Qt::AlignmentFlag>(alignment));
}

int QTextOption_Alignment(void* ptr){
	return static_cast<QTextOption*>(ptr)->alignment();
}

int QTextOption_Flags(void* ptr){
	return static_cast<QTextOption*>(ptr)->flags();
}

void QTextOption_SetAlignment(void* ptr, int alignment){
	static_cast<QTextOption*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextOption_SetFlags(void* ptr, int flags){
	static_cast<QTextOption*>(ptr)->setFlags(static_cast<QTextOption::Flag>(flags));
}

void QTextOption_SetTabStop(void* ptr, double tabStop){
	static_cast<QTextOption*>(ptr)->setTabStop(static_cast<qreal>(tabStop));
}

void QTextOption_SetTextDirection(void* ptr, int direction){
	static_cast<QTextOption*>(ptr)->setTextDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextOption_SetUseDesignMetrics(void* ptr, int enable){
	static_cast<QTextOption*>(ptr)->setUseDesignMetrics(enable != 0);
}

void QTextOption_SetWrapMode(void* ptr, int mode){
	static_cast<QTextOption*>(ptr)->setWrapMode(static_cast<QTextOption::WrapMode>(mode));
}

double QTextOption_TabStop(void* ptr){
	return static_cast<double>(static_cast<QTextOption*>(ptr)->tabStop());
}

int QTextOption_TextDirection(void* ptr){
	return static_cast<QTextOption*>(ptr)->textDirection();
}

int QTextOption_UseDesignMetrics(void* ptr){
	return static_cast<QTextOption*>(ptr)->useDesignMetrics();
}

int QTextOption_WrapMode(void* ptr){
	return static_cast<QTextOption*>(ptr)->wrapMode();
}

void QTextOption_DestroyQTextOption(void* ptr){
	static_cast<QTextOption*>(ptr)->~QTextOption();
}

