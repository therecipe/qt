#ifdef __cplusplus
extern "C" {
#endif

int QLCDNumber_IntValue(void* ptr);
int QLCDNumber_Mode(void* ptr);
int QLCDNumber_SegmentStyle(void* ptr);
void QLCDNumber_SetMode(void* ptr, int v);
void QLCDNumber_SetSegmentStyle(void* ptr, int v);
void QLCDNumber_SetSmallDecimalPoint(void* ptr, int v);
int QLCDNumber_SmallDecimalPoint(void* ptr);
void* QLCDNumber_NewQLCDNumber(void* parent);
int QLCDNumber_CheckOverflow2(void* ptr, int num);
int QLCDNumber_DigitCount(void* ptr);
void QLCDNumber_Display(void* ptr, char* s);
void QLCDNumber_Display3(void* ptr, int num);
void QLCDNumber_ConnectOverflow(void* ptr);
void QLCDNumber_DisconnectOverflow(void* ptr);
void QLCDNumber_SetBinMode(void* ptr);
void QLCDNumber_SetDecMode(void* ptr);
void QLCDNumber_SetDigitCount(void* ptr, int numDigits);
void QLCDNumber_SetHexMode(void* ptr);
void QLCDNumber_SetOctMode(void* ptr);
void QLCDNumber_DestroyQLCDNumber(void* ptr);

#ifdef __cplusplus
}
#endif