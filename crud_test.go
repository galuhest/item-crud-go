package crud

import "testing"

func TestGetUser(t *testing.T)	{
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCreateUser(t *testing.T)	{

}

func TestUpdateUser(t *testing.T)	{

}

func TestDeleteUser(t *testing.T)	{

}