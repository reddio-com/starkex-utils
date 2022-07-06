package verify

import "testing"

func Test_VerifyTransaction(t *testing.T) {
	verify, err := VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != true {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "5000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "1", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "2", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65d", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "1", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "419430", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cf", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743b")
	if err != nil {
		t.Error(err)
	}

	if verify != false {
		t.Error("verify failed")
	}

	verify, err = VerifyTransaction("0x38cae143fe6d2b8bdb7051f211744017d98f7e6a67e45a5dfc08759c119cf3c", "50000000000000", "12", "0x1142460171646987f20c714eda4b92812b22b811f56f27130937c267e29bd9e", "1", "0xc664b68afced392656ed8c4adaefa8e8ffbf65dc", "2", "4194303", "0x1876a0bb08e6cc5999e0cf6bea82b725c752b014380a092fa952ad025ba6cfa", "0x63cab8f658f0ef532a530b7f409446123744a2de8d85bf009d8a092bf6743be")
	if err != nil {
		t.Error(err)
	}

	if verify != true {
		t.Error("verify failed")
	}

}
