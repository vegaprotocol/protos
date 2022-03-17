# Changelog


## Unreleased (0.50.0)

### 🚨 Breaking changes
- [](https://github.com/vegaprotocol/protos/issues/) -

### 🗑️ Deprecation
- [](https://github.com/vegaprotocol/protos/pull/) -

### 🛠 Improvements
- [339](https://github.com/vegaprotocol/protos/pull/339) - Added position state message
- [341](https://github.com/vegaprotocol/protos/issues/341) - https://github.com/vegaprotocol/protos/issues/341
- [344](https://github.com/vegaprotocol/protos/issues/344) - Add the data structures for staking and `multisigcontrol` checkpoints
- [348](https://github.com/vegaprotocol/protos/issues/348) - Add candles V2 API
- [345](https://github.com/vegaprotocol/protos/issues/345) - Added proof of work to transaction

### 🐛 Fixes
- [](https://github.com/vegaprotocol/protos/pull/) -


## 0.49.2

### 🚨 Breaking changes
- [](https://github.com/vegaprotocol/protos/pull/) -

### 🗑️ Deprecation
- [](https://github.com/vegaprotocol/protos/pull/) -

### 🛠 Improvements
- [](https://github.com/vegaprotocol/protos/pull/) -

### 🐛 Fixes
- [](https://github.com/vegaprotocol/protos/pull/) -


## 0.49.1

### 🚨 Breaking changes
- [334](https://github.com/vegaprotocol/protos/issues/334) - Add ranking scores and reward score to node.


### 🛠 Improvements
- [325](https://github.com/vegaprotocol/protos/pull/325) - v2 data node API holds only new API methods
- [329](https://github.com/vegaprotocol/protos/pull/329) - remove unused request response types from v2 API
- [329](https://github.com/vegaprotocol/protos/pull/329) - Add proposal for market update and its validation.



## 0.49.0

### 🚨 Breaking changes
[298](https://github.com/vegaprotocol/protos/pull/298) - Add support for fractional order sizes

### 🛠 Improvements
- [285](https://github.com/vegaprotocol/protos/pull/285) - Update snapshot protos to store enough data to restore `corestate` service
- [296](https://github.com/vegaprotocol/protos/pull/296) - Remove old id generator fields from execution engine's snapshot
- [293](https://github.com/vegaprotocol/protos/pull/293) - Re-enable proto check
- [306](https://github.com/vegaprotocol/protos/pull/306) - Add initial v2 orders API
- [312](https://github.com/vegaprotocol/protos/pull/306) - Add rpc call to get network limits

### 🐛 Fixes
- [308](https://github.com/vegaprotocol/protos/pull/308) - Validate order price and fail if not positive
- [310](https://github.com/vegaprotocol/protos/pull/310) - Add `feesplitter` state to snapshot to allow markets to restore
- [319](https://github.com/vegaprotocol/protos/pull/319) - Fix incorrect data types on `MarketData` message definition

## 0.48.0

### 🚨 Breaking changes
- [314](https://github.com/vegaprotocol/protos/pull/314) - Add more data to submit transaction endpoints
- [218](https://github.com/vegaprotocol/protos/pull/) - Restructure EthereumConfig to separate staking and vesting contract addresses, plus add block height at which they have been added respectively
- [256](https://github.com/vegaprotocol/protos/pull/) - Rework freeform proposal protos so that they align with other proposals
- [290](https://github.com/vegaprotocol/protos/pull/290) - Generate stubs using new versions of `protoc-gen-xxx`

### 🗑️ Deprecation
- [](https://github.com/vegaprotocol/protos/pull/) -

### 🛠 Improvements
- [303](https://github.com/vegaprotocol/protos/pull/303) - Add types for CheckTransaction and CheckRawTransaction for vega
- [191](https://github.com/vegaprotocol/protos/pull/191) - Add details on transaction nonce and block_height documentation
- [196](https://github.com/vegaprotocol/protos/pull/199) - Get rid of risk result and change risk factors to string
- [212](https://github.com/vegaprotocol/protos/pull/212) - Add pagination field to DelegationRequest
- [213](https://github.com/vegaprotocol/protos/pull/217) - Add changes for ValidatorPerformance snapshot and events
- [219](https://github.com/vegaprotocol/protos/pull/219) - Include floating point consensus flags in the snapshot for a market
- [189](https://github.com/vegaprotocol/protos/pull/189) - Update offset to string and disable negative offset for buy side
- [225](https://github.com/vegaprotocol/protos/pull/225) - Added snapshot payload for floating point consensus engine
- [144](https://github.com/vegaprotocol/protos/pull/144) - Remove validation annotations
- [228](https://github.com/vegaprotocol/protos/pull/228) - Raw Score and Performance added to Node definition for validator performance reporting
- [230](https://github.com/vegaprotocol/protos/pull/230) - Implement the transfer commands
- [230](https://github.com/vegaprotocol/protos/pull/230) - Change the transfer command to have recurring transfer end epoch optional
- [236](https://github.com/vegaprotocol/protos/pull/236) - Implement CancelTransfer command
- [242](https://github.com/vegaprotocol/protos/pull/242) - Add Submit Raw Transaction endpoint and supporting messages
- [244](https://github.com/vegaprotocol/protos/pull/244) - Add account types and snapshot definitions for market and fee trackers
- [240](https://github.com/vegaprotocol/protos/pull/240) - Add transfers snapshot types
- [248](https://github.com/vegaprotocol/protos/pull/248) - Add transaction hash to events
- [250](https://github.com/vegaprotocol/protos/pull/250) - Rename transaction hash to tx hash
- [253](https://github.com/vegaprotocol/protos/pull/253) - Allow `OracleSpec` for internal oracle to be defined without public keys
- [261](https://github.com/vegaprotocol/protos/pull/261) - Add `staking_asset_total_supply` to staking accounts snapshot
- [266](https://github.com/vegaprotocol/protos/pull/266) - Improve validator score event and add validator ranking event and proto definitions for topology snapshot
- [263](https://github.com/vegaprotocol/protos/pull/263) - Add new validator commands
- [272](https://github.com/vegaprotocol/protos/pull/272) - Update snapshot protos to be able to handle the iavl exported nodes
- [276](https://github.com/vegaprotocol/protos/pull/276) - Remove maturity from future product, tick size from market and trading mode config from market
- [277](https://github.com/vegaprotocol/protos/pull/277) - Update validation for risk params for log normal risk model
- [287](https://github.com/vegaprotocol/protos/pull/287) - Add `QueryBalanceHistory` call and associated messages
- [300](https://github.com/vegaprotocol/protos/pull/300) - Add market data endpoints for trading data service version 2

### 🐛 Fixes
- [202](https://github.com/vegaprotocol/protos/pull/202) - Replaces Withdrawal Status Cancelled with Rejected which is more accurate.
- [214](https://github.com/vegaprotocol/protos/pull/214) - Rewording of transaction nonce and block_height


## 0.47.0
*2021-12-09*


### 🛠 Improvements
- [180](https://github.com/vegaprotocol/protos/pull/180) - Update `CHANGELOG.md` since GH Action implemented
- [162](https://github.com/vegaprotocol/protos/pull/162) - Update changelog for `v0.46.0`
- [159](https://github.com/vegaprotocol/protos/pull/159) - Update command and add transaction - Key rotate submission
- [164](https://github.com/vegaprotocol/protos/pull/164) - Remove Stream Start, add `ChainID` to header
- [165](https://github.com/vegaprotocol/protos/pull/165) - Update key rotate command and add key rotate event
- [167](https://github.com/vegaprotocol/protos/pull/167) - Update validator function
- [170](https://github.com/vegaprotocol/protos/pull/170) - Tidy up repo to align with team processes and workflows
- [171](https://github.com/vegaprotocol/protos/pull/171) - Update topology snapshot to reflect key rotations
- [175](https://github.com/vegaprotocol/protos/pull/175) - Add data node key rotations API
- [177](https://github.com/vegaprotocol/protos/pull/177) - Add key rotations checkpoint
- [186](https://github.com/vegaprotocol/protos/pull/186) - Proto definitions for floating point state variable
- [195](https://github.com/vegaprotocol/protos/pull/195) - Add network limits bus events & GRPC calls
- [188](https://github.com/vegaprotocol/protos/pull/188) - Add support for filtering & paginating rewards
- [196](https://github.com/vegaprotocol/protos/pull/196) - Add propose_xxx_enabled_from to network limits event
- [201](https://github.com/vegaprotocol/protos/pull/201) - Add chain event for stake total supply
- [207](https://github.com/vegaprotocol/protos/pull/207) - Adding new commands for Liquidity Provision Cancellation and Amendment

### 🐛 Fixes
- [173](https://github.com/vegaprotocol/protos/pull/173) - Rename node registration field
- [210](https://github.com/vegaprotocol/protos/pull/210) - Remove redundant Id field from Liquidity Provision Cancellation and Amendment messages.

## 0.46.0
*2021-11-22*

### 🛠 Improvements
- [158](https://github.com/vegaprotocol/protos/pull/158) - Release version `v0.46.0`
- [157](https://github.com/vegaprotocol/protos/pull/157) - Release version `v0.46.0`
- [155](https://github.com/vegaprotocol/protos/pull/155) - Add a checkpoint proto definitions for pending rewards
- [154](https://github.com/vegaprotocol/protos/pull/154) - Add Hello message
- [151](https://github.com/vegaprotocol/protos/pull/151) - Add changelog, linked PR and project board Github actions
- [135](https://github.com/vegaprotocol/protos/pull/135) - Add snapshot protos for topology
- [137](https://github.com/vegaprotocol/protos/pull/137) - Add new prices fields to price monitor
- [138](https://github.com/vegaprotocol/protos/pull/138) - Add oracle snapshot proto model
- [128](https://github.com/vegaprotocol/protos/pull/128) - Add liquidity snapshot types
- [141](https://github.com/vegaprotocol/protos/pull/141) - Add liquidity target snapshot
- [140](https://github.com/vegaprotocol/protos/pull/140) - Add proto definition for liquidity/supplied snapshot
- [145](https://github.com/vegaprotocol/protos/pull/145) - Add protos for freeform proposals
- [148](https://github.com/vegaprotocol/protos/pull/148) - Add TX Hash to SubmitTransactionResponse proto message
- [149](https://github.com/vegaprotocol/protos/pull/149) - Add transaction hash to SubmitTransactionResponse
- [147](https://github.com/vegaprotocol/protos/pull/147) - Add observers for delegation and rewards
- [146](https://github.com/vegaprotocol/protos/pull/146) - Add NewFreeform as a new proposal type
- [96](https://github.com/vegaprotocol/protos/pull/96) - Add key rotate submission command
- [166](https://github.com/vegaprotocol/protos/pull/166) - Update key rotate submission command and add key rotate event

### 🐛 Fixes
- [139](https://github.com/vegaprotocol/protos/pull/139) - Fix oracle data model
- [81](https://github.com/vegaprotocol/protos/pull/81) - Fix broken markdown in README
- [167](https://github.com/vegaprotocol/protos/pull/167) - Update key rotation validation function
- [173](https://github.com/vegaprotocol/protos/pull/173) - Rename node registration field
- [181](https://github.com/vegaprotocol/protos/pull/181) - Fix type for key rotations checkpoint proto

## 0.45.1
*2021-10-23*

### 🛠 Improvements
- [124](https://github.com/vegaprotocol/protos/pull/124) - Add event fwder snapshot model
- [126](https://github.com/vegaprotocol/protos/pull/126) - Add ID generator fields to execution type
- [130](https://github.com/vegaprotocol/protos/pull/130) - Add snapshot protos for stake verifier pending stakes
- [129](https://github.com/vegaprotocol/protos/pull/129) - data model for witness snapshot
- [132](https://github.com/vegaprotocol/protos/pull/132) - added payload for reconciliation time for delegation


## 0.45.0
*2021-10-19*

### 🛠 Improvements
- [114](https://github.com/vegaprotocol/protos/pull/114) - Restructure rewards snapshot model
- [117](https://github.com/vegaprotocol/protos/pull/117) - Add limit_state to snapshot protos
- [118](https://github.com/vegaprotocol/protos/pull/118) - Add spam snapshot data model
- [120](https://github.com/vegaprotocol/protos/pull/120) - Add proto definitions for notary snapshots
- [121](https://github.com/vegaprotocol/protos/pull/121) - Add replay protection snapshot proto
- [122](https://github.com/vegaprotocol/protos/pull/122) - Add protos for futures product
- [123](https://github.com/vegaprotocol/protos/pull/123) - Add snapshot proto definitions for future oracle data
- [80](https://github.com/vegaprotocol/protos/pull/80) - Converted int64 to strings
- [125](https://github.com/vegaprotocol/protos/pull/125) - Release version `v0.45.0`


## 0.44.0
*2021-10-07*

### 🛠 Improvements
- [86](https://github.com/vegaprotocol/protos/pull/86) - Remove the datanode trading proxy
- [85](https://github.com/vegaprotocol/protos/pull/85) - Add snapshot types
- [87](https://github.com/vegaprotocol/protos/pull/87) - Add auto-delegation data to checkpoint
- [89](https://github.com/vegaprotocol/protos/pull/89) - Added new message for EpochState
- [98](https://github.com/vegaprotocol/protos/pull/98) - Revert "epoch snapshot protos"
- [100](https://github.com/vegaprotocol/protos/pull/100) - Add collateral asset snapshot
- [99](https://github.com/vegaprotocol/protos/pull/99) - Add back the epoch protos message for snapshotting
- [103](https://github.com/vegaprotocol/protos/pull/103) - Cleanup package layout
- [94](https://github.com/vegaprotocol/protos/pull/94) - Added missing matching book fields
- [107](https://github.com/vegaprotocol/protos/pull/107) - add new fields to events and commands for node / validators
- [104](https://github.com/vegaprotocol/protos/pull/104) - Added IDGenerator message
- [108](https://github.com/vegaprotocol/protos/pull/108) - Add global reward pool account endpoint
- [109](https://github.com/vegaprotocol/protos/pull/109) - Remove delegation from party message
- [110](https://github.com/vegaprotocol/protos/pull/110) - Added asset actions to banking snapshot and added reward snapshot proto
- [113](https://github.com/vegaprotocol/protos/pull/113) - Add node-proposals to governance snapshot protos
- [115](https://github.com/vegaprotocol/protos/pull/115) - Release version `v0.44.0`

### 🐛 Fixes
- [91](https://github.com/vegaprotocol/protos/pull/91) - Fix `GRPC->REST` binding for SubmitTransaction


## 0.43.0
*2021-10-01*

### 🛠 Improvements
- [83](https://github.com/vegaprotocol/protos/pull/83) - Add `ORACLE_SOURCE_JSON` to OracleSource
- [106](https://github.com/vegaprotocol/protos/pull/106) - Release version `v0.43.0`

### 🐛 Fixes
- [78](https://github.com/vegaprotocol/protos/pull/78) - Change uint64 to string
- [84](https://github.com/vegaprotocol/protos/pull/84) - Fix epochs


## 0.42.0
*2021-09-10*

### 🛠 Improvements
- [15](https://github.com/vegaprotocol/protos/pull/15) - Remove all prepare endpoints
- [17](https://github.com/vegaprotocol/protos/pull/17) - Add types needed for checkpoints
- [21](https://github.com/vegaprotocol/protos/pull/21) - Remove `TXv1`
- [23](https://github.com/vegaprotocol/protos/pull/23) - feature/3726 event for reward payout
- [26](https://github.com/vegaprotocol/protos/pull/26) - Add rewards API messages
- [28](https://github.com/vegaprotocol/protos/pull/28) - Add validators API protos
- [30](https://github.com/vegaprotocol/protos/pull/30) - Add partyID to rewards event
- [33](https://github.com/vegaprotocol/protos/pull/33) - Remove all prepare functions
- [37](https://github.com/vegaprotocol/protos/pull/37) - Add REST endpoint for rewards information
- [38](https://github.com/vegaprotocol/protos/pull/38) - Add checkpoint event and data-node API updates
- [39](https://github.com/vegaprotocol/protos/pull/39) - Add bus event type enum
- [18](https://github.com/vegaprotocol/protos/pull/18) - Expose vega and ethereum pub keys in NodeRegistration
- [40](https://github.com/vegaprotocol/protos/pull/40) - Add proto definitions for delegation
- [41](https://github.com/vegaprotocol/protos/pull/41) - Move delegations APIs into the datanode tree
- [43](https://github.com/vegaprotocol/protos/pull/43) - Add reward history
- [45](https://github.com/vegaprotocol/protos/pull/45) - Removed account details and added a time field to reward message
- [34](https://github.com/vegaprotocol/protos/pull/34) - Add reference checks on proposal
- [1](https://github.com/vegaprotocol/protos/pull/1) - Remove duplicated Transaction helpers
- [48](https://github.com/vegaprotocol/protos/pull/48) - Add NewSignature func
- [50](https://github.com/vegaprotocol/protos/pull/50) - Add checkpoint restore command
- [49](https://github.com/vegaprotocol/protos/pull/49) - Add REST bindings for Validators API
- [51](https://github.com/vegaprotocol/protos/pull/51) - Generate latest develop
- [22](https://github.com/vegaprotocol/protos/pull/22) - Add staking event to chain event
- [54](https://github.com/vegaprotocol/protos/pull/54) - Add delegations REST
- [53](https://github.com/vegaprotocol/protos/pull/53) - Remove fee
- [56](https://github.com/vegaprotocol/protos/pull/56) - Add delegation types for checkpoints
- [57](https://github.com/vegaprotocol/protos/pull/57) - Added an event for validator score for rewarding
- [55](https://github.com/vegaprotocol/protos/pull/55) - Remove vega APIs
- [59](https://github.com/vegaprotocol/protos/pull/59) - Add node score
- [58](https://github.com/vegaprotocol/protos/pull/58) - Add target address to ERC20Withdrawal approval
- [60](https://github.com/vegaprotocol/protos/pull/60) - Add partiers stake API
- [62](https://github.com/vegaprotocol/protos/pull/62) - Add the delegations command to the commands checks
- [61](https://github.com/vegaprotocol/protos/pull/61) - Add delegation / undelegation / restor snapshot tx err
- [63](https://github.com/vegaprotocol/protos/pull/63) - Add support for restore command
- [31](https://github.com/vegaprotocol/protos/pull/31) - Move most of uint64 to strings
- [65](https://github.com/vegaprotocol/protos/pull/65) - Add epoch expiry and action
- [67](https://github.com/vegaprotocol/protos/pull/67) - Add node id to validator update event
- [69](https://github.com/vegaprotocol/protos/pull/69) - Move last block height to proxy
- [72](https://github.com/vegaprotocol/protos/pull/72) - Change delegation balance to string
- [73](https://github.com/vegaprotocol/protos/pull/73) - Feature/add delegations core apis
- [74](https://github.com/vegaprotocol/protos/pull/74) - Add epoch field to the checkpoint proto
- [75](https://github.com/vegaprotocol/protos/pull/75) - Add version to events
- [76](https://github.com/vegaprotocol/protos/pull/76) - Add block field
- [77](https://github.com/vegaprotocol/protos/pull/77) - Release version `v0.42.0`

### 🐛 Fixes
- [19](https://github.com/vegaprotocol/protos/pull/19) - Fix NewMarketSubmission with 0 value as decimal place
- [24](https://github.com/vegaprotocol/protos/pull/24) - Fix proto add missing bus definition
- [52](https://github.com/vegaprotocol/protos/pull/52) - Corrected asset typeFix REST API generation
- [47](https://github.com/vegaprotocol/protos/pull/47) - Corrected asset type


## 0.41.0
*2021-08-09*

### 🛠 Improvements
- [2](https://github.com/vegaprotocol/protos/pull/) - Add data node updates
- [7](https://github.com/vegaprotocol/protos/pull/) - Add license and readme
- [5](https://github.com/vegaprotocol/protos/pull/) - Add epoch message
- [11](https://github.com/vegaprotocol/protos/pull/) - Add StakingEvent
- [10](https://github.com/vegaprotocol/protos/pull/) - Create Jenkinsfile
- [12](https://github.com/vegaprotocol/protos/pull/12) - Release version `v0.41.0`
