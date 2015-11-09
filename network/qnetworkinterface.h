#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkInterface_NewQNetworkInterface();
void* QNetworkInterface_NewQNetworkInterface2(void* other);
int QNetworkInterface_Flags(void* ptr);
char* QNetworkInterface_HardwareAddress(void* ptr);
char* QNetworkInterface_HumanReadableName(void* ptr);
int QNetworkInterface_Index(void* ptr);
int QNetworkInterface_IsValid(void* ptr);
char* QNetworkInterface_Name(void* ptr);
void QNetworkInterface_Swap(void* ptr, void* other);
void QNetworkInterface_DestroyQNetworkInterface(void* ptr);

#ifdef __cplusplus
}
#endif