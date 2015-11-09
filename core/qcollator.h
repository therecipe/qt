#ifdef __cplusplus
extern "C" {
#endif

int QCollator_CaseSensitivity(void* ptr);
int QCollator_IgnorePunctuation(void* ptr);
int QCollator_NumericMode(void* ptr);
void QCollator_SetCaseSensitivity(void* ptr, int sensitivity);
void QCollator_SetIgnorePunctuation(void* ptr, int on);
void QCollator_SetNumericMode(void* ptr, int on);
void* QCollator_NewQCollator3(void* other);
void* QCollator_NewQCollator2(void* other);
void* QCollator_NewQCollator(void* locale);
void QCollator_SetLocale(void* ptr, void* locale);
void QCollator_Swap(void* ptr, void* other);
void QCollator_DestroyQCollator(void* ptr);
int QCollator_Compare3(void* ptr, void* s1, int len1, void* s2, int len2);
int QCollator_Compare(void* ptr, char* s1, char* s2);
int QCollator_Compare2(void* ptr, void* s1, void* s2);

#ifdef __cplusplus
}
#endif