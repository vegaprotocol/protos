package commands_test

import (
	"testing"

	"code.vegaprotocol.io/protos/commands"
	types "code.vegaprotocol.io/protos/vega"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"github.com/stretchr/testify/assert"
)

func TestCheckProposalSubmissionForUpdateAsset(t *testing.T) {
	t.Run("Submitting an asset change without new asset fails", TestUpdateAssetChangeSubmissionWithoutUpdateAssetFails)
	t.Run("Submitting an asset change without changes fails", TestUpdateAssetChangeSubmissionWithoutChangesFails)
	t.Run("Submitting an asset change without source fails", TestUpdateAssetChangeSubmissionWithoutSourceFails)
	t.Run("Submitting an asset change without name fails", testUpdateAssetChangeSubmissionWithoutNameFails)
	t.Run("Submitting an asset change with name succeeds", testUpdateAssetChangeSubmissionWithNameSucceeds)
	t.Run("Submitting an asset change without symbol fails", testUpdateAssetChangeSubmissionWithoutSymbolFails)
	t.Run("Submitting an asset change with symbol succeeds", testUpdateAssetChangeSubmissionWithSymbolSucceeds)
	t.Run("Submitting an asset change without decimal fails", testUpdateAssetChangeSubmissionWithoutDecimalsFails)
	t.Run("Submitting an asset change with decimal succeeds", testUpdateAssetChangeSubmissionWithDecimalsSucceeds)
	t.Run("Submitting an asset change without total supply fails", testUpdateAssetChangeSubmissionWithoutTotalSupplyFails)
	t.Run("Submitting an asset change with total supply succeeds", testUpdateAssetChangeSubmissionWithTotalSupplySucceeds)
	t.Run("Submitting an asset change with not-a-number total supply fails", testUpdateAssetChangeSubmissionWithNaNTotalSupplyFails)
	t.Run("Submitting an ERC20 asset change without ERC20 asset fails", testUpdateERC20AssetChangeSubmissionWithoutErc20AssetFails)
	t.Run("Submitting an ERC20 asset change with invalid lifetime limit fails", testUpdateERC20AssetChangeSubmissionWithInvalidLifetimeLimitFails)
	t.Run("Submitting an ERC20 asset change with valid lifetime limit succeeds", testUpdateERC20AssetChangeSubmissionWithValidLifetimeLimitSucceeds)
	t.Run("Submitting an ERC20 asset change with invalid withdrawal threshold fails", testUpdateERC20AssetChangeSubmissionWithInvalidWithdrawalThresholdFails)
	t.Run("Submitting an ERC20 asset change with valid withdrawal threshold succeeds", testUpdateERC20AssetChangeSubmissionWithValidWithdrawalThresholdSucceeds)
}
func TestUpdateAssetChangeSubmissionWithoutUpdateAssetFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset"), commands.ErrIsRequired)
}

func TestUpdateAssetChangeSubmissionWithoutChangesFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes"), commands.ErrIsRequired)
}

func TestUpdateAssetChangeSubmissionWithoutSourceFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.source"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithoutNameFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Name: "",
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.name"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithNameSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Name: "My built-in asset",
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.name"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithoutSymbolFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Symbol: "",
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.symbol"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithSymbolSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Symbol: "My symbol",
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.symbol"))
}

func testUpdateAssetChangeSubmissionWithoutDecimalsFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Decimals: 0,
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.decimals"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithDecimalsSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Decimals: RandomPositiveU64(),
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.decimals"))
}

func testUpdateAssetChangeSubmissionWithoutTotalSupplyFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						TotalSupply: "",
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.total_supply"), commands.ErrIsRequired)
}

func testUpdateAssetChangeSubmissionWithTotalSupplySucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						TotalSupply: "10000",
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.total_supply"))
}

func testUpdateAssetChangeSubmissionWithNaNTotalSupplyFails(t *testing.T) {
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
					Change: &types.ProposalTerms_UpdateAsset{
						UpdateAsset: &types.UpdateAsset{
							Changes: &types.AssetDetailsUpdate{
								TotalSupply: tc.value,
							},
						},
					},
				},
			})

			assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.total_supply"), tc.error)
		})
	}
}

func testUpdateERC20AssetChangeSubmissionWithoutErc20AssetFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Source: &types.AssetDetailsUpdate_Erc20{},
					},
				},
			},
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.change.update_asset.changes.source.erc20"), commands.ErrIsRequired)
}

func testUpdateERC20AssetChangeSubmissionWithInvalidLifetimeLimitFails(t *testing.T) {
	tcs := []struct {
		name  string
		err   error
		value string
	}{
		{
			name:  "Without lifetime limit",
			value: "",
			err:   commands.ErrIsRequired,
		}, {
			name:  "With not-a-number lifetime limit",
			value: "forty-two",
			err:   commands.ErrIsNotValidNumber,
		}, {
			name:  "With zero lifetime limit",
			value: "0",
			err:   commands.ErrMustBePositive,
		}, {
			name:  "With negative lifetime limit",
			value: "-10",
			err:   commands.ErrMustBePositive,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					Change: &types.ProposalTerms_UpdateAsset{
						UpdateAsset: &types.UpdateAsset{
							Changes: &types.AssetDetailsUpdate{
								Source: &types.AssetDetailsUpdate_Erc20{
									Erc20: &types.ERC20Update{
										LifetimeLimit: tc.value,
									},
								},
							},
						},
					},
				},
			})

			assert.Contains(tt, err.Get("proposal_submission.terms.change.update_asset.changes.source.erc20.lifetime_limit"), tc.err)
		})
	}
}

func testUpdateERC20AssetChangeSubmissionWithValidLifetimeLimitSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Source: &types.AssetDetailsUpdate_Erc20{
							Erc20: &types.ERC20Update{
								LifetimeLimit: "100",
							},
						},
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.source.erc20.lifetime_limit"))
}

func testUpdateERC20AssetChangeSubmissionWithInvalidWithdrawalThresholdFails(t *testing.T) {
	tcs := []struct {
		name  string
		err   error
		value string
	}{
		{
			name:  "Without withdraw threshold",
			value: "",
			err:   commands.ErrIsRequired,
		}, {
			name:  "With not-a-number withdraw threshold",
			value: "forty-two",
			err:   commands.ErrIsNotValidNumber,
		}, {
			name:  "With zero withdraw threshold",
			value: "0",
			err:   commands.ErrMustBePositive,
		}, {
			name:  "With negative withdraw threshold",
			value: "-10",
			err:   commands.ErrMustBePositive,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					Change: &types.ProposalTerms_UpdateAsset{
						UpdateAsset: &types.UpdateAsset{
							Changes: &types.AssetDetailsUpdate{
								Source: &types.AssetDetailsUpdate_Erc20{
									Erc20: &types.ERC20Update{
										WithdrawThreshold: tc.value,
									},
								},
							},
						},
					},
				},
			})

			assert.Contains(tt, err.Get("proposal_submission.terms.change.update_asset.changes.source.erc20.withdraw_threshold"), tc.err)
		})
	}
}

func testUpdateERC20AssetChangeSubmissionWithValidWithdrawalThresholdSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			Change: &types.ProposalTerms_UpdateAsset{
				UpdateAsset: &types.UpdateAsset{
					Changes: &types.AssetDetailsUpdate{
						Source: &types.AssetDetailsUpdate_Erc20{
							Erc20: &types.ERC20Update{
								WithdrawThreshold: "100",
							},
						},
					},
				},
			},
		},
	})

	assert.Empty(t, err.Get("proposal_submission.terms.change.update_asset.changes.source.erc20.withdraw_threshold"))
}
