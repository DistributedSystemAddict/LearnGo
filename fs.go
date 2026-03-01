package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ---------------------------------------------------------------------------
// Global flags
// ---------------------------------------------------------------------------

var (
	debugMode      = false
	printOpsFlag   = true
	printStateFlag = true
	printFinalFlag = true
)

func dprint(msg string) {
	if debugMode {
		fmt.Println(msg)
	}
}

// ---------------------------------------------------------------------------
// Bitmap — tracks allocation of inodes or data blocks
// ---------------------------------------------------------------------------

type Bitmap struct {
	size         int
	bmap         []int
	numAllocated int
}

func NewBitmap(size int) *Bitmap {
	return &Bitmap{
		size: size,
		bmap: make([]int, size),
	}
}

// Alloc finds the first free bit, marks it allocated, returns its index.
// Returns -1 if no free bit exists.
func (b *Bitmap) Alloc() int {
	for i := 0; i < len(b.bmap); i++ {
		if b.bmap[i] == 0 {
			b.bmap[i] = 1
			b.numAllocated++
			return i
		}
	}
	return -1
}

// Free releases the bit at position num.
func (b *Bitmap) Free(num int) {
	if b.bmap[num] != 1 {
		panic(fmt.Sprintf("bitmap.Free: bit %d is not allocated", num))
	}
	b.numAllocated--
	b.bmap[num] = 0
}

// MarkAllocated forcibly marks a bit as allocated (used for root init).
func (b *Bitmap) MarkAllocated(num int) {
	if b.bmap[num] != 0 {
		panic(fmt.Sprintf("bitmap.MarkAllocated: bit %d is already allocated", num))
	}
	b.numAllocated++
	b.bmap[num] = 1
}

// NumFree returns how many bits are still free.
func (b *Bitmap) NumFree() int {
	return b.size - b.numAllocated
}

// Dump returns a string like "10000000".
func (b *Bitmap) Dump() string {
	s := ""
	for _, v := range b.bmap {
		s += fmt.Sprintf("%d", v)
	}
	return s
}

// ---------------------------------------------------------------------------
// DirEntry — a single (name, inode-number) pair inside a directory block
// ---------------------------------------------------------------------------

type DirEntry struct {
	Name string
	Inum int
}

// ---------------------------------------------------------------------------
// Block — one data block (free | directory | file-data)
// ---------------------------------------------------------------------------

type Block struct {
	ftype   string // "d", "f", "free"
	dirUsed int
	maxUsed int
	dirList []DirEntry
	data    string
}

func NewBlock(ftype string) *Block {
	if ftype != "d" && ftype != "f" && ftype != "free" {
		panic("block: invalid ftype: " + ftype)
	}
	return &Block{
		ftype:   ftype,
		dirUsed: 0,
		maxUsed: 32,
		dirList: []DirEntry{},
		data:    "",
	}
}

// Dump returns a human-readable representation of the block's contents.
func (b *Block) Dump() string {
	if b.ftype == "free" {
		return "[]"
	} else if b.ftype == "d" {
		rc := ""
		for _, d := range b.dirList {
			short := fmt.Sprintf("(%s,%d)", d.Name, d.Inum)
			if rc == "" {
				rc = short
			} else {
				rc += " " + short
			}
		}
		return "[" + rc + "]"
	} else {
		return fmt.Sprintf("[%s]", b.data)
	}
}

// SetType transitions block from "free" to the given type.
func (b *Block) SetType(ftype string) {
	if b.ftype != "free" {
		panic("block.SetType: block is not free")
	}
	b.ftype = ftype
}

// AddData writes data into a file block.
func (b *Block) AddData(data string) {
	if b.ftype != "f" {
		panic("block.AddData: block is not a file")
	}
	b.data = data
}

// GetNumEntries returns how many directory entries exist.
func (b *Block) GetNumEntries() int {
	if b.ftype != "d" {
		panic("block.GetNumEntries: not a directory")
	}
	return b.dirUsed
}

// GetFreeEntries returns remaining capacity.
func (b *Block) GetFreeEntries() int {
	if b.ftype != "d" {
		panic("block.GetFreeEntries: not a directory")
	}
	return b.maxUsed - b.dirUsed
}

