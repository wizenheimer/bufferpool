package BufferPool

const pageSize = 5

type PageID int

// represents page on disk
type Page struct {
	id       PageID
	data     [pageSize]byte
	isDirty  bool
	pinCount int
}

// get the current pincount
func (pg *Page) PinCount() int {
	return pg.pinCount
}

// get the pageID for a given page
func (pg *Page) ID() PageID {
	return pg.id
}

// decrement the pinCount for the page
func (pg *Page) DecPinCount() {
	if pg.PinCount() > 0 {
		pg.pinCount -= 1
	}
}

// checks whether page is pinned
func (pg *Page) isPinned() bool {
	return pg.PinCount() > 0
}

// checks whether page is dirty
func (pg *Page) isPageDirty() bool {
	return pg.isDirty
}
