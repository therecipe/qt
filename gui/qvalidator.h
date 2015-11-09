#ifdef __cplusplus
extern "C" {
#endif

void QValidator_ConnectChanged(void* ptr);
void QValidator_DisconnectChanged(void* ptr);
void QValidator_SetLocale(void* ptr, void* locale);
void QValidator_DestroyQValidator(void* ptr);

#ifdef __cplusplus
}
#endif