// GetEntry returns the n-th directory entry.
func (b *Block) GetEntry(num int) DirEntry {
	if b.ftype != "d" {
		panic("block.GetEntry: not a directory")
	}
	if num >= b.dirUsed {
		panic("block.GetEntry: index out of range")
	}
	return b.dirList[num]
}

// AddDirEntry appends a (name, inum) entry to the directory.
func (b *Block) AddDirEntry(name string, inum int) {
	if b.ftype != "d" {
		panic("block.AddDirEntry: not a directory")
	}
	b.dirList = append(b.dirList, DirEntry{Name: name, Inum: inum})
	b.dirUsed++
	if b.dirUsed > b.maxUsed {
		panic("block.AddDirEntry: directory full")
	}
}

// DelDirEntry removes the entry whose basename matches name.
func (b *Block) DelDirEntry(name string) {
	if b.ftype != "d" {
		panic("block.DelDirEntry: not a directory")
	}
	parts := strings.Split(name, "/")
	dname := parts[len(parts)-1]
	for i := 0; i < len(b.dirList); i++ {
		if b.dirList[i].Name == dname {
			b.dirList = append(b.dirList[:i], b.dirList[i+1:]...)
			b.dirUsed--
			return
		}
	}
	panic("block.DelDirEntry: entry not found: " + name)
}

// DirEntryExists checks whether name already exists in the directory.
func (b *Block) DirEntryExists(name string) bool {
	if b.ftype != "d" {
		panic("block.DirEntryExists: not a directory")
	}
	for _, d := range b.dirList {
		if d.Name == name {
			return true
		}
	}
	return false
}

// FreeBlock releases the block back to "free" state.
func (b *Block) FreeBlock() {
	if b.ftype == "free" {
		panic("block.FreeBlock: already free")
	}
	if b.ftype == "d" {
		if b.dirUsed != 2 {
			panic("block.FreeBlock: directory still has entries beyond . and ..")
		}
		b.dirUsed = 0
	}
	b.data = ""
	b.ftype = "free"
	b.dirList = []DirEntry{}
}

// ---------------------------------------------------------------------------
// Inode — metadata for one file or directory
// ---------------------------------------------------------------------------

type Inode struct {
	ftype  string // "d", "f", "free"
	addr   int    // index of the data block (-1 = none)
	refCnt int
}

func NewInode() *Inode {
	return &Inode{ftype: "free", addr: -1, refCnt: 1}
}

func (n *Inode) SetAll(ftype string, addr int, refCnt int) {
	if ftype != "d" && ftype != "f" && ftype != "free" {
		panic("inode.SetAll: invalid ftype: " + ftype)
	}
	n.ftype = ftype
	n.addr = addr
	n.refCnt = refCnt
}

func (n *Inode) IncRefCnt()       { n.refCnt++ }
func (n *Inode) DecRefCnt()       { n.refCnt-- }
func (n *Inode) GetRefCnt() int   { return n.refCnt }
func (n *Inode) SetType(t string) { n.ftype = t }
func (n *Inode) SetAddr(b int)    { n.addr = b }
func (n *Inode) GetAddr() int     { return n.addr }
func (n *Inode) GetType() string  { return n.ftype }

// GetSize returns 0 if no data block, 1 otherwise (max 1 block per file).
func (n *Inode) GetSize() int {
	if n.addr == -1 {
		return 0
	}
	return 1
}

// FreeInode resets the inode to the free state.
func (n *Inode) FreeInode() {
	n.ftype = "free"
	n.addr = -1
}

// ---------------------------------------------------------------------------
// FS — the file-system simulator
// ---------------------------------------------------------------------------

type FS struct {
	numInodes int
	numData   int

	ibitmap *Bitmap
	inodes  []*Inode

	dbitmap *Bitmap
	data    []*Block

	root int // root inode number (always 0)

	// workload bookkeeping
	files      []string
	dirs       []string
	nameToInum map[string]int

	rng *rand.Rand // deterministic random source
}

