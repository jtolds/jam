package cache

import (
	"fmt"
)

type cappedMisraGries struct {
	k, freqCap int
	a          map[string]int
}

func newCappedMisraGries(k, freqCap int) (*cappedMisraGries, error) {
	if freqCap < 1 {
		return nil, fmt.Errorf("invalid cache frequency cap")
	}
	return &cappedMisraGries{
		k:       k,
		freqCap: freqCap,
		a:       map[string]int{},
	}, nil
}

func (m *cappedMisraGries) Observe(elem string) (capped bool) {
	if count, exists := m.a[elem]; exists {
		count++
		if count > m.freqCap {
			return true
		}
		m.a[elem] = count
		return false
	}

	if len(m.a) < m.k-1 {
		m.a[elem] = 1
		return false
	}

	for elem, count := range m.a {
		count--
		if count > 0 {
			m.a[elem] = count
		} else {
			delete(m.a, elem)
		}
	}
	return false
}

func (m *cappedMisraGries) Delete(elem string) {
	delete(m.a, elem)
}

func (m *cappedMisraGries) Save() (rv []*HitCount) {
	for elem, count := range m.a {
		rv = append(rv, &HitCount{Entry: elem, Count: int64(count)})
	}
	return rv
}

func (m *cappedMisraGries) Load(vals []*HitCount) {
	for _, v := range vals {
		m.a[v.Entry] = int(v.Count)
	}
}
