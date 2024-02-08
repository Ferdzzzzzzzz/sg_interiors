package mux

import (
	"net/http"
	"sg/core/form"
	"sg/core/v"
	"sg/tmpls/page"
)

func (s service) contactAction(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	msg := r.FormValue("msg")
	honey := r.FormValue("email_honey")

	if honey != "" {
		s.l.Println("honeypot triggered")
		http.Redirect(w, r, r.URL.Path, http.StatusBadRequest)
		return nil
	}

	// ====================================================================
	// Form Validation

	state := page.ContactForm{
		State: form.Success,
		Name:  form.Val(name),
		Email: form.Val(email),
		Msg:   form.Val(msg),
	}

	if v.IsEmpty(name) {
		state.State = form.Initial
		state.Name.Err = "Field is required"
	}

	if v.IsEmpty(email) {
		state.State = form.Initial
		state.Email.Err = "Field is required"
	} else if !v.IsValidEmail(email) {
		state.State = form.Initial
		state.Email.Err = "Invalid Email"
	}

	if state.State != form.Success {
		return page.Contact(state).Render(r.Context(), w)
	}

	// ====================================================================
	// Alert

	if err := s.alerter.Contact(name, email, msg); err != nil {
		s.l.Printf("ERROR: %s", err)
		s.l.Printf("Tried to submit contact form with: %s, %s, %s", name, email, msg)

		return page.Contact(page.ContactForm{
			State: form.Error,
		}).Render(r.Context(), w)
	}

	return page.Contact(page.ContactForm{
		State: form.Success,
	}).Render(r.Context(), w)
}
