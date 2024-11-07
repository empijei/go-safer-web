package auth_test

import (
	"context"
	"testing"

	auth "github.com/empijei/go-safer-web/auth/simple"
)

func TestGrantCheck(t *testing.T) {
	ctx := auth.Grant(context.Background(), "a", "b")
	cases := []struct {
		in      []string
		wantErr bool
	}{
		{[]string{}, false},
		{nil, false},
		{[]string{"a"}, false},
		{[]string{"b"}, false},
		{[]string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, true},
		{[]string{"c", "b"}, true},
		{[]string{"c"}, true},
	}
	for _, c := range cases {
		if _, err := auth.Check(ctx, c.in...); (err == nil && c.wantErr == true) || (err != nil && c.wantErr == false) {
			t.Errorf("Check(%v): got err:%v want error:%v", c.in, err, c.wantErr)
		}
	}
}

func TestCheckMust(t *testing.T) {
	// TODO test that if Grant or Check were never called we always return errors.

	t.Run("with grant", func(t *testing.T) {
		ctx := context.Background()
		if _, err := auth.Check(ctx, "a"); err == nil {
			t.Errorf("Check before grant, didn't get an error")
		}
		ctx = auth.Grant(ctx, "a", "b")
		ctx, err := auth.Check(ctx, "a", "b")
		if err != nil {
			t.Fatalf("Got unwanted error in setup calling Check: %q", err)
		}

		cases := []struct {
			in      []string
			wantErr bool
		}{
			{[]string{}, false},
			{nil, false},
			{[]string{"a"}, false},
			{[]string{"b"}, false},
			{[]string{"a", "b"}, false},
			{[]string{"a", "b", "c"}, true},
			{[]string{"c", "b"}, true},
			{[]string{"c"}, true},
		}
		for _, c := range cases {
			if err := auth.Must(ctx, c.in...); (err == nil && c.wantErr == true) || (err != nil && c.wantErr == false) {
				t.Errorf("Check(%v): got err:%v want error:%v", c.in, err, c.wantErr)
			}
		}
	})

	t.Run("without grant or check", func(t *testing.T) {
		check := func(ctx context.Context) {
			cases := []struct {
				in      []string
				wantErr bool
			}{
				{[]string{}, true},
				{nil, true},
				{[]string{"a"}, true},
			}
			t.Helper()
			for _, c := range cases {
				if err := auth.Must(ctx, c.in...); (err == nil && c.wantErr == true) || (err != nil && c.wantErr == false) {
					t.Errorf("Check(%v): got err:%v want error:%v", c.in, err, c.wantErr)
				}
			}
		}
		ctx := context.Background()
		check(ctx)
		ctx = auth.Grant(ctx, "a", "b")
		check(ctx)
	})
}
