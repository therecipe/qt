#include "qtextformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBrush>
#include <QTextFormat>
#include "_cgo_export.h"

class MyQTextFormat: public QTextFormat {
public:
};

QtObjectPtr QTextFormat_NewQTextFormat3(QtObjectPtr other){
	return new QTextFormat(*static_cast<QTextFormat*>(other));
}

void QTextFormat_SetObjectIndex(QtObjectPtr ptr, int index){
	static_cast<QTextFormat*>(ptr)->setObjectIndex(index);
}

QtObjectPtr QTextFormat_NewQTextFormat(){
	return new QTextFormat();
}

QtObjectPtr QTextFormat_NewQTextFormat2(int ty){
	return new QTextFormat(ty);
}

int QTextFormat_BoolProperty(QtObjectPtr ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->boolProperty(propertyId);
}

void QTextFormat_ClearBackground(QtObjectPtr ptr){
	static_cast<QTextFormat*>(ptr)->clearBackground();
}

void QTextFormat_ClearForeground(QtObjectPtr ptr){
	static_cast<QTextFormat*>(ptr)->clearForeground();
}

void QTextFormat_ClearProperty(QtObjectPtr ptr, int propertyId){
	static_cast<QTextFormat*>(ptr)->clearProperty(propertyId);
}

int QTextFormat_HasProperty(QtObjectPtr ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->hasProperty(propertyId);
}

int QTextFormat_IntProperty(QtObjectPtr ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->intProperty(propertyId);
}

int QTextFormat_IsBlockFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isBlockFormat();
}

int QTextFormat_IsCharFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isCharFormat();
}

int QTextFormat_IsEmpty(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isEmpty();
}

int QTextFormat_IsFrameFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isFrameFormat();
}

int QTextFormat_IsImageFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isImageFormat();
}

int QTextFormat_IsListFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isListFormat();
}

int QTextFormat_IsTableCellFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isTableCellFormat();
}

int QTextFormat_IsTableFormat(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isTableFormat();
}

int QTextFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->isValid();
}

int QTextFormat_LayoutDirection(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->layoutDirection();
}

void QTextFormat_Merge(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QTextFormat*>(ptr)->merge(*static_cast<QTextFormat*>(other));
}

int QTextFormat_ObjectIndex(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->objectIndex();
}

int QTextFormat_ObjectType(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->objectType();
}

char* QTextFormat_Property(QtObjectPtr ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->property(propertyId).toString().toUtf8().data();
}

int QTextFormat_PropertyCount(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->propertyCount();
}

void QTextFormat_SetBackground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QTextFormat*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetForeground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QTextFormat*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetLayoutDirection(QtObjectPtr ptr, int direction){
	static_cast<QTextFormat*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextFormat_SetObjectType(QtObjectPtr ptr, int ty){
	static_cast<QTextFormat*>(ptr)->setObjectType(ty);
}

void QTextFormat_SetProperty(QtObjectPtr ptr, int propertyId, char* value){
	static_cast<QTextFormat*>(ptr)->setProperty(propertyId, QVariant(value));
}

char* QTextFormat_StringProperty(QtObjectPtr ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->stringProperty(propertyId).toUtf8().data();
}

void QTextFormat_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QTextFormat*>(ptr)->swap(*static_cast<QTextFormat*>(other));
}

int QTextFormat_Type(QtObjectPtr ptr){
	return static_cast<QTextFormat*>(ptr)->type();
}

void QTextFormat_DestroyQTextFormat(QtObjectPtr ptr){
	static_cast<QTextFormat*>(ptr)->~QTextFormat();
}

