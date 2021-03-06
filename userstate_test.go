package permissions

import "testing"

func TestPerm(t *testing.T) {
	userstate := NewUserStateSimple()

	userstate.AddUser("bob", "hunter1", "bob@zombo.com")

	if !userstate.HasUser("bob") {
		t.Error("Error, user bob should exist")
	}

	if userstate.IsConfirmed("bob") {
		t.Error("Error, user bob should not be confirmed right now.")
	}

	userstate.MarkConfirmed("bob")

	if !userstate.IsConfirmed("bob") {
		t.Error("Error, user bob should be marked as confirmed right now.")
	}

	if userstate.IsAdmin("bob") {
		t.Error("Error, user bob should not have admin rights")
	}

	userstate.SetAdminStatus("bob")

	if !userstate.IsAdmin("bob") {
		t.Error("Error, user bob should have admin rights")
	}

	userstate.RemoveUser("bob")

	if userstate.HasUser("bob") {
		t.Error("Error, user bob should not exist")
	}
}

func TestPasswordBasic(t *testing.T) {
	userstate := NewUserStateSimple()

	// Assert that the default password algorithm is sha256
	if userstate.GetPasswordAlgo() != "sha256" {
		t.Error("Error, sha256 should be the default password algorithm")
	}

	// Set password algorithm
	userstate.SetPasswordAlgo("bcrypt")

	// Assert that the algorithm is now bcrypt
	if userstate.GetPasswordAlgo() != "bcrypt" {
		t.Error("Error, setting password algorithm failed")
	}

}

func TestPasswordAlgoMatching(t *testing.T) {
	userstate := NewUserStateSimple()
	// generate two different password using the same credentials but different algos
	userstate.SetPasswordAlgo("sha256")
	sha256_hash := userstate.HashPassword("testuser@example.com", "textpassword")
	userstate.SetPasswordAlgo("bcrypt")
	bcrypt_hash := userstate.HashPassword("testuser@example.com", "textpassword")

	// they shouldn't match
	if sha256_hash == bcrypt_hash {
		t.Error("Error, different algorithms should not have a password match")
	}
}
