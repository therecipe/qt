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

QtObjectPtr QMimeType_NewQMimeType(){
	return new QMimeType();
}

QtObjectPtr QMimeType_NewQMimeType2(QtObjectPtr other){
	return new QMimeType(*static_cast<QMimeType*>(other));
}

char* QMimeType_FilterString(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->filterString().toUtf8().data();
}

char* QMimeType_GenericIconName(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->genericIconName().toUtf8().data();
}

char* QMimeType_GlobPatterns(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->globPatterns().join("|").toUtf8().data();
}

char* QMimeType_IconName(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->iconName().toUtf8().data();
}

int QMimeType_Inherits(QtObjectPtr ptr, char* mimeTypeName){
	return static_cast<QMimeType*>(ptr)->inherits(QString(mimeTypeName));
}

int QMimeType_IsDefault(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->isDefault();
}

int QMimeType_IsValid(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->isValid();
}

char* QMimeType_Name(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->name().toUtf8().data();
}

void QMimeType_DestroyQMimeType(QtObjectPtr ptr){
	static_cast<QMimeType*>(ptr)->~QMimeType();
}

char* QMimeType_Aliases(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->aliases().join("|").toUtf8().data();
}

char* QMimeType_AllAncestors(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->allAncestors().join("|").toUtf8().data();
}

char* QMimeType_Comment(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->comment().toUtf8().data();
}

char* QMimeType_ParentMimeTypes(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->parentMimeTypes().join("|").toUtf8().data();
}

char* QMimeType_PreferredSuffix(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->preferredSuffix().toUtf8().data();
}

char* QMimeType_Suffixes(QtObjectPtr ptr){
	return static_cast<QMimeType*>(ptr)->suffixes().join("|").toUtf8().data();
}

void QMimeType_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QMimeType*>(ptr)->swap(*static_cast<QMimeType*>(other));
}

