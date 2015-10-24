#include "qjsonarray.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJsonValue>
#include <QJsonArray>
#include "_cgo_export.h"

class MyQJsonArray: public QJsonArray {
public:
};

void QJsonArray_Append(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QJsonArray*>(ptr)->append(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Contains(QtObjectPtr ptr, QtObjectPtr value){
	return static_cast<QJsonArray*>(ptr)->contains(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Count(QtObjectPtr ptr){
	return static_cast<QJsonArray*>(ptr)->count();
}

int QJsonArray_Empty(QtObjectPtr ptr){
	return static_cast<QJsonArray*>(ptr)->empty();
}

int QJsonArray_IsEmpty(QtObjectPtr ptr){
	return static_cast<QJsonArray*>(ptr)->isEmpty();
}

void QJsonArray_Pop_back(QtObjectPtr ptr){
	static_cast<QJsonArray*>(ptr)->pop_back();
}

void QJsonArray_Pop_front(QtObjectPtr ptr){
	static_cast<QJsonArray*>(ptr)->pop_front();
}

void QJsonArray_Prepend(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QJsonArray*>(ptr)->prepend(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_back(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QJsonArray*>(ptr)->push_back(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_front(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QJsonArray*>(ptr)->push_front(*static_cast<QJsonValue*>(value));
}

void QJsonArray_RemoveAt(QtObjectPtr ptr, int i){
	static_cast<QJsonArray*>(ptr)->removeAt(i);
}

void QJsonArray_RemoveFirst(QtObjectPtr ptr){
	static_cast<QJsonArray*>(ptr)->removeFirst();
}

void QJsonArray_RemoveLast(QtObjectPtr ptr){
	static_cast<QJsonArray*>(ptr)->removeLast();
}

int QJsonArray_Size(QtObjectPtr ptr){
	return static_cast<QJsonArray*>(ptr)->size();
}

void QJsonArray_DestroyQJsonArray(QtObjectPtr ptr){
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

