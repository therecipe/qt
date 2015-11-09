#ifdef __cplusplus
extern "C" {
#endif

void* QTextBoundaryFinder_NewQTextBoundaryFinder();
void* QTextBoundaryFinder_NewQTextBoundaryFinder3(int ty, char* stri);
void* QTextBoundaryFinder_NewQTextBoundaryFinder2(void* other);
int QTextBoundaryFinder_BoundaryReasons(void* ptr);
int QTextBoundaryFinder_IsAtBoundary(void* ptr);
int QTextBoundaryFinder_IsValid(void* ptr);
int QTextBoundaryFinder_Position(void* ptr);
void QTextBoundaryFinder_SetPosition(void* ptr, int position);
char* QTextBoundaryFinder_String(void* ptr);
void QTextBoundaryFinder_ToEnd(void* ptr);
int QTextBoundaryFinder_ToNextBoundary(void* ptr);
int QTextBoundaryFinder_ToPreviousBoundary(void* ptr);
void QTextBoundaryFinder_ToStart(void* ptr);
int QTextBoundaryFinder_Type(void* ptr);
void QTextBoundaryFinder_DestroyQTextBoundaryFinder(void* ptr);

#ifdef __cplusplus
}
#endif