package main

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	// 保存留待恢复的notifyUser
	saved := notifyUser
	defer func() {notifyUser = saved}()
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	const user = "yangchang@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyuser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("unexcepted notification message <<%s>>, " + "want substring %q", notifiedMsg, wantSubString)
	}
}
