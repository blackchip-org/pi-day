package main

import "math/big"

type Lft struct {
	q, r, s, t big.Int
}

func (t *Lft) extr(x *big.Int) *big.Rat {
	var n, d big.Int
	var r big.Rat
	return r.SetFrac(
		n.Add(n.Mul(&t.q, x), &t.r),
		d.Add(d.Mul(&t.s, x), &t.t))
}

var three = big.NewInt(3)
var four = big.NewInt(4)

func (t *Lft) Next() *big.Int {
	r := t.extr(three)
	var f big.Int
	return f.Div(r.Num(), r.Denom())
}

func (t *Lft) Safe(n *big.Int) bool {
	r := t.extr(four)
	var f big.Int
	if n.Cmp(f.Div(r.Num(), r.Denom())) == 0 {
		return true
	}
	return false
}

func (t *Lft) Comp(u *Lft) *Lft {
	var r Lft
	var a, b big.Int
	r.q.Add(a.Mul(&t.q, &u.q), b.Mul(&t.r, &u.s))
	r.r.Add(a.Mul(&t.q, &u.r), b.Mul(&t.r, &u.t))
	r.s.Add(a.Mul(&t.s, &u.q), b.Mul(&t.t, &u.s))
	r.t.Add(a.Mul(&t.s, &u.r), b.Mul(&t.t, &u.t))
	return &r
}

func (t *Lft) Prod(n *big.Int) *Lft {
	var r Lft
	r.q.SetInt64(10)
	r.r.Mul(r.r.SetInt64(-10), n)
	r.t.SetInt64(1)
	return r.Comp(t)
}
