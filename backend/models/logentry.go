package models

type LogEntry struct {
	Message        string `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8"`
	MessageRaw     string `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8"`
	StructuredData string `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tag            string `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8"`
	Sender         string `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8"`
	Groupings      string `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8"`
	Event          string `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8"`
	EventId        string `parquet:"name=EventId, type=BYTE_ARRAY, convertedtype=UTF8"`
	NanoTimeStamp  int64  `parquet:"name=NanoTimeStamp, type=INT64"`
	Namespace      string `parquet:"name=Namespace, type=BYTE_ARRAY, convertedtype=UTF8"`
}
