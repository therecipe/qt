#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBitArray_NewQBitArray();
QtObjectPtr QBitArray_NewQBitArray4(QtObjectPtr other);
QtObjectPtr QBitArray_NewQBitArray3(QtObjectPtr other);
QtObjectPtr QBitArray_NewQBitArray2(int size, int value);
int QBitArray_At(QtObjectPtr ptr, int i);
void QBitArray_Clear(QtObjectPtr ptr);
void QBitArray_ClearBit(QtObjectPtr ptr, int i);
int QBitArray_Count(QtObjectPtr ptr);
int QBitArray_Count2(QtObjectPtr ptr, int on);
int QBitArray_Fill(QtObjectPtr ptr, int value, int size);
void QBitArray_Fill2(QtObjectPtr ptr, int value, int begin, int end);
int QBitArray_IsEmpty(QtObjectPtr ptr);
int QBitArray_IsNull(QtObjectPtr ptr);
void QBitArray_Resize(QtObjectPtr ptr, int size);
void QBitArray_SetBit(QtObjectPtr ptr, int i);
void QBitArray_SetBit2(QtObjectPtr ptr, int i, int value);
int QBitArray_Size(QtObjectPtr ptr);
void QBitArray_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QBitArray_TestBit(QtObjectPtr ptr, int i);
int QBitArray_ToggleBit(QtObjectPtr ptr, int i);
void QBitArray_Truncate(QtObjectPtr ptr, int pos);

#ifdef __cplusplus
}
#endif