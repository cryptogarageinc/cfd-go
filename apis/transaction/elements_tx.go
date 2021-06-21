package transaction

import (
	"fmt"
	"strings"
	"unsafe"

	cfd "github.com/cryptogarageinc/cfd-go"
	"github.com/cryptogarageinc/cfd-go/apis/address"
	"github.com/cryptogarageinc/cfd-go/apis/descriptor"
	"github.com/cryptogarageinc/cfd-go/config"
	"github.com/cryptogarageinc/cfd-go/types"
	"github.com/pkg/errors"
)

const (
	EmptyBlinder string = types.EmptyBlinder
)

type ConfidentialTxApi interface {
	Create(version uint32, locktime uint32, txinList *[]types.InputConfidentialTxIn, txoutList *[]types.InputConfidentialTxOut, pegoutAddressList *[]string) (tx *types.ConfidentialTx, err error)
	Add(tx *types.ConfidentialTx, txinList *[]types.InputConfidentialTxIn, txoutList *[]types.InputConfidentialTxOut, pegoutAddressList *[]string) error
	Blind(tx *types.ConfidentialTx, txinList []types.BlindInputData, txoutList *[]types.BlindOutputData, option *types.BlindTxOption) error
	AddPubkeySign(tx *types.ConfidentialTx, outpoint *types.OutPoint, hashType types.HashType, pubkey *types.Pubkey, signature string) error
	AddPubkeySignByDescriptor(tx *types.ConfidentialTx, outpoint *types.OutPoint, outputDescriptor *types.Descriptor, signature string) error
	VerifySign(tx *types.ConfidentialTx, outpoint *types.OutPoint, txinUtxoList *[]types.UtxoData) (isVerify bool, reason string, err error)
	GetTxid(tx *types.ConfidentialTx) string
	GetPegoutAddress(tx *types.ConfidentialTx, index uint32) (pegoutAddress *types.Address, isPegoutOutput bool, err error)
	GetSighash(tx *types.ConfidentialTx, outpoint *types.OutPoint, sighashType types.SigHashType, utxoList *[]types.UtxoData) (sighash *types.ByteData, err error)
	GetAll(tx *types.ConfidentialTx, hasWitness bool) (data *types.TransactionData, txinList []types.ConfidentialTxIn, txoutList []types.ConfidentialTxOut, err error)
	GetTxIn(txHex string, outpoint *types.OutPoint) (txin *types.ConfidentialTxIn, err error)
}

func NewConfidentialTxApi() *ConfidentialTxApiImpl {
	return &ConfidentialTxApiImpl{}
}

// -------------------------------------
// API struct
// -------------------------------------

// ConfidentialTxApiImpl Create confidential transaction utility.
type ConfidentialTxApiImpl struct {
	Network *types.NetworkType
}

// -------------------------------------
// ConfidentialTxApiImpl
// -------------------------------------

func (u *ConfidentialTxApiImpl) Create(version uint32, locktime uint32, txinList *[]types.InputConfidentialTxIn, txoutList *[]types.InputConfidentialTxOut, pegoutAddressList *[]string) (tx *types.ConfidentialTx, err error) {
	if err = u.validConfig(); err != nil {
		return nil, err
	}
	txHandle, err := cfd.InitializeTransaction(u.Network.ToCfdValue(), version, locktime)
	if err != nil {
		return nil, err
	}
	defer cfd.FreeTransactionHandle(txHandle)

	if err = addConidentialTx(txHandle, *u.Network, locktime, txinList, txoutList, pegoutAddressList); err != nil {
		return nil, err
	}

	txHex, err := cfd.FinalizeTransaction(txHandle)
	if err != nil {
		return nil, err
	}
	if txoutList != nil {
		txHex, err = updateDirectNonce(txHandle, txHex, txoutList)
		if err != nil {
			return nil, err
		}
	}
	return &types.ConfidentialTx{Hex: txHex}, nil
}

