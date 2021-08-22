package gopwn

type GadgetType int

const (
	GADGET_TYPE_ROP GadgetType = iota
	GADGET_TYPE_JOP
	GADGET_TYPE_SYS
)
