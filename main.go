package main

func main() {
	var adu Adu
	adu.AddField(Header{len: 1, value: []byte{0x68}})
	adu.AddField(DataLen{len: 1})
	adu.AddField(Data{})
	adu.AddField(Tail{len: 1, value: []byte{0x86}})
	adu.Info()
}
