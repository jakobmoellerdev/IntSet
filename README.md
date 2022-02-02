# IntSet - Integer based Set based on a bit-vector
Every integer that is stored will be converted to a bit in a word in which its located. The words are sized after
the most optimal size on the given platform. (64-bit or 32-bit)

Since this is a set, we thus reduce existence of the number to the bit being set at the correct place in the bitfield at the given index.

Retrieving a number through the bitmask can be achieved by using the size of words multiplied by the word and adding the current bit as a number.