#ifdef __cplusplus
extern "C" {
#endif

void* QDomImplementation_NewQDomImplementation();
void* QDomImplementation_NewQDomImplementation2(void* x);
int QDomImplementation_HasFeature(void* ptr, char* feature, char* version);
int QDomImplementation_QDomImplementation_InvalidDataPolicy();
int QDomImplementation_IsNull(void* ptr);
void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(int policy);
void QDomImplementation_DestroyQDomImplementation(void* ptr);

#ifdef __cplusplus
}
#endif