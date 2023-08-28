package BufferPool

type DiskManager interface {
	ReadPage(id PageID) (*Page, error)
	WritePage(*Page) error
	AllocatePage() *PageID
	DeallocatePage(id PageID)
}
