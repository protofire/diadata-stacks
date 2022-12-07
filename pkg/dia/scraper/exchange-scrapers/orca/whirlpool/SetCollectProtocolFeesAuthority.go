// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetCollectProtocolFeesAuthority is the `setCollectProtocolFeesAuthority` instruction.
type SetCollectProtocolFeesAuthority struct {

	// [0] = [WRITE] whirlpoolsConfig
	//
	// [1] = [SIGNER] collectProtocolFeesAuthority
	//
	// [2] = [] newCollectProtocolFeesAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetCollectProtocolFeesAuthorityInstructionBuilder creates a new `SetCollectProtocolFeesAuthority` instruction builder.
func NewSetCollectProtocolFeesAuthorityInstructionBuilder() *SetCollectProtocolFeesAuthority {
	nd := &SetCollectProtocolFeesAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetWhirlpoolsConfigAccount sets the "whirlpoolsConfig" account.
func (inst *SetCollectProtocolFeesAuthority) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *SetCollectProtocolFeesAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig).WRITE()
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpoolsConfig" account.
func (inst *SetCollectProtocolFeesAuthority) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectProtocolFeesAuthorityAccount sets the "collectProtocolFeesAuthority" account.
func (inst *SetCollectProtocolFeesAuthority) SetCollectProtocolFeesAuthorityAccount(collectProtocolFeesAuthority ag_solanago.PublicKey) *SetCollectProtocolFeesAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collectProtocolFeesAuthority).SIGNER()
	return inst
}

// GetCollectProtocolFeesAuthorityAccount gets the "collectProtocolFeesAuthority" account.
func (inst *SetCollectProtocolFeesAuthority) GetCollectProtocolFeesAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetNewCollectProtocolFeesAuthorityAccount sets the "newCollectProtocolFeesAuthority" account.
func (inst *SetCollectProtocolFeesAuthority) SetNewCollectProtocolFeesAuthorityAccount(newCollectProtocolFeesAuthority ag_solanago.PublicKey) *SetCollectProtocolFeesAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(newCollectProtocolFeesAuthority)
	return inst
}

// GetNewCollectProtocolFeesAuthorityAccount gets the "newCollectProtocolFeesAuthority" account.
func (inst *SetCollectProtocolFeesAuthority) GetNewCollectProtocolFeesAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst SetCollectProtocolFeesAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetCollectProtocolFeesAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetCollectProtocolFeesAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetCollectProtocolFeesAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CollectProtocolFeesAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NewCollectProtocolFeesAuthority is not set")
		}
	}
	return nil
}

func (inst *SetCollectProtocolFeesAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetCollectProtocolFeesAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               whirlpoolsConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   collectProtocolFeesAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("newCollectProtocolFeesAuthority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj SetCollectProtocolFeesAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetCollectProtocolFeesAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetCollectProtocolFeesAuthorityInstruction declares a new SetCollectProtocolFeesAuthority instruction with the provided parameters and accounts.
func NewSetCollectProtocolFeesAuthorityInstruction(
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	collectProtocolFeesAuthority ag_solanago.PublicKey,
	newCollectProtocolFeesAuthority ag_solanago.PublicKey) *SetCollectProtocolFeesAuthority {
	return NewSetCollectProtocolFeesAuthorityInstructionBuilder().
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetCollectProtocolFeesAuthorityAccount(collectProtocolFeesAuthority).
		SetNewCollectProtocolFeesAuthorityAccount(newCollectProtocolFeesAuthority)
}
