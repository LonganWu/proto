package main

import (
	"fmt"
	"reflect"
)

type Adu struct {
	Fields []Field
}

func (adu *Adu) Info() {
	for _, v := range adu.Fields {
		of := reflect.TypeOf(v)
		fmt.Println("类型:", of, " 长度:", v.Length())
	}
}

func (adu *Adu) AddField(field Field) {
	adu.Fields = append(adu.Fields, field)
}

type Field interface {
	Length() uint16
}

type Header struct {
	len   byte
	value []byte
}

func (h Header) Length() uint16 {
	return uint16(h.len)
}

type DataLen struct {
	len byte
}

func (dl DataLen) Length() uint16 {
	return uint16(dl.len)
}
func (dl DataLen) SetByte() {

}

type Data struct {
	len  byte
	data []byte
}

func (d Data) Length() uint16 {
	return uint16(d.len)
}

type Tail struct {
	len   byte
	value []byte
}

func (t Tail) Length() uint16 {
	return uint16(t.len)
}
