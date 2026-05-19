package dataStructure

var (
	currentCount int
	hashMapList  [5][]entry
)

type entry struct {
	key   int
	value string
}

func newEntry(key int, value string) *entry {
	return &entry{key: key, value: value}
}

/*put (k,v)
get(k) : v
remove(k)
k: int
v:string
collision using chaining
linkedlist<entry>[]
{ll,ll,ll,ll}
*/

/*2 keys with the same function

1 way)Chaining: each cell pointed to a linked list, if there is a collision add to end of list
e.g array of 5 spaces , hash index value of 6 aand 11 leads to index 1 , store in linked list
2 way )Open addressing: find different slot for the second value
Probing ,1)Linear Probing : check each next slot for space _. hash(key) + i % table size until we find an empty slot
2) Quadratic Probing : hash(key) + (i * i) % table size , can lead to infinite loop
2) doubke hashing:*/

func Put(key int, value string) {
	hashCode := getHashCode(key)
	existingEntry := &hashMapList[hashCode]
	if existingEntry == nil {
		*existingEntry = make([]entry, 0)
	}

	for _, entry := range *existingEntry {
		if entry.key == key {
			entry.value = value
			return
		}
	}
	addedEntry := newEntry(key, value)
	*existingEntry = append(*existingEntry, *addedEntry)
	currentCount++
}

func Get(key int) string {
	hashCode := getHashCode(key)
	existingEntry := hashMapList[hashCode]

	if existingEntry != nil {
		for _, value := range existingEntry {
			if value.key == key {
				return value.value
			}
		}
	}
	return ""
}

func Remove(key int) {
	hashCode := getHashCode(key)
	existingEntry := hashMapList[hashCode]
	if existingEntry != nil {
		for index, entry := range existingEntry {
			if entry.key == key {
				existingEntry = append(existingEntry[:index], existingEntry[index+1:]...)
				hashMapList[hashCode] = existingEntry
				currentCount--
				return
			}
		}
	}
}

//sample usage
func getHashCode(input int) int {
	return input % len(hashMapList)
}

func FirstNonRepeatedCharacter(input string) string {
	inputArray := []rune(input)
	dictionary := make(map[rune]int, 0)

	for _, value := range inputArray {
		dictionary[value]++
	}

	for index, value := range dictionary {
		if value == 1 {
			return string(index)
		}
	}

	return ""
}
