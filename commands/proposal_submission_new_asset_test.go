package commands_test

import (
	"testing"

	"code.vegaprotocol.io/protos/commands"
	types "code.vegaprotocol.io/protos/vega"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"github.com/stretchr/testify/assert"
)

func TestCheckProposalSubmissionForNewAsset(t *testing.T) {
	t.Run("Submitting an asset change without change fails", testProposalSubmissionWithoutChangeFails)
	t.Run("Submitting an asset change without new asset fails", testAssetChangeSubmissionWithoutNewsAssetFails)
	t.Run("Submitting an asset change without changes fails", testAssetChangeSubmissionWithoutChangesFails)
	t.Run("Submitting an asset change without source fails", testAssetChangeSubmissionWithoutSourceFails)
	t.Run("Submitting an built-in asset change without built-in asset fails", testBuiltInAssetChangeSubmissionWithoutBuiltInAssetFails)
	t.Run("Submitting an built-in asset change without name fails", testBuiltInAssetChangeSubmissionWithoutNameFails)
	t.Run("Submitting an built-in asset change with name succeeds", testBuiltInAssetChangeSubmissionWithNameSucceeds)
	t.Run("Submitting an built-in asset change without symbol fails", testBuiltInAssetChangeSubmissionWithoutSymbolFails)
	t.Run("Submitting an built-in asset change with symbol succeeds", testBuiltInAssetChangeSubmissionWithSymbolSucceeds)
	t.Run("Submitting an built-in asset change without decimal fails", testBuiltInAssetChangeSubmissionWithoutDecimalsFails)
	t.Run("Submitting an built-in asset change with decimal succeeds", testBuiltInAssetChangeSubmissionWithDecimalsSucceeds)
	t.Run("Submitting an built-in asset change without total supply fails", testBuiltInAssetChangeSubmissionWithoutTotalSupplyFails)
	t.Run("Submitting an built-in asset change with total supply succeeds", testBuiltInAssetChangeSubmissionWithTotalSupplySucceeds)
	t.Run("Submitting an built-in asset change with not-a-number total supply fails", testBuiltInAssetChangeSubmissionWithNaNTotalSupplyFails)
	t.Run("Submitting an built-in asset change without max faucet amount fails", testBuiltInAssetChangeSubmissionWithoutMaxFaucetAmountMintFails)
	t.Run("Submitting an built-in asset change with max faucet amount succeeds", testBuiltInAssetChangeSubmissionWithMaxFaucetAmountMintSucceeds)
	t.Run("Submitting an built-in asset change with not-a-number max faucet amount fails", testBuiltInAssetChangeSubmissionWithNaNMaxFaucetAmountMintFails)
	t.Run("Submitting an ERC20 asset change without ERC20 asset fails", testERC20AssetChangeSubmissionWithoutErc20AssetFails)
	t.Run("Submitting an ERC20 asset change without contract address fails", testErc20AssetChangeSubmissionWithoutContractAddressFails)
	t.Run("Submitting an ERC20 asset change with contract address succeeds", testErc20AssetChangeSubmissionWithContractAddressSucceeds)
}

func testProposalSubmissionWithoutChangeFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change"), commands.ErrIsRequired)
}

func testAssetChangeSubmissionWithoutNewsAssetFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset"), commands.ErrIsRequired)
}

func testAssetChangeSubmissionWithoutChangesFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes"), commands.ErrIsRequired)
}

func testAssetChangeSubmissionWithoutSourceFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithoutBuiltInAssetFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_BuiltinAsset{},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithoutNameFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Name: "",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.name"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithNameSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Name: "My built-in asset",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.name"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithoutSymbolFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Symbol: "",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.symbol"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithSymbolSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Symbol: "My symbol",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.symbol"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithoutDecimalsFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Decimals: 0,
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.decimals"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithDecimalsSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Decimals: RandomPositiveU64(),
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.decimals"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithoutTotalSupplyFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						TotalSupply: "",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.total_supply"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithTotalSupplySucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						TotalSupply: "10000",
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.total_supply"), commands.ErrIsRequired)
	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.total_supply"), commands.ErrIsNotValidNumber)
	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.total_supply"), commands.ErrMustBePositive)
}

func testBuiltInAssetChangeSubmissionWithNaNTotalSupplyFails(t *testing.T) {
	testCases := []struct {
		msg   string
		value string
		error error
	}{
		{
			msg:   "with not-a-number value",
			value: "hello",
			error: commands.ErrIsNotValidNumber,
		}, {
			msg:   "with value of 0",
			value: "0",
			error: commands.ErrMustBePositive,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					Change: &types.ProposalTerms_NewAsset{
						NewAsset: &types.NewAsset{
							Changes: &types.AssetDetails{
								TotalSupply: tc.value,
								Source: &types.AssetDetails_BuiltinAsset{
									BuiltinAsset: &types.BuiltinAsset{},
								},
							},
						},
					},
				},
			})

			assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.total_supply"), tc.error)
		})
	}
}

func testBuiltInAssetChangeSubmissionWithoutMaxFaucetAmountMintFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{
								MaxFaucetAmountMint: "",
							},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.max_faucet_amount_mint"), commands.ErrIsRequired)
}

func testBuiltInAssetChangeSubmissionWithMaxFaucetAmountMintSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_BuiltinAsset{
							BuiltinAsset: &types.BuiltinAsset{
								MaxFaucetAmountMint: "10000",
							},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.max_faucet_amount_mint"), commands.ErrIsRequired)
	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.max_faucet_amount_mint"), commands.ErrIsNotValidNumber)
	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.max_faucet_amount_mint"), commands.ErrMustBePositive)
}

func testBuiltInAssetChangeSubmissionWithNaNMaxFaucetAmountMintFails(t *testing.T) {
	testCases := []struct {
		msg   string
		value string
		error error
	}{
		{
			msg:   "with not-a-number value",
			value: "hello",
			error: commands.ErrIsNotValidNumber,
		}, {
			msg:   "with value of 0",
			value: "0",
			error: commands.ErrMustBePositive,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					Change: &types.ProposalTerms_NewAsset{
						NewAsset: &types.NewAsset{
							Changes: &types.AssetDetails{
								Source: &types.AssetDetails_BuiltinAsset{
									BuiltinAsset: &types.BuiltinAsset{
										MaxFaucetAmountMint: tc.value,
									},
								},
							},
						},
					},
				},
			})

			assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.builtin_asset.max_faucet_amount_mint"), tc.error)
		})
	}
}

func testERC20AssetChangeSubmissionWithoutErc20AssetFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_Erc20{},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.erc20"), commands.ErrIsRequired)
}

func testErc20AssetChangeSubmissionWithoutContractAddressFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_Erc20{
							Erc20: &types.ERC20{
								ContractAddress: "",
							},
						},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.erc20.contract_address"), commands.ErrIsRequired)
}

func testErc20AssetChangeSubmissionWithContractAddressSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_NewAsset{
				NewAsset: &types.NewAsset{
					Changes: &types.AssetDetails{
						Source: &types.AssetDetails_Erc20{
							Erc20: &types.ERC20{
								ContractAddress: "My address",
							},
						},
					},
				},
			},
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.change.new_asset.changes.source.erc20.contract_address"), commands.ErrIsRequired)
}
