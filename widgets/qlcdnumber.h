#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLCDNumber_IntValue(QtObjectPtr ptr);
int QLCDNumber_Mode(QtObjectPtr ptr);
int QLCDNumber_SegmentStyle(QtObjectPtr ptr);
void QLCDNumber_SetMode(QtObjectPtr ptr, int v);
void QLCDNumber_SetSegmentStyle(QtObjectPtr ptr, int v);
void QLCDNumber_SetSmallDecimalPoint(QtObjectPtr ptr, int v);
int QLCDNumber_SmallDecimalPoint(QtObjectPtr ptr);
QtObjectPtr QLCDNumber_NewQLCDNumber(QtObjectPtr parent);
int QLCDNumber_CheckOverflow2(QtObjectPtr ptr, int num);
int QLCDNumber_DigitCount(QtObjectPtr ptr);
void QLCDNumber_Display(QtObjectPtr ptr, char* s);
void QLCDNumber_Display3(QtObjectPtr ptr, int num);
void QLCDNumber_ConnectOverflow(QtObjectPtr ptr);
void QLCDNumber_DisconnectOverflow(QtObjectPtr ptr);
void QLCDNumber_SetBinMode(QtObjectPtr ptr);
void QLCDNumber_SetDecMode(QtObjectPtr ptr);
void QLCDNumber_SetDigitCount(QtObjectPtr ptr, int numDigits);
void QLCDNumber_SetHexMode(QtObjectPtr ptr);
void QLCDNumber_SetOctMode(QtObjectPtr ptr);
void QLCDNumber_DestroyQLCDNumber(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif