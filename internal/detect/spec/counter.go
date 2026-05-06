package spec

type RejectCounter interface {
	Inc(reason RejectReason)
}
