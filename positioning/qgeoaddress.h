#ifdef __cplusplus
extern "C" {
#endif

void* QGeoAddress_NewQGeoAddress();
void* QGeoAddress_NewQGeoAddress2(void* other);
char* QGeoAddress_City(void* ptr);
void QGeoAddress_Clear(void* ptr);
char* QGeoAddress_Country(void* ptr);
char* QGeoAddress_CountryCode(void* ptr);
char* QGeoAddress_County(void* ptr);
char* QGeoAddress_District(void* ptr);
int QGeoAddress_IsEmpty(void* ptr);
int QGeoAddress_IsTextGenerated(void* ptr);
char* QGeoAddress_PostalCode(void* ptr);
void QGeoAddress_SetCity(void* ptr, char* city);
void QGeoAddress_SetCountry(void* ptr, char* country);
void QGeoAddress_SetCountryCode(void* ptr, char* countryCode);
void QGeoAddress_SetCounty(void* ptr, char* county);
void QGeoAddress_SetDistrict(void* ptr, char* district);
void QGeoAddress_SetPostalCode(void* ptr, char* postalCode);
void QGeoAddress_SetState(void* ptr, char* state);
void QGeoAddress_SetStreet(void* ptr, char* street);
void QGeoAddress_SetText(void* ptr, char* text);
char* QGeoAddress_Street(void* ptr);
char* QGeoAddress_Text(void* ptr);
void QGeoAddress_DestroyQGeoAddress(void* ptr);

#ifdef __cplusplus
}
#endif