func NewFS(numInodes, numData int, rng *rand.Rand) *FS {
	fs := &FS{
		numInodes:  numInodes,
		numData:    numData,
		ibitmap:    NewBitmap(numInodes),
		dbitmap:    NewBitmap(numData),
		root:       0,
		files:      []string{},
		dirs:       []string{"/"},
		nameToInum: map[string]int{"/": 0},
		rng:        rng,
	}

	fs.inodes = make([]*Inode, numInodes)
	for i := 0; i < numInodes; i++ {
		fs.inodes[i] = NewInode()
	}

	fs.data = make([]*Block, numData)
	for i := 0; i < numData; i++ {
		fs.data[i] = NewBlock("free")
	}

	// create root directory
	fs.ibitmap.MarkAllocated(fs.root)
	fs.inodes[fs.root].SetAll("d", 0, 2)
	fs.dbitmap.MarkAllocated(fs.root)
	fs.data[0].SetType("d")
	fs.data[0].AddDirEntry(".", fs.root)
	fs.data[0].AddDirEntry("..", fs.root)

	return fs
}

// Dump prints the entire file system state.
func (fs *FS) Dump() {
	fmt.Println("inode bitmap ", fs.ibitmap.Dump())
	fmt.Print("inodes       ")
	for i := 0; i < fs.numInodes; i++ {
		ftype := fs.inodes[i].GetType()
		if ftype == "free" {
			fmt.Print("[]")
		} else {
			fmt.Printf("[%s a:%d r:%d]", ftype, fs.inodes[i].GetAddr(), fs.inodes[i].GetRefCnt())
		}
	}
	fmt.Println()
	fmt.Println("data bitmap  ", fs.dbitmap.Dump())
	fmt.Print("data         ")
	for i := 0; i < fs.numData; i++ {
		fmt.Print(fs.data[i].Dump())
	}
	fmt.Println()
}

// makeName generates a random single-character file name.
func (fs *FS) makeName() string {
	p := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k',
		'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	return string(p[int(fs.rng.Float64()*float64(len(p)))])
}

func (fs *FS) inodeAlloc() int {
	return fs.ibitmap.Alloc()
}

func (fs *FS) inodeFree(inum int) {
	fs.ibitmap.Free(inum)
	fs.inodes[inum].FreeInode()
}

func (fs *FS) dataAlloc() int {
	return fs.dbitmap.Alloc()
}

func (fs *FS) dataFree(bnum int) {
	fs.dbitmap.Free(bnum)
	fs.data[bnum].FreeBlock()
}

// getParent returns the parent path of a given full path.
//
//	"/a/b/c" -> "/a/b"
//	"/c"     -> "/"
func (fs *FS) getParent(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) == 2 {
		return "/"
	}
	pname := ""
	for i := 1; i < len(parts)-1; i++ {
		pname += "/" + parts[i]
	}
	return pname
}

// -----------------------------------------------------------------------
// Core file-system operations
// -----------------------------------------------------------------------

// deleteFile unlinks a file or directory from the file system.
func (fs *FS) deleteFile(tfile string) int {
	if printOpsFlag {
		fmt.Printf("unlink(\"%s\");\n", tfile)
	}

	inum := fs.nameToInum[tfile]
	ftype := fs.inodes[inum].GetType()

	if fs.inodes[inum].GetRefCnt() == 1 {
		// free data block first
		dblock := fs.inodes[inum].GetAddr()
		if dblock != -1 {
			fs.dataFree(dblock)
		}
		// then free inode
		fs.inodeFree(inum)
	} else {
		fs.inodes[inum].DecRefCnt()
	}

	// remove entry from parent directory
	parent := fs.getParent(tfile)
	pinum := fs.nameToInum[parent]
	pblock := fs.inodes[pinum].GetAddr()
	// if deleting a directory, decrease parent refcnt
	if ftype == "d" {
		fs.inodes[pinum].DecRefCnt()
	}
	fs.data[pblock].DelDirEntry(tfile)

	// remove from internal tracking list
	for i, f := range fs.files {
		if f == tfile {
			fs.files = append(fs.files[:i], fs.files[i+1:]...)
			break
		}
	}
	return 0
}

// createLink creates a hard link from target to newfile inside parent.
func (fs *FS) createLink(target, newfile, parent string) int {
	parentInum := fs.nameToInum[parent]

	pblock := fs.inodes[parentInum].GetAddr()
	if fs.data[pblock].GetFreeEntries() <= 0 {
		dprint("*** createLink failed: no room in parent directory ***")
		return -1
	}

	if fs.data[pblock].DirEntryExists(newfile) {
		dprint("*** createLink failed: not a unique name ***")
		return -1
	}

	tinum := fs.nameToInum[target]
	fs.inodes[tinum].IncRefCnt()

	parts := strings.Split(newfile, "/")
	ename := parts[len(parts)-1]
	fs.data[pblock].AddDirEntry(ename, tinum)
	return tinum
}

