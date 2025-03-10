package gas

import (
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/store/models"

	"github.com/stretchr/testify/require"
)

func init() {
	// No need to wait 10 seconds in tests
	MaxStartTime = 1 * time.Second
}

func (b *BlockHistoryEstimator) CheckConnectivity(attempts []PriorAttempt) error {
	return b.checkConnectivity(attempts)
}

func BlockHistoryEstimatorFromInterface(bhe Estimator) *BlockHistoryEstimator {
	return bhe.(*BlockHistoryEstimator)
}

func SetRollingBlockHistory(bhe Estimator, blocks []Block) {
	bhe.(*BlockHistoryEstimator).blocksMu.Lock()
	defer bhe.(*BlockHistoryEstimator).blocksMu.Unlock()
	bhe.(*BlockHistoryEstimator).blocks = blocks
}

func GetRollingBlockHistory(bhe Estimator) []Block {
	return bhe.(*BlockHistoryEstimator).getBlocks()
}

func SetGasPrice(b *BlockHistoryEstimator, gp *assets.Wei) {
	b.priceMu.Lock()
	defer b.priceMu.Unlock()
	b.gasPrice = gp
}

func SetTipCap(b *BlockHistoryEstimator, gp *assets.Wei) {
	b.priceMu.Lock()
	defer b.priceMu.Unlock()
	b.tipCap = gp
}

func GetGasPrice(b *BlockHistoryEstimator) *assets.Wei {
	b.priceMu.RLock()
	defer b.priceMu.RUnlock()
	return b.gasPrice
}

func GetTipCap(b *BlockHistoryEstimator) *assets.Wei {
	b.priceMu.RLock()
	defer b.priceMu.RUnlock()
	return b.tipCap
}

func GetLatestBaseFee(b *BlockHistoryEstimator) *assets.Wei {
	b.latestMu.RLock()
	defer b.latestMu.RUnlock()
	if b.latest == nil {
		return nil
	}
	return b.latest.BaseFeePerGas
}

func SimulateStart(t *testing.T, b *BlockHistoryEstimator) {
	require.NoError(t, b.StartOnce("BlockHistoryEstimatorSimulatedStart", func() error { return nil }))
}

type MockConfig struct {
	BlockHistoryEstimatorBatchSizeF                 uint32
	BlockHistoryEstimatorBlockDelayF                uint16
	BlockHistoryEstimatorBlockHistorySizeF          uint16
	BlockHistoryEstimatorCheckInclusionBlocksF      uint16
	BlockHistoryEstimatorCheckInclusionPercentileF  uint16
	BlockHistoryEstimatorEIP1559FeeCapBufferBlocksF uint16
	BlockHistoryEstimatorTransactionPercentileF     uint16
	ChainTypeF                                      string
	DefaultHTTPTimeoutF                             models.Duration
	EvmEIP1559DynamicFeesF                          bool
	EvmGasBumpPercentF                              uint16
	EvmGasBumpThresholdF                            uint64
	EvmGasBumpWeiF                                  *assets.Wei
	EvmGasLimitMultiplierF                          float32
	EvmGasTipCapDefaultF                            *assets.Wei
	EvmGasTipCapMinimumF                            *assets.Wei
	EvmMaxGasPriceWeiF                              *assets.Wei
	EvmMinGasPriceWeiF                              *assets.Wei
	EvmGasPriceDefaultF                             *assets.Wei
}

func NewMockConfig() *MockConfig {
	return &MockConfig{}
}

func (m *MockConfig) BlockHistoryEstimatorBatchSize() uint32 {
	return m.BlockHistoryEstimatorBatchSizeF
}

func (m *MockConfig) BlockHistoryEstimatorBlockDelay() uint16 {
	return m.BlockHistoryEstimatorBlockDelayF
}

func (m *MockConfig) BlockHistoryEstimatorBlockHistorySize() uint16 {
	return m.BlockHistoryEstimatorBlockHistorySizeF
}

func (m *MockConfig) BlockHistoryEstimatorCheckInclusionPercentile() uint16 {
	return m.BlockHistoryEstimatorCheckInclusionPercentileF
}

func (m *MockConfig) BlockHistoryEstimatorCheckInclusionBlocks() uint16 {
	return m.BlockHistoryEstimatorCheckInclusionBlocksF
}

func (m *MockConfig) BlockHistoryEstimatorEIP1559FeeCapBufferBlocks() uint16 {
	return m.BlockHistoryEstimatorEIP1559FeeCapBufferBlocksF
}

func (m *MockConfig) BlockHistoryEstimatorTransactionPercentile() uint16 {
	return m.BlockHistoryEstimatorTransactionPercentileF
}

func (m *MockConfig) ChainType() config.ChainType {
	return config.ChainType(m.ChainTypeF)
}

func (m *MockConfig) DefaultHTTPTimeout() models.Duration {
	return m.DefaultHTTPTimeoutF
}

func (m *MockConfig) EvmEIP1559DynamicFees() bool {
	return m.EvmEIP1559DynamicFeesF
}

func (m *MockConfig) EvmFinalityDepth() uint32 {
	panic("not implemented") // TODO: Implement
}

func (m *MockConfig) EvmGasBumpPercent() uint16 {
	return m.EvmGasBumpPercentF
}

func (m *MockConfig) EvmGasBumpThreshold() uint64 {
	return m.EvmGasBumpThresholdF
}

func (m *MockConfig) EvmGasBumpWei() *assets.Wei {
	return m.EvmGasBumpWeiF
}

func (m *MockConfig) EvmGasFeeCapDefault() *assets.Wei {
	panic("not implemented") // TODO: Implement
}

func (m *MockConfig) EvmGasLimitMax() uint32 {
	panic("not implemented") // TODO: Implement
}

func (m *MockConfig) EvmGasLimitMultiplier() float32 {
	return m.EvmGasLimitMultiplierF
}

func (m *MockConfig) EvmGasPriceDefault() *assets.Wei {
	return m.EvmGasPriceDefaultF
}

func (m *MockConfig) EvmGasTipCapDefault() *assets.Wei {
	return m.EvmGasTipCapDefaultF
}

func (m *MockConfig) EvmGasTipCapMinimum() *assets.Wei {
	return m.EvmGasTipCapMinimumF
}

func (m *MockConfig) EvmMaxGasPriceWei() *assets.Wei {
	return m.EvmMaxGasPriceWeiF
}

func (m *MockConfig) EvmMinGasPriceWei() *assets.Wei {
	return m.EvmMinGasPriceWeiF
}

func (m *MockConfig) GasEstimatorMode() string {
	panic("not implemented") // TODO: Implement
}
