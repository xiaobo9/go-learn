package demo

import "log"

type (
	Point struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar Point                  // polar and Point denote different types
)

type TreeNode struct {
	left, right *TreeNode
	value       any
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct { /* Mutex fields */
	Mt int
}

func (m *Mutex) Lock()   { /* Lock implementation */ }
func (m *Mutex) Unlock() { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block

func test() {
	a := NewMutex{}
	log.Println(a)

	b := new(PtrMutex)
	log.Println(b)
}
