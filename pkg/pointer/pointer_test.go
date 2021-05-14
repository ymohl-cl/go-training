package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

Si P1 := &int{1021} alors

- P1 est stocke sur 8 octets a l'adresse [00 a 07] => ( ex: [0x0000]->[0x0007] p1 addrs)
- Sa valeur est une adresse (puisque c'est un pointer) vers une zone memoire qui represente un int, donc a l'addresse [08 a 16] => ( ex: [0x0008]->[0x0016] adresse du int )
- Donc a l'adresse [0x0008] => [0x0016] on peut trouver l'adresse du int, par example [0x4200] -> [0x4208]

- Pour finir, a l'adresse suivante [0x4200]->[0x4208] on peut lire 1021 <- la valeur du int

Syntaxtiquement:
	- Si j'utlise P1 j'obtiens sa valeur, soit l'adresse du int ([0x0008])
	- Si j'utilise *P1 j'obtiens la valeur vers laquelle mon pointer pointe, soit 1021
	- Si j'utilise &P1 j'obtiens l'adresse du pointer, soit ([0x0000])

*/

func TestPointer(t *testing.T) {
	t.Run("Should be ok, p1 and p2 are nil, because just declared", func(t *testing.T) {
		var p1 *int
		var p2 *int

		assert.Nil(t, p1)
		assert.Nil(t, p2)
	})

	t.Run("Ex00: Should repair the CopyValue method to enter in the ship", func(t *testing.T) {
		p1 := new(int)
		*p1 = 42

		// uncomment this assert, should be pass without change anything
		// assert.EqualValues(t, 42, *CopyValue(p1))
	})

	t.Run("Ex01: you should unmute the ship's bot with uncomment and fix the asserts", func(t *testing.T) {
		// Note with go, the memory is already clean !
		var p1 *int
		var p2 *int

		p1 = new(int)
		p2 = new(int)

		// assert.Nil(t, p1) || assert.NotNil(t, p1)
		// assert.Nil(t, p2) || assert.NotNil(t, p2)
		// assert.EqualValues(t, "replace me by the great value", *p1)
		// assert.EqualValues(t, "replace me by the geat value", *p2)

		assert.Equal(t, p1, p1) // useless, just to compilation ok
		assert.Equal(t, p2, p2) // useless, just to compilation ok
	})

	t.Run("Ex02: bad values !!! aie aie aie, help me to repair the ship's bot, don't let me !!", func(t *testing.T) {
		// Note with go, the memory is already clean !
		// Remenber the syntax :)
		p1 := new(int) // *int
		p2 := new(int) // *int
		p3 := new(int) // *int
		p4 := new(int) // *int
		p5 := new(int) // *int
		p6 := new(int) // *int

		*p1 = 42
		*p2 = 38
		*p3 = 67
		*p4 = *p3
		*p5 = *p1
		*p6 = *p2

		/* easy just to initialize the bot */
		// assert.EqualValues(t, 42, p1)
		// assert.EqualValues(t, 38, p2)
		// assert.EqualValues(t, 67, p3)

		// assert.EqualValues(t, 0, *p4)
		// assert.EqualValues(t, 0, *p5)
		// assert.EqualValues(t, 0, *p6)

		/* hum, keep calm and go it */
		// assert.EqualValues(t, 0, CopyValue(p1))
		// assert.EqualValues(t, 0, CopyValue(p5))
		// assert.EqualValues(t, 0, CopyValue(p2))

		/* something went wront, speed to fix ! */
		p4 = p3
		p1 = p2
		*p6 = *p3
		*p3 = *p5
		p5 = p4

		// assert.EqualValues(t, 0, p1)
		// assert.EqualValues(t, 0, p2)
		// assert.EqualValues(t, 0, p3)
		// assert.EqualValues(t, 0, p4)
		// assert.EqualValues(t, 0, p5)
		// assert.EqualValues(t, 0, p6)

		assert.Equal(t, p1, p1) // useless, just to compilation ok
		assert.Equal(t, p2, p2) // useless, just to compilation ok
		assert.Equal(t, p3, p3) // useless, just to compilation ok
		assert.Equal(t, p4, p4) // useless, just to compilation ok
		assert.Equal(t, p5, p5) // useless, just to compilation ok
		assert.Equal(t, p6, p6) // useless, just to compilation ok
	})

	t.Run("Ex03: bad syntax !!! last step to fly to the next galaxy", func(t *testing.T) {
		// Note with go, the memory is already clean !
		p1 := new(int)
		p2 := new(int)
		p3 := new(int)
		p4 := new(int)
		p5 := new(*int)
		p6 := new(int)

		*p1 = 42
		*p2 = 38
		*p3 = 67
		p4 = p1
		p5 = &p1
		p6 = p2

		/* start the motor ... */
		//		assert.EqualValues(t, p, Me(p1))      // You can't use twice time p1
		//		assert.EqualValues(t, p, CopyRef(p2)) // You can't use twice time p2
		//		assert.EqualValues(t, ?, &p4)        // You can't change the 'actual value'
		//		assert.EqualValues(t, ?, &p2)        // You can't change the 'actual value'

		/* Ok, just put the accelerator !! Enjoy */
		//	assert.EqualValues(t, p, p5) // expected should be a pointer (p?), probably, you should fix the p5 syntax
		p5 = &p4
		//		assert.EqualValues(t, p, **p5) // expected should be a pointer (p?), probably, you should fix the p5 syntax
		//		assert.EqualValues(t, 42, CopyValue(*p5)) // The expected is valid, how get 42 just with syntax update to 'actual value' ?

		assert.Equal(t, p1, p1) // useless, just to compilation ok
		assert.Equal(t, p2, p2) // useless, just to compilation ok
		assert.Equal(t, p3, p3) // useless, just to compilation ok
		assert.Equal(t, p4, p4) // useless, just to compilation ok
		assert.Equal(t, p5, p5) // useless, just to compilation ok
		assert.Equal(t, p6, p6) // useless, just to compilation ok
	})
}
