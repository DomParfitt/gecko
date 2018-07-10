package union

//Union struct representing a sum type
type Union struct {
	this *interface{}
	that *interface{}
}

//New union
func New(this, that *interface{}) Union {
	return Union{
		this: this,
		that: that,
	}
}

//Value of the union
func (u Union) Value() *interface{} {
	if u.this != nil {
		return u.this
	}

	return u.that
}