// createFile creates a new file or directory.
func (fs *FS) createFile(parent, newfile, ftype string) int {
	parentInum := fs.nameToInum[parent]

	pblock := fs.inodes[parentInum].GetAddr()
	if fs.data[pblock].GetFreeEntries() <= 0 {
		dprint("*** createFile failed: no room in parent directory ***")
		return -1
	}

	block := fs.inodes[parentInum].GetAddr()
	if fs.data[block].DirEntryExists(newfile) {
		dprint("*** createFile failed: not a unique name ***")
		return -1
	}

	inum := fs.inodeAlloc()
	if inum == -1 {
		dprint("*** createFile failed: no inodes left ***")
		return -1
	}

	fblock := -1
	refCnt := 1
	if ftype == "d" {
		refCnt = 2
		fblock = fs.dataAlloc()
		if fblock == -1 {
			dprint("*** createFile failed: no data blocks left ***")
			fs.inodeFree(inum)
			return -1
		}
		fs.data[fblock].SetType("d")
		fs.data[fblock].AddDirEntry(".", inum)
		fs.data[fblock].AddDirEntry("..", parentInum)
	}

	fs.inodes[inum].SetAll(ftype, fblock, refCnt)

	if ftype == "d" {
		fs.inodes[parentInum].IncRefCnt()
	}

	fs.data[pblock].AddDirEntry(newfile, inum)
	return inum
}

// writeFile appends one block of data to a file.
func (fs *FS) writeFile(tfile, data string) int {
	inum := fs.nameToInum[tfile]
	curSize := fs.inodes[inum].GetSize()
	dprint(fmt.Sprintf("writeFile: inum:%d cursize:%d refcnt:%d", inum, curSize, fs.inodes[inum].GetRefCnt()))

	if curSize == 1 {
		dprint("*** writeFile failed: file is full ***")
		return -1
	}

	fblock := fs.dataAlloc()
	if fblock == -1 {
		dprint("*** writeFile failed: no data blocks left ***")
		return -1
	}

	fs.data[fblock].SetType("f")
	fs.data[fblock].AddData(data)
	fs.inodes[inum].SetAddr(fblock)

	if printOpsFlag {
		fmt.Printf("fd=open(\"%s\", O_WRONLY|O_APPEND); write(fd, buf, BLOCKSIZE); close(fd);\n", tfile)
	}
	return 0
}

// -----------------------------------------------------------------------
// Random workload generators
// -----------------------------------------------------------------------

func (fs *FS) doDelete() int {
	dprint("doDelete")
	if len(fs.files) == 0 {
		return -1
	}
	dfile := fs.files[int(fs.rng.Float64()*float64(len(fs.files)))]
	dprint("try delete(" + dfile + ")")
	return fs.deleteFile(dfile)
}

func (fs *FS) doLink() int {
	dprint("doLink")
	if len(fs.files) == 0 {
		return -1
	}
	parent := fs.dirs[int(fs.rng.Float64()*float64(len(fs.dirs)))]
	nfile := fs.makeName()

	target := fs.files[int(fs.rng.Float64()*float64(len(fs.files)))]

	var fullName string
	if parent == "/" {
		fullName = parent + nfile
	} else {
		fullName = parent + "/" + nfile
	}

	dprint(fmt.Sprintf("try createLink(%s %s %s)", target, nfile, parent))
	inum := fs.createLink(target, nfile, parent)
	if inum >= 0 {
		fs.files = append(fs.files, fullName)
		fs.nameToInum[fullName] = inum
		if printOpsFlag {
			fmt.Printf("link(\"%s\", \"%s\");\n", target, fullName)
		}
		return 0
	}
	return -1
}

func (fs *FS) doCreate(ftype string) int {
	dprint("doCreate")
	parent := fs.dirs[int(fs.rng.Float64()*float64(len(fs.dirs)))]
	nfile := fs.makeName()

	var tlist *[]string
	if ftype == "d" {
		tlist = &fs.dirs
	} else {
		tlist = &fs.files
	}

	var fullName string
	if parent == "/" {
		fullName = parent + nfile
	} else {
		fullName = parent + "/" + nfile
	}

	dprint(fmt.Sprintf("try createFile(%s %s %s)", parent, nfile, ftype))
	inum := fs.createFile(parent, nfile, ftype)
	if inum >= 0 {
		*tlist = append(*tlist, fullName)
		fs.nameToInum[fullName] = inum
		if parent == "/" {
			parent = ""
		}
		if ftype == "d" {
			if printOpsFlag {
				fmt.Printf("mkdir(\"%s/%s\");\n", parent, nfile)
			}
		} else {
			if printOpsFlag {
				fmt.Printf("creat(\"%s/%s\");\n", parent, nfile)
			}
		}
		return 0
	}
	return -1
}