func (u *ConfidentialTxApiImpl) validConfig() error {
	if u.Network == nil {
		cfdConfig := config.GetCurrentCfdConfig()
		if !cfdConfig.Network.Valid() {
			return fmt.Errorf("CFD Error: NetworkType not set")
		}
		netType := cfdConfig.Network
		u.Network = &netType
	}
	if !u.Network.IsElements() {
		return fmt.Errorf("CFD Error: NetworkType is not elements")
	}
	return nil
}

func (t *ConfidentialTxApiImpl) Add(tx *types.ConfidentialTx, txinList *[]types.InputConfidentialTxIn, txoutList *[]types.InputConfidentialTxOut, pegoutAddressList *[]string) error {
	if err := t.validConfig(); err != nil {
		return err
	}
	txHandle, err := cfd.InitializeTransactionByHex(t.Network.ToCfdValue(), tx.Hex)
	if err != nil {
		return err
	}
	defer cfd.FreeTransactionHandle(txHandle)

	data, err := cfd.CfdGoGetConfidentialTxDataByHandle(txHandle)
	if err != nil {
		return err
	}

	if err = addConidentialTx(txHandle, *t.Network, data.LockTime, txinList, txoutList, pegoutAddressList); err != nil {
		return err
	}

	txHex, err := cfd.FinalizeTransaction(txHandle)
	if err != nil {
		return err
	}
	if txoutList != nil {
		txHex, err = updateDirectNonce(txHandle, txHex, txoutList)
		if err != nil {
			return err
		}
	}
	tx.Hex = txHex
	return nil
}

//func (t *ConfidentialTx) SetReissueAsset() error {
// FIXME need implements
//}

// Blind ...
func (t *ConfidentialTxApiImpl) Blind(tx *types.ConfidentialTx, txinList []types.BlindInputData, txoutList *[]types.BlindOutputData, option *types.BlindTxOption) error {
	lbtcAsset, _ := getDefaultBitcoinData()

	var err error
	if err = t.validConfig(); err != nil {
		return err
	}

	txHex := tx.Hex
	if option != nil && option.AppendDummyOutput {
		txHex, err = appendDummyOutput(txHex, *t.Network, &txinList)
		if err != nil {
			return err
		}
	}
	// convert list
	blindTxinList := make([]cfd.CfdBlindInputData, len(txinList))
	for index, input := range txinList {
		blindTxinList[index].Txid = input.OutPoint.Txid
		blindTxinList[index].Vout = input.OutPoint.Vout
		blindTxinList[index].Amount = input.Amount
		blindTxinList[index].ValueBlindFactor = input.ValueBlindFactor
		blindTxinList[index].Asset = input.Asset
		blindTxinList[index].AssetBlindFactor = input.AssetBlindFactor
		if input.IssuanceKey != nil {
			blindTxinList[index].AssetBlindingKey = input.IssuanceKey.AssetBlindingKey
			blindTxinList[index].TokenBlindingKey = input.IssuanceKey.TokenBlindingKey
		}
		if len(blindTxinList[index].Asset) == 0 {
			blindTxinList[index].Asset = lbtcAsset
		}
	}
	var blindOutputList []cfd.CfdBlindOutputData
	if txoutList != nil {
		blindOutputList = *(*[]cfd.CfdBlindOutputData)(unsafe.Pointer(txoutList))
	}
	blindOption := (*cfd.CfdBlindTxOption)(unsafe.Pointer(option))
	outputTx, err := cfd.CfdGoBlindRawTransaction(txHex, blindTxinList, blindOutputList, blindOption)
	if err != nil {
		return err
	}
	tx.Hex = outputTx
	return nil
}

// AddPubkeySign ...
func (t *ConfidentialTxApiImpl) AddPubkeySign(tx *types.ConfidentialTx, outpoint *types.OutPoint, hashType types.HashType, pubkey *types.Pubkey, signature string) error {
	if err := t.validConfig(); err != nil {
		return err
	}
	signParam := cfd.CfdSignParameter{
		Data:                signature,
		IsDerEncode:         false,
		SighashType:         int(cfd.KCfdSigHashAll),
		SighashAnyoneCanPay: false,
	}
	txHex, err := cfd.CfdGoAddTxPubkeyHashSign(t.Network.ToCfdValue(), tx.Hex, outpoint.Txid, outpoint.Vout, hashType.ToCfdValue(), pubkey.Hex, signParam)
	if err != nil {
		return errors.Wrap(err, "CT.AddPubkeySign error")
	}
	tx.Hex = txHex
	return nil
}

