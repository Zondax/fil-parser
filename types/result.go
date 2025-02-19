package types

//// ResultType represents the possible return types from FilecoinParser operations
//type ResultType interface {
//	*TxsParsedResult | *EventsParsedResult | *MultisigEvents
//}
//
//type Result[T ResultType] struct {
//	Result           T                 // Result contains the parsed data of type T
//	AggregatedErrors *AggregatedErrors // AggregatedErrors contains non-fatal errors collected during parsing
//}
//
//type AggregatedErrors struct {
//	Block  string
//	Errors []*ParsingError
//}
//
//func (ae *AggregatedErrors) AddParsingError(e *ParsingError) {
//	ae.Errors = append(ae.Errors, e)
//}
//
//type ParsingError struct {
//	Func string
//	Data map[string]string
//	Err  error
//}
//
//type MetadataError ParsingError
