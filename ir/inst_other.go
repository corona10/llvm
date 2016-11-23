package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ icmp ] ----------------------------------------------------------------

// InstICmp represents an icmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type InstICmp struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Integer condition code.
	cond IntPred
	// Operands.
	x, y value.Value
	// Type of the instruction.
	typ types.Type
}

// NewICmp returns a new icmp instruction based on the given integer condition
// code and operands.
func NewICmp(cond IntPred, x, y value.Value) *InstICmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len())
	}
	return &InstICmp{cond: cond, x: x, y: y, typ: typ}
}

// Type returns the type of the instruction.
func (i *InstICmp) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *InstICmp) Ident() string {
	return enc.Local(i.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstICmp) SetIdent(ident string) {
	i.ident = ident
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstICmp) LLVMString() string {
	x, y := i.X(), i.Y()
	return fmt.Sprintf("%s = icmp %s %s %s, %s %s",
		i.Ident(),
		i.Cond().LLVMString(),
		x.Type().LLVMString(),
		x.Ident(),
		y.Type().LLVMString(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstICmp) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstICmp) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Cond returns the integer condition code of the icmp instruction.
func (i *InstICmp) Cond() IntPred {
	return i.cond
}

// X returns the x operand of the icmp instruction.
func (i *InstICmp) X() value.Value {
	return i.x
}

// Y returns the y operand of the icmp instruction.
func (i *InstICmp) Y() value.Value {
	return i.y
}

// IntPred represents the set of condition codes of the icmp instruction.
type IntPred int

// Integer condition codes.
const (
	IntEq  IntPred = iota + 1 // eq: equal
	IntNe                     // ne: not equal
	IntUgt                    // ugt: unsigned greater than
	IntUge                    // uge: unsigned greater than or equal
	IntUlt                    // ult: unsigned less than
	IntUle                    // ule: unsigned less than or equal
	IntSgt                    // sgt: signed greater than
	IntSge                    // sge: signed greater than or equal
	IntSlt                    // slt: signed less than
	IntSle                    // sle: signed less than or equal
)

// LLVMString returns the LLVM syntax representation of the integer condition
// code.
func (cond IntPred) LLVMString() string {
	m := map[IntPred]string{
		IntEq:  "eq",
		IntNe:  "ne",
		IntUgt: "ugt",
		IntUge: "uge",
		IntUlt: "ult",
		IntUle: "ule",
		IntSgt: "sgt",
		IntSge: "sge",
		IntSlt: "slt",
		IntSle: "sle",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("<unknown integer condition code %d>", int(cond))
}

// --- [ fcmp ] ----------------------------------------------------------------

// InstFCmp represents an fcmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type InstFCmp struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Floating-point condition code.
	cond FloatPred
	// Operands.
	x, y value.Value
	// Type of the instruction.
	typ types.Type
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// condition code and operands.
func NewFCmp(cond FloatPred, x, y value.Value) *InstFCmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len())
	}
	return &InstFCmp{cond: cond, x: x, y: y, typ: typ}
}

