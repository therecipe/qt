#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder();
QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder3(int ty, char* stri);
QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder2(QtObjectPtr other);
int QTextBoundaryFinder_BoundaryReasons(QtObjectPtr ptr);
int QTextBoundaryFinder_IsAtBoundary(QtObjectPtr ptr);
int QTextBoundaryFinder_IsValid(QtObjectPtr ptr);
int QTextBoundaryFinder_Position(QtObjectPtr ptr);
void QTextBoundaryFinder_SetPosition(QtObjectPtr ptr, int position);
char* QTextBoundaryFinder_String(QtObjectPtr ptr);
void QTextBoundaryFinder_ToEnd(QtObjectPtr ptr);
int QTextBoundaryFinder_ToNextBoundary(QtObjectPtr ptr);
int QTextBoundaryFinder_ToPreviousBoundary(QtObjectPtr ptr);
void QTextBoundaryFinder_ToStart(QtObjectPtr ptr);
int QTextBoundaryFinder_Type(QtObjectPtr ptr);
void QTextBoundaryFinder_DestroyQTextBoundaryFinder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif