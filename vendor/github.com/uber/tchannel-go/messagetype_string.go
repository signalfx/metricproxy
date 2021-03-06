// generated by stringer -type=messageType; DO NOT EDIT

package tchannel

import "fmt"

const (
	_messageType_name_0 = "messageTypeInitReqmessageTypeInitResmessageTypeCallReqmessageTypeCallRes"
	_messageType_name_1 = "messageTypeCallReqContinuemessageTypeCallResContinue"
	_messageType_name_2 = "messageTypePingReqmessageTypePingRes"
	_messageType_name_3 = "messageTypeError"
)

var (
	_messageType_index_0 = [...]uint8{0, 18, 36, 54, 72}
	_messageType_index_1 = [...]uint8{0, 26, 52}
	_messageType_index_2 = [...]uint8{0, 18, 36}
	_messageType_index_3 = [...]uint8{0, 16}
)

func (i messageType) String() string {
	switch {
	case 1 <= i && i <= 4:
		i -= 1
		return _messageType_name_0[_messageType_index_0[i]:_messageType_index_0[i+1]]
	case 19 <= i && i <= 20:
		i -= 19
		return _messageType_name_1[_messageType_index_1[i]:_messageType_index_1[i+1]]
	case 208 <= i && i <= 209:
		i -= 208
		return _messageType_name_2[_messageType_index_2[i]:_messageType_index_2[i+1]]
	case i == 255:
		return _messageType_name_3
	default:
		return fmt.Sprintf("messageType(%d)", i)
	}
}