func (fs *FS) doAppend() int {
	dprint("doAppend")
	if len(fs.files) == 0 {
		return -1
	}
	afile := fs.files[int(fs.rng.Float64()*float64(len(fs.files)))]
	dprint("try writeFile(" + afile + ")")
	data := string(rune('a' + int(fs.rng.Float64()*26)))
	return fs.writeFile(afile, data)
}

// -----------------------------------------------------------------------
// Main simulation loop
// -----------------------------------------------------------------------

// Run executes the simulation for numRequests operations.
func (fs *FS) Run(numRequests int) {
	fmt.Println("Initial state")
	fmt.Println()
	fs.Dump()
	fmt.Println()

	for i := 0; i < numRequests; i++ {
		if !printOpsFlag {
			fmt.Println("Which operation took place?")
		}

		rc := -1
		for rc == -1 {
			r := fs.rng.Float64()
			if r < 0.3 {
				rc = fs.doAppend()
				dprint(fmt.Sprintf("doAppend rc:%d", rc))
			} else if r < 0.5 {
				rc = fs.doDelete()
				dprint(fmt.Sprintf("doDelete rc:%d", rc))
			} else if r < 0.7 {
				rc = fs.doLink()
				dprint(fmt.Sprintf("doLink rc:%d", rc))
			} else {
				if fs.rng.Float64() < 0.75 {
					rc = fs.doCreate("f")
					dprint(fmt.Sprintf("doCreate(f) rc:%d", rc))
				} else {
					rc = fs.doCreate("d")
					dprint(fmt.Sprintf("doCreate(d) rc:%d", rc))
				}
			}

			if fs.ibitmap.NumFree() == 0 {
				fmt.Println("File system out of inodes; rerun with more via command-line flag?")
				os.Exit(1)
			}
			if fs.dbitmap.NumFree() == 0 {
				fmt.Println("File system out of data blocks; rerun with more via command-line flag?")
				os.Exit(1)
			}
		}

		if printStateFlag {
			fmt.Println()
			fs.Dump()
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("  State of file system (inode bitmap, inodes, data bitmap, data)?")
			fmt.Println()
		}
	}

	if printFinalFlag {
		fmt.Println()
		fmt.Println("Summary of files, directories::")
		fmt.Println()
		fmt.Printf("  Files:       %v\n", fs.files)
		fmt.Printf("  Directories: %v\n", fs.dirs)
		fmt.Println()
	}
}

// ---------------------------------------------------------------------------
// main
// ---------------------------------------------------------------------------

func main() {
	seed := flag.Int64("s", 0, "the random seed")
	numInodes := flag.Int("i", 8, "number of inodes in file system")
	numData := flag.Int("d", 8, "number of data blocks in file system")
	numRequests := flag.Int("n", 10, "number of requests to simulate")
	reverse := flag.Bool("r", false, "instead of printing state, print ops")
	showFinal := flag.Bool("p", false, "print the final set of files/dirs")
	solve := flag.Bool("c", false, "compute answers for me")

	flag.Parse()

	fmt.Println("ARG seed", *seed)
	fmt.Println("ARG numInodes", *numInodes)
	fmt.Println("ARG numData", *numData)
	fmt.Println("ARG numRequests", *numRequests)
	fmt.Println("ARG reverse", *reverse)
	fmt.Println("ARG printFinal", *showFinal)
	fmt.Println()

	rng := rand.New(rand.NewSource(*seed))

	if *reverse {
		printStateFlag = false
		printOpsFlag = true
	} else {
		printStateFlag = true
		printOpsFlag = false
	}

	if *solve {
		printOpsFlag = true
		printStateFlag = true
	}

	printFinalFlag = *showFinal

	f := NewFS(*numInodes, *numData, rng)
	f.Run(*numRequests)
}
