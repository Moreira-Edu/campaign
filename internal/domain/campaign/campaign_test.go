package campaign

import (
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "Campaign x"
	content := "body"
	contacts := []string{"emailum@email.com", "emaildois@email.com"}

	campaing := NewCampaign(name, content, contacts)

	if campaing.Id != "2" {
		t.Errorf("expected 2")
	}
	if campaing.Name != name {
		t.Errorf("expected correct name")
	}
	if campaing.Content != content {
		t.Errorf("expected correct Content")
	}
	if len(campaing.Contacts) != 2 {
		t.Errorf("expected correct Content")
	}

}
