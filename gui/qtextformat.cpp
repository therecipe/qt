#include "qtextformat.h"
#include <QUrl>
#include <QModelIndex>
#include <QBrush>
#include <QColor>
#include <QString>
#include <QVariant>
#include <QTextFormat>
#include "_cgo_export.h"

class MyQTextFormat: public QTextFormat {
public:
};

void* QTextFormat_NewQTextFormat3(void* other){
	return new QTextFormat(*static_cast<QTextFormat*>(other));
}

void QTextFormat_SetObjectIndex(void* ptr, int index){
	static_cast<QTextFormat*>(ptr)->setObjectIndex(index);
}

void* QTextFormat_NewQTextFormat(){
	return new QTextFormat();
}

void* QTextFormat_NewQTextFormat2(int ty){
	return new QTextFormat(ty);
}

void* QTextFormat_Background(void* ptr){
	return new QBrush(static_cast<QTextFormat*>(ptr)->background());
}

int QTextFormat_BoolProperty(void* ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->boolProperty(propertyId);
}

void* QTextFormat_BrushProperty(void* ptr, int propertyId){
	return new QBrush(static_cast<QTextFormat*>(ptr)->brushProperty(propertyId));
}

void QTextFormat_ClearBackground(void* ptr){
	static_cast<QTextFormat*>(ptr)->clearBackground();
}

void QTextFormat_ClearForeground(void* ptr){
	static_cast<QTextFormat*>(ptr)->clearForeground();
}

void QTextFormat_ClearProperty(void* ptr, int propertyId){
	static_cast<QTextFormat*>(ptr)->clearProperty(propertyId);
}

void* QTextFormat_ColorProperty(void* ptr, int propertyId){
	return new QColor(static_cast<QTextFormat*>(ptr)->colorProperty(propertyId));
}

double QTextFormat_DoubleProperty(void* ptr, int propertyId){
	return static_cast<double>(static_cast<QTextFormat*>(ptr)->doubleProperty(propertyId));
}

void* QTextFormat_Foreground(void* ptr){
	return new QBrush(static_cast<QTextFormat*>(ptr)->foreground());
}

int QTextFormat_HasProperty(void* ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->hasProperty(propertyId);
}

int QTextFormat_IntProperty(void* ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->intProperty(propertyId);
}

int QTextFormat_IsBlockFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isBlockFormat();
}

int QTextFormat_IsCharFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isCharFormat();
}

int QTextFormat_IsEmpty(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isEmpty();
}

int QTextFormat_IsFrameFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isFrameFormat();
}

int QTextFormat_IsImageFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isImageFormat();
}

int QTextFormat_IsListFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isListFormat();
}

int QTextFormat_IsTableCellFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isTableCellFormat();
}

int QTextFormat_IsTableFormat(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isTableFormat();
}

int QTextFormat_IsValid(void* ptr){
	return static_cast<QTextFormat*>(ptr)->isValid();
}

int QTextFormat_LayoutDirection(void* ptr){
	return static_cast<QTextFormat*>(ptr)->layoutDirection();
}

void QTextFormat_Merge(void* ptr, void* other){
	static_cast<QTextFormat*>(ptr)->merge(*static_cast<QTextFormat*>(other));
}

int QTextFormat_ObjectIndex(void* ptr){
	return static_cast<QTextFormat*>(ptr)->objectIndex();
}

int QTextFormat_ObjectType(void* ptr){
	return static_cast<QTextFormat*>(ptr)->objectType();
}

void* QTextFormat_Property(void* ptr, int propertyId){
	return new QVariant(static_cast<QTextFormat*>(ptr)->property(propertyId));
}

int QTextFormat_PropertyCount(void* ptr){
	return static_cast<QTextFormat*>(ptr)->propertyCount();
}

void QTextFormat_SetBackground(void* ptr, void* brush){
	static_cast<QTextFormat*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetForeground(void* ptr, void* brush){
	static_cast<QTextFormat*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetLayoutDirection(void* ptr, int direction){
	static_cast<QTextFormat*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextFormat_SetObjectType(void* ptr, int ty){
	static_cast<QTextFormat*>(ptr)->setObjectType(ty);
}

void QTextFormat_SetProperty(void* ptr, int propertyId, void* value){
	static_cast<QTextFormat*>(ptr)->setProperty(propertyId, *static_cast<QVariant*>(value));
}

char* QTextFormat_StringProperty(void* ptr, int propertyId){
	return static_cast<QTextFormat*>(ptr)->stringProperty(propertyId).toUtf8().data();
}

void QTextFormat_Swap(void* ptr, void* other){
	static_cast<QTextFormat*>(ptr)->swap(*static_cast<QTextFormat*>(other));
}

int QTextFormat_Type(void* ptr){
	return static_cast<QTextFormat*>(ptr)->type();
}

void QTextFormat_DestroyQTextFormat(void* ptr){
	static_cast<QTextFormat*>(ptr)->~QTextFormat();
}

