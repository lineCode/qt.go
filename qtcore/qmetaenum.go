package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMetaEnum struct {
	*qtrt.CObject
}

func (this *QMetaEnum) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaEnum) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMetaEnumFromPointer(cthis unsafe.Pointer) *QMetaEnum {
	return &QMetaEnum{&qtrt.CObject{cthis}}
}
func (*QMetaEnum) NewFromPointer(cthis unsafe.Pointer) *QMetaEnum {
	return NewQMetaEnumFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:209
// index:0
// Public inline
// void QMetaEnum()
func NewQMetaEnum() *QMetaEnum {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMetaEnumC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMetaEnumFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:211
// index:0
// Public
// const char * name()
func (this *QMetaEnum) Name() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:212
// index:0
// Public
// bool isFlag()
func (this *QMetaEnum) IsFlag() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum6isFlagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:213
// index:0
// Public
// bool isScoped()
func (this *QMetaEnum) IsScoped() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum8isScopedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:215
// index:0
// Public
// int keyCount()
func (this *QMetaEnum) KeyCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum8keyCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:216
// index:0
// Public
// const char * key(int)
func (this *QMetaEnum) Key(index int) string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum3keyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:217
// index:0
// Public
// int value(int)
func (this *QMetaEnum) Value(index int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum5valueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:219
// index:0
// Public
// const char * scope()
func (this *QMetaEnum) Scope() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum5scopeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:221
// index:0
// Public
// int keyToValue(const char *, _Bool *)
func (this *QMetaEnum) KeyToValue(key string, ok unsafe.Pointer /*666*/) int {
	var convArg0 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum10keyToValueEPKcPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:222
// index:0
// Public
// const char * valueToKey(int)
func (this *QMetaEnum) ValueToKey(value int) string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum10valueToKeyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:223
// index:0
// Public
// int keysToValue(const char *, _Bool *)
func (this *QMetaEnum) KeysToValue(keys string, ok unsafe.Pointer /*666*/) int {
	var convArg0 = qtrt.CString(keys)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum11keysToValueEPKcPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:224
// index:0
// Public
// QByteArray valueToKeys(int)
func (this *QMetaEnum) ValueToKeys(value int) *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum11valueToKeysEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:226
// index:0
// Public inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaEnum) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum19enclosingMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:228
// index:0
// Public inline
// bool isValid()
func (this *QMetaEnum) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end