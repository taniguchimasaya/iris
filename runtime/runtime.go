// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, You can
// obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import (
	"math"
	"os"
	"time"

	"github.com/islisp-dev/iris/runtime/ilos"
)

var Time time.Time

func TopLevelHander(e ilos.Environment, c ilos.Instance) (ilos.Instance, ilos.Instance) {
	return nil, c
}

var TopLevel = ilos.NewEnvironment(
	ilos.NewStream(os.Stdin, nil, ilos.CharacterClass),
	ilos.NewStream(nil, os.Stdout, ilos.CharacterClass),
	ilos.NewStream(nil, os.Stderr, ilos.CharacterClass),
	ilos.NewFunction(ilos.NewSymbol("TOP-LEVEL-HANDLER"), TopLevelHander),
)

func defclass(name string, class ilos.Class) {
	symbol := ilos.NewSymbol(name)
	TopLevel.Class.Define(symbol, class)
}

func defspecial(name string, function interface{}) {
	symbol := ilos.NewSymbol(name)
	TopLevel.Special.Define(symbol, ilos.NewFunction(func2symbol(function), function))
}

func defun(name string, function interface{}) {
	symbol := ilos.NewSymbol(name)
	TopLevel.Function.Define(symbol, ilos.NewFunction(symbol, function))
}

func defgeneric(name string, function interface{}) {
	symbol := ilos.NewSymbol(name)
	lambdaList, _ := List(TopLevel, ilos.NewSymbol("FIRST"), ilos.NewSymbol("&REST"), ilos.NewSymbol("REST"))
	generic := ilos.NewGenericFunction(symbol, lambdaList, T, ilos.GenericFunctionClass)
	generic.(*ilos.GenericFunction).AddMethod(nil, lambdaList, []ilos.Class{ilos.StandardClassClass}, ilos.NewFunction(symbol, function))
	TopLevel.Function.Define(symbol, generic)
}

func defglobal(name string, value ilos.Instance) {
	symbol := ilos.NewSymbol(name)
	TopLevel.Variable.Define(symbol, value)
}

