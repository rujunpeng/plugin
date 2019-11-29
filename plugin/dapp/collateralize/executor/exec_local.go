// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	//"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/types"
	pty "github.com/33cn/plugin/plugin/dapp/collateralize/types"
)

func (c *Collateralize) execLocal(tx *types.Transaction, receipt *types.ReceiptData) (*types.LocalDBSet, error) {
	set := &types.LocalDBSet{}
	for _, item := range receipt.Logs {
		if item.Ty == pty.TyLogCollateralizeCreate || item.Ty == pty.TyLogCollateralizeBorrow || item.Ty == pty.TyLogCollateralizeAppend ||
			item.Ty == pty.TyLogCollateralizeRepay || item.Ty == pty.TyLogCollateralizeFeed || item.Ty == pty.TyLogCollateralizeRetrieve {
			var collateralizeLog pty.ReceiptCollateralize
			err := types.Decode(item.Log, &collateralizeLog)
			if err != nil {
				return nil, err
			}

			switch item.Ty {
			case pty.TyLogCollateralizeCreate:
				set.KV = append(set.KV, c.deleteCollateralizeStatus(collateralizeLog.PreStatus, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.deleteCollateralizeAddr(collateralizeLog.CreateAddr, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.addCollateralizeStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId, collateralizeLog.Index)...)
				set.KV = append(set.KV, c.addCollateralizeAddr(collateralizeLog.CreateAddr, collateralizeLog.CollateralizeId, collateralizeLog.Status, collateralizeLog.Index)...)
				break
			case pty.TyLogCollateralizeBorrow:
				set.KV = append(set.KV, c.addCollateralizeRecordStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId,
					collateralizeLog.RecordId, collateralizeLog.Index)...)
				set.KV = append(set.KV, c.addCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.CollateralizeId,
					collateralizeLog.RecordId, collateralizeLog.Index)...)
				break
			case pty.TyLogCollateralizeAppend:
				if collateralizeLog.Status == pty.CollateralizeUserStatusWarning {
					set.KV = append(set.KV, c.deleteCollateralizeRecordStatus(collateralizeLog.PreStatus, collateralizeLog.PreIndex)...)
					set.KV = append(set.KV, c.addCollateralizeRecordStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId,
						collateralizeLog.RecordId, collateralizeLog.Index)...)
					//set.KV = append(set.KV, c.deleteCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.PreIndex)...)
					//set.KV = append(set.KV, c.addCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.CollateralizeId,
					//	collateralizeLog.RecordId, collateralizeLog.Index)...)
				}
				break
			case pty.TyLogCollateralizeRepay:
				set.KV = append(set.KV, c.deleteCollateralizeRecordStatus(collateralizeLog.PreStatus, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.addCollateralizeRecordStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId,
					collateralizeLog.RecordId, collateralizeLog.Index)...)
				//set.KV = append(set.KV, c.deleteCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.PreIndex)...)
				break
			case pty.TyLogCollateralizeFeed:
				set.KV = append(set.KV, c.deleteCollateralizeRecordStatus(collateralizeLog.PreStatus, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.addCollateralizeRecordStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId,
					collateralizeLog.RecordId, collateralizeLog.Index)...)
				//set.KV = append(set.KV, c.deleteCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.PreIndex)...)
				//// 如果没有被清算，需要把地址索引更新
				//if collateralizeLog.Status == pty.CollateralizeUserStatusWarning || collateralizeLog.Status == pty.CollateralizeUserStatusExpire {
				//	set.KV = append(set.KV, c.addCollateralizeRecordAddr(collateralizeLog.AccountAddr, collateralizeLog.CollateralizeId,
				//		collateralizeLog.RecordId, collateralizeLog.Index)...)
				//}
				break
			case pty.TyLogCollateralizeRetrieve:
				set.KV = append(set.KV, c.deleteCollateralizeStatus(collateralizeLog.PreStatus, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.deleteCollateralizeAddr(collateralizeLog.CreateAddr, collateralizeLog.PreIndex)...)
				set.KV = append(set.KV, c.addCollateralizeStatus(collateralizeLog.Status, collateralizeLog.CollateralizeId, collateralizeLog.Index)...)
				set.KV = append(set.KV, c.addCollateralizeAddr(collateralizeLog.CreateAddr, collateralizeLog.CollateralizeId, collateralizeLog.Status, collateralizeLog.Index)...)
				break
			}
		}
	}
	return set, nil
}

// ExecLocal_Create Action
func (c *Collateralize) ExecLocal_Create(payload *pty.CollateralizeCreate, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Borrow Action
func (c *Collateralize) ExecLocal_Borrow(payload *pty.CollateralizeBorrow, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Repay Action
func (c *Collateralize) ExecLocal_Repay(payload *pty.CollateralizeRepay, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Repay Action
func (c *Collateralize) ExecLocal_Append(payload *pty.CollateralizeAppend, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Feed Action
func (c *Collateralize) ExecLocal_Feed(payload *pty.CollateralizeFeed, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Retrieve Action
func (c *Collateralize) ExecLocal_Retrieve(payload *pty.CollateralizeRetrieve, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}

// ExecLocal_Manage Action
func (c *Collateralize) ExecLocal_Manage(payload *pty.CollateralizeManage, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return c.execLocal(tx, receiptData)
}