// Type returns the type of the instruction.
func (i *InstFCmp) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *InstFCmp) Ident() string {
	return enc.Local(i.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstFCmp) SetIdent(ident string) {
	i.ident = ident
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstFCmp) LLVMString() string {
	x, y := i.X(), i.Y()
	return fmt.Sprintf("%s = fcmp %s %s %s, %s %s",
		i.Ident(),
		i.Cond().LLVMString(),
		x.Type().LLVMString(),
		x.Ident(),
		y.Type().LLVMString(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstFCmp) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstFCmp) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Cond returns the floating-point condition code of the fcmp instruction.
func (i *InstFCmp) Cond() FloatPred {
	return i.cond
}

// X returns the x operand of the fcmp instruction.
func (i *InstFCmp) X() value.Value {
	return i.x
}

// Y returns the y operand of the fcmp instruction.
func (i *InstFCmp) Y() value.Value {
	return i.y
}

// FloatPred represents the set of condition codes of the fcmp instruction.
type FloatPred int

// Floating-point condition codes.
const (
	FloatFalse FloatPred = iota + 1 // false: no comparison, always returns false
	FloatOeq                        // oeq: ordered and equal
	FloatOgt                        // ogt: ordered and greater than
	FloatOge                        // oge: ordered and greater than or equal
	FloatOlt                        // olt: ordered and less than
	FloatOle                        // ole: ordered and less than or equal
	FloatOne                        // one: ordered and not equal
	FloatOrd                        // ord: ordered (no nans)
	FloatUeq                        // ueq: unordered or equal
	FloatUgt                        // ugt: unordered or greater than
	FloatUge                        // uge: unordered or greater than or equal
	FloatUlt                        // ult: unordered or less than
	FloatUle                        // ule: unordered or less than or equal
	FloatUne                        // une: unordered or not equal
	FloatUno                        // uno: unordered (either nans)
	FloatTrue                       // true: no comparison, always returns true
)

// LLVMString returns the LLVM syntax representation of the floating-point
// condition code.
func (cond FloatPred) LLVMString() string {
	m := map[FloatPred]string{
		FloatFalse: "false",
		FloatOeq:   "oeq",
		FloatOgt:   "ogt",
		FloatOge:   "oge",
		FloatOlt:   "olt",
		FloatOle:   "ole",
		FloatOne:   "one",
		FloatOrd:   "ord",
		FloatUeq:   "ueq",
		FloatUgt:   "ugt",
		FloatUge:   "uge",
		FloatUlt:   "ult",
		FloatUle:   "ule",
		FloatUne:   "une",
		FloatUno:   "uno",
		FloatTrue:  "true",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("<unknown floating-point condition code %d>", int(cond))
}

// --- [ phi ] -----------------------------------------------------------------

// InstPHI represents a phi instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#phi-instruction
type InstPHI struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Incoming values.
	incs []*Incoming
	// Type of the instruction.
	typ types.Type
}

// NewPHI returns a new phi instruction based on the given incoming values.
func NewPHI(incs ...*Incoming) *InstPHI {
	if len(incs) < 1 {
		panic(fmt.Sprintf("invalid number of incoming values; expected > 0, got %d", len(incs)))
	}
	typ := incs[0].x.Type()
	return &InstPHI{incs: incs, typ: typ}
}

// Type returns the type of the instruction.
func (i *InstPHI) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *InstPHI) Ident() string {
	return enc.Local(i.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstPHI) SetIdent(ident string) {
	i.ident = ident
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstPHI) LLVMString() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s = phi %s ",
		i.Ident(),
		i.Type().LLVMString())
	for j, inc := range i.Incs() {
		if j != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "[ %s, %s ]",
			inc.X().Ident(),
			inc.Pred().Ident())
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (i *InstPHI) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstPHI) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Incs returns the incoming values of the phi instruction.
func (i *InstPHI) Incs() []*Incoming {
	return i.incs
}

// Incoming represents an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	x value.Value
	// Predecessor basic block of the incoming value.
	pred *BasicBlock
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x value.Value, pred *BasicBlock) *Incoming {
	return &Incoming{x: x, pred: pred}
}

// X returns the incoming value.
func (inc *Incoming) X() value.Value {
	return inc.x
}

// Pred returns the predecessor basic block of the incoming value.
func (inc *Incoming) Pred() *BasicBlock {
	return inc.pred
}

// --- [ select ] --------------------------------------------------------------

// InstSelect represents a select instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type InstSelect struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Selection condition.
	cond value.Value
	// Operands.
	x, y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *InstSelect {
	return &InstSelect{cond: cond, x: x, y: y}
}

// Type returns the type of the instruction.
func (i *InstSelect) Type() types.Type {
	return i.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (i *InstSelect) Ident() string {
	return enc.Local(i.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstSelect) SetIdent(ident string) {
	i.ident = ident
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstSelect) LLVMString() string {
	cond, x, y := i.Cond(), i.X(), i.Y()
	return fmt.Sprintf("%s = select %s %s, %s %s, %s %s",
		i.Ident(),
		cond.Type().LLVMString(),
		cond.Ident(),
		x.Type().LLVMString(),
		x.Ident(),
		y.Type().LLVMString(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstSelect) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstSelect) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Cond returns the selection condition of the select instruction.
func (i *InstSelect) Cond() value.Value {
	return i.cond
}

// X returns the x operand of the select instruction.
func (i *InstSelect) X() value.Value {
	return i.x
}

// Y returns the y operand of the select instruction.
func (i *InstSelect) Y() value.Value {
	return i.y
}

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a call instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type InstCall struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Callee.
	callee *Function
	// Function arguments.
	args []value.Value
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func NewCall(callee *Function, args ...value.Value) *InstCall {
	return &InstCall{callee: callee, args: args}
}

// Type returns the type of the instruction.
func (i *InstCall) Type() types.Type {
	return i.callee.typ.RetType()
}

// Ident returns the identifier associated with the instruction.
func (i *InstCall) Ident() string {
	return enc.Local(i.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstCall) SetIdent(ident string) {
	i.ident = ident
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstCall) LLVMString() string {
	buf := &bytes.Buffer{}
	typ := i.Type()
	if !typ.Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", i.Ident())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ.LLVMString(),
		i.Callee().Ident())
	for i, arg := range i.Args() {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			arg.Type().LLVMString(),
			arg.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (i *InstCall) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstCall) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Callee returns the callee of the call instruction.
func (i *InstCall) Callee() *Function {
	return i.callee
}

// Args returns the function arguments of the call instruction.
func (i *InstCall) Args() []value.Value {
	return i.args
}

// --- [ va_arg ] --------------------------------------------------------------

// --- [ landingpad ] ----------------------------------------------------------

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------
