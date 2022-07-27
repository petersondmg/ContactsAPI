package contact

import "capi/domain/entity"

type ContactView struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	CellPhone string `json:"cellphone"`
}

type ContactsPayload struct {
	Contacts []ContactView `json:"contacts"`
}

func (p *ContactsPayload) Entities() []*entity.Contact {
	contacts := make([]*entity.Contact, len(p.Contacts))
	for i, c := range p.Contacts {
		contacts[i] = &entity.Contact{
			ID:    c.ID,
			Name:  c.Name,
			Phone: c.CellPhone,
		}
	}
	return contacts
}

func NewContactView(c *entity.Contact) ContactView {
	return ContactView{
		ID:        c.ID,
		Name:      c.Name,
		CellPhone: c.Phone,
	}
}

func NewContactListView(contacts []*entity.Contact) []ContactView {
	list := make([]ContactView, len(contacts))
	for i, c := range contacts {
		list[i] = NewContactView(c)
	}
	return list
}