// AddPubkeySignByDescriptor ...
func (t *ConfidentialTxApiImpl) AddPubkeySignByDescriptor(tx *types.ConfidentialTx, outpoint *types.OutPoint, outputDescriptor *types.Descriptor, signature string) error {
	var err error
	if err = t.validConfig(); err != nil {
		return err
	}
	descUtil := descriptor.DescriptorApiImpl{Network: t.Network}
	data, _, _, err := descUtil.Parse(outputDescriptor)
	if err != nil {
		return err
	}
	if data.HashType != int(cfd.KCfdP2pkh) && data.HashType != int(cfd.KCfdP2wpkh) && data.HashType != int(cfd.KCfdP2shP2wpkh) {
		return fmt.Errorf("CFD Error: Descriptor hashType is not pubkeyHash")
	}

	hashType := types.NewHashType(data.HashType)
	var pubkey types.Pubkey
	if data.KeyType == int(cfd.KCfdDescriptorKeyPublic) {
		pubkey.Hex = data.Pubkey
	} else if data.KeyType == int(cfd.KCfdDescriptorKeyBip32) {
		pubkey.Hex, err = cfd.CfdGoGetPubkeyFromExtkey(data.ExtPubkey, t.Network.ToBitcoinType().ToCfdValue())
		if err != nil {
			return err
		}
	} else if data.KeyType == int(cfd.KCfdDescriptorKeyBip32Priv) {
		pubkey.Hex, err = cfd.CfdGoGetPubkeyFromExtkey(data.ExtPrivkey, t.Network.ToBitcoinType().ToCfdValue())
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("CFD Error: Descriptor keyType is not pubkeyHash")
	}
	return t.AddPubkeySign(tx, outpoint, hashType, &pubkey, signature)
}

// VerifySign ...
func (t *ConfidentialTxApiImpl) VerifySign(tx *types.ConfidentialTx, outpoint *types.OutPoint, txinUtxoList *[]types.UtxoData) (isVerify bool, reason string, err error) {
	if err := t.validConfig(); err != nil {
		return false, "", err
	}
	lbtcAsset, _ := getDefaultBitcoinData()
	utxoList := (*[]cfd.CfdUtxo)(unsafe.Pointer(txinUtxoList))
	for i, utxo := range *utxoList {
		if len(utxo.Asset) == 0 {
			(*utxoList)[i].Asset = lbtcAsset
		}
	}
	return cfd.CfdGoVerifySign(t.Network.ToCfdValue(), tx.Hex, *utxoList, outpoint.Txid, outpoint.Vout)
}

// GetTxid ...
func (t *ConfidentialTxApiImpl) GetTxid(tx *types.ConfidentialTx) string {
	if err := t.validConfig(); err != nil {
		return ""
	}
	handle, err := cfd.CfdGoInitializeTxDataHandle(t.Network.ToCfdValue(), tx.Hex)
	if err != nil {
		return ""
	}
	defer cfd.CfdGoFreeTxDataHandle(handle)

	data, err := cfd.CfdGoGetTxInfoByHandle(handle)
	if err != nil {
		return ""
	}
	return data.Txid
}

// GetPegoutAddress ...
func (t *ConfidentialTxApiImpl) GetPegoutAddress(tx *types.ConfidentialTx, index uint32) (pegoutAddress *types.Address, isPegoutOutput bool, err error) {
	if err := t.validConfig(); err != nil {
		return nil, false, err
	}
	addr, isPegoutOutput, err := cfd.CfdGoGetPegoutAddressFromTransaction(t.Network.ToCfdValue(), tx.Hex, index, t.Network.ToBitcoinType().ToCfdValue())
	if err != nil {
		return nil, false, err
	}
	btcNetwork := t.Network.ToBitcoinType()
	addrUtil := address.AddressApiImpl{Network: &btcNetwork}
	pegoutAddress, err = addrUtil.ParseAddress(addr)
	if err != nil {
		return nil, false, err
	}
	return pegoutAddress, isPegoutOutput, nil
}

