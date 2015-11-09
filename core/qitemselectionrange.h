#ifdef __cplusplus
extern "C" {
#endif

int QItemSelectionRange_Intersects(void* ptr, void* other);
void* QItemSelectionRange_NewQItemSelectionRange();
void* QItemSelectionRange_NewQItemSelectionRange2(void* other);
void* QItemSelectionRange_NewQItemSelectionRange4(void* index);
void* QItemSelectionRange_NewQItemSelectionRange3(void* topLeft, void* bottomRight);
int QItemSelectionRange_Bottom(void* ptr);
int QItemSelectionRange_Contains(void* ptr, void* index);
int QItemSelectionRange_Contains2(void* ptr, int row, int column, void* parentIndex);
int QItemSelectionRange_Height(void* ptr);
int QItemSelectionRange_IsEmpty(void* ptr);
int QItemSelectionRange_IsValid(void* ptr);
int QItemSelectionRange_Left(void* ptr);
void* QItemSelectionRange_Model(void* ptr);
void* QItemSelectionRange_Parent(void* ptr);
int QItemSelectionRange_Right(void* ptr);
int QItemSelectionRange_Top(void* ptr);
int QItemSelectionRange_Width(void* ptr);

#ifdef __cplusplus
}
#endif