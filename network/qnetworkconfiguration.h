#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkConfiguration_NewQNetworkConfiguration();
void* QNetworkConfiguration_NewQNetworkConfiguration2(void* other);
int QNetworkConfiguration_BearerType(void* ptr);
int QNetworkConfiguration_BearerTypeFamily(void* ptr);
char* QNetworkConfiguration_BearerTypeName(void* ptr);
char* QNetworkConfiguration_Identifier(void* ptr);
int QNetworkConfiguration_IsRoamingAvailable(void* ptr);
int QNetworkConfiguration_IsValid(void* ptr);
char* QNetworkConfiguration_Name(void* ptr);
int QNetworkConfiguration_Purpose(void* ptr);
void QNetworkConfiguration_Swap(void* ptr, void* other);
int QNetworkConfiguration_Type(void* ptr);
void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr);

#ifdef __cplusplus
}
#endif