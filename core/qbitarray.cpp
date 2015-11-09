#include "qbitarray.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBitArray>
#include "_cgo_export.h"

class MyQBitArray: public QBitArray {
public:
};

void* QBitArray_NewQBitArray(){
	return new QBitArray();
}

void* QBitArray_NewQBitArray4(void* other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray3(void* other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray2(int size, int value){
	return new QBitArray(size, value != 0);
}

int QBitArray_At(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->at(i);
}

void QBitArray_Clear(void* ptr){
	static_cast<QBitArray*>(ptr)->clear();
}

void QBitArray_ClearBit(void* ptr, int i){
	static_cast<QBitArray*>(ptr)->clearBit(i);
}

int QBitArray_Count(void* ptr){
	return static_cast<QBitArray*>(ptr)->count();
}

int QBitArray_Count2(void* ptr, int on){
	return static_cast<QBitArray*>(ptr)->count(on != 0);
}

int QBitArray_Fill(void* ptr, int value, int size){
	return static_cast<QBitArray*>(ptr)->fill(value != 0, size);
}

void QBitArray_Fill2(void* ptr, int value, int begin, int end){
	static_cast<QBitArray*>(ptr)->fill(value != 0, begin, end);
}

int QBitArray_IsEmpty(void* ptr){
	return static_cast<QBitArray*>(ptr)->isEmpty();
}

int QBitArray_IsNull(void* ptr){
	return static_cast<QBitArray*>(ptr)->isNull();
}

void QBitArray_Resize(void* ptr, int size){
	static_cast<QBitArray*>(ptr)->resize(size);
}

void QBitArray_SetBit(void* ptr, int i){
	static_cast<QBitArray*>(ptr)->setBit(i);
}

void QBitArray_SetBit2(void* ptr, int i, int value){
	static_cast<QBitArray*>(ptr)->setBit(i, value != 0);
}

int QBitArray_Size(void* ptr){
	return static_cast<QBitArray*>(ptr)->size();
}

void QBitArray_Swap(void* ptr, void* other){
	static_cast<QBitArray*>(ptr)->swap(*static_cast<QBitArray*>(other));
}

int QBitArray_TestBit(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->testBit(i);
}

int QBitArray_ToggleBit(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->toggleBit(i);
}

void QBitArray_Truncate(void* ptr, int pos){
	static_cast<QBitArray*>(ptr)->truncate(pos);
}

