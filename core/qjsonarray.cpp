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

void QJsonArray_Append(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->append(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Contains(void* ptr, void* value){
	return static_cast<QJsonArray*>(ptr)->contains(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Count(void* ptr){
	return static_cast<QJsonArray*>(ptr)->count();
}

int QJsonArray_Empty(void* ptr){
	return static_cast<QJsonArray*>(ptr)->empty();
}

void* QJsonArray_QJsonArray_FromStringList(char* list){
	return new QJsonArray(QJsonArray::fromStringList(QString(list).split("|", QString::SkipEmptyParts)));
}

int QJsonArray_IsEmpty(void* ptr){
	return static_cast<QJsonArray*>(ptr)->isEmpty();
}

void QJsonArray_Pop_back(void* ptr){
	static_cast<QJsonArray*>(ptr)->pop_back();
}

void QJsonArray_Pop_front(void* ptr){
	static_cast<QJsonArray*>(ptr)->pop_front();
}

void QJsonArray_Prepend(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->prepend(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_back(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->push_back(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_front(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->push_front(*static_cast<QJsonValue*>(value));
}

void QJsonArray_RemoveAt(void* ptr, int i){
	static_cast<QJsonArray*>(ptr)->removeAt(i);
}

void QJsonArray_RemoveFirst(void* ptr){
	static_cast<QJsonArray*>(ptr)->removeFirst();
}

void QJsonArray_RemoveLast(void* ptr){
	static_cast<QJsonArray*>(ptr)->removeLast();
}

int QJsonArray_Size(void* ptr){
	return static_cast<QJsonArray*>(ptr)->size();
}

void QJsonArray_DestroyQJsonArray(void* ptr){
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

