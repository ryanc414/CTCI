package objects

import (
	"container/list"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	math_rand "math/rand"
)

// The CircularArray type gives an array-like interface but, unlike a regular
// array, it can be efficiently rotated. Can be iterated like:

// for i := circ.First(); i != EndOfArray; i = circ.Next(i) {
//   // do something with circ.Data[i]...
// }
type CircularArray struct {
	Data   []interface{}
	offset int
}

const EndOfArray = -1

func InitCircularArray(data []interface{}) CircularArray {
	return CircularArray{Data: data, offset: 0}
}

// Return the index of the first element in the array
func (circ CircularArray) First() int {
	if len(circ.Data) == 0 {
		return EndOfArray
	} else {
		return circ.offset
	}
}

// Return the next index during iteration, given the current one.
func (circ CircularArray) Next(curr int) int {
	// bounds checking
	if curr < 0 || curr >= len(circ.Data) {
		panic("Index out of range")
	}

	next := (curr + 1) % len(circ.Data)
	if next == circ.offset {
		return EndOfArray
	} else {
		return next
	}
}

// Return a new circular array rotated by the amount specified.
func (circ CircularArray) Rotate(numPlaces int) CircularArray {
	return CircularArray{
		Data:   circ.Data,
		offset: (circ.offset + numPlaces) % len(circ.Data),
	}
}

// Seed the RNG so that different results are produced each time.
func SeedRng() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot send RNG")
	}

	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// Represents a direction of movement on a 2D grid.
type GridDirection int

const (
	Up = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

var GridDirections = [...]GridDirection{
	Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft,
}

type GridCoords struct {
	Row int
	Col int
}

// Return the next consecutive position in the given direction. Note: bounds
// checking is not performed here, so it is up to the caller to check if the
// position is within the grid or not.
func (position GridCoords) MoveDirection(direction GridDirection) GridCoords {
	switch direction {
	case Up:
		return GridCoords{Row: position.Row - 1, Col: position.Col}

	case UpRight:
		return GridCoords{Row: position.Row - 1, Col: position.Col + 1}

	case Right:
		return GridCoords{Row: position.Row, Col: position.Col + 1}

	case DownRight:
		return GridCoords{Row: position.Row + 1, Col: position.Col + 1}

	case Down:
		return GridCoords{Row: position.Row + 1, Col: position.Col}

	case DownLeft:
		return GridCoords{Row: position.Row + 1, Col: position.Col - 1}

	case Left:
		return GridCoords{Row: position.Row, Col: position.Col - 1}

	case UpLeft:
		return GridCoords{Row: position.Row - 1, Col: position.Col - 1}

	default:
		panic("Unexpected direction.")
	}
}

// Filesystem nodes and methods. This is only a very basic gist, a real
// filesystem would provide methods for copying and moving files, some helpers
// for iterating the filesystem, and support for symbolic links.

// Represent a directory in an in-memory file system.
type FileTreeDir struct {
	name       string
	parentDir  *FileTreeDir
	childDirs  map[string]*FileTreeDir
	childFiles map[string]*FileTreeFile
}

// Represent a file in an in-memory file system.
type FileTreeFile struct {
	name      string
	parentDir *FileTreeDir
	data      []byte
}

// Create the root directory.
func CreateRootDir(name string) *FileTreeDir {
	return &FileTreeDir{
		name:       name,
		childDirs:  make(map[string]*FileTreeDir),
		childFiles: make(map[string]*FileTreeFile),
	}
}

// Create a new file under an existing directory.
func (dir *FileTreeDir) CreateFile(
	name string, data []byte,
) (*FileTreeFile, error) {
	if dir.childFiles[name] != nil {
		return nil, errors.New("Filename already exists")
	}

	newFile := &FileTreeFile{
		name: name,
		data: data,
	}

	dir.childFiles[name] = newFile

	return newFile, nil
}

// Create a new dir under an existing directory.
func (dir *FileTreeDir) CreateDir(name string) (*FileTreeDir, error) {
	_, ok := dir.childDirs[name]
	if ok {
		return nil, errors.New("Dir already exists")
	}

	newDir := CreateRootDir(name)
	dir.childDirs[name] = newDir
	return newDir, nil
}

// List directory contents: the names of all directories and files.
func (dir *FileTreeDir) List() []string {
	numDirs := len(dir.childDirs)
	numFiles := len(dir.childFiles)
	names := make([]string, 0, numDirs+numFiles)

	for i := range dir.childDirs {
		names = append(names, dir.childDirs[i].name)
	}

	for i := range dir.childFiles {
		names = append(names, dir.childFiles[i].name)
	}

	return names
}

// Delete a file from a dir.
func (dir *FileTreeDir) DeleteFile(name string) error {
	_, ok := dir.childFiles[name]
	if !ok {
		return errors.New("No such file")
	}

	delete(dir.childFiles, name)
	return nil
}

// Delete a directory from a dir.
func (dir *FileTreeDir) DeleteDir(name string) error {
	childDir, ok := dir.childDirs[name]
	if !ok {
		return errors.New("No such dir")
	}

	for k := range childDir.childFiles {
		err := childDir.DeleteFile(k)
		if err != nil {
			return err
		}
	}

	for k := range childDir.childDirs {
		err := childDir.DeleteDir(k)
		if err != nil {
			return err
		}
	}

	delete(dir.childFiles, name)
	return nil
}

// Implement hash table with linked lists for collisions.

// This table will be super simple and store integers only. Integers will be
// "hashed" into n buckets by simple modulus. Within each bucket there will
// be a linked list of values.
type HashTable []*list.List

type HashTableNode struct {
	key   int
	value interface{}
}

func InitHashTable(numBuckets int) HashTable {
	hashTable := make([]*list.List, numBuckets)
	for i := range hashTable {
		hashTable[i] = list.New()
	}

	return hashTable
}

// Insert or update a key-value pair into the hash table.
func (table HashTable) SetItem(key int, value interface{}) {
	hash := table.hash(key)
	bucket := table[hash]

	// First, delete any existing item in the hash bucket.
	for el := bucket.Front(); el != nil; el = el.Next() {
		node := el.Value.(HashTableNode)
		if node.key == key {
			bucket.Remove(el)
			break
		}
	}

	newNode := HashTableNode{key: key, value: value}
	table[hash].PushFront(newNode)
}

// Get a value from the hash table given its key.
func (table HashTable) GetItem(key int) (interface{}, error) {
	hash := table.hash(key)
	bucket := table[hash]

	for el := bucket.Front(); el != nil; el = el.Next() {
		node := el.Value.(HashTableNode)
		if node.key == key {
			return node.value, nil
		}
	}

	return nil, errors.New("Not found")
}

// Delete a key from the hash table.
func (table HashTable) Delete(key int) error {
	hash := table.hash(key)
	bucket := table[hash]

	for el := bucket.Front(); el != nil; el = el.Next() {
		node := el.Value.(HashTableNode)
		if node.key == key {
			bucket.Remove(el)
			return nil
		}
	}

	return errors.New("Not found")
}

// Hash a key, returning an inteter n where 0 < n <= size of hash table.
func (table HashTable) hash(key int) int {
	return key % len(table)
}
