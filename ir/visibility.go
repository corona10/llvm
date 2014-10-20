package ir

// Visibility defines how the symbol of an executable or shared object may be
// accessed, but does not impose restrictions during link-editing of object
// files [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#visibility-styles
type Visibility uint8

// Visibility styles.
const (
	// VisibilityDefault denotes an exported symbol (STV_DEFAULT in the case of
	// ELF).
	VisibilityDefault = iota
	// VisibilityHidden denotes a hidden symbol (STV_HIDDEN in the case of ELF).
	VisibilityHidden
	// VisibilityProtected denotes a protected symbol (STV_PROTECTED in the case
	// of ELF), which is an exported symbol that cannot be overridden.
	VisibilityProtected
)
