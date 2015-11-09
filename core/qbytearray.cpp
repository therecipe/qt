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

void QByteArray_Clear(void* ptr){
	static_cast<QByteArray*>(ptr)->clear();
}

int QByteArray_IndexOf3(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(const_cast<const char*>(str), from);
}

int QByteArray_IsNull(void* ptr){
	return static_cast<QByteArray*>(ptr)->isNull();
}

int QByteArray_LastIndexOf(void* ptr, void* ba, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_LastIndexOf3(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(const_cast<const char*>(str), from);
}

void* QByteArray_NewQByteArray(){
	return new QByteArray();
}

void* QByteArray_NewQByteArray6(void* other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray5(void* other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray2(char* data, int size){
	return new QByteArray(const_cast<const char*>(data), size);
}

void* QByteArray_NewQByteArray3(int size, char* ch){
	return new QByteArray(size, *ch);
}

void* QByteArray_Append5(void* ptr, char* ch){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*ch));
}

void* QByteArray_Append(void* ptr, void* ba){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*static_cast<QByteArray*>(ba)));
}

void* QByteArray_Append2(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(QString(str)));
}

void* QByteArray_Append3(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str)));
}

void* QByteArray_Append4(void* ptr, char* str, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str), len));
}

int QByteArray_Capacity(void* ptr){
	return static_cast<QByteArray*>(ptr)->capacity();
}

void QByteArray_Chop(void* ptr, int n){
	static_cast<QByteArray*>(ptr)->chop(n);
}

int QByteArray_Contains3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->contains(*ch);
}

int QByteArray_Contains(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->contains(*static_cast<QByteArray*>(ba));
}

int QByteArray_Contains2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->contains(const_cast<const char*>(str));
}

int QByteArray_Count4(void* ptr){
	return static_cast<QByteArray*>(ptr)->count();
}

int QByteArray_Count3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->count(*ch);
}

int QByteArray_Count(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->count(*static_cast<QByteArray*>(ba));
}

int QByteArray_Count2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->count(const_cast<const char*>(str));
}

int QByteArray_EndsWith3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->endsWith(*ch);
}

int QByteArray_EndsWith(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->endsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_EndsWith2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->endsWith(const_cast<const char*>(str));
}

void* QByteArray_Fill(void* ptr, char* ch, int size){
	return new QByteArray(static_cast<QByteArray*>(ptr)->fill(*ch, size));
}

void* QByteArray_QByteArray_FromBase64(void* base64){
	return new QByteArray(QByteArray::fromBase64(*static_cast<QByteArray*>(base64)));
}

void* QByteArray_QByteArray_FromBase642(void* base64, int options){
	return new QByteArray(QByteArray::fromBase64(*static_cast<QByteArray*>(base64), static_cast<QByteArray::Base64Option>(options)));
}

void* QByteArray_QByteArray_FromHex(void* hexEncoded){
	return new QByteArray(QByteArray::fromHex(*static_cast<QByteArray*>(hexEncoded)));
}

void* QByteArray_QByteArray_FromPercentEncoding(void* input, char* percent){
	return new QByteArray(QByteArray::fromPercentEncoding(*static_cast<QByteArray*>(input), *percent));
}

void* QByteArray_QByteArray_FromRawData(char* data, int size){
	return new QByteArray(QByteArray::fromRawData(const_cast<const char*>(data), size));
}

int QByteArray_IndexOf4(void* ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*ch, from);
}

int QByteArray_IndexOf(void* ptr, void* ba, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_IndexOf2(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(QString(str), from);
}

int QByteArray_IsEmpty(void* ptr){
	return static_cast<QByteArray*>(ptr)->isEmpty();
}

int QByteArray_LastIndexOf4(void* ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*ch, from);
}

int QByteArray_LastIndexOf2(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(QString(str), from);
}

void* QByteArray_Left(void* ptr, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->left(len));
}

void* QByteArray_LeftJustified(void* ptr, int width, char* fill, int truncate){
	return new QByteArray(static_cast<QByteArray*>(ptr)->leftJustified(width, *fill, truncate != 0));
}

int QByteArray_Length(void* ptr){
	return static_cast<QByteArray*>(ptr)->length();
}

void* QByteArray_Mid(void* ptr, int pos, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->mid(pos, len));
}

void* QByteArray_QByteArray_Number(int n, int base){
	return new QByteArray(QByteArray::number(n, base));
}

void* QByteArray_Prepend4(void* ptr, char* ch){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(*ch));
}

void* QByteArray_Prepend(void* ptr, void* ba){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(*static_cast<QByteArray*>(ba)));
}

void* QByteArray_Prepend2(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(const_cast<const char*>(str)));
}

void* QByteArray_Prepend3(void* ptr, char* str, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(const_cast<const char*>(str), len));
}

void QByteArray_Push_back3(void* ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_back(*ch);
}

void QByteArray_Push_back(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->push_back(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_back2(void* ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_back(const_cast<const char*>(str));
}

void QByteArray_Push_front3(void* ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_front(*ch);
}

void QByteArray_Push_front(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->push_front(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_front2(void* ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_front(const_cast<const char*>(str));
}

void* QByteArray_Repeated(void* ptr, int times){
	return new QByteArray(static_cast<QByteArray*>(ptr)->repeated(times));
}

void QByteArray_Reserve(void* ptr, int size){
	static_cast<QByteArray*>(ptr)->reserve(size);
}

void QByteArray_Resize(void* ptr, int size){
	static_cast<QByteArray*>(ptr)->resize(size);
}

void* QByteArray_Right(void* ptr, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->right(len));
}

void* QByteArray_RightJustified(void* ptr, int width, char* fill, int truncate){
	return new QByteArray(static_cast<QByteArray*>(ptr)->rightJustified(width, *fill, truncate != 0));
}

void* QByteArray_SetNum(void* ptr, int n, int base){
	return new QByteArray(static_cast<QByteArray*>(ptr)->setNum(n, base));
}

int QByteArray_Size(void* ptr){
	return static_cast<QByteArray*>(ptr)->size();
}

void QByteArray_Squeeze(void* ptr){
	static_cast<QByteArray*>(ptr)->squeeze();
}

int QByteArray_StartsWith3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->startsWith(*ch);
}

int QByteArray_StartsWith(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->startsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_StartsWith2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->startsWith(const_cast<const char*>(str));
}

void QByteArray_Swap(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->swap(*static_cast<QByteArray*>(other));
}

void* QByteArray_ToBase64(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toBase64());
}

void* QByteArray_ToBase642(void* ptr, int options){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toBase64(static_cast<QByteArray::Base64Option>(options)));
}

void* QByteArray_ToHex(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toHex());
}

int QByteArray_ToInt(void* ptr, int ok, int base){
	return static_cast<QByteArray*>(ptr)->toInt(NULL, base);
}

void* QByteArray_ToPercentEncoding(void* ptr, void* exclude, void* include, char* percent){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toPercentEncoding(*static_cast<QByteArray*>(exclude), *static_cast<QByteArray*>(include), *percent));
}

void QByteArray_Truncate(void* ptr, int pos){
	static_cast<QByteArray*>(ptr)->truncate(pos);
}

void QByteArray_DestroyQByteArray(void* ptr){
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

void* QByteArray_Simplified(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->simplified());
}

void* QByteArray_ToLower(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toLower());
}

void* QByteArray_ToUpper(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toUpper());
}

void* QByteArray_Trimmed(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->trimmed());
}