func init() {
	defglobal("*PI*", ilos.Float(math.Pi))
	defglobal("*MOST-POSITIVE-FLOAT*", MostPositiveFloat)
	defglobal("*MOST-NEGATIVE-FLOAT*", MostNegativeFloat)
	defun("-", Substruct)
	defun("+", Add)
	defun("*", Multiply)
	defun("<", NumberLessThan)
	defun("<=", NumberLessThanOrEqual)
	defun("=", NumberEqual)
	defun(">", NumberGreaterThan)
	defun(">=", NumberGreaterThanOrEqual)
	defspecial("QUASIQUOTE", Quasiquote)
	defun("ABS", Abs)
	defspecial("AND", And)
	defun("APPEND", Append)
	defun("APPLY", Apply)
	defun("ARRAY-DIMENSIONS", ArrayDimensions)
	defun("AREF", Aref)
	defun("ASSOC", Assoc)
	// TODO: defspecial2("ASSURE", Assure)
	defun("ATAN", Atan)
	defun("ATAN2", Atan2)
	defun("ATANH", Atanh)
	defun("BASIC-ARRAY*-P", BasicArrayStarP)
	defun("BASIC-ARRAY-P", BasicArrayP)
	defun("BASIC-VECTOR-P", BasicVectorP)
	defspecial("BLOCK", Block)
	defun("CAR", Car)
	defspecial("CASE", Case)
	defspecial("CASE-USING", CaseUsing)
	defspecial("CATCH", Catch)
	defun("CDR", Cdr)
	defun("CEILING", Ceiling)
	defun("CERROR", Cerror)
	defun("CHAR-INDEX", CharIndex)
	defun("CHAR/=", CharNotEqual)
	defun("CHAR<", CharLessThan)
	defun("CHAR<=", CharLessThanOrEqual)
	defun("CHAR=", CharEqual)
	defun("CHAR>", CharGreaterThan)
	defun("CHAR>=", CharGreaterThanOrEqual)
	defun("CHARACTERP", Characterp)
	defspecial("CLASS", Class)
	defun("CLASS-OF", ClassOf)
	defun("CLOSE", Close)
	// SKIP defun2("COERCION", Coercion)
	defspecial("COND", Cond)
	defun("CONDITION-CONTINUABLE", ConditionContinuable)
	defun("CONS", Cons)
	defun("CONSP", Consp)
	defun("CONTINUE-CONDITION", ContinueCondition)
	defspecial("CONVERT", Convert)
	defun("COS", Cos)
	defun("COSH", Cosh)
	defgeneric("CREATE", Create) //TODO Change to generic function
	defun("CREATE-ARRAY", CreateArray)
	defun("CREATE-LIST", CreateList)
	defun("CREATE-STRING", CreateString)
	defun("CREATE-STRING-INPUT-STREAM", CreateStringInputStream)
	defun("CREATE-STRING-OUTPUT-STREAM", CreateStringOutputStream)
	defun("CREATE-VECTOR", CreateVector)
	defspecial("DEFCLASS", Defclass)
	defspecial("DEFCONSTANT", Defconstant)
	defspecial("DEFDYNAMIC", Defdynamic)
	defspecial("DEFGENERIC", Defgeneric)
	defspecial("DEFMETHOD", Defmethod)
	defspecial("DEFGLOBAL", Defglobal)
	defspecial("DEFMACRO", Defmacro)
	defspecial("DEFUN", Defun)
	defun("DIV", Div)
	defspecial("DYNAMIC", Dynamic)
	defspecial("DYNAMIC-LET", DynamicLet)
	defun("ELT", Elt)
	defun("EQ", Eq)
	defun("EQL", Eql)
	defun("EQUAL", Equal)
	defun("ERROR", Error)
	defun("ERROR-OUTPUT", ErrorOutput)
	defun("EXP", Exp)
	defun("EXPT", Expt)
	// TODO defun2("FILE-LENGTH", FileLength)
	// TODO defun2("FILE-POSITION", FilePosition)
	defun("FINISH-OUTPUT", FinishOutput)
	defspecial("FLET", Flet)
	defun("FLOAT", Float)
	defun("FLOATP", Floatp)
	defun("FLOOR", Floor)
	defspecial("FOR", For)
	defun("FORMAT", Format)
	defun("FORMAT-CHAR", FormatChar)
	defun("FORMAT-FLOAT", FormatFloat)
	defun("FORMAT-FRESH-LINE", FormatFreshLine)
	defun("FORMAT-INTEGER", FormatInteger)
	defun("FORMAT-OBJECT", FormatObject)
	defun("FORMAT-TAB", FormatTab)
	defun("FUNCALL", Funcall)
	defspecial("FUNCTION", Function)
	defun("FUNCTIONP", Functionp)
	defun("GAREF", Garef)
	defun("GCD", Gcd)
	defun("GENERAL-ARRAY*-P", GeneralArrayStarP)
	defun("GENERAL-VECTOR-P", GeneralVectorP)
	defun("GENERIC-FUNCTION-P", GenericFunctionP)
	defun("GENSYM", Gensym)
	defun("GET-INTERNAL-REAL-TIME", GetInternalRealTime)
	defun("GET-INTERNAL-RUN-TIME", GetInternalRunTime)
	defun("GET-OUTPUT-STREAM-STRING", GetOutputStreamString)
	defun("GET-UNIVERSAL-TIME", GetUniversalTime)
	defspecial("GO", Go)
	defun("IDENTITY", Identity)
	defspecial("IF", If)
	defspecial("IGNORE-ERRORS", IgnoreErrors)
	defgeneric("INITIALIZE-OBJECT", InitializeObject) // TODO change generic function
	defun("INPUT-STREAM-P", InputStreamP)
	defun("INSTANCEP", Instancep)
	defun("INTEGERP", Integerp)
	defun("INTERNAL-TIME-UNITS-PER-SECOND", InternalTimeUnitsPerSecond)
	defun("ISQRT", Isqrt)
	defspecial("LABELS", Labels)
	defspecial("LAMBDA", Lambda)
	defun("LCM", Lcm)
	defun("LENGTH", Length)
	defspecial("LET", Let)
	defspecial("LET*", LetStar)
	defun("LIST", List)
	defun("LISTP", Listp)
	defun("LOG", Log)
	defun("MAP-INTO", MapInto)
	defun("MAPC", Mapc)
	defun("MAPCAN", Mapcan)
	defun("MAPCAR", Mapcar)
	defun("MAPCON", Mapcon)
	defun("MAPL", Mapl)
	defun("MAPLIST", Maplist)
	defun("MAX", Max)
	defun("MEMBER", Member)
	defun("MIN", Min)
	defun("MOD", Mod)
	defglobal("NI-L", Nil)
	defun("NOT", Not)
	defun("NREVERSE", Nreverse)
	defun("NULL", Null)
	defun("NUMBERP", Numberp)
	defun("OPEN-INPUT-FILE", OpenInputFile)
	defun("OPEN-IO-FILE", OpenIoFile)
	defun("OPEN-OUTPUT-FILE", OpenOutputFile)
	defun("OPEN-STREAM-P", OpenStreamP)
	defspecial("OR", Or)
	// defun("FLUSH-OUTPUT", FlushOutput)
	defun("OUTPUT-STREAM-P", OutputStreamP)
	defun("PARSE-NUMBER", ParseNumber)
	defun("PREVIEW-CHAR", PreviewChar)
	defun("PROBE-FILE", ProbeFile)
	defspecial("PROGN", Progn)
	defun("PROPERTY", Property)
	defspecial("QUASIQUOTE", Quasiquote)
	defspecial("QUOTE", Quote)
	defun("QUOTIENT", Quotient)
	defun("READ", Read)
	defun("READ-BYTE", ReadByte)
	defun("READ-CHAR", ReadChar)
	defun("READ-LINE", ReadLine)
	defun("REMOVE-PROPERTY", RemoveProperty)
	defun("REPORT-CONDITION", ReportCondition)
	defspecial("RETURN-FROM", ReturnFrom)
	defun("REVERSE", Reverse)
	defun("ROUND", Round)
	defun("SET-AREF", SetAref)
	defun("(SETF AREF)", SetAref)
	defun("SET-CAR", SetCar)
	defun("(SETF CAR)", SetCar)
	defun("SET-CDR", SetCdr)
	defun("(SETF CDR)", SetCdr)
	defun("SET-DYNAMIC", SetDynamic)
	defun("(SETF DYNAMIC)", SetDynamic)
	defun("SET-ELT", SetElt)
	defun("(SETF ELT)", SetElt)
	// TODO defun2("SET-FILE-POSITION", SetFilePosition)
	defun("SET-GAREF", SetGaref)
	defun("(SETF GAREF)", SetGaref)
	defun("SET-PROPERTY", SetProperty)
	defun("(SETF PROPERTY)", SetProperty)
	defspecial("SETF", Setf)
	defspecial("SETQ", Setq)
	defun("SIGNAL-CONDITION", SignalCondition)
	// TODO defun2("SIMPLE-ERROR-FORMAT-ARGUMENTS", SimpleErrorFormatArguments)
	// TODO defun2("SIMPLE-ERROR-FORMAT-STRING", SimpleErrorFormatString)
	defun("SIN", Sin)
	defun("SINH", Sinh)
	defun("SQRT", Sqrt)
	defun("STANDARD-INPUT", StandardInput)
	defun("STANDARD-OUTPUT", StandardOutput)
	defun("STREAM-READY-P", StreamReadyP)
	defun("STREAMP", Streamp)
	defun("STRING-APPEND", StringAppend)
	defun("STRING-INDEX", StringIndex)
	defun("STRING/=", StringNotEqual)
	defun("STRING>", StringGreaterThan)
	defun("STRING>=", StringGreaterThanOrEqual)
	defun("STRING=", StringEqual)
	defun("STRING<", StringLessThan)
	defun("STRING<=", StringLessThanOrEqual)
	defun("STRINGP", Stringp)
	defun("SUBCLASSP", Subclassp)
	defun("SUBSEQ", Subseq)
	defun("SYMBOLP", Symbolp)
	defglobal("T", T)
	defspecial("TAGBODY", Tagbody)
	defspecial("TAN", Tan)
	defspecial("TANH", Tanh)
	// TODO defspecial2("THE", The)
	defspecial("THROW", Throw)
	defun("TRUNCATE", Truncate)
	// TODO defun1("UNDEFINED-ENTITY-NAME", UndefinedEntityName)
	// TODO defun2("UNDEFINED-ENTITY-NAMESPACE", UndefinedEntityNamespace)
	defspecial("UNWIND-PROTECT", UnwindProtect)
	defun("VECTOR", Vector)
	defspecial("WHILE", While)
	defspecial("WITH-ERROR-OUTPUT", WithErrorOutput)
	defspecial("WITH-HANDLER", WithHandler)
	defspecial("WITH-OPEN-INPUT-FILE", WithOpenInputFile)
	defspecial("WITH-OPEN-OUTPUT-FILE", WithOpenOutputFile)
	defspecial("WITH-STANDARD-INPUT", WithStandardInput)
	defspecial("WITH-STANDARD-OUTPUT", WithStandardOutput)
	defun("WRITE-BYTE", WriteByte)
	defclass("<OBJECT>", ilos.ObjectClass)
	defclass("<BUILT-IN-CLASS>", ilos.BuiltInClassClass)
	defclass("<STANDARD-CLASS>", ilos.StandardClassClass)
	defclass("<BASIC-ARRAY>", ilos.BasicArrayClass)
	defclass("<BASIC-ARRAY-STAR>", ilos.BasicArrayStarClass)
	defclass("<GENERAL-ARRAY-STAR>", ilos.GeneralArrayStarClass)
	defclass("<BASIC-VECTOR>", ilos.BasicVectorClass)
	defclass("<GENERAL-VECTOR>", ilos.GeneralVectorClass)
	defclass("<STRING>", ilos.StringClass)
	defclass("<CHARACTER>", ilos.CharacterClass)
	defclass("<FUNCTION>", ilos.FunctionClass)
	defclass("<GENERIC-FUNCTION>", ilos.GenericFunctionClass)
	defclass("<STANDARD-GENERIC-FUNCTION>", ilos.StandardGenericFunctionClass)
	defclass("<LIST>", ilos.ListClass)
	defclass("<CONS>", ilos.ConsClass)
	defclass("<NULL>", ilos.NullClass)
	defclass("<SYMBOL>", ilos.SymbolClass)
	defclass("<NUMBER>", ilos.NumberClass)
	defclass("<INTEGER>", ilos.IntegerClass)
	defclass("<FLOAT>", ilos.FloatClass)
	defclass("<SERIOUS-CONDITION>", ilos.SeriousConditionClass)
	defclass("<ERROR>", ilos.ErrorClass)
	defclass("<ARITHMETIC-ERROR>", ilos.ArithmeticErrorClass)
	defclass("<DIVISION-BY-ZERO>", ilos.DivisionByZeroClass)
	defclass("<FLOATING-POINT-ONDERFLOW>", ilos.FloatingPointOnderflowClass)
	defclass("<FLOATING-POINT-UNDERFLOW>", ilos.FloatingPointUnderflowClass)
	defclass("<CONTROL-ERROR>", ilos.ControlErrorClass)
	defclass("<PARSE-ERROR>", ilos.ParseErrorClass)
	defclass("<PROGRAM-ERROR>", ilos.ProgramErrorClass)
	defclass("<DOMAIN-ERROR>", ilos.DomainErrorClass)
	defclass("<UNDEFINED-ENTITY>", ilos.UndefinedEntityClass)
	defclass("<UNDEFINED-VARIABLE>", ilos.UndefinedVariableClass)
	defclass("<UNDEFINED-FUNCTION>", ilos.UndefinedFunctionClass)
	defclass("<SIMPLE-ERROR>", ilos.SimpleErrorClass)
	defclass("<STREAM-ERROR>", ilos.StreamErrorClass)
	defclass("<END-OF-STREAM>", ilos.EndOfStreamClass)
	defclass("<STORAGE-EXHAUSTED>", ilos.StorageExhaustedClass)
	defclass("<STANDARD-OBJECT>", ilos.StandardObjectClass)
	defclass("<STREAM>", ilos.StreamClass)
	Time = time.Now()
}
