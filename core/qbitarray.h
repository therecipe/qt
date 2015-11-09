#ifdef __cplusplus
extern "C" {
#endif

void* QBitArray_NewQBitArray();
void* QBitArray_NewQBitArray4(void* other);
void* QBitArray_NewQBitArray3(void* other);
void* QBitArray_NewQBitArray2(int size, int value);
int QBitArray_At(void* ptr, int i);
void QBitArray_Clear(void* ptr);
void QBitArray_ClearBit(void* ptr, int i);
int QBitArray_Count(void* ptr);
int QBitArray_Count2(void* ptr, int on);
int QBitArray_Fill(void* ptr, int value, int size);
void QBitArray_Fill2(void* ptr, int value, int begin, int end);
int QBitArray_IsEmpty(void* ptr);
int QBitArray_IsNull(void* ptr);
void QBitArray_Resize(void* ptr, int size);
void QBitArray_SetBit(void* ptr, int i);
void QBitArray_SetBit2(void* ptr, int i, int value);
int QBitArray_Size(void* ptr);
void QBitArray_Swap(void* ptr, void* other);
int QBitArray_TestBit(void* ptr, int i);
int QBitArray_ToggleBit(void* ptr, int i);
void QBitArray_Truncate(void* ptr, int pos);

#ifdef __cplusplus
}
#endif