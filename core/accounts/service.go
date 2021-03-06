package accounts

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"go-demo/infra/base"
	. "go-demo/services"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		IAccountService = new(accountService)
	})
}

type accountService struct {
}

func (a *accountService) CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error) {
	domain := accountDomain{}
	//验证输入参数
	if err := base.ValidateStruct(&dto); err != nil {
		return nil, err
	}
	//验证账户是否存在和幂等性
	acc := domain.GetAccountByUserIdAndType(dto.UserId, AccountType(dto.AccountType))
	if acc != nil {
		return acc, errors.New(
			fmt.Sprintf("用户的该类型账户已经存在：username=%s[%s],账户类型=%d",
				dto.Username, dto.UserId, dto.AccountType))
	}
	//执行账户创建的业务逻辑
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := AccountDTO{
		UserId:       dto.UserId,
		Username:     dto.Username,
		AccountType:  dto.AccountType,
		AccountName:  dto.AccountName,
		CurrencyCode: dto.CurrencyCode,
		Status:       1,
		Balance:      amount,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

func (a *accountService) Transfer(dto AccountTransferDTO) (TransferStatus, error) {
	//验证参数
	domain := accountDomain{}
	//验证输入参数
	if err := base.ValidateStruct(&dto); err != nil {
		return TransferStatusFailure, err
	}
	//执行转账逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return TransferStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == FlagTransferOut {
		if dto.ChangeType > 0 {
			return TransferStatusFailure,
				errors.New("如果changeFlag为支出，那么changeType必须小于0")
		}
	} else {
		if dto.ChangeType < 0 {
			return TransferStatusFailure,
				errors.New("如果changeFlag为收入,那么changeType必须大于0")
		}
	}

	status, err := domain.Transfer(dto)
	//转账成功，并且交易主体和交易目标不是同一个人，而且交易类型不是储值，则进行反向操作
	if status == TransferStatusSuccess && dto.TradeBody.AccountNo != dto.TradeTarget.AccountNo && dto.ChangeType != AccountStoreValue {
		backwardDto := dto
		backwardDto.TradeBody = dto.TradeTarget
		backwardDto.TradeTarget = dto.TradeBody
		backwardDto.ChangeType = -dto.ChangeType
		backwardDto.ChangeFlag = -dto.ChangeFlag
		status, err := domain.Transfer(backwardDto)
		return status, err
	}
	return status, err
}

func (a *accountService) StoreValue(dto AccountTransferDTO) (TransferStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = FlagTransferIn
	dto.ChangeType = AccountStoreValue
	return a.Transfer(dto)
}

func (a *accountService) GetEnvelopeAccountByUserId(userId string) *AccountDTO {
	domain := accountDomain{}
	account := domain.GetEnvelopeAccountByUserId(userId)
	return account
}

func (a *accountService) GetAccount(accountNo string) *AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}
