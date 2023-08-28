package BufferPool

// denotes the size of the buffer pool
const MaxPoolSize = 4

type BufferPoolManager struct {
	diskManager DiskManager
	pages       [MaxPoolSize]*Page // represent frames of fixed size
	replacer    *ClockReplacer     // for cache replacement
	freeList    []FrameID          // tracking available frames
	pageTable   map[PageId]FrameId // mapping frameIDs with pagesIDs
}

// for requesting a new Page
func (bp *BufferPoolManager) NewPage() *Page {

}

// for fetching a new Page
func (bp *BufferPoolManager) FetchPage(pageID PageID) *Page {

}

// write the page to disk
func (bp *BufferPoolManager) FlushPage(pageID PageID) bool {

}

// write all pages to disk
func (bp *BufferPoolManager) FlushAllPages() {

}

// delete the page from buffer pool
func (bp *BufferPoolManager) DeletePage(pageID PageID) error {

}

// marks the page safe for eviction
func (bp *BufferPoolManager) UnpinPage(pageID PageID, isDirty bool) error {

}
