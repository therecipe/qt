#include "qmimetype.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QMimeType>
#include "_cgo_export.h"

class MyQMimeType: public QMimeType {
public:
};

void* QMimeType_NewQMimeType(){
	return new QMimeType();
}

void* QMimeType_NewQMimeType2(void* other){
	return new QMimeType(*static_cast<QMimeType*>(other));
}

char* QMimeType_FilterString(void* ptr){
	return static_cast<QMimeType*>(ptr)->filterString().toUtf8().data();
}

char* QMimeType_GenericIconName(void* ptr){
	return static_cast<QMimeType*>(ptr)->genericIconName().toUtf8().data();
}

char* QMimeType_GlobPatterns(void* ptr){
	return static_cast<QMimeType*>(ptr)->globPatterns().join("|").toUtf8().data();
}

char* QMimeType_IconName(void* ptr){
	return static_cast<QMimeType*>(ptr)->iconName().toUtf8().data();
}

int QMimeType_Inherits(void* ptr, char* mimeTypeName){
	return static_cast<QMimeType*>(ptr)->inherits(QString(mimeTypeName));
}

int QMimeType_IsDefault(void* ptr){
	return static_cast<QMimeType*>(ptr)->isDefault();
}

int QMimeType_IsValid(void* ptr){
	return static_cast<QMimeType*>(ptr)->isValid();
}

char* QMimeType_Name(void* ptr){
	return static_cast<QMimeType*>(ptr)->name().toUtf8().data();
}

void QMimeType_DestroyQMimeType(void* ptr){
	static_cast<QMimeType*>(ptr)->~QMimeType();
}

char* QMimeType_Aliases(void* ptr){
	return static_cast<QMimeType*>(ptr)->aliases().join("|").toUtf8().data();
}

char* QMimeType_AllAncestors(void* ptr){
	return static_cast<QMimeType*>(ptr)->allAncestors().join("|").toUtf8().data();
}

char* QMimeType_Comment(void* ptr){
	return static_cast<QMimeType*>(ptr)->comment().toUtf8().data();
}

char* QMimeType_ParentMimeTypes(void* ptr){
	return static_cast<QMimeType*>(ptr)->parentMimeTypes().join("|").toUtf8().data();
}

char* QMimeType_PreferredSuffix(void* ptr){
	return static_cast<QMimeType*>(ptr)->preferredSuffix().toUtf8().data();
}

char* QMimeType_Suffixes(void* ptr){
	return static_cast<QMimeType*>(ptr)->suffixes().join("|").toUtf8().data();
}

void QMimeType_Swap(void* ptr, void* other){
	static_cast<QMimeType*>(ptr)->swap(*static_cast<QMimeType*>(other));
}

