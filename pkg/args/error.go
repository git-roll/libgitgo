package args

type RequisiteFailed struct {
    key Key
}

func (r RequisiteFailed) Error() string {
    return string(r.key) + " is required"
}
