#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QItemSelectionRange_Intersects(QtObjectPtr ptr, QtObjectPtr other);
QtObjectPtr QItemSelectionRange_NewQItemSelectionRange();
QtObjectPtr QItemSelectionRange_NewQItemSelectionRange2(QtObjectPtr other);
QtObjectPtr QItemSelectionRange_NewQItemSelectionRange4(QtObjectPtr index);
QtObjectPtr QItemSelectionRange_NewQItemSelectionRange3(QtObjectPtr topLeft, QtObjectPtr bottomRight);
int QItemSelectionRange_Bottom(QtObjectPtr ptr);
int QItemSelectionRange_Contains(QtObjectPtr ptr, QtObjectPtr index);
int QItemSelectionRange_Contains2(QtObjectPtr ptr, int row, int column, QtObjectPtr parentIndex);
int QItemSelectionRange_Height(QtObjectPtr ptr);
int QItemSelectionRange_IsEmpty(QtObjectPtr ptr);
int QItemSelectionRange_IsValid(QtObjectPtr ptr);
int QItemSelectionRange_Left(QtObjectPtr ptr);
QtObjectPtr QItemSelectionRange_Model(QtObjectPtr ptr);
QtObjectPtr QItemSelectionRange_Parent(QtObjectPtr ptr);
int QItemSelectionRange_Right(QtObjectPtr ptr);
int QItemSelectionRange_Top(QtObjectPtr ptr);
int QItemSelectionRange_Width(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif