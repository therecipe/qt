#include "qbytearray.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include "_cgo_export.h"

class MyQByteArray: public QByteArray {
public:
};

void QByteArray_Clear(QtObjectPtr ptr){
	static_cast<QByteArray*>(ptr)->clear();
}

int QByteArray_IndexOf3(QtObjectPtr ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(const_cast<const char*>(str), from);
}

int QByteArray_IsNull(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->isNull();
}

int QByteArray_LastIndexOf(QtObjectPtr ptr, QtObjectPtr ba, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_LastIndexOf3(QtObjectPtr ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(const_cast<const char*>(str), from);
}

QtObjectPtr QByteArray_NewQByteArray(){
	return new QByteArray();
}

QtObjectPtr QByteArray_NewQByteArray6(QtObjectPtr other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

QtObjectPtr QByteArray_NewQByteArray5(QtObjectPtr other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

QtObjectPtr QByteArray_NewQByteArray2(char* data, int size){
	return new QByteArray(const_cast<const char*>(data), size);
}

QtObjectPtr QByteArray_NewQByteArray3(int size, char* ch){
	return new QByteArray(size, *ch);
}

int QByteArray_Capacity(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->capacity();
}

void QByteArray_Chop(QtObjectPtr ptr, int n){
	static_cast<QByteArray*>(ptr)->chop(n);
}

int QByteArray_Contains3(QtObjectPtr ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->contains(*ch);
}

int QByteArray_Contains(QtObjectPtr ptr, QtObjectPtr ba){
	return static_cast<QByteArray*>(ptr)->contains(*static_cast<QByteArray*>(ba));
}

int QByteArray_Contains2(QtObjectPtr ptr, char* str){
	return static_cast<QByteArray*>(ptr)->contains(const_cast<const char*>(str));
}

int QByteArray_Count4(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->count();
}

int QByteArray_Count3(QtObjectPtr ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->count(*ch);
}

int QByteArray_Count(QtObjectPtr ptr, QtObjectPtr ba){
	return static_cast<QByteArray*>(ptr)->count(*static_cast<QByteArray*>(ba));
}

int QByteArray_Count2(QtObjectPtr ptr, char* str){
	return static_cast<QByteArray*>(ptr)->count(const_cast<const char*>(str));
}

int QByteArray_EndsWith3(QtObjectPtr ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->endsWith(*ch);
}

int QByteArray_EndsWith(QtObjectPtr ptr, QtObjectPtr ba){
	return static_cast<QByteArray*>(ptr)->endsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_EndsWith2(QtObjectPtr ptr, char* str){
	return static_cast<QByteArray*>(ptr)->endsWith(const_cast<const char*>(str));
}

int QByteArray_IndexOf4(QtObjectPtr ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*ch, from);
}

int QByteArray_IndexOf(QtObjectPtr ptr, QtObjectPtr ba, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_IndexOf2(QtObjectPtr ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(QString(str), from);
}

int QByteArray_IsEmpty(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->isEmpty();
}

int QByteArray_LastIndexOf4(QtObjectPtr ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*ch, from);
}

int QByteArray_LastIndexOf2(QtObjectPtr ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(QString(str), from);
}

int QByteArray_Length(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->length();
}

void QByteArray_Push_back3(QtObjectPtr ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_back(*ch);
}

void QByteArray_Push_back(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QByteArray*>(ptr)->push_back(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_back2(QtObjectPtr ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_back(const_cast<const char*>(str));
}

void QByteArray_Push_front3(QtObjectPtr ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_front(*ch);
}

void QByteArray_Push_front(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QByteArray*>(ptr)->push_front(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_front2(QtObjectPtr ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_front(const_cast<const char*>(str));
}

void QByteArray_Reserve(QtObjectPtr ptr, int size){
	static_cast<QByteArray*>(ptr)->reserve(size);
}

void QByteArray_Resize(QtObjectPtr ptr, int size){
	static_cast<QByteArray*>(ptr)->resize(size);
}

int QByteArray_Size(QtObjectPtr ptr){
	return static_cast<QByteArray*>(ptr)->size();
}

void QByteArray_Squeeze(QtObjectPtr ptr){
	static_cast<QByteArray*>(ptr)->squeeze();
}

int QByteArray_StartsWith3(QtObjectPtr ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->startsWith(*ch);
}

int QByteArray_StartsWith(QtObjectPtr ptr, QtObjectPtr ba){
	return static_cast<QByteArray*>(ptr)->startsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_StartsWith2(QtObjectPtr ptr, char* str){
	return static_cast<QByteArray*>(ptr)->startsWith(const_cast<const char*>(str));
}

void QByteArray_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QByteArray*>(ptr)->swap(*static_cast<QByteArray*>(other));
}

int QByteArray_ToInt(QtObjectPtr ptr, int ok, int base){
	return static_cast<QByteArray*>(ptr)->toInt(NULL, base);
}

void QByteArray_Truncate(QtObjectPtr ptr, int pos){
	static_cast<QByteArray*>(ptr)->truncate(pos);
}

void QByteArray_DestroyQByteArray(QtObjectPtr ptr){
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