// GetSighash ...
func (t *ConfidentialTxApiImpl) GetSighash(tx *types.ConfidentialTx, outpoint *types.OutPoint, sighashType types.SigHashType, utxoList *[]types.UtxoData) (sighash *types.ByteData, err error) {
	if err := t.validConfig(); err != nil {
		return nil, err
	}
	cfdNetType := t.Network.ToCfdValue()
	descUtil := descriptor.DescriptorApiImpl{Network: t.Network}
	var script *cfd.Script
	var pubkey *cfd.ByteData
	for _, utxo := range *utxoList {
		if utxo.Txid == outpoint.Txid && utxo.Vout == outpoint.Vout {
			desc := descUtil.NewDescriptorFromString(utxo.Descriptor)
			if desc == nil {
				return nil, fmt.Errorf("CFD Error: Invalid descriptor string")
			}
			data, _, _, err := descUtil.Parse(desc)
			if err != nil {
				return nil, err
			}

			if data.HashType != int(cfd.KCfdP2pkh) && data.HashType != int(cfd.KCfdP2wpkh) && data.HashType != int(cfd.KCfdP2shP2wpkh) {
				return nil, fmt.Errorf("CFD Error: Descriptor hashType is not pubkeyHash")
			}
			if data.KeyType == int(cfd.KCfdDescriptorKeyPublic) {
				pubkey = cfd.NewByteDataFromHexIgnoreError(data.Pubkey)
			} else if data.KeyType == int(cfd.KCfdDescriptorKeyBip32) {
				tempPubkey, err := cfd.CfdGoGetPubkeyFromExtkey(data.ExtPubkey, t.Network.ToBitcoinType().ToCfdValue())
				if err != nil {
					return nil, err
				}
				pubkey = cfd.NewByteDataFromHexIgnoreError(tempPubkey)
			} else if data.KeyType == int(cfd.KCfdDescriptorKeyBip32Priv) {
				tempPubkey, err := cfd.CfdGoGetPubkeyFromExtkey(data.ExtPrivkey, t.Network.ToBitcoinType().ToCfdValue())
				if err != nil {
					return nil, err
				}
				pubkey = cfd.NewByteDataFromHexIgnoreError(tempPubkey)
			} else if len(data.RedeemScript) != 0 {
				script = cfd.NewScriptFromHexIgnoreError(data.RedeemScript)
			} else {
				return nil, fmt.Errorf("CFD Error: Descriptor invalid")
			}
			break
		}
	}
	lbtcAsset, _ := getDefaultBitcoinData()
	txinUtxoList := (*[]cfd.CfdUtxo)(unsafe.Pointer(utxoList))
	for i, utxo := range *txinUtxoList {
		if len(utxo.Asset) == 0 {
			(*txinUtxoList)[i].Asset = lbtcAsset
		}
	}
	cfdSighashType := cfd.SigHashType{
		Type:         sighashType.Type,
		AnyoneCanPay: sighashType.AnyoneCanPay,
		Rangeproof:   sighashType.Rangeproof,
	}
	sighashHex, err := cfd.CfdGoGetSighash(cfdNetType, tx.Hex, *txinUtxoList, outpoint.Txid, outpoint.Vout, &cfdSighashType, pubkey, script, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return types.NewByteDataFromHexIgnoreError(sighashHex), nil
}

// GetAll ...
func (t *ConfidentialTxApiImpl) GetAll(tx *types.ConfidentialTx, hasWitness bool) (data *types.TransactionData, txinList []types.ConfidentialTxIn, txoutList []types.ConfidentialTxOut, err error) {
	if err := t.validConfig(); err != nil {
		return data, txinList, txoutList, err
	}
	cfdData, cfdTxins, cfdTxouts, err := cfd.GetConfidentialTxData(tx.Hex, hasWitness)
	if err != nil {
		return nil, nil, nil, err
	}
	return convertListData(&cfdData, cfdTxins, cfdTxouts)
}

func (t *ConfidentialTxApiImpl) GetTxIn(txHex string, outpoint *types.OutPoint) (txin *types.ConfidentialTxIn, err error) {
	handle, err := cfd.CfdGoInitializeTxDataHandle(types.LiquidV1.ToCfdValue(), txHex)
	if err != nil {
		return nil, err
	}
	defer cfd.CfdGoFreeTxDataHandle(handle)

	var tempTxin types.ConfidentialTxIn
	index, err := cfd.CfdGoGetTxInIndexByHandle(handle, outpoint.Txid, outpoint.Vout)
	txid, vout, sequence, scriptSig, err := cfd.CfdGoGetTxInByHandle(handle, index)
	if err != nil {
		return nil, err
	}
	entropy, nonce, assetAmount, assetValue, tokenAmount, tokenValue, assetRangeproof, tokenRangeproof, err := cfd.CfdGoGetTxInIssuanceInfoByHandle(handle, index)
	if err != nil {
		return nil, err
	}

	tempTxin.OutPoint.Txid = txid
	tempTxin.OutPoint.Vout = vout
	tempTxin.Sequence = sequence
	tempTxin.ScriptSig = scriptSig
	if len(assetValue) > 2 {
		tempTxin.Issuance.Entropy = entropy
		tempTxin.Issuance.Nonce = nonce
		tempTxin.Issuance.AssetAmount = assetAmount
		tempTxin.Issuance.AssetValue = assetValue
		tempTxin.Issuance.TokenAmount = tokenAmount
		tempTxin.Issuance.TokenValue = tokenValue
		tempTxin.IssuanceAmountRangeproof = assetRangeproof
		tempTxin.InflationKeysRangeproof = tokenRangeproof
	}

	wCount, err := cfd.CfdGoGetTxInWitnessCountByHandle(handle, int(cfd.KCfdTxWitnessStackNormal), index)
	if err != nil {
		return nil, err
	}
	wList := make([]string, wCount)
	for j := uint32(0); j < wCount; j++ {
		stackData, err := cfd.CfdGoGetTxInWitnessByHandle(handle, int(cfd.KCfdTxWitnessStackNormal), index, j)
		if err != nil {
			return nil, err
		}
		wList[j] = stackData
	}
	tempTxin.WitnessStack.Stack = wList

	pCount, err := cfd.CfdGoGetTxInWitnessCountByHandle(handle, int(cfd.KCfdTxWitnessStackPegin), index)
	if err != nil {
		return nil, err
	}
	pList := make([]string, pCount)
	for j := uint32(0); j < pCount; j++ {
		stackData, err := cfd.CfdGoGetTxInWitnessByHandle(handle, int(cfd.KCfdTxWitnessStackPegin), index, j)
		if err != nil {
			return nil, err
		}
		pList[j] = stackData
	}
	tempTxin.PeginWitness.Stack = pList

	return &tempTxin, nil
}

func convertListData(cfdData *cfd.TransactionData, cfdTxinList []cfd.ConfidentialTxIn, cfdTxoutList []cfd.ConfidentialTxOut) (data *types.TransactionData, txinList []types.ConfidentialTxIn, txoutList []types.ConfidentialTxOut, err error) {
	// FIXME need implement logic
	data = (*types.TransactionData)(unsafe.Pointer(cfdData))
	txinList = *(*[]types.ConfidentialTxIn)(unsafe.Pointer(&cfdTxinList))
	txoutList = *(*[]types.ConfidentialTxOut)(unsafe.Pointer(&cfdTxoutList))
	return data, txinList, txoutList, nil
}

//func (t *ConfidentialTx) AddMultisigSign() error {
// FIXME need implements
//}

func getDefaultBitcoinData() (lbtcAssetId, genesisBlockHash string) {
	cfdConfig := config.GetCurrentCfdConfig()
	if len(cfdConfig.BitcoinGenesisBlockHash) == 64 {
		genesisBlockHash = cfdConfig.BitcoinGenesisBlockHash
	}
	if len(cfdConfig.BitcoinAssetId) == 64 {
		lbtcAssetId = cfdConfig.BitcoinAssetId
	}
	return lbtcAssetId, genesisBlockHash
}

// appendDummyOutput ...
func appendDummyOutput(txHex string, network types.NetworkType, txinList *[]types.BlindInputData) (outputTx string, err error) {
	var blindTxInCount uint32
	var blindTxOutCount uint32
	outputTx = txHex
	// get all list
	_, _, txoutList, err := cfd.GetConfidentialTxDataAll(txHex, false, false, network.ToCfdValue())
	if err != nil {
		return "", err
	}

	if txinList != nil {
		for _, txin := range *txinList {
			if len(txin.ValueBlindFactor) == 64 && txin.ValueBlindFactor != EmptyBlinder {
				blindTxInCount += 1
			}
		}
	}

	var feeAsset string
	for _, txout := range txoutList {
		if len(txout.LockingScript) == 0 {
			// fee
			feeAsset = txout.Asset
		} else if len(txout.LockingScript) > 68 && strings.HasPrefix(txout.LockingScript, "6a") {
			// pegout
		} else if len(txout.CommitmentNonce) == 66 {
			// set confidential key
			blindTxOutCount += 1
		}
		// TODO(k-matsuzawa): Should we also count Outputs that directly specify Nonce?
	}

	if (blindTxInCount + blindTxOutCount) == 1 {
		// generate random confidential key
		nonce, _, _, err := cfd.CfdGoCreateKeyPair(true, network.ToBitcoinType().ToCfdValue())
		if err != nil {
			return "", err
		}
		outputTx, err = cfd.CfdGoAddConfidentialTxOut(txHex, feeAsset, 0, "", "", "6a", nonce)
		if err != nil {
			return "", err
		}
		if outputTx == txHex {
			return "", fmt.Errorf("CFD Error: fail logic")
		}
	} else if (blindTxInCount + blindTxOutCount) == 0 {
		// invalid blinding.
		return "", fmt.Errorf("CFD Error: blinding in/out not found")
		//} else if (blindTxInCount + blindTxOutCount) == 2 {
		//return "", fmt.Errorf("CFD Error: in/out=%d/%d", blindTxInCount, blindTxOutCount)
	}
	return outputTx, nil
}

// addConidentialTx ...
func addConidentialTx(txHandle uintptr, network types.NetworkType, locktime uint32, txinList *[]types.InputConfidentialTxIn, txoutList *[]types.InputConfidentialTxOut, pegoutAddressList *[]string) error {
	lbtcAsset, bitcoinGenesisBlockHash := getDefaultBitcoinData()
	btcTxApi := TransactionApiImpl{Network: network.ToBitcoinTypePointer()}

	var err error
	if txinList != nil {
		var bitcoinTxOut *types.TxOut
		for i := 0; i < len(*txinList); i++ {
			seq := (*txinList)[i].Sequence
			if seq == 0 {
				if locktime == 0 {
					seq = uint32(cfd.KCfdSequenceLockTimeFinal)
				} else {
					seq = uint32(cfd.KCfdSequenceLockTimeEnableMax)
				}
			}
			if (*txinList)[i].PeginInput != nil {
				btcTx := types.Transaction{Hex: (*txinList)[i].PeginInput.BitcoinTransaction}
				bitcoinTxOut, err = btcTxApi.GetTxOut(&btcTx, (*txinList)[i].OutPoint.Vout)
				if err != nil {
					return err
				}
				asset := (*txinList)[i].PeginInput.BitcoinAssetId
				if len(asset) == 0 {
					asset = lbtcAsset
				}
				genesisBlockHash := (*txinList)[i].PeginInput.BitcoinGenesisBlockHash
				if len(genesisBlockHash) == 0 {
					genesisBlockHash = bitcoinGenesisBlockHash
				}
				err = cfd.AddPeginInput(txHandle,
					(*txinList)[i].OutPoint.Txid, (*txinList)[i].OutPoint.Vout,
					bitcoinTxOut.Amount, asset, genesisBlockHash,
					(*txinList)[i].PeginInput.ClaimScript,
					(*txinList)[i].PeginInput.BitcoinTransaction,
					(*txinList)[i].PeginInput.TxOutProof,
				)
				if err == nil && (*txinList)[i].Sequence != 0 {
					err = cfd.UpdateTxInSequence(txHandle, (*txinList)[i].OutPoint.Txid, (*txinList)[i].OutPoint.Vout, seq)
				}
			} else {
				err = cfd.AddTransactionInput(txHandle, (*txinList)[i].OutPoint.Txid, (*txinList)[i].OutPoint.Vout, seq)
			}
			if err != nil {
				return err
			}
		}
	}

	if txoutList != nil {
		for i := 0; i < len(*txoutList); i++ {
			asset := (*txoutList)[i].Asset
			if len(asset) == 0 {
				asset = lbtcAsset
			}
			if (*txoutList)[i].PegoutInput != nil {
				var pubkey string
				if len((*txoutList)[i].PegoutInput.OnlineKey) == 64 {
					pubkey, err = cfd.CfdGoGetPubkeyFromPrivkey((*txoutList)[i].PegoutInput.OnlineKey, "", true)
				} else {
					pubkey, err = cfd.CfdGoGetPubkeyFromPrivkey("", (*txoutList)[i].PegoutInput.OnlineKey, true)
				}
				if err != nil {
					return err
				}
				genesisBlockHash := (*txoutList)[i].PegoutInput.BitcoinGenesisBlockHash
				if len(genesisBlockHash) == 0 {
					genesisBlockHash = bitcoinGenesisBlockHash
				}
				mainchainAddress, err := cfd.AddPegoutOutput(
					txHandle, asset, (*txoutList)[i].Amount,
					network.ToBitcoinType().ToCfdValue(), network.ToCfdValue(),
					genesisBlockHash, pubkey, (*txoutList)[i].PegoutInput.OnlineKey,
					(*txoutList)[i].PegoutInput.BitcoinOutputDescriptor,
					(*txoutList)[i].PegoutInput.Bip32Counter,
					(*txoutList)[i].PegoutInput.Whitelist)
				if pegoutAddressList != nil && err == nil {
					*pegoutAddressList = append(*pegoutAddressList, mainchainAddress)
				}
			} else if (*txoutList)[i].IsDestroy {
				err = cfd.AddTransactionOutput(txHandle, (*txoutList)[i].Amount, "", "6a", asset)
			} else if (*txoutList)[i].IsFee {
				err = cfd.AddTransactionOutput(txHandle, (*txoutList)[i].Amount, "", "", asset)
			} else {
				err = cfd.AddTransactionOutput(txHandle, (*txoutList)[i].Amount, (*txoutList)[i].Address, (*txoutList)[i].LockingScript, asset)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateDirectNonce(txHandle uintptr, txHex string, txoutList *[]types.InputConfidentialTxOut) (outputTxHex string, err error) {
	count := 0
	for i := 0; i < len(*txoutList); i++ {
		if len((*txoutList)[i].Nonce) == types.CommitmentHexDataSize {
			count += 1
		}
	}
	if count == 0 {
		return txHex, nil
	}

	outputTxHex = txHex
	for i := 0; i < len(*txoutList); i++ {
		if len((*txoutList)[i].Nonce) != types.CommitmentHexDataSize {
			// do nothing
		} else if ((*txoutList)[i].PegoutInput != nil) || (*txoutList)[i].IsFee || (len((*txoutList)[i].Address) > 0) {
			// do nothing
		} else if (*txoutList)[i].IsDestroy || (len((*txoutList)[i].LockingScript) > 0) {
			asset, satoshiAmount, valueCommitment, _, lockingScript, err := cfd.CfdGoGetConfidentialTxOutSimpleByHandle(txHandle, uint32(i))
			if err != nil {
				return "", err
			}
			outputTxHex, err = cfd.CfdGoUpdateConfidentialTxOut(outputTxHex, uint32(i), asset, satoshiAmount, valueCommitment, "", lockingScript, (*txoutList)[i].Nonce)
			if err != nil {
				return "", err
			}
		}
	}
	return outputTxHex, nil
}