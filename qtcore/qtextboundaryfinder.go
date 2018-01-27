package qtcore

// /usr/include/qt/QtCore/qtextboundaryfinder.h
// #include <qtextboundaryfinder.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QTextBoundaryFinder struct {
	*qtrt.CObject
}

func (this *QTextBoundaryFinder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBoundaryFinder) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextBoundaryFinderFromPointer(cthis unsafe.Pointer) *QTextBoundaryFinder {
	return &QTextBoundaryFinder{&qtrt.CObject{cthis}}
}
func (*QTextBoundaryFinder) NewFromPointer(cthis unsafe.Pointer) *QTextBoundaryFinder {
	return NewQTextBoundaryFinderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:54
// index:0
// Public
// void QTextBoundaryFinder()
func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:76
// index:1
// Public
// void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QString &)
func NewQTextBoundaryFinder_1(type_ int, string *QString) *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeERK7QString", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:77
// index:2
// Public
// void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QChar *, int, unsigned char *, int)
func NewQTextBoundaryFinder_2(type_ int, chars *QChar /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/, bufferSize int) *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = chars.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1, length, buffer, bufferSize)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:57
// index:0
// Public
// void ~QTextBoundaryFinder()
func DeleteQTextBoundaryFinder(*QTextBoundaryFinder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:79
// index:0
// Public inline
// bool isValid()
func (this *QTextBoundaryFinder) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:81
// index:0
// Public inline
// QTextBoundaryFinder::BoundaryType type()
func (this *QTextBoundaryFinder) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:82
// index:0
// Public
// QString string()
func (this *QTextBoundaryFinder) String() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder6stringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:84
// index:0
// Public
// void toStart()
func (this *QTextBoundaryFinder) ToStart() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder7toStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:85
// index:0
// Public
// void toEnd()
func (this *QTextBoundaryFinder) ToEnd() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder5toEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:86
// index:0
// Public
// int position()
func (this *QTextBoundaryFinder) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:87
// index:0
// Public
// void setPosition(int)
func (this *QTextBoundaryFinder) SetPosition(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder11setPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:89
// index:0
// Public
// int toNextBoundary()
func (this *QTextBoundaryFinder) ToNextBoundary() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder14toNextBoundaryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:90
// index:0
// Public
// int toPreviousBoundary()
func (this *QTextBoundaryFinder) ToPreviousBoundary() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder18toPreviousBoundaryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:92
// index:0
// Public
// bool isAtBoundary()
func (this *QTextBoundaryFinder) IsAtBoundary() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder12isAtBoundaryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:93
// index:0
// Public
// QTextBoundaryFinder::BoundaryReasons boundaryReasons()
func (this *QTextBoundaryFinder) BoundaryReasons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder15boundaryReasonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QTextBoundaryFinder__BoundaryType = int

const QTextBoundaryFinder__Grapheme QTextBoundaryFinder__BoundaryType = 0
const QTextBoundaryFinder__Word QTextBoundaryFinder__BoundaryType = 1
const QTextBoundaryFinder__Sentence QTextBoundaryFinder__BoundaryType = 2
const QTextBoundaryFinder__Line QTextBoundaryFinder__BoundaryType = 3

type QTextBoundaryFinder__BoundaryReason = int

const QTextBoundaryFinder__NotAtBoundary QTextBoundaryFinder__BoundaryReason = 0
const QTextBoundaryFinder__BreakOpportunity QTextBoundaryFinder__BoundaryReason = 31
const QTextBoundaryFinder__StartOfItem QTextBoundaryFinder__BoundaryReason = 32
const QTextBoundaryFinder__EndOfItem QTextBoundaryFinder__BoundaryReason = 64
const QTextBoundaryFinder__MandatoryBreak QTextBoundaryFinder__BoundaryReason = 128
const QTextBoundaryFinder__SoftHyphen QTextBoundaryFinder__BoundaryReason = 256

//  body block end
