package signals

import "github.com/rahul0tripathi/fastws/schema/clientCmd"

func ParseClientSignal(data []byte) clientCmd.ClientCmd {
	clientEvent := clientCmd.GetRootAsClientEvent(data, 0)
	switch clientEvent.Event() {
	//case clientCmd.SOMEEVENT:
	//	return clientCmd.SOMEEVENT
	default:
		return clientCmd.ClientCmdDEFAULT
	}
}
