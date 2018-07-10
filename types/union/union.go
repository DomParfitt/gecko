package union

type Union struct {
	this *interface{}
	that *interface{}
}

func NewUnion(this, that *interface{}) Union {
	return Union{
		this: this,
		that: that,
	}
}

func (u Union) Value() *interface{} {
	if u.this != nil {
		return u.this
	}

	return u.that
}
