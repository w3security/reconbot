package util

import (
	"testing"
)

// 更新到最新版本
func TestUpdatereconbotVersionToLatest(t *testing.T) {
	err := UpdatereconbotVersionToLatest(true)
	if err != nil {
		t.Error("fail TestupdateNucleiVersionToLatest")
	}
}
