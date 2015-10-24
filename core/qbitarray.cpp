#include "qbitarray.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QBitArray>
#include "_cgo_export.h"

class MyQBitArray: public QBitArray {
public:
};

QtObjectPtr QBitArray_NewQBitArray(){
	return new QBitArray();
}

QtObjectPtr QBitArray_NewQBitArray4(QtObjectPtr other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

QtObjectPtr QBitArray_NewQBitArray3(QtObjectPtr other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

QtObjectPtr QBitArray_NewQBitArray2(int size, int value){
	return new QBitArray(size, value != 0);
}

int QBitArray_At(QtObjectPtr ptr, int i){
	return static_cast<QBitArray*>(ptr)->at(i);
}

void QBitArray_Clear(QtObjectPtr ptr){
	static_cast<QBitArray*>(ptr)->clear();
}

void QBitArray_ClearBit(QtObjectPtr ptr, int i){
	static_cast<QBitArray*>(ptr)->clearBit(i);
}

int QBitArray_Count(QtObjectPtr ptr){
	return static_cast<QBitArray*>(ptr)->count();
}

int QBitArray_Count2(QtObjectPtr ptr, int on){
	return static_cast<QBitArray*>(ptr)->count(on != 0);
}

int QBitArray_Fill(QtObjectPtr ptr, int value, int size){
	return static_cast<QBitArray*>(ptr)->fill(value != 0, size);
}

void QBitArray_Fill2(QtObjectPtr ptr, int value, int begin, int end){
	static_cast<QBitArray*>(ptr)->fill(value != 0, begin, end);
}

int QBitArray_IsEmpty(QtObjectPtr ptr){
	return static_cast<QBitArray*>(ptr)->isEmpty();
}

int QBitArray_IsNull(QtObjectPtr ptr){
	return static_cast<QBitArray*>(ptr)->isNull();
}

void QBitArray_Resize(QtObjectPtr ptr, int size){
	static_cast<QBitArray*>(ptr)->resize(size);
}

void QBitArray_SetBit(QtObjectPtr ptr, int i){
	static_cast<QBitArray*>(ptr)->setBit(i);
}

void QBitArray_SetBit2(QtObjectPtr ptr, int i, int value){
	static_cast<QBitArray*>(ptr)->setBit(i, value != 0);
}

int QBitArray_Size(QtObjectPtr ptr){
	return static_cast<QBitArray*>(ptr)->size();
}

void QBitArray_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QBitArray*>(ptr)->swap(*static_cast<QBitArray*>(other));
}

int QBitArray_TestBit(QtObjectPtr ptr, int i){
	return static_cast<QBitArray*>(ptr)->testBit(i);
}

int QBitArray_ToggleBit(QtObjectPtr ptr, int i){
	return static_cast<QBitArray*>(ptr)->toggleBit(i);
}

void QBitArray_Truncate(QtObjectPtr ptr, int pos){
	static_cast<QBitArray*>(ptr)->truncate(pos);
}

