#ifdef __cplusplus
extern "C" {
#endif

void* QSourceLocation_NewQSourceLocation();
void* QSourceLocation_NewQSourceLocation2(void* other);
void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c);
int QSourceLocation_IsNull(void* ptr);
void QSourceLocation_SetUri(void* ptr, void* newUri);
void QSourceLocation_DestroyQSourceLocation(void* ptr);

#ifdef __cplusplus
}
#endif