#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoAddress_NewQGeoAddress();
QtObjectPtr QGeoAddress_NewQGeoAddress2(QtObjectPtr other);
char* QGeoAddress_City(QtObjectPtr ptr);
void QGeoAddress_Clear(QtObjectPtr ptr);
char* QGeoAddress_Country(QtObjectPtr ptr);
char* QGeoAddress_CountryCode(QtObjectPtr ptr);
char* QGeoAddress_County(QtObjectPtr ptr);
char* QGeoAddress_District(QtObjectPtr ptr);
int QGeoAddress_IsEmpty(QtObjectPtr ptr);
int QGeoAddress_IsTextGenerated(QtObjectPtr ptr);
char* QGeoAddress_PostalCode(QtObjectPtr ptr);
void QGeoAddress_SetCity(QtObjectPtr ptr, char* city);
void QGeoAddress_SetCountry(QtObjectPtr ptr, char* country);
void QGeoAddress_SetCountryCode(QtObjectPtr ptr, char* countryCode);
void QGeoAddress_SetCounty(QtObjectPtr ptr, char* county);
void QGeoAddress_SetDistrict(QtObjectPtr ptr, char* district);
void QGeoAddress_SetPostalCode(QtObjectPtr ptr, char* postalCode);
void QGeoAddress_SetState(QtObjectPtr ptr, char* state);
void QGeoAddress_SetStreet(QtObjectPtr ptr, char* street);
void QGeoAddress_SetText(QtObjectPtr ptr, char* text);
char* QGeoAddress_Street(QtObjectPtr ptr);
char* QGeoAddress_Text(QtObjectPtr ptr);
void QGeoAddress_DestroyQGeoAddress(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif